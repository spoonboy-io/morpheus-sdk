# OpenAPI\Client\LicenseApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getLicense()**](LicenseApi.md#getLicense) | **GET** /api/license | Get license
[**installLicense()**](LicenseApi.md#installLicense) | **POST** /api/license | Install license key
[**testLicense()**](LicenseApi.md#testLicense) | **POST** /api/license/test | Test license key
[**uninstallLicense()**](LicenseApi.md#uninstallLicense) | **DELETE** /api/license | Uninstall license key


## `getLicense()`

```php
getLicense(): \OpenAPI\Client\Model\License
```

Get license

Get details about the license that is currently installed on the appliance. This returns the license details, but not the key itself. Your license key will never be returned and should always be kept secret.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LicenseApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getLicense();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LicenseApi->getLicense: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\License**](../Model/License.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `installLicense()`

```php
installLicense($inline_object125): \OpenAPI\Client\Model\License
```

Install license key

Install a new license key. This will potentially change the enabled features and capabilities of your Morpheus appliance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LicenseApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object125 = new \OpenAPI\Client\Model\InlineObject125(); // \OpenAPI\Client\Model\InlineObject125

try {
    $result = $apiInstance->installLicense($inline_object125);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LicenseApi->installLicense: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object125** | [**\OpenAPI\Client\Model\InlineObject125**](../Model/InlineObject125.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\License**](../Model/License.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `testLicense()`

```php
testLicense($inline_object126): \OpenAPI\Client\Model\License
```

Test license key

This endpoint can be used to decode a license to see if it is valid and inspect the license settings, such as who it belongs to and the enabled features. This is only a test, it does not install the key, or make any changes to your appliance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LicenseApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object126 = new \OpenAPI\Client\Model\InlineObject126(); // \OpenAPI\Client\Model\InlineObject126

try {
    $result = $apiInstance->testLicense($inline_object126);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LicenseApi->testLicense: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object126** | [**\OpenAPI\Client\Model\InlineObject126**](../Model/InlineObject126.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\License**](../Model/License.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `uninstallLicense()`

```php
uninstallLicense(): \OpenAPI\Client\Model\Model200Success
```

Uninstall license key

Uninstall your appliance license, leaving the appliance with no license installed. This will change the enabled features and capabilities of your Morpheus appliance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LicenseApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->uninstallLicense();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LicenseApi->uninstallLicense: ', $e->getMessage(), PHP_EOL;
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
