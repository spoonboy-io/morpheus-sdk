# \WikiApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWiki**](WikiApi.md#AddWiki) | **Post** /api/wiki/pages | Create a Wiki Page
[**GetWiki**](WikiApi.md#GetWiki) | **Get** /api/wiki/pages/{id} | Retrieves a specific Wiki page
[**GetWikiApp**](WikiApi.md#GetWikiApp) | **Get** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**GetWikiCategories**](WikiApi.md#GetWikiCategories) | **Get** /api/wiki/categories | Retrieves all Wiki categories associated with the account
[**GetWikiCloud**](WikiApi.md#GetWikiCloud) | **Get** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**GetWikiCluster**](WikiApi.md#GetWikiCluster) | **Get** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**GetWikiGroup**](WikiApi.md#GetWikiGroup) | **Get** /api/groups/{id}/wiki | Retrieves a Group Wiki Page
[**GetWikiInstance**](WikiApi.md#GetWikiInstance) | **Get** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**GetWikiServer**](WikiApi.md#GetWikiServer) | **Get** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**ListWiki**](WikiApi.md#ListWiki) | **Get** /api/wiki/pages | Retrieves wiki pages associated with the account.
[**RemoveWiki**](WikiApi.md#RemoveWiki) | **Delete** /api/wiki/pages/{id} | Deletes a Wiki Page
[**UpdateWiki**](WikiApi.md#UpdateWiki) | **Put** /api/wiki/pages/{id} | Updates a Wiki Page
[**UpdateWikiApp**](WikiApi.md#UpdateWikiApp) | **Put** /api/apps/{id}/wiki | Update an App Wiki Page
[**UpdateWikiCloud**](WikiApi.md#UpdateWikiCloud) | **Put** /api/zones/{id}/wiki | Update a Cloud Wiki Page
[**UpdateWikiCluster**](WikiApi.md#UpdateWikiCluster) | **Put** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page
[**UpdateWikiGroup**](WikiApi.md#UpdateWikiGroup) | **Put** /api/groups/{id}/wiki | Update a Group Wiki Page
[**UpdateWikiInstance**](WikiApi.md#UpdateWikiInstance) | **Put** /api/instances/{id}/wiki | Update an Instance Wiki Page
[**UpdateWikiServer**](WikiApi.md#UpdateWikiServer) | **Put** /api/servers/{id}/wiki | Update a Server Wiki Page



## AddWiki

> map[string]interface{} AddWiki(ctx).InlineObject272(inlineObject272).Execute()

Create a Wiki Page



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
    inlineObject272 := *openapiclient.NewInlineObject272(*openapiclient.NewApiWikiPagesPage("Sample Doc", "info", "A sample document in **markdown**.")) // InlineObject272 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.AddWiki(context.Background()).InlineObject272(inlineObject272).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.AddWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWiki`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.AddWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject272** | [**InlineObject272**](InlineObject272.md) |  | 

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


## GetWiki

> InlineResponse200168 GetWiki(ctx, id).Execute()

Retrieves a specific Wiki page



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
    resp, r, err := api_client.WikiApi.GetWiki(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWiki`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWiki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiRequest struct via the builder pattern


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


## GetWikiApp

> InlineResponse200168 GetWikiApp(ctx, id).Execute()

Retrieves an App Wiki Page



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
    resp, r, err := api_client.WikiApi.GetWikiApp(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiApp`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiAppRequest struct via the builder pattern


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


## GetWikiCategories

> InlineResponse200169 GetWikiCategories(ctx).Phrase(phrase).PagePhrase(pagePhrase).Execute()

Retrieves all Wiki categories associated with the account



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
    pagePhrase := "pagePhrase_example" // string | If specified will return a partial match on page name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.GetWikiCategories(context.Background()).Phrase(phrase).PagePhrase(pagePhrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiCategories`: InlineResponse200169
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **pagePhrase** | **string** | If specified will return a partial match on page name | 

### Return type

[**InlineResponse200169**](inline_response_200_169.md)

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
    resp, r, err := api_client.WikiApi.GetWikiCloud(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiCloud`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiCloud`: %v\n", resp)
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
    resp, r, err := api_client.WikiApi.GetWikiCluster(context.Background(), clusterId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiCluster`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiCluster`: %v\n", resp)
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


## GetWikiGroup

> InlineResponse200168 GetWikiGroup(ctx, id).Execute()

Retrieves a Group Wiki Page



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
    resp, r, err := api_client.WikiApi.GetWikiGroup(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiGroup`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiGroupRequest struct via the builder pattern


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
    resp, r, err := api_client.WikiApi.GetWikiInstance(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiInstance`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiInstance`: %v\n", resp)
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


## GetWikiServer

> InlineResponse200168 GetWikiServer(ctx, id).Execute()

Retrieves a Server Wiki Page



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
    resp, r, err := api_client.WikiApi.GetWikiServer(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.GetWikiServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiServer`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.GetWikiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiServerRequest struct via the builder pattern


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


## ListWiki

> InlineResponse200168 ListWiki(ctx).Name(name).Phrase(phrase).Execute()

Retrieves wiki pages associated with the account.



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
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.ListWiki(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.ListWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWiki`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.ListWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

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


## RemoveWiki

> Model200Success RemoveWiki(ctx, id).Execute()

Deletes a Wiki Page



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
    resp, r, err := api_client.WikiApi.RemoveWiki(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.RemoveWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveWiki`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.RemoveWiki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWikiRequest struct via the builder pattern


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


## UpdateWiki

> map[string]interface{} UpdateWiki(ctx, id).InlineObject273(inlineObject273).Execute()

Updates a Wiki Page



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
    inlineObject273 := *openapiclient.NewInlineObject273() // InlineObject273 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.UpdateWiki(context.Background(), id).InlineObject273(inlineObject273).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWiki`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWiki`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject273** | [**InlineObject273**](InlineObject273.md) |  | 

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


## UpdateWikiApp

> map[string]interface{} UpdateWikiApp(ctx, id).InlineObject267(inlineObject267).Execute()

Update an App Wiki Page



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
    inlineObject267 := *openapiclient.NewInlineObject267() // InlineObject267 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.UpdateWikiApp(context.Background(), id).InlineObject267(inlineObject267).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWikiApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiApp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWikiApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject267** | [**InlineObject267**](InlineObject267.md) |  | 

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
    resp, r, err := api_client.WikiApi.UpdateWikiCloud(context.Background(), id).InlineObject274(inlineObject274).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWikiCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWikiCloud`: %v\n", resp)
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
    resp, r, err := api_client.WikiApi.UpdateWikiCluster(context.Background(), clusterId).InlineObject268(inlineObject268).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWikiCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWikiCluster`: %v\n", resp)
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


## UpdateWikiGroup

> map[string]interface{} UpdateWikiGroup(ctx, id).InlineObject269(inlineObject269).Execute()

Update a Group Wiki Page



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
    inlineObject269 := *openapiclient.NewInlineObject269() // InlineObject269 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.UpdateWikiGroup(context.Background(), id).InlineObject269(inlineObject269).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWikiGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWikiGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject269** | [**InlineObject269**](InlineObject269.md) |  | 

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
    resp, r, err := api_client.WikiApi.UpdateWikiInstance(context.Background(), id).InlineObject270(inlineObject270).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWikiInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWikiInstance`: %v\n", resp)
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


## UpdateWikiServer

> map[string]interface{} UpdateWikiServer(ctx, id).InlineObject271(inlineObject271).Execute()

Update a Server Wiki Page



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
    inlineObject271 := *openapiclient.NewInlineObject271() // InlineObject271 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WikiApi.UpdateWikiServer(context.Background(), id).InlineObject271(inlineObject271).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiApi.UpdateWikiServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiServer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WikiApi.UpdateWikiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject271** | [**InlineObject271**](InlineObject271.md) |  | 

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

