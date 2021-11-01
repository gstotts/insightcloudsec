package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

// CONSTANTS
///////////////////////////////////////////

// STRUCTS
///////////////////////////////////////////
type Server struct {
	ID           int    `json:"server_id"`
	Name         string `json:"server_name"`
	Host         string `json:"server_host"`
	Port         int    `json:"server_port"`
	Secure       int    `json:"secure"`
	Type         string `json:"server_type"`
	GlobalScope  bool   `json:"global_scope"`
	MappedGroups int    `json:"mapped_groups"`
}

type Servers struct {
	Servers []Server `json:"servers"`
}

// AUTH SERVER FUNCTIONS
///////////////////////////////////////////

func (c Client) ListAuthenticationServers() (Servers, error) {
	resp, err := c.makeRequest(http.MethodPost, "/v2/prototype/authenticationservers/list", nil)
	if err != nil {
		return Servers{}, err
	}

	var ret Servers
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Servers{}, err
	}

	return ret, nil
}
