package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceGroups_Create(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/resourcegroup/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("resourceGroups/createResourceGroupResponse.json"))
	})

	resp, err := client.ResourceGroups.Create("Test Resource Group", "Would you believe it's for grouping things?")
	assert.NoError(t, err)
	assert.Equal(t, "resourcegroup:3:", resp.ID)
	assert.Equal(t, "Test Resource Group", resp.Name)
	assert.Equal(t, 3, resp.ResourceGroupID)
	assert.Equal(t, "Would you believe it's for grouping things?", resp.Description)
	assert.Equal(t, "system", resp.Category)
	assert.Equal(t, "2022-04-10 05:59:34", resp.CreationTime)
	assert.Equal(t, "user", resp.GroupType)
	assert.Equal(t, "organization", resp.OwnerType)
	assert.Equal(t, []string{}, resp.NestedResourceGroups)
	teardown()
}

func TestResourceGroups_AddToGroup(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/resourcegroups/resources/add", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var req ResourcesToGroup
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, req, ResourcesToGroup{
			ResourceIDs:      []string{"instance:20:us-east-1:i-0000a0b11cd33e4:"},
			ResourceGroupIDs: []string{"resourcegroup:3:"},
		})
	})

	err := client.ResourceGroups.AddToGroup([]string{"instance:20:us-east-1:i-0000a0b11cd33e4:"}, "resourcegroup:3:")
	assert.NoError(t, err)
	teardown()
}

func TestResourceGroups_Delete(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/resources/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var req ResourceGroupIDsList
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, req, ResourceGroupIDsList{
			ResourceIDs: []string{"instance:20:us-east-1:i-0000a0b11cd33e4:"},
		})
	})

	err := client.ResourceGroups.Delete([]string{"instance:20:us-east-1:i-0000a0b11cd33e4:"})
	assert.NoError(t, err)
	teardown()
}
