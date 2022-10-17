package bitfinex

import (
	"context"
	"strings"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)


func tableBitFinexCurrency(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "bitfinex_currency",
		Description: "Configs related to currency and pairs",
		List: &plugin.ListConfig{
			Hydrate: listAllCurrencies,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "currency", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency").NullIfZero(), Description: "Currency for which Config is requested"},

			// Currency Data
			{Name: "label", Type: proto.ColumnType_STRING, Transform: transform.FromField("Label").NullIfZero(), Description: "Verbose friendly name for the currency"},
			{Name: "symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Symbol").NullIfZero(), Description: "API symbol for the currency"},
			{Name: "unit", Type: proto.ColumnType_STRING, Transform: transform.FromField("Unit").NullIfZero(), Description: "Unit of measure for the currency"},
			{Name: "pairs", Type: proto.ColumnType_STRING, Transform: transform.FromField("Pairs").Transform(sliceToCommaSepratedValues).NullIfZero(), Description: "All pairs for the currency"},
			{Name: "pools", Type: proto.ColumnType_STRING, Transform: transform.FromField("Pools").Transform(sliceToCommaSepratedValues).NullIfZero(),Description: "Network/Protocol which is used by currency"},
		    {Name: "explorers_base_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Explorers.BaseUri").NullIfZero(), Description: "Recognised block explorer Base URL for the currency"},
			{Name: "explorers_address_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Explorers.AddressUri").NullIfZero(), Description: "Recognised block explorer Address URL for the currency"},
		    {Name: "explorers_transaction_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Explorers.TransactionUri").NullIfZero(), Description: "Recognised block explorer Transaction URL for the currency"},

		
			
		},
	}
}


func listAllCurrencies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	
	client := rest.NewClient()

	currenciesRes, err := client.Currencies.Conf(true, true, true, true, true)
	if err != nil {
		plugin.Logger(ctx).Error("bitfinex_currency.listAllCurrencies", "lookup_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("bitfinex_currency.listAllCurrencies", "response",currenciesRes)

	for _, currency := range currenciesRes {
		d.StreamListItem(ctx, currency)
	}
	return nil,nil

}

func sliceToCommaSepratedValues(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Debug("sliceToCommaSepratedValues",d.Value)
	pairs := d.Value.([] string)

	if len(pairs) > 0{

		pairCommaSeparatedString := strings.Join(pairs[:],",")
		return pairCommaSeparatedString,nil

	} 
	return nil,nil;
}

