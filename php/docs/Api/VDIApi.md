# OpenAPI\Client\VDIApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addVDIApps()**](VDIApi.md#addVDIApps) | **POST** /api/vdi-apps | Creates a VDI App
[**addVDIGateways()**](VDIApi.md#addVDIGateways) | **POST** /api/vdi-gateways | Creates a VDI Gateway
[**addVDIPools()**](VDIApi.md#addVDIPools) | **POST** /api/vdi-pools | Creates a VDI Pool
[**addVdiAllocation()**](VDIApi.md#addVdiAllocation) | **POST** /api/vdi/{id}/allocate | Allocate Virtual Desktop
[**getVDIAllocations()**](VDIApi.md#getVDIAllocations) | **GET** /api/vdi-allocations/{id} | Retrieves a Specific VDI Allocation
[**getVDIApps()**](VDIApi.md#getVDIApps) | **GET** /api/vdi-apps/{id} | Retrieves a Specific VDI App
[**getVDIGateways()**](VDIApi.md#getVDIGateways) | **GET** /api/vdi-gateways/{id} | Retrieves a Specific VDI Gateway
[**getVDIPools()**](VDIApi.md#getVDIPools) | **GET** /api/vdi-pools/{id} | Retrieves a Specific VDI Pool
[**getVdi()**](VDIApi.md#getVdi) | **GET** /api/vdi/{id} | Get a Specific Virtual Desktop
[**listVDIAllocations()**](VDIApi.md#listVDIAllocations) | **GET** /api/vdi-allocations | Retrieves all VDI Allocations
[**listVDIApps()**](VDIApi.md#listVDIApps) | **GET** /api/vdi-apps | Retrieves all VDI Apps
[**listVDIGateways()**](VDIApi.md#listVDIGateways) | **GET** /api/vdi-gateways | Retrieves all VDI Gateways
[**listVDIPools()**](VDIApi.md#listVDIPools) | **GET** /api/vdi-pools | Retrieves all VDI Pools
[**listVdi()**](VDIApi.md#listVdi) | **GET** /api/vdi | List Virtual Desktops
[**removeVDIApps()**](VDIApi.md#removeVDIApps) | **DELETE** /api/vdi-apps/{id} | Deletes a VDI App
[**removeVDIGateways()**](VDIApi.md#removeVDIGateways) | **DELETE** /api/vdi-gateways/{id} | Deletes a VDI Gateway
[**removeVDIPools()**](VDIApi.md#removeVDIPools) | **DELETE** /api/vdi-pools/{id} | Deletes a VDI Pool
[**updateVDIApps()**](VDIApi.md#updateVDIApps) | **PUT** /api/vdi-apps/{id} | Updates a VDI App Configuration or Icon
[**updateVDIGateways()**](VDIApi.md#updateVDIGateways) | **PUT** /api/vdi-gateways/{id} | Updates a VDI Gateway Configuration
[**updateVDIPools()**](VDIApi.md#updateVDIPools) | **PUT** /api/vdi-pools/{id} | Updates a VDI Pool Configuration or Icon


## `addVDIApps()`

```php
addVDIApps($inline_object257): AnyOfObject200Success
```

Creates a VDI App

Creates a VDI app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object257 = new \OpenAPI\Client\Model\InlineObject257(); // \OpenAPI\Client\Model\InlineObject257

try {
    $result = $apiInstance->addVDIApps($inline_object257);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->addVDIApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object257** | [**\OpenAPI\Client\Model\InlineObject257**](../Model/InlineObject257.md)|  | [optional]

### Return type

[**AnyOfObject200Success**](../Model/AnyOfObject200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addVDIGateways()`

```php
addVDIGateways($inline_object259): AnyOfObject200Success
```

Creates a VDI Gateway

Creates a VDI gateway.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object259 = new \OpenAPI\Client\Model\InlineObject259(); // \OpenAPI\Client\Model\InlineObject259

try {
    $result = $apiInstance->addVDIGateways($inline_object259);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->addVDIGateways: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object259** | [**\OpenAPI\Client\Model\InlineObject259**](../Model/InlineObject259.md)|  | [optional]

### Return type

[**AnyOfObject200Success**](../Model/AnyOfObject200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addVDIPools()`

```php
addVDIPools($inline_object261): AnyOfObject200Success
```

Creates a VDI Pool

Creates a VDI pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object261 = new \OpenAPI\Client\Model\InlineObject261(); // \OpenAPI\Client\Model\InlineObject261

try {
    $result = $apiInstance->addVDIPools($inline_object261);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->addVDIPools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object261** | [**\OpenAPI\Client\Model\InlineObject261**](../Model/InlineObject261.md)|  | [optional]

### Return type

[**AnyOfObject200Success**](../Model/AnyOfObject200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addVdiAllocation()`

```php
addVdiAllocation($id): object
```

Allocate Virtual Desktop

This endpoint allocates a specific virtual desktop for use by your user. It will return the desktop and its allocation for your user, or an error if allocation fails, which will occur if the desktop is fully allocated already. If your user already has an allocation, the desktop and allocation will still be returned succesfully and the server does not make any changes.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->addVdiAllocation($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->addVdiAllocation: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

## `getVDIAllocations()`

```php
getVDIAllocations($id): \OpenAPI\Client\Model\InlineResponse200163
```

Retrieves a Specific VDI Allocation

Retrieves a specific VDI allocation.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getVDIAllocations($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->getVDIAllocations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200163**](../Model/InlineResponse200163.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getVDIApps()`

```php
getVDIApps($id): \OpenAPI\Client\Model\InlineResponse200160
```

Retrieves a Specific VDI App

Retrieves a specific VDI app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getVDIApps($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->getVDIApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200160**](../Model/InlineResponse200160.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getVDIGateways()`

```php
getVDIGateways($id): \OpenAPI\Client\Model\InlineResponse200161
```

Retrieves a Specific VDI Gateway

Retrieves a specific VDI gateway.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getVDIGateways($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->getVDIGateways: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200161**](../Model/InlineResponse200161.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getVDIPools()`

```php
getVDIPools($id): \OpenAPI\Client\Model\InlineResponse200162
```

Retrieves a Specific VDI Pool

Retrieves a specific VDI pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getVDIPools($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->getVDIPools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200162**](../Model/InlineResponse200162.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getVdi()`

```php
getVdi($id): \OpenAPI\Client\Model\InlineResponse200164
```

Get a Specific Virtual Desktop

This endpoint retrieves a specific virtual desktop along with the allocation for your user if one exists.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getVdi($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->getVdi: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200164**](../Model/InlineResponse200164.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listVDIAllocations()`

```php
listVDIAllocations($max, $offset, $sort, $direction, $phrase, $name, $id, $status, $pool_id, $user_id): object
```

Retrieves all VDI Allocations

Retrieves all VDI allocations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
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
$id = preparing; // string | Filter by allocation ID
$status = preparing; // string | Filter by allocation status
$pool_id = 1; // int | Filter by `VDI Pool` ID
$user_id = 6; // int | Filter by User ID

try {
    $result = $apiInstance->listVDIAllocations($max, $offset, $sort, $direction, $phrase, $name, $id, $status, $pool_id, $user_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->listVDIAllocations: ', $e->getMessage(), PHP_EOL;
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
 **id** | **string**| Filter by allocation ID | [optional]
 **status** | **string**| Filter by allocation status | [optional]
 **pool_id** | **int**| Filter by &#x60;VDI Pool&#x60; ID | [optional]
 **user_id** | **int**| Filter by User ID | [optional]

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

## `listVDIApps()`

```php
listVDIApps($max, $offset, $sort, $direction, $phrase, $name, $description): object
```

Retrieves all VDI Apps

Retrieves all VDI apps.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
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
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`

try {
    $result = $apiInstance->listVDIApps($max, $offset, $sort, $direction, $phrase, $name, $description);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->listVDIApps: ', $e->getMessage(), PHP_EOL;
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
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]

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

## `listVDIGateways()`

```php
listVDIGateways($max, $offset, $sort, $direction, $phrase, $name, $description): object
```

Retrieves all VDI Gateways

Retrieves all VDI gateways.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
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
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`

try {
    $result = $apiInstance->listVDIGateways($max, $offset, $sort, $direction, $phrase, $name, $description);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->listVDIGateways: ', $e->getMessage(), PHP_EOL;
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
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]

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

## `listVDIPools()`

```php
listVDIPools($max, $offset, $sort, $direction, $phrase, $name, $description, $enabled): object
```

Retrieves all VDI Pools

Retrieves all VDI pools.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
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
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`
$enabled = false; // bool | Filter by enabled

try {
    $result = $apiInstance->listVDIPools($max, $offset, $sort, $direction, $phrase, $name, $description, $enabled);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->listVDIPools: ', $e->getMessage(), PHP_EOL;
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
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **enabled** | **bool**| Filter by enabled | [optional]

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

## `listVdi()`

```php
listVdi($phrase, $max, $offset, $sort, $direction, $name, $description): object
```

List Virtual Desktops

This endpoint retrieves all virtual desktops along with the allocation for your user if one exists.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`

try {
    $result = $apiInstance->listVdi($phrase, $max, $offset, $sort, $direction, $name, $description);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->listVdi: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]

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

## `removeVDIApps()`

```php
removeVDIApps($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a VDI App

Deletes a specified VDI App.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeVDIApps($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->removeVDIApps: ', $e->getMessage(), PHP_EOL;
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

## `removeVDIGateways()`

```php
removeVDIGateways($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a VDI Gateway

Deletes a specified VDI Gateway.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeVDIGateways($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->removeVDIGateways: ', $e->getMessage(), PHP_EOL;
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

## `removeVDIPools()`

```php
removeVDIPools($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a VDI Pool

Deletes a specified VDI Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeVDIPools($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->removeVDIPools: ', $e->getMessage(), PHP_EOL;
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

## `updateVDIApps()`

```php
updateVDIApps($id, $inline_object258): AnyOfObject200Success
```

Updates a VDI App Configuration or Icon

Updates a VDI App configuration or icon.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object258 = new \OpenAPI\Client\Model\InlineObject258(); // \OpenAPI\Client\Model\InlineObject258

try {
    $result = $apiInstance->updateVDIApps($id, $inline_object258);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->updateVDIApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object258** | [**\OpenAPI\Client\Model\InlineObject258**](../Model/InlineObject258.md)|  | [optional]

### Return type

[**AnyOfObject200Success**](../Model/AnyOfObject200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateVDIGateways()`

```php
updateVDIGateways($id, $inline_object260): AnyOfObject200Success
```

Updates a VDI Gateway Configuration

Updates a VDI Gateway configuration or icon.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object260 = new \OpenAPI\Client\Model\InlineObject260(); // \OpenAPI\Client\Model\InlineObject260

try {
    $result = $apiInstance->updateVDIGateways($id, $inline_object260);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->updateVDIGateways: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object260** | [**\OpenAPI\Client\Model\InlineObject260**](../Model/InlineObject260.md)|  | [optional]

### Return type

[**AnyOfObject200Success**](../Model/AnyOfObject200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateVDIPools()`

```php
updateVDIPools($id, $inline_object262): AnyOfObject200Success
```

Updates a VDI Pool Configuration or Icon

Updates a VDI Pool configuration or icon.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\VDIApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object262 = new \OpenAPI\Client\Model\InlineObject262(); // \OpenAPI\Client\Model\InlineObject262

try {
    $result = $apiInstance->updateVDIPools($id, $inline_object262);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VDIApi->updateVDIPools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object262** | [**\OpenAPI\Client\Model\InlineObject262**](../Model/InlineObject262.md)|  | [optional]

### Return type

[**AnyOfObject200Success**](../Model/AnyOfObject200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
