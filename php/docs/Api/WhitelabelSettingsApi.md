# OpenAPI\Client\WhitelabelSettingsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getWhitelabelImage()**](WhitelabelSettingsApi.md#getWhitelabelImage) | **GET** /api/whitelabel-settings/images/{imageType} | Download Image
[**listWhitelabelSettings()**](WhitelabelSettingsApi.md#listWhitelabelSettings) | **GET** /api/whitelabel-settings | Get Whitelabel Settings
[**removeWhitelabelImage()**](WhitelabelSettingsApi.md#removeWhitelabelImage) | **DELETE** /api/whitelabel-settings/images/{imageType} | Reset Image
[**updateWhitelabelImages()**](WhitelabelSettingsApi.md#updateWhitelabelImages) | **POST** /api/whitelabel-settings/images | Update Images
[**updateWhitelabelSettings()**](WhitelabelSettingsApi.md#updateWhitelabelSettings) | **PUT** /api/whitelabel-settings | Update Whitelabel Settings


## `getWhitelabelImage()`

```php
getWhitelabelImage($image_type): \SplFileObject
```

Download Image

Downloads the specified image.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WhitelabelSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$image_type = headerLogo; // string | Valid image types

try {
    $result = $apiInstance->getWhitelabelImage($image_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WhitelabelSettingsApi->getWhitelabelImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_type** | **string**| Valid image types |

### Return type

[**\SplFileObject**](../Model/\SplFileObject.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `image/ico`, `image/jpeg`, `image/png`, `image/svg+xml`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listWhitelabelSettings()`

```php
listWhitelabelSettings(): \OpenAPI\Client\Model\InlineResponse200166
```

Get Whitelabel Settings

This endpoint retrieves whitelabel settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WhitelabelSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listWhitelabelSettings();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WhitelabelSettingsApi->listWhitelabelSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200166**](../Model/InlineResponse200166.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `removeWhitelabelImage()`

```php
removeWhitelabelImage($image_type): \OpenAPI\Client\Model\Model200Success
```

Reset Image

Resets the specified image to the Morpheus default.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WhitelabelSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$image_type = headerLogo; // string | Valid image types

try {
    $result = $apiInstance->removeWhitelabelImage($image_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WhitelabelSettingsApi->removeWhitelabelImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_type** | **string**| Valid image types |

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

## `updateWhitelabelImages()`

```php
updateWhitelabelImages($header_logo_file, $reset_header_logo, $footer_logo_file, $reset_footer_logo, $login_logo_file, $reset_login_logo, $favicon_file, $reset_favicon_logo): \OpenAPI\Client\Model\Model200Success
```

Update Images

Uploads whitelabel images. Expects multipart form data as the request format, not JSON.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WhitelabelSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$header_logo_file = "/path/to/file.txt"; // \SplFileObject | Header logo image file, valid image types `png|jpg|svg`
$reset_header_logo = True; // bool | Resets header logo to default
$footer_logo_file = "/path/to/file.txt"; // \SplFileObject | Footer logo image file, valid image types `png|jpg|svg`
$reset_footer_logo = True; // bool | Resets footer logo to default
$login_logo_file = "/path/to/file.txt"; // \SplFileObject | Login logo image file, valid image types `png|jpg|svg`
$reset_login_logo = True; // bool | Resets login logo to default
$favicon_file = "/path/to/file.txt"; // \SplFileObject | Favicon image file, valid image type ico
$reset_favicon_logo = True; // bool | Resets favicon logo to default

try {
    $result = $apiInstance->updateWhitelabelImages($header_logo_file, $reset_header_logo, $footer_logo_file, $reset_footer_logo, $login_logo_file, $reset_login_logo, $favicon_file, $reset_favicon_logo);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WhitelabelSettingsApi->updateWhitelabelImages: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **header_logo_file** | **\SplFileObject****\SplFileObject**| Header logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **reset_header_logo** | **bool**| Resets header logo to default | [optional]
 **footer_logo_file** | **\SplFileObject****\SplFileObject**| Footer logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **reset_footer_logo** | **bool**| Resets footer logo to default | [optional]
 **login_logo_file** | **\SplFileObject****\SplFileObject**| Login logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **reset_login_logo** | **bool**| Resets login logo to default | [optional]
 **favicon_file** | **\SplFileObject****\SplFileObject**| Favicon image file, valid image type ico | [optional]
 **reset_favicon_logo** | **bool**| Resets favicon logo to default | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateWhitelabelSettings()`

```php
updateWhitelabelSettings($inline_object265): \OpenAPI\Client\Model\Model200Success
```

Update Whitelabel Settings

Update Whitelabel Settings

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WhitelabelSettingsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object265 = new \OpenAPI\Client\Model\InlineObject265(); // \OpenAPI\Client\Model\InlineObject265

try {
    $result = $apiInstance->updateWhitelabelSettings($inline_object265);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WhitelabelSettingsApi->updateWhitelabelSettings: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object265** | [**\OpenAPI\Client\Model\InlineObject265**](../Model/InlineObject265.md)|  | [optional]

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
