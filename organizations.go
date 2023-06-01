// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package planetscale

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/planetscale/pkg/models/operations"
	"github.com/speakeasy-sdks/planetscale/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// organizations -
// <p>API endpoints for managing organizations.</p>
type organizations struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newOrganizations(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *organizations {
	return &organizations{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Get - Get an organization
//
// ### Authorization
// A OAuth token must have the following scopes in order to use this API endpoint:
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | User | `read_organizations` |
// | Organization | `read_organization` |
func (s *organizations) Get(ctx context.Context, name string) (*operations.GetAnOrganizationResponse, error) {
	request := operations.GetAnOrganizationRequest{
		Name: name,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAnOrganizationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.GetAnOrganization200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// List - List organizations
//
// ### Authorization
// A   OAuth token must have the following   scopes in order to use this API endpoint:
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | User | `read_organizations` |
func (s *organizations) List(ctx context.Context, page *float64, perPage *float64) (*operations.ListOrganizationsResponse, error) {
	request := operations.ListOrganizationsRequest{
		Page:    page,
		PerPage: perPage,
	}

	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/organizations"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListOrganizationsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ListOrganizations200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// ListRegions - List regions for an organization
//
// ### Authorization
// A   OAuth token must have the following   scopes in order to use this API endpoint:
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | User | `read_organizations` |
// | Organization | `read_organization` |
func (s *organizations) ListRegions(ctx context.Context, name string, page *float64, perPage *float64) (*operations.ListRegionsForAnOrganizationResponse, error) {
	request := operations.ListRegionsForAnOrganizationRequest{
		Name:    name,
		Page:    page,
		PerPage: perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{name}/region", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListRegionsForAnOrganizationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ListRegionsForAnOrganization200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// Update - Update an organization
//
// ### Authorization
// A   OAuth token must have the following   scopes in order to use this API endpoint:
//
// **OAuth Scopes**
//
//	| Resource | Scopes |
//
// | :------- | :---------- |
// | Organization | `write_organization` |
func (s *organizations) Update(ctx context.Context, name string, requestBody *operations.UpdateAnOrganizationRequestBody) (*operations.UpdateAnOrganizationResponse, error) {
	request := operations.UpdateAnOrganizationRequest{
		Name:        name,
		RequestBody: requestBody,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{name}", request, nil)
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
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateAnOrganizationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.UpdateAnOrganization200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}
