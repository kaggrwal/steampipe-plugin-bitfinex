# Table: bitfinex_currency

This includes the config data and pairs related to particular currency

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

```
pairs   
1INCH:USD,AAVE:USD,ADAUSD,AIXUSD,ALBT:USD,ALGUSD,AMPUSD,ANCUSD,ANTUSD,APENFT:USD,APEUSD,ATLAS:USD,ATOUSD,AVAX:USD,AXSUSD,B2MUSD,BALUSD,BAND:USD,BATUSD,BCHABC:USD,BCHN:USD,BEST:USD,BFTUSD,BMIUSD,BMNUSD,BNTUSD,BOBA:USD,BOOUSD,BOSON:USD,BSVUSD,BTCUSD,BTGUSD,BTSE:USD,BTTUSD,CCDUSD,CELUSD,CHEX:USD,CHSB:USD,CHZUSD,CLOUSD,COMP:USD,CONV:USD,CRVUSD,DAIUSD,DGBUSD,DOGE:USD,DORA:USD,DOTUSD,DSHUSD,DUSK:USD,DVFUSD,EDOUSD,EGLD:USD,ENJUSD,EOSUSD,ETCUSD,ETH2X:USD,ETHUSD,ETHW:USD,ETPUSD,EUSUSD,EUTUSD,EXOUSD,FBTUSD,FCLUSD,FETUSD,FILUSD,FORTH:USD,FTMUSD,FTTUSD,FUNUSD,GALA:USD,GNOUSD,GNTUSD,GRTUSD,GTXUSD,GXTUSD,HECUSD,HIXUSD,HMTUSD,ICEUSD,ICPUSD,IDXUSD,IOTUSD,IQXUSD,JASMY:USD,JSTUSD,KAIUSD,KANUSD,KNCUSD,KSMUSD,LEOUSD,LINK:USD,LRCUSD,LTCUSD,LUNA2:USD,LUNA:USD,LUXO:USD,LYMUSD,MATIC:USD,MIMUSD,MIRUSD,MKRUSD,MLNUSD,MNAUSD,MOBUSD,MXNT:USD,NEAR:USD,NEOUSD,NEXO:USD,OCEAN:USD,OMGUSD,OMNUSD,OXYUSD,PASUSD,PAXUSD,PLANETS:USD,PLUUSD,PNGUSD,PNKUSD,POLC:USD,POLIS:USD,QRDO:USD,QTFUSD,QTMUSD,RBTUSD,REEF:USD,REPUSD,REQUSD,RLYUSD,ROSE:USD,RRTUSD,SAND:USD,SENATE:USD,SGBUSD,SHFT:USD,SHIB:USD,SIDUS:USD,SMRUSD,SNTUSD,SNXUSD,SOLUSD,SPELL:USD,SRMUSD,STGUSD,STJUSD,SUKU:USD,SUNUSD,SUSHI:USD,SWEAT:USD,SXXUSD,TERRAUST:USD,THETA:USD,TLOS:USD,TRADE:USD,TREEB:USD,TRXUSD,TSDUSD,UDCUSD,UNIUSD,UOSUSD,USTUSD,UTKUSD,VEEUSD,VELO:USD,VETUSD,VRAUSD,VSYUSD,WAVES:USD,WAXUSD,WBTUSD,WILD:USD,WNCG:USD,WOOUSD,XAUT:USD,XCNUSD,XDCUSD,XLMUSD,XMRUSD,XRAUSD,XRDUSD,XRPUSD,XTZUSD,XVGUSD,YFIUSD,ZCNUSD,ZECUSD,ZILUSD,ZMTUSD,ZRXUSD   
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

+----------+
| currency |
+----------+
| BSV      |
| WBT      |
| BTC      |
| LNX      |
| BTG      |
| LBT      |
| BCH      |
| BCHN     |
+----------+

## Implementation notes- Table Structure

+---------------------------+-------+------------------------------------------------------------+
| column                    | type  | description                                                |
+---------------------------+-------+------------------------------------------------------------+
| _ctx                      | jsonb | Steampipe context in JSON form, e.g. connection_name.      |
| currency                  | text  | Currency for which Config is requested                     |
| explorers_address_url     | text  | Recognised block explorer Address URL for the currency     |
| explorers_base_url        | text  | Recognised block explorer Base URL for the currency        |
| explorers_transaction_url | text  | Recognised block explorer Transaction URL for the currency |
| label                     | text  | Verbose friendly name for the currency                     |
| pairs                     | text  | All pairs for the currency                                 |
| pools                     | text  | Network/Protocol which is used by currency                 |
| symbol                    | text  | API symbol for the currency                                |
| unit                      | text  | Unit of measure for the currency                           |
+---------------------------+-------+------------------------------------------------------------+



