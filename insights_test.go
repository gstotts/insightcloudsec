package insightcloudsec

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsights_Create(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/insights/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	// Shoud return error as Severity is not set
	assert.Error(t, (_, err := client.Insights.Create(Insight{
		Name:          "Test Insight",
		Description:   "Description goes here.",
		Tags:          nil,
		ResourceTypes: []string{"divvyorganizationservice"},
		Filters: []InsightFilter{{
			Name: "divvy.filter.cloud_trail_in_all_regions",
		}},
		Badges: nil,
	})))
	// Should return error as filters are not set
	assert.Error(t, client.Insights.Create(Insight{
		Name:          "Test Insight",
		Description:   "Description goes here.",
		Severity:      INSIGHT_SEVERITY_MINOR,
		Tags:          nil,
		ResourceTypes: []string{"divvyorganizationservice"},
		Badges:        nil,
	}))
	// Valid
	assert.NoError(t, client.Insights.Create(Insight{
		Name:          "Test Insight",
		Description:   "Description goes here.",
		Severity:      INSIGHT_SEVERITY_MINOR,
		Tags:          nil,
		ResourceTypes: []string{"divvyorganizationservice"},
		Filters: []InsightFilter{{
			Name: "divvy.filter.cloud_trail_in_all_regions",
		}},
		Badges: nil,
	}))
	teardown()
}

func TestInsights_Delete(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/insights/12/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.Insights.Delete(12)
	assert.NoError(t, err)
	teardown()

	// Non-HTTP-200 Response should thrown the error
	setup()
	mux.HandleFunc("/v2/public/insights/12/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	})
	err = client.Insights.Delete(12)
	assert.Error(t, err)
	teardown()
}

func TestInsights_GetInsight(t *testing.T) {}

func TestInsights_GetInsight7Days(t *testing.T) {}

func TestInsights_List(t *testing.T) {}

func TestInsights_ListPacks(t *testing.T) {}
