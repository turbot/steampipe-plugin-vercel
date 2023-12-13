## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#22](https://github.com/turbot/steampipe-plugin-vercel/pull/22))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#22](https://github.com/turbot/steampipe-plugin-vercel/pull/22))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-vercel/blob/main/docs/LICENSE). ([#22](https://github.com/turbot/steampipe-plugin-vercel/pull/22))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#21](https://github.com/turbot/steampipe-plugin-vercel/pull/21))

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#14](https://github.com/turbot/steampipe-plugin-vercel/pull/14))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#12](https://github.com/turbot/steampipe-plugin-vercel/pull/12))
- Recompiled plugin with Go version `1.21`. ([#12](https://github.com/turbot/steampipe-plugin-vercel/pull/12))

## v0.3.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#7](https://github.com/turbot/steampipe-plugin-vercel/pull/7))

## v0.2.0 [2022-11-24]

_What's new?_

- New tables added:
  - [vercel_deployment](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_deployment) ([#4](https://github.com/turbot/steampipe-plugin-vercel/pull/4))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.8](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v418-2022-09-08) which increases the default open file limit. ([#5](https://github.com/turbot/steampipe-plugin-vercel/pull/5))

## v0.1.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#2](https://github.com/turbot/steampipe-plugin-vercel/pull/2))
- Recompiled plugin with Go version `1.19`. ([#2](https://github.com/turbot/steampipe-plugin-vercel/pull/2))

## v0.0.1 [2022-07-19]

_What's new?_

- New tables added:
  - [vercel_dns_record](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_dns_record)
  - [vercel_domain](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_domain)
  - [vercel_project](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_project)
  - [vercel_secret](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_secret)
  - [vercel_team](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_team)
  - [vercel_user](https://hub.steampipe.io/plugins/turbot/vercel/tables/vercel_user)
