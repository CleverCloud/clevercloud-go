# WannabeAddonProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **string** | Id of the add-on provider | 
**Plan** | **string** | Id of the price plan | 
**LinkedApp** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the future add-on, for display. | [optional] 
**Region** | **string** | Region to provision the add-on in. | 
**Options** | Pointer to **map[string]string** | Options to add to the provision call. | [optional] 
**Version** | Pointer to **string** | Version of the add-on | [optional] 
**PaymentIntent** | Pointer to [**SetupIntentView**](SetupIntentView.md) |  | [optional] 
**PaymentMethodType** | Pointer to **string** | Payment method type | [optional] 
**SepaSourceId** | Pointer to **string** | Id of the SEPA debit source | [optional] 

## Methods

### NewWannabeAddonProvision

`func NewWannabeAddonProvision(providerId string, plan string, region string, ) *WannabeAddonProvision`

NewWannabeAddonProvision instantiates a new WannabeAddonProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWannabeAddonProvisionWithDefaults

`func NewWannabeAddonProvisionWithDefaults() *WannabeAddonProvision`

NewWannabeAddonProvisionWithDefaults instantiates a new WannabeAddonProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *WannabeAddonProvision) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *WannabeAddonProvision) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *WannabeAddonProvision) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetPlan

`func (o *WannabeAddonProvision) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *WannabeAddonProvision) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *WannabeAddonProvision) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetLinkedApp

`func (o *WannabeAddonProvision) GetLinkedApp() string`

GetLinkedApp returns the LinkedApp field if non-nil, zero value otherwise.

### GetLinkedAppOk

`func (o *WannabeAddonProvision) GetLinkedAppOk() (*string, bool)`

GetLinkedAppOk returns a tuple with the LinkedApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedApp

`func (o *WannabeAddonProvision) SetLinkedApp(v string)`

SetLinkedApp sets LinkedApp field to given value.

### HasLinkedApp

`func (o *WannabeAddonProvision) HasLinkedApp() bool`

HasLinkedApp returns a boolean if a field has been set.

### GetName

`func (o *WannabeAddonProvision) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WannabeAddonProvision) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WannabeAddonProvision) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WannabeAddonProvision) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *WannabeAddonProvision) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WannabeAddonProvision) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WannabeAddonProvision) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetOptions

`func (o *WannabeAddonProvision) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WannabeAddonProvision) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WannabeAddonProvision) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *WannabeAddonProvision) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVersion

`func (o *WannabeAddonProvision) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WannabeAddonProvision) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WannabeAddonProvision) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WannabeAddonProvision) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPaymentIntent

`func (o *WannabeAddonProvision) GetPaymentIntent() SetupIntentView`

GetPaymentIntent returns the PaymentIntent field if non-nil, zero value otherwise.

### GetPaymentIntentOk

`func (o *WannabeAddonProvision) GetPaymentIntentOk() (*SetupIntentView, bool)`

GetPaymentIntentOk returns a tuple with the PaymentIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntent

`func (o *WannabeAddonProvision) SetPaymentIntent(v SetupIntentView)`

SetPaymentIntent sets PaymentIntent field to given value.

### HasPaymentIntent

`func (o *WannabeAddonProvision) HasPaymentIntent() bool`

HasPaymentIntent returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *WannabeAddonProvision) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *WannabeAddonProvision) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *WannabeAddonProvision) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *WannabeAddonProvision) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetSepaSourceId

`func (o *WannabeAddonProvision) GetSepaSourceId() string`

GetSepaSourceId returns the SepaSourceId field if non-nil, zero value otherwise.

### GetSepaSourceIdOk

`func (o *WannabeAddonProvision) GetSepaSourceIdOk() (*string, bool)`

GetSepaSourceIdOk returns a tuple with the SepaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaSourceId

`func (o *WannabeAddonProvision) SetSepaSourceId(v string)`

SetSepaSourceId sets SepaSourceId field to given value.

### HasSepaSourceId

`func (o *WannabeAddonProvision) HasSepaSourceId() bool`

HasSepaSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


