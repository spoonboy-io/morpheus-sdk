# \IntegrationsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIntegrationSnowObjects**](IntegrationsApi.md#AddIntegrationSnowObjects) | **Post** /api/integrations/{id}/objects | Creates an Exposed ServiceNow Catalog Item
[**AddIntegrations**](IntegrationsApi.md#AddIntegrations) | **Post** /api/integrations | Creates an Integration
[**GetIntegrationInventory**](IntegrationsApi.md#GetIntegrationInventory) | **Get** /api/integrations/{id}/inventory/{inventoryId} | Get a Specific Integration Inventory
[**GetIntegrationObjects**](IntegrationsApi.md#GetIntegrationObjects) | **Get** /api/integrations/{id}/objects/{objectId} | Get a Specific ServiceNow Integration Object
[**GetIntegrationTypeOptionTypes**](IntegrationsApi.md#GetIntegrationTypeOptionTypes) | **Get** /api/integration-types/{id}/option-types | Retrieves a Option Types for a Specific Integration Type
[**GetIntegrationTypes**](IntegrationsApi.md#GetIntegrationTypes) | **Get** /api/integration-types/{id} | Retrieves a Specific Integration Type
[**GetIntegrations**](IntegrationsApi.md#GetIntegrations) | **Get** /api/integrations/{id} | Retrieves a Specific Integration
[**ListIntegrationInventory**](IntegrationsApi.md#ListIntegrationInventory) | **Get** /api/integrations/{id}/inventory | Get All Integration Inventory
[**ListIntegrationObjects**](IntegrationsApi.md#ListIntegrationObjects) | **Get** /api/integrations/{id}/objects | Get ServiceNow Integration Objects
[**ListIntegrationTypes**](IntegrationsApi.md#ListIntegrationTypes) | **Get** /api/integration-types | Retrieves all Integration Types
[**ListIntegrations**](IntegrationsApi.md#ListIntegrations) | **Get** /api/integrations | Retrieves all Integrations
[**RefreshIntegrations**](IntegrationsApi.md#RefreshIntegrations) | **Post** /api/integrations/{id}/refresh | Refresh an Integration
[**RemoveIntegrationObjects**](IntegrationsApi.md#RemoveIntegrationObjects) | **Delete** /api/integrations/{id}/objects/{objectId} | Deletes a ServiceNow Integration object
[**RemoveIntegrations**](IntegrationsApi.md#RemoveIntegrations) | **Delete** /api/integrations/{id} | Deletes an Integration
[**UpdateIntegrationInventory**](IntegrationsApi.md#UpdateIntegrationInventory) | **Put** /api/integrations/{id}/inventory/{inventoryId} | Updating an Integration Inventory
[**UpdateIntegrations**](IntegrationsApi.md#UpdateIntegrations) | **Put** /api/integrations/{id} | Updates an Integration



## AddIntegrationSnowObjects

> map[string]interface{} AddIntegrationSnowObjects(ctx, id).InlineObject100(inlineObject100).Execute()

Creates an Exposed ServiceNow Catalog Item



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
    inlineObject100 := *openapiclient.NewInlineObject100() // InlineObject100 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.AddIntegrationSnowObjects(context.Background(), id).InlineObject100(inlineObject100).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.AddIntegrationSnowObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIntegrationSnowObjects`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.AddIntegrationSnowObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIntegrationSnowObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject100** | [**InlineObject100**](InlineObject100.md) |  | 

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


## AddIntegrations

> map[string]interface{} AddIntegrations(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Creates an Integration



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.AddIntegrations(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.AddIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIntegrations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.AddIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

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


## GetIntegrationInventory

> InlineResponse20065 GetIntegrationInventory(ctx, id, inventoryId).Execute()

Get a Specific Integration Inventory



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
    id := int64(9) // int64 | Morpheus ID of the Integration
    inventoryId := int64(1) // int64 | Morpheus ID of the Integration Inventory

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.GetIntegrationInventory(context.Background(), id, inventoryId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationInventory`: InlineResponse20065
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Integration | 
**inventoryId** | **int64** | Morpheus ID of the Integration Inventory | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20065**](inline_response_200_65.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationObjects

> InlineResponse20064 GetIntegrationObjects(ctx, id, objectId).Execute()

Get a Specific ServiceNow Integration Object



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
    objectId := int64(122) // int64 | Morpheus ID of the Object being created or referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.GetIntegrationObjects(context.Background(), id, objectId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationObjects`: InlineResponse20064
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**objectId** | **int64** | Morpheus ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20064**](inline_response_200_64.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationTypeOptionTypes

> InlineResponse20062 GetIntegrationTypeOptionTypes(ctx, id).Execute()

Retrieves a Option Types for a Specific Integration Type



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
    resp, r, err := api_client.IntegrationsApi.GetIntegrationTypeOptionTypes(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationTypeOptionTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationTypeOptionTypes`: InlineResponse20062
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationTypeOptionTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTypeOptionTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationTypes

> InlineResponse20061 GetIntegrationTypes(ctx, id).Optiontypes(optiontypes).Execute()

Retrieves a Specific Integration Type



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
    optiontypes := true // bool | Pass `true` to include optionTypes in the response for each integration type (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.GetIntegrationTypes(context.Background(), id).Optiontypes(optiontypes).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationTypes`: InlineResponse20061
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **optiontypes** | **bool** | Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [default to false]

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrations

> map[string]interface{} GetIntegrations(ctx, id).Execute()

Retrieves a Specific Integration



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
    resp, r, err := api_client.IntegrationsApi.GetIntegrations(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsRequest struct via the builder pattern


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


## ListIntegrationInventory

> map[string]interface{} ListIntegrationInventory(ctx, id).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Get All Integration Inventory



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
    id := int64(9) // int64 | Morpheus ID of the Integration
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.ListIntegrationInventory(context.Background(), id).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationInventory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationInventoryRequest struct via the builder pattern


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


## ListIntegrationObjects

> InlineResponse20063 ListIntegrationObjects(ctx, id).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Value(value).RefId(refId).Execute()

Get ServiceNow Integration Objects



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
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    value := "value_example" // string | The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing `type=string`. This means the data value returned by the API will be a string instead of an object.  (optional)
    refId := int64(3) // int64 | If specified will return an exact match on refId (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.ListIntegrationObjects(context.Background(), id).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Value(value).RefId(refId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrationObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationObjects`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrationObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **value** | **string** | The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the data value returned by the API will be a string instead of an object.  | 
 **refId** | **int64** | If specified will return an exact match on refId | 

### Return type

[**InlineResponse20063**](inline_response_200_63.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationTypes

> map[string]interface{} ListIntegrationTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Optiontypes(optiontypes).Description(description).Category(category).Creatable(creatable).Enabled(enabled).HasCMDB(hasCMDB).HasCMDBDiscovery(hasCMDBDiscovery).HasCM(hasCM).HasDNS(hasDNS).HasApprovals(hasApprovals).IsPlugin(isPlugin).Execute()

Retrieves all Integration Types



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
    optiontypes := true // bool | Pass `true` to include optionTypes in the response for each integration type (optional) (default to false)
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    category := "category_example" // string | If specified will return an exact match on category (optional)
    creatable := true // bool | Filter by creatable (optional)
    enabled := false // bool | Filter by enabled (optional)
    hasCMDB := true // bool | Filter by hasCMDB (optional)
    hasCMDBDiscovery := true // bool | Filter by hasCMDBDiscovery (optional)
    hasCM := true // bool | Filter by hasCM (optional)
    hasDNS := true // bool | Filter by hasDNS (optional)
    hasApprovals := true // bool | Filter by hasApprovals (optional)
    isPlugin := true // bool | Filter by isPlugin (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.ListIntegrationTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Optiontypes(optiontypes).Description(description).Category(category).Creatable(creatable).Enabled(enabled).HasCMDB(hasCMDB).HasCMDBDiscovery(hasCMDBDiscovery).HasCM(hasCM).HasDNS(hasDNS).HasApprovals(hasApprovals).IsPlugin(isPlugin).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **optiontypes** | **bool** | Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [default to false]
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 
 **category** | **string** | If specified will return an exact match on category | 
 **creatable** | **bool** | Filter by creatable | 
 **enabled** | **bool** | Filter by enabled | 
 **hasCMDB** | **bool** | Filter by hasCMDB | 
 **hasCMDBDiscovery** | **bool** | Filter by hasCMDBDiscovery | 
 **hasCM** | **bool** | Filter by hasCM | 
 **hasDNS** | **bool** | Filter by hasDNS | 
 **hasApprovals** | **bool** | Filter by hasApprovals | 
 **isPlugin** | **bool** | Filter by isPlugin | 

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


## ListIntegrations

> AnyOfobjectmeta ListIntegrations(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Id(id).Url(url).Host(host).Port(port).Username(username).Version(version).WindowsVersion(windowsVersion).Status(status).Objects(objects).Execute()

Retrieves all Integrations



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
    id := int64(7) // int64 | Morpheus ID of the Object being created or referenced (optional)
    url := "https://example.com/testimage.ovf" // string | Download the file from a remote url. This can be used instead of uploading a local file. (optional)
    host := "host_example" // string | Filter by integration Host (optional)
    port := "port_example" // string | Filter by integration Port (optional)
    username := "administrator" // string | Username (optional)
    version := int64(5) // int64 | Filter by version number (userVersion) (optional)
    windowsVersion := "windowsVersion_example" // string | Filter by integration Windows Version (optional)
    status := "running" // string | The instance status for filtering. (optional)
    objects := true // bool | Include `objects=true` to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.ListIntegrations(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Id(id).Url(url).Host(host).Port(port).Username(username).Version(version).WindowsVersion(windowsVersion).Status(status).Objects(objects).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrations`: AnyOfobjectmeta
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **id** | **int64** | Morpheus ID of the Object being created or referenced | 
 **url** | **string** | Download the file from a remote url. This can be used instead of uploading a local file. | 
 **host** | **string** | Filter by integration Host | 
 **port** | **string** | Filter by integration Port | 
 **username** | **string** | Username | 
 **version** | **int64** | Filter by version number (userVersion) | 
 **windowsVersion** | **string** | Filter by integration Windows Version | 
 **status** | **string** | The instance status for filtering. | 
 **objects** | **bool** | Include &#x60;objects&#x3D;true&#x60; to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  | [default to false]

### Return type

[**AnyOfobjectmeta**](anyOf&lt;object,meta&gt;.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshIntegrations

> map[string]interface{} RefreshIntegrations(ctx, id).Execute()

Refresh an Integration



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
    resp, r, err := api_client.IntegrationsApi.RefreshIntegrations(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.RefreshIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshIntegrations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.RefreshIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshIntegrationsRequest struct via the builder pattern


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


## RemoveIntegrationObjects

> Model200Success RemoveIntegrationObjects(ctx, id, objectId).Execute()

Deletes a ServiceNow Integration object



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
    objectId := int64(122) // int64 | Morpheus ID of the Object being created or referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.RemoveIntegrationObjects(context.Background(), id, objectId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.RemoveIntegrationObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveIntegrationObjects`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.RemoveIntegrationObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**objectId** | **int64** | Morpheus ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIntegrationObjectsRequest struct via the builder pattern


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


## RemoveIntegrations

> Model200Success RemoveIntegrations(ctx, id).Execute()

Deletes an Integration



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
    resp, r, err := api_client.IntegrationsApi.RemoveIntegrations(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.RemoveIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveIntegrations`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.RemoveIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIntegrationsRequest struct via the builder pattern


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


## UpdateIntegrationInventory

> map[string]interface{} UpdateIntegrationInventory(ctx, id, inventoryId).InlineObject101(inlineObject101).Execute()

Updating an Integration Inventory



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
    id := int64(9) // int64 | Morpheus ID of the Integration
    inventoryId := int64(1) // int64 | Morpheus ID of the Integration Inventory
    inlineObject101 := *openapiclient.NewInlineObject101(*openapiclient.NewApiIntegrationsIdInventoryInventoryIdInventory()) // InlineObject101 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.UpdateIntegrationInventory(context.Background(), id, inventoryId).InlineObject101(inlineObject101).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UpdateIntegrationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIntegrationInventory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.UpdateIntegrationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Integration | 
**inventoryId** | **int64** | Morpheus ID of the Integration Inventory | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject101** | [**InlineObject101**](InlineObject101.md) |  | 

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


## UpdateIntegrations

> map[string]interface{} UpdateIntegrations(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Updates an Integration



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.UpdateIntegrations(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UpdateIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIntegrations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.UpdateIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

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

