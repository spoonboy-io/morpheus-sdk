# OpenAPI\Client\DeploymentsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addDeploymentFile()**](DeploymentsApi.md#addDeploymentFile) | **POST** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Upload a Deployment File
[**addDeploymentVersion()**](DeploymentsApi.md#addDeploymentVersion) | **POST** /api/deployments/{deploymentId}/versions | Create a new Deployment Version
[**addDeployments()**](DeploymentsApi.md#addDeployments) | **POST** /api/deployments | Create a new Deployment
[**deleteDeployment()**](DeploymentsApi.md#deleteDeployment) | **DELETE** /api/deployments/{deploymentId} | Delete a Deployment
[**deleteDeploymentFile()**](DeploymentsApi.md#deleteDeploymentFile) | **DELETE** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Delete a Deployment File
[**deleteDeploymentVersion()**](DeploymentsApi.md#deleteDeploymentVersion) | **DELETE** /api/deployments/{deploymentId}/versions/{id} | Delete a Deployment Version
[**getDeployment()**](DeploymentsApi.md#getDeployment) | **GET** /api/deployments/{deploymentId} | Get a Specific Deployment
[**getDeploymentVersion()**](DeploymentsApi.md#getDeploymentVersion) | **GET** /api/deployments/{deploymentId}/versions/{id} | Get a Specific Deployment Version
[**listDeploymentFiles()**](DeploymentsApi.md#listDeploymentFiles) | **GET** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | List Deployment Files
[**listDeploymentVersions()**](DeploymentsApi.md#listDeploymentVersions) | **GET** /api/deployments/{deploymentId}/versions | Get All Versions For a Deployment
[**listDeployments()**](DeploymentsApi.md#listDeployments) | **GET** /api/deployments | Get All Deployments
[**updateDeployment()**](DeploymentsApi.md#updateDeployment) | **PUT** /api/deployments/{deploymentId} | Updating a Deployment
[**updateDeploymentVersion()**](DeploymentsApi.md#updateDeploymentVersion) | **PUT** /api/deployments/{deploymentId}/versions/{id} | Updating a Deployment Version


## `addDeploymentFile()`

```php
addDeploymentFile($deployment_id, $id, $filepath, $file): \OpenAPI\Client\Model\Model200Success
```

Upload a Deployment File

This endpoint will upload a file for a specific deployment version. This will overwrite the file if one with the same name exists already.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$id = 1; // int | Morpheus ID of the Object being referenced
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.
$file = "/path/to/file.txt"; // \SplFileObject

try {
    $result = $apiInstance->addDeploymentFile($deployment_id, $id, $filepath, $file);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->addDeploymentFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **file** | **\SplFileObject****\SplFileObject**|  | [optional]

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

## `addDeploymentVersion()`

```php
addDeploymentVersion($deployment_id, $inline_object69): object
```

Create a new Deployment Version

This endpoint will create a new deployment version that is ready to have files uploaded to it. The default type is file, which has files directly uploaded via Morpheus. Alternatively, the type git or fetch can be used to just point to a repository or remote url.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$inline_object69 = new \OpenAPI\Client\Model\InlineObject69(); // \OpenAPI\Client\Model\InlineObject69

try {
    $result = $apiInstance->addDeploymentVersion($deployment_id, $inline_object69);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->addDeploymentVersion: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **inline_object69** | [**\OpenAPI\Client\Model\InlineObject69**](../Model/InlineObject69.md)|  | [optional]

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

## `addDeployments()`

```php
addDeployments($inline_object67): object
```

Create a new Deployment

This endpoint will create a new deployment that is ready to have versions added to it.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object67 = new \OpenAPI\Client\Model\InlineObject67(); // \OpenAPI\Client\Model\InlineObject67

try {
    $result = $apiInstance->addDeployments($inline_object67);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->addDeployments: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object67** | [**\OpenAPI\Client\Model\InlineObject67**](../Model/InlineObject67.md)|  | [optional]

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

## `deleteDeployment()`

```php
deleteDeployment($deployment_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Deployment

This endpoint will delete an existing deployment.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID

try {
    $result = $apiInstance->deleteDeployment($deployment_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->deleteDeployment: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |

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

## `deleteDeploymentFile()`

```php
deleteDeploymentFile($deployment_id, $id, $filepath, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Deployment File

This endpoint will delete an existing deployment file. To recursively delete a directory and all of its contents, the force parameter must be specified.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$id = 1; // int | Morpheus ID of the Object being referenced
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteDeploymentFile($deployment_id, $id, $filepath, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->deleteDeploymentFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
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

## `deleteDeploymentVersion()`

```php
deleteDeploymentVersion($deployment_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Deployment Version

This endpoint will delete an existing deployment version.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteDeploymentVersion($deployment_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->deleteDeploymentVersion: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
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

## `getDeployment()`

```php
getDeployment($deployment_id, $max_versions): \OpenAPI\Client\Model\InlineResponse20038
```

Get a Specific Deployment

This endpoint retrieves a specific deployment. By default the 5 most recent versions are returned, more can be returned by specifying the maxVersions parameter.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$max_versions = 6; // int | Max number of recent versions to return.

try {
    $result = $apiInstance->getDeployment($deployment_id, $max_versions);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->getDeployment: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **max_versions** | **int**| Max number of recent versions to return. | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20038**](../Model/InlineResponse20038.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getDeploymentVersion()`

```php
getDeploymentVersion($deployment_id, $id): \OpenAPI\Client\Model\InlineResponse20039
```

Get a Specific Deployment Version

This endpoint retrieves a specific deployment version.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getDeploymentVersion($deployment_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->getDeploymentVersion: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20039**](../Model/InlineResponse20039.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listDeploymentFiles()`

```php
listDeploymentFiles($deployment_id, $id, $filepath, $max, $offset, $phrase, $version): object
```

List Deployment Files

This endpoint returns a list of files for a specific deployment version. This only applies to deploy type file. Files are sorted alphabetically, with directories appearing at the beginning of the list.  The filepath parameter can be specified to search for specific files or directories. To list files under a directory, use a trailing / in the filepath parameter.  To list a specific file, provide it's full path.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$id = 1; // int | Morpheus ID of the Object being referenced
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$version = 5; // int | Filter by version number (userVersion)

try {
    $result = $apiInstance->listDeploymentFiles($deployment_id, $id, $filepath, $max, $offset, $phrase, $version);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->listDeploymentFiles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]

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

## `listDeploymentVersions()`

```php
listDeploymentVersions($deployment_id, $max, $offset, $phrase, $version, $type, $date_created, $last_updated): object
```

Get All Versions For a Deployment

This endpoint returns a paginated list of versions for a specific deployment.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$version = 5; // int | Filter by version number (userVersion)
$type = file; // string | Filter by type (deployType), file, git, fetch
$date_created = 2019-01-01; // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)

try {
    $result = $apiInstance->listDeploymentVersions($deployment_id, $max, $offset, $phrase, $version, $type, $date_created, $last_updated);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->listDeploymentVersions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **type** | **string**| Filter by type (deployType), file, git, fetch | [optional]
 **date_created** | **string**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

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

## `listDeployments()`

```php
listDeployments($max, $offset, $phrase, $name, $description, $date_created, $last_updated): object
```

Get All Deployments

This endpoint returns a paginated list of deployments.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`
$date_created = 2019-01-01; // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)

try {
    $result = $apiInstance->listDeployments($max, $offset, $phrase, $name, $description, $date_created, $last_updated);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->listDeployments: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **date_created** | **string**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

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

## `updateDeployment()`

```php
updateDeployment($deployment_id, $inline_object68): object
```

Updating a Deployment

This endpoint will update an existing deployment.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$inline_object68 = new \OpenAPI\Client\Model\InlineObject68(); // \OpenAPI\Client\Model\InlineObject68

try {
    $result = $apiInstance->updateDeployment($deployment_id, $inline_object68);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->updateDeployment: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **inline_object68** | [**\OpenAPI\Client\Model\InlineObject68**](../Model/InlineObject68.md)|  | [optional]

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

## `updateDeploymentVersion()`

```php
updateDeploymentVersion($deployment_id, $id, $inline_object70): object
```

Updating a Deployment Version

This endpoint will update an existing deployment version.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploymentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$deployment_id = 4; // int | Deployment ID
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object70 = new \OpenAPI\Client\Model\InlineObject70(); // \OpenAPI\Client\Model\InlineObject70

try {
    $result = $apiInstance->updateDeploymentVersion($deployment_id, $id, $inline_object70);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploymentsApi->updateDeploymentVersion: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object70** | [**\OpenAPI\Client\Model\InlineObject70**](../Model/InlineObject70.md)|  | [optional]

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
