# \LoadBalancersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancer**](LoadBalancersApi.md#CreateLoadBalancer) | **Post** /api/load-balancers | Create a Load Balancer
[**CreateLoadBalancerMonitor**](LoadBalancersApi.md#CreateLoadBalancerMonitor) | **Post** /api/load-balancers/{loadBalancerId}/monitors | Create a Load Balancer Monitor
[**CreateLoadBalancerPool**](LoadBalancersApi.md#CreateLoadBalancerPool) | **Post** /api/load-balancers/{loadBalancerId}/pools | Create a Load Balancer Pool
[**CreateLoadBalancerPoolNode**](LoadBalancersApi.md#CreateLoadBalancerPoolNode) | **Post** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Create a Load Balancer Pool Node
[**CreateLoadBalancerProfile**](LoadBalancersApi.md#CreateLoadBalancerProfile) | **Post** /api/load-balancers/{loadBalancerId}/profiles | Create a Load Balancer Profile
[**CreateLoadBalancerVirtualServer**](LoadBalancersApi.md#CreateLoadBalancerVirtualServer) | **Post** /api/load-balancers/{loadBalancerId}/virtual-servers | Create a Load Balancer Virtual Server
[**DeleteLoadBalancer**](LoadBalancersApi.md#DeleteLoadBalancer) | **Delete** /api/load-balancers/{loadBalancerId} | Delete a Load Balancer
[**DeleteLoadBalancerMonitor**](LoadBalancersApi.md#DeleteLoadBalancerMonitor) | **Delete** /api/load-balancers/{loadBalancerId}/monitors/{id} | Delete a Load Balancer Monitor
[**DeleteLoadBalancerPool**](LoadBalancersApi.md#DeleteLoadBalancerPool) | **Delete** /api/load-balancers/{loadBalancerId}/pools/{id} | Delete a Load Balancer Pool
[**DeleteLoadBalancerPoolNode**](LoadBalancersApi.md#DeleteLoadBalancerPoolNode) | **Delete** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Delete a Load Balancer Pool Node
[**DeleteLoadBalancerProfile**](LoadBalancersApi.md#DeleteLoadBalancerProfile) | **Delete** /api/load-balancers/{loadBalancerId}/profiles/{id} | Delete a Load Balancer Profile
[**DeleteLoadBalancerVirtualServer**](LoadBalancersApi.md#DeleteLoadBalancerVirtualServer) | **Delete** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Delete a Load Balancer Virtual Server
[**GetLoadBalancer**](LoadBalancersApi.md#GetLoadBalancer) | **Get** /api/load-balancers/{loadBalancerId} | Get a Specific Load Balancer
[**GetLoadBalancerMonitor**](LoadBalancersApi.md#GetLoadBalancerMonitor) | **Get** /api/load-balancers/{loadBalancerId}/monitors/{id} | Get a Specific Load Balancer Monitor
[**GetLoadBalancerPool**](LoadBalancersApi.md#GetLoadBalancerPool) | **Get** /api/load-balancers/{loadBalancerId}/pools/{id} | Get a Specific Load Balancer Pool
[**GetLoadBalancerPoolNode**](LoadBalancersApi.md#GetLoadBalancerPoolNode) | **Get** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Get a Specific Load Balancer Pool Node
[**GetLoadBalancerProfile**](LoadBalancersApi.md#GetLoadBalancerProfile) | **Get** /api/load-balancers/{loadBalancerId}/profiles/{id} | Get a Specific Load Balancer Profile
[**GetLoadBalancerType**](LoadBalancersApi.md#GetLoadBalancerType) | **Get** /api/load-balancer-types/{id} | Get a Specific Load Balancer Type
[**GetLoadBalancerVirtualServer**](LoadBalancersApi.md#GetLoadBalancerVirtualServer) | **Get** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Get a Specific Load Balancer Virtual Server
[**ListLoadBalancerMonitors**](LoadBalancersApi.md#ListLoadBalancerMonitors) | **Get** /api/load-balancers/{loadBalancerId}/monitors | Get All Load Balancer Monitors For Load Balancer
[**ListLoadBalancerPoolNodes**](LoadBalancersApi.md#ListLoadBalancerPoolNodes) | **Get** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Get All Load Balancer Pool Nodes For Load Balancer Pool
[**ListLoadBalancerPools**](LoadBalancersApi.md#ListLoadBalancerPools) | **Get** /api/load-balancers/{loadBalancerId}/pools | Get All Load Balancer Pools For Load Balancer
[**ListLoadBalancerProfiles**](LoadBalancersApi.md#ListLoadBalancerProfiles) | **Get** /api/load-balancers/{loadBalancerId}/profiles | Get All Load Balancer Profiles For Load Balancer
[**ListLoadBalancerTypes**](LoadBalancersApi.md#ListLoadBalancerTypes) | **Get** /api/load-balancer-types | Get All Load Balancer Types
[**ListLoadBalancerVirtualServers**](LoadBalancersApi.md#ListLoadBalancerVirtualServers) | **Get** /api/load-balancers/{loadBalancerId}/virtual-servers | Get All Load Balancer Virtual Servers For Load Balancer
[**ListLoadBalancers**](LoadBalancersApi.md#ListLoadBalancers) | **Get** /api/load-balancers | Get All Load Balancers
[**RefreshLoadBalancer**](LoadBalancersApi.md#RefreshLoadBalancer) | **Put** /api/load-balancers/{loadBalancerId}/refresh | Refresh a Load Balancer
[**UpdateLoadBalancer**](LoadBalancersApi.md#UpdateLoadBalancer) | **Put** /api/load-balancers/{loadBalancerId} | Update a Load Balancer
[**UpdateLoadBalancerMonitor**](LoadBalancersApi.md#UpdateLoadBalancerMonitor) | **Put** /api/load-balancers/{loadBalancerId}/monitors/{id} | Update a Load Balancer Monitor
[**UpdateLoadBalancerPool**](LoadBalancersApi.md#UpdateLoadBalancerPool) | **Put** /api/load-balancers/{loadBalancerId}/pools/{id} | Update a Load Balancer Pool
[**UpdateLoadBalancerPoolNode**](LoadBalancersApi.md#UpdateLoadBalancerPoolNode) | **Put** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Update a Load Balancer Pool Node
[**UpdateLoadBalancerProfile**](LoadBalancersApi.md#UpdateLoadBalancerProfile) | **Put** /api/load-balancers/{loadBalancerId}/profiles/{id} | Update a Load Balancer Profile
[**UpdateLoadBalancerVirtualServer**](LoadBalancersApi.md#UpdateLoadBalancerVirtualServer) | **Put** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Update a Load Balancer Virtual Server



## CreateLoadBalancer

> InlineResponse20078 CreateLoadBalancer(ctx).InlineObject127(inlineObject127).Execute()

Create a Load Balancer



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
    inlineObject127 := *openapiclient.NewInlineObject127() // InlineObject127 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.CreateLoadBalancer(context.Background()).InlineObject127(inlineObject127).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.CreateLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancer`: InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.CreateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject127** | [**InlineObject127**](InlineObject127.md) |  | 

### Return type

[**InlineResponse20078**](inline_response_200_78.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancerMonitor

> map[string]interface{} CreateLoadBalancerMonitor(ctx, loadBalancerId).InlineObject129(inlineObject129).Execute()

Create a Load Balancer Monitor



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    inlineObject129 := *openapiclient.NewInlineObject129() // InlineObject129 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.CreateLoadBalancerMonitor(context.Background(), loadBalancerId).InlineObject129(inlineObject129).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.CreateLoadBalancerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerMonitor`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.CreateLoadBalancerMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject129** | [**InlineObject129**](InlineObject129.md) |  | 

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


## CreateLoadBalancerPool

> map[string]interface{} CreateLoadBalancerPool(ctx, loadBalancerId).InlineObject131(inlineObject131).Execute()

Create a Load Balancer Pool



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    inlineObject131 := *openapiclient.NewInlineObject131() // InlineObject131 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.CreateLoadBalancerPool(context.Background(), loadBalancerId).InlineObject131(inlineObject131).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.CreateLoadBalancerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerPool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.CreateLoadBalancerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject131** | [**InlineObject131**](InlineObject131.md) |  | 

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


## CreateLoadBalancerPoolNode

> map[string]interface{} CreateLoadBalancerPoolNode(ctx, loadBalancerPoolId).InlineObject137(inlineObject137).Execute()

Create a Load Balancer Pool Node



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
    loadBalancerPoolId := float32(4) // float32 | Load Balancer Pool ID
    inlineObject137 := *openapiclient.NewInlineObject137() // InlineObject137 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.CreateLoadBalancerPoolNode(context.Background(), loadBalancerPoolId).InlineObject137(inlineObject137).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.CreateLoadBalancerPoolNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerPoolNode`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.CreateLoadBalancerPoolNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerPoolId** | **float32** | Load Balancer Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerPoolNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject137** | [**InlineObject137**](InlineObject137.md) |  | 

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


## CreateLoadBalancerProfile

> map[string]interface{} CreateLoadBalancerProfile(ctx, loadBalancerId).InlineObject133(inlineObject133).Execute()

Create a Load Balancer Profile



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    inlineObject133 := *openapiclient.NewInlineObject133() // InlineObject133 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.CreateLoadBalancerProfile(context.Background(), loadBalancerId).InlineObject133(inlineObject133).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.CreateLoadBalancerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.CreateLoadBalancerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject133** | [**InlineObject133**](InlineObject133.md) |  | 

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


## CreateLoadBalancerVirtualServer

> InlineResponse20082 CreateLoadBalancerVirtualServer(ctx, loadBalancerId).InlineObject135(inlineObject135).Execute()

Create a Load Balancer Virtual Server



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    inlineObject135 := *openapiclient.NewInlineObject135() // InlineObject135 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.CreateLoadBalancerVirtualServer(context.Background(), loadBalancerId).InlineObject135(inlineObject135).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.CreateLoadBalancerVirtualServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancerVirtualServer`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.CreateLoadBalancerVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject135** | [**InlineObject135**](InlineObject135.md) |  | 

### Return type

[**InlineResponse20082**](inline_response_200_82.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancer

> Model200Success DeleteLoadBalancer(ctx, loadBalancerId).Execute()

Delete a Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DeleteLoadBalancer(context.Background(), loadBalancerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DeleteLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DeleteLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerRequest struct via the builder pattern


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


## DeleteLoadBalancerMonitor

> Model200Success DeleteLoadBalancerMonitor(ctx, loadBalancerId, id).Execute()

Delete a Load Balancer Monitor



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DeleteLoadBalancerMonitor(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DeleteLoadBalancerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerMonitor`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DeleteLoadBalancerMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerMonitorRequest struct via the builder pattern


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


## DeleteLoadBalancerPool

> Model200Success DeleteLoadBalancerPool(ctx, loadBalancerId, id).Execute()

Delete a Load Balancer Pool



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DeleteLoadBalancerPool(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DeleteLoadBalancerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerPool`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DeleteLoadBalancerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerPoolRequest struct via the builder pattern


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


## DeleteLoadBalancerPoolNode

> Model200Success DeleteLoadBalancerPoolNode(ctx, loadBalancerPoolId, id).Execute()

Delete a Load Balancer Pool Node



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
    loadBalancerPoolId := float32(4) // float32 | Load Balancer Pool ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DeleteLoadBalancerPoolNode(context.Background(), loadBalancerPoolId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DeleteLoadBalancerPoolNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerPoolNode`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DeleteLoadBalancerPoolNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerPoolId** | **float32** | Load Balancer Pool ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerPoolNodeRequest struct via the builder pattern


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


## DeleteLoadBalancerProfile

> Model200Success DeleteLoadBalancerProfile(ctx, loadBalancerId, id).Execute()

Delete a Load Balancer Profile



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DeleteLoadBalancerProfile(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DeleteLoadBalancerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerProfile`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DeleteLoadBalancerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerProfileRequest struct via the builder pattern


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


## DeleteLoadBalancerVirtualServer

> Model200Success DeleteLoadBalancerVirtualServer(ctx, loadBalancerId, id).Execute()

Delete a Load Balancer Virtual Server



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DeleteLoadBalancerVirtualServer(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DeleteLoadBalancerVirtualServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoadBalancerVirtualServer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DeleteLoadBalancerVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerVirtualServerRequest struct via the builder pattern


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


## GetLoadBalancer

> InlineResponse20078 GetLoadBalancer(ctx, loadBalancerId).Execute()

Get a Specific Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancer(context.Background(), loadBalancerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancer`: InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20078**](inline_response_200_78.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerMonitor

> InlineResponse20079 GetLoadBalancerMonitor(ctx, loadBalancerId, id).Execute()

Get a Specific Load Balancer Monitor



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancerMonitor(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerMonitor`: InlineResponse20079
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancerMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20079**](inline_response_200_79.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerPool

> InlineResponse20080 GetLoadBalancerPool(ctx, loadBalancerId, id).Execute()

Get a Specific Load Balancer Pool



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancerPool(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerPool`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20080**](inline_response_200_80.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerPoolNode

> InlineResponse20083 GetLoadBalancerPoolNode(ctx, loadBalancerPoolId, id).Execute()

Get a Specific Load Balancer Pool Node



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
    loadBalancerPoolId := float32(4) // float32 | Load Balancer Pool ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancerPoolNode(context.Background(), loadBalancerPoolId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancerPoolNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerPoolNode`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancerPoolNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerPoolId** | **float32** | Load Balancer Pool ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerPoolNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20083**](inline_response_200_83.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerProfile

> InlineResponse20081 GetLoadBalancerProfile(ctx, loadBalancerId, id).Execute()

Get a Specific Load Balancer Profile



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancerProfile(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerProfile`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20081**](inline_response_200_81.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerType

> InlineResponse20077 GetLoadBalancerType(ctx, id).Execute()

Get a Specific Load Balancer Type



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
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancerType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancerType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerType`: InlineResponse20077
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20077**](inline_response_200_77.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerVirtualServer

> InlineResponse20082 GetLoadBalancerVirtualServer(ctx, loadBalancerId, id).Execute()

Get a Specific Load Balancer Virtual Server



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.GetLoadBalancerVirtualServer(context.Background(), loadBalancerId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.GetLoadBalancerVirtualServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerVirtualServer`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.GetLoadBalancerVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20082**](inline_response_200_82.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoadBalancerMonitors

> map[string]interface{} ListLoadBalancerMonitors(ctx, loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()

Get All Load Balancer Monitors For Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancerMonitors(context.Background(), loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancerMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancerMonitors`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancerMonitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## ListLoadBalancerPoolNodes

> map[string]interface{} ListLoadBalancerPoolNodes(ctx, loadBalancerPoolId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get All Load Balancer Pool Nodes For Load Balancer Pool



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
    loadBalancerPoolId := float32(4) // float32 | Load Balancer Pool ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancerPoolNodes(context.Background(), loadBalancerPoolId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancerPoolNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancerPoolNodes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancerPoolNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerPoolId** | **float32** | Load Balancer Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerPoolNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
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


## ListLoadBalancerPools

> map[string]interface{} ListLoadBalancerPools(ctx, loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()

Get All Load Balancer Pools For Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancerPools(context.Background(), loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancerPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancerPools`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancerPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## ListLoadBalancerProfiles

> map[string]interface{} ListLoadBalancerProfiles(ctx, loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()

Get All Load Balancer Profiles For Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancerProfiles(context.Background(), loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancerProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancerProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancerProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## ListLoadBalancerTypes

> map[string]interface{} ListLoadBalancerTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).OptionTypes(optionTypes).Phrase(phrase).Name(name).Code(code).Execute()

Get All Load Balancer Types



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
    optionTypes := true // bool | Pass true to include optionTypes in the response for each entry. (optional) (default to false)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancerTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).OptionTypes(optionTypes).Phrase(phrase).Name(name).Code(code).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancerTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancerTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **optionTypes** | **bool** | Pass true to include optionTypes in the response for each entry. | [default to false]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 

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


## ListLoadBalancerVirtualServers

> map[string]interface{} ListLoadBalancerVirtualServers(ctx, loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).VipName(vipName).VipAddress(vipAddress).VipHostname(vipHostname).Execute()

Get All Load Balancer Virtual Servers For Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    vipName := "lb-114" // string | If specified will return an exact match on vipName (optional)
    vipAddress := "192.168.123.44" // string | If specified will return an exact match on vipAddress (optional)
    vipHostname := "mylb" // string | If specified will return an exact match on vipHostname (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancerVirtualServers(context.Background(), loadBalancerId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).VipName(vipName).VipAddress(vipAddress).VipHostname(vipHostname).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancerVirtualServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancerVirtualServers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancerVirtualServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerVirtualServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **vipName** | **string** | If specified will return an exact match on vipName | 
 **vipAddress** | **string** | If specified will return an exact match on vipAddress | 
 **vipHostname** | **string** | If specified will return an exact match on vipHostname | 

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


## ListLoadBalancers

> map[string]interface{} ListLoadBalancers(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()

Get All Load Balancers



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
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.ListLoadBalancers(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.ListLoadBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.ListLoadBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## RefreshLoadBalancer

> map[string]interface{} RefreshLoadBalancer(ctx, loadBalancerId).Execute()

Refresh a Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.RefreshLoadBalancer(context.Background(), loadBalancerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.RefreshLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshLoadBalancer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.RefreshLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshLoadBalancerRequest struct via the builder pattern


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


## UpdateLoadBalancer

> InlineResponse20078 UpdateLoadBalancer(ctx, loadBalancerId).InlineObject128(inlineObject128).Execute()

Update a Load Balancer



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    inlineObject128 := *openapiclient.NewInlineObject128() // InlineObject128 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.UpdateLoadBalancer(context.Background(), loadBalancerId).InlineObject128(inlineObject128).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.UpdateLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancer`: InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.UpdateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject128** | [**InlineObject128**](InlineObject128.md) |  | 

### Return type

[**InlineResponse20078**](inline_response_200_78.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancerMonitor

> map[string]interface{} UpdateLoadBalancerMonitor(ctx, loadBalancerId, id).InlineObject130(inlineObject130).Execute()

Update a Load Balancer Monitor



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject130 := *openapiclient.NewInlineObject130() // InlineObject130 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.UpdateLoadBalancerMonitor(context.Background(), loadBalancerId, id).InlineObject130(inlineObject130).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.UpdateLoadBalancerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancerMonitor`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.UpdateLoadBalancerMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject130** | [**InlineObject130**](InlineObject130.md) |  | 

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


## UpdateLoadBalancerPool

> map[string]interface{} UpdateLoadBalancerPool(ctx, loadBalancerId, id).InlineObject132(inlineObject132).Execute()

Update a Load Balancer Pool



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject132 := *openapiclient.NewInlineObject132() // InlineObject132 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.UpdateLoadBalancerPool(context.Background(), loadBalancerId, id).InlineObject132(inlineObject132).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.UpdateLoadBalancerPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancerPool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.UpdateLoadBalancerPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject132** | [**InlineObject132**](InlineObject132.md) |  | 

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


## UpdateLoadBalancerPoolNode

> map[string]interface{} UpdateLoadBalancerPoolNode(ctx, loadBalancerPoolId, id).InlineObject138(inlineObject138).Execute()

Update a Load Balancer Pool Node



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
    loadBalancerPoolId := float32(4) // float32 | Load Balancer Pool ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject138 := *openapiclient.NewInlineObject138() // InlineObject138 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.UpdateLoadBalancerPoolNode(context.Background(), loadBalancerPoolId, id).InlineObject138(inlineObject138).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.UpdateLoadBalancerPoolNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancerPoolNode`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.UpdateLoadBalancerPoolNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerPoolId** | **float32** | Load Balancer Pool ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerPoolNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject138** | [**InlineObject138**](InlineObject138.md) |  | 

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


## UpdateLoadBalancerProfile

> map[string]interface{} UpdateLoadBalancerProfile(ctx, loadBalancerId, id).InlineObject134(inlineObject134).Execute()

Update a Load Balancer Profile



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject134 := *openapiclient.NewInlineObject134() // InlineObject134 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.UpdateLoadBalancerProfile(context.Background(), loadBalancerId, id).InlineObject134(inlineObject134).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.UpdateLoadBalancerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancerProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.UpdateLoadBalancerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject134** | [**InlineObject134**](InlineObject134.md) |  | 

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


## UpdateLoadBalancerVirtualServer

> InlineResponse20082 UpdateLoadBalancerVirtualServer(ctx, loadBalancerId, id).InlineObject136(inlineObject136).Execute()

Update a Load Balancer Virtual Server



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
    loadBalancerId := float32(4) // float32 | Load Balancer ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject136 := *openapiclient.NewInlineObject136() // InlineObject136 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.UpdateLoadBalancerVirtualServer(context.Background(), loadBalancerId, id).InlineObject136(inlineObject136).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.UpdateLoadBalancerVirtualServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancerVirtualServer`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.UpdateLoadBalancerVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **float32** | Load Balancer ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject136** | [**InlineObject136**](InlineObject136.md) |  | 

### Return type

[**InlineResponse20082**](inline_response_200_82.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

