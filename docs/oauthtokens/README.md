# OAuthTokens

## Overview


<p>API endpoints for managing OAuth tokens.</p>


### Available Operations

* [Create](#create) - Create or renew an OAuth token

## Create


<p>Create an OAuth token from an authorization grant code, or refresh an OAuth token from a refresh token</p>

### Authorization
A service token   must have the following access   in order to use this API endpoint:

**Service Token Accesses**
  `write_oauth_tokens`



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
    res, err := s.OAuthTokens.Create(ctx, "architecto", "magnam", &operations.CreateOrRenewAnOauthTokenRequestBody{
        ClientID: "et",
        ClientSecret: "excepturi",
        Code: planetscale.String("ullam"),
        GrantType: operations.CreateOrRenewAnOauthTokenRequestBodyGrantTypeRefreshToken,
        RedirectURI: planetscale.String("quos"),
        RefreshToken: planetscale.String("sint"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrRenewAnOauthToken200ApplicationJSONObject != nil {
        // handle response
    }
}
```
