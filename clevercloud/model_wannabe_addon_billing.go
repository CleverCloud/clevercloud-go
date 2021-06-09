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

// WannabeAddonBilling struct for WannabeAddonBilling
type WannabeAddonBilling struct {
	Cost *float32 `json:"cost,omitempty"`
}

// NewWannabeAddonBilling instantiates a new WannabeAddonBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWannabeAddonBilling() *WannabeAddonBilling {
	this := WannabeAddonBilling{}
	return &this
}

// NewWannabeAddonBillingWithDefaults instantiates a new WannabeAddonBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWannabeAddonBillingWithDefaults() *WannabeAddonBilling {
	this := WannabeAddonBilling{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *WannabeAddonBilling) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WannabeAddonBilling) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *WannabeAddonBilling) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *WannabeAddonBilling) SetCost(v float32) {
	o.Cost = &v
}

func (o WannabeAddonBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	return json.Marshal(toSerialize)
}

type NullableWannabeAddonBilling struct {
	value *WannabeAddonBilling
	isSet bool
}

func (v NullableWannabeAddonBilling) Get() *WannabeAddonBilling {
	return v.value
}

func (v *NullableWannabeAddonBilling) Set(val *WannabeAddonBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableWannabeAddonBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableWannabeAddonBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWannabeAddonBilling(val *WannabeAddonBilling) *NullableWannabeAddonBilling {
	return &NullableWannabeAddonBilling{value: val, isSet: true}
}

func (v NullableWannabeAddonBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWannabeAddonBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
