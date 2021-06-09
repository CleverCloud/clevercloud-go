# OAuthRightsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Almighty** | Pointer to **bool** |  | [optional] 
**AccessOrganisations** | Pointer to **bool** |  | [optional] 
**ManageOrganisations** | Pointer to **bool** |  | [optional] 
**ManageOrganisationsServices** | Pointer to **bool** |  | [optional] 
**ManageOrganisationsApplications** | Pointer to **bool** |  | [optional] 
**ManageOrganisationsMembers** | Pointer to **bool** |  | [optional] 
**AccessOrganisationsBills** | Pointer to **bool** |  | [optional] 
**AccessOrganisationsCreditCount** | Pointer to **bool** |  | [optional] 
**AccessOrganisationsConsumptionStatistics** | Pointer to **bool** |  | [optional] 
**AccessPersonalInformation** | Pointer to **bool** |  | [optional] 
**ManagePersonalInformation** | Pointer to **bool** |  | [optional] 
**ManageSshKeys** | Pointer to **bool** |  | [optional] 

## Methods

### NewOAuthRightsView

`func NewOAuthRightsView() *OAuthRightsView`

NewOAuthRightsView instantiates a new OAuthRightsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRightsViewWithDefaults

`func NewOAuthRightsViewWithDefaults() *OAuthRightsView`

NewOAuthRightsViewWithDefaults instantiates a new OAuthRightsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlmighty

`func (o *OAuthRightsView) GetAlmighty() bool`

GetAlmighty returns the Almighty field if non-nil, zero value otherwise.

### GetAlmightyOk

`func (o *OAuthRightsView) GetAlmightyOk() (*bool, bool)`

GetAlmightyOk returns a tuple with the Almighty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlmighty

`func (o *OAuthRightsView) SetAlmighty(v bool)`

SetAlmighty sets Almighty field to given value.

### HasAlmighty

`func (o *OAuthRightsView) HasAlmighty() bool`

HasAlmighty returns a boolean if a field has been set.

### GetAccessOrganisations

`func (o *OAuthRightsView) GetAccessOrganisations() bool`

GetAccessOrganisations returns the AccessOrganisations field if non-nil, zero value otherwise.

### GetAccessOrganisationsOk

`func (o *OAuthRightsView) GetAccessOrganisationsOk() (*bool, bool)`

GetAccessOrganisationsOk returns a tuple with the AccessOrganisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOrganisations

`func (o *OAuthRightsView) SetAccessOrganisations(v bool)`

SetAccessOrganisations sets AccessOrganisations field to given value.

### HasAccessOrganisations

`func (o *OAuthRightsView) HasAccessOrganisations() bool`

HasAccessOrganisations returns a boolean if a field has been set.

### GetManageOrganisations

`func (o *OAuthRightsView) GetManageOrganisations() bool`

GetManageOrganisations returns the ManageOrganisations field if non-nil, zero value otherwise.

### GetManageOrganisationsOk

`func (o *OAuthRightsView) GetManageOrganisationsOk() (*bool, bool)`

GetManageOrganisationsOk returns a tuple with the ManageOrganisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageOrganisations

`func (o *OAuthRightsView) SetManageOrganisations(v bool)`

SetManageOrganisations sets ManageOrganisations field to given value.

### HasManageOrganisations

`func (o *OAuthRightsView) HasManageOrganisations() bool`

HasManageOrganisations returns a boolean if a field has been set.

### GetManageOrganisationsServices

`func (o *OAuthRightsView) GetManageOrganisationsServices() bool`

GetManageOrganisationsServices returns the ManageOrganisationsServices field if non-nil, zero value otherwise.

### GetManageOrganisationsServicesOk

`func (o *OAuthRightsView) GetManageOrganisationsServicesOk() (*bool, bool)`

GetManageOrganisationsServicesOk returns a tuple with the ManageOrganisationsServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageOrganisationsServices

`func (o *OAuthRightsView) SetManageOrganisationsServices(v bool)`

SetManageOrganisationsServices sets ManageOrganisationsServices field to given value.

### HasManageOrganisationsServices

`func (o *OAuthRightsView) HasManageOrganisationsServices() bool`

HasManageOrganisationsServices returns a boolean if a field has been set.

### GetManageOrganisationsApplications

`func (o *OAuthRightsView) GetManageOrganisationsApplications() bool`

GetManageOrganisationsApplications returns the ManageOrganisationsApplications field if non-nil, zero value otherwise.

### GetManageOrganisationsApplicationsOk

`func (o *OAuthRightsView) GetManageOrganisationsApplicationsOk() (*bool, bool)`

GetManageOrganisationsApplicationsOk returns a tuple with the ManageOrganisationsApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageOrganisationsApplications

`func (o *OAuthRightsView) SetManageOrganisationsApplications(v bool)`

SetManageOrganisationsApplications sets ManageOrganisationsApplications field to given value.

### HasManageOrganisationsApplications

`func (o *OAuthRightsView) HasManageOrganisationsApplications() bool`

HasManageOrganisationsApplications returns a boolean if a field has been set.

### GetManageOrganisationsMembers

`func (o *OAuthRightsView) GetManageOrganisationsMembers() bool`

GetManageOrganisationsMembers returns the ManageOrganisationsMembers field if non-nil, zero value otherwise.

### GetManageOrganisationsMembersOk

`func (o *OAuthRightsView) GetManageOrganisationsMembersOk() (*bool, bool)`

GetManageOrganisationsMembersOk returns a tuple with the ManageOrganisationsMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageOrganisationsMembers

`func (o *OAuthRightsView) SetManageOrganisationsMembers(v bool)`

SetManageOrganisationsMembers sets ManageOrganisationsMembers field to given value.

### HasManageOrganisationsMembers

`func (o *OAuthRightsView) HasManageOrganisationsMembers() bool`

HasManageOrganisationsMembers returns a boolean if a field has been set.

### GetAccessOrganisationsBills

`func (o *OAuthRightsView) GetAccessOrganisationsBills() bool`

GetAccessOrganisationsBills returns the AccessOrganisationsBills field if non-nil, zero value otherwise.

### GetAccessOrganisationsBillsOk

`func (o *OAuthRightsView) GetAccessOrganisationsBillsOk() (*bool, bool)`

GetAccessOrganisationsBillsOk returns a tuple with the AccessOrganisationsBills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOrganisationsBills

`func (o *OAuthRightsView) SetAccessOrganisationsBills(v bool)`

SetAccessOrganisationsBills sets AccessOrganisationsBills field to given value.

### HasAccessOrganisationsBills

`func (o *OAuthRightsView) HasAccessOrganisationsBills() bool`

HasAccessOrganisationsBills returns a boolean if a field has been set.

### GetAccessOrganisationsCreditCount

`func (o *OAuthRightsView) GetAccessOrganisationsCreditCount() bool`

GetAccessOrganisationsCreditCount returns the AccessOrganisationsCreditCount field if non-nil, zero value otherwise.

### GetAccessOrganisationsCreditCountOk

`func (o *OAuthRightsView) GetAccessOrganisationsCreditCountOk() (*bool, bool)`

GetAccessOrganisationsCreditCountOk returns a tuple with the AccessOrganisationsCreditCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOrganisationsCreditCount

`func (o *OAuthRightsView) SetAccessOrganisationsCreditCount(v bool)`

SetAccessOrganisationsCreditCount sets AccessOrganisationsCreditCount field to given value.

### HasAccessOrganisationsCreditCount

`func (o *OAuthRightsView) HasAccessOrganisationsCreditCount() bool`

HasAccessOrganisationsCreditCount returns a boolean if a field has been set.

### GetAccessOrganisationsConsumptionStatistics

`func (o *OAuthRightsView) GetAccessOrganisationsConsumptionStatistics() bool`

GetAccessOrganisationsConsumptionStatistics returns the AccessOrganisationsConsumptionStatistics field if non-nil, zero value otherwise.

### GetAccessOrganisationsConsumptionStatisticsOk

`func (o *OAuthRightsView) GetAccessOrganisationsConsumptionStatisticsOk() (*bool, bool)`

GetAccessOrganisationsConsumptionStatisticsOk returns a tuple with the AccessOrganisationsConsumptionStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOrganisationsConsumptionStatistics

`func (o *OAuthRightsView) SetAccessOrganisationsConsumptionStatistics(v bool)`

SetAccessOrganisationsConsumptionStatistics sets AccessOrganisationsConsumptionStatistics field to given value.

### HasAccessOrganisationsConsumptionStatistics

`func (o *OAuthRightsView) HasAccessOrganisationsConsumptionStatistics() bool`

HasAccessOrganisationsConsumptionStatistics returns a boolean if a field has been set.

### GetAccessPersonalInformation

`func (o *OAuthRightsView) GetAccessPersonalInformation() bool`

GetAccessPersonalInformation returns the AccessPersonalInformation field if non-nil, zero value otherwise.

### GetAccessPersonalInformationOk

`func (o *OAuthRightsView) GetAccessPersonalInformationOk() (*bool, bool)`

GetAccessPersonalInformationOk returns a tuple with the AccessPersonalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPersonalInformation

`func (o *OAuthRightsView) SetAccessPersonalInformation(v bool)`

SetAccessPersonalInformation sets AccessPersonalInformation field to given value.

### HasAccessPersonalInformation

`func (o *OAuthRightsView) HasAccessPersonalInformation() bool`

HasAccessPersonalInformation returns a boolean if a field has been set.

### GetManagePersonalInformation

`func (o *OAuthRightsView) GetManagePersonalInformation() bool`

GetManagePersonalInformation returns the ManagePersonalInformation field if non-nil, zero value otherwise.

### GetManagePersonalInformationOk

`func (o *OAuthRightsView) GetManagePersonalInformationOk() (*bool, bool)`

GetManagePersonalInformationOk returns a tuple with the ManagePersonalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagePersonalInformation

`func (o *OAuthRightsView) SetManagePersonalInformation(v bool)`

SetManagePersonalInformation sets ManagePersonalInformation field to given value.

### HasManagePersonalInformation

`func (o *OAuthRightsView) HasManagePersonalInformation() bool`

HasManagePersonalInformation returns a boolean if a field has been set.

### GetManageSshKeys

`func (o *OAuthRightsView) GetManageSshKeys() bool`

GetManageSshKeys returns the ManageSshKeys field if non-nil, zero value otherwise.

### GetManageSshKeysOk

`func (o *OAuthRightsView) GetManageSshKeysOk() (*bool, bool)`

GetManageSshKeysOk returns a tuple with the ManageSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageSshKeys

`func (o *OAuthRightsView) SetManageSshKeys(v bool)`

SetManageSshKeys sets ManageSshKeys field to given value.

### HasManageSshKeys

`func (o *OAuthRightsView) HasManageSshKeys() bool`

HasManageSshKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


