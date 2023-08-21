# \VDIApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVDIApps**](VDIApi.md#AddVDIApps) | **Post** /api/vdi-apps | Creates a VDI App
[**AddVDIGateways**](VDIApi.md#AddVDIGateways) | **Post** /api/vdi-gateways | Creates a VDI Gateway
[**AddVDIPools**](VDIApi.md#AddVDIPools) | **Post** /api/vdi-pools | Creates a VDI Pool
[**AddVdiAllocation**](VDIApi.md#AddVdiAllocation) | **Post** /api/vdi/{id}/allocate | Allocate Virtual Desktop
[**GetVDIAllocations**](VDIApi.md#GetVDIAllocations) | **Get** /api/vdi-allocations/{id} | Retrieves a Specific VDI Allocation
[**GetVDIApps**](VDIApi.md#GetVDIApps) | **Get** /api/vdi-apps/{id} | Retrieves a Specific VDI App
[**GetVDIGateways**](VDIApi.md#GetVDIGateways) | **Get** /api/vdi-gateways/{id} | Retrieves a Specific VDI Gateway
[**GetVDIPools**](VDIApi.md#GetVDIPools) | **Get** /api/vdi-pools/{id} | Retrieves a Specific VDI Pool
[**GetVdi**](VDIApi.md#GetVdi) | **Get** /api/vdi/{id} | Get a Specific Virtual Desktop
[**ListVDIAllocations**](VDIApi.md#ListVDIAllocations) | **Get** /api/vdi-allocations | Retrieves all VDI Allocations
[**ListVDIApps**](VDIApi.md#ListVDIApps) | **Get** /api/vdi-apps | Retrieves all VDI Apps
[**ListVDIGateways**](VDIApi.md#ListVDIGateways) | **Get** /api/vdi-gateways | Retrieves all VDI Gateways
[**ListVDIPools**](VDIApi.md#ListVDIPools) | **Get** /api/vdi-pools | Retrieves all VDI Pools
[**ListVdi**](VDIApi.md#ListVdi) | **Get** /api/vdi | List Virtual Desktops
[**RemoveVDIApps**](VDIApi.md#RemoveVDIApps) | **Delete** /api/vdi-apps/{id} | Deletes a VDI App
[**RemoveVDIGateways**](VDIApi.md#RemoveVDIGateways) | **Delete** /api/vdi-gateways/{id} | Deletes a VDI Gateway
[**RemoveVDIPools**](VDIApi.md#RemoveVDIPools) | **Delete** /api/vdi-pools/{id} | Deletes a VDI Pool
[**UpdateVDIApps**](VDIApi.md#UpdateVDIApps) | **Put** /api/vdi-apps/{id} | Updates a VDI App Configuration or Icon
[**UpdateVDIGateways**](VDIApi.md#UpdateVDIGateways) | **Put** /api/vdi-gateways/{id} | Updates a VDI Gateway Configuration
[**UpdateVDIPools**](VDIApi.md#UpdateVDIPools) | **Put** /api/vdi-pools/{id} | Updates a VDI Pool Configuration or Icon



## AddVDIApps

> AnyOfobject200Success AddVDIApps(ctx).InlineObject257(inlineObject257).Execute()

Creates a VDI App



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
    inlineObject257 := *openapiclient.NewInlineObject257("TODO") // InlineObject257 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.AddVDIApps(context.Background()).InlineObject257(inlineObject257).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.AddVDIApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVDIApps`: AnyOfobject200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.AddVDIApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVDIAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject257** | [**InlineObject257**](InlineObject257.md) |  | 

### Return type

[**AnyOfobject200Success**](anyOf&lt;object,200-success&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVDIGateways

> AnyOfobject200Success AddVDIGateways(ctx).InlineObject259(inlineObject259).Execute()

Creates a VDI Gateway



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
    inlineObject259 := *openapiclient.NewInlineObject259("TODO") // InlineObject259 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.AddVDIGateways(context.Background()).InlineObject259(inlineObject259).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.AddVDIGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVDIGateways`: AnyOfobject200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.AddVDIGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVDIGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject259** | [**InlineObject259**](InlineObject259.md) |  | 

### Return type

[**AnyOfobject200Success**](anyOf&lt;object,200-success&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVDIPools

> AnyOfobject200Success AddVDIPools(ctx).InlineObject261(inlineObject261).Execute()

Creates a VDI Pool



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
    inlineObject261 := *openapiclient.NewInlineObject261("TODO") // InlineObject261 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.AddVDIPools(context.Background()).InlineObject261(inlineObject261).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.AddVDIPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVDIPools`: AnyOfobject200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.AddVDIPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVDIPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject261** | [**InlineObject261**](InlineObject261.md) |  | 

### Return type

[**AnyOfobject200Success**](anyOf&lt;object,200-success&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVdiAllocation

> map[string]interface{} AddVdiAllocation(ctx, id).Execute()

Allocate Virtual Desktop



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
    resp, r, err := api_client.VDIApi.AddVdiAllocation(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.AddVdiAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVdiAllocation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.AddVdiAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVdiAllocationRequest struct via the builder pattern


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


## GetVDIAllocations

> InlineResponse200163 GetVDIAllocations(ctx, id).Execute()

Retrieves a Specific VDI Allocation



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
    resp, r, err := api_client.VDIApi.GetVDIAllocations(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.GetVDIAllocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVDIAllocations`: InlineResponse200163
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.GetVDIAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVDIAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200163**](inline_response_200_163.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVDIApps

> InlineResponse200160 GetVDIApps(ctx, id).Execute()

Retrieves a Specific VDI App



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
    resp, r, err := api_client.VDIApi.GetVDIApps(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.GetVDIApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVDIApps`: InlineResponse200160
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.GetVDIApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVDIAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200160**](inline_response_200_160.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVDIGateways

> InlineResponse200161 GetVDIGateways(ctx, id).Execute()

Retrieves a Specific VDI Gateway



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
    resp, r, err := api_client.VDIApi.GetVDIGateways(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.GetVDIGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVDIGateways`: InlineResponse200161
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.GetVDIGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVDIGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200161**](inline_response_200_161.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVDIPools

> InlineResponse200162 GetVDIPools(ctx, id).Execute()

Retrieves a Specific VDI Pool



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
    resp, r, err := api_client.VDIApi.GetVDIPools(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.GetVDIPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVDIPools`: InlineResponse200162
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.GetVDIPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVDIPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200162**](inline_response_200_162.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdi

> InlineResponse200164 GetVdi(ctx, id).Execute()

Get a Specific Virtual Desktop



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
    resp, r, err := api_client.VDIApi.GetVdi(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.GetVdi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdi`: InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.GetVdi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200164**](inline_response_200_164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVDIAllocations

> map[string]interface{} ListVDIAllocations(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Id(id).Status(status).PoolId(poolId).UserId(userId).Execute()

Retrieves all VDI Allocations



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
    id := "preparing" // string | Filter by allocation ID (optional)
    status := "preparing" // string | Filter by allocation status (optional)
    poolId := int64(1) // int64 | Filter by `VDI Pool` ID (optional)
    userId := int64(6) // int64 | Filter by User ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.ListVDIAllocations(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Id(id).Status(status).PoolId(poolId).UserId(userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.ListVDIAllocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVDIAllocations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.ListVDIAllocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVDIAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **id** | **string** | Filter by allocation ID | 
 **status** | **string** | Filter by allocation status | 
 **poolId** | **int64** | Filter by &#x60;VDI Pool&#x60; ID | 
 **userId** | **int64** | Filter by User ID | 

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


## ListVDIApps

> map[string]interface{} ListVDIApps(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Execute()

Retrieves all VDI Apps



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
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.ListVDIApps(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.ListVDIApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVDIApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.ListVDIApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVDIAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 

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


## ListVDIGateways

> map[string]interface{} ListVDIGateways(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Execute()

Retrieves all VDI Gateways



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
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.ListVDIGateways(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.ListVDIGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVDIGateways`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.ListVDIGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVDIGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 

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


## ListVDIPools

> map[string]interface{} ListVDIPools(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Enabled(enabled).Execute()

Retrieves all VDI Pools



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
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    enabled := false // bool | Filter by enabled (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.ListVDIPools(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Enabled(enabled).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.ListVDIPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVDIPools`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.ListVDIPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVDIPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 
 **enabled** | **bool** | Filter by enabled | 

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


## ListVdi

> map[string]interface{} ListVdi(ctx).Phrase(phrase).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Description(description).Execute()

List Virtual Desktops



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
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.ListVdi(context.Background()).Phrase(phrase).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Description(description).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.ListVdi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVdi`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.ListVdi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVdiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 

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


## RemoveVDIApps

> Model200Success RemoveVDIApps(ctx, id).Execute()

Deletes a VDI App



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
    resp, r, err := api_client.VDIApi.RemoveVDIApps(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.RemoveVDIApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVDIApps`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.RemoveVDIApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVDIAppsRequest struct via the builder pattern


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


## RemoveVDIGateways

> Model200Success RemoveVDIGateways(ctx, id).Execute()

Deletes a VDI Gateway



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
    resp, r, err := api_client.VDIApi.RemoveVDIGateways(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.RemoveVDIGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVDIGateways`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.RemoveVDIGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVDIGatewaysRequest struct via the builder pattern


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


## RemoveVDIPools

> Model200Success RemoveVDIPools(ctx, id).Execute()

Deletes a VDI Pool



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
    resp, r, err := api_client.VDIApi.RemoveVDIPools(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.RemoveVDIPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVDIPools`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.RemoveVDIPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVDIPoolsRequest struct via the builder pattern


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


## UpdateVDIApps

> AnyOfobject200Success UpdateVDIApps(ctx, id).InlineObject258(inlineObject258).Execute()

Updates a VDI App Configuration or Icon



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
    inlineObject258 := *openapiclient.NewInlineObject258("TODO") // InlineObject258 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.UpdateVDIApps(context.Background(), id).InlineObject258(inlineObject258).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.UpdateVDIApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVDIApps`: AnyOfobject200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.UpdateVDIApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVDIAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject258** | [**InlineObject258**](InlineObject258.md) |  | 

### Return type

[**AnyOfobject200Success**](anyOf&lt;object,200-success&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVDIGateways

> AnyOfobject200Success UpdateVDIGateways(ctx, id).InlineObject260(inlineObject260).Execute()

Updates a VDI Gateway Configuration



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
    inlineObject260 := *openapiclient.NewInlineObject260("TODO") // InlineObject260 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.UpdateVDIGateways(context.Background(), id).InlineObject260(inlineObject260).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.UpdateVDIGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVDIGateways`: AnyOfobject200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.UpdateVDIGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVDIGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject260** | [**InlineObject260**](InlineObject260.md) |  | 

### Return type

[**AnyOfobject200Success**](anyOf&lt;object,200-success&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVDIPools

> AnyOfobject200Success UpdateVDIPools(ctx, id).InlineObject262(inlineObject262).Execute()

Updates a VDI Pool Configuration or Icon



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
    inlineObject262 := *openapiclient.NewInlineObject262(*openapiclient.NewApiVdiPoolsIdVdiPool()) // InlineObject262 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VDIApi.UpdateVDIPools(context.Background(), id).InlineObject262(inlineObject262).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VDIApi.UpdateVDIPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVDIPools`: AnyOfobject200Success
    fmt.Fprintf(os.Stdout, "Response from `VDIApi.UpdateVDIPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVDIPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject262** | [**InlineObject262**](InlineObject262.md) |  | 

### Return type

[**AnyOfobject200Success**](anyOf&lt;object,200-success&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

