# Table: bitfinex_currency

This includes the config data and pairs related to particular currency.

## Examples

### To get pairs for a currency

```sql
select
  pairs
from
  bitfinex_currency
where
  currency = 'USD';
```
     
### To get all currency that contains bitcoin in label/description

```sql
select
  currency
from
  bitfinex_currency
where
  label like '%Bitcoin%';
```






