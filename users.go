package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var _ Users = (*users)(nil)

type Users interface {
	Create(user User) (UserListDetails, error)
	CreateAPIUser(api_user APIUser) (APIUserResponse, error)
	CreateSAMLUser(saml_user SAMLUser) (UserListDetails, error)
	CurrentUserInfo() (UserListDetails, error)
	ConvertToAPIOnly(user_id int) (APIKey_Response, error)
	Get2FAStatus(user_id int32) (UsersMFAStatus, error)
	Enable2FACurrentUser() (OTP, error)
	Disable2FA(user_id int32) error
	Delete(user_resource_id string) error
	DeleteByUsername(username string) error
	List() (UserList, error)
	SetConsoleAccess(user_id int, access bool) error
}

type users struct {
	client *Client
}

type User struct {
	// For use when creating a user.
	Name              string `json:"name"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmPassword   string `json:"confirm_password"`
	AccessLevel       string `json:"access_level"`
	TwoFactorRequired bool   `json:"two_factor_required"`
}

type APIUser struct {
	// For use when creating an API only user.
	Name               string `json:"name"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	AuthenticationType string `json:"authentication_type"`
}

type APIUserResponse struct {
	ID       int    `json:"user_id"`
	OrgID    int    `json:"organization_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	APIKey   string `json:"api_key"`
}

type APIKey_Response struct {
	ID     string `json:"user_id"`
	APIKey string `json:"api_key"`
}

type SAMLUser struct {
	// For use when creating an SAML user.
	Name                   string `json:"name"`
	Username               string `json:"username"`
	Email                  string `json:"email"`
	AccessLevel            string `json:"access_level"`
	AuthenticationType     string `json:"authentication_type"`
	AuthenticationServerID int32  `json:"authentication_server_id"`
}

type UserListDetails struct {
	// For use with data returned from individual users in a UserList struct.
	Username               string   `json:"username"`
	ID                     int      `json:"user_id"`
	Created                string   `json:"create_date"`
	Name                   string   `json:"name"`
	Email                  string   `json:"email_address"`
	ResourceID             string   `json:"resource_id"`
	TwoFactorEnabled       bool     `json:"two_factor_enabled"`
	TwoFactorRequired      bool     `json:"two_factor_required"`
	FailedLoginAttempts    int      `json:"consecutive_failed_login_attempts,omitempty"`
	LastLogin              string   `json:"last_login_time,omitempty"`
	Suspended              bool     `json:"suspended,omitempty"`
	NavigationBlacklist    []string `json:"navigation_blacklist"`
	RequirePWReset         bool     `json:"require_pw_reset"`
	ConsoleAccessDenied    bool     `json:"console_access_denied,omitempty"`
	ActiveAPIKey           bool     `json:"active_api_key_present,omitempty"`
	Org                    string   `json:"organization_name"`
	OrgID                  int      `json:"organization_id"`
	DomainAdmin            bool     `json:"domain_admin"`
	DomainViewer           bool     `json:"domain_viewer"`
	OrgAdmin               bool     `json:"organization_admin"`
	Groups                 int      `json:"groups,omitempty"`
	AuthPluginExists       bool     `json:"auth_plugin_exists,omitempty"`
	OwnedResources         int      `json:"owned_resources,omitempty"`
	TempPassword           string   `json:"temporary_pw,omitempty"`
	TempPasswordExpiration string   `json:"temp_pw_expiration,omitempty"`
}

type UserList struct {
	// For use with data returned from a listing of users.
	Users []UserListDetails `json:"users"`
	Count int               `json:"total_count"`
}

type UserIDPayload struct {
	UserID int32 `json:"user_id,omitempty"`
}

type UserIDPayloadString struct {
	UserID string `json:"user_id"`
}

type UsersMFAStatus struct {
	Enabled  bool `json:"enabled"`
	Required bool `json:"required"`
}

type OTP struct {
	Secret string `json:"otp_secret"`
}

type Success struct {
	Success bool `json:"success"`
}

// USER FUNCTIONS

func (c *users) List() (UserList, error) {
	// List all InsightCloudSec users

	resp, err := c.client.makeRequest(http.MethodGet, "/v2/public/users/list", nil)
	if err != nil {
		return UserList{}, err
	}

	var ret UserList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserList{}, err
	}

	return ret, nil
}

func (c *users) Create(user User) (UserListDetails, error) {
	// Creates an InsightCloudSec User account

	if user.Password == "" || user.AccessLevel == "" || user.Name == "" || user.Username == "" || user.Email == "" {
		return UserListDetails{}, fmt.Errorf("[-] user's name, username, email, password and accesslevel must be set")
	}

	// If user.ConfirmPassword is empty, make it the same as user.Password
	if user.ConfirmPassword == "" {
		user.ConfirmPassword = user.Password
	}

	// Validate AccessLevel settings
	if user.AccessLevel != "BASIC_USER" && user.AccessLevel != "ORGANIZATION_ADMIN" && user.AccessLevel != "DOMAIN_VIEWER" && user.AccessLevel != "DOMAIN_ADMIN" {
		return UserListDetails{}, fmt.Errorf("[-] user.AccessLevel must be one of: BASIC_USER, ORGANIZATION_ADMIN, DOMAIN_VIEWER, or DOMAIN_ADMIN")
	}

	data, err := json.Marshal(user)
	if err != nil {
		return UserListDetails{}, err
	}

	resp, err := c.client.makeRequest(http.MethodPost, "/v2/public/user/create", bytes.NewBuffer(data))
	if err != nil {
		return UserListDetails{}, err
	}

	var ret UserListDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserListDetails{}, err
	}

	return ret, nil
}

func (c *users) CreateAPIUser(api_user APIUser) (APIUserResponse, error) {
	// Creates an InsightCloudSec API Only User

	if api_user.Username == "" || api_user.Email == "" || api_user.Name == "" {
		return APIUserResponse{}, fmt.Errorf("[-] user's name, username, email must be set")
	}

	api_user.AuthenticationType = "internal"

	data, err := json.Marshal(api_user)
	if err != nil {
		return APIUserResponse{}, err
	}

	resp, err := c.client.makeRequest(http.MethodPost, "/v2/public/user/create_api_only_user", bytes.NewBuffer(data))
	if err != nil {
		return APIUserResponse{}, err
	}
	var ret APIUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return APIUserResponse{}, err
	}
	return ret, nil
}

func (u *users) CreateSAMLUser(saml_user SAMLUser) (UserListDetails, error) {
	payload, err := json.Marshal(saml_user)
	if err != nil {
		return UserListDetails{}, err
	}

	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/create", bytes.NewBuffer(payload))
	if err != nil {
		return UserListDetails{}, err
	}

	var ret UserListDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserListDetails{}, err
	}
	return ret, err
}

func (u *users) Delete(user_resource_id string) error {
	// Deletes the user corresponding to the given user_resource_id.
	//
	// Example usage:  client.DeleteUser("divvyuser:7")

	resp, err := u.client.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/user/%s/delete", user_resource_id), nil)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	return nil
}

func (c *users) DeleteByUsername(username string) error {
	// Deletes the user corresponding to the given username.
	//
	// Example usage: client.DeleteUserByUsername("jdoe")

	users, err := c.List()
	if err != nil {
		return err
	}

	var id string
	for _, user := range users.Users {
		if user.Username == strings.ToLower(username) {
			id = user.ResourceID
		}
	}

	if id == "" {
		return fmt.Errorf("[-] ERROR: Username not found")
	}

	err = c.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *users) CurrentUserInfo() (UserListDetails, error) {
	resp, err := u.client.makeRequest(http.MethodGet, "/v2/public/user/info", nil)
	if err != nil {
		return UserListDetails{}, err
	}

	var user UserListDetails
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return UserListDetails{}, err
	}

	return user, nil
}

func (u *users) Get2FAStatus(user_id int32) (UsersMFAStatus, error) {
	id := UserIDPayload{UserID: user_id}
	payload, err := json.Marshal(id)
	if err != nil {
		return UsersMFAStatus{}, err
	}

	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/tfa_state", bytes.NewBuffer(payload))
	if err != nil {
		return UsersMFAStatus{}, err
	}

	var ret UsersMFAStatus
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UsersMFAStatus{}, err
	}

	return ret, err
}

func (u *users) Enable2FACurrentUser() (OTP, error) {
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/tfa_enable", nil)
	if err != nil {
		return OTP{}, err
	}
	var ret OTP
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return OTP{}, err
	}
	return ret, nil
}

func (u *users) Disable2FA(user_id int32) error {
	payload, err := json.Marshal(UserIDPayload{UserID: user_id})
	if err != nil {
		return err
	}
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/tfa_disable", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	var ret Success
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return err
	}

	if ret.Success == true {
		return nil
	}

	return fmt.Errorf("ERROR: API Returned a failure attempting to disable")
}

func (u *users) ConvertToAPIOnly(user_id int) (APIKey_Response, error) {
	payload, err := json.Marshal(UserIDPayloadString{UserID: strconv.Itoa(user_id)})
	if err != nil {
		return APIKey_Response{}, err
	}
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/update_to_api_only_user", bytes.NewBuffer(payload))
	if err != nil {
		return APIKey_Response{}, err
	}

	var ret APIKey_Response
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return APIKey_Response{}, err
	}
	return ret, nil
}

func (u *users) SetConsoleAccess(user_id int, access bool) error {
	payload, err := json.Marshal(fmt.Sprintf("{\n\"user_id\": \"%d\",\n\"console_access_denied\": \"%t\"\n}", user_id, access))
	if err != nil {
		return err
	}
	_, err = u.client.makeRequest(http.MethodPost, "/v2/public/user/update_console_access", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	return nil
}
