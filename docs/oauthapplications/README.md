# OAuthApplications

## Overview


<p>API endpoints for fetching OAuth applications.</p>


### Available Operations

* [Delete](#delete) - Delete an OAuth token
* [Get](#get) - Get an OAuth application
* [GetToken](#gettoken) - Get an OAuth token
* [List](#list) - List OAuth applications
* [ListTokens](#listtokens) - List OAuth tokens

## Delete


### Authorization
A service token   must have the following access   in order to use this API endpoint:

**Service Token Accesses**
  `delete_oauth_tokens`



### Example Usage

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
    req := operations.DeleteAnOauthTokenRequest{
        ApplicationID: "quis",
        Organization: "totam",
        TokenID: "dignissimos",
    }

    res, err := s.OAuthApplications.Delete(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Get


### Authorization
A service token   must have the following access   in order to use this API endpoint:

**Service Token Accesses**
  `read_oauth_applications`



### Example Usage

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
    req := operations.GetAnOauthApplicationRequest{
        ApplicationID: "eaque",
        Organization: "quis",
    }

    res, err := s.OAuthApplications.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAnOauthApplication200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetToken


### Authorization
A service token   must have the following access   in order to use this API endpoint:

**Service Token Accesses**
  `read_oauth_tokens`



### Example Usage

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
    req := operations.GetAnOauthTokenRequest{
        ApplicationID: "nesciunt",
        Organization: "eos",
        TokenID: "perferendis",
    }

    res, err := s.OAuthApplications.GetToken(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAnOauthToken200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## List


### Authorization
A service token   must have the following access   in order to use this API endpoint:

**Service Token Accesses**
  `read_oauth_applications`



### Example Usage

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
    req := operations.ListOauthApplicationsRequest{
        Organization: "dolores",
        Page: planetscale.Float64(7936.98),
        PerPage: planetscale.Float64(4634.51),
    }

    res, err := s.OAuthApplications.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListOauthApplications200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListTokens


<p>List OAuth tokens created by an OAuth application</p>

### Authorization
A service token   must have the following access   in order to use this API endpoint:

**Service Token Accesses**
  `read_oauth_tokens`



### Example Usage

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
    req := operations.ListOauthTokensRequest{
        ApplicationID: "dolor",
        Organization: "vero",
        Page: planetscale.Float64(3453.52),
        PerPage: planetscale.Float64(9441.2),
    }

    res, err := s.OAuthApplications.ListTokens(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListOauthTokens200ApplicationJSONObject != nil {
        // handle response
    }
}
```
