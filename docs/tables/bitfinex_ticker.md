# Table: bitfinex_ticker

This includes the ticker data related to particular currency and pairs 

Note: It uses the format "t[Symbol]" (i.e. tBTCUSD, tETHUSD, tBTCUST, ...) for currency pairs and "f[Symbol]" (e.g. fUSD, fBTC, fETH, ...) for funding currencies. To get the the configs/data related to currencies and pairs refer [table_bitfinex_currency](/plugins/kaggrwal/bitfinex/docs/tables/tables_bitfinex_currency.md)


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


## Implementation notes- Table Structure

+-----------------------+------------------+--------------------------------------------------------------------------------------------------+
| column                | type             | description                                                                                      |
+-----------------------+------------------+--------------------------------------------------------------------------------------------------+
| _ctx                  | jsonb            | Steampipe context in JSON form, e.g. connection_name.                                            |
| ask                   | double precision | Price of last lowest ask                                                                         |
| ask_period            | bigint           | Ask period covered in days (funding tickers only)                                                |
| ask_size              | double precision | Sum of the 25 lowest ask sizes                                                                   |
| bid                   | double precision | Price of last highest bid                                                                        |
| bid_period            | bigint           | Bid period covered in days (funding tickers only)                                                |
| bid_size              | double precision | Sum of the 25 highest bid sizes                                                                  |
| daily_change          | double precision | Amount that the last price has changed since yesterday                                           |
| daily_change_relative | double precision | Relative price change since yesterday (*100 for percentage change)                               |
| frr                   | double precision | Flash Return Rate - average of all fixed rate funding over the last hour (funding tickers only). |
| frr_amount_available  | double precision | The amount of funding that is available at the Flash Return Rate (funding tickers only)          |
| high                  | double precision | Daily high                                                                                       |
| last_price            | double precision | Price of the last trade                                                                          |
| low                   | double precision | Daily low                                                                                        |
| symbol                | text             | Symbol for which Ticker Data is requested                                                        |
| volume                | double precision | Daily volume                                                                                     |
+-----------------------+------------------+--------------------------------------------------------------------------------------------------+


