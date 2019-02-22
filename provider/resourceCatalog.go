package provider

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/aferlim/terraform-provider-example/client/catalog"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCatalog() *schema.Resource {
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
			"project_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "A description of an item",
			},
			"conversion_rate": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
			"external_payment": {
				Type:        schema.TypeBool,
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
		ID:             fmt.Sprintf("%v", rand.Intn(1000)),
		Name:           d.Get("name").(string),
		ProjectID:      d.Get("project_id").(int),
		ConversionRate: d.Get("conversion_rate").(int),
		ExternalPaymet: d.Get("external_payment").(bool),
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}

	d.SetId(item.ID)
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

	d.SetId(itemID)
	d.Set("code", itemID)
	d.Set("name", item.Name)
	d.Set("project_id", item.ProjectID)
	d.Set("conversion_rate", item.ConversionRate)
	d.Set("external_payment", item.ExternalPaymet)
	return nil
}

func resourceUpdateCatalog(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).CatalogClient

	item := catalog.Catalog{
		ID:             d.Id(),
		Name:           d.Get("name").(string),
		ProjectID:      d.Get("project_id").(int),
		ConversionRate: d.Get("conversion_rate").(int),
		ExternalPaymet: d.Get("external_payment").(bool),
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
