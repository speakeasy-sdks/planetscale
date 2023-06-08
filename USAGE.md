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
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Databases.Create(ctx, "corrupti", &operations.CreateADatabaseRequestBody{
        Name: "Kelvin Sporer",
        Notes: planetscale.String("corrupti"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateADatabase201ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->