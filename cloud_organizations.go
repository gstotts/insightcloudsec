package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

var _ CloudOrgs = (*cloudorgs)(nil)

type CloudOrgs interface {
	Add(CloudOrg) (CloudOrg, error)
	List() ([]Domains, error)
}

type cloudorgs struct {
	client *Client
}

type CloudOrg struct {
}

type Domains struct {
	OrgID          int    `json:"organization_id"`
	Status         int    `json:"status"`
	AutoBadge      bool   `json:"auto_badge"`
	AutoAdd        bool   `json:"auto_add"`
	AddedTimestamp string `json:"added_timestamp"`
	Failures       int    `json:"failures"`
	CloudTypeID    string `json:"cloud_type_id"`
	DomainID       string `json:"domain_id"`
	Projects       int    `json:"projects"`
	DomainName     string `json:"domain_name"`
}

func (o *cloudorgs) Add(c CloudOrg) (CloudOrg, error) {
	return CloudOrg{}, nil
}

func (o *cloudorgs) List() ([]Domains, error) {
	// Lists all cloud organizations enrolled in InsightCloudSec
	resp, err := o.client.makeRequest(http.MethodGet, "/v2/public/cloud/domains", nil)
	if err != nil {
		return nil, err
	}

	var ret []Domains
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
