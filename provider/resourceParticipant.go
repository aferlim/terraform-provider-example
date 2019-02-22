package provider

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/aferlim/terraform-provider-example/client/participant"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceParticipant() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the resource, also acts as it's unique ID",
			},
			"login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the resource, also acts as it's unique ID",
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"pasword": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"customer_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "A description of an item",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "A description of an item",
			},
			"project_configuration_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
			"active": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "An optional list of tags, represented as a key, value pair",
			},
		},
		Create: resourceCreateParticipant,
		Read:   resourceReadParticipant,
		Update: resourceUpdateParticipant,
		Delete: resourceDeleteParticipant,
		Exists: resourceExistsParticipant,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateParticipant(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).ParticipantClient

	item := participant.Participant{
		ID:                     fmt.Sprintf("%v", rand.Intn(1000)),
		Name:                   d.Get("name").(string),
		Login:                  d.Get("login").(string),
		Email:                  d.Get("email").(string),
		Password:               d.Get("password").(string),
		ProjectID:              d.Get("project_id").(int),
		CustomerID:             d.Get("customer_id").(int),
		ProjectConfigurationID: d.Get("project_configuration_id").(int),
		Active:                 d.Get("active").(bool),
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return err
	}

	d.SetId(item.ID)
	return nil
}

func resourceReadParticipant(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).ParticipantClient

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
	d.Set("name").(string)
	d.Set("login").(string)
	d.Set("email").(string)
	d.Set("password").(string)
	d.Set("project_id").(int)
	d.Set("customer_id").(int)
	d.Set("project_configuration_id").(int)
	d.Set("active").(bool)
	return nil
}

func resourceUpdateParticipant(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).ParticipantClient

	item := participant.Participant{
		ID:                     d.Id(),
		Name:                   d.Get("name").(string),
		Login:                  d.Get("login").(string),
		Email:                  d.Get("email").(string),
		Password:               d.Get("password").(string),
		ProjectID:              d.Get("project_id").(int),
		CustomerID:             d.Get("customer_id").(int),
		ProjectConfigurationID: d.Get("project_configuration_id").(int),
		Active:                 d.Get("active").(bool),
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteParticipant(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*AllClients).ParticipantClient

	itemID := d.Id()

	err := apiClient.DeleteItem(itemID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsParticipant(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*AllClients).ParticipantClient

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
