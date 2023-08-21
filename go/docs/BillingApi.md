# \BillingApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingAccount**](BillingApi.md#GetBillingAccount) | **Get** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it
[**GetBillingInstancesIdentifier**](BillingApi.md#GetBillingInstancesIdentifier) | **Get** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.
[**GetBillingServersIdentifier**](BillingApi.md#GetBillingServersIdentifier) | **Get** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.
[**GetBillingZoneIdentifier**](BillingApi.md#GetBillingZoneIdentifier) | **Get** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.
[**ListBillingAccount**](BillingApi.md#ListBillingAccount) | **Get** /api/billing/account | Retrieves billing information for the requesting user&#39;s account.
[**ListBillingInstances**](BillingApi.md#ListBillingInstances) | **Get** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account.
[**ListBillingServers**](BillingApi.md#ListBillingServers) | **Get** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.
[**ListBillingZone**](BillingApi.md#ListBillingZone) | **Get** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account.



## GetBillingAccount

> map[string]interface{} GetBillingAccount(ctx, id).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()

This endpoint will retrieve a specific account by id if the user has permission to access it



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeComputeServers := true // bool | Optional ability to exclude compute servers (optional) (default to true)
    includeInstances := true // bool | Optional ability to exclude instances (optional) (default to true)
    includeDiscoveredServers := true // bool | Optional ability to exclude discovered servers (optional) (default to true)
    includeLoadBalancers := true // bool | Optional ability to exclude load balancers (optional) (default to true)
    includeVirtualImages := true // bool | Optional ability to exclude virtual images (optional) (default to true)
    includeSnapshots := true // bool | Optional ability to exclude snapshots (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetBillingAccount(context.Background(), id).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBillingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBillingAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeComputeServers** | **bool** | Optional ability to exclude compute servers | [default to true]
 **includeInstances** | **bool** | Optional ability to exclude instances | [default to true]
 **includeDiscoveredServers** | **bool** | Optional ability to exclude discovered servers | [default to true]
 **includeLoadBalancers** | **bool** | Optional ability to exclude load balancers | [default to true]
 **includeVirtualImages** | **bool** | Optional ability to exclude virtual images | [default to true]
 **includeSnapshots** | **bool** | Optional ability to exclude snapshots | [default to true]

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


## GetBillingInstancesIdentifier

> map[string]interface{} GetBillingInstancesIdentifier(ctx, identifier).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()

Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.



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
    identifier := "identifier_example" // string | Morpheus UUID or ID of the Object being created or referenced
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeTenants := true // bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to false)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetBillingInstancesIdentifier(context.Background(), identifier).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBillingInstancesIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingInstancesIdentifier`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBillingInstancesIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Morpheus UUID or ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingInstancesIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeTenants** | **bool** | Optional ability to include all subtenant billing information when calling from a master tenant user | [default to false]
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 

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


## GetBillingServersIdentifier

> map[string]interface{} GetBillingServersIdentifier(ctx, identifier).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()

Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.



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
    identifier := "identifier_example" // string | Morpheus UUID or ID of the Object being created or referenced
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeTenants := true // bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to false)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetBillingServersIdentifier(context.Background(), identifier).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBillingServersIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingServersIdentifier`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBillingServersIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Morpheus UUID or ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingServersIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeTenants** | **bool** | Optional ability to include all subtenant billing information when calling from a master tenant user | [default to false]
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 

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


## GetBillingZoneIdentifier

> map[string]interface{} GetBillingZoneIdentifier(ctx, identifier).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()

Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.



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
    identifier := "identifier_example" // string | Morpheus UUID or ID of the Object being created or referenced
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeComputeServers := true // bool | Optional ability to exclude compute servers (optional) (default to true)
    includeInstances := true // bool | Optional ability to exclude instances (optional) (default to true)
    includeDiscoveredServers := true // bool | Optional ability to exclude discovered servers (optional) (default to true)
    includeLoadBalancers := true // bool | Optional ability to exclude load balancers (optional) (default to true)
    includeVirtualImages := true // bool | Optional ability to exclude virtual images (optional) (default to true)
    includeSnapshots := true // bool | Optional ability to exclude snapshots (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.GetBillingZoneIdentifier(context.Background(), identifier).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBillingZoneIdentifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingZoneIdentifier`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBillingZoneIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Morpheus UUID or ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingZoneIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeComputeServers** | **bool** | Optional ability to exclude compute servers | [default to true]
 **includeInstances** | **bool** | Optional ability to exclude instances | [default to true]
 **includeDiscoveredServers** | **bool** | Optional ability to exclude discovered servers | [default to true]
 **includeLoadBalancers** | **bool** | Optional ability to exclude load balancers | [default to true]
 **includeVirtualImages** | **bool** | Optional ability to exclude virtual images | [default to true]
 **includeSnapshots** | **bool** | Optional ability to exclude snapshots | [default to true]

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


## ListBillingAccount

> map[string]interface{} ListBillingAccount(ctx).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()

Retrieves billing information for the requesting user's account.



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeComputeServers := true // bool | Optional ability to exclude compute servers (optional) (default to true)
    includeInstances := true // bool | Optional ability to exclude instances (optional) (default to true)
    includeDiscoveredServers := true // bool | Optional ability to exclude discovered servers (optional) (default to true)
    includeLoadBalancers := true // bool | Optional ability to exclude load balancers (optional) (default to true)
    includeVirtualImages := true // bool | Optional ability to exclude virtual images (optional) (default to true)
    includeSnapshots := true // bool | Optional ability to exclude snapshots (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.ListBillingAccount(context.Background()).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListBillingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListBillingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeComputeServers** | **bool** | Optional ability to exclude compute servers | [default to true]
 **includeInstances** | **bool** | Optional ability to exclude instances | [default to true]
 **includeDiscoveredServers** | **bool** | Optional ability to exclude discovered servers | [default to true]
 **includeLoadBalancers** | **bool** | Optional ability to exclude load balancers | [default to true]
 **includeVirtualImages** | **bool** | Optional ability to exclude virtual images | [default to true]
 **includeSnapshots** | **bool** | Optional ability to exclude snapshots | [default to true]

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


## ListBillingInstances

> map[string]interface{} ListBillingInstances(ctx).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()

Retrieves billing information for all instances on the requestor's account.



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeTenants := true // bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to false)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.ListBillingInstances(context.Background()).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListBillingInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingInstances`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListBillingInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeTenants** | **bool** | Optional ability to include all subtenant billing information when calling from a master tenant user | [default to false]
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 

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


## ListBillingServers

> map[string]interface{} ListBillingServers(ctx).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()

Retrieves billing information for all servers (container hosts) on the requestor's account.



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeTenants := true // bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to false)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.ListBillingServers(context.Background()).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeTenants(includeTenants).AccountId(accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListBillingServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingServers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListBillingServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeTenants** | **bool** | Optional ability to include all subtenant billing information when calling from a master tenant user | [default to false]
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 

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


## ListBillingZone

> map[string]interface{} ListBillingZone(ctx).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()

Retrieves billing information for all zones on the requestor's account.



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
    startDate := "2019-01-01" // string | Filter by startDate greater than or equal to a specified date (optional)
    endDate := "2020-01-01" // string | Filter by endDate less than or equal to a specified date (optional)
    includeUsages := true // bool | Optional ability to suppress the usage records (optional) (default to true)
    maxUsages := int64(789) // int64 | Optional ability to limit the usages returned (optional)
    offsetUsages := int64(789) // int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    includeComputeServers := true // bool | Optional ability to exclude compute servers (optional) (default to true)
    includeInstances := true // bool | Optional ability to exclude instances (optional) (default to true)
    includeDiscoveredServers := true // bool | Optional ability to exclude discovered servers (optional) (default to true)
    includeLoadBalancers := true // bool | Optional ability to exclude load balancers (optional) (default to true)
    includeVirtualImages := true // bool | Optional ability to exclude virtual images (optional) (default to true)
    includeSnapshots := true // bool | Optional ability to exclude snapshots (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingApi.ListBillingZone(context.Background()).StartDate(startDate).EndDate(endDate).IncludeUsages(includeUsages).MaxUsages(maxUsages).OffsetUsages(offsetUsages).IncludeComputeServers(includeComputeServers).IncludeInstances(includeInstances).IncludeDiscoveredServers(includeDiscoveredServers).IncludeLoadBalancers(includeLoadBalancers).IncludeVirtualImages(includeVirtualImages).IncludeSnapshots(includeSnapshots).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListBillingZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingZone`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListBillingZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Filter by startDate greater than or equal to a specified date | 
 **endDate** | **string** | Filter by endDate less than or equal to a specified date | 
 **includeUsages** | **bool** | Optional ability to suppress the usage records | [default to true]
 **maxUsages** | **int64** | Optional ability to limit the usages returned | 
 **offsetUsages** | **int64** | Optional ability to offset the usages returned, for use with maxUsages to paginate | 
 **includeComputeServers** | **bool** | Optional ability to exclude compute servers | [default to true]
 **includeInstances** | **bool** | Optional ability to exclude instances | [default to true]
 **includeDiscoveredServers** | **bool** | Optional ability to exclude discovered servers | [default to true]
 **includeLoadBalancers** | **bool** | Optional ability to exclude load balancers | [default to true]
 **includeVirtualImages** | **bool** | Optional ability to exclude virtual images | [default to true]
 **includeSnapshots** | **bool** | Optional ability to exclude snapshots | [default to true]

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

