![image](https://hub.steampipe.io/images/plugins/turbot/vercel-social-graphic.png)

# Vercel Plugin for Steampipe

Use SQL to query projects, teams, domains and more from Vercel.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/vercel)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/vercel/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-vercel/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install vercel
```

Configure your credentials in `~/.steampipe/config/vercel.spc`:

```hcl
connection "vercel" {
  plugin    = "vercel"
  api_token = "YwbeYCAYfpdPKSj9yd18JUXX"
  team      = "mycompany" # Optional
}
```

Run steampipe:

```shell
steampipe query
```

Query your projects:

```sql
select
  name,
  framework,
  updated_at
from
  vercel_project;
```

```
+------------------+-----------+---------------------------+
| name             | framework | updated_at                |
+------------------+-----------+---------------------------+
| my-homepage-io   | nextjs    | 2022-07-12T14:04:42-04:00 |
| another-site-com | nextjs    | 2022-07-12T14:28:38-04:00 |
+------------------+-----------+---------------------------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-vercel.git
cd steampipe-plugin-vercel
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/vercel.spc
```

Try it!

```
steampipe query
> .inspect vercel
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-vercel/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-vercel/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Vercel Plugin](https://github.com/turbot/steampipe-plugin-vercel/labels/help%20wanted)
