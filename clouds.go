package insightcloudsec

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Cloud struct {
	ID                  int                   `json:"id"`
	Name                string                `json:"name"`
	CloudTypeID         string                `json:"cloud_type_id"`
	AccountID           string                `json:"account_id"`
	Created             ICSTime               `json:"creation_time"`
	Status              string                `json:"status"`
	BadgeCount          int                   `json:"badge_count"`
	ResourceCount       int                   `json:"resource_count"`
	LastRefreshed       ICSTime               `json:"last_refreshed"`
	RoleARN             string                `json:"role_arn"`
	GroupResourceID     string                `json:"group_resource_id:"`
	ResourceID          string                `json:"resource_id:"`
	FailedResourceTypes []FailedResourceTypes `json:"failed_resource_types"`
	EDHRole             string                `json:"event_driven_harvest_role"`
	StrategyID          int                   `json:"strategy_id"`
	CloudOrgID          string                `json:"cloud_organization_id"`
}

type FailedResourceTypes struct {
	Type        string   `json:"resource_type"`
	Permissions []string `json:"permissions"`
}

type CloudList struct {
	Clouds []Cloud `json:"clouds"`
}

type CloudType struct {
	ID     string `json:"cloud_type_id"`
	Name   string `json:"name"`
	Access string `json:"cloud_access"`
}

type CloudTypesList struct {
	CloudTypes []CloudType `json:"clouds"`
}

type CloudRegion struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	ResourceID            string `json:"resource_id"`
	Status                string `json:"status"`
	HarvestRateMultiplier int    `json:"harvest_rate_multiplier"`
}

type CloudRegionList struct {
	Regions []CloudRegion `json:"regions"`
}

type HarvestingStrategy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OrgID       int    `json:"organization_id"`
	OrgServices int    `json:"organization_services"`
	Default     bool   `json:"type_default"`
	CloudTypeID string `json:"cloud_type_id"`
}

type HarvestingStrategyList struct {
	Strategies []HarvestingStrategy `json:"strategies"`
}

type ICSTime time.Time
// Need to Handle Time appropriately given how the API returns.

func (j *ICSTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*j = ICSTime(t)
	return nil
}

func (j *ICSTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}


// CLOUD FUNCTIONS
///////////////////////////////////////////
func (c Client) List_Clouds() (*CloudList, error) {
	// Return a CloudList item containing all the clouds from the API.
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/clouds/list", nil)
	if err != nil {
		return nil, err
	}

	var ret CloudList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}


func (c Client) List_Cloud_Types() (*CloudTypesList, error) {
	// Returns a CloudTypesList item containing all the cloud types from the API.
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil)
	if err != nil {
		return nil, err
	}

	var ret CloudTypesList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
