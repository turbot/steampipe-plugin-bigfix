package bigfix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
		TableMap: map[string]*plugin.Table{
			"bigfix_action":   tableBigFixAction(ctx),
			"bigfix_analysis": tableBigFixAnalysis(ctx),
			"bigfix_computer": tableBigFixComputer(ctx),
			"bigfix_fixlet":   tableBigFixFixlet(ctx),
			"bigfix_property": tableBigFixProperty(ctx),
			"bigfix_role":     tableBigFixRole(ctx),
			"bigfix_site":     tableBigFixSite(ctx),
			"bigfix_task":     tableBigFixTask(ctx),
		},
	}
}
