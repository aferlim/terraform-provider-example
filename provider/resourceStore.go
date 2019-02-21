package provider

import (
	"fmt"
	"strings"

	"github.com/aferlim/terraform-provider-example/client/store"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceStore() *schema.Resource {
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
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"VendorId": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "A description of an item",
			},
			"projectConfigurationId": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
			"visible": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
			"parameters": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "An optional list of tags, represented as a key, value pair",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
		Create: resourceCreateStore,
		Read:   resourceReadStore,
		Update: resourceUpdateStore,
		Delete: resourceDeleteStore,
		Exists: resourceExistsStore,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateStore(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).StoreClient

	tags := []string{"parameter 1", "parameter 2", "parameter 3"}

	item := store.Store{
		Name:                   d.Get("name").(string),
		Description:            d.Get("description").(string),
		VendorID:               d.Get("vendorId").(int),
		ProjectConfigurationID: d.Get("projectConfigurationId").(int),
		Visible:                d.Get("visible").(int),
		Parameters:             tags,
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}

	d.Set("id", item)
	return nil
}

func resourceReadStore(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).StoreClient

	itemID := d.Id()
	item, err := apiClient.GetItem(itemID)
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
	d.Set("description", item.Description)
	d.Set("vendorId", item.VendorID)
	d.Set("projectConfigurationId", item.ProjectConfigurationID)
	d.Set("visible", item.Visible)
	d.Set("parameters", item.Parameters)
	return nil
}

func resourceUpdateStore(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).StoreClient

	tags := []string{"parameter 1", "parameter 2", "parameter 3"}

	item := store.Store{
		Name:                   d.Get("name").(string),
		Description:            d.Get("description").(string),
		VendorID:               d.Get("vendorId").(int),
		ProjectConfigurationID: d.Get("projectConfigurationId").(int),
		Visible:                d.Get("visible").(int),
		Parameters:             tags,
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteStore(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).StoreClient

	itemID := d.Id()

	err := apiClient.DeleteItem(itemID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsStore(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*AllClients).StoreClient

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
