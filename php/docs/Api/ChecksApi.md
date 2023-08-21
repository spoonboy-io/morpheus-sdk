# OpenAPI\Client\ChecksApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCheckApps()**](ChecksApi.md#addCheckApps) | **POST** /api/monitoring/apps | Create a New Check App
[**addCheckGroups()**](ChecksApi.md#addCheckGroups) | **POST** /api/monitoring/groups | Create a New Check Group
[**addChecks()**](ChecksApi.md#addChecks) | **POST** /api/monitoring/checks | Create a New Check
[**deleteCheckApps()**](ChecksApi.md#deleteCheckApps) | **DELETE** /api/monitoring/apps/{id} | Delete a Specific Check App
[**deleteCheckGroups()**](ChecksApi.md#deleteCheckGroups) | **DELETE** /api/monitoring/groups/{id} | Delete a Specific Check Group
[**deleteChecks()**](ChecksApi.md#deleteChecks) | **DELETE** /api/monitoring/checks/{id} | Delete a Specific Check
[**getCheckApps()**](ChecksApi.md#getCheckApps) | **GET** /api/monitoring/apps/{id} | Get a Specific Check App
[**getCheckGroups()**](ChecksApi.md#getCheckGroups) | **GET** /api/monitoring/groups/{id} | Get a Specific Check Group
[**getCheckTypes()**](ChecksApi.md#getCheckTypes) | **GET** /api/monitoring/check-types/{id} | Get a Specific Check Type
[**getChecks()**](ChecksApi.md#getChecks) | **GET** /api/monitoring/checks/{id} | Get a Specific Check
[**listCheckApps()**](ChecksApi.md#listCheckApps) | **GET** /api/monitoring/apps | List All Check Apps
[**listCheckGroups()**](ChecksApi.md#listCheckGroups) | **GET** /api/monitoring/groups | List All Check Groups
[**listCheckTypes()**](ChecksApi.md#listCheckTypes) | **GET** /api/monitoring/check-types | List All Check Types
[**listChecks()**](ChecksApi.md#listChecks) | **GET** /api/monitoring/checks | List All Checks
[**updateCheckApps()**](ChecksApi.md#updateCheckApps) | **PUT** /api/monitoring/apps/{id} | Update Check App
[**updateCheckGroups()**](ChecksApi.md#updateCheckGroups) | **PUT** /api/monitoring/groups/{id} | Update Check Group
[**updateChecks()**](ChecksApi.md#updateChecks) | **PUT** /api/monitoring/checks/{id} | Updates a Check
[**updateMuteAllCheckApps()**](ChecksApi.md#updateMuteAllCheckApps) | **PUT** /api/monitoring/apps/mute-all | Mute All Check Apps
[**updateMuteAllCheckGroups()**](ChecksApi.md#updateMuteAllCheckGroups) | **PUT** /api/monitoring/groups/mute-all | Mute All Check Groups
[**updateMuteAllChecks()**](ChecksApi.md#updateMuteAllChecks) | **PUT** /api/monitoring/checks/mute-all | Mute All Checks
[**updateMuteCheckApps()**](ChecksApi.md#updateMuteCheckApps) | **PUT** /api/monitoring/apps/{id}/mute | Mute Check App
[**updateMuteCheckGroups()**](ChecksApi.md#updateMuteCheckGroups) | **PUT** /api/monitoring/groups/{id}/mute | Mute Check Group
[**updateMuteChecks()**](ChecksApi.md#updateMuteChecks) | **PUT** /api/monitoring/checks/{id}/mute | Mute Check


## `addCheckApps()`

```php
addCheckApps($inline_object29): object
```

Create a New Check App

Create a new check app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object29 = new \OpenAPI\Client\Model\InlineObject29(); // \OpenAPI\Client\Model\InlineObject29

try {
    $result = $apiInstance->addCheckApps($inline_object29);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->addCheckApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object29** | [**\OpenAPI\Client\Model\InlineObject29**](../Model/InlineObject29.md)|  | [optional]

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

## `addCheckGroups()`

```php
addCheckGroups($inline_object37): object
```

Create a New Check Group

Create a new check group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object37 = new \OpenAPI\Client\Model\InlineObject37(); // \OpenAPI\Client\Model\InlineObject37

try {
    $result = $apiInstance->addCheckGroups($inline_object37);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->addCheckGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object37** | [**\OpenAPI\Client\Model\InlineObject37**](../Model/InlineObject37.md)|  | [optional]

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

## `addChecks()`

```php
addChecks($inline_object33): object
```

Create a New Check

Create a new monitoring check.  SSH tunneling options allow the different check types to tunnel to a host via a proxy, and execute checks relative to the proxy. A SSH tunnel can use your account generated public and private key-pairs or SSH password. It is strongly recommended to use a key-pair.  To enable SSH tunneling for a check, add `tunnelOn`, `sshHost`, `sshUser`, and optionally, `sshPort` and `sshPassword` parameters to any check type config as seen earlier in the Check Types section.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object33 = new \OpenAPI\Client\Model\InlineObject33(); // \OpenAPI\Client\Model\InlineObject33

try {
    $result = $apiInstance->addChecks($inline_object33);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->addChecks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object33** | [**\OpenAPI\Client\Model\InlineObject33**](../Model/InlineObject33.md)|  | [optional]

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

## `deleteCheckApps()`

```php
deleteCheckApps($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Specific Check App

Delete an existing monitoring check app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteCheckApps($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->deleteCheckApps: ', $e->getMessage(), PHP_EOL;
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

## `deleteCheckGroups()`

```php
deleteCheckGroups($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Specific Check Group

Delete an existing monitoring check group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteCheckGroups($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->deleteCheckGroups: ', $e->getMessage(), PHP_EOL;
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

## `deleteChecks()`

```php
deleteChecks($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Specific Check

Delete an existing monitoring check.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteChecks($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->deleteChecks: ', $e->getMessage(), PHP_EOL;
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

## `getCheckApps()`

```php
getCheckApps($id): \OpenAPI\Client\Model\InlineResponse20016
```

Get a Specific Check App

Get details about a specific monitoring check app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCheckApps($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->getCheckApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20016**](../Model/InlineResponse20016.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getCheckGroups()`

```php
getCheckGroups($id): \OpenAPI\Client\Model\InlineResponse20019
```

Get a Specific Check Group

Get details about a specific monitoring check group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCheckGroups($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->getCheckGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20019**](../Model/InlineResponse20019.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getCheckTypes()`

```php
getCheckTypes($id): \OpenAPI\Client\Model\InlineResponse20018
```

Get a Specific Check Type

Get details about a specific monitoring check type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCheckTypes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->getCheckTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20018**](../Model/InlineResponse20018.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getChecks()`

```php
getChecks($id): \OpenAPI\Client\Model\InlineResponse20017
```

Get a Specific Check

Get details about a specific monitoring check.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getChecks($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->getChecks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20017**](../Model/InlineResponse20017.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listCheckApps()`

```php
listCheckApps($max, $offset, $sort, $name, $phrase, $status, $last_updated, $deleted): object
```

List All Check Apps

Get a list of check apps.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$status = running; // string | The instance status for filtering.
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$deleted = true; // bool | If true, only deleted resources or instances in pending removal status are returned.

try {
    $result = $apiInstance->listCheckApps($max, $offset, $sort, $name, $phrase, $status, $last_updated, $deleted);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->listCheckApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **status** | **string**| The instance status for filtering. | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

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

## `listCheckGroups()`

```php
listCheckGroups($max, $offset, $sort, $name, $phrase, $status, $last_updated, $deleted): object
```

List All Check Groups

Get a list of check groups.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$status = running; // string | The instance status for filtering.
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$deleted = true; // bool | If true, only deleted resources or instances in pending removal status are returned.

try {
    $result = $apiInstance->listCheckGroups($max, $offset, $sort, $name, $phrase, $status, $last_updated, $deleted);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->listCheckGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **status** | **string**| The instance status for filtering. | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

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

## `listCheckTypes()`

```php
listCheckTypes($max, $offset, $sort, $name, $phrase): object
```

List All Check Types

Get a list of monitoring check types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listCheckTypes($max, $offset, $sort, $name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->listCheckTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
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

## `listChecks()`

```php
listChecks($max, $offset, $sort, $name, $phrase, $status, $last_updated, $deleted): object
```

List All Checks

Get a list of monitoring checks.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$status = healthy; // string | Filter by health status.
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$deleted = true; // bool | If true, only deleted resources or instances in pending removal status are returned.

try {
    $result = $apiInstance->listChecks($max, $offset, $sort, $name, $phrase, $status, $last_updated, $deleted);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->listChecks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **status** | **string**| Filter by health status. | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

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

## `updateCheckApps()`

```php
updateCheckApps($id, $inline_object31): object
```

Update Check App

Update an existing monitoring check app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object31 = new \OpenAPI\Client\Model\InlineObject31(); // \OpenAPI\Client\Model\InlineObject31

try {
    $result = $apiInstance->updateCheckApps($id, $inline_object31);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateCheckApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object31** | [**\OpenAPI\Client\Model\InlineObject31**](../Model/InlineObject31.md)|  | [optional]

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

## `updateCheckGroups()`

```php
updateCheckGroups($id, $inline_object38): object
```

Update Check Group

Update an existing monitoring check group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object38 = new \OpenAPI\Client\Model\InlineObject38(); // \OpenAPI\Client\Model\InlineObject38

try {
    $result = $apiInstance->updateCheckGroups($id, $inline_object38);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateCheckGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object38** | [**\OpenAPI\Client\Model\InlineObject38**](../Model/InlineObject38.md)|  | [optional]

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

## `updateChecks()`

```php
updateChecks($id, $inline_object35): object
```

Updates a Check

Updates a monitoring check.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object35 = new \OpenAPI\Client\Model\InlineObject35(); // \OpenAPI\Client\Model\InlineObject35

try {
    $result = $apiInstance->updateChecks($id, $inline_object35);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateChecks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object35** | [**\OpenAPI\Client\Model\InlineObject35**](../Model/InlineObject35.md)|  | [optional]

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

## `updateMuteAllCheckApps()`

```php
updateMuteAllCheckApps($inline_object30): object
```

Mute All Check Apps

Mute all existing check apps.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object30 = new \OpenAPI\Client\Model\InlineObject30(); // \OpenAPI\Client\Model\InlineObject30

try {
    $result = $apiInstance->updateMuteAllCheckApps($inline_object30);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateMuteAllCheckApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object30** | [**\OpenAPI\Client\Model\InlineObject30**](../Model/InlineObject30.md)|  | [optional]

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

## `updateMuteAllCheckGroups()`

```php
updateMuteAllCheckGroups($inline_object40): object
```

Mute All Check Groups

Mute all existing check groups.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object40 = new \OpenAPI\Client\Model\InlineObject40(); // \OpenAPI\Client\Model\InlineObject40

try {
    $result = $apiInstance->updateMuteAllCheckGroups($inline_object40);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateMuteAllCheckGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object40** | [**\OpenAPI\Client\Model\InlineObject40**](../Model/InlineObject40.md)|  | [optional]

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

## `updateMuteAllChecks()`

```php
updateMuteAllChecks($inline_object34): object
```

Mute All Checks

Mute all existing checks.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object34 = new \OpenAPI\Client\Model\InlineObject34(); // \OpenAPI\Client\Model\InlineObject34

try {
    $result = $apiInstance->updateMuteAllChecks($inline_object34);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateMuteAllChecks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object34** | [**\OpenAPI\Client\Model\InlineObject34**](../Model/InlineObject34.md)|  | [optional]

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

## `updateMuteCheckApps()`

```php
updateMuteCheckApps($id, $inline_object32): object
```

Mute Check App

Mute an existing check app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object32 = new \OpenAPI\Client\Model\InlineObject32(); // \OpenAPI\Client\Model\InlineObject32

try {
    $result = $apiInstance->updateMuteCheckApps($id, $inline_object32);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateMuteCheckApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object32** | [**\OpenAPI\Client\Model\InlineObject32**](../Model/InlineObject32.md)|  | [optional]

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

## `updateMuteCheckGroups()`

```php
updateMuteCheckGroups($id, $inline_object39): object
```

Mute Check Group

Mute an existing check group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object39 = new \OpenAPI\Client\Model\InlineObject39(); // \OpenAPI\Client\Model\InlineObject39

try {
    $result = $apiInstance->updateMuteCheckGroups($id, $inline_object39);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateMuteCheckGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object39** | [**\OpenAPI\Client\Model\InlineObject39**](../Model/InlineObject39.md)|  | [optional]

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

## `updateMuteChecks()`

```php
updateMuteChecks($id, $inline_object36): object
```

Mute Check

Mute an existing check.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ChecksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object36 = new \OpenAPI\Client\Model\InlineObject36(); // \OpenAPI\Client\Model\InlineObject36

try {
    $result = $apiInstance->updateMuteChecks($id, $inline_object36);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ChecksApi->updateMuteChecks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object36** | [**\OpenAPI\Client\Model\InlineObject36**](../Model/InlineObject36.md)|  | [optional]

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
