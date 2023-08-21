# \GuidanceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGuidanceSettings**](GuidanceSettingsApi.md#GetGuidanceSettings) | **Get** /api/guidance-settings | Get Guidance Settings
[**UpdateGuidanceSettings**](GuidanceSettingsApi.md#UpdateGuidanceSettings) | **Put** /api/guidance-settings | Update Guidance Settings



## GetGuidanceSettings

> InlineResponse20047 GetGuidanceSettings(ctx).Execute()

Get Guidance Settings



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
    resp, r, err := api_client.GuidanceSettingsApi.GetGuidanceSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceSettingsApi.GetGuidanceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGuidanceSettings`: InlineResponse20047
    fmt.Fprintf(os.Stdout, "Response from `GuidanceSettingsApi.GetGuidanceSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuidanceSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuidanceSettings

> Model200Success UpdateGuidanceSettings(ctx).InlineObject79(inlineObject79).Execute()

Update Guidance Settings



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
    inlineObject79 := *openapiclient.NewInlineObject79() // InlineObject79 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GuidanceSettingsApi.UpdateGuidanceSettings(context.Background()).InlineObject79(inlineObject79).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceSettingsApi.UpdateGuidanceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGuidanceSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `GuidanceSettingsApi.UpdateGuidanceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuidanceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject79** | [**InlineObject79**](InlineObject79.md) |  | 

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

