---
title: "Steampipe Table: bigfix_action - Query BigFix Action using SQL"
description: "Allows users to query BigFix Action data, providing details such as action ID, name, title, description, relevance, and more. This table is useful for action management, security audits, and operational troubleshooting."
folder: "Actions"
---

# Table: bigfix_action - Query BigFix Action using SQL

The BigFix Action represents an executable item in the BigFix platform that defines specific operations to be performed on endpoints. It contains information about the action's identity, content, relevance, and execution details. This table provides comprehensive visibility into all actions managed by BigFix.

## Table Usage Guide

The `bigfix_action` table in Steampipe provides you with information about actions managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query action-specific details, including action ID, name, title, description, relevance, and execution parameters. You can utilize this table to gather insights on action execution, security policies, and compliance requirements. The schema outlines the various attributes of the BigFix action, including relevance expressions, script commands, and execution settings.

## Examples

### Basic action information
Discover the segments that provide details about actions in your BigFix environment, including their titles and descriptions. This can be useful for auditing action execution and maintaining security compliance.

```sql+postgres
select
  name,
  id,
  title,
  description,
  relevance
from
  bigfix_action;
```

```sql+sqlite
select
  name,
  id,
  title,
  description,
  relevance
from
  bigfix_action;
```

### Actions by category
Identify actions by their category to understand the distribution of different action types in your environment.

```sql+postgres
select
  category,
  count(*) as action_count
from
  bigfix_action
group by
  category
order by
  action_count desc;
```

```sql+sqlite
select
  category,
  count(*) as action_count
from
  bigfix_action
group by
  category
order by
  action_count desc;
```

### Actions with specific relevance
Find actions with specific relevance expressions to understand action targeting and identify potential issues.

```sql+postgres
select
  name,
  title,
  relevance,
  category
from
  bigfix_action
where
  relevance is not null;
```

```sql+sqlite
select
  name,
  title,
  relevance,
  category
from
  bigfix_action
where
  relevance is not null;
```

### Actions with script commands
Analyze actions with script commands to understand action behavior and identify any security concerns.

```sql+postgres
select
  name,
  title,
  script_commands
from
  bigfix_action
where
  script_commands is not null;
```

```sql+sqlite
select
  name,
  title,
  script_commands
from
  bigfix_action
where
  script_commands is not null;
```

### Actions by source
Analyze actions by their source to understand action origins and identify any missing or outdated content.

```sql+postgres
select
  source,
  count(*) as action_count
from
  bigfix_action
group by
  source
order by
  action_count desc;
```

```sql+sqlite
select
  source,
  count(*) as action_count
from
  bigfix_action
group by
  source
order by
  action_count desc;
```

### Actions with post-action behavior
Examine post-action behavior for actions to understand action completion and identify any misconfigurations.

```sql+postgres
select
  name,
  title,
  post_action_behavior
from
  bigfix_action
where
  post_action_behavior is not null;
```

```sql+sqlite
select
  name,
  title,
  post_action_behavior
from
  bigfix_action
where
  post_action_behavior is not null;
```

### Recent actions
Find recently created or modified actions to understand current action execution and identify any new security policies.

```sql+postgres
select
  name,
  title,
  last_modified,
  source_release_date
from
  bigfix_action
order by
  last_modified desc
limit 10;
```

```sql+sqlite
select
  name,
  title,
  last_modified,
  source_release_date
from
  bigfix_action
order by
  last_modified desc
limit 10;
```

### Actions with MIME fields
Identify actions with MIME fields to understand content formatting and identify any issues.

```sql+postgres
select
  name,
  title,
  mime_fields
from
  bigfix_action
where
  mime_fields is not null;
```

```sql+sqlite
select
  name,
  title,
  mime_fields
from
  bigfix_action
where
  mime_fields is not null;
```

### Actions with specific targets
Find actions with specific targets to understand action scope and identify any targeting issues.

```sql+postgres
select
  name,
  title,
  target
from
  bigfix_action
where
  target is not null;
```

```sql+sqlite
select
  name,
  title,
  target
from
  bigfix_action
where
  target is not null;
```
