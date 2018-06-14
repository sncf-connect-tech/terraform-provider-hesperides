package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/voyages-sncf-technologies/terraform-provider-hesperides/hesperides"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: hesperides.Provider})
}
