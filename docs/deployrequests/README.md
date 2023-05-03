# DeployRequests

## Overview


<p>API endpoints for managing database deploy requests.</p>


### Available Operations

* [Cancel](#cancel) - Cancel a queued deploy request
* [Close](#close) - Close a deploy request
* [CompleteErroredDeploy](#completeerroreddeploy) - Complete an errored deploy
* [CompleteGatedDeploy](#completegateddeploy) - Complete a gated deploy request
* [CompleteRevert](#completerevert) - Complete a revert
* [Create](#create) - Create a deploy request
* [Get](#get) - Get a deploy request
* [GetDeployment](#getdeployment) - Get a deployment
* [GetQueue](#getqueue) - Get a deploy queue
* [List](#list) - List deploy requests
* [ListOperations](#listoperations) - List deploy operations
* [Queue](#queue) - Queue a deploy request
* [SkipRevertPeriod](#skiprevertperiod) - Skip revert period
* [Update](#update) - Update auto-apply for deploy request

## Cancel


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.Cancel(ctx, operations.CancelAQueuedDeployRequestRequest{
        Database: "architecto",
        Number: "architecto",
        Organization: "repudiandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CancelAQueuedDeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Close


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `write_deploy_requests` |
| Database | `write_deploy_requests` |

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
    res, err := s.DeployRequests.Close(ctx, operations.CloseADeployRequestRequest{
        RequestBody: &operations.CloseADeployRequestRequestBody{
            State: operations.CloseADeployRequestRequestBodyStateEnumClosed.ToPointer(),
        },
        Database: "ullam",
        Number: "expedita",
        Organization: "nihil",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CloseADeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CompleteErroredDeploy


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.CompleteErroredDeploy(ctx, operations.CompleteAnErroredDeployRequest{
        Database: "repellat",
        Number: "quibusdam",
        Organization: "sed",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompleteAnErroredDeploy200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CompleteGatedDeploy


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.CompleteGatedDeploy(ctx, operations.CompleteAGatedDeployRequestRequest{
        Database: "saepe",
        Number: "pariatur",
        Organization: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompleteAGatedDeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CompleteRevert


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.CompleteRevert(ctx, operations.CompleteARevertRequest{
        Database: "consequuntur",
        Number: "praesentium",
        Organization: "natus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompleteARevert200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Create


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_requests`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `write_deploy_requests` |
| Database | `write_deploy_requests` |

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
    res, err := s.DeployRequests.Create(ctx, operations.CreateADeployRequestRequest{
        RequestBody: &operations.CreateADeployRequestRequestBody{
            Branch: planetscale.String("magni"),
            IntoBranch: planetscale.String("sunt"),
            Notes: planetscale.String("quo"),
            Number: planetscale.String("illum"),
        },
        Database: "pariatur",
        Organization: "maxime",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateADeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Get


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_deploy_requests` |
| Database | `read_deploy_requests` |

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
    res, err := s.DeployRequests.Get(ctx, operations.GetADeployRequestRequest{
        Database: "ea",
        Number: "excepturi",
        Organization: "odit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetADeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetDeployment


<p>Get the deployment for a deploy request</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_deploy_requests` |
| Database | `read_deploy_requests` |

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
    res, err := s.DeployRequests.GetDeployment(ctx, operations.GetADeploymentRequest{
        Database: "ea",
        Number: "accusantium",
        Organization: "ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetADeployment200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetQueue


<p>Get the deploy queue for a database</p>



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
    res, err := s.DeployRequests.GetQueue(ctx, operations.GetADeployQueueRequest{
        Database: "maiores",
        Organization: "quidem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetADeployQueue200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## List


<p>List deploy requests for a database</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_deploy_requests` |
| Database | `read_deploy_requests` |

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
    res, err := s.DeployRequests.List(ctx, operations.ListDeployRequestsRequest{
        Database: "ipsam",
        Organization: "voluptate",
        Page: planetscale.Float64(4200.75),
        PerPage: planetscale.Float64(7220.56),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDeployRequests200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListOperations


<p>List deploy operations for a deploy request</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `read_deploy_requests` |
| Database | `read_deploy_requests` |

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
    res, err := s.DeployRequests.ListOperations(ctx, operations.ListDeployOperationsRequest{
        Database: "eaque",
        Number: "pariatur",
        Organization: "nemo",
        Page: planetscale.Float64(9755.22),
        PerPage: planetscale.Float64(166.27),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDeployOperations200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Queue


### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.Queue(ctx, operations.QueueADeployRequestRequest{
        Database: "fugiat",
        Number: "amet",
        Organization: "aut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.QueueADeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## SkipRevertPeriod


<p>Skips the revert period for a deploy request</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.SkipRevertPeriod(ctx, operations.SkipRevertPeriodRequest{
        Database: "cumque",
        Number: "corporis",
        Organization: "hic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SkipRevertPeriod200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Update


<p>Enables or disabled the auto-apply setting for a deploy request</p>

### Authorization
A service token or OAuth token must have the following access or scopes in order to use this API endpoint:

**Service Token Accesses**
  `read_deploy_request`, `create_deploy_request`

**OAuth Scopes**

  | Resource | Scopes |
| :------- | :---------- |
| Organization | `deploy_deploy_requests` |
| Database | `deploy_deploy_requests` |

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
    res, err := s.DeployRequests.Update(ctx, operations.UpdateAutoApplyForDeployRequestRequest{
        Database: "libero",
        Number: "nobis",
        Organization: "dolores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAutoApplyForDeployRequest200ApplicationJSONObject != nil {
        // handle response
    }
}
```
