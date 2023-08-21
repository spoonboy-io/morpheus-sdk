# \SecurityGroupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecurityGroupLocations**](SecurityGroupsApi.md#AddSecurityGroupLocations) | **Post** /api/security-groups/{id}/locations | Creates a Security Group Location
[**AddSecurityGroupRules**](SecurityGroupsApi.md#AddSecurityGroupRules) | **Post** /api/security-groups/{id}/rules | Creates a Security Group Rule
[**AddSecurityGroups**](SecurityGroupsApi.md#AddSecurityGroups) | **Post** /api/security-groups | Creates a Security Group
[**GetSecurityGroupRules**](SecurityGroupsApi.md#GetSecurityGroupRules) | **Get** /api/security-groups/{id}/rules/{sgId} | Retrieves a Specific Security Group Rule
[**GetSecurityGroups**](SecurityGroupsApi.md#GetSecurityGroups) | **Get** /api/security-groups/{id} | Retrieves a Specific Security Group
[**ListSecurityGroupRules**](SecurityGroupsApi.md#ListSecurityGroupRules) | **Get** /api/security-groups/{id}/rules | Retrieves all Security Group Rules
[**ListSecurityGroups**](SecurityGroupsApi.md#ListSecurityGroups) | **Get** /api/security-groups | Retrieves all Security Groups
[**RemoveSecurityGroupLocations**](SecurityGroupsApi.md#RemoveSecurityGroupLocations) | **Delete** /api/security-groups/{id}/locations/{locationId} | Deletes a Security Group Location
[**RemoveSecurityGroupRules**](SecurityGroupsApi.md#RemoveSecurityGroupRules) | **Delete** /api/security-groups/{id}/rules/{sgId} | Deletes a Security Group Rule
[**RemoveSecurityGroups**](SecurityGroupsApi.md#RemoveSecurityGroups) | **Delete** /api/security-groups/{id} | Deletes a Security Group
[**UpdateSecurityGroupRules**](SecurityGroupsApi.md#UpdateSecurityGroupRules) | **Put** /api/security-groups/{id}/rules/{sgId} | Updates a Security Group Rule
[**UpdateSecurityGroups**](SecurityGroupsApi.md#UpdateSecurityGroups) | **Put** /api/security-groups/{id} | Updating a Security Group



## AddSecurityGroupLocations

> map[string]interface{} AddSecurityGroupLocations(ctx, id).InlineObject215(inlineObject215).Execute()

Creates a Security Group Location



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
    inlineObject215 := *openapiclient.NewInlineObject215(*openapiclient.NewApiSecurityGroupsIdLocationsSecurityGroupLocation(int64(1), "TODO")) // InlineObject215 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.AddSecurityGroupLocations(context.Background(), id).InlineObject215(inlineObject215).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.AddSecurityGroupLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSecurityGroupLocations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.AddSecurityGroupLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSecurityGroupLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject215** | [**InlineObject215**](InlineObject215.md) |  | 

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


## AddSecurityGroupRules

> map[string]interface{} AddSecurityGroupRules(ctx, id).InlineObject216(inlineObject216).Execute()

Creates a Security Group Rule



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
    inlineObject216 := *openapiclient.NewInlineObject216(*openapiclient.NewApiSecurityGroupsIdRulesRule("Protocol_example", "RuleType_example")) // InlineObject216 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.AddSecurityGroupRules(context.Background(), id).InlineObject216(inlineObject216).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.AddSecurityGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSecurityGroupRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.AddSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject216** | [**InlineObject216**](InlineObject216.md) |  | 

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


## AddSecurityGroups

> InlineResponse200133 AddSecurityGroups(ctx).InlineObject213(inlineObject213).Execute()

Creates a Security Group



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
    inlineObject213 := *openapiclient.NewInlineObject213(*openapiclient.NewApiSecurityGroupsSecurityGroup("Name_example", int64(3))) // InlineObject213 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.AddSecurityGroups(context.Background()).InlineObject213(inlineObject213).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.AddSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSecurityGroups`: InlineResponse200133
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.AddSecurityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject213** | [**InlineObject213**](InlineObject213.md) |  | 

### Return type

[**InlineResponse200133**](inline_response_200_133.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroupRules

> InlineResponse200135 GetSecurityGroupRules(ctx, id, sgId).Execute()

Retrieves a Specific Security Group Rule



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
    sgId := float32(2352) // float32 | Morpheus ID of the security group rule being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.GetSecurityGroupRules(context.Background(), id, sgId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.GetSecurityGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityGroupRules`: InlineResponse200135
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.GetSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**sgId** | **float32** | Morpheus ID of the security group rule being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200135**](inline_response_200_135.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityGroups

> InlineResponse200134 GetSecurityGroups(ctx, id).Execute()

Retrieves a Specific Security Group



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
    resp, r, err := api_client.SecurityGroupsApi.GetSecurityGroups(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.GetSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityGroups`: InlineResponse200134
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.GetSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200134**](inline_response_200_134.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityGroupRules

> map[string]interface{} ListSecurityGroupRules(ctx, id).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Security Group Rules



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.ListSecurityGroupRules(context.Background(), id).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.ListSecurityGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityGroupRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.ListSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupRulesRequest struct via the builder pattern


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


## ListSecurityGroups

> map[string]interface{} ListSecurityGroups(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Security Groups



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
    resp, r, err := api_client.SecurityGroupsApi.ListSecurityGroups(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.ListSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.ListSecurityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupsRequest struct via the builder pattern


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


## RemoveSecurityGroupLocations

> Model200Success RemoveSecurityGroupLocations(ctx, id, locationId).Execute()

Deletes a Security Group Location



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
    locationId := float32(2) // float32 | The ID of the location

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.RemoveSecurityGroupLocations(context.Background(), id, locationId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.RemoveSecurityGroupLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecurityGroupLocations`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.RemoveSecurityGroupLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**locationId** | **float32** | The ID of the location | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecurityGroupLocationsRequest struct via the builder pattern


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


## RemoveSecurityGroupRules

> Model200Success RemoveSecurityGroupRules(ctx, id, sgId).Execute()

Deletes a Security Group Rule



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
    sgId := float32(2352) // float32 | Morpheus ID of the security group rule being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.RemoveSecurityGroupRules(context.Background(), id, sgId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.RemoveSecurityGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecurityGroupRules`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.RemoveSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**sgId** | **float32** | Morpheus ID of the security group rule being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecurityGroupRulesRequest struct via the builder pattern


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


## RemoveSecurityGroups

> Model200Success RemoveSecurityGroups(ctx, id).Execute()

Deletes a Security Group



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
    resp, r, err := api_client.SecurityGroupsApi.RemoveSecurityGroups(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.RemoveSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecurityGroups`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.RemoveSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecurityGroupsRequest struct via the builder pattern


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


## UpdateSecurityGroupRules

> map[string]interface{} UpdateSecurityGroupRules(ctx, id, sgId).InlineObject217(inlineObject217).Execute()

Updates a Security Group Rule



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
    sgId := float32(2352) // float32 | Morpheus ID of the security group rule being referenced
    inlineObject217 := *openapiclient.NewInlineObject217(*openapiclient.NewApiSecurityGroupsIdRulesSgIdRule("Protocol_example", "RuleType_example")) // InlineObject217 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.UpdateSecurityGroupRules(context.Background(), id, sgId).InlineObject217(inlineObject217).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.UpdateSecurityGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityGroupRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.UpdateSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**sgId** | **float32** | Morpheus ID of the security group rule being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject217** | [**InlineObject217**](InlineObject217.md) |  | 

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


## UpdateSecurityGroups

> InlineResponse200133 UpdateSecurityGroups(ctx, id).InlineObject214(inlineObject214).Execute()

Updating a Security Group



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
    inlineObject214 := *openapiclient.NewInlineObject214(*openapiclient.NewApiSecurityGroupsIdSecurityGroup()) // InlineObject214 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityGroupsApi.UpdateSecurityGroups(context.Background(), id).InlineObject214(inlineObject214).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.UpdateSecurityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityGroups`: InlineResponse200133
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.UpdateSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject214** | [**InlineObject214**](InlineObject214.md) |  | 

### Return type

[**InlineResponse200133**](inline_response_200_133.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

