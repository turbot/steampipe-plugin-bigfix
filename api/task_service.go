package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"resty.dev/v3"
)

// TaskService encapsulates the API logic for task-related operations
type TaskService struct {
	client *Client
}

// NewTaskService creates a new TaskService
func NewTaskService(client *Client) *TaskService {
	return &TaskService{
		client: client,
	}
}

// List retrieves all tasks for a specific site
func (ts *TaskService) List(ctx context.Context, siteName string, siteType string) ([]model.Task, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/tasks/external/" + url.PathEscape(siteName)
	case "operator":
		endpoint = "/api/tasks/operator/" + url.PathEscape(siteName)
	case "master":
		endpoint = "/api/tasks/master"
	case "action":
		endpoint = "/api/tasks/action/" + url.PathEscape(siteName)
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := ts.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ts.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ts.client.BaseURL + ":" + strconv.Itoa(ts.client.PortNumber) + endpoint)
	}, "bigfix_task_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks for site %s (%s): %w", siteName, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for tasks response
	var result model.TaskListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Set site information for each task
	for i := range result.Tasks {
		result.Tasks[i].SiteName = siteName
		result.Tasks[i].SiteType = siteType
	}

	plugin.Logger(ctx).Debug("API response tasks:", result.Tasks)

	return result.Tasks, nil
}

// Get retrieves a specific task detail
func (ts *TaskService) Get(ctx context.Context, siteName string, siteType string, taskID int) (*model.Task, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/task/external/" + url.PathEscape(siteName) + "/" + strconv.Itoa(taskID)
	case "operator":
		endpoint = "/api/task/operator/" + url.PathEscape(siteName) + "/" + strconv.Itoa(taskID)
	case "master":
		endpoint = "/api/task/master/" + strconv.Itoa(taskID)
	case "action":
		endpoint = "/api/task/action/" + url.PathEscape(siteName) + "/" + strconv.Itoa(taskID)
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := ts.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ts.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ts.client.BaseURL + ":" + strconv.Itoa(ts.client.PortNumber) + endpoint)
	}, "bigfix_task_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch task %d for site %s (%s): %w", taskID, siteName, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for task detail response
	var result model.TaskDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert to Task model
	resourceURL := ts.client.BaseURL + ":" + strconv.Itoa(ts.client.PortNumber) + endpoint
	task := result.Task.ToTask(taskID, resourceURL, siteName, siteType)

	plugin.Logger(ctx).Debug("API response task:", task)

	return task, nil
}
