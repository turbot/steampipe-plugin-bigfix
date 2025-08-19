---
title: "Steampipe Table: bigfix_fixlet - Query BigFix Fixlets using SQL"
description: "Allows users to query BigFix Fixlet data, providing details such as fixlet ID, name, title, description, relevance, actions, and more. This table is useful for fixlet management, security audits, and operational troubleshooting."
folder: "Fixlets"
---

# Table: bigfix_fixlet - Query BigFix Fixlets using SQL

The BigFix Fixlet represents a content item in the BigFix platform that defines security patches, software updates, or configuration changes. It contains information about the fixlet's identity, content, relevance, associated actions, and deployment details. This table provides comprehensive visibility into all fixlets managed by BigFix.

## Table Usage Guide

The `bigfix_fixlet` table in Steampipe provides you with information about fixlets managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query fixlet-specific details, including fixlet ID, name, title, description, relevance, and associated actions. You can utilize this table to gather insights on patch management, security policies, and compliance requirements. The schema outlines the various attributes of the BigFix fixlet, including relevance expressions, actions, and deployment settings.

## Examples

### Basic fixlet information
Discover the segments that provide details about fixlets in your BigFix environment, including their titles and descriptions. This can be useful for auditing patch management and maintaining security compliance.

```sql+postgres
select
  name,
  id,
  title,
  description,
  site_name,
  site_type
from
  bigfix_fixlet;
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
  bigfix_fixlet;
```

### Fixlets by category
Identify fixlets by their category to understand the distribution of different patch types in your environment.

```sql+postgres
select
  category,
  count(*) as fixlet_count
from
  bigfix_fixlet
group by
  category
order by
  fixlet_count desc;
```

```sql+sqlite
select
  category,
  count(*) as fixlet_count
from
  bigfix_fixlet
group by
  category
order by
  fixlet_count desc;
```

### Fixlets with high severity
Find fixlets with high severity to understand critical security patches and identify potential risks.

```sql+postgres
select
  name,
  title,
  source_severity,
  site_name,
  site_type
from
  bigfix_fixlet
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
  bigfix_fixlet
where
  source_severity = 'Critical'
order by
  site_name;
```

### Fixlets with specific relevance
Find fixlets with specific relevance expressions to understand patch targeting and identify potential issues.

```sql+postgres
select
  name,
  title,
  relevance,
  site_name,
  site_type
from
  bigfix_fixlet
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
  bigfix_fixlet
where
  relevance is not null;
```

### Fixlets with actions
Analyze fixlets with their associated actions to understand patch deployment and identify any issues.

```sql+postgres
select
  name,
  title,
  actions,
  site_name,
  site_type
from
  bigfix_fixlet
where
  actions is not null;
```

```sql+sqlite
select
  name,
  title,
  actions,
  site_name,
  site_type
from
  bigfix_fixlet
where
  actions is not null;
```

### Fixlets by source
Analyze fixlets by their source to understand patch origins and identify any missing or outdated content.

```sql+postgres
select
  source,
  count(*) as fixlet_count
from
  bigfix_fixlet
group by
  source
order by
  fixlet_count desc;
```

```sql+sqlite
select
  source,
  count(*) as fixlet_count
from
  bigfix_fixlet
group by
  source
order by
  fixlet_count desc;
```

### Fixlets with large download sizes
Identify fixlets with large download sizes to understand resource requirements and identify potential performance issues.

```sql+postgres
select
  name,
  title,
  download_size,
  site_name,
  site_type
from
  bigfix_fixlet
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
  bigfix_fixlet
where
  download_size > 1000000
order by
  download_size desc;
```

### Recent fixlets
Find recently created or modified fixlets to understand current patch management and identify any new security policies.

```sql+postgres
select
  name,
  title,
  last_modified,
  source_release_date
from
  bigfix_fixlet
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
  bigfix_fixlet
order by
  last_modified desc
limit 10;
```

### Fixlets with MIME fields
Identify fixlets with MIME fields to understand content formatting and identify any issues.

```sql+postgres
select
  name,
  title,
  mime_fields
from
  bigfix_fixlet
where
  mime_fields is not null;
```

```sql+sqlite
select
  name,
  title,
  mime_fields
from
  bigfix_fixlet
where
  mime_fields is not null;
```

### Fixlets in external sites
Focus on fixlets in external sites to understand external patch sources and their configuration.

```sql+postgres
select
  name,
  title,
  description,
  source,
  source_release_date
from
  bigfix_fixlet
where
  site_type = 'external'
order by
  name;
```

```sql+sqlite
select
  name,
  title,
  description,
  source,
  source_release_date
from
  bigfix_fixlet
where
  site_type = 'external'
order by
  name;
```
