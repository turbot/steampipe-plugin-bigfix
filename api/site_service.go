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

// SiteService handles site-related API operations
type SiteService struct {
	client *Client
}

// NewSiteService creates a new site service
func NewSiteService(client *Client) *SiteService {
	return &SiteService{client: client}
}

// List retrieves a list of sites from the BigFix API
func (ss *SiteService) List() ([]model.Site, error) {
	endpoint := "/api/sites"

	// Perform the request with retry logic and limiter tag
	resp, err := ss.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ss.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ss.client.BaseURL + ":" + strconv.Itoa(ss.client.PortNumber) + endpoint)
	}, "bigfix_site_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch sites: %w", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the XML response
	var result model.SiteListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert XML sites to Site models
	var sites []model.Site

	// Add external sites
	for _, externalSite := range result.ExternalSites {
		site := externalSite.ToSite()
		sites = append(sites, *site)
	}

	// Add operator sites
	for _, operatorSite := range result.OperatorSites {
		site := operatorSite.ToSite()
		sites = append(sites, *site)
	}

	// Add action sites
	for _, actionSite := range result.ActionSites {
		site := actionSite.ToSite()
		sites = append(sites, *site)
	}

	return sites, nil
}

// Get retrieves a single site by name and type
// The siteType should be one of: "external", "operator", "master" (for action site)
func (ss *SiteService) Get(ctx context.Context, name string, siteType string) (*model.Site, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/site/external/" + url.PathEscape(name)
	case "operator":
		endpoint = "/api/site/operator/" + url.PathEscape(name)
	case "master", "action":
		endpoint = "/api/site/master"
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := ss.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ss.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ss.client.BaseURL + ":" + strconv.Itoa(ss.client.PortNumber) + endpoint)
	}, "bigfix_site_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch site %s (%s): %w", name, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for site detail response
	var result model.SiteDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert the appropriate site detail to Site model
	var site *model.Site
	switch siteType {
	case "external":
		if result.ExternalSite != nil {
			site = result.ExternalSite.ToSite()
		}
	case "operator":
		if result.OperatorSite != nil {
			site = result.OperatorSite.ToSite()
		}
	case "master", "action":
		if result.ActionSite != nil {
			site = result.ActionSite.ToSite()
		}
	}

	if site == nil {
		return nil, fmt.Errorf("site %s (%s) not found", name, siteType)
	}

	// Set the resource URL for the site
	site.Resource = ss.client.BaseURL + ":" + strconv.Itoa(ss.client.PortNumber) + endpoint

	plugin.Logger(ctx).Debug("API response site:", site)

	return site, nil
}

// GetPermissions retrieves permissions for a specific site
func (ss *SiteService) GetPermissions(ctx context.Context, name string, siteType string) ([]model.SitePermission, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/site/external/" + url.PathEscape(name) + "/permissions"
	case "operator":
		endpoint = "/api/site/operator/" + url.PathEscape(name) + "/permissions"
	case "master":
		endpoint = "/api/site/master/permissions"
	case "action":
		// Action sites might not have permissions endpoint, but let's try the action pattern
		endpoint = "/api/site/action/" + url.PathEscape(name) + "/permissions"
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := ss.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ss.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ss.client.BaseURL + ":" + strconv.Itoa(ss.client.PortNumber) + endpoint)
	}, "bigfix_site_permissions")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch permissions for site %s (%s): %w", name, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for permissions response
	var result model.SitePermissionsResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	plugin.Logger(ctx).Debug("API response permissions:", result.Permissions)

	return result.Permissions, nil
}

// GetFiles retrieves files for a specific site
func (ss *SiteService) GetFiles(ctx context.Context, name string, siteType string) ([]model.SiteFile, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/site/external/" + url.PathEscape(name) + "/files"
	case "operator":
		endpoint = "/api/site/operator/" + url.PathEscape(name) + "/files"
	case "master":
		endpoint = "/api/site/master/files"
	case "action":
		endpoint = "/api/site/action/" + url.PathEscape(name) + "/files"
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := ss.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return ss.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(ss.client.BaseURL + ":" + strconv.Itoa(ss.client.PortNumber) + endpoint)
	}, "bigfix_site_files")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch files for site %s (%s): %w", name, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for files response
	var result model.SiteFilesResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	plugin.Logger(ctx).Debug("API response files:", result.Files)

	return result.Files, nil
}
