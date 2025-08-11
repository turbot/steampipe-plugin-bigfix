package bigfix

import (
	"context"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBigFixRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_role",
		Description: "BigFix Role",
		List: &plugin.ListConfig{
			Hydrate: listBigFixRoles,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Required},
			},
			Hydrate: getBigFixRole,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBigFixRole,
				IgnoreConfig: &plugin.IgnoreConfig{
					ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the role.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource",
				Description: "The resource URL of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified",
				Description: "The last modified timestamp of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "master_operator",
				Description: "Whether the role has master operator privileges.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "custom_content",
				Description: "Whether the role can access custom content.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "show_other_actions",
				Description: "Whether the role can show other actions.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "stop_other_actions",
				Description: "Whether the role can stop other actions.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "can_create_actions",
				Description: "Whether the role can create actions.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "post_action_behavior_privilege",
				Description: "The post action behavior privilege for the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "action_script_commands_privilege",
				Description: "The action script commands privilege for the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "can_send_multiple_refresh",
				Description: "Whether the role can send multiple refresh commands.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "can_submit_queries",
				Description: "Whether the role can submit queries.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "can_lock",
				Description: "Whether the role can lock resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "unmanaged_asset_privilege",
				Description: "The unmanaged asset privilege for the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "interface_logins",
				Description: "The interface login permissions for the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("InterfaceLogins"),
			},
		},
	}
}

func listBigFixRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_role.listBigFixRoles", "service_creation_error", err)
		return nil, err
	}

	// Get all roles
	roles, err := client.Role.List(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_role.listBigFixRoles", "api_err", err)
		return nil, err
	}

	// Stream the roles
	for _, role := range roles {
		d.StreamListItem(ctx, role)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the role from the hydrate data
	var roleID int

	if h.Item != nil {
		role := h.Item.(model.Role)
		roleID = role.ID
	}

	if idQual := d.EqualsQuals["id"]; idQual != nil {
		roleID = int(idQual.GetInt64Value())
	}

	if roleID == 0 {
		return nil, nil
	}

	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_role.getBigFixRole", "service_creation_error", err)
		return nil, err
	}

	// Get the role detail
	role, err := client.Role.Get(ctx, roleID)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_role.getBigFixRole", "api_error", err)
		return nil, err
	}

	return role, nil
}
