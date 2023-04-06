// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListDatabasePromotionRequestsRequest struct {
	// The name of the database
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// The name of the organization the database belongs to
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
	// If provided, specifies the page offset of returned results
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
	// If provided, specifies the number of returned results
	PerPage *float64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListDatabasePromotionRequestsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Promotion requests for the database
	ListDatabasePromotionRequests200ApplicationJSONObject map[string]map[string]interface{}
}
