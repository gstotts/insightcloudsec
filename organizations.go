package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////
type Organization struct {
	Name string `json:"organization_name"`
}

type OrganizationList struct {
	Orgs []Organizations
}

type Organizations struct {
	Name string `json:"name"`
	ID   int    `json:"organization_id"`
}

// ORGANIZATION CLIENT FUNCTIONS
///////////////////////////////////////////
func (c Client) CreateOrganization(name string) error {
	data, err := json.Marshal(Organization{Name: name})
	if err != nil {
		return fmt.Errorf("[-] error marshalling organization")
	}
	resp, err := c.makeRequest(http.MethodPost, "/v2/prototype/domain/organization/create", bytes.NewBuffer(data))
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c Client) SwitchOrganization(name string) error {
	data, err := json.Marshal(Organization{Name: name})
	if err != nil {
		return fmt.Errorf("[-] error marshalling organization")
	}
	resp, err := c.makeRequest(http.MethodPost, "/v2/prototype/domain/switch_organization", bytes.NewBuffer(data))
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c Client) EditOrganizationName(resource_id int, name string) error {
	data, err := json.Marshal(Organization{Name: name})
	if err != nil {
		return fmt.Errorf("[-] error marshalling organization")
	}
	resp, err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/domain/organization/%d/update", resource_id), bytes.NewBuffer(data))
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c Client) DeleteOrganization(resource_id int) error {
	resp, err := c.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/domain/organization/%d/delete", resource_id), nil)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c Client) ListOrganizations() (OrganizationList, error) {
	resp, err := c.makeRequest(http.MethodGet, "/v2/prototype/domain/organizations/get", nil)
	if err != nil {
		return OrganizationList{}, err
	}

	var ret OrganizationList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return OrganizationList{}, err
	}
	return ret, nil
}
