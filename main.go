package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/pluralsh/steampipe-plugin-aws/aws"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: aws.Plugin})
}
