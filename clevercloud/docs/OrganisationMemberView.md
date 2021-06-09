# OrganisationMemberView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to [**OrganisationMemberUserView**](OrganisationMemberUserView.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Job** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganisationMemberView

`func NewOrganisationMemberView() *OrganisationMemberView`

NewOrganisationMemberView instantiates a new OrganisationMemberView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationMemberViewWithDefaults

`func NewOrganisationMemberViewWithDefaults() *OrganisationMemberView`

NewOrganisationMemberViewWithDefaults instantiates a new OrganisationMemberView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *OrganisationMemberView) GetMember() OrganisationMemberUserView`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OrganisationMemberView) GetMemberOk() (*OrganisationMemberUserView, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OrganisationMemberView) SetMember(v OrganisationMemberUserView)`

SetMember sets Member field to given value.

### HasMember

`func (o *OrganisationMemberView) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetRole

`func (o *OrganisationMemberView) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganisationMemberView) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganisationMemberView) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganisationMemberView) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetJob

`func (o *OrganisationMemberView) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *OrganisationMemberView) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *OrganisationMemberView) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *OrganisationMemberView) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


