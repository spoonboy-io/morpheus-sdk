# \LicenseApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicense**](LicenseApi.md#GetLicense) | **Get** /api/license | Get license
[**InstallLicense**](LicenseApi.md#InstallLicense) | **Post** /api/license | Install license key
[**TestLicense**](LicenseApi.md#TestLicense) | **Post** /api/license/test | Test license key
[**UninstallLicense**](LicenseApi.md#UninstallLicense) | **Delete** /api/license | Uninstall license key



## GetLicense

> License GetLicense(ctx).Execute()

Get license



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
    resp, r, err := api_client.LicenseApi.GetLicense(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


### Return type

[**License**](license.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallLicense

> License InstallLicense(ctx).InlineObject125(inlineObject125).Execute()

Install license key



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
    inlineObject125 := *openapiclient.NewInlineObject125("License_example") // InlineObject125 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.InstallLicense(context.Background()).InlineObject125(inlineObject125).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.InstallLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.InstallLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstallLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject125** | [**InlineObject125**](InlineObject125.md) |  | 

### Return type

[**License**](license.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestLicense

> License TestLicense(ctx).InlineObject126(inlineObject126).Execute()

Test license key



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
    inlineObject126 := *openapiclient.NewInlineObject126("License_example") // InlineObject126 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.TestLicense(context.Background()).InlineObject126(inlineObject126).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.TestLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.TestLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject126** | [**InlineObject126**](InlineObject126.md) |  | 

### Return type

[**License**](license.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallLicense

> Model200Success UninstallLicense(ctx).Execute()

Uninstall license key



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
    resp, r, err := api_client.LicenseApi.UninstallLicense(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.UninstallLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UninstallLicense`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.UninstallLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallLicenseRequest struct via the builder pattern


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

