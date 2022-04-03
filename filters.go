// Last Reviewed: 2-Apr-2022
// InsightCloudSec Version at time of review: 22.2

package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////

type Filter_Registry_Result struct {
	ID                  string                   `json:"filter_id"`
	Name                string                   `json:"name"`
	Description         string                   `json:"description"`
	Supported_Resources []string                 `json:"supported_resources"`
	Supports_Common     bool                     `json:"supports_common"`
	Supported_Clouds    []string                 `json:"supported_clouds"`
	Settings_Config     []Filter_Registry_Config `json:"settings_config"`
}

type Filter_Registry_Config struct {
	Field_Type   string                           `json:"field_type"`
	Name         string                           `json:"name"`
	Display_Name string                           `json:"display_name"`
	Description  interface{}                      `json:"description"` // Should be string but divvy.query.shared_file_system_lifecycle_policy was returning an array for a boolean field
	Options      []string                         `json:"options"`
	Choices      []Filter_Registry_Config_Choices `json:"choices"`
	State_Hash   string                           `json:"_state_hash"`
}

type Filter_Registry_Config_Choices struct {
	Value         string `json:"value"`
	Display_Value string `json:"display_value"`
}

// FUNCTIONS
///////////////////////////////////////////

func (c Client) Get_Filter_Registry() (map[string]Filter_Registry_Result, error) {
	// Returns a map of filter registry results
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/insights/filter-registry", nil)
	if err != nil {
		return nil, err
	}

	var fr map[string]Filter_Registry_Result
	if err := json.NewDecoder(resp.Body).Decode(&fr); err != nil {
		return nil, err
	}

	return fr, nil
}
