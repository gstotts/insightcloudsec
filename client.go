package insightcloudsec-go

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

type SessionInfo struct {
	ID   string `json:"session_id"`
	User string `json:"username"`
}

type Client struct {
	APIKey     string
	BaseURL    string
	httpClient *http.Client
}

// newClient returns a new client for use with the API
func NewClient() (*Client, error) {
	baseURL := os.Getenv("INSIGHTCLOUDSEC_BASE_URL")
	apiKey := os.Getenv("INSIGHTCLOUDSEC_API_KEY")
	client := http.DefaultClient

	// Prompt for missing baseURL if no env set
	reader := bufio.NewReader(os.Stdin)
	if baseURL == "" {
		fmt.Print("BaseURL: ")
		baseURL, _ = reader.ReadString('\n')
		baseURL = strings.TrimSpace(baseURL)
	}

	// Prompt for username and password if no apikey env set
	if apiKey == "" {
		fmt.Print("Username: ")
		user, _ := reader.ReadString('\n')
		user = strings.TrimSpace(user)

		fmt.Print("Password: ")
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))

		var data = []byte(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", user, string(bytePassword)))
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s/v2/public/user/login", baseURL), bytes.NewBuffer(data))

		req.Header.Set("Content-Type", "application/json;UTF-8")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode >= http.StatusBadRequest {
			if resp.StatusCode == 401 {
				fmt.Println("\n[!] Error: Invalid username or account locked.")
				os.Exit(1)
			} else if resp.StatusCode == 500 {
				fmt.Println("\n[!] Error: Invalid password.")
				os.Exit(1)
			} else {
				return nil, APIRequestError{
					StatusCode: resp.StatusCode,
					Message:    resp.Status,
				}
			}
		}
		defer resp.Body.Close()

		var session SessionInfo
		if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
			return nil, err
		}
		apiKey = session.ID
		fmt.Println("\n[+] User successfully logged in.")
	}

	return &Client{
		APIKey:     apiKey,
		BaseURL:    baseURL,
		httpClient: client,
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
	req.Header.Set("X-Auth-Token", c.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, APIRequestError{
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
		}
	}
	return resp, nil
}
