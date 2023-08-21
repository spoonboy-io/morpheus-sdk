# OpenAPI\Client\MonitoringSettingsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMonitoringSettings()**](MonitoringSettingsApi.md#getMonitoringSettings) | **GET** /api/monitoring-settings | Get Monitoring Settings
[**updateMonitoringSettings()**](MonitoringSettingsApi.md#updateMonitoringSettings) | **PUT** /api/monitoring-settings | Update Monitoring Settings


## `getMonitoringSettings()`

```php
getMonitoringSettings(): \OpenAPI\Client\Model\InlineResponse20085
```

Get Monitoring Settings

This endpoint retrieves monitoring settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\MonitoringSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getMonitoringSettings();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MonitoringSettingsApi->getMonitoringSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20085**](../Model/InlineResponse20085.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateMonitoringSettings()`

```php
updateMonitoringSettings($inline_object141): Model200Success
```

Update Monitoring Settings

Update Monitoring Settings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\MonitoringSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object141 = new \OpenAPI\Client\Model\InlineObject141(); // \OpenAPI\Client\Model\InlineObject141

try {
    $result = $apiInstance->updateMonitoringSettings($inline_object141);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MonitoringSettingsApi->updateMonitoringSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object141** | [**\OpenAPI\Client\Model\InlineObject141**](../Model/InlineObject141.md)|  | [optional]

### Return type

[**Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
