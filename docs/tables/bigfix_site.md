---
title: "Steampipe Table: bigfix_site - Query BigFix Site using SQL"
description: "Allows users to query BigFix Site data, providing details such as site name, type, display name, permissions, files, and more. This table is useful for site management, security audits, and operational troubleshooting."
folder: "Sites"
---

# Table: bigfix_site - Query BigFix Site using SQL

The BigFix Site represents a collection of content and policies in the BigFix platform. It contains information about the site's identity, type, permissions, files, and configuration. This table provides comprehensive visibility into all sites managed by BigFix.

## Table Usage Guide

The `bigfix_site` table in Steampipe provides you with information about sites managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query site-specific details, including site name, type, display name, permissions, and files. You can utilize this table to gather insights on site configuration, access control, and content management. The schema outlines the various attributes of the BigFix site, including permissions, files, and subscription settings.

## Examples

### Basic site information
Discover the segments that provide details about sites in your BigFix environment, including their types and display names. This can be useful for auditing site configuration and maintaining security compliance.

```sql+postgres
select
  name,
  type,
  display_name,
  description,
  subscription_mode
from
  bigfix_site;
```

```sql+sqlite
select
  name,
  type,
  display_name,
  description,
  subscription_mode
from
  bigfix_site;
```

### Sites by type
Identify sites by their type to understand the distribution of different site types in your environment.

```sql+postgres
select
  type,
  count(*) as site_count
from
  bigfix_site
group by
  type
order by
  site_count desc;
```

```sql+sqlite
select
  type,
  count(*) as site_count
from
  bigfix_site
group by
  type
order by
  site_count desc;
```

### Sites with specific subscription modes
Find sites with specific subscription modes to understand content distribution and access patterns.

```sql+postgres
select
  name,
  type,
  subscription_mode,
  display_name
from
  bigfix_site
where
  subscription_mode = 'All'
order by
  name;
```

```sql+sqlite
select
  name,
  type,
  subscription_mode,
  display_name
from
  bigfix_site
where
  subscription_mode = 'All'
order by
  name;
```

### Sites with permissions
Examine permissions for sites to understand access control and identify any security issues.

```sql+postgres
select
  name,
  type,
  permissions
from
  bigfix_site
where
  permissions is not null;
```

```sql+sqlite
select
  name,
  type,
  permissions
from
  bigfix_site
where
  permissions is not null;
```

### Sites with files
Analyze files in sites to understand content distribution and identify any missing or outdated content.

```sql+postgres
select
  name,
  type,
  files
from
  bigfix_site
where
  files is not null;
```

```sql+sqlite
select
  name,
  type,
  files
from
  bigfix_site
where
  files is not null;
```

### External sites
Focus on external sites to understand external content sources and their configuration.

```sql+postgres
select
  name,
  display_name,
  description,
  subscription_mode,
  gather_url
from
  bigfix_site
where
  type = 'external'
order by
  name;
```

```sql+sqlite
select
  name,
  display_name,
  description,
  subscription_mode,
  gather_url
from
  bigfix_site
where
  type = 'external'
order by
  name;
```

### Operator sites
Examine operator sites to understand operator-specific content and configuration.

```sql+postgres
select
  name,
  display_name,
  description,
  subscription_mode
from
  bigfix_site
where
  type = 'operator'
order by
  name;
```

```sql+sqlite
select
  name,
  display_name,
  description,
  subscription_mode
from
  bigfix_site
where
  type = 'operator'
order by
  name;
```

### Master sites
Identify master sites to understand the main content sources and their configuration.

```sql+postgres
select
  name,
  display_name,
  description,
  subscription_mode,
  gather_url
from
  bigfix_site
where
  type = 'master'
order by
  name;
```

```sql+sqlite
select
  name,
  display_name,
  description,
  subscription_mode,
  gather_url
from
  bigfix_site
where
  type = 'master'
order by
  name;
```

### Sites with global read permissions
Find sites with global read permissions to understand access control and identify potential security risks.

```sql+postgres
select
  name,
  type,
  global_read_permission,
  display_name
from
  bigfix_site
where
  global_read_permission = true
order by
  name;
```

```sql+sqlite
select
  name,
  type,
  global_read_permission,
  display_name
from
  bigfix_site
where
  global_read_permission = true
order by
  name;
```
