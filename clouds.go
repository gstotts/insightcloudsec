package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var _ Clouds = (*clouds)(nil)

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

type Clouds interface {
	AddAWSCloud(cloud_data AWSCloudAccount) (Cloud, error)
	AddAzureCloud(cloud_data AzureCloudAccount) (Cloud, error)
	AddGCPCloud(cloud_data GCPCloudAccount) (Cloud, error)
	Delete(cloud_resource_id string) error
	Update(id int, cloud_data CloudAccountParameters) (Cloud, error)
	List() (CloudList, error)
	ListHarvestingStrategies() ([]HarvestingStrategy, error)
	ListProvisioningClouds() (CloudList, error)
	ListRegions(target Cloud) (CloudRegionList, error)
	EnableRegionByName(target Cloud, region string) error
	DisableRegionByName(target Cloud, region string) error
	ListTypes() (CloudTypesList, error)
	GetByName(name string) (Cloud, error)
	GetByID(id int) (Cloud, error)
	QueueStatus() (QueueStatus, error)
	SystemStatus() (SystemStatus, error)
	PauseHarvesting(targets Cloud) error
	ResumeHarvesting(targets Cloud) error
}

type clouds struct {
	client *Client
}

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
	CloudOrgDomainName  string                `json:"cloud_organization_domain_name,omitempty"`
	CloudOrgNickname    string                `json:"cloud_organization_nickname,omitempty"`
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
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	ResourceID            string  `json:"resource_id"`
	Status                string  `json:"status"`
	HarvestRateMultiplier float32 `json:"harvest_rate_multiplier"`
}

type CloudRegionList struct {
	Regions []CloudRegion `json:"regions"`
}

type HarvestingStrategy struct {
	ID                       int      `json:"id"`
	Name                     string   `json:"name"`
	OrgID                    int      `json:"organization_id"`
	OrgServices              int      `json:"organization_services"`
	Default                  bool     `json:"type_default"`
	CloudTypeID              string   `json:"cloud_type_id"`
	SystemDefined            bool     `json:"system_defined"`
	DisabledRegions          []string `json:"disabled_regions"`
	DynamicScheduling        bool     `json:"dynamic_scheduling"`
	DynamicSchedulingEnabled bool     `json:"dynamic_scheduling_enabled"`
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
	Project               string             `json:"project,omitempty"`
}

type QueueStatus struct {
	P0                int       `json:"p0"`
	P1                int       `json:"p1"`
	P2                int       `json:"p2"`
	P3                int       `json:"p3"`
	SlowestJobs       []SlowJob `json:"slowest_jobs"`
	ProcessTime       TimeStats `json:"process_time"`
	ProcessTimeP0     TimeStats `json:"process_time_p0"`
	ProcessTimeP1     TimeStats `json:"process_time_p1"`
	ProcessTimeP2     TimeStats `json:"process_time_p2"`
	ProcessTimeP3     TimeStats `json:"process_time_p3"`
	Workers           int       `json:"workers"`
	QueueWait         TimeStats `json:"queue_wait"`
	QueueWaitP0       TimeStats `json:"queue_wait_p0"`
	QueueWaitP1       TimeStats `json:"queue_wait_p1"`
	QueueWaitP2       TimeStats `json:"queue_wait_p2"`
	QueueWaitP3       TimeStats `json:"queue_wait_p3"`
	QueueWaitAll      TimeStats `json:"queue_wait_all"`
	SchedulerInternal int       `json:"scheduler_internal"`
}

type SlowJob struct {
	Name     string
	Duration float64
}

type SystemStatus struct {
	Diagnostics []map[string]interface{} `json:"diagnostics"`
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
	Count    int     `json:"count"`
	Min      float32 `json:"min"`
	Max      float32 `json:"max"`
	Sum      float32 `json:"sum"`
	SumSQ    float64 `json:"sumsq"`
	StdDev   float64 `json:"stddev"`
	Average  float64 `json:"average"`
	Current  float64 `json:"current,omitempty"`
	Variance float64 `json:"variance,omitempty"`
}

// CLOUD ACCOUNT SETUP FUNCTIONS
///////////////////////////////////////////

func (s *clouds) AddAWSCloud(cloud_data AWSCloudAccount) (Cloud, error) {
	err := validateAWSCloud(cloud_data)
	if err != nil {
		return Cloud{}, err
	}

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/prototype/cloud/add", cloud_data)
	if err != nil {
		return Cloud{}, err
	}

	var ret Cloud
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Cloud{}, err
	}

	return ret, nil
}

func (s *clouds) AddAzureCloud(cloud_data AzureCloudAccount) (Cloud, error) {
	err := validateAzureCloud(cloud_data)
	if err != nil {
		return Cloud{}, err
	}

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/prototype/cloud/add", cloud_data)
	if err != nil {
		return Cloud{}, err
	}

	var ret Cloud
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Cloud{}, err
	}

	return ret, nil
}

func (s *clouds) AddGCPCloud(cloud_data GCPCloudAccount) (Cloud, error) {
	err := validateGCPCloud(cloud_data)
	if err != nil {
		return Cloud{}, err

	}

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/prototype/cloud/add", cloud_data)
	if err != nil {
		return Cloud{}, err
	}

	var ret Cloud
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Cloud{}, err
	}

	return ret, nil
}

func (s *clouds) Update(id int, cloud_data CloudAccountParameters) (Cloud, error) {
	if cloud_data.CloudType == AWS_CLOUD_TYPE {
		err := validateAWSCloud(AWSCloudAccount{cloud_data})
		if err != nil {
			return Cloud{}, err
		}
	} else if cloud_data.CloudType == AZURE_CLOUD_TYPE {
		err := validateAzureCloud(AzureCloudAccount{cloud_data})
		if err != nil {
			return Cloud{}, err
		}
	} else if cloud_data.CloudType == GCP_CLOUD_TYPE {
		err := validateGCPCloud(GCPCloudAccount{cloud_data})
		if err != nil {
			return Cloud{}, err
		}
	} else {
		return Cloud{}, fmt.Errorf("[-] ERROR: Invalid cloud type to update: %s", cloud_data.CloudType)
	}

	resp, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/cloud/%d/update", id), cloud_data)
	if err != nil {
		return Cloud{}, err
	}

	var ret Cloud
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Cloud{}, err
	}

	return ret, nil
}

func validateAWSCloud(cloud_data AWSCloudAccount) error {
	if cloud_data.CreationParameters.CloudType != AWS_CLOUD_TYPE {
		return fmt.Errorf("[-] ERROR: cloud account must be of type AWS to use, not %s", cloud_data.CreationParameters.CloudType)
	}

	if cloud_data.CreationParameters.AuthType == STS_ASSUME_AUTH {
		// If using STS Assume Role, make sure secret and key are set
		if cloud_data.CreationParameters.ApiKeyOrCert == "" || cloud_data.CreationParameters.SecretKey == "" {
			return fmt.Errorf("[-] ERROR: assume role AWS accounts require a secret and key are set")
		}
	}

	// Make sure AWS properties exist only, otherwise return error
	if cloud_data.CreationParameters.TenantID != "" || cloud_data.CreationParameters.SubscriptionID != "" || cloud_data.CreationParameters.AppID != "" || cloud_data.CreationParameters.GCPAuth.Type != "" {
		return fmt.Errorf("[-] ERROR: cloud account of type AWS must not have TenantID, SubscriptionID or AppID set")
	}

	return nil
}

func validateAzureCloud(cloud_data AzureCloudAccount) error {
	if cloud_data.CreationParameters.CloudType != AZURE_CLOUD_TYPE {
		return fmt.Errorf("[-] ERROR: cloud account must be of type AZURE_ARM to use, not %s", cloud_data.CreationParameters.CloudType)
	}
	if cloud_data.CreationParameters.AuthType == STANDARD_AUTH && cloud_data.CreationParameters.ApiKeyOrCert == "" {
		return fmt.Errorf("[-] ERROR: azure cloud of AuthType standard requires ApiKeyOrCert be set")
	} else if cloud_data.CreationParameters.AuthType == CERT_AUTH && (cloud_data.CreationParameters.ApiKeyOrCert == "" || cloud_data.CreationParameters.CertificateThumbprint == "") {
		// If using cert auth, make sure pem and thumbprint set
		return fmt.Errorf("[-] ERROR: azure cloud of AuthType client_certificate requires ApiKeyOrCert and CertificateThumbprint be set")
	} else if cloud_data.CreationParameters.AuthType != STANDARD_AUTH && cloud_data.CreationParameters.AuthType != CERT_AUTH {
		return fmt.Errorf("[-] ERROR: azure cloud accounts must use authtype standard or client_certificate, not %s", cloud_data.CreationParameters.AuthType)
	}

	// Make sure Azure properties exist only, otherwise eliminate
	if cloud_data.CreationParameters.RoleArn != "" || cloud_data.CreationParameters.SecretKey != "" || cloud_data.CreationParameters.SessionName != "" || cloud_data.CreationParameters.Duration != 0 || cloud_data.CreationParameters.AccountNumber != "" || cloud_data.CreationParameters.ExternalID != "" || cloud_data.CreationParameters.GCPAuth.Type != "" {
		return fmt.Errorf("[-] ERROR: cloud account of type AZURE_ARM must not have RoleArn, SecretKey, SessionName, Duration, AccountNumber or ExternalID set")
	}

	// Make sure tenant_id and

	return nil
}

func validateGCPCloud(cloud_data GCPCloudAccount) error {
	if cloud_data.CreationParameters.CloudType != GCP_CLOUD_TYPE {
		return fmt.Errorf("[-] ERROR: cloud account must be of type GCE to use, not %s", cloud_data.CreationParameters.CloudType)
	}
	// Validate required GCE settings exist
	if cloud_data.CreationParameters.GCPAuth.ProjectID == "" || cloud_data.CreationParameters.GCPAuth.Type == "" {
		return fmt.Errorf("[-] ERROR: cloud account of type GCE requires ProjectID and API Credentials be set")
	}
	// Throw error if other settings exist
	if cloud_data.CreationParameters.AuthType != "" || cloud_data.CreationParameters.ApiKeyOrCert != "" || cloud_data.CreationParameters.SecretKey != "" || cloud_data.CreationParameters.RoleArn != "" || cloud_data.CreationParameters.SessionName != "" || cloud_data.CreationParameters.ExternalID != "" || cloud_data.CreationParameters.TenantID != "" || cloud_data.CreationParameters.SubscriptionID != "" || cloud_data.CreationParameters.AppID != "" {
		return fmt.Errorf("[-] ERROR: cloud account of type GCE must not have AuthType, ApiKey, SecretKey, RoleArn, SessionName, ExternalID, TenantID, SubscriptionID or AppID set")
	}

	return nil
}

func (s *clouds) Delete(cloud_resource_id string) error {
	resp, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/cloud/%s/delete", cloud_resource_id), nil)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}

// MANAGING CLOUD FUNCTIONS
///////////////////////////////////////////

func (s *clouds) List() (CloudList, error) {
	// Return a CloudList item containing all the clouds from the API.
	resp, err := s.client.makeRequest(http.MethodGet, "/v2/public/clouds/list", nil)
	if err != nil {
		return CloudList{}, err
	}

	var ret CloudList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return CloudList{}, err
	}

	return ret, nil
}

func (s *clouds) GetByName(name string) (Cloud, error) {
	// Returns the specific cloud of the name given.
	all_clouds, err := s.List()
	if err != nil {
		return Cloud{}, err
	}

	for _, cloud := range all_clouds.Clouds {
		if strings.EqualFold(cloud.Name, name) {
			return cloud, nil
		}
	}

	return Cloud{}, fmt.Errorf("[-] ERROR: Cloud Named %s Not Found", name)
}

func (s *clouds) GetByID(id int) (Cloud, error) {
	// Returns the specific cloud of the ID given.
	all_clouds, _ := s.List()
	for _, cloud := range all_clouds.Clouds {
		if id == cloud.ID {
			return cloud, nil
		}
	}

	return Cloud{}, fmt.Errorf("[-] ERROR: Cloud of ID %d Not Found", id)
}

func (s *clouds) ListTypes() (CloudTypesList, error) {
	// Returns a CloudTypesList item containing all the cloud types from the API.
	resp, err := s.client.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil)
	if err != nil {
		return CloudTypesList{}, err
	}

	var ret CloudTypesList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return CloudTypesList{}, err
	}

	return ret, nil
}

func (s *clouds) ListProvisioningClouds() (CloudList, error) {
	// Returns a list of provisioning clouds.
	resp, err := s.client.makeRequest(http.MethodGet, "/v2/public/clouds/provisioning/list", nil)
	if err != nil {
		return CloudList{}, err
	}

	var ret CloudList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return CloudList{}, err
	}
	return ret, nil
}

func (s *clouds) QueueStatus() (QueueStatus, error) {
	// Returns the queue status statistics.
	resp, err := s.client.makeRequest(http.MethodGet, "/v2/prototype/diagnostics/queues/status/get", nil)
	if err != nil {
		return QueueStatus{}, err
	}

	var ret QueueStatus
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return QueueStatus{}, err
	}
	return ret, nil
}

func (s *clouds) SystemStatus() (SystemStatus, error) {
	// Returns the system status diagnostics as a map[string]interface{}
	resp, err := s.client.makeRequest(http.MethodGet, "/v2/prototype/diagnostics/system/status/get", nil)
	if err != nil {
		return SystemStatus{}, err
	}

	var ret SystemStatus
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return SystemStatus{}, err
	}
	return ret, nil
}

func (s *clouds) ListHarvestingStrategies() ([]HarvestingStrategy, error) {
	// Returns a HarvestingStrategyList item containing all the cloud harvesting strategies from the API.
	resp, err := s.client.makeRequest(http.MethodGet, "/v2/harvestingstrategy/strategy", nil)
	if err != nil {
		return nil, err
	}

	var ret HarvestingStrategyList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.Strategies, nil
}

func (s *clouds) ListRegions(target Cloud) (CloudRegionList, error) {
	// Returns a CloudRegionList for the given Cloud.
	var ret CloudRegionList
	resp, err := s.client.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/cloud/%s/regions/list", target.ResourceID), nil)
	if err != nil {
		return CloudRegionList{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return CloudRegionList{}, err
	}

	return ret, nil
}

func (s *clouds) DisableRegionByName(target Cloud, region string) error {
	regions, err := s.client.Clouds.ListRegions(target)
	if err != nil {
		return err
	}

	var resource_id string
	for _, list_region := range regions.Regions {
		if strings.EqualFold(list_region.Name, region) {
			resource_id = list_region.ResourceID
		}
	}

	if resource_id == "" {
		return fmt.Errorf("[-] ERROR: Region Named %s Not Found", region)
	}

	_, err = s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/cloud/region/%s/disable", resource_id), nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *clouds) EnableRegionByName(target Cloud, region string) error {
	regions, err := s.client.Clouds.ListRegions(target)
	if err != nil {
		return err
	}

	var resource_id string
	for _, list_region := range regions.Regions {
		if strings.EqualFold(list_region.Name, region) {
			resource_id = list_region.ResourceID
		}
	}

	if resource_id == "" {
		return fmt.Errorf("[-] ERROR: Region Named %s Not Found", region)
	}

	_, err = s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/cloud/region/%s/enable", resource_id), nil)
	if err != nil {
		return err
	}
	return nil
}

func setHarvestingStatus(target Cloud, status string) map[string]interface{} {
	payload := map[string]interface{}{
		"resource_ids": []string{target.ResourceID},
		"status":       status,
	}

	return payload
}

func (s *clouds) PauseHarvesting(target Cloud) error {
	_, err := s.client.makeRequest(http.MethodPost, "/v2/public/clouds/status/set", setHarvestingStatus(target, "PAUSED"))
	if err != nil {
		return err
	}

	return nil
}

func (s *clouds) ResumeHarvesting(target Cloud) error {

	_, err := s.client.makeRequest(http.MethodPost, "/v2/public/clouds/status/set", setHarvestingStatus(target, "DEFAULT"))
	if err != nil {
		return err
	}

	return nil
}
