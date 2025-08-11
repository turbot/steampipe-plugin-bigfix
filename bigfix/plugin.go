package bigfix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/rate_limiter"
)

const pluginName = "steampipe-plugin-bigfix"

// Plugin returns the BigFix plugin definition.
func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		RateLimiters: []*rate_limiter.Definition{
			{
				Name:       "bigfix_api",
				FillRate:   300,
				BucketSize: 300,
				Scope:      []string{"connection"},
			},
		},
		TableMap: map[string]*plugin.Table{
			"bigfix_computer": tableBigFixComputer(ctx),
			"bigfix_site":     tableBigFixSite(ctx),
			"bigfix_analysis": tableBigFixAnalysis(ctx),
			"bigfix_task":     tableBigFixTask(ctx),
			"bigfix_action":   tableBigFixAction(ctx),
			"bigfix_fixlet":   tableBigFixFixlet(ctx),
			"bigfix_property": tableBigFixProperty(ctx),
			"bigfix_role":     tableBigFixRole(ctx),
		},
	}
}
