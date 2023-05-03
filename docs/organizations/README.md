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
            APIKeyHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Organizations.Get(ctx, operations.GetAnOrganizationRequest{
        Name: "Mr. Lee Funk III",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAnOrganization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Organizations.List(ctx, operations.ListOrganizationsRequest{
        Page: planetscale.Float64(4878.38),
        PerPage: planetscale.Float64(3117.96),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListOrganizations200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Organizations.ListRegions(ctx, operations.ListRegionsForAnOrganizationRequest{
        Name: "Rickey Wolf",
        Page: planetscale.Float64(1796.03),
        PerPage: planetscale.Float64(5424.99),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRegionsForAnOrganization200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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
            APIKeyHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Organizations.Update(ctx, operations.UpdateAnOrganizationRequest{
        RequestBody: &operations.UpdateAnOrganizationRequestBody{
            BillingEmail: planetscale.String("sit"),
            RequireAdminForProductionAccess: planetscale.Bool(false),
        },
        Name: "Stephen Roberts",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAnOrganization200ApplicationJSONObject != nil {
        // handle response
    }
}
```
