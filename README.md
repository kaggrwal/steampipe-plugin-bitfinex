![image]()

# Bitfinex Plugin for Steampipe

Use SQL to query data from Bitfinex


- **[Get started →](https://hub.steampipe.io/plugins/kaggrwal/bitfinex)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/kaggrwal/bitfinex/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/kaggrwal/steampipe-plugin-bitfinex/issues)


## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install kaggrwal/bitfinex
```


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

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/kaggrwal/steampipe-plugin-bitfinex.git
cd steampipe-plugin-bitfinex
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
```

Try it!
```
steampipe query
> .inspect bitfinex_ticker
```

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/kaggrwal/steampipe-plugin-bitfinex/blob/main/LICENSE).

 Further reading:

* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

