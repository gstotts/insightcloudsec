package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

var (
	BOT_SEVERITY_RANGES = []string{BOT_SEVERITY_HIGH, BOT_SEVERITY_MEDIUM, BOT_SEVERITY_LOW}
	BOT_CATEGORIES      = []string{BOT_CATEGORY_SECURITY, BOT_CATEGORY_OPTIMIZATION, BOT_CATEGORY_CURATION, BOT_CATEGORY_BEST_PRACTICES, BOT_CATEOGRY_MISC}
	BOT_STATES          = []string{BOT_STATE_RUNNING, BOT_STATE_PAUSED, BOT_STATE_ARCHIVED, BOT_STATE_SCANNING}
)

var _ Bots = (*bots)(nil)

type Bots interface {
	ArchiveBot(id string) error
	Create(bot_data Bot) (BotResults, error)
	GetBotByID(id string) (BotResults, error)
	EnableBot(id string) error
	List() (BotList, error)
	PauseBot(id string) error
}

type bots struct {
	client *Client
}

type Bot struct {
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Severity        string          `json:"severity"`
	Category        string          `json:"category"`
	OnDemandEnabled bool            `json:"ondemand_enabled"`
	State           string          `json:"state"`
	Instructions    BotInstructions `json:"instructions"`
}

type BotList struct {
	Bots  []BotResults `json:"bots"`
	Count int          `json:"count"`
}

type BotResults struct {
	ResourceID            string          `json:"resource_id"`
	Name                  string          `json:"name"`
	Description           string          `json:"description"`
	Notes                 string          `json:"notes"`
	InsightID             string          `json:"insight_id"`
	Source                string          `json:"source"`
	InsightName           string          `json:"insight_name"`
	Owner                 string          `json:"owner"`
	OwnerName             string          `json:"owner_name"`
	State                 string          `json:"state"`
	DateCreated           string          `json:"date_created"`
	DateModified          string          `json:"date_modified"`
	Category              string          `json:"category"`
	BadgeScopeOperator    string          `json:"badge_scope_operator"`
	Instructions          BotInstructions `json:"instructions"`
	Schedule              BotSchedule     `json:"schedule"`
	HookpointCreated      bool            `json:"hookpoint_created"`
	HookpointModified     bool            `json:"hookpoint_modified"`
	HookpointTagsModified bool            `json:"hookpoint_tags_modified"`
	HookpointDestroyed    bool            `json:"hookpoint_destroyed"`
	NextScheduled         float32         `json:"next_scheduled_run"`
	Valid                 bool            `json:"valid"`
	EventFailures         BotErrors       `json:"event_failures"`
	Severity              string          `json:"severity"`
	DetailedLogging       bool            `json:"detailed_logging"`
	Version               int             `json:"version"`
	ExemptionsCount       int             `json:"exemptions_count"`
}

type BotErrors struct {
	Errors       int `json:"errors"`
	Timeouts     int `json:"timeouts"`
	InvalidPerms int `json:"invalid_perms"`
}

type BotSchedule struct {
	Type         string               `json:"_type"`
	TimeOfDay    BotScheduleTimeOfDay `json:"time_of_day"`
	DayOfMonth   int                  `json:"day_of_month"`
	DayOfWeek    int                  `json:"day_of_week"`
	ExcludeDays  []int                `json:"exclude_days"`
	MinuteOfHour int                  `json:"minute_of_hour"`
	SecondOfHour int                  `json:"second_of_hour"`
}

type BotScheduleTimeOfDay struct {
	Type   string `json:"_type"`
	Second int    `json:"second"`
	Minute int    `json:"minute"`
	Hour   int    `json:"hour"`
}

type BotInstructions struct {
	ResourceTypes []string            `json:"resource_types"`
	Filters       []BotFilter         `json:"filters"`
	Actions       []BotAction         `json:"actions"`
	Groups        []string            `json:"groups"`
	Badges        []map[string]string `json:"badges"`
}

type BotFilter struct {
	Config map[string]interface{} `json:"config"`
	Name   string                 `json:"name"`
}

type BotAction struct {
	RunWhenResultIs bool        `json:"run_when_result_is"`
	Config          interface{} `json:"config"`
	Name            string      `json:"name"`
}

// RESOURCE FUNCTIONS
///////////////////////////////////////////

func (s *bots) Create(bot_data Bot) (BotResults, error) {
	err := validateBot(bot_data)
	if err != nil {
		return BotResults{}, nil
	}

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/public/botfactory/bot/create", bot_data)
	if err != nil {
		return BotResults{}, err
	}

	var ret BotResults
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return BotResults{}, err
	}

	return ret, nil
}

func (s *bots) List() (BotList, error) {
	default_body := make(map[string]interface{})
	default_body["filters"] = []string{}
	default_body["offset"] = 0

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/public/botfactory/list", default_body)
	if err != nil {
		return BotList{}, err
	}

	var ret BotList
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return BotList{}, err
	}

	return ret, nil
}

func (s *bots) ArchiveBot(id string) error {
	_, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/archive", id), nil)
	return err
}

func (s *bots) PauseBot(id string) error {
	// Function pauses the bot of the given Resource ID and returns error if failed
	_, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/pause", id), nil)
	return err
}

func (s *bots) EnableBot(id string) error {
	// Function enables the both of the given Resource ID and returns error if failed
	_, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/resume", id), nil)
	return err
}

func (s *bots) GetBotByID(id string) (BotResults, error) {
	resp, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/botfactory/%s/get", id), nil)
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
