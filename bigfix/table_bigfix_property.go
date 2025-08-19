package bigfix

import (
	"context"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBigFixProperty(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_property",
		Description: "BigFix Property.",
		List: &plugin.ListConfig{
			Hydrate: listBigFixProperties,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Required},
			},
			Hydrate: getBigFixProperty,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the property.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource",
				Description: "The resource URL of the property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified",
				Description: "The last modified timestamp of the property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_reserved",
				Description: "Whether the property is reserved.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "definition",
				Description: "The definition of the property.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixProperty,
			},
		},
	}
}

func listBigFixProperties(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_property.listBigFixProperties", "service_creation_error", err)
		return nil, err
	}

	// Get all properties
	properties, err := client.Property.List(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_property.listBigFixProperties", "api_err", err)
		return nil, err
	}

	// Stream the properties
	for _, property := range properties {
		d.StreamListItem(ctx, property)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixProperty(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Get the property from the hydrate data
	var propertyID int

	if h.Item != nil {
		property := h.Item.(model.BigFixProperty)
		propertyID = property.ID
	}

	if idQual := d.EqualsQuals["id"]; idQual != nil {
		propertyID = int(idQual.GetInt64Value())
	}

	if propertyID == 0 {
		return nil, nil
	}

	// Create the service
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_property.getBigFixProperty", "service_creation_error", err)
		return nil, err
	}

	// Get the property detail
	property, err := client.Property.Get(ctx, propertyID)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_property.getBigFixProperty", "api_error", err)
		return nil, err
	}

	return property, nil
}
