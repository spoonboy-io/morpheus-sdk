# OpenAPI\Client\ProvisioningLicensesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addProvisioningLicense()**](ProvisioningLicensesApi.md#addProvisioningLicense) | **POST** /api/provisioning-licenses | Create a License
[**getProvisioningLicense()**](ProvisioningLicensesApi.md#getProvisioningLicense) | **GET** /api/provisioning-licenses/{id} | Get a Specific License
[**getProvisioningLicenseReservations()**](ProvisioningLicensesApi.md#getProvisioningLicenseReservations) | **GET** /api/provisioning-licenses/{id}/reservations | Get Reservations for Specific License
[**listProvisioningLicenses()**](ProvisioningLicensesApi.md#listProvisioningLicenses) | **GET** /api/provisioning-licenses | Get All Licenses
[**removeProvisioningLicense()**](ProvisioningLicensesApi.md#removeProvisioningLicense) | **DELETE** /api/provisioning-licenses/{id} | Delete a License
[**updateProvisioningLicense()**](ProvisioningLicensesApi.md#updateProvisioningLicense) | **PUT** /api/provisioning-licenses/{id} | Update a License


## `addProvisioningLicense()`

```php
addProvisioningLicense($inline_object202): Model200Success
```

Create a License

Use this command to create a new license.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningLicensesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object202 = new \OpenAPI\Client\Model\InlineObject202(); // \OpenAPI\Client\Model\InlineObject202

try {
    $result = $apiInstance->addProvisioningLicense($inline_object202);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningLicensesApi->addProvisioningLicense: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object202** | [**\OpenAPI\Client\Model\InlineObject202**](../Model/InlineObject202.md)|  | [optional]

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

## `getProvisioningLicense()`

```php
getProvisioningLicense($id): \OpenAPI\Client\Model\InlineResponse200126
```

Get a Specific License

This endpoint retrieves a specific license.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningLicensesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getProvisioningLicense($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningLicensesApi->getProvisioningLicense: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200126**](../Model/InlineResponse200126.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getProvisioningLicenseReservations()`

```php
getProvisioningLicenseReservations($id): \OpenAPI\Client\Model\InlineResponse200127
```

Get Reservations for Specific License

This endpoint retrieves all reservations for a specific license. Each time a license is applied to a new server, a reservation is created, reducing the available copies for the license.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningLicensesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getProvisioningLicenseReservations($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningLicensesApi->getProvisioningLicenseReservations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200127**](../Model/InlineResponse200127.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listProvisioningLicenses()`

```php
listProvisioningLicenses($max, $offset, $sort, $direction, $phrase, $name, $license_type, $license_version, $org_name, $full_name): object
```

Get All Licenses

This endpoint retrieves all licenses.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningLicensesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$license_type = win; // string | If specified will return an exact match on licenseType code
$license_version = 2012 R2 Standard; // string | If specified will return an exact match on licenseVersion
$org_name = Acme Motors; // string | If specified will return an exact match on orgName
$full_name = Bugs Bunny; // string | If specified will return an exact match on fullName

try {
    $result = $apiInstance->listProvisioningLicenses($max, $offset, $sort, $direction, $phrase, $name, $license_type, $license_version, $org_name, $full_name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningLicensesApi->listProvisioningLicenses: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **license_type** | **string**| If specified will return an exact match on licenseType code | [optional]
 **license_version** | **string**| If specified will return an exact match on licenseVersion | [optional]
 **org_name** | **string**| If specified will return an exact match on orgName | [optional]
 **full_name** | **string**| If specified will return an exact match on fullName | [optional]

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

## `removeProvisioningLicense()`

```php
removeProvisioningLicense($id): \OpenAPI\Client\Model\Model200Success
```

Delete a License

Will delete a license.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningLicensesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeProvisioningLicense($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningLicensesApi->removeProvisioningLicense: ', $e->getMessage(), PHP_EOL;
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

## `updateProvisioningLicense()`

```php
updateProvisioningLicense($id, $inline_object203): Model200Success
```

Update a License

Use this command to update an existing license.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ProvisioningLicensesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object203 = new \OpenAPI\Client\Model\InlineObject203(); // \OpenAPI\Client\Model\InlineObject203

try {
    $result = $apiInstance->updateProvisioningLicense($id, $inline_object203);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ProvisioningLicensesApi->updateProvisioningLicense: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object203** | [**\OpenAPI\Client\Model\InlineObject203**](../Model/InlineObject203.md)|  | [optional]

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
