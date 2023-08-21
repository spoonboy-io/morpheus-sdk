# OpenAPI\Client\StorageApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addStorageBuckets()**](StorageApi.md#addStorageBuckets) | **POST** /api/storage-buckets | Creates a Storage Bucket
[**addStorageServers()**](StorageApi.md#addStorageServers) | **POST** /api/storage-servers | Creates a Storage Server
[**addStorageVolumes()**](StorageApi.md#addStorageVolumes) | **POST** /api/storage-volumes | Creates a Storage Volume
[**getStorageBuckets()**](StorageApi.md#getStorageBuckets) | **GET** /api/storage-buckets/{id} | Retrieves a Specific Storage Bucket
[**getStorageServerTypes()**](StorageApi.md#getStorageServerTypes) | **GET** /api/storage-server-types/{id} | Retrieves a Specific Storage Server Type
[**getStorageServers()**](StorageApi.md#getStorageServers) | **GET** /api/storage-servers/{id} | Retrieves a Specific Storage Server
[**getStorageVolumeTypes()**](StorageApi.md#getStorageVolumeTypes) | **GET** /api/storage-volume-types/{id} | Retrieves a Specific Storage Volume Type
[**getStorageVolumes()**](StorageApi.md#getStorageVolumes) | **GET** /api/storage-volumes/{id} | Retrieves a Specific Storage Volume
[**listStorageBuckets()**](StorageApi.md#listStorageBuckets) | **GET** /api/storage-buckets | Retrieves all Storage Buckets
[**listStorageServerTypes()**](StorageApi.md#listStorageServerTypes) | **GET** /api/storage-server-types | Retrieves all Storage Server Types
[**listStorageServers()**](StorageApi.md#listStorageServers) | **GET** /api/storage-servers | Retrieves all Storage Servers
[**listStorageVolumeTypes()**](StorageApi.md#listStorageVolumeTypes) | **GET** /api/storage-volume-types | Retrieves all Storage Volume Types
[**listStorageVolumes()**](StorageApi.md#listStorageVolumes) | **GET** /api/storage-volumes | Retrieves all Storage Volumes
[**removeStorageBuckets()**](StorageApi.md#removeStorageBuckets) | **DELETE** /api/storage-buckets/{id} | Deletes a Storage Bucket
[**removeStorageServers()**](StorageApi.md#removeStorageServers) | **DELETE** /api/storage-servers/{id} | Deletes a Storage Server
[**removeStorageVolumes()**](StorageApi.md#removeStorageVolumes) | **DELETE** /api/storage-volumes/{id} | Deletes a Storage Volume
[**updateStorageBuckets()**](StorageApi.md#updateStorageBuckets) | **PUT** /api/storage-buckets/{id} | Updates a Storage Bucket
[**updateStorageServers()**](StorageApi.md#updateStorageServers) | **PUT** /api/storage-servers/{id} | Updates a Storage Server
[**updateStorageVolumes()**](StorageApi.md#updateStorageVolumes) | **PUT** /api/storage-volumes/{id} | Updates a Storage Volume


## `addStorageBuckets()`

```php
addStorageBuckets($inline_object232): object
```

Creates a Storage Bucket

Creates a storage bucket.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object232 = new \OpenAPI\Client\Model\InlineObject232(); // \OpenAPI\Client\Model\InlineObject232

try {
    $result = $apiInstance->addStorageBuckets($inline_object232);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->addStorageBuckets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object232** | [**\OpenAPI\Client\Model\InlineObject232**](../Model/InlineObject232.md)|  | [optional]

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

## `addStorageServers()`

```php
addStorageServers($inline_object234): object
```

Creates a Storage Server

Creates a storage server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object234 = new \OpenAPI\Client\Model\InlineObject234(); // \OpenAPI\Client\Model\InlineObject234

try {
    $result = $apiInstance->addStorageServers($inline_object234);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->addStorageServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object234** | [**\OpenAPI\Client\Model\InlineObject234**](../Model/InlineObject234.md)|  | [optional]

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

## `addStorageVolumes()`

```php
addStorageVolumes($inline_object236): object
```

Creates a Storage Volume

This endpoint allows creating a Storage Volume. Only certain types of storage servers support creating and deleting storage volumes, such as 3Par and Isilon. Configuration options vary by `Storage Volume Type`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object236 = new \OpenAPI\Client\Model\InlineObject236(); // \OpenAPI\Client\Model\InlineObject236

try {
    $result = $apiInstance->addStorageVolumes($inline_object236);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->addStorageVolumes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object236** | [**\OpenAPI\Client\Model\InlineObject236**](../Model/InlineObject236.md)|  | [optional]

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

## `getStorageBuckets()`

```php
getStorageBuckets($id): \OpenAPI\Client\Model\InlineResponse200145
```

Retrieves a Specific Storage Bucket

Retrieves a specific storage bucket.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStorageBuckets($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->getStorageBuckets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200145**](../Model/InlineResponse200145.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStorageServerTypes()`

```php
getStorageServerTypes($id): \OpenAPI\Client\Model\InlineResponse200146
```

Retrieves a Specific Storage Server Type

Retrieves a specific storage server type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStorageServerTypes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->getStorageServerTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200146**](../Model/InlineResponse200146.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStorageServers()`

```php
getStorageServers($id): \OpenAPI\Client\Model\InlineResponse200147
```

Retrieves a Specific Storage Server

Retrieves a specific storage server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStorageServers($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->getStorageServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200147**](../Model/InlineResponse200147.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStorageVolumeTypes()`

```php
getStorageVolumeTypes($id): \OpenAPI\Client\Model\InlineResponse200148
```

Retrieves a Specific Storage Volume Type

Retrieves a specific storage volume type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStorageVolumeTypes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->getStorageVolumeTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200148**](../Model/InlineResponse200148.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStorageVolumes()`

```php
getStorageVolumes($id): \OpenAPI\Client\Model\InlineResponse200149
```

Retrieves a Specific Storage Volume

Retrieves a specific storage volume.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStorageVolumes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->getStorageVolumes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200149**](../Model/InlineResponse200149.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listStorageBuckets()`

```php
listStorageBuckets($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Storage Buckets

Retrieves all storage buckets.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
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

try {
    $result = $apiInstance->listStorageBuckets($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->listStorageBuckets: ', $e->getMessage(), PHP_EOL;
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

## `listStorageServerTypes()`

```php
listStorageServerTypes($max, $offset, $sort, $direction, $name, $code, $phrase): object
```

Retrieves all Storage Server Types

Fetch a paginated list of available storage server types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listStorageServerTypes($max, $offset, $sort, $direction, $name, $code, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->listStorageServerTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

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

## `listStorageServers()`

```php
listStorageServers($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Storage Servers

Retrieves all storage servers.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
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

try {
    $result = $apiInstance->listStorageServers($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->listStorageServers: ', $e->getMessage(), PHP_EOL;
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

## `listStorageVolumeTypes()`

```php
listStorageVolumeTypes($max, $offset, $sort, $direction, $name, $code, $phrase): object
```

Retrieves all Storage Volume Types

Fetch a paginated list of available storage volume types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listStorageVolumeTypes($max, $offset, $sort, $direction, $name, $code, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->listStorageVolumeTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

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

## `listStorageVolumes()`

```php
listStorageVolumes($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Storage Volumes

Retrieves all storage volumes.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
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

try {
    $result = $apiInstance->listStorageVolumes($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->listStorageVolumes: ', $e->getMessage(), PHP_EOL;
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

## `removeStorageBuckets()`

```php
removeStorageBuckets($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Storage Bucket

Deletes a specified storage bucket.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeStorageBuckets($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->removeStorageBuckets: ', $e->getMessage(), PHP_EOL;
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

## `removeStorageServers()`

```php
removeStorageServers($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Storage Server

Deletes a specified storage server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeStorageServers($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->removeStorageServers: ', $e->getMessage(), PHP_EOL;
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

## `removeStorageVolumes()`

```php
removeStorageVolumes($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Storage Volume

Deletes a specified storage volume.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeStorageVolumes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->removeStorageVolumes: ', $e->getMessage(), PHP_EOL;
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

## `updateStorageBuckets()`

```php
updateStorageBuckets($id, $inline_object233): object
```

Updates a Storage Bucket

Updates a storage bucket.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object233 = new \OpenAPI\Client\Model\InlineObject233(); // \OpenAPI\Client\Model\InlineObject233

try {
    $result = $apiInstance->updateStorageBuckets($id, $inline_object233);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->updateStorageBuckets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object233** | [**\OpenAPI\Client\Model\InlineObject233**](../Model/InlineObject233.md)|  | [optional]

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

## `updateStorageServers()`

```php
updateStorageServers($id, $inline_object235): object
```

Updates a Storage Server

Updates a storage server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object235 = new \OpenAPI\Client\Model\InlineObject235(); // \OpenAPI\Client\Model\InlineObject235

try {
    $result = $apiInstance->updateStorageServers($id, $inline_object235);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->updateStorageServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object235** | [**\OpenAPI\Client\Model\InlineObject235**](../Model/InlineObject235.md)|  | [optional]

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

## `updateStorageVolumes()`

```php
updateStorageVolumes($id, $inline_object237): object
```

Updates a Storage Volume

Updates a storage volume.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\StorageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object237 = new \OpenAPI\Client\Model\InlineObject237(); // \OpenAPI\Client\Model\InlineObject237

try {
    $result = $apiInstance->updateStorageVolumes($id, $inline_object237);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling StorageApi->updateStorageVolumes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object237** | [**\OpenAPI\Client\Model\InlineObject237**](../Model/InlineObject237.md)|  | [optional]

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
