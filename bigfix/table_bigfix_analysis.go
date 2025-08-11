package bigfix

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
)

func tableBigFixAnalysis(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_analysis",
		Description: "BigFix Analysis",
		List: &plugin.ListConfig{
			ParentHydrate: listBigFixSites,
			Hydrate:       listBigFixAnalyses,
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
			Hydrate: getBigFixAnalysis,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBigFixAnalysis,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the analysis.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the analysis.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "site_name",
				Description: "The name of the site containing the analysis.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "site_type",
				Description: "The type of the site containing the analysis.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource",
				Description: "The resource URL of the analysis.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified",
				Description: "The last modified timestamp of the analysis.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "title",
				Description: "The title of the analysis.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "description",
				Description: "The description of the analysis.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "relevance",
				Description: "The relevance expressions of the analysis.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "category",
				Description: "The category of the analysis.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "source",
				Description: "The source of the analysis.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "source_release_date",
				Description: "The source release date of the analysis.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "delay",
				Description: "The evaluation delay of the analysis.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "mime_fields",
				Description: "MIME fields of the analysis.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixAnalysis,
			},
			{
				Name:        "properties",
				Description: "Properties defined in the analysis.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixAnalysis,
			},
		},
	}
}

func listBigFixAnalyses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_analysis.listBigFixAnalyses", "service_creation_error", err)
		return nil, err
	}

	// Get site from parent hydrate
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

	// Fetch analyses for the site
	analyses, err := client.Analysis.List(ctx, site.Name, site.Type)
	if err != nil {
		// In the case of parent hydrate the Ignore config is not being honored.
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("bigfix_analysis.listBigFixAnalyses", "api_err", err)
		return nil, err
	}

	// Return each analysis as a separate row
	for _, analysis := range analyses {
		d.StreamListItem(ctx, analysis)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixAnalysis(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_analysis.getBigFixAnalysis", "service_creation_error", err)
		return nil, err
	}

	var siteName, siteType string
	var analysisID int

	if h.Item != nil {
		site := h.Item.(model.Analysis)
		siteName = site.SiteName
		siteType = site.SiteType
		analysisID = site.ID
	}

	if nameQual := d.EqualsQuals["site_name"]; nameQual != nil {
		siteName = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["site_type"]; typeQual != nil {
		siteType = typeQual.GetStringValue()
	}
	if idQual := d.EqualsQuals["id"]; idQual != nil {
		analysisID = int(idQual.GetInt64Value())
	}

	if siteName == "" || siteType == "" || analysisID == 0 {
		return nil, nil
	}

	// Fetch the specific analysis
	analysis, err := client.Analysis.Get(ctx, siteName, siteType, analysisID)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_analysis.getBigFixAnalysis", "api_err", err)
		return nil, err
	}

	return analysis, nil
}
