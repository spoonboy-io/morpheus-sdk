# \ClusterLayoutsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClusterLayoutClone**](ClusterLayoutsApi.md#AddClusterLayoutClone) | **Post** /api/library/cluster-layouts/{id}/clone | Clone a Cluster Layout
[**AddClusterLayouts**](ClusterLayoutsApi.md#AddClusterLayouts) | **Post** /api/library/cluster-layouts | Create a Cluster Layout
[**DeleteClusterLayout**](ClusterLayoutsApi.md#DeleteClusterLayout) | **Delete** /api/library/cluster-layouts/{id} | Delete a Cluster Layout
[**GetClusterLayout**](ClusterLayoutsApi.md#GetClusterLayout) | **Get** /api/library/cluster-layouts/{id} | Get a Specific Cluster Layout
[**ListClusterLayouts**](ClusterLayoutsApi.md#ListClusterLayouts) | **Get** /api/library/cluster-layouts | Get All Cluster Layouts
[**UpdateClusterLayout**](ClusterLayoutsApi.md#UpdateClusterLayout) | **Put** /api/library/cluster-layouts/{id} | Update a Cluster Layout



## AddClusterLayoutClone

> SuccessId AddClusterLayoutClone(ctx, id).Name(name).Description(description).ComputeVersion(computeVersion).Execute()

Clone a Cluster Layout



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
    name := "New Name" // string | Name of cluster layout. Defaults to Copy of <cloned layout name> (optional)
    description := "New Description" // string | Description of cluster layout. Defaults to cloned layout description (optional)
    computeVersion := "2.0" // string | Version of cluster layout. Defaults to cloned layout version (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLayoutsApi.AddClusterLayoutClone(context.Background(), id).Name(name).Description(description).ComputeVersion(computeVersion).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsApi.AddClusterLayoutClone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClusterLayoutClone`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsApi.AddClusterLayoutClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterLayoutCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of cluster layout. Defaults to Copy of &lt;cloned layout name&gt; | 
 **description** | **string** | Description of cluster layout. Defaults to cloned layout description | 
 **computeVersion** | **string** | Version of cluster layout. Defaults to cloned layout version | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClusterLayouts

> SuccessId AddClusterLayouts(ctx).InlineObject50(inlineObject50).Execute()

Create a Cluster Layout



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
    inlineObject50 := *openapiclient.NewInlineObject50() // InlineObject50 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLayoutsApi.AddClusterLayouts(context.Background()).InlineObject50(inlineObject50).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsApi.AddClusterLayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClusterLayouts`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsApi.AddClusterLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject50** | [**InlineObject50**](InlineObject50.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterLayout

> Model200Success DeleteClusterLayout(ctx, id).Execute()

Delete a Cluster Layout



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
    resp, r, err := api_client.ClusterLayoutsApi.DeleteClusterLayout(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsApi.DeleteClusterLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterLayout`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsApi.DeleteClusterLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterLayoutRequest struct via the builder pattern


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


## GetClusterLayout

> InlineResponse20025 GetClusterLayout(ctx, id).Execute()

Get a Specific Cluster Layout



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
    resp, r, err := api_client.ClusterLayoutsApi.GetClusterLayout(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsApi.GetClusterLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterLayout`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsApi.GetClusterLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterLayouts

> map[string]interface{} ListClusterLayouts(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Cluster Layouts



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
    provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLayoutsApi.ListClusterLayouts(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsApi.ListClusterLayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterLayouts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsApi.ListClusterLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusterLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 
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


## UpdateClusterLayout

> SuccessId UpdateClusterLayout(ctx, id).InlineObject51(inlineObject51).Execute()

Update a Cluster Layout



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
    inlineObject51 := *openapiclient.NewInlineObject51() // InlineObject51 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterLayoutsApi.UpdateClusterLayout(context.Background(), id).InlineObject51(inlineObject51).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsApi.UpdateClusterLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterLayout`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsApi.UpdateClusterLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject51** | [**InlineObject51**](InlineObject51.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

