package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/dns"
	"github.com/chronark/vercel-go/endpoints/domain"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableVercelDnsRecord(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_dns_record",
		Description: "DNS records in the Vercel account.",
		List: &plugin.ListConfig{
			ParentHydrate: listDomain,
			Hydrate:       listDnsRecord,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "domain_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Domain.Name"), Description: "Domain name the record belongs to."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the DNS record."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the DNS record."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Type of the DNS record."},
			{Name: "ttl", Type: proto.ColumnType_INT, Description: "Time To Live of the DNS record."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the DNS record was created."},
			{Name: "creator", Type: proto.ColumnType_STRING, Description: "Creator of the DNS record."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the DNS record."},
			{Name: "mx_priority", Type: proto.ColumnType_INT, Description: "MX priority of the DNS record."},
			{Name: "priority", Type: proto.ColumnType_INT, Description: "Priority of the DNS record."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the DNS record."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the DNS record was created."},
		}),
	}
}

type dnsRecordRow struct {
	Domain domain.Domain
	dns.Record
}

func listDnsRecord(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_dns.listDnsRecord", "connection_error", err)
		return nil, err
	}

	parentDomain := h.Item.(domain.Domain)
	opts := dns.ListDnsRequest{Domain: parentDomain.Name, Limit: 100}

	for {
		plugin.Logger(ctx).Debug("vercel_dns.listDnsRecord", "opts", opts)
		res, err := conn.Dns.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("vercel_dns.listDnsRecord", "query_error", err)
			return nil, err
		}
		for _, i := range res.Records {
			d.StreamListItem(ctx, dnsRecordRow{parentDomain, i})
		}
		plugin.Logger(ctx).Debug("vercel_dns.listDnsRecord", "pagination", res.Pagination)
		if res.Pagination.Next == 0 {
			break
		}
		opts.Until = int(res.Pagination.Next)
	}

	return nil, nil
}
