# OpenAPI\Client\DeploysApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addInstanceDeploy()**](DeploysApi.md#addInstanceDeploy) | **POST** /api/instances/{id}/deploys | Deploy to an Instance
[**deletedeploy()**](DeploysApi.md#deletedeploy) | **DELETE** /api/deploys/{id} | Delete a Deploy
[**getInstanceDeploys()**](DeploysApi.md#getInstanceDeploys) | **GET** /api/instances/{id}/deploys | Get all Deploys for an Instance
[**listDeploys()**](DeploysApi.md#listDeploys) | **GET** /api/deploys | Get all Deploys
[**runDeploy()**](DeploysApi.md#runDeploy) | **POST** /api/deploys/{id}/deploy | Run a Deploy
[**updateDeploy()**](DeploysApi.md#updateDeploy) | **PUT** /api/deploys/{id} | Update a Deploy


## `addInstanceDeploy()`

```php
addInstanceDeploy($id, $inline_object92): \OpenAPI\Client\Model\InlineResponse20040
```

Deploy to an Instance

This endpoint will deploy the specified deployment version to specified instance. The version to deploy can be identified with deploymentId and version or with versionId alone.  By default, the deployment is executed right away. To prevent this so that it can be run manually later on.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploysApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object92 = new \OpenAPI\Client\Model\InlineObject92(); // \OpenAPI\Client\Model\InlineObject92

try {
    $result = $apiInstance->addInstanceDeploy($id, $inline_object92);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploysApi->addInstanceDeploy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object92** | [**\OpenAPI\Client\Model\InlineObject92**](../Model/InlineObject92.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20040**](../Model/InlineResponse20040.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deletedeploy()`

```php
deletedeploy($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Deploy

This endpoint will delete an archived instance deploy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploysApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deletedeploy($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploysApi->deletedeploy: ', $e->getMessage(), PHP_EOL;
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

## `getInstanceDeploys()`

```php
getInstanceDeploys($id, $max, $offset, $phrase, $name, $deployment_id, $instance_name, $instance_id, $version, $version_id, $created_by_id, $deploy_type, $date_created, $last_updated, $deploy_date, $status): object
```

Get all Deploys for an Instance

This endpoint retrieves all deploys for a specific instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploysApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$deployment_id = 5; // int | Filter by deployment id
$instance_name = instance1; // string | Filter by instance name
$instance_id = 94; // int | The Instance ID for Filtering
$version = 5; // int | Filter by version number (userVersion)
$version_id = 5; // int | Filter by deployment version id
$created_by_id = 63; // int | Filter by owner (user) id
$deploy_type = file; // string | Filter by type (deployType), file, git, fetch
$date_created = 2019-01-01; // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$deploy_date = 2020-01-01; // string | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified
$status = deploying; // string | Filter by status

try {
    $result = $apiInstance->getInstanceDeploys($id, $max, $offset, $phrase, $name, $deployment_id, $instance_name, $instance_id, $version, $version_id, $created_by_id, $deploy_type, $date_created, $last_updated, $deploy_date, $status);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploysApi->getInstanceDeploys: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **deployment_id** | **int**| Filter by deployment id | [optional]
 **instance_name** | **string**| Filter by instance name | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **version_id** | **int**| Filter by deployment version id | [optional]
 **created_by_id** | **int**| Filter by owner (user) id | [optional]
 **deploy_type** | **string**| Filter by type (deployType), file, git, fetch | [optional]
 **date_created** | **string**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deploy_date** | **string**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional]
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

## `listDeploys()`

```php
listDeploys($max, $offset, $phrase, $name, $deployment_id, $instance_name, $instance_id, $version, $version_id, $created_by_id, $deploy_type, $date_created, $last_updated, $deploy_date, $status): object
```

Get all Deploys

This endpoint retrieves all deploys.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploysApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$deployment_id = 5; // int | Filter by deployment id
$instance_name = instance1; // string | Filter by instance name
$instance_id = 94; // int | The Instance ID for Filtering
$version = 5; // int | Filter by version number (userVersion)
$version_id = 5; // int | Filter by deployment version id
$created_by_id = 63; // int | Filter by owner (user) id
$deploy_type = file; // string | Filter by type (deployType), file, git, fetch
$date_created = 2019-01-01; // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$deploy_date = 2020-01-01; // string | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified
$status = deploying; // string | Filter by status

try {
    $result = $apiInstance->listDeploys($max, $offset, $phrase, $name, $deployment_id, $instance_name, $instance_id, $version, $version_id, $created_by_id, $deploy_type, $date_created, $last_updated, $deploy_date, $status);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploysApi->listDeploys: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **deployment_id** | **int**| Filter by deployment id | [optional]
 **instance_name** | **string**| Filter by instance name | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **version_id** | **int**| Filter by deployment version id | [optional]
 **created_by_id** | **int**| Filter by owner (user) id | [optional]
 **deploy_type** | **string**| Filter by type (deployType), file, git, fetch | [optional]
 **date_created** | **string**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deploy_date** | **string**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional]
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

## `runDeploy()`

```php
runDeploy($id, $inline_object73): \OpenAPI\Client\Model\InlineResponse20040
```

Run a Deploy

This endpoint will run an existing instance deploy. This is for running a new staged deploy or to rollback to previous version by re-running a deploy that is archived.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploysApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object73 = new \OpenAPI\Client\Model\InlineObject73(); // \OpenAPI\Client\Model\InlineObject73

try {
    $result = $apiInstance->runDeploy($id, $inline_object73);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploysApi->runDeploy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object73** | [**\OpenAPI\Client\Model\InlineObject73**](../Model/InlineObject73.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20040**](../Model/InlineResponse20040.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateDeploy()`

```php
updateDeploy($id, $inline_object72): \OpenAPI\Client\Model\InlineResponse20040
```

Update a Deploy

This endpoint will update an existing deploy. This is typically only needed to change settings on a deploy that is staged, before it is run.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\DeploysApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object72 = new \OpenAPI\Client\Model\InlineObject72(); // \OpenAPI\Client\Model\InlineObject72

try {
    $result = $apiInstance->updateDeploy($id, $inline_object72);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DeploysApi->updateDeploy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object72** | [**\OpenAPI\Client\Model\InlineObject72**](../Model/InlineObject72.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20040**](../Model/InlineResponse20040.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
