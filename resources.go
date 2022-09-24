// Last Reviewed: 2-Apr-2022
// InsightCloudSec Version at time of review: 22.2

package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var _ Resources = (*resources)(nil)

type Resources interface {
	Query(q Query) (Query_Results, error)
	GetDetails(resource_id string) (Resource_Details, error)
	GetAssociations(resource_id string) (Resource_Associations, error)
	ListTags(resource_id string) ([]Tag, error)
	List_Settings(resource_id string) (Resource_Settings, error)
	SetOwner(resource_ids []string, owner_resource_id string) error
}

type resources struct {
	client *Client
}

type Query struct {
	// Built off of the Query v3-ETL endpoint of the InsightCloudSec API
	Badges                 *[]Badge        `json:"badges,omitempty"`
	Badge_Filter_Operator  string          `json:"badge_filter_operator,omitempty"`
	Filters                *[]Query_Filter `json:"filters,omitempty"`
	Insight                string          `json:"insight,omitempty"`
	Limit                  int32           `json:"limit"`
	Offset                 int32           `json:"offset,omitempty"`
	OrderBy                string          `json:"order_by,omitempty"`
	Scopes                 []string        `json:"scopes,omitempty"`
	Selected_Resource_Type string          `json:"selected_resource_type,omitempty"`
	Tags                   *[]Tag          `json:"tags,omitempty"`
	Cursor                 string          `json:"cursor,omitempty"`
}

type Query_Filter struct {
	// The name and configuration (as strings) of a query filter from the filter registry
	Name   string `json:"name"`
	Config string `json:"config"`
}

type Query_Results struct {
	// The result response from the query provided
	Counts                 map[string]int     `json:"counts,omitempty"`
	Selected_Resource_Type string             `json:"selected_resource_type"`
	Supported_Types        []string           `json:"supported_types"`
	Resources              []Resource_Results `json:"resources"`
	Scopes                 []string           `json:"scopes"`
	Limit                  int32              `json:"limit"`
	Offset                 int32              `json:"offset"`
	Order_By               string             `json:"order_by"`
	Filters                []Query_Filter     `json:"filters"`
	Next_Cursor            string             `json:"next_cursor"`
}

type Resource_Details struct {
	Dependencies map[string]([]Dependency_Details) `json:"dependencies"`
	Details      Resource_Results                  `json:"details"`
}

type Dependency_Details struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Resource_ID string `json:"resource_id"`
	Type        string `json:"type"`
}

type Resource_Results struct {
	Resource_Type                        string                               `json:"resource_type"`
	Access_Analyzer                      Access_Analyzer                      `json:"accessanalyzer,omitempty"`
	Access_List_Flow_Log                 Access_List_Flow_Log                 `json:"accesslistflowlog,omitempty"`
	Airflow_Environment                  Airflow_Environment                  `json:"airflowenvironment,omitempty"`
	API_Accounting_Config                API_Accounting_Config                `json:"apiaccountingconfig,omitempty"`
	App_Runner_Service                   App_Runner_Service                   `json:"apprunnerservice,omitempty"`
	App_Server                           App_Server                           `json:"appserver,omitempty"`
	Autoscaling_Group                    Autoscaling_Group                    `json:"autoscalinggroup,omitempty"`
	Autoscaling_Launch_Configuration     Autoscaling_Launch_Configuration     `json:"autoscalinglaunchconfiguration,omitempty"`
	AWS_Config                           AWS_Config                           `json:"awsconfig,omitempty"`
	Backend_Service                      Backend_Service                      `json:"backendservice,omitempty"`
	Backup_Vault                         Backup_Vault                         `json:"backupvault,omitempty"`
	Batch_Environment                    Batch_Environment                    `json:"batchenvironment,omitempty"`
	Batch_Pool                           Batch_Pool                           `json:"batchpool,omitempty"`
	BigData_Instance                     BigData_Instance                     `json:"bigdatainstance,omitempty"`
	BigData_Snapshot                     BigData_Snapshot                     `json:"bigdatasnapshot,omitempty"`
	BigData_Workspace                    BigData_Workspace                    `json:"bigdataworkspace,omitempty"`
	Broker_Instance                      Broker_Instance                      `json:"brokerinstance,omitempty"`
	Build_Project                        Build_Project                        `json:"buildproject,omitempty"`
	CloudWatch_Destination               CloudWatch_Destination               `json:"cloudwatchdestination,omitempty"`
	Cold_Storage                         Cold_Storage                         `json:"coldstorage,omitempty"`
	Container                            Container                            `json:"container,omitempty"`
	Container_Cluster                    Container_Cluster                    `json:"containercluster,omitempty"`
	Container_Deployment                 Container_Deployment                 `json:"containerdeployment,omitempty"`
	Container_Image                      Container_Image                      `json:"containerimage,omitempty"`
	Container_Instance                   Container_Instance                   `json:"containerinstance,omitempty"`
	Container_Node_Group                 Container_Node_Group                 `json:"containernodegroup,omitempty"`
	Container_Registry                   Container_Registry                   `json:"containerregistry,omitempty"`
	Container_Service                    Container_Service                    `json:"containerservice,omitempty"`
	Content_Delivery_Network             Content_Delivery_Network             `json:"contentdeliverynetwork,omitempty"`
	Data_Analytics_Workspace             Data_Analytics_Workspace             `json:"dataanalyticsworkspace,omitempty"`
	Database                             Database                             `json:"database,omitempty"`
	Databrick_Workspace                  Databrick_Workspace                  `json:"databricksworkspace,omitempty"`
	Data_Factory                         Data_Factory                         `json:"datafactory,omitempty"`
	Data_Lake_Storage                    Data_Lake_Storage                    `json:"datalakestorage,omitempty"`
	Datastore                            Datastore                            `json:"datastore,omitempty"`
	Datastream                           Datastream                           `json:"datastream,omitempty"`
	Datasync_Task                        Datasync_Task                        `json:"datasynctask,omitempty"`
	DB_Cluster                           DB_Cluster                           `json:"dbcluster,omitempty"`
	DB_Instance                          DB_Instance                          `json:"dbinstance,omitempty"`
	DB_Migration_Instance                DB_Migration_Instance                `json:"dbmigrationinstance,omitempty"`
	DB_Proxy                             DB_Proxy                             `json:"dbproxy,omitempty"`
	DB_Snapshot                          DB_Snapshot                          `json:"dbsnapshot,omitempty"`
	DDOS_Protection                      DDOS_Protection                      `json:"ddosprotection,omitempty"`
	Delivery_Stream                      Delivery_Stream                      `json:"deliverystream,omitempty"`
	Diagnostic_Setting                   Diagnostic_Setting                   `json:"diagnosticsetting,omitempty"`
	DirectConnect                        DirectConnect                        `json:"directconnect,omitempty"`
	Directory_Service                    Directory_Service                    `json:"directoryservice,omitempty"`
	Distributed_Table                    Distributed_Table                    `json:"distributedtable,omitempty"`
	Distributed_Table_Cluster            Distributed_Table_Cluster            `json:"distributedtablecluster,omitempty"`
	Divvy_Organization_Service           Divvy_Organization_Service           `json:"divvyorganizationservice,omitempty"`
	DNS_Zone                             DNS_Zone                             `json:"dnszone,omitempty"`
	Domain_Group                         Domain_Group                         `json:"domaingroup,omitempty"`
	Domain_User                          Domain_User                          `json:"domainuser,omitempty"`
	ECS_Task_Definition                  ECS_Task_Definition                  `json:"ecstaskdefinition,omitempty"`
	Email_Service_Domain                 Email_Service_Domain                 `json:"emailservicedomain,omitempty"`
	ES_Instance                          ES_Instance                          `json:"esinstance,omitempty"`
	ETL_Data_Catalog                     ETL_Data_Catalog                     `json:"etldatacatalog,omitempty"`
	ETL_Security_Config                  ETL_Security_Config                  `json:"etlsecurityconfig,omitempty"`
	Event_Subscription                   Event_Subscription                   `json:"eventsubscription,omitempty"`
	Fileshare                            Fileshare                            `json:"fileshare,omitempty"`
	Forwarding_Rule                      Forwarding_Rule                      `json:"forwardingrule,omitempty"`
	GCP_Stackdriver_Sink                 GCP_Stackdriver_Sink                 `json:"gcpstackdriversink,omitempty"`
	Global_Loadbalancer                  Global_Loadbalancer                  `json:"globalloadbalancer,omitempty"`
	Graph_Api                            Graph_Api                            `json:"graphapi,omitempty"`
	Hypervisor                           Hypervisor                           `json:"hypervisor,omitempty"`
	Identity_Provider                    Identity_Provider                    `json:"identityprovider,omitempty"`
	Instance                             Instance                             `json:"instance,omitempty"`
	Instance_Reservation                 Instance_Reservation                 `json:"instancereservation,omitempty"`
	Internet_Gateway                     Internet_Gateway                     `json:"internetgateway,omitempty"`
	K8s_Cluster_Role                     K8s_Cluster_Role                     `json:"kubernetesclusterrole,omitempty"`
	K8s_Config_Map                       K8s_Config_Map                       `json:"kubernetesconfigmap,omitempty"`
	K8s_Control_Plane                    K8s_Control_Plane                    `json:"kubernetescontrolplane,omitempty"`
	K8s_Cron_Job                         K8s_Cron_Job                         `json:"kubernetescronjob,omitempty"`
	K8s_Daemon_Set                       K8s_Daemon_Set                       `json:"kubernetesdaemonset,omitempty"`
	K8s_Ingress                          K8s_Ingress                          `json:"kubernetesingress,omitempty"`
	K8s_Job                              K8s_Job                              `json:"kubernetesjob,omitempty"`
	K8s_Mutating_Webhook_Configuration   K8s_Mutating_Webhook_Configuration   `json:"kubernetesmutatingwebhookconfiguration,omitempty"`
	K8s_Namespace                        K8s_Namespace                        `json:"kubernetesnamespace,omitempty"`
	K8s_Network_Policy                   K8s_Network_Policy                   `json:"kubernetesnetworkpolicy,omitempty"`
	K8s_Persistent_Volume                K8s_Persistent_Volume                `json:"kubernetespersistentvolume,omitempty"`
	K8s_Replica_Set                      K8s_Replica_Set                      `json:"kubernetesreplicaset,omitempty"`
	K8s_Role                             K8s_Role                             `json:"kubernetesrole,omitempty"`
	K8s_Secret                           K8s_Secret                           `json:"kubernetessecret,omitempty"`
	K8s_Service                          K8s_Service                          `json:"kubernetesservice,omitempty"`
	K8s_Service_Account                  K8s_Service_Account                  `json:"kubernetesserviceaccount,omitempty"`
	K8s_Stateful_Set                     K8s_Stateful_Set                     `json:"kubernetesstatefulset,omitempty"`
	K8s_Validating_Webhook_Configuration K8s_Validating_Webhook_Configuration `json:"kubernetesvalidatingwebhookconfiguration,omitempty"`
	Lightsail                            Lightsail                            `json:"lightsail,omitempty"`
	Loadbalancer                         Loadbalancer                         `json:"loadbalancer,omitempty"`
	Log_Analytics_Workspace              Log_Analytics_Workspace              `json:"loganalyticsworkspace,omitempty"`
	Logic_App                            Logic_App                            `json:"logicapp,omitempty"`
	MapReduce_Cluster                    MapReduce_Cluster                    `json:"mapreducecluster,omitempty"`
	MC_Database_Cluster                  MC_Database_Cluster                  `json:"mcdatabasecluster,omitempty"`
	MC_Instance                          MC_Instance                          `json:"mcinstance,omitempty"`
	MC_Snapshot                          MC_Snapshot                          `json:"mcsnapshot,omitempty"`
	Message_Queue                        Message_Queue                        `json:"messagequeue,omitempty"`
	ML_Instance                          ML_Instance                          `json:"mlinstance,omitempty"`
	NAT_Gateway                          NAT_Gateway                          `json:"natgateway,omitempty"`
	Network_Address_Group                Network_Address_Group                `json:"networkaddressgroup,omitempty"`
	Network_Endpoint                     Network_Endpoint                     `json:"networkendpoint,omitempty"`
	Network_Endpoint_Service             Network_Endpoint_Service             `json:"networkendpointservice,omitempty"`
	Network_Firewall                     Network_Firewall                     `json:"networkfirewall,omitempty"`
	Network_Firewall_Rule                Network_Firewall_Rule                `json:"networkfirewallrule,omitempty"`
	Network_Firewall_Rule_List           Network_Firewall_Rule_List           `json:"networkfirewallrulelist,omitempty"`
	Network_Flow_Log                     Network_Flow_Log                     `json:"networkflowlog,omitempty"`
	Network_Interface                    Network_Interface                    `json:"networkinterface,omitempty"`
	Network_Peer                         Network_Peer                         `json:"networkpeer,omitempty"`
	Notification_Subscription            Notification_Subscription            `json:"notificationsubscription,omitempty"`
	Notification_Topic                   Notification_Topic                   `json:"notificationtopic,omitempty"`
	Pod                                  Pod                                  `json:"pod,omitempty"`
	Pod_Security_Policy                  Pod_Security_Policy                  `json:"podsecuritypolicy,omitempty"`
	Private_Image                        Private_Image                        `json:"privateimage,omitempty"`
	Private_Network                      Private_Network                      `json:"privatenetwork,omitempty"`
	Private_Subnet                       Private_Subnet                       `json:"privatesubnet,omitempty"`
	Public_IP                            Public_IP                            `json:"publicip,omitempty"`
	Query_Log_Config                     Query_Log_Config                     `json:"querylogconfig,omitempty"`
	Recycle_Bin_Rule                     Recycle_Bin_Rule                     `json:"recyclebinrule,omitempty"`
	Resource_Access_List                 Resource_Access_List                 `json:"resourceaccesslist,omitempty"`
	Resource_Access_List_Rule            Resource_Access_List_Rule            `json:"resourceaccesslistrule,omitempty"`
	Resource_Group                       Resource_Group                       `json:"resourcegroup,omitempty"`
	REST_Api                             REST_Api                             `json:"restapi,omitempty"`
	REST_Api_Domain                      REST_Api_Domain                      `json:"restapidomain,omitempty"`
	REST_Api_Key                         REST_Api_Key                         `json:"restapikey,omitempty"`
	REST_API_Stage                       REST_API_Stage                       `json:"restapistage,omitempty"`
	Route                                Route                                `json:"route,omitempty"`
	Route_Table                          Route_Table                          `json:"routetable,omitempty"`
	Search_Cluster                       Search_Cluster                       `json:"searchcluster,omitempty"`
	Search_Index                         Search_Index                         `json:"searchindex,omitempty"`
	Secret                               Secret                               `json:"secret,omitempty"`
	Secure_File_Transfer                 Secure_File_Transfer                 `json:"securefiletransfer,omitempty"`
	Serverless_Application               Serverless_Application               `json:"serverlessapplication,omitempty"`
	Serverless_Function                  Serverless_Function                  `json:"serverlessfunction,omitempty"`
	Serverless_Layer                     Serverless_Layer                     `json:"serverlesslayer,omitempty"`
	Service_Access_Key                   Service_Access_Key                   `json:"serviceaccesskey,omitempty"`
	Service_Access_Point                 Service_Access_Point                 `json:"serviceaccesspoint,omitempty"`
	Service_Alarm                        Service_Alarm                        `json:"servicealarm,omitempty"`
	Service_App                          Service_App                          `json:"serviceapp,omitempty"`
	Service_Certificate                  Service_Certificate                  `json:"servicecertificate,omitempty"`
	Service_Check                        Service_Check                        `json:"servicecheck,omitempty"`
	Service_Control_Policy               Service_Control_Policy               `json:"servicecontrolpolicy,omitempty"`
	Service_Cost                         Service_Cost                         `json:"servicecost,omitempty"`
	Service_Dataset                      Service_Dataset                      `json:"servicedataset,omitempty"`
	Service_Detector                     Service_Detector                     `json:"servicedetector,omitempty"`
	Service_Domain                       Service_Domain                       `json:"servicedomain,omitempty"`
	Service_Encryption_Key               Service_Encryption_Key               `json:"serviceencryptionkey,omitempty"`
	Service_Encryption_Key_Vault         Service_Encryption_Key_Vault         `json:"serviceencryptionkeyvault,omitempty"`
	Service_Event_Bus                    Service_Event_Bus                    `json:"serviceeventbus,omitempty"`
	Service_Event_Rule                   Service_Event_Rule                   `json:"serviceeventrule,omitempty"`
	Service_Group                        Service_Group                        `json:"servicegroup,omitempty"`
	Service_Limit                        Service_Limit                        `json:"servicelimit,omitempty"`
	Service_Log_Group                    Service_Log_Group                    `json:"serviceloggroup,omitempty"`
	Service_Outpost                      Service_Outpost                      `json:"serviceoutpost,omitempty"`
	Service_Policy                       Service_Policy                       `json:"servicepolicy,omitempty"`
	Service_Region                       Service_Region                       `json:"serviceregion,omitempty"`
	Service_Role                         Service_Role                         `json:"servicerole,omitempty"`
	Service_User                         Service_User                         `json:"serviceuser,omitempty"`
	Shared_Filesystem                    Shared_Filesystem                    `json:"sharedfilesystem,omitempty"`
	Shared_Gallery                       Shared_Gallery                       `json:"sharedgallery,omitempty"`
	Shared_Gallery_Image                 Shared_Gallery_Image                 `json:"sharedgalleryimage,omitempty"`
	Shared_Gallery_Image_Version         Shared_Gallery_Image_Version         `json:"sharedgalleryimageversion,omitempty"`
	Site_To_Site_VPN                     Site_To_Site_VPN                     `json:"sitetositevpn,omitempty"`
	Snapshot                             Snapshot                             `json:"snapshot,omitempty"`
	Spanner                              Spanner                              `json:"spanner,omitempty"`
	Spanner_Database                     Spanner_Database                     `json:"spannerdatabase,omitempty"`
	SSH_Keypair                          SSH_Keypair                          `json:"sshkeypair,omitempty"`
	SSM_Document                         SSM_Document                         `json:"ssmdocument,omitempty"`
	Stack_Template                       Stack_Template                       `json:"stacktemplate,omitempty"`
	Step_Function                        Step_Function                        `json:"stepfunction,omitempty"`
	Storage_Account                      Storage_Account                      `json:"storageaccount,omitempty"`
	Storage_Container                    Storage_Container                    `json:"storagecontainer,omitempty"`
	Storage_Gateway                      Storage_Gateway                      `json:"storagegateway,omitempty"`
	Stored_Parameter                     Stored_Parameter                     `json:"storedparameter,omitempty"`
	Stream_Instance                      Stream_Instance                      `json:"streaminstance,omitempty"`
	Target_Proxy                         Target_Proxy                         `json:"targetproxy,omitempty"`
	Threat_Finding                       Threat_Finding                       `json:"threatfinding,omitempty"`
	Threat_Finding_Resource              Threat_Finding_Resource              `json:"threatfindingresource,omitempty"`
	Timeseries_Database                  Timeseries_Database                  `json:"timeseriesdatabase,omitempty"`
	Traffic_Mirror_Target                Traffic_Mirror_Target                `json:"trafficmirrortarget,omitempty"`
	Transcoding_Pipeline                 Transcoding_Pipeline                 `json:"transcodingpipeline,omitempty"`
	Transcription_Job                    Transcription_Job                    `json:"transcriptionjob,omitempty"`
	Transit_Gateway                      Transit_Gateway                      `json:"transitgateway,omitempty"`
	User_Pool                            User_Pool                            `json:"userpool,omitempty"`
	Video_Stream                         Video_Stream                         `json:"videostream,omitempty"`
	Virtual_Private_Gateway              Virtual_Private_Gateway              `json:"virtualprivategateway,omitempty"`
	Volume                               Volume                               `json:"volume,omitempty"`
	WAF                                  WAF                                  `json:"waf,omitempty"`
	Web_App                              Web_App                              `json:"webapp,omitempty"`
	Web_App_Group                        Web_App_Group                        `json:"webappgroup,omitempty"`
	Workspace                            Workspace                            `json:"workspace,omitempty"`
}

type Resource_Settings struct {
	Settings []interface{} `json:"setting_list"`
}

type Set_Resource_Owner_Request struct {
	Resource_IDs      []string `json:"resource_ids"`
	Owner_Resource_ID string   `json:"owner_resource_id"`
}

type Resource_Associations struct {
	Resource_Groups []Resource_Group `json:"resource_groups"`
}

// FUNCTIONS
///////////////////////////////////////////

func (c *resources) Query(q Query) (Query_Results, error) {
	// Queries InsightCloudSec for resources with the given query (using v3-ETL endpoint of the API)

	if q.Badge_Filter_Operator != "" {
		err := validateBadgeFilterOperator(q.Badge_Filter_Operator)
		if err != nil {
			return Query_Results{}, err
		}
	}

	if q.Limit == 0 {
		q.Limit = 1000
	} else {
		validateQueryLimit(q.Limit)
	}

	resp, err := c.client.makeRequest(http.MethodPost, "/v3/public/resource/etl-query", q)
	if err != nil {
		return Query_Results{}, err
	}

	var ret Query_Results
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Query_Results{}, err
	}

	return ret, nil
}

func (c *resources) GetDetails(resource_id string) (Resource_Details, error) {
	// Given a resource_id as a string, it returns the resource details and dependencies
	resp, err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/resource/%s/detail", resource_id), nil)
	if err != nil {
		return Resource_Details{}, err
	}

	var ret Resource_Details
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Resource_Details{}, err
	}

	return ret, nil
}

func (c *resources) SetOwner(resource_ids []string, owner_resource_id string) error {
	// Given a list of resource ids as strings and an owner_resource_id as string, it sets the given user as the owner of the list
	_, err := c.client.makeRequest(http.MethodPost, "/v2/public/resource/owner/set", Set_Resource_Owner_Request{Resource_IDs: resource_ids, Owner_Resource_ID: owner_resource_id})
	if err != nil {
		return err
	}

	return nil
}

func (c *resources) GetAssociations(resource_id string) (Resource_Associations, error) {
	resp, err := c.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/resource/%s/associations/get", resource_id), nil)
	if err != nil {
		return Resource_Associations{}, err
	}

	var ret Resource_Associations
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Resource_Associations{}, nil
	}

	return ret, nil
}

func (c *resources) ListTags(resource_id string) ([]Tag, error) {
	resp, err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/resource/%s/tags/list", resource_id), nil)
	if err != nil {
		return []Tag{}, err
	}

	var ret Tags_Response
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return []Tag{}, nil
	}
	return ret.Tags, nil
}

func (c *resources) List_Settings(resource_id string) (Resource_Settings, error) {
	resp, err := c.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/resource/%s/settings/list", resource_id), nil)
	if err != nil {
		return Resource_Settings{}, err
	}

	var ret Resource_Settings
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Resource_Settings{}, nil
	}
	return ret, nil
}

func validateBadgeFilterOperator(b string) error {
	// Validation Function for Query.Badge_Filter_Operator
	if strings.ToUpper(b) != "OR" && strings.ToUpper(b) != "AND" {
		return ValidationError{
			ItemToValidate: "BadgeFilterOperator",
			ExpectedValues: []string{"OR", "AND"},
		}
	}

	return nil
}

func validateQueryLimit(l int32) error {
	// Validation Function for Query.Limit
	if l < 1 || l > 1000 {
		return ValidationError{
			ItemToValidate: "Limit",
			ExpectedValues: []string{"0-1000"},
		}
	}

	return nil
}
