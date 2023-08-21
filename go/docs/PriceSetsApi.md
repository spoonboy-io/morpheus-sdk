# \PriceSetsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPriceSets**](PriceSetsApi.md#AddPriceSets) | **Post** /api/price-sets | Creates a Price Set
[**DeactivatePriceSets**](PriceSetsApi.md#DeactivatePriceSets) | **Put** /api/price-sets/{id}/deactivate | Deactivates a Price Set
[**GetPriceSets**](PriceSetsApi.md#GetPriceSets) | **Get** /api/price-sets/{id} | Retrieves a Specific Price Set
[**ListPriceSets**](PriceSetsApi.md#ListPriceSets) | **Get** /api/price-sets | Retrieves all Price Sets
[**UpdatePriceSets**](PriceSetsApi.md#UpdatePriceSets) | **Put** /api/price-sets/{id} | Updates a Price Set



## AddPriceSets

> map[string]interface{} AddPriceSets(ctx).InlineObject198(inlineObject198).Execute()

Creates a Price Set



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
    inlineObject198 := *openapiclient.NewInlineObject198(*openapiclient.NewApiPriceSetsPriceSet("testName", "priceSet1", "PriceUnit_example", "Type_example")) // InlineObject198 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PriceSetsApi.AddPriceSets(context.Background()).InlineObject198(inlineObject198).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceSetsApi.AddPriceSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPriceSets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PriceSetsApi.AddPriceSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPriceSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject198** | [**InlineObject198**](InlineObject198.md) |  | 

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


## DeactivatePriceSets

> Model200Success DeactivatePriceSets(ctx, id).Execute()

Deactivates a Price Set



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
    resp, r, err := api_client.PriceSetsApi.DeactivatePriceSets(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceSetsApi.DeactivatePriceSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivatePriceSets`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `PriceSetsApi.DeactivatePriceSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePriceSetsRequest struct via the builder pattern


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


## GetPriceSets

> InlineResponse200123 GetPriceSets(ctx, id).Execute()

Retrieves a Specific Price Set



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
    resp, r, err := api_client.PriceSetsApi.GetPriceSets(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceSetsApi.GetPriceSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriceSets`: InlineResponse200123
    fmt.Fprintf(os.Stdout, "Response from `PriceSetsApi.GetPriceSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200123**](inline_response_200_123.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPriceSets

> map[string]interface{} ListPriceSets(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).IncludeInactive(includeInactive).Type_(type_).Execute()

Retrieves all Price Sets



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
    includeInactive := true // bool | If true, include inactive prices in the results (optional)
    type_ := "type__example" // string | Filter by type code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PriceSetsApi.ListPriceSets(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).IncludeInactive(includeInactive).Type_(type_).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceSetsApi.ListPriceSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPriceSets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PriceSetsApi.ListPriceSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPriceSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **includeInactive** | **bool** | If true, include inactive prices in the results | 
 **type_** | **string** | Filter by type code | 

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


## UpdatePriceSets

> map[string]interface{} UpdatePriceSets(ctx, id).InlineObject199(inlineObject199).Execute()

Updates a Price Set



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
    inlineObject199 := *openapiclient.NewInlineObject199(*openapiclient.NewApiPriceSetsIdPriceSet()) // InlineObject199 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PriceSetsApi.UpdatePriceSets(context.Background(), id).InlineObject199(inlineObject199).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceSetsApi.UpdatePriceSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePriceSets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PriceSetsApi.UpdatePriceSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject199** | [**InlineObject199**](InlineObject199.md) |  | 

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

