package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const HostURL string = "http://localhost:8001"

// Client
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	APIKey     string
	Auth       AuthStruct

	AuthenticationServers AuthenticationServers
	Badges                Badges
	Bots                  Bots
	Clouds                Clouds
	Filters               Filters
	Insights              Insights
	Organizations         Organizations
	Resources             Resources
	ResourceGroups        ResourceGroups
	Users                 Users
}

type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"user_name"`
	Email            string `json:"user_email"`
	SessionID        string `json:"session_id"`
	Timeout          int    `json:"session_timeout"`
	DomainAdmin      bool   `json:"domain_admin"`
	CustomerID       string `json:"customer_id"`
	DomainViewer     bool   `json:"domain_veiwer"`
	AuthPluginExists bool   `json:"auth_plugin_exists"`
}

type APIErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
	Traceback    string `json:"traceback"`
}

// NewClient
func NewClient(host, username, pass, apikey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if apikey != nil {
		c.APIKey = *apikey
	}

	if username == nil || pass == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *pass,
	}

	resp, err := c.Login()
	if err != nil {
		return nil, err
	}

	c.Token = resp.SessionID
	c.AuthenticationServers = authServers{client: &c}
	c.Badges = &badges{client: &c}
	c.Bots = &bots{client: &c}
	c.Clouds = &clouds{client: &c}
	c.Filters = &filters{client: &c}
	c.Insights = &insights{client: &c}
	c.Users = &users{client: &c}
	c.Organizations = &orgs{client: &c}
	c.Resources = &resources{client: &c}
	c.ResourceGroups = &rsgroup{client: &c}
	return &c, nil
}

// Login to InsightCloudSec
func (c *Client) Login() (*AuthResponse, error) {

	// Verify AuthStruct is not blank
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("missing username and password")
	}

	// Make request
	body, err := c.makeRequest(http.MethodPost, "/v2/public/user/login", c.Auth, nil)
	if err != nil {
		return nil, errors.New("unable to login")
	}

	// Unmarshal data
	resp := AuthResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) makeRawRequest(method, path string, data interface{}, authtoken *string) (*http.Response, error) {

	// Set token
	token := c.Token
	if authtoken != nil {
		token = *authtoken
	}

	// Marshal Data for Payload
	byte_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Build Request
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.HostURL, path), bytes.NewBuffer(byte_data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("API-Key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "insightcloudsec-client-go")

	// Get Response
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) makeRequest(method, path string, data interface{}, authtoken *string) ([]byte, error) {

	resp, err := c.makeRawRequest(method, path, data, authtoken)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var api_error APIErrorResponse
		_ = json.Unmarshal(body, &api_error)
		return nil, fmt.Errorf("\n      HTTP Status: %d,\n   API Error Type: %s,\nAPI Error Message: %s", resp.StatusCode, api_error.ErrorType, api_error.ErrorMessage)
	}

	return body, err
}
