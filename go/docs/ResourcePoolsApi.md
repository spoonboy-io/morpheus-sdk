# \ResourcePoolsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourcePoolGroup**](ResourcePoolsApi.md#CreateResourcePoolGroup) | **Post** /api/resource-pools/groups | Create a Resource Pool Group
[**DeleteResourcePoolGroup**](ResourcePoolsApi.md#DeleteResourcePoolGroup) | **Delete** /api/resource-pools/groups/{id} | Delete a Resource Pool Group
[**GetResourcePoolGroups**](ResourcePoolsApi.md#GetResourcePoolGroups) | **Get** /api/resource-pools/groups | Get all Resource Pool Groups
[**GetresourcePoolGroup**](ResourcePoolsApi.md#GetresourcePoolGroup) | **Get** /api/resource-pools/groups/{id} | Get a Specific Resource Pool Group
[**UpdateResourcePoolGroup**](ResourcePoolsApi.md#UpdateResourcePoolGroup) | **Put** /api/resource-pools/groups/{id} | Update a Resource Pool Group



## CreateResourcePoolGroup

> InlineResponse200132 CreateResourcePoolGroup(ctx).InlineObject206(inlineObject206).Execute()

Create a Resource Pool Group



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
    inlineObject206 := *openapiclient.NewInlineObject206() // InlineObject206 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsApi.CreateResourcePoolGroup(context.Background()).InlineObject206(inlineObject206).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsApi.CreateResourcePoolGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourcePoolGroup`: InlineResponse200132
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsApi.CreateResourcePoolGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourcePoolGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject206** | [**InlineObject206**](InlineObject206.md) |  | 

### Return type

[**InlineResponse200132**](inline_response_200_132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourcePoolGroup

> Model200Success DeleteResourcePoolGroup(ctx, id).Execute()

Delete a Resource Pool Group



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
    resp, r, err := api_client.ResourcePoolsApi.DeleteResourcePoolGroup(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsApi.DeleteResourcePoolGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteResourcePoolGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsApi.DeleteResourcePoolGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourcePoolGroupRequest struct via the builder pattern


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


## GetResourcePoolGroups

> InlineResponse200131 GetResourcePoolGroups(ctx).Execute()

Get all Resource Pool Groups



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
    resp, r, err := api_client.ResourcePoolsApi.GetResourcePoolGroups(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsApi.GetResourcePoolGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourcePoolGroups`: InlineResponse200131
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsApi.GetResourcePoolGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePoolGroupsRequest struct via the builder pattern


### Return type

[**InlineResponse200131**](inline_response_200_131.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetresourcePoolGroup

> InlineResponse200132 GetresourcePoolGroup(ctx, id).Execute()

Get a Specific Resource Pool Group



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
    resp, r, err := api_client.ResourcePoolsApi.GetresourcePoolGroup(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsApi.GetresourcePoolGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetresourcePoolGroup`: InlineResponse200132
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsApi.GetresourcePoolGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetresourcePoolGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200132**](inline_response_200_132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourcePoolGroup

> InlineResponse200132 UpdateResourcePoolGroup(ctx, id).InlineObject207(inlineObject207).Execute()

Update a Resource Pool Group



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
    inlineObject207 := *openapiclient.NewInlineObject207() // InlineObject207 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcePoolsApi.UpdateResourcePoolGroup(context.Background(), id).InlineObject207(inlineObject207).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolsApi.UpdateResourcePoolGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourcePoolGroup`: InlineResponse200132
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolsApi.UpdateResourcePoolGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourcePoolGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject207** | [**InlineObject207**](InlineObject207.md) |  | 

### Return type

[**InlineResponse200132**](inline_response_200_132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

