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

// Summary struct for Summary
type Summary struct {
	User          UserSummary           `json:"user,omitempty"`
	Organisations []OrganisationSummary `json:"organisations,omitempty"`
}
