package bigfix

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
)

func tableBigFixTask(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_task",
		Description: "BigFix Task contains multi-step deployment workflows and complex remediation procedures with coordinated fixlets and actions for sophisticated deployments.",
		List: &plugin.ListConfig{
			ParentHydrate: listBigFixSites,
			Hydrate:       listBigFixTasks,
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
			Hydrate: getBigFixTask,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBigFixTask,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the task.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the task.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "site_name",
				Description: "The name of the site containing the task.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "site_type",
				Description: "The type of the site containing the task.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource",
				Description: "The resource URL of the task.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified",
				Description: "The last modified timestamp of the task.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "title",
				Description: "The title of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "description",
				Description: "The description of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "relevance",
				Description: "The relevance expressions of the task.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "category",
				Description: "The category of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "download_size",
				Description: "The download size of the task in bytes.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "source",
				Description: "The source of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "source_release_date",
				Description: "The source release date of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "source_severity",
				Description: "The source severity of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "delay",
				Description: "The execution delay of the task.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "mime_fields",
				Description: "MIME fields of the task.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "default_action",
				Description: "The default action of the task.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixTask,
			},
			{
				Name:        "actions",
				Description: "All actions available in the task.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixTask,
			},
		},
	}
}

func listBigFixTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_task.listBigFixTasks", "service_creation_error", err)
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

	// Fetch tasks for the site
	tasks, err := client.Task.List(ctx, site.Name, site.Type)
	if err != nil {
		// In the case of parent hydrate the Ignore config is not being honored.
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("bigfix_task.listBigFixTasks", "api_err", err)
		return nil, err
	}

	// Return each task as a separate row
	for _, task := range tasks {
		d.StreamListItem(ctx, task)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixTask(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_task.getBigFixTask", "service_creation_error", err)
		return nil, err
	}

	var siteName, siteType string
	var taskID int

	if h.Item != nil {
		site := h.Item.(model.Task)
		siteName = site.SiteName
		siteType = site.SiteType
		taskID = site.ID
	}

	if nameQual := d.EqualsQuals["site_name"]; nameQual != nil {
		siteName = nameQual.GetStringValue()
	}
	if typeQual := d.EqualsQuals["site_type"]; typeQual != nil {
		siteType = typeQual.GetStringValue()
	}
	if idQual := d.EqualsQuals["id"]; idQual != nil {
		taskID = int(idQual.GetInt64Value())
	}

	if siteName == "" || siteType == "" || taskID == 0 {
		return nil, nil
	}

	// Fetch the specific task
	task, err := client.Task.Get(ctx, siteName, siteType, taskID)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_task.getBigFixTask", "api_err", err)
		return nil, err
	}

	return task, nil
}
