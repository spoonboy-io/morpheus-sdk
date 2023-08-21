# \InstancesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstance**](InstancesApi.md#AddInstance) | **Post** /api/instances | Create an Instance
[**BackupInstance**](InstancesApi.md#BackupInstance) | **Put** /api/instances/{id}/backup | Backup an instance
[**BackupsInstance**](InstancesApi.md#BackupsInstance) | **Get** /api/instances/{id}/backups | Get list of backups for an Instance
[**CancelExpirationInstance**](InstancesApi.md#CancelExpirationInstance) | **Put** /api/instances/{id}/cancel-expiration | Cancel Expiration of an Instance
[**CancelRemovalInstance**](InstancesApi.md#CancelRemovalInstance) | **Put** /api/instances/{id}/cancel-removal | Cancel Removal of an Instance
[**CancelShutdownInstance**](InstancesApi.md#CancelShutdownInstance) | **Put** /api/instances/{id}/cancel-shutdown | Cancel Shutdown of an Instance
[**CloneImageInstance**](InstancesApi.md#CloneImageInstance) | **Put** /api/instances/{id}/clone-image | Clone to Image
[**CloneInstance**](InstancesApi.md#CloneInstance) | **Put** /api/instances/{id}/clone | Clone an Instance
[**CreateInstanceSchedule**](InstancesApi.md#CreateInstanceSchedule) | **Post** /api/instances/{id}/schedules | Create a new Instance Schedule
[**DeleteAllSnapshotsInstance**](InstancesApi.md#DeleteAllSnapshotsInstance) | **Delete** /api/instances/{id}/delete-all-snapshots | Delete All Snapshots of Instance
[**DeleteAllSnapshotsInstanceContainer**](InstancesApi.md#DeleteAllSnapshotsInstanceContainer) | **Delete** /api/instances/{id}/delete-container-snapshots/{containerId} | Delete All Snapshots of Instance Container
[**DeleteInstance**](InstancesApi.md#DeleteInstance) | **Delete** /api/instances/{id} | Delete an instance
[**DeleteInstanceSchedule**](InstancesApi.md#DeleteInstanceSchedule) | **Delete** /api/instances/{id}/schedules/{scheduleId} | Deletes an Instance Schedule
[**DeleteSnapshotInstance**](InstancesApi.md#DeleteSnapshotInstance) | **Delete** /api/snapshots/{id} | Delete Snapshot of Instance
[**EjectInstance**](InstancesApi.md#EjectInstance) | **Put** /api/instances/{id}/eject | Eject an instance
[**ExtendExpirationInstance**](InstancesApi.md#ExtendExpirationInstance) | **Put** /api/instances/{id}/extend-expiration | Extend Expiration of an Instance
[**ExtendShutdownInstance**](InstancesApi.md#ExtendShutdownInstance) | **Put** /api/instances/{id}/extend-shutdown | Extend Shutdown of an Instance
[**GetEnvVariables**](InstancesApi.md#GetEnvVariables) | **Get** /api/instances/{id}/envs | Get Env Variables
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /api/instances/{id} | Retrieves a Specific Instance
[**GetInstanceContainers**](InstancesApi.md#GetInstanceContainers) | **Get** /api/instances/{id}/containers | Get Container Details
[**GetInstanceHistory**](InstancesApi.md#GetInstanceHistory) | **Get** /api/instances/{id}/history | Get Instance History
[**GetInstanceSchedule**](InstancesApi.md#GetInstanceSchedule) | **Get** /api/instances/{id}/schedules/{scheduleId} | Get a Specific Instance Schedule
[**GetInstanceSchedules**](InstancesApi.md#GetInstanceSchedules) | **Get** /api/instances/{id}/schedules | Get all Instance Schedules
[**GetInstanceThreshold**](InstancesApi.md#GetInstanceThreshold) | **Get** /api/instances/{id}/threshold | Get an Instance Scale Threshold
[**GetInstanceTypeProvisioning**](InstancesApi.md#GetInstanceTypeProvisioning) | **Get** /api/instance-types/{id} | Get Specific Instance Type for Provisioning
[**GetPrepareApplyInstance**](InstancesApi.md#GetPrepareApplyInstance) | **Get** /api/instances/{id}/prepare-apply | Prepare To Apply an Instance
[**GetSnapshotInstance**](InstancesApi.md#GetSnapshotInstance) | **Get** /api/snapshots/{id} | Get a Specific Snapshot
[**GetStateInstance**](InstancesApi.md#GetStateInstance) | **Get** /api/instances/{id}/state | Get State of an Instance
[**GetValidateApplyInstance**](InstancesApi.md#GetValidateApplyInstance) | **Post** /api/instances/{id}/validate-apply | Validate Apply State for an Instance
[**GetWikiInstance**](InstancesApi.md#GetWikiInstance) | **Get** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**ImportSnapshotInstance**](InstancesApi.md#ImportSnapshotInstance) | **Put** /api/instances/{id}/import-snapshot | Import Snapshot of an Instance
[**LinkedCloneSnapshotInstance**](InstancesApi.md#LinkedCloneSnapshotInstance) | **Put** /api/instances/{id}/linked-clone/{snapshotId} | Create Linked Clone of Instance Snapshot
[**ListInstanceServicePlans**](InstancesApi.md#ListInstanceServicePlans) | **Get** /api/instances/service-plans | Get Available Service Plans for an Instance
[**ListInstanceTypesProvisioning**](InstancesApi.md#ListInstanceTypesProvisioning) | **Get** /api/instance-types | Get All Instance Types for Provisioning
[**ListInstances**](InstancesApi.md#ListInstances) | **Get** /api/instances | Get All Instances
[**ListSecurityGroupsInstance**](InstancesApi.md#ListSecurityGroupsInstance) | **Get** /api/instances/{id}/security-groups | Get Security Groups for an Instance
[**LockInstance**](InstancesApi.md#LockInstance) | **Put** /api/instances/{id}/lock | Lock an Instance
[**RefreshStateInstance**](InstancesApi.md#RefreshStateInstance) | **Post** /api/instances/{id}/refresh | Refresh State of an Instance
[**RemoveInstancesFromControl**](InstancesApi.md#RemoveInstancesFromControl) | **Post** /api/instances/remove-from-control | Remove From Control
[**ResizeInstance**](InstancesApi.md#ResizeInstance) | **Put** /api/instances/{id}/resize | Resize an Instance
[**RestartInstance**](InstancesApi.md#RestartInstance) | **Put** /api/instances/{id}/restart | Restart an instance
[**RevertSnapshotInstance**](InstancesApi.md#RevertSnapshotInstance) | **Put** /api/instances/{id}/revert-snapshot/{snapshotId} | Revert Instance to Snapshot
[**RunWorkflowInstance**](InstancesApi.md#RunWorkflowInstance) | **Put** /api/instances/{id}/workflow | Run Workflow on an Instance
[**SetApplyInstance**](InstancesApi.md#SetApplyInstance) | **Post** /api/instances/{id}/apply | Apply State of an Instance
[**SetInstanceSecurityGroups**](InstancesApi.md#SetInstanceSecurityGroups) | **Post** /api/instances/{id}/security-groups | Set Security Groups for an Instance
[**SnapshotInstance**](InstancesApi.md#SnapshotInstance) | **Put** /api/instances/{id}/snapshot | Snapshot an Instance
[**SnapshotsInstance**](InstancesApi.md#SnapshotsInstance) | **Get** /api/instances/{id}/snapshots | Get list of snapshots for an Instance
[**StartInstance**](InstancesApi.md#StartInstance) | **Put** /api/instances/{id}/start | Start an instance
[**StopInstance**](InstancesApi.md#StopInstance) | **Put** /api/instances/{id}/stop | Stop an instance
[**SuspendInstance**](InstancesApi.md#SuspendInstance) | **Put** /api/instances/{id}/suspend | Suspend an instance
[**UnlockInstance**](InstancesApi.md#UnlockInstance) | **Put** /api/instances/{id}/unlock | Unlock an Instance
[**UpdateInstance**](InstancesApi.md#UpdateInstance) | **Put** /api/instances/{id} | Updating an Instance
[**UpdateInstanceNetworkInterface**](InstancesApi.md#UpdateInstanceNetworkInterface) | **Put** /api/instances/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for an Instance&#39;s Network
[**UpdateInstanceSchedule**](InstancesApi.md#UpdateInstanceSchedule) | **Put** /api/instances/{id}/schedules/{scheduleId} | Updating an Instance Schedule
[**UpdateInstanceThreshold**](InstancesApi.md#UpdateInstanceThreshold) | **Put** /api/instances/{id}/threshold | Updates an Instance Scale Threshold
[**UpdateWikiInstance**](InstancesApi.md#UpdateWikiInstance) | **Put** /api/instances/{id}/wiki | Update an Instance Wiki Page



## AddInstance

> map[string]interface{} AddInstance(ctx).InstanceCreate(instanceCreate).Execute()

Create an Instance



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
    instanceCreate := *openapiclient.NewInstanceCreate(*openapiclient.NewInstanceCreateInstance("myinstance", *openapiclient.NewInstanceCreateInstanceSite(int64(2)), *openapiclient.NewInstanceCreateInstanceInstanceType("Ubuntu"), *openapiclient.NewInstanceCreateInstanceLayout(int64(105)), *openapiclient.NewInstanceCreateInstancePlan(int64(75))), "TODO") // InstanceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.AddInstance(context.Background()).InstanceCreate(instanceCreate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.AddInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.AddInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceCreate** | [**InstanceCreate**](InstanceCreate.md) |  | 

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


## BackupInstance

> Model200Success BackupInstance(ctx, id).Execute()

Backup an instance



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
    resp, r, err := api_client.InstancesApi.BackupInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.BackupInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.BackupInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupInstanceRequest struct via the builder pattern


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


## BackupsInstance

> InstanceBackups BackupsInstance(ctx, id).Execute()

Get list of backups for an Instance



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
    resp, r, err := api_client.InstancesApi.BackupsInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.BackupsInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupsInstance`: InstanceBackups
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.BackupsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceBackups**](instanceBackups.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelExpirationInstance

> Model200Success CancelExpirationInstance(ctx, id).Execute()

Cancel Expiration of an Instance



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
    resp, r, err := api_client.InstancesApi.CancelExpirationInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CancelExpirationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelExpirationInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CancelExpirationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelExpirationInstanceRequest struct via the builder pattern


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


## CancelRemovalInstance

> Model200Success CancelRemovalInstance(ctx, id).Execute()

Cancel Removal of an Instance



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
    resp, r, err := api_client.InstancesApi.CancelRemovalInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CancelRemovalInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelRemovalInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CancelRemovalInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRemovalInstanceRequest struct via the builder pattern


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


## CancelShutdownInstance

> Model200Success CancelShutdownInstance(ctx, id).Execute()

Cancel Shutdown of an Instance



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
    resp, r, err := api_client.InstancesApi.CancelShutdownInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CancelShutdownInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelShutdownInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CancelShutdownInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelShutdownInstanceRequest struct via the builder pattern


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


## CloneImageInstance

> Model200Success CloneImageInstance(ctx, id).InstancesCloneImage(instancesCloneImage).Execute()

Clone to Image



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
    instancesCloneImage := *openapiclient.NewInstancesCloneImage() // InstancesCloneImage |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.CloneImageInstance(context.Background(), id).InstancesCloneImage(instancesCloneImage).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CloneImageInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneImageInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CloneImageInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneImageInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instancesCloneImage** | [**InstancesCloneImage**](InstancesCloneImage.md) |  | 

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


## CloneInstance

> Model200Success CloneInstance(ctx, id).InstanceClone(instanceClone).Execute()

Clone an Instance



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
    instanceClone := *openapiclient.NewInstanceClone() // InstanceClone |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.CloneInstance(context.Background(), id).InstanceClone(instanceClone).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CloneInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CloneInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceClone** | [**InstanceClone**](InstanceClone.md) |  | 

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


## CreateInstanceSchedule

> map[string]interface{} CreateInstanceSchedule(ctx, id).InlineObject96(inlineObject96).Execute()

Create a new Instance Schedule



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
    id := int64(42) // int64 | Instance ID
    inlineObject96 := *openapiclient.NewInlineObject96() // InlineObject96 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.CreateInstanceSchedule(context.Background(), id).InlineObject96(inlineObject96).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateInstanceSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstanceSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CreateInstanceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject96** | [**InlineObject96**](InlineObject96.md) |  | 

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


## DeleteAllSnapshotsInstance

> Model200Success DeleteAllSnapshotsInstance(ctx, id).Execute()

Delete All Snapshots of Instance



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
    resp, r, err := api_client.InstancesApi.DeleteAllSnapshotsInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteAllSnapshotsInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllSnapshotsInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteAllSnapshotsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllSnapshotsInstanceRequest struct via the builder pattern


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


## DeleteAllSnapshotsInstanceContainer

> Model200Success DeleteAllSnapshotsInstanceContainer(ctx, id, containerId).Execute()

Delete All Snapshots of Instance Container



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
    containerId := float32(4) // float32 | Container ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteAllSnapshotsInstanceContainer(context.Background(), id, containerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteAllSnapshotsInstanceContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllSnapshotsInstanceContainer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteAllSnapshotsInstanceContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**containerId** | **float32** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllSnapshotsInstanceContainerRequest struct via the builder pattern


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


## DeleteInstance

> Model200Success DeleteInstance(ctx, id).PreserveVolumes(preserveVolumes).KeepBackups(keepBackups).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()

Delete an instance



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
    preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
    keepBackups := "on" // string | Preserve copy of backups (optional) (default to "off")
    releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
    releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteInstance(context.Background(), id).PreserveVolumes(preserveVolumes).KeepBackups(keepBackups).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
 **keepBackups** | **string** | Preserve copy of backups | [default to &quot;off&quot;]
 **releaseFloatingIps** | **string** | Release Floating IPs | [default to &quot;on&quot;]
 **releaseEIPs** | **string** | Alias for releaseFloatingIps | [default to &quot;on&quot;]
 **force** | **string** | Force Delete | [default to &quot;off&quot;]

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


## DeleteInstanceSchedule

> Model200Success DeleteInstanceSchedule(ctx, id, scheduleId).Execute()

Deletes an Instance Schedule



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
    id := int64(42) // int64 | Instance ID
    scheduleId := int64(1) // int64 | Instance Schedule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteInstanceSchedule(context.Background(), id, scheduleId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteInstanceSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstanceSchedule`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteInstanceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance ID | 
**scheduleId** | **int64** | Instance Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceScheduleRequest struct via the builder pattern


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


## DeleteSnapshotInstance

> Model200Success DeleteSnapshotInstance(ctx, id).Execute()

Delete Snapshot of Instance



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
    resp, r, err := api_client.InstancesApi.DeleteSnapshotInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteSnapshotInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnapshotInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.DeleteSnapshotInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotInstanceRequest struct via the builder pattern


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


## EjectInstance

> Model200Success EjectInstance(ctx, id).Execute()

Eject an instance



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
    resp, r, err := api_client.InstancesApi.EjectInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.EjectInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EjectInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.EjectInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiEjectInstanceRequest struct via the builder pattern


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


## ExtendExpirationInstance

> Model200Success ExtendExpirationInstance(ctx, id).Execute()

Extend Expiration of an Instance



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
    resp, r, err := api_client.InstancesApi.ExtendExpirationInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ExtendExpirationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendExpirationInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ExtendExpirationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendExpirationInstanceRequest struct via the builder pattern


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


## ExtendShutdownInstance

> Model200Success ExtendShutdownInstance(ctx, id).Execute()

Extend Shutdown of an Instance



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
    resp, r, err := api_client.InstancesApi.ExtendShutdownInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ExtendShutdownInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendShutdownInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ExtendShutdownInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendShutdownInstanceRequest struct via the builder pattern


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


## GetEnvVariables

> InlineResponse20057 GetEnvVariables(ctx, id).Execute()

Get Env Variables



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
    resp, r, err := api_client.InstancesApi.GetEnvVariables(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetEnvVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvVariables`: InlineResponse20057
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetEnvVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> InlineResponse20056 GetInstance(ctx, id).Details(details).Execute()

Retrieves a Specific Instance



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
    details := true // bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstance(context.Background(), id).Details(details).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: InlineResponse20056
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **bool** | Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [default to false]

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceContainers

> map[string]interface{} GetInstanceContainers(ctx, id).Execute()

Get Container Details



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
    resp, r, err := api_client.InstancesApi.GetInstanceContainers(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceContainers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceContainersRequest struct via the builder pattern


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


## GetInstanceHistory

> map[string]interface{} GetInstanceHistory(ctx, id).Phrase(phrase).ContainerId(containerId).ServerId(serverId).ZoneId(zoneId).Execute()

Get Instance History



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
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    containerId := int64(135) // int64 | The Container ID for Filtering (optional)
    serverId := int64(97) // int64 | The Server ID for Filtering (optional)
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstanceHistory(context.Background(), id).Phrase(phrase).ContainerId(containerId).ServerId(serverId).ZoneId(zoneId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceHistory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **containerId** | **int64** | The Container ID for Filtering | 
 **serverId** | **int64** | The Server ID for Filtering | 
 **zoneId** | **int64** | The Zone ID for Filtering | 

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


## GetInstanceSchedule

> InlineResponse20059 GetInstanceSchedule(ctx, id, scheduleId).Execute()

Get a Specific Instance Schedule



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
    id := int64(42) // int64 | Instance ID
    scheduleId := int64(1) // int64 | Instance Schedule ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstanceSchedule(context.Background(), id, scheduleId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceSchedule`: InlineResponse20059
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance ID | 
**scheduleId** | **int64** | Instance Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceSchedules

> map[string]interface{} GetInstanceSchedules(ctx, id).Execute()

Get all Instance Schedules



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
    id := int64(42) // int64 | Instance ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstanceSchedules(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceSchedulesRequest struct via the builder pattern


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


## GetInstanceThreshold

> InlineResponse20058 GetInstanceThreshold(ctx, id).Execute()

Get an Instance Scale Threshold



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
    resp, r, err := api_client.InstancesApi.GetInstanceThreshold(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceThreshold`: InlineResponse20058
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeProvisioning

> map[string]interface{} GetInstanceTypeProvisioning(ctx, id).Execute()

Get Specific Instance Type for Provisioning



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
    resp, r, err := api_client.InstancesApi.GetInstanceTypeProvisioning(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceTypeProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceTypeProvisioning`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceTypeProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeProvisioningRequest struct via the builder pattern


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


## GetPrepareApplyInstance

> map[string]interface{} GetPrepareApplyInstance(ctx, id).Execute()

Prepare To Apply an Instance



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
    resp, r, err := api_client.InstancesApi.GetPrepareApplyInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetPrepareApplyInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrepareApplyInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetPrepareApplyInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrepareApplyInstanceRequest struct via the builder pattern


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


## GetSnapshotInstance

> Snapshot GetSnapshotInstance(ctx, id).Execute()

Get a Specific Snapshot



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
    resp, r, err := api_client.InstancesApi.GetSnapshotInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetSnapshotInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshotInstance`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetSnapshotInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Snapshot**](snapshot.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStateInstance

> map[string]interface{} GetStateInstance(ctx, id).Execute()

Get State of an Instance



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
    resp, r, err := api_client.InstancesApi.GetStateInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetStateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStateInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetStateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStateInstanceRequest struct via the builder pattern


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


## GetValidateApplyInstance

> map[string]interface{} GetValidateApplyInstance(ctx, id).InlineObject98(inlineObject98).Execute()

Validate Apply State for an Instance



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
    inlineObject98 := *openapiclient.NewInlineObject98() // InlineObject98 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetValidateApplyInstance(context.Background(), id).InlineObject98(inlineObject98).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetValidateApplyInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidateApplyInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetValidateApplyInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidateApplyInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject98** | [**InlineObject98**](InlineObject98.md) |  | 

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


## GetWikiInstance

> InlineResponse200168 GetWikiInstance(ctx, id).Execute()

Retrieves an Instance Wiki Page



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
    resp, r, err := api_client.InstancesApi.GetWikiInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetWikiInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiInstance`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetWikiInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200168**](inline_response_200_168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSnapshotInstance

> Model200Success ImportSnapshotInstance(ctx, id).InlineObject93(inlineObject93).Execute()

Import Snapshot of an Instance



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
    inlineObject93 := *openapiclient.NewInlineObject93() // InlineObject93 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ImportSnapshotInstance(context.Background(), id).InlineObject93(inlineObject93).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ImportSnapshotInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSnapshotInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ImportSnapshotInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSnapshotInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject93** | [**InlineObject93**](InlineObject93.md) |  | 

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


## LinkedCloneSnapshotInstance

> Model200Success LinkedCloneSnapshotInstance(ctx, id, snapshotId).Execute()

Create Linked Clone of Instance Snapshot



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
    snapshotId := float32(4) // float32 | Snapshot ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.LinkedCloneSnapshotInstance(context.Background(), id, snapshotId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.LinkedCloneSnapshotInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkedCloneSnapshotInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.LinkedCloneSnapshotInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**snapshotId** | **float32** | Snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkedCloneSnapshotInstanceRequest struct via the builder pattern


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


## ListInstanceServicePlans

> InlineResponse20060 ListInstanceServicePlans(ctx).ZoneId(zoneId).LayoutId(layoutId).SiteId(siteId).Execute()

Get Available Service Plans for an Instance



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
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    layoutId := int64(98) // int64 | The Layout ID for Filtering (optional)
    siteId := int64(7) // int64 | The Site ID for Filtering (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ListInstanceServicePlans(context.Background()).ZoneId(zoneId).LayoutId(layoutId).SiteId(siteId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListInstanceServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceServicePlans`: InlineResponse20060
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListInstanceServicePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **layoutId** | **int64** | The Layout ID for Filtering | 
 **siteId** | **int64** | The Site ID for Filtering | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceTypesProvisioning

> map[string]interface{} ListInstanceTypesProvisioning(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Featured(featured).Details(details).Execute()

Get All Instance Types for Provisioning



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
    code := "azr" // string | If specified will return an exact match on code (optional)
    featured := false // bool | Filter by featured (optional)
    details := true // bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ListInstanceTypesProvisioning(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Featured(featured).Details(details).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListInstanceTypesProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceTypesProvisioning`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListInstanceTypesProvisioning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypesProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **featured** | **bool** | Filter by featured | 
 **details** | **bool** | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | 

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


## ListInstances

> map[string]interface{} ListInstances(ctx).Max(max).Offset(offset).Name(name).Phrase(phrase).InstanceType(instanceType).LastUpdated(lastUpdated).CreatedBy(createdBy).AgentInstalled(agentInstalled).Status(status).Environment(environment).ShowDeleted(showDeleted).Deleted(deleted).ExpireDate(expireDate).ExpireDateMin(expireDateMin).ExpireDays(expireDays).ExpireDaysMin(expireDaysMin).ShutdownDate(shutdownDate).ShutdownDateMin(shutdownDateMin).ShutdownDays(shutdownDays).ShutdownDaysMin(shutdownDaysMin).Labels(labels).AllLabels(allLabels).Tags(tags).Metadata(metadata).Details(details).Execute()

Get All Instances



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
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    instanceType := "ubuntu" // string | The Instance Type Code for Filtering (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    createdBy := int64(63) // int64 | The User ID for Filtering (optional)
    agentInstalled := true // bool | Filter instances by if agent is installed or not on the associated servers. (optional)
    status := "running" // string | The instance status for filtering. (optional)
    environment := "lab" // string | The environment for filtering. (optional)
    showDeleted := true // bool | If true, includes instances in pending removal status. (optional) (default to false)
    deleted := true // bool | If true, only deleted resources or instances in pending removal status are returned. (optional)
    expireDate := "2019-01-01" // string | Filter by expireDate less than or equal to specified date (optional)
    expireDateMin := "2019-01-01" // string | Filter expireDate greater than or equal to the specified date (optional)
    expireDays := "20" // string | Filter by expireDays less than or equal to the specified value (optional)
    expireDaysMin := "20" // string | Filter by expireDays greater than or equal to the specified value (optional)
    shutdownDate := "2019-01-01" // string | Filter by shutdownDate less than equal to the specified date (optional)
    shutdownDateMin := "2019-01-01" // string | Filter by shutdownDate greater than or equal to the specified date (optional)
    shutdownDays := "20" // string | Filter by shutdownDays less than or equal to the specified value (optional)
    shutdownDaysMin := "20" // string | Filter by shutdownDays greater than or equal to the specified value (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    tags := "tags.env=qa&tags.env=test" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
    metadata := "tags.env=qa&tags.env=test" // string | Alias for tags (optional)
    details := true // bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ListInstances(context.Background()).Max(max).Offset(offset).Name(name).Phrase(phrase).InstanceType(instanceType).LastUpdated(lastUpdated).CreatedBy(createdBy).AgentInstalled(agentInstalled).Status(status).Environment(environment).ShowDeleted(showDeleted).Deleted(deleted).ExpireDate(expireDate).ExpireDateMin(expireDateMin).ExpireDays(expireDays).ExpireDaysMin(expireDaysMin).ShutdownDate(shutdownDate).ShutdownDateMin(shutdownDateMin).ShutdownDays(shutdownDays).ShutdownDaysMin(shutdownDaysMin).Labels(labels).AllLabels(allLabels).Tags(tags).Metadata(metadata).Details(details).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstances`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **instanceType** | **string** | The Instance Type Code for Filtering | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **createdBy** | **int64** | The User ID for Filtering | 
 **agentInstalled** | **bool** | Filter instances by if agent is installed or not on the associated servers. | 
 **status** | **string** | The instance status for filtering. | 
 **environment** | **string** | The environment for filtering. | 
 **showDeleted** | **bool** | If true, includes instances in pending removal status. | [default to false]
 **deleted** | **bool** | If true, only deleted resources or instances in pending removal status are returned. | 
 **expireDate** | **string** | Filter by expireDate less than or equal to specified date | 
 **expireDateMin** | **string** | Filter expireDate greater than or equal to the specified date | 
 **expireDays** | **string** | Filter by expireDays less than or equal to the specified value | 
 **expireDaysMin** | **string** | Filter by expireDays greater than or equal to the specified value | 
 **shutdownDate** | **string** | Filter by shutdownDate less than equal to the specified date | 
 **shutdownDateMin** | **string** | Filter by shutdownDate greater than or equal to the specified date | 
 **shutdownDays** | **string** | Filter by shutdownDays less than or equal to the specified value | 
 **shutdownDaysMin** | **string** | Filter by shutdownDays greater than or equal to the specified value | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **tags** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
 **metadata** | **string** | Alias for tags | 
 **details** | **bool** | Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [default to false]

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


## ListSecurityGroupsInstance

> map[string]interface{} ListSecurityGroupsInstance(ctx, id).Execute()

Get Security Groups for an Instance



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
    resp, r, err := api_client.InstancesApi.ListSecurityGroupsInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListSecurityGroupsInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityGroupsInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListSecurityGroupsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupsInstanceRequest struct via the builder pattern


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


## LockInstance

> Model200Success LockInstance(ctx, id).Execute()

Lock an Instance



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
    resp, r, err := api_client.InstancesApi.LockInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.LockInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.LockInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockInstanceRequest struct via the builder pattern


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


## RefreshStateInstance

> Model200Success RefreshStateInstance(ctx, id).Execute()

Refresh State of an Instance



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
    resp, r, err := api_client.InstancesApi.RefreshStateInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RefreshStateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshStateInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RefreshStateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshStateInstanceRequest struct via the builder pattern


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


## RemoveInstancesFromControl

> SuccessMessage RemoveInstancesFromControl(ctx).InlineObject99(inlineObject99).Execute()

Remove From Control



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
    inlineObject99 := *openapiclient.NewInlineObject99() // InlineObject99 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.RemoveInstancesFromControl(context.Background()).InlineObject99(inlineObject99).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RemoveInstancesFromControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInstancesFromControl`: SuccessMessage
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RemoveInstancesFromControl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInstancesFromControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject99** | [**InlineObject99**](InlineObject99.md) |  | 

### Return type

[**SuccessMessage**](successMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeInstance

> map[string]interface{} ResizeInstance(ctx, id).InstanceResize(instanceResize).Execute()

Resize an Instance



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
    instanceResize := *openapiclient.NewInstanceResize() // InstanceResize |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ResizeInstance(context.Background(), id).InstanceResize(instanceResize).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ResizeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ResizeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceResize** | [**InstanceResize**](InstanceResize.md) |  | 

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


## RestartInstance

> map[string]interface{} RestartInstance(ctx, id).Execute()

Restart an instance



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
    resp, r, err := api_client.InstancesApi.RestartInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RestartInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RestartInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartInstanceRequest struct via the builder pattern


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


## RevertSnapshotInstance

> Model200Success RevertSnapshotInstance(ctx, id, snapshotId).Execute()

Revert Instance to Snapshot



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
    snapshotId := float32(4) // float32 | Snapshot ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.RevertSnapshotInstance(context.Background(), id, snapshotId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RevertSnapshotInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertSnapshotInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RevertSnapshotInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**snapshotId** | **float32** | Snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertSnapshotInstanceRequest struct via the builder pattern


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


## RunWorkflowInstance

> Model200Success RunWorkflowInstance(ctx, id).WorkflowId(workflowId).WorkflowName(workflowName).InstanceWorkflow(instanceWorkflow).Execute()

Run Workflow on an Instance



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
    workflowId := int64(3) // int64 | ID of the workflow to execute (optional)
    workflowName := "myworkflow" // string | Name of the workflow to execute (optional)
    instanceWorkflow := *openapiclient.NewInstanceWorkflow() // InstanceWorkflow |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.RunWorkflowInstance(context.Background(), id).WorkflowId(workflowId).WorkflowName(workflowName).InstanceWorkflow(instanceWorkflow).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RunWorkflowInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunWorkflowInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RunWorkflowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunWorkflowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowId** | **int64** | ID of the workflow to execute | 
 **workflowName** | **string** | Name of the workflow to execute | 
 **instanceWorkflow** | [**InstanceWorkflow**](InstanceWorkflow.md) |  | 

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


## SetApplyInstance

> Model200Success SetApplyInstance(ctx, id).InlineObject91(inlineObject91).Execute()

Apply State of an Instance



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
    inlineObject91 := *openapiclient.NewInlineObject91() // InlineObject91 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.SetApplyInstance(context.Background(), id).InlineObject91(inlineObject91).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.SetApplyInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetApplyInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.SetApplyInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApplyInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject91** | [**InlineObject91**](InlineObject91.md) |  | 

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


## SetInstanceSecurityGroups

> map[string]interface{} SetInstanceSecurityGroups(ctx, id).InlineObject94(inlineObject94).Execute()

Set Security Groups for an Instance



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
    inlineObject94 := *openapiclient.NewInlineObject94() // InlineObject94 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.SetInstanceSecurityGroups(context.Background(), id).InlineObject94(inlineObject94).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.SetInstanceSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetInstanceSecurityGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.SetInstanceSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInstanceSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject94** | [**InlineObject94**](InlineObject94.md) |  | 

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


## SnapshotInstance

> Model200Success SnapshotInstance(ctx, id).InstanceSnapshot(instanceSnapshot).Execute()

Snapshot an Instance



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
    instanceSnapshot := *openapiclient.NewInstanceSnapshot() // InstanceSnapshot |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.SnapshotInstance(context.Background(), id).InstanceSnapshot(instanceSnapshot).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.SnapshotInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.SnapshotInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceSnapshot** | [**InstanceSnapshot**](InstanceSnapshot.md) |  | 

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


## SnapshotsInstance

> InstanceSnapshots SnapshotsInstance(ctx, id).Execute()

Get list of snapshots for an Instance



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
    resp, r, err := api_client.InstancesApi.SnapshotsInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.SnapshotsInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotsInstance`: InstanceSnapshots
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.SnapshotsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceSnapshots**](instanceSnapshots.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartInstance

> map[string]interface{} StartInstance(ctx, id).Execute()

Start an instance



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
    resp, r, err := api_client.InstancesApi.StartInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.StartInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.StartInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


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


## StopInstance

> map[string]interface{} StopInstance(ctx, id).Execute()

Stop an instance



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
    resp, r, err := api_client.InstancesApi.StopInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.StopInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.StopInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


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


## SuspendInstance

> Model200Success SuspendInstance(ctx, id).Execute()

Suspend an instance



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
    resp, r, err := api_client.InstancesApi.SuspendInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.SuspendInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.SuspendInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendInstanceRequest struct via the builder pattern


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


## UnlockInstance

> Model200Success UnlockInstance(ctx, id).Execute()

Unlock an Instance



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
    resp, r, err := api_client.InstancesApi.UnlockInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UnlockInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockInstance`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.UnlockInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockInstanceRequest struct via the builder pattern


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


## UpdateInstance

> map[string]interface{} UpdateInstance(ctx, id).InstanceUpdate(instanceUpdate).Execute()

Updating an Instance



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
    instanceUpdate := *openapiclient.NewInstanceUpdate() // InstanceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateInstance(context.Background(), id).InstanceUpdate(instanceUpdate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceUpdate** | [**InstanceUpdate**](InstanceUpdate.md) |  | 

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


## UpdateInstanceNetworkInterface

> map[string]interface{} UpdateInstanceNetworkInterface(ctx, id, networkInterfaceId).NetworkInterfaceUpdate(networkInterfaceUpdate).Execute()

Updating a label for an Instance's Network



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
    networkInterfaceId := float32(7) // float32 | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
    networkInterfaceUpdate := *openapiclient.NewNetworkInterfaceUpdate() // NetworkInterfaceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateInstanceNetworkInterface(context.Background(), id, networkInterfaceId).NetworkInterfaceUpdate(networkInterfaceUpdate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstanceNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceNetworkInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.UpdateInstanceNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**networkInterfaceId** | **float32** | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkInterfaceUpdate** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md) |  | 

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


## UpdateInstanceSchedule

> map[string]interface{} UpdateInstanceSchedule(ctx, id, scheduleId).InlineObject97(inlineObject97).Execute()

Updating an Instance Schedule



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
    id := int64(9) // int64 | Instance ID
    scheduleId := int64(1) // int64 | Instance Schedule ID
    inlineObject97 := *openapiclient.NewInlineObject97(*openapiclient.NewInstanceScheduleUpdate()) // InlineObject97 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateInstanceSchedule(context.Background(), id, scheduleId).InlineObject97(inlineObject97).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstanceSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.UpdateInstanceSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance ID | 
**scheduleId** | **int64** | Instance Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject97** | [**InlineObject97**](InlineObject97.md) |  | 

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


## UpdateInstanceThreshold

> map[string]interface{} UpdateInstanceThreshold(ctx, id).InlineObject95(inlineObject95).Execute()

Updates an Instance Scale Threshold



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
    inlineObject95 := *openapiclient.NewInlineObject95(*openapiclient.NewApiInstancesIdThresholdInstanceThreshold()) // InlineObject95 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateInstanceThreshold(context.Background(), id).InlineObject95(inlineObject95).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstanceThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceThreshold`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.UpdateInstanceThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject95** | [**InlineObject95**](InlineObject95.md) |  | 

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


## UpdateWikiInstance

> map[string]interface{} UpdateWikiInstance(ctx, id).InlineObject270(inlineObject270).Execute()

Update an Instance Wiki Page



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
    inlineObject270 := *openapiclient.NewInlineObject270() // InlineObject270 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateWikiInstance(context.Background(), id).InlineObject270(inlineObject270).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateWikiInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.UpdateWikiInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject270** | [**InlineObject270**](InlineObject270.md) |  | 

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

