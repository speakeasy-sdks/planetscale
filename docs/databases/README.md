# Databases

## Overview


<p>API endpoints for managing databases within an organization.</p>


### Available Operations

* [Create](#create) - Create a database
* [Delete](#delete) - Delete a database
* [Get](#get) - Get a database
* [List](#list) - List databases
* [ListPromotionRequests](#listpromotionrequests) - List database promotion requests
* [ListReadOnlyRegions](#listreadonlyregions) - List read-only regions
* [ListRegions](#listregions) - List database regions
* [Update](#update) - Update database settings

## Create


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `create_databases`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `create_databases` |

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
    req := operations.CreateADatabaseRequest{
        RequestBody: &operations.CreateADatabaseRequestBody{
            Name: "Sophia Wintheiser",
            Notes: planetscale.String("nam"),
        },
        Organization: "id",
    }

    res, err := s.Databases.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateADatabase201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Delete


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `delete_database`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `delete_databases` |
| Database | `delete_database` |

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
    req := operations.DeleteADatabaseRequest{
        Name: "Jaime Will",
        Organization: "nisi",
    }

    res, err := s.Databases.Delete(ctx, req)
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
  `read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_databases` |
| Database | `read_database` |

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
    req := operations.GetADatabaseRequest{
        Name: "Ada Moen IV",
        Organization: "magnam",
    }

    res, err := s.Databases.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetADatabase200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## List


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_databases` |

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
    req := operations.ListDatabasesRequest{
        Organization: "distinctio",
        Page: planetscale.Float64(6601.74),
        PerPage: planetscale.Float64(2879.91),
    }

    res, err := s.Databases.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDatabases200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListPromotionRequests


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_databases` |
| Database | `read_database` |

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
    req := operations.ListDatabasePromotionRequestsRequest{
        Name: "Laurie Mosciski",
        Organization: "vero",
        Page: planetscale.Float64(1354.74),
        PerPage: planetscale.Float64(1028.63),
    }

    res, err := s.Databases.ListPromotionRequests(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDatabasePromotionRequests200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListReadOnlyRegions


<p>List read-only regions for the database’s default branch</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |

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
    req := operations.ListReadOnlyRegionsRequest{
        Name: "Frances Marks",
        Organization: "quos",
        Page: planetscale.Float64(5743.25),
        PerPage: planetscale.Float64(336.25),
    }

    res, err := s.Databases.ListReadOnlyRegions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListReadOnlyRegions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListRegions


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_databases` |
| Database | `read_database` |

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
    req := operations.ListDatabaseRegionsRequest{
        Name: "Abel O'Hara",
        Organization: "dolor",
        Page: planetscale.Float64(8965.47),
        PerPage: planetscale.Float64(1412.64),
    }

    res, err := s.Databases.ListRegions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDatabaseRegions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Update


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `write_database`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `write_databases` |
| Database | `write_database` |

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
    req := operations.UpdateDatabaseSettingsRequest{
        RequestBody: &operations.UpdateDatabaseSettingsRequestBody{
            AllowDataBranching: planetscale.Bool(false),
            AutomaticMigrations: planetscale.Bool(false),
            DefaultBranch: planetscale.String("nemo"),
            InsightsRawQueries: planetscale.Bool(false),
            MigrationFramework: planetscale.String("quasi"),
            MigrationTableName: planetscale.String("iure"),
            Notes: planetscale.String("doloribus"),
            ProductionBranchWebConsole: planetscale.Bool(false),
            RequireApprovalForDeploy: planetscale.Bool(false),
            RestrictBranchRegion: planetscale.Bool(false),
        },
        Name: "Frederick Schoen",
        Organization: "in",
    }

    res, err := s.Databases.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDatabaseSettings200ApplicationJSONObject != nil {
        // handle response
    }
}
```
