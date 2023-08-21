# \TenantsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTenant**](TenantsApi.md#AddTenant) | **Post** /api/accounts | Create a Tenant
[**AddUserTenant**](TenantsApi.md#AddUserTenant) | **Post** /api/accounts/{accountId}/users | Create a User For a Tenant
[**CreateTenantSubtenantGroup**](TenantsApi.md#CreateTenantSubtenantGroup) | **Post** /api/accounts/{accountId}/groups | Create a Group for Subtenant
[**GetTenant**](TenantsApi.md#GetTenant) | **Get** /api/accounts/{id} | Get tenant
[**GetTenantSubtenantGroup**](TenantsApi.md#GetTenantSubtenantGroup) | **Get** /api/accounts/{accountId}/groups/{id} | Get a Specific Group for Subtenant
[**ListTenantSubtenantGroups**](TenantsApi.md#ListTenantSubtenantGroups) | **Get** /api/accounts/{accountId}/groups | Get Subtenant Groups
[**ListTenants**](TenantsApi.md#ListTenants) | **Get** /api/accounts | List All Tenants
[**ListTenantsAvailableRoles**](TenantsApi.md#ListTenantsAvailableRoles) | **Get** /api/accounts/available-roles | List available roles for a tenant
[**RemoveTenant**](TenantsApi.md#RemoveTenant) | **Delete** /api/accounts/{id} | Delete a Specific Tenant
[**RemoveTenantSubtenantGroup**](TenantsApi.md#RemoveTenantSubtenantGroup) | **Delete** /api/accounts/{accountId}/groups/{id} | Delete a Group for Subtenant
[**UpdateTenant**](TenantsApi.md#UpdateTenant) | **Put** /api/accounts/{id} | Update tenant
[**UpdateTenantSubtenantGroup**](TenantsApi.md#UpdateTenantSubtenantGroup) | **Put** /api/accounts/{accountId}/groups/{id} | Updating a Group for Subtenant
[**UpdateTenantSubtenantGroupZones**](TenantsApi.md#UpdateTenantSubtenantGroupZones) | **Put** /api/accounts/{accountId}/groups/{id}/update-zones | Updating Group Zones for Subtenant



## AddTenant

> map[string]interface{} AddTenant(ctx).InlineObject238(inlineObject238).Execute()

Create a Tenant



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
    inlineObject238 := *openapiclient.NewInlineObject238(*openapiclient.NewApiAccountsAccount("Test Tenant")) // InlineObject238 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.AddTenant(context.Background()).InlineObject238(inlineObject238).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.AddTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.AddTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject238** | [**InlineObject238**](InlineObject238.md) |  | 

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


## AddUserTenant

> map[string]interface{} AddUserTenant(ctx, accountId).InlineObject243(inlineObject243).Execute()

Create a User For a Tenant



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    inlineObject243 := *openapiclient.NewInlineObject243(*openapiclient.NewUserCreate("jsmith", "jsmith@email.com", "Passw0rd!", []openapiclient.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites{*openapiclient.NewApiBlueprintsIdUpdatePermissionsResourcePermissionSites()})) // InlineObject243 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.AddUserTenant(context.Background(), accountId).InlineObject243(inlineObject243).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.AddUserTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.AddUserTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject243** | [**InlineObject243**](InlineObject243.md) |  | 

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


## CreateTenantSubtenantGroup

> InlineResponse200152 CreateTenantSubtenantGroup(ctx, accountId).InlineObject240(inlineObject240).Execute()

Create a Group for Subtenant



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    inlineObject240 := *openapiclient.NewInlineObject240(*openapiclient.NewApiAccountsAccountIdGroupsGroup("Name_example")) // InlineObject240 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.CreateTenantSubtenantGroup(context.Background(), accountId).InlineObject240(inlineObject240).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.CreateTenantSubtenantGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenantSubtenantGroup`: InlineResponse200152
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.CreateTenantSubtenantGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantSubtenantGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject240** | [**InlineObject240**](InlineObject240.md) |  | 

### Return type

[**InlineResponse200152**](inline_response_200_152.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenant

> InlineResponse200150 GetTenant(ctx, id).Execute()

Get tenant



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
    resp, r, err := api_client.TenantsApi.GetTenant(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.GetTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenant`: InlineResponse200150
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200150**](inline_response_200_150.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantSubtenantGroup

> InlineResponse200153 GetTenantSubtenantGroup(ctx, accountId, id).Execute()

Get a Specific Group for Subtenant



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.GetTenantSubtenantGroup(context.Background(), accountId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.GetTenantSubtenantGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantSubtenantGroup`: InlineResponse200153
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.GetTenantSubtenantGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantSubtenantGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200153**](inline_response_200_153.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTenantSubtenantGroups

> map[string]interface{} ListTenantSubtenantGroups(ctx, accountId).Phrase(phrase).Name(name).LastUpdated(lastUpdated).Execute()

Get Subtenant Groups



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.ListTenantSubtenantGroups(context.Background(), accountId).Phrase(phrase).Name(name).LastUpdated(lastUpdated).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.ListTenantSubtenantGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantSubtenantGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.ListTenantSubtenantGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTenantSubtenantGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 

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


## ListTenants

> map[string]interface{} ListTenants(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).LastUpdated(lastUpdated).Execute()

List All Tenants



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
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.ListTenants(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).LastUpdated(lastUpdated).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.ListTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenants`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.ListTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 

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


## ListTenantsAvailableRoles

> TenantsAvailableRoles ListTenantsAvailableRoles(ctx).Execute()

List available roles for a tenant



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
    resp, r, err := api_client.TenantsApi.ListTenantsAvailableRoles(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.ListTenantsAvailableRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantsAvailableRoles`: TenantsAvailableRoles
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.ListTenantsAvailableRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTenantsAvailableRolesRequest struct via the builder pattern


### Return type

[**TenantsAvailableRoles**](tenantsAvailableRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTenant

> Model200Success RemoveTenant(ctx, id).RemoveResources(removeResources).Execute()

Delete a Specific Tenant



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
    removeResources := true // bool | Remove Resources. This will delete all the managed resources in the tenant. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.RemoveTenant(context.Background(), id).RemoveResources(removeResources).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.RemoveTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTenant`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.RemoveTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeResources** | **bool** | Remove Resources. This will delete all the managed resources in the tenant. | [default to false]

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


## RemoveTenantSubtenantGroup

> Model200Success RemoveTenantSubtenantGroup(ctx, accountId, id).Execute()

Delete a Group for Subtenant



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.RemoveTenantSubtenantGroup(context.Background(), accountId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.RemoveTenantSubtenantGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTenantSubtenantGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.RemoveTenantSubtenantGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTenantSubtenantGroupRequest struct via the builder pattern


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


## UpdateTenant

> InlineResponse200151 UpdateTenant(ctx, id).InlineObject239(inlineObject239).Execute()

Update tenant



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
    inlineObject239 := *openapiclient.NewInlineObject239(map[string]interface{}(123)) // InlineObject239 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.UpdateTenant(context.Background(), id).InlineObject239(inlineObject239).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenant`: InlineResponse200151
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject239** | [**InlineObject239**](InlineObject239.md) |  | 

### Return type

[**InlineResponse200151**](inline_response_200_151.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantSubtenantGroup

> InlineResponse200152 UpdateTenantSubtenantGroup(ctx, accountId, id).InlineObject241(inlineObject241).Execute()

Updating a Group for Subtenant



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject241 := *openapiclient.NewInlineObject241(*openapiclient.NewApiAccountsAccountIdGroupsIdGroup()) // InlineObject241 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.UpdateTenantSubtenantGroup(context.Background(), accountId, id).InlineObject241(inlineObject241).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.UpdateTenantSubtenantGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantSubtenantGroup`: InlineResponse200152
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.UpdateTenantSubtenantGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantSubtenantGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject241** | [**InlineObject241**](InlineObject241.md) |  | 

### Return type

[**InlineResponse200152**](inline_response_200_152.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantSubtenantGroupZones

> Model200Success UpdateTenantSubtenantGroupZones(ctx, accountId, id).InlineObject242(inlineObject242).Execute()

Updating Group Zones for Subtenant



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
    accountId := int64(7) // int64 | The ID of the subtenant account
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject242 := *openapiclient.NewInlineObject242(*openapiclient.NewApiAccountsAccountIdGroupsIdUpdateZonesGroup([]openapiclient.ApiStorageVolumesStorageVolumeStorageServer{*openapiclient.NewApiStorageVolumesStorageVolumeStorageServer(int64(123))})) // InlineObject242 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.UpdateTenantSubtenantGroupZones(context.Background(), accountId, id).InlineObject242(inlineObject242).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.UpdateTenantSubtenantGroupZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantSubtenantGroupZones`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.UpdateTenantSubtenantGroupZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantSubtenantGroupZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject242** | [**InlineObject242**](InlineObject242.md) |  | 

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

