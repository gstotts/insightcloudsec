package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var _ Users = (*users)(nil)

type Users interface {
	Create(user User) (UserListDetails, error)
	CreateAPIUser(api_user APIUser) (APIUserResponse, error)
	CurrentUserInfo() (UserListDetails, error)
	Delete(user_resource_id string) error
	DeleteByUsername(username string) error
	List() (UserList, error)
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
	FailedLoginAttempts int      `json:"consecutive_failed_login_attempts,omitempty"`
	LastLogin           string   `json:"last_login_time,omitempty"`
	Suspended           bool     `json:"suspended,omitempty"`
	NavigationBlacklist []string `json:"navigation_blacklist"`
	RequirePWReset      bool     `json:"require_pw_reset"`
	ConsoleAccessDenied bool     `json:"console_access_denied,omitempty"`
	ActiveAPIKey        bool     `json:"active_api_key_present,omitempty"`
	Org                 string   `json:"organization_name"`
	OrgID               int      `json:"organization_id"`
	DomainAdmin         bool     `json:"domain_admin"`
	DomainViewer        bool     `json:"domain_viewer"`
	OrgAdmin            bool     `json:"organization_admin"`
	Groups              int      `json:"groups,omitempty"`
	OwnedResources      int      `json:"owned_resources,omitempty"`
}

type UserList struct {
	// For use with data returned from a listing of users.
	Users []UserListDetails `json:"users"`
	Count int               `json:"total_count"`
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

func (c *users) Delete(user_resource_id string) error {
	// Deletes the user corresponding to the given user_resource_id.
	//
	// Example usage:  client.DeleteUser("divvyuser:7")

	resp, err := c.client.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/user/%s/delete", user_resource_id), nil)
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
