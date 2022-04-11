package insightcloudsec

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_List(t *testing.T) {
	setup()

	teardown()
}

func TestUsers_Create(t *testing.T) {
	setup()

	teardown()
}

func TestUsers_CreateAPIUser(t *testing.T) {
	setup()

	teardown()
}

func TestUsers_CreateSAMLUser(t *testing.T) {
	setup()

	teardown()
}

func TestUsers_Delete(t *testing.T) {
	setup()

	teardown()
}

func TestUsers_DeleteByUsername(t *testing.T) {
	setup()

	teardown()
}

func TestUsers_CurrentUserInfo(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/info", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/currentUserInfoResponse.json"))
	})

	resp, err := client.Users.CurrentUserInfo()
	assert.NoError(t, err)
	assert.Equal(t, "divvyuser:6:", resp.ResourceID)
	assert.Equal(t, "Lord Farquad", resp.Name)
	assert.Equal(t, 6, resp.ID)
	assert.Equal(t, false, resp.OrgAdmin)
	assert.Equal(t, true, resp.DomainAdmin)
	assert.Equal(t, false, resp.DomainViewer)
	assert.Equal(t, "farquad@dulock.org", resp.Email)
	assert.Equal(t, "farquad", resp.Username)
	assert.Equal(t, "Dulock", resp.Org)
	assert.Equal(t, 3, resp.OrgID)
	assert.Equal(t, true, resp.TwoFactorEnabled)
	assert.Equal(t, false, resp.TwoFactorRequired)
	assert.Equal(t, false, resp.AuthPluginExists)
	assert.Equal(t, []string{}, resp.NavigationBlacklist)
	assert.Equal(t, false, resp.RequirePWReset)
	assert.Equal(t, "2021-10-11 11:38:38", resp.Created)
	teardown()
}

func TestUsers_Get2FAStatus(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/tfa_state", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/get2FAStatus.json"))
	})

	resp, err := client.Users.Get2FAStatus(2)
	assert.NoError(t, err)
	assert.Equal(t, true, resp.Enabled)
	assert.Equal(t, false, resp.Required)

	teardown()
}
