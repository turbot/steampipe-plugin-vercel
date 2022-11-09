# Table: vercel_deployment

List deployments in your account.

## Examples

### List recent deployments

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
where
  created_at > now() - interval '2 weeks'
order by 
  created_at desc
```
