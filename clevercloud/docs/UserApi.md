# \UserApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AskForPasswordResetViaForm**](UserApi.md#AskForPasswordResetViaForm) | **Post** /password_forgotten | 
[**AuthorizePaypalTransaction**](UserApi.md#AuthorizePaypalTransaction) | **Put** /invoice/external/paypal/{bid} | 
[**CancelPaypalTransaction**](UserApi.md#CancelPaypalTransaction) | **Delete** /invoice/external/paypal/{bid} | 
[**ConfirmPasswordResetRequest**](UserApi.md#ConfirmPasswordResetRequest) | **Get** /password_forgotten/{key} | 
[**CreateUserFromForm**](UserApi.md#CreateUserFromForm) | **Post** /users | 
[**DeleteGithubLink**](UserApi.md#DeleteGithubLink) | **Delete** /github/link | 
[**FinsihGithubSignup**](UserApi.md#FinsihGithubSignup) | **Post** /github/signup | 
[**GetApplications**](UserApi.md#GetApplications) | **Get** /users/{id}/applications | 
[**GetEnv**](UserApi.md#GetEnv) | **Get** /application/{appId}/environment | 
[**GetGitInfo**](UserApi.md#GetGitInfo) | **Get** /users/{userId}/git-info | 
[**GetGithub**](UserApi.md#GetGithub) | **Get** /github | 
[**GetGithubApplications**](UserApi.md#GetGithubApplications) | **Get** /github/applications | 
[**GetGithubCallback**](UserApi.md#GetGithubCallback) | **Get** /github/callback | 
[**GetGithubEmails**](UserApi.md#GetGithubEmails) | **Get** /github/emails | 
[**GetGithubKeys**](UserApi.md#GetGithubKeys) | **Get** /github/keys | 
[**GetGithubLink**](UserApi.md#GetGithubLink) | **Get** /github/link | 
[**GetGithubLogin**](UserApi.md#GetGithubLogin) | **Get** /github/login | 
[**GetGithubUsername**](UserApi.md#GetGithubUsername) | **Get** /github/username | 
[**GetLoginForm**](UserApi.md#GetLoginForm) | **Get** /session/login | 
[**GetLoginForm1**](UserApi.md#GetLoginForm1) | **Get** /sessions/login | 
[**GetPasswordForgottenForm**](UserApi.md#GetPasswordForgottenForm) | **Get** /password_forgotten | 
[**GetSignupForm**](UserApi.md#GetSignupForm) | **Get** /session/signup | 
[**GetSignupForm1**](UserApi.md#GetSignupForm1) | **Get** /sessions/signup | 
[**GetUserById**](UserApi.md#GetUserById) | **Get** /users/{id} | 
[**GithubSignup**](UserApi.md#GithubSignup) | **Get** /github/signup | 
[**Login**](UserApi.md#Login) | **Post** /session/login | 
[**Login1**](UserApi.md#Login1) | **Post** /sessions/login | 
[**MfaLogin**](UserApi.md#MfaLogin) | **Post** /session/mfa_login | 
[**MfaLogin1**](UserApi.md#MfaLogin1) | **Post** /sessions/mfa_login | 
[**PostGithubRedeploy**](UserApi.md#PostGithubRedeploy) | **Post** /github/redeploy | 
[**ResetPasswordForgotten**](UserApi.md#ResetPasswordForgotten) | **Post** /password_forgotten/{key} | 
[**UpdateEnv**](UserApi.md#UpdateEnv) | **Put** /application/{appId}/environment | 
[**UpdateInvoice**](UserApi.md#UpdateInvoice) | **Post** /invoice/external/{bid} | 



## AskForPasswordResetViaForm

> string AskForPasswordResetViaForm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AskForPasswordResetViaFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AskForPasswordResetViaFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | **optional.String**|  | 
 **dropTokens** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizePaypalTransaction

> AuthorizePaypalTransaction(ctx, bid, paymentData)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 
**paymentData** | [**PaymentData**](PaymentData.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPaypalTransaction

> CancelPaypalTransaction(ctx, bid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPasswordResetRequest

> string ConfirmPasswordResetRequest(ctx, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 
 **optional** | ***ConfirmPasswordResetRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfirmPasswordResetRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserFromForm

> CreateUserFromForm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateUserFromFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserFromFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationKey** | **optional.String**|  | 
 **addonBetaInvitationKey** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **pass** | **optional.String**|  | 
 **urlNext** | **optional.String**|  | 
 **terms** | **optional.String**|  | 
 **subscriptionSource** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGithubLink

> Message DeleteGithubLink(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinsihGithubSignup

> string FinsihGithubSignup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FinsihGithubSignupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FinsihGithubSignupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **otherId** | **optional.String**|  | 
 **otherEmail** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **autoLink** | **optional.String**|  | 
 **terms** | **optional.String**|  | 
 **invitationKey** | **optional.String**|  | 
 **mfaKind** | **optional.String**|  | 
 **mfaAttempt** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> []ApplicationView GetApplications(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnv

> string GetEnv(ctx, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
 **optional** | ***GetEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitInfo

> string GetGitInfo(ctx, userId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithub

> OAuthTransactionView GetGithub(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**OAuthTransactionView**](OAuthTransactionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubApplications

> []OAuthApplicationView GetGithubApplications(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]OAuthApplicationView**](OAuthApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubCallback

> GetGithubCallback(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGithubCallbackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGithubCallbackOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccOAuthData** | **optional.String**|  | 
 **code** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **error_** | **optional.String**|  | 
 **errorDescription** | **optional.String**|  | 
 **errorUri** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubEmails

> []string GetGithubEmails(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubKeys

> []SshKeyView GetGithubKeys(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SshKeyView**](SshKeyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubLink

> string GetGithubLink(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGithubLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGithubLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**|  | 
 **redirectUrl** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubLogin

> GetGithubLogin(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGithubLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGithubLoginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUrl** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **invitationKey** | **optional.String**|  | 
 **subscriptionSource** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubUsername

> string GetGithubUsername(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginForm

> string GetLoginForm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLoginFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoginFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secondaryEmailKey** | **optional.String**|  | 
 **deletionKey** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginForm1

> string GetLoginForm1(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLoginForm1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoginForm1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secondaryEmailKey** | **optional.String**|  | 
 **deletionKey** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordForgottenForm

> string GetPasswordForgottenForm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPasswordForgottenFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPasswordForgottenFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignupForm

> string GetSignupForm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSignupFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSignupFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationKey** | **optional.String**|  | 
 **urlNext** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignupForm1

> string GetSignupForm1(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSignupForm1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSignupForm1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationKey** | **optional.String**|  | 
 **urlNext** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> UserView GetUserById(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**UserView**](UserView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GithubSignup

> GithubSignup(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GithubSignupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GithubSignupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUrl** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **invitationKey** | **optional.String**|  | 
 **subscriptionSource** | **optional.String**|  | 
 **terms** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> string Login(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 
 **pass** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login1

> string Login1(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Login1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Login1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 
 **pass** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MfaLogin

> []OAuthApplicationView MfaLogin(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MfaLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MfaLoginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaKind** | **optional.String**|  | 
 **mfaAttempt** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **authId** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 

### Return type

[**[]OAuthApplicationView**](OAuthApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MfaLogin1

> []OAuthApplicationView MfaLogin1(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MfaLogin1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MfaLogin1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaKind** | **optional.String**|  | 
 **mfaAttempt** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **authId** | **optional.String**|  | 
 **fromAuthorize** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **cleverFlavor** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 

### Return type

[**[]OAuthApplicationView**](OAuthApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGithubRedeploy

> Message PostGithubRedeploy(ctx, githubWebhookPayload, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubWebhookPayload** | [**GithubWebhookPayload**](GithubWebhookPayload.md)|  | 
 **optional** | ***PostGithubRedeployOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostGithubRedeployOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAgent** | **optional.String**|  | 
 **xGithubEvent** | **optional.String**|  | 
 **xHubSignature** | **optional.String**|  | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordForgotten

> string ResetPasswordForgotten(ctx, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 
 **optional** | ***ResetPasswordForgottenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResetPasswordForgottenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pass** | **optional.String**|  | 
 **pass2** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnv

> Message UpdateEnv(ctx, appId, body, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**body** | **string**|  | 
 **optional** | ***UpdateEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **optional.String**|  | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> UpdateInvoice(ctx, bid, endOfInvoiceResponse)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 
**endOfInvoiceResponse** | [**EndOfInvoiceResponse**](EndOfInvoiceResponse.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

