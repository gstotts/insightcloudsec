package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var _ Iac = (*iac)(nil)

type Iac interface {
	Get_Configs() ([]IACConfigs, error)
	Scan(data IACScan, readable bool) (interface{}, error)
	Get_Scan_Results(id int32, readable bool) (interface{}, error)
}

type iac struct {
	client *Client
}

type IACConfigs struct {
	ID               int      `json:"id"`
	OrganizationID   int      `json:"organization_id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	PackID           int      `json:"pack_id"`
	Source           string   `json:"source"`
	InsightBlacklist []string `json:"insights_blacklist"`
	InsightsWarnOnly []string `json:"insights_warn_only"`
	SlackChannel     string   `json:"slack_channel"`
	EmailRecipients  []string `json:"email_recipients"`
	LastBuildAt      string   `json:"last_build_at"`
	CreatedAt        string   `json:"created_at"`
	LastModified     string   `json:"last_modified"`
	TotalBuilds      int      `json:"total_builds"`
	SuccessCount     int      `json:"success_count"`
	FailureCount     int      `json:"failure_count"`
	ConsumptionURL   string   `json:"consumption_url"`
	Favorite         bool     `json:"favorite"`
}

type IACScan struct {
	Provider     string `json:"iac_provider"`
	AuthorName   string `json:"author_name"`
	ConfigName   string `json:"config_name"`
	ScanName     string `json:"scan_name"`
	ScanTemplate string `json:"scan_template"`
}

type IACScanResponse struct {
	BuildID         int32                  `json:"build_id"`
	ConfigName      string                 `json:"config_name"`
	Details         IACScanResponseDetails `json:"details"`
	Errors          []string               `json:"errors"`
	Message         string                 `json:"message"`
	ResourceMapping IACResourceMapping     `json:"resource_mapping"`
	ScanResults     string                 `json:"scan_results"`
	Stacktrace      string                 `json:"stacktrace"`
	Status          string                 `json:"status"`
	Success         bool                   `json:"success"`
}

type IACScanResponseDetails struct {
	FailedInsights   []IACInsightDetails `json:"failed_insights"`
	FailedResources  int                 `json:"failed_resources"`
	PassedInsights   []IACInsightDetails `json:"passed_insights"`
	PassedResources  int                 `json:"passed_resources"`
	SkippedInsights  []IACInsightDetails `json:"skipped_insigfhts"`
	SkippedResources int                 `json:"skipped_resources"`
	TotalInsights    int                 `json:"total_insights"`
	TotalResources   int                 `json:"total_resources"`
	WarnedInsights   []IACInsightDetails `json:"warned_insights"`
	WarnedResources  int                 `json:"warned_resources"`
}

type IACResourceMapping struct{}

type IACInsightDetails struct {
	Description   string   `json:"Description"`
	Failure       []string `json:"failure"`
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Notes         string   `json:"notes"`
	ResourceTypes []string `json:"resource_types"`
	Severity      int      `json:"severity"`
	Source        string   `json:"source"`
	Success       []string `json:"success"`
	WarnOnly      bool     `json:"warn_only"`
	Warning       []string `json:"warning"`
}

type IACScanResponseReadable struct {
	BuildID     int                      `json:"build_id"`
	ConfigName  string                   `json:"config_name"`
	Message     string                   `json:"message"`
	ScanResults string                   `json:"scan_results"`
	Success     bool                     `json:"success"`
	Status      string                   `json:"status"`
	Errors      []string                 `json:"errors"`
	Resources   []IACScanResourceDetails `json:"resources"`
}

type IACScanResourceDetails struct {
	Source       string  `json:"resource_source"`
	Address      string  `json:"resource_address"`
	Name         string  `json:"resource_name"`
	Type         string  `json:"resource_type"`
	Region       string  `json:"resource_region"`
	PassedRules  []Rules `json:"passed_rules"`
	WarningRules []Rules `json:"warning_rules"`
	FailedRules  []Rules `json:"failed_rules"`
	Stacktrace   string  `json:"stacktrace"`
}

type Rules struct {
	Name        string `json:"rule_name"`
	Description string `json:"rule_description"`
	Detail      string `json:"rule_detail"`
}

type IACResults struct {
}

type IACResultsReadable struct {
}

func (c *iac) Get_Configs() ([]IACConfigs, error) {
	// Returns a list of all IAC Configs from the API
	resp, err := c.client.makeRequest(http.MethodGet, "/v3/iac/configs", nil)
	if err != nil {
		return nil, err
	}

	var ret []IACConfigs
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *iac) Scan(data IACScan, readable bool) (interface{}, error) {
	// Initiates a scan and returns results
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("[-] ERROR: Marshal error: %s", err)
	}

	resp, err := c.client.makeRequest(http.MethodPost, fmt.Sprintf("/v3/iac/scan?readable=%t", readable), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	if readable {
		var ret IACScanResponseReadable
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return nil, err
		}

		return ret, nil
	} else {
		var ret IACScanResponseDetails
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (c *iac) Get_Scan_Results(id int32, readable bool) (interface{}, error) {
	// Returns scan results for given build id

	return nil, nil
}
