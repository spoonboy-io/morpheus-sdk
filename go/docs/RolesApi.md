# \RolesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoles**](RolesApi.md#AddRoles) | **Post** /api/roles | Create role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /api/roles/{id} | Delete role
[**GetRole**](RolesApi.md#GetRole) | **Get** /api/roles/{id} | Get role
[**ListRoles**](RolesApi.md#ListRoles) | **Get** /api/roles | List roles
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /api/roles/{id} | Update role
[**UpdateRoleBlueprintAccess**](RolesApi.md#UpdateRoleBlueprintAccess) | **Put** /api/roles/{id}/update-blueprint | Customizing Blueprint Access
[**UpdateRoleCatalogItemTypeAccess**](RolesApi.md#UpdateRoleCatalogItemTypeAccess) | **Put** /api/roles/{id}/update-catalog-item-type | Customizing Catalog Item Type Access
[**UpdateRoleCloudAccess**](RolesApi.md#UpdateRoleCloudAccess) | **Put** /api/roles/{id}/update-cloud | Customizing Cloud Access
[**UpdateRoleGroupAccess**](RolesApi.md#UpdateRoleGroupAccess) | **Put** /api/roles/{id}/update-group | Customizing Group Access
[**UpdateRoleInstanceTypeAccess**](RolesApi.md#UpdateRoleInstanceTypeAccess) | **Put** /api/roles/{id}/update-instance-type | Customizing Instance Type Access
[**UpdateRolePermission**](RolesApi.md#UpdateRolePermission) | **Put** /api/roles/{id}/update-permission | Updating Role Permissions
[**UpdateRolePersonaAccess**](RolesApi.md#UpdateRolePersonaAccess) | **Put** /api/roles/{id}/update-persona | Customizing Persona Access
[**UpdateRoleReportTypeAccess**](RolesApi.md#UpdateRoleReportTypeAccess) | **Put** /api/roles/{id}/update-report-type | Customizing Report Type Access
[**UpdateRoleTaskAccess**](RolesApi.md#UpdateRoleTaskAccess) | **Put** /api/roles/{id}/update-task | Customizing Task Access
[**UpdateRoleVDIPoolAccess**](RolesApi.md#UpdateRoleVDIPoolAccess) | **Put** /api/roles/{id}/update-vdi-pool | Customizing VDI Pool Access
[**UpdateRoleWorkflowAccess**](RolesApi.md#UpdateRoleWorkflowAccess) | **Put** /api/roles/{id}/update-task-set | Customizing Workflow Access



## AddRoles

> Role AddRoles(ctx).IncludeDefaultAccess(includeDefaultAccess).InlineObject208(inlineObject208).Execute()

Create role



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
    includeDefaultAccess := true // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)
    inlineObject208 := *openapiclient.NewInlineObject208(*openapiclient.NewApiRolesRole("Authority_example")) // InlineObject208 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.AddRoles(context.Background()).IncludeDefaultAccess(includeDefaultAccess).InlineObject208(inlineObject208).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoles`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDefaultAccess** | **bool** | Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | 
 **inlineObject208** | [**InlineObject208**](InlineObject208.md) |  | 

### Return type

[**Role**](role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> Model200Success DeleteRole(ctx, id).Execute()

Delete role



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
    resp, r, err := api_client.RolesApi.DeleteRole(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRole`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## GetRole

> Role GetRole(ctx, id).IncludeDefaultAccess(includeDefaultAccess).Execute()

Get role



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
    includeDefaultAccess := true // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.GetRole(context.Background(), id).IncludeDefaultAccess(includeDefaultAccess).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDefaultAccess** | **bool** | Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | 

### Return type

[**Role**](role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> map[string]interface{} ListRoles(ctx).Phrase(phrase).Max(max).Offset(offset).Sort(sort).Direction(direction).Authority(authority).RoleType(roleType).Execute()

List roles



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
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    authority := "authority_example" // string | Filter by authority (optional)
    roleType := "roleType_example" // string | Filter by roleType (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.ListRoles(context.Background()).Phrase(phrase).Max(max).Offset(offset).Sort(sort).Direction(direction).Authority(authority).RoleType(roleType).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **authority** | **string** | Filter by authority | 
 **roleType** | **string** | Filter by roleType | 

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


## UpdateRole

> Role UpdateRole(ctx, id).IncludeDefaultAccess(includeDefaultAccess).InlineObject209(inlineObject209).Execute()

Update role



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
    includeDefaultAccess := true // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)
    inlineObject209 := *openapiclient.NewInlineObject209(*openapiclient.NewApiRolesIdRole()) // InlineObject209 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.UpdateRole(context.Background(), id).IncludeDefaultAccess(includeDefaultAccess).InlineObject209(inlineObject209).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDefaultAccess** | **bool** | Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | 
 **inlineObject209** | [**InlineObject209**](InlineObject209.md) |  | 

### Return type

[**Role**](role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleBlueprintAccess

> map[string]interface{} UpdateRoleBlueprintAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Blueprint Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleBlueprintAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleBlueprintAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleBlueprintAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleBlueprintAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleBlueprintAccessRequest struct via the builder pattern


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


## UpdateRoleCatalogItemTypeAccess

> map[string]interface{} UpdateRoleCatalogItemTypeAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Catalog Item Type Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleCatalogItemTypeAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleCatalogItemTypeAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleCatalogItemTypeAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleCatalogItemTypeAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleCatalogItemTypeAccessRequest struct via the builder pattern


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


## UpdateRoleCloudAccess

> map[string]interface{} UpdateRoleCloudAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Cloud Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleCloudAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleCloudAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleCloudAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleCloudAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleCloudAccessRequest struct via the builder pattern


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


## UpdateRoleGroupAccess

> map[string]interface{} UpdateRoleGroupAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Group Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleGroupAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleGroupAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleGroupAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleGroupAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleGroupAccessRequest struct via the builder pattern


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


## UpdateRoleInstanceTypeAccess

> map[string]interface{} UpdateRoleInstanceTypeAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Instance Type Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleInstanceTypeAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleInstanceTypeAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleInstanceTypeAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleInstanceTypeAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleInstanceTypeAccessRequest struct via the builder pattern


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


## UpdateRolePermission

> map[string]interface{} UpdateRolePermission(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Updating Role Permissions



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
    resp, r, err := api_client.RolesApi.UpdateRolePermission(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRolePermission`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRolePermissionRequest struct via the builder pattern


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


## UpdateRolePersonaAccess

> map[string]interface{} UpdateRolePersonaAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Persona Access



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
    resp, r, err := api_client.RolesApi.UpdateRolePersonaAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRolePersonaAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRolePersonaAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRolePersonaAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRolePersonaAccessRequest struct via the builder pattern


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


## UpdateRoleReportTypeAccess

> map[string]interface{} UpdateRoleReportTypeAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Report Type Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleReportTypeAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleReportTypeAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleReportTypeAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleReportTypeAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleReportTypeAccessRequest struct via the builder pattern


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


## UpdateRoleTaskAccess

> map[string]interface{} UpdateRoleTaskAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Task Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleTaskAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleTaskAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleTaskAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleTaskAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleTaskAccessRequest struct via the builder pattern


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


## UpdateRoleVDIPoolAccess

> map[string]interface{} UpdateRoleVDIPoolAccess(ctx, id).InlineObject210(inlineObject210).Execute()

Customizing VDI Pool Access



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
    inlineObject210 := *openapiclient.NewInlineObject210(int32(123), "Access_example") // InlineObject210 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.UpdateRoleVDIPoolAccess(context.Background(), id).InlineObject210(inlineObject210).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleVDIPoolAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleVDIPoolAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleVDIPoolAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleVDIPoolAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject210** | [**InlineObject210**](InlineObject210.md) |  | 

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


## UpdateRoleWorkflowAccess

> map[string]interface{} UpdateRoleWorkflowAccess(ctx, id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Customizing Workflow Access



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
    resp, r, err := api_client.RolesApi.UpdateRoleWorkflowAccess(context.Background(), id).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRoleWorkflowAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleWorkflowAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRoleWorkflowAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleWorkflowAccessRequest struct via the builder pattern


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

