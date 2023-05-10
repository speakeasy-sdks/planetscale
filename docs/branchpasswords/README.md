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
    res, err := s.BranchPasswords.Create(ctx, operations.CreateABranchPasswordRequest{
        RequestBody: &operations.CreateABranchPasswordRequestBody{
            ReadOnlyRegionID: planetscale.String("illum"),
            Role: operations.CreateABranchPasswordRequestBodyRoleEnumWriter.ToPointer(),
            TTL: planetscale.Float64(6235.64),
        },
        Branch: "deserunt",
        Database: "suscipit",
        Organization: "iure",
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
    res, err := s.BranchPasswords.Delete(ctx, operations.DeleteABranchPasswordRequest{
        Branch: "magnam",
        Database: "debitis",
        ID: "0f467cc8-796e-4d15-9a05-dfc2ddf7cc78",
        Organization: "porro",
    })
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
        Branch: "dolorum",
        Database: "dicta",
        ID: "ba928fc8-1674-42cb-b392-05929396fea7",
        Organization: "corporis",
        ReadOnlyRegionID: planetscale.String("iste"),
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
        Branch: "iure",
        Database: "saepe",
        Organization: "quidem",
        Page: planetscale.Float64(992.8),
        PerPage: planetscale.Float64(602.25),
        ReadOnlyRegionID: planetscale.String("reiciendis"),
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
            ReadOnlyRegionID: planetscale.String("est"),
        },
        Branch: "mollitia",
        Database: "laborum",
        ID: "2352c595-5907-4aff-9a3a-2fa946773925",
        Organization: "vitae",
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
            DisplayName: "laborum",
            ReadOnlyRegionID: planetscale.String("animi"),
        },
        Branch: "enim",
        Database: "odit",
        ID: "c3f5ad01-9da1-4ffe-b8f0-97b0074f1547",
        Organization: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateABranchPassword200ApplicationJSONObject != nil {
        // handle response
    }
}
```
