package insightcloudsec

import (
	"encoding/json"
	"net/http"
	"strings"
)

type CloudOrganizations interface {
	Create(cloud_type string, creds string, nickname string, auto_add bool, auto_remove bool, domain_name string, parent_folder_id []string, remove_suspended bool, skip_prefixes []string) (Cloud_Organization, error)
	List() ([]Cloud_Organization, error)
}

type corgs struct {
	client *Client
}

type Cloud_Organization_Create struct {
	// For use in creating Cloud Organizations
	Cloud_Type       string   `json:"cloud_type"` // AWS, AZURE_ARM, GCE
	Credentials      string   `json:"credentials"`
	Nickname         string   `json:"nickname"`
	Auto_Add         bool     `json:"auto_add,omitempty"`    //GCP Only
	Auto_Badge       bool     `json:"auto_badge,omitempty"`  //GCP Only
	Auto_Remove      bool     `json:"auto_remove,omitempty"` // GCP Only
	Domain_Name      string   `json:"domain_name"`
	Parent_Folder_ID []string `json:"parent_folder_id"`
	Remove_Suspended bool     `json:"remove_suspended,omitempty"` // AWS Only
	Skip_Prefixes    []string `json:"skip_prefixes,omitempty"`
}

type Cloud_Organization struct {
	Organization_ID int    `json:"organization_id"`
	Status          int    `json:"status"`
	Auto_Badge      bool   `jso:"auto_badge"`
	Auto_Add        bool   `json:"auto_add"`
	Added_Timestamp string `json:"added_timestamp"`
	Failures        int    `json:"failures"`
	Cloud_Type_ID   string `json:"cloud_type_id"`
	Domain_ID       string `json:"domain_id"`
	Projects        int    `json:"projects"`
	Domain_Name     string `json:"domain_name"`
}

type Cloud_Organizations_List struct {
	Domains []Cloud_Organization `json:"domains"`
}

// Functions
///////////////////////////////////////////

func (c *corgs) Create(cloud_type string, creds string, nickname string, auto_add bool, auto_remove bool, domain_name string, parent_folder_id []string, remove_suspended bool, skip_prefixes []string) (Cloud_Organization, error) {
	// Creates a cloud organization

	if creds == "" {
		return Cloud_Organization{}, ValidationError{
			ItemToValidate: "creds",
			ExpectedValues: []string{"Credentials are required for cloud organization creation"},
		}
	}

	if nickname == "" {
		return Cloud_Organization{}, ValidationError{
			ItemToValidate: "nickname",
			ExpectedValues: []string{"Nicknames are required for cloud organization creation"},
		}
	}

	var cloud_org Cloud_Organization_Create
	cloud_type = strings.ToUpper(cloud_type)
	if cloud_type == "AWS" {
		cloud_org = create_AWS_Cloud_Org(creds, nickname, domain_name, parent_folder_id, remove_suspended, skip_prefixes)
	} else if cloud_type == "AZURE_ARM" {
		cloud_org = create_Azure_Cloud_Org(creds, nickname, domain_name, parent_folder_id, skip_prefixes)
	} else if cloud_type == "GCE" {
		cloud_org = create_GCE_Cloud_Org(creds, nickname, auto_add, auto_remove, domain_name, parent_folder_id, skip_prefixes)
	} else {
		return Cloud_Organization{}, ValidationError{
			ItemToValidate: "cloud_type",
			ExpectedValues: []string{"AWS", "AZURE_ARM", "GCE"},
		}
	}

	resp, err := c.client.makeRequest(http.MethodPost, "/v2/public/cloud/domain/add", cloud_org)
	if err != nil {
		return Cloud_Organization{}, err
	}

	var ret Cloud_Organization
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Cloud_Organization{}, err
	}

	return ret, nil
}

func create_AWS_Cloud_Org(creds string, nickname string, domain_name string, parent_folder_id []string, remove_suspended bool, skip_prefixes []string) Cloud_Organization_Create {
	// Creates AWS Cloud Organization
	return Cloud_Organization_Create{
		Cloud_Type:       "AWS",
		Credentials:      creds,
		Nickname:         nickname,
		Domain_Name:      domain_name,
		Parent_Folder_ID: parent_folder_id,
		Remove_Suspended: remove_suspended,
		Skip_Prefixes:    skip_prefixes,
	}
}

func create_Azure_Cloud_Org(creds string, nickname string, domain_name string, parent_folder_id []string, skip_prefixes []string) Cloud_Organization_Create {
	return Cloud_Organization_Create{
		Cloud_Type:       "AZURE_ARM",
		Credentials:      creds,
		Nickname:         nickname,
		Domain_Name:      domain_name,
		Parent_Folder_ID: parent_folder_id,
		Skip_Prefixes:    skip_prefixes,
	}
}

func create_GCE_Cloud_Org(creds string, nickname string, auto_add bool, auto_remove bool, domain_name string, parent_folder_id []string, skip_prefixes []string) Cloud_Organization_Create {
	return Cloud_Organization_Create{
		Cloud_Type:       "GCE",
		Credentials:      creds,
		Nickname:         nickname,
		Auto_Add:         auto_add,
		Auto_Badge:       false,
		Auto_Remove:      auto_remove,
		Domain_Name:      domain_name,
		Parent_Folder_ID: parent_folder_id,
		Skip_Prefixes:    skip_prefixes,
	}
}

func (c *corgs) List() ([]Cloud_Organization, error) {
	// Returns a list of cloud organizations
	resp, err := c.client.makeRequest(http.MethodGet, "/v2/public/cloud/domains", nil)
	if err != nil {
		return []Cloud_Organization{}, err
	}

	var ret Cloud_Organizations_List
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return []Cloud_Organization{}, err
	}

	return ret.Domains, nil
}

func (c Client) Delete() error {
	return nil
}

func (c Client) Edit() error {
	return nil
}
