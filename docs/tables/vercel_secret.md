# Table: vercel_secret

Secrets are a legacy way to encrypt variables in Vercel, and have been deprecated. All environment variables are now encrypted by default.

## Examples

### List all secrets

```sql
select
  project_id,
  name,
  uid,
  created_at
from
  vercel_secret
```

### Secrets more than 1 year old

```sql
select
  project_id,
  name,
  uid,
  created_at
from
  vercel_secret
where
  created_at < now() - interval '1 year'
```

### Secrets used by environment variables

```sql
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
  and e ->> 'value' = s.uid
```
