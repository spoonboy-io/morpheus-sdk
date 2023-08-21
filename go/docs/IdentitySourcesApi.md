# \IdentitySourcesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdentitySources**](IdentitySourcesApi.md#AddIdentitySources) | **Post** /api/user-sources | Creates an Identity Source
[**GetIdentitySources**](IdentitySourcesApi.md#GetIdentitySources) | **Get** /api/user-sources/{id} | Retrieves a Specific Identity Source
[**ListIdentitySources**](IdentitySourcesApi.md#ListIdentitySources) | **Get** /api/user-sources | Retrieves all Identity Sources
[**RemoveIdentitySources**](IdentitySourcesApi.md#RemoveIdentitySources) | **Delete** /api/user-sources/{id} | Deletes an Identity Source
[**UpdateIdentitySourceSubdomains**](IdentitySourcesApi.md#UpdateIdentitySourceSubdomains) | **Put** /api/user-sources/{id}/subdomain | Updates an Identity Source Subdomain
[**UpdateIdentitySources**](IdentitySourcesApi.md#UpdateIdentitySources) | **Put** /api/user-sources/{id} | Updates an Identity Source



## AddIdentitySources

> map[string]interface{} AddIdentitySources(ctx).AccountId(accountId).UserSourceCreate(userSourceCreate).Execute()

Creates an Identity Source



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
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    userSourceCreate := *openapiclient.NewUserSourceCreate(*openapiclient.NewUserSourceCreateUserSource("mydomain AD", "Type_example", *openapiclient.NewUserSourceCreateUserSourceDefaultAccountRole(int64(19)))) // UserSourceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitySourcesApi.AddIdentitySources(context.Background()).AccountId(accountId).UserSourceCreate(userSourceCreate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourcesApi.AddIdentitySources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdentitySources`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentitySourcesApi.AddIdentitySources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIdentitySourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 
 **userSourceCreate** | [**UserSourceCreate**](UserSourceCreate.md) |  | 

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


## GetIdentitySources

> InlineResponse20051 GetIdentitySources(ctx, id).Execute()

Retrieves a Specific Identity Source



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
    resp, r, err := api_client.IdentitySourcesApi.GetIdentitySources(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourcesApi.GetIdentitySources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentitySources`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `IdentitySourcesApi.GetIdentitySources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentitySources

> map[string]interface{} ListIdentitySources(ctx).Type_(type_).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).AccountId(accountId).Execute()

Retrieves all Identity Sources



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
    type_ := "type__example" // string | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.  (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitySourcesApi.ListIdentitySources(context.Background()).Type_(type_).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).AccountId(accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourcesApi.ListIdentitySources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentitySources`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentitySourcesApi.ListIdentitySources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitySourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 

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


## RemoveIdentitySources

> Model200Success RemoveIdentitySources(ctx, id).Execute()

Deletes an Identity Source



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
    resp, r, err := api_client.IdentitySourcesApi.RemoveIdentitySources(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourcesApi.RemoveIdentitySources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveIdentitySources`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `IdentitySourcesApi.RemoveIdentitySources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIdentitySourcesRequest struct via the builder pattern


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


## UpdateIdentitySourceSubdomains

> map[string]interface{} UpdateIdentitySourceSubdomains(ctx, id).InlineObject82(inlineObject82).Execute()

Updates an Identity Source Subdomain



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
    inlineObject82 := *openapiclient.NewInlineObject82("mynewdomain") // InlineObject82 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitySourcesApi.UpdateIdentitySourceSubdomains(context.Background(), id).InlineObject82(inlineObject82).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourcesApi.UpdateIdentitySourceSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentitySourceSubdomains`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentitySourcesApi.UpdateIdentitySourceSubdomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitySourceSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject82** | [**InlineObject82**](InlineObject82.md) |  | 

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


## UpdateIdentitySources

> map[string]interface{} UpdateIdentitySources(ctx, id).UserSourceCreate(userSourceCreate).Execute()

Updates an Identity Source



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
    userSourceCreate := *openapiclient.NewUserSourceCreate(*openapiclient.NewUserSourceCreateUserSource("mydomain AD", "Type_example", *openapiclient.NewUserSourceCreateUserSourceDefaultAccountRole(int64(19)))) // UserSourceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentitySourcesApi.UpdateIdentitySources(context.Background(), id).UserSourceCreate(userSourceCreate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourcesApi.UpdateIdentitySources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentitySources`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentitySourcesApi.UpdateIdentitySources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitySourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSourceCreate** | [**UserSourceCreate**](UserSourceCreate.md) |  | 

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

