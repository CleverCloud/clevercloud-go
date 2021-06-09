# SuperNovaInstanceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**AppPort** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**SuperNovaFlavor**](SuperNovaFlavor.md) |  | [optional] 
**Commit** | Pointer to **string** |  | [optional] 
**DeployNumber** | Pointer to **int32** |  | [optional] 
**DeployId** | Pointer to **string** |  | [optional] 
**InstanceNumber** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewSuperNovaInstanceView

`func NewSuperNovaInstanceView() *SuperNovaInstanceView`

NewSuperNovaInstanceView instantiates a new SuperNovaInstanceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperNovaInstanceViewWithDefaults

`func NewSuperNovaInstanceViewWithDefaults() *SuperNovaInstanceView`

NewSuperNovaInstanceViewWithDefaults instantiates a new SuperNovaInstanceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuperNovaInstanceView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuperNovaInstanceView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuperNovaInstanceView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuperNovaInstanceView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *SuperNovaInstanceView) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SuperNovaInstanceView) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SuperNovaInstanceView) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *SuperNovaInstanceView) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetIp

`func (o *SuperNovaInstanceView) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SuperNovaInstanceView) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SuperNovaInstanceView) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SuperNovaInstanceView) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAppPort

`func (o *SuperNovaInstanceView) GetAppPort() int32`

GetAppPort returns the AppPort field if non-nil, zero value otherwise.

### GetAppPortOk

`func (o *SuperNovaInstanceView) GetAppPortOk() (*int32, bool)`

GetAppPortOk returns a tuple with the AppPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPort

`func (o *SuperNovaInstanceView) SetAppPort(v int32)`

SetAppPort sets AppPort field to given value.

### HasAppPort

`func (o *SuperNovaInstanceView) HasAppPort() bool`

HasAppPort returns a boolean if a field has been set.

### GetState

`func (o *SuperNovaInstanceView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SuperNovaInstanceView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SuperNovaInstanceView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SuperNovaInstanceView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFlavor

`func (o *SuperNovaInstanceView) GetFlavor() SuperNovaFlavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *SuperNovaInstanceView) GetFlavorOk() (*SuperNovaFlavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *SuperNovaInstanceView) SetFlavor(v SuperNovaFlavor)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *SuperNovaInstanceView) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetCommit

`func (o *SuperNovaInstanceView) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *SuperNovaInstanceView) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *SuperNovaInstanceView) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *SuperNovaInstanceView) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetDeployNumber

`func (o *SuperNovaInstanceView) GetDeployNumber() int32`

GetDeployNumber returns the DeployNumber field if non-nil, zero value otherwise.

### GetDeployNumberOk

`func (o *SuperNovaInstanceView) GetDeployNumberOk() (*int32, bool)`

GetDeployNumberOk returns a tuple with the DeployNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployNumber

`func (o *SuperNovaInstanceView) SetDeployNumber(v int32)`

SetDeployNumber sets DeployNumber field to given value.

### HasDeployNumber

`func (o *SuperNovaInstanceView) HasDeployNumber() bool`

HasDeployNumber returns a boolean if a field has been set.

### GetDeployId

`func (o *SuperNovaInstanceView) GetDeployId() string`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *SuperNovaInstanceView) GetDeployIdOk() (*string, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *SuperNovaInstanceView) SetDeployId(v string)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *SuperNovaInstanceView) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.

### GetInstanceNumber

`func (o *SuperNovaInstanceView) GetInstanceNumber() int32`

GetInstanceNumber returns the InstanceNumber field if non-nil, zero value otherwise.

### GetInstanceNumberOk

`func (o *SuperNovaInstanceView) GetInstanceNumberOk() (*int32, bool)`

GetInstanceNumberOk returns a tuple with the InstanceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNumber

`func (o *SuperNovaInstanceView) SetInstanceNumber(v int32)`

SetInstanceNumber sets InstanceNumber field to given value.

### HasInstanceNumber

`func (o *SuperNovaInstanceView) HasInstanceNumber() bool`

HasInstanceNumber returns a boolean if a field has been set.

### GetDisplayName

`func (o *SuperNovaInstanceView) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SuperNovaInstanceView) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SuperNovaInstanceView) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SuperNovaInstanceView) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCreationDate

`func (o *SuperNovaInstanceView) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SuperNovaInstanceView) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SuperNovaInstanceView) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SuperNovaInstanceView) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


