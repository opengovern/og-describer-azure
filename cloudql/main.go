package main

import (
	"github.com/opengovern/og-describer-azure/cloudql/azure"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: azure.Plugin})
}
