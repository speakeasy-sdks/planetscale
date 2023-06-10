// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPromotionRequestRequest struct {
	// The name of the database the branch belongs to
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The name of the branch
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// The name of the organization the branch belongs to
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
}

type GetAPromotionRequest200ApplicationJSON struct {
}

type GetAPromotionRequestResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The promotion request for a branch
	GetAPromotionRequest200ApplicationJSONObject map[string]GetAPromotionRequest200ApplicationJSON
}
