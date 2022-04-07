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
		fmt.Fprintf(w, `
		{
			"servers": [
			  {
				"server_id": 1,
				"server_name": "Okta SSO",
				"server_host": "",
				"server_port": 0,
				"secure": 1,
				"server_type": "saml",
				"global_scope": false,
				"mapped_groups": 0
			  },
			  {
				"server_id": 2,
				"server_name": "Azure AD",
				"server_host": "",
				"server_port": 0,
				"secure": 1,
				"server_type": "saml",
				"global_scope": true,
				"mapped_groups": 0
			  },
			  {
				"server_id": 3,
				"server_name": "Rapid7 Okta",
				"server_host": "",
				"server_port": 0,
				"secure": 1,
				"server_type": "saml",
				"global_scope": null,
				"mapped_groups": 1
			  }
			]
		  }`)
	})

	resp, err := client.AuthenticationServers.List()
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
