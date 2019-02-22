package campaign

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Client is the rest Api from Items
type Client struct {
	BaseClient
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, token string) *Client {

	return &Client{
		NewBase(
			hostname,
			token,
		)}
}

// GetAll Retrieves all of the Items from the server
func (c *Client) GetAll() (*[]Campaign, error) {

	body, err := c.httpRequest("campaign", "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	items := []Campaign{}
	err = json.NewDecoder(body).Decode(&items)

	if err != nil {
		return nil, err
	}

	return &items, nil
}

// GetItem gets an item with a specific name from the server
func (c *Client) GetItem(id string) (*Campaign, error) {

	body, err := c.httpRequest(fmt.Sprintf("campaign/%v", id), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	item := &Campaign{}
	err = json.NewDecoder(body).Decode(item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

// NewItem creates a new Item
func (c *Client) NewItem(item *Campaign) error {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)

	if err != nil {
		return err
	}

	_, err = c.httpRequest("campaign", "POST", buf)

	if err != nil {
		return err
	}

	return nil
}

// UpdateItem updates the values of an item
func (c *Client) UpdateItem(item *Campaign) error {
	buf := bytes.Buffer{}

	err := json.NewEncoder(&buf).Encode(item)

	if err != nil {
		return err
	}
	_, err = c.httpRequest("campaign", "PUT", buf)

	if err != nil {
		return err
	}
	return nil
}

// DeleteItem removes an item from the server
func (c *Client) DeleteItem(id string) error {

	_, err := c.httpRequest(fmt.Sprintf("campaign/%v", id), "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}
