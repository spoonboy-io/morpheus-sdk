# OpenAPI\Client\ServicePlansApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**activateServicePlans()**](ServicePlansApi.md#activateServicePlans) | **PUT** /api/service-plans/{id}/activate | Activates a Service Plan
[**addServicePlans()**](ServicePlansApi.md#addServicePlans) | **POST** /api/service-plans | Creates a Service Plan
[**deactivateServicePlans()**](ServicePlansApi.md#deactivateServicePlans) | **PUT** /api/service-plans/{id}/deactivate | Deactivates a Service Plan
[**getServicePlans()**](ServicePlansApi.md#getServicePlans) | **GET** /api/service-plans/{id} | Retrieves a Specific Service Plan
[**listServicePlans()**](ServicePlansApi.md#listServicePlans) | **GET** /api/service-plans | Retrieves all Service Plans
[**removeServicePlans()**](ServicePlansApi.md#removeServicePlans) | **DELETE** /api/service-plans/{id} | Deletes a Service Plan
[**updateServicePlans()**](ServicePlansApi.md#updateServicePlans) | **PUT** /api/service-plans/{id} | Updates a Service Plan


## `activateServicePlans()`

```php
activateServicePlans($id): \OpenAPI\Client\Model\Model200Success
```

Activates a Service Plan

Activates a service plan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->activateServicePlans($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->activateServicePlans: ', $e->getMessage(), PHP_EOL;
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

## `addServicePlans()`

```php
addServicePlans($inline_object228): object
```

Creates a Service Plan

Creates a service plan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object228 = new \OpenAPI\Client\Model\InlineObject228(); // \OpenAPI\Client\Model\InlineObject228

try {
    $result = $apiInstance->addServicePlans($inline_object228);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->addServicePlans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object228** | [**\OpenAPI\Client\Model\InlineObject228**](../Model/InlineObject228.md)|  | [optional]

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

## `deactivateServicePlans()`

```php
deactivateServicePlans($id): \OpenAPI\Client\Model\Model200Success
```

Deactivates a Service Plan

Deactivates a service plan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deactivateServicePlans($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->deactivateServicePlans: ', $e->getMessage(), PHP_EOL;
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

## `getServicePlans()`

```php
getServicePlans($id): \OpenAPI\Client\Model\InlineResponse200142
```

Retrieves a Specific Service Plan

Retrieves a specific service plan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getServicePlans($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->getServicePlans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200142**](../Model/InlineResponse200142.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listServicePlans()`

```php
listServicePlans($max, $offset, $sort, $direction, $phrase, $name, $include_zones, $provision_type_id, $include_inactive): object
```

Retrieves all Service Plans

Retrieves all service plans.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
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
$include_zones = false; // bool | Indicates including zones in the service plan response payload
$provision_type_id = 22; // int | Provision type filter, restricts query to only load service plans of specified provision type
$include_inactive = true; // bool | If true, include inactive prices in the results

try {
    $result = $apiInstance->listServicePlans($max, $offset, $sort, $direction, $phrase, $name, $include_zones, $provision_type_id, $include_inactive);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->listServicePlans: ', $e->getMessage(), PHP_EOL;
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
 **include_zones** | **bool**| Indicates including zones in the service plan response payload | [optional] [default to false]
 **provision_type_id** | **int**| Provision type filter, restricts query to only load service plans of specified provision type | [optional]
 **include_inactive** | **bool**| If true, include inactive prices in the results | [optional]

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

## `removeServicePlans()`

```php
removeServicePlans($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Service Plan

Deletes a specified service plan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeServicePlans($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->removeServicePlans: ', $e->getMessage(), PHP_EOL;
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

## `updateServicePlans()`

```php
updateServicePlans($id, $inline_object229): object
```

Updates a Service Plan

Updates a service plan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServicePlansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object229 = new \OpenAPI\Client\Model\InlineObject229(); // \OpenAPI\Client\Model\InlineObject229

try {
    $result = $apiInstance->updateServicePlans($id, $inline_object229);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServicePlansApi->updateServicePlans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object229** | [**\OpenAPI\Client\Model\InlineObject229**](../Model/InlineObject229.md)|  | [optional]

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
