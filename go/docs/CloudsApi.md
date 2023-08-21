# \CloudsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCloudResourcePool**](CloudsApi.md#AddCloudResourcePool) | **Post** /api/zones/{zoneId}/resource-pools | Creates a Specified Resource Pool for Specified Cloud
[**AddClouds**](CloudsApi.md#AddClouds) | **Post** /api/zones | Creates a Cloud
[**GetCloudDatastores**](CloudsApi.md#GetCloudDatastores) | **Get** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**GetCloudFolders**](CloudsApi.md#GetCloudFolders) | **Get** /api/zones/{zoneId}/folders/{id} | Retrieves a Resource Folder for Specified Cloud
[**GetCloudResourcePools**](CloudsApi.md#GetCloudResourcePools) | **Get** /api/zones/{zoneId}/resource-pools/{id} | Retrieves a Resource Pool for Specified Cloud
[**GetCloudTypes**](CloudsApi.md#GetCloudTypes) | **Get** /api/zone-types/{id} | Retrieves a Specific Cloud Type
[**GetClouds**](CloudsApi.md#GetClouds) | **Get** /api/zones/{id} | Retrieves a Specific Cloud
[**GetWikiCloud**](CloudsApi.md#GetWikiCloud) | **Get** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**ListCloudDatastores**](CloudsApi.md#ListCloudDatastores) | **Get** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**ListCloudFolders**](CloudsApi.md#ListCloudFolders) | **Get** /api/zones/{zoneId}/folders | Retrieves all resource folders for Specified Cloud
[**ListCloudResourcePools**](CloudsApi.md#ListCloudResourcePools) | **Get** /api/zones/{zoneId}/resource-pools | Retrieves all Resource Pools for Specified Cloud
[**ListCloudSecurityGroups**](CloudsApi.md#ListCloudSecurityGroups) | **Get** /api/zones/{id}/security-groups | Retrieves all Security Groups for a Cloud
[**ListCloudTypes**](CloudsApi.md#ListCloudTypes) | **Get** /api/zone-types | Retrieves all Cloud Types
[**ListClouds**](CloudsApi.md#ListClouds) | **Get** /api/zones | Retrieves all Clouds
[**RefreshClouds**](CloudsApi.md#RefreshClouds) | **Post** /api/zones/{id}/refresh | Refreshes a Cloud
[**RemoveCloudResourcePools**](CloudsApi.md#RemoveCloudResourcePools) | **Delete** /api/zones/{zoneId}/resource-pools/{id} | Deletes a Resource Pool for Specified Cloud
[**RemoveClouds**](CloudsApi.md#RemoveClouds) | **Delete** /api/zones/{id} | Deletes a Cloud
[**UpdateCloudDatastores**](CloudsApi.md#UpdateCloudDatastores) | **Put** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**UpdateCloudFolders**](CloudsApi.md#UpdateCloudFolders) | **Put** /api/zones/{zoneId}/folders/{id} | Updates a Resource Folder for Specified Cloud
[**UpdateCloudLogo**](CloudsApi.md#UpdateCloudLogo) | **Post** /api/zones/{id}/update-logo | Update Logo For Cloud
[**UpdateCloudResourcePool**](CloudsApi.md#UpdateCloudResourcePool) | **Put** /api/zones/{zoneId}/resource-pools/{id} | Updates a Specified Resource Pool for Specified Cloud
[**UpdateCloudSecurityGroups**](CloudsApi.md#UpdateCloudSecurityGroups) | **Post** /api/zones/{id}/security-groups | Sets Security Groups for a Cloud
[**UpdateClouds**](CloudsApi.md#UpdateClouds) | **Put** /api/zones/{id} | Updates a Cloud
[**UpdateWikiCloud**](CloudsApi.md#UpdateWikiCloud) | **Put** /api/zones/{id}/wiki | Update a Cloud Wiki Page



## AddCloudResourcePool

> InlineResponse20024 AddCloudResourcePool(ctx, zoneId).InlineObject45(inlineObject45).Execute()

Creates a Specified Resource Pool for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    inlineObject45 := *openapiclient.NewInlineObject45(*openapiclient.NewApiZonesZoneIdResourcePoolsResourcePool("Name_example", "TODO")) // InlineObject45 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.AddCloudResourcePool(context.Background(), zoneId).InlineObject45(inlineObject45).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.AddCloudResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCloudResourcePool`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.AddCloudResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCloudResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject45** | [**InlineObject45**](InlineObject45.md) |  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClouds

> map[string]interface{} AddClouds(ctx).InlineObject41(inlineObject41).Execute()

Creates a Cloud



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
    inlineObject41 := *openapiclient.NewInlineObject41(*openapiclient.NewZoneCreate("My Cloud", "TODO", int64(3))) // InlineObject41 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.AddClouds(context.Background()).InlineObject41(inlineObject41).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.AddClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClouds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.AddClouds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject41** | [**InlineObject41**](InlineObject41.md) |  | 

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


## GetCloudDatastores

> map[string]interface{} GetCloudDatastores(ctx, zoneId, id).Execute()

Retrieves a Datastore for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetCloudDatastores(context.Background(), zoneId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetCloudDatastores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDatastores`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetCloudDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDatastoresRequest struct via the builder pattern


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


## GetCloudFolders

> map[string]interface{} GetCloudFolders(ctx, zoneId, id).Execute()

Retrieves a Resource Folder for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetCloudFolders(context.Background(), zoneId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetCloudFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudFolders`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetCloudFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudFoldersRequest struct via the builder pattern


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


## GetCloudResourcePools

> map[string]interface{} GetCloudResourcePools(ctx, zoneId, id).Execute()

Retrieves a Resource Pool for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.GetCloudResourcePools(context.Background(), zoneId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetCloudResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudResourcePools`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetCloudResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudResourcePoolsRequest struct via the builder pattern


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


## GetCloudTypes

> InlineResponse20020 GetCloudTypes(ctx, id).Execute()

Retrieves a Specific Cloud Type



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
    resp, r, err := api_client.CloudsApi.GetCloudTypes(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetCloudTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudTypes`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetCloudTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClouds

> InlineResponse20021 GetClouds(ctx, id).Execute()

Retrieves a Specific Cloud



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
    resp, r, err := api_client.CloudsApi.GetClouds(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClouds`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiCloud

> InlineResponse200168 GetWikiCloud(ctx, id).Execute()

Retrieves a Cloud Wiki Page



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
    resp, r, err := api_client.CloudsApi.GetWikiCloud(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.GetWikiCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiCloud`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.GetWikiCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiCloudRequest struct via the builder pattern


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


## ListCloudDatastores

> map[string]interface{} ListCloudDatastores(ctx, zoneId).Name(name).Phrase(phrase).Max(max).Sort(sort).Direction(direction).Execute()

Retrieves all Datastores for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.ListCloudDatastores(context.Background(), zoneId).Name(name).Phrase(phrase).Max(max).Sort(sort).Direction(direction).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListCloudDatastores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudDatastores`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListCloudDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]

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


## ListCloudFolders

> map[string]interface{} ListCloudFolders(ctx, zoneId).Name(name).Phrase(phrase).Max(max).Execute()

Retrieves all resource folders for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.ListCloudFolders(context.Background(), zoneId).Name(name).Phrase(phrase).Max(max).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListCloudFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudFolders`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListCloudFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]

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


## ListCloudResourcePools

> map[string]interface{} ListCloudResourcePools(ctx, zoneId).Name(name).Phrase(phrase).Max(max).Execute()

Retrieves all Resource Pools for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.ListCloudResourcePools(context.Background(), zoneId).Name(name).Phrase(phrase).Max(max).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListCloudResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudResourcePools`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListCloudResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]

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


## ListCloudSecurityGroups

> map[string]interface{} ListCloudSecurityGroups(ctx, id).Execute()

Retrieves all Security Groups for a Cloud



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
    resp, r, err := api_client.CloudsApi.ListCloudSecurityGroups(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListCloudSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudSecurityGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListCloudSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudSecurityGroupsRequest struct via the builder pattern


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


## ListCloudTypes

> map[string]interface{} ListCloudTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).Execute()

Retrieves all Cloud Types



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
    code := "azr" // string | If specified will return an exact match on code (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.ListCloudTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListCloudTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListCloudTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 

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


## ListClouds

> map[string]interface{} ListClouds(ctx).LastUpdated(lastUpdated).Type_(type_).GroupId(groupId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Clouds



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
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    type_ := "type__example" // string | If specified will return all zones by cloud type code. Refer to `Zone Types` API for up to date listings.  (optional)
    groupId := int64(13) // int64 | If specified will return all zones assigned to a server group by id. Accepts multiple values. (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.ListClouds(context.Background()).LastUpdated(lastUpdated).Type_(type_).GroupId(groupId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.ListClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClouds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.ListClouds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **type_** | **string** | If specified will return all zones by cloud type code. Refer to &#x60;Zone Types&#x60; API for up to date listings.  | 
 **groupId** | **int64** | If specified will return all zones assigned to a server group by id. Accepts multiple values. | 
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


## RefreshClouds

> Model200Success RefreshClouds(ctx, id).InlineObject47(inlineObject47).Execute()

Refreshes a Cloud



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
    inlineObject47 := *openapiclient.NewInlineObject47() // InlineObject47 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.RefreshClouds(context.Background(), id).InlineObject47(inlineObject47).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.RefreshClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshClouds`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.RefreshClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject47** | [**InlineObject47**](InlineObject47.md) |  | 

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


## RemoveCloudResourcePools

> Model200Success RemoveCloudResourcePools(ctx, zoneId, id).Execute()

Deletes a Resource Pool for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.RemoveCloudResourcePools(context.Background(), zoneId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.RemoveCloudResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCloudResourcePools`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.RemoveCloudResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCloudResourcePoolsRequest struct via the builder pattern


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


## RemoveClouds

> Model200Success RemoveClouds(ctx, id).RemoveResources(removeResources).Execute()

Deletes a Cloud



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
    removeResources := true // bool | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.RemoveClouds(context.Background(), id).RemoveResources(removeResources).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.RemoveClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveClouds`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.RemoveClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeResources** | **bool** | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [default to false]

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


## UpdateCloudDatastores

> InlineResponse20022 UpdateCloudDatastores(ctx, zoneId, id).InlineObject43(inlineObject43).Execute()

Updates a Specified Datastore for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject43 := *openapiclient.NewInlineObject43(*openapiclient.NewApiZonesZoneIdDataStoresIdDatastore()) // InlineObject43 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateCloudDatastores(context.Background(), zoneId, id).InlineObject43(inlineObject43).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateCloudDatastores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudDatastores`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateCloudDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject43** | [**InlineObject43**](InlineObject43.md) |  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudFolders

> InlineResponse20023 UpdateCloudFolders(ctx, zoneId, id).InlineObject44(inlineObject44).Execute()

Updates a Resource Folder for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject44 := *openapiclient.NewInlineObject44(*openapiclient.NewApiZonesZoneIdFoldersIdFolder()) // InlineObject44 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateCloudFolders(context.Background(), zoneId, id).InlineObject44(inlineObject44).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateCloudFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudFolders`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateCloudFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject44** | [**InlineObject44**](InlineObject44.md) |  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudLogo

> Model200Success UpdateCloudLogo(ctx, id).Logo(logo).DarkLogo(darkLogo).Execute()

Update Logo For Cloud



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
    logo := os.NewFile(1234, "some_file") // *os.File | Logo File png,jpg,svg (optional)
    darkLogo := os.NewFile(1234, "some_file") // *os.File | Dark Logo File png,jpg,svg (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateCloudLogo(context.Background(), id).Logo(logo).DarkLogo(darkLogo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateCloudLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudLogo`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateCloudLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logo** | ***os.File** | Logo File png,jpg,svg | 
 **darkLogo** | ***os.File** | Dark Logo File png,jpg,svg | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudResourcePool

> InlineResponse20024 UpdateCloudResourcePool(ctx, zoneId, id).InlineObject46(inlineObject46).Execute()

Updates a Specified Resource Pool for Specified Cloud



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
    zoneId := float32(7) // float32 | The ID of the cloud
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject46 := *openapiclient.NewInlineObject46(*openapiclient.NewApiZonesZoneIdResourcePoolsIdResourcePool()) // InlineObject46 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateCloudResourcePool(context.Background(), zoneId, id).InlineObject46(inlineObject46).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateCloudResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudResourcePool`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateCloudResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject46** | [**InlineObject46**](InlineObject46.md) |  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudSecurityGroups

> map[string]interface{} UpdateCloudSecurityGroups(ctx, id).InlineObject49(inlineObject49).Execute()

Sets Security Groups for a Cloud



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
    inlineObject49 := *openapiclient.NewInlineObject49([]int64{int64(123)}) // InlineObject49 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateCloudSecurityGroups(context.Background(), id).InlineObject49(inlineObject49).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateCloudSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudSecurityGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateCloudSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject49** | [**InlineObject49**](InlineObject49.md) |  | 

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


## UpdateClouds

> map[string]interface{} UpdateClouds(ctx, id).InlineObject42(inlineObject42).Execute()

Updates a Cloud



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
    inlineObject42 := *openapiclient.NewInlineObject42(*openapiclient.NewApiZonesIdZone("My Cloud", "ZoneType_example", int64(3), map[string]interface{}({"id":1}))) // InlineObject42 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateClouds(context.Background(), id).InlineObject42(inlineObject42).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateClouds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClouds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject42** | [**InlineObject42**](InlineObject42.md) |  | 

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


## UpdateWikiCloud

> map[string]interface{} UpdateWikiCloud(ctx, id).InlineObject274(inlineObject274).Execute()

Update a Cloud Wiki Page



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
    inlineObject274 := *openapiclient.NewInlineObject274() // InlineObject274 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsApi.UpdateWikiCloud(context.Background(), id).InlineObject274(inlineObject274).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsApi.UpdateWikiCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudsApi.UpdateWikiCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject274** | [**InlineObject274**](InlineObject274.md) |  | 

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

