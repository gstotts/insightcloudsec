package insightcloudsec

import (
	"encoding/json"
	"net/http"
)

var _ AuthenticationServers = (*authServers)(nil)

type AuthenticationServers interface {
	List() (AuthenticationServersList, error)
}

type authServers struct {
	client *Client
}

type AuthenticationServer struct {
	ID           int    `json:"server_id"`
	Name         string `json:"server_name"`
	Host         string `json:"server_host"`
	Port         int    `json:"server_port"`
	Secure       int    `json:"secure"`
	Type         string `json:"server_type"`
	GlobalScope  bool   `json:"global_scope"`
	MappedGroups int    `json:"mapped_groups"`
}

type AuthenticationServersList struct {
	Servers []AuthenticationServer `json:"servers"`
}

func (s *authServers) List() (*AuthenticationServersList, error) {
	body, err := s.client.makeRequest(http.MethodPost, "/v2/prototype/authenticationservers/list", nil, nil)
	if err != nil {
		return AuthenticationServersList{}, err
	}

	var ret AuthenticationServersList
	if err := json.Unmarshal(body, &ret); err != nil {
		return AuthenticationServersList{}, err
	}

	return ret, nil
}
