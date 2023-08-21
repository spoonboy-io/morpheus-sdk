# \UsageApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsages**](UsageApi.md#ListUsages) | **Get** /api/usage | Retrieves Usage Records



## ListUsages

> Usages ListUsages(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).StartDate(startDate).EndDate(endDate).Execute()

Retrieves Usage Records



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsageApi.ListUsages(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).StartDate(startDate).EndDate(endDate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.ListUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsages`: Usages
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.ListUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 

### Return type

[**Usages**](usages.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

