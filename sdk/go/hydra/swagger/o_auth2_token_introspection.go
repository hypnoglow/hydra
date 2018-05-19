/*
 * ORY Hydra - Cloud Native OAuth 2.0 and OpenID Connect Server
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here. Keep in mind that this document reflects the latest branch, always. Support for versioned documentation is coming in the future.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// https://tools.ietf.org/html/rfc7662
type OAuth2TokenIntrospection struct {

	// Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).
	Active bool `json:"active,omitempty"`

	Aud []string `json:"aud,omitempty"`

	// ClientID is aclient identifier for the OAuth 2.0 client that requested this token.
	ClientId string `json:"client_id,omitempty"`

	// Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty"`

	// Extra is arbitrary data set by the session.
	Ext map[string]interface{} `json:"ext,omitempty"`

	// Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.
	Iat int64 `json:"iat,omitempty"`

	// IssuerURL is a string representing the issuer of this token
	Iss string `json:"iss,omitempty"`

	// NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.
	Nbf int64 `json:"nbf,omitempty"`

	// Scope is a JSON string containing a space-separated list of scopes associated with this token.
	Scope string `json:"scope,omitempty"`

	// Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.
	Sub string `json:"sub,omitempty"`

	// TokenType is the introspected token's type, for example `access_token` or `refresh_token`.
	TokenType string `json:"token_type,omitempty"`

	// Username is a human-readable identifier for the resource owner who authorized this token.
	Username string `json:"username,omitempty"`
}
