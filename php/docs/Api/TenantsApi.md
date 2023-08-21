# OpenAPI\Client\TenantsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addTenant()**](TenantsApi.md#addTenant) | **POST** /api/accounts | Create a Tenant
[**addUserTenant()**](TenantsApi.md#addUserTenant) | **POST** /api/accounts/{accountId}/users | Create a User For a Tenant
[**createTenantSubtenantGroup()**](TenantsApi.md#createTenantSubtenantGroup) | **POST** /api/accounts/{accountId}/groups | Create a Group for Subtenant
[**getTenant()**](TenantsApi.md#getTenant) | **GET** /api/accounts/{id} | Get tenant
[**getTenantSubtenantGroup()**](TenantsApi.md#getTenantSubtenantGroup) | **GET** /api/accounts/{accountId}/groups/{id} | Get a Specific Group for Subtenant
[**listTenantSubtenantGroups()**](TenantsApi.md#listTenantSubtenantGroups) | **GET** /api/accounts/{accountId}/groups | Get Subtenant Groups
[**listTenants()**](TenantsApi.md#listTenants) | **GET** /api/accounts | List All Tenants
[**listTenantsAvailableRoles()**](TenantsApi.md#listTenantsAvailableRoles) | **GET** /api/accounts/available-roles | List available roles for a tenant
[**removeTenant()**](TenantsApi.md#removeTenant) | **DELETE** /api/accounts/{id} | Delete a Specific Tenant
[**removeTenantSubtenantGroup()**](TenantsApi.md#removeTenantSubtenantGroup) | **DELETE** /api/accounts/{accountId}/groups/{id} | Delete a Group for Subtenant
[**updateTenant()**](TenantsApi.md#updateTenant) | **PUT** /api/accounts/{id} | Update tenant
[**updateTenantSubtenantGroup()**](TenantsApi.md#updateTenantSubtenantGroup) | **PUT** /api/accounts/{accountId}/groups/{id} | Updating a Group for Subtenant
[**updateTenantSubtenantGroupZones()**](TenantsApi.md#updateTenantSubtenantGroupZones) | **PUT** /api/accounts/{accountId}/groups/{id}/update-zones | Updating Group Zones for Subtenant


## `addTenant()`

```php
addTenant($inline_object238): object
```

Create a Tenant

Create a new tenant. This new account will be a sub-tenant with the master tenant as its parent.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object238 = new \OpenAPI\Client\Model\InlineObject238(); // \OpenAPI\Client\Model\InlineObject238

try {
    $result = $apiInstance->addTenant($inline_object238);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->addTenant: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object238** | [**\OpenAPI\Client\Model\InlineObject238**](../Model/InlineObject238.md)|  | [optional]

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

## `addUserTenant()`

```php
addUserTenant($account_id, $inline_object243): object
```

Create a User For a Tenant

Create a User For a Tenant.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$inline_object243 = new \OpenAPI\Client\Model\InlineObject243(); // \OpenAPI\Client\Model\InlineObject243

try {
    $result = $apiInstance->addUserTenant($account_id, $inline_object243);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->addUserTenant: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **inline_object243** | [**\OpenAPI\Client\Model\InlineObject243**](../Model/InlineObject243.md)|  | [optional]

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

## `createTenantSubtenantGroup()`

```php
createTenantSubtenantGroup($account_id, $inline_object240): \OpenAPI\Client\Model\InlineResponse200152
```

Create a Group for Subtenant

Create a Group for Subtenant.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$inline_object240 = new \OpenAPI\Client\Model\InlineObject240(); // \OpenAPI\Client\Model\InlineObject240

try {
    $result = $apiInstance->createTenantSubtenantGroup($account_id, $inline_object240);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->createTenantSubtenantGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **inline_object240** | [**\OpenAPI\Client\Model\InlineObject240**](../Model/InlineObject240.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200152**](../Model/InlineResponse200152.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getTenant()`

```php
getTenant($id): \OpenAPI\Client\Model\InlineResponse200150
```

Get tenant

Get details about a tenant

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getTenant($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->getTenant: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200150**](../Model/InlineResponse200150.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getTenantSubtenantGroup()`

```php
getTenantSubtenantGroup($account_id, $id): \OpenAPI\Client\Model\InlineResponse200153
```

Get a Specific Group for Subtenant

This endpoint retrieves a specific group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getTenantSubtenantGroup($account_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->getTenantSubtenantGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200153**](../Model/InlineResponse200153.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listTenantSubtenantGroups()`

```php
listTenantSubtenantGroups($account_id, $phrase, $name, $last_updated): object
```

Get Subtenant Groups

Groups belonging to a subtenant can be managed by the master account.  This endpoint retrieves all groups and a list of zones associated with the group by id.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)

try {
    $result = $apiInstance->listTenantSubtenantGroups($account_id, $phrase, $name, $last_updated);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->listTenantSubtenantGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

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

## `listTenants()`

```php
listTenants($max, $offset, $sort, $direction, $phrase, $name, $last_updated): object
```

List All Tenants

Get a list of tenants. A tenant is also referred to as an account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)

try {
    $result = $apiInstance->listTenants($max, $offset, $sort, $direction, $phrase, $name, $last_updated);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->listTenants: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

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

## `listTenantsAvailableRoles()`

```php
listTenantsAvailableRoles(): \OpenAPI\Client\Model\TenantsAvailableRoles
```

List available roles for a tenant

Get a list of available roles that can be assigned as the default base role for a sub tenant account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listTenantsAvailableRoles();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->listTenantsAvailableRoles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\TenantsAvailableRoles**](../Model/TenantsAvailableRoles.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `removeTenant()`

```php
removeTenant($id, $remove_resources): \OpenAPI\Client\Model\Model200Success
```

Delete a Specific Tenant

Delete an existing tenant. This action is not reversible and will result in the removal of all data pertaining to this tenant as well as potentially any provisioned assets depending on the value of `removeResources`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$remove_resources = false; // bool | Remove Resources. This will delete all the managed resources in the tenant.

try {
    $result = $apiInstance->removeTenant($id, $remove_resources);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->removeTenant: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_resources** | **bool**| Remove Resources. This will delete all the managed resources in the tenant. | [optional] [default to false]

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

## `removeTenantSubtenantGroup()`

```php
removeTenantSubtenantGroup($account_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Group for Subtenant

If a group has zones or servers still tied to it, a delete action will fail.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeTenantSubtenantGroup($account_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->removeTenantSubtenantGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
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

## `updateTenant()`

```php
updateTenant($id, $inline_object239): \OpenAPI\Client\Model\InlineResponse200151
```

Update tenant

Update an existing tenant.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object239 = new \OpenAPI\Client\Model\InlineObject239(); // \OpenAPI\Client\Model\InlineObject239

try {
    $result = $apiInstance->updateTenant($id, $inline_object239);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->updateTenant: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object239** | [**\OpenAPI\Client\Model\InlineObject239**](../Model/InlineObject239.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200151**](../Model/InlineResponse200151.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateTenantSubtenantGroup()`

```php
updateTenantSubtenantGroup($account_id, $id, $inline_object241): \OpenAPI\Client\Model\InlineResponse200152
```

Updating a Group for Subtenant

Updating a Group for Subtenant.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object241 = new \OpenAPI\Client\Model\InlineObject241(); // \OpenAPI\Client\Model\InlineObject241

try {
    $result = $apiInstance->updateTenantSubtenantGroup($account_id, $id, $inline_object241);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->updateTenantSubtenantGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object241** | [**\OpenAPI\Client\Model\InlineObject241**](../Model/InlineObject241.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200152**](../Model/InlineResponse200152.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateTenantSubtenantGroupZones()`

```php
updateTenantSubtenantGroupZones($account_id, $id, $inline_object242): \OpenAPI\Client\Model\Model200Success
```

Updating Group Zones for Subtenant

This will update the zones that are assigned to the group. Any zones that are not passed in the zones parameter will be removed from the group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\TenantsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 7; // int | The ID of the subtenant account
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object242 = new \OpenAPI\Client\Model\InlineObject242(); // \OpenAPI\Client\Model\InlineObject242

try {
    $result = $apiInstance->updateTenantSubtenantGroupZones($account_id, $id, $inline_object242);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TenantsApi->updateTenantSubtenantGroupZones: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object242** | [**\OpenAPI\Client\Model\InlineObject242**](../Model/InlineObject242.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
