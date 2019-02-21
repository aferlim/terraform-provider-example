package provider

import (
	iacitem "github.com/aferlim/terraform-provider-example/client/iac-item"
)

// AllClients provides all provider clients
type AllClients struct {
	ItemsClient iacitem.Client
}
