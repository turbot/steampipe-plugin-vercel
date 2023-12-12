---
title: "Steampipe Table: vercel_secret - Query Vercel Secrets using SQL"
description: "Allows users to query Vercel Secrets, providing detailed information about secrets stored in Vercel platform."
---

# Table: vercel_secret - Query Vercel Secrets using SQL

Vercel Secrets is a feature within the Vercel platform that allows users to store sensitive data, like API keys, securely. These secrets can be used in environment variables for Vercel projects, ensuring that sensitive data is not exposed in your code or Vercel logs. They provide an additional layer of security for your applications, keeping your sensitive data safe and secure.

## Table Usage Guide

The `vercel_secret` table provides insights into the secrets stored within Vercel. As a developer or security analyst, explore secret-specific details through this table, including secret names, created timestamp, and the projects that use these secrets. Utilize it to uncover information about secrets, such as those that are outdated, unused, or associated with specific projects, ensuring the security and integrity of your applications.

## Examples

### List all secrets
Discover the segments that contain confidential data by identifying instances where certain projects have created secrets. This can help in managing and reviewing the configuration for data security across different projects.

```sql+postgres
select
  project_id,
  name,
  uid,
  created_at
from
  vercel_secret;
```

```sql+sqlite
select
  project_id,
  name,
  uid,
  created_at
from
  vercel_secret;
```

### Secrets more than 1 year old
Identify older secrets within your project that may pose a security risk. This query is useful for maintaining good security hygiene by pinpointing secrets that have been in use for over a year.

```sql+postgres
select
  project_id,
  name,
  uid,
  created_at
from
  vercel_secret
where
  created_at < now() - interval '1 year';
```

```sql+sqlite
select
  project_id,
  name,
  uid,
  created_at
from
  vercel_secret
where
  created_at < datetime('now', '-1 year');
```

### Secrets used by environment variables
Discover the secrets that are being utilized by your environment variables in your projects. This can help in understanding the linkage between your projects and the secrets, enhancing your project's security and management.

```sql+postgres
select
  p.name as project_name,
  e ->> 'key' as env_var,
  e ->> 'type' as env_var_type,
  s.name as secret_name,
  s.uid as secret_uid,
  s.created_at as secret_created_at
from
  vercel_project as p,
  jsonb_array_elements(env) as e,
  vercel_secret as s
where
  e ->> 'type' = 'secret'
  and e ->> 'value' = s.uid;
```

```sql+sqlite
select
  p.name as project_name,
  json_extract(e.value, '$.key') as env_var,
  json_extract(e.value, '$.type') as env_var_type,
  s.name as secret_name,
  s.uid as secret_uid,
  s.created_at as secret_created_at
from
  vercel_project as p,
  json_each(env) as e,
  vercel_secret as s
where
  json_extract(e.value, '$.type') = 'secret'
  and json_extract(e.value, '$.value') = s.uid;
```