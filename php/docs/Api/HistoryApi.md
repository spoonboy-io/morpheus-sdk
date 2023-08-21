# OpenAPI\Client\HistoryApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHistory()**](HistoryApi.md#getHistory) | **GET** /api/processes/{id} | Retrieves a Specific Process
[**getHistoryEvent()**](HistoryApi.md#getHistoryEvent) | **GET** /api/processes/events/{id} | Retrieves a Specific Process Event
[**listHistory()**](HistoryApi.md#listHistory) | **GET** /api/processes | Retrieves Process History


## `getHistory()`

```php
getHistory($id): \OpenAPI\Client\Model\InlineResponse20048
```

Retrieves a Specific Process

Retrieves a specific process.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HistoryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getHistory($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HistoryApi->getHistory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20048**](../Model/InlineResponse20048.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getHistoryEvent()`

```php
getHistoryEvent($id): \OpenAPI\Client\Model\InlineResponse20049
```

Retrieves a Specific Process Event

Retrieves a specific process event.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HistoryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getHistoryEvent($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HistoryApi->getHistoryEvent: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20049**](../Model/InlineResponse20049.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listHistory()`

```php
listHistory($instance_id, $container_id, $server_id, $zone_id, $app_id, $phrase): object
```

Retrieves Process History

Retrieves process history for objects

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HistoryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_id = 94; // int | The Instance ID for Filtering
$container_id = 135; // int | The Container ID for Filtering
$server_id = 97; // int | The Server ID for Filtering
$zone_id = 3; // int | The Zone ID for Filtering
$app_id = 5; // int | The App ID for Filtering
$phrase = 'phrase_example'; // string | Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error

try {
    $result = $apiInstance->listHistory($instance_id, $container_id, $server_id, $zone_id, $app_id, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HistoryApi->listHistory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **app_id** | **int**| The App ID for Filtering | [optional]
 **phrase** | **string**| Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error | [optional]

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
