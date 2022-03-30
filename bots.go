package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CONSTANTS
///////////////////////////////////////////

const (
	// Bot Severities
	BOT_SEVERITY_LOW    = "low"
	BOT_SEVERITY_MEDIUM = "medium"
	BOT_SEVERITY_HIGH   = "high"

	// Bot Categories
	BOT_CATEGORY_SECURITY       = "Security"
	BOT_CATEGORY_OPTIMIZATION   = "Optimization"
	BOT_CATEGORY_CURATION       = "Curation"
	BOT_CATEGORY_BEST_PRACTICES = "Best Practices"
	BOT_CATEOGRY_MISC           = "Miscellaneous"

	// Bot States
	BOT_STATE_RUNNING  = "RUNNING"
	BOT_STATE_ARCHIVED = "ARCHIVED"
	BOT_STATE_SCANNING = "SCANNING"
	BOT_STATE_PAUSED   = "PAUSED"
)

// VARIABLES FOR VALIDATIONS
///////////////////////////////////////////
var (
	BOT_SEVERITY_RANGES = []string{BOT_SEVERITY_HIGH, BOT_SEVERITY_MEDIUM, BOT_SEVERITY_LOW}
	BOT_CATEGORIES      = []string{BOT_CATEGORY_SECURITY, BOT_CATEGORY_OPTIMIZATION, BOT_CATEGORY_CURATION, BOT_CATEGORY_BEST_PRACTICES, BOT_CATEOGRY_MISC}
	BOT_STATES          = []string{BOT_STATE_RUNNING, BOT_STATE_PAUSED, BOT_STATE_ARCHIVED, BOT_STATE_SCANNING}
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
	err := validateBot(bot_data)
	if err != nil {
		return BotResults{}, nil
	}

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

func (c Client) ArchiveBot(id string) error {
	_, err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/archive", id), nil)
	return err
}

func (c Client) PauseBot(id string) error {
	// Function pauses the bot of the given Resource ID and returns error if failed
	_, err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/pause", id), nil)
	return err
}

func (c Client) EnableBot(id string) error {
	// Function enables the both of the given Resource ID and returns error if failed
	_, err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/resume", id), nil)
	return err
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

func validateBot(b Bot) error {
	if !isInSlice(b.Severity, BOT_SEVERITY_RANGES) {
		return fmt.Errorf("[-] ERROR: Bot Severity must be one of %s", BOT_SEVERITY_RANGES)
	}

	if !isInSlice(b.Category, BOT_CATEGORIES) {
		return fmt.Errorf("[-] ERROR: Bot Category must be one of %s", BOT_CATEGORIES)
	}
	return nil
}
