package bigfix

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBigFixFixlet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_fixlet",
		Description: "BigFix Fixlet.",
		List: &plugin.ListConfig{
			ParentHydrate: listBigFixSites,
			Hydrate:       listBigFixFixlets,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "site_name", Require: plugin.Optional},
				{Name: "site_type", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "site_name", Require: plugin.Required},
				{Name: "site_type", Require: plugin.Required},
				{Name: "id", Require: plugin.Required},
			},
			Hydrate: getBigFixFixlet,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBigFixFixlet,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the fixlet.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the fixlet.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "site_name",
				Description: "The name of the site containing the fixlet.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "site_type",
				Description: "The type of the site containing the fixlet.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource",
				Description: "The resource URL of the fixlet.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified",
				Description: "The last modified timestamp of the fixlet.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "title",
				Description: "The title of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "description",
				Description: "The description of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "relevance",
				Description: "The relevance expressions of the fixlet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "category",
				Description: "The category of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "download_size",
				Description: "The download size of the fixlet in bytes.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "source",
				Description: "The source of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "source_id",
				Description: "The source ID of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "source_release_date",
				Description: "The source release date of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "source_severity",
				Description: "The source severity of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "cve_names",
				Description: "The CVE names associated with the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "mime_fields",
				Description: "MIME fields of the fixlet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "delay",
				Description: "The evaluation delay of the fixlet.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "default_action",
				Description: "The default action of the fixlet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixFixlet,
			},
			{
				Name:        "actions",
				Description: "All actions available in the fixlet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixFixlet,
			},
		},
	}
}

func listBigFixFixlets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the site from the parent hydrate
	site := h.Item.(model.Site)

	// Check if optional key quals are provided to filter the results
	var targetSiteName, targetSiteType string
	if nameQual := d.EqualsQuals["site_name"]; nameQual != nil {
		targetSiteName = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["site_type"]; typeQual != nil {
		targetSiteType = typeQual.GetStringValue()
	}

	// If optional quals are provided, only fetch if they match the current site
	if targetSiteName != "" && targetSiteName != site.Name {
		return nil, nil
	}
	if targetSiteType != "" && targetSiteType != site.Type {
		return nil, nil
	}

	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_fixlet.listBigFixFixlets", "service_creation_error", err)
		return nil, err
	}

	// Get the fixlets for this site
	fixlets, err := client.Fixlet.List(ctx, site.Name, site.Type)
	if err != nil {
		// In the case of parent hydrate the Ignore config is not being honored.
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("bigfix_fixlet.listBigFixFixlets", "api_err", err)
		return nil, err
	}

	// Stream the fixlets
	for _, fixlet := range fixlets {
		d.StreamListItem(ctx, fixlet)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixFixlet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the fixlet from the hydrate data
	var siteName, siteType string
	var fixletID int

	if h.Item != nil {
		fixlet := h.Item.(model.Fixlet)
		siteName = fixlet.SiteName
		siteType = fixlet.SiteType
		fixletID = fixlet.ID
	}

	if nameQual := d.EqualsQuals["site_name"]; nameQual != nil {
		siteName = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["site_type"]; typeQual != nil {
		siteType = typeQual.GetStringValue()
	}
	if idQual := d.EqualsQuals["id"]; idQual != nil {
		fixletID = int(idQual.GetInt64Value())
	}

	if siteName == "" || siteType == "" || fixletID == 0 {
		return nil, nil
	}

	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_fixlet.getBigFixFixlet", "service_creation_error", err)
		return nil, err
	}

	// Get the fixlet detail
	fixlet, err := client.Fixlet.Get(ctx, siteName, siteType, fixletID)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_fixlet.getBigFixFixlet", "api_error", err)
		return nil, err
	}

	return fixlet, nil
}
