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

// AnalysisService encapsulates the API logic for analysis-related operations
type AnalysisService struct {
	client *Client
}

// NewAnalysisService creates a new AnalysisService
func NewAnalysisService(client *Client) *AnalysisService {
	return &AnalysisService{
		client: client,
	}
}

// List retrieves all analyses for a specific site
func (as *AnalysisService) List(ctx context.Context, siteName string, siteType string) ([]model.Analysis, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/analyses/external/" + url.PathEscape(siteName)
	case "operator":
		endpoint = "/api/analyses/operator/" + url.PathEscape(siteName)
	case "master":
		endpoint = "/api/analyses/master"
	case "action":
		endpoint = "/api/analyses/action/" + url.PathEscape(siteName)
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := as.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return as.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(as.client.BaseURL + ":" + strconv.Itoa(as.client.PortNumber) + endpoint)
	}, "bigfix_analysis_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch analyses for site %s (%s): %w", siteName, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for analyses response
	var result model.AnalysisListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Set site information for each analysis
	for i := range result.Analyses {
		result.Analyses[i].SiteName = siteName
		result.Analyses[i].SiteType = siteType
	}

	plugin.Logger(ctx).Debug("API response analyses:", result.Analyses)

	return result.Analyses, nil
}

// Get retrieves a specific analysis detail
func (as *AnalysisService) Get(ctx context.Context, siteName string, siteType string, analysisID int) (*model.Analysis, error) {
	var endpoint string

	switch siteType {
	case "external":
		endpoint = "/api/analysis/external/" + url.PathEscape(siteName) + "/" + strconv.Itoa(analysisID)
	case "operator":
		endpoint = "/api/analysis/operator/" + url.PathEscape(siteName) + "/" + strconv.Itoa(analysisID)
	case "master":
		endpoint = "/api/analysis/master/" + strconv.Itoa(analysisID)
	case "action":
		endpoint = "/api/analysis/action/" + url.PathEscape(siteName) + "/" + strconv.Itoa(analysisID)
	default:
		return nil, fmt.Errorf("invalid site type: %s. Must be one of: external, operator, master, action", siteType)
	}

	// Perform the request with retry logic and limiter tag
	resp, err := as.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return as.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(as.client.BaseURL + ":" + strconv.Itoa(as.client.PortNumber) + endpoint)
	}, "bigfix_analysis_get")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch analysis %d for site %s (%s): %w", analysisID, siteName, siteType, err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for analysis detail response
	var result model.AnalysisDetailResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert to Analysis model
	resourceURL := as.client.BaseURL + ":" + strconv.Itoa(as.client.PortNumber) + endpoint
	analysis := result.Analysis.ToAnalysis(analysisID, resourceURL, siteName, siteType)

	plugin.Logger(ctx).Debug("API response analysis:", analysis)

	return analysis, nil
}
