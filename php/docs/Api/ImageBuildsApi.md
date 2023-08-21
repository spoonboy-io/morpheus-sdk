# OpenAPI\Client\ImageBuildsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addBootScript()**](ImageBuildsApi.md#addBootScript) | **POST** /api/boot-scripts | Create a Boot Script
[**addImageBuild()**](ImageBuildsApi.md#addImageBuild) | **POST** /api/image-builds | Create an Image Build
[**addPreseedScript()**](ImageBuildsApi.md#addPreseedScript) | **POST** /api/preseed-scripts | Create a Preseed Script
[**deleteBootScript()**](ImageBuildsApi.md#deleteBootScript) | **DELETE** /api/boot-scripts/{id} | Delete a Boot Script
[**deleteImageBuild()**](ImageBuildsApi.md#deleteImageBuild) | **DELETE** /api/image-builds/{id} | Delete an Image Build
[**deletePreseedScript()**](ImageBuildsApi.md#deletePreseedScript) | **DELETE** /api/preseed-scripts/{id} | Delete a Preseed Script
[**executeImageBuild()**](ImageBuildsApi.md#executeImageBuild) | **POST** /api/image-builds/{id}/run | Run an Image Build
[**getBootScript()**](ImageBuildsApi.md#getBootScript) | **GET** /api/boot-scripts/{id} | Get a Specific Boot Script
[**getImageBuild()**](ImageBuildsApi.md#getImageBuild) | **GET** /api/image-builds/{id} | Get a Specific Image Build
[**getImageBuildExecutions()**](ImageBuildsApi.md#getImageBuildExecutions) | **GET** /api/image-builds/{id}/list-executions | List Image Build Executions
[**getPreseedScript()**](ImageBuildsApi.md#getPreseedScript) | **GET** /api/preseed-scripts/{id} | Get a Specific Preseed Script
[**listBootScripts()**](ImageBuildsApi.md#listBootScripts) | **GET** /api/boot-scripts | Boot Scripts
[**listImageBuilds()**](ImageBuildsApi.md#listImageBuilds) | **GET** /api/image-builds | Get All Image Builds
[**listPreseedScripts()**](ImageBuildsApi.md#listPreseedScripts) | **GET** /api/preseed-scripts | Preseed Scripts
[**updateBootScript()**](ImageBuildsApi.md#updateBootScript) | **PUT** /api/boot-scripts/{id} | Update a Boot Script
[**updateImageBuild()**](ImageBuildsApi.md#updateImageBuild) | **PUT** /api/image-builds/{id} | Update an Image Build
[**updatePreseedScript()**](ImageBuildsApi.md#updatePreseedScript) | **PUT** /api/preseed-scripts/{id} | Update a Preseed Script


## `addBootScript()`

```php
addBootScript($inline_object83): object
```

Create a Boot Script

Create a Boot Script

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object83 = new \OpenAPI\Client\Model\InlineObject83(); // \OpenAPI\Client\Model\InlineObject83

try {
    $result = $apiInstance->addBootScript($inline_object83);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->addBootScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object83** | [**\OpenAPI\Client\Model\InlineObject83**](../Model/InlineObject83.md)|  | [optional]

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

## `addImageBuild()`

```php
addImageBuild($inline_object85): object
```

Create an Image Build

Create an Image Build

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object85 = new \OpenAPI\Client\Model\InlineObject85(); // \OpenAPI\Client\Model\InlineObject85

try {
    $result = $apiInstance->addImageBuild($inline_object85);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->addImageBuild: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object85** | [**\OpenAPI\Client\Model\InlineObject85**](../Model/InlineObject85.md)|  | [optional]

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

## `addPreseedScript()`

```php
addPreseedScript($inline_object196): object
```

Create a Preseed Script

Create a Preseed Script

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object196 = new \OpenAPI\Client\Model\InlineObject196(); // \OpenAPI\Client\Model\InlineObject196

try {
    $result = $apiInstance->addPreseedScript($inline_object196);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->addPreseedScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object196** | [**\OpenAPI\Client\Model\InlineObject196**](../Model/InlineObject196.md)|  | [optional]

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

## `deleteBootScript()`

```php
deleteBootScript($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Boot Script

Will delete a boot script from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteBootScript($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->deleteBootScript: ', $e->getMessage(), PHP_EOL;
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

## `deleteImageBuild()`

```php
deleteImageBuild($id): \OpenAPI\Client\Model\Model200Success
```

Delete an Image Build

Will delete an image build from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteImageBuild($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->deleteImageBuild: ', $e->getMessage(), PHP_EOL;
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

## `deletePreseedScript()`

```php
deletePreseedScript($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Preseed Script

Will delete a preseed script from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deletePreseedScript($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->deletePreseedScript: ', $e->getMessage(), PHP_EOL;
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

## `executeImageBuild()`

```php
executeImageBuild($id): \OpenAPI\Client\Model\Model200Success
```

Run an Image Build

Running an image build is done asynchronously. This api will kick off the new execution and update the image build status.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->executeImageBuild($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->executeImageBuild: ', $e->getMessage(), PHP_EOL;
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

## `getBootScript()`

```php
getBootScript($id): \OpenAPI\Client\Model\InlineResponse20052
```

Get a Specific Boot Script

This endpoint retrieves a specific boot script.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getBootScript($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->getBootScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20052**](../Model/InlineResponse20052.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getImageBuild()`

```php
getImageBuild($id): \OpenAPI\Client\Model\InlineResponse20053
```

Get a Specific Image Build

This endpoint retrieves a specific image build.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getImageBuild($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->getImageBuild: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20053**](../Model/InlineResponse20053.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getImageBuildExecutions()`

```php
getImageBuildExecutions($id, $build_number, $status): object
```

List Image Build Executions

List all executions for an image build. This same info is also returned by the image build GET api, which returns the 100 most recent executions.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$build_number = 4; // int | If specified will return an exact match on buildNumber
$status = running; // string | Filter by status

try {
    $result = $apiInstance->getImageBuildExecutions($id, $build_number, $status);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->getImageBuildExecutions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **build_number** | **int**| If specified will return an exact match on buildNumber | [optional]
 **status** | **string**| Filter by status | [optional]

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

## `getPreseedScript()`

```php
getPreseedScript($id): \OpenAPI\Client\Model\InlineResponse200122
```

Get a Specific Preseed Script

This endpoint retrieves a specific preseed script.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPreseedScript($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->getPreseedScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200122**](../Model/InlineResponse200122.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listBootScripts()`

```php
listBootScripts($name, $phrase): object
```

Boot Scripts

Boot Scripts are used in the Image Builder service. See Image Builds.  This endpoint retrieves all boot scripts associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listBootScripts($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->listBootScripts: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

## `listImageBuilds()`

```php
listImageBuilds($name, $phrase): object
```

Get All Image Builds

This endpoint retrieves all image builds associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listImageBuilds($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->listImageBuilds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

## `listPreseedScripts()`

```php
listPreseedScripts($name, $phrase): object
```

Preseed Scripts

Preseed Scripts are used in the Image Builder service. See Image Builds.  This endpoint retrieves all preseed scripts associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listPreseedScripts($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->listPreseedScripts: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

## `updateBootScript()`

```php
updateBootScript($id, $inline_object84): object
```

Update a Boot Script

Update a Boot Script

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object84 = new \OpenAPI\Client\Model\InlineObject84(); // \OpenAPI\Client\Model\InlineObject84

try {
    $result = $apiInstance->updateBootScript($id, $inline_object84);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->updateBootScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object84** | [**\OpenAPI\Client\Model\InlineObject84**](../Model/InlineObject84.md)|  | [optional]

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

## `updateImageBuild()`

```php
updateImageBuild($id, $inline_object86): \OpenAPI\Client\Model\InlineResponse20054
```

Update an Image Build

Update an Image Build

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object86 = new \OpenAPI\Client\Model\InlineObject86(); // \OpenAPI\Client\Model\InlineObject86

try {
    $result = $apiInstance->updateImageBuild($id, $inline_object86);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->updateImageBuild: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object86** | [**\OpenAPI\Client\Model\InlineObject86**](../Model/InlineObject86.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20054**](../Model/InlineResponse20054.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updatePreseedScript()`

```php
updatePreseedScript($id, $inline_object197): object
```

Update a Preseed Script

Update a Preseed Script

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ImageBuildsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object197 = new \OpenAPI\Client\Model\InlineObject197(); // \OpenAPI\Client\Model\InlineObject197

try {
    $result = $apiInstance->updatePreseedScript($id, $inline_object197);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ImageBuildsApi->updatePreseedScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object197** | [**\OpenAPI\Client\Model\InlineObject197**](../Model/InlineObject197.md)|  | [optional]

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
