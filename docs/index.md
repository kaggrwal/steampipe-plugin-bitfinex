---
organization: kaggrwal
category: ["internet"]
icon_url: "/images/plugins/kaggrwal/bitfinex.svg"
brand_color: "#16B157"
display_name: Bitfinex
short_name: bitfinex
description: Steampipe plugin for querying data from bitfinex
og_description: Query Bitfinex with SQL! Open source CLI. No DB required. 
og_image: "/images/plugins/kaggrwal/bitfinex-social-graphic.png"
---

# BITFINEX + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[BITFINEX](https://www.bitfinex.com/about) a cryptocurrency exchange owned and operated by iFinex Inc.

For example: To get the Ticker Data

```sql
select
  bid,
  ask
from
  bitfinex_ticker
where
  symbol = 'tLTCUSD';
```

```
+--------+--------+
| bid    | ask    |
+--------+--------+
| 51.851 | 51.884 |
+--------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/kaggrwal/bitfinex/docs/tables)**

## Get started

### Install

Download and install the latest BITFINEX plugin:

```bash
steampipe plugin install kaggrwal/bitfinex
```

### Configuration

No configuration is needed. Installing the latest bitfinex plugin will create a config file (`~/.steampipe/config/bitfinex.spc`) with a single connection named `bitfinex`:

```hcl
connection "bitfinex" {
  plugin = "kaggrwal/bitfinex"
}
```

## Get involved

* Open source: https://github.com/kaggrwal/steampipe-plugin-bitfinex
* Community: [Slack Channel](https://steampipe.io/community/join)
