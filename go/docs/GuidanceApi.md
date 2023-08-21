# \GuidanceApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteGuidances**](GuidanceApi.md#ExecuteGuidances) | **Put** /api/guidance/{id}/execute | Executes a Specific Guidance Recommendation
[**GetGuidanceStats**](GuidanceApi.md#GetGuidanceStats) | **Get** /api/guidance/stats | Retrieves Guidance Stats
[**GetGuidanceTypes**](GuidanceApi.md#GetGuidanceTypes) | **Get** /api/guidance/types | Retrieves Guidance Types
[**GetGuidances**](GuidanceApi.md#GetGuidances) | **Get** /api/guidance/{id} | Retrieves a Specific Guidance Recommendation
[**IgnoreGuidances**](GuidanceApi.md#IgnoreGuidances) | **Put** /api/guidance/{id}/ignore | Ignores a Specific Guidance Recommendation
[**ListGuidances**](GuidanceApi.md#ListGuidances) | **Get** /api/guidance | Retrieves all Guidance Recommendations



## ExecuteGuidances

> Model200Success ExecuteGuidances(ctx, id).Execute()

Executes a Specific Guidance Recommendation



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
    resp, r, err := api_client.GuidanceApi.ExecuteGuidances(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceApi.ExecuteGuidances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteGuidances`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `GuidanceApi.ExecuteGuidances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteGuidancesRequest struct via the builder pattern


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


## GetGuidanceStats

> InlineResponse20045 GetGuidanceStats(ctx).Execute()

Retrieves Guidance Stats



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
    resp, r, err := api_client.GuidanceApi.GetGuidanceStats(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceApi.GetGuidanceStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGuidanceStats`: InlineResponse20045
    fmt.Fprintf(os.Stdout, "Response from `GuidanceApi.GetGuidanceStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuidanceStatsRequest struct via the builder pattern


### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuidanceTypes

> InlineResponse20046 GetGuidanceTypes(ctx).Execute()

Retrieves Guidance Types



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
    resp, r, err := api_client.GuidanceApi.GetGuidanceTypes(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceApi.GetGuidanceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGuidanceTypes`: InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `GuidanceApi.GetGuidanceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuidanceTypesRequest struct via the builder pattern


### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuidances

> InlineResponse20044 GetGuidances(ctx, id).Execute()

Retrieves a Specific Guidance Recommendation



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
    resp, r, err := api_client.GuidanceApi.GetGuidances(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceApi.GetGuidances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGuidances`: InlineResponse20044
    fmt.Fprintf(os.Stdout, "Response from `GuidanceApi.GetGuidances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuidancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IgnoreGuidances

> Model200Success IgnoreGuidances(ctx, id).Execute()

Ignores a Specific Guidance Recommendation



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
    resp, r, err := api_client.GuidanceApi.IgnoreGuidances(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceApi.IgnoreGuidances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IgnoreGuidances`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `GuidanceApi.IgnoreGuidances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiIgnoreGuidancesRequest struct via the builder pattern


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


## ListGuidances

> map[string]interface{} ListGuidances(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Severity(severity).Type_(type_).State(state).Execute()

Retrieves all Guidance Recommendations



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
    severity := "severity_example" // string | Filter by severity (optional)
    type_ := "type__example" // string | Filter by Guidance Type.  See `Retrieves Guidance Types` for most up to date list of types. (optional)
    state := "state_example" // string | Filter by state, restricts query to only load discoveries by state (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GuidanceApi.ListGuidances(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Severity(severity).Type_(type_).State(state).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuidanceApi.ListGuidances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGuidances`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GuidanceApi.ListGuidances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGuidancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **severity** | **string** | Filter by severity | 
 **type_** | **string** | Filter by Guidance Type.  See &#x60;Retrieves Guidance Types&#x60; for most up to date list of types. | 
 **state** | **string** | Filter by state, restricts query to only load discoveries by state | 

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

