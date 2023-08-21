# \InvoicesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoiceLineItems**](InvoicesApi.md#GetInvoiceLineItems) | **Get** /api/invoice-line-items/{id} | Get a Specific Invoice Line Item
[**GetInvoices**](InvoicesApi.md#GetInvoices) | **Get** /api/invoices/{id} | Get a Specific Invoice
[**ListInvoiceLineItems**](InvoicesApi.md#ListInvoiceLineItems) | **Get** /api/invoice-line-items | List All Invoice Line Items
[**ListInvoices**](InvoicesApi.md#ListInvoices) | **Get** /api/invoices | List All Invoices
[**UpdateInvoices**](InvoicesApi.md#UpdateInvoices) | **Put** /api/invoices/{id} | Update Invoice Tags



## GetInvoiceLineItems

> map[string]interface{} GetInvoiceLineItems(ctx, id).Execute()

Get a Specific Invoice Line Item



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
    resp, r, err := api_client.InvoicesApi.GetInvoiceLineItems(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoiceLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceLineItems`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoiceLineItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceLineItemsRequest struct via the builder pattern


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


## GetInvoices

> map[string]interface{} GetInvoices(ctx, id).Execute()

Get a Specific Invoice



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
    resp, r, err := api_client.InvoicesApi.GetInvoices(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


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


## ListInvoiceLineItems

> map[string]interface{} ListInvoiceLineItems(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).StartDate(startDate).EndDate(endDate).Period(period).RefType(refType).RefId(refId).ZoneId(zoneId).SiteId(siteId).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).UserId(userId).ProjectId(projectId).Active(active).AccountId(accountId).IncludeTotals(includeTotals).Execute()

List All Invoice Line Items



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    period := "201901" // string | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  (optional)
    refType := "ComputeSite" // string | If specified will return an exact match on refType.  (optional)
    refId := int64(3) // int64 | If specified will return an exact match on refId (optional)
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    siteId := int64(7) // int64 | The Site ID for Filtering (optional)
    instanceId := int64(94) // int64 | The Instance ID for Filtering (optional)
    containerId := int64(135) // int64 | The Container ID for Filtering (optional)
    serverId := int64(97) // int64 | The Server ID for Filtering (optional)
    userId := int64(6) // int64 | Filter by User ID (optional)
    projectId := int64(1) // int64 | The Project ID for Filtering (optional)
    active := false // bool | True or False flag for Active (optional)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    includeTotals := true // bool | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.ListInvoiceLineItems(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).StartDate(startDate).EndDate(endDate).Period(period).RefType(refType).RefId(refId).ZoneId(zoneId).SiteId(siteId).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).UserId(userId).ProjectId(projectId).Active(active).AccountId(accountId).IncludeTotals(includeTotals).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ListInvoiceLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoiceLineItems`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ListInvoiceLineItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoiceLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **period** | **string** | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | 
 **refType** | **string** | If specified will return an exact match on refType.  | 
 **refId** | **int64** | If specified will return an exact match on refId | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **siteId** | **int64** | The Site ID for Filtering | 
 **instanceId** | **int64** | The Instance ID for Filtering | 
 **containerId** | **int64** | The Container ID for Filtering | 
 **serverId** | **int64** | The Server ID for Filtering | 
 **userId** | **int64** | Filter by User ID | 
 **projectId** | **int64** | The Project ID for Filtering | 
 **active** | **bool** | True or False flag for Active | 
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 
 **includeTotals** | **bool** | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [default to false]

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


## ListInvoices

> map[string]interface{} ListInvoices(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).StartDate(startDate).EndDate(endDate).Period(period).RefType(refType).RefId(refId).RefStatus(refStatus).ZoneId(zoneId).SiteId(siteId).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).UserId(userId).ProjectId(projectId).Active(active).AccountId(accountId).IncludeLineItems(includeLineItems).IncludeTotals(includeTotals).Tags(tags).Execute()

List All Invoices



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
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "refName")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    period := "201901" // string | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  (optional)
    refType := "ComputeSite" // string | If specified will return an exact match on refType.  (optional)
    refId := int64(3) // int64 | If specified will return an exact match on refId (optional)
    refStatus := "provisioned" // string | If specified, will filter on the associated StorageVolume status. This is only applicable whn `refType=StorageVolume`.  (optional)
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    siteId := int64(7) // int64 | The Site ID for Filtering (optional)
    instanceId := int64(94) // int64 | The Instance ID for Filtering (optional)
    containerId := int64(135) // int64 | The Container ID for Filtering (optional)
    serverId := int64(97) // int64 | The Server ID for Filtering (optional)
    userId := int64(6) // int64 | Filter by User ID (optional)
    projectId := int64(1) // int64 | The Project ID for Filtering (optional)
    active := false // bool | True or False flag for Active (optional)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    includeLineItems := true // bool | Pass true to include the list of `lineItems` for each invoice. Only `lineItemCount` is returned by default.  (optional) (default to false)
    includeTotals := true // bool | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`.  (optional) (default to false)
    tags := "tags.env=qa&tags.env=test" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.ListInvoices(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).StartDate(startDate).EndDate(endDate).Period(period).RefType(refType).RefId(refId).RefStatus(refStatus).ZoneId(zoneId).SiteId(siteId).InstanceId(instanceId).ContainerId(containerId).ServerId(serverId).UserId(userId).ProjectId(projectId).Active(active).AccountId(accountId).IncludeLineItems(includeLineItems).IncludeTotals(includeTotals).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;refName&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **period** | **string** | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | 
 **refType** | **string** | If specified will return an exact match on refType.  | 
 **refId** | **int64** | If specified will return an exact match on refId | 
 **refStatus** | **string** | If specified, will filter on the associated StorageVolume status. This is only applicable whn &#x60;refType&#x3D;StorageVolume&#x60;.  | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **siteId** | **int64** | The Site ID for Filtering | 
 **instanceId** | **int64** | The Instance ID for Filtering | 
 **containerId** | **int64** | The Container ID for Filtering | 
 **serverId** | **int64** | The Server ID for Filtering | 
 **userId** | **int64** | Filter by User ID | 
 **projectId** | **int64** | The Project ID for Filtering | 
 **active** | **bool** | True or False flag for Active | 
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 
 **includeLineItems** | **bool** | Pass true to include the list of &#x60;lineItems&#x60; for each invoice. Only &#x60;lineItemCount&#x60; is returned by default.  | [default to false]
 **includeTotals** | **bool** | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [default to false]
 **tags** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 

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


## UpdateInvoices

> map[string]interface{} UpdateInvoices(ctx, id).InlineObject102(inlineObject102).Execute()

Update Invoice Tags



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
    inlineObject102 := *openapiclient.NewInlineObject102() // InlineObject102 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.UpdateInvoices(context.Background(), id).InlineObject102(inlineObject102).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.UpdateInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInvoices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.UpdateInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject102** | [**InlineObject102**](InlineObject102.md) |  | 

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

