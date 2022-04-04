package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////

type Badge_Creation_Request struct {
	Resource_IDs []string `json:"target_resource_ids"`
	Badges       []Badge  `json:"badges"`
}

type Badge struct {
	// The key and value of a given badge for use with filters, insights, etc.
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Badges struct {
	Badges []Badge `json:"badges"`
}

// FUNCTIONS
///////////////////////////////////////////

func (c Client) Create_Badge(target_resource_ids []string, badge_data []Badge) error {
	// Creates a badge for target resource ids of key and value pairings provided in map
	data, err := json.Marshal(Badge_Creation_Request{
		Resource_IDs: target_resource_ids,
		Badges:       badge_data,
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

func (c Client) Update_Cloud_Badges(org_resource_id string, badge_data Badges) error {
	// Updates cloud badges for given organization but overwrites any existing. USE WITH CAUTION
	data, err := json.Marshal(badge_data)
	if err != nil {
		return err
	}

	_, err = c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s/update", org_resource_id), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Delete_Badges(badges Badges) error {
	// Deletes given list of badges
	data, err := json.Marshal(badges)
	if err != nil {
		return err
	}

	_, err = c.makeRequest(http.MethodPost, "/v2/public/badges/delete", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}
