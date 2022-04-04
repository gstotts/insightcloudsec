package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////

type Badge struct {
	// The key and value of a given badge for use with filters, insights, etc.
	Key            string `json:"key"`
	Value          string `json:"value"`
	Auto_Generated bool   `json:"auto_generated,omitempty"`
}

type Badges struct {
	Badges []Badge `json:"badges"`
}

type Badge_Request struct {
	Org_Resource_IDs []string `json:"target_resource_ids"`
	Badges           []Badge  `json:"badges"`
}

type Badge_Count_Response struct {
	Resource_Count []interface{} `json:"resource_count"`
}

// FUNCTIONS
///////////////////////////////////////////

func (c Client) Create_Badge(target_org_resource_ids []string, badge_data map[string]string) error {
	// Creates a badge for target organization resource ids of key and value pairings provided in map
	data, err := json.Marshal(Badge_Request{
		Org_Resource_IDs: target_org_resource_ids,
		Badges:           createBadgesFromMap(badge_data),
	})
	if err != nil {
		return err
	}

	_, err = c.makeRequest(http.MethodPost, "/v2/public/badges/create", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Update_Cloud_Badges(org_resource_id string, badge_data map[string]string) error {
	// Updates cloud badges for given organization but overwrites any existing. USE WITH CAUTION
	data, err := json.Marshal(createBadgesFromMap(badge_data))
	if err != nil {
		return err
	}

	_, err = c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s/update", org_resource_id), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Delete_Badges(target_org_resource_ids []string, badges map[string]string) error {
	// Deletes given list of badges defined as a map of key/values.
	data := Badge_Request{
		Org_Resource_IDs: target_org_resource_ids,
		Badges:           createBadgesFromMap(badges),
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = c.makeRequest(http.MethodPost, "/v2/public/badges/delete", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	return nil
}

func (c Client) List_Clouds_With_Badges(badges map[string]string) ([]Cloud, error) {
	// Returns a list of cloud accounts what contain the given badges defined as a map of key / values.
	data, err := json.Marshal(badges)
	if err != nil {
		return []Cloud{}, err
	}

	resp, err := c.makeRequest(http.MethodPost, "/v2/public/badge/clouds/list", bytes.NewBuffer(data))
	if err != nil {
		return []Cloud{}, err
	}

	var clouds []Cloud
	if err := json.NewDecoder(resp.Body).Decode(&clouds); err != nil {
		return []Cloud{}, err
	}

	return clouds, nil
}

func (c Client) List_Resource_Badges(org_resource_id string) ([]Badge, error) {
	// Returns a list of resource badges for a given organization
	resp, err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s/list", org_resource_id), nil)
	if err != nil {
		return []Badge{}, err
	}

	var ret []Badge
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return []Badge{}, err
	}

	return ret, nil
}

func (c Client) List_Resources_Badge_Count() (Badge_Count_Response, error) {
	// Returns a list of badge counts for all resources.
	resp, err := c.makeRequest(http.MethodPost, "/v2/public/badges/count", nil)
	if err != nil {
		return Badge_Count_Response{}, err
	}

	var ret Badge_Count_Response
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Badge_Count_Response{}, err
	}

	return ret, nil
}

func createBadgesFromMap(m map[string]string) []Badge {
	badges := []Badge{}
	for badge, value := range m {
		item := Badge{
			Key:   badge,
			Value: value,
		}
		badges = append(badges, item)
	}

	return badges
}
