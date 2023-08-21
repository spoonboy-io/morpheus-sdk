# \IncidentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIncident**](IncidentsApi.md#AddIncident) | **Post** /api/monitoring/incidents | Create a New Incident
[**DeleteIncidents**](IncidentsApi.md#DeleteIncidents) | **Delete** /api/monitoring/incidents/{id} | Close a Specific Incident
[**GetIncidents**](IncidentsApi.md#GetIncidents) | **Get** /api/monitoring/incidents/{id} | Get a Specific Incident
[**ListIncidents**](IncidentsApi.md#ListIncidents) | **Get** /api/monitoring/incidents | List All Incidents
[**UpdateIncidents**](IncidentsApi.md#UpdateIncidents) | **Put** /api/monitoring/incidents/{id} | Update Incident
[**UpdateIncidentsReopen**](IncidentsApi.md#UpdateIncidentsReopen) | **Get** /api/monitoring/incidents/{id}/reopen | Reopen a Specific Incident
[**UpdateMuteAllIncidents**](IncidentsApi.md#UpdateMuteAllIncidents) | **Put** /api/monitoring/incidents/mute-all | Mute All Incidents
[**UpdateMuteIncidents**](IncidentsApi.md#UpdateMuteIncidents) | **Put** /api/monitoring/incidents/{id}/mute | Mute Incident



## AddIncident

> map[string]interface{} AddIncident(ctx).InlineObject87(inlineObject87).Execute()

Create a New Incident



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
    inlineObject87 := *openapiclient.NewInlineObject87(*openapiclient.NewApiMonitoringIncidentsIncident("Incident Name")) // InlineObject87 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.AddIncident(context.Background()).InlineObject87(inlineObject87).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.AddIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIncident`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.AddIncident`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject87** | [**InlineObject87**](InlineObject87.md) |  | 

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


## DeleteIncidents

> Model200Success DeleteIncidents(ctx, id).Execute()

Close a Specific Incident



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
    resp, r, err := api_client.IncidentsApi.DeleteIncidents(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIncidents`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.DeleteIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentsRequest struct via the builder pattern


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


## GetIncidents

> InlineResponse20055 GetIncidents(ctx, id).Execute()

Get a Specific Incident



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
    resp, r, err := api_client.IncidentsApi.GetIncidents(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidents`: InlineResponse20055
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncidents

> map[string]interface{} ListIncidents(ctx).Max(max).Offset(offset).Status(status).Severity(severity).Execute()

List All Incidents



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
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    status := "running" // string | The instance status for filtering. (optional)
    severity := "severity_example" // string | Filter by severity (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.ListIncidents(context.Background()).Max(max).Offset(offset).Status(status).Severity(severity).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ListIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **status** | **string** | The instance status for filtering. | 
 **severity** | **string** | Filter by severity | 

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


## UpdateIncidents

> map[string]interface{} UpdateIncidents(ctx, id).InlineObject88(inlineObject88).Execute()

Update Incident



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
    inlineObject88 := *openapiclient.NewInlineObject88(*openapiclient.NewApiMonitoringIncidentsIdIncident()) // InlineObject88 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.UpdateIncidents(context.Background(), id).InlineObject88(inlineObject88).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncidents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject88** | [**InlineObject88**](InlineObject88.md) |  | 

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


## UpdateIncidentsReopen

> SuccessMessage UpdateIncidentsReopen(ctx, id).Execute()

Reopen a Specific Incident



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
    resp, r, err := api_client.IncidentsApi.UpdateIncidentsReopen(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentsReopen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncidentsReopen`: SuccessMessage
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentsReopen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncidentsReopenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessMessage**](successMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMuteAllIncidents

> map[string]interface{} UpdateMuteAllIncidents(ctx).InlineObject90(inlineObject90).Execute()

Mute All Incidents



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
    inlineObject90 := *openapiclient.NewInlineObject90(false) // InlineObject90 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.UpdateMuteAllIncidents(context.Background()).InlineObject90(inlineObject90).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateMuteAllIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteAllIncidents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateMuteAllIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteAllIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject90** | [**InlineObject90**](InlineObject90.md) |  | 

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


## UpdateMuteIncidents

> map[string]interface{} UpdateMuteIncidents(ctx, id).InlineObject89(inlineObject89).Execute()

Mute Incident



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
    inlineObject89 := *openapiclient.NewInlineObject89(false) // InlineObject89 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.UpdateMuteIncidents(context.Background(), id).InlineObject89(inlineObject89).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateMuteIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteIncidents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateMuteIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject89** | [**InlineObject89**](InlineObject89.md) |  | 

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

