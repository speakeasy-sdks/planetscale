# ListBranchesRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Database`                                                 | *string*                                                   | :heavy_check_mark:                                         | The name of the database the branch belongs to             |
| `Organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the branch belongs to         |
| `Page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `PerPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |