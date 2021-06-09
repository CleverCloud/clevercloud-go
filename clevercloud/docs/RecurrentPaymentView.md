# RecurrentPaymentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserView**](UserView.md) |  | [optional] 
**Target** | Pointer to [**OwnerView**](OwnerView.md) |  | [optional] 
**ThresholdAmount** | Pointer to **float32** |  | [optional] 
**MonthlyAmount** | Pointer to **float32** |  | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewRecurrentPaymentView

`func NewRecurrentPaymentView() *RecurrentPaymentView`

NewRecurrentPaymentView instantiates a new RecurrentPaymentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrentPaymentViewWithDefaults

`func NewRecurrentPaymentViewWithDefaults() *RecurrentPaymentView`

NewRecurrentPaymentViewWithDefaults instantiates a new RecurrentPaymentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *RecurrentPaymentView) GetUser() UserView`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RecurrentPaymentView) GetUserOk() (*UserView, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RecurrentPaymentView) SetUser(v UserView)`

SetUser sets User field to given value.

### HasUser

`func (o *RecurrentPaymentView) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTarget

`func (o *RecurrentPaymentView) GetTarget() OwnerView`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RecurrentPaymentView) GetTargetOk() (*OwnerView, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RecurrentPaymentView) SetTarget(v OwnerView)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RecurrentPaymentView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetThresholdAmount

`func (o *RecurrentPaymentView) GetThresholdAmount() float32`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *RecurrentPaymentView) GetThresholdAmountOk() (*float32, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *RecurrentPaymentView) SetThresholdAmount(v float32)`

SetThresholdAmount sets ThresholdAmount field to given value.

### HasThresholdAmount

`func (o *RecurrentPaymentView) HasThresholdAmount() bool`

HasThresholdAmount returns a boolean if a field has been set.

### GetMonthlyAmount

`func (o *RecurrentPaymentView) GetMonthlyAmount() float32`

GetMonthlyAmount returns the MonthlyAmount field if non-nil, zero value otherwise.

### GetMonthlyAmountOk

`func (o *RecurrentPaymentView) GetMonthlyAmountOk() (*float32, bool)`

GetMonthlyAmountOk returns a tuple with the MonthlyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyAmount

`func (o *RecurrentPaymentView) SetMonthlyAmount(v float32)`

SetMonthlyAmount sets MonthlyAmount field to given value.

### HasMonthlyAmount

`func (o *RecurrentPaymentView) HasMonthlyAmount() bool`

HasMonthlyAmount returns a boolean if a field has been set.

### GetThreshold

`func (o *RecurrentPaymentView) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RecurrentPaymentView) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RecurrentPaymentView) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RecurrentPaymentView) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCurrency

`func (o *RecurrentPaymentView) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RecurrentPaymentView) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RecurrentPaymentView) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RecurrentPaymentView) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetToken

`func (o *RecurrentPaymentView) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RecurrentPaymentView) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RecurrentPaymentView) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RecurrentPaymentView) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


