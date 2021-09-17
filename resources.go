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
	Instance InstanceResource `json:"instance"`
	Type     string           `json:"resource_type"`
}

type InstanceResource struct {
	Type               string               `json:"instance_type"`
	ImageID            string               `json:"image_id"`
	KeyName            string               `json:"key_name"`
	LaunchTime         string               `json:"launch_time"`
	ID                 string               `json:"instance_id"`
	Platform           string               `json:"platform"`
	Tenancy            string               `json:"tenancy"`
	DetailedMonitoring bool                 `json:"detailed_monitoring"`
	Common             CommonResourceValues `json:"common"`
}

type CommonResourceValues struct {
	//Common Attributes
	Account          string `json:"account"`
	Name             string `json:"resource_name"`
	OrgServiceID     int    `json:"organization_service_id"`
	AvailabilityZone string `json:"availablility_zone"`
	Region           string `json:"region"`
	ID               string `json:"resource_id"`
	Cloud            string `json:"cloud"`
	Type             string `json:"type"`
}

// QUERY FUNCTIONS
///////////////////////////////////////////
func (c Client) Query(q *Query) (*QueryResult, error) {
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

	return ret, nil
}

// RESOURCE FUNCTIONS
///////////////////////////////////////////
