package insightcloudsec

// STRUCTS
// /////////////////////////////////////////
type Common_Resource_Values struct {
	//Common Attributes
	Resource_ID             string            `json:"resource_id"`
	Resource_Name           string            `json:"resource_name"`
	Type                    string            `json:"resource_type"`
	Cloud                   string            `json:"cloud"`
	Account                 string            `json:"account"`
	Account_ID              string            `json:"account_id"`
	Account_Status          string            `json:"account_status"`
	Organization_Service_ID int               `json:"organization_service_id"`
	Availability_Zone       string            `json:"availability_zone"`
	Region                  string            `json:"region"`
	Creation_Timestamp      string            `json:"creation_timestamp"`
	Discovered_Timestamp    string            `json:"discovered_timestamp"`
	Modified_Timestamp      string            `json:"modified_timestamp"`
	Namespace_ID            string            `json:"namespace_id"`
	Tags                    map[string]string `json:"tags"`
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
	Common              Common_Resource_Values `json:"common"`
	State               string                 `json:"state"`
	Instance_Type       string                 `json:"instance_type"`
	Endpoint_Address    string                 `json:"endpoint_address"`
	Endpoint_Port       int                    `json:"endpoint_port"`
	Version             string                 `json:"version"`
	Nodes               []string               `json:"nodes"`
	VPC_ID              string                 `json:"vpc_id"`
	Subnet_Group_Name   string                 `json:"subnet_group_name"`
	Encrypted           bool                   `json:"encrypted"`
	Publicly_Accessible bool                   `json:"publicly_accessible"`
	SSL_Requried        bool                   `json:"ssl_required"`
	Create_Time         string                 `json:"create_time"`
	Logging             bool                   `json:"logging"`
	Logging_Bucket      string                 `json:"logging_bucket,omitempty"`
	Parameter_Groups    []string               `json:"parameter_groups"`
}

type BigData_Snapshot struct {
	Common        Common_Resource_Values `json:"common"`
	Snapshot_ID   string                 `json:"snapshot_id"`
	State         string                 `json:"state"`
	Snapshot_Type string                 `json:"snapshot_type"`
	Nodes         []string               `json:"nodes"`
	Instance_Type string                 `json:"instance_type"`
	Encrypted     bool                   `json:"encrypted"`
	Create_Time   string                 `json:"create_time"`
}

type BigData_Workspace struct {
	Common                    Common_Resource_Values `json:"common"`
	Name                      string                 `json:"name"`
	State                     string                 `json:"state"`
	Region_Name               string                 `json:"region_name"`
	Double_Encryption_Enabled bool                   `json:"double_encryption_enabled"`
	SQL_Administrator_Login   string                 `json:"sql_administrator_login"`
	Workspace_Type            string                 `json:"workspace_type"`
}

type Broker_Instance struct {
	Common                  Common_Resource_Values `json:"common"`
	Instance_Type           string                 `json:"instance_type"`
	Instance_ID             string                 `json:"instance_id"`
	State                   string                 `json:"state"`
	Engine                  string                 `json:"engine"`
	Engine_Version          string                 `json:"engine_version"`
	Nodes                   []string               `json:"nodes"`
	Endpoint_Address        string                 `json:"endpoint_address"`
	Publicly_Accessible     bool                   `json:"publicly_accessible"`
	General_Logs            bool                   `json:"general_logs"`
	Audit_Logs              bool                   `json:"audit_logs"`
	Authentication_Strategy string                 `json:"authentication_strategy"`
	Auto_Minor_Upgrades     bool                   `json:"auto_minor_upgrades"`
}

type Build_Project struct {
	Common              Common_Resource_Values `json:"common"`
	Description         string                 `json:"description"`
	Creation_Date       string                 `json:"creation_date"`
	Build_Type          string                 `json:"build_type"`
	Build_Image         string                 `json:"build_image"`
	Privilege_Mode      string                 `json:"privilege_mode"`
	Cache_Type          string                 `json:"cache_type"`
	Key_Resource_ID     string                 `json:"key_resource_id"`
	Network_Resource_ID string                 `json:"network_resource_id"`
}

type CloudWatch_Destination struct {
	Common           Common_Resource_Values `json:"common"`
	Destination_Name string                 `json:"destination_name"`
	ARN              string                 `json:"arn"`
	Target_ARN       string                 `json:"target_arn"`
	Role_ARN         string                 `json:"role_arn"`
	Access_Policy    string                 `json:"access_policy"`
	Trusted_Accounts []string               `json:"trusted_accounts"`
	Creation_Time    string                 `json:"creation_time"`
}

type Cold_Storage struct {
	Common               Common_Resource_Values `json:"common"`
	Creation_Date        string                 `json:"creation_date"`
	Last_Inventory_Date  string                 `json:"last_inventory_date"`
	Size_In_Bytes        int                    `json:"size_in_bytes"`
	Number_Of_Archives   int                    `json:"number_of_archives"`
	Lock_Creation_Date   string                 `json:"lock_creation_date"`
	Lock_Expiration_Date string                 `json:"lock_expiration_date"`
	Lock_State           string                 `json:"lock_state"`
}

type Container struct {
	Common                      Common_Resource_Values `json:"common"`
	Name                        string                 `json:"name"`
	Pod_Name                    string                 `json:"pod_name"`
	Pod_Resource_ID             string                 `json:"pod_resource_id"`
	Namespace                   string                 `json:"namespace"`
	Image                       string                 `json:"image"`
	Image_Pull_Policy           string                 `json:"image_pull_policy"`
	Repository                  string                 `json:"repository"`
	Version                     string                 `json:"version"`
	Raw_Image_Tag               string                 `json:"raw_image_tag"`
	Privileged                  bool                   `json:"privileged"`
	Command                     string                 `json:"command"`
	Args                        []string               `json:"args"`
	Log_Driver                  string                 `json:"log_driver"`
	Log_Group_Name              string                 `json:"log_group_name"`
	Log_Group_ID                string                 `json:"log_group_resource_id"`
	Task_Definition_Resource_ID string                 `json:"task_definition_resource_id"`
	Restart_Count               int                    `json:"restart_count"`
	TTY                         int                    `json:"tty"`
	Working_Dir                 string                 `json:"working_dir"`
}

type Container_Cluster struct {
	Common                      Common_Resource_Values `json:"common"`
	Name                        string                 `json:"name"`
	ARN                         string                 `json:"arn"`
	Res_Type                    string                 `json:"res_type"`
	Created_At                  string                 `json:"created_at"`
	Endpoint                    string                 `json:"endpoint"`
	Fargate                     bool                   `json:"fargate"`
	Role_ARN                    string                 `json:"role_arn"`
	Region_Name                 string                 `json:"region_name"`
	Network_Resource_ID         string                 `json:"network_resource_id"`
	Status                      string                 `json:"status"`
	Version                     string                 `json:"version"`
	Platform_Version            string                 `json:"platform_version"`
	Monitoring                  bool                   `json:"moniotring"`
	Logging                     bool                   `json:"logging"`
	Logging_Types               []string               `json:"logging_types,omitempty"`
	Endpoint_Public_Access      bool                   `json:"endpoint_public_access"`
	Endpoint_Private_Access     bool                   `json:"endpoint_private_access"`
	Public_Access_CIDRs         []string               `json:"public_access_cidrs,omitempty"`
	Key_Resource_ID             string                 `json:"key_resource_id"`
	Master_Auth_Network_Enabled bool                   `json:"master_auth_network_enabled"`
}

type Container_Deployment struct {
	Common                      Common_Resource_Values `json:"common"`
	Name                        string                 `json:"name"`
	Namespace                   string                 `json:"namespace"`
	ARN                         string                 `json:"arn"`
	Last_Status                 string                 `json:"last_status"`
	Desired_Status              string                 `json:"desired_status"`
	Launch_Type                 string                 `json:"launch_type"`
	Connectivity                string                 `json:"connectivity"`
	Create_Time                 string                 `json:"create_time"`
	Task_Definition_Resource_ID string                 `json:"task_definition_resource_id"`
	Paused                      bool                   `json:"paused"`
	Replicas                    string                 `json:"replicas,omitempty"`
	Available_Replicas          []string               `json:"available_replicas,omitempty"`
	Unavailable_Replicas        []string               `json:"unavailable_replicas,omitempty"`
	Ready_Replicas              []string               `json:"ready_replicas,omitempty"`
	Updated_Replicas            []string               `json:"updated_replicas,omitempty"`
	Observed_Generation         string                 `json:"observed_generation,omitempty"`
	Collision_Count             int                    `json:"collision_count"`
}

type Container_Image struct {
	Common        Common_Resource_Values `json:"common"`
	Digest        string                 `json:"digest"`
	SHA256        string                 `json:"sha256"`
	Push_Time     string                 `json:"push_time"`
	Last_Scanned  string                 `json:"last_scanned"`
	Registry_ID   string                 `json:"registry_id"`
	Registry_Name string                 `json:"registry_name"`
	Finding_Count int                    `json:"finding_count"`
	Critical      int                    `json:"critical"`
	High          int                    `json:"high"`
	Medium        int                    `json:"medium"`
	Low           int                    `json:"low"`
	Size          int                    `json:"size"`
	Image_Tags    map[string]string      `json:"image_tags"`
}

type Container_Instance struct {
	Common                    Common_Resource_Values `json:"common"`
	Instance_Resource_ID      string                 `json:"instance_resource_id"`
	Network_Resource_ID       string                 `json:"network_resource_id"`
	Pod_CIDR                  string                 `json:"pod_cidr"`
	Pod_Count                 int                    `json:"pod_count"`
	Internal_IP_Address       string                 `json:"internal_ip_address"`
	External_IP_Address       string                 `json:"external_ip_address"`
	Hostname                  string                 `json:"hostname"`
	Architecture              string                 `json:"architecture"`
	Boot_ID                   string                 `json:"boot_id"`
	Container_Runtime_Version string                 `json:"container_runtime_version"`
	OS                        string                 `json:"operating_system"`
	OS_Image                  string                 `json:"os_image"`
	Ready                     bool                   `json:"ready"`
	Unschedulable             bool                   `json:"unschedulable"`
}

type Container_Node_Group struct {
	Common Common_Resource_Values `json:"common"`
}

type Container_Registry struct {
	Common              Common_Resource_Values `json:"common"`
	Name                string                 `json:"name"`
	Create_Time         string                 `json:"create_time"`
	Status              string                 `json:"status,omitempty"`
	Trusted_Accounts    []string               `json:"trusted_accounts,omitempty"`
	Lifecycle_Policy    string                 `json:"lifescycle_policy,omitempty"`
	Registry_ID         string                 `json:"registry_id"`
	Namespace_ID        string                 `json:"namespace_id"`
	Publicly_Accessible bool                   `json:"publicly_accessible"`
	Scan_On_Push        bool                   `json:"scan_on_push:"`
	Image_Count         int                    `json:"image_count"`
	Encryption_Type     string                 `json:"encryption_type"`
	Key_Resource_ID     string                 `json:"key_resource_id,omitempty"`
	Tag_Mutability      bool                   `json:"tag_mutability"`
}

type Container_Service struct {
	Common                 Common_Resource_Values `json:"common"`
	Create_Time            string                 `json:"create_time"`
	Role_Resource_ID       string                 `json:"role_resource_id"`
	Role_Name              string                 `json:"role_name"`
	Cluster_ID             string                 `json:"cluster_id"`
	Task_Resource_ID       string                 `json:"task_resource_id"`
	Platform_Version       string                 `json:"platform_version"`
	Scheduling_Strategy    string                 `json:"scheduling_strategy"`
	Assign_Public_IP       bool                   `json:"assign_public_ip"`
	Enable_ECS_Tags        bool                   `json:"enable_ecs_tags"`
	Enable_Execute_Command bool                   `json:"enable_execute_command"`
	Created_By             string                 `json:"created_by"`
	Desired_Count          int                    `json:"desired_count"`
	Running_Count          int                    `json:"running_count"`
	Pending_Count          int                    `json:"pending_count"`
}

type Content_Delivery_Network struct {
	Common                   Common_Resource_Values `json:"common"`
	Distribution_ID          string                 `json:"distribution_id"`
	Domain_Name              string                 `json:"domain_name"`
	Alternate_Domain_Names   []string               `json:"alternate_domain_names"`
	Status                   string                 `json:"status"`
	Delivery_Method          string                 `json:"delivery_method"`
	State                    string                 `json:"state"`
	Root_Object              string                 `json:"root_object"`
	HTTP_Versions            []string               `json:"http_versions"`
	IPV6_Enabled             bool                   `json:"ipv6_enabled"`
	Last_Modified            string                 `json:"last_modified"`
	Log_Bucket               string                 `json:"log_bucket"`
	Origins                  []string               `json:"origins"`
	Security_Policy          string                 `json:"security_policy"`
	Certificate              string                 `json:"certificate"`
	Web_ACL_ID               string                 `json:"web_acl_id"`
	Price_Class              string                 `json:"price_class"`
	Comment                  string                 `json:"comment"`
	ARN                      string                 `json:"arn"`
	Logging                  bool                   `json:"logging"`
	Cookie_Logging           bool                   `json:"cookie_logging"`
	Origin_Access_Identities []string               `json:"origin_access_identities"`
	HTTPS_Required           bool                   `json:"https_required"`
	Viewer_Protocol_Policy   string                 `json:"viewer_protocol_poilcy"`
	Geo_Whitelist            []string               `json:"geo_whitelist,omitempty"`
	Geo_Blacklist            []string               `json:"geo_blacklist,omitempty"`
}

type Data_Analytics_Workspace struct {
	Common          Common_Resource_Values `json:"common"`
	Workspace_ID    string                 `json:"workspace_id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	State           string                 `json:"state"`
	Encrypted       bool                   `json:"encrypted"`
	Key_Resource_ID string                 `json:"key_resource_id,omitempty"`
	Output_Location string                 `json:"output_location,omitempty"`
	Create_Time     string                 `json:"create_time"`
	Requester_Pays  bool                   `json:"requester_pays"`
	Metrics_Enabled bool                   `json:"metrics_enabled"`
}

type Database struct {
	Common               Common_Resource_Values `json:"common"`
	Database_Type        string                 `json:"database_type"`
	Encrypted            bool                   `json:"encrypted"`
	Character_Set        string                 `json:"character_set"`
	Collation            bool                   `json:"collation"`
	Instance_Resource_ID string                 `json:"instance_resource_id"`
}

type Databrick_Workspace struct {
	Common          Common_Resource_Values `json:"common"`
	State           string                 `json:"state"`
	Encryption_Type string                 `json:"encryption_type"`
}

type Data_Factory struct {
	Common Common_Resource_Values `json:"common"`
}

type Data_Lake_Storage struct {
	Common Common_Resource_Values `json:"common"`
}

type Datastore struct {
	Common Common_Resource_Values `json:"common"`
}

type Datastream struct {
	Common Common_Resource_Values `json:"common"`
}

type Datasync_Task struct {
	Common Common_Resource_Values `json:"common"`
}

type DB_Cluster struct {
	Common Common_Resource_Values `json:"common"`
}

type DB_Instance struct {
	Common Common_Resource_Values `json:"common"`
}

type DB_Migration_Instance struct {
	Common Common_Resource_Values `json:"common"`
}

type DB_Proxy struct {
	Common Common_Resource_Values `json:"common"`
}

type DB_Snapshot struct {
	Common Common_Resource_Values `json:"common"`
}

type DDOS_Protection struct {
	Common Common_Resource_Values `json:"common"`
}

type Delivery_Stream struct {
	Common Common_Resource_Values `json:"common"`
}

type Diagnostic_Setting struct {
	Common Common_Resource_Values `json:"common"`
}

type DirectConnect struct {
	Common Common_Resource_Values `json:"common"`
}

type Directory_Service struct {
	Common Common_Resource_Values `json:"common"`
}

type Distributed_Table struct {
	Common Common_Resource_Values `json:"common"`
}

type Distributed_Table_Cluster struct {
	Common Common_Resource_Values `json:"common"`
}

type Divvy_Organization_Service struct {
	Common                  Common_Resource_Values `json:"common"`
	Resource_ID             string                 `json:"resource_id"`
	Organization_Service_ID string                 `json:"organization_service_id"`
	Account_ID              string                 `json:"account_id,omitempty"`
	Tenant_ID               string                 `json:"tenant_id,omitempty"`
	Payer_Account           string                 `json:"payer_acccount"`
	Name                    string                 `json:"name"`
	Added_Timestamp         string                 `json:"added_timestamp"`
	Bots                    int                    `json:"bots"`
	Resources               int                    `json:"resources"`
}

type DNS_Zone struct {
	Common     Common_Resource_Values `json:"common"`
	Domain     string                 `json:"domain"`
	Comment    string                 `json:"comment,omitempty"`
	Is_Private bool                   `json:"is_private_zone"`
	Records    int                    `json:"records"`
}

type Domain_Group struct {
	Common Common_Resource_Values `json:"common"`
}

type Domain_User struct {
	Common Common_Resource_Values `json:"common"`
}

type ECS_Task_Definition struct {
	Common             Common_Resource_Values `json:"common"`
	Name               string                 `json:"name"`
	Version            string                 `json:"version"`
	ARN                string                 `json:"arn"`
	Status             string                 `json:"status"`
	Network_Mode       string                 `json:"network_mode"`
	Launch_Type        string                 `json:"launch_type"`
	Execution_Role_ARN string                 `json:"execution_role_arn"`
	CPU                string                 `json:"cpu"`
	Memory             string                 `json:"memory"`
	Family             string                 `json:"family"`
	Created_At         string                 `json:"created_at"`
	Container_Count    int                    `json:"container_count"`
}

type Email_Service_Domain struct {
	Common Common_Resource_Values `json:"common"`
}

type ES_Instance struct {
	Common Common_Resource_Values `json:"common"`
}

type ETL_Data_Catalog struct {
	Common Common_Resource_Values `json:"common"`
}

type ETL_Security_Config struct {
	Common Common_Resource_Values `json:"common"`
}

type Event_Subscription struct {
	Common Common_Resource_Values `json:"common"`
}

type Fileshare struct {
	Common Common_Resource_Values `json:"common"`
}

type Forwarding_Rule struct {
	Common Common_Resource_Values `json:"common"`
}

type GCP_Stackdriver_Sink struct {
	Common Common_Resource_Values `json:"common"`
}

type Global_Loadbalancer struct {
	Common Common_Resource_Values `json:"common"`
}

type Graph_Api struct {
	Common Common_Resource_Values `json:"common"`
}

type Hypervisor struct {
	Common Common_Resource_Values `json:"common"`
}

type Identity_Provider struct {
	Common Common_Resource_Values `json:"common"`
}

type InstanceAssociation struct {
	Common InstanceAssociationDetails `json:"common"`
}

type InstanceAssociationDetails struct {
	Resource_ID   string `json:"resource_id"`
	Resource_Name string `json:"resource_name"`
}

type Instance struct {
	Common                                               Common_Resource_Values `json:"common"`
	Instance_ID                                          string                 `json:"instance_id"`
	Instance_Type                                        string                 `json:"instance_type"`
	Launch_Time                                          string                 `json:"launch_time"`
	Platform                                             string                 `json:"platform"`
	State                                                string                 `json:"state"`
	Image_ID                                             string                 `json:"image_id"`
	Public_IP_Address                                    string                 `json:"public_ip_address,omitempty"`
	Private_IP_Address                                   string                 `json:"private_ip_address"`
	Network_Resource_ID                                  string                 `json:"network_resource_id"`
	Subnet_Resource_ID                                   string                 `json:"subnet_resource_id"`
	Object_ID                                            string                 `json:"object_id"`
	Key_Name                                             string                 `json:"key_name,omitempty"`
	Role_Name                                            string                 `json:"role_name,omitempty"`
	Role_Resource_ID                                     string                 `json:"role_resource_id"`
	Termination_Protection                               string                 `json:"termination_protection,omitempty"`
	VM_Extensions                                        string                 `json:"vm_extensions,omitempty"`
	JIT_Access_Policy                                    string                 `json:"jit_access_policy"`
	Root_Device_Type                                     string                 `json:"root_device_type,omitempty"`
	Tenancy                                              string                 `json:"tenancy,omitempty"`
	Detailed_Monitoring                                  bool                   `json:"detailed_monitoring,omitempty"`
	Secondary_Private_IP_Addresses                       []string               `json:"secondary_private_ip_addresses,omitempty"`
	Secondary_Public_IP_Addresses                        []string               `json:"secondary_public_ip_addresses,omitempty"`
	Architecture                                         string                 `json:"architecture,omitempty"`
	AWS_Instance_Metadata_Service_V2_Required            bool                   `json:"aws_instance_metadata_service_v2_required,omitempty"`
	AWS_Instance_Metadata_Service_Hop_Limit              int                    `json:"aws_instance_metadata_service_hop_limit,omitempty"`
	AWS_Instance_Metadata_Service_Endpoint_Enabled       string                 `json:"aws_instance_metadata_service_endpoint_enabled,omitempty"`
	AWS_Instance_Metadata_Service_Endpoint_Config_Status string                 `json:"aws_instance_metadata_service_endpoint_config_status,omitempty"`
	State_Transition_Reason                              string                 `json:"state_transition_reason,omitempty"`
}

type Instance_Reservation struct {
	Common Common_Resource_Values `json:"common"`
}

type Internet_Gateway struct {
	Common Common_Resource_Values `json:"common"`
}

type IPAddress struct {
	IPAddress string `json:"ip_address"`
	Type      string `json:"ip_address_type"`
}

type K8s_Cluster_Role struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Config_Map struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Control_Plane struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Cron_Job struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Daemon_Set struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Ingress struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Job struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Mutating_Webhook_Configuration struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Namespace struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Network_Policy struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Persistent_Volume struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Replica_Set struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Role struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Secret struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Service struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Service_Account struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Stateful_Set struct {
	Common Common_Resource_Values `json:"common"`
}

type K8s_Validating_Webhook_Configuration struct {
	Common Common_Resource_Values `json:"common"`
}

type Lightsail struct {
	Common Common_Resource_Values `json:"common"`
}

type Loadbalancer struct {
	Common Common_Resource_Values `json:"common"`
}

type Log_Analytics_Workspace struct {
	Common Common_Resource_Values `json:"common"`
}

type Logic_App struct {
	Common Common_Resource_Values `json:"common"`
}

type MapReduce_Cluster struct {
	Common Common_Resource_Values `json:"common"`
}

type MC_Database_Cluster struct {
	Common Common_Resource_Values `json:"common"`
}

type MC_Instance struct {
	Common Common_Resource_Values `json:"common"`
}

type MC_Snapshot struct {
	Common Common_Resource_Values `json:"common"`
}

type Message_Queue struct {
	Common Common_Resource_Values `json:"common"`
}

type ML_Instance struct {
	Common Common_Resource_Values `json:"common"`
}

type NAT_Gateway struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Address_Group struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Endpoint struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Endpoint_Service struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Firewall struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Firewall_Rule struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Firewall_Rule_List struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Flow_Log struct {
	Common Common_Resource_Values `json:"common"`
}

type Network_Interface struct {
	Common             Common_Resource_Values `json:"common"`
	Description        string                 `json:"description"`
	DeviceIndex        int                    `json:"device_index"`
	InstanceID         string                 `json:"instance_id"`
	InstanceResourceID string                 `json:"instance_resource_id"`
	IPAddresses        []IPAddress            `json:"ip_addresses"`
	MacAddress         string                 `json:"mac_address"`
	NetworkInterfaceID string                 `json:"network_interface_id"`
	NetworkResourceID  string                 `json:"network_resource_id"`
	SubnetID           string                 `json:"subnet_id"`
	SubnetResourceID   string                 `json:"subnet_resource_id"`
}

type Network_Peer struct {
	Common Common_Resource_Values `json:"common"`
}

type Notification_Subscription struct {
	Common Common_Resource_Values `json:"common"`
}

type Notification_Topic struct {
	Common Common_Resource_Values `json:"common"`
}

type Pod struct {
	Common Common_Resource_Values `json:"common"`
}

type Pod_Security_Policy struct {
	Common Common_Resource_Values `json:"common"`
}

type Private_Image struct {
	Common Common_Resource_Values `json:"common"`
}

type Private_Network struct {
	Common Common_Resource_Values `json:"common"`
}

type Private_Subnet struct {
	Common                  Common_Resource_Values `json:"common,omitempty"`
	Subnet_ID               string                 `json:"subnet_id"`
	Availability_Zone       string                 `json:"availability_zone"`
	CIDR                    string                 `json:"cidr"`
	Available_IPs           int                    `json:"available_ips"`
	Network_Resource_ID     string                 `json:"network_resource_id"`
	Network_ID              string                 `json:"network_id"`
	Route_Table_Resource_ID string                 `json:"route_table_resource_id"`
	Public                  bool                   `json:"public"`
}

type Public_IP struct {
	Common Common_Resource_Values `json:"common"`
}

type Query_Log_Config struct {
	Common Common_Resource_Values `json:"common"`
}

type Recycle_Bin_Rule struct {
	Common Common_Resource_Values `json:"common"`
}

type Resource_Access_List struct {
	Common Common_Resource_Values `json:"common"`
}

type Resource_Access_List_Rule struct {
	Common Common_Resource_Values `json:"common"`
}

type Resource_Group struct {
	Common            Common_Resource_Values `json:"common,omitempty"`
	Resource_ID       string                 `json:"resource_id"`
	ID                string                 `json:"id,omitempty"`
	Name              string                 `json:"name"`
	Group_Type        string                 `json:"group_type"`
	Owner_Type        string                 `json:"owner_type"`
	Resource_Count    int                    `json:"resource_count"`
	Owner_Resource_ID string                 `json:"owner_resource_id,omitempty"`
}

type REST_Api struct {
	Common Common_Resource_Values `json:"common"`
}

type REST_Api_Domain struct {
	Common Common_Resource_Values `json:"common"`
}

type REST_Api_Key struct {
	Common Common_Resource_Values `json:"common"`
}

type REST_API_Stage struct {
	Common Common_Resource_Values `json:"common"`
}

type Route struct {
	Common Common_Resource_Values `json:"common"`
}

type Route_Table struct {
	Common Common_Resource_Values `json:"common"`
}

type Search_Cluster struct {
	Common Common_Resource_Values `json:"common"`
}

type Search_Index struct {
	Common Common_Resource_Values `json:"common"`
}

type Secret struct {
	Common              Common_Resource_Values `json:"common"`
	ARN                 string                 `json:"arn"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description,omitempty"`
	Key_Resource_ID     string                 `json:"key_resource_id,omitempty"`
	Region_Name         string                 `json:"region_name"`
	Rotation_Enabled    bool                   `json:"rotation_enabled"`
	Rotation_Days       int                    `json:"rotation_days,omitempty"`
	Rotation_Lambda_ARN string                 `json:"rotation_lambda_arn"`
	Last_Accessed_Date  string                 `json:"last_accessed_date"`
	Last_Changed_Date   string                 `json:"last_changed_date"`
	Creation_Date       string                 `json:"creation_date,omitempty"`
	Activation_Date     string                 `json:"activation_date,omitempty"`
	Deleted_Date        string                 `json:"deleted_date,omitempty"`
	Expiration_Date     string                 `json:"expiration_date,omitempty"`
	Enabled             bool                   `json:"enabled,omitempty"`
	Customer_Managed    bool                   `json:"customer_managed,omitempty"`
	Content_Type        string                 `json:"content_type,omitempty"`
	Parent_Vault        string                 `json:"parent_vault,omitempty"`
}

type Secure_File_Transfer struct {
	Common Common_Resource_Values `json:"common"`
}

type Serverless_Application struct {
	Common Common_Resource_Values `json:"common"`
}

type Serverless_Function struct {
	Common Common_Resource_Values `json:"common"`
}

type Serverless_Layer struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Access_Key struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Access_Point struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Alarm struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_App struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Certificate struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Check struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Control_Policy struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Cost struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Dataset struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Detector struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Domain struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Encryption_Key struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Encryption_Key_Vault struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Event_Bus struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Event_Rule struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Group struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Limit struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Log_Group struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Outpost struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Policy struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Region struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_Role struct {
	Common Common_Resource_Values `json:"common"`
}

type Service_User struct {
	Common Common_Resource_Values `json:"common"`
}

type Shared_Filesystem struct {
	Common Common_Resource_Values `json:"common"`
}

type Shared_Gallery struct {
	Common Common_Resource_Values `json:"common"`
}

type Shared_Gallery_Image struct {
	Common Common_Resource_Values `json:"common"`
}

type Shared_Gallery_Image_Version struct {
	Common Common_Resource_Values `json:"common"`
}

type Site_To_Site_VPN struct {
	Common Common_Resource_Values `json:"common"`
}

type Snapshot struct {
	Common Common_Resource_Values `json:"common"`
}

type Spanner struct {
	Common Common_Resource_Values `json:"common"`
}

type Spanner_Database struct {
	Common Common_Resource_Values `json:"common"`
}

type SSH_Keypair struct {
	Common Common_Resource_Values `json:"common"`
}

type SSM_Document struct {
	Common Common_Resource_Values `json:"common"`
}

type Stack_Template struct {
	Common Common_Resource_Values `json:"common"`
}

type Step_Function struct {
	Common Common_Resource_Values `json:"common"`
}

type Storage_Account struct {
	Common Common_Resource_Values `json:"common"`
}

type Storage_Container struct {
	Common Common_Resource_Values `json:"common"`
}

type Storage_Gateway struct {
	Common Common_Resource_Values `json:"common"`
}

type Stored_Parameter struct {
	Common Common_Resource_Values `json:"common"`
}

type Stream_Instance struct {
	Common Common_Resource_Values `json:"common"`
}

type Target_Proxy struct {
	Common Common_Resource_Values `json:"common"`
}

type Threat_Finding struct {
	Common Common_Resource_Values `json:"common"`
}

type Threat_Finding_Resource struct {
	Common Common_Resource_Values `json:"common"`
}

type Timeseries_Database struct {
	Common Common_Resource_Values `json:"common"`
}

type Traffic_Mirror_Target struct {
	Common Common_Resource_Values `json:"common"`
}

type Transcoding_Pipeline struct {
	Common Common_Resource_Values `json:"common"`
}

type Transcription_Job struct {
	Common Common_Resource_Values `json:"common"`
}

type Transit_Gateway struct {
	Common Common_Resource_Values `json:"common"`
}

type User_Pool struct {
	Common Common_Resource_Values `json:"common"`
}

type Video_Stream struct {
	Common Common_Resource_Values `json:"common"`
}

type Virtual_Private_Gateway struct {
	Common Common_Resource_Values `json:"common"`
}

type Volume struct {
	Common                Common_Resource_Values `json:"common"`
	Volume_ID             string                 `json:"volume_id"`
	Volume_Type           string                 `json:"volume_type"`
	Size                  int                    `json:"size"`
	State                 string                 `json:"state"`
	Rated_IOPS            string                 `json:"rated_iops,omitempty"`
	Encrypted             bool                   `json:"encrypted"`
	Delete_On_Termination bool                   `json:"delete_on_termination"`
	Attach_State          string                 `json:"attach_state"`
	Attach_Device_Name    string                 `json:"attach_device_name,omitempty"`
	Instance_Association  InstanceAssociation    `json:"instance_association"`
	Creation_Time         string                 `json:"creation_time"`
}

type WAF struct {
	Common Common_Resource_Values `json:"common"`
}

type Web_App struct {
	Common Common_Resource_Values `json:"common"`
}

type Web_App_Group struct {
	Common Common_Resource_Values `json:"common"`
}

type Workspace struct {
	Common Common_Resource_Values `json:"common"`
}
