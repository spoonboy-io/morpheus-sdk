# OpenAPI\Client\HealthApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**acknowledgeHealthAlarm()**](HealthApi.md#acknowledgeHealthAlarm) | **PUT** /api/health/alarms/{id}/acknowledge | Acknowledge a Health Alarm
[**acknowledgeHealthAlarms()**](HealthApi.md#acknowledgeHealthAlarms) | **PUT** /api/health/alarms/acknowledge | Acknowledge Many Health Alarms
[**exportHealthLogs()**](HealthApi.md#exportHealthLogs) | **GET** /api/health/logs/export | Export Appliance Health Logs
[**getHealthAlarms()**](HealthApi.md#getHealthAlarms) | **GET** /api/health/alarms/{id} | Retrieves a Specific Appliance Health Alarm
[**listHealth()**](HealthApi.md#listHealth) | **GET** /api/health | Retrieves Appliance Health
[**listHealthAlarms()**](HealthApi.md#listHealthAlarms) | **GET** /api/health/alarms | Retrieves Appliance Health Alarms
[**listHealthLogs()**](HealthApi.md#listHealthLogs) | **GET** /api/health/logs | Retrieves Appliance Health Logs
[**ping()**](HealthApi.md#ping) | **GET** /api/ping | Basic information about current Morpheus Installation


## `acknowledgeHealthAlarm()`

```php
acknowledgeHealthAlarm($id, $inline_object81): \OpenAPI\Client\Model\Model200Success
```

Acknowledge a Health Alarm

Acknowledge a specific health alarm.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object81 = new \OpenAPI\Client\Model\InlineObject81(); // \OpenAPI\Client\Model\InlineObject81

try {
    $result = $apiInstance->acknowledgeHealthAlarm($id, $inline_object81);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->acknowledgeHealthAlarm: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object81** | [**\OpenAPI\Client\Model\InlineObject81**](../Model/InlineObject81.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `acknowledgeHealthAlarms()`

```php
acknowledgeHealthAlarms($inline_object80): \OpenAPI\Client\Model\Model200Success
```

Acknowledge Many Health Alarms

Acknowledge health alarms.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object80 = new \OpenAPI\Client\Model\InlineObject80(); // \OpenAPI\Client\Model\InlineObject80

try {
    $result = $apiInstance->acknowledgeHealthAlarms($inline_object80);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->acknowledgeHealthAlarms: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object80** | [**\OpenAPI\Client\Model\InlineObject80**](../Model/InlineObject80.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `exportHealthLogs()`

```php
exportHealthLogs($max, $offset, $sort, $direction, $phrase, $name, $acknowledged, $start_date, $end_date, $reverse): \SplFileObject
```

Export Appliance Health Logs

This endpoint downloads the morpheus appliance logs as a file attachment. By default, the most recent 10,000 log entries are returned, with the newest at the end of the file. The format for each log entry is `timestamp` `level` `message`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
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
$acknowledged = false; // bool | True or False flag for Acknowledged items
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$reverse = True; // bool | Reverse order of records. This `true` by default when sort and direction are not passed, but `false` by default if either is passed. This means that by default the newest log entries are the bottom of the file.

try {
    $result = $apiInstance->exportHealthLogs($max, $offset, $sort, $direction, $phrase, $name, $acknowledged, $start_date, $end_date, $reverse);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->exportHealthLogs: ', $e->getMessage(), PHP_EOL;
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
 **acknowledged** | **bool**| True or False flag for Acknowledged items | [optional]
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **reverse** | **bool**| Reverse order of records. This &#x60;true&#x60; by default when sort and direction are not passed, but &#x60;false&#x60; by default if either is passed. This means that by default the newest log entries are the bottom of the file. | [optional]

### Return type

[**\SplFileObject**](../Model/\SplFileObject.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/octet-stream`, `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getHealthAlarms()`

```php
getHealthAlarms($id): object
```

Retrieves a Specific Appliance Health Alarm

This endpoint will retrieve a specific health alarm by ID.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getHealthAlarms($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->getHealthAlarms: ', $e->getMessage(), PHP_EOL;
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

## `listHealth()`

```php
listHealth(): object
```

Retrieves Appliance Health

This endpoint retrieves health info about the appliance such as cpu, memory and database usage. Elasticsearch statistics and queue usage are also returned.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listHealth();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->listHealth: ', $e->getMessage(), PHP_EOL;
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

## `listHealthAlarms()`

```php
listHealthAlarms($max, $offset, $sort, $direction, $phrase, $name, $acknowledged): object
```

Retrieves Appliance Health Alarms

This endpoint retrieves all health alarms, which are Operation notifications from Cloud and other Service Integrations. These alarms are not generated by the appliance, but synced and displayed for visibility. By default only open alarms are returned. Open alarms are those that have not yet been acknowledged.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
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
$acknowledged = false; // bool | True or False flag for Acknowledged items

try {
    $result = $apiInstance->listHealthAlarms($max, $offset, $sort, $direction, $phrase, $name, $acknowledged);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->listHealthAlarms: ', $e->getMessage(), PHP_EOL;
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
 **acknowledged** | **bool**| True or False flag for Acknowledged items | [optional]

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

## `listHealthLogs()`

```php
listHealthLogs($max, $offset, $sort, $direction, $phrase, $name, $acknowledged, $start_date, $end_date): object
```

Retrieves Appliance Health Logs

This endpoint retrieves all health logs. These are the logs of the remote appliance itself. These logs show all ui activity and are useful for troubleshooting and auditing. Stack traces are filtered for Morpheus services. Complete stack traces can be found in `/var/log/morpheus/morpheus-ui/current`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
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
$acknowledged = false; // bool | True or False flag for Acknowledged items
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date

try {
    $result = $apiInstance->listHealthLogs($max, $offset, $sort, $direction, $phrase, $name, $acknowledged, $start_date, $end_date);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->listHealthLogs: ', $e->getMessage(), PHP_EOL;
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
 **acknowledged** | **bool**| True or False flag for Acknowledged items | [optional]
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]

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

## `ping()`

```php
ping(): \OpenAPI\Client\Model\Ping
```

Basic information about current Morpheus Installation

This endpoint can be used to check the remote appliance build version and some other basic information.  This is an unsecured endpoint and does not require authorization. However, build version will not be returned unless you are authenticated with a valid access token.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HealthApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->ping();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->ping: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\Ping**](../Model/Ping.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
