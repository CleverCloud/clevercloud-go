# PriceWithTaxInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **int64** |  | [optional] 
**TaxFreePrice** | Pointer to **int64** |  | [optional] 
**VatRatio** | Pointer to **int64** |  | [optional] 
**InvoicedOwner** | Pointer to **string** |  | [optional] 

## Methods

### NewPriceWithTaxInfo

`func NewPriceWithTaxInfo() *PriceWithTaxInfo`

NewPriceWithTaxInfo instantiates a new PriceWithTaxInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithTaxInfoWithDefaults

`func NewPriceWithTaxInfoWithDefaults() *PriceWithTaxInfo`

NewPriceWithTaxInfoWithDefaults instantiates a new PriceWithTaxInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PriceWithTaxInfo) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceWithTaxInfo) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceWithTaxInfo) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceWithTaxInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxFreePrice

`func (o *PriceWithTaxInfo) GetTaxFreePrice() int64`

GetTaxFreePrice returns the TaxFreePrice field if non-nil, zero value otherwise.

### GetTaxFreePriceOk

`func (o *PriceWithTaxInfo) GetTaxFreePriceOk() (*int64, bool)`

GetTaxFreePriceOk returns a tuple with the TaxFreePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFreePrice

`func (o *PriceWithTaxInfo) SetTaxFreePrice(v int64)`

SetTaxFreePrice sets TaxFreePrice field to given value.

### HasTaxFreePrice

`func (o *PriceWithTaxInfo) HasTaxFreePrice() bool`

HasTaxFreePrice returns a boolean if a field has been set.

### GetVatRatio

`func (o *PriceWithTaxInfo) GetVatRatio() int64`

GetVatRatio returns the VatRatio field if non-nil, zero value otherwise.

### GetVatRatioOk

`func (o *PriceWithTaxInfo) GetVatRatioOk() (*int64, bool)`

GetVatRatioOk returns a tuple with the VatRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRatio

`func (o *PriceWithTaxInfo) SetVatRatio(v int64)`

SetVatRatio sets VatRatio field to given value.

### HasVatRatio

`func (o *PriceWithTaxInfo) HasVatRatio() bool`

HasVatRatio returns a boolean if a field has been set.

### GetInvoicedOwner

`func (o *PriceWithTaxInfo) GetInvoicedOwner() string`

GetInvoicedOwner returns the InvoicedOwner field if non-nil, zero value otherwise.

### GetInvoicedOwnerOk

`func (o *PriceWithTaxInfo) GetInvoicedOwnerOk() (*string, bool)`

GetInvoicedOwnerOk returns a tuple with the InvoicedOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedOwner

`func (o *PriceWithTaxInfo) SetInvoicedOwner(v string)`

SetInvoicedOwner sets InvoicedOwner field to given value.

### HasInvoicedOwner

`func (o *PriceWithTaxInfo) HasInvoicedOwner() bool`

HasInvoicedOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


