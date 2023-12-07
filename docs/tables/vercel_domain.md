---
title: "Steampipe Table: vercel_domain - Query Vercel Domains using SQL"
description: "Allows users to query Vercel Domains, specifically providing details about the domain name, verification status, and associated records."
---

# Table: vercel_domain - Query Vercel Domains using SQL

Vercel Domains is a feature within Vercel that allows you to manage and configure your custom domains for your Vercel deployments. It provides a centralized way to set up and manage domains, including DNS records, SSL certificates, and more. Vercel Domains helps you maintain the health and performance of your custom domains and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `vercel_domain` table provides insights into the custom domains within Vercel. As a DevOps engineer, explore domain-specific details through this table, including verification status, associated DNS records, and SSL certificates. Utilize it to uncover information about domains, such as those with invalid SSL certificates, the verification status of domains, and the configuration of DNS records.

## Examples

### List all domains
Discover the segments that consist of all domains, including their verification status and the dates they were created and will expire. This can help you manage and track your domains effectively.

```sql+postgres
select
  name,
  verified,
  created_at,
  expires_at
from
  vercel_domain;
```

```sql+sqlite
select
  name,
  verified,
  created_at,
  expires_at
from
  vercel_domain;
```

### Domains expiring in the next 90 days
Determine the domains that are due to expire within the next 90 days. This query is useful for proactive management of domains, ensuring they are renewed on time to prevent service disruptions.
Lists domains expiring soon. Domains managed outside of Vercel will not be included in results.


```sql+postgres
select
  name,
  expires_at
from
  vercel_domain
where
  expires_at < now() + interval '90 days';
```

```sql+sqlite
select
  name,
  expires_at
from
  vercel_domain
where
  expires_at < datetime('now', '+90 days');
```