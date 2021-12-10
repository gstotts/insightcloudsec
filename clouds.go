package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CONSTANTS
///////////////////////////////////////////
const (
	// Cloud Type Constants
	AWS_CLOUD_TYPE   = "AWS"
	AZURE_CLOUD_TYPE = "AZURE_ARM"
	GCP_CLOUD_TYPE   = "GCE"

	// Cloud Authentication Type Constants
	STS_ASSUME_AUTH      = "assume_role"
	INSTANCE_ASSUME_AUTH = "instance_assume_role"
	STANDARD_AUTH        = "standard"
	CERT_AUTH            = "client_certificate"
	GCP_SERVICE_ACCOUNT  = "service_account"

	// GCP Default URIs
	GCP_DEFAULT_AUTH_URI  = "https://accounts.google.com/o/oauth2/auth"
	GCP_DEFAULT_TOKEN_URI = "https://accounts.google.com/o/oauth2/token"
)

// STRUCTS
///////////////////////////////////////////
type Cloud struct {
	ID                  int                   `json:"id"`
	Name                string                `json:"name"`
	CloudTypeID         string                `json:"cloud_type_id"`
	AccountID           string                `json:"account_id"`
	Created             string                `json:"creation_time"`
	Status              string                `json:"status"`
	BadgeCount          int                   `json:"badge_count,omitempty"`
	ResourceCount       int                   `json:"resource_count,omitempty"`
	LastRefreshed       string                `json:"last_refreshed"`
	RoleARN             string                `json:"role_arn,omitempty"`
	GroupResourceID     string                `json:"group_resource_id"`
	ResourceID          string                `json:"resource_id"`
	FailedResourceTypes []FailedResourceTypes `json:"failed_resource_types,omitempty"`
	EDHRole             string                `json:"event_driven_harvest_role,omitempty"`
	StrategyID          int                   `json:"strategy_id,omitempty"`
	CloudOrgID          string                `json:"cloud_organization_id,omitempty"`
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

type AzureCloudAccount struct {
	CreationParameters CloudAccountParameters `json:"creation_params"`
}

type GCPCloudAccount struct {
	CreationParameters CloudAccountParameters `json:"creation_params"`
}

type GCPAccountApiCreds struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderx509CertURL string `json:"auth_provider_x509_cert_url"`
	Clientx509CertUrl       string `json:"client_x509_cert_url"`
}

type CloudAccountParameters struct {
	CloudType             string             `json:"cloud_type"`
	AuthType              string             `json:"authentication_type,omitempty"`
	Name                  string             `json:"name"`
	AccountNumber         string             `json:"account_number,omitempty"`
	ApiKeyOrCert          string             `json:"api_key,omitempty"`
	SecretKey             string             `json:"secret_key,omitempty"`
	RoleArn               string             `json:"role_arn,omitempty"`
	ExternalID            string             `json:"external_id,omitempty"`
	Duration              int                `json:"duration,omitempty"`
	SessionName           string             `json:"session_name,omitempty"`
	TenantID              string             `json:"tenant_id,omitempty"`
	AppID                 string             `json:"app_id,omitempty"`
	SubscriptionID        string             `json:"subscription_id,omitempty"`
	CertificateThumbprint string             `json:"certificate_thumbprint,omitempty"`
	GCPAuth               GCPAccountApiCreds `json:"api_credentials,omitempty"`
}

type QueueStatus struct {
	P0          int       `json:"p0"`
	P1          int       `json:"p1"`
	P2          int       `json:"p2"`
	SlowestJobs []SlowJob `json:"slowest_jobs"`
	ProcessTime TimeStats `json:"process_time"`
	Workers     int       `json:"workers"`
	QueueWait   TimeStats `json:"queue_wait"`
}

type SlowJob struct {
	Name     string
	Duration float64
}

// Custom Unmarshal to separate out the array into name and duration
func (s *SlowJob) UnmarshalJSON(b []byte) error {
	var v []interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	s.Name, _ = v[0].(string)
	s.Duration, _ = v[1].(float64)
	return nil
}

type TimeStats struct {
	Count   int     `json:"count"`
	Min     float32 `json:"min"`
	Max     float32 `json:"max"`
	Sum     float32 `json:"sum"`
	SumSQ   float64 `json:"sumsq"`
	StdDev  float64 `json:"stddev"`
	Average float64 `json:"average"`
}

// CLOUD ACCOUNT SETUP FUNCTIONS
///////////////////////////////////////////

func (c Client) AddAWSCloud(cloud_data AWSCloudAccount) (Cloud, error) {
	if cloud_data.CreationParameters.CloudType != AWS_CLOUD_TYPE {
		return Cloud{}, fmt.Errorf("[-] ERROR: cloud account must be of type AWS to use, not %s", cloud_data.CreationParameters.CloudType)
	}

	if cloud_data.CreationParameters.AuthType == STS_ASSUME_AUTH {
		// If using STS Assume Role, make sure secret and key are set
		if cloud_data.CreationParameters.ApiKeyOrCert == "" || cloud_data.CreationParameters.SecretKey == "" {
			return Cloud{}, fmt.Errorf("[-] ERROR: assume role AWS accounts require a secret and key are set")
		}
	}

	// Make sure AWS properties exist only, otherwise return error
	if cloud_data.CreationParameters.TenantID != "" || cloud_data.CreationParameters.SubscriptionID != "" || cloud_data.CreationParameters.AppID != "" || cloud_data.CreationParameters.GCPAuth.Type != "" {
		return Cloud{}, fmt.Errorf("[-] ERROR: cloud account of type AWS must not have TenantID, SubscriptionID or AppID set")
	}

	data, err := json.Marshal(cloud_data)
	if err != nil {
		return Cloud{}, err
	}

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

func (c Client) AddAzureCloud(cloud_data AzureCloudAccount) (Cloud, error) {
	if cloud_data.CreationParameters.CloudType != AZURE_CLOUD_TYPE {
		return Cloud{}, fmt.Errorf("[-] ERROR: cloud account must be of type AZURE_ARM to use, not %s", cloud_data.CreationParameters.CloudType)
	}

	if cloud_data.CreationParameters.AuthType == STANDARD_AUTH && cloud_data.CreationParameters.ApiKeyOrCert == "" {
		return Cloud{}, fmt.Errorf("[-] ERROR: azure cloud of AuthType standard requires ApiKeyOrCert be set")
	} else if cloud_data.CreationParameters.AuthType == CERT_AUTH && (cloud_data.CreationParameters.ApiKeyOrCert == "" || cloud_data.CreationParameters.CertificateThumbprint == "") {
		// If using cert auth, make sure pem and thumbprint set
		return Cloud{}, fmt.Errorf("[-] ERROR: azure cloud of AuthType client_certificate requires ApiKeyOrCert and CertificateThumbprint be set")
	} else if cloud_data.CreationParameters.AuthType != STANDARD_AUTH && cloud_data.CreationParameters.AuthType != CERT_AUTH {
		return Cloud{}, fmt.Errorf("[-] ERROR: azure cloud accounts must use authtype standard or client_certificate, not %s", cloud_data.CreationParameters.AuthType)
	}

	// Make sure Azure properties exist only, otherwise eliminate
	if cloud_data.CreationParameters.RoleArn != "" || cloud_data.CreationParameters.SecretKey != "" || cloud_data.CreationParameters.SessionName != "" || cloud_data.CreationParameters.Duration != 0 || cloud_data.CreationParameters.AccountNumber != "" || cloud_data.CreationParameters.ExternalID != "" || cloud_data.CreationParameters.GCPAuth.Type != "" {
		return Cloud{}, fmt.Errorf("[-] ERROR: cloud account of type AZURE_ARM must not have RoleArn, SecretKey, SessionName, Duration, AccountNumber or ExternalID set")
	}

	data, err := json.Marshal(cloud_data)
	if err != nil {
		return Cloud{}, err
	}

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

func (c Client) AddGCPCloud(cloud_data GCPCloudAccount) (Cloud, error) {
	if cloud_data.CreationParameters.CloudType != GCP_CLOUD_TYPE {
		return Cloud{}, fmt.Errorf("[-] ERROR: cloud account must be of type GCE to use, not %s", cloud_data.CreationParameters.CloudType)
	}

	data, err := json.Marshal(cloud_data)
	if err != nil {
		return Cloud{}, err
	}

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

// CLOUD ORG FUNCTIONS
///////////////////////////////////////////

// MANAGING CLOUD FUNCTIONS
///////////////////////////////////////////

func (c Client) ListClouds() ([]Cloud, error) {
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

func (c Client) ListCloudTypes() (CloudTypesList, error) {
	// Returns a CloudTypesList item containing all the cloud types from the API.
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil)
	if err != nil {
		return CloudTypesList{}, err
	}

	var ret CloudTypesList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return CloudTypesList{}, err
	}

	return ret, nil
}

func (c Client) ListProvisioningClouds() (CloudList, error) {
	// Returns a list of provisioning clouds.
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/clouds/provisioning/list", nil)
	if err != nil {
		return CloudList{}, err
	}

	var ret CloudList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return CloudList{}, err
	}
	return ret, nil
}

func (c Client) QueueStatus() (QueueStatus, error) {
	// Returns the queue status statistics.
	resp, err := c.makeRequest(http.MethodGet, "/v2/prototype/diagnostics/queues/status/get", nil)
	if err != nil {
		return QueueStatus{}, err
	}

	var ret QueueStatus
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return QueueStatus{}, err
	}
	return ret, nil
}

func (c Client) ListHarvestingStrategies() ([]HarvestingStrategy, error) {
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

func (c Client) ListCloudRegions(target Cloud) ([]CloudRegion, error) {
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
