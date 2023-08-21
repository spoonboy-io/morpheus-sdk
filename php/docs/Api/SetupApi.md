# OpenAPI\Client\SetupApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**setup()**](SetupApi.md#setup) | **POST** /api/setup | Setup appliance


## `setup()`

```php
setup($unknown_base_type): \OpenAPI\Client\Model\Setup
```

Setup appliance

Initialize a freshly installed appliance to create the master tenant and System Admin user.  Authorization is not required.  This operation can only be executed successfully once.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\SetupApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$unknown_base_type = {"applianceName":"The Matrix","accountName":"Meta Cortex Corporation","firstName":"Thomas","lastName":"Anderson","email":"tanderson@mccorp.com","username":"tanderson","password":"QnW}cg}8}<~:P9YU"}; // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->setup($unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SetupApi->setup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Setup**](../Model/Setup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
