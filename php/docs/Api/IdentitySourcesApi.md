# OpenAPI\Client\IdentitySourcesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIdentitySources()**](IdentitySourcesApi.md#addIdentitySources) | **POST** /api/user-sources | Creates an Identity Source
[**getIdentitySources()**](IdentitySourcesApi.md#getIdentitySources) | **GET** /api/user-sources/{id} | Retrieves a Specific Identity Source
[**listIdentitySources()**](IdentitySourcesApi.md#listIdentitySources) | **GET** /api/user-sources | Retrieves all Identity Sources
[**removeIdentitySources()**](IdentitySourcesApi.md#removeIdentitySources) | **DELETE** /api/user-sources/{id} | Deletes an Identity Source
[**updateIdentitySourceSubdomains()**](IdentitySourcesApi.md#updateIdentitySourceSubdomains) | **PUT** /api/user-sources/{id}/subdomain | Updates an Identity Source Subdomain
[**updateIdentitySources()**](IdentitySourcesApi.md#updateIdentitySources) | **PUT** /api/user-sources/{id} | Updates an Identity Source


## `addIdentitySources()`

```php
addIdentitySources($account_id, $user_source_create): object
```

Creates an Identity Source

Creates an identity source.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IdentitySourcesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
$user_source_create = new \OpenAPI\Client\Model\UserSourceCreate(); // \OpenAPI\Client\Model\UserSourceCreate

try {
    $result = $apiInstance->addIdentitySources($account_id, $user_source_create);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IdentitySourcesApi->addIdentitySources: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **user_source_create** | [**\OpenAPI\Client\Model\UserSourceCreate**](../Model/UserSourceCreate.md)|  | [optional]

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

## `getIdentitySources()`

```php
getIdentitySources($id): \OpenAPI\Client\Model\InlineResponse20051
```

Retrieves a Specific Identity Source

Retrieves a specific identity source.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IdentitySourcesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getIdentitySources($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IdentitySourcesApi->getIdentitySources: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20051**](../Model/InlineResponse20051.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listIdentitySources()`

```php
listIdentitySources($type, $max, $offset, $sort, $direction, $phrase, $name, $account_id): object
```

Retrieves all Identity Sources

Retrieves all identity sources.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IdentitySourcesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$type = 'type_example'; // string | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.

try {
    $result = $apiInstance->listIdentitySources($type, $max, $offset, $sort, $direction, $phrase, $name, $account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IdentitySourcesApi->listIdentitySources: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type** | **string**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings. | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

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

## `removeIdentitySources()`

```php
removeIdentitySources($id): \OpenAPI\Client\Model\Model200Success
```

Deletes an Identity Source

Deletes a specified identity source.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IdentitySourcesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeIdentitySources($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IdentitySourcesApi->removeIdentitySources: ', $e->getMessage(), PHP_EOL;
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

## `updateIdentitySourceSubdomains()`

```php
updateIdentitySourceSubdomains($id, $inline_object82): object
```

Updates an Identity Source Subdomain

Updates an identity source subdomain.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IdentitySourcesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object82 = new \OpenAPI\Client\Model\InlineObject82(); // \OpenAPI\Client\Model\InlineObject82

try {
    $result = $apiInstance->updateIdentitySourceSubdomains($id, $inline_object82);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IdentitySourcesApi->updateIdentitySourceSubdomains: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object82** | [**\OpenAPI\Client\Model\InlineObject82**](../Model/InlineObject82.md)|  | [optional]

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

## `updateIdentitySources()`

```php
updateIdentitySources($id, $user_source_create): object
```

Updates an Identity Source

Updates an identity source.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IdentitySourcesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$user_source_create = new \OpenAPI\Client\Model\UserSourceCreate(); // \OpenAPI\Client\Model\UserSourceCreate

try {
    $result = $apiInstance->updateIdentitySources($id, $user_source_create);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IdentitySourcesApi->updateIdentitySources: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **user_source_create** | [**\OpenAPI\Client\Model\UserSourceCreate**](../Model/UserSourceCreate.md)|  | [optional]

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
