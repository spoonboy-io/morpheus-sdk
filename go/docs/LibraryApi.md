# \LibraryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileTemplate**](LibraryApi.md#AddFileTemplate) | **Post** /api/library/container-templates | Create a File Template
[**AddInstanceType**](LibraryApi.md#AddInstanceType) | **Post** /api/library/instance-types | Create an Instance Type
[**AddLayout**](LibraryApi.md#AddLayout) | **Post** /api/library/instance-types/{instanceTypeId}/layouts | Create a Layout
[**AddNodeType**](LibraryApi.md#AddNodeType) | **Post** /api/library/container-types | Create a Node Type
[**AddOptionList**](LibraryApi.md#AddOptionList) | **Post** /api/library/option-type-lists | Create an Option List
[**AddOptionType**](LibraryApi.md#AddOptionType) | **Post** /api/library/option-types | Create an Input
[**AddScript**](LibraryApi.md#AddScript) | **Post** /api/library/container-scripts | Create a Script
[**AddSpecTemplate**](LibraryApi.md#AddSpecTemplate) | **Post** /api/library/spec-templates | Create a Spec Template
[**AddVirtualImage**](LibraryApi.md#AddVirtualImage) | **Post** /api/virtual-images | Create a Virtual Image
[**AddVirtualImageFile**](LibraryApi.md#AddVirtualImageFile) | **Post** /api/virtual-images/{virtualImageId}/upload | Upload Virtual Image File
[**DeleteFileTemplate**](LibraryApi.md#DeleteFileTemplate) | **Delete** /api/library/container-templates/{id} | Delete a File Template
[**DeleteInstanceType**](LibraryApi.md#DeleteInstanceType) | **Delete** /api/library/instance-types/{instanceTypeId} | Delete an Instance Type
[**DeleteLayout**](LibraryApi.md#DeleteLayout) | **Delete** /api/library/layouts/{id} | Delete a Layout
[**DeleteNodeType**](LibraryApi.md#DeleteNodeType) | **Delete** /api/library/container-types/{id} | Delete a Node Type
[**DeleteOptionList**](LibraryApi.md#DeleteOptionList) | **Delete** /api/library/option-type-lists/{id} | Delete an Option List
[**DeleteOptionType**](LibraryApi.md#DeleteOptionType) | **Delete** /api/library/option-types/{id} | Delete an Input
[**DeleteScript**](LibraryApi.md#DeleteScript) | **Delete** /api/library/container-scripts/{id} | Delete a Script
[**DeleteSpecTemplate**](LibraryApi.md#DeleteSpecTemplate) | **Delete** /api/library/spec-templates/{id} | Delete a Spec Template
[**GetFileTemplate**](LibraryApi.md#GetFileTemplate) | **Get** /api/library/container-templates/{id} | Get a Specific File Template
[**GetInput**](LibraryApi.md#GetInput) | **Get** /api/library/option-types/{id} | Get A Specific Input
[**GetInstanceType**](LibraryApi.md#GetInstanceType) | **Get** /api/library/instance-types/{instanceTypeId} | Get a Specific Instance Type
[**GetLayout**](LibraryApi.md#GetLayout) | **Get** /api/library/layouts/{id} | Get a Specific Layout
[**GetNodeType**](LibraryApi.md#GetNodeType) | **Get** /api/library/container-types/{id} | Get a Specific Node Type
[**GetOptionList**](LibraryApi.md#GetOptionList) | **Get** /api/library/option-type-lists/{id} | Get a Specific Option List
[**GetOptionListItems**](LibraryApi.md#GetOptionListItems) | **Get** /api/library/option-type-lists/{id}/items | List Items for a Specific Option List
[**GetScript**](LibraryApi.md#GetScript) | **Get** /api/library/container-scripts/{id} | Get a Specific Script
[**GetSecurityPackageType**](LibraryApi.md#GetSecurityPackageType) | **Get** /api/security-package-types/{id} | Retrieves a Specific Security Package Type
[**GetSpecTemplate**](LibraryApi.md#GetSpecTemplate) | **Get** /api/library/spec-templates/{id} | Get a Specific Spec Template
[**GetVirtualImage**](LibraryApi.md#GetVirtualImage) | **Get** /api/virtual-images/{virtualImageId} | Get a Specific Virtual Image
[**ListFileTemplates**](LibraryApi.md#ListFileTemplates) | **Get** /api/library/container-templates | Get All File Templates
[**ListInputs**](LibraryApi.md#ListInputs) | **Get** /api/library/option-types | Get All Inputs
[**ListInstanceTypes**](LibraryApi.md#ListInstanceTypes) | **Get** /api/library/instance-types | Get All Instance Types
[**ListLayouts**](LibraryApi.md#ListLayouts) | **Get** /api/library/layouts | Get All Layouts
[**ListLayoutsForInstanceType**](LibraryApi.md#ListLayoutsForInstanceType) | **Get** /api/library/instance-types/{instanceTypeId}/layouts | Get All Layouts For an Instance Type
[**ListNodeTypes**](LibraryApi.md#ListNodeTypes) | **Get** /api/library/container-types | Get All Node Types
[**ListOptionLists**](LibraryApi.md#ListOptionLists) | **Get** /api/library/option-type-lists | Get All Option Lists
[**ListScripts**](LibraryApi.md#ListScripts) | **Get** /api/library/container-scripts | Get All Scripts
[**ListSecurityPackageTypes**](LibraryApi.md#ListSecurityPackageTypes) | **Get** /api/security-package-types | Retrieves all Security Package Types
[**ListSpecTemplates**](LibraryApi.md#ListSpecTemplates) | **Get** /api/library/spec-templates | Get All Spec Templates
[**ListVirtualImageLocations**](LibraryApi.md#ListVirtualImageLocations) | **Get** /api/virtual-images/{virtualImageId}/locations | Get a List of Virtual Image Locations
[**ListVirtualImages**](LibraryApi.md#ListVirtualImages) | **Get** /api/virtual-images | Get List of Virtual Images
[**RemoveSecurityScans**](LibraryApi.md#RemoveSecurityScans) | **Delete** /api/security-scans/{id} | Deletes a Security Scan
[**RemoveVirtualImage**](LibraryApi.md#RemoveVirtualImage) | **Delete** /api/virtual-images/{virtualImageId} | Delete a Virtual Image
[**RemoveVirtualImageFile**](LibraryApi.md#RemoveVirtualImageFile) | **Delete** /api/virtual-images/{virtualImageId}/files | Remove Virtual Image File
[**RemoveVirtualImageLocation**](LibraryApi.md#RemoveVirtualImageLocation) | **Delete** /api/virtual-images/{virtualImageId}/locations/{id} | Delete a Virtual Image Location
[**SetInstanceTypeFeatured**](LibraryApi.md#SetInstanceTypeFeatured) | **Put** /api/library/instance-types/{instanceTypeId}/toggle-featured | Toggle Featured For Instance Type
[**UpdateFileTemplate**](LibraryApi.md#UpdateFileTemplate) | **Put** /api/library/container-templates/{id} | Update a File Template
[**UpdateInstanceType**](LibraryApi.md#UpdateInstanceType) | **Put** /api/library/instance-types/{instanceTypeId} | Update an Instance Type
[**UpdateInstanceTypeLogo**](LibraryApi.md#UpdateInstanceTypeLogo) | **Post** /api/library/instance-types/{instanceTypeId}/update-logo | Update Logo For Instance Type
[**UpdateLayout**](LibraryApi.md#UpdateLayout) | **Put** /api/library/layouts/{id} | Update a Layout
[**UpdateLayoutPermissions**](LibraryApi.md#UpdateLayoutPermissions) | **Put** /api/library/layouts/{id}/permissions | Update Layout Permissions
[**UpdateNodeType**](LibraryApi.md#UpdateNodeType) | **Put** /api/library/container-types/{id} | Update a Node Type
[**UpdateOptionList**](LibraryApi.md#UpdateOptionList) | **Put** /api/library/option-type-lists/{id} | Update an Option List
[**UpdateOptionType**](LibraryApi.md#UpdateOptionType) | **Put** /api/library/option-types/{id} | Update an Input
[**UpdateScript**](LibraryApi.md#UpdateScript) | **Put** /api/library/container-scripts/{id} | Update a Script
[**UpdateSpecTemplate**](LibraryApi.md#UpdateSpecTemplate) | **Put** /api/library/spec-templates/{id} | Update a Spec Template
[**UpdateVirtualImage**](LibraryApi.md#UpdateVirtualImage) | **Put** /api/virtual-images/{virtualImageId} | Update a Virtual Image



## AddFileTemplate

> SuccessId AddFileTemplate(ctx).InlineObject111(inlineObject111).Execute()

Create a File Template



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
    inlineObject111 := *openapiclient.NewInlineObject111() // InlineObject111 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddFileTemplate(context.Background()).InlineObject111(inlineObject111).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddFileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFileTemplate`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddFileTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject111** | [**InlineObject111**](InlineObject111.md) |  | 

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


## AddInstanceType

> Model200Success AddInstanceType(ctx).InlineObject113(inlineObject113).Execute()

Create an Instance Type



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
    inlineObject113 := *openapiclient.NewInlineObject113() // InlineObject113 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddInstanceType(context.Background()).InlineObject113(inlineObject113).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInstanceType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddInstanceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject113** | [**InlineObject113**](InlineObject113.md) |  | 

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


## AddLayout

> SuccessId AddLayout(ctx, instanceTypeId).InlineObject115(inlineObject115).Execute()

Create a Layout



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type
    inlineObject115 := *openapiclient.NewInlineObject115() // InlineObject115 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddLayout(context.Background(), instanceTypeId).InlineObject115(inlineObject115).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLayout`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject115** | [**InlineObject115**](InlineObject115.md) |  | 

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


## AddNodeType

> map[string]interface{} AddNodeType(ctx).InlineObject109(inlineObject109).Execute()

Create a Node Type



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
    inlineObject109 := *openapiclient.NewInlineObject109() // InlineObject109 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddNodeType(context.Background()).InlineObject109(inlineObject109).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddNodeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNodeType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddNodeType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNodeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject109** | [**InlineObject109**](InlineObject109.md) |  | 

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


## AddOptionList

> Model200Success AddOptionList(ctx).InlineObject119(inlineObject119).Execute()

Create an Option List



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
    inlineObject119 := *openapiclient.NewInlineObject119() // InlineObject119 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddOptionList(context.Background()).InlineObject119(inlineObject119).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddOptionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOptionList`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddOptionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject119** | [**InlineObject119**](InlineObject119.md) |  | 

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


## AddOptionType

> SuccessId AddOptionType(ctx).InlineObject121(inlineObject121).Execute()

Create an Input



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
    inlineObject121 := *openapiclient.NewInlineObject121() // InlineObject121 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddOptionType(context.Background()).InlineObject121(inlineObject121).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddOptionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOptionType`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddOptionType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOptionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject121** | [**InlineObject121**](InlineObject121.md) |  | 

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


## AddScript

> ScriptSuccessId AddScript(ctx).InlineObject107(inlineObject107).Execute()

Create a Script



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
    inlineObject107 := *openapiclient.NewInlineObject107() // InlineObject107 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddScript(context.Background()).InlineObject107(inlineObject107).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScript`: ScriptSuccessId
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject107** | [**InlineObject107**](InlineObject107.md) |  | 

### Return type

[**ScriptSuccessId**](scriptSuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSpecTemplate

> SuccessId AddSpecTemplate(ctx).InlineObject123(inlineObject123).Execute()

Create a Spec Template



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
    inlineObject123 := *openapiclient.NewInlineObject123() // InlineObject123 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddSpecTemplate(context.Background()).InlineObject123(inlineObject123).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddSpecTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSpecTemplate`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddSpecTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSpecTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject123** | [**InlineObject123**](InlineObject123.md) |  | 

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


## AddVirtualImage

> map[string]interface{} AddVirtualImage(ctx).InlineObject263(inlineObject263).Execute()

Create a Virtual Image



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
    inlineObject263 := *openapiclient.NewInlineObject263(*openapiclient.NewVirtualImageCreate()) // InlineObject263 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddVirtualImage(context.Background()).InlineObject263(inlineObject263).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddVirtualImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualImage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddVirtualImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject263** | [**InlineObject263**](InlineObject263.md) |  | 

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


## AddVirtualImageFile

> Model200Success AddVirtualImageFile(ctx, virtualImageId).Filename(filename).Url(url).Body(body).Execute()

Upload Virtual Image File



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
    virtualImageId := float32(4) // float32 | Virtual Image ID
    filename := "testimage.ovf" // string | The name of the file (optional)
    url := "https://example.com/testimage.ovf" // string | Download the file from a remote url. This can be used instead of uploading a local file. (optional)
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.AddVirtualImageFile(context.Background(), virtualImageId).Filename(filename).Url(url).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.AddVirtualImageFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualImageFile`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.AddVirtualImageFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualImageFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** | The name of the file | 
 **url** | **string** | Download the file from a remote url. This can be used instead of uploading a local file. | 
 **body** | ***os.File** |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileTemplate

> Model200Success DeleteFileTemplate(ctx, id).Execute()

Delete a File Template



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
    resp, r, err := api_client.LibraryApi.DeleteFileTemplate(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteFileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFileTemplate`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteFileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileTemplateRequest struct via the builder pattern


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


## DeleteInstanceType

> Model200Success DeleteInstanceType(ctx, instanceTypeId).Execute()

Delete an Instance Type



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.DeleteInstanceType(context.Background(), instanceTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstanceType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceTypeRequest struct via the builder pattern


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


## DeleteLayout

> Model200Success DeleteLayout(ctx, id).Execute()

Delete a Layout



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
    resp, r, err := api_client.LibraryApi.DeleteLayout(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLayout`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLayoutRequest struct via the builder pattern


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


## DeleteNodeType

> Model200Success DeleteNodeType(ctx, id).Execute()

Delete a Node Type



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
    resp, r, err := api_client.LibraryApi.DeleteNodeType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteNodeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNodeType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteNodeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeTypeRequest struct via the builder pattern


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


## DeleteOptionList

> Model200Success DeleteOptionList(ctx, id).Execute()

Delete an Option List



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
    resp, r, err := api_client.LibraryApi.DeleteOptionList(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteOptionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOptionList`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteOptionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionListRequest struct via the builder pattern


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


## DeleteOptionType

> Model200Success DeleteOptionType(ctx, id).Execute()

Delete an Input



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
    resp, r, err := api_client.LibraryApi.DeleteOptionType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteOptionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOptionType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteOptionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionTypeRequest struct via the builder pattern


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


## DeleteScript

> Model200Success DeleteScript(ctx, id).Execute()

Delete a Script



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
    resp, r, err := api_client.LibraryApi.DeleteScript(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteScript`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScriptRequest struct via the builder pattern


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


## DeleteSpecTemplate

> Model200Success DeleteSpecTemplate(ctx, id).Execute()

Delete a Spec Template



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
    resp, r, err := api_client.LibraryApi.DeleteSpecTemplate(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteSpecTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSpecTemplate`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.DeleteSpecTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpecTemplateRequest struct via the builder pattern


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


## GetFileTemplate

> InlineResponse20070 GetFileTemplate(ctx, id).Execute()

Get a Specific File Template



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
    resp, r, err := api_client.LibraryApi.GetFileTemplate(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetFileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileTemplate`: InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetFileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20070**](inline_response_200_70.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInput

> InlineResponse20075 GetInput(ctx, id).Execute()

Get A Specific Input



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
    resp, r, err := api_client.LibraryApi.GetInput(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetInput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInput`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetInput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20075**](inline_response_200_75.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceType

> InlineResponse20071 GetInstanceType(ctx, instanceTypeId).Execute()

Get a Specific Instance Type



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.GetInstanceType(context.Background(), instanceTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceType`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLayout

> InlineResponse20072 GetLayout(ctx, id).Execute()

Get a Specific Layout



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
    resp, r, err := api_client.LibraryApi.GetLayout(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLayout`: InlineResponse20072
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeType

> InlineResponse20069 GetNodeType(ctx, id).Execute()

Get a Specific Node Type



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
    resp, r, err := api_client.LibraryApi.GetNodeType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetNodeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeType`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetNodeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20069**](inline_response_200_69.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionList

> InlineResponse20073 GetOptionList(ctx, id).Execute()

Get a Specific Option List



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
    resp, r, err := api_client.LibraryApi.GetOptionList(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetOptionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptionList`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetOptionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionListItems

> InlineResponse20074 GetOptionListItems(ctx, id).Execute()

List Items for a Specific Option List



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
    resp, r, err := api_client.LibraryApi.GetOptionListItems(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetOptionListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptionListItems`: InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetOptionListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20074**](inline_response_200_74.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScript

> InlineResponse20068 GetScript(ctx, id).Execute()

Get a Specific Script



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
    resp, r, err := api_client.LibraryApi.GetScript(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScript`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20068**](inline_response_200_68.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityPackageType

> InlineResponse200136 GetSecurityPackageType(ctx, id).Execute()

Retrieves a Specific Security Package Type



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
    resp, r, err := api_client.LibraryApi.GetSecurityPackageType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetSecurityPackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityPackageType`: InlineResponse200136
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetSecurityPackageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPackageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200136**](inline_response_200_136.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecTemplate

> InlineResponse20076 GetSpecTemplate(ctx, id).Execute()

Get a Specific Spec Template



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
    resp, r, err := api_client.LibraryApi.GetSpecTemplate(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetSpecTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecTemplate`: InlineResponse20076
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetSpecTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20076**](inline_response_200_76.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualImage

> InlineResponse200165 GetVirtualImage(ctx, virtualImageId).Execute()

Get a Specific Virtual Image



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
    virtualImageId := float32(4) // float32 | Virtual Image ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.GetVirtualImage(context.Background(), virtualImageId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetVirtualImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualImage`: InlineResponse200165
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetVirtualImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200165**](inline_response_200_165.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFileTemplates

> map[string]interface{} ListFileTemplates(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).FileName(fileName).Execute()

Get All File Templates



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    fileName := "testtext.txt" // string | Filename filter, restricts query to only load file template matching fileName specified (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListFileTemplates(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).FileName(fileName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListFileTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFileTemplates`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListFileTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFileTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **fileName** | **string** | Filename filter, restricts query to only load file template matching fileName specified | 

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


## ListInputs

> map[string]interface{} ListInputs(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Labels(labels).AllLabels(allLabels).FieldName(fieldName).FieldContext(fieldContext).FieldLabel(fieldLabel).Execute()

Get All Inputs



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    fieldName := "cloudName" // string | Field Name filter, restricts query to only load type matching fieldName specified (optional)
    fieldContext := "config.customOptions" // string | Field Context filter, restricts query to only load type matching fieldContext specified (optional)
    fieldLabel := "DB Version" // string | Field Label filter, restricts query to only load type matching fieldLabel specified (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListInputs(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Labels(labels).AllLabels(allLabels).FieldName(fieldName).FieldContext(fieldContext).FieldLabel(fieldLabel).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListInputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInputs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListInputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **fieldName** | **string** | Field Name filter, restricts query to only load type matching fieldName specified | 
 **fieldContext** | **string** | Field Context filter, restricts query to only load type matching fieldContext specified | 
 **fieldLabel** | **string** | Field Label filter, restricts query to only load type matching fieldLabel specified | 

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


## ListInstanceTypes

> map[string]interface{} ListInstanceTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Featured(featured).Labels(labels).AllLabels(allLabels).Details(details).Execute()

Get All Instance Types



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    details := true // bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListInstanceTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Featured(featured).Labels(labels).AllLabels(allLabels).Details(details).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListInstanceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListInstanceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypesRequest struct via the builder pattern


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
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
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


## ListLayouts

> map[string]interface{} ListLayouts(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Layouts



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
    provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListLayouts(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListLayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLayouts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
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


## ListLayoutsForInstanceType

> map[string]interface{} ListLayoutsForInstanceType(ctx, instanceTypeId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Layouts For an Instance Type



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)
    provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListLayoutsForInstanceType(context.Background(), instanceTypeId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListLayoutsForInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLayoutsForInstanceType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListLayoutsForInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLayoutsForInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
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


## ListNodeTypes

> map[string]interface{} ListNodeTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Node Types



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
    provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListNodeTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListNodeTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodeTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListNodeTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNodeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
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


## ListOptionLists

> map[string]interface{} ListOptionLists(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()

Get All Option Lists



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListOptionLists(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListOptionLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOptionLists`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListOptionLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptionListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## ListScripts

> map[string]interface{} ListScripts(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Labels(labels).AllLabels(allLabels).ScriptType(scriptType).ScriptPhase(scriptPhase).Execute()

Get All Scripts



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    scriptType := "scriptType_example" // string | Script type code filter, restricts query to only load scripts of specified type (optional)
    scriptPhase := "scriptPhase_example" // string | Script phase filter, restricts query to only load scripts of specified phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListScripts(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Labels(labels).AllLabels(allLabels).ScriptType(scriptType).ScriptPhase(scriptPhase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScripts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **scriptType** | **string** | Script type code filter, restricts query to only load scripts of specified type | 
 **scriptPhase** | **string** | Script phase filter, restricts query to only load scripts of specified phase | 

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


## ListSecurityPackageTypes

> map[string]interface{} ListSecurityPackageTypes(ctx).Execute()

Retrieves all Security Package Types



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
    resp, r, err := api_client.LibraryApi.ListSecurityPackageTypes(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListSecurityPackageTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityPackageTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListSecurityPackageTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityPackageTypesRequest struct via the builder pattern


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


## ListSpecTemplates

> map[string]interface{} ListSpecTemplates(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()

Get All Spec Templates



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
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListSpecTemplates(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListSpecTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpecTemplates`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListSpecTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpecTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
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


## ListVirtualImageLocations

> map[string]interface{} ListVirtualImageLocations(ctx, virtualImageId).Execute()

Get a List of Virtual Image Locations



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
    virtualImageId := float32(4) // float32 | Virtual Image ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListVirtualImageLocations(context.Background(), virtualImageId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListVirtualImageLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualImageLocations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListVirtualImageLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualImageLocationsRequest struct via the builder pattern


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


## ListVirtualImages

> map[string]interface{} ListVirtualImages(ctx).Max(max).Offset(offset).Name(name).Phrase(phrase).LastUpdated(lastUpdated).FilterType(filterType).ImageType(imageType).Tags(tags).Labels(labels).AllLabels(allLabels).Execute()

Get List of Virtual Images



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
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    filterType := "System" // string | Filter by type, \"User\", \"System\", \"Synced\", or \"All\" (optional) (default to "User")
    imageType := "vmware" // string | Filter by image type code, \"vmware\", \"ami\", etc (optional)
    tags := "tags.env=qa&tags.env=test" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.ListVirtualImages(context.Background()).Max(max).Offset(offset).Name(name).Phrase(phrase).LastUpdated(lastUpdated).FilterType(filterType).ImageType(imageType).Tags(tags).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.ListVirtualImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualImages`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.ListVirtualImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **filterType** | **string** | Filter by type, \&quot;User\&quot;, \&quot;System\&quot;, \&quot;Synced\&quot;, or \&quot;All\&quot; | [default to &quot;User&quot;]
 **imageType** | **string** | Filter by image type code, \&quot;vmware\&quot;, \&quot;ami\&quot;, etc | 
 **tags** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
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


## RemoveSecurityScans

> Model200Success RemoveSecurityScans(ctx, id).Execute()

Deletes a Security Scan



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
    resp, r, err := api_client.LibraryApi.RemoveSecurityScans(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.RemoveSecurityScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecurityScans`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.RemoveSecurityScans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecurityScansRequest struct via the builder pattern


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


## RemoveVirtualImage

> Model200Success RemoveVirtualImage(ctx, virtualImageId).Execute()

Delete a Virtual Image



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
    virtualImageId := float32(4) // float32 | Virtual Image ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.RemoveVirtualImage(context.Background(), virtualImageId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.RemoveVirtualImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVirtualImage`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.RemoveVirtualImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVirtualImageRequest struct via the builder pattern


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


## RemoveVirtualImageFile

> Model200Success RemoveVirtualImageFile(ctx, virtualImageId).Filename(filename).Execute()

Remove Virtual Image File



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
    virtualImageId := float32(4) // float32 | Virtual Image ID
    filename := "testimage.ovf" // string | The name of the file (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.RemoveVirtualImageFile(context.Background(), virtualImageId).Filename(filename).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.RemoveVirtualImageFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVirtualImageFile`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.RemoveVirtualImageFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVirtualImageFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** | The name of the file | 

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


## RemoveVirtualImageLocation

> Model200Success RemoveVirtualImageLocation(ctx, virtualImageId, id).Execute()

Delete a Virtual Image Location



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
    virtualImageId := float32(4) // float32 | Virtual Image ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.RemoveVirtualImageLocation(context.Background(), virtualImageId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.RemoveVirtualImageLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVirtualImageLocation`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.RemoveVirtualImageLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVirtualImageLocationRequest struct via the builder pattern


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


## SetInstanceTypeFeatured

> Model200Success SetInstanceTypeFeatured(ctx, instanceTypeId).Execute()

Toggle Featured For Instance Type



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.SetInstanceTypeFeatured(context.Background(), instanceTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.SetInstanceTypeFeatured``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetInstanceTypeFeatured`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.SetInstanceTypeFeatured`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInstanceTypeFeaturedRequest struct via the builder pattern


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


## UpdateFileTemplate

> Model200Success UpdateFileTemplate(ctx, id).InlineObject112(inlineObject112).Execute()

Update a File Template



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
    inlineObject112 := *openapiclient.NewInlineObject112() // InlineObject112 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateFileTemplate(context.Background(), id).InlineObject112(inlineObject112).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateFileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFileTemplate`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateFileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject112** | [**InlineObject112**](InlineObject112.md) |  | 

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


## UpdateInstanceType

> Model200Success UpdateInstanceType(ctx, instanceTypeId).InlineObject114(inlineObject114).Execute()

Update an Instance Type



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type
    inlineObject114 := *openapiclient.NewInlineObject114() // InlineObject114 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateInstanceType(context.Background(), instanceTypeId).InlineObject114(inlineObject114).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject114** | [**InlineObject114**](InlineObject114.md) |  | 

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


## UpdateInstanceTypeLogo

> Model200Success UpdateInstanceTypeLogo(ctx, instanceTypeId).Logo(logo).DarkLogo(darkLogo).Execute()

Update Logo For Instance Type



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
    instanceTypeId := float32(2) // float32 | The ID of the instance type
    logo := os.NewFile(1234, "some_file") // *os.File | Logo File png,jpg,svg (optional)
    darkLogo := os.NewFile(1234, "some_file") // *os.File | Dark Logo File png,jpg,svg (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateInstanceTypeLogo(context.Background(), instanceTypeId).Logo(logo).DarkLogo(darkLogo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateInstanceTypeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInstanceTypeLogo`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateInstanceTypeLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **float32** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTypeLogoRequest struct via the builder pattern


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


## UpdateLayout

> Model200Success UpdateLayout(ctx, id).InlineObject117(inlineObject117).Execute()

Update a Layout



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
    inlineObject117 := *openapiclient.NewInlineObject117() // InlineObject117 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateLayout(context.Background(), id).InlineObject117(inlineObject117).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLayout`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject117** | [**InlineObject117**](InlineObject117.md) |  | 

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


## UpdateLayoutPermissions

> Model200Success UpdateLayoutPermissions(ctx, id).InlineObject118(inlineObject118).Execute()

Update Layout Permissions



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
    inlineObject118 := *openapiclient.NewInlineObject118() // InlineObject118 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateLayoutPermissions(context.Background(), id).InlineObject118(inlineObject118).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateLayoutPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLayoutPermissions`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateLayoutPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLayoutPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject118** | [**InlineObject118**](InlineObject118.md) |  | 

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


## UpdateNodeType

> Model200Success UpdateNodeType(ctx, id).InlineObject110(inlineObject110).Execute()

Update a Node Type



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
    inlineObject110 := *openapiclient.NewInlineObject110() // InlineObject110 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateNodeType(context.Background(), id).InlineObject110(inlineObject110).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateNodeType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateNodeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject110** | [**InlineObject110**](InlineObject110.md) |  | 

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


## UpdateOptionList

> Model200Success UpdateOptionList(ctx, id).InlineObject120(inlineObject120).Execute()

Update an Option List



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
    inlineObject120 := *openapiclient.NewInlineObject120() // InlineObject120 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateOptionList(context.Background(), id).InlineObject120(inlineObject120).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateOptionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOptionList`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateOptionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject120** | [**InlineObject120**](InlineObject120.md) |  | 

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


## UpdateOptionType

> Model200Success UpdateOptionType(ctx, id).InlineObject122(inlineObject122).Execute()

Update an Input



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
    inlineObject122 := *openapiclient.NewInlineObject122() // InlineObject122 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateOptionType(context.Background(), id).InlineObject122(inlineObject122).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateOptionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOptionType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateOptionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject122** | [**InlineObject122**](InlineObject122.md) |  | 

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


## UpdateScript

> ScriptSuccessId UpdateScript(ctx, id).InlineObject108(inlineObject108).Execute()

Update a Script



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
    inlineObject108 := *openapiclient.NewInlineObject108() // InlineObject108 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateScript(context.Background(), id).InlineObject108(inlineObject108).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScript`: ScriptSuccessId
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject108** | [**InlineObject108**](InlineObject108.md) |  | 

### Return type

[**ScriptSuccessId**](scriptSuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpecTemplate

> Model200Success UpdateSpecTemplate(ctx, id).InlineObject124(inlineObject124).Execute()

Update a Spec Template



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
    inlineObject124 := *openapiclient.NewInlineObject124() // InlineObject124 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateSpecTemplate(context.Background(), id).InlineObject124(inlineObject124).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateSpecTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpecTemplate`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateSpecTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpecTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject124** | [**InlineObject124**](InlineObject124.md) |  | 

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


## UpdateVirtualImage

> map[string]interface{} UpdateVirtualImage(ctx, virtualImageId).InlineObject264(inlineObject264).Execute()

Update a Virtual Image



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
    virtualImageId := float32(4) // float32 | Virtual Image ID
    inlineObject264 := *openapiclient.NewInlineObject264(*openapiclient.NewVirtualImageUpdate()) // InlineObject264 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LibraryApi.UpdateVirtualImage(context.Background(), virtualImageId).InlineObject264(inlineObject264).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.UpdateVirtualImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualImage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.UpdateVirtualImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **float32** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject264** | [**InlineObject264**](InlineObject264.md) |  | 

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

