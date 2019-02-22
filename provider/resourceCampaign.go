package provider

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/aferlim/terraform-provider-example/client/campaign"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCampaign() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the resource, also acts as it's unique ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"clientId": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "A description of an item",
			},
			"externalPoints": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
		},
		Create: resourceCreateCampaign,
		Read:   resourceReadCampaign,
		Update: resourceUpdateCampaign,
		Delete: resourceDeleteCampaign,
		Exists: resourceExistsCampaign,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CampaignClient

	item := campaign.Campaign{
		ID:             string(rand.Intn(1000)),
		Name:           d.Get("name").(string),
		ClientID:       d.Get("clientId").(int),
		ExternalPoints: d.Get("externalPoints").(int),
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}

	d.SetId(item.ID)
	return nil
}

func resourceReadCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CampaignClient

	itemID := d.Id()
	item, err := apiClient.GetItem(itemID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", itemID)
		}
	}

	d.SetId(itemID)
	d.Set("code", itemID)
	d.Set("name", item.Name)
	d.Set("clientId", item.ClientID)
	d.Set("externalPoints", item.ExternalPoints)
	return nil
}

func resourceUpdateCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CampaignClient

	itemID := d.Id()

	item := campaign.Campaign{
		ID:             itemID,
		Name:           d.Get("name").(string),
		ClientID:       d.Get("clientId").(int),
		ExternalPoints: d.Get("externalPoints").(int),
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CampaignClient

	itemID := d.Id()

	err := apiClient.DeleteItem(itemID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsCampaign(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*AllClients).CampaignClient

	itemID := d.Id()
	_, err := apiClient.GetItem(itemID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err

	}
	return true, nil
}
