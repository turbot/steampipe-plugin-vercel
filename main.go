package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-vercel/vercel"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: vercel.Plugin})
}
