# OpenAPI\Client\InvoicesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getInvoiceLineItems()**](InvoicesApi.md#getInvoiceLineItems) | **GET** /api/invoice-line-items/{id} | Get a Specific Invoice Line Item
[**getInvoices()**](InvoicesApi.md#getInvoices) | **GET** /api/invoices/{id} | Get a Specific Invoice
[**listInvoiceLineItems()**](InvoicesApi.md#listInvoiceLineItems) | **GET** /api/invoice-line-items | List All Invoice Line Items
[**listInvoices()**](InvoicesApi.md#listInvoices) | **GET** /api/invoices | List All Invoices
[**updateInvoices()**](InvoicesApi.md#updateInvoices) | **PUT** /api/invoices/{id} | Update Invoice Tags


## `getInvoiceLineItems()`

```php
getInvoiceLineItems($id): object
```

Get a Specific Invoice Line Item

Get details about a specific invoice line item.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InvoicesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getInvoiceLineItems($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InvoicesApi->getInvoiceLineItems: ', $e->getMessage(), PHP_EOL;
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

## `getInvoices()`

```php
getInvoices($id): object
```

Get a Specific Invoice

Get details about a specific invoice.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InvoicesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getInvoices($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InvoicesApi->getInvoices: ', $e->getMessage(), PHP_EOL;
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

## `listInvoiceLineItems()`

```php
listInvoiceLineItems($max, $offset, $sort, $direction, $phrase, $name, $start_date, $end_date, $period, $ref_type, $ref_id, $zone_id, $site_id, $instance_id, $container_id, $server_id, $user_id, $project_id, $active, $account_id, $include_totals): object
```

List All Invoice Line Items

This endpoint retrieves a list of all invoice line items for the specified parameters.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InvoicesApi(
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
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$period = 201901; // string | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.
$ref_type = ComputeSite; // string | If specified will return an exact match on refType.
$ref_id = 3; // int | If specified will return an exact match on refId
$zone_id = 3; // int | The Zone ID for Filtering
$site_id = 7; // int | The Site ID for Filtering
$instance_id = 94; // int | The Instance ID for Filtering
$container_id = 135; // int | The Container ID for Filtering
$server_id = 97; // int | The Server ID for Filtering
$user_id = 6; // int | Filter by User ID
$project_id = 1; // int | The Project ID for Filtering
$active = false; // bool | True or False flag for Active
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
$include_totals = false; // bool | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`.

try {
    $result = $apiInstance->listInvoiceLineItems($max, $offset, $sort, $direction, $phrase, $name, $start_date, $end_date, $period, $ref_type, $ref_id, $zone_id, $site_id, $instance_id, $container_id, $server_id, $user_id, $project_id, $active, $account_id, $include_totals);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InvoicesApi->listInvoiceLineItems: ', $e->getMessage(), PHP_EOL;
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
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **period** | **string**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM. | [optional]
 **ref_type** | **string**| If specified will return an exact match on refType. | [optional]
 **ref_id** | **int**| If specified will return an exact match on refId | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **user_id** | **int**| Filter by User ID | [optional]
 **project_id** | **int**| The Project ID for Filtering | [optional]
 **active** | **bool**| True or False flag for Active | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **include_totals** | **bool**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;. | [optional] [default to false]

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

## `listInvoices()`

```php
listInvoices($max, $offset, $sort, $direction, $phrase, $name, $start_date, $end_date, $period, $ref_type, $ref_id, $ref_status, $zone_id, $site_id, $instance_id, $container_id, $server_id, $user_id, $project_id, $active, $account_id, $include_line_items, $include_totals, $tags): object
```

List All Invoices

This endpoint retrieves a list of invoices for the specified parameters.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InvoicesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'refName'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$period = 201901; // string | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.
$ref_type = ComputeSite; // string | If specified will return an exact match on refType.
$ref_id = 3; // int | If specified will return an exact match on refId
$ref_status = provisioned; // string | If specified, will filter on the associated StorageVolume status. This is only applicable whn `refType=StorageVolume`.
$zone_id = 3; // int | The Zone ID for Filtering
$site_id = 7; // int | The Site ID for Filtering
$instance_id = 94; // int | The Instance ID for Filtering
$container_id = 135; // int | The Container ID for Filtering
$server_id = 97; // int | The Server ID for Filtering
$user_id = 6; // int | Filter by User ID
$project_id = 1; // int | The Project ID for Filtering
$active = false; // bool | True or False flag for Active
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
$include_line_items = false; // bool | Pass true to include the list of `lineItems` for each invoice. Only `lineItemCount` is returned by default.
$include_totals = false; // bool | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`.
$tags = tags.env=qa&tags.env=test; // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)

try {
    $result = $apiInstance->listInvoices($max, $offset, $sort, $direction, $phrase, $name, $start_date, $end_date, $period, $ref_type, $ref_id, $ref_status, $zone_id, $site_id, $instance_id, $container_id, $server_id, $user_id, $project_id, $active, $account_id, $include_line_items, $include_totals, $tags);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InvoicesApi->listInvoices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;refName&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **period** | **string**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM. | [optional]
 **ref_type** | **string**| If specified will return an exact match on refType. | [optional]
 **ref_id** | **int**| If specified will return an exact match on refId | [optional]
 **ref_status** | **string**| If specified, will filter on the associated StorageVolume status. This is only applicable whn &#x60;refType&#x3D;StorageVolume&#x60;. | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **user_id** | **int**| Filter by User ID | [optional]
 **project_id** | **int**| The Project ID for Filtering | [optional]
 **active** | **bool**| True or False flag for Active | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **include_line_items** | **bool**| Pass true to include the list of &#x60;lineItems&#x60; for each invoice. Only &#x60;lineItemCount&#x60; is returned by default. | [optional] [default to false]
 **include_totals** | **bool**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;. | [optional] [default to false]
 **tags** | **string**| Filter by tags (metadata). This allows filtering by a tag name and value(s) | [optional]

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

## `updateInvoices()`

```php
updateInvoices($id, $inline_object102): object
```

Update Invoice Tags

Update, Add, or Remove invoice tag(s).

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InvoicesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object102 = new \OpenAPI\Client\Model\InlineObject102(); // \OpenAPI\Client\Model\InlineObject102

try {
    $result = $apiInstance->updateInvoices($id, $inline_object102);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InvoicesApi->updateInvoices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object102** | [**\OpenAPI\Client\Model\InlineObject102**](../Model/InlineObject102.md)|  | [optional]

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
