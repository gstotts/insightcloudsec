package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////
type Query struct {
	Scopes               []string `json:"scopes"`
	Filters              []string `json:"filters"`
	Offset               int      `json:"offset"`
	Limit                int      `json:"limit"`
	OrderBy              string   `json:"order_by"`
	SelectedResourceType string   `json:"selected_resource_type"`
}

type QueryResult struct {
	Scopes               []string       `json:"scopes"`
	Filters              []QueryFilter  `json:"filters"`
	Offset               int            `json:"offset"`
	OrderBy              string         `json:"order_by"`
	Counts               map[string]int `json:"counts"`
	SelectedResourceType string         `json:"selected_resource_type"`
	Resources            []Resource     `json:"resources"`
	SupportedTypes       []string       `json:"supported_resources"`
}

type QueryFilter struct {
	Config QueryFilterConfig
	Name   string
}

type QueryFilterConfig struct {
}

type Resource struct {
	Dependencies map[string]string `json:"dependencies"`
	Details      map[string]string `json:"details"`
}

// QUERY FUNCTIONS
///////////////////////////////////////////
func (c Client) Query(q *Query) (map[string]int, error) {
	if q.Filters == nil {
		q.Filters = make([]string, 0)
	}
	if q.Scopes == nil {
		q.Scopes = make([]string, 0)
	}

	data, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(http.MethodPost, "/v2/public/resource/query", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var ret *QueryResult
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.Counts, nil
}

// RESOURCE FUNCTIONS
///////////////////////////////////////////
