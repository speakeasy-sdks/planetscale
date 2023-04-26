<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/planetscale"
	"github.com/speakeasy-sdks/planetscale/pkg/models/operations"
)

func main() {
    s := planetscale.New(
        planetscale.WithSecurity(shared.Security{
            APIKeyHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateADatabaseRequest{
        RequestBody: &operations.CreateADatabaseRequestBody{
            Name: "Terrence Rau",
            Notes: planetscale.String("nulla"),
        },
        Organization: "corrupti",
    }

    res, err := s.Databases.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateADatabase201ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->