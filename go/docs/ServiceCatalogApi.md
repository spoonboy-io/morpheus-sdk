# \ServiceCatalogApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCatalogCart**](ServiceCatalogApi.md#AddCatalogCart) | **Post** /api/catalog/checkout | Checkout Catalog Cart
[**AddCatalogCartItem**](ServiceCatalogApi.md#AddCatalogCartItem) | **Put** /api/catalog/cart/items | Add Catalog Item to Cart
[**AddCatalogOrder**](ServiceCatalogApi.md#AddCatalogOrder) | **Post** /api/catalog/orders | Place Catalog Order
[**DeleteCatalogCart**](ServiceCatalogApi.md#DeleteCatalogCart) | **Delete** /api/catalog/cart | Clear Catalog Cart
[**DeleteCatalogCartItem**](ServiceCatalogApi.md#DeleteCatalogCartItem) | **Delete** /api/catalog/cart/items/{id} | Remove a Catalog Item From Cart
[**DeleteCatalogItem**](ServiceCatalogApi.md#DeleteCatalogItem) | **Delete** /api/catalog/items/{id} | Delete a Catalog Inventory Item
[**GetCatalogItem**](ServiceCatalogApi.md#GetCatalogItem) | **Get** /api/catalog/items/{id} | Get a Specific Catalog Inventory Item
[**GetCatalogType**](ServiceCatalogApi.md#GetCatalogType) | **Get** /api/catalog/types/{id} | Get a Specific Catalog Type
[**ListCatalogCart**](ServiceCatalogApi.md#ListCatalogCart) | **Get** /api/catalog/cart | Get Catalog Cart
[**ListCatalogItems**](ServiceCatalogApi.md#ListCatalogItems) | **Get** /api/catalog/items | List Catalog Inventory Items
[**ListCatalogTypes**](ServiceCatalogApi.md#ListCatalogTypes) | **Get** /api/catalog/types | List Catalog Types



## AddCatalogCart

> map[string]interface{} AddCatalogCart(ctx).Execute()

Checkout Catalog Cart



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
    resp, r, err := api_client.ServiceCatalogApi.AddCatalogCart(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.AddCatalogCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCatalogCart`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.AddCatalogCart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddCatalogCartRequest struct via the builder pattern


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


## AddCatalogCartItem

> map[string]interface{} AddCatalogCartItem(ctx).Validate(validate).CatalogCartItemCreate(catalogCartItemCreate).Execute()

Add Catalog Item to Cart



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
    validate := true // bool | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory (optional) (default to false)
    catalogCartItemCreate := *openapiclient.NewCatalogCartItemCreate() // CatalogCartItemCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceCatalogApi.AddCatalogCartItem(context.Background()).Validate(validate).CatalogCartItemCreate(catalogCartItemCreate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.AddCatalogCartItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCatalogCartItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.AddCatalogCartItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCatalogCartItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool** | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [default to false]
 **catalogCartItemCreate** | [**CatalogCartItemCreate**](CatalogCartItemCreate.md) |  | 

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


## AddCatalogOrder

> map[string]interface{} AddCatalogOrder(ctx).Validate(validate).InlineObject227(inlineObject227).Execute()

Place Catalog Order



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
    validate := true // bool | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory (optional) (default to false)
    inlineObject227 := *openapiclient.NewInlineObject227() // InlineObject227 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceCatalogApi.AddCatalogOrder(context.Background()).Validate(validate).InlineObject227(inlineObject227).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.AddCatalogOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCatalogOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.AddCatalogOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCatalogOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool** | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [default to false]
 **inlineObject227** | [**InlineObject227**](InlineObject227.md) |  | 

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


## DeleteCatalogCart

> Model200Success DeleteCatalogCart(ctx).Execute()

Clear Catalog Cart



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
    resp, r, err := api_client.ServiceCatalogApi.DeleteCatalogCart(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.DeleteCatalogCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCatalogCart`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.DeleteCatalogCart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogCartRequest struct via the builder pattern


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


## DeleteCatalogCartItem

> Model200Success DeleteCatalogCartItem(ctx, id).Execute()

Remove a Catalog Item From Cart



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
    resp, r, err := api_client.ServiceCatalogApi.DeleteCatalogCartItem(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.DeleteCatalogCartItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCatalogCartItem`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.DeleteCatalogCartItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogCartItemRequest struct via the builder pattern


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


## DeleteCatalogItem

> Model200Success DeleteCatalogItem(ctx, id).PreserveVolumes(preserveVolumes).KeepBackups(keepBackups).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).RemoveInstances(removeInstances).Force(force).Execute()

Delete a Catalog Inventory Item



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
    preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
    keepBackups := "on" // string | Preserve copy of backups (optional) (default to "off")
    releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
    releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
    removeInstances := "off" // string | Remove Instances. Only applies to type `blueprint` (Apps) (optional) (default to "on")
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceCatalogApi.DeleteCatalogItem(context.Background(), id).PreserveVolumes(preserveVolumes).KeepBackups(keepBackups).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).RemoveInstances(removeInstances).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.DeleteCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCatalogItem`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.DeleteCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
 **keepBackups** | **string** | Preserve copy of backups | [default to &quot;off&quot;]
 **releaseFloatingIps** | **string** | Release Floating IPs | [default to &quot;on&quot;]
 **releaseEIPs** | **string** | Alias for releaseFloatingIps | [default to &quot;on&quot;]
 **removeInstances** | **string** | Remove Instances. Only applies to type &#x60;blueprint&#x60; (Apps) | [default to &quot;on&quot;]
 **force** | **string** | Force Delete | [default to &quot;off&quot;]

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


## GetCatalogItem

> InlineResponse200140 GetCatalogItem(ctx, id).Execute()

Get a Specific Catalog Inventory Item



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
    resp, r, err := api_client.ServiceCatalogApi.GetCatalogItem(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.GetCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogItem`: InlineResponse200140
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.GetCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200140**](inline_response_200_140.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogType

> map[string]interface{} GetCatalogType(ctx, id).Execute()

Get a Specific Catalog Type



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
    resp, r, err := api_client.ServiceCatalogApi.GetCatalogType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.GetCatalogType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.GetCatalogType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogTypeRequest struct via the builder pattern


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


## ListCatalogCart

> InlineResponse200139 ListCatalogCart(ctx).Execute()

Get Catalog Cart



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
    resp, r, err := api_client.ServiceCatalogApi.ListCatalogCart(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.ListCatalogCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogCart`: InlineResponse200139
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.ListCatalogCart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogCartRequest struct via the builder pattern


### Return type

[**InlineResponse200139**](inline_response_200_139.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogItems

> map[string]interface{} ListCatalogItems(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

List Catalog Inventory Items



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceCatalogApi.ListCatalogItems(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.ListCatalogItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogItems`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.ListCatalogItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## ListCatalogTypes

> map[string]interface{} ListCatalogTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Featured(featured).Execute()

List Catalog Types



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
    featured := false // bool | Filter by featured (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceCatalogApi.ListCatalogTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Featured(featured).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCatalogApi.ListCatalogTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceCatalogApi.ListCatalogTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **featured** | **bool** | Filter by featured | 

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

