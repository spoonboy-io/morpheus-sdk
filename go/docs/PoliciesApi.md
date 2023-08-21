# \PoliciesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicies**](PoliciesApi.md#AddPolicies) | **Post** /api/policies | Creates a Policy
[**AddPoliciesCloud**](PoliciesApi.md#AddPoliciesCloud) | **Post** /api/zones/{cloudId}/policies | Creates a Policy for a Cloud
[**AddPoliciesGroup**](PoliciesApi.md#AddPoliciesGroup) | **Post** /api/groups/{groupId}/policies | Creates a Policy for a Group
[**GetPolicies**](PoliciesApi.md#GetPolicies) | **Get** /api/policies/{id} | Retrieves a Specific Policy
[**GetPoliciesCloud**](PoliciesApi.md#GetPoliciesCloud) | **Get** /api/zones/{cloudId}/policies/{id} | Retrieves a Specific Policy for a Cloud
[**GetPoliciesGroup**](PoliciesApi.md#GetPoliciesGroup) | **Get** /api/groups/{groupId}/policies/{id} | Retrieves a Specific Policy for a Group
[**ListPolicies**](PoliciesApi.md#ListPolicies) | **Get** /api/policies | Retrieves all Policies
[**ListPoliciesCloud**](PoliciesApi.md#ListPoliciesCloud) | **Get** /api/zones/{cloudId}/policies | Retrieves Policies for a Cloud
[**ListPoliciesGroup**](PoliciesApi.md#ListPoliciesGroup) | **Get** /api/groups/{groupId}/policies | Retrieves Policies for a Group
[**ListPolicyTypes**](PoliciesApi.md#ListPolicyTypes) | **Get** /api/policy-types | Retrieves all Policy Types
[**RemovePolicies**](PoliciesApi.md#RemovePolicies) | **Delete** /api/policies/{id} | Deletes a Policy
[**RemovePoliciesCloud**](PoliciesApi.md#RemovePoliciesCloud) | **Delete** /api/zones/{cloudId}/policies/{id} | Deletes a Policy for a Cloud
[**RemovePoliciesGroup**](PoliciesApi.md#RemovePoliciesGroup) | **Delete** /api/groups/{groupId}/policies/{id} | Deletes a Policy for a Group
[**UpdatePolicies**](PoliciesApi.md#UpdatePolicies) | **Put** /api/policies/{id} | Updates a Policy
[**UpdatePoliciesCloud**](PoliciesApi.md#UpdatePoliciesCloud) | **Put** /api/zones/{cloudId}/policies/{id} | Updates a Policy for a Cloud
[**UpdatePoliciesGroup**](PoliciesApi.md#UpdatePoliciesGroup) | **Put** /api/groups/{groupId}/policies/{id} | Updates a Policy for a Group



## AddPolicies

> map[string]interface{} AddPolicies(ctx).InlineObject184(inlineObject184).Execute()

Creates a Policy



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
    inlineObject184 := *openapiclient.NewInlineObject184(*openapiclient.NewPolicyCreate("Sample Policy", *openapiclient.NewPolicyCreatePolicyType(), "TODO")) // InlineObject184 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.AddPolicies(context.Background()).InlineObject184(inlineObject184).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.AddPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.AddPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject184** | [**InlineObject184**](InlineObject184.md) |  | 

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


## AddPoliciesCloud

> map[string]interface{} AddPoliciesCloud(ctx, cloudId).InlineObject188(inlineObject188).Execute()

Creates a Policy for a Cloud



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
    cloudId := int64(7) // int64 | The ID of the cloud
    inlineObject188 := *openapiclient.NewInlineObject188(*openapiclient.NewPolicyCloudCreate("Sample Policy", *openapiclient.NewPolicyCloudCreatePolicyType())) // InlineObject188 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.AddPoliciesCloud(context.Background(), cloudId).InlineObject188(inlineObject188).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.AddPoliciesCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPoliciesCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.AddPoliciesCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int64** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPoliciesCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject188** | [**InlineObject188**](InlineObject188.md) |  | 

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


## AddPoliciesGroup

> map[string]interface{} AddPoliciesGroup(ctx, groupId).InlineObject186(inlineObject186).Execute()

Creates a Policy for a Group



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
    groupId := int64(7) // int64 | The ID of the group
    inlineObject186 := *openapiclient.NewInlineObject186(*openapiclient.NewPolicyGroupCreate("Sample Policy", *openapiclient.NewPolicyGroupCreatePolicyType())) // InlineObject186 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.AddPoliciesGroup(context.Background(), groupId).InlineObject186(inlineObject186).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.AddPoliciesGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPoliciesGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.AddPoliciesGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPoliciesGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject186** | [**InlineObject186**](InlineObject186.md) |  | 

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


## GetPolicies

> map[string]interface{} GetPolicies(ctx, id).Execute()

Retrieves a Specific Policy



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
    resp, r, err := api_client.PoliciesApi.GetPolicies(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


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


## GetPoliciesCloud

> map[string]interface{} GetPoliciesCloud(ctx, cloudId, id).Execute()

Retrieves a Specific Policy for a Cloud



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
    cloudId := int64(7) // int64 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPoliciesCloud(context.Background(), cloudId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPoliciesCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoliciesCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPoliciesCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int64** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesCloudRequest struct via the builder pattern


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


## GetPoliciesGroup

> map[string]interface{} GetPoliciesGroup(ctx, groupId, id).Execute()

Retrieves a Specific Policy for a Group



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
    groupId := int64(7) // int64 | The ID of the group
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.GetPoliciesGroup(context.Background(), groupId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.GetPoliciesGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoliciesGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.GetPoliciesGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesGroupRequest struct via the builder pattern


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


## ListPolicies

> map[string]interface{} ListPolicies(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Policies



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
    resp, r, err := api_client.PoliciesApi.ListPolicies(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


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


## ListPoliciesCloud

> map[string]interface{} ListPoliciesCloud(ctx, cloudId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves Policies for a Cloud



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
    cloudId := int64(7) // int64 | The ID of the cloud
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.ListPoliciesCloud(context.Background(), cloudId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ListPoliciesCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPoliciesCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ListPoliciesCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int64** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesCloudRequest struct via the builder pattern


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


## ListPoliciesGroup

> map[string]interface{} ListPoliciesGroup(ctx, groupId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves Policies for a Group



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
    groupId := int64(7) // int64 | The ID of the group
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.ListPoliciesGroup(context.Background(), groupId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ListPoliciesGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPoliciesGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ListPoliciesGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesGroupRequest struct via the builder pattern


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


## ListPolicyTypes

> map[string]interface{} ListPolicyTypes(ctx).Execute()

Retrieves all Policy Types



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
    resp, r, err := api_client.PoliciesApi.ListPolicyTypes(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.ListPolicyTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.ListPolicyTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyTypesRequest struct via the builder pattern


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


## RemovePolicies

> Model200Success RemovePolicies(ctx, id).Execute()

Deletes a Policy



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
    resp, r, err := api_client.PoliciesApi.RemovePolicies(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.RemovePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePolicies`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.RemovePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePoliciesRequest struct via the builder pattern


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


## RemovePoliciesCloud

> Model200Success RemovePoliciesCloud(ctx, cloudId, id).Execute()

Deletes a Policy for a Cloud



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
    cloudId := int64(7) // int64 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.RemovePoliciesCloud(context.Background(), cloudId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.RemovePoliciesCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePoliciesCloud`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.RemovePoliciesCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int64** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePoliciesCloudRequest struct via the builder pattern


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


## RemovePoliciesGroup

> Model200Success RemovePoliciesGroup(ctx, groupId, id).Execute()

Deletes a Policy for a Group



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
    groupId := int64(7) // int64 | The ID of the group
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.RemovePoliciesGroup(context.Background(), groupId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.RemovePoliciesGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePoliciesGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.RemovePoliciesGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePoliciesGroupRequest struct via the builder pattern


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


## UpdatePolicies

> map[string]interface{} UpdatePolicies(ctx, id).InlineObject185(inlineObject185).Execute()

Updates a Policy



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
    inlineObject185 := *openapiclient.NewInlineObject185(*openapiclient.NewPolicyUpdate()) // InlineObject185 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.UpdatePolicies(context.Background(), id).InlineObject185(inlineObject185).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.UpdatePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.UpdatePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject185** | [**InlineObject185**](InlineObject185.md) |  | 

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


## UpdatePoliciesCloud

> map[string]interface{} UpdatePoliciesCloud(ctx, cloudId, id).InlineObject189(inlineObject189).Execute()

Updates a Policy for a Cloud



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
    cloudId := int64(7) // int64 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject189 := *openapiclient.NewInlineObject189(*openapiclient.NewPolicyCloudUpdate()) // InlineObject189 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.UpdatePoliciesCloud(context.Background(), cloudId, id).InlineObject189(inlineObject189).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.UpdatePoliciesCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePoliciesCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.UpdatePoliciesCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int64** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoliciesCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject189** | [**InlineObject189**](InlineObject189.md) |  | 

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


## UpdatePoliciesGroup

> map[string]interface{} UpdatePoliciesGroup(ctx, groupId, id).InlineObject187(inlineObject187).Execute()

Updates a Policy for a Group



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
    groupId := int64(7) // int64 | The ID of the group
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject187 := *openapiclient.NewInlineObject187(*openapiclient.NewPolicyGroupUpdate()) // InlineObject187 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.UpdatePoliciesGroup(context.Background(), groupId, id).InlineObject187(inlineObject187).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.UpdatePoliciesGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePoliciesGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.UpdatePoliciesGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoliciesGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject187** | [**InlineObject187**](InlineObject187.md) |  | 

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

