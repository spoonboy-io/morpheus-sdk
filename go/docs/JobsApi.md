# \JobsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddJobs**](JobsApi.md#AddJobs) | **Post** /api/jobs | Creates a Job
[**ExecuteJobs**](JobsApi.md#ExecuteJobs) | **Put** /api/jobs/{id}/execute | Executes a Specific Job
[**GetJobExecutionEvents**](JobsApi.md#GetJobExecutionEvents) | **Get** /api/job-executions/{id}/events/{eventId} | Retrieves a Specific Job Execution Event
[**GetJobExecutions**](JobsApi.md#GetJobExecutions) | **Get** /api/job-executions/{id} | Retrieves a Specific Job Execution
[**GetJobs**](JobsApi.md#GetJobs) | **Get** /api/jobs/{id} | Retrieves a Specific Job
[**ListJobExecutions**](JobsApi.md#ListJobExecutions) | **Get** /api/job-executions | Retrieves all Job Executions
[**ListJobs**](JobsApi.md#ListJobs) | **Get** /api/jobs | Retrieves all Jobs
[**RemoveJobs**](JobsApi.md#RemoveJobs) | **Delete** /api/jobs/{id} | Deletes a Job
[**UpdateJobs**](JobsApi.md#UpdateJobs) | **Put** /api/jobs/{id} | Updates a Job



## AddJobs

> map[string]interface{} AddJobs(ctx).InlineObject103(inlineObject103).Execute()

Creates a Job



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
    inlineObject103 := *openapiclient.NewInlineObject103("TODO") // InlineObject103 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobsApi.AddJobs(context.Background()).InlineObject103(inlineObject103).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.AddJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.AddJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject103** | [**InlineObject103**](InlineObject103.md) |  | 

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


## ExecuteJobs

> Model200Success ExecuteJobs(ctx, id).CustomConfig(customConfig).Execute()

Executes a Specific Job



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
    customConfig := "customConfig=%7Bfoo%3Abar%7D" // string | Optional custom config (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobsApi.ExecuteJobs(context.Background(), id).CustomConfig(customConfig).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ExecuteJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteJobs`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.ExecuteJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customConfig** | **string** | Optional custom config | 

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


## GetJobExecutionEvents

> map[string]interface{} GetJobExecutionEvents(ctx, id, eventId).Execute()

Retrieves a Specific Job Execution Event



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
    eventId := int64(323) // int64 | ID of the job execution event

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobsApi.GetJobExecutionEvents(context.Background(), id, eventId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobExecutionEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobExecutionEvents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobExecutionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**eventId** | **int64** | ID of the job execution event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobExecutionEventsRequest struct via the builder pattern


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


## GetJobExecutions

> map[string]interface{} GetJobExecutions(ctx, id).Execute()

Retrieves a Specific Job Execution



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
    resp, r, err := api_client.JobsApi.GetJobExecutions(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobExecutions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobExecutionsRequest struct via the builder pattern


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


## GetJobs

> map[string]interface{} GetJobs(ctx, id).Execute()

Retrieves a Specific Job



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
    resp, r, err := api_client.JobsApi.GetJobs(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


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


## ListJobExecutions

> map[string]interface{} ListJobExecutions(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Job Executions



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
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobsApi.ListJobExecutions(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ListJobExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobExecutions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.ListJobExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 

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


## ListJobs

> map[string]interface{} ListJobs(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).ItemSource(itemSource).Labels(labels).AllLabels(allLabels).Execute()

Retrieves all Jobs



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
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    itemSource := "itemSource_example" // string | Source filter, restricts query to only load jobs of specified source. (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobsApi.ListJobs(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).ItemSource(itemSource).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ListJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.ListJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **itemSource** | **string** | Source filter, restricts query to only load jobs of specified source. | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

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


## RemoveJobs

> Model200Success RemoveJobs(ctx, id).Execute()

Deletes a Job



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
    resp, r, err := api_client.JobsApi.RemoveJobs(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.RemoveJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveJobs`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.RemoveJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveJobsRequest struct via the builder pattern


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


## UpdateJobs

> map[string]interface{} UpdateJobs(ctx, id).InlineObject104(inlineObject104).Execute()

Updates a Job



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
    inlineObject104 := *openapiclient.NewInlineObject104(*openapiclient.NewApiJobsIdJob()) // InlineObject104 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobsApi.UpdateJobs(context.Background(), id).InlineObject104(inlineObject104).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.UpdateJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.UpdateJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject104** | [**InlineObject104**](InlineObject104.md) |  | 

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

