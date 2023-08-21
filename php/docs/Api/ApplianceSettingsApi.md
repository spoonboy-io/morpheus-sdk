# OpenAPI\Client\ApplianceSettingsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**listApplianceSettings()**](ApplianceSettingsApi.md#listApplianceSettings) | **GET** /api/appliance-settings | Get Appliance Settings
[**reindex()**](ApplianceSettingsApi.md#reindex) | **POST** /api/appliance-settings/reindex | Reindex Search
[**setApplianceSettingsMaintenanceMode()**](ApplianceSettingsApi.md#setApplianceSettingsMaintenanceMode) | **POST** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**updateApplianceSettings()**](ApplianceSettingsApi.md#updateApplianceSettings) | **PUT** /api/appliance-settings | Update Appliance Settings


## `listApplianceSettings()`

```php
listApplianceSettings(): \OpenAPI\Client\Model\InlineResponse200
```

Get Appliance Settings

This endpoint retrieves appliance settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ApplianceSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listApplianceSettings();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ApplianceSettingsApi->listApplianceSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200**](../Model/InlineResponse200.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `reindex()`

```php
reindex(): \OpenAPI\Client\Model\Model200Success
```

Reindex Search

Reindex all searchable data

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ApplianceSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->reindex();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ApplianceSettingsApi->reindex: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

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

## `setApplianceSettingsMaintenanceMode()`

```php
setApplianceSettingsMaintenanceMode($enabled): \OpenAPI\Client\Model\Model200Success
```

Toggle Maintenance Mode

This endpoint allows toggling the appliance maintenance mode.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ApplianceSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$enabled = True; // bool | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa.

try {
    $result = $apiInstance->setApplianceSettingsMaintenanceMode($enabled);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ApplianceSettingsApi->setApplianceSettingsMaintenanceMode: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool**| Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | [optional]

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

## `updateApplianceSettings()`

```php
updateApplianceSettings($inline_object2): \OpenAPI\Client\Model\Model200Success
```

Update Appliance Settings

Update Appliance Settings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ApplianceSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object2 = new \OpenAPI\Client\Model\InlineObject2(); // \OpenAPI\Client\Model\InlineObject2

try {
    $result = $apiInstance->updateApplianceSettings($inline_object2);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ApplianceSettingsApi->updateApplianceSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object2** | [**\OpenAPI\Client\Model\InlineObject2**](../Model/InlineObject2.md)|  | [optional]

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
