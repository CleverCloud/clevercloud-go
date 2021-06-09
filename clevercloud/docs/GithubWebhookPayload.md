# GithubWebhookPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** |  | [optional] 
**After** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**GithubWebhookRepository**](GithubWebhookRepository.md) |  | [optional] 
**Sender** | Pointer to [**GithubWebhookSender**](GithubWebhookSender.md) |  | [optional] 
**Pusher** | Pointer to [**GithubWebhookPusher**](GithubWebhookPusher.md) |  | [optional] 
**HeadCommit** | Pointer to [**GithubCommit**](GithubCommit.md) |  | [optional] 

## Methods

### NewGithubWebhookPayload

`func NewGithubWebhookPayload() *GithubWebhookPayload`

NewGithubWebhookPayload instantiates a new GithubWebhookPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookPayloadWithDefaults

`func NewGithubWebhookPayloadWithDefaults() *GithubWebhookPayload`

NewGithubWebhookPayloadWithDefaults instantiates a new GithubWebhookPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GithubWebhookPayload) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GithubWebhookPayload) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GithubWebhookPayload) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GithubWebhookPayload) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAfter

`func (o *GithubWebhookPayload) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *GithubWebhookPayload) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *GithubWebhookPayload) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *GithubWebhookPayload) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetRepository

`func (o *GithubWebhookPayload) GetRepository() GithubWebhookRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GithubWebhookPayload) GetRepositoryOk() (*GithubWebhookRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GithubWebhookPayload) SetRepository(v GithubWebhookRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GithubWebhookPayload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSender

`func (o *GithubWebhookPayload) GetSender() GithubWebhookSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GithubWebhookPayload) GetSenderOk() (*GithubWebhookSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GithubWebhookPayload) SetSender(v GithubWebhookSender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *GithubWebhookPayload) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetPusher

`func (o *GithubWebhookPayload) GetPusher() GithubWebhookPusher`

GetPusher returns the Pusher field if non-nil, zero value otherwise.

### GetPusherOk

`func (o *GithubWebhookPayload) GetPusherOk() (*GithubWebhookPusher, bool)`

GetPusherOk returns a tuple with the Pusher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPusher

`func (o *GithubWebhookPayload) SetPusher(v GithubWebhookPusher)`

SetPusher sets Pusher field to given value.

### HasPusher

`func (o *GithubWebhookPayload) HasPusher() bool`

HasPusher returns a boolean if a field has been set.

### GetHeadCommit

`func (o *GithubWebhookPayload) GetHeadCommit() GithubCommit`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *GithubWebhookPayload) GetHeadCommitOk() (*GithubCommit, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *GithubWebhookPayload) SetHeadCommit(v GithubCommit)`

SetHeadCommit sets HeadCommit field to given value.

### HasHeadCommit

`func (o *GithubWebhookPayload) HasHeadCommit() bool`

HasHeadCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


