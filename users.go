package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var _ Users = (*users)(nil)

type Users interface {
	Create(user User) (UserDetails, error)
	CreateAPIUser(api_user APIUser) (APIUserResponse, error)
	CreateSAMLUser(saml_user SAMLUser) (UserDetails, error)
	CurrentUserInfo() (UserDetails, error)
	ConvertToAPIOnly(user_id int) (APIKey_Response, error)
	Get2FAStatus(user_id int) (UsersMFAStatus, error)
	GetUserByID(user_id int) (UserDetails, error)
	GetUserByUsername(username string) (UserDetails, error)
	Enable2FACurrentUser() (OTP, error)
	Disable2FA(user_id int) error
	DeactivateAPIKeys(user_id int) error
	Delete(user_resource_id string) error
	DeleteByUsername(username string) error
	List() (UserList, error)
	ListDomainAdmins() (UserList, error)
	ListAll() (UserList, error)
	SetConsoleAccess(user_id int, access bool) error
	UpdateUserInfo(user_id int, name string, username string, email string, access_level string) (UserDetails, error)
	EditAccessLevel(user_id int, current string, desired string) (UserDetails, error)
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

type UserDetails struct {
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
	Users []UserDetails `json:"users"`
	Count int           `json:"total_count"`
}

type UserIDPayload struct {
	UserID int `json:"user_id,omitempty"`
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

type ConsoleDeniedRequest struct {
	UserID string `json:"user_id"`
	Access bool   `json:"console_access_denied"`
}

type UserInfoUpdate struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessLevel string `json:"access_level"`
}

type AccessLevelChange struct {
	Current string `json:"current_access_level"`
	Desired string `json:"new_access_level"`
}

// USER FUNCTIONS

func (u *users) List() (UserList, error) {
	// List all InsightCloudSec users (non-Domain Admins)
	resp, err := u.client.makeRequest(http.MethodGet, "/v2/public/users/list", nil)
	if err != nil {
		return UserList{}, err
	}

	var ret UserList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserList{}, err
	}

	return ret, nil
}

func (u *users) ListDomainAdmins() (UserList, error) {
	// List domain admins in InsightCloudSec
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/prototype/domains/admins/list", nil)
	if err != nil {
		return UserList{}, err
	}

	var ret UserList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserList{}, err
	}
	ret.Count = len(ret.Users)

	return ret, nil
}

func (u *users) ListAll() (UserList, error) {
	//  Return one list of both domain admins and regular users
	regular_users, err := u.List()
	if err != nil {
		return UserList{}, err
	}
	domain_admins, err := u.ListDomainAdmins()
	if err != nil {
		return UserList{}, err
	}

	var combined UserList
	combined.Count = regular_users.Count + domain_admins.Count
	combined.Users = append(regular_users.Users, domain_admins.Users...)

	return combined, nil
}

func (u *users) Create(user User) (UserDetails, error) {
	// Creates an InsightCloudSec User account

	// If required values are empty, return error
	if user.AccessLevel == "" || user.Name == "" || user.Username == "" || user.Email == "" {
		return UserDetails{}, fmt.Errorf("[-] user's name, username, email, password and accesslevel must be set")
	}

	// If user.ConfirmPassword is empty, make it the same as user.Password
	if user.ConfirmPassword == "" {
		user.ConfirmPassword = user.Password
	}

	// Validate AccessLevel settings
	if user.AccessLevel != "BASIC_USER" && user.AccessLevel != "ORGANIZATION_ADMIN" && user.AccessLevel != "DOMAIN_VIEWER" && user.AccessLevel != "DOMAIN_ADMIN" {
		return UserDetails{}, fmt.Errorf("[-] user.AccessLevel must be one of: BASIC_USER, ORGANIZATION_ADMIN, DOMAIN_VIEWER, or DOMAIN_ADMIN")
	}

	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/create", user)
	if err != nil {
		return UserDetails{}, err
	}

	var ret UserDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserDetails{}, err
	}

	return ret, nil
}

func (u *users) CreateAPIUser(api_user APIUser) (APIUserResponse, error) {
	// Creates an InsightCloudSec API Only User

	if api_user.Username == "" || api_user.Email == "" || api_user.Name == "" {
		return APIUserResponse{}, fmt.Errorf("[-] user's name, username, email must be set")
	}

	api_user.AuthenticationType = "internal"

	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/create_api_only_user", api_user)
	if err != nil {
		return APIUserResponse{}, err
	}
	var ret APIUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return APIUserResponse{}, err
	}
	return ret, nil
}

func (u *users) CreateSAMLUser(saml_user SAMLUser) (UserDetails, error) {
	// Creates a SAML User
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/create", saml_user)
	if err != nil {
		return UserDetails{}, err
	}

	var ret UserDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserDetails{}, err
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

func (u *users) DeleteByUsername(username string) error {
	// Deletes the user corresponding to the given username.
	//
	// Example usage: client.DeleteUserByUsername("jdoe")

	users, err := u.ListAll()
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

	err = u.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *users) CurrentUserInfo() (UserDetails, error) {
	resp, err := u.client.makeRequest(http.MethodGet, "/v2/public/user/info", nil)
	if err != nil {
		return UserDetails{}, err
	}

	var user UserDetails
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return UserDetails{}, err
	}

	return user, nil
}

func (u *users) Get2FAStatus(user_id int) (UsersMFAStatus, error) {
	// Gets the 2FA status for user of given user_id
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/tfa_state", UserIDPayload{UserID: user_id})
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
	// Enables 2FA for current user
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

func (u *users) Disable2FA(user_id int) error {
	// Disables 2FA for user of given user_id
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/tfa_disable", UserIDPayload{UserID: user_id})
	if err != nil {
		return err
	}

	var ret Success
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return err
	}

	if ret.Success {
		return nil
	}

	return fmt.Errorf("ERROR: API Returned a failure attempting to disable")
}

func (u *users) ConvertToAPIOnly(user_id int) (APIKey_Response, error) {
	// Converts user of given user_id to API Only User
	resp, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/update_to_api_only_user", UserIDPayloadString{UserID: strconv.Itoa(user_id)})
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
	//Sets the console access for given user_id
	_, err := u.client.makeRequest(http.MethodPost, "/v2/public/user/update_console_access", ConsoleDeniedRequest{UserID: strconv.Itoa(user_id), Access: access})
	if err != nil {
		return err
	}

	return nil
}

func (u *users) DeactivateAPIKeys(user_id int) error {
	// Deactivates the API Keys for user of given user_id
	_, err := u.client.makeRequest(http.MethodPost, "/v2/public/apikey/deactivate", UserIDPayloadString{UserID: strconv.Itoa(user_id)})
	if err != nil {
		return err
	}
	return nil
}

func (u *users) GetUserByUsername(username string) (UserDetails, error) {
	listOfUsers, err := u.client.Users.ListAll()
	if err != nil {
		return UserDetails{}, err
	}

	for i, user := range listOfUsers.Users {
		if user.Username == username {
			return listOfUsers.Users[i], nil
		}
	}

	return UserDetails{}, fmt.Errorf("[-] ERROR: Found no user with username %s", username)
}

func (u *users) GetUserByID(user_id int) (UserDetails, error) {
	listOfUsers, err := u.client.Users.ListAll()
	if err != nil {
		return UserDetails{}, err
	}

	for i, user := range listOfUsers.Users {
		if user.ID == user_id {
			return listOfUsers.Users[i], nil
		}
	}

	return UserDetails{}, fmt.Errorf("[-] ERROR: Found no user with user_id: %d", user_id)
}

func (u *users) UpdateUserInfo(user_id int, name string, username string, email string, access_level string) (UserDetails, error) {
	payload := UserInfoUpdate{Name: name, Username: username, Email: email, AccessLevel: access_level}
	resp, err := u.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/user/divvyuser:%d:/update", user_id), payload)
	if err != nil {
		return UserDetails{}, err
	}

	var ret UserDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserDetails{}, err
	}

	return ret, nil
}

func (u *users) EditAccessLevel(user_id int, current string, desired string) (UserDetails, error) {
	payload := AccessLevelChange{Current: current, Desired: desired}
	resp, err := u.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/user/divvyuser:%d:/edit-access-level", user_id), payload)
	if err != nil {
		return UserDetails{}, err
	}

	var ret UserDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserDetails{}, err
	}

	return ret, nil
}
