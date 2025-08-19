package bigfix

import (
	"context"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableBigFixSite(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_site",
		Description: "BigFix Site contains content repositories and organizational structures with configurations, permissions, and subscription settings for content management.",
		List: &plugin.ListConfig{
			Hydrate: listBigFixSites,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Required},
				{Name: "type", Require: plugin.Required},
			},
			Hydrate: getBigFixSite,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBigFixSitePermissions,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
			{
				Func: getBigFixSiteFiles,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the site.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "The display name of the site.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixSite,
				Transform:   transform.FromField("DisplayName"),
			},
			{
				Name:        "type",
				Description: "The type of the site (action, external, operator).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the site.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixSite,
			},
			{
				Name:        "global_read_permission",
				Description: "Whether the site has global read permission.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getBigFixSite,
				Transform:   transform.FromField("GlobalReadPermission"),
			},
			{
				Name:        "subscription_mode",
				Description: "The subscription mode of the site.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixSite,
				Transform:   transform.FromField("SubscriptionMode"),
			},
			{
				Name:        "gather_url",
				Description: "The gather URL of the site.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixSite,
				Transform:   transform.FromField("GatherURL"),
			},
			{
				Name:        "permissions",
				Description: "Permissions associated with the site.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixSitePermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "files",
				Description: "Files associated with the site.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixSiteFiles,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "resource",
				Description: "The resource URL of the site.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixSite,
				Transform:   transform.FromField("Resource"),
			},
		},
	}
}

func listBigFixSites(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.listBigFixSites", "service_creation_error", err)
		return nil, err
	}

	sites, err := client.Site.List()
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.listBigFixSites", "api_err", err)
		return nil, err
	}

	for _, site := range sites {
		d.StreamListItem(ctx, site)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixSite(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.getBigFixSite", "service_creation_error", err)
		return nil, err
	}

	var name, siteType string

	// Try to get name and type from quals; when hydrating from list, use h.Item
	if nameQual := d.EqualsQuals["name"]; nameQual != nil {
		name = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["type"]; typeQual != nil {
		siteType = typeQual.GetStringValue()
	}

	// When hydrating columns for list items - use the site from h.Item
	if h != nil && h.Item != nil {
		if site, ok := h.Item.(model.Site); ok {
			name = site.Name
			siteType = site.Type
		}
	}

	if name == "" || siteType == "" {
		return nil, nil
	}

	// Fetch a single site with detailed information
	site, err := client.Site.Get(ctx, name, siteType)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.getBigFixSite", "api_err", err)
		return nil, err
	}

	return site, nil
}

func getBigFixSitePermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.getBigFixSitePermissions", "service_creation_error", err)
		return nil, err
	}

	var name, siteType string

	// Try to get name and type from quals; when hydrating from list, use h.Item
	if nameQual := d.EqualsQuals["name"]; nameQual != nil {
		name = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["type"]; typeQual != nil {
		siteType = typeQual.GetStringValue()
	}

	// When hydrating columns for list items - use the site from h.Item
	if h != nil && h.Item != nil {
		if site, ok := h.Item.(model.Site); ok {
			name = site.Name
			siteType = site.Type
		}
	}

	if name == "" || siteType == "" {
		return nil, nil
	}

	// Fetch permissions for the site
	permissions, err := client.Site.GetPermissions(ctx, name, siteType)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.getBigFixSitePermissions", "api_err", err)
		return nil, err
	}

	return permissions, nil
}

func getBigFixSiteFiles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.getBigFixSiteFiles", "service_creation_error", err)
		return nil, err
	}

	var name, siteType string

	// Try to get name and type from quals; when hydrating from list, use h.Item
	if nameQual := d.EqualsQuals["name"]; nameQual != nil {
		name = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["type"]; typeQual != nil {
		siteType = typeQual.GetStringValue()
	}

	// When hydrating columns for list items - use the site from h.Item
	if h != nil && h.Item != nil {
		if site, ok := h.Item.(model.Site); ok {
			name = site.Name
			siteType = site.Type
		}
	}

	if name == "" || siteType == "" {
		return nil, nil
	}

	// Fetch files for the site
	files, err := client.Site.GetFiles(ctx, name, siteType)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_site.getBigFixSiteFiles", "api_err", err)
		return nil, err
	}

	return files, nil
}
