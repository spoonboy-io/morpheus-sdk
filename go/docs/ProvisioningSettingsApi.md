# \ProvisioningSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProvisioningSettings**](ProvisioningSettingsApi.md#ListProvisioningSettings) | **Get** /api/provisioning-settings | Get Provisioning Settings
[**UpdateProvisioningSettings**](ProvisioningSettingsApi.md#UpdateProvisioningSettings) | **Put** /api/provisioning-settings | Update Provisioning Settings



## ListProvisioningSettings

> InlineResponse200128 ListProvisioningSettings(ctx).Execute()

Get Provisioning Settings



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
    resp, r, err := api_client.ProvisioningSettingsApi.ListProvisioningSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningSettingsApi.ListProvisioningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProvisioningSettings`: InlineResponse200128
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningSettingsApi.ListProvisioningSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProvisioningSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse200128**](inline_response_200_128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvisioningSettings

> Model200Success UpdateProvisioningSettings(ctx).InlineObject204(inlineObject204).Execute()

Update Provisioning Settings



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
    inlineObject204 := *openapiclient.NewInlineObject204() // InlineObject204 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvisioningSettingsApi.UpdateProvisioningSettings(context.Background()).InlineObject204(inlineObject204).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningSettingsApi.UpdateProvisioningSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProvisioningSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningSettingsApi.UpdateProvisioningSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProvisioningSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject204** | [**InlineObject204**](InlineObject204.md) |  | 

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

