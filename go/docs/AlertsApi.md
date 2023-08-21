# \AlertsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAlerts**](AlertsApi.md#AddAlerts) | **Post** /api/monitoring/alerts | Create a New Alert
[**DeleteAlerts**](AlertsApi.md#DeleteAlerts) | **Delete** /api/monitoring/alerts/{id} | Delete a Specific Alert
[**GetAlerts**](AlertsApi.md#GetAlerts) | **Get** /api/monitoring/alerts/{id} | Get a Specific Alert
[**ListAlerts**](AlertsApi.md#ListAlerts) | **Get** /api/monitoring/alerts | List All Alerts
[**UpdateAlerts**](AlertsApi.md#UpdateAlerts) | **Put** /api/monitoring/alerts/{id} | Update Alert



## AddAlerts

> map[string]interface{} AddAlerts(ctx).InlineObject(inlineObject).Execute()

Create a New Alert



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
    inlineObject := *openapiclient.NewInlineObject(*openapiclient.NewApiMonitoringAlertsAlert("My Alert")) // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.AddAlerts(context.Background()).InlineObject(inlineObject).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.AddAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAlerts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.AddAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlerts

> Model200Success DeleteAlerts(ctx, id).Execute()

Delete a Specific Alert



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
    resp, r, err := api_client.AlertsApi.DeleteAlerts(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlerts`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsRequest struct via the builder pattern


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


## GetAlerts

> map[string]interface{} GetAlerts(ctx, id).Execute()

Get a Specific Alert



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
    resp, r, err := api_client.AlertsApi.GetAlerts(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlerts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListAlerts

> map[string]interface{} ListAlerts(ctx).Max(max).Offset(offset).LastUpdated(lastUpdated).Execute()

List All Alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlerts(context.Background()).Max(max).Offset(offset).LastUpdated(lastUpdated).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlerts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 

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


## UpdateAlerts

> Model200Success UpdateAlerts(ctx, id).InlineObject1(inlineObject1).Execute()

Update Alert



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
    inlineObject1 := *openapiclient.NewInlineObject1(*openapiclient.NewApiMonitoringAlertsIdAlert()) // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlerts(context.Background(), id).InlineObject1(inlineObject1).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlerts`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

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

