package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

var _ Iac = (*iac)(nil)

type Iac interface {
	Get_Configs() ([]IACConfigs, error)
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
