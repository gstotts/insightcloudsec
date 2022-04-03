package insightcloudsec

// STRUCTS
///////////////////////////////////////////
type Common_Resource_Values struct {
	//Common Attributes
	ID                   string            `json:"resource_id"`
	Name                 string            `json:"resource_name"`
	Type                 string            `json:"type"`
	Cloud                string            `json:"cloud"`
	Account              string            `json:"account"`
	Account_ID           string            `json:"account_id"`
	Account_Status       string            `json:"account_status"`
	Org_Service_ID       int               `json:"organization_service_id"`
	Availability_Zone    string            `json:"availablility_zone"`
	Region               string            `json:"region"`
	Creation_Timestamp   string            `json:"creation_timestamp"`
	Discovered_Timestamp string            `json:"discovered_timestamp"`
	Modified_Timestamp   string            `json:"modified_timestamp"`
	Namespace_ID         string            `json:"namespace_id"`
	Tags                 map[string]string `json:"tags"`
}

type Access_Analyzer struct {
	Common                Common_Resource_Values `json:"common"`
	Analyzer_ID           string                 `json:"analyzer_id"`
	Account_Mapping       string                 `json:"account_mapping"`
	Active_Finding_Count  int                    `json:"active_finding_count"`
	Public_Finding_Count  int                    `json:"public_finding_count"`
	Cross_Account_Count   int                    `json:"cross_account_count"`
	Unknown_Account_Count int                    `json:"unknown_account_count"`
}

type Access_List_Flow_Log struct {
	Common                     Common_Resource_Values `json:"common"`
	Name                       string                 `json:"name"`
	Region_Name                string                 `json:"region_name"`
	Storage_ID                 string                 `json:"storage_id"`
	Target_Resource_ID         string                 `json:"target_resource_id"`
	Provisioning_State         string                 `json:"provisioning_state"`
	Enabled                    bool                   `json:"enabled"`
	Retention                  bool                   `json:"retention_enabled"`
	Retention_Time             string                 `json:"retention_time"`
	Traffic_Analytics_Enabled  bool                   `json:"traffic_analytics_enabled"`
	Traffic_Analytics_Interval string                 `json:"traffic_analytics_interval"`
	Version                    string                 `json:"version"`
}

type Airflow_Environment struct {
	Common                     Common_Resource_Values `json:"common"`
	Webserver_Access_Mode      string                 `json:"webserver_access_mode"`
	Webserver_URL              string                 `json:"webserver_url"`
	Status                     string                 `json:"status"`
	Environment_Class          string                 `json:"environment_class"`
	Max_Workers                int                    `json:"max_workers"`
	Encrypted                  bool                   `json:"encrypted"`
	Key_Resource_ID            string                 `json:"key_resource_id"`
	Logging_Config             string                 `json:"logging_config"`
	Execution_Role_Resource_ID string                 `json:"execution_role_resource_id"`
	Service_Role_Resource_ID   string                 `json:"service_role_resource_id"`
}

type API_Accounting_Config struct {
	Common                        Common_Resource_Values `json:"common"`
	Accounting_Config_ID          string                 `json:"accounting_config_id"`
	Name                          string                 `json:"name"`
	Multi_Region                  bool                   `json:"multi_region"`
	Parent_Resource_ID            string                 `json:"parent_resource_id"`
	Is_Logging                    bool                   `json:"is_logging"`
	Is_Organization_Trail         bool                   `json:"is_organization_trail"`
	Include_Global_Service_Events bool                   `json:"include_global_service_events"`
	Storage_Container_Name        string                 `json:"storage_container_name"`
	Management_Events             string                 `json:"management_events,omitempty"`
	SNS_Topic                     string                 `json:"sns_topic_name"`
	CloudWatch_Group_ARN          string                 `json:"cloud_watch_group_arn"`
	Retention_Days                int                    `json:"retention_days"`
	Locked                        bool                   `json:"locked"`
}

type App_Runner_Service struct {
	Common          Common_Resource_Values `json:"common"`
	Service_ID      string                 `json:"service_id"`
	Status          string                 `json:"status"`
	Repository      string                 `json:"repository"`
	Repository_Type string                 `json:"repository_type"`
	Cores           int                    `json:"cores"`
	Memory          int                    `json:"memory"`
}

type App_Server struct {
	Common             Common_Resource_Values `json:"common"`
	App_Server_ID      string                 `json:"app_server_id"`
	State              string                 `json:"state"`
	Server_Type        string                 `json:"server_type"`
	Instance_Count     int                    `json:"instance_count"`
	Max_Instance_Count int                    `json:"max_instance_count"`
	App_Count          int                    `json:"app_count"`
}

type Autoscaling_Group struct {
	Common                    Common_Resource_Values `json:"common"`
	Group_ID                  string                 `json:"group_id"`
	Create_Time               string                 `json:"create_time"`
	Health_Check_Grace_Period int                    `json:"health_check_grace_period"`
	Multi_AZ                  bool                   `json:"multi_az"`
	Min_Size                  int                    `json:"min_size"`
	Max_Size                  int                    `json:"max_size"`
	Desired_Capacity          int                    `json:"desired_capacity"`
	New_Instance_Protection   bool                   `json:"new_instance_protection"`
	Default_Cooldown          int                    `json:"default_cooldown"`
	Upgrade_Policy            string                 `json:"upgrade_policy"`
	Suspended_Processes       []string               `json:"suspended_processes"`
}

type Autoscaling_Launch_Configuration struct {
	Common                  Common_Resource_Values `json:"common"`
	Name                    string                 `json:"name"`
	Image_ID                string                 `json:"image_id"`
	Instance_Type           string                 `json:"instance_type"`
	IAM_Role                string                 `json:"identity_management_role"`
	Region_Name             string                 `json:"region_name"`
	Create_Time             string                 `json:"create_time"`
	Monitoring              bool                   `json:"monitoring"`
	Block_Storage_Optimized bool                   `json:"block_storage_optimized"`
	Associate_IP            string                 `json:"associate_ip"`
	RAM_ID                  string                 `json:"ram_id"`
	Kernel_ID               string                 `json:"kernel_id"`
}

type AWS_Config struct {
	Common                        Common_Resource_Values `json:"common"`
	Resource_ID                   string                 `json:"resource_id"`
	Delivery_Channel_Created      bool                   `json:"delivery_channel_created"`
	Confiuration_Recorder_Created bool                   `json:"configuration_recorder_created"`
	Auditing_Has_Begun            bool                   `json:"auditing_has_begun"`
	Auditing_Enabled              bool                   `json:"auditing_enabled"`
	Cross_Account                 bool                   `json:"cross_account"`
	Unknown_Account               bool                   `json:"unknown_account"`
}

type Backend_Service struct {
	Common               Common_Resource_Values `json:"common"`
	Kind                 string                 `json:"kind"`
	Storage_Container_ID string                 `json:"storage_container_resource_id"`
	PortName             string                 `json:"port_name"`
	Port                 string                 `json:"port"`
	Created_Time         string                 `json:"created_time"`
	Scheme               string                 `json:"scheme"`
}

type Backup_Vault struct {
	Common          Common_Resource_Values `json:"common"`
	Name            string                 `json:"name"`
	Create_Time     string                 `json:"create_time"`
	Recovery_Points int                    `json:"recovery_points"`
	Policy          string                 `json:"policy"`
	Public          bool                   `json:"public"`
	Key_Resource_ID string                 `json:"key_resource_id"`
}

type Batch_Environment struct {
	Common          Common_Resource_Values `json:"common"`
	Name            string                 `json:"name"`
	Region          string                 `json:"region_name"`
	Endpoint        string                 `json:"endpoint"`
	State           string                 `json:"state"`
	Allocation_Type string                 `json:"allocation_type"`
	Public_Access   bool                   `json:"public_access"`
	Min_CPUs        int                    `json:"minimum_cpus"`
	Max_CPUs        int                    `json:"maximum_cpus"`
	Pool_Type       string                 `json:"pool_type"`
}

type Batch_Pool struct {
	Common                   Common_Resource_Values `json:"common"`
	Name                     string                 `json:"name"`
	Region                   string                 `json:"region"`
	State                    string                 `json:"state"`
	VM_Size                  string                 `json:"vm_size"`
	Autoscaling              string                 `json:"autoscaling"`
	Inter_Node_Communication string                 `json:"inter_node_communication"`
}

type BigData_Instance struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type BigData_Snapshot struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type BigData_Workspace struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Broker_Instance struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Build_Project struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type CloudWatch_Destination struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Cold_Storage struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Cluster struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Deployment struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Image struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Instance struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Node_Group struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Registry struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Service struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Container_Delivery_Network struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Data_Analytics_Workspace struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Database struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Databricks_Workspace struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Data_Factory struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Data_Lake_Storage struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

type Datastore struct {
	Common Common_Resource_Values `json:"common,omitempty"`
}

// "datastream",
// "datasynctask",
// "dbcluster",
// "dbinstance",
// "dbmigrationinstance",
// "dbproxy",
// "dbsnapshot",
// "ddosprotection",
// "deliverystream",
// "diagnosticsetting",
// "directconnect",
// "directoryservice",
// "distributedtable",
// "distributedtablecluster",
// "divvyorganizationservice",
// "dnszone",
// "domaingroup",
// "domainuser",
// "ecstaskdefinition",
// "emailservicedomain",
// "esinstance",
// "etldatacatalog",
// "etlsecurityconfig",
// "eventsubscription",
// "fileshare",
// "forwardingrule",
// "gcpstackdriversink",
// "globalloadbalancer",
// "graphapi",
// "hypervisor",
// "identityprovider",
// "instance",
// "instancereservation",
// "internetgateway",
// "kubernetesclusterrole",
// "kubernetesconfigmap",
// "kubernetescontrolplane",
// "kubernetescronjob",
// "kubernetesdaemonset",
// "kubernetesingress",
// "kubernetesjob",
// "kubernetesmutatingwebhookconfiguration",
// "kubernetesnamespace",
// "kubernetesnetworkpolicy",
// "kubernetespersistentvolume",
// "kubernetesreplicaset",
// "kubernetesrole",
// "kubernetessecret",
// "kubernetesservice",
// "kubernetesserviceaccount",
// "kubernetesstatefulset",
// "kubernetesvalidatingwebhookconfiguration",
// "lightsail",
// "loadbalancer",
// "loganalyticsworkspace",
// "logicapp",
// "mapreducecluster",
// "mcdatabasecluster",
// "mcinstance",
// "mcsnapshot",
// "messagequeue",
// "mlinstance",
// "natgateway",
// "networkaddressgroup",
// "networkendpoint",
// "networkendpointservice",
// "networkfirewall",
// "networkfirewallrule",
// "networkfirewallrulelist",
// "networkflowlog",
// "networkinterface",
// "networkpeer",
// "notificationsubscription",
// "notificationtopic",
// "pod",
// "podsecuritypolicy",
// "privateimage",
// "privatenetwork",

// "publicip",
// "querylogconfig",
// "recyclebinrule",
// "resourceaccesslist",
// "resourceaccesslistrule",
// "resourcegroup",
// "restapi",
// "restapidomain",
// "restapikey",
// "restapistage",
// "route",
// "routetable",
// "searchcluster",
// "searchindex",
// "secret",
// "securefiletransfer",
// "serverlessapplication",
// "serverlessfunction",
// "serverlesslayer",
// "serviceaccesskey",
// "serviceaccesspoint",
// "servicealarm",
// "serviceapp",
// "servicecertificate",
// "servicecheck",
// "servicecontrolpolicy",
// "servicecost",
// "servicedataset",
// "servicedetector",
// "servicedomain",
// "serviceencryptionkey",
// "serviceencryptionkeyvault",
// "serviceeventbus",
// "serviceeventrule",
// "servicegroup",
// "servicelimit",
// "serviceloggroup",
// "serviceoutpost",
// "servicepolicy",
// "serviceregion",
// "servicerole",
// "serviceuser",
// "sharedfilesystem",
// "sharedgallery",
// "sharedgalleryimage",
// "sharedgalleryimageversion",
// "sitetositevpn",
// "snapshot",
// "spanner",
// "spannerdatabase",
// "sshkeypair",
// "ssmdocument",
// "stacktemplate",
// "stepfunction",
// "storageaccount",
// "storagecontainer",
// "storagegateway",
// "storedparameter",
// "streaminstance",
// "targetproxy",
// "threatfinding",
// "threatfindingresource",
// "timeseriesdatabase",
// "trafficmirrortarget",
// "transcodingpipeline",
// "transcriptionjob",
// "transitgateway",
// "userpool",
// "videostream",
// "virtualprivategateway",
// "volume",
// "waf",
// "webapp",
// "webappgroup",
// "workspace"

type Private_Subnet struct {
	Common                  Common_Resource_Values `json:"common"`
	Subnet_ID               string                 `json:"subnet_id"`
	Availability_Zone       string                 `json:"availability_zone"`
	CIDR                    string                 `json:"cidr"`
	Available_IPs           int                    `json:"available_ips"`
	Network_Resource_ID     string                 `json:"network_resource_id"`
	Network_ID              string                 `json:"network_id"`
	Route_Table_Resource_ID string                 `json:"route_table_resource_id"`
	Public                  bool                   `json:"public"`
}
