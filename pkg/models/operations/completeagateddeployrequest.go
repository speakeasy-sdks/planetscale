// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CompleteAGatedDeployRequestRequest struct {
	// The name of the deploy request's database
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The number of the deploy request
	Number string `pathParam:"style=simple,explode=false,name=number"`
	// The name of the deploy request's organization
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
}

type CompleteAGatedDeployRequestResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Returns the deploy request whose deployment has been completed
	CompleteAGatedDeployRequest200ApplicationJSONObject map[string]map[string]interface{}
}
