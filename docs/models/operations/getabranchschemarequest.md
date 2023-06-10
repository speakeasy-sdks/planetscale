# GetABranchSchemaRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Database`                                            | *string*                                              | :heavy_check_mark:                                    | The name of the database the branch belongs to        |
| `Name`                                                | *string*                                              | :heavy_check_mark:                                    | The name of the branch                                |
| `Organization`                                        | *string*                                              | :heavy_check_mark:                                    | The name of the organization the branch belongs to    |
| `Keyspace`                                            | **string*                                             | :heavy_minus_sign:                                    | If provided, the schema for this keyspace is returned |