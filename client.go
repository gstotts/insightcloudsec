package insightcloudsec

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

// STRUCTS
///////////////////////////////////////////
type SessionInfo struct {
	ID   string `json:"session_id"`
	User string `json:"username"`
}

type Client struct {
	APIKey     string
	BaseURL    string
	SessionID  string
	HttpClient *http.Client
}

// HELPER FUNCTIONS
///////////////////////////////////////////
func __getBaseURL() string {
	baseURL := os.Getenv("INSIGHTCLOUDSEC_BASE_URL")
	// Prompt for missing baseURL if no env set
	reader := bufio.NewReader(os.Stdin)
	if baseURL == "" {
		fmt.Print("BaseURL: ")
		baseURL, _ = reader.ReadString('\n')
		baseURL = strings.TrimSpace(baseURL)
	}

	return baseURL
}

func __getSessionID(baseURL string, user string, pass string) (string, error) {
	client := http.DefaultClient
	var data = []byte(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", user, pass))
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/v2/public/user/login", baseURL), bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json;UTF-8")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		if resp.StatusCode == 401 {
			fmt.Println("\n[!] Error: Invalid username or account locked.")
			os.Exit(1)
		} else if resp.StatusCode == 500 {
			fmt.Println("\n[!] Error: Invalid password.")
			os.Exit(1)
		} else {
			return "", APIRequestError{
				StatusCode: resp.StatusCode,
				Message:    resp.Status,
			}
		}
	}
	defer resp.Body.Close()

	var session SessionInfo
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return "", err
	}
	fmt.Println("\n[+] User successfully logged in.")
	return session.ID, nil
}

// CLIENT FUNCTIONS
///////////////////////////////////////////

// Creates a new InsightCloudSec client.  If no information is passed, it will request interactively.
func NewClient() (*Client, error) {
	baseURL := __getBaseURL()
	apiKey := os.Getenv("INSIGHTCLOUDSEC_API_KEY")
	client := http.DefaultClient
	sessionID := ""

	reader := bufio.NewReader(os.Stdin)
	// Prompt for username and password if no apikey env set
	if apiKey == "" {
		fmt.Print("Username: ")
		user, _ := reader.ReadString('\n')
		user = strings.TrimSpace(user)

		fmt.Print("Password: ")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		var err error
		sessionID, err = __getSessionID(baseURL, user, string(bytePassword))
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		APIKey:     apiKey,
		BaseURL:    baseURL,
		HttpClient: client,
		SessionID:  sessionID,
	}, nil
}

// Creates a new InsightCloudSec client using the provided Api-Key
func NewClientWithKey(apiKey string) (*Client, error) {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    __getBaseURL(),
		HttpClient: http.DefaultClient,
		SessionID:  "",
	}, nil
}

// Creates a new InsightCloudSec client using the provided credentials
func NewClientWithCreds(user string, pass string) (*Client, error) {
	baseURL := __getBaseURL()
	sessionID, err := __getSessionID(baseURL, user, pass)
	if err != nil {
		return nil, err
	}
	return &Client{
		APIKey:     "",
		BaseURL:    baseURL,
		HttpClient: http.DefaultClient,
		SessionID:  sessionID,
	}, nil
}

func (c Client) makeRequest(method, path string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.BaseURL, path),
		data,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.APIKey == "" {
		req.Header.Set("X-Auth-Token", c.SessionID)
	} else {
		req.Header.Set("Api-Key", c.APIKey)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil || (resp.StatusCode != 200 && resp.StatusCode != 201) {
		return nil, APIRequestError{
			Request:    *req,
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
		}
	}
	return resp, nil
}
