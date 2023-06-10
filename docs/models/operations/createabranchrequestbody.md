# CreateABranchRequestBody


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `BackupID`                                                           | **string*                                                            | :heavy_minus_sign:                                                   | If provided, restores the backup's schema and data to the new branch |
| `Name`                                                               | *string*                                                             | :heavy_check_mark:                                                   | The name of the branch                                               |
| `ParentBranch`                                                       | *string*                                                             | :heavy_check_mark:                                                   | Parent branch                                                        |