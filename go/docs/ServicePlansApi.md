# \ServicePlansApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateServicePlans**](ServicePlansApi.md#ActivateServicePlans) | **Put** /api/service-plans/{id}/activate | Activates a Service Plan
[**AddServicePlans**](ServicePlansApi.md#AddServicePlans) | **Post** /api/service-plans | Creates a Service Plan
[**DeactivateServicePlans**](ServicePlansApi.md#DeactivateServicePlans) | **Put** /api/service-plans/{id}/deactivate | Deactivates a Service Plan
[**GetServicePlans**](ServicePlansApi.md#GetServicePlans) | **Get** /api/service-plans/{id} | Retrieves a Specific Service Plan
[**ListServicePlans**](ServicePlansApi.md#ListServicePlans) | **Get** /api/service-plans | Retrieves all Service Plans
[**RemoveServicePlans**](ServicePlansApi.md#RemoveServicePlans) | **Delete** /api/service-plans/{id} | Deletes a Service Plan
[**UpdateServicePlans**](ServicePlansApi.md#UpdateServicePlans) | **Put** /api/service-plans/{id} | Updates a Service Plan



## ActivateServicePlans

> Model200Success ActivateServicePlans(ctx, id).Execute()

Activates a Service Plan



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
    resp, r, err := api_client.ServicePlansApi.ActivateServicePlans(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.ActivateServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateServicePlans`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.ActivateServicePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateServicePlansRequest struct via the builder pattern


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


## AddServicePlans

> map[string]interface{} AddServicePlans(ctx).InlineObject228(inlineObject228).Execute()

Creates a Service Plan



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
    inlineObject228 := *openapiclient.NewInlineObject228(*openapiclient.NewApiServicePlansServicePlan("Name_example", "Code_example", float32(123), float32(123), []openapiclient.ApiServicePlansServicePlanProvisionType{*openapiclient.NewApiServicePlansServicePlanProvisionType(int64(123))})) // InlineObject228 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePlansApi.AddServicePlans(context.Background()).InlineObject228(inlineObject228).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.AddServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServicePlans`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.AddServicePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject228** | [**InlineObject228**](InlineObject228.md) |  | 

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


## DeactivateServicePlans

> Model200Success DeactivateServicePlans(ctx, id).Execute()

Deactivates a Service Plan



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
    resp, r, err := api_client.ServicePlansApi.DeactivateServicePlans(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.DeactivateServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateServicePlans`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.DeactivateServicePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateServicePlansRequest struct via the builder pattern


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


## GetServicePlans

> InlineResponse200142 GetServicePlans(ctx, id).Execute()

Retrieves a Specific Service Plan



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
    resp, r, err := api_client.ServicePlansApi.GetServicePlans(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.GetServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicePlans`: InlineResponse200142
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.GetServicePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200142**](inline_response_200_142.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServicePlans

> map[string]interface{} ListServicePlans(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).IncludeZones(includeZones).ProvisionTypeId(provisionTypeId).IncludeInactive(includeInactive).Execute()

Retrieves all Service Plans



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
    includeZones := true // bool | Indicates including zones in the service plan response payload (optional) (default to false)
    provisionTypeId := int64(22) // int64 | Provision type filter, restricts query to only load service plans of specified provision type (optional)
    includeInactive := true // bool | If true, include inactive prices in the results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePlansApi.ListServicePlans(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).IncludeZones(includeZones).ProvisionTypeId(provisionTypeId).IncludeInactive(includeInactive).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.ListServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServicePlans`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.ListServicePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **includeZones** | **bool** | Indicates including zones in the service plan response payload | [default to false]
 **provisionTypeId** | **int64** | Provision type filter, restricts query to only load service plans of specified provision type | 
 **includeInactive** | **bool** | If true, include inactive prices in the results | 

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


## RemoveServicePlans

> Model200Success RemoveServicePlans(ctx, id).Execute()

Deletes a Service Plan



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
    resp, r, err := api_client.ServicePlansApi.RemoveServicePlans(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.RemoveServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveServicePlans`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.RemoveServicePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServicePlansRequest struct via the builder pattern


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


## UpdateServicePlans

> map[string]interface{} UpdateServicePlans(ctx, id).InlineObject229(inlineObject229).Execute()

Updates a Service Plan



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
    inlineObject229 := *openapiclient.NewInlineObject229(*openapiclient.NewApiServicePlansIdServicePlan()) // InlineObject229 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePlansApi.UpdateServicePlans(context.Background(), id).InlineObject229(inlineObject229).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePlansApi.UpdateServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServicePlans`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicePlansApi.UpdateServicePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject229** | [**InlineObject229**](InlineObject229.md) |  | 

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

