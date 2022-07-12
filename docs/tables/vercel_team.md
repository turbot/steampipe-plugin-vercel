# Table: vercel_team

Teams allow groups of users to work together. This table includes all teams the authenticated user has access to.

## Examples

### List all teams

```sql
select
  id,
  slug,
  name,
  description
from
  vercel_team
```

### Get role of the authenticated user in each team

```sql
select
  name,
  membership ->> 'role' as role
from
  vercel_team
```

## Number of invoiced seats per team

```sql
select
  name,
  billing -> 'invoiceItems' -> 'teamSeats' ->> 'quantity' as seats
from
  vercel_team
```
