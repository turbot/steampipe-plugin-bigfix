package bigfix

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-bigfix/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func NewService(ctx context.Context, d *plugin.QueryData) (*api.Client, error) {
	config := GetConfig(d.Connection)

	if config.ServerName == nil {
		return nil, fmt.Errorf("server_name is required")
	}

	if config.UserName == nil {
		return nil, fmt.Errorf("user_name is required")
	}

	if config.Password == nil {
		return nil, fmt.Errorf("password is required")
	}

	if config.Port == nil {
		return nil, fmt.Errorf("port is required")
	}

	// Default insecure_skip_verify to false if not specified
	insecureSkipVerify := false
	if config.InsecureSkipVerify != nil {
		insecureSkipVerify = *config.InsecureSkipVerify
	}

	// Default request timeout to 120 seconds if not specified
	requestTimeout := 120 * time.Second
	if config.RequestTimeout != nil {
		requestTimeout = time.Duration(*config.RequestTimeout) * time.Second
	}

	client := api.NewClient(*config.ServerName, *config.UserName, *config.Password, *config.Port, insecureSkipVerify, requestTimeout)

	// Apply configuration overrides
	if config.MaxRetries != nil {
		client = client.WithMaxRetries(*config.MaxRetries)
	}

	if config.MinRetryDelay != nil {
		client = client.WithMinDelay(time.Duration(*config.MinRetryDelay) * time.Millisecond)
	}

	return client, nil
}
