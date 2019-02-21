package provider

import (
	"fmt"
	"strings"

	"github.com/aferlim/terraform-provider-example/client/catalog"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCatalog() *schema.Resource {
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
			"projectId": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "A description of an item",
			},
			"conversionRate": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
		},
		Create: resourceCreateCatalog,
		Read:   resourceReadCatalog,
		Update: resourceUpdateCatalog,
		Delete: resourceDeleteCatalog,
		Exists: resourceExistsCatalog,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateCatalog(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CatalogClient

	item := catalog.Catalog{
		Name:           d.Get("name").(string),
		ProjectID:      d.Get("projectId").(int),
		ConversionRate: d.Get("conversionRate").(int),
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}

	d.Set("id", item)
	return nil
}

func resourceReadCatalog(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CatalogClient

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
	d.Set("projectId", item.ProjectID)
	d.Set("conversionRate", item.ConversionRate)
	return nil
}

func resourceUpdateCatalog(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CatalogClient

	item := catalog.Catalog{
		Name:           d.Get("name").(string),
		ProjectID:      d.Get("projectId").(int),
		ConversionRate: d.Get("conversionRate").(int),
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteCatalog(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CatalogClient

	itemID := d.Id()

	err := apiClient.DeleteItem(itemID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsCatalog(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*AllClients).CatalogClient

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
