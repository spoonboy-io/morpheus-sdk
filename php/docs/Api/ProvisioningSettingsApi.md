# OpenAPI\Client\ProvisioningSettingsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**listProvisioningSettings()**](ProvisioningSettingsApi.md#listProvisioningSettings) | **GET** /api/provisioning-settings | Get Provisioning Settings
[**updateProvisioningSettings()**](ProvisioningSettingsApi.md#updateProvisioningSettings) | **PUT** /api/provisioning-settings | Update Provisioning Settings


## `listProvisioningSettings()`

```php
listProvisioningSettings(): \OpenAPI\Client\Model\InlineResponse200128
```

Get Provisioning Settings

This endpoint retrieves provisioning settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listProvisioningSettings();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningSettingsApi->listProvisioningSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200128**](../Model/InlineResponse200128.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateProvisioningSettings()`

```php
updateProvisioningSettings($inline_object204): \OpenAPI\Client\Model\Model200Success
```

Update Provisioning Settings

Update Provisioning Settings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object204 = new \OpenAPI\Client\Model\InlineObject204(); // \OpenAPI\Client\Model\InlineObject204

try {
    $result = $apiInstance->updateProvisioningSettings($inline_object204);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningSettingsApi->updateProvisioningSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object204** | [**\OpenAPI\Client\Model\InlineObject204**](../Model/InlineObject204.md)|  | [optional]

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
