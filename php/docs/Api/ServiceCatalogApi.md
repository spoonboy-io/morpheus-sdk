# OpenAPI\Client\ServiceCatalogApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCatalogCart()**](ServiceCatalogApi.md#addCatalogCart) | **POST** /api/catalog/checkout | Checkout Catalog Cart
[**addCatalogCartItem()**](ServiceCatalogApi.md#addCatalogCartItem) | **PUT** /api/catalog/cart/items | Add Catalog Item to Cart
[**addCatalogOrder()**](ServiceCatalogApi.md#addCatalogOrder) | **POST** /api/catalog/orders | Place Catalog Order
[**deleteCatalogCart()**](ServiceCatalogApi.md#deleteCatalogCart) | **DELETE** /api/catalog/cart | Clear Catalog Cart
[**deleteCatalogCartItem()**](ServiceCatalogApi.md#deleteCatalogCartItem) | **DELETE** /api/catalog/cart/items/{id} | Remove a Catalog Item From Cart
[**deleteCatalogItem()**](ServiceCatalogApi.md#deleteCatalogItem) | **DELETE** /api/catalog/items/{id} | Delete a Catalog Inventory Item
[**getCatalogItem()**](ServiceCatalogApi.md#getCatalogItem) | **GET** /api/catalog/items/{id} | Get a Specific Catalog Inventory Item
[**getCatalogType()**](ServiceCatalogApi.md#getCatalogType) | **GET** /api/catalog/types/{id} | Get a Specific Catalog Type
[**listCatalogCart()**](ServiceCatalogApi.md#listCatalogCart) | **GET** /api/catalog/cart | Get Catalog Cart
[**listCatalogItems()**](ServiceCatalogApi.md#listCatalogItems) | **GET** /api/catalog/items | List Catalog Inventory Items
[**listCatalogTypes()**](ServiceCatalogApi.md#listCatalogTypes) | **GET** /api/catalog/types | List Catalog Types


## `addCatalogCart()`

```php
addCatalogCart(): object
```

Checkout Catalog Cart

Use this command to checkout, finalizing your cart and placing an order. This converts each item in the cart to an inventory item, changing the status from IN_CART to ORDERED and potentially starts the provisioning process for each item.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->addCatalogCart();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->addCatalogCart: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

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

## `addCatalogCartItem()`

```php
addCatalogCartItem($validate, $catalog_cart_item_create): object
```

Add Catalog Item to Cart

Use this command to add an item to your service catalog cart.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$validate = false; // bool | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory
$catalog_cart_item_create = {"item":{"type":{"name":"example"},"config":{"appName":"My App"}}}; // \OpenAPI\Client\Model\CatalogCartItemCreate

try {
    $result = $apiInstance->addCatalogCartItem($validate, $catalog_cart_item_create);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->addCatalogCartItem: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] [default to false]
 **catalog_cart_item_create** | [**\OpenAPI\Client\Model\CatalogCartItemCreate**](../Model/CatalogCartItemCreate.md)|  | [optional]

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

## `addCatalogOrder()`

```php
addCatalogOrder($validate, $inline_object227): object
```

Place Catalog Order

This will place an order for the specified items, adding items to the inventory right away, without using the cart.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$validate = false; // bool | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory
$inline_object227 = new \OpenAPI\Client\Model\InlineObject227(); // \OpenAPI\Client\Model\InlineObject227

try {
    $result = $apiInstance->addCatalogOrder($validate, $inline_object227);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->addCatalogOrder: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] [default to false]
 **inline_object227** | [**\OpenAPI\Client\Model\InlineObject227**](../Model/InlineObject227.md)|  | [optional]

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

## `deleteCatalogCart()`

```php
deleteCatalogCart(): \OpenAPI\Client\Model\Model200Success
```

Clear Catalog Cart

Use this command to empty your cart, deleting all the items in it.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->deleteCatalogCart();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->deleteCatalogCart: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

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

## `deleteCatalogCartItem()`

```php
deleteCatalogCartItem($id): \OpenAPI\Client\Model\Model200Success
```

Remove a Catalog Item From Cart

Will remove a catalog item that is currently in the cart.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteCatalogCartItem($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->deleteCatalogCartItem: ', $e->getMessage(), PHP_EOL;
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

## `deleteCatalogItem()`

```php
deleteCatalogItem($id, $preserve_volumes, $keep_backups, $release_floating_ips, $release_ei_ps, $remove_instances, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Catalog Inventory Item

Will delete a catalog inventory item, which by default will deprovision any associated any instances and servers.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$preserve_volumes = on; // string | Preserve Volumes
$keep_backups = on; // string | Preserve copy of backups
$release_floating_ips = off; // string | Release Floating IPs
$release_ei_ps = off; // string | Alias for releaseFloatingIps
$remove_instances = off; // string | Remove Instances. Only applies to type `blueprint` (Apps)
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteCatalogItem($id, $preserve_volumes, $keep_backups, $release_floating_ips, $release_ei_ps, $remove_instances, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->deleteCatalogItem: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **preserve_volumes** | **string**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **keep_backups** | **string**| Preserve copy of backups | [optional] [default to &#39;off&#39;]
 **release_floating_ips** | **string**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **release_ei_ps** | **string**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
 **remove_instances** | **string**| Remove Instances. Only applies to type &#x60;blueprint&#x60; (Apps) | [optional] [default to &#39;on&#39;]
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

## `getCatalogItem()`

```php
getCatalogItem($id): \OpenAPI\Client\Model\InlineResponse200140
```

Get a Specific Catalog Inventory Item

This endpoint retrieves a specific catalog inventory item.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCatalogItem($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->getCatalogItem: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200140**](../Model/InlineResponse200140.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getCatalogType()`

```php
getCatalogType($id): object
```

Get a Specific Catalog Type

This endpoint retrieves a specific catalog item type. This also returns an array of associated optionTypes that are used to configure the catalog item.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getCatalogType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->getCatalogType: ', $e->getMessage(), PHP_EOL;
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

## `listCatalogCart()`

```php
listCatalogCart(): \OpenAPI\Client\Model\InlineResponse200139
```

Get Catalog Cart

This endpoint retrieves the current catalog cart and all the items in it.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listCatalogCart();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->listCatalogCart: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200139**](../Model/InlineResponse200139.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listCatalogItems()`

```php
listCatalogItems($max, $offset, $sort, $direction, $phrase, $name): object
```

List Catalog Inventory Items

This endpoint retrieves a list of the catalog inventory items.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
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

try {
    $result = $apiInstance->listCatalogItems($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->listCatalogItems: ', $e->getMessage(), PHP_EOL;
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

## `listCatalogTypes()`

```php
listCatalogTypes($max, $offset, $sort, $direction, $phrase, $name, $featured): object
```

List Catalog Types

This endpoint retrieves the types available for ordering.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ServiceCatalogApi(
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
$featured = false; // bool | Filter by featured

try {
    $result = $apiInstance->listCatalogTypes($max, $offset, $sort, $direction, $phrase, $name, $featured);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ServiceCatalogApi->listCatalogTypes: ', $e->getMessage(), PHP_EOL;
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
 **featured** | **bool**| Filter by featured | [optional]

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
