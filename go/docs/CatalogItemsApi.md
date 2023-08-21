# \CatalogItemsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCatalogItemType**](CatalogItemsApi.md#AddCatalogItemType) | **Post** /api/catalog-item-types | Create a Catalog Item Type
[**GetCatalogItemType**](CatalogItemsApi.md#GetCatalogItemType) | **Get** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type
[**ListCatalogItemTypes**](CatalogItemsApi.md#ListCatalogItemTypes) | **Get** /api/catalog-item-types | Get All Catalog Item Types
[**RemoveCatalogItemType**](CatalogItemsApi.md#RemoveCatalogItemType) | **Delete** /api/catalog-item-types/{id} | Delete a Catalog Item Type
[**UpdateCatalogItemType**](CatalogItemsApi.md#UpdateCatalogItemType) | **Put** /api/catalog-item-types/{id} | Update a Catalog Item Type
[**UpdateCatalogItemTypeLogo**](CatalogItemsApi.md#UpdateCatalogItemTypeLogo) | **Put** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type



## AddCatalogItemType

> map[string]interface{} AddCatalogItemType(ctx).InlineObject24(inlineObject24).Execute()

Create a Catalog Item Type



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
    inlineObject24 := *openapiclient.NewInlineObject24() // InlineObject24 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogItemsApi.AddCatalogItemType(context.Background()).InlineObject24(inlineObject24).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogItemsApi.AddCatalogItemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCatalogItemType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CatalogItemsApi.AddCatalogItemType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCatalogItemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject24** | [**InlineObject24**](InlineObject24.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemType

> InlineResponse20015 GetCatalogItemType(ctx, id).Execute()

Get a Specific Catalog Item Type



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
    resp, r, err := api_client.CatalogItemsApi.GetCatalogItemType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogItemsApi.GetCatalogItemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogItemType`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `CatalogItemsApi.GetCatalogItemType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogItemTypes

> map[string]interface{} ListCatalogItemTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Enabled(enabled).Featured(featured).Labels(labels).AllLabels(allLabels).Code(code).Execute()

Get All Catalog Item Types



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
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    enabled := false // bool | Filter by enabled (optional)
    featured := false // bool | Filter by featured (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogItemsApi.ListCatalogItemTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Description(description).Enabled(enabled).Featured(featured).Labels(labels).AllLabels(allLabels).Code(code).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogItemsApi.ListCatalogItemTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogItemTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CatalogItemsApi.ListCatalogItemTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogItemTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 
 **enabled** | **bool** | Filter by enabled | 
 **featured** | **bool** | Filter by featured | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **code** | **string** | If specified will return an exact match on code | 

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


## RemoveCatalogItemType

> Model200Success RemoveCatalogItemType(ctx, id).Execute()

Delete a Catalog Item Type



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
    resp, r, err := api_client.CatalogItemsApi.RemoveCatalogItemType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogItemsApi.RemoveCatalogItemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCatalogItemType`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `CatalogItemsApi.RemoveCatalogItemType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCatalogItemTypeRequest struct via the builder pattern


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


## UpdateCatalogItemType

> map[string]interface{} UpdateCatalogItemType(ctx, id).InlineObject25(inlineObject25).Execute()

Update a Catalog Item Type



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
    inlineObject25 := *openapiclient.NewInlineObject25() // InlineObject25 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogItemsApi.UpdateCatalogItemType(context.Background(), id).InlineObject25(inlineObject25).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogItemsApi.UpdateCatalogItemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogItemType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CatalogItemsApi.UpdateCatalogItemType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject25** | [**InlineObject25**](InlineObject25.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItemTypeLogo

> map[string]interface{} UpdateCatalogItemTypeLogo(ctx, id).CatalogItemTypeLogo(catalogItemTypeLogo).CatalogItemTypeDarkLogo(catalogItemTypeDarkLogo).Execute()

Update Logo For Catalog Item Type



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
    catalogItemTypeLogo := os.NewFile(1234, "some_file") // *os.File | Logo File png,jpg,svg (optional)
    catalogItemTypeDarkLogo := os.NewFile(1234, "some_file") // *os.File | Dark Logo File png,jpg,svg (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogItemsApi.UpdateCatalogItemTypeLogo(context.Background(), id).CatalogItemTypeLogo(catalogItemTypeLogo).CatalogItemTypeDarkLogo(catalogItemTypeDarkLogo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogItemsApi.UpdateCatalogItemTypeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogItemTypeLogo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CatalogItemsApi.UpdateCatalogItemTypeLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemTypeLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catalogItemTypeLogo** | ***os.File** | Logo File png,jpg,svg | 
 **catalogItemTypeDarkLogo** | ***os.File** | Dark Logo File png,jpg,svg | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

