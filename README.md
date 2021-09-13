# insightcloudsec
Go Module for Interacting with InsightCloudSec API

### Examples


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

	clouds, err := c.List_Clouds()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, cloud := range clouds.Clouds {
		fmt.Printf("Name: %s\n", cloud.Name)
		regions, _ := c.List_Cloud_Regions(cloud)
		fmt.Println("Regions:")
		for _, region := range regions.Regions {
			fmt.Printf("- %s\n", region.Name)
		}
	}
}
```
</details>

