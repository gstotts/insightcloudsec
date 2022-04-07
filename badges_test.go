package insightcloudsec

import (
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

func TestBadges_Create(t *testing.T) {}
