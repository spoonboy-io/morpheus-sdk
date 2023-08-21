# OpenAPI\Client\SSLCertificatesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCertificate()**](SSLCertificatesApi.md#addCertificate) | **POST** /api/certificates | Create a Certificate
[**deleteCertificate()**](SSLCertificatesApi.md#deleteCertificate) | **DELETE** /api/certificates/{id} | Delete a Certificate
[**getCertificate()**](SSLCertificatesApi.md#getCertificate) | **GET** /api/certificates/{id} | Get a Specific Certificate
[**listCertificates()**](SSLCertificatesApi.md#listCertificates) | **GET** /api/certificates | Get All SSL Certificates
[**updateCertificate()**](SSLCertificatesApi.md#updateCertificate) | **PUT** /api/certificates/{id} | Update a Certificate


## `addCertificate()`

```php
addCertificate($inline_object230): object
```

Create a Certificate

Create a Certificate

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SSLCertificatesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object230 = new \OpenAPI\Client\Model\InlineObject230(); // \OpenAPI\Client\Model\InlineObject230

try {
    $result = $apiInstance->addCertificate($inline_object230);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SSLCertificatesApi->addCertificate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object230** | [**\OpenAPI\Client\Model\InlineObject230**](../Model/InlineObject230.md)|  | [optional]

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

## `deleteCertificate()`

```php
deleteCertificate($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Certificate

Will delete a certificate from the system and make it no longer usable.  If a certificate is actively in use, a delete will fail.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SSLCertificatesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteCertificate($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SSLCertificatesApi->deleteCertificate: ', $e->getMessage(), PHP_EOL;
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

## `getCertificate()`

```php
getCertificate($id): \OpenAPI\Client\Model\InlineResponse200144
```

Get a Specific Certificate

This endpoint retrieves a specific certificate.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SSLCertificatesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCertificate($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SSLCertificatesApi->getCertificate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200144**](../Model/InlineResponse200144.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listCertificates()`

```php
listCertificates($name): \OpenAPI\Client\Model\InlineResponse200143
```

Get All SSL Certificates

This endpoint retrieves all SSL certificates associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SSLCertificatesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listCertificates($name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SSLCertificatesApi->listCertificates: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200143**](../Model/InlineResponse200143.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateCertificate()`

```php
updateCertificate($id, $inline_object231): \OpenAPI\Client\Model\InlineResponse200144
```

Update a Certificate

Update a Certificate.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SSLCertificatesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object231 = new \OpenAPI\Client\Model\InlineObject231(); // \OpenAPI\Client\Model\InlineObject231

try {
    $result = $apiInstance->updateCertificate($id, $inline_object231);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SSLCertificatesApi->updateCertificate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object231** | [**\OpenAPI\Client\Model\InlineObject231**](../Model/InlineObject231.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200144**](../Model/InlineResponse200144.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
