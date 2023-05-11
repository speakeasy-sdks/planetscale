// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package planetscale

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/planetscale/pkg/models/operations"
	"github.com/speakeasy-sdks/planetscale/pkg/utils"
	"net/http"
)

// databases -
// <p>API endpoints for managing databases within an organization.</p>
type databases struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newDatabases(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *databases {
	return &databases{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Create - Create a database
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`create_databases`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `create_databases` |
func (s *databases) Create(ctx context.Context, organization string, requestBody *operations.CreateADatabaseRequestBody) (*operations.CreateADatabaseResponse, error) {
	request := operations.CreateADatabaseRequest{
		Organization: organization,
		RequestBody:  requestBody,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateADatabaseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateADatabase201ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// Delete - Delete a database
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`delete_database`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `delete_databases` |
// | Database | `delete_database` |
func (s *databases) Delete(ctx context.Context, name string, organization string) (*operations.DeleteADatabaseResponse, error) {
	request := operations.DeleteADatabaseRequest{
		Name:         name,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteADatabaseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// Get - Get a database
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_databases` |
// | Database | `read_database` |
func (s *databases) Get(ctx context.Context, name string, organization string) (*operations.GetADatabaseResponse, error) {
	request := operations.GetADatabaseRequest{
		Name:         name,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetADatabaseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetADatabase200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// List - List databases
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_databases` |
func (s *databases) List(ctx context.Context, organization string, page *float64, perPage *float64) (*operations.ListDatabasesResponse, error) {
	request := operations.ListDatabasesRequest{
		Organization: organization,
		Page:         page,
		PerPage:      perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListDatabasesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListDatabases200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// ListPromotionRequests - List database promotion requests
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_databases` |
// | Database | `read_database` |
func (s *databases) ListPromotionRequests(ctx context.Context, name string, organization string, page *float64, perPage *float64) (*operations.ListDatabasePromotionRequestsResponse, error) {
	request := operations.ListDatabasePromotionRequestsRequest{
		Name:         name,
		Organization: organization,
		Page:         page,
		PerPage:      perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{name}/promotion-requests", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListDatabasePromotionRequestsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListDatabasePromotionRequests200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// ListReadOnlyRegions - List read-only regions
//
// <p>List read-only regions for the database’s default branch</p>
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_branches` |
// | Database | `read_branches` |
func (s *databases) ListReadOnlyRegions(ctx context.Context, name string, organization string, page *float64, perPage *float64) (*operations.ListReadOnlyRegionsResponse, error) {
	request := operations.ListReadOnlyRegionsRequest{
		Name:         name,
		Organization: organization,
		Page:         page,
		PerPage:      perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{name}/read-only-regions", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListReadOnlyRegionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListReadOnlyRegions200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// ListRegions - List database regions
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_database`, `delete_database`, `write_database`, `read_branch`, `delete_branch`, `create_branch`, `delete_production_branch`, `connect_branch`, `connect_production_branch`, `delete_branch_password`, `delete_production_branch_password`, `read_deploy_request`, `create_deploy_request`, `approve_deploy_request`, `read_comment`, `create_comment`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_databases` |
// | Database | `read_database` |
func (s *databases) ListRegions(ctx context.Context, name string, organization string, page *float64, perPage *float64) (*operations.ListDatabaseRegionsResponse, error) {
	request := operations.ListDatabaseRegionsRequest{
		Name:         name,
		Organization: organization,
		Page:         page,
		PerPage:      perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{name}/regions", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListDatabaseRegionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListDatabaseRegions200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// Update - Update database settings
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`write_database`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `write_databases` |
// | Database | `write_database` |
func (s *databases) Update(ctx context.Context, name string, organization string, requestBody *operations.UpdateDatabaseSettingsRequestBody) (*operations.UpdateDatabaseSettingsResponse, error) {
	request := operations.UpdateDatabaseSettingsRequest{
		Name:         name,
		Organization: organization,
		RequestBody:  requestBody,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateDatabaseSettingsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateDatabaseSettings200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}
