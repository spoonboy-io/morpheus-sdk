# \LogSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogSettingsSyslogRules**](LogSettingsApi.md#AddLogSettingsSyslogRules) | **Post** /api/log-settings/syslog-rules | Create a New Syslog Rule
[**DeleteLogSettingsSyslogRules**](LogSettingsApi.md#DeleteLogSettingsSyslogRules) | **Delete** /api/log-settings/syslog-rules/{id} | Delete a Specific Syslog Rule
[**ListLogSettings**](LogSettingsApi.md#ListLogSettings) | **Get** /api/log-settings | List All Log Settings
[**UpdateLogSettings**](LogSettingsApi.md#UpdateLogSettings) | **Put** /api/log-settings | Update Log Settings



## AddLogSettingsSyslogRules

> Model200Success AddLogSettingsSyslogRules(ctx).InlineObject140(inlineObject140).Execute()

Create a New Syslog Rule



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
    inlineObject140 := *openapiclient.NewInlineObject140(*openapiclient.NewApiLogSettingsSyslogRulesSyslogRule("foo", "*.* @@bar.com:8080")) // InlineObject140 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogSettingsApi.AddLogSettingsSyslogRules(context.Background()).InlineObject140(inlineObject140).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsApi.AddLogSettingsSyslogRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogSettingsSyslogRules`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LogSettingsApi.AddLogSettingsSyslogRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogSettingsSyslogRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject140** | [**InlineObject140**](InlineObject140.md) |  | 

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


## DeleteLogSettingsSyslogRules

> Model200Success DeleteLogSettingsSyslogRules(ctx, id).Execute()

Delete a Specific Syslog Rule



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogSettingsApi.DeleteLogSettingsSyslogRules(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsApi.DeleteLogSettingsSyslogRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogSettingsSyslogRules`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LogSettingsApi.DeleteLogSettingsSyslogRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogSettingsSyslogRulesRequest struct via the builder pattern


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


## ListLogSettings

> InlineResponse20084 ListLogSettings(ctx).Execute()

List All Log Settings



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
    resp, r, err := api_client.LogSettingsApi.ListLogSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsApi.ListLogSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogSettings`: InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `LogSettingsApi.ListLogSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse20084**](inline_response_200_84.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogSettings

> Model200Success UpdateLogSettings(ctx).InlineObject139(inlineObject139).Execute()

Update Log Settings



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
    inlineObject139 := *openapiclient.NewInlineObject139() // InlineObject139 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogSettingsApi.UpdateLogSettings(context.Background()).InlineObject139(inlineObject139).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsApi.UpdateLogSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LogSettingsApi.UpdateLogSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject139** | [**InlineObject139**](InlineObject139.md) |  | 

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

