# OpenAPI\Client\DashboardApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard()**](DashboardApi.md#dashboard) | **GET** /api/dashboard | Overview of Dashboard information


## `dashboard()`

```php
dashboard(): \OpenAPI\Client\Model\Dashboard
```

Overview of Dashboard information

This endpoint can be used to view dashboard information about the remote Morpheus appliance. This is an overview and summary of data available to the user that can be used to render a dashboard.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DashboardApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->dashboard();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DashboardApi->dashboard: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\Dashboard**](../Model/Dashboard.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
