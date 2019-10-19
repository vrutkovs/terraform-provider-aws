package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vrutkovs/terraform-provider-aws/v3/aws"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aws.Provider})
}
