package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadges_createBadgesFromMap(t *testing.T) {
	want := []Badge{
		{
			Key:            "SKU",
			Value:          "1234-1234-1234444",
			Auto_Generated: false,
		},
		{
			Key:            "Name",
			Value:          "Test Badging 453",
			Auto_Generated: false,
		},
		{
			Key:            "Dog",
			Value:          "Cat",
			Auto_Generated: false,
		},
	}

	input := map[string]string{"SKU": "1234-1234-1234444", "Name": "Test Badging 453", "Dog": "Cat"}
	results := createBadgesFromMap(input)

	assert.ElementsMatch(t, want, results)
}

func TestBadges_Create(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/badges/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	})

	err := client.Badges.Create([]string{"divvyorganizationservice:1"}, map[string]string{"BadgeKey1": "BadgeValue1", "BadgeKey2": "BadgeValue2"})

	assert.NoError(t, err)
	teardown()
}

func TestBadges_Update(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/badges/divvyserviceorganization:5/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		var body_object BadgeRequest
		err := json.NewDecoder(r.Body).Decode(&body_object)
		assert.NoError(t, err)
		assert.Equal(t,
			BadgeRequest{
				Org_Resource_IDs: []string{"divvyorganizationservice:5"},
				Badges:           []Badge{{Key: "name", Value: "barf"}},
			},
			body_object)
	})

	client.Badges.Update("divvyorganizationservice:5", map[string]string{"name": "barf"})

	teardown()
}

func TestBadges_Delete(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/badges/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		var body_object BadgeRequest
		err := json.NewDecoder(r.Body).Decode(&body_object)
		assert.NoError(t, err)
		assert.Equal(t,
			BadgeRequest{
				Org_Resource_IDs: []string{"divvyorganizationservice:5"},
				Badges:           []Badge{{Key: "Name", Value: "Barf"}},
			},
			body_object)
	})

	client.Badges.Delete([]string{"divvyorganizationservice:5"}, map[string]string{"Name": "Barf"})
	teardown()
}

func TestBadgges_ListResourceBadges(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/badges/divvyorganizationservice:1/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("badges/listResourceBadges.json"))
	})

	list, err := client.Badges.ListResourceBadges("divvyorganizationservice:1")
	want := []Badge{
		{
			Key:            "BadgeKey1",
			Value:          "BadgeValue1",
			Auto_Generated: false,
		},
		{
			Key:            "BadgeKey2",
			Value:          "BadgeValue2",
			Auto_Generated: false,
		},
	}

	assert.NoError(t, err)
	assert.ElementsMatch(t, want, list)
	teardown()
}

func TestBadgges_ListCloudsWithBadges(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/badge/clouds/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("badges/listCloudsWithBadges.json"))
	})

	list, err := client.Badges.ListCloudsWithBadges(nil)
	want := []CloudBadges{
		{
			Resource_ID: "divvyorganizationservice:1",
			Name:        "lonestar",
		},
		{
			Resource_ID: "divvyorganizationservice:2",
			Name:        "darkHelmet123",
		},
		{
			Resource_ID: "divvyorganizationservice:4",
			Name:        "the_schwartz",
		},
	}

	assert.NoError(t, err)
	assert.ElementsMatch(t, want, list)
	teardown()
}

func TestBadgges_ListResourceBadgeCount(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/badges/count", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("badges/listResourceBadgeCounts.json"))
	})

	list, err := client.Badges.ListResourceBadgeCount([]string{"divvyorganizationservice:3"})
	want := Resource_Count{
		Resource_Count: []BadgeResourceCount{
			{
				Resource_ID: "divvyorganizationservice:3",
				Count:       4,
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, want, list)
	teardown()
}
