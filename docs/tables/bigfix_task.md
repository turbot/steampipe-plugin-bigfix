---
title: "Steampipe Table: bigfix_task - Query BigFix Task using SQL"
description: "Allows users to query BigFix Task data, providing details such as task ID, name, title, description, relevance, computers, and more. This table is useful for task management, security audits, and operational troubleshooting."
folder: "Task Management"
---

# Table: bigfix_task - Query BigFix Task using SQL

The BigFix Task represents an executable content item in the BigFix platform that defines actions to be performed on endpoints. It contains information about the task's identity, content, relevance, target computers, and execution details. This table provides comprehensive visibility into all tasks managed by BigFix.

## Table Usage Guide

The `bigfix_task` table in Steampipe provides you with information about tasks managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query task-specific details, including task ID, name, title, description, relevance, and target computers. You can utilize this table to gather insights on task execution, security policies, and compliance requirements. The schema outlines the various attributes of the BigFix task, including relevance expressions, default actions, and computer targeting.

## Examples

### Basic task information
Discover the segments that provide details about tasks in your BigFix environment, including their titles and descriptions. This can be useful for auditing task execution and maintaining security compliance.

```sql+postgres
select
  name,
  id,
  title,
  description,
  site_name,
  site_type
from
  bigfix_task;
```

```sql+sqlite
select
  name,
  id,
  title,
  description,
  site_name,
  site_type
from
  bigfix_task;
```

### Tasks by category
Identify tasks by their category to understand the distribution of different task types in your environment.

```sql+postgres
select
  category,
  count(*) as task_count
from
  bigfix_task
group by
  category
order by
  task_count desc;
```

```sql+sqlite
select
  category,
  count(*) as task_count
from
  bigfix_task
group by
  category
order by
  task_count desc;
```

### Tasks with high severity
Find tasks with high severity to understand critical security policies and identify potential risks.

```sql+postgres
select
  name,
  title,
  source_severity,
  site_name,
  site_type
from
  bigfix_task
where
  source_severity = 'Critical'
order by
  site_name;
```

```sql+sqlite
select
  name,
  title,
  source_severity,
  site_name,
  site_type
from
  bigfix_task
where
  source_severity = 'Critical'
order by
  site_name;
```

### Tasks with specific relevance
Find tasks with specific relevance expressions to understand task targeting and identify potential issues.

```sql+postgres
select
  name,
  title,
  relevance,
  site_name,
  site_type
from
  bigfix_task
where
  relevance is not null;
```

```sql+sqlite
select
  name,
  title,
  relevance,
  site_name,
  site_type
from
  bigfix_task
where
  relevance is not null;
```

### Tasks with computers
Analyze tasks with their target computers to understand task scope and identify any targeting issues.

```sql+postgres
select
  name,
  title,
  computers,
  site_name,
  site_type
from
  bigfix_task
where
  computers is not null;
```

```sql+sqlite
select
  name,
  title,
  computers,
  site_name,
  site_type
from
  bigfix_task
where
  computers is not null;
```

### Tasks by source
Analyze tasks by their source to understand task origins and identify any missing or outdated content.

```sql+postgres
select
  source,
  count(*) as task_count
from
  bigfix_task
group by
  source
order by
  task_count desc;
```

```sql+sqlite
select
  source,
  count(*) as task_count
from
  bigfix_task
group by
  source
order by
  task_count desc;
```

### Tasks with default actions
Examine default actions for tasks to understand task behavior and identify any misconfigurations.

```sql+postgres
select
  name,
  title,
  default_action
from
  bigfix_task
where
  default_action is not null;
```

```sql+sqlite
select
  name,
  title,
  default_action
from
  bigfix_task
where
  default_action is not null;
```

### Recent tasks
Find recently created or modified tasks to understand current task execution and identify any new security policies.

```sql+postgres
select
  name,
  title,
  last_modified,
  source_release_date
from
  bigfix_task
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
  bigfix_task
order by
  last_modified desc
limit 10;
```

### Tasks with large download sizes
Identify tasks with large download sizes to understand resource requirements and identify potential performance issues.

```sql+postgres
select
  name,
  title,
  download_size,
  site_name,
  site_type
from
  bigfix_task
where
  download_size > 1000000
order by
  download_size desc;
```

```sql+sqlite
select
  name,
  title,
  download_size,
  site_name,
  site_type
from
  bigfix_task
where
  download_size > 1000000
order by
  download_size desc;
```
