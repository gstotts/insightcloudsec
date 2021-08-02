package insightcloudsec

import (
	"encoding/json"
	"fmt"
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
	// Returns a list of all Insights from the API
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


func (c *Client) GetInsight(insight_id interface{}, insight_source string) (*Insight, error) {
	// Returns the specific Insight associated with the Insight ID and the Source provided
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/insights/%s/%s", insight_id.(string), insight_source), nil)
	if err != nil {
		return nil, err
	}

	var ret Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
