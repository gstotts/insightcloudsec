package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	_userAgent = "insightcloudsec-client"
)

// Config provides configuration details to the client

type Config struct {
	BaseURL    string
	ApiKey     string
	Headers    http.Header
	HTTPClient *http.Client
}

// DefaultConfig returns a default configuration structure

func DefaultConfig() *Config {
	config := &Config{
		BaseURL:    os.Getenv("INSIGHTCLOUDSEC_BASE_URL"),
		ApiKey:     os.Getenv("INSIGHTCLOUDSEC_API_KEY"),
		HTTPClient: http.DefaultClient,
	}

	return config
}

// Client is the InsightCloudSec API client.

type Client struct {
	baseURL *url.URL
	apikey  string
	http    *http.Client

	AuthenticationServers AuthenticationServers
	Badges                Badges
	Bots                  Bots
	Clouds                Clouds
	CloudOrgs             CloudOrganizations
	Filters               Filters
	Insights              Insights
	Organizations         Organizations
	Resources             Resources
	ResourceGroups        ResourceGroups
	Users                 Users
}

// NewClient creates a new InsightCloudSec API client
func NewClient(cfg *Config) (*Client, error) {
	config := DefaultConfig()

	if cfg != nil {
		if cfg.BaseURL != "" {
			config.BaseURL = cfg.BaseURL
		}
		if cfg.ApiKey != "" {
			config.ApiKey = cfg.ApiKey
		}
		for k, v := range cfg.Headers {
			config.Headers[k] = v
		}
		if cfg.HTTPClient != nil {
			config.HTTPClient = cfg.HTTPClient
		}
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid baseURL address: %w", err)
	}

	if config.ApiKey == "" {
		return nil, fmt.Errorf("missing API token")
	}

	client := &Client{
		baseURL: baseURL,
		apikey:  config.ApiKey,
		http:    config.HTTPClient,
	}

	client.AuthenticationServers = &authServers{client: client}
	client.Badges = &badges{client: client}
	client.Bots = &bots{client: client}
	client.Clouds = &clouds{client: client}
	client.Filters = &filters{client: client}
	client.Insights = &insights{client: client}
	client.Users = &users{client: client}
	client.Organizations = &orgs{client: client}
	client.Resources = &resources{client: client}
	client.ResourceGroups = &rsgroup{client: client}
	client.CloudOrgs = &corgs{client: client}
	return client, nil
}

func (c Client) makeRequest(method, path string, data interface{}) (*http.Response, error) {

	// Marshall json if data is not nil
	byte_data, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("[-] ERROR: Marshal error: %s", err)
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.baseURL, path),
		bytes.NewBuffer(byte_data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-Key", c.apikey)
	req.Header.Set("User-Agent", _userAgent)

	resp, err := c.http.Do(req)
	if err != nil || (resp.StatusCode >= 300) {
		return nil, APIRequestError{
			Request:    *req,
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
		}
	}

	return resp, nil
}

func (c Client) Close() error {
	_, err := c.makeRequest(http.MethodPost, "/v2/public/user/logout", nil)
	if err != nil {
		return err
	}

	return nil
}
