# \SecurityPackagesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecurityPackages**](SecurityPackagesApi.md#AddSecurityPackages) | **Post** /api/security-packages | Creates a Security Package
[**GetSecurityPackages**](SecurityPackagesApi.md#GetSecurityPackages) | **Get** /api/security-packages/{id} | Retrieves a Specific Security Package
[**ListSecurityPackages**](SecurityPackagesApi.md#ListSecurityPackages) | **Get** /api/security-packages | Retrieves all Security Packages
[**RemoveSecurityPackages**](SecurityPackagesApi.md#RemoveSecurityPackages) | **Delete** /api/security-packages/{id} | Deletes a Security Package
[**UpdateSecurityPackages**](SecurityPackagesApi.md#UpdateSecurityPackages) | **Put** /api/security-packages/{id} | Updates a Security Package



## AddSecurityPackages

> map[string]interface{} AddSecurityPackages(ctx).InlineObject218(inlineObject218).Execute()

Creates a Security Package



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
    inlineObject218 := *openapiclient.NewInlineObject218(*openapiclient.NewApiSecurityPackagesSecurityPackage("Sample Security Package", "http://10.0.2.2:8080/public-archives/download/SCAP/scap-security-guide-0.1.51.zip")) // InlineObject218 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityPackagesApi.AddSecurityPackages(context.Background()).InlineObject218(inlineObject218).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPackagesApi.AddSecurityPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSecurityPackages`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityPackagesApi.AddSecurityPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSecurityPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject218** | [**InlineObject218**](InlineObject218.md) |  | 

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


## GetSecurityPackages

> map[string]interface{} GetSecurityPackages(ctx, id).Execute()

Retrieves a Specific Security Package



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
    resp, r, err := api_client.SecurityPackagesApi.GetSecurityPackages(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPackagesApi.GetSecurityPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityPackages`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityPackagesApi.GetSecurityPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPackagesRequest struct via the builder pattern


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


## ListSecurityPackages

> map[string]interface{} ListSecurityPackages(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()

Retrieves all Security Packages



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
    resp, r, err := api_client.SecurityPackagesApi.ListSecurityPackages(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPackagesApi.ListSecurityPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityPackages`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityPackagesApi.ListSecurityPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityPackagesRequest struct via the builder pattern


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


## RemoveSecurityPackages

> Model200Success RemoveSecurityPackages(ctx, id).Execute()

Deletes a Security Package



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
    resp, r, err := api_client.SecurityPackagesApi.RemoveSecurityPackages(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPackagesApi.RemoveSecurityPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecurityPackages`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `SecurityPackagesApi.RemoveSecurityPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecurityPackagesRequest struct via the builder pattern


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


## UpdateSecurityPackages

> map[string]interface{} UpdateSecurityPackages(ctx, id).InlineObject219(inlineObject219).Execute()

Updates a Security Package



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
    inlineObject219 := *openapiclient.NewInlineObject219(*openapiclient.NewApiSecurityPackagesIdSecurityPackage()) // InlineObject219 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityPackagesApi.UpdateSecurityPackages(context.Background(), id).InlineObject219(inlineObject219).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityPackagesApi.UpdateSecurityPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityPackages`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityPackagesApi.UpdateSecurityPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject219** | [**InlineObject219**](InlineObject219.md) |  | 

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

