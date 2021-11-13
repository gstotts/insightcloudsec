package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS

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

type SAMLUser struct {
	// For use when creating an SAML user.
	Name                   string `json:"name"`
	Username               string `json:"username"`
	Email                  string `json:"email"`
	AccessLevel            string `json:"access_level"`
	AuthenticationType     string `json:"authentication_type"`
	AuthenticationServerID int    `json:"authentication_server_id"`
}

type UserListDetails struct {
	// For use with data returned from individual users in a UserList struct.
	Username            string   `json:"username"`
	ID                  int      `json:"user_id"`
	Created             string   `json:"created_date"`
	Name                string   `json:"name"`
	Email               string   `json:"email_address"`
	ResourceID          string   `json:"resource_id"`
	TwoFactorEnabled    bool     `json:"two_factor_enabled"`
	TwoFactorRequired   bool     `json:"two_factor_required"`
	FailedLoginAttempts int      `json:"consecutive_failed_login_attempts"`
	LastLogin           string   `json:"last_login_time"`
	Suspended           bool     `json:"suspended"`
	NavigationBlacklist []string `json:"navigation_blacklist"`
	RequirePWReset      bool     `json:"require_pw_reset"`
	ConsoleAccessDenied bool     `json:"console_access_denied"`
	ActiveAPIKey        bool     `json:"active_api_key_present"`
	Org                 string   `json:"organization_name"`
	OrgID               int      `json:"organization_id"`
	DomainAdmin         bool     `json:"domain_admin"`
	DomainViewer        bool     `json:"domain_viewer"`
	OrgAdmin            bool     `json:"organization_admin"`
	Groups              int      `json:"groups"`
	OwnedResources      int      `json:"owned_resources"`
}

type UserList struct {
	// For use with data returned from a listing of users.
	Users []UserListDetails `json:"users"`
	Count int               `json:"total_count"`
}

// USER FUNCTIONS

func (c Client) ListUsers() ([]UserListDetails, error) {
	// List all InsightCloudSec users

	resp, err := c.makeRequest(http.MethodGet, "/v2/public/users/list", nil)
	if err != nil {
		return nil, err
	}

	var ret UserList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret.Users, nil
}

func (c Client) CreateUser(user User) (UserListDetails, error) {
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

	resp, err := c.makeRequest(http.MethodPost, "/v2/public/user/create", bytes.NewBuffer(data))
	if err != nil {
		return UserListDetails{}, err
	}

	var ret UserListDetails
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserListDetails{}, err
	}

	return ret, nil
}

func (c Client) CreateAPIUser(api_user APIUser) (APIUserResponse, error) {
	// Creates an InsightCloudSec API Only User

	if api_user.Username == "" || api_user.Email == "" || api_user.Name == "" {
		return APIUserResponse{}, fmt.Errorf("[-] user's name, username, email must be set")
	}

	api_user.AuthenticationType = "internal"

	data, err := json.Marshal(api_user)
	if err != nil {
		return APIUserResponse{}, err
	}

	resp, err := c.makeRequest(http.MethodPost, "/v2/public/user/create_api_only_user", bytes.NewBuffer(data))
	if err != nil {
		return APIUserResponse{}, err
	}
	var ret APIUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return APIUserResponse{}, err
	}
	return ret, nil
}
