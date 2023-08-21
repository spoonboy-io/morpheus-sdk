# \ApplianceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApplianceSettings**](ApplianceSettingsApi.md#ListApplianceSettings) | **Get** /api/appliance-settings | Get Appliance Settings
[**Reindex**](ApplianceSettingsApi.md#Reindex) | **Post** /api/appliance-settings/reindex | Reindex Search
[**SetApplianceSettingsMaintenanceMode**](ApplianceSettingsApi.md#SetApplianceSettingsMaintenanceMode) | **Post** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**UpdateApplianceSettings**](ApplianceSettingsApi.md#UpdateApplianceSettings) | **Put** /api/appliance-settings | Update Appliance Settings



## ListApplianceSettings

> InlineResponse200 ListApplianceSettings(ctx).Execute()

Get Appliance Settings



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
    resp, r, err := api_client.ApplianceSettingsApi.ListApplianceSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsApi.ListApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplianceSettings`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsApi.ListApplianceSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplianceSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reindex

> Model200Success Reindex(ctx).Execute()

Reindex Search



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
    resp, r, err := api_client.ApplianceSettingsApi.Reindex(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsApi.Reindex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Reindex`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsApi.Reindex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReindexRequest struct via the builder pattern


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


## SetApplianceSettingsMaintenanceMode

> Model200Success SetApplianceSettingsMaintenanceMode(ctx).Enabled(enabled).Execute()

Toggle Maintenance Mode



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
    enabled := true // bool | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplianceSettingsApi.SetApplianceSettingsMaintenanceMode(context.Background()).Enabled(enabled).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsApi.SetApplianceSettingsMaintenanceMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetApplianceSettingsMaintenanceMode`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsApi.SetApplianceSettingsMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetApplianceSettingsMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | 

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


## UpdateApplianceSettings

> Model200Success UpdateApplianceSettings(ctx).InlineObject2(inlineObject2).Execute()

Update Appliance Settings



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
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplianceSettingsApi.UpdateApplianceSettings(context.Background()).InlineObject2(inlineObject2).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsApi.UpdateApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplianceSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsApi.UpdateApplianceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

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

