## v1.1.0 [2026-11-13]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v6.2.1](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v621-2026-11-12), which lowers plugin memory use, with the query cache fixes for missing or duplicate rows from [v6.2.0](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v620-2026-11-11). ([#30](https://github.com/turbot/steampipe-plugin-bigfix/pull/30))

## v1.0.1 [2025-08-19]

_Bug fixes_

- Fixed the plugin's brand color.

## v1.0.0 [2025-08-19]

_What's new?_

- New tables added
  - [bigfix_action](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_action)
  - [bigfix_analysis](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_analysis)
  - [bigfix_computer](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_computer)
  - [bigfix_fixlet](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_fixlet)
  - [bigfix_property](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_property)
  - [bigfix_role](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_role)
  - [bigfix_site](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_site)
  - [bigfix_task](https://hub.steampipe.io/plugins/turbot/bigfix/tables/bigfix_task)