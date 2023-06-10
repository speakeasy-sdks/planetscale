# ListDeployOperationsRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Database`                                                 | *string*                                                   | :heavy_check_mark:                                         | The name of the database the deploy request belongs to     |
| `Number`                                                   | *string*                                                   | :heavy_check_mark:                                         | The number of the deploy request                           |
| `Organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the deploy request belongs to |
| `Page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `PerPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |