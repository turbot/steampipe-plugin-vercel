---
title: "Steampipe Table: vercel_team - Query Vercel Teams using SQL"
description: "Allows users to query Vercel Teams, specifically providing details about each team's information, including ID, name, slug, description, and more."
---

# Table: vercel_team - Query Vercel Teams using SQL

A Vercel Team is a collaborative workspace in Vercel, a deployment and hosting platform. It allows multiple users to work together on projects, share resources, and manage permissions. Teams can be used to manage both open-source projects and commercial projects in a shared workspace.

## Table Usage Guide

The `vercel_team` table provides insights into teams within Vercel. As a developer or DevOps engineer, explore team-specific details through this table, including team ID, name, slug, description, and more. Utilize it to uncover information about teams, such as their creation time, collaboration details, and the users associated with each team.

## Examples

### List all teams
Explore which teams are currently set up in your Vercel environment. This can help in managing team access and permissions effectively.

```sql+postgres
select
  id,
  slug,
  name,
  description
from
  vercel_team;
```

```sql+sqlite
select
  id,
  slug,
  name,
  description
from
  vercel_team;
```

### Get role of the authenticated user in each team
The first query allows you to find out the role of the authenticated user in each team, which can be useful in managing team permissions and roles. The second query provides information on the number of invoiced seats per team, which can be essential for budgeting and resource allocation.

```sql+postgres
select
  name,
  membership ->> 'role' as role
from
  vercel_team;
```

```sql+sqlite
select
  name,
  json_extract(membership, '$.role') as role
from
  vercel_team;
```

## Number of invoiced seats per team

```sql+postgres
select
  name,
  billing -> 'invoiceItems' -> 'teamSeats' ->> 'quantity' as seats
from
  vercel_team;
```

```sql+sqlite
select
  name,
  json_extract(json_extract(json_extract(billing, '$.invoiceItems'), '$.teamSeats'), '$.quantity') as seats
from
  vercel_team;
```