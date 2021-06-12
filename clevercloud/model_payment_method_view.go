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

// PaymentMethodView struct for PaymentMethodView
type PaymentMethodView struct {
	OwnerId        string `json:"ownerId,omitempty"`
	Type           string `json:"type,omitempty"`
	Token          string `json:"token,omitempty"`
	ImageUrl       string `json:"imageUrl,omitempty"`
	IsDefault      bool   `json:"isDefault,omitempty"`
	HolderName     string `json:"holderName,omitempty"`
	Number         string `json:"number,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	IsExpired      bool   `json:"isExpired,omitempty"`
	CardType       string `json:"cardType,omitempty"`
	Email          string `json:"email,omitempty"`
	BankCode       string `json:"bankCode,omitempty"`
	BranchCode     string `json:"branchCode,omitempty"`
	Country        string `json:"country,omitempty"`
	Fingerprint    string `json:"fingerprint,omitempty"`
}
