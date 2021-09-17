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
	Instance InstanceResource `json:"instance,omitempty"`
	Volume   VolumeResource   `json:"volume,omitempty"`
	Type     string           `json:"resource_type"`
}

type CommonResourceValues struct {
	//Common Attributes
	ID               string            `json:"resource_id"`
	Name             string            `json:"resource_name"`
	Type             string            `json:"type"`
	Cloud            string            `json:"cloud"`
	Account          string            `json:"account"`
	AccountID        string            `json:"account_id"`
	AccountStatus    string            `json:"account_status"`
	OrgServiceID     int               `json:"organization_service_id"`
	AvailabilityZone string            `json:"availablility_zone"`
	Region           string            `json:"region"`
	Created          string            `json:"creation_timestamp"`
	Discovered       string            `json:"discovered_timestamp"`
	Modified         string            `json:"modified_timestamp"`
	NamespaceID      string            `json:"namespace_id"`
	Tags             map[string]string `json:"tags"`
}

type InstanceResource struct {
	Common                                         CommonResourceValues `json:"common"`
	ID                                             string               `json:"instance_id"`
	Type                                           string               `json:"instance_type"`
	LaunchTime                                     string               `json:"launch_time"`
	Platform                                       string               `json:"platform"`
	State                                          string               `json:"state"`
	ImageID                                        string               `json:"image_id"`
	PublicIPAddress                                string               `json:"public_ip_address,omitempty"`
	PrivateIPAddress                               string               `json:"private_ip_address"`
	NetworkResourceID                              string               `json:"network_resource_id"`
	SubnetResourceID                               string               `json:"subnet_resource_id"`
	ObjectID                                       string               `json:"object_id"`
	KeyName                                        string               `json:"key_name,omitempty"`
	RoleName                                       string               `json:"role_name,omitempty"`
	TerminationProtection                          string               `json:"termination_protection,omitempty"`
	VMExtensions                                   string               `json:"vm_extensions,omitempty"`
	JITAccessPolicy                                string               `json:"jit_access_policy"`
	RootDeviceType                                 string               `json:"root_device_type,omitempty"`
	Tenancy                                        string               `json:"tenancy,omitempty"`
	DetailedMonitoring                             bool                 `json:"detailed_monitoring,omitempty"`
	SecondaryPrivateIPAddresses                    []string             `json:"secondary_private_ip_addresses,omitempty"`
	SecondaryPublicIPAddresses                     []string             `json:"secondary_public_ip_addresses,omitempty"`
	Architecture                                   string               `json:"architecture,omitempty"`
	AWSInstanceMetadataServiceV2Required           bool                 `json:"aws_instance_metadata_service_v2_required,omitempty"`
	AWSInstanceMetadataServiceHopLimit             int                  `json:"aws_instance_metadata_hop_limit,omitempty"`
	AWSInstanceMetadataServiceEndpointEnabled      string               `json:"aws_instance_metadata_service_endpoint_enabled,omitempty"`
	AWSInstanceMetadataServiceEndpointConfigStatus string               `json:"aws_instance_metadata_service_endpoint_config_status,omitempty"`
	StateTransitionReason                          string               `json:"state_transition_reason,omitempty"`
}

type VolumeResource struct {
	Common              CommonResourceValues `json:"common"`
	ID                  string               `json:"volume_id"`
	Type                string               `json:"volume_type"`
	Size                int                  `json:"size"`
	State               string               `json:"state"`
	IOPS                string               `json:"rated_iops,omitempty"`
	Encrypted           bool                 `json:"encrypted"`
	DeleteOnTermination bool                 `json:"delete_on_termination"`
	AttachState         string               `json:"attach_state"`
	AttachDeviceName    string               `json:"attach_device_name,omitempty"`
	InstanceAssociation InstanceAssociation  `json:"instance_association"`
	Created             string               `json:"creation_time"`
}

type InstanceAssociation struct {
	Common InstanceAssociationDetails `json:"common"`
}

type InstanceAssociationDetails struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
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
