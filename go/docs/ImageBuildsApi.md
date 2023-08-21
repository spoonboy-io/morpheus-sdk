# \ImageBuildsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBootScript**](ImageBuildsApi.md#AddBootScript) | **Post** /api/boot-scripts | Create a Boot Script
[**AddImageBuild**](ImageBuildsApi.md#AddImageBuild) | **Post** /api/image-builds | Create an Image Build
[**AddPreseedScript**](ImageBuildsApi.md#AddPreseedScript) | **Post** /api/preseed-scripts | Create a Preseed Script
[**DeleteBootScript**](ImageBuildsApi.md#DeleteBootScript) | **Delete** /api/boot-scripts/{id} | Delete a Boot Script
[**DeleteImageBuild**](ImageBuildsApi.md#DeleteImageBuild) | **Delete** /api/image-builds/{id} | Delete an Image Build
[**DeletePreseedScript**](ImageBuildsApi.md#DeletePreseedScript) | **Delete** /api/preseed-scripts/{id} | Delete a Preseed Script
[**ExecuteImageBuild**](ImageBuildsApi.md#ExecuteImageBuild) | **Post** /api/image-builds/{id}/run | Run an Image Build
[**GetBootScript**](ImageBuildsApi.md#GetBootScript) | **Get** /api/boot-scripts/{id} | Get a Specific Boot Script
[**GetImageBuild**](ImageBuildsApi.md#GetImageBuild) | **Get** /api/image-builds/{id} | Get a Specific Image Build
[**GetImageBuildExecutions**](ImageBuildsApi.md#GetImageBuildExecutions) | **Get** /api/image-builds/{id}/list-executions | List Image Build Executions
[**GetPreseedScript**](ImageBuildsApi.md#GetPreseedScript) | **Get** /api/preseed-scripts/{id} | Get a Specific Preseed Script
[**ListBootScripts**](ImageBuildsApi.md#ListBootScripts) | **Get** /api/boot-scripts | Boot Scripts
[**ListImageBuilds**](ImageBuildsApi.md#ListImageBuilds) | **Get** /api/image-builds | Get All Image Builds
[**ListPreseedScripts**](ImageBuildsApi.md#ListPreseedScripts) | **Get** /api/preseed-scripts | Preseed Scripts
[**UpdateBootScript**](ImageBuildsApi.md#UpdateBootScript) | **Put** /api/boot-scripts/{id} | Update a Boot Script
[**UpdateImageBuild**](ImageBuildsApi.md#UpdateImageBuild) | **Put** /api/image-builds/{id} | Update an Image Build
[**UpdatePreseedScript**](ImageBuildsApi.md#UpdatePreseedScript) | **Put** /api/preseed-scripts/{id} | Update a Preseed Script



## AddBootScript

> map[string]interface{} AddBootScript(ctx).InlineObject83(inlineObject83).Execute()

Create a Boot Script



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject83 := *openapiclient.NewInlineObject83() // InlineObject83 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.AddBootScript(context.Background()).InlineObject83(inlineObject83).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.AddBootScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBootScript`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.AddBootScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBootScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject83** | [**InlineObject83**](InlineObject83.md) |  | 

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


## AddImageBuild

> map[string]interface{} AddImageBuild(ctx).InlineObject85(inlineObject85).Execute()

Create an Image Build



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject85 := *openapiclient.NewInlineObject85() // InlineObject85 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.AddImageBuild(context.Background()).InlineObject85(inlineObject85).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.AddImageBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddImageBuild`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.AddImageBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddImageBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject85** | [**InlineObject85**](InlineObject85.md) |  | 

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


## AddPreseedScript

> map[string]interface{} AddPreseedScript(ctx).InlineObject196(inlineObject196).Execute()

Create a Preseed Script



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject196 := *openapiclient.NewInlineObject196() // InlineObject196 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.AddPreseedScript(context.Background()).InlineObject196(inlineObject196).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.AddPreseedScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPreseedScript`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.AddPreseedScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPreseedScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject196** | [**InlineObject196**](InlineObject196.md) |  | 

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


## DeleteBootScript

> Model200Success DeleteBootScript(ctx, id).Execute()

Delete a Boot Script



### Example

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
    resp, r, err := api_client.ImageBuildsApi.DeleteBootScript(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.DeleteBootScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBootScript`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.DeleteBootScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBootScriptRequest struct via the builder pattern


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


## DeleteImageBuild

> Model200Success DeleteImageBuild(ctx, id).Execute()

Delete an Image Build



### Example

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
    resp, r, err := api_client.ImageBuildsApi.DeleteImageBuild(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.DeleteImageBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteImageBuild`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.DeleteImageBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageBuildRequest struct via the builder pattern


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


## DeletePreseedScript

> Model200Success DeletePreseedScript(ctx, id).Execute()

Delete a Preseed Script



### Example

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
    resp, r, err := api_client.ImageBuildsApi.DeletePreseedScript(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.DeletePreseedScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePreseedScript`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.DeletePreseedScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePreseedScriptRequest struct via the builder pattern


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


## ExecuteImageBuild

> Model200Success ExecuteImageBuild(ctx, id).Execute()

Run an Image Build



### Example

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
    resp, r, err := api_client.ImageBuildsApi.ExecuteImageBuild(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.ExecuteImageBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteImageBuild`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.ExecuteImageBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteImageBuildRequest struct via the builder pattern


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


## GetBootScript

> InlineResponse20052 GetBootScript(ctx, id).Execute()

Get a Specific Boot Script



### Example

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
    resp, r, err := api_client.ImageBuildsApi.GetBootScript(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.GetBootScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBootScript`: InlineResponse20052
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.GetBootScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBootScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageBuild

> InlineResponse20053 GetImageBuild(ctx, id).Execute()

Get a Specific Image Build



### Example

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
    resp, r, err := api_client.ImageBuildsApi.GetImageBuild(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.GetImageBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageBuild`: InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.GetImageBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageBuildExecutions

> map[string]interface{} GetImageBuildExecutions(ctx, id).BuildNumber(buildNumber).Status(status).Execute()

List Image Build Executions



### Example

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
    buildNumber := int64(4) // int64 | If specified will return an exact match on buildNumber (optional)
    status := "running" // string | Filter by status (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.GetImageBuildExecutions(context.Background(), id).BuildNumber(buildNumber).Status(status).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.GetImageBuildExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageBuildExecutions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.GetImageBuildExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageBuildExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildNumber** | **int64** | If specified will return an exact match on buildNumber | 
 **status** | **string** | Filter by status | 

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


## GetPreseedScript

> InlineResponse200122 GetPreseedScript(ctx, id).Execute()

Get a Specific Preseed Script



### Example

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
    resp, r, err := api_client.ImageBuildsApi.GetPreseedScript(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.GetPreseedScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreseedScript`: InlineResponse200122
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.GetPreseedScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreseedScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200122**](inline_response_200_122.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBootScripts

> map[string]interface{} ListBootScripts(ctx).Name(name).Phrase(phrase).Execute()

Boot Scripts



### Example

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
    resp, r, err := api_client.ImageBuildsApi.ListBootScripts(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.ListBootScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBootScripts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.ListBootScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBootScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## ListImageBuilds

> map[string]interface{} ListImageBuilds(ctx).Name(name).Phrase(phrase).Execute()

Get All Image Builds



### Example

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
    resp, r, err := api_client.ImageBuildsApi.ListImageBuilds(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.ListImageBuilds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageBuilds`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.ListImageBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImageBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## ListPreseedScripts

> map[string]interface{} ListPreseedScripts(ctx).Name(name).Phrase(phrase).Execute()

Preseed Scripts



### Example

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
    resp, r, err := api_client.ImageBuildsApi.ListPreseedScripts(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.ListPreseedScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPreseedScripts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.ListPreseedScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPreseedScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## UpdateBootScript

> map[string]interface{} UpdateBootScript(ctx, id).InlineObject84(inlineObject84).Execute()

Update a Boot Script



### Example

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
    inlineObject84 := *openapiclient.NewInlineObject84() // InlineObject84 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.UpdateBootScript(context.Background(), id).InlineObject84(inlineObject84).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.UpdateBootScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBootScript`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.UpdateBootScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBootScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject84** | [**InlineObject84**](InlineObject84.md) |  | 

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


## UpdateImageBuild

> InlineResponse20054 UpdateImageBuild(ctx, id).InlineObject86(inlineObject86).Execute()

Update an Image Build



### Example

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
    inlineObject86 := *openapiclient.NewInlineObject86() // InlineObject86 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.UpdateImageBuild(context.Background(), id).InlineObject86(inlineObject86).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.UpdateImageBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImageBuild`: InlineResponse20054
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.UpdateImageBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject86** | [**InlineObject86**](InlineObject86.md) |  | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreseedScript

> map[string]interface{} UpdatePreseedScript(ctx, id).InlineObject197(inlineObject197).Execute()

Update a Preseed Script



### Example

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
    inlineObject197 := *openapiclient.NewInlineObject197() // InlineObject197 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImageBuildsApi.UpdatePreseedScript(context.Background(), id).InlineObject197(inlineObject197).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageBuildsApi.UpdatePreseedScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePreseedScript`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageBuildsApi.UpdatePreseedScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePreseedScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject197** | [**InlineObject197**](InlineObject197.md) |  | 

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

