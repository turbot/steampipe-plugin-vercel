package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/domain"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableVercelDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_domain",
		Description: "Domains in the Vercel account.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the domain."},
			// Other columns
			{Name: "bought_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("BoughtAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the domain was bought."},
			{Name: "config_verified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ConfigVerifiedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the domain configuration was verified."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the domain was created."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ExpiresAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the domain expires."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "ID of the domain."},
			{Name: "ns_verified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("NsVerifiedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the name server was verified."},
			{Name: "service_type", Type: proto.ColumnType_STRING, Description: "Service provided by the domain, e.g. external."},
			{Name: "transferred_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("TransferredAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the domain was created."},
			{Name: "txt_verified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("TxtVerifiedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the domain was created."},
			{Name: "verification_record", Type: proto.ColumnType_STRING, Description: "Verification record for the domain."},
			{Name: "cdn_enabled", Type: proto.ColumnType_BOOL, Description: "If true, then the Content Delivery Network is enabled for this domain."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Description: "True if the domain is verified."},
			{Name: "name_servers", Type: proto.ColumnType_JSON, Description: "Name servers for the domain."},
			{Name: "renew", Type: proto.ColumnType_BOOL, Description: "True if the domain should auto-renew."},
			{Name: "intended_name_servers", Type: proto.ColumnType_JSON, Description: "Intended name servers for the domain."},
			{Name: "creator", Type: proto.ColumnType_JSON, Description: "Creator of the domain."},
			{Name: "zone", Type: proto.ColumnType_BOOL, Description: "Zone of the domain."},
		}),
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_domain.listDomain", "connection_error", err)
		return nil, err
	}

	opts := domain.ListDomainsRequest{Limit: 100}
	for {
		plugin.Logger(ctx).Debug("vercel_domain.listDomain", "opts", opts)
		res, err := conn.Domain.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("vercel_domain.listDomain", "query_error", err)
			return nil, err
		}
		for _, i := range res.Domains {
			d.StreamListItem(ctx, i)
		}
		plugin.Logger(ctx).Debug("vercel_domain.listDomain", "pagination", res.Pagination)
		if res.Pagination.Next == 0 {
			break
		}
		opts.Until = int(res.Pagination.Next)
	}

	return nil, nil
}
