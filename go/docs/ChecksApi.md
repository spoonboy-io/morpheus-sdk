# \ChecksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCheckApps**](ChecksApi.md#AddCheckApps) | **Post** /api/monitoring/apps | Create a New Check App
[**AddCheckGroups**](ChecksApi.md#AddCheckGroups) | **Post** /api/monitoring/groups | Create a New Check Group
[**AddChecks**](ChecksApi.md#AddChecks) | **Post** /api/monitoring/checks | Create a New Check
[**DeleteCheckApps**](ChecksApi.md#DeleteCheckApps) | **Delete** /api/monitoring/apps/{id} | Delete a Specific Check App
[**DeleteCheckGroups**](ChecksApi.md#DeleteCheckGroups) | **Delete** /api/monitoring/groups/{id} | Delete a Specific Check Group
[**DeleteChecks**](ChecksApi.md#DeleteChecks) | **Delete** /api/monitoring/checks/{id} | Delete a Specific Check
[**GetCheckApps**](ChecksApi.md#GetCheckApps) | **Get** /api/monitoring/apps/{id} | Get a Specific Check App
[**GetCheckGroups**](ChecksApi.md#GetCheckGroups) | **Get** /api/monitoring/groups/{id} | Get a Specific Check Group
[**GetCheckTypes**](ChecksApi.md#GetCheckTypes) | **Get** /api/monitoring/check-types/{id} | Get a Specific Check Type
[**GetChecks**](ChecksApi.md#GetChecks) | **Get** /api/monitoring/checks/{id} | Get a Specific Check
[**ListCheckApps**](ChecksApi.md#ListCheckApps) | **Get** /api/monitoring/apps | List All Check Apps
[**ListCheckGroups**](ChecksApi.md#ListCheckGroups) | **Get** /api/monitoring/groups | List All Check Groups
[**ListCheckTypes**](ChecksApi.md#ListCheckTypes) | **Get** /api/monitoring/check-types | List All Check Types
[**ListChecks**](ChecksApi.md#ListChecks) | **Get** /api/monitoring/checks | List All Checks
[**UpdateCheckApps**](ChecksApi.md#UpdateCheckApps) | **Put** /api/monitoring/apps/{id} | Update Check App
[**UpdateCheckGroups**](ChecksApi.md#UpdateCheckGroups) | **Put** /api/monitoring/groups/{id} | Update Check Group
[**UpdateChecks**](ChecksApi.md#UpdateChecks) | **Put** /api/monitoring/checks/{id} | Updates a Check
[**UpdateMuteAllCheckApps**](ChecksApi.md#UpdateMuteAllCheckApps) | **Put** /api/monitoring/apps/mute-all | Mute All Check Apps
[**UpdateMuteAllCheckGroups**](ChecksApi.md#UpdateMuteAllCheckGroups) | **Put** /api/monitoring/groups/mute-all | Mute All Check Groups
[**UpdateMuteAllChecks**](ChecksApi.md#UpdateMuteAllChecks) | **Put** /api/monitoring/checks/mute-all | Mute All Checks
[**UpdateMuteCheckApps**](ChecksApi.md#UpdateMuteCheckApps) | **Put** /api/monitoring/apps/{id}/mute | Mute Check App
[**UpdateMuteCheckGroups**](ChecksApi.md#UpdateMuteCheckGroups) | **Put** /api/monitoring/groups/{id}/mute | Mute Check Group
[**UpdateMuteChecks**](ChecksApi.md#UpdateMuteChecks) | **Put** /api/monitoring/checks/{id}/mute | Mute Check



## AddCheckApps

> map[string]interface{} AddCheckApps(ctx).InlineObject29(inlineObject29).Execute()

Create a New Check App



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
    inlineObject29 := *openapiclient.NewInlineObject29(*openapiclient.NewApiMonitoringAppsMonitorApp("My Check App")) // InlineObject29 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.AddCheckApps(context.Background()).InlineObject29(inlineObject29).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.AddCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCheckApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.AddCheckApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject29** | [**InlineObject29**](InlineObject29.md) |  | 

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


## AddCheckGroups

> map[string]interface{} AddCheckGroups(ctx).InlineObject37(inlineObject37).Execute()

Create a New Check Group



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
    inlineObject37 := *openapiclient.NewInlineObject37(*openapiclient.NewApiMonitoringGroupsCheckGroup("My Check Group")) // InlineObject37 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.AddCheckGroups(context.Background()).InlineObject37(inlineObject37).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.AddCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCheckGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.AddCheckGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCheckGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject37** | [**InlineObject37**](InlineObject37.md) |  | 

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


## AddChecks

> map[string]interface{} AddChecks(ctx).InlineObject33(inlineObject33).Execute()

Create a New Check



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
    inlineObject33 := *openapiclient.NewInlineObject33("TODO") // InlineObject33 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.AddChecks(context.Background()).InlineObject33(inlineObject33).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.AddChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddChecks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.AddChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject33** | [**InlineObject33**](InlineObject33.md) |  | 

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


## DeleteCheckApps

> Model200Success DeleteCheckApps(ctx, id).Execute()

Delete a Specific Check App



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
    resp, r, err := api_client.ChecksApi.DeleteCheckApps(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.DeleteCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCheckApps`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.DeleteCheckApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckAppsRequest struct via the builder pattern


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


## DeleteCheckGroups

> Model200Success DeleteCheckGroups(ctx, id).Execute()

Delete a Specific Check Group



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
    resp, r, err := api_client.ChecksApi.DeleteCheckGroups(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.DeleteCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCheckGroups`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.DeleteCheckGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckGroupsRequest struct via the builder pattern


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


## DeleteChecks

> Model200Success DeleteChecks(ctx, id).Execute()

Delete a Specific Check



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
    resp, r, err := api_client.ChecksApi.DeleteChecks(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.DeleteChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteChecks`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.DeleteChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChecksRequest struct via the builder pattern


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


## GetCheckApps

> InlineResponse20016 GetCheckApps(ctx, id).Execute()

Get a Specific Check App



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
    resp, r, err := api_client.ChecksApi.GetCheckApps(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.GetCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckApps`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.GetCheckApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckGroups

> InlineResponse20019 GetCheckGroups(ctx, id).Execute()

Get a Specific Check Group



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
    resp, r, err := api_client.ChecksApi.GetCheckGroups(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.GetCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckGroups`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.GetCheckGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckTypes

> InlineResponse20018 GetCheckTypes(ctx, id).Execute()

Get a Specific Check Type



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
    resp, r, err := api_client.ChecksApi.GetCheckTypes(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.GetCheckTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckTypes`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.GetCheckTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChecks

> InlineResponse20017 GetChecks(ctx, id).Execute()

Get a Specific Check



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
    resp, r, err := api_client.ChecksApi.GetChecks(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.GetChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChecks`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.GetChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCheckApps

> map[string]interface{} ListCheckApps(ctx).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()

List All Check Apps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    status := "running" // string | The instance status for filtering. (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted := true // bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.ListCheckApps(context.Background()).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ListCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ListCheckApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **status** | **string** | The instance status for filtering. | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **deleted** | **bool** | If true, only deleted resources or instances in pending removal status are returned. | 

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


## ListCheckGroups

> map[string]interface{} ListCheckGroups(ctx).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()

List All Check Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    status := "running" // string | The instance status for filtering. (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted := true // bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.ListCheckGroups(context.Background()).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ListCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ListCheckGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **status** | **string** | The instance status for filtering. | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **deleted** | **bool** | If true, only deleted resources or instances in pending removal status are returned. | 

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


## ListCheckTypes

> map[string]interface{} ListCheckTypes(ctx).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Execute()

List All Check Types



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
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.ListCheckTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ListCheckTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ListCheckTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

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


## ListChecks

> map[string]interface{} ListChecks(ctx).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()

List All Checks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    status := "healthy" // string | Filter by health status. (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted := true // bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.ListChecks(context.Background()).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ListChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChecks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ListChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **status** | **string** | Filter by health status. | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **deleted** | **bool** | If true, only deleted resources or instances in pending removal status are returned. | 

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


## UpdateCheckApps

> map[string]interface{} UpdateCheckApps(ctx, id).InlineObject31(inlineObject31).Execute()

Update Check App



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
    inlineObject31 := *openapiclient.NewInlineObject31(*openapiclient.NewApiMonitoringAppsIdMonitorApp()) // InlineObject31 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateCheckApps(context.Background(), id).InlineObject31(inlineObject31).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCheckApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateCheckApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject31** | [**InlineObject31**](InlineObject31.md) |  | 

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


## UpdateCheckGroups

> map[string]interface{} UpdateCheckGroups(ctx, id).InlineObject38(inlineObject38).Execute()

Update Check Group



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
    inlineObject38 := *openapiclient.NewInlineObject38(*openapiclient.NewApiMonitoringGroupsIdCheckGroup()) // InlineObject38 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateCheckGroups(context.Background(), id).InlineObject38(inlineObject38).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCheckGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateCheckGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject38** | [**InlineObject38**](InlineObject38.md) |  | 

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


## UpdateChecks

> map[string]interface{} UpdateChecks(ctx, id).InlineObject35(inlineObject35).Execute()

Updates a Check



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
    inlineObject35 := *openapiclient.NewInlineObject35("TODO") // InlineObject35 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateChecks(context.Background(), id).InlineObject35(inlineObject35).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChecks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject35** | [**InlineObject35**](InlineObject35.md) |  | 

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


## UpdateMuteAllCheckApps

> map[string]interface{} UpdateMuteAllCheckApps(ctx).InlineObject30(inlineObject30).Execute()

Mute All Check Apps



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
    inlineObject30 := *openapiclient.NewInlineObject30(false) // InlineObject30 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateMuteAllCheckApps(context.Background()).InlineObject30(inlineObject30).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateMuteAllCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteAllCheckApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateMuteAllCheckApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteAllCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject30** | [**InlineObject30**](InlineObject30.md) |  | 

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


## UpdateMuteAllCheckGroups

> map[string]interface{} UpdateMuteAllCheckGroups(ctx).InlineObject40(inlineObject40).Execute()

Mute All Check Groups



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
    inlineObject40 := *openapiclient.NewInlineObject40(false) // InlineObject40 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateMuteAllCheckGroups(context.Background()).InlineObject40(inlineObject40).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateMuteAllCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteAllCheckGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateMuteAllCheckGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteAllCheckGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject40** | [**InlineObject40**](InlineObject40.md) |  | 

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


## UpdateMuteAllChecks

> map[string]interface{} UpdateMuteAllChecks(ctx).InlineObject34(inlineObject34).Execute()

Mute All Checks



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
    inlineObject34 := *openapiclient.NewInlineObject34(false) // InlineObject34 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateMuteAllChecks(context.Background()).InlineObject34(inlineObject34).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateMuteAllChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteAllChecks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateMuteAllChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteAllChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject34** | [**InlineObject34**](InlineObject34.md) |  | 

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


## UpdateMuteCheckApps

> map[string]interface{} UpdateMuteCheckApps(ctx, id).InlineObject32(inlineObject32).Execute()

Mute Check App



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
    inlineObject32 := *openapiclient.NewInlineObject32(false) // InlineObject32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateMuteCheckApps(context.Background(), id).InlineObject32(inlineObject32).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateMuteCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteCheckApps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateMuteCheckApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject32** | [**InlineObject32**](InlineObject32.md) |  | 

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


## UpdateMuteCheckGroups

> map[string]interface{} UpdateMuteCheckGroups(ctx, id).InlineObject39(inlineObject39).Execute()

Mute Check Group



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
    inlineObject39 := *openapiclient.NewInlineObject39(false) // InlineObject39 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateMuteCheckGroups(context.Background(), id).InlineObject39(inlineObject39).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateMuteCheckGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteCheckGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateMuteCheckGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteCheckGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject39** | [**InlineObject39**](InlineObject39.md) |  | 

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


## UpdateMuteChecks

> map[string]interface{} UpdateMuteChecks(ctx, id).InlineObject36(inlineObject36).Execute()

Mute Check



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
    inlineObject36 := *openapiclient.NewInlineObject36(false) // InlineObject36 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChecksApi.UpdateMuteChecks(context.Background(), id).InlineObject36(inlineObject36).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.UpdateMuteChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMuteChecks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.UpdateMuteChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMuteChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject36** | [**InlineObject36**](InlineObject36.md) |  | 

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

