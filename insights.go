package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	INSIGHT_SEVERITY_CRITICAL = 5
	INSIGHT_SEVERITY_SEVERE   = 4
	INSIGHT_SEVERITY_MAJOR    = 3
	INSIGHT_SEVERITY_MODERATE = 2
	INSIGHT_SEVERITY_MINOR    = 1
)

var _ Insights = (*insights)(nil)

type Insights interface {
	Create(i Insight) error
	Edit(i Insight) (*Insight, error)
	Delete(insight_id int) error
	Get_Insight(insight_id int, insight_source string) (*Insight, error)
	Get_Insight_7_Days(insight_id int, insight_source string) (map[string]int, error)
	List() ([]Insight, error)
	List_Packs() ([]InsightPack, error)
}

type insights struct {
	client *Client
}

type Insight struct {
	ID                  int             `json:"insight_id,omitempty"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	TemplateID          int             `json:"template_id"`
	OrgID               int             `json:"organization_id,omitempty"`
	OwnerResourceID     string          `json:"owner_resource_id,omitempty"`
	Severity            int             `json:"severity"`
	Scopes              []string        `json:"scopes"`
	Tags                []string        `json:"tags"`
	ResourceTypes       []string        `json:"resource_types"`
	Filters             []InsightFilter `json:"filters"`
	Timeseries          bool            `json:"timeseries,omitempty"`
	TimeseriesCache     int             `json:"timeseries_cache,omitempty"`
	Badges              []Badge         `json:"badges,omitempty"`
	BadgeFilterOperator string          `json:"badge_filter_operator,omitempty"`
}

type InsightFilter struct {
	Name        string                 `json:"name"`
	Config      map[string]interface{} `json:"config"`
	Collections map[string]interface{} `json:"collections"`
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

func (c *insights) Create(i Insight) (*Insight, error) {
	// Creates an Insight in InsightCloudSec given the insight object with appropriate configs.  Returns an error if insight creation fails.

	// Make sure severity is set
	if i.Severity == 0 {
		return nil, fmt.Errorf("[-] ERROR: Insight Severity must be set")
	}

	// Make sure Filters are set
	if i.Filters == nil {
		return nil, fmt.Errorf("[-] ERROR: Insight filters must be set")
	}

	// Clean up any empty config and collection fields for filters so they return empty object in json
	for idx, filter := range i.Filters {
		if filter.Config == nil {
			i.Filters[idx].Config = make(map[string]interface{})
		}
		if filter.Collections == nil {
			i.Filters[idx].Collections = make(map[string]interface{})
		}
	}

	// Clean up any empty scopes so they return an empty object in json
	if i.Scopes == nil {
		i.Scopes = make([]string, 0)
	}

	data, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("[-] ERROR: Marshal error: %s", err)
	}

	resp, err := c.client.makeRequest(http.MethodPost, "/v2/public/insights/create", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var ret Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *insights) Edit(i Insight) (*Insight, error) {
	current_data, err := c.client.Insights.Get_Insight(i.ID, "custom")
	if err != nil {
		return nil, err
	}

	// Cleanup any missing required fields for update using existing fields
	if i.Name == "" {
		i.Name = current_data.Name
	}
	if i.Description == "" {
		i.Description = current_data.Description
	}
	if i.ResourceTypes == nil {
		i.ResourceTypes = current_data.ResourceTypes
	}
	if i.OwnerResourceID == "" {
		i.OwnerResourceID = current_data.OwnerResourceID
	}
	if i.TemplateID == 0 {
		i.TemplateID = current_data.TemplateID
	}
	if i.Severity == 0 {
		i.Severity = current_data.Severity
	}
	if i.Filters == nil {
		i.Filters = current_data.Filters
	}

	// Clean up any empty config and collection fields for filters so they return empty object in json
	for idx, filter := range i.Filters {
		if filter.Config == nil {
			i.Filters[idx].Config = make(map[string]interface{})
		}
		if filter.Collections == nil {
			i.Filters[idx].Collections = make(map[string]interface{})
		}
	}

	// Clean up any empty scopes so they return an empty object in json
	if i.Scopes == nil {
		i.Scopes = make([]string, 0)
	}

	data, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("[-] ERROR: Marshal error: %s", err)
	}

	resp, err := c.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/insights/%d/edit", i.ID), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var ret Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *insights) List() ([]Insight, error) {
	// Returns a list of all Insights from the API
	resp, err := c.client.makeRequest(http.MethodGet, "/v2/public/insights/list", nil)
	if err != nil {
		return nil, err
	}

	var ret []Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *insights) Get_Insight(insight_id int, insight_source string) (*Insight, error) {
	// Returns the specific Insight associated with the Insight ID and the Source provided
	resp, err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/insights/%d/%s", insight_id, insight_source), nil)
	if err != nil {
		return nil, err
	}

	var ret Insight
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *insights) Delete(insight_id int) error {
	// Deletes the Insight for the given id.  Returns an error if fails.
	resp, err := c.client.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/public/insights/%d/delete", insight_id), nil)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *insights) Get_Insight_7_Days(insight_id int, insight_source string) (map[string]int, error) {
	// Returns the 7 Day View of Insight associated with the Insight ID and the Source provided
	resp, err := c.client.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/insights/%d/%s/insight-data-7-days", insight_id, insight_source), nil)
	if err != nil {
		return nil, err
	}

	var ret map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *insights) List_Packs() ([]InsightPack, error) {
	// Returns a list of all Insight Packs from the API
	resp, err := c.client.makeRequest(http.MethodGet, "/v2/public/insights/packs/list", nil)
	if err != nil {
		return nil, err
	}

	var ret []InsightPack
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
