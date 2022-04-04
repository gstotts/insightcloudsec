package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Badge_Creation_Request struct {
	Resource_IDs []string `json:"target_resource_ids"`
	Badges       []Badge  `json:"badges"`
}

type Badge struct {
	// The key and value of a given badge for use with filters, insights, etc.
	Key   string `json:"key"`
	Value string `json:"value"`
}

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
