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
			AccountNumber: "123456789101",
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
			AccountNumber: "123456789101",
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
			AccountNumber: "123456789101",
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
			AccountNumber: "123456789101",
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

func TestClouds_validateGCPCloud(t *testing.T) {

}

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
	mux.HandleFunc("/v2/public/clouds/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/listClouds.json"))
	})
	resp, err := client.Clouds.List()
	assert.NoError(t, err)
	assert.Equal(t, "my-cloud-1", resp.Clouds[0].Name)
	assert.Equal(t, "my-cloud-2", resp.Clouds[1].Name)
	assert.Equal(t, "my-azure-cloud", resp.Clouds[2].Name)
	assert.Equal(t, 1, resp.Clouds[0].ID)
	assert.Equal(t, 2, resp.Clouds[1].ID)
	assert.Equal(t, 8, resp.Clouds[2].ID)
	assert.Equal(t, "AWS", resp.Clouds[0].CloudTypeID)
	assert.Equal(t, "AWS", resp.Clouds[1].CloudTypeID)
	assert.Equal(t, "AZURE_ARM", resp.Clouds[2].CloudTypeID)
	assert.Equal(t, "012345678910", resp.Clouds[0].AccountID)
	assert.Equal(t, "012345678911", resp.Clouds[1].AccountID)
	assert.Equal(t, "0aa0a000-0000-0000-aaaa-00aa0aaa000a", resp.Clouds[2].AccountID)
	assert.Equal(t, "2022-04-05 03:35:06", resp.Clouds[0].Created)
	assert.Equal(t, "2022-04-05 03:35:07", resp.Clouds[1].Created)
	assert.Equal(t, "2021-12-16 22:32:04", resp.Clouds[2].Created)
	assert.Equal(t, "PAUSED", resp.Clouds[0].Status)
	assert.Equal(t, "PAUSED", resp.Clouds[1].Status)
	assert.Equal(t, "PAUSED", resp.Clouds[2].Status)
	assert.Equal(t, 2, resp.Clouds[0].BadgeCount)
	assert.Equal(t, 4, resp.Clouds[1].BadgeCount)
	assert.Equal(t, 0, resp.Clouds[2].BadgeCount)
	assert.Equal(t, 424, resp.Clouds[0].ResourceCount)
	assert.Equal(t, 176, resp.Clouds[1].ResourceCount)
	assert.Equal(t, 42, resp.Clouds[2].ResourceCount)
	assert.Equal(t, "2022-04-12 02:06:47", resp.Clouds[0].LastRefreshed)
	assert.Equal(t, "2022-04-12 02:06:47", resp.Clouds[1].LastRefreshed)
	assert.Equal(t, "2022-04-12 02:06:47", resp.Clouds[2].LastRefreshed)
	assert.Equal(t, "master_role", resp.Clouds[0].RoleARN)
	assert.Equal(t, "master_role", resp.Clouds[1].RoleARN)
	assert.Equal(t, "", resp.Clouds[2].RoleARN)
	assert.Equal(t, "divvyorganizationservice:1", resp.Clouds[0].GroupResourceID)
	assert.Equal(t, "divvyorganizationservice:2", resp.Clouds[1].GroupResourceID)
	assert.Equal(t, "divvyorganizationservice:8", resp.Clouds[2].GroupResourceID)
	assert.Equal(t, "divvyorganizationservice:1", resp.Clouds[0].ResourceID)
	assert.Equal(t, "divvyorganizationservice:2", resp.Clouds[1].ResourceID)
	assert.Equal(t, "divvyorganizationservice:8", resp.Clouds[2].ResourceID)
	assert.Equal(t, "idle", resp.Clouds[0].EDHRole)
	assert.Equal(t, "idle", resp.Clouds[1].EDHRole)
	assert.Equal(t, "idle", resp.Clouds[2].EDHRole)
	assert.Equal(t, 1, resp.Clouds[0].StrategyID)
	assert.Equal(t, 1, resp.Clouds[1].StrategyID)
	assert.Equal(t, 4, resp.Clouds[2].StrategyID)
	assert.Equal(t, "o-00aaaaaaaa", resp.Clouds[0].CloudOrgID)
	assert.Equal(t, "o-00aaaaaaaa", resp.Clouds[1].CloudOrgID)
	assert.Equal(t, "", resp.Clouds[2].CloudOrgID)
	assert.Equal(t, "o-00aaaaaaaa", resp.Clouds[0].CloudOrgDomainName)
	assert.Equal(t, "o-00aaaaaaaa", resp.Clouds[1].CloudOrgDomainName)
	assert.Equal(t, "", resp.Clouds[2].CloudOrgDomainName)
	assert.Equal(t, "AWS Org Test", resp.Clouds[0].CloudOrgNickname)
	assert.Equal(t, "AWS Org Test", resp.Clouds[1].CloudOrgNickname)
	assert.Equal(t, "", resp.Clouds[2].CloudOrgNickname)
	teardown()
}

func TestClouds_ListHarvestingStrategies(t *testing.T) {
	setup()

	teardown()
}

func TestClouds_ListRegions(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/cloud/divvyorganizationservice:8/regions/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/listRegions.json"))
	})

	resp, err := client.Clouds.ListRegions(Cloud{
		ResourceID: "divvyorganizationservice:8",
	})
	assert.NoError(t, err)
	assert.Equal(t, "brazilsouth", resp.Regions[3].ID)
	assert.Equal(t, "brazilsouth", resp.Regions[3].Name)
	assert.Equal(t, "serviceregion:8:brazilsouth:", resp.Regions[3].ResourceID)
	assert.Equal(t, "ACTIVE", resp.Regions[3].Status)
	assert.Equal(t, float32(1.0), resp.Regions[3].HarvestRateMultiplier)
	teardown()
}

func TestClouds_ListTypes(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/cloudtypes/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/listCloudTypes.json"))
	})
	resp, err := client.Clouds.ListTypes()
	assert.NoError(t, err)
	// Verifying just one since all utilze same json marshalling
	assert.Equal(t, "AliCloud", resp.CloudTypes[0].Name)
	assert.Equal(t, "ALICLOUD", resp.CloudTypes[0].ID)
	assert.Equal(t, "public", resp.CloudTypes[0].Access)
	teardown()
}

func TestClouds_GetByName(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/clouds/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/listClouds.json"))
	})
	resp, err := client.Clouds.GetByName("my-cloud-2")
	assert.NoError(t, err)
	assert.Equal(t, "012345678911", resp.AccountID)
	teardown()
}

func TestClouds_GetByID(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/clouds/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/listClouds.json"))
	})

	resp, err := client.Clouds.GetByID(8)
	assert.NoError(t, err)
	assert.Equal(t, "my-azure-cloud", resp.Name)
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
