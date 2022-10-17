package bitfinex

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-bitfinex",
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"bitfinex_ticker": tableBitFinexTicker(ctx),
			"bitfinex_currency": tableBitFinexCurrency(ctx),
		},
	}
	return p
}
