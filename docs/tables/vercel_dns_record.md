# Table: vercel_dns_record

List DNS records in your account.

## Examples

### List all DNS records for all domains

```sql
select
  domain_name,
  name,
  type,
  value,
  ttl
from
  vercel_dns_record
```

### List all A records for all domains

```sql
select
  domain_name,
  type,
  value,
  ttl
from
  vercel_dns_record
where
  type = 'A'
```
