# OpenAPI\Client\LogSettingsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addLogSettingsSyslogRules()**](LogSettingsApi.md#addLogSettingsSyslogRules) | **POST** /api/log-settings/syslog-rules | Create a New Syslog Rule
[**deleteLogSettingsSyslogRules()**](LogSettingsApi.md#deleteLogSettingsSyslogRules) | **DELETE** /api/log-settings/syslog-rules/{id} | Delete a Specific Syslog Rule
[**listLogSettings()**](LogSettingsApi.md#listLogSettings) | **GET** /api/log-settings | List All Log Settings
[**updateLogSettings()**](LogSettingsApi.md#updateLogSettings) | **PUT** /api/log-settings | Update Log Settings


## `addLogSettingsSyslogRules()`

```php
addLogSettingsSyslogRules($inline_object140): \OpenAPI\Client\Model\Model200Success
```

Create a New Syslog Rule

Creates a new syslog rule. This command will also update existing syslog rule by specified name if already exists

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LogSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object140 = new \OpenAPI\Client\Model\InlineObject140(); // \OpenAPI\Client\Model\InlineObject140

try {
    $result = $apiInstance->addLogSettingsSyslogRules($inline_object140);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LogSettingsApi->addLogSettingsSyslogRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object140** | [**\OpenAPI\Client\Model\InlineObject140**](../Model/InlineObject140.md)|  | [optional]

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

## `deleteLogSettingsSyslogRules()`

```php
deleteLogSettingsSyslogRules($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Specific Syslog Rule

Will delete the syslog rule matching the specified name.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LogSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLogSettingsSyslogRules($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LogSettingsApi->deleteLogSettingsSyslogRules: ', $e->getMessage(), PHP_EOL;
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

## `listLogSettings()`

```php
listLogSettings(): \OpenAPI\Client\Model\InlineResponse20084
```

List All Log Settings

This endpoint retrieves log settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LogSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listLogSettings();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LogSettingsApi->listLogSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20084**](../Model/InlineResponse20084.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateLogSettings()`

```php
updateLogSettings($inline_object139): Model200Success
```

Update Log Settings

Update Log Settings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LogSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object139 = new \OpenAPI\Client\Model\InlineObject139(); // \OpenAPI\Client\Model\InlineObject139

try {
    $result = $apiInstance->updateLogSettings($inline_object139);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LogSettingsApi->updateLogSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object139** | [**\OpenAPI\Client\Model\InlineObject139**](../Model/InlineObject139.md)|  | [optional]

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
