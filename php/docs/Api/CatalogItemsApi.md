# OpenAPI\Client\CatalogItemsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCatalogItemType()**](CatalogItemsApi.md#addCatalogItemType) | **POST** /api/catalog-item-types | Create a Catalog Item Type
[**getCatalogItemType()**](CatalogItemsApi.md#getCatalogItemType) | **GET** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type
[**listCatalogItemTypes()**](CatalogItemsApi.md#listCatalogItemTypes) | **GET** /api/catalog-item-types | Get All Catalog Item Types
[**removeCatalogItemType()**](CatalogItemsApi.md#removeCatalogItemType) | **DELETE** /api/catalog-item-types/{id} | Delete a Catalog Item Type
[**updateCatalogItemType()**](CatalogItemsApi.md#updateCatalogItemType) | **PUT** /api/catalog-item-types/{id} | Update a Catalog Item Type
[**updateCatalogItemTypeLogo()**](CatalogItemsApi.md#updateCatalogItemTypeLogo) | **PUT** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type


## `addCatalogItemType()`

```php
addCatalogItemType($inline_object24): object
```

Create a Catalog Item Type

Use this command to create a catalog item type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CatalogItemsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object24 = new \OpenAPI\Client\Model\InlineObject24(); // \OpenAPI\Client\Model\InlineObject24

try {
    $result = $apiInstance->addCatalogItemType($inline_object24);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CatalogItemsApi->addCatalogItemType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object24** | [**\OpenAPI\Client\Model\InlineObject24**](../Model/InlineObject24.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getCatalogItemType()`

```php
getCatalogItemType($id): \OpenAPI\Client\Model\InlineResponse20015
```

Get a Specific Catalog Item Type

This endpoint retrieves a specific category item type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CatalogItemsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCatalogItemType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CatalogItemsApi->getCatalogItemType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20015**](../Model/InlineResponse20015.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listCatalogItemTypes()`

```php
listCatalogItemTypes($max, $offset, $sort, $direction, $phrase, $name, $description, $enabled, $featured, $labels, $all_labels, $code): object
```

Get All Catalog Item Types

This endpoint retrieves all catalog item types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CatalogItemsApi(
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
$description = The desription of my cool object; // string | Filter by description, wildcard may be specified as %. eg. `example-%`
$enabled = false; // bool | Filter by enabled
$featured = false; // bool | Filter by featured
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$code = azr; // string | If specified will return an exact match on code

try {
    $result = $apiInstance->listCatalogItemTypes($max, $offset, $sort, $direction, $phrase, $name, $description, $enabled, $featured, $labels, $all_labels, $code);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CatalogItemsApi->listCatalogItemTypes: ', $e->getMessage(), PHP_EOL;
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
 **description** | **string**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **enabled** | **bool**| Filter by enabled | [optional]
 **featured** | **bool**| Filter by featured | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]

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

## `removeCatalogItemType()`

```php
removeCatalogItemType($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Catalog Item Type

Will delete a catalog item type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CatalogItemsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeCatalogItemType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CatalogItemsApi->removeCatalogItemType: ', $e->getMessage(), PHP_EOL;
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

## `updateCatalogItemType()`

```php
updateCatalogItemType($id, $inline_object25): object
```

Update a Catalog Item Type

Use this command to update an existing catalog item type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CatalogItemsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object25 = new \OpenAPI\Client\Model\InlineObject25(); // \OpenAPI\Client\Model\InlineObject25

try {
    $result = $apiInstance->updateCatalogItemType($id, $inline_object25);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CatalogItemsApi->updateCatalogItemType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object25** | [**\OpenAPI\Client\Model\InlineObject25**](../Model/InlineObject25.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`, `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateCatalogItemTypeLogo()`

```php
updateCatalogItemTypeLogo($id, $catalog_item_type_logo, $catalog_item_type_dark_logo): object
```

Update Logo For Catalog Item Type

Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CatalogItemsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$catalog_item_type_logo = "/path/to/file.txt"; // \SplFileObject | Logo File png,jpg,svg
$catalog_item_type_dark_logo = "/path/to/file.txt"; // \SplFileObject | Dark Logo File png,jpg,svg

try {
    $result = $apiInstance->updateCatalogItemTypeLogo($id, $catalog_item_type_logo, $catalog_item_type_dark_logo);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CatalogItemsApi->updateCatalogItemTypeLogo: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **catalog_item_type_logo** | **\SplFileObject****\SplFileObject**| Logo File png,jpg,svg | [optional]
 **catalog_item_type_dark_logo** | **\SplFileObject****\SplFileObject**| Dark Logo File png,jpg,svg | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
