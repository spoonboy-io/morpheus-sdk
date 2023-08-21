# \AuthenticationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForgotPassword**](AuthenticationApi.md#ForgotPassword) | **Post** /api/forgot/send-email | Request a reset password email
[**GetAccessToken**](AuthenticationApi.md#GetAccessToken) | **Post** /oauth/token | Provides authentication via username and password
[**ResetPassword**](AuthenticationApi.md#ResetPassword) | **Post** /api/forgot/reset-password | Reset user password
[**Whoami**](AuthenticationApi.md#Whoami) | **Get** /api/whoami | Retrieves information about current user roles and permissions



## ForgotPassword

> InlineResponse2007 ForgotPassword(ctx).InlineObject11(inlineObject11).Execute()

Request a reset password email



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject11 := *openapiclient.NewInlineObject11("Username_example") // InlineObject11 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.ForgotPassword(context.Background()).InlineObject11(inlineObject11).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.ForgotPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForgotPassword`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.ForgotPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject11** | [**InlineObject11**](InlineObject11.md) |  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessToken

> AccessToken GetAccessToken(ctx).ClientId(clientId).GrantType(grantType).Scope(scope).Username(username).Password(password).RefreshToken(refreshToken).Execute()

Provides authentication via username and password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli.
    grantType := "grantType_example" // string | OAuth Grant Type, use password.
    scope := "scope_example" // string | OAuth token scope, use write.
    username := "username_example" // string | Specified as \\\"username\\\" or \\\"tenantId\\\\username\\\" with proper HTML encoding and used in conjunction with `password`. Not utilized with `refresh_token`. (optional)
    password := "password_example" // string | The Password for defined `username`. Must have proper HTML encoding (optional)
    refreshToken := TODO // interface{} | The `refresh_token` value from a previous API generation can be utilized instead of `username` and `password`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.GetAccessToken(context.Background()).ClientId(clientId).GrantType(grantType).Scope(scope).Username(username).Password(password).RefreshToken(refreshToken).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.GetAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessToken`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli. | 
 **grantType** | **string** | OAuth Grant Type, use password. | 
 **scope** | **string** | OAuth token scope, use write. | 
 **username** | **string** | Specified as \\\&quot;username\\\&quot; or \\\&quot;tenantId\\\\username\\\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;. Not utilized with &#x60;refresh_token&#x60;. | 
 **password** | **string** | The Password for defined &#x60;username&#x60;. Must have proper HTML encoding | 
 **refreshToken** | [**interface{}**](interface{}.md) | The &#x60;refresh_token&#x60; value from a previous API generation can be utilized instead of &#x60;username&#x60; and &#x60;password&#x60;. | 

### Return type

[**AccessToken**](accessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> InlineResponse2006 ResetPassword(ctx).InlineObject10(inlineObject10).Execute()

Reset user password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject10 := *openapiclient.NewInlineObject10("Token_example", "Password_example") // InlineObject10 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.ResetPassword(context.Background()).InlineObject10(inlineObject10).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPassword`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.ResetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject10** | [**InlineObject10**](InlineObject10.md) |  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Whoami

> InlineResponse200167 Whoami(ctx).Execute()

Retrieves information about current user roles and permissions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.Whoami(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Whoami``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Whoami`: InlineResponse200167
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.Whoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiRequest struct via the builder pattern


### Return type

[**InlineResponse200167**](inline_response_200_167.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

