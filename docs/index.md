---
organization: Turbot
category: ["security", "endpoint management"]
icon_url: "/images/plugins/turbot/bigfix.svg"
brand_color: "#FF6B35"
display_name: "BigFix"
name: "bigfix"
description: "Steampipe plugin for querying computers, sites, analyses, tasks, actions, fixlets, properties, and roles from BigFix."
og_description: "Query BigFix with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/bigfix-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# BigFix + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[BigFix](https://www.hcltech.com/software/bigfix) is an endpoint management and security platform that provides real-time visibility and control over endpoints across your organization.

For example:

```sql
select
  name,
  id,
  os,
  last_report_time
from
  bigfix_computer
```

```
+-----------------+----+------------------+---------------------+
| name            | id | os               | last_report_time    |
+-----------------+----+------------------+---------------------+
| DESKTOP-ABC123  | 1  | Windows 10 Pro  | 2024-01-15 10:30:00 |
| LAPTOP-XYZ789   | 2  | macOS 13.2.1    | 2024-01-15 09:45:00 |
| SERVER-DEF456   | 3  | Ubuntu 22.04    | 2024-01-15 11:15:00 |
+-----------------+----+------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/bigfix/tables)**

## Get started

### Install

Download and install the latest BigFix plugin:

```bash
steampipe plugin install bigfix
```

### Credentials

| Item        | Description                                                                                         |
| ----------- | --------------------------------------------------------------------------------------------------- |
| Credentials | BigFix server credentials (username/password) are required for authentication.                      |
| Permissions | The user must have appropriate permissions to access BigFix API endpoints.                          |
| Radius      | Each connection represents a single BigFix server instance.                                         |
| Resolution  | Credentials must be explicitly set in the Steampipe config file (`~/.steampipe/config/bigfix.spc`). |

### Configuration

Installing the latest bigfix plugin will create a config file (`~/.steampipe/config/bigfix.spc`) with a single connection named `bigfix`:

```hcl
connection "bigfix" {
  plugin = "bigfix"

  # `server_name` defines the BigFix server hostname or IP address.
  # This is required for connecting to the BigFix server.
  #server_name = "bigfix.example.com"

  # `port` defines the port number for the BigFix server.
  # Defaults to 52311 if not specified.
  #port = 52311

  # `username` defines the username for BigFix authentication.
  # This is required for connecting to the BigFix server.
  #username = "admin"

  # `password` defines the password for BigFix authentication.
  # This is required for connecting to the BigFix server.
  #password = "your_password"

  # The maximum number of attempts (including the initial call) Steampipe will
  # make for failing API calls. Defaults to 3 and must be greater than or equal to 1.
  #max_retries = 3

  # The minimum retry delay in milliseconds after which retries will be performed.
  # This delay is also used as a base value when calculating the exponential backoff retry times.
  # Defaults to 100ms and must be greater than or equal to 1ms.
  #min_retry_delay = 100

  # List of additional BigFix error messages to ignore for all queries.
  # When encountering these errors, the API call will not be retried and empty results will be returned.
  # By default, "not found" errors are ignored and will still be ignored even if this argument is not set.
  #ignore_error_messages = ["Access Denied", "Unauthorized", "Invalid credentials"]

  # Whether to skip TLS certificate verification when connecting to the BigFix server.
  # This should only be used in development or testing environments with self-signed certificates.
  # Defaults to false for security.
  #insecure_skip_verify = false

  # The request timeout in seconds for API requests to the BigFix server.
  # This is useful for environments with slow network connections or large datasets.
  # Defaults to 120 seconds.
  #request_timeout = 120
}
```
