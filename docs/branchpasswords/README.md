# BranchPasswords

## Overview


<p>API endpoints for managing Branch Passwords.</p>


### Available Operations

* [Create](#create) - Create a branch password
* [Delete](#delete) - Delete a branch password
* [Get](#get) - Get a branch password
* [List](#list) - List branch passwords
* [Renew](#renew) - Renew a branch password
* [Update](#update) - Update a branch password

## Create


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `connect_production_branch`, `create_branch_password`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `manage_passwords` |
| Database | `manage_passwords` |
| Branch | `manage_passwords` |

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
    res, err := s.BranchPasswords.Create(ctx, "illum", "vel", "error", &operations.CreateABranchPasswordRequestBody{
        ReadOnlyRegionID: planetscale.String("deserunt"),
        Role: operations.CreateABranchPasswordRequestBodyRoleWriter.ToPointer(),
        TTL: planetscale.Float64(4375.87),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateABranchPassword201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Delete


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `delete_production_branch_password`, `delete_branch_password`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |

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
    res, err := s.BranchPasswords.Delete(ctx, "magnam", "debitis", "ipsa", "delectus")
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
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `manage_passwords` |
| Database | `manage_passwords` |
| Branch | `manage_passwords` |

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
    res, err := s.BranchPasswords.Get(ctx, operations.GetABranchPasswordRequest{
        Branch: "tempora",
        Database: "suscipit",
        ID: "7cc8796e-d151-4a05-9fc2-ddf7cc78ca1b",
        Organization: "officia",
        ReadOnlyRegionID: planetscale.String("occaecati"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetABranchPassword200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## List


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `manage_passwords` |
| Database | `manage_passwords` |
| Branch | `manage_passwords` |

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
    res, err := s.BranchPasswords.List(ctx, operations.ListBranchPasswordsRequest{
        Branch: "fugit",
        Database: "deleniti",
        Organization: "hic",
        Page: planetscale.Float64(7586.16),
        PerPage: planetscale.Float64(5218.48),
        ReadOnlyRegionID: planetscale.String("beatae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListBranchPasswords200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Renew


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `connect_production_branch`, `create_branch_password`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `manage_passwords` |
| Database | `manage_passwords` |
| Branch | `manage_passwords` |

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
    res, err := s.BranchPasswords.Renew(ctx, operations.RenewABranchPasswordRequest{
        RequestBody: &operations.RenewABranchPasswordRequestBody{
            ReadOnlyRegionID: planetscale.String("commodi"),
        },
        Branch: "molestiae",
        Database: "modi",
        ID: "2cb73920-5929-4396-bea7-596eb10faaa2",
        Organization: "dolorem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RenewABranchPassword200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Update


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `connect_production_branch`, `create_branch_password`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `manage_passwords` |
| Database | `manage_passwords` |
| Branch | `manage_passwords` |

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
    res, err := s.BranchPasswords.Update(ctx, operations.UpdateABranchPasswordRequest{
        RequestBody: &operations.UpdateABranchPasswordRequestBody{
            DisplayName: "corporis",
            ReadOnlyRegionID: planetscale.String("explicabo"),
        },
        Branch: "nobis",
        Database: "enim",
        ID: "955907af-f1a3-4a2f-a946-7739251aa52c",
        Organization: "sequi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateABranchPassword200ApplicationJSONObject != nil {
        // handle response
    }
}
```
