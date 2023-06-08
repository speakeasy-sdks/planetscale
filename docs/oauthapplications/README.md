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
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.OAuthApplications.Delete(ctx, "omnis", "molestiae", "perferendis")
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
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.OAuthApplications.Get(ctx, "nihil", "magnam")
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
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.OAuthApplications.GetToken(ctx, "distinctio", "id", "labore")
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
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.OAuthApplications.List(ctx, "labore", 3834.62, 6180.16)
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
            APIKeyHeader: "",
        }),
    )

    ctx := context.Background()
    res, err := s.OAuthApplications.ListTokens(ctx, "nobis", "eum", 8784.53, 1354.74)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListOauthTokens200ApplicationJSONObject != nil {
        // handle response
    }
}
```
