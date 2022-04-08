package insightcloudsec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResources_validateBadgeFilterOperator(t *testing.T) {
	// AND and OR are good
	assert.NoError(t, validateBadgeFilterOperator("AND"))
	assert.NoError(t, validateBadgeFilterOperator("and"))
	assert.NoError(t, validateBadgeFilterOperator("aNd"))
	assert.NoError(t, validateBadgeFilterOperator("OR"))
	assert.NoError(t, validateBadgeFilterOperator("or"))
	assert.NoError(t, validateBadgeFilterOperator("oR"))

	// Others should error
	assert.Error(t, validateBadgeFilterOperator("bug"))
	assert.Error(t, validateBadgeFilterOperator("false"))
}

func TestResources_validateQueryLimit(t *testing.T) {
	// Values between 1-1000 are good
	assert.NoError(t, validateQueryLimit(3))
	assert.NoError(t, validateQueryLimit(768))
	// 0 returns error
	assert.Error(t, validateQueryLimit(0))
	// greater than 1000 returns error
	assert.Error(t, validateQueryLimit(1500))
}

func TestResources_Query(t *testing.T) {

}

func TestResources_GetDetails(t *testing.T) {

}

func TestResources_SetOwner(t *testing.T) {

}

func TestResources_GetAssociations(t *testing.T) {

}

func TestResources_ListTags(t *testing.T) {

}

func TestResources_ListSettings(t *testing.T) {

}
