# OpenAPI\Client\OptionsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getOptionSourceData()**](OptionsApi.md#getOptionSourceData) | **GET** /api/options/{optionSource} | Get Option Source Data
[**listCodeRepositories()**](OptionsApi.md#listCodeRepositories) | **GET** /api/options/codeRepositories | Retrieves a list of Code/GIT Repositories
[**listOptionNetworkOptions()**](OptionsApi.md#listOptionNetworkOptions) | **GET** /api/options/zoneNetworkOptions | Retrieves network options by zone/cloud
[**listOptionValues()**](OptionsApi.md#listOptionValues) | **GET** /api/options/list | Retrieves input option values


## `getOptionSourceData()`

```php
getOptionSourceData($option_source): object
```

Get Option Source Data

Returns a list of name/value pairs for option-type models. Some option-types depend on input data for proper representation. This typically includes zoneId or siteId for the item being provisioned as request parameters or sometimes previous option type parameters. Each option returned has a `value`, which is often the `id`, but may be a `code` or other attribute.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\OptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$option_source = keyPairs; // string | `optionSource` to be listed

try {
    $result = $apiInstance->getOptionSourceData($option_source);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling OptionsApi->getOptionSourceData: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **option_source** | **string**| &#x60;optionSource&#x60; to be listed |

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

## `listCodeRepositories()`

```php
listCodeRepositories($integration_id): Model200Success
```

Retrieves a list of Code/GIT Repositories

Retrieves a list of Code/GIT Repositories

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\OptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$integration_id = 33; // int | Filter by an integration Id.

try {
    $result = $apiInstance->listCodeRepositories($integration_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling OptionsApi->listCodeRepositories: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integration_id** | **int**| Filter by an integration Id. | [optional]

### Return type

[**Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listOptionNetworkOptions()`

```php
listOptionNetworkOptions($zone_id, $provision_type_id): ZoneNetworkOptions
```

Retrieves network options by zone/cloud

This endpoint can be used to see which network options are available for a given cloud (zoneId) and provision type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\OptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 3; // int | The Zone ID for Filtering
$provision_type_id = 22; // int | Provision type filter, restricts query to only load service plans of specified provision type

try {
    $result = $apiInstance->listOptionNetworkOptions($zone_id, $provision_type_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling OptionsApi->listOptionNetworkOptions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **provision_type_id** | **int**| Provision type filter, restricts query to only load service plans of specified provision type | [optional]

### Return type

[**ZoneNetworkOptions**](../Model/ZoneNetworkOptions.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listOptionValues()`

```php
listOptionValues($option_type_id, $config): object
```

Retrieves input option values

Retrieves all input option values.  Can be used with parameters to supply dependent input values.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\OptionsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$option_type_id = 56; // int | Input or Option Type ID
$config = {"config.fieldName1":"value1"}; // object | Input parameters are required if the input is dependent on them.  Fields must be prefixed with `config.`

try {
    $result = $apiInstance->listOptionValues($option_type_id, $config);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling OptionsApi->listOptionValues: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **option_type_id** | **int**| Input or Option Type ID |
 **config** | [**object**](../Model/.md)| Input parameters are required if the input is dependent on them.  Fields must be prefixed with &#x60;config.&#x60; | [optional]

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
