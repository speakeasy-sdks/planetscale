# ListOauthTokensRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ApplicationID`                                               | *string*                                                      | :heavy_check_mark:                                            | The ID of the OAuth application                               |
| `Organization`                                                | *string*                                                      | :heavy_check_mark:                                            | The name of the organization the OAuth application belongs to |
| `Page`                                                        | **float64*                                                    | :heavy_minus_sign:                                            | If provided, specifies the page offset of returned results    |
| `PerPage`                                                     | **float64*                                                    | :heavy_minus_sign:                                            | If provided, specifies the number of returned results         |