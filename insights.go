package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

type Insight struct {
	ID              int                      `json:"insight_id"`
	Name            string                   `json:"name"`
	Description     string                   `json:"description"`
	TemplateID      int                      `json:"template_id"`
	OrgID           int                      `json:"organization_id"`
	Severity        int                      `json:"severity"`
	Scopes          []string                 `json:"scopes"`
	Tags            []string                 `json:"tags"`
	ResourceTypes   []string                 `json:"resource_types"`
	Filters         []map[string]interface{} `json:"filters"`
	Timeseries      bool                     `json:"timeseries"`
	TimeseriesCache int                      `json:"timeseries_cache"`
}

func (c *Client) ListInsights() ([]Insight, error) {
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/insights/list", nil)
	if err != nil {
		return nil, err
	}

	var ret []Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
