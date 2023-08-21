# OpenAPI\Client\SecurityScansApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getSecurityScans()**](SecurityScansApi.md#getSecurityScans) | **GET** /api/security-scans/{id} | Retrieves a Specific Security Scan
[**listSecurityScans()**](SecurityScansApi.md#listSecurityScans) | **GET** /api/security-scans | Retrieves all Security Scans


## `getSecurityScans()`

```php
getSecurityScans($id, $results): object
```

Retrieves a Specific Security Scan

Retrieves a specific security scan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityScansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$results = true; // bool | Include the `results` object in the response under the security scan. This is a potentially very large object containing the raw results of the scan.

try {
    $result = $apiInstance->getSecurityScans($id, $results);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityScansApi->getSecurityScans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **results** | **bool**| Include the &#x60;results&#x60; object in the response under the security scan. This is a potentially very large object containing the raw results of the scan. | [optional] [default to false]

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

## `listSecurityScans()`

```php
listSecurityScans($max, $offset, $sort, $direction, $phrase, $security_package_id, $server_id, $results): object
```

Retrieves all Security Scans

Retrieves all security scans.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityScansApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'scanDate'; // string | Sort order, the name of the property to sort by
$direction = desc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description of security package
$security_package_id = 56; // int | Filter results by security package id(s). This parameter can be passed multiple times to match more than one id.
$server_id = 97; // int | The Server ID for Filtering
$results = true; // bool | Include the `results` object in the response under each security scan. This is a potentially very large object containing the raw results of the scan.

try {
    $result = $apiInstance->listSecurityScans($max, $offset, $sort, $direction, $phrase, $security_package_id, $server_id, $results);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityScansApi->listSecurityScans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;scanDate&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;desc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description of security package | [optional]
 **security_package_id** | **int**| Filter results by security package id(s). This parameter can be passed multiple times to match more than one id. | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **results** | **bool**| Include the &#x60;results&#x60; object in the response under each security scan. This is a potentially very large object containing the raw results of the scan. | [optional] [default to false]

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
