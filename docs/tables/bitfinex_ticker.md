# Table: bitfinex_ticker

This includes the ticker data related to particular currency and pairs. 

Note: It uses the format "t[Symbol]" (i.e. tBTCUSD, tETHUSD, tBTCUST, ...) for currency pairs and "f[Symbol]" (e.g. fUSD, fBTC, fETH, ...) for funding currencies. To get the the configs/data related to currencies and pairs refer [table_bitfinex_currency](/plugins/kaggrwal/bitfinex/docs/tables/tables_bitfinex_currency.md).


## Examples

### To get whole ticker data for a symbol

```sql
select
  *
from
  bitfinex_ticker
where
  symbol = 'tLTCUSD';
```

### To get Volume for particular pair

```sql
select
  volume
from
  bitfinex_ticker
where
  symbol = 'fUSD';
```


### To get the change in price since yesterday for a particular pair

```sql
select
  daily_change
from
  bitfinex_ticker
where
  symbol = 'fUSD';
```


### To get all pairs whose ask is between 2 and 10

```sql
select
  *
from
  bitfinex_ticker
where
  ask > 2 AND ask < 10;
```


### To get highest price with the pair for the particular day

```sql
select
  high,symbol
from
  bitfinex_ticker
where
  order by high desc limit 1;
```


