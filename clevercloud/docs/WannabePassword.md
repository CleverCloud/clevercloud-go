# WannabePassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 
**DropTokens** | Pointer to **bool** |  | [optional] 

## Methods

### NewWannabePassword

`func NewWannabePassword() *WannabePassword`

NewWannabePassword instantiates a new WannabePassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWannabePasswordWithDefaults

`func NewWannabePasswordWithDefaults() *WannabePassword`

NewWannabePasswordWithDefaults instantiates a new WannabePassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *WannabePassword) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *WannabePassword) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *WannabePassword) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *WannabePassword) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *WannabePassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *WannabePassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *WannabePassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *WannabePassword) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetDropTokens

`func (o *WannabePassword) GetDropTokens() bool`

GetDropTokens returns the DropTokens field if non-nil, zero value otherwise.

### GetDropTokensOk

`func (o *WannabePassword) GetDropTokensOk() (*bool, bool)`

GetDropTokensOk returns a tuple with the DropTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropTokens

`func (o *WannabePassword) SetDropTokens(v bool)`

SetDropTokens sets DropTokens field to given value.

### HasDropTokens

`func (o *WannabePassword) HasDropTokens() bool`

HasDropTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


