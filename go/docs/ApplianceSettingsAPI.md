# \ApplianceSettingsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApplianceSettings**](ApplianceSettingsAPI.md#ListApplianceSettings) | **Get** /api/appliance-settings | Get Appliance Settings
[**Reindex**](ApplianceSettingsAPI.md#Reindex) | **Post** /api/appliance-settings/reindex | Reindex Search
[**SetApplianceSettingsMaintenanceMode**](ApplianceSettingsAPI.md#SetApplianceSettingsMaintenanceMode) | **Post** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**UpdateApplianceSettings**](ApplianceSettingsAPI.md#UpdateApplianceSettings) | **Put** /api/appliance-settings | Update Appliance Settings



## ListApplianceSettings

> ListApplianceSettings200Response ListApplianceSettings(ctx).Execute()

Get Appliance Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceSettingsAPI.ListApplianceSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsAPI.ListApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplianceSettings`: ListApplianceSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsAPI.ListApplianceSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplianceSettingsRequest struct via the builder pattern


### Return type

[**ListApplianceSettings200Response**](ListApplianceSettings200Response.md)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceSettingsAPI.Reindex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsAPI.Reindex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Reindex`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsAPI.Reindex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReindexRequest struct via the builder pattern


### Return type

[**Model200Success**](Model200Success.md)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    enabled := true // bool | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceSettingsAPI.SetApplianceSettingsMaintenanceMode(context.Background()).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsAPI.SetApplianceSettingsMaintenanceMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetApplianceSettingsMaintenanceMode`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsAPI.SetApplianceSettingsMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetApplianceSettingsMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplianceSettings

> Model200Success UpdateApplianceSettings(ctx).UpdateApplianceSettingsRequest(updateApplianceSettingsRequest).Execute()

Update Appliance Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateApplianceSettingsRequest := *openapiclient.NewUpdateApplianceSettingsRequest() // UpdateApplianceSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplianceSettingsAPI.UpdateApplianceSettings(context.Background()).UpdateApplianceSettingsRequest(updateApplianceSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplianceSettingsAPI.UpdateApplianceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplianceSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ApplianceSettingsAPI.UpdateApplianceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApplianceSettingsRequest** | [**UpdateApplianceSettingsRequest**](UpdateApplianceSettingsRequest.md) |  | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

