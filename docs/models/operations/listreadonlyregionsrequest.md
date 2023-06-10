# ListReadOnlyRegionsRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The name of the database                                   |
| `Organization`                                             | *string*                                                   | :heavy_check_mark:                                         | The name of the organization the database belongs to       |
| `Page`                                                     | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the page offset of returned results |
| `PerPage`                                                  | **float64*                                                 | :heavy_minus_sign:                                         | If provided, specifies the number of returned results      |