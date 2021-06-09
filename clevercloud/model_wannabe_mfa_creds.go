/*
 * Clever-Cloud API
 *
 * Public API for managing Clever-Cloud data and products
 *
 * API version: 1.0.1
 * Contact: support@clever-cloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clevercloud

import (
	"encoding/json"
)

// WannabeMFACreds struct for WannabeMFACreds
type WannabeMFACreds struct {
	RevokeTokens *bool   `json:"revokeTokens,omitempty"`
	Code         *string `json:"code,omitempty"`
}

// NewWannabeMFACreds instantiates a new WannabeMFACreds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWannabeMFACreds() *WannabeMFACreds {
	this := WannabeMFACreds{}
	return &this
}

// NewWannabeMFACredsWithDefaults instantiates a new WannabeMFACreds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWannabeMFACredsWithDefaults() *WannabeMFACreds {
	this := WannabeMFACreds{}
	return &this
}

// GetRevokeTokens returns the RevokeTokens field value if set, zero value otherwise.
func (o *WannabeMFACreds) GetRevokeTokens() bool {
	if o == nil || o.RevokeTokens == nil {
		var ret bool
		return ret
	}
	return *o.RevokeTokens
}

// GetRevokeTokensOk returns a tuple with the RevokeTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WannabeMFACreds) GetRevokeTokensOk() (*bool, bool) {
	if o == nil || o.RevokeTokens == nil {
		return nil, false
	}
	return o.RevokeTokens, true
}

// HasRevokeTokens returns a boolean if a field has been set.
func (o *WannabeMFACreds) HasRevokeTokens() bool {
	if o != nil && o.RevokeTokens != nil {
		return true
	}

	return false
}

// SetRevokeTokens gets a reference to the given bool and assigns it to the RevokeTokens field.
func (o *WannabeMFACreds) SetRevokeTokens(v bool) {
	o.RevokeTokens = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *WannabeMFACreds) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WannabeMFACreds) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *WannabeMFACreds) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *WannabeMFACreds) SetCode(v string) {
	o.Code = &v
}

func (o WannabeMFACreds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RevokeTokens != nil {
		toSerialize["revokeTokens"] = o.RevokeTokens
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableWannabeMFACreds struct {
	value *WannabeMFACreds
	isSet bool
}

func (v NullableWannabeMFACreds) Get() *WannabeMFACreds {
	return v.value
}

func (v *NullableWannabeMFACreds) Set(val *WannabeMFACreds) {
	v.value = val
	v.isSet = true
}

func (v NullableWannabeMFACreds) IsSet() bool {
	return v.isSet
}

func (v *NullableWannabeMFACreds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWannabeMFACreds(val *WannabeMFACreds) *NullableWannabeMFACreds {
	return &NullableWannabeMFACreds{value: val, isSet: true}
}

func (v NullableWannabeMFACreds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWannabeMFACreds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
