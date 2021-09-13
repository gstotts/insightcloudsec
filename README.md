# insightcloudsec
Go Module for Interacting with InsightCloudSec API

### Examples

<details><summary>List Clouds</summary>

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

	cs, err := c.List_Clouds()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, c := range cs.Clouds {
		fmt.Printf("          Name: %s\n", c.Name)
		fmt.Printf("Resource Count: %d\n\n", c.ResourceCount)
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

	hs, err := c.List_Harvesting_Strategies()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, s := range hs.Strategies {
		fmt.Printf("Name: %s\n", s.Name)
	}
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

	insights, err := c.List_Insights()
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
<details><summary>Get Insight</summary>

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

	insightID := 2
	insightSource := "backoffice"

	details, err := c.Get_Insight(insightID, insightSource)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(details.Name)
}	
```
</details>
<details><summary>Get Insight 7 Day Stats</summary>

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

	insightID := 2
	insightSource := "backoffice"

	details, err := c.Get_Insight_7_Days(insightID, insightSource)
	if err != nil {
		fmt.Println(err)
	}
	for date, info := range details {
		fmt.Println(date, info)
	}
}	
```
</details>
