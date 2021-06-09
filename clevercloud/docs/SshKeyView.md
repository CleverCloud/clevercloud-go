# SshKeyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 

## Methods

### NewSshKeyView

`func NewSshKeyView() *SshKeyView`

NewSshKeyView instantiates a new SshKeyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyViewWithDefaults

`func NewSshKeyViewWithDefaults() *SshKeyView`

NewSshKeyViewWithDefaults instantiates a new SshKeyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshKeyView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeyView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeyView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SshKeyView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *SshKeyView) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeyView) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeyView) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKeyView) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetFingerprint

`func (o *SshKeyView) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SshKeyView) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SshKeyView) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SshKeyView) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


