# OpenAPI\Client\ActivityApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**listActivity()**](ActivityApi.md#listActivity) | **GET** /api/activity | Retrieves Activity


## `listActivity()`

```php
listActivity($max, $offset, $sort, $order, $phrase, $name, $user_id, $tenant_id, $timeframe, $start, $end): object
```

Retrieves Activity

Retrieves activity.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ActivityApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$user_id = 6; // int | Filter by User ID
$tenant_id = 3; // float | Filter by Tenant ID. Only available to the master account.
$timeframe = 'month'; // string | Filter by a timeframe (ex - today, yesterday, week, month, 3months)
$start = new \DateTime("2013-10-20T19:20:30+01:00"); // \DateTime | Filter by activity on or after a date(time). Default is 1 month prior
$end = new \DateTime("2013-10-20T19:20:30+01:00"); // \DateTime | Filter by activity on or before a date(time). Default is current date

try {
    $result = $apiInstance->listActivity($max, $offset, $sort, $order, $phrase, $name, $user_id, $tenant_id, $timeframe, $start, $end);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ActivityApi->listActivity: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **user_id** | **int**| Filter by User ID | [optional]
 **tenant_id** | **float**| Filter by Tenant ID. Only available to the master account. | [optional]
 **timeframe** | **string**| Filter by a timeframe (ex - today, yesterday, week, month, 3months) | [optional] [default to &#39;month&#39;]
 **start** | **\DateTime**| Filter by activity on or after a date(time). Default is 1 month prior | [optional]
 **end** | **\DateTime**| Filter by activity on or before a date(time). Default is current date | [optional]

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
