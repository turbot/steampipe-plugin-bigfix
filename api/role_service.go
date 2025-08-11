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

// RoleService encapsulates the API logic for role-related operations
type RoleService struct {
	client *Client
}

// NewRoleService creates a new RoleService
func NewRoleService(client *Client) *RoleService {
	return &RoleService{
		client: client,
	}
}

// List retrieves all roles
func (rs *RoleService) List(ctx context.Context) ([]model.Role, error) {
	endpoint := "/api/roles"

	// Perform the request with retry logic and limiter tag
	resp, err := rs.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return rs.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(rs.client.BaseURL + ":" + strconv.Itoa(rs.client.PortNumber) + endpoint)
	}, "bigfix_role_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch roles: %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for roles response
	var result model.RoleListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	plugin.Logger(ctx).Debug("API response roles:", result.Roles)

	return result.Roles, nil
}

// Get retrieves a specific role detail
func (rs *RoleService) Get(ctx context.Context, roleID int) (*model.Role, error) {
	endpoint := "/api/role/" + strconv.Itoa(roleID)

	// Perform the request with retry logic and limiter tag
	resp, err := rs.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return rs.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(rs.client.BaseURL + ":" + strconv.Itoa(rs.client.PortNumber) + endpoint)
	}, "bigfix_role_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch role %d: %w", roleID, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for role detail response
	var result model.RoleDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Set the resource URL for the role
	result.Role.Resource = rs.client.BaseURL + ":" + strconv.Itoa(rs.client.PortNumber) + endpoint

	plugin.Logger(ctx).Debug("API response role:", result.Role)

	return &result.Role, nil
}
