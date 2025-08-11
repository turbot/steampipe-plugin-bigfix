package main

import (
	"github.com/turbot/steampipe-plugin-bigfix/bigfix"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: bigfix.Plugin})
}
