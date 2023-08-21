# OpenAPI\Client\KeyPairsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addKeyPairs()**](KeyPairsApi.md#addKeyPairs) | **POST** /api/key-pairs | Creates a Key Pair
[**generateKeyPairs()**](KeyPairsApi.md#generateKeyPairs) | **POST** /api/key-pairs/generate | Generates a Key Pair
[**getKeyPairs()**](KeyPairsApi.md#getKeyPairs) | **GET** /api/key-pairs/{id} | Retrieves a Specific Key Pair
[**removeKeyPairs()**](KeyPairsApi.md#removeKeyPairs) | **DELETE** /api/key-pairs/{id} | Deletes a Key Pair


## `addKeyPairs()`

```php
addKeyPairs($inline_object105): object
```

Creates a Key Pair

Creates a Key Pair.  Private keys are typically optional.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\KeyPairsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object105 = new \OpenAPI\Client\Model\InlineObject105(); // \OpenAPI\Client\Model\InlineObject105

try {
    $result = $apiInstance->addKeyPairs($inline_object105);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling KeyPairsApi->addKeyPairs: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object105** | [**\OpenAPI\Client\Model\InlineObject105**](../Model/InlineObject105.md)|  | [optional]

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

## `generateKeyPairs()`

```php
generateKeyPairs($inline_object106): \OpenAPI\Client\Model\InlineResponse20066
```

Generates a Key Pair

Generates a Key Pair.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\KeyPairsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object106 = new \OpenAPI\Client\Model\InlineObject106(); // \OpenAPI\Client\Model\InlineObject106

try {
    $result = $apiInstance->generateKeyPairs($inline_object106);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling KeyPairsApi->generateKeyPairs: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object106** | [**\OpenAPI\Client\Model\InlineObject106**](../Model/InlineObject106.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20066**](../Model/InlineResponse20066.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getKeyPairs()`

```php
getKeyPairs($id): \OpenAPI\Client\Model\InlineResponse20067
```

Retrieves a Specific Key Pair

Retrieves a specific key pair.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\KeyPairsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getKeyPairs($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling KeyPairsApi->getKeyPairs: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20067**](../Model/InlineResponse20067.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `removeKeyPairs()`

```php
removeKeyPairs($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Key Pair

Deletes a specified key pair.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\KeyPairsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeKeyPairs($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling KeyPairsApi->removeKeyPairs: ', $e->getMessage(), PHP_EOL;
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
