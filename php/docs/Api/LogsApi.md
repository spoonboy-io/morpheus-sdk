# OpenAPI\Client\LogsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**listLogs()**](LogsApi.md#listLogs) | **GET** /api/logs | Retrieves Logs


## `listLogs()`

```php
listLogs($max, $offset, $sort, $order, $query, $message, $source_type, $type_code, $object_id, $token, $level, $start_ms, $end_ms, $start_date_time, $end_date_time, $containers, $servers, $cluster_id): Log
```

Retrieves Logs

Retrieves logs based on filters provided.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LogsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$query = 'query_example'; // string | Alias for phrase
$message = 'message_example'; // string | Filter by message
$source_type = 'source_type_example'; // string | Filter by source type
$type_code = 'type_code_example'; // string | Filter by code type
$object_id = 15; // int | Filter by objectId
$token = 'token_example'; // string | Filter by token
$level = ERROR; // string | Filter by log level. Multiple values can be passed pipe delimited.
$start_ms = 1657309873; // int | Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified.
$end_ms = 1657309873; // int | Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified.
$start_date_time = 2019-03-06T17:52:29+0000; // \DateTime | Start Date timestamp (ISO 8601)
$end_date_time = 2019-03-06T17:52:29+0000; // \DateTime | End Date timestamp (ISO 8601)
$containers = 135; // int | The Container ID(s) for filtering. Accepts multiple values.
$servers = 135; // int | The Server ID(s) for filtering. Accepts multiple values.
$cluster_id = 135; // int | The Cluster ID(s) for filtering. Accepts multiple values.

try {
    $result = $apiInstance->listLogs($max, $offset, $sort, $order, $query, $message, $source_type, $type_code, $object_id, $token, $level, $start_ms, $end_ms, $start_date_time, $end_date_time, $containers, $servers, $cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LogsApi->listLogs: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **query** | **string**| Alias for phrase | [optional]
 **message** | **string**| Filter by message | [optional]
 **source_type** | **string**| Filter by source type | [optional]
 **type_code** | **string**| Filter by code type | [optional]
 **object_id** | **int**| Filter by objectId | [optional]
 **token** | **string**| Filter by token | [optional]
 **level** | **string**| Filter by log level. Multiple values can be passed pipe delimited. | [optional]
 **start_ms** | **int**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. | [optional]
 **end_ms** | **int**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. | [optional]
 **start_date_time** | **\DateTime**| Start Date timestamp (ISO 8601) | [optional]
 **end_date_time** | **\DateTime**| End Date timestamp (ISO 8601) | [optional]
 **containers** | **int**| The Container ID(s) for filtering. Accepts multiple values. | [optional]
 **servers** | **int**| The Server ID(s) for filtering. Accepts multiple values. | [optional]
 **cluster_id** | **int**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional]

### Return type

[**Log**](../Model/Log.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
