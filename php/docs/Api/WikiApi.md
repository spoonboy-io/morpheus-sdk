# OpenAPI\Client\WikiApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addWiki()**](WikiApi.md#addWiki) | **POST** /api/wiki/pages | Create a Wiki Page
[**getWiki()**](WikiApi.md#getWiki) | **GET** /api/wiki/pages/{id} | Retrieves a specific Wiki page
[**getWikiApp()**](WikiApi.md#getWikiApp) | **GET** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**getWikiCategories()**](WikiApi.md#getWikiCategories) | **GET** /api/wiki/categories | Retrieves all Wiki categories associated with the account
[**getWikiCloud()**](WikiApi.md#getWikiCloud) | **GET** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**getWikiCluster()**](WikiApi.md#getWikiCluster) | **GET** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**getWikiGroup()**](WikiApi.md#getWikiGroup) | **GET** /api/groups/{id}/wiki | Retrieves a Group Wiki Page
[**getWikiInstance()**](WikiApi.md#getWikiInstance) | **GET** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**getWikiServer()**](WikiApi.md#getWikiServer) | **GET** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**listWiki()**](WikiApi.md#listWiki) | **GET** /api/wiki/pages | Retrieves wiki pages associated with the account.
[**removeWiki()**](WikiApi.md#removeWiki) | **DELETE** /api/wiki/pages/{id} | Deletes a Wiki Page
[**updateWiki()**](WikiApi.md#updateWiki) | **PUT** /api/wiki/pages/{id} | Updates a Wiki Page
[**updateWikiApp()**](WikiApi.md#updateWikiApp) | **PUT** /api/apps/{id}/wiki | Update an App Wiki Page
[**updateWikiCloud()**](WikiApi.md#updateWikiCloud) | **PUT** /api/zones/{id}/wiki | Update a Cloud Wiki Page
[**updateWikiCluster()**](WikiApi.md#updateWikiCluster) | **PUT** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page
[**updateWikiGroup()**](WikiApi.md#updateWikiGroup) | **PUT** /api/groups/{id}/wiki | Update a Group Wiki Page
[**updateWikiInstance()**](WikiApi.md#updateWikiInstance) | **PUT** /api/instances/{id}/wiki | Update an Instance Wiki Page
[**updateWikiServer()**](WikiApi.md#updateWikiServer) | **PUT** /api/servers/{id}/wiki | Update a Server Wiki Page


## `addWiki()`

```php
addWiki($inline_object272): object
```

Create a Wiki Page

Creates a Wiki Page

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object272 = new \OpenAPI\Client\Model\InlineObject272(); // \OpenAPI\Client\Model\InlineObject272

try {
    $result = $apiInstance->addWiki($inline_object272);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->addWiki: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object272** | [**\OpenAPI\Client\Model\InlineObject272**](../Model/InlineObject272.md)|  | [optional]

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

## `getWiki()`

```php
getWiki($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves a specific Wiki page

This endpoint retrieves a specific wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWiki($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->getWiki: ', $e->getMessage(), PHP_EOL;
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


$apiInstance = new OpenAPI\Client\Api\WikiApi(
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
    echo 'Exception when calling WikiApi->getWikiApp: ', $e->getMessage(), PHP_EOL;
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

## `getWikiCategories()`

```php
getWikiCategories($phrase, $page_phrase): \OpenAPI\Client\Model\InlineResponse200169
```

Retrieves all Wiki categories associated with the account

This endpoint retrieves all Wiki categories associated with the account. The results are not paginated. The categories returned are those of the found pages.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$page_phrase = 'page_phrase_example'; // string | If specified will return a partial match on page name

try {
    $result = $apiInstance->getWikiCategories($phrase, $page_phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->getWikiCategories: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **page_phrase** | **string**| If specified will return a partial match on page name | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200169**](../Model/InlineResponse200169.md)

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


$apiInstance = new OpenAPI\Client\Api\WikiApi(
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
    echo 'Exception when calling WikiApi->getWikiCloud: ', $e->getMessage(), PHP_EOL;
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

## `getWikiCluster()`

```php
getWikiCluster($cluster_id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves a Cluster Wiki Page

This endpoint retrieves a cluster Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getWikiCluster($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->getWikiCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

## `getWikiGroup()`

```php
getWikiGroup($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves a Group Wiki Page

This endpoint retrieves a group Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWikiGroup($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->getWikiGroup: ', $e->getMessage(), PHP_EOL;
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

## `getWikiInstance()`

```php
getWikiInstance($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves an Instance Wiki Page

This endpoint retrieves an instance Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWikiInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->getWikiInstance: ', $e->getMessage(), PHP_EOL;
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

## `getWikiServer()`

```php
getWikiServer($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves a Server Wiki Page

This endpoint retrieves a server Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWikiServer($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->getWikiServer: ', $e->getMessage(), PHP_EOL;
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

## `listWiki()`

```php
listWiki($name, $phrase): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves wiki pages associated with the account.

This endpoint retrieves wiki pages associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listWiki($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->listWiki: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

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

## `removeWiki()`

```php
removeWiki($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Wiki Page

Deletes the specified Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeWiki($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->removeWiki: ', $e->getMessage(), PHP_EOL;
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

## `updateWiki()`

```php
updateWiki($id, $inline_object273): object
```

Updates a Wiki Page

Updates a Wiki Page

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object273 = new \OpenAPI\Client\Model\InlineObject273(); // \OpenAPI\Client\Model\InlineObject273

try {
    $result = $apiInstance->updateWiki($id, $inline_object273);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->updateWiki: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object273** | [**\OpenAPI\Client\Model\InlineObject273**](../Model/InlineObject273.md)|  | [optional]

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


$apiInstance = new OpenAPI\Client\Api\WikiApi(
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
    echo 'Exception when calling WikiApi->updateWikiApp: ', $e->getMessage(), PHP_EOL;
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


$apiInstance = new OpenAPI\Client\Api\WikiApi(
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
    echo 'Exception when calling WikiApi->updateWikiCloud: ', $e->getMessage(), PHP_EOL;
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

## `updateWikiCluster()`

```php
updateWikiCluster($cluster_id, $inline_object268): object
```

Update a Cluster Wiki Page

Updates a cluster Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object268 = new \OpenAPI\Client\Model\InlineObject268(); // \OpenAPI\Client\Model\InlineObject268

try {
    $result = $apiInstance->updateWikiCluster($cluster_id, $inline_object268);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->updateWikiCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object268** | [**\OpenAPI\Client\Model\InlineObject268**](../Model/InlineObject268.md)|  | [optional]

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

## `updateWikiGroup()`

```php
updateWikiGroup($id, $inline_object269): object
```

Update a Group Wiki Page

Updates a group Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object269 = new \OpenAPI\Client\Model\InlineObject269(); // \OpenAPI\Client\Model\InlineObject269

try {
    $result = $apiInstance->updateWikiGroup($id, $inline_object269);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->updateWikiGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object269** | [**\OpenAPI\Client\Model\InlineObject269**](../Model/InlineObject269.md)|  | [optional]

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

## `updateWikiInstance()`

```php
updateWikiInstance($id, $inline_object270): object
```

Update an Instance Wiki Page

Updates an instance Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object270 = new \OpenAPI\Client\Model\InlineObject270(); // \OpenAPI\Client\Model\InlineObject270

try {
    $result = $apiInstance->updateWikiInstance($id, $inline_object270);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->updateWikiInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object270** | [**\OpenAPI\Client\Model\InlineObject270**](../Model/InlineObject270.md)|  | [optional]

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

## `updateWikiServer()`

```php
updateWikiServer($id, $inline_object271): object
```

Update a Server Wiki Page

Updates a server Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\WikiApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object271 = new \OpenAPI\Client\Model\InlineObject271(); // \OpenAPI\Client\Model\InlineObject271

try {
    $result = $apiInstance->updateWikiServer($id, $inline_object271);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WikiApi->updateWikiServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object271** | [**\OpenAPI\Client\Model\InlineObject271**](../Model/InlineObject271.md)|  | [optional]

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
