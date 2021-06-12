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

// WannabeUser struct for WannabeUser
type WannabeUser struct {
	Email              string `json:"email,omitempty"`
	Name               string `json:"name,omitempty"`
	Password           string `json:"password,omitempty"`
	Phone              string `json:"phone,omitempty"`
	Address            string `json:"address,omitempty"`
	City               string `json:"city,omitempty"`
	Zipcode            string `json:"zipcode,omitempty"`
	Country            string `json:"country,omitempty"`
	Lang               string `json:"lang,omitempty"`
	Terms              bool   `json:"terms,omitempty"`
	SubscriptionSource string `json:"subscriptionSource,omitempty"`
}
