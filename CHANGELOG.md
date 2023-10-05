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
