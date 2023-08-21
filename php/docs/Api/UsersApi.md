# OpenAPI\Client\UsersApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addUser()**](UsersApi.md#addUser) | **POST** /api/users | Create a New User
[**deleteUser()**](UsersApi.md#deleteUser) | **DELETE** /api/users/{id} | Delete a User
[**deleteUserSettingsAccessToken()**](UsersApi.md#deleteUserSettingsAccessToken) | **PUT** /api/user-settings/clear-access-token | Revoke API Access Token
[**deleteUserSettingsAvatar()**](UsersApi.md#deleteUserSettingsAvatar) | **DELETE** /api/user-settings/avatar | Delete Avatar
[**deleteUserSettingsDesktopBackground()**](UsersApi.md#deleteUserSettingsDesktopBackground) | **DELETE** /api/user-settings/desktop-background | Delete Desktop Background
[**getUser()**](UsersApi.md#getUser) | **GET** /api/users/{id} | Get a Specific User
[**getUserPermissions()**](UsersApi.md#getUserPermissions) | **GET** /api/users/{id}/permissions | Get a Specific User Permissions
[**getUserSettingsApiClients()**](UsersApi.md#getUserSettingsApiClients) | **GET** /api/user-settings/api-clients | Get Available API Clients
[**listUserSettings()**](UsersApi.md#listUserSettings) | **GET** /api/user-settings | User Settings
[**listUsers()**](UsersApi.md#listUsers) | **GET** /api/users | List All Users
[**listUsersAvailableRoles()**](UsersApi.md#listUsersAvailableRoles) | **GET** /api/users/available-roles | List available roles for a user
[**updateUserSettings()**](UsersApi.md#updateUserSettings) | **PUT** /api/user-settings | Update User Settings
[**updateUserSettingsAccessToken()**](UsersApi.md#updateUserSettingsAccessToken) | **PUT** /api/user-settings/regenerate-access-token | Regenerate API Access Token
[**updateUserSettingsAvatar()**](UsersApi.md#updateUserSettingsAvatar) | **POST** /api/user-settings/avatar | Update Avatar
[**updateUserSettingsDesktopBackground()**](UsersApi.md#updateUserSettingsDesktopBackground) | **POST** /api/user-settings/desktop-background | Update Desktop Background
[**updateUsers()**](UsersApi.md#updateUsers) | **PUT** /api/users/{id} | Update user
[**whoami()**](UsersApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions


## `addUser()`

```php
addUser($account_id, $inline_object255): object
```

Create a New User

Create a new user.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 56; // int | Tenant ID, create user in a sub tenant account instead of your own.
$inline_object255 = new \OpenAPI\Client\Model\InlineObject255(); // \OpenAPI\Client\Model\InlineObject255

try {
    $result = $apiInstance->addUser($account_id, $inline_object255);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->addUser: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Tenant ID, create user in a sub tenant account instead of your own. | [optional]
 **inline_object255** | [**\OpenAPI\Client\Model\InlineObject255**](../Model/InlineObject255.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteUser()`

```php
deleteUser($id): \OpenAPI\Client\Model\Model200Success
```

Delete a User

Delete an existing user.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteUser($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->deleteUser: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteUserSettingsAccessToken()`

```php
deleteUserSettingsAccessToken($user_id, $client_id): \OpenAPI\Client\Model\Model200Success
```

Revoke API Access Token

This endpoint revokes your API access token for the specified client.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user
$client_id = morph-cli; // string | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.

try {
    $result = $apiInstance->deleteUserSettingsAccessToken($user_id, $client_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->deleteUserSettingsAccessToken: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **client_id** | **string**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteUserSettingsAvatar()`

```php
deleteUserSettingsAvatar($user_id): \OpenAPI\Client\Model\Model200Success
```

Delete Avatar

Delete your avatar image.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user

try {
    $result = $apiInstance->deleteUserSettingsAvatar($user_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->deleteUserSettingsAvatar: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteUserSettingsDesktopBackground()`

```php
deleteUserSettingsDesktopBackground($user_id): \OpenAPI\Client\Model\Model200Success
```

Delete Desktop Background

Delete your desktop background image.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user

try {
    $result = $apiInstance->deleteUserSettingsDesktopBackground($user_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->deleteUserSettingsDesktopBackground: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getUser()`

```php
getUser($id, $include_access): \OpenAPI\Client\Model\InlineResponse200158
```

Get a Specific User

This endpoint will retrieve a specific user by id if the user has permission to access the user.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$include_access = true; // bool | Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s).

try {
    $result = $apiInstance->getUser($id, $include_access);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->getUser: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **include_access** | **bool**| Include &#x60;access&#x60; information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200158**](../Model/InlineResponse200158.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getUserPermissions()`

```php
getUserPermissions($id): \OpenAPI\Client\Model\InlineResponse200159
```

Get a Specific User Permissions

This will list all the permissions for a specific user.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getUserPermissions($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->getUserPermissions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200159**](../Model/InlineResponse200159.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getUserSettingsApiClients()`

```php
getUserSettingsApiClients($user_id): \OpenAPI\Client\Model\InlineResponse200157
```

Get Available API Clients

This endpoint retrieves a list of available API clients.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user

try {
    $result = $apiInstance->getUserSettingsApiClients($user_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->getUserSettingsApiClients: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200157**](../Model/InlineResponse200157.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listUserSettings()`

```php
listUserSettings($user_id): UserSettings
```

User Settings

Provides API for managing your own user settings and api access tokens.  This endpoint retrieves your user settings and API access token information.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user

try {
    $result = $apiInstance->listUserSettings($user_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->listUserSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**UserSettings**](../Model/UserSettings.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listUsers()`

```php
listUsers($max, $offset, $sort, $direction, $phrase, $username, $email, $role_id, $last_updated, $account_id, $global): object
```

List All Users

This endpoint retrieves all users in the current user's tenant account. Master tenant users with permission to manage subtenants can use `global=true` to find users across all tenants.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on username or email
$username = administrator; // string | Username
$email = mytest@email.com; // string | E-Mail Address
$role_id = 13; // int | Filter by Role ID (supports multiple values)
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
$global = false; // bool | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users.

try {
    $result = $apiInstance->listUsers($max, $offset, $sort, $direction, $phrase, $username, $email, $role_id, $last_updated, $account_id, $global);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->listUsers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on username or email | [optional]
 **username** | **string**| Username | [optional]
 **email** | **string**| E-Mail Address | [optional]
 **role_id** | **int**| Filter by Role ID (supports multiple values) | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **global** | **bool**| Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] [default to false]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listUsersAvailableRoles()`

```php
listUsersAvailableRoles($account_id): \OpenAPI\Client\Model\UsersAvailableRoles
```

List available roles for a user

Get a list of roles that can be assigned to a user.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 56; // int | Tenant ID, find roles available to users of a sub tenant account.

try {
    $result = $apiInstance->listUsersAvailableRoles($account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->listUsersAvailableRoles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Tenant ID, find roles available to users of a sub tenant account. | [optional]

### Return type

[**\OpenAPI\Client\Model\UsersAvailableRoles**](../Model/UsersAvailableRoles.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateUserSettings()`

```php
updateUserSettings($user_id, $inline_object252): \OpenAPI\Client\Model\Model200Success
```

Update User Settings

This endpoint allows updating user settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user
$inline_object252 = new \OpenAPI\Client\Model\InlineObject252(); // \OpenAPI\Client\Model\InlineObject252

try {
    $result = $apiInstance->updateUserSettings($user_id, $inline_object252);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->updateUserSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **inline_object252** | [**\OpenAPI\Client\Model\InlineObject252**](../Model/InlineObject252.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateUserSettingsAccessToken()`

```php
updateUserSettingsAccessToken($user_id, $client_id): UserSettingsRegenerateAccessToken
```

Regenerate API Access Token

This endpoint regenerates your API access token for the specified client. If a current token exists, it is revoked and a new token is returned.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user
$client_id = morph-cli; // string | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.

try {
    $result = $apiInstance->updateUserSettingsAccessToken($user_id, $client_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->updateUserSettingsAccessToken: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **client_id** | **string**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional]

### Return type

[**UserSettingsRegenerateAccessToken**](../Model/UserSettingsRegenerateAccessToken.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateUserSettingsAvatar()`

```php
updateUserSettingsAvatar($user_id, $user_avatar): \OpenAPI\Client\Model\Model200Success
```

Update Avatar

This endpoint updates your avatar image. Expects multipart form data as the request format, not JSON.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user
$user_avatar = 'user_avatar_example'; // string

try {
    $result = $apiInstance->updateUserSettingsAvatar($user_id, $user_avatar);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->updateUserSettingsAvatar: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **user_avatar** | **string**|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateUserSettingsDesktopBackground()`

```php
updateUserSettingsDesktopBackground($user_id, $user_desktop_background): \OpenAPI\Client\Model\Model200Success
```

Update Desktop Background

This endpoint updates your desktop background image that is used in the Virtual Desktop persona. Expects multipart form data as the request format, not JSON.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$user_id = 56; // int | ID of User (Only available to the master tenant.)  Default is the current user
$user_desktop_background = 'user_desktop_background_example'; // string

try {
    $result = $apiInstance->updateUserSettingsDesktopBackground($user_id, $user_desktop_background);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->updateUserSettingsDesktopBackground: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **user_desktop_background** | **string**|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateUsers()`

```php
updateUsers($id, $inline_object256): object
```

Update user

Update an existing user.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object256 = new \OpenAPI\Client\Model\InlineObject256(); // \OpenAPI\Client\Model\InlineObject256

try {
    $result = $apiInstance->updateUsers($id, $inline_object256);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->updateUsers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object256** | [**\OpenAPI\Client\Model\InlineObject256**](../Model/InlineObject256.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `whoami()`

```php
whoami(): \OpenAPI\Client\Model\InlineResponse200167
```

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->whoami();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsersApi->whoami: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200167**](../Model/InlineResponse200167.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
