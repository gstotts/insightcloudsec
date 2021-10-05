package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type ResourceDetails struct {
	Dependencies ResourceDependencies `json:"dependencies"`
	Details      Resource             `json:"details"`
}

type ResourceDependencies struct {
}

// See resource_structs.go for individual resource types
type Resource struct {
	Type                     string                           `json:"resource_type"`
	AccessFlowLog            AccessFlowLogResource            `json:"accesslistflowlog,omitempty"`
	AccessAnalyzer           AccessAnalyzerResource           `json:"accessanalyzer,omitempty"`
	AirflowEnvironment       AirflowEnvironmentResource       `json:"airflowenvironment,omitempty"`
	APIAccountingConfig      APIAccountingConfigResource      `json:"apiaccountingconfig,omitempty"`
	AppRunnerService         AppRunnerServiceResource         `json:"apprunnerservice,omitempty"`
	AppServer                AppServerResource                `json:"appserver,omitempty"`
	AutoscalingGroup         AutoscalingGroupResource         `json:"autoscalinggroup,omitempty"`
	AutoscalingLaunchConfig  AutoscalingLaunchConfigResource  `json:"autoscalinglaunchconfiguration,omitempty"`
	AWSConfig                AWSConfigResource                `json:"awsconfig,omitempty"`
	BackendService           BackendServiceResource           `json:"backendservice,omitempty"`
	BackupVault              BackupVaultResource              `json:"backupvault,omitempty"`
	BatchEnvironment         BatchEnvironmentResource         `json:"batchenvironment,omitempty"`
	BatchPool                BatchPoolResource                `json:"batchpool,omitempty"`
	BigDataInstance          BigDataInstanceResource          `json:"bigdatainstance,omitempty"`
	BigDataSnapshot          BigDataSnapshotResource          `json:"bigdatasnapshot,omitempty"`
	BigDataWorkspace         BigDataWorkspaceResource         `json:"bigdataworkspace,omitempty"`
	BrokerInstance           BrokerInstanceResource           `json:"brokerinstance,omitempty"`
	BuildProject             BuildProjectResource             `json:"buildproject,omitempty"`
	CloudwatchDestination    CloudwatchDestinationResource    `json:"cloudwatchdestination,omitempty"`
	ColdStorage              ColdStorageResource              `json:"coldstorage,omitempty"`
	Container                ContainerResource                `json:"container,omitempty"`
	ContainerCluster         ContainerClusterResource         `json:"containercluster,omitempty"`
	ContainerDeployment      ContainerDeploymentResource      `json:"containerdeployment,omitempty"`
	ContainerImage           ContainerImageResource           `json:"containerimage,omitempty"`
	ContainerInstance        ContainerInstanceResource        `json:"containerinstance,omitempty"`
	ContainerRegistry        ContainerRegistryResource        `json:"containerregistry,omitempty"`
	ContainerService         ContainerServiceResource         `json:"containerservice,omitempty"`
	CDN                      ContentDeliveryNetworkResource   `json:"contentdeliverynetwork,omitempty"`
	DivvyOrganizationService DivvyOrganizationServiceResource `json:"divvyorganizationservice,omitempty"`
	Instance                 InstanceResource                 `json:"instance,omitempty"`
	Volume                   VolumeResource                   `json:"volume,omitempty"`
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
func (c Client) Detail_Resource(id string) (*Resource, error) {
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/resource/%s/detail", id), nil)
	if err != nil {
		return nil, err
	}

	var ret *Resource
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
