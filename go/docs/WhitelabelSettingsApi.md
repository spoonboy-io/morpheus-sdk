# \WhitelabelSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWhitelabelImage**](WhitelabelSettingsApi.md#GetWhitelabelImage) | **Get** /api/whitelabel-settings/images/{imageType} | Download Image
[**ListWhitelabelSettings**](WhitelabelSettingsApi.md#ListWhitelabelSettings) | **Get** /api/whitelabel-settings | Get Whitelabel Settings
[**RemoveWhitelabelImage**](WhitelabelSettingsApi.md#RemoveWhitelabelImage) | **Delete** /api/whitelabel-settings/images/{imageType} | Reset Image
[**UpdateWhitelabelImages**](WhitelabelSettingsApi.md#UpdateWhitelabelImages) | **Post** /api/whitelabel-settings/images | Update Images
[**UpdateWhitelabelSettings**](WhitelabelSettingsApi.md#UpdateWhitelabelSettings) | **Put** /api/whitelabel-settings | Update Whitelabel Settings



## GetWhitelabelImage

> *os.File GetWhitelabelImage(ctx, imageType).Execute()

Download Image



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
    imageType := "headerLogo" // string | Valid image types

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WhitelabelSettingsApi.GetWhitelabelImage(context.Background(), imageType).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelabelSettingsApi.GetWhitelabelImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWhitelabelImage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `WhitelabelSettingsApi.GetWhitelabelImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageType** | **string** | Valid image types | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhitelabelImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/ico, image/jpeg, image/png, image/svg+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhitelabelSettings

> InlineResponse200166 ListWhitelabelSettings(ctx).Execute()

Get Whitelabel Settings



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
    resp, r, err := api_client.WhitelabelSettingsApi.ListWhitelabelSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelabelSettingsApi.ListWhitelabelSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWhitelabelSettings`: InlineResponse200166
    fmt.Fprintf(os.Stdout, "Response from `WhitelabelSettingsApi.ListWhitelabelSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWhitelabelSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse200166**](inline_response_200_166.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWhitelabelImage

> Model200Success RemoveWhitelabelImage(ctx, imageType).Execute()

Reset Image



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
    imageType := "headerLogo" // string | Valid image types

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WhitelabelSettingsApi.RemoveWhitelabelImage(context.Background(), imageType).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelabelSettingsApi.RemoveWhitelabelImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveWhitelabelImage`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `WhitelabelSettingsApi.RemoveWhitelabelImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageType** | **string** | Valid image types | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWhitelabelImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhitelabelImages

> Model200Success UpdateWhitelabelImages(ctx).HeaderLogoFile(headerLogoFile).ResetHeaderLogo(resetHeaderLogo).FooterLogoFile(footerLogoFile).ResetFooterLogo(resetFooterLogo).LoginLogoFile(loginLogoFile).ResetLoginLogo(resetLoginLogo).FaviconFile(faviconFile).ResetFaviconLogo(resetFaviconLogo).Execute()

Update Images



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
    headerLogoFile := os.NewFile(1234, "some_file") // *os.File | Header logo image file, valid image types `png|jpg|svg` (optional)
    resetHeaderLogo := true // bool | Resets header logo to default (optional)
    footerLogoFile := os.NewFile(1234, "some_file") // *os.File | Footer logo image file, valid image types `png|jpg|svg` (optional)
    resetFooterLogo := true // bool | Resets footer logo to default (optional)
    loginLogoFile := os.NewFile(1234, "some_file") // *os.File | Login logo image file, valid image types `png|jpg|svg` (optional)
    resetLoginLogo := true // bool | Resets login logo to default (optional)
    faviconFile := os.NewFile(1234, "some_file") // *os.File | Favicon image file, valid image type ico (optional)
    resetFaviconLogo := true // bool | Resets favicon logo to default (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WhitelabelSettingsApi.UpdateWhitelabelImages(context.Background()).HeaderLogoFile(headerLogoFile).ResetHeaderLogo(resetHeaderLogo).FooterLogoFile(footerLogoFile).ResetFooterLogo(resetFooterLogo).LoginLogoFile(loginLogoFile).ResetLoginLogo(resetLoginLogo).FaviconFile(faviconFile).ResetFaviconLogo(resetFaviconLogo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelabelSettingsApi.UpdateWhitelabelImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWhitelabelImages`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `WhitelabelSettingsApi.UpdateWhitelabelImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhitelabelImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **headerLogoFile** | ***os.File** | Header logo image file, valid image types &#x60;png|jpg|svg&#x60; | 
 **resetHeaderLogo** | **bool** | Resets header logo to default | 
 **footerLogoFile** | ***os.File** | Footer logo image file, valid image types &#x60;png|jpg|svg&#x60; | 
 **resetFooterLogo** | **bool** | Resets footer logo to default | 
 **loginLogoFile** | ***os.File** | Login logo image file, valid image types &#x60;png|jpg|svg&#x60; | 
 **resetLoginLogo** | **bool** | Resets login logo to default | 
 **faviconFile** | ***os.File** | Favicon image file, valid image type ico | 
 **resetFaviconLogo** | **bool** | Resets favicon logo to default | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhitelabelSettings

> Model200Success UpdateWhitelabelSettings(ctx).InlineObject265(inlineObject265).Execute()

Update Whitelabel Settings



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
    inlineObject265 := *openapiclient.NewInlineObject265() // InlineObject265 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WhitelabelSettingsApi.UpdateWhitelabelSettings(context.Background()).InlineObject265(inlineObject265).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelabelSettingsApi.UpdateWhitelabelSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWhitelabelSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `WhitelabelSettingsApi.UpdateWhitelabelSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhitelabelSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject265** | [**InlineObject265**](InlineObject265.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

