# FlavorView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Mem** | Pointer to **int32** |  | [optional] 
**Cpus** | Pointer to **int32** |  | [optional] 
**Gpus** | Pointer to **int32** |  | [optional] 
**Disk** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**Microservice** | Pointer to **bool** |  | [optional] 
**MachineLearning** | Pointer to **bool** |  | [optional] 
**Nice** | Pointer to **int32** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**Memory** | Pointer to [**ValueWithUnit**](ValueWithUnit.md) |  | [optional] 
**Rbdimage** | Pointer to **string** |  | [optional] 

## Methods

### NewFlavorView

`func NewFlavorView() *FlavorView`

NewFlavorView instantiates a new FlavorView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorViewWithDefaults

`func NewFlavorViewWithDefaults() *FlavorView`

NewFlavorViewWithDefaults instantiates a new FlavorView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlavorView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlavorView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlavorView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlavorView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMem

`func (o *FlavorView) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *FlavorView) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *FlavorView) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *FlavorView) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetCpus

`func (o *FlavorView) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *FlavorView) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *FlavorView) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *FlavorView) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetGpus

`func (o *FlavorView) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *FlavorView) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *FlavorView) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *FlavorView) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetDisk

`func (o *FlavorView) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *FlavorView) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *FlavorView) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *FlavorView) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetPrice

`func (o *FlavorView) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FlavorView) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FlavorView) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FlavorView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAvailable

`func (o *FlavorView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *FlavorView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *FlavorView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *FlavorView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetMicroservice

`func (o *FlavorView) GetMicroservice() bool`

GetMicroservice returns the Microservice field if non-nil, zero value otherwise.

### GetMicroserviceOk

`func (o *FlavorView) GetMicroserviceOk() (*bool, bool)`

GetMicroserviceOk returns a tuple with the Microservice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroservice

`func (o *FlavorView) SetMicroservice(v bool)`

SetMicroservice sets Microservice field to given value.

### HasMicroservice

`func (o *FlavorView) HasMicroservice() bool`

HasMicroservice returns a boolean if a field has been set.

### GetMachineLearning

`func (o *FlavorView) GetMachineLearning() bool`

GetMachineLearning returns the MachineLearning field if non-nil, zero value otherwise.

### GetMachineLearningOk

`func (o *FlavorView) GetMachineLearningOk() (*bool, bool)`

GetMachineLearningOk returns a tuple with the MachineLearning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearning

`func (o *FlavorView) SetMachineLearning(v bool)`

SetMachineLearning sets MachineLearning field to given value.

### HasMachineLearning

`func (o *FlavorView) HasMachineLearning() bool`

HasMachineLearning returns a boolean if a field has been set.

### GetNice

`func (o *FlavorView) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *FlavorView) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *FlavorView) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *FlavorView) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetPriceId

`func (o *FlavorView) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *FlavorView) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *FlavorView) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *FlavorView) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetMemory

`func (o *FlavorView) GetMemory() ValueWithUnit`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *FlavorView) GetMemoryOk() (*ValueWithUnit, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *FlavorView) SetMemory(v ValueWithUnit)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *FlavorView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetRbdimage

`func (o *FlavorView) GetRbdimage() string`

GetRbdimage returns the Rbdimage field if non-nil, zero value otherwise.

### GetRbdimageOk

`func (o *FlavorView) GetRbdimageOk() (*string, bool)`

GetRbdimageOk returns a tuple with the Rbdimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbdimage

`func (o *FlavorView) SetRbdimage(v string)`

SetRbdimage sets Rbdimage field to given value.

### HasRbdimage

`func (o *FlavorView) HasRbdimage() bool`

HasRbdimage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


