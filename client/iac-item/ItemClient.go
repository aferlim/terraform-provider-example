package iacitem

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ItemClient is the rest Api from Items
type ItemClient struct {
	BaseClient
}

// GetAll Retrieves all of the Items from the server
func (c *ItemClient) GetAll() (*[]Item, error) {

	body, err := c.httpRequest("items", "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	items := []Item{}
	err = json.NewDecoder(body).Decode(&items)

	if err != nil {
		return nil, err
	}

	return &items, nil
}

// GetItem gets an item with a specific name from the server
func (c *ItemClient) GetItem(name string) (*Item, error) {

	body, err := c.httpRequest(fmt.Sprintf("items/%v", name), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	item := &Item{}
	err = json.NewDecoder(body).Decode(item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

// NewItem creates a new Item
func (c *ItemClient) NewItem(item *Item) error {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)

	if err != nil {
		return err
	}

	_, err = c.httpRequest("items", "POST", buf)

	if err != nil {
		return err
	}

	return nil
}

// UpdateItem updates the values of an item
func (c *ItemClient) UpdateItem(item *Item) error {
	buf := bytes.Buffer{}

	err := json.NewEncoder(&buf).Encode(item)

	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("items/%s", item.Name), "PUT", buf)

	if err != nil {
		return err
	}
	return nil
}

// DeleteItem removes an item from the server
func (c *ItemClient) DeleteItem(itemName string) error {

	_, err := c.httpRequest(fmt.Sprintf("items/%s", itemName), "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}
