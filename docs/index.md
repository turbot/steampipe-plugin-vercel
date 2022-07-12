---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/vercel.svg"
brand_color: "#000000"
display_name: "Vercel"
short_name: "vercel"
description: "Steampipe plugin to query projects, teams, domains and more from Vercel."
og_description: "Query Vercel with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/vercel-social-graphic.png"
---

# Vercel + Steampipe

[Vercel](https://vercel.com) is a cloud hosting service for frontend frameworks and static sites.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List projects in your Vercel account:

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/vercel/tables)**

## Get started

### Install

Download and install the latest Vercel plugin:

```bash
steampipe plugin install vercel
```

### Configuration

Installing the latest vercel plugin will create a config file (`~/.steampipe/config/vercel.spc`) with a single connection named `vercel`:

```hcl
connection "vercel" {
  plugin    = "vercel"
  api_token = "YwbeYCAYfpdPKSj9yd18JUXX"
  team      = "mycompany"
}
```

- `api_token` - [API token](https://vercel.com/support/articles/how-do-i-use-a-vercel-api-access-token) to access your account.
- `team` - Optional team to target.

Environment variables are also available as an alternate configuration method:
* `VERCEL_API_TOKEN`
* `VERCEL_TEAM`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-vercel
- Community: [Slack Channel](https://steampipe.io/community/join)
