# EndOfInvoiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEndOfInvoiceError

`func NewEndOfInvoiceError() *EndOfInvoiceError`

NewEndOfInvoiceError instantiates a new EndOfInvoiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndOfInvoiceErrorWithDefaults

`func NewEndOfInvoiceErrorWithDefaults() *EndOfInvoiceError`

NewEndOfInvoiceErrorWithDefaults instantiates a new EndOfInvoiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EndOfInvoiceError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EndOfInvoiceError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EndOfInvoiceError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EndOfInvoiceError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *EndOfInvoiceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EndOfInvoiceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EndOfInvoiceError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EndOfInvoiceError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessages

`func (o *EndOfInvoiceError) GetMessages() map[string]string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *EndOfInvoiceError) GetMessagesOk() (*map[string]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *EndOfInvoiceError) SetMessages(v map[string]string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *EndOfInvoiceError) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


