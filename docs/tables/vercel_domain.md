# Table: vercel_domain

List domains in your account.

## Examples

### List all domains

```sql
select
  name,
  verified,
  created_at,
  expires_at
from
  vercel_domain;
```

### Domains expiring in the next 90 days

Lists domains expiring soon. Domains managed outside of Vercel will not be included in results.

```sql
select
  name,
  expires_at
from
  vercel_domain
where
  expires_at < now() + interval '90 days';
```
