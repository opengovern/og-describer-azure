package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDNSRecordSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_dns_record_set",
		Description: "Azure DNS Record Sets",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDNSRecordSet,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDNSRecordSet,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:      "subscription",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Subscription"),
			},
			{
				Name:        "zone_id",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the DNS zone.",
				Transform:   transform.FromField("Description.DNSZoneID")},
			{
				Name:        "zone_name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the DNS zone.",
				Transform:   transform.FromField("Description.DNSZoneName")},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the DNS zone.",
				Transform:   transform.FromField("Description.DNSRecordSet.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a DNS zone uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DNSRecordSet.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DNSRecordSet.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the DNS Record Set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DNSRecordSet.Type")},
			{
				Name:        "a_records",
				Description: "The list of A records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.ARecords")},
			{
				Name:        "aaaa_records",
				Description: "The list of AAAA records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.AaaaRecords")},
			{
				Name:        "caa_records",
				Description: "The list of CAA records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.CaaRecords")},
			{
				Name:        "cname_record",
				Description: "The list of CNAME record in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.CnameRecord")},
			{
				Name:        "mx_records",
				Description: "The list of MX records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.MxRecords")},
			{
				Name:        "ns_records",
				Description: "The list of NS records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.NsRecords")},
			{
				Name:        "soa_records",
				Description: "The list of SOA records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.SoaRecord")},
			{
				Name:        "srv_records",
				Description: "The list of SRV records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.SrvRecords")},
			{
				Name:        "txt_records",
				Description: "The list of TXT records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.TxtRecords")},
			{
				Name:        "fqdn",
				Description: "Fully qualified domain name of the record set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.Fqdn")},
			{
				Name:        "ttl",
				Description: "The TTL (time-to-live) of the records in the record set.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.TTL")},
			{
				Name:        "target_resource",
				Description: "A reference to an azure resource from where the dns resource value is taken.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.TargetResource")},
			{
				Name:        "provisioning_state",
				Description: "provisioning State of the record set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.ProvisioningState")},
			{
				Name:        "metadata",
				Description: "The list of CNAME records in the record set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DNSRecordSet.Properties.Metadata")},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
