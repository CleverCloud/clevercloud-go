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

// OrganisationMemberUserView struct for OrganisationMemberUserView
type OrganisationMemberUserView struct {
	Id     *string `json:"id,omitempty"`
	Email  *string `json:"email,omitempty"`
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}

// NewOrganisationMemberUserView instantiates a new OrganisationMemberUserView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationMemberUserView() *OrganisationMemberUserView {
	this := OrganisationMemberUserView{}
	return &this
}

// NewOrganisationMemberUserViewWithDefaults instantiates a new OrganisationMemberUserView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationMemberUserViewWithDefaults() *OrganisationMemberUserView {
	this := OrganisationMemberUserView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganisationMemberUserView) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationMemberUserView) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganisationMemberUserView) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganisationMemberUserView) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganisationMemberUserView) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationMemberUserView) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganisationMemberUserView) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganisationMemberUserView) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganisationMemberUserView) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationMemberUserView) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganisationMemberUserView) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganisationMemberUserView) SetName(v string) {
	o.Name = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *OrganisationMemberUserView) GetAvatar() string {
	if o == nil || o.Avatar == nil {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationMemberUserView) GetAvatarOk() (*string, bool) {
	if o == nil || o.Avatar == nil {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *OrganisationMemberUserView) HasAvatar() bool {
	if o != nil && o.Avatar != nil {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *OrganisationMemberUserView) SetAvatar(v string) {
	o.Avatar = &v
}

func (o OrganisationMemberUserView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	return json.Marshal(toSerialize)
}

type NullableOrganisationMemberUserView struct {
	value *OrganisationMemberUserView
	isSet bool
}

func (v NullableOrganisationMemberUserView) Get() *OrganisationMemberUserView {
	return v.value
}

func (v *NullableOrganisationMemberUserView) Set(val *OrganisationMemberUserView) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationMemberUserView) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationMemberUserView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationMemberUserView(val *OrganisationMemberUserView) *NullableOrganisationMemberUserView {
	return &NullableOrganisationMemberUserView{value: val, isSet: true}
}

func (v NullableOrganisationMemberUserView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationMemberUserView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
