![image]()

# Bitfinex Plugin for Steampipe

Use SQL to query data from Bitfinex


## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install bitfinex
```

Run a query:
```sql
select volume from bitfinex_ticker where symbol = 'fUSD';
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

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

