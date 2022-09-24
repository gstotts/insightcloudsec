package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var _ Organizations = (*orgs)(nil)

type Organizations interface {
	Create(name string) error
	Edit_Name(resource_id int, name string) error
	Delete(resource_id int) error
	List() ([]organizations, error)
	Switch(name string) error
}

type orgs struct {
	client *Client
}

type organization struct {
	Name string `json:"organization_name"`
}

type organizations struct {
	Name string `json:"name"`
	ID   int    `json:"organization_id"`
}

func (c *orgs) Create(name string) error {
	// Creates an organization for the given name
	resp, err := c.client.makeRequest(http.MethodPost, "/v2/prototype/domain/organization/create", organization{Name: name})
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *orgs) Switch(name string) error {
	// Switches organization to the given organization name
	resp, err := c.client.makeRequest(http.MethodPost, "/v2/prototype/domain/switch_organization", organization{Name: name})
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *orgs) Edit_Name(resource_id int, name string) error {
	// Renames an organization of given resource_id to the given name
	resp, err := c.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/domain/organization/divvyorganization:%d/update", resource_id), organization{Name: name})
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *orgs) Delete(resource_id int) error {
	// Deletes an organization of given resource_id
	resp, err := c.client.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/domain/organization/divvyorganization:%d/delete", resource_id), nil)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *orgs) List() ([]organizations, error) {
	// Lists all organizations
	resp, err := c.client.makeRequest(http.MethodGet, "/v2/prototype/domain/organizations/get", nil)
	if err != nil {
		return []organizations{}, err
	}

	var ret []organizations
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return []organizations{}, err
	}
	return ret, nil
}
