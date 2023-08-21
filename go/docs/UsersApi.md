# \UsersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](UsersApi.md#AddUser) | **Post** /api/users | Create a New User
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /api/users/{id} | Delete a User
[**DeleteUserSettingsAccessToken**](UsersApi.md#DeleteUserSettingsAccessToken) | **Put** /api/user-settings/clear-access-token | Revoke API Access Token
[**DeleteUserSettingsAvatar**](UsersApi.md#DeleteUserSettingsAvatar) | **Delete** /api/user-settings/avatar | Delete Avatar
[**DeleteUserSettingsDesktopBackground**](UsersApi.md#DeleteUserSettingsDesktopBackground) | **Delete** /api/user-settings/desktop-background | Delete Desktop Background
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/users/{id} | Get a Specific User
[**GetUserPermissions**](UsersApi.md#GetUserPermissions) | **Get** /api/users/{id}/permissions | Get a Specific User Permissions
[**GetUserSettingsApiClients**](UsersApi.md#GetUserSettingsApiClients) | **Get** /api/user-settings/api-clients | Get Available API Clients
[**ListUserSettings**](UsersApi.md#ListUserSettings) | **Get** /api/user-settings | User Settings
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /api/users | List All Users
[**ListUsersAvailableRoles**](UsersApi.md#ListUsersAvailableRoles) | **Get** /api/users/available-roles | List available roles for a user
[**UpdateUserSettings**](UsersApi.md#UpdateUserSettings) | **Put** /api/user-settings | Update User Settings
[**UpdateUserSettingsAccessToken**](UsersApi.md#UpdateUserSettingsAccessToken) | **Put** /api/user-settings/regenerate-access-token | Regenerate API Access Token
[**UpdateUserSettingsAvatar**](UsersApi.md#UpdateUserSettingsAvatar) | **Post** /api/user-settings/avatar | Update Avatar
[**UpdateUserSettingsDesktopBackground**](UsersApi.md#UpdateUserSettingsDesktopBackground) | **Post** /api/user-settings/desktop-background | Update Desktop Background
[**UpdateUsers**](UsersApi.md#UpdateUsers) | **Put** /api/users/{id} | Update user
[**Whoami**](UsersApi.md#Whoami) | **Get** /api/whoami | Retrieves information about current user roles and permissions



## AddUser

> map[string]interface{} AddUser(ctx).AccountId(accountId).InlineObject255(inlineObject255).Execute()

Create a New User



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
    accountId := int64(789) // int64 | Tenant ID, create user in a sub tenant account instead of your own. (optional)
    inlineObject255 := *openapiclient.NewInlineObject255(*openapiclient.NewUserCreate("jsmith", "jsmith@email.com", "Passw0rd!", []openapiclient.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites{*openapiclient.NewApiBlueprintsIdUpdatePermissionsResourcePermissionSites()})) // InlineObject255 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.AddUser(context.Background()).AccountId(accountId).InlineObject255(inlineObject255).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.AddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.AddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Tenant ID, create user in a sub tenant account instead of your own. | 
 **inlineObject255** | [**InlineObject255**](InlineObject255.md) |  | 

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


## DeleteUser

> Model200Success DeleteUser(ctx, id).Execute()

Delete a User



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
    resp, r, err := api_client.UsersApi.DeleteUser(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DeleteUserSettingsAccessToken

> Model200Success DeleteUserSettingsAccessToken(ctx).UserId(userId).ClientId(clientId).Execute()

Revoke API Access Token



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    clientId := "morph-cli" // string | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteUserSettingsAccessToken(context.Background()).UserId(userId).ClientId(clientId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUserSettingsAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserSettingsAccessToken`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteUserSettingsAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **clientId** | **string** | Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | 

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


## DeleteUserSettingsAvatar

> Model200Success DeleteUserSettingsAvatar(ctx).UserId(userId).Execute()

Delete Avatar



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteUserSettingsAvatar(context.Background()).UserId(userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUserSettingsAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserSettingsAvatar`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteUserSettingsAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

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


## DeleteUserSettingsDesktopBackground

> Model200Success DeleteUserSettingsDesktopBackground(ctx).UserId(userId).Execute()

Delete Desktop Background



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteUserSettingsDesktopBackground(context.Background()).UserId(userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUserSettingsDesktopBackground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserSettingsDesktopBackground`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteUserSettingsDesktopBackground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsDesktopBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

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


## GetUser

> InlineResponse200158 GetUser(ctx, id).IncludeAccess(includeAccess).Execute()

Get a Specific User



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
    includeAccess := true // bool | Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUser(context.Background(), id).IncludeAccess(includeAccess).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: InlineResponse200158
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAccess** | **bool** | Include &#x60;access&#x60; information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). | 

### Return type

[**InlineResponse200158**](inline_response_200_158.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPermissions

> InlineResponse200159 GetUserPermissions(ctx, id).Execute()

Get a Specific User Permissions



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
    resp, r, err := api_client.UsersApi.GetUserPermissions(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPermissions`: InlineResponse200159
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsApiClients

> InlineResponse200157 GetUserSettingsApiClients(ctx).UserId(userId).Execute()

Get Available API Clients



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUserSettingsApiClients(context.Background()).UserId(userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserSettingsApiClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettingsApiClients`: InlineResponse200157
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserSettingsApiClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsApiClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

### Return type

[**InlineResponse200157**](inline_response_200_157.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserSettings

> UserSettings ListUserSettings(ctx).UserId(userId).Execute()

User Settings



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUserSettings(context.Background()).UserId(userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserSettings`: UserSettings
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

### Return type

[**UserSettings**](userSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> map[string]interface{} ListUsers(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Username(username).Email(email).RoleId(roleId).LastUpdated(lastUpdated).AccountId(accountId).Global(global).Execute()

List All Users



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
    phrase := "phrase_example" // string | Search phrase for partial matches on username or email (optional)
    username := "administrator" // string | Username (optional)
    email := "mytest@email.com" // string | E-Mail Address (optional)
    roleId := int64(13) // int64 | Filter by Role ID (supports multiple values) (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    global := true // bool | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUsers(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Username(username).Email(email).RoleId(roleId).LastUpdated(lastUpdated).AccountId(accountId).Global(global).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on username or email | 
 **username** | **string** | Username | 
 **email** | **string** | E-Mail Address | 
 **roleId** | **int64** | Filter by Role ID (supports multiple values) | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 
 **global** | **bool** | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [default to false]

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


## ListUsersAvailableRoles

> UsersAvailableRoles ListUsersAvailableRoles(ctx).AccountId(accountId).Execute()

List available roles for a user



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
    accountId := int64(789) // int64 | Tenant ID, find roles available to users of a sub tenant account. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUsersAvailableRoles(context.Background()).AccountId(accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsersAvailableRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsersAvailableRoles`: UsersAvailableRoles
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsersAvailableRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersAvailableRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Tenant ID, find roles available to users of a sub tenant account. | 

### Return type

[**UsersAvailableRoles**](usersAvailableRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettings

> Model200Success UpdateUserSettings(ctx).UserId(userId).InlineObject252(inlineObject252).Execute()

Update User Settings



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    inlineObject252 := *openapiclient.NewInlineObject252() // InlineObject252 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUserSettings(context.Background()).UserId(userId).InlineObject252(inlineObject252).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **inlineObject252** | [**InlineObject252**](InlineObject252.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsAccessToken

> UserSettingsRegenerateAccessToken UpdateUserSettingsAccessToken(ctx).UserId(userId).ClientId(clientId).Execute()

Regenerate API Access Token



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    clientId := "morph-cli" // string | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUserSettingsAccessToken(context.Background()).UserId(userId).ClientId(clientId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUserSettingsAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserSettingsAccessToken`: UserSettingsRegenerateAccessToken
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUserSettingsAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **clientId** | **string** | Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | 

### Return type

[**UserSettingsRegenerateAccessToken**](userSettingsRegenerateAccessToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsAvatar

> Model200Success UpdateUserSettingsAvatar(ctx).UserId(userId).UserAvatar(userAvatar).Execute()

Update Avatar



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    userAvatar := "userAvatar_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUserSettingsAvatar(context.Background()).UserId(userId).UserAvatar(userAvatar).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUserSettingsAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserSettingsAvatar`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUserSettingsAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **userAvatar** | **string** |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsDesktopBackground

> Model200Success UpdateUserSettingsDesktopBackground(ctx).UserId(userId).UserDesktopBackground(userDesktopBackground).Execute()

Update Desktop Background



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
    userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    userDesktopBackground := "userDesktopBackground_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUserSettingsDesktopBackground(context.Background()).UserId(userId).UserDesktopBackground(userDesktopBackground).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUserSettingsDesktopBackground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserSettingsDesktopBackground`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUserSettingsDesktopBackground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsDesktopBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **userDesktopBackground** | **string** |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsers

> map[string]interface{} UpdateUsers(ctx, id).InlineObject256(inlineObject256).Execute()

Update user



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
    inlineObject256 := *openapiclient.NewInlineObject256(*openapiclient.NewApiUsersIdUser()) // InlineObject256 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUsers(context.Background(), id).InlineObject256(inlineObject256).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUsers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject256** | [**InlineObject256**](InlineObject256.md) |  | 

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


## Whoami

> InlineResponse200167 Whoami(ctx).Execute()

Retrieves information about current user roles and permissions



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
    resp, r, err := api_client.UsersApi.Whoami(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.Whoami``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Whoami`: InlineResponse200167
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.Whoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiRequest struct via the builder pattern


### Return type

[**InlineResponse200167**](inline_response_200_167.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

