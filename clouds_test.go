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
	mux.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc("/v2/prototype/diagnostics/queues/status/get", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/getQueueStatus.json"))
	})

	resp, err := client.Clouds.QueueStatus()
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.SchedulerInternal)
	assert.Equal(t, 2, resp.Workers)
	assert.Equal(t, []int{0, 0, 0, 0}, []int{resp.P0, resp.P1, resp.P2, resp.P3})
	assert.Equal(t, 0, resp.QueueWait.Count)
	assert.Equal(t, []float32{0, 0, 0}, []float32{resp.QueueWait.Min, resp.QueueWait.Max, resp.QueueWait.Sum})
	assert.Equal(t, []float64{0, 0, 0}, []float64{resp.QueueWait.SumSQ, resp.QueueWait.StdDev, resp.QueueWait.Average})
	assert.Equal(t, 159, resp.QueueWaitP0.Count)
	assert.Equal(t, []float32{0, 2.393037, 93.75721199999997}, []float32{resp.QueueWaitP0.Min, resp.QueueWaitP0.Max, resp.QueueWaitP0.Sum})
	assert.Equal(t, []float64{74.16969662233997, 0.34571581342216895, 0.5896680000000004}, []float64{resp.QueueWaitP0.SumSQ, resp.QueueWaitP0.StdDev, resp.QueueWaitP0.Average})
	assert.Equal(t, 0.11951942365015193, resp.QueueWaitP0.Variance)
	assert.Equal(t, 1.002351, resp.QueueWaitP0.Current)
	// The remaining all use the same embeded structs for stats so just verifying can grab one value from each
	assert.Equal(t, []float64{0, 0, 1.5798686302096074}, []float64{resp.QueueWaitP2.StdDev, resp.QueueWaitP3.StdDev, resp.QueueWaitAll.StdDev})
	assert.Equal(t, []int{482, 159, 322, 0, 0}, []int{resp.ProcessTime.Count, resp.ProcessTimeP0.Count, resp.ProcessTimeP1.Count, resp.ProcessTimeP2.Count, resp.ProcessTimeP3.Count})
	assert.Equal(t, "BackOfficeInsightHarvester", resp.SlowestJobs[0].Name)
	assert.Equal(t, 35.0, resp.SlowestJobs[0].Duration)
	teardown()
}
