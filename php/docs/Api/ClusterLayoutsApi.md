# OpenAPI\Client\ClusterLayoutsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addClusterLayoutClone()**](ClusterLayoutsApi.md#addClusterLayoutClone) | **POST** /api/library/cluster-layouts/{id}/clone | Clone a Cluster Layout
[**addClusterLayouts()**](ClusterLayoutsApi.md#addClusterLayouts) | **POST** /api/library/cluster-layouts | Create a Cluster Layout
[**deleteClusterLayout()**](ClusterLayoutsApi.md#deleteClusterLayout) | **DELETE** /api/library/cluster-layouts/{id} | Delete a Cluster Layout
[**getClusterLayout()**](ClusterLayoutsApi.md#getClusterLayout) | **GET** /api/library/cluster-layouts/{id} | Get a Specific Cluster Layout
[**listClusterLayouts()**](ClusterLayoutsApi.md#listClusterLayouts) | **GET** /api/library/cluster-layouts | Get All Cluster Layouts
[**updateClusterLayout()**](ClusterLayoutsApi.md#updateClusterLayout) | **PUT** /api/library/cluster-layouts/{id} | Update a Cluster Layout


## `addClusterLayoutClone()`

```php
addClusterLayoutClone($id, $name, $description, $compute_version): \OpenAPI\Client\Model\SuccessId
```

Clone a Cluster Layout

Use this command to clone a cluster layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClusterLayoutsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$name = New Name; // string | Name of cluster layout. Defaults to Copy of <cloned layout name>
$description = New Description; // string | Description of cluster layout. Defaults to cloned layout description
$compute_version = 2.0; // string | Version of cluster layout. Defaults to cloned layout version

try {
    $result = $apiInstance->addClusterLayoutClone($id, $name, $description, $compute_version);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClusterLayoutsApi->addClusterLayoutClone: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **name** | **string**| Name of cluster layout. Defaults to Copy of &lt;cloned layout name&gt; | [optional]
 **description** | **string**| Description of cluster layout. Defaults to cloned layout description | [optional]
 **compute_version** | **string**| Version of cluster layout. Defaults to cloned layout version | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addClusterLayouts()`

```php
addClusterLayouts($inline_object50): \OpenAPI\Client\Model\SuccessId
```

Create a Cluster Layout

Use this command to create a cluster layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClusterLayoutsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object50 = new \OpenAPI\Client\Model\InlineObject50(); // \OpenAPI\Client\Model\InlineObject50

try {
    $result = $apiInstance->addClusterLayouts($inline_object50);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClusterLayoutsApi->addClusterLayouts: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object50** | [**\OpenAPI\Client\Model\InlineObject50**](../Model/InlineObject50.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteClusterLayout()`

```php
deleteClusterLayout($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Cluster Layout

Will delete a cluster layout

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClusterLayoutsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteClusterLayout($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClusterLayoutsApi->deleteClusterLayout: ', $e->getMessage(), PHP_EOL;
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

## `getClusterLayout()`

```php
getClusterLayout($id): \OpenAPI\Client\Model\InlineResponse20025
```

Get a Specific Cluster Layout

This endpoint retrieves a specific cluster layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClusterLayoutsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getClusterLayout($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClusterLayoutsApi->getClusterLayout: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20025**](../Model/InlineResponse20025.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listClusterLayouts()`

```php
listClusterLayouts($max, $offset, $sort, $direction, $phrase, $provision_type, $labels, $all_labels): object
```

Get All Cluster Layouts

This endpoint retrieves all cluster layouts.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClusterLayoutsApi(
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
$provision_type = 'provision_type_example'; // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listClusterLayouts($max, $offset, $sort, $direction, $phrase, $provision_type, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClusterLayoutsApi->listClusterLayouts: ', $e->getMessage(), PHP_EOL;
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
 **provision_type** | **string**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings. | [optional]
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

## `updateClusterLayout()`

```php
updateClusterLayout($id, $inline_object51): \OpenAPI\Client\Model\SuccessId
```

Update a Cluster Layout

Use this command to update an existing cluster layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClusterLayoutsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object51 = new \OpenAPI\Client\Model\InlineObject51(); // \OpenAPI\Client\Model\InlineObject51

try {
    $result = $apiInstance->updateClusterLayout($id, $inline_object51);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClusterLayoutsApi->updateClusterLayout: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object51** | [**\OpenAPI\Client\Model\InlineObject51**](../Model/InlineObject51.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
