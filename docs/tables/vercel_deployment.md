# Table: vercel_deployment

List deployments in your account.

## Examples

### List all deployments

```sql
select
  name as project,
  to_char(created_at, 'YYYY-MM-DD HH24:mm') as created_at,
  now() - ready > interval '1 min' as ready,
  url,
  creator ->> 'email' as creator,
  meta ->> 'githubCommitMessage' as commit
from
  vercel_deployment
order by 
  created_at desc;
```
