# OpenAPI\Client\InstancesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addInstance()**](InstancesApi.md#addInstance) | **POST** /api/instances | Create an Instance
[**backupInstance()**](InstancesApi.md#backupInstance) | **PUT** /api/instances/{id}/backup | Backup an instance
[**backupsInstance()**](InstancesApi.md#backupsInstance) | **GET** /api/instances/{id}/backups | Get list of backups for an Instance
[**cancelExpirationInstance()**](InstancesApi.md#cancelExpirationInstance) | **PUT** /api/instances/{id}/cancel-expiration | Cancel Expiration of an Instance
[**cancelRemovalInstance()**](InstancesApi.md#cancelRemovalInstance) | **PUT** /api/instances/{id}/cancel-removal | Cancel Removal of an Instance
[**cancelShutdownInstance()**](InstancesApi.md#cancelShutdownInstance) | **PUT** /api/instances/{id}/cancel-shutdown | Cancel Shutdown of an Instance
[**cloneImageInstance()**](InstancesApi.md#cloneImageInstance) | **PUT** /api/instances/{id}/clone-image | Clone to Image
[**cloneInstance()**](InstancesApi.md#cloneInstance) | **PUT** /api/instances/{id}/clone | Clone an Instance
[**createInstanceSchedule()**](InstancesApi.md#createInstanceSchedule) | **POST** /api/instances/{id}/schedules | Create a new Instance Schedule
[**deleteAllSnapshotsInstance()**](InstancesApi.md#deleteAllSnapshotsInstance) | **DELETE** /api/instances/{id}/delete-all-snapshots | Delete All Snapshots of Instance
[**deleteAllSnapshotsInstanceContainer()**](InstancesApi.md#deleteAllSnapshotsInstanceContainer) | **DELETE** /api/instances/{id}/delete-container-snapshots/{containerId} | Delete All Snapshots of Instance Container
[**deleteInstance()**](InstancesApi.md#deleteInstance) | **DELETE** /api/instances/{id} | Delete an instance
[**deleteInstanceSchedule()**](InstancesApi.md#deleteInstanceSchedule) | **DELETE** /api/instances/{id}/schedules/{scheduleId} | Deletes an Instance Schedule
[**deleteSnapshotInstance()**](InstancesApi.md#deleteSnapshotInstance) | **DELETE** /api/snapshots/{id} | Delete Snapshot of Instance
[**ejectInstance()**](InstancesApi.md#ejectInstance) | **PUT** /api/instances/{id}/eject | Eject an instance
[**extendExpirationInstance()**](InstancesApi.md#extendExpirationInstance) | **PUT** /api/instances/{id}/extend-expiration | Extend Expiration of an Instance
[**extendShutdownInstance()**](InstancesApi.md#extendShutdownInstance) | **PUT** /api/instances/{id}/extend-shutdown | Extend Shutdown of an Instance
[**getEnvVariables()**](InstancesApi.md#getEnvVariables) | **GET** /api/instances/{id}/envs | Get Env Variables
[**getInstance()**](InstancesApi.md#getInstance) | **GET** /api/instances/{id} | Retrieves a Specific Instance
[**getInstanceContainers()**](InstancesApi.md#getInstanceContainers) | **GET** /api/instances/{id}/containers | Get Container Details
[**getInstanceHistory()**](InstancesApi.md#getInstanceHistory) | **GET** /api/instances/{id}/history | Get Instance History
[**getInstanceSchedule()**](InstancesApi.md#getInstanceSchedule) | **GET** /api/instances/{id}/schedules/{scheduleId} | Get a Specific Instance Schedule
[**getInstanceSchedules()**](InstancesApi.md#getInstanceSchedules) | **GET** /api/instances/{id}/schedules | Get all Instance Schedules
[**getInstanceThreshold()**](InstancesApi.md#getInstanceThreshold) | **GET** /api/instances/{id}/threshold | Get an Instance Scale Threshold
[**getInstanceTypeProvisioning()**](InstancesApi.md#getInstanceTypeProvisioning) | **GET** /api/instance-types/{id} | Get Specific Instance Type for Provisioning
[**getPrepareApplyInstance()**](InstancesApi.md#getPrepareApplyInstance) | **GET** /api/instances/{id}/prepare-apply | Prepare To Apply an Instance
[**getSnapshotInstance()**](InstancesApi.md#getSnapshotInstance) | **GET** /api/snapshots/{id} | Get a Specific Snapshot
[**getStateInstance()**](InstancesApi.md#getStateInstance) | **GET** /api/instances/{id}/state | Get State of an Instance
[**getValidateApplyInstance()**](InstancesApi.md#getValidateApplyInstance) | **POST** /api/instances/{id}/validate-apply | Validate Apply State for an Instance
[**getWikiInstance()**](InstancesApi.md#getWikiInstance) | **GET** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**importSnapshotInstance()**](InstancesApi.md#importSnapshotInstance) | **PUT** /api/instances/{id}/import-snapshot | Import Snapshot of an Instance
[**linkedCloneSnapshotInstance()**](InstancesApi.md#linkedCloneSnapshotInstance) | **PUT** /api/instances/{id}/linked-clone/{snapshotId} | Create Linked Clone of Instance Snapshot
[**listInstanceServicePlans()**](InstancesApi.md#listInstanceServicePlans) | **GET** /api/instances/service-plans | Get Available Service Plans for an Instance
[**listInstanceTypesProvisioning()**](InstancesApi.md#listInstanceTypesProvisioning) | **GET** /api/instance-types | Get All Instance Types for Provisioning
[**listInstances()**](InstancesApi.md#listInstances) | **GET** /api/instances | Get All Instances
[**listSecurityGroupsInstance()**](InstancesApi.md#listSecurityGroupsInstance) | **GET** /api/instances/{id}/security-groups | Get Security Groups for an Instance
[**lockInstance()**](InstancesApi.md#lockInstance) | **PUT** /api/instances/{id}/lock | Lock an Instance
[**refreshStateInstance()**](InstancesApi.md#refreshStateInstance) | **POST** /api/instances/{id}/refresh | Refresh State of an Instance
[**removeInstancesFromControl()**](InstancesApi.md#removeInstancesFromControl) | **POST** /api/instances/remove-from-control | Remove From Control
[**resizeInstance()**](InstancesApi.md#resizeInstance) | **PUT** /api/instances/{id}/resize | Resize an Instance
[**restartInstance()**](InstancesApi.md#restartInstance) | **PUT** /api/instances/{id}/restart | Restart an instance
[**revertSnapshotInstance()**](InstancesApi.md#revertSnapshotInstance) | **PUT** /api/instances/{id}/revert-snapshot/{snapshotId} | Revert Instance to Snapshot
[**runWorkflowInstance()**](InstancesApi.md#runWorkflowInstance) | **PUT** /api/instances/{id}/workflow | Run Workflow on an Instance
[**setApplyInstance()**](InstancesApi.md#setApplyInstance) | **POST** /api/instances/{id}/apply | Apply State of an Instance
[**setInstanceSecurityGroups()**](InstancesApi.md#setInstanceSecurityGroups) | **POST** /api/instances/{id}/security-groups | Set Security Groups for an Instance
[**snapshotInstance()**](InstancesApi.md#snapshotInstance) | **PUT** /api/instances/{id}/snapshot | Snapshot an Instance
[**snapshotsInstance()**](InstancesApi.md#snapshotsInstance) | **GET** /api/instances/{id}/snapshots | Get list of snapshots for an Instance
[**startInstance()**](InstancesApi.md#startInstance) | **PUT** /api/instances/{id}/start | Start an instance
[**stopInstance()**](InstancesApi.md#stopInstance) | **PUT** /api/instances/{id}/stop | Stop an instance
[**suspendInstance()**](InstancesApi.md#suspendInstance) | **PUT** /api/instances/{id}/suspend | Suspend an instance
[**unlockInstance()**](InstancesApi.md#unlockInstance) | **PUT** /api/instances/{id}/unlock | Unlock an Instance
[**updateInstance()**](InstancesApi.md#updateInstance) | **PUT** /api/instances/{id} | Updating an Instance
[**updateInstanceNetworkInterface()**](InstancesApi.md#updateInstanceNetworkInterface) | **PUT** /api/instances/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for an Instance&#39;s Network
[**updateInstanceSchedule()**](InstancesApi.md#updateInstanceSchedule) | **PUT** /api/instances/{id}/schedules/{scheduleId} | Updating an Instance Schedule
[**updateInstanceThreshold()**](InstancesApi.md#updateInstanceThreshold) | **PUT** /api/instances/{id}/threshold | Updates an Instance Scale Threshold
[**updateWikiInstance()**](InstancesApi.md#updateWikiInstance) | **PUT** /api/instances/{id}/wiki | Update an Instance Wiki Page


## `addInstance()`

```php
addInstance($instance_create): object
```

Create an Instance

Create an Instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_create = {"zoneId":15,"instance":{"name":"test2","cloud":"vcenter1 - homelab","site":{"id":1},"type":"ubuntu","instanceType":{"code":"ubuntu"},"instanceContext":"dev","labels":["ubuntuserver"],"layout":{"id":1399,"code":"vmware-ubuntu-20.04-single"},"plan":{"id":361,"code":"vm-512","name":"1 CPU, 512MB Memory"},"userGroup":{"id":1}},"config":{"resourcePoolId":75,"hostId":"","vmwareFolderId":"group-v90002","nestedVirtualization":"off","createUser":true},"volumes":[{"id":-1,"rootVolume":true,"name":"root","size":10,"sizeId":null,"storageType":1,"datastoreId":59}],"networkInterfaces":[{"network":{"id":"network-110"},"networkInterfaceTypeId":4}],"layoutSize":1}; // \OpenAPI\Client\Model\InstanceCreate

try {
    $result = $apiInstance->addInstance($instance_create);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->addInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_create** | [**\OpenAPI\Client\Model\InstanceCreate**](../Model/InstanceCreate.md)|  | [optional]

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

## `backupInstance()`

```php
backupInstance($id): \OpenAPI\Client\Model\Model200Success
```

Backup an instance

Backup an instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->backupInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->backupInstance: ', $e->getMessage(), PHP_EOL;
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

## `backupsInstance()`

```php
backupsInstance($id): \OpenAPI\Client\Model\InstanceBackups
```

Get list of backups for an Instance

Get list of backups for an Instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->backupsInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->backupsInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InstanceBackups**](../Model/InstanceBackups.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `cancelExpirationInstance()`

```php
cancelExpirationInstance($id): \OpenAPI\Client\Model\Model200Success
```

Cancel Expiration of an Instance

This operation will cancel the expiration of an instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->cancelExpirationInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->cancelExpirationInstance: ', $e->getMessage(), PHP_EOL;
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

## `cancelRemovalInstance()`

```php
cancelRemovalInstance($id): \OpenAPI\Client\Model\Model200Success
```

Cancel Removal of an Instance

This operation will undo the delete of an instance that is pending removal.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->cancelRemovalInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->cancelRemovalInstance: ', $e->getMessage(), PHP_EOL;
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

## `cancelShutdownInstance()`

```php
cancelShutdownInstance($id): \OpenAPI\Client\Model\Model200Success
```

Cancel Shutdown of an Instance

This operation will cancel the shutdown of an instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->cancelShutdownInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->cancelShutdownInstance: ', $e->getMessage(), PHP_EOL;
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

## `cloneImageInstance()`

```php
cloneImageInstance($id, $instances_clone_image): \OpenAPI\Client\Model\Model200Success
```

Clone to Image

This endpoint allows creating an image template from an existing instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$instances_clone_image = {"templateName":"Example Image","zoneFolder":"group-v81"}; // \OpenAPI\Client\Model\InstancesCloneImage

try {
    $result = $apiInstance->cloneImageInstance($id, $instances_clone_image);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->cloneImageInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instances_clone_image** | [**\OpenAPI\Client\Model\InstancesCloneImage**](../Model/InstancesCloneImage.md)|  | [optional]

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

## `cloneInstance()`

```php
cloneInstance($id, $instance_clone): \OpenAPI\Client\Model\Model200Success
```

Clone an Instance

One can easily clone an instance and all containers within that instance. The containers are backed up via the backup services and used as a snapshot to produce a clone of the instance. It is possible to clone this app instance into an entirely different availability zone.  This endpoint also supports all of the same parameters as instance creation, so you can override any configuration options when provisioning the clone.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$instance_clone = {"name":"New Name","group":{"id":1}}; // \OpenAPI\Client\Model\InstanceClone

try {
    $result = $apiInstance->cloneInstance($id, $instance_clone);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->cloneInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_clone** | [**\OpenAPI\Client\Model\InstanceClone**](../Model/InstanceClone.md)|  | [optional]

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

## `createInstanceSchedule()`

```php
createInstanceSchedule($id, $inline_object96): object
```

Create a new Instance Schedule

Create a new schedule for a specific instance.  This creates an instance scaling threshold that only applies during a defined schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 42; // int | Instance ID
$inline_object96 = new \OpenAPI\Client\Model\InlineObject96(); // \OpenAPI\Client\Model\InlineObject96

try {
    $result = $apiInstance->createInstanceSchedule($id, $inline_object96);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->createInstanceSchedule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **inline_object96** | [**\OpenAPI\Client\Model\InlineObject96**](../Model/InlineObject96.md)|  | [optional]

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

## `deleteAllSnapshotsInstance()`

```php
deleteAllSnapshotsInstance($id): \OpenAPI\Client\Model\Model200Success
```

Delete All Snapshots of Instance

Delete All Snapshots attached to Instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteAllSnapshotsInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->deleteAllSnapshotsInstance: ', $e->getMessage(), PHP_EOL;
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

## `deleteAllSnapshotsInstanceContainer()`

```php
deleteAllSnapshotsInstanceContainer($id, $container_id): \OpenAPI\Client\Model\Model200Success
```

Delete All Snapshots of Instance Container

Delete All Snapshots attached to Instance Container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$container_id = 4; // float | Container ID

try {
    $result = $apiInstance->deleteAllSnapshotsInstanceContainer($id, $container_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->deleteAllSnapshotsInstanceContainer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **container_id** | **float**| Container ID |

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

## `deleteInstance()`

```php
deleteInstance($id, $preserve_volumes, $keep_backups, $release_floating_ips, $release_ei_ps, $force): \OpenAPI\Client\Model\Model200Success
```

Delete an instance

Will delete an instance and all associated monitors and backups.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
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
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteInstance($id, $preserve_volumes, $keep_backups, $release_floating_ips, $release_ei_ps, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->deleteInstance: ', $e->getMessage(), PHP_EOL;
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

## `deleteInstanceSchedule()`

```php
deleteInstanceSchedule($id, $schedule_id): \OpenAPI\Client\Model\Model200Success
```

Deletes an Instance Schedule

Deletes a specified instance scaling schedule

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 42; // int | Instance ID
$schedule_id = 1; // int | Instance Schedule ID

try {
    $result = $apiInstance->deleteInstanceSchedule($id, $schedule_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->deleteInstanceSchedule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **schedule_id** | **int**| Instance Schedule ID |

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

## `deleteSnapshotInstance()`

```php
deleteSnapshotInstance($id): \OpenAPI\Client\Model\Model200Success
```

Delete Snapshot of Instance

Delete snapshot of instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteSnapshotInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->deleteSnapshotInstance: ', $e->getMessage(), PHP_EOL;
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

## `ejectInstance()`

```php
ejectInstance($id): \OpenAPI\Client\Model\Model200Success
```

Eject an instance

This will eject any ISO media on all containers in the instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->ejectInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->ejectInstance: ', $e->getMessage(), PHP_EOL;
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

## `extendExpirationInstance()`

```php
extendExpirationInstance($id): \OpenAPI\Client\Model\Model200Success
```

Extend Expiration of an Instance

This operation will extend the expiration of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->extendExpirationInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->extendExpirationInstance: ', $e->getMessage(), PHP_EOL;
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

## `extendShutdownInstance()`

```php
extendShutdownInstance($id): \OpenAPI\Client\Model\Model200Success
```

Extend Shutdown of an Instance

This operation will extend the shutdown of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->extendShutdownInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->extendShutdownInstance: ', $e->getMessage(), PHP_EOL;
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

## `getEnvVariables()`

```php
getEnvVariables($id): \OpenAPI\Client\Model\InlineResponse20057
```

Get Env Variables

This gets all the environment variables associated with the instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getEnvVariables($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getEnvVariables: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20057**](../Model/InlineResponse20057.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getInstance()`

```php
getInstance($id, $details): \OpenAPI\Client\Model\InlineResponse20056
```

Retrieves a Specific Instance

Retrieves a specific instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$details = true; // bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2.

try {
    $result = $apiInstance->getInstance($id, $details);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **details** | **bool**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] [default to false]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20056**](../Model/InlineResponse20056.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getInstanceContainers()`

```php
getInstanceContainers($id): object
```

Get Container Details

This can be valuable for evaluating the details of the compute server(s) running on an instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getInstanceContainers($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstanceContainers: ', $e->getMessage(), PHP_EOL;
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

## `getInstanceHistory()`

```php
getInstanceHistory($id, $phrase, $container_id, $server_id, $zone_id): object
```

Get Instance History

This endpoint retrieves the process history for a specific instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$container_id = 135; // int | The Container ID for Filtering
$server_id = 97; // int | The Server ID for Filtering
$zone_id = 3; // int | The Zone ID for Filtering

try {
    $result = $apiInstance->getInstanceHistory($id, $phrase, $container_id, $server_id, $zone_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstanceHistory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]

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

## `getInstanceSchedule()`

```php
getInstanceSchedule($id, $schedule_id): \OpenAPI\Client\Model\InlineResponse20059
```

Get a Specific Instance Schedule

This endpoint retrieves a specific instance scaling schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 42; // int | Instance ID
$schedule_id = 1; // int | Instance Schedule ID

try {
    $result = $apiInstance->getInstanceSchedule($id, $schedule_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstanceSchedule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **schedule_id** | **int**| Instance Schedule ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20059**](../Model/InlineResponse20059.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getInstanceSchedules()`

```php
getInstanceSchedules($id): object
```

Get all Instance Schedules

This endpoint retrieves all the scaling threshold schedules for a specific instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 42; // int | Instance ID

try {
    $result = $apiInstance->getInstanceSchedules($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstanceSchedules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |

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

## `getInstanceThreshold()`

```php
getInstanceThreshold($id): \OpenAPI\Client\Model\InlineResponse20058
```

Get an Instance Scale Threshold

Retrieves the scale threshold settings for a specific instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getInstanceThreshold($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstanceThreshold: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20058**](../Model/InlineResponse20058.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getInstanceTypeProvisioning()`

```php
getInstanceTypeProvisioning($id): object
```

Get Specific Instance Type for Provisioning

Fetch an instance type by ID.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getInstanceTypeProvisioning($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getInstanceTypeProvisioning: ', $e->getMessage(), PHP_EOL;
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

## `getPrepareApplyInstance()`

```php
getPrepareApplyInstance($id): object
```

Prepare To Apply an Instance

This endpoint provides a way to view the current instance configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPrepareApplyInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getPrepareApplyInstance: ', $e->getMessage(), PHP_EOL;
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

## `getSnapshotInstance()`

```php
getSnapshotInstance($id): \OpenAPI\Client\Model\Snapshot
```

Get a Specific Snapshot

This endpoint retrieves a specific snapshot.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getSnapshotInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getSnapshotInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Snapshot**](../Model/Snapshot.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStateInstance()`

```php
getStateInstance($id): object
```

Get State of an Instance

This endpoint provides a way to view the state of an instance. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStateInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getStateInstance: ', $e->getMessage(), PHP_EOL;
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

## `getValidateApplyInstance()`

```php
getValidateApplyInstance($id, $inline_object98): object
```

Validate Apply State for an Instance

This endpoint provides a way to validate instance configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object98 = new \OpenAPI\Client\Model\InlineObject98(); // \OpenAPI\Client\Model\InlineObject98

try {
    $result = $apiInstance->getValidateApplyInstance($id, $inline_object98);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getValidateApplyInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object98** | [**\OpenAPI\Client\Model\InlineObject98**](../Model/InlineObject98.md)|  | [optional]

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

## `getWikiInstance()`

```php
getWikiInstance($id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves an Instance Wiki Page

This endpoint retrieves an instance Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWikiInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->getWikiInstance: ', $e->getMessage(), PHP_EOL;
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

## `importSnapshotInstance()`

```php
importSnapshotInstance($id, $inline_object93): \OpenAPI\Client\Model\Model200Success
```

Import Snapshot of an Instance

It is possible to import a snapshot of an instance. This creates a Virtual Image of the instance as it currently exists.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object93 = new \OpenAPI\Client\Model\InlineObject93(); // \OpenAPI\Client\Model\InlineObject93

try {
    $result = $apiInstance->importSnapshotInstance($id, $inline_object93);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->importSnapshotInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object93** | [**\OpenAPI\Client\Model\InlineObject93**](../Model/InlineObject93.md)|  | [optional]

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

## `linkedCloneSnapshotInstance()`

```php
linkedCloneSnapshotInstance($id, $snapshot_id): \OpenAPI\Client\Model\Model200Success
```

Create Linked Clone of Instance Snapshot

It is possible to create a linked clone of an Instance Snapshot.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$snapshot_id = 4; // float | Snapshot ID

try {
    $result = $apiInstance->linkedCloneSnapshotInstance($id, $snapshot_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->linkedCloneSnapshotInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **snapshot_id** | **float**| Snapshot ID |

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

## `listInstanceServicePlans()`

```php
listInstanceServicePlans($zone_id, $layout_id, $site_id): \OpenAPI\Client\Model\InlineResponse20060
```

Get Available Service Plans for an Instance

This endpoint retrieves all the Service Plans available for the specified cloud and instance layout. The response includes details about the plans and their configuration options. It may be used to get the list of available plans when creating a new instance or resizing an existing instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$zone_id = 3; // int | The Zone ID for Filtering
$layout_id = 98; // int | The Layout ID for Filtering
$site_id = 7; // int | The Site ID for Filtering

try {
    $result = $apiInstance->listInstanceServicePlans($zone_id, $layout_id, $site_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->listInstanceServicePlans: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **layout_id** | **int**| The Layout ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20060**](../Model/InlineResponse20060.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listInstanceTypesProvisioning()`

```php
listInstanceTypesProvisioning($max, $offset, $sort, $direction, $phrase, $name, $code, $featured, $details): object
```

Get All Instance Types for Provisioning

Fetch the list of available instance types. These can vary in range from database containers, to web containers, to custom containers.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
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
$featured = false; // bool | Filter by featured
$details = true; // bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default.

try {
    $result = $apiInstance->listInstanceTypesProvisioning($max, $offset, $sort, $direction, $phrase, $name, $code, $featured, $details);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->listInstanceTypesProvisioning: ', $e->getMessage(), PHP_EOL;
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
 **featured** | **bool**| Filter by featured | [optional]
 **details** | **bool**| Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | [optional]

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

## `listInstances()`

```php
listInstances($max, $offset, $name, $phrase, $instance_type, $last_updated, $created_by, $agent_installed, $status, $environment, $show_deleted, $deleted, $expire_date, $expire_date_min, $expire_days, $expire_days_min, $shutdown_date, $shutdown_date_min, $shutdown_days, $shutdown_days_min, $labels, $all_labels, $tags, $metadata, $details): object
```

Get All Instances

This endpoint retrieves a paginated list of instances.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$instance_type = ubuntu; // string | The Instance Type Code for Filtering
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$created_by = 63; // int | The User ID for Filtering
$agent_installed = true; // bool | Filter instances by if agent is installed or not on the associated servers.
$status = running; // string | The instance status for filtering.
$environment = lab; // string | The environment for filtering.
$show_deleted = true; // bool | If true, includes instances in pending removal status.
$deleted = true; // bool | If true, only deleted resources or instances in pending removal status are returned.
$expire_date = 2019-01-01; // string | Filter by expireDate less than or equal to specified date
$expire_date_min = 2019-01-01; // string | Filter expireDate greater than or equal to the specified date
$expire_days = 20; // string | Filter by expireDays less than or equal to the specified value
$expire_days_min = 20; // string | Filter by expireDays greater than or equal to the specified value
$shutdown_date = 2019-01-01; // string | Filter by shutdownDate less than equal to the specified date
$shutdown_date_min = 2019-01-01; // string | Filter by shutdownDate greater than or equal to the specified date
$shutdown_days = 20; // string | Filter by shutdownDays less than or equal to the specified value
$shutdown_days_min = 20; // string | Filter by shutdownDays greater than or equal to the specified value
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$tags = tags.env=qa&tags.env=test; // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)
$metadata = tags.env=qa&tags.env=test; // string | Alias for tags
$details = true; // bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2.

try {
    $result = $apiInstance->listInstances($max, $offset, $name, $phrase, $instance_type, $last_updated, $created_by, $agent_installed, $status, $environment, $show_deleted, $deleted, $expire_date, $expire_date_min, $expire_days, $expire_days_min, $shutdown_date, $shutdown_date_min, $shutdown_days, $shutdown_days_min, $labels, $all_labels, $tags, $metadata, $details);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->listInstances: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **instance_type** | **string**| The Instance Type Code for Filtering | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **created_by** | **int**| The User ID for Filtering | [optional]
 **agent_installed** | **bool**| Filter instances by if agent is installed or not on the associated servers. | [optional]
 **status** | **string**| The instance status for filtering. | [optional]
 **environment** | **string**| The environment for filtering. | [optional]
 **show_deleted** | **bool**| If true, includes instances in pending removal status. | [optional] [default to false]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]
 **expire_date** | **string**| Filter by expireDate less than or equal to specified date | [optional]
 **expire_date_min** | **string**| Filter expireDate greater than or equal to the specified date | [optional]
 **expire_days** | **string**| Filter by expireDays less than or equal to the specified value | [optional]
 **expire_days_min** | **string**| Filter by expireDays greater than or equal to the specified value | [optional]
 **shutdown_date** | **string**| Filter by shutdownDate less than equal to the specified date | [optional]
 **shutdown_date_min** | **string**| Filter by shutdownDate greater than or equal to the specified date | [optional]
 **shutdown_days** | **string**| Filter by shutdownDays less than or equal to the specified value | [optional]
 **shutdown_days_min** | **string**| Filter by shutdownDays greater than or equal to the specified value | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **tags** | **string**| Filter by tags (metadata). This allows filtering by a tag name and value(s) | [optional]
 **metadata** | **string**| Alias for tags | [optional]
 **details** | **bool**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] [default to false]

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

## `listSecurityGroupsInstance()`

```php
listSecurityGroupsInstance($id): object
```

Get Security Groups for an Instance

This returns a list of all of the security groups applied to an instance and whether the firewall is enabled.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->listSecurityGroupsInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->listSecurityGroupsInstance: ', $e->getMessage(), PHP_EOL;
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

## `lockInstance()`

```php
lockInstance($id): \OpenAPI\Client\Model\Model200Success
```

Lock an Instance

This will lock the instance. While locked, instances may not be removed.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->lockInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->lockInstance: ', $e->getMessage(), PHP_EOL;
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

## `refreshStateInstance()`

```php
refreshStateInstance($id): \OpenAPI\Client\Model\Model200Success
```

Refresh State of an Instance

This endpoint provides a way to refresh the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->refreshStateInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->refreshStateInstance: ', $e->getMessage(), PHP_EOL;
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

## `removeInstancesFromControl()`

```php
removeInstancesFromControl($inline_object99): \OpenAPI\Client\Model\SuccessMessage
```

Remove From Control

Will delete a brownfield instance (or instances) asynchronously (Only deletes records local to Morpheus, actual VMs remain unchanged).

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object99 = new \OpenAPI\Client\Model\InlineObject99(); // \OpenAPI\Client\Model\InlineObject99

try {
    $result = $apiInstance->removeInstancesFromControl($inline_object99);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->removeInstancesFromControl: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object99** | [**\OpenAPI\Client\Model\InlineObject99**](../Model/InlineObject99.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessMessage**](../Model/SuccessMessage.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `resizeInstance()`

```php
resizeInstance($id, $instance_resize): object
```

Resize an Instance

It is possible to resize containers within an instance by increasing their memory plan or storage limit. This is done by assigning a new service plan to the container. This endpoint also allows for NIC reconfiguration by passing a new array of networkInterfaces

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$instance_resize = {"instance":{"id":1,"plan":{"id":15}},"volumes":[{"id":-1,"rootVolume":true,"name":"root","size":20,"sizeId":null,"storageType":null,"datastoreId":null}],"deleteOriginalVolumes":true,"servicePlanOptions":{"maxCores":1,"coresPerSocket":1,"maxMemory":536870912},"networkInterfaces":[{"id":534,"name":"eth0","network":{"id":"network-20"}},{"name":"eth1","network":{"id":"network-20"}}]}; // \OpenAPI\Client\Model\InstanceResize

try {
    $result = $apiInstance->resizeInstance($id, $instance_resize);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->resizeInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_resize** | [**\OpenAPI\Client\Model\InstanceResize**](../Model/InstanceResize.md)|  | [optional]

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

## `restartInstance()`

```php
restartInstance($id): object
```

Restart an instance

This will restart all containers running within an instance. This includes rebuilding the environment variables and applying settings to the docker containers.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->restartInstance: ', $e->getMessage(), PHP_EOL;
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

## `revertSnapshotInstance()`

```php
revertSnapshotInstance($id, $snapshot_id): \OpenAPI\Client\Model\Model200Success
```

Revert Instance to Snapshot

It is possible to restore an Instance to a snapshot.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$snapshot_id = 4; // float | Snapshot ID

try {
    $result = $apiInstance->revertSnapshotInstance($id, $snapshot_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->revertSnapshotInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **snapshot_id** | **float**| Snapshot ID |

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

## `runWorkflowInstance()`

```php
runWorkflowInstance($id, $workflow_id, $workflow_name, $instance_workflow): \OpenAPI\Client\Model\Model200Success
```

Run Workflow on an Instance

This will run a provisioning workflow on all containers in an instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$workflow_id = 3; // int | ID of the workflow to execute
$workflow_name = myworkflow; // string | Name of the workflow to execute
$instance_workflow = {"taskSet":{"customOptions":{"foo":"bar"}}}; // \OpenAPI\Client\Model\InstanceWorkflow

try {
    $result = $apiInstance->runWorkflowInstance($id, $workflow_id, $workflow_name, $instance_workflow);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->runWorkflowInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **workflow_id** | **int**| ID of the workflow to execute | [optional]
 **workflow_name** | **string**| Name of the workflow to execute | [optional]
 **instance_workflow** | [**\OpenAPI\Client\Model\InstanceWorkflow**](../Model/InstanceWorkflow.md)|  | [optional]

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

## `setApplyInstance()`

```php
setApplyInstance($id, $inline_object91): \OpenAPI\Client\Model\Model200Success
```

Apply State of an Instance

This endpoint provides a way to apply the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object91 = new \OpenAPI\Client\Model\InlineObject91(); // \OpenAPI\Client\Model\InlineObject91

try {
    $result = $apiInstance->setApplyInstance($id, $inline_object91);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->setApplyInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object91** | [**\OpenAPI\Client\Model\InlineObject91**](../Model/InlineObject91.md)|  | [optional]

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

## `setInstanceSecurityGroups()`

```php
setInstanceSecurityGroups($id, $inline_object94): object
```

Set Security Groups for an Instance

Set Security Groups for an Instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object94 = new \OpenAPI\Client\Model\InlineObject94(); // \OpenAPI\Client\Model\InlineObject94

try {
    $result = $apiInstance->setInstanceSecurityGroups($id, $inline_object94);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->setInstanceSecurityGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object94** | [**\OpenAPI\Client\Model\InlineObject94**](../Model/InlineObject94.md)|  | [optional]

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

## `snapshotInstance()`

```php
snapshotInstance($id, $instance_snapshot): \OpenAPI\Client\Model\Model200Success
```

Snapshot an Instance

This endpoint will create a snapshot of an instance. This is done asychronously, so the ID of the snapshot is not returned.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$instance_snapshot = {"snapshot":{"name":"snapshot-test","description":"A snapshot created via the Morpheus api"}}; // \OpenAPI\Client\Model\InstanceSnapshot

try {
    $result = $apiInstance->snapshotInstance($id, $instance_snapshot);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->snapshotInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_snapshot** | [**\OpenAPI\Client\Model\InstanceSnapshot**](../Model/InstanceSnapshot.md)|  | [optional]

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

## `snapshotsInstance()`

```php
snapshotsInstance($id): \OpenAPI\Client\Model\InstanceSnapshots
```

Get list of snapshots for an Instance

Get list of snapshots for an Instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->snapshotsInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->snapshotsInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InstanceSnapshots**](../Model/InstanceSnapshots.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `startInstance()`

```php
startInstance($id): object
```

Start an instance

This will start all containers running within an instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->startInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->startInstance: ', $e->getMessage(), PHP_EOL;
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

## `stopInstance()`

```php
stopInstance($id): object
```

Stop an instance

This will stop all containers running within an instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->stopInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->stopInstance: ', $e->getMessage(), PHP_EOL;
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

## `suspendInstance()`

```php
suspendInstance($id): \OpenAPI\Client\Model\Model200Success
```

Suspend an instance

This will suspend all containers in the instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->suspendInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->suspendInstance: ', $e->getMessage(), PHP_EOL;
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

## `unlockInstance()`

```php
unlockInstance($id): \OpenAPI\Client\Model\Model200Success
```

Unlock an Instance

This will unlock the instance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->unlockInstance($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->unlockInstance: ', $e->getMessage(), PHP_EOL;
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

## `updateInstance()`

```php
updateInstance($id, $instance_update): object
```

Updating an Instance

Updating an Instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$instance_update = {"instance":{"description":"my new redis","addTags":[{"name":"hello","value":"world"},{"name":"flash","value":"bang"}],"removeTags":[{"name":"oldTag"}]},"config":{"customOptions":{"dbfoldername":"data"}}}; // \OpenAPI\Client\Model\InstanceUpdate

try {
    $result = $apiInstance->updateInstance($id, $instance_update);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->updateInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_update** | [**\OpenAPI\Client\Model\InstanceUpdate**](../Model/InstanceUpdate.md)|  | [optional]

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

## `updateInstanceNetworkInterface()`

```php
updateInstanceNetworkInterface($id, $network_interface_id, $network_interface_update): object
```

Updating a label for an Instance's Network

Updating an Instance's Network's Label

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$network_interface_id = 7; // float | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
$network_interface_update = {"name":"newLabelName"}; // \OpenAPI\Client\Model\NetworkInterfaceUpdate

try {
    $result = $apiInstance->updateInstanceNetworkInterface($id, $network_interface_id, $network_interface_update);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->updateInstanceNetworkInterface: ', $e->getMessage(), PHP_EOL;
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

## `updateInstanceSchedule()`

```php
updateInstanceSchedule($id, $schedule_id, $inline_object97): object
```

Updating an Instance Schedule

This endpoint provides updating of an instance schedule

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 9; // int | Instance ID
$schedule_id = 1; // int | Instance Schedule ID
$inline_object97 = new \OpenAPI\Client\Model\InlineObject97(); // \OpenAPI\Client\Model\InlineObject97

try {
    $result = $apiInstance->updateInstanceSchedule($id, $schedule_id, $inline_object97);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->updateInstanceSchedule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **schedule_id** | **int**| Instance Schedule ID |
 **inline_object97** | [**\OpenAPI\Client\Model\InlineObject97**](../Model/InlineObject97.md)|  | [optional]

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

## `updateInstanceThreshold()`

```php
updateInstanceThreshold($id, $inline_object95): object
```

Updates an Instance Scale Threshold

Updates the scale threshold settings for a specific instance

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object95 = new \OpenAPI\Client\Model\InlineObject95(); // \OpenAPI\Client\Model\InlineObject95

try {
    $result = $apiInstance->updateInstanceThreshold($id, $inline_object95);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->updateInstanceThreshold: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object95** | [**\OpenAPI\Client\Model\InlineObject95**](../Model/InlineObject95.md)|  | [optional]

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

## `updateWikiInstance()`

```php
updateWikiInstance($id, $inline_object270): object
```

Update an Instance Wiki Page

Updates an instance Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\InstancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object270 = new \OpenAPI\Client\Model\InlineObject270(); // \OpenAPI\Client\Model\InlineObject270

try {
    $result = $apiInstance->updateWikiInstance($id, $inline_object270);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling InstancesApi->updateWikiInstance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object270** | [**\OpenAPI\Client\Model\InlineObject270**](../Model/InlineObject270.md)|  | [optional]

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
