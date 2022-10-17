package bitfinex

import (
	"context"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableBitFinexTicker(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bitfinex_ticker",
		Description: "Ticker Data for any pair from Bitfinex",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("symbol"),
			Hydrate: getSingleTicker,
		},
		List: &plugin.ListConfig{
			Hydrate: listAllTickers,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol").NullIfZero(), Description: "Symbol for which Ticker Data is requested"},

			// Ticker Data
			{Name: "frr", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Frr"), Description: "Flash Return Rate - average of all fixed rate funding over the last hour (funding tickers only)."},
			{Name: "bid", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Bid"), Description: "Price of last highest bid"},
			{Name: "bid_period", Type: proto.ColumnType_INT, Transform: transform.FromField("BidPeriod"), Description: "Bid period covered in days (funding tickers only)"},
			{Name: "bid_size", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("BidSize"), Description: "Sum of the 25 highest bid sizes"},
			{Name: "ask", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Ask"),Description: "Price of last lowest ask"},
			{Name: "ask_period", Type: proto.ColumnType_INT, Transform: transform.FromField("AskPeriod"), Description: "Ask period covered in days (funding tickers only)"},
			{Name: "ask_size", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("AskSize"), Description: "Sum of the 25 lowest ask sizes"},
			{Name: "daily_change", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("DailyChange"), Description: "Amount that the last price has changed since yesterday"},
			{Name: "daily_change_relative", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("DailyChangePerc"), Description: "Relative price change since yesterday (*100 for percentage change)"},
			{Name: "last_price", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("LastPrice"), Description: "Price of the last trade"},
			{Name: "volume", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Volume"), Description: "Daily volume"},
			{Name: "high", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("High"), Description: "Daily high"},
			{Name: "low", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Low"), Description: "Daily low"},
			{Name: "frr_amount_available", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("FrrAmountAvailable"), Description: "The amount of funding that is available at the Flash Return Rate (funding tickers only)"},

			
		},
	}
}



func getSingleTicker(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	symbol := quals["symbol"].GetStringValue()
	
	
	client := rest.NewClient()

	tickerRes, err := client.Tickers.Get(symbol)
	if err != nil {
		plugin.Logger(ctx).Error("bitfinex_ticker.getSingleTicker", "lookup_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("bitfinex_ticker.getSingleTicker", "response",tickerRes)
		
	return tickerRes,nil
}

func listAllTickers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	
	client := rest.NewClient()

	tickerRes, err := client.Tickers.All()
	if err != nil {
		plugin.Logger(ctx).Error("bitfinex_ticker.listAllTickers", "lookup_error", err)
		return nil, err
	}
	plugin.Logger(ctx).Debug("bitfinex_ticker.listAllTickers", "response",tickerRes)

	for _, ticker := range tickerRes {
		d.StreamListItem(ctx, ticker)
	}
	return nil,nil

}


