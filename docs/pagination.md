Some of the endpoints provided are paginated and the Meta struct needs to be utilised to fetch all available
data. 

Below is an example of fetching paginated country data and looping over all available pages:

```go
package main

import (
"context"
"fmt"
"github.com/myuz1/sportmonks-v3-test"
)

func main() {
	client := sportmonks.NewDefaultHTTPClient("YOUR_TOKEN_GOES_HERE")

	includes := []string{"leagues"}

	_, meta, err := client.Countries(context.Background(), 1, includes)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for i := 1; i <= meta.Pagination.TotalPages; i++ {
		countries, _, err := client.Countries(context.Background(), i, includes)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// Do something with countries variable
	}
}
```
