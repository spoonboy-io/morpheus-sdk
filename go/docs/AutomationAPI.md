# \AutomationAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExecuteSchedules**](AutomationAPI.md#AddExecuteSchedules) | **Post** /api/execute-schedules | Creates a Execute Schedule
[**ExecuteExecutionRequest**](AutomationAPI.md#ExecuteExecutionRequest) | **Post** /api/execution-request/execute | Executes an Execution Request
[**GetExecuteSchedules**](AutomationAPI.md#GetExecuteSchedules) | **Get** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**GetExecutionRequest**](AutomationAPI.md#GetExecutionRequest) | **Get** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**ListExecuteSchedules**](AutomationAPI.md#ListExecuteSchedules) | **Get** /api/execute-schedules | Retrieves all Execute Schedules
[**RemoveExecuteSchedules**](AutomationAPI.md#RemoveExecuteSchedules) | **Delete** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**UpdateExecuteSchedules**](AutomationAPI.md#UpdateExecuteSchedules) | **Put** /api/execute-schedules/{id} | Updates a Execute Schedule



## AddExecuteSchedules

> AddExecuteSchedules200Response AddExecuteSchedules(ctx).AddExecuteSchedulesRequest(addExecuteSchedulesRequest).Execute()

Creates a Execute Schedule



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
    addExecuteSchedulesRequest := *openapiclient.NewAddExecuteSchedulesRequest(*openapiclient.NewAddExecuteSchedulesRequestSchedule("Sample Execution")) // AddExecuteSchedulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.AddExecuteSchedules(context.Background()).AddExecuteSchedulesRequest(addExecuteSchedulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.AddExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExecuteSchedules`: AddExecuteSchedules200Response
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.AddExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExecuteSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addExecuteSchedulesRequest** | [**AddExecuteSchedulesRequest**](AddExecuteSchedulesRequest.md) |  | 

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteExecutionRequest

> ExecuteExecutionRequest200Response ExecuteExecutionRequest(ctx).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).ExecuteExecutionRequestRequest(executeExecutionRequestRequest).Execute()

Executes an Execution Request



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
    instanceId := int64(94) // int64 | The Instance ID for Filtering (optional)
    containerId := int64(135) // int64 | The Container ID for Filtering (optional)
    serverId := int64(97) // int64 | The Server ID for Filtering (optional)
    executeExecutionRequestRequest := *openapiclient.NewExecuteExecutionRequestRequest("uname -a") // ExecuteExecutionRequestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.ExecuteExecutionRequest(context.Background()).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).ExecuteExecutionRequestRequest(executeExecutionRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.ExecuteExecutionRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteExecutionRequest`: ExecuteExecutionRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.ExecuteExecutionRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteExecutionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **int64** | The Instance ID for Filtering | 
 **containerId** | **int64** | The Container ID for Filtering | 
 **serverId** | **int64** | The Server ID for Filtering | 
 **executeExecutionRequestRequest** | [**ExecuteExecutionRequestRequest**](ExecuteExecutionRequestRequest.md) |  | 

### Return type

[**ExecuteExecutionRequest200Response**](ExecuteExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecuteSchedules

> GetExecuteSchedules200Response GetExecuteSchedules(ctx, id).Execute()

Retrieves a Specific Execute Schedule



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.GetExecuteSchedules(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.GetExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecuteSchedules`: GetExecuteSchedules200Response
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.GetExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecuteSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetExecuteSchedules200Response**](GetExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionRequest

> GetExecutionRequest200Response GetExecutionRequest(ctx, uniqueId).Execute()

Retrieves a Specific Execution Request



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
    uniqueId := "c56f75d0-a59a-4566-92e3-4dc716c5d076" // string | The Unique ID of the execution request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.GetExecutionRequest(context.Background(), uniqueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.GetExecutionRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutionRequest`: GetExecutionRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.GetExecutionRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqueId** | **string** | The Unique ID of the execution request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetExecutionRequest200Response**](GetExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecuteSchedules

> ListExecuteSchedules200Response ListExecuteSchedules(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Execute Schedules



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
    max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.ListExecuteSchedules(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.ListExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExecuteSchedules`: ListExecuteSchedules200Response
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.ListExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExecuteSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 

### Return type

[**ListExecuteSchedules200Response**](ListExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExecuteSchedules

> Model200Success RemoveExecuteSchedules(ctx, id).Execute()

Deletes a Execute Schedule



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.RemoveExecuteSchedules(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.RemoveExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExecuteSchedules`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.RemoveExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExecuteSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateExecuteSchedules

> AddExecuteSchedules200Response UpdateExecuteSchedules(ctx, id).UpdateExecuteSchedulesRequest(updateExecuteSchedulesRequest).Execute()

Updates a Execute Schedule



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    updateExecuteSchedulesRequest := *openapiclient.NewUpdateExecuteSchedulesRequest(*openapiclient.NewUpdateExecuteSchedulesRequestSchedule()) // UpdateExecuteSchedulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomationAPI.UpdateExecuteSchedules(context.Background(), id).UpdateExecuteSchedulesRequest(updateExecuteSchedulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.UpdateExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExecuteSchedules`: AddExecuteSchedules200Response
    fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.UpdateExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExecuteSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExecuteSchedulesRequest** | [**UpdateExecuteSchedulesRequest**](UpdateExecuteSchedulesRequest.md) |  | 

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

