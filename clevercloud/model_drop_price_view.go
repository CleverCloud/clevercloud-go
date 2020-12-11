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

// DropPriceView struct for DropPriceView
type DropPriceView struct {
	Currency string  `json:"currency,omitempty"`
	Value    float32 `json:"value,omitempty"`
}