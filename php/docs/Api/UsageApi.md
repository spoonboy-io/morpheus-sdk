# OpenAPI\Client\UsageApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**listUsages()**](UsageApi.md#listUsages) | **GET** /api/usage | Retrieves Usage Records


## `listUsages()`

```php
listUsages($max, $offset, $sort, $direction, $phrase, $start_date, $end_date): \OpenAPI\Client\Model\Usages
```

Retrieves Usage Records

Retrieves a paginated list of usage records. The usages are scoped to only include resources you have access to.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\UsageApi(
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
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date

try {
    $result = $apiInstance->listUsages($max, $offset, $sort, $direction, $phrase, $start_date, $end_date);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsageApi->listUsages: ', $e->getMessage(), PHP_EOL;
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
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]

### Return type

[**\OpenAPI\Client\Model\Usages**](../Model/Usages.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
