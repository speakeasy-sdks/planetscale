# ListDeployRequestsRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Database`                                                 | *string*                                                   | :heavy_check_mark:                                         | The name of the deploy request's database                  |
| `Organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the deploy request's organization              |
| `Page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `PerPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |