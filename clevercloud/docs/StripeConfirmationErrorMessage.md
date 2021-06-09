# StripeConfirmationErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PaymentIntent** | Pointer to [**SetupIntentView**](SetupIntentView.md) |  | [optional] 

## Methods

### NewStripeConfirmationErrorMessage

`func NewStripeConfirmationErrorMessage() *StripeConfirmationErrorMessage`

NewStripeConfirmationErrorMessage instantiates a new StripeConfirmationErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeConfirmationErrorMessageWithDefaults

`func NewStripeConfirmationErrorMessageWithDefaults() *StripeConfirmationErrorMessage`

NewStripeConfirmationErrorMessageWithDefaults instantiates a new StripeConfirmationErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeConfirmationErrorMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeConfirmationErrorMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeConfirmationErrorMessage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StripeConfirmationErrorMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *StripeConfirmationErrorMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StripeConfirmationErrorMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StripeConfirmationErrorMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StripeConfirmationErrorMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *StripeConfirmationErrorMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeConfirmationErrorMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeConfirmationErrorMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeConfirmationErrorMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPaymentIntent

`func (o *StripeConfirmationErrorMessage) GetPaymentIntent() SetupIntentView`

GetPaymentIntent returns the PaymentIntent field if non-nil, zero value otherwise.

### GetPaymentIntentOk

`func (o *StripeConfirmationErrorMessage) GetPaymentIntentOk() (*SetupIntentView, bool)`

GetPaymentIntentOk returns a tuple with the PaymentIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntent

`func (o *StripeConfirmationErrorMessage) SetPaymentIntent(v SetupIntentView)`

SetPaymentIntent sets PaymentIntent field to given value.

### HasPaymentIntent

`func (o *StripeConfirmationErrorMessage) HasPaymentIntent() bool`

HasPaymentIntent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


