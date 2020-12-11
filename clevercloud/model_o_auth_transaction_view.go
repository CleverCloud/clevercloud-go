/*
 * Clever-Cloud API
 *
 * Public API for managing Clever-Cloud data and products
 *
 * API version: 1.0.1
 * Contact: support@clever-cloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package clevercloud

// OAuthTransactionView struct for OAuthTransactionView
type OAuthTransactionView struct {
	TransactionId string `json:"transactionId,omitempty"`
	RedirectUri   string `json:"redirectUri,omitempty"`
}