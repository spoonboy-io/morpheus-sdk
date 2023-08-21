# OpenAPI\Client\CloudsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCloudResourcePool()**](CloudsApi.md#addCloudResourcePool) | **POST** /api/zones/{zoneId}/resource-pools | Creates a Specified Resource Pool for Specified Cloud
[**addClouds()**](CloudsApi.md#addClouds) | **POST** /api/zones | Creates a Cloud
[**getCloudDatastores()**](CloudsApi.md#getCloudDatastores) | **GET** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**getCloudFolders()**](CloudsApi.md#getCloudFolders) | **GET** /api/zones/{zoneId}/folders/{id} | Retrieves a Resource Folder for Specified Cloud
[**getCloudResourcePools()**](CloudsApi.md#getCloudResourcePools) | **GET** /api/zones/{zoneId}/resource-pools/{id} | Retrieves a Resource Pool for Specified Cloud
[**getCloudTypes()**](CloudsApi.md#getCloudTypes) | **GET** /api/zone-types/{id} | Retrieves a Specific Cloud Type
[**getClouds()**](CloudsApi.md#getClouds) | **GET** /api/zones/{id} | Retrieves a Specific Cloud
[**getWikiCloud()**](CloudsApi.md#getWikiCloud) | **GET** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**listCloudDatastores()**](CloudsApi.md#listCloudDatastores) | **GET** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**listCloudFolders()**](CloudsApi.md#listCloudFolders) | **GET** /api/zones/{zoneId}/folders | Retrieves all resource folders for Specified Cloud
[**listCloudResourcePools()**](CloudsApi.md#listCloudResourcePools) | **GET** /api/zones/{zoneId}/resource-pools | Retrieves all Resource Pools for Specified Cloud
[**listCloudSecurityGroups()**](CloudsApi.md#listCloudSecurityGroups) | **GET** /api/zones/{id}/security-groups | Retrieves all Security Groups for a Cloud
[**listCloudTypes()**](CloudsApi.md#listCloudTypes) | **GET** /api/zone-types | Retrieves all Cloud Types
[**listClouds()**](CloudsApi.md#listClouds) | **GET** /api/zones | Retrieves all Clouds
[**refreshClouds()**](CloudsApi.md#refreshClouds) | **POST** /api/zones/{id}/refresh | Refreshes a Cloud
[**removeCloudResourcePools()**](CloudsApi.md#removeCloudResourcePools) | **DELETE** /api/zones/{zoneId}/resource-pools/{id} | Deletes a Resource Pool for Specified Cloud
[**removeClouds()**](CloudsApi.md#removeClouds) | **DELETE** /api/zones/{id} | Deletes a Cloud
[**updateCloudDatastores()**](CloudsApi.md#updateCloudDatastores) | **PUT** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**updateCloudFolders()**](CloudsApi.md#updateCloudFolders) | **PUT** /api/zones/{zoneId}/folders/{id} | Updates a Resource Folder for Specified Cloud
[**updateCloudLogo()**](CloudsApi.md#updateCloudLogo) | **POST** /api/zones/{id}/update-logo | Update Logo For Cloud
[**updateCloudResourcePool()**](CloudsApi.md#updateCloudResourcePool) | **PUT** /api/zones/{zoneId}/resource-pools/{id} | Updates a Specified Resource Pool for Specified Cloud
[**updateCloudSecurityGroups()**](CloudsApi.md#updateCloudSecurityGroups) | **POST** /api/zones/{id}/security-groups | Sets Security Groups for a Cloud
[**updateClouds()**](CloudsApi.md#updateClouds) | **PUT** /api/zones/{id} | Updates a Cloud
[**updateWikiCloud()**](CloudsApi.md#updateWikiCloud) | **PUT** /api/zones/{id}/wiki | Update a Cloud Wiki Page


## `addCloudResourcePool()`

```php
addCloudResourcePool($zone_id, $inline_object45): \OpenAPI\Client\Model\InlineResponse20024
```

Creates a Specified Resource Pool for Specified Cloud

Creates a resource pool for specified cloud. Only certain types of clouds support creating and deleting resource pools. Configuration options vary by type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$inline_object45 = new \OpenAPI\Client\Model\InlineObject45(); // \OpenAPI\Client\Model\InlineObject45

try {
    $result = $apiInstance->addCloudResourcePool($zone_id, $inline_object45);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->addCloudResourcePool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **inline_object45** | [**\OpenAPI\Client\Model\InlineObject45**](../Model/InlineObject45.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20024**](../Model/InlineResponse20024.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addClouds()`

```php
addClouds($inline_object41): object
```

Creates a Cloud

Creates a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object41 = new \OpenAPI\Client\Model\InlineObject41(); // \OpenAPI\Client\Model\InlineObject41

try {
    $result = $apiInstance->addClouds($inline_object41);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->addClouds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object41** | [**\OpenAPI\Client\Model\InlineObject41**](../Model/InlineObject41.md)|  | [optional]

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

## `getCloudDatastores()`

```php
getCloudDatastores($zone_id, $id): object
```

Retrieves a Datastore for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific datastore under a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCloudDatastores($zone_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->getCloudDatastores: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
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

## `getCloudFolders()`

```php
getCloudFolders($zone_id, $id): object
```

Retrieves a Resource Folder for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific folder under a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCloudFolders($zone_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->getCloudFolders: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
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

## `getCloudResourcePools()`

```php
getCloudResourcePools($zone_id, $id): object
```

Retrieves a Resource Pool for Specified Cloud

This endpoint retrieves a specific resource pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCloudResourcePools($zone_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->getCloudResourcePools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
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

## `getCloudTypes()`

```php
getCloudTypes($id): \OpenAPI\Client\Model\InlineResponse20020
```

Retrieves a Specific Cloud Type

Retrieves a specific cloud type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCloudTypes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->getCloudTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20020**](../Model/InlineResponse20020.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClouds()`

```php
getClouds($id): \OpenAPI\Client\Model\InlineResponse20021
```

Retrieves a Specific Cloud

Retrieves a specific cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getClouds($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->getClouds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20021**](../Model/InlineResponse20021.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getWikiCloud()`

```php
getWikiCloud($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves a Cloud Wiki Page

This endpoint retrieves a cloud Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWikiCloud($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->getWikiCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200168**](../Model/InlineResponse200168.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listCloudDatastores()`

```php
listCloudDatastores($zone_id, $name, $phrase, $max, $sort, $direction): object
```

Retrieves all Datastores for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all data stores under a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort

try {
    $result = $apiInstance->listCloudDatastores($zone_id, $name, $phrase, $max, $sort, $direction);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->listCloudDatastores: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]

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

## `listCloudFolders()`

```php
listCloudFolders($zone_id, $name, $phrase, $max): object
```

Retrieves all resource folders for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource folders under a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records

try {
    $result = $apiInstance->listCloudFolders($zone_id, $name, $phrase, $max);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->listCloudFolders: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]

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

## `listCloudResourcePools()`

```php
listCloudResourcePools($zone_id, $name, $phrase, $max): object
```

Retrieves all Resource Pools for Specified Cloud

Resource Pools can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource pools under a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records

try {
    $result = $apiInstance->listCloudResourcePools($zone_id, $name, $phrase, $max);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->listCloudResourcePools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]

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

## `listCloudSecurityGroups()`

```php
listCloudSecurityGroups($id): object
```

Retrieves all Security Groups for a Cloud

Retrieves all security groups for a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->listCloudSecurityGroups($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->listCloudSecurityGroups: ', $e->getMessage(), PHP_EOL;
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

## `listCloudTypes()`

```php
listCloudTypes($max, $offset, $sort, $direction, $name, $code, $phrase, $provision_type): object
```

Retrieves all Cloud Types

Fetch a paginated list of available cloud types. This returns the configuration options for each type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
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
$provision_type = 'provision_type_example'; // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.

try {
    $result = $apiInstance->listCloudTypes($max, $offset, $sort, $direction, $name, $code, $phrase, $provision_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->listCloudTypes: ', $e->getMessage(), PHP_EOL;
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
 **provision_type** | **string**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings. | [optional]

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

## `listClouds()`

```php
listClouds($last_updated, $type, $group_id, $max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Clouds

Retrieves all clouds.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$type = 'type_example'; // string | If specified will return all zones by cloud type code. Refer to `Zone Types` API for up to date listings.
$group_id = 13; // int | If specified will return all zones assigned to a server group by id. Accepts multiple values.
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listClouds($last_updated, $type, $group_id, $max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->listClouds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **type** | **string**| If specified will return all zones by cloud type code. Refer to &#x60;Zone Types&#x60; API for up to date listings. | [optional]
 **group_id** | **int**| If specified will return all zones assigned to a server group by id. Accepts multiple values. | [optional]
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

## `refreshClouds()`

```php
refreshClouds($id, $inline_object47): \OpenAPI\Client\Model\Model200Success
```

Refreshes a Cloud

Refreshes a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object47 = new \OpenAPI\Client\Model\InlineObject47(); // \OpenAPI\Client\Model\InlineObject47

try {
    $result = $apiInstance->refreshClouds($id, $inline_object47);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->refreshClouds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object47** | [**\OpenAPI\Client\Model\InlineObject47**](../Model/InlineObject47.md)|  | [optional]

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

## `removeCloudResourcePools()`

```php
removeCloudResourcePools($zone_id, $id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Resource Pool for Specified Cloud

Deletes a resource pool for specified Cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeCloudResourcePools($zone_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->removeCloudResourcePools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
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

## `removeClouds()`

```php
removeClouds($id, $remove_resources): \OpenAPI\Client\Model\Model200Success
```

Deletes a Cloud

Deletes a specified Cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$remove_resources = false; // bool | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute.

try {
    $result = $apiInstance->removeClouds($id, $remove_resources);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->removeClouds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_resources** | **bool**| Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [optional] [default to false]

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

## `updateCloudDatastores()`

```php
updateCloudDatastores($zone_id, $id, $inline_object43): \OpenAPI\Client\Model\InlineResponse20022
```

Updates a Specified Datastore for Specified Cloud

Updates a datastore for specified cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object43 = new \OpenAPI\Client\Model\InlineObject43(); // \OpenAPI\Client\Model\InlineObject43

try {
    $result = $apiInstance->updateCloudDatastores($zone_id, $id, $inline_object43);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateCloudDatastores: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object43** | [**\OpenAPI\Client\Model\InlineObject43**](../Model/InlineObject43.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20022**](../Model/InlineResponse20022.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateCloudFolders()`

```php
updateCloudFolders($zone_id, $id, $inline_object44): \OpenAPI\Client\Model\InlineResponse20023
```

Updates a Resource Folder for Specified Cloud

Updates a resource folder for specified cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object44 = new \OpenAPI\Client\Model\InlineObject44(); // \OpenAPI\Client\Model\InlineObject44

try {
    $result = $apiInstance->updateCloudFolders($zone_id, $id, $inline_object44);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateCloudFolders: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object44** | [**\OpenAPI\Client\Model\InlineObject44**](../Model/InlineObject44.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20023**](../Model/InlineResponse20023.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateCloudLogo()`

```php
updateCloudLogo($id, $logo, $dark_logo): \OpenAPI\Client\Model\Model200Success
```

Update Logo For Cloud

Use this command to update the logo and dark logo images for a cloud. This endpoint expects multipart form data as the request format, not JSON.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$logo = "/path/to/file.txt"; // \SplFileObject | Logo File png,jpg,svg
$dark_logo = "/path/to/file.txt"; // \SplFileObject | Dark Logo File png,jpg,svg

try {
    $result = $apiInstance->updateCloudLogo($id, $logo, $dark_logo);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateCloudLogo: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **logo** | **\SplFileObject****\SplFileObject**| Logo File png,jpg,svg | [optional]
 **dark_logo** | **\SplFileObject****\SplFileObject**| Dark Logo File png,jpg,svg | [optional]

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

## `updateCloudResourcePool()`

```php
updateCloudResourcePool($zone_id, $id, $inline_object46): \OpenAPI\Client\Model\InlineResponse20024
```

Updates a Specified Resource Pool for Specified Cloud

Updates a resource pool for specified cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 7; // float | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object46 = new \OpenAPI\Client\Model\InlineObject46(); // \OpenAPI\Client\Model\InlineObject46

try {
    $result = $apiInstance->updateCloudResourcePool($zone_id, $id, $inline_object46);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateCloudResourcePool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object46** | [**\OpenAPI\Client\Model\InlineObject46**](../Model/InlineObject46.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20024**](../Model/InlineResponse20024.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateCloudSecurityGroups()`

```php
updateCloudSecurityGroups($id, $inline_object49): object
```

Sets Security Groups for a Cloud

Sets security groups for acloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object49 = new \OpenAPI\Client\Model\InlineObject49(); // \OpenAPI\Client\Model\InlineObject49

try {
    $result = $apiInstance->updateCloudSecurityGroups($id, $inline_object49);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateCloudSecurityGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object49** | [**\OpenAPI\Client\Model\InlineObject49**](../Model/InlineObject49.md)|  | [optional]

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

## `updateClouds()`

```php
updateClouds($id, $inline_object42): object
```

Updates a Cloud

Updates a cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object42 = new \OpenAPI\Client\Model\InlineObject42(); // \OpenAPI\Client\Model\InlineObject42

try {
    $result = $apiInstance->updateClouds($id, $inline_object42);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateClouds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object42** | [**\OpenAPI\Client\Model\InlineObject42**](../Model/InlineObject42.md)|  | [optional]

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

## `updateWikiCloud()`

```php
updateWikiCloud($id, $inline_object274): object
```

Update a Cloud Wiki Page

Updates a cloud Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CloudsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object274 = new \OpenAPI\Client\Model\InlineObject274(); // \OpenAPI\Client\Model\InlineObject274

try {
    $result = $apiInstance->updateWikiCloud($id, $inline_object274);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CloudsApi->updateWikiCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object274** | [**\OpenAPI\Client\Model\InlineObject274**](../Model/InlineObject274.md)|  | [optional]

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
