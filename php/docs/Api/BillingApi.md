# OpenAPI\Client\BillingApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getBillingAccount()**](BillingApi.md#getBillingAccount) | **GET** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it
[**getBillingInstancesIdentifier()**](BillingApi.md#getBillingInstancesIdentifier) | **GET** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.
[**getBillingServersIdentifier()**](BillingApi.md#getBillingServersIdentifier) | **GET** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.
[**getBillingZoneIdentifier()**](BillingApi.md#getBillingZoneIdentifier) | **GET** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.
[**listBillingAccount()**](BillingApi.md#listBillingAccount) | **GET** /api/billing/account | Retrieves billing information for the requesting user&#39;s account.
[**listBillingInstances()**](BillingApi.md#listBillingInstances) | **GET** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account.
[**listBillingServers()**](BillingApi.md#listBillingServers) | **GET** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.
[**listBillingZone()**](BillingApi.md#listBillingZone) | **GET** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account.


## `getBillingAccount()`

```php
getBillingAccount($id, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots): object
```

This endpoint will retrieve a specific account by id if the user has permission to access it

Will retrieve billing information for a specific tenant, if it is the current account or a sub account of the requesting user's account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_compute_servers = true; // bool | Optional ability to exclude compute servers
$include_instances = true; // bool | Optional ability to exclude instances
$include_discovered_servers = true; // bool | Optional ability to exclude discovered servers
$include_load_balancers = true; // bool | Optional ability to exclude load balancers
$include_virtual_images = true; // bool | Optional ability to exclude virtual images
$include_snapshots = true; // bool | Optional ability to exclude snapshots

try {
    $result = $apiInstance->getBillingAccount($id, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->getBillingAccount: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] [default to true]
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] [default to true]
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] [default to true]
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] [default to true]
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] [default to true]
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] [default to true]

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

## `getBillingInstancesIdentifier()`

```php
getBillingInstancesIdentifier($identifier, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id): object
```

Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$identifier = 'identifier_example'; // string | Morpheus UUID or ID of the Object being created or referenced
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_tenants = false; // bool | Optional ability to include all subtenant billing information when calling from a master tenant user
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.

try {
    $result = $apiInstance->getBillingInstancesIdentifier($identifier, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->getBillingInstancesIdentifier: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string**| Morpheus UUID or ID of the Object being created or referenced |
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

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

## `getBillingServersIdentifier()`

```php
getBillingServersIdentifier($identifier, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id): object
```

Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$identifier = 'identifier_example'; // string | Morpheus UUID or ID of the Object being created or referenced
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_tenants = false; // bool | Optional ability to include all subtenant billing information when calling from a master tenant user
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.

try {
    $result = $apiInstance->getBillingServersIdentifier($identifier, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->getBillingServersIdentifier: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string**| Morpheus UUID or ID of the Object being created or referenced |
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

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

## `getBillingZoneIdentifier()`

```php
getBillingZoneIdentifier($identifier, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots): object
```

Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$identifier = 'identifier_example'; // string | Morpheus UUID or ID of the Object being created or referenced
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_compute_servers = true; // bool | Optional ability to exclude compute servers
$include_instances = true; // bool | Optional ability to exclude instances
$include_discovered_servers = true; // bool | Optional ability to exclude discovered servers
$include_load_balancers = true; // bool | Optional ability to exclude load balancers
$include_virtual_images = true; // bool | Optional ability to exclude virtual images
$include_snapshots = true; // bool | Optional ability to exclude snapshots

try {
    $result = $apiInstance->getBillingZoneIdentifier($identifier, $start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->getBillingZoneIdentifier: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string**| Morpheus UUID or ID of the Object being created or referenced |
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] [default to true]
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] [default to true]
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] [default to true]
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] [default to true]
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] [default to true]
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] [default to true]

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

## `listBillingAccount()`

```php
listBillingAccount($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots): object
```

Retrieves billing information for the requesting user's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_compute_servers = true; // bool | Optional ability to exclude compute servers
$include_instances = true; // bool | Optional ability to exclude instances
$include_discovered_servers = true; // bool | Optional ability to exclude discovered servers
$include_load_balancers = true; // bool | Optional ability to exclude load balancers
$include_virtual_images = true; // bool | Optional ability to exclude virtual images
$include_snapshots = true; // bool | Optional ability to exclude snapshots

try {
    $result = $apiInstance->listBillingAccount($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->listBillingAccount: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] [default to true]
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] [default to true]
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] [default to true]
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] [default to true]
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] [default to true]
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] [default to true]

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

## `listBillingInstances()`

```php
listBillingInstances($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id): object
```

Retrieves billing information for all instances on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_tenants = false; // bool | Optional ability to include all subtenant billing information when calling from a master tenant user
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.

try {
    $result = $apiInstance->listBillingInstances($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->listBillingInstances: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

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

## `listBillingServers()`

```php
listBillingServers($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id): object
```

Retrieves billing information for all servers (container hosts) on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_tenants = false; // bool | Optional ability to include all subtenant billing information when calling from a master tenant user
$account_id = 3; // int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.

try {
    $result = $apiInstance->listBillingServers($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_tenants, $account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->listBillingServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

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

## `listBillingZone()`

```php
listBillingZone($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots): object
```

Retrieves billing information for all zones on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BillingApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$start_date = 2019-01-01; // string | Filter by startDate greater than or equal to a specified date
$end_date = 2020-01-01; // string | Filter by endDate less than or equal to a specified date
$include_usages = true; // bool | Optional ability to suppress the usage records
$max_usages = 56; // int | Optional ability to limit the usages returned
$offset_usages = 56; // int | Optional ability to offset the usages returned, for use with maxUsages to paginate
$include_compute_servers = true; // bool | Optional ability to exclude compute servers
$include_instances = true; // bool | Optional ability to exclude instances
$include_discovered_servers = true; // bool | Optional ability to exclude discovered servers
$include_load_balancers = true; // bool | Optional ability to exclude load balancers
$include_virtual_images = true; // bool | Optional ability to exclude virtual images
$include_snapshots = true; // bool | Optional ability to exclude snapshots

try {
    $result = $apiInstance->listBillingZone($start_date, $end_date, $include_usages, $max_usages, $offset_usages, $include_compute_servers, $include_instances, $include_discovered_servers, $include_load_balancers, $include_virtual_images, $include_snapshots);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BillingApi->listBillingZone: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **string**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **string**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] [default to true]
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] [default to true]
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] [default to true]
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] [default to true]
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] [default to true]
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] [default to true]
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] [default to true]

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
