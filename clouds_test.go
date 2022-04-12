package insightcloudsec

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClouds_validateAWSCloud(t *testing.T) {
	// Wrong Cloud Type
	wrong_cloud := AWSCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:     "GCP",
			AuthType:      "assume_role",
			Name:          "Test Cloud AWS Bad 1",
			AccountNumber: "1234567891011",
			RoleArn:       "",
			Duration:      0,
			ApiKeyOrCert:  "1231241241234123",
			SecretKey:     "1231241241234123",
			SessionName:   "InsightCloudSec Test",
		},
	}
	assert.Error(t, validateAWSCloud(wrong_cloud))

	// No Key or Secret Data
	no_key_or_secret := AWSCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:     "AWS",
			AuthType:      "assume_role",
			Name:          "Test Cloud AWS Bad 2",
			AccountNumber: "1234567891011",
			RoleArn:       "",
			Duration:      0,
			SessionName:   "InsightCloudSec Test",
		},
	}
	assert.Error(t, validateAWSCloud(no_key_or_secret))

	// Other cloud properties set
	other_props := AWSCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:     "AWS",
			AuthType:      "assume_role",
			Name:          "Test Cloud AWS Bad 3",
			AccountNumber: "1234567891011",
			ApiKeyOrCert:  "1231241241234123",
			SecretKey:     "1231241241234123",
			RoleArn:       "",
			TenantID:      "tenant-id",
			AppID:         "app-id",
			Project:       "project-name",
		},
	}
	assert.Error(t, validateAWSCloud(other_props))

	// Passing example
	good := AWSCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:     "AWS",
			AuthType:      "assume_role",
			Name:          "Test Cloud AWS Bad 1",
			AccountNumber: "1234567891011",
			RoleArn:       "",
			Duration:      0,
			ApiKeyOrCert:  "1231241241234123",
			SecretKey:     "1231241241234123",
			SessionName:   "InsightCloudSec Test",
		},
	}
	assert.NoError(t, validateAWSCloud(good))
}

func TestClouds_validateAzureCloud(t *testing.T) {
	// Wrong Cloud Type
	wrong := AzureCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:      "AWS",
			AuthType:       "assume_role",
			Name:           "Test Cloud AWS Bad 1",
			ApiKeyOrCert:   "1231241241234123",
			TenantID:       "tenant_id",
			SubscriptionID: "sub_id",
			AppID:          "app_id",
		},
	}
	assert.Error(t, validateAzureCloud(wrong))

	// Invalid Auth Type
	wrong2 := AzureCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:      "AZURE_ARM",
			AuthType:       "assume_role",
			Name:           "Test Cloud AWS Bad 1",
			ApiKeyOrCert:   "1231241241234123",
			TenantID:       "tenant_id",
			SubscriptionID: "sub_id",
			AppID:          "app_id",
		},
	}
	assert.Error(t, validateAzureCloud(wrong2))

	// Missing required parameters for auth
	wrong3 := AzureCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:      "AZURE_ARM",
			AuthType:       STANDARD_AUTH,
			Name:           "Test Cloud AWS Bad 1",
			TenantID:       "tenant_id",
			SubscriptionID: "sub_id",
			AppID:          "app_id",
		},
	}
	assert.Error(t, validateAzureCloud(wrong3))

	// Other cloud parameters exist
	wrong4 := AzureCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:      "AZURE_ARM",
			AuthType:       STANDARD_AUTH,
			Name:           "Test Cloud AWS Bad 1",
			ApiKeyOrCert:   "1231241241234123",
			RoleArn:        "arn:asdfasdf:asdfasdfasdF:asdfasdfsad/asdfasdf",
			TenantID:       "tenant_id",
			SubscriptionID: "sub_id",
			AppID:          "app_id",
		},
	}
	assert.Error(t, validateAzureCloud(wrong4))

	//Good
	good := AzureCloudAccount{
		CreationParameters: CloudAccountParameters{
			CloudType:      "AZURE_ARM",
			AuthType:       STANDARD_AUTH,
			Name:           "Test Cloud AWS Bad 1",
			ApiKeyOrCert:   "1231241241234123",
			TenantID:       "tenant_id",
			SubscriptionID: "sub_id",
			AppID:          "app_id",
		},
	}
	assert.NoError(t, validateAzureCloud(good))
}

func TestClouds_validateGCPCloud(t *testing.T) {}

func TestClouds_AddAWSCloud(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/cloud/add", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	teardown()
}

func TestClouds_AddAzureCloud(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/cloud/add", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	teardown()
}

func TestClouds_AddGCPCloud(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/cloud/add", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	teardown()
}

func TestClouds_Delete(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/cloud/divvyorganizationservice:1/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.Clouds.Delete("divvyorganizationservice:1")
	assert.NoError(t, err)
	teardown()

	// Test when error returned from API
	setup()
	mux.HandleFunc("/v2/public/cloud/divvyorganizationservice:1/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusForbidden)
	})
	err = client.Clouds.Delete("divvyorganizationservice:1")
	assert.Error(t, err)
	teardown()
}

func TestClouds_Update(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_List(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/cloud/divvyorganizationservice:1/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/listClouds.json"))
	})

	teardown()
}

func TestClouds_ListHarvestingStrategies(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_ListProvisioningClouds(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_ListRegions(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_ListTypes(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_GetByName(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_GetByID(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_QueueStatus(t *testing.T) {
	setup()

	teardown()
}
