# OpenAPI\Client\PriceSetsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPriceSets()**](PriceSetsApi.md#addPriceSets) | **POST** /api/price-sets | Creates a Price Set
[**deactivatePriceSets()**](PriceSetsApi.md#deactivatePriceSets) | **PUT** /api/price-sets/{id}/deactivate | Deactivates a Price Set
[**getPriceSets()**](PriceSetsApi.md#getPriceSets) | **GET** /api/price-sets/{id} | Retrieves a Specific Price Set
[**listPriceSets()**](PriceSetsApi.md#listPriceSets) | **GET** /api/price-sets | Retrieves all Price Sets
[**updatePriceSets()**](PriceSetsApi.md#updatePriceSets) | **PUT** /api/price-sets/{id} | Updates a Price Set


## `addPriceSets()`

```php
addPriceSets($inline_object198): object
```

Creates a Price Set

Creates a price set.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PriceSetsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object198 = new \OpenAPI\Client\Model\InlineObject198(); // \OpenAPI\Client\Model\InlineObject198

try {
    $result = $apiInstance->addPriceSets($inline_object198);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PriceSetsApi->addPriceSets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object198** | [**\OpenAPI\Client\Model\InlineObject198**](../Model/InlineObject198.md)|  | [optional]

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

## `deactivatePriceSets()`

```php
deactivatePriceSets($id): \OpenAPI\Client\Model\Model200Success
```

Deactivates a Price Set

Deactivates a price set. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PriceSetsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deactivatePriceSets($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PriceSetsApi->deactivatePriceSets: ', $e->getMessage(), PHP_EOL;
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

## `getPriceSets()`

```php
getPriceSets($id): \OpenAPI\Client\Model\InlineResponse200123
```

Retrieves a Specific Price Set

Retrieves a specific price set.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PriceSetsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPriceSets($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PriceSetsApi->getPriceSets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200123**](../Model/InlineResponse200123.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listPriceSets()`

```php
listPriceSets($max, $offset, $sort, $direction, $phrase, $name, $include_inactive, $type): object
```

Retrieves all Price Sets

Retrieves all price sets.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PriceSetsApi(
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
$include_inactive = true; // bool | If true, include inactive prices in the results
$type = 'type_example'; // string | Filter by type code

try {
    $result = $apiInstance->listPriceSets($max, $offset, $sort, $direction, $phrase, $name, $include_inactive, $type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PriceSetsApi->listPriceSets: ', $e->getMessage(), PHP_EOL;
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
 **include_inactive** | **bool**| If true, include inactive prices in the results | [optional]
 **type** | **string**| Filter by type code | [optional]

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

## `updatePriceSets()`

```php
updatePriceSets($id, $inline_object199): object
```

Updates a Price Set

Updates a price set.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PriceSetsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object199 = new \OpenAPI\Client\Model\InlineObject199(); // \OpenAPI\Client\Model\InlineObject199

try {
    $result = $apiInstance->updatePriceSets($id, $inline_object199);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PriceSetsApi->updatePriceSets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object199** | [**\OpenAPI\Client\Model\InlineObject199**](../Model/InlineObject199.md)|  | [optional]

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
