# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**Organisations** | Pointer to [**[]OrganisationSummary**](OrganisationSummary.md) |  | [optional] 

## Methods

### NewSummary

`func NewSummary() *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *Summary) GetUser() UserSummary`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Summary) GetUserOk() (*UserSummary, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Summary) SetUser(v UserSummary)`

SetUser sets User field to given value.

### HasUser

`func (o *Summary) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganisations

`func (o *Summary) GetOrganisations() []OrganisationSummary`

GetOrganisations returns the Organisations field if non-nil, zero value otherwise.

### GetOrganisationsOk

`func (o *Summary) GetOrganisationsOk() (*[]OrganisationSummary, bool)`

GetOrganisationsOk returns a tuple with the Organisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisations

`func (o *Summary) SetOrganisations(v []OrganisationSummary)`

SetOrganisations sets Organisations field to given value.

### HasOrganisations

`func (o *Summary) HasOrganisations() bool`

HasOrganisations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


