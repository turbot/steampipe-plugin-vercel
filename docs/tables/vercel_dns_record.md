---
title: "Steampipe Table: vercel_dns_record - Query Vercel DNS Records using SQL"
description: "Allows users to query DNS Records in Vercel, specifically the details of each record such as ID, name, type, value, and more."
---

# Table: vercel_dns_record - Query Vercel DNS Records using SQL

Vercel DNS Records are configurations that provide information about which IP addresses are associated with a domain name. They are essential for directing web traffic to the correct servers. Vercel provides a platform for DNS management, allowing users to create, update, and delete DNS records.

## Table Usage Guide

The `vercel_dns_record` table provides insights into DNS Records within Vercel. As a network administrator, explore DNS record details through this table, including record types, associated values, and related metadata. Utilize it to uncover information about DNS records, such as their configurations, associated domain names, and the status of each record.

## Examples

### List all DNS records for all domains
Explore which DNS records are associated with your domains. This can be useful in managing and troubleshooting your network, ensuring the correct routing of internet traffic to your domains.

```sql
select
  domain_name,
  name,
  type,
  value,
  ttl
from
  vercel_dns_record;
```

### List all A records for all domains
Explore which domains have been assigned an IPv4 address (A record). This is useful for understanding the distribution of IP addresses across your domains, aiding in network management and troubleshooting.

```sql
select
  domain_name,
  type,
  value,
  ttl
from
  vercel_dns_record
where
  type = 'A';
```