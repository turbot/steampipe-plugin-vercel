# Table: vercel_project

List projects in your account.

## Examples

### List all projects

```sql
select
  name,
  framework,
  updated_at
from
  vercel_project
```

### Projects not updated in the last year

```sql
select
  name,
  framework,
  updated_at
from
  vercel_project
where
  updated_at < now() - interval '1 year'
```

### Latest deployments

```sql
select
  to_timestamp((d ->> 'createdAt')::bigint / 1000) as created_at,
  name,
  d ->> 'url' as url,
  d -> 'creator' ->> 'username' as creator_username
from
  vercel_project as p,
  jsonb_array_elements(latest_deployments) as d
order by
  created_at desc
```

### Current production target by project

```sql
select
  name,
  to_timestamp((targets -> 'production' ->> 'createdAt')::bigint / 1000) as created_at,
  targets -> 'production' ->> 'url' as url,
  targets -> 'production' -> 'meta' ->> 'githubOrg' as github_org,
  targets -> 'production' -> 'meta' ->> 'githubRepo' as github_repo,
  targets -> 'production' -> 'meta' ->> 'githubCommitSha' as github_commit_sha,
  targets -> 'production' -> 'meta' ->> 'githubCommitAuthorName' as github_commit_author_name,
  targets -> 'production' -> 'meta' ->> 'githubCommitMessage' as github_commit_message
from
  vercel_project
order by
  name
```

### List all project environment variables

```sql
select
  name,
  e ->> 'key',
  e ->> 'target'
from
  vercel_project as p,
  jsonb_array_elements(env) as e
```

### Environment variables that are not encrypted

```sql
select
  name,
  e ->> 'key',
  e ->> 'type'
from
  vercel_project as p,
  jsonb_array_elements(env) as e
where
  e ->> 'type' != 'encrypted'
```

### Production environment variables older than 90 days

```sql
select
  name,
  e ->> 'key',
  e ->> 'target',
  to_timestamp((e ->> 'createdAt')::bigint / 1000)
from
  vercel_project as p,
  jsonb_array_elements(env) as e
where
  e -> 'target' ? 'production'
  and to_timestamp((e ->> 'createdAt')::bigint / 1000) < now() - interval '90 days'
```

### Projects by framework environment variables older than 90 days

```sql
select
  framework,
  count(id)
from
  vercel_project
group by
  framework
order by
  count desc
```
