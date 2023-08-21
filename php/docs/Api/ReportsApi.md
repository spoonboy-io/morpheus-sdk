# OpenAPI\Client\ReportsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**downloadReports()**](ReportsApi.md#downloadReports) | **GET** /api/reports/download/{id}{format} | Downloads a specific report result as a file attachment
[**getReportTypes()**](ReportsApi.md#getReportTypes) | **GET** /api/report-types | This endpoint retrieves all available report types
[**getReports()**](ReportsApi.md#getReports) | **GET** /api/reports/{id} | This endpoint returns a specific report result
[**listReports()**](ReportsApi.md#listReports) | **GET** /api/reports | Returns all reports
[**removeReports()**](ReportsApi.md#removeReports) | **DELETE** /api/reports/{id} | This endpoint will delete a report result
[**runReports()**](ReportsApi.md#runReports) | **POST** /api/reports | Run specified report


## `downloadReports()`

```php
downloadReports($id, $format): \OpenAPI\Client\Model\Model200Success
```

Downloads a specific report result as a file attachment

This endpoint downloads a specific report result as a file attachment. The default file format is `.json`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ReportsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$format = '.json'; // string | Format of the rendered report file, `.json` or `.csv`.

try {
    $result = $apiInstance->downloadReports($id, $format);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->downloadReports: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **format** | **string**| Format of the rendered report file, &#x60;.json&#x60; or &#x60;.csv&#x60;. | [default to &#39;.json&#39;]

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

## `getReportTypes()`

```php
getReportTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $category): \OpenAPI\Client\Model\InlineResponse200129
```

This endpoint retrieves all available report types

This endpoint retrieves all available report types. A report type has optionTypes that define the parameters available when executing a report of that type. The sample response has been abbreviated.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ReportsApi(
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
$code = azr; // string | If specified will return an exact match on code
$category = 'category_example'; // string | If specified will return an exact match on category

try {
    $result = $apiInstance->getReportTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $category);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->getReportTypes: ', $e->getMessage(), PHP_EOL;
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
 **code** | **string**| If specified will return an exact match on code | [optional]
 **category** | **string**| If specified will return an exact match on category | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200129**](../Model/InlineResponse200129.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getReports()`

```php
getReports($id): \OpenAPI\Client\Model\InlineResponse200130
```

This endpoint returns a specific report result

This endpoint retrieves a specific report result. The response includes the result data as rows which can be used to render the report. Each report type will have sections for data and headers that vary by type, use Download a Specific Report to get the results organized by section.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ReportsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getReports($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->getReports: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200130**](../Model/InlineResponse200130.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listReports()`

```php
listReports($max, $offset, $sort, $direction, $phrase, $name, $report_type, $category): object
```

Returns all reports

This endpoint returns all reports. This is results of reports that have been executed in the past.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ReportsApi(
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
$report_type = 'report_type_example'; // string | If specified will return an exact match on report type code, accepts multiple values
$category = 'category_example'; // string | If specified will return an exact match on category

try {
    $result = $apiInstance->listReports($max, $offset, $sort, $direction, $phrase, $name, $report_type, $category);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->listReports: ', $e->getMessage(), PHP_EOL;
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
 **report_type** | **string**| If specified will return an exact match on report type code, accepts multiple values | [optional]
 **category** | **string**| If specified will return an exact match on category | [optional]

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

## `removeReports()`

```php
removeReports($id): \OpenAPI\Client\Model\Model200Success
```

This endpoint will delete a report result

This endpoint will delete a report result.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ReportsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeReports($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->removeReports: ', $e->getMessage(), PHP_EOL;
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

## `runReports()`

```php
runReports($inline_object205): object
```

Run specified report

This endpoint execute the specified report type and create a new report result.  The available parameters vary by report type. Refer to the defined `inputs` for each report.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ReportsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object205 = new \OpenAPI\Client\Model\InlineObject205(); // \OpenAPI\Client\Model\InlineObject205

try {
    $result = $apiInstance->runReports($inline_object205);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ReportsApi->runReports: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object205** | [**\OpenAPI\Client\Model\InlineObject205**](../Model/InlineObject205.md)|  | [optional]

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
