// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateABranchPasswordRequestBodyRoleEnum - The database role of the password (i.e. admin)
type CreateABranchPasswordRequestBodyRoleEnum string

const (
	CreateABranchPasswordRequestBodyRoleEnumReader     CreateABranchPasswordRequestBodyRoleEnum = "reader"
	CreateABranchPasswordRequestBodyRoleEnumWriter     CreateABranchPasswordRequestBodyRoleEnum = "writer"
	CreateABranchPasswordRequestBodyRoleEnumAdmin      CreateABranchPasswordRequestBodyRoleEnum = "admin"
	CreateABranchPasswordRequestBodyRoleEnumReadwriter CreateABranchPasswordRequestBodyRoleEnum = "readwriter"
)

func (e *CreateABranchPasswordRequestBodyRoleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "reader":
		fallthrough
	case "writer":
		fallthrough
	case "admin":
		fallthrough
	case "readwriter":
		*e = CreateABranchPasswordRequestBodyRoleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateABranchPasswordRequestBodyRoleEnum: %s", s)
	}
}

type CreateABranchPasswordRequestBody struct {
	// A read-only region of the database branch
	ReadOnlyRegionID *string `json:"read_only_region_id,omitempty"`
	// The database role of the password (i.e. admin)
	Role *CreateABranchPasswordRequestBodyRoleEnum `json:"role,omitempty"`
	// Time to live (in seconds) for the password. The password will be invalid when TTL has passed
	TTL *float64 `json:"ttl,omitempty"`
}

type CreateABranchPasswordRequest struct {
	RequestBody *CreateABranchPasswordRequestBody `request:"mediaType=application/json"`
	// The name of the branch the password belongs to
	Branch string `pathParam:"style=simple,explode=false,name=branch"`
	// The name of the database the password belongs to
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The name of the organization the password belongs to
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
}

type CreateABranchPasswordResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Creates a database branch password
	CreateABranchPassword201ApplicationJSONObject map[string]map[string]interface{}
}
