package insightcloudsec

// STRUCTS
///////////////////////////////////////////
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

type AccessAnalyzerResource struct {
	Common                   CommonResourceValues `json:"common"`
	ID                       string               `json:"analyzer_id"`
	Mapping                  string               `json:"account_mapping"`
	ActiveFindingCount       int                  `json:"active_finding_count"`
	PublicFindingCount       int                  `json:"public_finding_count"`
	CrossAccountFindingCount int                  `json:"cross_account_count"`
	ThirdPartyFindingCount   int                  `json:"unknown_account_count"`
}

type AccessFlowLogResource struct {
	Common                   CommonResourceValues `json:"common"`
	Name                     string               `json:"name"`
	RegionName               string               `json:"region_name"`
	StorageID                string               `json:"storage_id"`
	TargetResourceID         string               `json:"target_resource_id"`
	ProvisioningState        string               `json:"provisioning_state"`
	Enabled                  bool                 `json:"enabled"`
	Retention                bool                 `json:"retention_enabled"`
	RetentionTime            string               `json:"retention_time"`
	TrafficAnalytics         bool                 `json:"traffic_analytics_enabled"`
	TrafficAnalyticsInterval string               `json:"raffic_analytics_interval"`
	Version                  string               `json:"version"`
}

type AirflowEnvironmentResource struct {
	Common              CommonResourceValues `json:"common"`
	WebserverAccessMode string               `json:"webserver_access_mode"`
	WebserverURL        string               `json:"webserver_url"`
	Status              string               `json:"status"`
	Class               string               `json:"environment_class"`
	MaxWorkers          int                  `json:"max_workers"`
	Encrypted           bool                 `json:"encrypted"`
	EncryptionKeyID     string               `json:"key_resource_id"`
	LoggingConfig       string               `json:"logging_config"`
	ExecutionRoleID     string               `json:"execution_role_resource_id"`
	ServiceRoleID       string               `json:"service_role_resource_id"`
}

type APIAccountingConfigResource struct {
	Common                     CommonResourceValues `json:"common"`
	ID                         string               `json:"accounting_config_id"`
	Name                       string               `json:"name"`
	MultiRegion                bool                 `json:"multi_region"`
	ParentResourceID           string               `json:"parent_resource_id"`
	IsLogging                  bool                 `json:"is_logging"`
	IsOrgTrail                 bool                 `json:"is_organization_trail"`
	IncludeGlobalServiceEvents bool                 `json:"include_global_service_events"`
	StorageContainerName       string               `json:"storage_container_name"`
	ManagementEvents           string               `json:"management_events,omitempty"`
	SNSTopic                   string               `json:"sns_topic_name"`
	LogGroup                   string               `json:"cloud_watch_group_arn"`
	RetentionDays              int                  `json:"retention_days"`
	Locked                     bool                 `json:"locked"`
}

type AppRunnerServiceResource struct {
	Common         CommonResourceValues `json:"commmon"`
	ID             string               `json:"service_id"`
	Status         string               `json:"status"`
	Repository     string               `json:"repository"`
	RepositoryType string               `json:"repository_type"`
	Cores          int                  `json:"cores"`
	Memory         int                  `json:"memory"`
}

type AppServerResource struct {
	Common           CommonResourceValues `json:"commmon"`
	ID               string               `json:"app_server_id"`
	State            string               `json:"state"`
	Type             string               `json:"server_type"`
	InstanceCount    int                  `json:"instance_count"`
	MaxInstanceCount int                  `json:"max_instance_count"`
	AppCount         int                  `json:"app_count"`
}

type AutoscalingGroupResource struct {
	Common                 CommonResourceValues `json:"commmon"`
	ID                     string               `json:"group_id"`
	Created                string               `json:"create_time"`
	HealthCheckGracePeriod int                  `json:"health_check_grace_period"`
	MultiAZ                bool                 `json:"multi_az"`
	MinSize                int                  `json:"min_size"`
	MaxSize                int                  `json:"max_size"`
	DesiredCapacity        int                  `json:"desired_capacity"`
	NewInstanceProtection  bool                 `json:"new_instance_protection"`
	DefaultCooldown        int                  `json:"default_cooldown"`
	UpgradePolicy          string               `json:"upgrade_policy"`
	SuspendedProcesses     []string             `json:"suspended_processes"`
}

type AutoscalingLaunchConfigResource struct {
	Common                CommonResourceValues `json:"commmon"`
	Name                  string               `json:"name"`
	ImageID               string               `json:"image_id"`
	InstanceType          string               `json:"instance_type"`
	IAMRole               string               `json:"identity_management_role"`
	Region                string               `json:"region_name"`
	Created               string               `json:"create_time"`
	Monitoring            bool                 `json:"monitoring"`
	BlockStorageOptimized bool                 `json:"block_storage_optimized"`
	AssociateIP           string               `json:"associate_ip"`
	RAMID                 string               `json:"ram_id"`
	KernelID              string               `json:"kernel_id"`
}

type AWSConfigResource struct {
	Common                      CommonResourceValues `json:"commmon"`
	ID                          string               `json:"resource_id"`
	DeliveryChannelCreated      bool                 `json:"delivery_channel_created"`
	ConfiurationRecorderCreated bool                 `json:"configuration_recorder_created"`
	AuditingBegun               bool                 `json:"auditing_has_begun"`
	AuditingEnabled             bool                 `json:"auditing_enabled"`
	CrossAccount                bool                 `json:"cross_account"`
	UnknownAccount              bool                 `json:"unknown_account"`
}

type BackendServiceResource struct {
	Common             CommonResourceValues `json:"common"`
	Kind               string               `json:"kind"`
	StorageContainerID string               `json:"storage_container_resource_id"`
	PortName           string               `json:"port_name"`
	Port               string               `json:"port"`
	Created            string               `json:"created_time"`
	Scheme             string               `json:"scheme"`
}

type BackupVaultResource struct {
	Common          CommonResourceValues `json:"common"`
	Name            string               `json:"name"`
	Created         string               `json:"create_time"`
	RecoveryPoints  int                  `json:"recovery_points"`
	Policy          string               `json:"policy"`
	Public          bool                 `json:"public"`
	EncryptionKeyID string               `json:"key_resource_id"`
}

type BatchEnvironmentResource struct {
	Common         CommonResourceValues `json:"common"`
	Name           string               `json:"name"`
	Region         string               `json:"region_name"`
	Endpoint       string               `json:"endpoint"`
	State          string               `json:"state"`
	AllocationType string               `json:"allocation_type"`
	PublicAccess   bool                 `json:"public_access"`
	MinVCPUs       int                  `json:"minimum_cpus"`
	MaxVCPUs       int                  `json:"maximum_cpus"`
	PoolType       string               `json:"pool_type"`
}

type BatchPoolResource struct {
	Common                 CommonResourceValues `json:"common"`
	Name                   string               `json:"name"`
	Region                 string               `json:"region"`
	State                  string               `json:"state"`
	InstanceSize           string               `json:"vm_size"`
	Autoscaling            string               `json:"autoscaling"`
	InterNodeCommunication string               `json:"inter_node_communication"`
}

type BigDataInstanceResource struct {
	Common          CommonResourceValues `json:"common"`
	State           string               `json:"state"`
	Type            string               `json:"instance_type"`
	EndpointAddress string               `json:"endpoint_address"`
	EndpointPort    int                  `json:"endpoint_port"`
	Version         string               `json:"version"`
	Nodes           []string             `json:"nodes"`
	VPCID           string               `json:"vpc_id"`
	SubnetGroupName string               `json:"subnet_group_name"`
	Encrypted       bool                 `json:"encrypted"`
	PublicAccess    bool                 `json:"publicly_accessible"`
	SSLRequried     bool                 `json:"ssl_required"`
	Created         string               `json:"create_time"`
	Logging         bool                 `json:"logging"`
	LoggingBucket   string               `json:"logging_bucket,omitempty"`
	ParameterGroups []string             `json:"parameter_groups"`
}

type BigDataSnapshotResource struct {
	Common       CommonResourceValues `json:"common"`
	ID           string               `json:"snapshot_id"`
	State        string               `json:"state"`
	Type         string               `json:"snapshot_type"`
	Nodes        []string             `json:"nodes"`
	InstanceType string               `json:"instance_type"`
	Encrypted    bool                 `json:"encrypted"`
	Created      string               `json:"create_time"`
}

type BigDataWorkspaceResource struct {
	Common           CommonResourceValues `json:"common"`
	Name             string               `json:"name"`
	State            string               `json:"state"`
	Region           string               `json:"region_name"`
	DoubleEncryption bool                 `json:"double_encryption_enabled"`
	SQLAdmin         string               `json:"sql_administrator_login"`
	Type             string               `json:"workspace_type"`
}

type BrokerInstanceResource struct {
	Common            CommonResourceValues `json:"common"`
	Type              string               `json:"instance_type"`
	ID                string               `json:"instance_id"`
	State             string               `json:"state"`
	Engine            string               `json:"engine"`
	EngineVersion     string               `json:"engine_version"`
	Nodes             []string             `json:"nodes"`
	EndpointAddress   string               `json:"endpoint_address"`
	PublicAccess      bool                 `json:"publicly_accessible"`
	GeneralLogging    bool                 `json:"general_logs"`
	AuditLogging      bool                 `json:"audit_logs"`
	AuthStrategy      string               `json:"authentication_strategy"`
	AutoMinorUpgrades bool                 `json:"auto_minor_upgrades"`
}

type BuildProjectResource struct {
	Common            CommonResourceValues `json:"common"`
	Description       string               `json:"description"`
	Created           string               `json:"creation_date"`
	BuildType         string               `json:"build_type"`
	BuildImage        string               `json:"build_image"`
	PrivilegeMode     string               `json:"privilege_mode"`
	CacheType         string               `json:"cache_type"`
	EncryptionKeyID   string               `json:"key_resource_id"`
	NetworkResourceID string               `json:"network_resource_id"`
}

type CloudwatchDestinationResource struct {
	Common          CommonResourceValues `json:"common"`
	Name            string               `json:"destination_name"`
	ARN             string               `json:"arn"`
	TargetARN       string               `json:"target_arn"`
	RoleARN         string               `json:"role_arn"`
	AccessPolicy    string               `json:"access_policy"`
	TrustedAccounts []string             `json:"trusted_accounts"`
	Created         string               `json:"creation_time"`
}

type ColdStorageResource struct {
	Common             CommonResourceValues `json:"common"`
	DateCreated        string               `json:"creation_date"`
	LastInventoryDate  string               `json:"last_inventory_date"`
	Size               int                  `json:"size_in_bytes"`
	ArchivesCount      int                  `json:"number_of_archives"`
	LockCreationDate   string               `json:"lock_creation_date"`
	LockExpirationDate string               `json:"lock_expiration_date"`
	LockState          string               `json:"lock_state"`
}

type ContainerResource struct {
	Common           CommonResourceValues `json:"common"`
	Name             string               `json:"name"`
	PodName          string               `json:"pod_name"`
	PodID            string               `json:"pod_resource_id"`
	Namespace        string               `json:"namespace"`
	Image            string               `json:"image"`
	ImagePullPolicy  string               `json:"image_pull_policy"`
	Repository       string               `json:"repository"`
	Version          string               `json:"version"`
	RawImageTag      string               `json:"raw_image_tag"`
	Privileged       bool                 `json:"privileged"`
	Command          string               `json:"command"`
	Args             []string             `json:"args"`
	LogDriver        string               `json:"log_driver"`
	LogGroupName     string               `json:"log_group_name"`
	LogGroupID       string               `json:"log_group_resource_id"`
	TaskDefinitionID string               `json:"task_definition_resource_id"`
	RestartCount     int                  `json:"restart_count"`
	TTY              int                  `json:"tty"`
	WorkingDirectory string               `json:"working_dir"`
}

type ContainerClusterResource struct {
	Common            CommonResourceValues `json:"common"`
	Name              string               `json:"name"`
	ARN               string               `json:"arn"`
	Type              string               `json:"res_type"`
	Created           string               `json:"created_at"`
	Endpoint          string               `json:"endpoint"`
	Fargate           bool                 `json:"fargate"`
	RoleARN           string               `json:"role_arn"`
	Region            string               `json:"region_name"`
	NetworkResourceID string               `json:"network_resource_id"`
	Status            string               `json:"status"`
	Version           string               `json:"version"`
	PlatformVersion   string               `json:"platform_version"`
	Monitoring        bool                 `json:"moniotring"`
	Logging           bool                 `json:"logging"`
	LoggingTypes      []string             `json:"logging_types,omitempty"`
	PublicAccess      bool                 `json:"endpoint_public_access"`
	PrivateAccess     bool                 `json:"endpoint_private_access"`
	PublicCIDRs       []string             `json:"public_access_cidrs,omitempty"`
	EncryptionKeyID   string               `json:"key_resource_id"`
	MasterAuthNetwork bool                 `json:"master_auth_network_enabled"`
}

type ContainerDeploymentResource struct {
	Common              CommonResourceValues `json:"common"`
	Name                string               `json:"name"`
	Namespace           string               `json:"namespace"`
	ARN                 string               `json:"arn"`
	LastStatus          string               `json:"last_status"`
	DesiredStatus       string               `json:"desired_status"`
	LaunchType          string               `json:"launch_type"`
	Connectivity        string               `json:"connectivity"`
	Created             string               `json:"create_time"`
	TaskDefinitionID    string               `json:"task_definition_resource_id"`
	Paused              bool                 `json:"paused"`
	Replicas            string               `json:"replicas,omitempty"`
	AvailableReplicas   []string             `json:"available_replicas,omitempty"`
	UnavailableReplicas []string             `json:"unavailable_replicas,omitempty"`
	ReadyReplicas       []string             `json:"ready_replicas,omitempty"`
	UpdatedReplicas     []string             `json:"updated_replicas,omitempty"`
	ObservedGeneration  string               `json:"observed_generation,omitempty"`
	CollisionCount      int                  `json:"collision_count"`
}

type ContainerImageResource struct {
	Common        CommonResourceValues `json:"common"`
	Digest        string               `json:"digest"`
	HashAlgorithm string               `json:"sha256"`
	PushTime      string               `json:"push_time"`
	LastScanned   string               `json:"last_scanned"`
	RegistryID    string               `json:"registry_id"`
	RegistryName  string               `json:"registry_name"`
	FindingCount  int                  `json:"finding_count"`
	Critical      int                  `json:"critical"`
	High          int                  `json:"high"`
	Medium        int                  `json:"medium"`
	Low           int                  `json:"low"`
	Size          int                  `json:"size"`
	Tags          map[string]string    `json:"image_tags"`
}

type ContainerInstanceResource struct {
	Common            CommonResourceValues `json:"common"`
	ID                string               `json:"instance_resource_id"`
	NetworkResourceID string               `json:"network_resource_id"`
	PodCIDR           string               `json:"pod_cidr"`
	PodCount          int                  `json:"pod_count"`
	PrivateIPAddress  string               `json:"internal_ip_address"`
	PublicIPAddress   string               `json:"external_ip_address"`
	Hostname          string               `json:"hostname"`
	Architecture      string               `json:"architecture"`
	BootID            string               `json:"boot_id"`
	RuntimeVersion    string               `json:"container_runtime_version"`
	OS                string               `json:"operating_system"`
	OSImage           string               `json:"os_image"`
	Ready             bool                 `json:"ready"`
	Unschedulable     bool                 `json:"unschedulable"`
}

type ContainerRegistryResource struct {
	Common          CommonResourceValues `json:"common"`
	Name            string               `json:"name"`
	Created         string               `json:"create_time"`
	Status          string               `json:"status,omitempty"`
	TrustedAccounts []string             `json:"trusted_accounts,omitempty"`
	LifecyclePolicy string               `json:"lifescycle_policy,omitempty"`
	ID              string               `json:"registry_id"`
	NamespaceID     string               `json:"namespace_id"`
	PublicAccess    bool                 `json:"publicly_accessible"`
	ScanOnPush      bool                 `json:"scan_on_push:"`
	ImageCount      int                  `json:"image_count"`
	EncryptionType  string               `json:"encryption_type"`
	EncryptionKeyID string               `json:"key_resource_id,omitempty"`
	TagMutability   bool                 `json:"tag_mutability"`
}

type ContainerServiceResource struct {
	Common               CommonResourceValues `json:"common"`
	Created              string               `json:"create_time"`
	RoleID               string               `json:"role_resource_id"`
	RoleName             string               `json:"role_name"`
	ClusterID            string               `json:"cluster_id"`
	TaskID               string               `json:"task_resource_id"`
	PlatformVersion      string               `json:"platform_version"`
	SchedulingStrategy   string               `json:"scheduling_strategy"`
	AssignPublicIP       bool                 `json:"assign_public_ip"`
	EnableECSTags        bool                 `json:"enable_ecs_tags"`
	EnableExecuteCommand bool                 `json:"enable_execute_command"`
	CreatedBy            string               `json:"created_by"`
	DesiredCount         int                  `json:"desired_count"`
	RunningCount         int                  `json:"running_count"`
	PendingCount         int                  `json:"pending_count"`
}

type ContentDeliveryNetworkResource struct {
	Common                 CommonResourceValues `json:"common"`
	DistributionID         string               `json:"distribution_id"`
	Domain                 string               `json:"domain_name"`
	AlternateDomains       []string             `json:"alternate_domain_names"`
	Status                 string               `json:"status"`
	DeliveryMethod         string               `json:"delivery_method"`
	State                  string               `json:"state"`
	RootObject             string               `json:"root_object"`
	HTTPVersions           []string             `json:"http_versions"`
	IPV6Enabled            bool                 `json:"ipv6_enabled"`
	LastModified           string               `json:"last_modified"`
	LogBucket              string               `json:"log_bucket"`
	Origins                []string             `json:"origins"`
	SecurityPolicy         string               `json:"security_policy"`
	Certificate            string               `json:"certificate"`
	WebACLID               string               `json:"web_acl_id"`
	PriceClass             string               `json:"price_class"`
	Comment                string               `json:"comment"`
	ARN                    string               `json:"arn"`
	Logging                bool                 `json:"logging"`
	CookieLogging          bool                 `json:"cookie_logging"`
	OriginAccessIdentities []string             `json:"origin_access_identities"`
	HTTPSRequired          bool                 `json:"https_required"`
	ViewerProtocolPolicy   string               `json:"viewer_protocol_poilcy"`
	GeoWhitelist           []string             `json:"geo_whitelist,omitempty"`
	GeoBlacklist           []string             `json:"geo_blacklist,omitempty"`
}

type DivvyOrganizationServiceResource struct {
	Common         CommonResourceValues `json:"common"`
	ID             string               `json:"resource_id"`
	OrgServiceID   string               `json:"organization_service_id"`
	AccountID      string               `json:"account_id,omitempty"`
	TenantID       string               `json:"tenant_id,omitempty"`
	PayerAccount   string               `json:"payer_acccount"`
	Name           string               `json:"name"`
	AddedTimestamp string               `json:"added_timestamp"`
	Bots           int                  `json:"bots"`
	Resources      int                  `json:"resources"`
}

type InstanceAssociation struct {
	Common InstanceAssociationDetails `json:"common"`
}

type InstanceAssociationDetails struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
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
