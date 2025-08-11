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

// ComputerService handles computer-related API operations
type ComputerService struct {
	client *Client
}

// NewComputerService creates a new computer service
func NewComputerService(client *Client) *ComputerService {
	return &ComputerService{client: client}
}

// List retrieves a list of computers from the BigFix API
func (cs *ComputerService) List() ([]model.Computer, error) {
	// Build the endpoint URL with filtered fields
	params := url.Values{}
	params.Add("fields", "ID,Name,OS,LastReportTime,CPU,IPAddress")
	endpoint := "/api/computers?" + params.Encode()

	// Perform the request with retry logic and limiter tag
	resp, err := cs.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return cs.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(cs.client.BaseURL + ":" + strconv.Itoa(cs.client.PortNumber) + endpoint)
	}, "bigfix_computer_list")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch computers: %w", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the XML response
	var result model.BESAPI
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert XML computers to Computer models
	computers := make([]model.Computer, 0, len(result.Computers))
	for _, computerXML := range result.Computers {
		computer, err := computerXML.ToComputer()
		if err != nil {
			return nil, fmt.Errorf("failed to convert computer XML to model: %w", err)
		}
		computers = append(computers, *computer)
	}

	return computers, nil
}

// Get retrieves a single computer by ID.
//
// The fields parameter is optional. When provided, it will be URL-encoded and
// passed as the value of the `fields` query parameter to filter the response.
// Examples (pass them unencoded; they will be encoded automatically):
// - Property<Name=Computer Name,OS,Last Report Time>
// - Property<Analysis&Name=Analysis1&Computer Name,&OS,Analysis2&Last Report Time>
func (cs *ComputerService) Get(ctx context.Context, id int) (*model.Computer, error) {
	// Build endpoint with optional fields filter
	endpoint := "/api/computer/" + fmt.Sprintf("%d", id) + "?fields" // Get all properties

	// Perform the request with retry logic and limiter tag
	resp, err := cs.client.executeWithRetryDefaultWithLimiter(func() (*resty.Response, error) {
		return cs.client.Resty.R().
			SetHeader("Accept", "application/xml").
			Get(cs.client.BaseURL + ":" + strconv.Itoa(cs.client.PortNumber) + endpoint)
	}, "bigfix_computer_get")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch computer %d: %w", id, err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML for single computer response with properties
	var result model.ComputerXMLResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	// Convert XML response to Computer model
	computer, err := result.Computer.ToComputer()
	if err != nil {
		return nil, fmt.Errorf("failed to convert XML response to computer model: %w", err)
	}

	// Ensure the ID is set from the URL parameter
	computer.ID = id

	plugin.Logger(ctx).Debug("API response computer:", computer)

	return computer, nil
}
