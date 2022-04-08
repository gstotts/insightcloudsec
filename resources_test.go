package insightcloudsec

import (
	"fmt"
	"net/http"
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

func TestResources_InstanceQuery(t *testing.T) {
	setup()
	mux.HandleFunc("/v3/public/resource/etl-query", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("resources/queryResponse.json"))
	})

	resource, err := client.Resources.Query(Query{
		Limit:                  1000,
		Selected_Resource_Type: "instance",
	})

	assert.NoError(t, err)
	assert.Equal(t, "instance:18:us-east-1:i-12300000000000:", resource.Resources[0].Instance.Common.Resource_ID)
	assert.Equal(t, "mega_maid", resource.Resources[0].Instance.Common.Resource_Name)
	assert.Equal(t, "Spaceballs", resource.Resources[0].Instance.Common.Account)
	assert.Equal(t, "1234567891011", resource.Resources[0].Instance.Common.Account_ID)
	assert.Equal(t, "instance", resource.Resources[0].Instance.Common.Type)
	assert.Equal(t, "AWS", resource.Resources[0].Instance.Common.Cloud)
	teardown()
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
