# OpenAPI\Client\RolesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addRoles()**](RolesApi.md#addRoles) | **POST** /api/roles | Create role
[**deleteRole()**](RolesApi.md#deleteRole) | **DELETE** /api/roles/{id} | Delete role
[**getRole()**](RolesApi.md#getRole) | **GET** /api/roles/{id} | Get role
[**listRoles()**](RolesApi.md#listRoles) | **GET** /api/roles | List roles
[**updateRole()**](RolesApi.md#updateRole) | **PUT** /api/roles/{id} | Update role
[**updateRoleBlueprintAccess()**](RolesApi.md#updateRoleBlueprintAccess) | **PUT** /api/roles/{id}/update-blueprint | Customizing Blueprint Access
[**updateRoleCatalogItemTypeAccess()**](RolesApi.md#updateRoleCatalogItemTypeAccess) | **PUT** /api/roles/{id}/update-catalog-item-type | Customizing Catalog Item Type Access
[**updateRoleCloudAccess()**](RolesApi.md#updateRoleCloudAccess) | **PUT** /api/roles/{id}/update-cloud | Customizing Cloud Access
[**updateRoleGroupAccess()**](RolesApi.md#updateRoleGroupAccess) | **PUT** /api/roles/{id}/update-group | Customizing Group Access
[**updateRoleInstanceTypeAccess()**](RolesApi.md#updateRoleInstanceTypeAccess) | **PUT** /api/roles/{id}/update-instance-type | Customizing Instance Type Access
[**updateRolePermission()**](RolesApi.md#updateRolePermission) | **PUT** /api/roles/{id}/update-permission | Updating Role Permissions
[**updateRolePersonaAccess()**](RolesApi.md#updateRolePersonaAccess) | **PUT** /api/roles/{id}/update-persona | Customizing Persona Access
[**updateRoleReportTypeAccess()**](RolesApi.md#updateRoleReportTypeAccess) | **PUT** /api/roles/{id}/update-report-type | Customizing Report Type Access
[**updateRoleTaskAccess()**](RolesApi.md#updateRoleTaskAccess) | **PUT** /api/roles/{id}/update-task | Customizing Task Access
[**updateRoleVDIPoolAccess()**](RolesApi.md#updateRoleVDIPoolAccess) | **PUT** /api/roles/{id}/update-vdi-pool | Customizing VDI Pool Access
[**updateRoleWorkflowAccess()**](RolesApi.md#updateRoleWorkflowAccess) | **PUT** /api/roles/{id}/update-task-set | Customizing Workflow Access


## `addRoles()`

```php
addRoles($include_default_access, $inline_object208): Role
```

Create role

Create a new role.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$include_default_access = true; // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`
$inline_object208 = new \OpenAPI\Client\Model\InlineObject208(); // \OpenAPI\Client\Model\InlineObject208

try {
    $result = $apiInstance->addRoles($include_default_access, $inline_object208);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->addRoles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include_default_access** | **bool**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60; | [optional]
 **inline_object208** | [**\OpenAPI\Client\Model\InlineObject208**](../Model/InlineObject208.md)|  | [optional]

### Return type

[**Role**](../Model/Role.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteRole()`

```php
deleteRole($id): \OpenAPI\Client\Model\Model200Success
```

Delete role

Delete an existing role. A role cannot be deleted while it is still in use.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteRole($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->deleteRole: ', $e->getMessage(), PHP_EOL;
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

## `getRole()`

```php
getRole($id, $include_default_access): \OpenAPI\Client\Model\Role
```

Get role

Get details about a role

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$include_default_access = true; // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`

try {
    $result = $apiInstance->getRole($id, $include_default_access);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->getRole: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **include_default_access** | **bool**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60; | [optional]

### Return type

[**\OpenAPI\Client\Model\Role**](../Model/Role.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listRoles()`

```php
listRoles($phrase, $max, $offset, $sort, $direction, $authority, $role_type): object
```

List roles

Get a list of roles.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$authority = 'authority_example'; // string | Filter by authority
$role_type = 'role_type_example'; // string | Filter by roleType

try {
    $result = $apiInstance->listRoles($phrase, $max, $offset, $sort, $direction, $authority, $role_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->listRoles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **authority** | **string**| Filter by authority | [optional]
 **role_type** | **string**| Filter by roleType | [optional]

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

## `updateRole()`

```php
updateRole($id, $include_default_access, $inline_object209): Role
```

Update role

Update an existing role.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$include_default_access = true; // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`
$inline_object209 = new \OpenAPI\Client\Model\InlineObject209(); // \OpenAPI\Client\Model\InlineObject209

try {
    $result = $apiInstance->updateRole($id, $include_default_access, $inline_object209);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRole: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **include_default_access** | **bool**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60; | [optional]
 **inline_object209** | [**\OpenAPI\Client\Model\InlineObject209**](../Model/InlineObject209.md)|  | [optional]

### Return type

[**Role**](../Model/Role.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateRoleBlueprintAccess()`

```php
updateRoleBlueprintAccess($id, $unknown_base_type): object
```

Customizing Blueprint Access

Customizing Blueprint Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleBlueprintAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleBlueprintAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleCatalogItemTypeAccess()`

```php
updateRoleCatalogItemTypeAccess($id, $unknown_base_type): object
```

Customizing Catalog Item Type Access

Customizing Catalog Item Type Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleCatalogItemTypeAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleCatalogItemTypeAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleCloudAccess()`

```php
updateRoleCloudAccess($id, $unknown_base_type): object
```

Customizing Cloud Access

Customizing Cloud Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleCloudAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleCloudAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleGroupAccess()`

```php
updateRoleGroupAccess($id, $unknown_base_type): object
```

Customizing Group Access

Customizing Group Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleGroupAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleGroupAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleInstanceTypeAccess()`

```php
updateRoleInstanceTypeAccess($id, $unknown_base_type): object
```

Customizing Instance Type Access

Customizing Instance Type Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleInstanceTypeAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleInstanceTypeAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRolePermission()`

```php
updateRolePermission($id, $unknown_base_type): object
```

Updating Role Permissions

Update a feature permission or default permission category (group, cloud, persona, ect.)

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = {"permissionCode":"admin-users","access":"read"}; // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRolePermission($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRolePermission: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRolePersonaAccess()`

```php
updateRolePersonaAccess($id, $unknown_base_type): object
```

Customizing Persona Access

Customizing Persona Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRolePersonaAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRolePersonaAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleReportTypeAccess()`

```php
updateRoleReportTypeAccess($id, $unknown_base_type): object
```

Customizing Report Type Access

Customizing Report Type Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleReportTypeAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleReportTypeAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleTaskAccess()`

```php
updateRoleTaskAccess($id, $unknown_base_type): object
```

Customizing Task Access

Customizing Task Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleTaskAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleTaskAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `updateRoleVDIPoolAccess()`

```php
updateRoleVDIPoolAccess($id, $inline_object210): object
```

Customizing VDI Pool Access

Customizing VDI Pool Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object210 = new \OpenAPI\Client\Model\InlineObject210(); // \OpenAPI\Client\Model\InlineObject210

try {
    $result = $apiInstance->updateRoleVDIPoolAccess($id, $inline_object210);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleVDIPoolAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object210** | [**\OpenAPI\Client\Model\InlineObject210**](../Model/InlineObject210.md)|  | [optional]

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

## `updateRoleWorkflowAccess()`

```php
updateRoleWorkflowAccess($id, $unknown_base_type): object
```

Customizing Workflow Access

Customizing Workflow Access

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\RolesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateRoleWorkflowAccess($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RolesApi->updateRoleWorkflowAccess: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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
