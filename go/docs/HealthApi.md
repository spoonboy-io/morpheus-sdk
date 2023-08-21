# \HealthApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcknowledgeHealthAlarm**](HealthApi.md#AcknowledgeHealthAlarm) | **Put** /api/health/alarms/{id}/acknowledge | Acknowledge a Health Alarm
[**AcknowledgeHealthAlarms**](HealthApi.md#AcknowledgeHealthAlarms) | **Put** /api/health/alarms/acknowledge | Acknowledge Many Health Alarms
[**ExportHealthLogs**](HealthApi.md#ExportHealthLogs) | **Get** /api/health/logs/export | Export Appliance Health Logs
[**GetHealthAlarms**](HealthApi.md#GetHealthAlarms) | **Get** /api/health/alarms/{id} | Retrieves a Specific Appliance Health Alarm
[**ListHealth**](HealthApi.md#ListHealth) | **Get** /api/health | Retrieves Appliance Health
[**ListHealthAlarms**](HealthApi.md#ListHealthAlarms) | **Get** /api/health/alarms | Retrieves Appliance Health Alarms
[**ListHealthLogs**](HealthApi.md#ListHealthLogs) | **Get** /api/health/logs | Retrieves Appliance Health Logs
[**Ping**](HealthApi.md#Ping) | **Get** /api/ping | Basic information about current Morpheus Installation



## AcknowledgeHealthAlarm

> Model200Success AcknowledgeHealthAlarm(ctx, id).InlineObject81(inlineObject81).Execute()

Acknowledge a Health Alarm



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
    inlineObject81 := *openapiclient.NewInlineObject81(*openapiclient.NewApiHealthAlarmsIdAcknowledgeAlarm(false)) // InlineObject81 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.AcknowledgeHealthAlarm(context.Background(), id).InlineObject81(inlineObject81).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.AcknowledgeHealthAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcknowledgeHealthAlarm`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.AcknowledgeHealthAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeHealthAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject81** | [**InlineObject81**](InlineObject81.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcknowledgeHealthAlarms

> Model200Success AcknowledgeHealthAlarms(ctx).InlineObject80(inlineObject80).Execute()

Acknowledge Many Health Alarms



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
    inlineObject80 := *openapiclient.NewInlineObject80(*openapiclient.NewApiHealthAlarmsAcknowledgeAlarm(false)) // InlineObject80 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.AcknowledgeHealthAlarms(context.Background()).InlineObject80(inlineObject80).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.AcknowledgeHealthAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcknowledgeHealthAlarms`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.AcknowledgeHealthAlarms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeHealthAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject80** | [**InlineObject80**](InlineObject80.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportHealthLogs

> *os.File ExportHealthLogs(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Acknowledged(acknowledged).StartDate(startDate).EndDate(endDate).Reverse(reverse).Execute()

Export Appliance Health Logs



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
    acknowledged := false // bool | True or False flag for Acknowledged items (optional)
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    reverse := true // bool | Reverse order of records. This `true` by default when sort and direction are not passed, but `false` by default if either is passed. This means that by default the newest log entries are the bottom of the file.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.ExportHealthLogs(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Acknowledged(acknowledged).StartDate(startDate).EndDate(endDate).Reverse(reverse).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.ExportHealthLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportHealthLogs`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.ExportHealthLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportHealthLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **acknowledged** | **bool** | True or False flag for Acknowledged items | 
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **reverse** | **bool** | Reverse order of records. This &#x60;true&#x60; by default when sort and direction are not passed, but &#x60;false&#x60; by default if either is passed. This means that by default the newest log entries are the bottom of the file.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthAlarms

> map[string]interface{} GetHealthAlarms(ctx, id).Execute()

Retrieves a Specific Appliance Health Alarm



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
    resp, r, err := api_client.HealthApi.GetHealthAlarms(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.GetHealthAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthAlarms`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.GetHealthAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthAlarmsRequest struct via the builder pattern


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


## ListHealth

> map[string]interface{} ListHealth(ctx).Execute()

Retrieves Appliance Health



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
    resp, r, err := api_client.HealthApi.ListHealth(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.ListHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHealth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.ListHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHealthRequest struct via the builder pattern


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


## ListHealthAlarms

> map[string]interface{} ListHealthAlarms(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Acknowledged(acknowledged).Execute()

Retrieves Appliance Health Alarms



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
    acknowledged := false // bool | True or False flag for Acknowledged items (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.ListHealthAlarms(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Acknowledged(acknowledged).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.ListHealthAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHealthAlarms`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.ListHealthAlarms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHealthAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **acknowledged** | **bool** | True or False flag for Acknowledged items | 

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


## ListHealthLogs

> map[string]interface{} ListHealthLogs(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Acknowledged(acknowledged).StartDate(startDate).EndDate(endDate).Execute()

Retrieves Appliance Health Logs



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
    acknowledged := false // bool | True or False flag for Acknowledged items (optional)
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.ListHealthLogs(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Acknowledged(acknowledged).StartDate(startDate).EndDate(endDate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.ListHealthLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHealthLogs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.ListHealthLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHealthLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **acknowledged** | **bool** | True or False flag for Acknowledged items | 
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 

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


## Ping

> Ping Ping(ctx).Execute()

Basic information about current Morpheus Installation



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
    resp, r, err := api_client.HealthApi.Ping(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ping`: Ping
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

[**Ping**](ping.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

