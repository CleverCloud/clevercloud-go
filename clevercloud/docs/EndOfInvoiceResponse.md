# EndOfInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**PosData** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**BtcPrice** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**InvoiceTime** | Pointer to **int64** |  | [optional] 
**CurrentTime** | Pointer to **int64** |  | [optional] 
**ExpirationTime** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to [**EndOfInvoiceError**](EndOfInvoiceError.md) |  | [optional] 

## Methods

### NewEndOfInvoiceResponse

`func NewEndOfInvoiceResponse() *EndOfInvoiceResponse`

NewEndOfInvoiceResponse instantiates a new EndOfInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndOfInvoiceResponseWithDefaults

`func NewEndOfInvoiceResponseWithDefaults() *EndOfInvoiceResponse`

NewEndOfInvoiceResponseWithDefaults instantiates a new EndOfInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndOfInvoiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndOfInvoiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndOfInvoiceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndOfInvoiceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *EndOfInvoiceResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EndOfInvoiceResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EndOfInvoiceResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EndOfInvoiceResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPosData

`func (o *EndOfInvoiceResponse) GetPosData() string`

GetPosData returns the PosData field if non-nil, zero value otherwise.

### GetPosDataOk

`func (o *EndOfInvoiceResponse) GetPosDataOk() (*string, bool)`

GetPosDataOk returns a tuple with the PosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosData

`func (o *EndOfInvoiceResponse) SetPosData(v string)`

SetPosData sets PosData field to given value.

### HasPosData

`func (o *EndOfInvoiceResponse) HasPosData() bool`

HasPosData returns a boolean if a field has been set.

### GetStatus

`func (o *EndOfInvoiceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndOfInvoiceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndOfInvoiceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EndOfInvoiceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBtcPrice

`func (o *EndOfInvoiceResponse) GetBtcPrice() string`

GetBtcPrice returns the BtcPrice field if non-nil, zero value otherwise.

### GetBtcPriceOk

`func (o *EndOfInvoiceResponse) GetBtcPriceOk() (*string, bool)`

GetBtcPriceOk returns a tuple with the BtcPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtcPrice

`func (o *EndOfInvoiceResponse) SetBtcPrice(v string)`

SetBtcPrice sets BtcPrice field to given value.

### HasBtcPrice

`func (o *EndOfInvoiceResponse) HasBtcPrice() bool`

HasBtcPrice returns a boolean if a field has been set.

### GetPrice

`func (o *EndOfInvoiceResponse) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EndOfInvoiceResponse) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EndOfInvoiceResponse) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EndOfInvoiceResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *EndOfInvoiceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EndOfInvoiceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EndOfInvoiceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EndOfInvoiceResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInvoiceTime

`func (o *EndOfInvoiceResponse) GetInvoiceTime() int64`

GetInvoiceTime returns the InvoiceTime field if non-nil, zero value otherwise.

### GetInvoiceTimeOk

`func (o *EndOfInvoiceResponse) GetInvoiceTimeOk() (*int64, bool)`

GetInvoiceTimeOk returns a tuple with the InvoiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTime

`func (o *EndOfInvoiceResponse) SetInvoiceTime(v int64)`

SetInvoiceTime sets InvoiceTime field to given value.

### HasInvoiceTime

`func (o *EndOfInvoiceResponse) HasInvoiceTime() bool`

HasInvoiceTime returns a boolean if a field has been set.

### GetCurrentTime

`func (o *EndOfInvoiceResponse) GetCurrentTime() int64`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *EndOfInvoiceResponse) GetCurrentTimeOk() (*int64, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *EndOfInvoiceResponse) SetCurrentTime(v int64)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *EndOfInvoiceResponse) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *EndOfInvoiceResponse) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *EndOfInvoiceResponse) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *EndOfInvoiceResponse) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *EndOfInvoiceResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetError

`func (o *EndOfInvoiceResponse) GetError() EndOfInvoiceError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EndOfInvoiceResponse) GetErrorOk() (*EndOfInvoiceError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EndOfInvoiceResponse) SetError(v EndOfInvoiceError)`

SetError sets Error field to given value.

### HasError

`func (o *EndOfInvoiceResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


