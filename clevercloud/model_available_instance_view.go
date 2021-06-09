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

// AvailableInstanceView struct for AvailableInstanceView
type AvailableInstanceView struct {
	Type          *string              `json:"type,omitempty"`
	Version       *string              `json:"version,omitempty"`
	Name          *string              `json:"name,omitempty"`
	Variant       *InstanceVariantView `json:"variant,omitempty"`
	Description   *string              `json:"description,omitempty"`
	Enabled       *bool                `json:"enabled,omitempty"`
	ComingSoon    *bool                `json:"comingSoon,omitempty"`
	MaxInstances  *int32               `json:"maxInstances,omitempty"`
	Tags          *[]string            `json:"tags,omitempty"`
	Deployments   *[]string            `json:"deployments,omitempty"`
	Flavors       *[]FlavorView        `json:"flavors,omitempty"`
	DefaultFlavor *FlavorView          `json:"defaultFlavor,omitempty"`
	BuildFlavor   *FlavorView          `json:"buildFlavor,omitempty"`
}

// NewAvailableInstanceView instantiates a new AvailableInstanceView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableInstanceView() *AvailableInstanceView {
	this := AvailableInstanceView{}
	return &this
}

// NewAvailableInstanceViewWithDefaults instantiates a new AvailableInstanceView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableInstanceViewWithDefaults() *AvailableInstanceView {
	this := AvailableInstanceView{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AvailableInstanceView) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AvailableInstanceView) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AvailableInstanceView) SetName(v string) {
	o.Name = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetVariant() InstanceVariantView {
	if o == nil || o.Variant == nil {
		var ret InstanceVariantView
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetVariantOk() (*InstanceVariantView, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasVariant() bool {
	if o != nil && o.Variant != nil {
		return true
	}

	return false
}

// SetVariant gets a reference to the given InstanceVariantView and assigns it to the Variant field.
func (o *AvailableInstanceView) SetVariant(v InstanceVariantView) {
	o.Variant = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AvailableInstanceView) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AvailableInstanceView) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetComingSoon returns the ComingSoon field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetComingSoon() bool {
	if o == nil || o.ComingSoon == nil {
		var ret bool
		return ret
	}
	return *o.ComingSoon
}

// GetComingSoonOk returns a tuple with the ComingSoon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetComingSoonOk() (*bool, bool) {
	if o == nil || o.ComingSoon == nil {
		return nil, false
	}
	return o.ComingSoon, true
}

// HasComingSoon returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasComingSoon() bool {
	if o != nil && o.ComingSoon != nil {
		return true
	}

	return false
}

// SetComingSoon gets a reference to the given bool and assigns it to the ComingSoon field.
func (o *AvailableInstanceView) SetComingSoon(v bool) {
	o.ComingSoon = &v
}

// GetMaxInstances returns the MaxInstances field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetMaxInstances() int32 {
	if o == nil || o.MaxInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxInstances
}

// GetMaxInstancesOk returns a tuple with the MaxInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetMaxInstancesOk() (*int32, bool) {
	if o == nil || o.MaxInstances == nil {
		return nil, false
	}
	return o.MaxInstances, true
}

// HasMaxInstances returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasMaxInstances() bool {
	if o != nil && o.MaxInstances != nil {
		return true
	}

	return false
}

// SetMaxInstances gets a reference to the given int32 and assigns it to the MaxInstances field.
func (o *AvailableInstanceView) SetMaxInstances(v int32) {
	o.MaxInstances = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AvailableInstanceView) SetTags(v []string) {
	o.Tags = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetDeployments() []string {
	if o == nil || o.Deployments == nil {
		var ret []string
		return ret
	}
	return *o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetDeploymentsOk() (*[]string, bool) {
	if o == nil || o.Deployments == nil {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasDeployments() bool {
	if o != nil && o.Deployments != nil {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []string and assigns it to the Deployments field.
func (o *AvailableInstanceView) SetDeployments(v []string) {
	o.Deployments = &v
}

// GetFlavors returns the Flavors field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetFlavors() []FlavorView {
	if o == nil || o.Flavors == nil {
		var ret []FlavorView
		return ret
	}
	return *o.Flavors
}

// GetFlavorsOk returns a tuple with the Flavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetFlavorsOk() (*[]FlavorView, bool) {
	if o == nil || o.Flavors == nil {
		return nil, false
	}
	return o.Flavors, true
}

// HasFlavors returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasFlavors() bool {
	if o != nil && o.Flavors != nil {
		return true
	}

	return false
}

// SetFlavors gets a reference to the given []FlavorView and assigns it to the Flavors field.
func (o *AvailableInstanceView) SetFlavors(v []FlavorView) {
	o.Flavors = &v
}

// GetDefaultFlavor returns the DefaultFlavor field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetDefaultFlavor() FlavorView {
	if o == nil || o.DefaultFlavor == nil {
		var ret FlavorView
		return ret
	}
	return *o.DefaultFlavor
}

// GetDefaultFlavorOk returns a tuple with the DefaultFlavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetDefaultFlavorOk() (*FlavorView, bool) {
	if o == nil || o.DefaultFlavor == nil {
		return nil, false
	}
	return o.DefaultFlavor, true
}

// HasDefaultFlavor returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasDefaultFlavor() bool {
	if o != nil && o.DefaultFlavor != nil {
		return true
	}

	return false
}

// SetDefaultFlavor gets a reference to the given FlavorView and assigns it to the DefaultFlavor field.
func (o *AvailableInstanceView) SetDefaultFlavor(v FlavorView) {
	o.DefaultFlavor = &v
}

// GetBuildFlavor returns the BuildFlavor field value if set, zero value otherwise.
func (o *AvailableInstanceView) GetBuildFlavor() FlavorView {
	if o == nil || o.BuildFlavor == nil {
		var ret FlavorView
		return ret
	}
	return *o.BuildFlavor
}

// GetBuildFlavorOk returns a tuple with the BuildFlavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableInstanceView) GetBuildFlavorOk() (*FlavorView, bool) {
	if o == nil || o.BuildFlavor == nil {
		return nil, false
	}
	return o.BuildFlavor, true
}

// HasBuildFlavor returns a boolean if a field has been set.
func (o *AvailableInstanceView) HasBuildFlavor() bool {
	if o != nil && o.BuildFlavor != nil {
		return true
	}

	return false
}

// SetBuildFlavor gets a reference to the given FlavorView and assigns it to the BuildFlavor field.
func (o *AvailableInstanceView) SetBuildFlavor(v FlavorView) {
	o.BuildFlavor = &v
}

func (o AvailableInstanceView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Variant != nil {
		toSerialize["variant"] = o.Variant
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ComingSoon != nil {
		toSerialize["comingSoon"] = o.ComingSoon
	}
	if o.MaxInstances != nil {
		toSerialize["maxInstances"] = o.MaxInstances
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Deployments != nil {
		toSerialize["deployments"] = o.Deployments
	}
	if o.Flavors != nil {
		toSerialize["flavors"] = o.Flavors
	}
	if o.DefaultFlavor != nil {
		toSerialize["defaultFlavor"] = o.DefaultFlavor
	}
	if o.BuildFlavor != nil {
		toSerialize["buildFlavor"] = o.BuildFlavor
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableInstanceView struct {
	value *AvailableInstanceView
	isSet bool
}

func (v NullableAvailableInstanceView) Get() *AvailableInstanceView {
	return v.value
}

func (v *NullableAvailableInstanceView) Set(val *AvailableInstanceView) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableInstanceView) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableInstanceView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableInstanceView(val *AvailableInstanceView) *NullableAvailableInstanceView {
	return &NullableAvailableInstanceView{value: val, isSet: true}
}

func (v NullableAvailableInstanceView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableInstanceView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
