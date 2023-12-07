---
title: "Steampipe Table: vercel_user - Query Vercel User using SQL"
description: "Allows users to query User data in Vercel, providing insights into user account details, including email, username, name, and more."
---

# Table: vercel_user - Query Vercel User using SQL

Vercel User is a fundamental entity within the Vercel platform that represents an individual user account. Each user has a unique username and email, along with additional personal details. Vercel User is the primary entity for authentication and authorization within the Vercel platform.

## Table Usage Guide

The `vercel_user` table provides insights into individual user accounts within the Vercel platform. As a DevOps engineer or a security analyst, explore user-specific details through this table, including email, username, and other personal details. Utilize it to uncover information about user accounts, such as their creation date, bio, and the verification of their email.

## Examples

### Get user information
Explore the user profiles within your Vercel account to better manage and understand your user base. This can be particularly useful for auditing user access and permissions within your account.

```sql+postgres
select
  *
from
  vercel_user;
```

```sql+sqlite
select
  *
from
  vercel_user;
```