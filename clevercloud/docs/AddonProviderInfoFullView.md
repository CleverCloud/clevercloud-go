# AddonProviderInfoFullView

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
**Plans** | Pointer to [**[]AddonPlanView**](AddonPlanView.md) |  | [optional] 
**Features** | Pointer to [**[]AddonFeatureView**](AddonFeatureView.md) |  | [optional] 

## Methods

### NewAddonProviderInfoFullView

`func NewAddonProviderInfoFullView() *AddonProviderInfoFullView`

NewAddonProviderInfoFullView instantiates a new AddonProviderInfoFullView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonProviderInfoFullViewWithDefaults

`func NewAddonProviderInfoFullViewWithDefaults() *AddonProviderInfoFullView`

NewAddonProviderInfoFullViewWithDefaults instantiates a new AddonProviderInfoFullView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonProviderInfoFullView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonProviderInfoFullView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonProviderInfoFullView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonProviderInfoFullView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddonProviderInfoFullView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonProviderInfoFullView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonProviderInfoFullView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddonProviderInfoFullView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWebsite

`func (o *AddonProviderInfoFullView) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AddonProviderInfoFullView) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AddonProviderInfoFullView) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AddonProviderInfoFullView) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetSupportEmail

`func (o *AddonProviderInfoFullView) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *AddonProviderInfoFullView) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *AddonProviderInfoFullView) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *AddonProviderInfoFullView) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetGooglePlusName

`func (o *AddonProviderInfoFullView) GetGooglePlusName() string`

GetGooglePlusName returns the GooglePlusName field if non-nil, zero value otherwise.

### GetGooglePlusNameOk

`func (o *AddonProviderInfoFullView) GetGooglePlusNameOk() (*string, bool)`

GetGooglePlusNameOk returns a tuple with the GooglePlusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlusName

`func (o *AddonProviderInfoFullView) SetGooglePlusName(v string)`

SetGooglePlusName sets GooglePlusName field to given value.

### HasGooglePlusName

`func (o *AddonProviderInfoFullView) HasGooglePlusName() bool`

HasGooglePlusName returns a boolean if a field has been set.

### GetTwitterName

`func (o *AddonProviderInfoFullView) GetTwitterName() string`

GetTwitterName returns the TwitterName field if non-nil, zero value otherwise.

### GetTwitterNameOk

`func (o *AddonProviderInfoFullView) GetTwitterNameOk() (*string, bool)`

GetTwitterNameOk returns a tuple with the TwitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterName

`func (o *AddonProviderInfoFullView) SetTwitterName(v string)`

SetTwitterName sets TwitterName field to given value.

### HasTwitterName

`func (o *AddonProviderInfoFullView) HasTwitterName() bool`

HasTwitterName returns a boolean if a field has been set.

### GetAnalyticsId

`func (o *AddonProviderInfoFullView) GetAnalyticsId() string`

GetAnalyticsId returns the AnalyticsId field if non-nil, zero value otherwise.

### GetAnalyticsIdOk

`func (o *AddonProviderInfoFullView) GetAnalyticsIdOk() (*string, bool)`

GetAnalyticsIdOk returns a tuple with the AnalyticsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsId

`func (o *AddonProviderInfoFullView) SetAnalyticsId(v string)`

SetAnalyticsId sets AnalyticsId field to given value.

### HasAnalyticsId

`func (o *AddonProviderInfoFullView) HasAnalyticsId() bool`

HasAnalyticsId returns a boolean if a field has been set.

### GetShortDesc

`func (o *AddonProviderInfoFullView) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *AddonProviderInfoFullView) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *AddonProviderInfoFullView) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *AddonProviderInfoFullView) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### GetLongDesc

`func (o *AddonProviderInfoFullView) GetLongDesc() string`

GetLongDesc returns the LongDesc field if non-nil, zero value otherwise.

### GetLongDescOk

`func (o *AddonProviderInfoFullView) GetLongDescOk() (*string, bool)`

GetLongDescOk returns a tuple with the LongDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDesc

`func (o *AddonProviderInfoFullView) SetLongDesc(v string)`

SetLongDesc sets LongDesc field to given value.

### HasLongDesc

`func (o *AddonProviderInfoFullView) HasLongDesc() bool`

HasLongDesc returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AddonProviderInfoFullView) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AddonProviderInfoFullView) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AddonProviderInfoFullView) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AddonProviderInfoFullView) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *AddonProviderInfoFullView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddonProviderInfoFullView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddonProviderInfoFullView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddonProviderInfoFullView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOpenInNewTab

`func (o *AddonProviderInfoFullView) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *AddonProviderInfoFullView) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *AddonProviderInfoFullView) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.

### HasOpenInNewTab

`func (o *AddonProviderInfoFullView) HasOpenInNewTab() bool`

HasOpenInNewTab returns a boolean if a field has been set.

### GetCanUpgrade

`func (o *AddonProviderInfoFullView) GetCanUpgrade() bool`

GetCanUpgrade returns the CanUpgrade field if non-nil, zero value otherwise.

### GetCanUpgradeOk

`func (o *AddonProviderInfoFullView) GetCanUpgradeOk() (*bool, bool)`

GetCanUpgradeOk returns a tuple with the CanUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgrade

`func (o *AddonProviderInfoFullView) SetCanUpgrade(v bool)`

SetCanUpgrade sets CanUpgrade field to given value.

### HasCanUpgrade

`func (o *AddonProviderInfoFullView) HasCanUpgrade() bool`

HasCanUpgrade returns a boolean if a field has been set.

### GetRegions

`func (o *AddonProviderInfoFullView) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AddonProviderInfoFullView) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AddonProviderInfoFullView) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AddonProviderInfoFullView) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetPlans

`func (o *AddonProviderInfoFullView) GetPlans() []AddonPlanView`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AddonProviderInfoFullView) GetPlansOk() (*[]AddonPlanView, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AddonProviderInfoFullView) SetPlans(v []AddonPlanView)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AddonProviderInfoFullView) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetFeatures

`func (o *AddonProviderInfoFullView) GetFeatures() []AddonFeatureView`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AddonProviderInfoFullView) GetFeaturesOk() (*[]AddonFeatureView, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AddonProviderInfoFullView) SetFeatures(v []AddonFeatureView)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AddonProviderInfoFullView) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


