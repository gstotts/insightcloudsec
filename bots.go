package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////
type Bot struct {
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Severity        string          `json:"severity"`
	Category        string          `json:"category"`
	OnDemandEnabled bool            `json:"ondemand_enabled"`
	State           string          `json:"state"`
	Instructions    BotInstructions `json:"instructions"`
}

type BotList struct{}

type BotResults struct {
	ResourceID         string          `json:"resource_id"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Notes              string          `json:"notes"`
	InsightID          string          `json:"insight_id"`
	Source             string          `json:"source"`
	InsightName        string          `json:"insight_name"`
	Owner              string          `json:"owner"`
	OwnerName          string          `json:"owner_name"`
	State              string          `json:"state"`
	DateCreated        string          `json:"date_created"`
	DateModified       string          `json:"date_modified"`
	Category           string          `json:"category"`
	BadgeScopeOperator string          `json:"badge_scope_operator"`
	Instructions       BotInstructions `json:"instructions"`
	Valid              bool            `json:"valid"`
	Errors             []interface{}   `json:"errors"`
	Severity           string          `json:"severity"`
	DetailedLogging    bool            `json:"detailed_logging"`
	Scope              []string        `json:"scope"`
}

type BotInstructions struct {
	ResourceTypes       []string            `json:"resource_types"`
	Filters             []BotFilter         `json:"filters"`
	Actions             []BotAction         `json:"actions"`
	Groups              []string            `json:"groups"`
	Badges              []map[string]string `json:"badges"`
	Hookpoints          []string            `json:"hookpoints"`
	Schedule            string              `json:"schedule,omitempty"`
	ScheduleDescription string              `json:"schedule_description,omitempty"`
}

type BotFilter struct {
	Config interface{} `json:"config"`
	Name   string      `json:"name"`
}

type BotAction struct {
	RunWhenResultIs bool        `json:"run_when_result_is"`
	Config          interface{} `json:"config"`
	Name            string      `json:"name"`
}

// RESOURCE FUNCTIONS
///////////////////////////////////////////

func (c Client) CreateBot(bot_data Bot) (BotResults, error) {
	data, err := json.Marshal(bot_data)
	if err != nil {
		return BotResults{}, err
	}

	resp, err := c.makeRequest(http.MethodPost, "/v2/public/botfactory/bot/create", bytes.NewBuffer(data))
	if err != nil {
		return BotResults{}, err
	}

	var ret BotResults
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return BotResults{}, err
	}

	return ret, nil
}

func (c Client) GetBotByID(id string) (BotResults, error) {
	resp, err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/get", id), nil)
	if err != nil {
		return BotResults{}, err
	}

	var ret BotResults
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return BotResults{}, err
	}

	return ret, nil
}
