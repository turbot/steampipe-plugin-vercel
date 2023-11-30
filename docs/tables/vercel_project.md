---
title: "Steampipe Table: vercel_project - Query Vercel Projects using SQL"
description: "Allows users to query Vercel Projects, specifically the details of each project such as name, id, type, and owner. This can provide insights into project configurations and ownership."
---

# Table: vercel_project - Query Vercel Projects using SQL

A Vercel Project is a workspace where you can deploy your applications or websites. Each project is linked to a Git repository and contains settings for deployments, domains, environment variables, and more. Projects can be owned by an individual or a team, and they provide the basis for continuous deployment in Vercel.

## Table Usage Guide

The `vercel_project` table provides insights into projects within Vercel. As a DevOps engineer, explore project-specific details through this table, including project names, types, and owners. Utilize it to uncover information about projects, such as their configurations, associated Git repositories, and their continuous deployment settings.

## Examples

### List all projects
Explore the various projects, including their respective frameworks and last updated dates, to keep track of the latest changes and developments. This can help in understanding the current state of each project and facilitate strategic planning.

```sql
select
  name,
  framework,
  updated_at
from
  vercel_project;
```

### Projects not updated in the last year
Identify projects that have not seen any updates in the past year. This can be useful to determine which projects may be inactive or outdated.

```sql
select
  name,
  framework,
  updated_at
from
  vercel_project
where
  updated_at < now() - interval '1 year';
```

### Latest deployments
Gain insights into the most recent project deployments, including when they were created, by whom, and their associated URLs. This is particularly useful for tracking project updates and ensuring accountability within your team.

```sql
select
  to_timestamp((d ->> 'createdAt')::bigint / 1000) as created_at,
  name,
  d ->> 'url' as url,
  d -> 'creator' ->> 'username' as creator_username
from
  vercel_project as p,
  jsonb_array_elements(latest_deployments) as d
order by
  created_at desc;
```

### Current production target by project
This query is useful for gaining insights into the production targets of different projects. It arranges them in order of their names, providing a clear view of the production status, including the GitHub repository details, which can aid in project management and progress tracking.

```sql
select
  name,
  to_timestamp((targets -> 'production' ->> 'createdAt')::bigint / 1000) as created_at,
  targets -> 'production' ->> 'url' as url,
  targets -> 'production' -> 'meta' ->> 'githubOrg' as github_org,
  targets -> 'production' -> 'meta' ->> 'githubRepo' as github_repo,
  targets -> 'production' -> 'meta' ->> 'githubCommitSha' as github_commit_sha,
  targets -> 'production' -> 'meta' ->> 'githubCommitAuthorName' as github_commit_author_name,
  targets -> 'production' -> 'meta' ->> 'githubCommitMessage' as github_commit_message
from
  vercel_project
order by
  name;
```

### List all project environment variables
Explore the environmental variables associated with each project to understand how configurations are set and managed. This is useful for auditing and maintaining consistency across various projects.

```sql
select
  name,
  e ->> 'key',
  e ->> 'target'
from
  vercel_project as p,
  jsonb_array_elements(env) as e;
```

### Environment variables that are not encrypted
Gain insights into the environment variables within your Vercel project that are not encrypted. This query is useful for identifying potential security risks in your project's configuration.

```sql
select
  name,
  e ->> 'key',
  e ->> 'type'
from
  vercel_project as p,
  jsonb_array_elements(env) as e
where
  e ->> 'type' != 'encrypted';
```

### Production environment variables older than 90 days
Explore which production environment variables have not been updated in the last 90 days. This can help identify potential areas of neglect or outdated configurations in your project.

```sql
select
  name,
  e ->> 'key',
  e ->> 'target',
  to_timestamp((e ->> 'createdAt')::bigint / 1000)
from
  vercel_project as p,
  jsonb_array_elements(env) as e
where
  e -> 'target' ? 'production'
  and to_timestamp((e ->> 'createdAt')::bigint / 1000) < now() - interval '90 days';
```

### Projects by framework environment variables older than 90 days
Explore which frameworks are most commonly used in your Vercel projects. This can help you understand the popularity and usage of different frameworks within your projects.

```sql
select
  framework,
  count(id)
from
  vercel_project
group by
  framework
order by
  count desc;
```