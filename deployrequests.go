// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package planetscale

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/planetscale/pkg/models/operations"
	"github.com/speakeasy-sdks/planetscale/pkg/utils"
	"net/http"
)

// deployRequests -
// <p>API endpoints for managing database deploy requests.</p>
type deployRequests struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newDeployRequests(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *deployRequests {
	return &deployRequests{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Cancel - Cancel a queued deploy request
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) Cancel(ctx context.Context, database string, number string, organization string) (*operations.CancelAQueuedDeployRequestResponse, error) {
	request := operations.CancelAQueuedDeployRequestRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/cancel", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.CancelAQueuedDeployRequestResponse{
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

			res.CancelAQueuedDeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Close - Close a deploy request
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `write_deploy_requests` |
// | Database | `write_deploy_requests` |
func (s *deployRequests) Close(ctx context.Context, database string, number string, organization string, requestBody *operations.CloseADeployRequestRequestBody) (*operations.CloseADeployRequestResponse, error) {
	request := operations.CloseADeployRequestRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
		RequestBody:  requestBody,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}", request, nil)
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

	res := &operations.CloseADeployRequestResponse{
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

			res.CloseADeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CompleteErroredDeploy - Complete an errored deploy
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) CompleteErroredDeploy(ctx context.Context, database string, number string, organization string) (*operations.CompleteAnErroredDeployResponse, error) {
	request := operations.CompleteAnErroredDeployRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/complete-deploy", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.CompleteAnErroredDeployResponse{
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

			res.CompleteAnErroredDeploy200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CompleteGatedDeploy - Complete a gated deploy request
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) CompleteGatedDeploy(ctx context.Context, database string, number string, organization string) (*operations.CompleteAGatedDeployRequestResponse, error) {
	request := operations.CompleteAGatedDeployRequestRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/apply-deploy", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.CompleteAGatedDeployRequestResponse{
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

			res.CompleteAGatedDeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CompleteRevert - Complete a revert
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) CompleteRevert(ctx context.Context, database string, number string, organization string) (*operations.CompleteARevertResponse, error) {
	request := operations.CompleteARevertRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/revert", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.CompleteARevertResponse{
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

			res.CompleteARevert200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Create - Create a deploy request
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_requests`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `write_deploy_requests` |
// | Database | `write_deploy_requests` |
func (s *deployRequests) Create(ctx context.Context, database string, organization string, requestBody *operations.CreateADeployRequestRequestBody) (*operations.CreateADeployRequestResponse, error) {
	request := operations.CreateADeployRequestRequest{
		Database:     database,
		Organization: organization,
		RequestBody:  requestBody,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests", request, nil)
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

	res := &operations.CreateADeployRequestResponse{
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

			res.CreateADeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Get - Get a deploy request
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_deploy_requests` |
// | Database | `read_deploy_requests` |
func (s *deployRequests) Get(ctx context.Context, database string, number string, organization string) (*operations.GetADeployRequestResponse, error) {
	request := operations.GetADeployRequestRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}", request, nil)
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

	res := &operations.GetADeployRequestResponse{
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

			res.GetADeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetDeployment - Get a deployment
//
// <p>Get the deployment for a deploy request</p>
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_deploy_requests` |
// | Database | `read_deploy_requests` |
func (s *deployRequests) GetDeployment(ctx context.Context, database string, number string, organization string) (*operations.GetADeploymentResponse, error) {
	request := operations.GetADeploymentRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment", request, nil)
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

	res := &operations.GetADeploymentResponse{
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

			res.GetADeployment200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetQueue - Get a deploy queue
//
// <p>Get the deploy queue for a database</p>
func (s *deployRequests) GetQueue(ctx context.Context, database string, organization string) (*operations.GetADeployQueueResponse, error) {
	request := operations.GetADeployQueueRequest{
		Database:     database,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-queue", request, nil)
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

	res := &operations.GetADeployQueueResponse{
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

			res.GetADeployQueue200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// List - List deploy requests
//
// <p>List deploy requests for a database</p>
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_deploy_requests` |
// | Database | `read_deploy_requests` |
func (s *deployRequests) List(ctx context.Context, database string, organization string, page *float64, perPage *float64) (*operations.ListDeployRequestsResponse, error) {
	request := operations.ListDeployRequestsRequest{
		Database:     database,
		Organization: organization,
		Page:         page,
		PerPage:      perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests", request, nil)
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

	res := &operations.ListDeployRequestsResponse{
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

			res.ListDeployRequests200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListOperations - List deploy operations
//
// <p>List deploy operations for a deploy request</p>
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `read_deploy_requests` |
// | Database | `read_deploy_requests` |
func (s *deployRequests) ListOperations(ctx context.Context, request operations.ListDeployOperationsRequest) (*operations.ListDeployOperationsResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/operations", request, nil)
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

	res := &operations.ListDeployOperationsResponse{
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

			res.ListDeployOperations200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Queue - Queue a deploy request
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) Queue(ctx context.Context, database string, number string, organization string) (*operations.QueueADeployRequestResponse, error) {
	request := operations.QueueADeployRequestRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/deploy", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.QueueADeployRequestResponse{
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

			res.QueueADeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// SkipRevertPeriod - Skip revert period
//
// <p>Skips the revert period for a deploy request</p>
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) SkipRevertPeriod(ctx context.Context, database string, number string, organization string) (*operations.SkipRevertPeriodResponse, error) {
	request := operations.SkipRevertPeriodRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/skip-revert", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.SkipRevertPeriodResponse{
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

			res.SkipRevertPeriod200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// Update - Update auto-apply for deploy request
//
// <p>Enables or disabled the auto-apply setting for a deploy request</p>
//
// ### Authorization
// A service token or OAuth token must have the following access or scopes in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_deploy_request`, `create_deploy_request`
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `deploy_deploy_requests` |
// | Database | `deploy_deploy_requests` |
func (s *deployRequests) Update(ctx context.Context, database string, number string, organization string) (*operations.UpdateAutoApplyForDeployRequestResponse, error) {
	request := operations.UpdateAutoApplyForDeployRequestRequest{
		Database:     database,
		Number:       number,
		Organization: organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/databases/{database}/deploy-requests/{number}/auto-apply", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.UpdateAutoApplyForDeployRequestResponse{
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

			res.UpdateAutoApplyForDeployRequest200ApplicationJSONObject = out
		}
	}

	return res, nil
}
