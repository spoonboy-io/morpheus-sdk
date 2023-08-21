# \SSLCertificatesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCertificate**](SSLCertificatesApi.md#AddCertificate) | **Post** /api/certificates | Create a Certificate
[**DeleteCertificate**](SSLCertificatesApi.md#DeleteCertificate) | **Delete** /api/certificates/{id} | Delete a Certificate
[**GetCertificate**](SSLCertificatesApi.md#GetCertificate) | **Get** /api/certificates/{id} | Get a Specific Certificate
[**ListCertificates**](SSLCertificatesApi.md#ListCertificates) | **Get** /api/certificates | Get All SSL Certificates
[**UpdateCertificate**](SSLCertificatesApi.md#UpdateCertificate) | **Put** /api/certificates/{id} | Update a Certificate



## AddCertificate

> map[string]interface{} AddCertificate(ctx).InlineObject230(inlineObject230).Execute()

Create a Certificate



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
    inlineObject230 := *openapiclient.NewInlineObject230() // InlineObject230 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLCertificatesApi.AddCertificate(context.Background()).InlineObject230(inlineObject230).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.AddCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.AddCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject230** | [**InlineObject230**](InlineObject230.md) |  | 

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


## DeleteCertificate

> Model200Success DeleteCertificate(ctx, id).Execute()

Delete a Certificate



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
    resp, r, err := api_client.SSLCertificatesApi.DeleteCertificate(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.DeleteCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertificate`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.DeleteCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


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


## GetCertificate

> InlineResponse200144 GetCertificate(ctx, id).Execute()

Get a Specific Certificate



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
    resp, r, err := api_client.SSLCertificatesApi.GetCertificate(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.GetCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificate`: InlineResponse200144
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200144**](inline_response_200_144.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificates

> InlineResponse200143 ListCertificates(ctx).Name(name).Execute()

Get All SSL Certificates



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLCertificatesApi.ListCertificates(context.Background()).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.ListCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCertificates`: InlineResponse200143
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.ListCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 

### Return type

[**InlineResponse200143**](inline_response_200_143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> InlineResponse200144 UpdateCertificate(ctx, id).InlineObject231(inlineObject231).Execute()

Update a Certificate



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
    inlineObject231 := *openapiclient.NewInlineObject231() // InlineObject231 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLCertificatesApi.UpdateCertificate(context.Background(), id).InlineObject231(inlineObject231).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.UpdateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificate`: InlineResponse200144
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.UpdateCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject231** | [**InlineObject231**](InlineObject231.md) |  | 

### Return type

[**InlineResponse200144**](inline_response_200_144.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

