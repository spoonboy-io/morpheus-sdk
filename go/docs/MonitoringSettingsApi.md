# \MonitoringSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitoringSettings**](MonitoringSettingsApi.md#GetMonitoringSettings) | **Get** /api/monitoring-settings | Get Monitoring Settings
[**UpdateMonitoringSettings**](MonitoringSettingsApi.md#UpdateMonitoringSettings) | **Put** /api/monitoring-settings | Update Monitoring Settings



## GetMonitoringSettings

> InlineResponse20085 GetMonitoringSettings(ctx).Execute()

Get Monitoring Settings



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
    resp, r, err := api_client.MonitoringSettingsApi.GetMonitoringSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSettingsApi.GetMonitoringSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringSettings`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSettingsApi.GetMonitoringSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitoringSettings

> Model200Success UpdateMonitoringSettings(ctx).InlineObject141(inlineObject141).Execute()

Update Monitoring Settings



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
    inlineObject141 := *openapiclient.NewInlineObject141() // InlineObject141 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitoringSettingsApi.UpdateMonitoringSettings(context.Background()).InlineObject141(inlineObject141).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringSettingsApi.UpdateMonitoringSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitoringSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `MonitoringSettingsApi.UpdateMonitoringSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject141** | [**InlineObject141**](InlineObject141.md) |  | 

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

