---
title: "Steampipe Table: vercel_deployment - Query Vercel Deployments using SQL"
description: "Allows users to query Vercel Deployments, specifically deployment details such as deployment ID, status, and URL, providing insights into application deployment patterns and potential issues."
---

# Table: vercel_deployment - Query Vercel Deployments using SQL

A Vercel Deployment is a version of your application that is accessible via a URL. Each deployment is assigned a unique URL that points to a specific version of your application. Vercel Deployments are immutable, meaning once they are created, they cannot be changed or deleted.

## Table Usage Guide

The `vercel_deployment` table provides insights into deployments within the Vercel platform. As a DevOps engineer, explore deployment-specific details through this table, including deployment status, creation time, and associated metadata. Utilize it to uncover information about deployments, such as those that failed, those that are currently active, and the historical record of deployments.

## Examples

### List recent deployments
Explore recent project deployments to gain insights into their status, associated URLs, creators, and relevant commit messages and references. This can be particularly useful for tracking development progress and identifying potential issues in a timely manner.

```sql+postgres
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
  created_at desc;
```

```sql+sqlite
select
  name as project,
  state,
  url,
  json_extract(creator, '$.email') as creator,
  json_extract(meta, '$.githubCommitMessage') as commit_message,
  json_extract(meta, '$.githubCommitRef') as commit_ref
from
  vercel_deployment
where
  created_at > datetime('now', '-14 day')
order by
  created_at desc;
```

### List recent deployments that are in ERROR state
This example allows you to identify recent projects that have encountered errors during deployment. This can be useful for quickly pinpointing problematic deployments and addressing issues in a timely manner.

```sql+postgres
select
  name as project,
  state,
  url,
  creator ->> 'email' as creator,
  meta ->> 'githubCommitMessage' as commit_message
from
  vercel_deployment
where
  created_at > now() - interval '2 weeks'
  and state = 'ERROR'
order by
  created_at desc;
```

```sql+sqlite
select
  name as project,
  state,
  url,
  json_extract(creator, '$.email') as creator,
  json_extract(meta, '$.githubCommitMessage') as commit_message
from
  vercel_deployment
where
  created_at > datetime('now', '-14 day')
  and state = 'ERROR'
order by
  created_at desc;
```