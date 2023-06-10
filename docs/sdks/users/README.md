# Users

## Overview


<p>API endpoints for fetching user information.</p>


### Available Operations

* [Current](#current) - Get current user

## Current


<p>Get the user associated with this service token</p>



### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/planetscale"
)

func main() {
    s := planetscale.New(
        planetscale.WithSecurity(shared.Security{
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Users.Current(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCurrentUser200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCurrentUserResponse](../../models/operations/getcurrentuserresponse.md), error**

