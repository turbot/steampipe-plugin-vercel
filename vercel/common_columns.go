package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/user"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "user_uid",
			Description: "Unique identifier of the user.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getUserUid,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getUserUidMemoized = plugin.HydrateFunc(getUserInfo).Memoize(memoize.WithCacheKeyFunction(getUserUidCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getUserUid(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	res, err := getUserUidMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	user := res.(user.User)
	
	return user.Uid, nil
}

// Build a cache key for the call to getUserUidCacheKey.
func getUserUidCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getUserUid"
	return key, nil
}

func getUserInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_user.listUser", "connection_error", err)
		return nil, err
	}

	res, err := conn.User.Get()
	if err != nil {
		plugin.Logger(ctx).Error("vercel_user.listUser", "query_error", err)
		return nil, err
	}
	return res.User, nil
}
