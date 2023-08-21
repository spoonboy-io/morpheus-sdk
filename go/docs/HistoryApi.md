# \HistoryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistory**](HistoryApi.md#GetHistory) | **Get** /api/processes/{id} | Retrieves a Specific Process
[**GetHistoryEvent**](HistoryApi.md#GetHistoryEvent) | **Get** /api/processes/events/{id} | Retrieves a Specific Process Event
[**ListHistory**](HistoryApi.md#ListHistory) | **Get** /api/processes | Retrieves Process History



## GetHistory

> InlineResponse20048 GetHistory(ctx, id).Execute()

Retrieves a Specific Process



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
    resp, r, err := api_client.HistoryApi.GetHistory(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryEvent

> InlineResponse20049 GetHistoryEvent(ctx, id).Execute()

Retrieves a Specific Process Event



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
    resp, r, err := api_client.HistoryApi.GetHistoryEvent(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistoryEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoryEvent`: InlineResponse20049
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistoryEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistory

> map[string]interface{} ListHistory(ctx).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).ZoneId(zoneId).AppId(appId).Phrase(phrase).Execute()

Retrieves Process History



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
    instanceId := int64(94) // int64 | The Instance ID for Filtering (optional)
    containerId := int64(135) // int64 | The Container ID for Filtering (optional)
    serverId := int64(97) // int64 | The Server ID for Filtering (optional)
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    appId := int64(5) // int64 | The App ID for Filtering (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HistoryApi.ListHistory(context.Background()).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).ZoneId(zoneId).AppId(appId).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **int64** | The Instance ID for Filtering | 
 **containerId** | **int64** | The Container ID for Filtering | 
 **serverId** | **int64** | The Server ID for Filtering | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **appId** | **int64** | The App ID for Filtering | 
 **phrase** | **string** | Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

