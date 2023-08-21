# OpenAPI\Client\HostsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHost()**](HostsApi.md#getHost) | **GET** /api/servers/{id} | Get a Specific Host
[**getHostSnpshots()**](HostsApi.md#getHostSnpshots) | **GET** /api/servers/{id}/snapshots | Get list of snapshots for a Host
[**getHostType()**](HostsApi.md#getHostType) | **GET** /api/server-types/{id} | Get a Specific Host Type
[**getWikiServer()**](HostsApi.md#getWikiServer) | **GET** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**listHostTypes()**](HostsApi.md#listHostTypes) | **GET** /api/server-types | Host Types
[**listHosts()**](HostsApi.md#listHosts) | **GET** /api/servers | Get All Hosts
[**listServerServicePlans()**](HostsApi.md#listServerServicePlans) | **GET** /api/servers/service-plans | Get Available Service Plans for a Host
[**removeHost()**](HostsApi.md#removeHost) | **DELETE** /api/servers/{id} | Delete a Host
[**restartHost()**](HostsApi.md#restartHost) | **PUT** /api/servers/{id}/restart | Restart a Host
[**startHost()**](HostsApi.md#startHost) | **PUT** /api/servers/{id}/start | Start a Host
[**stopHost()**](HostsApi.md#stopHost) | **PUT** /api/servers/{id}/stop | Stop a Host
[**updateHost()**](HostsApi.md#updateHost) | **PUT** /api/servers/{id} | Updating a Host
[**updateHostAssignTenant()**](HostsApi.md#updateHostAssignTenant) | **PUT** /api/servers/{id}/assign-account | Assign To Tenant
[**updateHostCloud()**](HostsApi.md#updateHostCloud) | **PUT** /api/servers/change-cloud | Change Server Cloud
[**updateHostExecuteWorkflow()**](HostsApi.md#updateHostExecuteWorkflow) | **PUT** /api/servers/{id}/workflow | Run Workflow on a Host
[**updateHostInstallAgent()**](HostsApi.md#updateHostInstallAgent) | **PUT** /api/servers/{id}/install-agent | Install Agent
[**updateHostManaged()**](HostsApi.md#updateHostManaged) | **PUT** /api/servers/{id}/make-managed | Convert To Managed
[**updateHostResize()**](HostsApi.md#updateHostResize) | **PUT** /api/servers/{id}/resize | Resize a Host
[**updateHostUpgradeAgent()**](HostsApi.md#updateHostUpgradeAgent) | **PUT** /api/servers/{id}/upgrade | Upgrade Agent
[**updateServerNetworkInterface()**](HostsApi.md#updateServerNetworkInterface) | **PUT** /api/servers/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for a Server&#39;s Network
[**updateWikiServer()**](HostsApi.md#updateWikiServer) | **PUT** /api/servers/{id}/wiki | Update a Server Wiki Page


## `getHost()`

```php
getHost($id): \OpenAPI\Client\Model\InlineResponse200137
```

Get a Specific Host

This endpoint retrieves a specific host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getHost($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->getHost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200137**](../Model/InlineResponse200137.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getHostSnpshots()`

```php
getHostSnpshots($id): \OpenAPI\Client\Model\InlineResponse200138
```

Get list of snapshots for a Host

Get list of snapshots for a Host

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getHostSnpshots($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->getHostSnpshots: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200138**](../Model/InlineResponse200138.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getHostType()`

```php
getHostType($id): \OpenAPI\Client\Model\InlineResponse20050
```

Get a Specific Host Type

This endpoint will retrieve a specific host type by id

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getHostType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->getHostType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20050**](../Model/InlineResponse20050.md)

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


$apiInstance = new OpenAPI\Client\Api\HostsApi(
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
    echo 'Exception when calling HostsApi->getWikiServer: ', $e->getMessage(), PHP_EOL;
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

## `listHostTypes()`

```php
listHostTypes($max, $offset, $sort, $direction, $name, $code, $phrase, $provision_type, $zone_type, $creatable): object
```

Host Types

Fetch a paginated list of available host types. This returns the configuration options for each type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$provision_type = 'provision_type_example'; // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.
$zone_type = 'zone_type_example'; // string | Filter by Cloud Type code.
$creatable = True; // bool | Filter by creatable flag. This is whether or not it can be provisioned.

try {
    $result = $apiInstance->listHostTypes($max, $offset, $sort, $direction, $name, $code, $phrase, $provision_type, $zone_type, $creatable);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->listHostTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **provision_type** | **string**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings. | [optional]
 **zone_type** | **string**| Filter by Cloud Type code. | [optional]
 **creatable** | **bool**| Filter by creatable flag. This is whether or not it can be provisioned. | [optional]

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

## `listHosts()`

```php
listHosts($name, $phrase, $zone_id, $site_id, $cluster_id, $managed, $server_type, $power_state, $ip, $vm, $vm_hypervisor, $bare_metal_host, $status, $agent_installed, $max, $offset, $last_updated, $created_by, $labels, $all_labels, $tags, $metadata, $uuid, $external_id, $internal_id, $external_uniquel_id): object
```

Get All Hosts

This endpoint retrieves a paginated list of hosts.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$zone_id = 3; // int | The Zone ID for Filtering
$site_id = 7; // int | The Site ID for Filtering
$cluster_id = 135; // int | The Cluster ID(s) for filtering. Accepts multiple values.
$managed = false; // bool | Filter by managed (true) or unmanaged (false)
$server_type = vmwareHypervisor; // string | Filter by server type code
$power_state = off; // string | Filter by power status
$ip = 192.168.1.45; // string | Filter by IP address
$vm = false; // bool | Filter to show only Virtual Machines (true)
$vm_hypervisor = false; // bool | Filter to show only VM Hypervisors (true)
$bare_metal_host = false; // bool | Filter to show only Baremetal Servers
$status = 'status_example'; // string | Filter by status
$agent_installed = true; // bool | Filter by agent installed (true)
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$created_by = 63; // int | The User ID for Filtering
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$tags = tags.env=qa&tags.env=test; // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)
$metadata = tags.env=qa&tags.env=test; // string | Alias for tags
$uuid = 'uuid_example'; // string | Filter by UUID
$external_id = 'external_id_example'; // string | Filter by External ID
$internal_id = 'internal_id_example'; // string | Filter by Internal ID
$external_uniquel_id = 'external_uniquel_id_example'; // string | Filter by External Unique ID

try {
    $result = $apiInstance->listHosts($name, $phrase, $zone_id, $site_id, $cluster_id, $managed, $server_type, $power_state, $ip, $vm, $vm_hypervisor, $bare_metal_host, $status, $agent_installed, $max, $offset, $last_updated, $created_by, $labels, $all_labels, $tags, $metadata, $uuid, $external_id, $internal_id, $external_uniquel_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->listHosts: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]
 **cluster_id** | **int**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional]
 **managed** | **bool**| Filter by managed (true) or unmanaged (false) | [optional]
 **server_type** | **string**| Filter by server type code | [optional]
 **power_state** | **string**| Filter by power status | [optional]
 **ip** | **string**| Filter by IP address | [optional]
 **vm** | **bool**| Filter to show only Virtual Machines (true) | [optional]
 **vm_hypervisor** | **bool**| Filter to show only VM Hypervisors (true) | [optional]
 **bare_metal_host** | **bool**| Filter to show only Baremetal Servers | [optional]
 **status** | **string**| Filter by status | [optional]
 **agent_installed** | **bool**| Filter by agent installed (true) | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **created_by** | **int**| The User ID for Filtering | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **tags** | **string**| Filter by tags (metadata). This allows filtering by a tag name and value(s) | [optional]
 **metadata** | **string**| Alias for tags | [optional]
 **uuid** | **string**| Filter by UUID | [optional]
 **external_id** | **string**| Filter by External ID | [optional]
 **internal_id** | **string**| Filter by Internal ID | [optional]
 **external_uniquel_id** | **string**| Filter by External Unique ID | [optional]

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

## `listServerServicePlans()`

```php
listServerServicePlans($zone_id, $server_type_id, $site_id): \OpenAPI\Client\Model\InlineResponse200141
```

Get Available Service Plans for a Host

This endpoint retrieves all the Service Plans available for the specified cloud and host type. It may be used to get the list of available plans when creating a new host or resizing an existing host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 3; // int | The Zone ID for Filtering
$server_type_id = 5; // int | The ID of the Host Type
$site_id = 7; // int | The Site ID for Filtering

try {
    $result = $apiInstance->listServerServicePlans($zone_id, $server_type_id, $site_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->listServerServicePlans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **int**| The Zone ID for Filtering |
 **server_type_id** | **int**| The ID of the Host Type | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200141**](../Model/InlineResponse200141.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `removeHost()`

```php
removeHost($id, $remove_resources, $remove_instances, $preserve_volumes, $release_floating_ips, $release_ei_ps, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Host

Will delete a host asynchronously.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$remove_resources = off; // string | Remove Resources
$remove_instances = on; // string | Remove Instances
$preserve_volumes = on; // string | Preserve Volumes
$release_floating_ips = off; // string | Release Floating IPs
$release_ei_ps = off; // string | Alias for releaseFloatingIps
$force = on; // string | Force Delete

try {
    $result = $apiInstance->removeHost($id, $remove_resources, $remove_instances, $preserve_volumes, $release_floating_ips, $release_ei_ps, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->removeHost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_resources** | **string**| Remove Resources | [optional] [default to &#39;on&#39;]
 **remove_instances** | **string**| Remove Instances | [optional] [default to &#39;off&#39;]
 **preserve_volumes** | **string**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **release_floating_ips** | **string**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **release_ei_ps** | **string**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
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

## `restartHost()`

```php
restartHost($id): object
```

Restart a Host

This will restart a host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartHost($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->restartHost: ', $e->getMessage(), PHP_EOL;
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

## `startHost()`

```php
startHost($id): \OpenAPI\Client\Model\Model200Success
```

Start a Host

This will start a host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->startHost($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->startHost: ', $e->getMessage(), PHP_EOL;
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

## `stopHost()`

```php
stopHost($id): \OpenAPI\Client\Model\Model200Success
```

Stop a Host

This will stop a host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->stopHost($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->stopHost: ', $e->getMessage(), PHP_EOL;
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

## `updateHost()`

```php
updateHost($id, $inline_object220): \OpenAPI\Client\Model\InlineResponse200137
```

Updating a Host

Updating a Host

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object220 = new \OpenAPI\Client\Model\InlineObject220(); // \OpenAPI\Client\Model\InlineObject220

try {
    $result = $apiInstance->updateHost($id, $inline_object220);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object220** | [**\OpenAPI\Client\Model\InlineObject220**](../Model/InlineObject220.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200137**](../Model/InlineResponse200137.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateHostAssignTenant()`

```php
updateHostAssignTenant($id, $account_id, $inline_object221): object
```

Assign To Tenant

This will change the ownership of the host to the specified Tenant account. This is only available to Master Tenant users.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$account_id = 3; // int | ID of the Tenant
$inline_object221 = new \OpenAPI\Client\Model\InlineObject221(); // \OpenAPI\Client\Model\InlineObject221

try {
    $result = $apiInstance->updateHostAssignTenant($id, $account_id, $inline_object221);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostAssignTenant: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **account_id** | **int**| ID of the Tenant | [optional]
 **inline_object221** | [**\OpenAPI\Client\Model\InlineObject221**](../Model/InlineObject221.md)|  | [optional]

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

## `updateHostCloud()`

```php
updateHostCloud($inline_object226): object
```

Change Server Cloud

This api call is reserved for migrating servers from one cloud to another. This could be due to moving clusters or resource pool scoping of a server without losing the data.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object226 = new \OpenAPI\Client\Model\InlineObject226(); // \OpenAPI\Client\Model\InlineObject226

try {
    $result = $apiInstance->updateHostCloud($inline_object226);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object226** | [**\OpenAPI\Client\Model\InlineObject226**](../Model/InlineObject226.md)|  | [optional]

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

## `updateHostExecuteWorkflow()`

```php
updateHostExecuteWorkflow($id, $workflow_id, $workflow_name, $inline_object225): \OpenAPI\Client\Model\Model200Success
```

Run Workflow on a Host

This will run a provisioning workflow on a host.  For operational workflows, see Execute a Workflow.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$workflow_id = 3; // int | ID of the workflow to execute
$workflow_name = myworkflow; // string | Name of the workflow to execute
$inline_object225 = new \OpenAPI\Client\Model\InlineObject225(); // \OpenAPI\Client\Model\InlineObject225

try {
    $result = $apiInstance->updateHostExecuteWorkflow($id, $workflow_id, $workflow_name, $inline_object225);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostExecuteWorkflow: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **workflow_id** | **int**| ID of the workflow to execute | [optional]
 **workflow_name** | **string**| Name of the workflow to execute | [optional]
 **inline_object225** | [**\OpenAPI\Client\Model\InlineObject225**](../Model/InlineObject225.md)|  | [optional]

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

## `updateHostInstallAgent()`

```php
updateHostInstallAgent($id, $inline_object222): object
```

Install Agent

This will make the host a managed server, and install the agent.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object222 = new \OpenAPI\Client\Model\InlineObject222(); // \OpenAPI\Client\Model\InlineObject222

try {
    $result = $apiInstance->updateHostInstallAgent($id, $inline_object222);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostInstallAgent: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object222** | [**\OpenAPI\Client\Model\InlineObject222**](../Model/InlineObject222.md)|  | [optional]

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

## `updateHostManaged()`

```php
updateHostManaged($id, $inline_object223): object
```

Convert To Managed

This will make the host a managed server, and install the agent.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object223 = new \OpenAPI\Client\Model\InlineObject223(); // \OpenAPI\Client\Model\InlineObject223

try {
    $result = $apiInstance->updateHostManaged($id, $inline_object223);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostManaged: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object223** | [**\OpenAPI\Client\Model\InlineObject223**](../Model/InlineObject223.md)|  | [optional]

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

## `updateHostResize()`

```php
updateHostResize($id, $inline_object224): object
```

Resize a Host

Will resize a host asynchronously. This endpoint also allows for NIC reconfiguration by passing a new array of `networkInterfaces`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object224 = new \OpenAPI\Client\Model\InlineObject224(); // \OpenAPI\Client\Model\InlineObject224

try {
    $result = $apiInstance->updateHostResize($id, $inline_object224);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostResize: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object224** | [**\OpenAPI\Client\Model\InlineObject224**](../Model/InlineObject224.md)|  | [optional]

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

## `updateHostUpgradeAgent()`

```php
updateHostUpgradeAgent($id): \OpenAPI\Client\Model\Model200Success
```

Upgrade Agent

This will upgrade the version of the agent installed on the host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->updateHostUpgradeAgent($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateHostUpgradeAgent: ', $e->getMessage(), PHP_EOL;
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

## `updateServerNetworkInterface()`

```php
updateServerNetworkInterface($id, $network_interface_id, $network_interface_update): object
```

Updating a label for a Server's Network

Updating a Server's Network's Label

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\HostsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$network_interface_id = 7; // float | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
$network_interface_update = {"name":"newLabelName"}; // \OpenAPI\Client\Model\NetworkInterfaceUpdate

try {
    $result = $apiInstance->updateServerNetworkInterface($id, $network_interface_id, $network_interface_update);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HostsApi->updateServerNetworkInterface: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **network_interface_id** | **float**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced |
 **network_interface_update** | [**\OpenAPI\Client\Model\NetworkInterfaceUpdate**](../Model/NetworkInterfaceUpdate.md)|  | [optional]

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


$apiInstance = new OpenAPI\Client\Api\HostsApi(
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
    echo 'Exception when calling HostsApi->updateWikiServer: ', $e->getMessage(), PHP_EOL;
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
