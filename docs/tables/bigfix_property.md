---
title: "Steampipe Table: bigfix_property - Query BigFix Property using SQL"
description: "Allows users to query BigFix Property data, providing details such as property ID, name, definition, and more. This table is useful for property management, configuration audits, and operational troubleshooting."
folder: "Property Management"
---

# Table: bigfix_property - Query BigFix Property using SQL

The BigFix Property represents a configuration item in the BigFix platform that defines custom properties and their values. It contains information about the property's identity, definition, and usage. This table provides comprehensive visibility into all properties managed by BigFix.

## Table Usage Guide

The `bigfix_property` table in Steampipe provides you with information about properties managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query property-specific details, including property ID, name, definition, and reservation status. You can utilize this table to gather insights on configuration management, custom properties, and system settings. The schema outlines the various attributes of the BigFix property, including definitions, reservation status, and modification timestamps.

## Examples

### Basic property information

Discover the segments that provide details about properties in your BigFix environment, including their names and definitions. This can be useful for auditing configuration and maintaining system compliance.

```sql+postgres
select
  name,
  id,
  definition,
  is_reserved
from
  bigfix_property;
```

```sql+sqlite
select
  name,
  id,
  definition,
  is_reserved
from
  bigfix_property;
```

### Properties by reservation status

Identify properties by their reservation status to understand the distribution of reserved vs. custom properties in your environment.

```sql+postgres
select
  is_reserved,
  count(*) as property_count
from
  bigfix_property
group by
  is_reserved
order by
  property_count desc;
```

```sql+sqlite
select
  is_reserved,
  count(*) as property_count
from
  bigfix_property
group by
  is_reserved
order by
  property_count desc;
```

### Reserved properties

Find reserved properties to understand system-defined properties and identify any customizations.

```sql+postgres
select
  name,
  id,
  definition,
  last_modified
from
  bigfix_property
where
  is_reserved = 1
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  definition,
  last_modified
from
  bigfix_property
where
  is_reserved = 1
order by
  name;
```

### Custom properties

Identify custom properties to understand user-defined configurations and identify any potential issues.

```sql+postgres
select
  name,
  id,
  definition,
  last_modified
from
  bigfix_property
where
  is_reserved = 0
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  definition,
  last_modified
from
  bigfix_property
where
  is_reserved = 0
order by
  name;
```

### Properties with specific definitions

Find properties with specific definitions to understand property usage and identify any misconfigurations.

```sql+postgres
select
  name,
  id,
  definition
from
  bigfix_property
where
  definition is not null
  and definition != '';
```

```sql+sqlite
select
  name,
  id,
  definition
from
  bigfix_property
where
  definition is not null
  and definition != '';
```

### Recent property modifications

Find recently modified properties to understand current configuration changes and identify any new customizations.

```sql+postgres
select
  name,
  id,
  last_modified,
  is_reserved
from
  bigfix_property
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
  is_reserved
from
  bigfix_property
where
  last_modified is not null
order by
  last_modified desc
limit 10;
```

### Properties by name pattern

Search for properties by name pattern to understand property naming conventions and identify related configurations.

```sql+postgres
select
  name,
  id,
  definition,
  is_reserved
from
  bigfix_property
where
  name like '%config%'
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  definition,
  is_reserved
from
  bigfix_property
where
  name like '%config%'
order by
  name;
```

### Properties with long definitions

Identify properties with long definitions to understand complex configurations and identify any potential issues.

```sql+postgres
select
  name,
  id,
  length(definition) as definition_length,
  is_reserved
from
  bigfix_property
where
  definition is not null
order by
  definition_length desc
limit 10;
```

```sql+sqlite
select
  name,
  id,
  length(definition) as definition_length,
  is_reserved
from
  bigfix_property
where
  definition is not null
order by
  definition_length desc
limit 10;
```

### Properties without definitions

Find properties without definitions to identify potential configuration gaps and incomplete setups.

```sql+postgres
select
  name,
  id,
  is_reserved,
  last_modified
from
  bigfix_property
where
  definition is null
  or definition = ''
order by
  name;
```

```sql+sqlite
select
  name,
  id,
  is_reserved,
  last_modified
from
  bigfix_property
where
  definition is null
  or definition = ''
order by
  name;
```
