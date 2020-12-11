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

// AddonView struct for AddonView
type AddonView struct {
	Id           string                `json:"id,omitempty"`
	Name         string                `json:"name,omitempty"`
	RealId       string                `json:"realId,omitempty"`
	Region       string                `json:"region,omitempty"`
	Provider     AddonProviderInfoView `json:"provider,omitempty"`
	Plan         AddonPlanView         `json:"plan,omitempty"`
	CreationDate int64                 `json:"creationDate,omitempty"`
	ConfigKeys   []string              `json:"configKeys,omitempty"`
}