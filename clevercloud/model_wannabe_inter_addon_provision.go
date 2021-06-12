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

// WannabeInterAddonProvision struct for WannabeInterAddonProvision
type WannabeInterAddonProvision struct {
	OrganisationId string            `json:"organisationId,omitempty"`
	UserId         string            `json:"userId,omitempty"`
	ProviderId     string            `json:"providerId,omitempty"`
	AddonId        string            `json:"addonId,omitempty"`
	Plan           string            `json:"plan,omitempty"`
	Name           string            `json:"name,omitempty"`
	Region         string            `json:"region,omitempty"`
	Options        map[string]string `json:"options,omitempty"`
}
