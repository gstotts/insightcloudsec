# insightcloudsec
Go Module for Interacting with InsightCloudSec API

### Examples

<details><summary>List Clouds</summary>

```go
package main

import (
	"fmt"

	"github.com/gstotts/insightcloudsec"
)

func main() {
	// Get a client
	c, err := insightcloudsec.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	clouds, err := c.ListClouds()
	if err != nil {
		fmt.Println(err)
	}
	for _, cloud := range clouds {
		fmt.Println(cloud.Name)
	}
}
```
</details>
<details><summary>List Cloud Types</summary>

```go
package main

import (
	"fmt"

	"github.com/gstotts/insightcloudsec"
)

func main() {
	// Get a client
	c, err := insightcloudsec.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	types, err := c.ListCloudTypes()
	if err != nil {
		fmt.Println(err)
	}
	for _, t := range types {
		fmt.Println(t.Name)
	}
}
```
</details>
<details><summary>List Cloud Regions</summary>
	
```go
package main

import (
	"fmt"
	"os"

	"github.com/gstotts/insightcloudsec"
)

func main() {
	// Get a client
	c, err := insightcloudsec.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	clouds, err := c.ListClouds()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, cloud := range clouds.Clouds {
		fmt.Printf("Name: %s\n", cloud.Name)
		regions, _ := c.ListCloudRegions(cloud)
		fmt.Println("Regions:")
		for _, region := range regions {
			fmt.Printf("- %s\n", region.Name)
		}
	}
}
```
</details>
<details><summary>List Harvesting Strategies</summary>

```go
package main

import (
	"fmt"
	"os"

	"github.com/gstotts/insightcloudsec"
)

func main() {
	// Get a client
	c, err := insightcloudsec.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	hs, err := c.ListHarvestingStrategies()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, s := range hs {
		fmt.Printf("Name: %s\n", s.Name)
	}
}
```
</details>
<details><summary>Create Insight</summary>

```go
package main

import (
	"github.com/gstotts/insightcloudsec"
)

func main() {

	filter := insightcloudsec.InsightFilter{}
	filter.Name = "divvy.filter.cloud_trail_in_all_regions"
	insight := insightcloudsec.Insight{}
	insight.Name = "Test Divvy Insight"
	insight.Description = "Test Divvy Insight Description"
	insight.Severity = insightcloudsec.INSIGHT_SEVERITY_MINOR
	insight.Tags = nil
	insight.Badges = nil
	insight.BadgeFilterOperator = ""
	insight.ResourceTypes = []string{"divvyorganizationservice"}
	insight.Filters = []insightcloudsec.InsightFilter{filter}

	ics, _ := insightcloudsec.NewClient()
	ics.CreateInsight(insight)

}

```
</details>
<details><summary>List Insights</summary>
	
```go
package main

import (
	"fmt"
	"os"

	"github.com/gstotts/insightcloudsec"
)

func main() {
	// Get a client
	c, err := insightcloudsec.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	insights, err := c.ListInsights()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, insight := range insights {
		fmt.Printf("       Name: %s\n", insight.Name)
		fmt.Printf("Description: %s\n\n", insight.Description)
	}
}
```
</details>

<details><summary>Add Azure Cloud</summary>

```go
package main

import (
	"fmt"

	"github.com/gstotts/insightcloudsec"
)

func main() {
	ics, _ := insightcloudsec.NewClient()

	azure_cloud := insightcloudsec.AzureCloudAccount{
		CreationParameters: insightcloudsec.CloudAccountParameters{
			CloudType:      insightcloudsec.AZURE_CLOUD_TYPE,
			AuthType:       insightcloudsec.STANDARD_AUTH,
			Name:           "Azure Test",
			ApiKeyOrCert:   "1111111111111111111111111111",
			TenantID:       "01234567-1234-1234-1234-012345678901",
			SubscriptionID: "01234567-1234-1234-1234-012345678901",
			AppID:          "01234567-1234-1234-1234-012345678901",
		},
	}
	account, err := ics.AddAzureCloud(azure_cloud)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}

```
</details>