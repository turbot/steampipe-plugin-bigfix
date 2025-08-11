package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"resty.dev/v3"
)

// ActionService encapsulates the API logic for action-related operations
type ActionService struct {
	client *Client
}

// NewActionService creates a new ActionService
func NewActionService(client *Client) *ActionService {
	return &ActionService{
		client: client,
	}
}

// List retrieves all actions
func (as *ActionService) List(ctx context.Context) ([]model.Action, error) {
	endpoint := "/api/actions"

	// Perform the request with retry logic and limiter tag
	resp, err := as.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return as.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(as.client.BaseURL + ":" + strconv.Itoa(as.client.PortNumber) + endpoint)
	}, "bigfix_action_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch actions: %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for actions response
	var result model.ActionListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Actions don't have site information in the API response

	plugin.Logger(ctx).Debug("API response actions:", result.Actions)

	return result.Actions, nil
}

// Get retrieves a specific action detail
func (as *ActionService) Get(ctx context.Context, actionID int) (*model.Action, error) {
	endpoint := "/api/action/" + strconv.Itoa(actionID)

	// Perform the request with retry logic and limiter tag
	resp, err := as.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return as.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(as.client.BaseURL + ":" + strconv.Itoa(as.client.PortNumber) + endpoint)
	}, "bigfix_action_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch action %d: %w", actionID, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for action detail response
	var result model.ActionDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert to Action model
	resourceURL := as.client.BaseURL + ":" + strconv.Itoa(as.client.PortNumber) + endpoint
	action := result.Action.ToAction(actionID, resourceURL)

	plugin.Logger(ctx).Debug("API response action:", action)

	return action, nil
}
