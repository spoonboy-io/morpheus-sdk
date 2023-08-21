# OpenAPI\Client\AppsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addAppInstance()**](AppsApi.md#addAppInstance) | **POST** /api/apps/{id}/add-instance | Add Existing Instance to App
[**addAppUndoDelete()**](AppsApi.md#addAppUndoDelete) | **PUT** /api/apps/{id}/cancel-removal | Undo Delete of an App
[**addApps()**](AppsApi.md#addApps) | **POST** /api/apps | Create an App
[**applyAppState()**](AppsApi.md#applyAppState) | **POST** /api/apps/{id}/apply | Apply State of an App
[**deleteApp()**](AppsApi.md#deleteApp) | **DELETE** /api/apps/{id} | Delete an App
[**getApp()**](AppsApi.md#getApp) | **GET** /api/apps/{id} | Get a Specific App
[**getAppSecurityGroups()**](AppsApi.md#getAppSecurityGroups) | **GET** /api/apps/{id}/security-groups | Get Security Groups for an App
[**getAppState()**](AppsApi.md#getAppState) | **GET** /api/apps/{id}/state | Get State of an App
[**getWikiApp()**](AppsApi.md#getWikiApp) | **GET** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**listApps()**](AppsApi.md#listApps) | **GET** /api/apps | Get All Apps
[**prepareAppApply()**](AppsApi.md#prepareAppApply) | **GET** /api/apps/{id}/prepare-apply | Prepare To Apply an App
[**refreshAppState()**](AppsApi.md#refreshAppState) | **POST** /api/apps/{id}/refresh | Refresh State of an App
[**removeAppInstance()**](AppsApi.md#removeAppInstance) | **POST** /api/apps/{id}/remove-instance | Remove Instance from App
[**setAppSecurityGroups()**](AppsApi.md#setAppSecurityGroups) | **POST** /api/apps/{id}/security-groups | Set Security Groups for an App
[**updateApp()**](AppsApi.md#updateApp) | **PUT** /api/apps/{id} | Updating an App
[**updateWikiApp()**](AppsApi.md#updateWikiApp) | **PUT** /api/apps/{id}/wiki | Update an App Wiki Page
[**validateAppState()**](AppsApi.md#validateAppState) | **POST** /api/apps/{id}/validate-apply | Validate Apply State for an App


## `addAppInstance()`

```php
addAppInstance($id, $inline_object3): \OpenAPI\Client\Model\InlineResponse2003
```

Add Existing Instance to App

Add Existing Instance to App

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object3 = new \OpenAPI\Client\Model\InlineObject3(); // \OpenAPI\Client\Model\InlineObject3

try {
    $result = $apiInstance->addAppInstance($id, $inline_object3);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->addAppInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object3** | [**\OpenAPI\Client\Model\InlineObject3**](../Model/InlineObject3.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2003**](../Model/InlineResponse2003.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addAppUndoDelete()`

```php
addAppUndoDelete($id): \OpenAPI\Client\Model\InlineResponse2003
```

Undo Delete of an App

This operation will undo the delete of an app that is pending removal.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->addAppUndoDelete($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->addAppUndoDelete: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse2003**](../Model/InlineResponse2003.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addApps()`

```php
addApps($app_create): \OpenAPI\Client\Model\InlineResponse2002
```

Create an App

Create an App

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$app_create = {"blueprintId":"existing","name":"sample","description":"A sample app","group":{"id":1},"defaultCloud":{"id":19}}; // \OpenAPI\Client\Model\AppCreate

try {
    $result = $apiInstance->addApps($app_create);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->addApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_create** | [**\OpenAPI\Client\Model\AppCreate**](../Model/AppCreate.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2002**](../Model/InlineResponse2002.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `applyAppState()`

```php
applyAppState($id, $inline_object4): \OpenAPI\Client\Model\Model200Success
```

Apply State of an App

This endpoint provides a way to apply the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object4 = new \OpenAPI\Client\Model\InlineObject4(); // \OpenAPI\Client\Model\InlineObject4

try {
    $result = $apiInstance->applyAppState($id, $inline_object4);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->applyAppState: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object4** | [**\OpenAPI\Client\Model\InlineObject4**](../Model/InlineObject4.md)|  | [optional]

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

## `deleteApp()`

```php
deleteApp($id, $remove_instances, $preserve_volumes, $keep_backups, $release_floating_ips, $release_ei_ps, $force): \OpenAPI\Client\Model\Model200Success
```

Delete an App

Will delete an app. Use removeInstances=on to also delete the instances in the app and all associated monitors and backups.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$remove_instances = on; // string | Remove Instances
$preserve_volumes = on; // string | Preserve Volumes
$keep_backups = on; // string | Preserve copy of backups
$release_floating_ips = off; // string | Release Floating IPs
$release_ei_ps = off; // string | Alias for releaseFloatingIps
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteApp($id, $remove_instances, $preserve_volumes, $keep_backups, $release_floating_ips, $release_ei_ps, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->deleteApp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_instances** | **string**| Remove Instances | [optional] [default to &#39;off&#39;]
 **preserve_volumes** | **string**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **keep_backups** | **string**| Preserve copy of backups | [optional] [default to &#39;off&#39;]
 **release_floating_ips** | **string**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **release_ei_ps** | **string**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
 **force** | **string**| Force Delete | [optional] [default to &#39;off&#39;]

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

## `getApp()`

```php
getApp($id): \OpenAPI\Client\Model\InlineResponse2003
```

Get a Specific App

This endpoint retrieves a specific app.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getApp($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->getApp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse2003**](../Model/InlineResponse2003.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getAppSecurityGroups()`

```php
getAppSecurityGroups($id): object
```

Get Security Groups for an App

This returns a list of all of the security groups applied to an app and whether the firewall is enabled.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getAppSecurityGroups($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->getAppSecurityGroups: ', $e->getMessage(), PHP_EOL;
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

## `getAppState()`

```php
getAppState($id): AppState
```

Get State of an App

This endpoint provides a way to view the state of an app. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getAppState($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->getAppState: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AppState**](../Model/AppState.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getWikiApp()`

```php
getWikiApp($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves an App Wiki Page

This endpoint retrieves an app Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWikiApp($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->getWikiApp: ', $e->getMessage(), PHP_EOL;
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

## `listApps()`

```php
listApps($max, $offset, $name, $phrase, $created_by, $show_deleted, $labels, $all_labels): object
```

Get All Apps

This endpoint retrieves a paginated list of apps.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$created_by = 63; // int | The User ID for Filtering
$show_deleted = true; // bool | If true, includes instances in pending removal status.
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listApps($max, $offset, $name, $phrase, $created_by, $show_deleted, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->listApps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **created_by** | **int**| The User ID for Filtering | [optional]
 **show_deleted** | **bool**| If true, includes instances in pending removal status. | [optional] [default to false]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]

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

## `prepareAppApply()`

```php
prepareAppApply($id): AppPrepareApply
```

Prepare To Apply an App

This endpoint provides a way to view the current app configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->prepareAppApply($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->prepareAppApply: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AppPrepareApply**](../Model/AppPrepareApply.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `refreshAppState()`

```php
refreshAppState($id, $body): \OpenAPI\Client\Model\Model200Success
```

Refresh State of an App

This endpoint provides a way to refresh the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$body = new \stdClass; // object

try {
    $result = $apiInstance->refreshAppState($id, $body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->refreshAppState: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **body** | **object**|  | [optional]

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

## `removeAppInstance()`

```php
removeAppInstance($id, $inline_object5): \OpenAPI\Client\Model\InlineResponse2003
```

Remove Instance from App

Remove Instance from App

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object5 = new \OpenAPI\Client\Model\InlineObject5(); // \OpenAPI\Client\Model\InlineObject5

try {
    $result = $apiInstance->removeAppInstance($id, $inline_object5);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->removeAppInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object5** | [**\OpenAPI\Client\Model\InlineObject5**](../Model/InlineObject5.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2003**](../Model/InlineResponse2003.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `setAppSecurityGroups()`

```php
setAppSecurityGroups($id, $unknown_base_type): object
```

Set Security Groups for an App

Set Security Groups for an App

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = {"securityGroupIds":[19,2]}; // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->setAppSecurityGroups($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->setAppSecurityGroups: ', $e->getMessage(), PHP_EOL;
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

## `updateApp()`

```php
updateApp($id, $app_update): \OpenAPI\Client\Model\InlineResponse2003
```

Updating an App

This endpoint provides updating of some basic app settings.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$app_update = {"name":"My Sample App","description":"A new description of this app"}; // \OpenAPI\Client\Model\AppUpdate

try {
    $result = $apiInstance->updateApp($id, $app_update);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->updateApp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **app_update** | [**\OpenAPI\Client\Model\AppUpdate**](../Model/AppUpdate.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2003**](../Model/InlineResponse2003.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateWikiApp()`

```php
updateWikiApp($id, $inline_object267): object
```

Update an App Wiki Page

Updates an app Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object267 = new \OpenAPI\Client\Model\InlineObject267(); // \OpenAPI\Client\Model\InlineObject267

try {
    $result = $apiInstance->updateWikiApp($id, $inline_object267);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->updateWikiApp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object267** | [**\OpenAPI\Client\Model\InlineObject267**](../Model/InlineObject267.md)|  | [optional]

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

## `validateAppState()`

```php
validateAppState($id, $inline_object6): object
```

Validate Apply State for an App

This endpoint provides a way to validate app configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AppsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object6 = new \OpenAPI\Client\Model\InlineObject6(); // \OpenAPI\Client\Model\InlineObject6

try {
    $result = $apiInstance->validateAppState($id, $inline_object6);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AppsApi->validateAppState: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object6** | [**\OpenAPI\Client\Model\InlineObject6**](../Model/InlineObject6.md)|  | [optional]

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
