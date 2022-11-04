// Last Reviewed: 2-Apr-2022
// InsightCloudSec Version at time of review: 22.2

package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

var _ Filters = (*filters)(nil)

type Filters interface {
	Get_Registry() (map[string]FilterRegistryResult, error)
}

type filters struct {
	client *Client
}

type FilterRegistryResult struct {
	ID                  string                 `json:"Filterid"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	Supported_Resources []string               `json:"supported_resources"`
	Supports_Common     bool                   `json:"supports_common"`
	Supported_Clouds    []string               `json:"supported_clouds"`
	Settings_Config     []FilterRegistryConfig `json:"settings_config"`
}

type FilterRegistryConfig struct {
	Field_Type   string                         `json:"field_type"`
	Name         string                         `json:"name"`
	Display_Name string                         `json:"display_name"`
	Description  interface{}                    `json:"description"` // Should be string but divvy.query.shared_file_system_lifecycle_policy was returning an array for a boolean field
	Options      []string                       `json:"options"`
	Choices      []FilterRegistryConfig_Choices `json:"choices"`
	State_Hash   string                         `json:"_state_hash"`
}

type FilterRegistryConfig_Choices struct {
	Value         string `json:"value"`
	Display_Value string `json:"display_value"`
}

// FUNCTIONS
///////////////////////////////////////////

func (f *filters) Get_Registry() (map[string]FilterRegistryResult, error) {
	// Returns a map of filter registry results
	body, err := f.client.makeRequest(http.MethodGet, "/v2/public/insights/filter-registry", nil, nil)
	if err != nil {
		return nil, err
	}

	var fr map[string]FilterRegistryResult
	if err := json.Unmarshal(body, &fr); err != nil {
		return nil, err
	}

	return fr, nil
}
