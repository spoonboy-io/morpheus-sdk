# \ClustersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCluster**](ClustersApi.md#AddCluster) | **Post** /api/clusters | Create a Cluster
[**AddClusterNamespace**](ClustersApi.md#AddClusterNamespace) | **Post** /api/clusters/{clusterId}/namespaces | Add Namespace (Kubernetes)
[**AddClusterWorker**](ClustersApi.md#AddClusterWorker) | **Post** /api/clusters/{clusterId}/servers | Add Worker
[**ApplyTemplate**](ClustersApi.md#ApplyTemplate) | **Post** /api/clusters/{clusterId}/apply-template | Apply Template to Cluster (Kubernetes)
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /api/clusters/{clusterId} | Delete a Cluster
[**DeleteClusterContainer**](ClustersApi.md#DeleteClusterContainer) | **Delete** /api/clusters/{clusterId}/containers/{id} | Delete Container
[**DeleteClusterDeployment**](ClustersApi.md#DeleteClusterDeployment) | **Delete** /api/clusters/{clusterId}/deployments/{id} | Delete Deployment
[**DeleteClusterJob**](ClustersApi.md#DeleteClusterJob) | **Delete** /api/clusters/{clusterId}/jobs/{id} | Delete a Job
[**DeleteClusterNamespace**](ClustersApi.md#DeleteClusterNamespace) | **Delete** /api/clusters/{clusterId}/namespaces/{id} | Delete a Namespace (Kubernetes)
[**DeleteClusterService**](ClustersApi.md#DeleteClusterService) | **Delete** /api/clusters/{clusterId}/services/{id} | Delete a Service
[**DeleteClusterStatefulSet**](ClustersApi.md#DeleteClusterStatefulSet) | **Delete** /api/clusters/{clusterId}/statefulsets/{id} | Delete a Stateful Set
[**DeleteClusterVolume**](ClustersApi.md#DeleteClusterVolume) | **Delete** /api/clusters/{clusterId}/volumes/{id} | Delete a Volume
[**DeleteClusterWorker**](ClustersApi.md#DeleteClusterWorker) | **Delete** /api/clusters/{clusterId}/servers/{id} | Delete a Worker
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /api/clusters/{clusterId} | Get a Specific Cluster
[**GetClusterApiConfig**](ClustersApi.md#GetClusterApiConfig) | **Get** /api/clusters/{clusterId}/api-config | Get API Config
[**GetClusterDatastore**](ClustersApi.md#GetClusterDatastore) | **Get** /api/clusters/{clusterId}/datastores/{id} | Get a Specific Datastore
[**GetClusterHistory**](ClustersApi.md#GetClusterHistory) | **Get** /api/clusters/{clusterId}/history | Get Cluster History
[**GetClusterHistoryDetail**](ClustersApi.md#GetClusterHistoryDetail) | **Get** /api/clusters/{clusterId}/history/{id} | Get Cluster History Details
[**GetClusterHistoryEventDetail**](ClustersApi.md#GetClusterHistoryEventDetail) | **Get** /api/clusters/{clusterId}/history/events/{id} | Get Cluster History Event
[**GetClusterMasters**](ClustersApi.md#GetClusterMasters) | **Get** /api/clusters/{clusterId}/masters | Get Masters (Kubernetes)
[**GetClusterNamespace**](ClustersApi.md#GetClusterNamespace) | **Get** /api/clusters/{clusterId}/namespaces/{id} | Get Namespace (Kubernetes)
[**GetClusterNamespaces**](ClustersApi.md#GetClusterNamespaces) | **Get** /api/clusters/{clusterId}/namespaces | List Namespaces (Kubernetes)
[**GetClusterUpgradeVersions**](ClustersApi.md#GetClusterUpgradeVersions) | **Get** /api/clusters/{clusterId}/upgrade-cluster | Get Cluster Upgrade Versions (Kubernetes)
[**GetWikiCluster**](ClustersApi.md#GetWikiCluster) | **Get** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**ListClusterContainers**](ClustersApi.md#ListClusterContainers) | **Get** /api/clusters/{clusterId}/containers | Get Containers
[**ListClusterDatastores**](ClustersApi.md#ListClusterDatastores) | **Get** /api/clusters/{clusterId}/datastores | Get Datastores
[**ListClusterDeployments**](ClustersApi.md#ListClusterDeployments) | **Get** /api/clusters/{clusterId}/deployments | Get Deployments
[**ListClusterJobs**](ClustersApi.md#ListClusterJobs) | **Get** /api/clusters/{clusterId}/jobs | Get Jobs
[**ListClusterPods**](ClustersApi.md#ListClusterPods) | **Get** /api/clusters/{clusterId}/pods | Get Pods
[**ListClusterServices**](ClustersApi.md#ListClusterServices) | **Get** /api/clusters/{clusterId}/services | Get Services
[**ListClusterStatefulSets**](ClustersApi.md#ListClusterStatefulSets) | **Get** /api/clusters/{clusterId}/statefulsets | Get Stateful Sets
[**ListClusterTypes**](ClustersApi.md#ListClusterTypes) | **Get** /api/cluster-types | Get All Cluster Types
[**ListClusterVolumes**](ClustersApi.md#ListClusterVolumes) | **Get** /api/clusters/{clusterId}/volumes | Get Volumes
[**ListClusterWorkers**](ClustersApi.md#ListClusterWorkers) | **Get** /api/clusters/{clusterId}/workers | Get Workers
[**ListClusters**](ClustersApi.md#ListClusters) | **Get** /api/clusters | Get All Clusters
[**RestartClusterContainer**](ClustersApi.md#RestartClusterContainer) | **Put** /api/clusters/{clusterId}/containers/{id}/restart | Restart a Container
[**RestartClusterDeployment**](ClustersApi.md#RestartClusterDeployment) | **Put** /api/clusters/{clusterId}/deployments/{id}/restart | Restart a Deployment
[**RestartClusterPod**](ClustersApi.md#RestartClusterPod) | **Put** /api/clusters/{clusterId}/pods/{id}/restart | Restart a Pod
[**RestartClusterStatefulSet**](ClustersApi.md#RestartClusterStatefulSet) | **Put** /api/clusters/{clusterId}/statefulsets/{id}/restart | Restart a Stateful Set
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Put** /api/clusters/{clusterId} | Update Cluster
[**UpdateClusterDatastore**](ClustersApi.md#UpdateClusterDatastore) | **Put** /api/clusters/{clusterId}/datastores/{id} | Update Datastore
[**UpdateClusterNamespace**](ClustersApi.md#UpdateClusterNamespace) | **Put** /api/clusters/{clusterId}/namespaces/{id} | Update Namespace (Kubernetes)
[**UpdateClusterPermissions**](ClustersApi.md#UpdateClusterPermissions) | **Put** /api/clusters/{clusterId}/permissions | Update Cluster Permissions
[**UpdateClusterUpgradeVersions**](ClustersApi.md#UpdateClusterUpgradeVersions) | **Post** /api/clusters/{clusterId}/upgrade-cluster | Upgrade a Cluster (Kubernetes)
[**UpdateClusterWorkerCount**](ClustersApi.md#UpdateClusterWorkerCount) | **Put** /api/clusters/{clusterId}/worker-count | Update Worker Count
[**UpdateWikiCluster**](ClustersApi.md#UpdateWikiCluster) | **Put** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page



## AddCluster

> map[string]interface{} AddCluster(ctx).InlineObject52(inlineObject52).Execute()

Create a Cluster



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
    inlineObject52 := *openapiclient.NewInlineObject52() // InlineObject52 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.AddCluster(context.Background()).InlineObject52(inlineObject52).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.AddCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.AddCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject52** | [**InlineObject52**](InlineObject52.md) |  | 

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


## AddClusterNamespace

> map[string]interface{} AddClusterNamespace(ctx, clusterId).InlineObject56(inlineObject56).Execute()

Add Namespace (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    inlineObject56 := *openapiclient.NewInlineObject56() // InlineObject56 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.AddClusterNamespace(context.Background(), clusterId).InlineObject56(inlineObject56).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.AddClusterNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClusterNamespace`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.AddClusterNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject56** | [**InlineObject56**](InlineObject56.md) |  | 

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


## AddClusterWorker

> map[string]interface{} AddClusterWorker(ctx, clusterId).InlineObject59(inlineObject59).Execute()

Add Worker



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
    clusterId := int32(5) // int32 | The ID of the cluster
    inlineObject59 := *openapiclient.NewInlineObject59() // InlineObject59 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.AddClusterWorker(context.Background(), clusterId).InlineObject59(inlineObject59).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.AddClusterWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClusterWorker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.AddClusterWorker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterWorkerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject59** | [**InlineObject59**](InlineObject59.md) |  | 

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


## ApplyTemplate

> ClusterApplyTemplate ApplyTemplate(ctx, clusterId).InlineObject54(inlineObject54).Execute()

Apply Template to Cluster (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    inlineObject54 := *openapiclient.NewInlineObject54() // InlineObject54 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ApplyTemplate(context.Background(), clusterId).InlineObject54(inlineObject54).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ApplyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyTemplate`: ClusterApplyTemplate
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ApplyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject54** | [**InlineObject54**](InlineObject54.md) |  | 

### Return type

[**ClusterApplyTemplate**](clusterApplyTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> Model200Success DeleteCluster(ctx, clusterId).RemoveInstances(removeInstances).RemoveResources(removeResources).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()

Delete a Cluster



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
    clusterId := int32(5) // int32 | The ID of the cluster
    removeInstances := "on" // string | Remove Instances (optional) (default to "off")
    removeResources := "off" // string | Remove Resources (optional) (default to "on")
    preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
    releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
    releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteCluster(context.Background(), clusterId).RemoveInstances(removeInstances).RemoveResources(removeResources).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeInstances** | **string** | Remove Instances | [default to &quot;off&quot;]
 **removeResources** | **string** | Remove Resources | [default to &quot;on&quot;]
 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
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


## DeleteClusterContainer

> Model200Success DeleteClusterContainer(ctx, clusterId, id).Force(force).Execute()

Delete Container



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterContainer(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterContainer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteClusterDeployment

> Model200Success DeleteClusterDeployment(ctx, clusterId, id).Force(force).Execute()

Delete Deployment



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterDeployment(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterDeployment`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteClusterJob

> Model200Success DeleteClusterJob(ctx, clusterId, id).Force(force).Execute()

Delete a Job



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterJob(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterJob`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteClusterNamespace

> Model200Success DeleteClusterNamespace(ctx, clusterId, id).Force(force).Execute()

Delete a Namespace (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterNamespace(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterNamespace`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteClusterService

> Model200Success DeleteClusterService(ctx, clusterId, id).Force(force).Execute()

Delete a Service



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterService(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterService`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteClusterStatefulSet

> SuccessError DeleteClusterStatefulSet(ctx, clusterId, id).Force(force).Execute()

Delete a Stateful Set



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterStatefulSet(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterStatefulSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterStatefulSet`: SuccessError
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterStatefulSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterStatefulSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **string** | Force Delete | [default to &quot;off&quot;]

### Return type

[**SuccessError**](successError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterVolume

> Model200Success DeleteClusterVolume(ctx, clusterId, id).Force(force).Execute()

Delete a Volume



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterVolume(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterVolume`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteClusterWorker

> Model200Success DeleteClusterWorker(ctx, clusterId, id).Force(force).Execute()

Delete a Worker



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteClusterWorker(context.Background(), clusterId, id).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteClusterWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterWorker`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteClusterWorker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterWorkerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCluster

> InlineResponse20026 GetCluster(ctx, clusterId).Execute()

Get a Specific Cluster



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetCluster(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterApiConfig

> ClusterApiConfig GetClusterApiConfig(ctx, clusterId).Execute()

Get API Config



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterApiConfig(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterApiConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterApiConfig`: ClusterApiConfig
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterApiConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterApiConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterApiConfig**](clusterApiConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterDatastore

> InlineResponse20027 GetClusterDatastore(ctx, clusterId, id).Execute()

Get a Specific Datastore



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterDatastore(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterDatastore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterDatastore`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterHistory

> map[string]interface{} GetClusterHistory(ctx, clusterId).Execute()

Get Cluster History



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterHistory(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterHistory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterHistoryRequest struct via the builder pattern


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


## GetClusterHistoryDetail

> InlineResponse20028 GetClusterHistoryDetail(ctx, clusterId, id).Execute()

Get Cluster History Details



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterHistoryDetail(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterHistoryDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterHistoryDetail`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterHistoryDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterHistoryDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterHistoryEventDetail

> InlineResponse20029 GetClusterHistoryEventDetail(ctx, clusterId, id).Execute()

Get Cluster History Event



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterHistoryEventDetail(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterHistoryEventDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterHistoryEventDetail`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterHistoryEventDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterHistoryEventDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterMasters

> InlineResponse20030 GetClusterMasters(ctx, clusterId).Phrase(phrase).Execute()

Get Masters (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterMasters(context.Background(), clusterId).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterMasters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterMasters`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterMasters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterMastersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNamespace

> InlineResponse20031 GetClusterNamespace(ctx, clusterId, id).Execute()

Get Namespace (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterNamespace(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNamespace`: InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNamespaces

> map[string]interface{} GetClusterNamespaces(ctx, clusterId).Execute()

List Namespaces (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterNamespaces(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNamespaces`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNamespacesRequest struct via the builder pattern


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


## GetClusterUpgradeVersions

> InlineResponse20032 GetClusterUpgradeVersions(ctx, clusterId).Execute()

Get Cluster Upgrade Versions (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterUpgradeVersions(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterUpgradeVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterUpgradeVersions`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterUpgradeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterUpgradeVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiCluster

> InlineResponse200168 GetWikiCluster(ctx, clusterId).Execute()

Retrieves a Cluster Wiki Page



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetWikiCluster(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetWikiCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiCluster`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetWikiCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiClusterRequest struct via the builder pattern


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


## ListClusterContainers

> map[string]interface{} ListClusterContainers(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()

Get Containers



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    resourceLevel := "resourceLevel_example" // string | Resource level filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterContainers(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterContainers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **resourceLevel** | **string** | Resource level filter | 

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


## ListClusterDatastores

> map[string]interface{} ListClusterDatastores(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Name(name).Code(code).HideInactive(hideInactive).Execute()

Get Datastores



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)
    hideInactive := true // bool | If true restricts query to only load active datastores (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterDatastores(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Name(name).Code(code).HideInactive(hideInactive).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterDatastores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterDatastores`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **hideInactive** | **bool** | If true restricts query to only load active datastores | [default to false]

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


## ListClusterDeployments

> map[string]interface{} ListClusterDeployments(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()

Get Deployments



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    resourceLevel := "resourceLevel_example" // string | Resource level filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterDeployments(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterDeployments`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **resourceLevel** | **string** | Resource level filter | 

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


## ListClusterJobs

> map[string]interface{} ListClusterJobs(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Execute()

Get Jobs



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterJobs(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterJobs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

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


## ListClusterPods

> map[string]interface{} ListClusterPods(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()

Get Pods



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    resourceLevel := "resourceLevel_example" // string | Resource level filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterPods(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterPods`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterPods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **resourceLevel** | **string** | Resource level filter | 

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


## ListClusterServices

> map[string]interface{} ListClusterServices(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Execute()

Get Services



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterServices(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterServices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

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


## ListClusterStatefulSets

> map[string]interface{} ListClusterStatefulSets(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()

Get Stateful Sets



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
    clusterId := int32(5) // int32 | The ID of the cluster
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    resourceLevel := "resourceLevel_example" // string | Resource level filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterStatefulSets(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).ResourceLevel(resourceLevel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterStatefulSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterStatefulSets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterStatefulSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterStatefulSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **resourceLevel** | **string** | Resource level filter | 

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


## ListClusterTypes

> map[string]interface{} ListClusterTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProviderType(providerType).Execute()

Get All Cluster Types



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
    providerType := "providerType_example" // string | Filter by `Provider Type` code.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProviderType(providerType).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusterTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **providerType** | **string** | Filter by &#x60;Provider Type&#x60; code.  | 

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


## ListClusterVolumes

> map[string]interface{} ListClusterVolumes(ctx, clusterId).Execute()

Get Volumes



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterVolumes(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterVolumes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterVolumesRequest struct via the builder pattern


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


## ListClusterWorkers

> map[string]interface{} ListClusterWorkers(ctx, clusterId).Execute()

Get Workers



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
    clusterId := int32(5) // int32 | The ID of the cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterWorkers(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterWorkers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterWorkers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterWorkers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterWorkersRequest struct via the builder pattern


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


## ListClusters

> map[string]interface{} ListClusters(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).ZoneId(zoneId).TypeId(typeId).Labels(labels).AllLabels(allLabels).Execute()

Get All Clusters



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
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    typeId := int64(3) // int64 | Type filter, restricts query to only load clusters of a specified cluster type (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusters(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).ZoneId(zoneId).TypeId(typeId).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **typeId** | **int64** | Type filter, restricts query to only load clusters of a specified cluster type | 
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


## RestartClusterContainer

> SuccessError RestartClusterContainer(ctx, clusterId, id).Execute()

Restart a Container



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.RestartClusterContainer(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RestartClusterContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartClusterContainer`: SuccessError
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.RestartClusterContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartClusterContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuccessError**](successError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartClusterDeployment

> SuccessError RestartClusterDeployment(ctx, clusterId, id).Execute()

Restart a Deployment



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.RestartClusterDeployment(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RestartClusterDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartClusterDeployment`: SuccessError
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.RestartClusterDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartClusterDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuccessError**](successError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartClusterPod

> SuccessError RestartClusterPod(ctx, clusterId, id).Execute()

Restart a Pod



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.RestartClusterPod(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RestartClusterPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartClusterPod`: SuccessError
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.RestartClusterPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartClusterPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuccessError**](successError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartClusterStatefulSet

> SuccessError RestartClusterStatefulSet(ctx, clusterId, id).Execute()

Restart a Stateful Set



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.RestartClusterStatefulSet(context.Background(), clusterId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RestartClusterStatefulSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartClusterStatefulSet`: SuccessError
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.RestartClusterStatefulSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartClusterStatefulSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuccessError**](successError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> map[string]interface{} UpdateCluster(ctx, clusterId).InlineObject53(inlineObject53).Execute()

Update Cluster



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
    clusterId := int32(5) // int32 | The ID of the cluster
    inlineObject53 := *openapiclient.NewInlineObject53() // InlineObject53 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateCluster(context.Background(), clusterId).InlineObject53(inlineObject53).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject53** | [**InlineObject53**](InlineObject53.md) |  | 

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


## UpdateClusterDatastore

> map[string]interface{} UpdateClusterDatastore(ctx, clusterId, id).InlineObject55(inlineObject55).Execute()

Update Datastore



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject55 := *openapiclient.NewInlineObject55() // InlineObject55 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateClusterDatastore(context.Background(), clusterId, id).InlineObject55(inlineObject55).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateClusterDatastore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterDatastore`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateClusterDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject55** | [**InlineObject55**](InlineObject55.md) |  | 

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


## UpdateClusterNamespace

> map[string]interface{} UpdateClusterNamespace(ctx, clusterId, id).InlineObject57(inlineObject57).Execute()

Update Namespace (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject57 := *openapiclient.NewInlineObject57() // InlineObject57 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateClusterNamespace(context.Background(), clusterId, id).InlineObject57(inlineObject57).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateClusterNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterNamespace`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateClusterNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject57** | [**InlineObject57**](InlineObject57.md) |  | 

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


## UpdateClusterPermissions

> map[string]interface{} UpdateClusterPermissions(ctx, clusterId).InlineObject58(inlineObject58).Execute()

Update Cluster Permissions



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
    clusterId := int32(5) // int32 | The ID of the cluster
    inlineObject58 := *openapiclient.NewInlineObject58(*openapiclient.NewClusterUpdatePermissions()) // InlineObject58 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateClusterPermissions(context.Background(), clusterId).InlineObject58(inlineObject58).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateClusterPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterPermissions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateClusterPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject58** | [**InlineObject58**](InlineObject58.md) |  | 

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


## UpdateClusterUpgradeVersions

> Model200Success UpdateClusterUpgradeVersions(ctx, clusterId).TargetVersion(targetVersion).Execute()

Upgrade a Cluster (Kubernetes)



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
    clusterId := int32(5) // int32 | The ID of the cluster
    targetVersion := "1.21.14" // string | Target version for cluster after upgrade

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateClusterUpgradeVersions(context.Background(), clusterId).TargetVersion(targetVersion).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateClusterUpgradeVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterUpgradeVersions`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateClusterUpgradeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterUpgradeVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetVersion** | **string** | Target version for cluster after upgrade | 

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


## UpdateClusterWorkerCount

> Model200Success UpdateClusterWorkerCount(ctx, clusterId).WorkerCount(workerCount).Execute()

Update Worker Count



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
    clusterId := int32(5) // int32 | The ID of the cluster
    workerCount := int64(5) // int64 | The target number of worker nodes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateClusterWorkerCount(context.Background(), clusterId).WorkerCount(workerCount).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateClusterWorkerCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterWorkerCount`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateClusterWorkerCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterWorkerCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerCount** | **int64** | The target number of worker nodes | 

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


## UpdateWikiCluster

> map[string]interface{} UpdateWikiCluster(ctx, clusterId).InlineObject268(inlineObject268).Execute()

Update a Cluster Wiki Page



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
    clusterId := int32(5) // int32 | The ID of the cluster
    inlineObject268 := *openapiclient.NewInlineObject268() // InlineObject268 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateWikiCluster(context.Background(), clusterId).InlineObject268(inlineObject268).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateWikiCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateWikiCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject268** | [**InlineObject268**](InlineObject268.md) |  | 

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

