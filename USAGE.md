<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/planetscale"
    "github.com/speakeasy-sdks/planetscale/pkg/models/shared"
    "github.com/speakeasy-sdks/planetscale/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.CreateADatabaseRequest{
        RequestBody: &operations.CreateADatabaseRequestBody{
            Name: "corrupti",
            Notes: "provident",
        },
        Organization: "distinctio",
    }

    ctx := context.Background()
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