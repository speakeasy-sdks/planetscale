// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetCurrentUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Returns the current user that is associated with this service token
	GetCurrentUser200ApplicationJSONObject map[string]map[string]interface{}
}
