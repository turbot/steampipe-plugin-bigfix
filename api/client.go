package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"resty.dev/v3"
)

const BaseURL = "https://%s"

// Client is a reusable HTTP client for the BigFix API using Resty.
type Client struct {
	Resty      *resty.Client
	BaseURL    string
	PortNumber int
	MaxRetries int
	MinDelay   time.Duration
	rand       *rand.Rand

	// Service clients
	Computer *ComputerService
	Site     *SiteService
	Analysis *AnalysisService
	Task     *TaskService
	Action   *ActionService
	Fixlet   *FixletService
	Property *PropertyService
	Role     *RoleService
}

// NewClient returns a new Client with a Resty client and the BigFix API base URL.
func NewClient(serverName, userName, password string, port int, insecureSkipVerify bool, timeout time.Duration) *Client {
	client := resty.New()

	// Configure timeouts - increased to handle large datasets
	client.SetTimeout(timeout)

	client.SetBasicAuth(userName, password)

	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: insecureSkipVerify})

	// Basic retry configuration if available
	if err := client.SetRetryCount(3); err != nil {
		log.Printf("[WARNING] Could not set retry count: %v", err)
	}

	baseUrl := fmt.Sprintf(BaseURL, serverName)

	bigfixClient := &Client{
		Resty:      client,
		BaseURL:    baseUrl,
		PortNumber: port,
		MaxRetries: 3,                                               // Default to 3 retries
		MinDelay:   100 * time.Millisecond,                          // Default minimum delay
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())), // Modern random source
	}

	// Initialize service clients
	bigfixClient.Computer = NewComputerService(bigfixClient)
	bigfixClient.Site = NewSiteService(bigfixClient)
	bigfixClient.Analysis = NewAnalysisService(bigfixClient)
	bigfixClient.Task = NewTaskService(bigfixClient)
	bigfixClient.Action = NewActionService(bigfixClient)
	bigfixClient.Fixlet = NewFixletService(bigfixClient)
	bigfixClient.Property = NewPropertyService(bigfixClient)
	bigfixClient.Role = NewRoleService(bigfixClient)

	return bigfixClient
}

// WithMaxRetries sets the maximum number of retries for the client
func (c *Client) WithMaxRetries(maxRetries int) *Client {
	c.MaxRetries = maxRetries
	return c
}

// WithMinDelay sets the minimum delay for backoff calculations
func (c *Client) WithMinDelay(minDelay time.Duration) *Client {
	c.MinDelay = minDelay
	return c
}

// BackoffDelay returns the duration to wait before the next attempt should be
// made. Returns an error if unable get a duration.
func (c *Client) BackoffDelay(attempt int, err error) (time.Duration, error) {
	minDelay := c.MinDelay

	// The calculated jitter will be between [0.8, 1.2)
	var jitter = float64(c.rand.Intn(120-80)+80) / 100

	retryTime := time.Duration(int(float64(int(minDelay.Nanoseconds())*int(math.Pow(3, float64(attempt)))) * jitter))

	// Cap retry time at 5 minutes to avoid too long a wait
	if retryTime > time.Duration(5*time.Minute) {
		retryTime = time.Duration(5 * time.Minute)
	}

	// Low level method to log retries since we don't have context etc here.
	// Logging is helpful for visibility into retries and choke points in using
	// the API.
	log.Printf("[INFO] BackoffDelay: attempt=%d, retryTime=%s, err=%v", attempt, retryTime.String(), err)

	return retryTime, nil
}

// executeWithRetry performs an HTTP request with manual retry logic
func (c *Client) executeWithRetry(request func() (*resty.Response, error), maxRetries int) (*resty.Response, error) {
	var lastErr error
	var resp *resty.Response

	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("[REQUEST] Attempt %d/%d", attempt, maxRetries)

		resp, lastErr = request()

		if lastErr == nil && resp != nil {
			statusCode := resp.StatusCode()
			log.Printf("[RESPONSE] Status: %d", statusCode)

			// Success - no need to retry
			if statusCode >= 200 && statusCode < 300 {
				return resp, nil
			}

			// Check if we should retry based on status code
			shouldRetry := false
			switch statusCode {
			case 429: // Rate limited
				log.Printf("[RETRY] Rate limited (429), retrying...")
				shouldRetry = true
			case 408: // Request timeout
				log.Printf("[RETRY] Request timeout (408), retrying...")
				shouldRetry = true
			case 500, 502, 503, 504: // Server errors
				log.Printf("[RETRY] Server error (%d), retrying...", statusCode)
				shouldRetry = true
			default:
				if statusCode >= 400 && statusCode < 500 {
					log.Printf("[NO RETRY] Client error (%d), not retrying", statusCode)
					return resp, fmt.Errorf("HTTP client error: %d %s", statusCode, resp.Status())
				}
			}

			if !shouldRetry {
				return resp, fmt.Errorf("HTTP error: %d %s", statusCode, resp.Status())
			}
		} else if lastErr != nil {
			log.Printf("[RETRY] Network error: %v", lastErr)
		}

		// Don't sleep after the last attempt
		if attempt < maxRetries {
			backoff, err := c.BackoffDelay(attempt, lastErr)
			if err != nil {
				log.Printf("[ERROR] Failed to calculate backoff delay: %v", err)
				backoff = 1 * time.Second // fallback
			}

			time.Sleep(backoff)
		}
	}

	if lastErr != nil {
		return resp, fmt.Errorf("request failed after %d attempts: %w", maxRetries, lastErr)
	}

	return resp, fmt.Errorf("request failed after %d attempts with status %d", maxRetries, resp.StatusCode())
}

// executeWithRetryDefaultWithLimiter performs an HTTP request using the client's default retry settings with limiter tag
func (c *Client) executeWithRetryDefaultWithLimiter(request func() (*resty.Response, error), limiterTag string) (*resty.Response, error) {
	// Add limiter tag to the request
	limiterRequest := func() (*resty.Response, error) {
		return request()
	}
	return c.executeWithRetry(limiterRequest, c.MaxRetries)
}

// Backward compatibility methods - these delegate to the new service-based API

// ListComputers provides backward compatibility for existing code
// Deprecated: Use client.Computer.List() instead
func (c *Client) ListComputers() ([]model.Computer, error) {
	return c.Computer.List()
}

// GetComputer provides backward compatibility for existing code
// Deprecated: Use client.Computer.Get() instead
func (c *Client) GetComputer(ctx context.Context, id int) (*model.Computer, error) {
	return c.Computer.Get(ctx, id)
}

// ListSites provides backward compatibility for existing code
// Deprecated: Use client.Site.List() instead
func (c *Client) ListSites() ([]model.Site, error) {
	return c.Site.List()
}

// GetSite provides backward compatibility for existing code
// Deprecated: Use client.Site.Get() instead
func (c *Client) GetSite(ctx context.Context, name string, siteType string) (*model.Site, error) {
	return c.Site.Get(ctx, name, siteType)
}

// GetSitePermissions provides backward compatibility for existing code
// Deprecated: Use client.Site.GetPermissions() instead
func (c *Client) GetSitePermissions(ctx context.Context, name string, siteType string) ([]model.SitePermission, error) {
	return c.Site.GetPermissions(ctx, name, siteType)
}

// GetSiteFiles provides backward compatibility for existing code
// Deprecated: Use client.Site.GetFiles() instead
func (c *Client) GetSiteFiles(ctx context.Context, name string, siteType string) ([]model.SiteFile, error) {
	return c.Site.GetFiles(ctx, name, siteType)
}

// GetSiteAnalyses provides backward compatibility for existing code
// Deprecated: Use client.Analysis.List() instead
func (c *Client) GetSiteAnalyses(ctx context.Context, name string, siteType string) ([]model.Analysis, error) {
	return c.Analysis.List(ctx, name, siteType)
}

// GetSiteAnalysis provides backward compatibility for existing code
// Deprecated: Use client.Analysis.Get() instead
func (c *Client) GetSiteAnalysis(ctx context.Context, name string, siteType string, analysisID int) (*model.Analysis, error) {
	return c.Analysis.Get(ctx, name, siteType, analysisID)
}

// GetSiteTasks provides backward compatibility for existing code
// Deprecated: Use client.Task.List() instead
func (c *Client) GetSiteTasks(ctx context.Context, name string, siteType string) ([]model.Task, error) {
	return c.Task.List(ctx, name, siteType)
}

// GetSiteTask provides backward compatibility for existing code
// Deprecated: Use client.Task.Get() instead
func (c *Client) GetSiteTask(ctx context.Context, name string, siteType string, taskID int) (*model.Task, error) {
	return c.Task.Get(ctx, name, siteType, taskID)
}
