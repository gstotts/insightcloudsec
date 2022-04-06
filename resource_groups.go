package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var _ ResourceGroups = (*rsgroup)(nil)

type ResourceGroups interface {
	Create(name string, description string) (ResourceGroup, error)
	AddToGroup(resource_ids []string, group_name string) error
	Delete(resource_ids []string) error
}

type rsgroup struct {
	client *Client
}

type ResourceGroupConifg struct {
	Name        string `json:"group_name"`
	Description string `json:"group_description"`
	OwnerType   string `json:"group_owner_type"`
}

type ResourceGroup struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	ResourceGroupID      string   `json:"resource_group_id"`
	Description          string   `json:"description"`
	Category             string   `json:"category"`
	CreationTime         string   `json:"creation_time"`
	GroupType            string   `json:"group_type"`
	OwnerType            string   `json:"owner_type"`
	NestedResourceGroups []string `json:"nested_resource_groups"`
}

type ResourceGroupIDsList struct {
	ResourceIDs []string `json:"resource_ids"`
}

type ResourcesToGroup struct {
	ResourceIDs      []string `json:"resource_ids"`
	ResourceGroupIDs []string `json:"resource_group_ids"`
}

func (r *rsgroup) Create(name string, description string) (ResourceGroup, error) {
	config := ResourceGroupConifg{
		Name:        name,
		Description: description,
		OwnerType:   "organization",
	}

	data, err := json.Marshal(config)
	if err != nil {
		return ResourceGroup{}, err
	}

	resp, err := r.client.makeRequest(http.MethodPost, "/v2/public/resourcegroup/create", bytes.NewBuffer(data))
	if err != nil {
		return ResourceGroup{}, err
	}

	var ret ResourceGroup
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return ResourceGroup{}, err
	}

	return ret, nil
}

func (r *rsgroup) AddToGroup(resource_ids []string, group_name string) error {
	config := ResourcesToGroup{
		ResourceGroupIDs: []string{group_name},
		ResourceIDs:      resource_ids,
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	_, err = r.client.makeRequest(http.MethodPost, "/v2/prototype/resourcegroups/add", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return err
}

func (r *rsgroup) Delete(resource_ids []string) error {
	config := ResourceGroupIDsList{
		ResourceIDs: resource_ids,
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	_, err = r.client.makeRequest(http.MethodPost, "/v2/prototype/resources/delete", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}
