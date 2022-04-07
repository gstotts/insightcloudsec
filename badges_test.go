package insightcloudsec

import (
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

	assert.Equal(t, want, results)
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
	assert.Equal(t, want, list)
	teardown()
}
