---
title: "Steampipe Table: bigfix_computer - Query BigFix Computer using SQL"
description: "Allows users to query BigFix Computer data, providing details such as computer ID, name, OS, IP address, last report time, and more. This table is useful for endpoint management, security audits, and operational troubleshooting."
folder: "Computers"
---

# Table: bigfix_computer - Query BigFix Computer using SQL

The BigFix Computer represents an endpoint device managed by the BigFix platform. It contains information about the computer's identity, operating system, network configuration, hardware details, and client settings. This table provides comprehensive visibility into all computers managed by BigFix.

## Table Usage Guide

The `bigfix_computer` table in Steampipe provides you with information about computers managed by BigFix. This table allows you, as a DevOps engineer or security analyst, to query computer-specific details, including computer ID, name, operating system, IP address, and last report time. You can utilize this table to gather insights on endpoint health, security posture, and compliance status. The schema outlines the various attributes of the BigFix computer, including hardware information, network details, and client settings.

## Examples

### Basic computer information
Discover the segments that provide details about computers in your BigFix environment, including their operating systems and last report times. This can be useful for auditing endpoint health and maintaining security compliance.

```sql+postgres
select
  name,
  id,
  os,
  ip_address,
  last_report_time
from
  bigfix_computer;
```

```sql+sqlite
select
  name,
  id,
  os,
  ip_address,
  last_report_time
from
  bigfix_computer;
```

### Computers with specific operating systems
Identify computers running specific operating systems to help with patch management and security assessments.

```sql+postgres
select
  name,
  id,
  os,
  os_family,
  os_name,
  os_version
from
  bigfix_computer
where
  os_family = 'Windows';
```

```sql+sqlite
select
  name,
  id,
  os,
  os_family,
  os_name,
  os_version
from
  bigfix_computer
where
  os_family = 'Windows';
```

### Computers that haven't reported recently
Find computers that haven't reported to BigFix recently, which could indicate connectivity issues or security concerns.

```sql+postgres
select
  name,
  id,
  last_report_time,
  os
from
  bigfix_computer
where
  last_report_time < now() - interval '7 days'
order by
  last_report_time;
```

```sql+sqlite
select
  name,
  id,
  last_report_time,
  os
from
  bigfix_computer
where
  last_report_time < datetime('now', '-7 days')
order by
  last_report_time;
```

### Computers with specific hardware configurations
Analyze computers based on their hardware specifications to identify systems that may need upgrades or have specific requirements.

```sql+postgres
select
  name,
  id,
  cpu,
  ram,
  total_size_of_system,
  free_space_on_system
from
  bigfix_computer
where
  ram < 8192
order by
  ram;
```

```sql+sqlite
select
  name,
  id,
  cpu,
  ram,
  total_size_of_system,
  free_space_on_system
from
  bigfix_computer
where
  ram < 8192
order by
  ram;
```

### Network information for computers
Get detailed network information for all computers to understand network topology and identify potential network issues.

```sql+postgres
select
  name,
  id,
  ip_address,
  ipv6_address,
  dns_name,
  mac_address,
  subnet_address
from
  bigfix_computer
where
  ip_address is not null;
```

```sql+sqlite
select
  name,
  id,
  ip_address,
  ipv6_address,
  dns_name,
  mac_address,
  subnet_address
from
  bigfix_computer
where
  ip_address is not null;
```

### Client settings for computers
Examine client settings for computers to understand configuration and identify any misconfigurations.

```sql+postgres
select
  name,
  id,
  client_settings
from
  bigfix_computer
where
  client_settings is not null;
```

```sql+sqlite
select
  name,
  id,
  client_settings
from
  bigfix_computer
where
  client_settings is not null;
```

### Computers by device type
Categorize computers by their device type to understand the distribution of different endpoint types in your environment.

```sql+postgres
select
  device_type,
  count(*) as computer_count
from
  bigfix_computer
group by
  device_type
order by
  computer_count desc;
```

```sql+sqlite
select
  device_type,
  count(*) as computer_count
from
  bigfix_computer
group by
  device_type
order by
  computer_count desc;
```

### Computers with specific license types
Identify computers with specific license types to ensure proper licensing compliance.

```sql+postgres
select
  name,
  id,
  license_type,
  os
from
  bigfix_computer
where
  license_type is not null
order by
  license_type;
```

```sql+sqlite
select
  name,
  id,
  license_type,
  os
from
  bigfix_computer
where
  license_type is not null
order by
  license_type;
```
