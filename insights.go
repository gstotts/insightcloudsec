package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////
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

type BackofficeMetadata struct {
	PackID       int    `json:"pack_id"`
	PackName     string `json:"pack_name"`
	TemplateID   int    `json:"template_id"`
	TemplateName string `json:"template_name"`
	Description  string `json:"description"`
	Order        int    `json:"order"`
}

type InsightPack struct {
	ID                  int                  `json:"pack_id"`
	OrgID               int                  `json:"organization_id"`
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	Source              string               `json:"source"`
	LogoURL             string               `json:"logo_url"`
	InsertedAt          string               `json:"inserted_at"`
	UpdatedAt           string               `json:"updated_at"`
	Backoffice          []int                `json:"backoffice"`
	Backoffice_Metadata []BackofficeMetadata `json:"backoffice_metadata"`
	Custom              []int                `json:"custom"`
}

// INSIGHT FUNCTIONS
///////////////////////////////////////////

func (c *Client) List_Insights() ([]Insight, error) {
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

func (c *Client) Get_Insight(insight_id int, insight_source string) (*Insight, error) {
	// Returns the specific Insight associated with the Insight ID and the Source provided
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/insights/%d/%s", insight_id, insight_source), nil)
	if err != nil {
		return nil, err
	}

	var ret Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) Get_Insight_7_Days(insight_id int, insight_source string) (map[string]int, error) {
	// Returns the 7 Day View of Insight associated with the Insight ID and the Source provided
	resp, err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/insights/%d/%s/insight-data-7-days", insight_id, insight_source), nil)
	if err != nil {
		return nil, err
	}

	var ret map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// PACK FUNCTIONS
///////////////////////////////////////////

func (c *Client) List_Packs() ([]InsightPack, error) {
	// Returns a list of all Insight Packs from the API
	resp, err := c.makeRequest(http.MethodGet, "/v2/public/insights/packs/list", nil)
	if err != nil {
		return nil, err
	}

	var ret []InsightPack
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// FILTER FUNCTIONS
///////////////////////////////////////////
