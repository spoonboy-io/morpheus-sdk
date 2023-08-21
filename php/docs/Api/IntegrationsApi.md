# OpenAPI\Client\IntegrationsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIntegrationSnowObjects()**](IntegrationsApi.md#addIntegrationSnowObjects) | **POST** /api/integrations/{id}/objects | Creates an Exposed ServiceNow Catalog Item
[**addIntegrations()**](IntegrationsApi.md#addIntegrations) | **POST** /api/integrations | Creates an Integration
[**getIntegrationInventory()**](IntegrationsApi.md#getIntegrationInventory) | **GET** /api/integrations/{id}/inventory/{inventoryId} | Get a Specific Integration Inventory
[**getIntegrationObjects()**](IntegrationsApi.md#getIntegrationObjects) | **GET** /api/integrations/{id}/objects/{objectId} | Get a Specific ServiceNow Integration Object
[**getIntegrationTypeOptionTypes()**](IntegrationsApi.md#getIntegrationTypeOptionTypes) | **GET** /api/integration-types/{id}/option-types | Retrieves a Option Types for a Specific Integration Type
[**getIntegrationTypes()**](IntegrationsApi.md#getIntegrationTypes) | **GET** /api/integration-types/{id} | Retrieves a Specific Integration Type
[**getIntegrations()**](IntegrationsApi.md#getIntegrations) | **GET** /api/integrations/{id} | Retrieves a Specific Integration
[**listIntegrationInventory()**](IntegrationsApi.md#listIntegrationInventory) | **GET** /api/integrations/{id}/inventory | Get All Integration Inventory
[**listIntegrationObjects()**](IntegrationsApi.md#listIntegrationObjects) | **GET** /api/integrations/{id}/objects | Get ServiceNow Integration Objects
[**listIntegrationTypes()**](IntegrationsApi.md#listIntegrationTypes) | **GET** /api/integration-types | Retrieves all Integration Types
[**listIntegrations()**](IntegrationsApi.md#listIntegrations) | **GET** /api/integrations | Retrieves all Integrations
[**refreshIntegrations()**](IntegrationsApi.md#refreshIntegrations) | **POST** /api/integrations/{id}/refresh | Refresh an Integration
[**removeIntegrationObjects()**](IntegrationsApi.md#removeIntegrationObjects) | **DELETE** /api/integrations/{id}/objects/{objectId} | Deletes a ServiceNow Integration object
[**removeIntegrations()**](IntegrationsApi.md#removeIntegrations) | **DELETE** /api/integrations/{id} | Deletes an Integration
[**updateIntegrationInventory()**](IntegrationsApi.md#updateIntegrationInventory) | **PUT** /api/integrations/{id}/inventory/{inventoryId} | Updating an Integration Inventory
[**updateIntegrations()**](IntegrationsApi.md#updateIntegrations) | **PUT** /api/integrations/{id} | Updates an Integration


## `addIntegrationSnowObjects()`

```php
addIntegrationSnowObjects($id, $inline_object100): object
```

Creates an Exposed ServiceNow Catalog Item

This endpoint creates an Exposed Catalog Item. This is an integration object of type `catalog` that references a `Catalog Item Type.`

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object100 = new \OpenAPI\Client\Model\InlineObject100(); // \OpenAPI\Client\Model\InlineObject100

try {
    $result = $apiInstance->addIntegrationSnowObjects($id, $inline_object100);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->addIntegrationSnowObjects: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object100** | [**\OpenAPI\Client\Model\InlineObject100**](../Model/InlineObject100.md)|  | [optional]

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

## `addIntegrations()`

```php
addIntegrations($unknown_base_type): object
```

Creates an Integration

Creates an integration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->addIntegrations($unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->addIntegrations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `getIntegrationInventory()`

```php
getIntegrationInventory($id, $inventory_id): \OpenAPI\Client\Model\InlineResponse20065
```

Get a Specific Integration Inventory

This endpoint retrieves a specific integration inventory. Only certain types of integrations support this operation, such as Ansible Tower.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 9; // int | Morpheus ID of the Integration
$inventory_id = 1; // int | Morpheus ID of the Integration Inventory

try {
    $result = $apiInstance->getIntegrationInventory($id, $inventory_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->getIntegrationInventory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Integration |
 **inventory_id** | **int**| Morpheus ID of the Integration Inventory |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20065**](../Model/InlineResponse20065.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getIntegrationObjects()`

```php
getIntegrationObjects($id, $object_id): \OpenAPI\Client\Model\InlineResponse20064
```

Get a Specific ServiceNow Integration Object

This endpoint retrieves a specific ServiceNow integration object.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$object_id = 122; // int | Morpheus ID of the Object being created or referenced

try {
    $result = $apiInstance->getIntegrationObjects($id, $object_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->getIntegrationObjects: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **object_id** | **int**| Morpheus ID of the Object being created or referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20064**](../Model/InlineResponse20064.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getIntegrationTypeOptionTypes()`

```php
getIntegrationTypeOptionTypes($id): \OpenAPI\Client\Model\InlineResponse20062
```

Retrieves a Option Types for a Specific Integration Type

This endpoint will retrieve the list of option types for a specific integration type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getIntegrationTypeOptionTypes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->getIntegrationTypeOptionTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20062**](../Model/InlineResponse20062.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getIntegrationTypes()`

```php
getIntegrationTypes($id, $optiontypes): \OpenAPI\Client\Model\InlineResponse20061
```

Retrieves a Specific Integration Type

Retrieves a specific integration type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$optiontypes = false; // bool | Pass `true` to include optionTypes in the response for each integration type

try {
    $result = $apiInstance->getIntegrationTypes($id, $optiontypes);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->getIntegrationTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **optiontypes** | **bool**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] [default to false]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20061**](../Model/InlineResponse20061.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getIntegrations()`

```php
getIntegrations($id): object
```

Retrieves a Specific Integration

Retrieves a specific integration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getIntegrations($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->getIntegrations: ', $e->getMessage(), PHP_EOL;
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

## `listIntegrationInventory()`

```php
listIntegrationInventory($id, $max, $offset, $sort, $direction, $phrase, $name): object
```

Get All Integration Inventory

This endpoint retrieves a list of inventory for a specific integration. Only certain types of integrations support this operation, such as Ansible Tower.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 9; // int | Morpheus ID of the Integration
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listIntegrationInventory($id, $max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->listIntegrationInventory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Integration |
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

## `listIntegrationObjects()`

```php
listIntegrationObjects($id, $max, $offset, $sort, $direction, $phrase, $name, $value, $ref_id): \OpenAPI\Client\Model\InlineResponse20063
```

Get ServiceNow Integration Objects

This endpoint retrieves a list of exposed objects for a specific ServiceNow integration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$value = 'value_example'; // string | The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing `type=string`. This means the data value returned by the API will be a string instead of an object.
$ref_id = 3; // int | If specified will return an exact match on refId

try {
    $result = $apiInstance->listIntegrationObjects($id, $max, $offset, $sort, $direction, $phrase, $name, $value, $ref_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->listIntegrationObjects: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **value** | **string**| The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the data value returned by the API will be a string instead of an object. | [optional]
 **ref_id** | **int**| If specified will return an exact match on refId | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20063**](../Model/InlineResponse20063.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listIntegrationTypes()`

```php
listIntegrationTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $optiontypes, $description, $category, $creatable, $enabled, $has_cmdb, $has_cmdb_discovery, $has_cm, $has_dns, $has_approvals, $is_plugin): object
```

Retrieves all Integration Types

An Integration Type is specific third party software that the appliance can be configured to integrate with, such as Ansible, Chef, Git, ServiceNOW, etc.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
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
$code = azr; // string | If specified will return an exact match on code
$optiontypes = false; // bool | Pass `true` to include optionTypes in the response for each integration type
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`
$category = 'category_example'; // string | If specified will return an exact match on category
$creatable = True; // bool | Filter by creatable
$enabled = false; // bool | Filter by enabled
$has_cmdb = True; // bool | Filter by hasCMDB
$has_cmdb_discovery = True; // bool | Filter by hasCMDBDiscovery
$has_cm = True; // bool | Filter by hasCM
$has_dns = True; // bool | Filter by hasDNS
$has_approvals = True; // bool | Filter by hasApprovals
$is_plugin = True; // bool | Filter by isPlugin

try {
    $result = $apiInstance->listIntegrationTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $optiontypes, $description, $category, $creatable, $enabled, $has_cmdb, $has_cmdb_discovery, $has_cm, $has_dns, $has_approvals, $is_plugin);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->listIntegrationTypes: ', $e->getMessage(), PHP_EOL;
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
 **code** | **string**| If specified will return an exact match on code | [optional]
 **optiontypes** | **bool**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] [default to false]
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **category** | **string**| If specified will return an exact match on category | [optional]
 **creatable** | **bool**| Filter by creatable | [optional]
 **enabled** | **bool**| Filter by enabled | [optional]
 **has_cmdb** | **bool**| Filter by hasCMDB | [optional]
 **has_cmdb_discovery** | **bool**| Filter by hasCMDBDiscovery | [optional]
 **has_cm** | **bool**| Filter by hasCM | [optional]
 **has_dns** | **bool**| Filter by hasDNS | [optional]
 **has_approvals** | **bool**| Filter by hasApprovals | [optional]
 **is_plugin** | **bool**| Filter by isPlugin | [optional]

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

## `listIntegrations()`

```php
listIntegrations($max, $offset, $sort, $direction, $phrase, $name, $id, $url, $host, $port, $username, $version, $windows_version, $status, $objects): AnyOfObjectMeta
```

Retrieves all Integrations

Retrieves all integrations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
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
$id = 7; // int | Morpheus ID of the Object being created or referenced
$url = https://example.com/testimage.ovf; // string | Download the file from a remote url. This can be used instead of uploading a local file.
$host = 'host_example'; // string | Filter by integration Host
$port = 'port_example'; // string | Filter by integration Port
$username = administrator; // string | Username
$version = 5; // int | Filter by version number (userVersion)
$windows_version = 'windows_version_example'; // string | Filter by integration Windows Version
$status = running; // string | The instance status for filtering.
$objects = false; // bool | Include `objects=true` to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.

try {
    $result = $apiInstance->listIntegrations($max, $offset, $sort, $direction, $phrase, $name, $id, $url, $host, $port, $username, $version, $windows_version, $status, $objects);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->listIntegrations: ', $e->getMessage(), PHP_EOL;
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
 **id** | **int**| Morpheus ID of the Object being created or referenced | [optional]
 **url** | **string**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional]
 **host** | **string**| Filter by integration Host | [optional]
 **port** | **string**| Filter by integration Port | [optional]
 **username** | **string**| Username | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **windows_version** | **string**| Filter by integration Windows Version | [optional]
 **status** | **string**| The instance status for filtering. | [optional]
 **objects** | **bool**| Include &#x60;objects&#x3D;true&#x60; to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2. | [optional] [default to false]

### Return type

[**AnyOfObjectMeta**](../Model/AnyOfObjectMeta.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `refreshIntegrations()`

```php
refreshIntegrations($id): object
```

Refresh an Integration

This endpoint will refresh an existing Integration. Only some types support this and will actually perform an action as a result.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->refreshIntegrations($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->refreshIntegrations: ', $e->getMessage(), PHP_EOL;
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

## `removeIntegrationObjects()`

```php
removeIntegrationObjects($id, $object_id): \OpenAPI\Client\Model\Model200Success
```

Deletes a ServiceNow Integration object

Deletes a specified ServiceNow integration object.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$object_id = 122; // int | Morpheus ID of the Object being created or referenced

try {
    $result = $apiInstance->removeIntegrationObjects($id, $object_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->removeIntegrationObjects: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **object_id** | **int**| Morpheus ID of the Object being created or referenced |

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

## `removeIntegrations()`

```php
removeIntegrations($id): \OpenAPI\Client\Model\Model200Success
```

Deletes an Integration

Deletes a specified integration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeIntegrations($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->removeIntegrations: ', $e->getMessage(), PHP_EOL;
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

## `updateIntegrationInventory()`

```php
updateIntegrationInventory($id, $inventory_id, $inline_object101): object
```

Updating an Integration Inventory

This endpoint provides updating of integration inventory

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 9; // int | Morpheus ID of the Integration
$inventory_id = 1; // int | Morpheus ID of the Integration Inventory
$inline_object101 = new \OpenAPI\Client\Model\InlineObject101(); // \OpenAPI\Client\Model\InlineObject101

try {
    $result = $apiInstance->updateIntegrationInventory($id, $inventory_id, $inline_object101);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->updateIntegrationInventory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Integration |
 **inventory_id** | **int**| Morpheus ID of the Integration Inventory |
 **inline_object101** | [**\OpenAPI\Client\Model\InlineObject101**](../Model/InlineObject101.md)|  | [optional]

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

## `updateIntegrations()`

```php
updateIntegrations($id, $unknown_base_type): object
```

Updates an Integration

Updates an integration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IntegrationsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = new \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE(); // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateIntegrations($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IntegrationsApi->updateIntegrations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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
