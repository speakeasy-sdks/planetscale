# Organizations

## Overview


<p>API endpoints for managing organizations.</p>


### Available Operations

* [Get](#get) - Get an organization
* [List](#list) - List organizations
* [ListRegions](#listregions) - List regions for an organization
* [Update](#update) - Update an organization

## Get


### Authorization
A OAuth token must have the following scopes in order to use this API endpoint:

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| User | `read_organizations` |
| Organization | `read_organization` |

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
    res, err := s.Organizations.Get(ctx, "accusantium")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAnOrganization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the organization                          |


### Response

**[*operations.GetAnOrganizationResponse](../../models/operations/getanorganizationresponse.md), error**


## List


### Authorization
A   OAuth token must have the following   scopes in order to use this API endpoint:

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| User | `read_organizations` |

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
    res, err := s.Organizations.List(ctx, 6532.01, 9689.62)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListOrganizations200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListOrganizationsResponse](../../models/operations/listorganizationsresponse.md), error**


## ListRegions


### Authorization
A   OAuth token must have the following   scopes in order to use this API endpoint:

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| User | `read_organizations` |
| Organization | `read_organization` |

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
    res, err := s.Organizations.ListRegions(ctx, "mollitia", 3209.97, 4314.18)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRegionsForAnOrganization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The name of the organization                               |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListRegionsForAnOrganizationResponse](../../models/operations/listregionsforanorganizationresponse.md), error**


## Update


### Authorization
A   OAuth token must have the following   scopes in order to use this API endpoint:

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `write_organization` |

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
    res, err := s.Organizations.Update(ctx, "dolor", &operations.UpdateAnOrganizationRequestBody{
        BillingEmail: planetscale.String("necessitatibus"),
        RequireAdminForProductionAccess: planetscale.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAnOrganization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `name`                                                                                                    | *string*                                                                                                  | :heavy_check_mark:                                                                                        | The name of the organization                                                                              |
| `requestBody`                                                                                             | [*operations.UpdateAnOrganizationRequestBody](../../models/operations/updateanorganizationrequestbody.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |


### Response

**[*operations.UpdateAnOrganizationResponse](../../models/operations/updateanorganizationresponse.md), error**

