<div align="center">
    <img src="https://user-images.githubusercontent.com/6267663/230379169-dae2f415-423f-4791-8310-8b3304fd449d.svg" width="100">
    <h1>PlanetScale Go SDK</h1>
   <p>The world’s most advanced serverless MySQL platform</p>
   <a href="https://api-docs.planetscale.com/reference/getting-started-with-planetscale-api"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/planetscale-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/planetscale-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div> 

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/planetscale-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [BranchPasswords](docs/branchpasswords/README.md)

* [Create](docs/branchpasswords/README.md#create) - Create a branch password
* [Delete](docs/branchpasswords/README.md#delete) - Delete a branch password
* [Get](docs/branchpasswords/README.md#get) - Get a branch password
* [List](docs/branchpasswords/README.md#list) - List branch passwords
* [Renew](docs/branchpasswords/README.md#renew) - Renew a branch password
* [Update](docs/branchpasswords/README.md#update) - Update a branch password

### [Branches](docs/branches/README.md)

* [CancelDemotionRequest](docs/branches/README.md#canceldemotionrequest) - Cancel or deny a demotion request
* [Create](docs/branches/README.md#create) - Create a branch
* [CreatePromotionRequest](docs/branches/README.md#createpromotionrequest) - Create a promotion request
* [Delete](docs/branches/README.md#delete) - Delete a branch
* [Demote](docs/branches/README.md#demote) - Demote a branch
* [Get](docs/branches/README.md#get) - Get a branch
* [GetDemotionRequest](docs/branches/README.md#getdemotionrequest) - Get a demotion request
* [GetPromotionRequest](docs/branches/README.md#getpromotionrequest) - Get a promotion request
* [GetSchema](docs/branches/README.md#getschema) - Get a branch schema
* [GetStatus](docs/branches/README.md#getstatus) - Get branch status
* [List](docs/branches/README.md#list) - List branches

### [Databases](docs/databases/README.md)

* [Create](docs/databases/README.md#create) - Create a database
* [Delete](docs/databases/README.md#delete) - Delete a database
* [Get](docs/databases/README.md#get) - Get a database
* [List](docs/databases/README.md#list) - List databases
* [ListPromotionRequests](docs/databases/README.md#listpromotionrequests) - List database promotion requests
* [ListReadOnlyRegions](docs/databases/README.md#listreadonlyregions) - List read-only regions
* [ListRegions](docs/databases/README.md#listregions) - List database regions
* [Update](docs/databases/README.md#update) - Update database settings

### [DeployRequests](docs/deployrequests/README.md)

* [Cancel](docs/deployrequests/README.md#cancel) - Cancel a queued deploy request
* [Close](docs/deployrequests/README.md#close) - Close a deploy request
* [CompleteErroredDeploy](docs/deployrequests/README.md#completeerroreddeploy) - Complete an errored deploy
* [CompleteGatedDeploy](docs/deployrequests/README.md#completegateddeploy) - Complete a gated deploy request
* [CompleteRevert](docs/deployrequests/README.md#completerevert) - Complete a revert
* [Create](docs/deployrequests/README.md#create) - Create a deploy request
* [Get](docs/deployrequests/README.md#get) - Get a deploy request
* [GetDeployment](docs/deployrequests/README.md#getdeployment) - Get a deployment
* [GetQueue](docs/deployrequests/README.md#getqueue) - Get a deploy queue
* [List](docs/deployrequests/README.md#list) - List deploy requests
* [ListOperations](docs/deployrequests/README.md#listoperations) - List deploy operations
* [Queue](docs/deployrequests/README.md#queue) - Queue a deploy request
* [SkipRevertPeriod](docs/deployrequests/README.md#skiprevertperiod) - Skip revert period
* [Update](docs/deployrequests/README.md#update) - Update auto-apply for deploy request

### [OAuthApplications](docs/oauthapplications/README.md)

* [Delete](docs/oauthapplications/README.md#delete) - Delete an OAuth token
* [Get](docs/oauthapplications/README.md#get) - Get an OAuth application
* [GetToken](docs/oauthapplications/README.md#gettoken) - Get an OAuth token
* [List](docs/oauthapplications/README.md#list) - List OAuth applications
* [ListTokens](docs/oauthapplications/README.md#listtokens) - List OAuth tokens

### [OAuthTokens](docs/oauthtokens/README.md)

* [Create](docs/oauthtokens/README.md#create) - Create or renew an OAuth token

### [Organizations](docs/organizations/README.md)

* [Get](docs/organizations/README.md#get) - Get an organization
* [List](docs/organizations/README.md#list) - List organizations
* [ListRegions](docs/organizations/README.md#listregions) - List regions for an organization
* [Update](docs/organizations/README.md#update) - Update an organization

### [Users](docs/users/README.md)

* [Current](docs/users/README.md#current) - Get current user
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
