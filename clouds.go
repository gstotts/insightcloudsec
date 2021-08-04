package insightcloudsec

import "time"

type Cloud struct {
	ID                  int                   `json:"id"`
	Name                string                `json:"name"`
	CloudTypeID         string                `json:"cloud_type_id"`
	AccountID           string                `json:"account_id"`
	Created             time.Time             `json:"creation_time"`
	Status              string                `json:"status"`
	BadgeCount          int                   `json:"badge_count"`
	ResourceCount       int                   `json:"resource_count"`
	LastRefreshed       time.Time             `json:"last_refreshed"`
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
	Access string `json:"cloud_acces"`
}

type CloudTypesList struct {
	CloudsTypes []CloudType `json:"clouds"`
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

type HarevestingStrategyList struct {
	Strategies []HarvestingStrategy `json:"strategies"`
}
