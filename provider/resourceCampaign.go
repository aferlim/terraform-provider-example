package provider

import (
	"fmt"
	"strings"

	"github.com/aferlim/terraform-provider-example/client/campaign"
	iacitem "github.com/aferlim/terraform-provider-example/client/iac-item"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCampaign() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
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
		Name:           d.Get("name").(string),
		ClientID:       d.Get("clientId").(int),
		ExternalPoints: d.Get("externalPoints").(int),
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}

	d.Set("id", item)
	return nil
}

func resourceReadCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CampaignClient

	itemID := d.Id()
	item, err := apiClient.GetItem(d.ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", itemID)
		}
	}

	d.SetId(item.ID)
	d.Set("id", item.ID)
	d.Set("name", item.Name)
	d.Set("clientId", item.ClientID)
	d.Set("externalPoints", item.ExternalPoints)
	return nil
}

func resourceUpdateCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).ItemsClient

	tfTags := d.Get("tags").(*schema.Set).List()
	tags := make([]string, len(tfTags))
	for i, tfTag := range tfTags {
		tags[i] = tfTag.(string)
	}

	item := iacitem.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        tags,
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteCampaign(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).ItemsClient

	itemID := d.Id()

	err := apiClient.DeleteItem(itemID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsCampaign(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*AllClients).ItemsClient

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
