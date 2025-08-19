---
title: "Steampipe Table: bigfix_analysis - Query BigFix Analysis using SQL"
description: "Allows users to query BigFix Analysis data, providing details such as analysis ID, name, title, description, relevance, and more. This table is useful for content management, security audits, and operational troubleshooting."
folder: "Analyses"
---

# Table: bigfix_analysis - Query BigFix Analysis using SQL

The BigFix Analysis represents a content item in the BigFix platform that defines relevance expressions and properties. It contains information about the analysis's identity, content, relevance, and metadata. This table provides comprehensive visibility into all analyses managed by BigFix.

## Table Usage Guide

The `bigfix_analysis` table in Steampipe provides you with information about analyses managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query analysis-specific details, including analysis ID, name, title, description, relevance, and category. You can utilize this table to gather insights on content relevance, security policies, and compliance requirements. The schema outlines the various attributes of the BigFix analysis, including relevance expressions, properties, and metadata.

## Examples

### Basic analysis information
Discover the segments that provide details about analyses in your BigFix environment, including their titles and descriptions. This can be useful for auditing content and maintaining security compliance.

```sql+postgres
select
  name,
  id,
  title,
  description,
  site_name,
  site_type
from
  bigfix_analysis;
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
  bigfix_analysis;
```

### Analyses by category
Identify analyses by their category to understand the distribution of different content types in your environment.

```sql+postgres
select
  category,
  count(*) as analysis_count
from
  bigfix_analysis
group by
  category
order by
  analysis_count desc;
```

```sql+sqlite
select
  category,
  count(*) as analysis_count
from
  bigfix_analysis
group by
  category
order by
  analysis_count desc;
```

### Analyses with specific relevance
Find analyses with specific relevance expressions to understand content targeting and identify potential issues.

```sql+postgres
select
  name,
  title,
  relevance,
  site_name,
  site_type
from
  bigfix_analysis
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
  bigfix_analysis
where
  relevance is not null;
```

### Analyses by source
Analyze analyses by their source to understand content origins and identify any missing or outdated content.

```sql+postgres
select
  source,
  count(*) as analysis_count
from
  bigfix_analysis
group by
  source
order by
  analysis_count desc;
```

```sql+sqlite
select
  source,
  count(*) as analysis_count
from
  bigfix_analysis
group by
  source
order by
  analysis_count desc;
```

### Analyses in external sites
Focus on analyses in external sites to understand external content sources and their configuration.

```sql+postgres
select
  name,
  title,
  description,
  source,
  source_release_date
from
  bigfix_analysis
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
  bigfix_analysis
where
  site_type = 'external'
order by
  name;
```

### Analyses with properties
Examine properties for analyses to understand configuration and identify any misconfigurations.

```sql+postgres
select
  name,
  title,
  properties
from
  bigfix_analysis
where
  properties is not null;
```

```sql+sqlite
select
  name,
  title,
  properties
from
  bigfix_analysis
where
  properties is not null;
```

### Recent analyses
Find recently created or modified analyses to understand current content and identify any new security policies.

```sql+postgres
select
  name,
  title,
  last_modified,
  source_release_date
from
  bigfix_analysis
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
  bigfix_analysis
order by
  last_modified desc
limit 10;
```

### Analyses with MIME fields
Identify analyses with MIME fields to understand content formatting and identify any issues.

```sql+postgres
select
  name,
  title,
  mime_fields
from
  bigfix_analysis
where
  mime_fields is not null;
```

```sql+sqlite
select
  name,
  title,
  mime_fields
from
  bigfix_analysis
where
  mime_fields is not null;
```
