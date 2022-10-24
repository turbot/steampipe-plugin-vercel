# Table: vercel_deployment

List deployments in your account.

## Examples

### List all deployments

```sql
select
  name as project,
  state,
  url,
  creator ->> 'email' as creator,
  meta ->> 'githubCommitMessage' as commit_message,
  meta ->> 'githubCommitRef' as commit_ref
from
  vercel_deployment
order by 
  created_at desc
```
