# OpenAPI\Client\ResourcePoolsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**createResourcePoolGroup()**](ResourcePoolsApi.md#createResourcePoolGroup) | **POST** /api/resource-pools/groups | Create a Resource Pool Group
[**deleteResourcePoolGroup()**](ResourcePoolsApi.md#deleteResourcePoolGroup) | **DELETE** /api/resource-pools/groups/{id} | Delete a Resource Pool Group
[**getResourcePoolGroups()**](ResourcePoolsApi.md#getResourcePoolGroups) | **GET** /api/resource-pools/groups | Get all Resource Pool Groups
[**getresourcePoolGroup()**](ResourcePoolsApi.md#getresourcePoolGroup) | **GET** /api/resource-pools/groups/{id} | Get a Specific Resource Pool Group
[**updateResourcePoolGroup()**](ResourcePoolsApi.md#updateResourcePoolGroup) | **PUT** /api/resource-pools/groups/{id} | Update a Resource Pool Group


## `createResourcePoolGroup()`

```php
createResourcePoolGroup($inline_object206): \OpenAPI\Client\Model\InlineResponse200132
```

Create a Resource Pool Group

Use this command to create a resource pool group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ResourcePoolsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object206 = new \OpenAPI\Client\Model\InlineObject206(); // \OpenAPI\Client\Model\InlineObject206

try {
    $result = $apiInstance->createResourcePoolGroup($inline_object206);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ResourcePoolsApi->createResourcePoolGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object206** | [**\OpenAPI\Client\Model\InlineObject206**](../Model/InlineObject206.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200132**](../Model/InlineResponse200132.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteResourcePoolGroup()`

```php
deleteResourcePoolGroup($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Resource Pool Group

Will delete a Resource Pool Group from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ResourcePoolsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteResourcePoolGroup($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ResourcePoolsApi->deleteResourcePoolGroup: ', $e->getMessage(), PHP_EOL;
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

## `getResourcePoolGroups()`

```php
getResourcePoolGroups(): \OpenAPI\Client\Model\InlineResponse200131
```

Get all Resource Pool Groups

This endpoint retrieves all Resource Pool Groups associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ResourcePoolsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getResourcePoolGroups();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ResourcePoolsApi->getResourcePoolGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200131**](../Model/InlineResponse200131.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getresourcePoolGroup()`

```php
getresourcePoolGroup($id): \OpenAPI\Client\Model\InlineResponse200132
```

Get a Specific Resource Pool Group

This endpoint retrieves a specific Resource Pool Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ResourcePoolsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getresourcePoolGroup($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ResourcePoolsApi->getresourcePoolGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200132**](../Model/InlineResponse200132.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateResourcePoolGroup()`

```php
updateResourcePoolGroup($id, $inline_object207): \OpenAPI\Client\Model\InlineResponse200132
```

Update a Resource Pool Group

Use this command to update an existing resource pool Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ResourcePoolsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object207 = new \OpenAPI\Client\Model\InlineObject207(); // \OpenAPI\Client\Model\InlineObject207

try {
    $result = $apiInstance->updateResourcePoolGroup($id, $inline_object207);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ResourcePoolsApi->updateResourcePoolGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object207** | [**\OpenAPI\Client\Model\InlineObject207**](../Model/InlineObject207.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200132**](../Model/InlineResponse200132.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
