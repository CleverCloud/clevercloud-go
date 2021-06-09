# AddonProviderInfoView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**SupportEmail** | Pointer to **string** |  | [optional] 
**GooglePlusName** | Pointer to **string** |  | [optional] 
**TwitterName** | Pointer to **string** |  | [optional] 
**AnalyticsId** | Pointer to **string** |  | [optional] 
**ShortDesc** | Pointer to **string** |  | [optional] 
**LongDesc** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OpenInNewTab** | Pointer to **bool** |  | [optional] 
**CanUpgrade** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddonProviderInfoView

`func NewAddonProviderInfoView() *AddonProviderInfoView`

NewAddonProviderInfoView instantiates a new AddonProviderInfoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonProviderInfoViewWithDefaults

`func NewAddonProviderInfoViewWithDefaults() *AddonProviderInfoView`

NewAddonProviderInfoViewWithDefaults instantiates a new AddonProviderInfoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonProviderInfoView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonProviderInfoView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonProviderInfoView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonProviderInfoView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddonProviderInfoView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonProviderInfoView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonProviderInfoView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddonProviderInfoView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWebsite

`func (o *AddonProviderInfoView) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AddonProviderInfoView) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AddonProviderInfoView) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AddonProviderInfoView) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetSupportEmail

`func (o *AddonProviderInfoView) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *AddonProviderInfoView) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *AddonProviderInfoView) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *AddonProviderInfoView) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetGooglePlusName

`func (o *AddonProviderInfoView) GetGooglePlusName() string`

GetGooglePlusName returns the GooglePlusName field if non-nil, zero value otherwise.

### GetGooglePlusNameOk

`func (o *AddonProviderInfoView) GetGooglePlusNameOk() (*string, bool)`

GetGooglePlusNameOk returns a tuple with the GooglePlusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlusName

`func (o *AddonProviderInfoView) SetGooglePlusName(v string)`

SetGooglePlusName sets GooglePlusName field to given value.

### HasGooglePlusName

`func (o *AddonProviderInfoView) HasGooglePlusName() bool`

HasGooglePlusName returns a boolean if a field has been set.

### GetTwitterName

`func (o *AddonProviderInfoView) GetTwitterName() string`

GetTwitterName returns the TwitterName field if non-nil, zero value otherwise.

### GetTwitterNameOk

`func (o *AddonProviderInfoView) GetTwitterNameOk() (*string, bool)`

GetTwitterNameOk returns a tuple with the TwitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterName

`func (o *AddonProviderInfoView) SetTwitterName(v string)`

SetTwitterName sets TwitterName field to given value.

### HasTwitterName

`func (o *AddonProviderInfoView) HasTwitterName() bool`

HasTwitterName returns a boolean if a field has been set.

### GetAnalyticsId

`func (o *AddonProviderInfoView) GetAnalyticsId() string`

GetAnalyticsId returns the AnalyticsId field if non-nil, zero value otherwise.

### GetAnalyticsIdOk

`func (o *AddonProviderInfoView) GetAnalyticsIdOk() (*string, bool)`

GetAnalyticsIdOk returns a tuple with the AnalyticsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsId

`func (o *AddonProviderInfoView) SetAnalyticsId(v string)`

SetAnalyticsId sets AnalyticsId field to given value.

### HasAnalyticsId

`func (o *AddonProviderInfoView) HasAnalyticsId() bool`

HasAnalyticsId returns a boolean if a field has been set.

### GetShortDesc

`func (o *AddonProviderInfoView) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *AddonProviderInfoView) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *AddonProviderInfoView) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *AddonProviderInfoView) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### GetLongDesc

`func (o *AddonProviderInfoView) GetLongDesc() string`

GetLongDesc returns the LongDesc field if non-nil, zero value otherwise.

### GetLongDescOk

`func (o *AddonProviderInfoView) GetLongDescOk() (*string, bool)`

GetLongDescOk returns a tuple with the LongDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDesc

`func (o *AddonProviderInfoView) SetLongDesc(v string)`

SetLongDesc sets LongDesc field to given value.

### HasLongDesc

`func (o *AddonProviderInfoView) HasLongDesc() bool`

HasLongDesc returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AddonProviderInfoView) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AddonProviderInfoView) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AddonProviderInfoView) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AddonProviderInfoView) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *AddonProviderInfoView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddonProviderInfoView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddonProviderInfoView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddonProviderInfoView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOpenInNewTab

`func (o *AddonProviderInfoView) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *AddonProviderInfoView) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *AddonProviderInfoView) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.

### HasOpenInNewTab

`func (o *AddonProviderInfoView) HasOpenInNewTab() bool`

HasOpenInNewTab returns a boolean if a field has been set.

### GetCanUpgrade

`func (o *AddonProviderInfoView) GetCanUpgrade() bool`

GetCanUpgrade returns the CanUpgrade field if non-nil, zero value otherwise.

### GetCanUpgradeOk

`func (o *AddonProviderInfoView) GetCanUpgradeOk() (*bool, bool)`

GetCanUpgradeOk returns a tuple with the CanUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgrade

`func (o *AddonProviderInfoView) SetCanUpgrade(v bool)`

SetCanUpgrade sets CanUpgrade field to given value.

### HasCanUpgrade

`func (o *AddonProviderInfoView) HasCanUpgrade() bool`

HasCanUpgrade returns a boolean if a field has been set.

### GetRegions

`func (o *AddonProviderInfoView) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AddonProviderInfoView) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AddonProviderInfoView) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AddonProviderInfoView) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


