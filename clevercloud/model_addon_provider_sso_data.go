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

// AddonProviderSSOData struct for AddonProviderSSOData
type AddonProviderSSOData struct {
	Url      *string `json:"url,omitempty"`
	Macaroon *string `json:"macaroon,omitempty"`
}

// NewAddonProviderSSOData instantiates a new AddonProviderSSOData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonProviderSSOData() *AddonProviderSSOData {
	this := AddonProviderSSOData{}
	return &this
}

// NewAddonProviderSSODataWithDefaults instantiates a new AddonProviderSSOData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonProviderSSODataWithDefaults() *AddonProviderSSOData {
	this := AddonProviderSSOData{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AddonProviderSSOData) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonProviderSSOData) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AddonProviderSSOData) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AddonProviderSSOData) SetUrl(v string) {
	o.Url = &v
}

// GetMacaroon returns the Macaroon field value if set, zero value otherwise.
func (o *AddonProviderSSOData) GetMacaroon() string {
	if o == nil || o.Macaroon == nil {
		var ret string
		return ret
	}
	return *o.Macaroon
}

// GetMacaroonOk returns a tuple with the Macaroon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonProviderSSOData) GetMacaroonOk() (*string, bool) {
	if o == nil || o.Macaroon == nil {
		return nil, false
	}
	return o.Macaroon, true
}

// HasMacaroon returns a boolean if a field has been set.
func (o *AddonProviderSSOData) HasMacaroon() bool {
	if o != nil && o.Macaroon != nil {
		return true
	}

	return false
}

// SetMacaroon gets a reference to the given string and assigns it to the Macaroon field.
func (o *AddonProviderSSOData) SetMacaroon(v string) {
	o.Macaroon = &v
}

func (o AddonProviderSSOData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Macaroon != nil {
		toSerialize["macaroon"] = o.Macaroon
	}
	return json.Marshal(toSerialize)
}

type NullableAddonProviderSSOData struct {
	value *AddonProviderSSOData
	isSet bool
}

func (v NullableAddonProviderSSOData) Get() *AddonProviderSSOData {
	return v.value
}

func (v *NullableAddonProviderSSOData) Set(val *AddonProviderSSOData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonProviderSSOData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonProviderSSOData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonProviderSSOData(val *AddonProviderSSOData) *NullableAddonProviderSSOData {
	return &NullableAddonProviderSSOData{value: val, isSet: true}
}

func (v NullableAddonProviderSSOData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonProviderSSOData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
