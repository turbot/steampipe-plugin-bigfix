package bigfix

import (
	"context"

	"github.com/turbot/steampipe-plugin-bigfix/api/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableBigFixComputer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bigfix_computer",
		Description: "BigFix Computer.",
		List: &plugin.ListConfig{
			Hydrate: listBigFixComputers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getBigFixComputer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"not found"}),
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
			},
			{
				Name:        "id",
				Description: "The id of the computer.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "os",
				Description: "The operating system.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("OS"),
			},
			{
				Name:        "cpu",
				Description: "The CPU information.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("CPU"),
			},
			{
				Name:        "ip_address",
				Description: "The IP address of the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("IPAddress"),
			},
			{
				Name:        "ipv6_address",
				Description: "The IPv6 address of the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("IPv6Address"),
			},
			{
				Name:        "dns_name",
				Description: "The DNS name of the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("DNSName"),
			},
			{
				Name:        "mac_address",
				Description: "The MAC address of the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("MACAddress"),
			},
			{
				Name:        "os_family",
				Description: "The operating system family.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("OSFamily"),
			},
			{
				Name:        "os_name",
				Description: "The operating system name.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("OSName"),
			},
			{
				Name:        "os_version",
				Description: "The operating system version.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("OSVersion"),
			},
			{
				Name:        "user_name",
				Description: "The user name associated with the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("UserName"),
			},
			{
				Name:        "ram",
				Description: "The amount of RAM in the computer.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("RAM"),
			},
			{
				Name:        "locked",
				Description: "Whether the computer is locked.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("Locked"),
			},
			{
				Name:        "bes_relay_selection",
				Description: "The BES relay selection method.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("BESRelaySelection"),
			},
			{
				Name:        "relay",
				Description: "The relay information.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("Relay"),
			},
			{
				Name:        "distance_to_bes_relay",
				Description: "The distance to BES relay.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("DistanceToBESRelay"),
			},
			{
				Name:        "agent_type",
				Description: "The agent type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("AgentType"),
			},
			{
				Name:        "device_type",
				Description: "The device type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("DeviceType"),
			},
			{
				Name:        "agent_version",
				Description: "The agent version.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("AgentVersion"),
			},
			{
				Name:        "computer_type",
				Description: "The computer type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("ComputerType"),
			},
			{
				Name:        "license_type",
				Description: "The license type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("LicenseType"),
			},
			{
				Name:        "free_space_on_system",
				Description: "The free space on system drive.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("FreeSpaceOnSystem"),
			},
			{
				Name:        "total_size_of_system",
				Description: "The total size of system drive.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("TotalSizeOfSystem"),
			},
			{
				Name:        "bios",
				Description: "The BIOS information.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("BIOS"),
			},
			{
				Name:        "subnet_address",
				Description: "The subnet address.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("SubnetAddress"),
			},
			{
				Name:        "client_settings",
				Description: "Array of client settings name-value pairs.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("ClientSettings"),
			},
			{
				Name:        "subscribed_sites",
				Description: "Array of subscribed sites.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("SubscribedSites"),
			},
			{
				Name:        "last_report_time",
				Description: "The last time the computer reported to the BigFix server.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastReportTime"),
			},
			// Keep this column for any custom properties that may have been added to the computer but are not captured in the specific columns above
			{
				Name:        "properties",
				Description: "All raw properties from the BigFix API.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBigFixComputer,
				Transform:   transform.FromField("Properties"),
			},
		},
	}
}

func listBigFixComputers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_computer.listBigFixComputers", "service_creation_error", err)
		return nil, err
	}

	computers, err := client.Computer.List()
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_computer.listBigFixComputers", "api_err", err)
		return nil, err
	}

	for _, computer := range computers {
		d.StreamListItem(ctx, computer)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getBigFixComputer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := NewService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_computer.getBigFixComputer", "service_creation_error", err)
		return nil, err
	}

	// Try to get id from quals; when hydrating from list, use h.Item
	var id int64
	if qual := d.EqualsQuals["id"]; qual != nil {
		id = qual.GetInt64Value()
	} else if h != nil && h.Item != nil {
		// When hydrating columns for list items - use reflection to get ID from model.Computer
		if computer, ok := h.Item.(model.Computer); ok {
			id = int64(computer.ID)
		}
	}
	if id == 0 {
		return nil, nil
	}

	// Fetch a single computer with detailed Property fields
	computer, err := client.Computer.Get(ctx, int(id))
	if err != nil {
		plugin.Logger(ctx).Error("bigfix_computer.getBigFixComputer", "api_err", err)
		return nil, err
	}

	return computer, nil
}
