package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/kaggrwal/steampipe-plugin-bitfinex/bitfinex"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: bitfinex.Plugin})
}
