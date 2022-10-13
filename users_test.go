package insightcloudsec

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_List(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/listUsersResponse.json"))
	})
	mux.HandleFunc("/v2/prototype/domains/admins/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/DAUsersResponse.json"))
	})

	resp, err := client.Users.List()
	assert.NoError(t, err)
	assert.Equal(t, "divvyuser:2:", resp.Users[0].ResourceID)
	assert.Equal(t, "divvyuser:9:", resp.Users[1].ResourceID)
	assert.Equal(t, "divvyuser:10:", resp.Users[2].ResourceID)
	assert.Equal(t, "Timmy Testington", resp.Users[0].Name)
	assert.Equal(t, "Billy Bobb", resp.Users[1].Name)
	assert.Equal(t, "Mitchell Jacks", resp.Users[2].Name)
	assert.Equal(t, 2, resp.Users[0].ID)
	assert.Equal(t, 9, resp.Users[1].ID)
	assert.Equal(t, 10, resp.Users[2].ID)
	assert.Equal(t, false, resp.Users[0].OrgAdmin)
	assert.Equal(t, false, resp.Users[1].OrgAdmin)
	assert.Equal(t, true, resp.Users[2].OrgAdmin)
	assert.Equal(t, false, resp.Users[0].DomainAdmin)
	assert.Equal(t, false, resp.Users[1].DomainAdmin)
	assert.Equal(t, false, resp.Users[2].DomainAdmin)
	assert.Equal(t, false, resp.Users[0].DomainViewer)
	assert.Equal(t, false, resp.Users[1].DomainViewer)
	assert.Equal(t, false, resp.Users[2].DomainViewer)
	assert.Equal(t, "testington@testers.com", resp.Users[0].Email)
	assert.Equal(t, "bbobb@bingo.xyz", resp.Users[1].Email)
	assert.Equal(t, "mjacks@xyz.org", resp.Users[2].Email)
	assert.Equal(t, "testuser1", resp.Users[0].Username)
	assert.Equal(t, "bbobb", resp.Users[1].Username)
	assert.Equal(t, "mjacks", resp.Users[2].Username)
	assert.Equal(t, "Default Organization", resp.Users[0].Org)
	assert.Equal(t, "Default Organization", resp.Users[1].Org)
	assert.Equal(t, "Default Organization", resp.Users[2].Org)
	assert.Equal(t, 1, resp.Users[0].OrgID)
	assert.Equal(t, 1, resp.Users[1].OrgID)
	assert.Equal(t, 1, resp.Users[2].OrgID)
	assert.Equal(t, false, resp.Users[0].TwoFactorEnabled)
	assert.Equal(t, true, resp.Users[1].TwoFactorEnabled)
	assert.Equal(t, false, resp.Users[2].TwoFactorEnabled)
	assert.Equal(t, false, resp.Users[0].TwoFactorRequired)
	assert.Equal(t, true, resp.Users[1].TwoFactorRequired)
	assert.Equal(t, false, resp.Users[2].TwoFactorRequired)
	assert.Equal(t, 1, resp.Users[0].Groups)
	assert.Equal(t, 0, resp.Users[1].Groups)
	assert.Equal(t, 0, resp.Users[2].Groups)
	assert.Equal(t, 0, resp.Users[0].OwnedResources)
	assert.Equal(t, 0, resp.Users[1].OwnedResources)
	assert.Equal(t, 5, resp.Users[2].OwnedResources)
	assert.Equal(t, 0, resp.Users[0].FailedLoginAttempts)
	assert.Equal(t, 0, resp.Users[1].FailedLoginAttempts)
	assert.Equal(t, 0, resp.Users[2].FailedLoginAttempts)
	assert.Equal(t, false, resp.Users[0].Suspended)
	assert.Equal(t, false, resp.Users[1].Suspended)
	assert.Equal(t, false, resp.Users[2].Suspended)
	assert.Equal(t, "2022-04-11 15:23:19", resp.Users[0].LastLogin)
	assert.Equal(t, []string{}, resp.Users[0].NavigationBlacklist)
	assert.Equal(t, []string{}, resp.Users[1].NavigationBlacklist)
	assert.Equal(t, []string{}, resp.Users[2].NavigationBlacklist)
	assert.Equal(t, false, resp.Users[0].RequirePWReset)
	assert.Equal(t, true, resp.Users[1].RequirePWReset)
	assert.Equal(t, false, resp.Users[2].RequirePWReset)
	assert.Equal(t, false, resp.Users[0].ConsoleAccessDenied)
	assert.Equal(t, false, resp.Users[1].ConsoleAccessDenied)
	assert.Equal(t, false, resp.Users[2].ConsoleAccessDenied)
	assert.Equal(t, true, resp.Users[0].ActiveAPIKey)
	assert.Equal(t, false, resp.Users[1].ActiveAPIKey)
	assert.Equal(t, false, resp.Users[2].ActiveAPIKey)
	assert.Equal(t, "2021-11-02 21:27:39", resp.Users[0].Created)
	assert.Equal(t, "2022-04-11 15:11:34", resp.Users[1].Created)
	assert.Equal(t, "2022-04-11 15:12:34", resp.Users[2].Created)
	assert.Equal(t, 3, resp.Count)
	teardown()
}

func TestUsers_Create(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/createUserResponse.json"))
	})

	resp, err := client.Users.Create(User{
		Name:              "Boaty McBoatFace",
		Username:          "Boatface",
		Email:             "boat@boatface.com",
		Password:          "password",
		ConfirmPassword:   "password",
		AccessLevel:       "BASIC_USER",
		TwoFactorRequired: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, 11, resp.ID)
	assert.Equal(t, 1, resp.OrgID)
	assert.Equal(t, "Boaty McBoatFace", resp.Name)
	assert.Equal(t, "Boatface", resp.Username)
	assert.Equal(t, "f7.,Pcd!RBm|", resp.TempPassword)
	assert.Equal(t, "2022-04-11T19:34:34Z", resp.TempPasswordExpiration)
	teardown()
}

func TestUsers_CreateAPIUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/createAPIUserResponse.json"))
	})

	resp, err := client.Users.CreateAPIUser(APIUser{
		Name:               "Boaty McBoatface",
		Username:           "Boatface",
		Email:              "boat@boatface.com",
		AuthenticationType: "internal",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Boaty McBoatface", resp.Name)
	assert.Equal(t, 13, resp.ID)
	assert.Equal(t, 1, resp.OrgID)
	assert.Equal(t, "Boatface", resp.Username)
	assert.Equal(t, "FakeKEY====================================ENDHere", resp.APIKey)
	teardown()
}

func TestUsers_CreateSAMLUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/createUserSAMLResponse.json"))
	})

	resp, err := client.Users.CreateSAMLUser(SAMLUser{
		Name:                   "Boat McBoatFace",
		Username:               "Boatface",
		Email:                  "boat@boatface.com",
		AccessLevel:            "BASIC_USER",
		AuthenticationType:     "saml",
		AuthenticationServerID: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, "divvyuser:12:", resp.ResourceID)
	assert.Equal(t, "Boaty McBoatFace", resp.Name)
	assert.Equal(t, "Boatface", resp.Username)
	assert.Equal(t, 12, resp.ID)
	assert.Equal(t, false, resp.OrgAdmin)
	assert.Equal(t, false, resp.DomainAdmin)
	assert.Equal(t, false, resp.DomainViewer)
	assert.Equal(t, "boat@boatface.com", resp.Email)
	assert.Equal(t, "Boatface", resp.Username)
	assert.Equal(t, "Default Organization", resp.Org)
	assert.Equal(t, 1, resp.OrgID)
	assert.Equal(t, false, resp.TwoFactorEnabled)
	assert.Equal(t, false, resp.TwoFactorRequired)
	assert.Equal(t, []string{}, resp.NavigationBlacklist)
	teardown()
}

func TestUsers_Delete(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/user/divvyuser:2:/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.Users.Delete("divvyuser:2:")
	assert.NoError(t, err)

	teardown()
}

func TestUsers_DeleteByUsername(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/listUsersResponse.json"))
	})
	mux.HandleFunc("/v2/prototype/domains/admins/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/DAUsersResponse.json"))
	})
	mux.HandleFunc("/v2/prototype/user/divvyuser:2:/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.Users.DeleteByUsername("testuser1")
	assert.NoError(t, err)

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

func TestUsers_Enable2FACurrentUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/tfa_enable", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\n\"otp_secret\": \"123456789\"\n}")
	})
	resp, err := client.Users.Enable2FACurrentUser()
	assert.NoError(t, err)
	assert.Equal(t, "123456789", resp.Secret)

	teardown()
}

func TestUsers_Disable2FA(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/tfa_disable", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\n\"success\": true\n}")
	})
	err := client.Users.Disable2FA(3)
	assert.NoError(t, err)
	teardown()

	setup()
	mux.HandleFunc("/v2/public/user/tfa_disable", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\n\"success\": false\n}")
	})
	err = client.Users.Disable2FA(3)
	assert.Error(t, err)
	teardown()
}

func TestUsers_CovertToAPIOnly(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/update_to_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\n\"user_id\": \"2\",\n\"api_key\": \"Aa11111111111caa11a1a111-AaaAa1_aaa11AAaAA1111\"}")
	})
	resp, err := client.Users.ConvertToAPIOnly(2)
	assert.NoError(t, err)
	assert.Equal(t, "2", resp.ID)
	assert.Equal(t, "Aa11111111111caa11a1a111-AaaAa1_aaa11AAaAA1111", resp.APIKey)
	teardown()
}

func TestUsers_SetConsoleAccess(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/update_console_access", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.Users.SetConsoleAccess(2, true)
	assert.NoError(t, err)
	teardown()
}

func TestUsers_DeactivateAPIKeys(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/apikey/deactivate", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.Users.DeactivateAPIKeys(2)
	assert.NoError(t, err)
	teardown()
}

func TestUsers_GetUserByUsername(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/listUsersResponse.json"))
	})
	mux.HandleFunc("/v2/prototype/domains/admins/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/DAUsersResponse.json"))
	})
	
	user, err := client.Users.GetUserByUsername("bbobb")
	assert.NoError(t, err)
	assert.Equal(t, "bbobb", user.Username)
	assert.Equal(t, "bbobb@bingo.xyz", user.Email)

	_, err = client.Users.GetUserByUsername("ddddd")
	assert.Error(t, err)
	teardown()
}

func TestUsers_GetUserByID(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/listUsersResponse.json"))
	})
	mux.HandleFunc("/v2/prototype/domains/admins/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/DAUsersResponse.json"))
	})
	
	user, err := client.Users.GetUserByID(9)
	assert.NoError(t, err)
	assert.Equal(t, 9, user.ID)
	assert.Equal(t, "bbobb@bingo.xyz", user.Email)

	_, err = client.Users.GetUserByID(22)
	assert.Error(t, err)
	teardown()
}
