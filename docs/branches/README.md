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
