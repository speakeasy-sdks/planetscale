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
    res, err := s.Databases.Create(ctx, "praesentium", &operations.CreateADatabaseRequestBody{
        Name: "Grady Botsford",
        Notes: planetscale.String("veritatis"),
    })
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
    res, err := s.Databases.Delete(ctx, "itaque", "incidunt")
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
    res, err := s.Databases.Get(ctx, "enim", "consequatur")
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
    res, err := s.Databases.List(ctx, "est", 8423.42, 1317.97)
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
    res, err := s.Databases.ListPromotionRequests(ctx, "deserunt", "distinctio", 8413.86, 2894.06)
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
    res, err := s.Databases.ListReadOnlyRegions(ctx, "modi", "qui", 3978.21, 5865.13)
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
    res, err := s.Databases.ListRegions(ctx, "quos", "perferendis", 1649.4, 8289.4)
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
    res, err := s.Databases.Update(ctx, "ipsam", "alias", &operations.UpdateDatabaseSettingsRequestBody{
        AllowDataBranching: planetscale.Bool(false),
        AutomaticMigrations: planetscale.Bool(false),
        DefaultBranch: planetscale.String("fugit"),
        InsightsRawQueries: planetscale.Bool(false),
        MigrationFramework: planetscale.String("dolorum"),
        MigrationTableName: planetscale.String("excepturi"),
        Notes: planetscale.String("tempora"),
        ProductionBranchWebConsole: planetscale.Bool(false),
        RequireApprovalForDeploy: planetscale.Bool(false),
        RestrictBranchRegion: planetscale.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDatabaseSettings200ApplicationJSONObject != nil {
        // handle response
    }
}
```
