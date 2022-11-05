package insightcloudsec

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthServers_List(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/authenticationservers/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("authServers/list.json"))
	})

	resp, err := &client.AuthenticationServers.List()
	want := AuthenticationServersList{
		Servers: []AuthenticationServer{
			{
				ID:           1,
				Name:         "Okta SSO",
				Host:         "",
				Port:         0,
				Secure:       1,
				Type:         "saml",
				GlobalScope:  false,
				MappedGroups: 0,
			},
			{
				ID:           2,
				Name:         "Azure AD",
				Host:         "",
				Port:         0,
				Secure:       1,
				Type:         "saml",
				GlobalScope:  true,
				MappedGroups: 0,
			},
			{
				ID:           3,
				Name:         "Rapid7 Okta",
				Host:         "",
				Port:         0,
				Secure:       1,
				Type:         "saml",
				GlobalScope:  false,
				MappedGroups: 1,
			},
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, resp)
	}

	teardown()
}
