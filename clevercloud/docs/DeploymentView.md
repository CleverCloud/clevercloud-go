# DeploymentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to **string** |  | [optional] 
**Cause** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to **int32** |  | [optional] 
**Author** | Pointer to [**Author**](Author.md) |  | [optional] 

## Methods

### NewDeploymentView

`func NewDeploymentView() *DeploymentView`

NewDeploymentView instantiates a new DeploymentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentViewWithDefaults

`func NewDeploymentViewWithDefaults() *DeploymentView`

NewDeploymentViewWithDefaults instantiates a new DeploymentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentView) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentView) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentView) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *DeploymentView) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeploymentView) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeploymentView) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeploymentView) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDate

`func (o *DeploymentView) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DeploymentView) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DeploymentView) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *DeploymentView) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetState

`func (o *DeploymentView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAction

`func (o *DeploymentView) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeploymentView) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeploymentView) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeploymentView) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentView) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentView) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentView) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentView) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCause

`func (o *DeploymentView) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *DeploymentView) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *DeploymentView) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *DeploymentView) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInstances

`func (o *DeploymentView) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *DeploymentView) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *DeploymentView) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *DeploymentView) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetAuthor

`func (o *DeploymentView) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DeploymentView) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DeploymentView) SetAuthor(v Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DeploymentView) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


