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

// AddonApplicationInfo struct for AddonApplicationInfo
type AddonApplicationInfo struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Config      map[string]string `json:"config,omitempty"`
	CallbackUrl string            `json:"callback_url,omitempty"`
	OwnerEmail  string            `json:"owner_email,omitempty"`
	OwnerId     string            `json:"owner_id,omitempty"`
	OwnerName   string            `json:"owner_name,omitempty"`
	OwnerEmails []string          `json:"owner_emails,omitempty"`
	Region      string            `json:"region,omitempty"`
	Domains     []string          `json:"domains,omitempty"`
}
