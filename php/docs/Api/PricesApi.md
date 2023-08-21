# OpenAPI\Client\PricesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPrices()**](PricesApi.md#addPrices) | **POST** /api/prices | Creates a Price
[**deactivatePrices()**](PricesApi.md#deactivatePrices) | **PUT** /api/prices/{id}/deactivate | Deactivates a Price
[**getPrices()**](PricesApi.md#getPrices) | **GET** /api/prices/{id} | Retrieves a Specific Price
[**listPrices()**](PricesApi.md#listPrices) | **GET** /api/prices | Retrieves all Prices
[**updatePrices()**](PricesApi.md#updatePrices) | **PUT** /api/prices/{id} | Updates a Price


## `addPrices()`

```php
addPrices($inline_object200): object
```

Creates a Price

Creates a price.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PricesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object200 = new \OpenAPI\Client\Model\InlineObject200(); // \OpenAPI\Client\Model\InlineObject200

try {
    $result = $apiInstance->addPrices($inline_object200);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PricesApi->addPrices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object200** | [**\OpenAPI\Client\Model\InlineObject200**](../Model/InlineObject200.md)|  | [optional]

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

## `deactivatePrices()`

```php
deactivatePrices($id): \OpenAPI\Client\Model\Model200Success
```

Deactivates a Price

Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PricesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deactivatePrices($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PricesApi->deactivatePrices: ', $e->getMessage(), PHP_EOL;
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

## `getPrices()`

```php
getPrices($id): \OpenAPI\Client\Model\InlineResponse200124
```

Retrieves a Specific Price

Retrieves a specific price.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PricesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPrices($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PricesApi->getPrices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200124**](../Model/InlineResponse200124.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listPrices()`

```php
listPrices($max, $offset, $sort, $direction, $phrase, $name, $include_inactive, $price_type, $platform, $currency, $type): object
```

Retrieves all Prices

Retrieves all prices.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PricesApi(
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
$price_type = 'price_type_example'; // string | Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software
$platform = linux; // string | Restricts query to only load only prices with specified platform
$currency = USD; // string | Restricts query to only load only prices with specified currency
$type = 'type_example'; // string | Filter by type code

try {
    $result = $apiInstance->listPrices($max, $offset, $sort, $direction, $phrase, $name, $include_inactive, $price_type, $platform, $currency, $type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PricesApi->listPrices: ', $e->getMessage(), PHP_EOL;
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
 **price_type** | **string**| Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software | [optional]
 **platform** | **string**| Restricts query to only load only prices with specified platform | [optional]
 **currency** | **string**| Restricts query to only load only prices with specified currency | [optional]
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

## `updatePrices()`

```php
updatePrices($id, $inline_object201): object
```

Updates a Price

Updates a price.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PricesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object201 = new \OpenAPI\Client\Model\InlineObject201(); // \OpenAPI\Client\Model\InlineObject201

try {
    $result = $apiInstance->updatePrices($id, $inline_object201);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PricesApi->updatePrices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object201** | [**\OpenAPI\Client\Model\InlineObject201**](../Model/InlineObject201.md)|  | [optional]

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
