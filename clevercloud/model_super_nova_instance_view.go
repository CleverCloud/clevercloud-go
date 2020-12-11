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

// SuperNovaInstanceView struct for SuperNovaInstanceView
type SuperNovaInstanceView struct {
	Id             string          `json:"id,omitempty"`
	AppId          string          `json:"appId,omitempty"`
	Ip             string          `json:"ip,omitempty"`
	AppPort        int32           `json:"appPort,omitempty"`
	State          string          `json:"state,omitempty"`
	Flavor         SuperNovaFlavor `json:"flavor,omitempty"`
	Commit         string          `json:"commit,omitempty"`
	DeployNumber   int32           `json:"deployNumber,omitempty"`
	DeployId       string          `json:"deployId,omitempty"`
	InstanceNumber int32           `json:"instanceNumber,omitempty"`
	DisplayName    string          `json:"displayName,omitempty"`
	CreationDate   int64           `json:"creationDate,omitempty"`
}