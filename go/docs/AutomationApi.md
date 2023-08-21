# \AutomationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExecuteSchedules**](AutomationApi.md#AddExecuteSchedules) | **Post** /api/execute-schedules | Creates a Execute Schedule
[**AddPowerScheduleInstances**](AutomationApi.md#AddPowerScheduleInstances) | **Put** /api/power-schedules/{id}/add-instances | Add Instances to a Power Schedule
[**AddPowerScheduleServers**](AutomationApi.md#AddPowerScheduleServers) | **Put** /api/power-schedules/{id}/add-servers | Add Servers to a Power Schedule
[**AddPowerSchedules**](AutomationApi.md#AddPowerSchedules) | **Post** /api/power-schedules | Creates a Power Schedule
[**AddScaleThresholds**](AutomationApi.md#AddScaleThresholds) | **Post** /api/scale-thresholds | Creates a Scale Threshold
[**AddTasks**](AutomationApi.md#AddTasks) | **Post** /api/tasks | Creates a Task
[**AddWorkflows**](AutomationApi.md#AddWorkflows) | **Post** /api/task-sets | Creates a Workflow
[**DeletePowerScheduleInstances**](AutomationApi.md#DeletePowerScheduleInstances) | **Put** /api/power-schedules/{id}/remove-instances | Remove Instances from a Power Schedule
[**DeletePowerScheduleServers**](AutomationApi.md#DeletePowerScheduleServers) | **Put** /api/power-schedules/{id}/remove-servers | Remove Servers from a Power Schedule
[**ExecuteExecutionRequest**](AutomationApi.md#ExecuteExecutionRequest) | **Post** /api/execution-request/execute | Executes an Execution Request
[**ExecuteTasks**](AutomationApi.md#ExecuteTasks) | **Post** /api/tasks/{id}/execute | Executes a Task
[**ExecuteWorkflows**](AutomationApi.md#ExecuteWorkflows) | **Post** /api/task-sets/{id}/execute | Executes a Workflow
[**GetExecuteSchedules**](AutomationApi.md#GetExecuteSchedules) | **Get** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**GetExecutionRequest**](AutomationApi.md#GetExecutionRequest) | **Get** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**GetPowerSchedules**](AutomationApi.md#GetPowerSchedules) | **Get** /api/power-schedules/{id} | Retrieves a Specific Power Schedule
[**GetScaleThresholds**](AutomationApi.md#GetScaleThresholds) | **Get** /api/scale-thresholds/{id} | Retrieves a Specific Scale Threshold
[**GetTaskTypes**](AutomationApi.md#GetTaskTypes) | **Get** /api/task-types/{id} | Retrieves a Specific Task Type
[**GetTasks**](AutomationApi.md#GetTasks) | **Get** /api/tasks/{id} | Retrieves a Specific Task
[**GetWorkflows**](AutomationApi.md#GetWorkflows) | **Get** /api/task-sets/{id} | Retrieves a Specific Workflow
[**ListExecuteSchedules**](AutomationApi.md#ListExecuteSchedules) | **Get** /api/execute-schedules | Retrieves all Execute Schedules
[**ListPowerSchedules**](AutomationApi.md#ListPowerSchedules) | **Get** /api/power-schedules | Retrieves all Power Schedules
[**ListScaleThresholds**](AutomationApi.md#ListScaleThresholds) | **Get** /api/scale-thresholds | Retrieves all Scale Thresholds
[**ListTaskTypes**](AutomationApi.md#ListTaskTypes) | **Get** /api/task-types | Retrieves all Task Types
[**ListTasks**](AutomationApi.md#ListTasks) | **Get** /api/tasks | Retrieves all Tasks
[**ListWorkflows**](AutomationApi.md#ListWorkflows) | **Get** /api/task-sets | Retrieves all Workflows
[**RemoveExecuteSchedules**](AutomationApi.md#RemoveExecuteSchedules) | **Delete** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**RemovePowerSchedules**](AutomationApi.md#RemovePowerSchedules) | **Delete** /api/power-schedules/{id} | Deletes a Power Schedule
[**RemoveScaleThresholds**](AutomationApi.md#RemoveScaleThresholds) | **Delete** /api/scale-thresholds/{id} | Deletes a Scale Threshold
[**RemoveTasks**](AutomationApi.md#RemoveTasks) | **Delete** /api/tasks/{id} | Deletes a Task
[**RemoveWorkflows**](AutomationApi.md#RemoveWorkflows) | **Delete** /api/task-sets/{id} | Deletes a Workflow
[**UpdateExecuteSchedules**](AutomationApi.md#UpdateExecuteSchedules) | **Put** /api/execute-schedules/{id} | Updates a Execute Schedule
[**UpdatePowerSchedules**](AutomationApi.md#UpdatePowerSchedules) | **Put** /api/power-schedules/{id} | Updates a Power Schedule
[**UpdateScaleThresholds**](AutomationApi.md#UpdateScaleThresholds) | **Put** /api/scale-thresholds/{id} | Updates a Scale Threshold
[**UpdateTasks**](AutomationApi.md#UpdateTasks) | **Put** /api/tasks/{id} | Updates a Task
[**UpdateWorkflows**](AutomationApi.md#UpdateWorkflows) | **Put** /api/task-sets/{id} | Updates a Workflow



## AddExecuteSchedules

> map[string]interface{} AddExecuteSchedules(ctx).InlineObject12(inlineObject12).Execute()

Creates a Execute Schedule



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
    inlineObject12 := *openapiclient.NewInlineObject12(*openapiclient.NewApiExecuteSchedulesSchedule("Sample Execution")) // InlineObject12 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddExecuteSchedules(context.Background()).InlineObject12(inlineObject12).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExecuteSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExecuteSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject12** | [**InlineObject12**](InlineObject12.md) |  | 

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


## AddPowerScheduleInstances

> Model200SuccessExpanded AddPowerScheduleInstances(ctx, id).InlineObject192(inlineObject192).Execute()

Add Instances to a Power Schedule



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
    inlineObject192 := *openapiclient.NewInlineObject192([]int64{int64(123)}) // InlineObject192 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddPowerScheduleInstances(context.Background(), id).InlineObject192(inlineObject192).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddPowerScheduleInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPowerScheduleInstances`: Model200SuccessExpanded
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddPowerScheduleInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPowerScheduleInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject192** | [**InlineObject192**](InlineObject192.md) |  | 

### Return type

[**Model200SuccessExpanded**](200-success-expanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPowerScheduleServers

> Model200SuccessExpanded AddPowerScheduleServers(ctx, id).InlineObject193(inlineObject193).Execute()

Add Servers to a Power Schedule



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
    inlineObject193 := *openapiclient.NewInlineObject193([]int64{int64(123)}) // InlineObject193 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddPowerScheduleServers(context.Background(), id).InlineObject193(inlineObject193).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddPowerScheduleServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPowerScheduleServers`: Model200SuccessExpanded
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddPowerScheduleServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPowerScheduleServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject193** | [**InlineObject193**](InlineObject193.md) |  | 

### Return type

[**Model200SuccessExpanded**](200-success-expanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPowerSchedules

> map[string]interface{} AddPowerSchedules(ctx).InlineObject190(inlineObject190).Execute()

Creates a Power Schedule



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
    inlineObject190 := *openapiclient.NewInlineObject190(*openapiclient.NewApiPowerSchedulesSchedule("Sample Threshold")) // InlineObject190 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddPowerSchedules(context.Background()).InlineObject190(inlineObject190).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddPowerSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPowerSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddPowerSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPowerSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject190** | [**InlineObject190**](InlineObject190.md) |  | 

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


## AddScaleThresholds

> map[string]interface{} AddScaleThresholds(ctx).InlineObject211(inlineObject211).Execute()

Creates a Scale Threshold



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
    inlineObject211 := *openapiclient.NewInlineObject211(*openapiclient.NewApiScaleThresholdsScaleThreshold("Sample Threshold")) // InlineObject211 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddScaleThresholds(context.Background()).InlineObject211(inlineObject211).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddScaleThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScaleThresholds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddScaleThresholds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddScaleThresholdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject211** | [**InlineObject211**](InlineObject211.md) |  | 

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


## AddTasks

> map[string]interface{} AddTasks(ctx).InlineObject246(inlineObject246).Execute()

Creates a Task



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
    inlineObject246 := *openapiclient.NewInlineObject246(*openapiclient.NewApiTasksTask("Sample Task", *openapiclient.NewApiTasksTaskTaskType("Code_example"), "ExecuteTarget_example")) // InlineObject246 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddTasks(context.Background()).InlineObject246(inlineObject246).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTasks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject246** | [**InlineObject246**](InlineObject246.md) |  | 

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


## AddWorkflows

> map[string]interface{} AddWorkflows(ctx).InlineObject249(inlineObject249).Execute()

Creates a Workflow



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
    inlineObject249 := *openapiclient.NewInlineObject249(*openapiclient.NewApiTaskSetsTaskSet("Sample Workflow")) // InlineObject249 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.AddWorkflows(context.Background()).InlineObject249(inlineObject249).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.AddWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWorkflows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.AddWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject249** | [**InlineObject249**](InlineObject249.md) |  | 

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


## DeletePowerScheduleInstances

> Model200SuccessExpanded DeletePowerScheduleInstances(ctx, id).InlineObject194(inlineObject194).Execute()

Remove Instances from a Power Schedule



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
    inlineObject194 := *openapiclient.NewInlineObject194([]int64{int64(123)}) // InlineObject194 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.DeletePowerScheduleInstances(context.Background(), id).InlineObject194(inlineObject194).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.DeletePowerScheduleInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePowerScheduleInstances`: Model200SuccessExpanded
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.DeletePowerScheduleInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePowerScheduleInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject194** | [**InlineObject194**](InlineObject194.md) |  | 

### Return type

[**Model200SuccessExpanded**](200-success-expanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePowerScheduleServers

> Model200SuccessExpanded DeletePowerScheduleServers(ctx, id).InlineObject195(inlineObject195).Execute()

Remove Servers from a Power Schedule



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
    inlineObject195 := *openapiclient.NewInlineObject195([]int64{int64(123)}) // InlineObject195 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.DeletePowerScheduleServers(context.Background(), id).InlineObject195(inlineObject195).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.DeletePowerScheduleServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePowerScheduleServers`: Model200SuccessExpanded
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.DeletePowerScheduleServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePowerScheduleServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject195** | [**InlineObject195**](InlineObject195.md) |  | 

### Return type

[**Model200SuccessExpanded**](200-success-expanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteExecutionRequest

> InlineResponse2008 ExecuteExecutionRequest(ctx).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).InlineObject14(inlineObject14).Execute()

Executes an Execution Request



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
    inlineObject14 := *openapiclient.NewInlineObject14("uname -a") // InlineObject14 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExecuteExecutionRequest(context.Background()).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).InlineObject14(inlineObject14).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExecuteExecutionRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteExecutionRequest`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExecuteExecutionRequest`: %v\n", resp)
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
 **inlineObject14** | [**InlineObject14**](InlineObject14.md) |  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteTasks

> map[string]interface{} ExecuteTasks(ctx, id).InlineObject248(inlineObject248).Execute()

Executes a Task



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
    inlineObject248 := *openapiclient.NewInlineObject248(*openapiclient.NewApiTasksIdExecuteJob()) // InlineObject248 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExecuteTasks(context.Background(), id).InlineObject248(inlineObject248).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExecuteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteTasks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExecuteTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject248** | [**InlineObject248**](InlineObject248.md) |  | 

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


## ExecuteWorkflows

> map[string]interface{} ExecuteWorkflows(ctx, id).InlineObject251(inlineObject251).Execute()

Executes a Workflow



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
    inlineObject251 := *openapiclient.NewInlineObject251(*openapiclient.NewApiTasksIdExecuteJob()) // InlineObject251 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExecuteWorkflows(context.Background(), id).InlineObject251(inlineObject251).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExecuteWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteWorkflows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExecuteWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject251** | [**InlineObject251**](InlineObject251.md) |  | 

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


## GetExecuteSchedules

> map[string]interface{} GetExecuteSchedules(ctx, id).Execute()

Retrieves a Specific Execute Schedule



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
    resp, r, err := api_client.AutomationApi.GetExecuteSchedules(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecuteSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetExecuteSchedules`: %v\n", resp)
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionRequest

> map[string]interface{} GetExecutionRequest(ctx, uniqueId).Execute()

Retrieves a Specific Execution Request



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
    uniqueId := "c56f75d0-a59a-4566-92e3-4dc716c5d076" // string | The Unique ID of the execution request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.GetExecutionRequest(context.Background(), uniqueId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetExecutionRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutionRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetExecutionRequest`: %v\n", resp)
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPowerSchedules

> map[string]interface{} GetPowerSchedules(ctx, id).Execute()

Retrieves a Specific Power Schedule



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
    resp, r, err := api_client.AutomationApi.GetPowerSchedules(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetPowerSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPowerSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetPowerSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPowerSchedulesRequest struct via the builder pattern


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


## GetScaleThresholds

> map[string]interface{} GetScaleThresholds(ctx, id).Execute()

Retrieves a Specific Scale Threshold



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
    resp, r, err := api_client.AutomationApi.GetScaleThresholds(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetScaleThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScaleThresholds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetScaleThresholds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScaleThresholdsRequest struct via the builder pattern


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


## GetTaskTypes

> InlineResponse200156 GetTaskTypes(ctx, id).Execute()

Retrieves a Specific Task Type



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
    resp, r, err := api_client.AutomationApi.GetTaskTypes(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetTaskTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskTypes`: InlineResponse200156
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetTaskTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200156**](inline_response_200_156.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> map[string]interface{} GetTasks(ctx, id).Execute()

Retrieves a Specific Task



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
    resp, r, err := api_client.AutomationApi.GetTasks(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


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


## GetWorkflows

> map[string]interface{} GetWorkflows(ctx, id).Execute()

Retrieves a Specific Workflow



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
    resp, r, err := api_client.AutomationApi.GetWorkflows(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsRequest struct via the builder pattern


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


## ListExecuteSchedules

> map[string]interface{} ListExecuteSchedules(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Execute Schedules



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
    resp, r, err := api_client.AutomationApi.ListExecuteSchedules(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ListExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExecuteSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ListExecuteSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExecuteSchedulesRequest struct via the builder pattern


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


## ListPowerSchedules

> map[string]interface{} ListPowerSchedules(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Power Schedules



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
    resp, r, err := api_client.AutomationApi.ListPowerSchedules(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ListPowerSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPowerSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ListPowerSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPowerSchedulesRequest struct via the builder pattern


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


## ListScaleThresholds

> map[string]interface{} ListScaleThresholds(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Scale Thresholds



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
    resp, r, err := api_client.AutomationApi.ListScaleThresholds(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ListScaleThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScaleThresholds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ListScaleThresholds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScaleThresholdsRequest struct via the builder pattern


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


## ListTaskTypes

> InlineResponse200155 ListTaskTypes(ctx).Name(name).Code(code).Execute()

Retrieves all Task Types



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
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ListTaskTypes(context.Background()).Name(name).Code(code).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ListTaskTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaskTypes`: InlineResponse200155
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ListTaskTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaskTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 

### Return type

[**InlineResponse200155**](inline_response_200_155.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> map[string]interface{} ListTasks(ctx).Type_(type_).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()

Retrieves all Tasks



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
    type_ := "type__example" // string | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.  (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ListTasks(context.Background()).Type_(type_).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTasks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## ListWorkflows

> map[string]interface{} ListWorkflows(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Type_(type_).Execute()

Retrieves all Workflows



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    type_ := "provision" // string | Filter by Workflow Type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ListWorkflows(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Type_(type_).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ListWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **type_** | **string** | Filter by Workflow Type | 

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
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.RemoveExecuteSchedules(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.RemoveExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExecuteSchedules`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.RemoveExecuteSchedules`: %v\n", resp)
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

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePowerSchedules

> Model200Success RemovePowerSchedules(ctx, id).Execute()

Deletes a Power Schedule



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
    resp, r, err := api_client.AutomationApi.RemovePowerSchedules(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.RemovePowerSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePowerSchedules`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.RemovePowerSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePowerSchedulesRequest struct via the builder pattern


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


## RemoveScaleThresholds

> Model200Success RemoveScaleThresholds(ctx, id).Execute()

Deletes a Scale Threshold



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
    resp, r, err := api_client.AutomationApi.RemoveScaleThresholds(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.RemoveScaleThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveScaleThresholds`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.RemoveScaleThresholds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveScaleThresholdsRequest struct via the builder pattern


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


## RemoveTasks

> Model200Success RemoveTasks(ctx, id).Execute()

Deletes a Task



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
    resp, r, err := api_client.AutomationApi.RemoveTasks(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.RemoveTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTasks`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.RemoveTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTasksRequest struct via the builder pattern


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


## RemoveWorkflows

> Model200Success RemoveWorkflows(ctx, id).Execute()

Deletes a Workflow



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
    resp, r, err := api_client.AutomationApi.RemoveWorkflows(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.RemoveWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveWorkflows`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.RemoveWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWorkflowsRequest struct via the builder pattern


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


## UpdateExecuteSchedules

> map[string]interface{} UpdateExecuteSchedules(ctx, id).InlineObject13(inlineObject13).Execute()

Updates a Execute Schedule



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
    inlineObject13 := *openapiclient.NewInlineObject13(*openapiclient.NewApiExecuteSchedulesIdSchedule()) // InlineObject13 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.UpdateExecuteSchedules(context.Background(), id).InlineObject13(inlineObject13).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.UpdateExecuteSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExecuteSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.UpdateExecuteSchedules`: %v\n", resp)
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

 **inlineObject13** | [**InlineObject13**](InlineObject13.md) |  | 

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


## UpdatePowerSchedules

> map[string]interface{} UpdatePowerSchedules(ctx, id).InlineObject191(inlineObject191).Execute()

Updates a Power Schedule



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
    inlineObject191 := *openapiclient.NewInlineObject191(*openapiclient.NewApiPowerSchedulesIdSchedule()) // InlineObject191 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.UpdatePowerSchedules(context.Background(), id).InlineObject191(inlineObject191).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.UpdatePowerSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePowerSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.UpdatePowerSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePowerSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject191** | [**InlineObject191**](InlineObject191.md) |  | 

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


## UpdateScaleThresholds

> map[string]interface{} UpdateScaleThresholds(ctx, id).InlineObject212(inlineObject212).Execute()

Updates a Scale Threshold



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
    inlineObject212 := *openapiclient.NewInlineObject212(*openapiclient.NewApiScaleThresholdsIdScaleThreshold()) // InlineObject212 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.UpdateScaleThresholds(context.Background(), id).InlineObject212(inlineObject212).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.UpdateScaleThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScaleThresholds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.UpdateScaleThresholds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScaleThresholdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject212** | [**InlineObject212**](InlineObject212.md) |  | 

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


## UpdateTasks

> map[string]interface{} UpdateTasks(ctx, id).InlineObject247(inlineObject247).Execute()

Updates a Task



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
    inlineObject247 := *openapiclient.NewInlineObject247(*openapiclient.NewApiTasksIdTask()) // InlineObject247 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.UpdateTasks(context.Background(), id).InlineObject247(inlineObject247).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.UpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTasks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.UpdateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject247** | [**InlineObject247**](InlineObject247.md) |  | 

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


## UpdateWorkflows

> map[string]interface{} UpdateWorkflows(ctx, id).InlineObject250(inlineObject250).Execute()

Updates a Workflow



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
    inlineObject250 := *openapiclient.NewInlineObject250(*openapiclient.NewApiTaskSetsIdTaskSet()) // InlineObject250 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.UpdateWorkflows(context.Background(), id).InlineObject250(inlineObject250).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.UpdateWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflows`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.UpdateWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject250** | [**InlineObject250**](InlineObject250.md) |  | 

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

