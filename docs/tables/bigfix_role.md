---
title: "Steampipe Table: bigfix_role - Query BigFix Role using SQL"
description: "Allows users to query BigFix Role data, providing details such as role ID, name, permissions, interface logins, and more. This table is useful for role management, access control audits, and security troubleshooting."
folder: "Roles"
---

# Table: bigfix_role - Query BigFix Role using SQL

The BigFix Role represents a user role in the BigFix platform that defines permissions and access rights. It contains information about the role's identity, privileges, interface access, and security settings. This table provides comprehensive visibility into all roles managed by BigFix.

## Table Usage Guide

The `bigfix_role` table in Steampipe provides you with information about roles managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query role-specific details, including role ID, name, permissions, and interface access. You can utilize this table to gather insights on access control, security policies, and compliance requirements. The schema outlines the various attributes of the BigFix role, including privileges, interface logins, and security settings.

## Examples

### Basic role information
Discover the segments that provide details about roles in your BigFix environment, including their names and permissions. This can be useful for auditing access control and maintaining security compliance.

```sql+postgres
select
  name,
  id,
  master_operator,
  can_create_actions,
  interface_logins
from
  bigfix_role;
```

```sql+sqlite
select
  name,
  id,
  master_operator,
  can_create_actions,
  interface_logins
from
  bigfix_role;
```

### Roles by master operator status
Identify roles by their master operator status to understand the distribution of administrative roles in your environment.

```sql+postgres
select
  master_operator,
  count(*) as role_count
from
  bigfix_role
group by
  master_operator
order by
  role_count desc;
```

```sql+sqlite
select
  master_operator,
  count(*) as role_count
from
  bigfix_role
group by
  master_operator
order by
  role_count desc;
```

### Master operator roles
Find master operator roles to understand administrative privileges and identify potential security risks.

```sql+postgres
select
  name,
  id,
  can_create_actions,
  can_submit_queries,
  interface_logins
from
  bigfix_role
where
  master_operator = 1
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  can_create_actions,
  can_submit_queries,
  interface_logins
from
  bigfix_role
where
  master_operator = 1
order by
  name;
```

### Roles with action creation privileges
Identify roles with action creation privileges to understand who can create executable content and identify potential security concerns.

```sql+postgres
select
  name,
  id,
  can_create_actions,
  show_other_actions,
  stop_other_actions
from
  bigfix_role
where
  can_create_actions = 1
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  can_create_actions,
  show_other_actions,
  stop_other_actions
from
  bigfix_role
where
  can_create_actions = 1
order by
  name;
```

### Roles with API access
Find roles with API access to understand programmatic access privileges and identify potential security risks.

```sql+postgres
select
  name,
  id,
  interface_logins
from
  bigfix_role
where
  json_extract(interface_logins, '$.api') = true
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  interface_logins
from
  bigfix_role
where
  json_extract(interface_logins, '$.api') = true
order by
  name;
```

### Roles with console access
Identify roles with console access to understand administrative interface privileges and identify potential security concerns.

```sql+postgres
select
  name,
  id,
  interface_logins
from
  bigfix_role
where
  json_extract(interface_logins, '$.console') = true
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  interface_logins
from
  bigfix_role
where
  json_extract(interface_logins, '$.console') = true
order by
  name;
```

### Roles with query submission privileges
Find roles with query submission privileges to understand who can submit queries and identify potential security risks.

```sql+postgres
select
  name,
  id,
  can_submit_queries,
  can_lock,
  interface_logins
from
  bigfix_role
where
  can_submit_queries = 1
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  can_submit_queries,
  can_lock,
  interface_logins
from
  bigfix_role
where
  can_submit_queries = 1
order by
  name;
```

### Roles with custom content access
Identify roles with custom content access to understand content management privileges and identify potential security concerns.

```sql+postgres
select
  name,
  id,
  custom_content,
  interface_logins
from
  bigfix_role
where
  custom_content = 1
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  custom_content,
  interface_logins
from
  bigfix_role
where
  custom_content = 1
order by
  name;
```

### Recent role modifications
Find recently modified roles to understand current access control changes and identify any new security policies.

```sql+postgres
select
  name,
  id,
  last_modified,
  master_operator
from
  bigfix_role
where
  last_modified is not null
order by
  last_modified desc
limit 10;
```

```sql+sqlite
select
  name,
  id,
  last_modified,
  master_operator
from
  bigfix_role
where
  last_modified is not null
order by
  last_modified desc
limit 10;
```

### Roles with specific privileges
Search for roles with specific privileges to understand permission patterns and identify potential security issues.

```sql+postgres
select
  name,
  id,
  post_action_behavior_privilege,
  action_script_commands_privilege,
  unmanaged_asset_privilege
from
  bigfix_role
where
  post_action_behavior_privilege is not null
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  post_action_behavior_privilege,
  action_script_commands_privilege,
  unmanaged_asset_privilege
from
  bigfix_role
where
  post_action_behavior_privilege is not null
order by
  name;
```
