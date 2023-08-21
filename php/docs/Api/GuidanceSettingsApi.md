# OpenAPI\Client\GuidanceSettingsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getGuidanceSettings()**](GuidanceSettingsApi.md#getGuidanceSettings) | **GET** /api/guidance-settings | Get Guidance Settings
[**updateGuidanceSettings()**](GuidanceSettingsApi.md#updateGuidanceSettings) | **PUT** /api/guidance-settings | Update Guidance Settings


## `getGuidanceSettings()`

```php
getGuidanceSettings(): \OpenAPI\Client\Model\InlineResponse20047
```

Get Guidance Settings

This endpoint retrieves guidance settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getGuidanceSettings();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceSettingsApi->getGuidanceSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20047**](../Model/InlineResponse20047.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateGuidanceSettings()`

```php
updateGuidanceSettings($inline_object79): Model200Success
```

Update Guidance Settings

Update Guidance Settings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object79 = new \OpenAPI\Client\Model\InlineObject79(); // \OpenAPI\Client\Model\InlineObject79

try {
    $result = $apiInstance->updateGuidanceSettings($inline_object79);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceSettingsApi->updateGuidanceSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object79** | [**\OpenAPI\Client\Model\InlineObject79**](../Model/InlineObject79.md)|  | [optional]

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
