// Last Reviewed: 2-Apr-2022
// InsightCloudSec Version at time of review: 22.2

package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

// STRUCTS
///////////////////////////////////////////

type Query struct {
	// Built off of the Query v3-ETL endpoint of the InsightCloudSec API
	Badges                 *[]Badge        `json:"badges,omitempty"`
	Badge_Filter_Operator  string          `json:"badge_filter_operator,omitempty"`
	Filters                *[]Query_Filter `json:"filters,omitempty"`
	Insight                string          `json:"insight,omitempty"`
	Limit                  int32           `json:"limit"`
	Offset                 int32           `json:"offset,omitempty"`
	OrderBy                string          `json:"order_by,omitempty"`
	Scopes                 []string        `json:"scopes,omitempty"`
	Selected_Resource_Type string          `json:"selected_resource_type,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	Cursor                 string          `json:"cursor,omitempty"`
}

type Badge struct {
	// The key and value of a given badge for use with filters, insights, etc.
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Query_Filter struct {
	// The name and configuration of a query filter from the filter registry
	Name   string `json:"name"`
	Config string `json:"config"`
}

type Query_Results struct {
	// The result response from the query provided
	Counts                 map[string]int     `json:"counts,omitempty"`
	Selected_Resource_Type string             `json:"selected_resource_type"`
	Supported_Types        []string           `json:"supported_types"`
	Resources              []Resource_Results `json:"resources"`
	Scopes                 []string           `json:"scopes"`
	Limit                  int32              `json:"limit"`
	Offset                 int32              `json:"offset"`
	Order_By               string             `json:"order_by"`
	Filters                []Query_Filter     `json:"filters"`
	Next_Cursor            string             `json:"next_cursor"`
}

type Resource_Results struct {
	Resource_Type                    string                           `json:"resource_type"`
	Access_Analyzer                  Access_Analyzer                  `json:"accessanalyzer,omitempty"`
	Access_List_Flow_Log             Access_List_Flow_Log             `json:"accesslistflowlog,omitempty"`
	Airflow_Environment              Airflow_Environment              `json:"airflowenvironment,omitempty"`
	API_Accounting_Config            API_Accounting_Config            `json:"apiaccountingconfig,omitempty"`
	App_Runner_Service               App_Runner_Service               `json:"apprunnerservice,omitempty"`
	App_Server                       App_Server                       `json:"appserver,omitempty"`
	Autoscaling_Group                Autoscaling_Group                `json:"autoscalinggroup,omitempty"`
	Autoscaling_Launch_Configuration Autoscaling_Launch_Configuration `json:"autoscalinglaunchconfiguration,omitempty"`
	AWS_Config                       AWS_Config                       `json:"awsconfig,omitempty"`
	Backend_Service                  Backend_Service                  `json:"backendservice,omitempty"`
	Backup_Vault                     Backup_Vault                     `json:"backupvault,omitempty"`
	Batch_Environment                Batch_Environment                `json:"batchenvironment,omitempty"`
	Batch_Pool                       Batch_Pool                       `json:"batchpool,omitempty"`
}

// FUNCTIONS
///////////////////////////////////////////

func (c Client) Query_Resources(q Query) (Query_Results, error) {
	// Queries InsightCloudSec for resources with the given query (using v3-ETL endpoint of the API)

	// Verify Badge_Filter_Operator is appropriate
	if q.Badge_Filter_Operator != "" {
		err := validateBadgeFilterOperator(q.Badge_Filter_Operator)
		if err != nil {
			return Query_Results{}, err
		}
	}

	// Verify required Limit is within requirements
	if q.Limit == 0 {
		q.Limit = 1000
	} else {
		validateQueryLimit(q.Limit)
	}

	data, err := json.Marshal(q)
	if err != nil {
		return Query_Results{}, err
	}

	resp, err := c.makeRequest(http.MethodPost, "/v3/public/resource/etl-query", bytes.NewBuffer(data))
	if err != nil {
		return Query_Results{}, err
	}

	var ret Query_Results
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Query_Results{}, err
	}

	return ret, nil
}

// Validation Function for Query.Badge_Filter_Operator
func validateBadgeFilterOperator(b string) error {
	if strings.ToUpper(b) != "OR" || strings.ToUpper(b) != "AND" {
		return ValidationError{
			ItemToValidate: "BadgeFilterOperator",
			ExpectedValues: []string{"OR", "AND"},
		}
	} else {
		return nil
	}
}

// Validation Function for Query.Limit
func validateQueryLimit(l int32) error {
	if l < 0 || l > 1000 {
		return ValidationError{
			ItemToValidate: "Limit",
			ExpectedValues: []string{"0-1000"},
		}
	} else {
		return nil
	}
}
