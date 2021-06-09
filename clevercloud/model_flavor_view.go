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

// FlavorView struct for FlavorView
type FlavorView struct {
	Name            *string        `json:"name,omitempty"`
	Mem             *int32         `json:"mem,omitempty"`
	Cpus            *int32         `json:"cpus,omitempty"`
	Gpus            *int32         `json:"gpus,omitempty"`
	Disk            *int32         `json:"disk,omitempty"`
	Price           *float32       `json:"price,omitempty"`
	Available       *bool          `json:"available,omitempty"`
	Microservice    *bool          `json:"microservice,omitempty"`
	MachineLearning *bool          `json:"machine_learning,omitempty"`
	Nice            *int32         `json:"nice,omitempty"`
	PriceId         *string        `json:"price_id,omitempty"`
	Memory          *ValueWithUnit `json:"memory,omitempty"`
	Rbdimage        *string        `json:"rbdimage,omitempty"`
}

// NewFlavorView instantiates a new FlavorView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorView() *FlavorView {
	this := FlavorView{}
	return &this
}

// NewFlavorViewWithDefaults instantiates a new FlavorView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorViewWithDefaults() *FlavorView {
	this := FlavorView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FlavorView) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FlavorView) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FlavorView) SetName(v string) {
	o.Name = &v
}

// GetMem returns the Mem field value if set, zero value otherwise.
func (o *FlavorView) GetMem() int32 {
	if o == nil || o.Mem == nil {
		var ret int32
		return ret
	}
	return *o.Mem
}

// GetMemOk returns a tuple with the Mem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetMemOk() (*int32, bool) {
	if o == nil || o.Mem == nil {
		return nil, false
	}
	return o.Mem, true
}

// HasMem returns a boolean if a field has been set.
func (o *FlavorView) HasMem() bool {
	if o != nil && o.Mem != nil {
		return true
	}

	return false
}

// SetMem gets a reference to the given int32 and assigns it to the Mem field.
func (o *FlavorView) SetMem(v int32) {
	o.Mem = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *FlavorView) GetCpus() int32 {
	if o == nil || o.Cpus == nil {
		var ret int32
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetCpusOk() (*int32, bool) {
	if o == nil || o.Cpus == nil {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *FlavorView) HasCpus() bool {
	if o != nil && o.Cpus != nil {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int32 and assigns it to the Cpus field.
func (o *FlavorView) SetCpus(v int32) {
	o.Cpus = &v
}

// GetGpus returns the Gpus field value if set, zero value otherwise.
func (o *FlavorView) GetGpus() int32 {
	if o == nil || o.Gpus == nil {
		var ret int32
		return ret
	}
	return *o.Gpus
}

// GetGpusOk returns a tuple with the Gpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetGpusOk() (*int32, bool) {
	if o == nil || o.Gpus == nil {
		return nil, false
	}
	return o.Gpus, true
}

// HasGpus returns a boolean if a field has been set.
func (o *FlavorView) HasGpus() bool {
	if o != nil && o.Gpus != nil {
		return true
	}

	return false
}

// SetGpus gets a reference to the given int32 and assigns it to the Gpus field.
func (o *FlavorView) SetGpus(v int32) {
	o.Gpus = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *FlavorView) GetDisk() int32 {
	if o == nil || o.Disk == nil {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetDiskOk() (*int32, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *FlavorView) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *FlavorView) SetDisk(v int32) {
	o.Disk = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *FlavorView) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *FlavorView) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *FlavorView) SetPrice(v float32) {
	o.Price = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *FlavorView) GetAvailable() bool {
	if o == nil || o.Available == nil {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetAvailableOk() (*bool, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *FlavorView) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *FlavorView) SetAvailable(v bool) {
	o.Available = &v
}

// GetMicroservice returns the Microservice field value if set, zero value otherwise.
func (o *FlavorView) GetMicroservice() bool {
	if o == nil || o.Microservice == nil {
		var ret bool
		return ret
	}
	return *o.Microservice
}

// GetMicroserviceOk returns a tuple with the Microservice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetMicroserviceOk() (*bool, bool) {
	if o == nil || o.Microservice == nil {
		return nil, false
	}
	return o.Microservice, true
}

// HasMicroservice returns a boolean if a field has been set.
func (o *FlavorView) HasMicroservice() bool {
	if o != nil && o.Microservice != nil {
		return true
	}

	return false
}

// SetMicroservice gets a reference to the given bool and assigns it to the Microservice field.
func (o *FlavorView) SetMicroservice(v bool) {
	o.Microservice = &v
}

// GetMachineLearning returns the MachineLearning field value if set, zero value otherwise.
func (o *FlavorView) GetMachineLearning() bool {
	if o == nil || o.MachineLearning == nil {
		var ret bool
		return ret
	}
	return *o.MachineLearning
}

// GetMachineLearningOk returns a tuple with the MachineLearning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetMachineLearningOk() (*bool, bool) {
	if o == nil || o.MachineLearning == nil {
		return nil, false
	}
	return o.MachineLearning, true
}

// HasMachineLearning returns a boolean if a field has been set.
func (o *FlavorView) HasMachineLearning() bool {
	if o != nil && o.MachineLearning != nil {
		return true
	}

	return false
}

// SetMachineLearning gets a reference to the given bool and assigns it to the MachineLearning field.
func (o *FlavorView) SetMachineLearning(v bool) {
	o.MachineLearning = &v
}

// GetNice returns the Nice field value if set, zero value otherwise.
func (o *FlavorView) GetNice() int32 {
	if o == nil || o.Nice == nil {
		var ret int32
		return ret
	}
	return *o.Nice
}

// GetNiceOk returns a tuple with the Nice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetNiceOk() (*int32, bool) {
	if o == nil || o.Nice == nil {
		return nil, false
	}
	return o.Nice, true
}

// HasNice returns a boolean if a field has been set.
func (o *FlavorView) HasNice() bool {
	if o != nil && o.Nice != nil {
		return true
	}

	return false
}

// SetNice gets a reference to the given int32 and assigns it to the Nice field.
func (o *FlavorView) SetNice(v int32) {
	o.Nice = &v
}

// GetPriceId returns the PriceId field value if set, zero value otherwise.
func (o *FlavorView) GetPriceId() string {
	if o == nil || o.PriceId == nil {
		var ret string
		return ret
	}
	return *o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetPriceIdOk() (*string, bool) {
	if o == nil || o.PriceId == nil {
		return nil, false
	}
	return o.PriceId, true
}

// HasPriceId returns a boolean if a field has been set.
func (o *FlavorView) HasPriceId() bool {
	if o != nil && o.PriceId != nil {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given string and assigns it to the PriceId field.
func (o *FlavorView) SetPriceId(v string) {
	o.PriceId = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *FlavorView) GetMemory() ValueWithUnit {
	if o == nil || o.Memory == nil {
		var ret ValueWithUnit
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetMemoryOk() (*ValueWithUnit, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *FlavorView) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given ValueWithUnit and assigns it to the Memory field.
func (o *FlavorView) SetMemory(v ValueWithUnit) {
	o.Memory = &v
}

// GetRbdimage returns the Rbdimage field value if set, zero value otherwise.
func (o *FlavorView) GetRbdimage() string {
	if o == nil || o.Rbdimage == nil {
		var ret string
		return ret
	}
	return *o.Rbdimage
}

// GetRbdimageOk returns a tuple with the Rbdimage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorView) GetRbdimageOk() (*string, bool) {
	if o == nil || o.Rbdimage == nil {
		return nil, false
	}
	return o.Rbdimage, true
}

// HasRbdimage returns a boolean if a field has been set.
func (o *FlavorView) HasRbdimage() bool {
	if o != nil && o.Rbdimage != nil {
		return true
	}

	return false
}

// SetRbdimage gets a reference to the given string and assigns it to the Rbdimage field.
func (o *FlavorView) SetRbdimage(v string) {
	o.Rbdimage = &v
}

func (o FlavorView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Mem != nil {
		toSerialize["mem"] = o.Mem
	}
	if o.Cpus != nil {
		toSerialize["cpus"] = o.Cpus
	}
	if o.Gpus != nil {
		toSerialize["gpus"] = o.Gpus
	}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Microservice != nil {
		toSerialize["microservice"] = o.Microservice
	}
	if o.MachineLearning != nil {
		toSerialize["machine_learning"] = o.MachineLearning
	}
	if o.Nice != nil {
		toSerialize["nice"] = o.Nice
	}
	if o.PriceId != nil {
		toSerialize["price_id"] = o.PriceId
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Rbdimage != nil {
		toSerialize["rbdimage"] = o.Rbdimage
	}
	return json.Marshal(toSerialize)
}

type NullableFlavorView struct {
	value *FlavorView
	isSet bool
}

func (v NullableFlavorView) Get() *FlavorView {
	return v.value
}

func (v *NullableFlavorView) Set(val *FlavorView) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorView) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorView(val *FlavorView) *NullableFlavorView {
	return &NullableFlavorView{value: val, isSet: true}
}

func (v NullableFlavorView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
