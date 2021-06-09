# DropCountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **float32** |  | [optional] 
**DropPrice** | Pointer to [**DropPriceView**](DropPriceView.md) |  | [optional] 

## Methods

### NewDropCountView

`func NewDropCountView() *DropCountView`

NewDropCountView instantiates a new DropCountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropCountViewWithDefaults

`func NewDropCountViewWithDefaults() *DropCountView`

NewDropCountViewWithDefaults instantiates a new DropCountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *DropCountView) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DropCountView) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DropCountView) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DropCountView) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetCount

`func (o *DropCountView) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DropCountView) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DropCountView) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DropCountView) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDropPrice

`func (o *DropCountView) GetDropPrice() DropPriceView`

GetDropPrice returns the DropPrice field if non-nil, zero value otherwise.

### GetDropPriceOk

`func (o *DropCountView) GetDropPriceOk() (*DropPriceView, bool)`

GetDropPriceOk returns a tuple with the DropPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPrice

`func (o *DropCountView) SetDropPrice(v DropPriceView)`

SetDropPrice sets DropPrice field to given value.

### HasDropPrice

`func (o *DropCountView) HasDropPrice() bool`

HasDropPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


