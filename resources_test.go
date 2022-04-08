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
	assert.Equal(t, "instance", resource.Resources[0].Resource_Type)
	assert.Equal(t, "instance:18:us-east-1:i-12300000000000:", resource.Resources[0].Instance.Common.Resource_ID)
	assert.Equal(t, "mega_maid", resource.Resources[0].Instance.Common.Resource_Name)
	assert.Equal(t, "instance", resource.Resources[0].Instance.Common.Type)
	assert.Equal(t, "AWS", resource.Resources[0].Instance.Common.Cloud)
	assert.Equal(t, "Spaceballs", resource.Resources[0].Instance.Common.Account)
	assert.Equal(t, "1234567891011", resource.Resources[0].Instance.Common.Account_ID)
	assert.Equal(t, "PAUSED", resource.Resources[0].Instance.Common.Account_Status)
	assert.Equal(t, 18, resource.Resources[0].Instance.Common.Organization_Service_ID)
	assert.Equal(t, "us-east-1c", resource.Resources[0].Instance.Common.Availability_Zone)
	assert.Equal(t, "us-east-1", resource.Resources[0].Instance.Common.Region)
	assert.Equal(t, "2022-02-18 05:35:47", resource.Resources[0].Instance.Common.Creation_Timestamp)
	assert.Equal(t, "2022-04-05 04:13:31", resource.Resources[0].Instance.Common.Discovered_Timestamp)
	assert.Equal(t, "2022-04-05 04:13:31", resource.Resources[0].Instance.Common.Modified_Timestamp)
	assert.Equal(t, "arn:aws:ec2:us-east-1:1234567891011:instance/i-12300000000000", resource.Resources[0].Instance.Common.Namespace_ID)
	assert.Equal(t, map[string]string{"Name": "mega_maid"}, resource.Resources[0].Instance.Common.Tags)
	assert.Equal(t, "i-12300000000000", resource.Resources[0].Instance.Instance_ID)
	assert.Equal(t, "t2.micro", resource.Resources[0].Instance.Instance_Type)
	assert.Equal(t, "2022-02-18T05:35:47Z", resource.Resources[0].Instance.Launch_Time)
	assert.Equal(t, "linux", resource.Resources[0].Instance.Platform)
	assert.Equal(t, "running", resource.Resources[0].Instance.State)
	assert.Equal(t, "ebs", resource.Resources[0].Instance.Root_Device_Type)
	assert.Equal(t, "mega_maid_key", resource.Resources[0].Instance.Key_Name)
	assert.Equal(t, "ami-123456789101112", resource.Resources[0].Instance.Image_ID)
	assert.Equal(t, "10.1.2.3", resource.Resources[0].Instance.Private_IP_Address)

	assert.ElementsMatch(t, []string{"divvyorganizationservice:18", "resourcegroup:1:", "resourcegroup:2:"}, resource.Scopes)
	assert.Equal(t, int32(1000), resource.Limit)
	assert.Equal(t, int32(2), resource.Offset)
	assert.Equal(t, "", resource.Order_By)
	assert.Equal(t, []Query_Filter{}, resource.Filters)
	assert.Equal(t, "", resource.Next_Cursor)
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
