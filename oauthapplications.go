// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package planetscale

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/planetscale/pkg/models/operations"
	"github.com/speakeasy-sdks/planetscale/pkg/utils"
	"net/http"
)

// oAuthApplications -
// <p>API endpoints for fetching OAuth applications.</p>
type oAuthApplications struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newOAuthApplications(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *oAuthApplications {
	return &oAuthApplications{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Delete - Delete an OAuth token
//
// ### Authorization
// A service token   must have the following access   in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`delete_oauth_tokens`
func (s *oAuthApplications) Delete(ctx context.Context, applicationID string, organization string, tokenID string) (*operations.DeleteAnOauthTokenResponse, error) {
	request := operations.DeleteAnOauthTokenRequest{
		ApplicationID: applicationID,
		Organization:  organization,
		TokenID:       tokenID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

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

	res := &operations.DeleteAnOauthTokenResponse{
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

// Get - Get an OAuth application
//
// ### Authorization
// A service token   must have the following access   in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_oauth_applications`
func (s *oAuthApplications) Get(ctx context.Context, applicationID string, organization string) (*operations.GetAnOauthApplicationResponse, error) {
	request := operations.GetAnOauthApplicationRequest{
		ApplicationID: applicationID,
		Organization:  organization,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/oauth-applications/{application_id}", request, nil)
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
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAnOauthApplicationResponse{
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

			res.GetAnOauthApplication200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// GetToken - Get an OAuth token
//
// ### Authorization
// A service token   must have the following access   in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_oauth_tokens`
func (s *oAuthApplications) GetToken(ctx context.Context, applicationID string, organization string, tokenID string) (*operations.GetAnOauthTokenResponse, error) {
	request := operations.GetAnOauthTokenRequest{
		ApplicationID: applicationID,
		Organization:  organization,
		TokenID:       tokenID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}", request, nil)
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
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAnOauthTokenResponse{
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

			res.GetAnOauthToken200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// List - List OAuth applications
//
// ### Authorization
// A service token   must have the following access   in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_oauth_applications`
func (s *oAuthApplications) List(ctx context.Context, organization string, page *float64, perPage *float64) (*operations.ListOauthApplicationsResponse, error) {
	request := operations.ListOauthApplicationsRequest{
		Organization: organization,
		Page:         page,
		PerPage:      perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/oauth-applications", request, nil)
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
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListOauthApplicationsResponse{
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

			res.ListOauthApplications200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}

// ListTokens - List OAuth tokens
//
// <p>List OAuth tokens created by an OAuth application</p>
//
// ### Authorization
// A service token   must have the following access   in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`read_oauth_tokens`
func (s *oAuthApplications) ListTokens(ctx context.Context, applicationID string, organization string, page *float64, perPage *float64) (*operations.ListOauthTokensResponse, error) {
	request := operations.ListOauthTokensRequest{
		ApplicationID: applicationID,
		Organization:  organization,
		Page:          page,
		PerPage:       perPage,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/oauth-applications/{application_id}/tokens", request, nil)
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
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListOauthTokensResponse{
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

			res.ListOauthTokens200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}
