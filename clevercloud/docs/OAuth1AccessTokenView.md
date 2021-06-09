# OAuth1AccessTokenView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Consumer** | Pointer to [**OAuth1ConsumerView**](OAuth1ConsumerView.md) |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**LastUtilisation** | Pointer to **int64** |  | [optional] 
**Rights** | Pointer to [**OAuthRightsView**](OAuthRightsView.md) |  | [optional] 

## Methods

### NewOAuth1AccessTokenView

`func NewOAuth1AccessTokenView() *OAuth1AccessTokenView`

NewOAuth1AccessTokenView instantiates a new OAuth1AccessTokenView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth1AccessTokenViewWithDefaults

`func NewOAuth1AccessTokenViewWithDefaults() *OAuth1AccessTokenView`

NewOAuth1AccessTokenViewWithDefaults instantiates a new OAuth1AccessTokenView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *OAuth1AccessTokenView) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OAuth1AccessTokenView) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OAuth1AccessTokenView) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OAuth1AccessTokenView) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetConsumer

`func (o *OAuth1AccessTokenView) GetConsumer() OAuth1ConsumerView`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *OAuth1AccessTokenView) GetConsumerOk() (*OAuth1ConsumerView, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *OAuth1AccessTokenView) SetConsumer(v OAuth1ConsumerView)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *OAuth1AccessTokenView) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetCreationDate

`func (o *OAuth1AccessTokenView) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *OAuth1AccessTokenView) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *OAuth1AccessTokenView) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *OAuth1AccessTokenView) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastUtilisation

`func (o *OAuth1AccessTokenView) GetLastUtilisation() int64`

GetLastUtilisation returns the LastUtilisation field if non-nil, zero value otherwise.

### GetLastUtilisationOk

`func (o *OAuth1AccessTokenView) GetLastUtilisationOk() (*int64, bool)`

GetLastUtilisationOk returns a tuple with the LastUtilisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUtilisation

`func (o *OAuth1AccessTokenView) SetLastUtilisation(v int64)`

SetLastUtilisation sets LastUtilisation field to given value.

### HasLastUtilisation

`func (o *OAuth1AccessTokenView) HasLastUtilisation() bool`

HasLastUtilisation returns a boolean if a field has been set.

### GetRights

`func (o *OAuth1AccessTokenView) GetRights() OAuthRightsView`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *OAuth1AccessTokenView) GetRightsOk() (*OAuthRightsView, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *OAuth1AccessTokenView) SetRights(v OAuthRightsView)`

SetRights sets Rights field to given value.

### HasRights

`func (o *OAuth1AccessTokenView) HasRights() bool`

HasRights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


