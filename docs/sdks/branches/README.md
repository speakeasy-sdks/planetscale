# Branches

## Overview


<p>API endpoints for managing Branches.</p>


### Available Operations

* [CancelDemotionRequest](#canceldemotionrequest) - Cancel or deny a demotion request
* [Create](#create) - Create a branch
* [CreatePromotionRequest](#createpromotionrequest) - Create a promotion request
* [Delete](#delete) - Delete a branch
* [Demote](#demote) - Demote a branch
* [Get](#get) - Get a branch
* [GetDemotionRequest](#getdemotionrequest) - Get a demotion request
* [GetPromotionRequest](#getpromotionrequest) - Get a promotion request
* [GetSchema](#getschema) - Get a branch schema
* [GetStatus](#getstatus) - Get branch status
* [List](#list) - List branches

## CancelDemotionRequest


<p>Cancels or denies a demotion request for a database branch</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `demote_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `demote_branches` |
| Database | `demote_branches` |

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
    res, err := s.Branches.CancelDemotionRequest(ctx, "tenetur", "ipsam", "id")
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
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.CancelOrDenyADemotionRequestResponse](../../models/operations/cancelordenyademotionrequestresponse.md), error**


## Create


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `create_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `write_branches` |
| Database | `write_branches` |

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
    res, err := s.Branches.Create(ctx, "possimus", "aut", &operations.CreateABranchRequestBody{
        BackupID: planetscale.String("quasi"),
        Name: "Dr. Jake Pacocha",
        ParentBranch: "vero",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateABranch201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `database`                                                                                  | *string*                                                                                    | :heavy_check_mark:                                                                          | The name of the database the branch belongs to                                              |
| `organization`                                                                              | *string*                                                                                    | :heavy_check_mark:                                                                          | The name of the organization the branch belongs to                                          |
| `requestBody`                                                                               | [*operations.CreateABranchRequestBody](../../models/operations/createabranchrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |


### Response

**[*operations.CreateABranchResponse](../../models/operations/createabranchresponse.md), error**


## CreatePromotionRequest


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `connect_production_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `promote_branches` |
| Database | `promote_branches` |

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
    res, err := s.Branches.CreatePromotionRequest(ctx, "nihil", "praesentium", "voluptatibus")
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAPromotionRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.CreateAPromotionRequestResponse](../../models/operations/createapromotionrequestresponse.md), error**


## Delete


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `delete_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `delete_branches`, `delete_production_branches` |
| Database | `delete_branches`, `delete_production_branches` |
| Branch | `delete_branch` |

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
    res, err := s.Branches.Delete(ctx, "ipsa", "omnis", "voluptate")
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
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.DeleteABranchResponse](../../models/operations/deleteabranchresponse.md), error**


## Demote


<p>Demote a branch from production to development</p>

### Authorization
A   OAuth token must have the following   scopes in order to use this API endpoint:

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `demote_branches` |
| Database | `demote_branches` |

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
    res, err := s.Branches.Demote(ctx, 7392.64, 199.87, 391.87)
    if err != nil {
        log.Fatal(err)
    }

    if res.DemoteABranch200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *float64*                                             | :heavy_check_mark:                                    | N/A                                                   |
| `name`                                                | *float64*                                             | :heavy_check_mark:                                    | N/A                                                   |
| `organization`                                        | *float64*                                             | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.DemoteABranchResponse](../../models/operations/demoteabranchresponse.md), error**


## Get


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |

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
    res, err := s.Branches.Get(ctx, "reprehenderit", "ut", "maiores")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetABranch200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.GetABranchResponse](../../models/operations/getabranchresponse.md), error**


## GetDemotionRequest


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `connect_branch`, `read_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |

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
    res, err := s.Branches.GetDemotionRequest(ctx, "dicta", "corporis", "dolore")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetADemotionRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.GetADemotionRequestResponse](../../models/operations/getademotionrequestresponse.md), error**


## GetPromotionRequest


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `connect_production_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `promote_branches` |
| Database | `promote_branches` |

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
    res, err := s.Branches.GetPromotionRequest(ctx, "iusto", "dicta", "harum")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAPromotionRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.GetAPromotionRequestResponse](../../models/operations/getapromotionrequestresponse.md), error**


## GetSchema


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |

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
    res, err := s.Branches.GetSchema(ctx, "enim", "accusamus", "commodi", "repudiandae")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetABranchSchema200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |
| `keyspace`                                            | **string*                                             | :heavy_minus_sign:                                    | If provided, the schema for this keyspace is returned |


### Response

**[*operations.GetABranchSchemaResponse](../../models/operations/getabranchschemaresponse.md), error**


## GetStatus


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |

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
    res, err := s.Branches.GetStatus(ctx, "quae", "ipsum", "quidem")
    if err != nil {
        log.Fatal(err)
    }

    if res.GetBranchStatus200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |


### Response

**[*operations.GetBranchStatusResponse](../../models/operations/getbranchstatusresponse.md), error**


## List


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |

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
    res, err := s.Branches.List(ctx, "molestias", "excepturi", 8651.03, 2653.89)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListBranches200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `database`                                                 | *string*                                                   | :heavy_check_mark:                                         | The name of the database the branch belongs to             |
| `organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the branch belongs to         |
| `page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `perPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |


### Response

**[*operations.ListBranchesResponse](../../models/operations/listbranchesresponse.md), error**

