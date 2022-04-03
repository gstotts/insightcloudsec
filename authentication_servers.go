package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

// CONSTANTS
///////////////////////////////////////////

// STRUCTS
///////////////////////////////////////////
type Authentication_Server struct {
	ID           int    `json:"server_id"`
	Name         string `json:"server_name"`
	Host         string `json:"server_host"`
	Port         int    `json:"server_port"`
	Secure       int    `json:"secure"`
	Type         string `json:"server_type"`
	GlobalScope  bool   `json:"global_scope"`
	MappedGroups int    `json:"mapped_groups"`
}

type Authentication_Servers struct {
	Servers []Authentication_Server `json:"servers"`
}

// AUTH SERVER FUNCTIONS
///////////////////////////////////////////

func (c Client) List_Authentication_Servers() (Authentication_Servers, error) {
	resp, err := c.makeRequest(http.MethodPost, "/v2/prototype/authenticationservers/list", nil)
	if err != nil {
		return Authentication_Servers{}, err
	}

	var ret Authentication_Servers
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Authentication_Servers{}, err
	}

	return ret, nil
}
