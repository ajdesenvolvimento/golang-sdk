/*
 * MELI Markeplace SDK
 *
 * This is a the codebase to generate a SDK for Open Platform Marketplace
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package meli
// InlineObject struct for InlineObject
type InlineObject struct {
	GrantType string `json:"grant_type,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RedirectUri string `json:"redirect_uri,omitempty"`
	Code string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
