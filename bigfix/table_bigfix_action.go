package bigfix

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBigFixAction(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_action",
		Description: "BigFix Action",
		List: &plugin.ListConfig{
			Hydrate: listBigFixActions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Required},
			},
			Hydrate: getBigFixAction,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBigFixAction,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the action.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the action.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource",
				Description: "The resource URL of the action.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified",
				Description: "The last modified timestamp of the action.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "title",
				Description: "The title of the action.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "relevance",
				Description: "The relevance expression of the action.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "action_script",
				Description: "The action script content.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "success_criteria",
				Description: "The success criteria for the action.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "settings",
				Description: "The settings configuration for the action.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "settings_locks",
				Description: "The settings locks for the action.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "target",
				Description: "The target configuration for the action.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixAction,
			},
			{
				Name:        "is_urgent",
				Description: "Whether the action is marked as urgent.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getBigFixAction,
			},
		},
	}
}

func listBigFixActions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_action.listBigFixActions", "service_creation_error", err)
		return nil, err
	}

	// Get all actions
	actions, err := client.Action.List(ctx)
	if err != nil {
		// In the case of parent hydrate the Ignore config is not being honored.
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("bigfix_action.listBigFixActions", "api_err", err)
		return nil, err
	}

	// Stream the actions
	for _, action := range actions {
		d.StreamListItem(ctx, action)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixAction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the action from the hydrate data
	var actionID int

	if h.Item != nil {
		action := h.Item.(model.Action)
		actionID = action.ID
	}

	if idQual := d.EqualsQuals["id"]; idQual != nil {
		actionID = int(idQual.GetInt64Value())
	}

	if actionID == 0 {
		return nil, nil
	}

	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_action.getBigFixAction", "service_creation_error", err)
		return nil, err
	}

	// Get the action detail
	action, err := client.Action.Get(ctx, actionID)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_action.getBigFixAction", "api_error", err)
		return nil, err
	}

	return action, nil
}
