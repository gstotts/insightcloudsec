package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////
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
	GroupResourceID     string                `json:"group_resource_id"`
	ResourceID          string                `json:"resource_id"`
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

type AWSCloudAccount struct {
	CreationParameters CloudAccountParameters `json:"creation_params"`
}

type CloudAccountParameters struct {
	CloudType     string `json:"cloud_type"`
	AuthType      string `json:"authentication_type"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	ApiKey        string `json:"api_key,omitempty"`
	SecretKey     string `json:"secret_key,omitempty"`
	RoleArn       string `json:"role_arn"`
	ExternalID    string `json:"external_id"`
	Duration      int    `json:"duration"`
	SessionName   string `json:"session_name"`
}

// CLOUD FUNCTIONS
///////////////////////////////////////////

func (c Client) Add_AWS_Cloud(cloud_data AWSCloudAccount) (Cloud, error) {
	if cloud_data.CreationParameters.AuthType == "assume_role" {
		// If using STS Assume Role, make sure secret and key are set
		if cloud_data.CreationParameters.ApiKey == "" || cloud_data.CreationParameters.SecretKey == "" {
			return Cloud{}, fmt.Errorf("[-] ERROR: assume role AWS accounts require a secret and key are set")
		}
	}
	fmt.Println(cloud_data)
	data, err := json.Marshal(cloud_data)
	if err != nil {
		return Cloud{}, err
	}
	fmt.Println(data)

	resp, err := c.makeRequest(http.MethodPost, "/v2/prototype/cloud/add", bytes.NewBuffer(data))
	if err != nil {
		return Cloud{}, err
	}

	var ret Cloud
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Cloud{}, err
	}

	return ret, nil
}

func (c Client) List_Clouds() ([]Cloud, error) {
	// Return a CloudList item containing all the clouds from the API.
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/clouds/list", nil)
	if err != nil {
		return nil, err
	}

	var ret CloudList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.Clouds, nil
}

func (c Client) List_Cloud_Types() ([]CloudType, error) {
	// Returns a CloudTypesList item containing all the cloud types from the API.
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil)
	if err != nil {
		return nil, err
	}

	var ret CloudTypesList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.CloudTypes, nil
}

func (c Client) List_Harvesting_Strategies() ([]HarvestingStrategy, error) {
	// Returns a HarvestingStrategyList item containing all the cloud harvesting strategies from the API.
	resp, err := c.makeRequest(http.MethodGet, "/v2/harvestingstrategy/strategy", nil)
	if err != nil {
		return nil, err
	}

	var ret HarvestingStrategyList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.Strategies, nil
}

func (c Client) List_Cloud_Regions(target Cloud) ([]CloudRegion, error) {
	// Returns a CloudRegionList for the given Cloud.
	var ret CloudRegionList
	fmt.Println(target.ResourceID)
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/cloud/%s/regions/list", target.ResourceID), nil)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.Regions, nil
}
