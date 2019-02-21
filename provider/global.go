package provider

import (
	"github.com/aferlim/terraform-provider-example/client/campaign"
	"github.com/aferlim/terraform-provider-example/client/catalog"
	iacitem "github.com/aferlim/terraform-provider-example/client/iac-item"
	"github.com/aferlim/terraform-provider-example/client/store"
)

// AllClients provides all provider clients
type AllClients struct {
	ItemsClient    *iacitem.Client
	CampaignClient *campaign.Client
	CatalogClient  *catalog.Client
	StoreClient    *store.Client
}

// InstanceClients isntance all the clients
func InstanceClients(address string, token string) (interface{}, error) {
	return &AllClients{
		iacitem.NewClient(address, token),
		campaign.NewClient(address, token),
		catalog.NewClient(address, token),
		store.NewClient(address, token),
	}, nil
}
