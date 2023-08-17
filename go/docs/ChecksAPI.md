# \ChecksAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCheckApps**](ChecksAPI.md#AddCheckApps) | **Post** /api/monitoring/apps | Create a New Check App
[**ListCheckApps**](ChecksAPI.md#ListCheckApps) | **Get** /api/monitoring/apps | List All Check Apps



## AddCheckApps

> AddCheckApps200Response AddCheckApps(ctx).AddCheckAppsRequest(addCheckAppsRequest).Execute()

Create a New Check App



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    addCheckAppsRequest := *openapiclient.NewAddCheckAppsRequest(*openapiclient.NewAddCheckAppsRequestMonitorApp("My Check App")) // AddCheckAppsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksAPI.AddCheckApps(context.Background()).AddCheckAppsRequest(addCheckAppsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.AddCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCheckApps`: AddCheckApps200Response
    fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.AddCheckApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCheckAppsRequest** | [**AddCheckAppsRequest**](AddCheckAppsRequest.md) |  | 

### Return type

[**AddCheckApps200Response**](AddCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCheckApps

> ListCheckApps200Response ListCheckApps(ctx).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()

List All Check Apps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    status := "running" // string | The instance status for filtering. (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted := true // bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksAPI.ListCheckApps(context.Background()).Max(max).Offset(offset).Sort(sort).Name(name).Phrase(phrase).Status(status).LastUpdated(lastUpdated).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ListCheckApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckApps`: ListCheckApps200Response
    fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ListCheckApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **status** | **string** | The instance status for filtering. | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **deleted** | **bool** | If true, only deleted resources or instances in pending removal status are returned. | 

### Return type

[**ListCheckApps200Response**](ListCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

