// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package planetscale

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/planetscale/pkg/models/operations"
	"github.com/speakeasy-sdks/planetscale/pkg/utils"
	"net/http"
)

// oAuthTokens -
// <p>API endpoints for managing OAuth tokens.</p>
type oAuthTokens struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newOAuthTokens(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *oAuthTokens {
	return &oAuthTokens{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Create - Create or renew an OAuth token
//
// <p>Create an OAuth token from an authorization grant code, or refresh an OAuth token from a refresh token</p>
//
// ### Authorization
// A service token   must have the following access   in order to use this API endpoint:
//
// **Service Token Accesses**
//
//	`write_oauth_tokens`
func (s *oAuthTokens) Create(ctx context.Context, id string, organization string, requestBody *operations.CreateOrRenewAnOauthTokenRequestBody) (*operations.CreateOrRenewAnOauthTokenResponse, error) {
	request := operations.CreateOrRenewAnOauthTokenRequest{
		ID:           id,
		Organization: organization,
		RequestBody:  requestBody,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization}/oauth-applications/{id}/token", request, nil)
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

	res := &operations.CreateOrRenewAnOauthTokenResponse{
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

			res.CreateOrRenewAnOauthToken200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 500:
	}

	return res, nil
}
