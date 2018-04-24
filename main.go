package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/victorsalaun/terraform-provider-hesperides/hesperides"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: hesperides.Provider})
}
