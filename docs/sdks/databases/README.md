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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `organization`                                                                                  | *string*                                                                                        | :heavy_check_mark:                                                                              | The name of the organization the database belongs to                                            |
| `requestBody`                                                                                   | [*operations.CreateADatabaseRequestBody](../../models/operations/createadatabaserequestbody.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |


### Response

**[*operations.CreateADatabaseResponse](../../models/operations/createadatabaseresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the database                              |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the database belongs to  |


### Response

**[*operations.DeleteADatabaseResponse](../../models/operations/deleteadatabaseresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the database                              |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the database belongs to  |


### Response

**[*operations.GetADatabaseResponse](../../models/operations/getadatabaseresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the database belongs to       |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListDatabasesResponse](../../models/operations/listdatabasesresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The name of the database                                   |
| `organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the database belongs to       |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListDatabasePromotionRequestsResponse](../../models/operations/listdatabasepromotionrequestsresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The name of the database                                   |
| `organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the database belongs to       |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListReadOnlyRegionsResponse](../../models/operations/listreadonlyregionsresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The name of the database                                   |
| `organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the database belongs to       |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListDatabaseRegionsResponse](../../models/operations/listdatabaseregionsresponse.md), error**


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
            APIKeyHeader: "",
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

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |
| `name`                                                                                                        | *string*                                                                                                      | :heavy_check_mark:                                                                                            | The name of the database                                                                                      |
| `organization`                                                                                                | *string*                                                                                                      | :heavy_check_mark:                                                                                            | The name of the organization the database belongs to                                                          |
| `requestBody`                                                                                                 | [*operations.UpdateDatabaseSettingsRequestBody](../../models/operations/updatedatabasesettingsrequestbody.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |


### Response

**[*operations.UpdateDatabaseSettingsResponse](../../models/operations/updatedatabasesettingsresponse.md), error**

