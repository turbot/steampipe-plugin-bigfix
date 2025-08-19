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

// FixletService encapsulates the API logic for fixlet-related operations
type FixletService struct {
	client *Client
}

// NewFixletService creates a new FixletService
func NewFixletService(client *Client) *FixletService {
	return &FixletService{
		client: client,
	}
}

// List retrieves all fixlets for a specific site
func (fs *FixletService) List(ctx context.Context, siteName string, siteType string) ([]model.Fixlet, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/fixlets/external/" + url.PathEscape(siteName)
	case "operator":
		endpoint = "/api/fixlets/operator/" + url.PathEscape(siteName)
	case "master":
		endpoint = "/api/fixlets/master"
	case "action":
		endpoint = "/api/fixlets/action/" + url.PathEscape(siteName)
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := fs.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return fs.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(fs.client.BaseURL + ":" + strconv.Itoa(fs.client.PortNumber) + endpoint)
	}, "bigfix_fixlet_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch fixlets for site %s (%s): %w", siteName, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for fixlets response
	var result model.FixletListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Set site information for each fixlet
	for i := range result.Fixlets {
		result.Fixlets[i].SiteName = siteName
		result.Fixlets[i].SiteType = siteType
	}

	plugin.Logger(ctx).Debug("API response fixlets:", result.Fixlets)

	return result.Fixlets, nil
}

// Get retrieves a specific fixlet detail
func (fs *FixletService) Get(ctx context.Context, siteName string, siteType string, fixletID int) (*model.Fixlet, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/fixlet/external/" + url.PathEscape(siteName) + "/" + strconv.Itoa(fixletID)
	case "operator":
		endpoint = "/api/fixlet/operator/" + url.PathEscape(siteName) + "/" + strconv.Itoa(fixletID)
	case "master":
		endpoint = "/api/fixlet/master/" + strconv.Itoa(fixletID)
	case "action":
		endpoint = "/api/fixlet/action/" + url.PathEscape(siteName) + "/" + strconv.Itoa(fixletID)
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := fs.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return fs.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(fs.client.BaseURL + ":" + strconv.Itoa(fs.client.PortNumber) + endpoint)
	}, "bigfix_fixlet_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch fixlet %d for site %s (%s): %w", fixletID, siteName, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for fixlet detail response
	var result model.FixletDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert to Fixlet model
	resourceURL := fs.client.BaseURL + ":" + strconv.Itoa(fs.client.PortNumber) + endpoint
	fixlet := result.Fixlet.ToFixlet(fixletID, resourceURL, siteName, siteType)

	plugin.Logger(ctx).Debug("API response fixlet:", fixlet)

	return fixlet, nil
}
