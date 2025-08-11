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

// PropertyService encapsulates the API logic for property-related operations
type PropertyService struct {
	client *Client
}

// NewPropertyService creates a new PropertyService
func NewPropertyService(client *Client) *PropertyService {
	return &PropertyService{
		client: client,
	}
}

// List retrieves all properties
func (ps *PropertyService) List(ctx context.Context) ([]model.BigFixProperty, error) {
	endpoint := "/api/properties"

	// Perform the request with retry logic and limiter tag
	resp, err := ps.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ps.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ps.client.BaseURL + ":" + strconv.Itoa(ps.client.PortNumber) + endpoint)
	}, "bigfix_property_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch properties: %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for properties response
	var result model.BigFixPropertyListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	plugin.Logger(ctx).Debug("API response properties:", result.Properties)

	return result.Properties, nil
}

// Get retrieves a specific property detail
func (ps *PropertyService) Get(ctx context.Context, propertyID int) (*model.BigFixProperty, error) {
	endpoint := "/api/property/" + strconv.Itoa(propertyID)

	// Perform the request with retry logic and limiter tag
	resp, err := ps.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ps.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ps.client.BaseURL + ":" + strconv.Itoa(ps.client.PortNumber) + endpoint)
	}, "bigfix_property_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch property %d: %w", propertyID, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for property detail response
	var result model.BigFixPropertyDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert to BigFixProperty model
	resourceURL := ps.client.BaseURL + ":" + strconv.Itoa(ps.client.PortNumber) + endpoint
	property := result.Property.ToBigFixProperty(propertyID, resourceURL, result.Property.Name, 0) // IsReserved not available in detail response

	plugin.Logger(ctx).Debug("API response property:", property)

	return property, nil
}
