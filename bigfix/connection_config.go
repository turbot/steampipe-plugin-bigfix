package bigfix

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type BigFixConfig struct {
	MaxRetries          *int     `hcl:"max_retries,optional"`
	MinRetryDelay       *int64   `hcl:"min_retry_delay,optional"`
	ServerName          *string  `hcl:"server_name,optional"`
	Port                *int     `hcl:"port,optional"`
	UserName            *string  `hcl:"user_name,optional"`
	Password            *string  `hcl:"password,optional"`
	IgnoreErrorMessages []string `hcl:"ignore_error_messages,optional"`
	InsecureSkipVerify  *bool    `hcl:"insecure_skip_verify,optional"`
	RequestTimeout      *int64   `hcl:"request_timeout,optional"`
}

func ConfigInstance() interface{} {
	return &BigFixConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) BigFixConfig {
	if connection == nil || connection.Config == nil {
		return BigFixConfig{}
	}
	config, _ := connection.Config.(BigFixConfig)

	return config
}
