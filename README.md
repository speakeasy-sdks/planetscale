<div align="center">
    <img src="https://user-images.githubusercontent.com/6267663/230379169-dae2f415-423f-4791-8310-8b3304fd449d.svg" width="100">
    <h1>Planetscale Go SDK</h1>
   <p>The world’s most advanced serverless MySQL platform</p>
   <a href="https://api-docs.planetscale.com/reference/getting-started-with-planetscale-api"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/planetscale-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/planetscale-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/planetscale-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/planetscale-go?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/planetscale
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### BranchPasswords

* `Create` - Create a branch password
* `Delete` - Delete a branch password
* `Get` - Get a branch password
* `List` - List branch passwords
* `Renew` - Renew a branch password
* `Update` - Update a branch password

### Branches

* `CancelDemotionRequest` - Cancel or deny a demotion request
* `Create` - Create a branch
* `CreatePromotionRequest` - Create a promotion request
* `Delete` - Delete a branch
* `Demote` - Demote a branch
* `Get` - Get a branch
* `GetDemotionRequest` - Get a demotion request
* `GetPromotionRequest` - Get a promotion request
* `GetSchema` - Get a branch schema
* `GetStatus` - Get branch status
* `List` - List branches

### Databases

* `Create` - Create a database
* `Delete` - Delete a database
* `Get` - Get a database
* `List` - List databases
* `ListPromotionRequests` - List database promotion requests
* `ListReadOnlyRegions` - List read-only regions
* `ListRegions` - List database regions
* `Update` - Update database settings

### DeployRequests

* `Cancel` - Cancel a queued deploy request
* `Close` - Close a deploy request
* `CompleteErroredDeploy` - Complete an errored deploy
* `CompleteGatedDeploy` - Complete a gated deploy request
* `CompleteRevert` - Complete a revert
* `Create` - Create a deploy request
* `Get` - Get a deploy request
* `GetDeployment` - Get a deployment
* `GetQueue` - Get a deploy queue
* `List` - List deploy requests
* `ListOperations` - List deploy operations
* `Queue` - Queue a deploy request
* `SkipRevertPeriod` - Skip revert period
* `Update` - Update auto-apply for deploy request

### OAuthApplications

* `Delete` - Delete an OAuth token
* `Get` - Get an OAuth application
* `GetToken` - Get an OAuth token
* `List` - List OAuth applications
* `ListTokens` - List OAuth tokens

### OAuthTokens

* `Renew` - Create or renew an OAuth token

### Organizations

* `Get` - Get an organization
* `List` - List organizations
* `ListRegions` - List regions for an organization
* `Update` - Update an organization

### Users

* `Current` - Get current user
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
