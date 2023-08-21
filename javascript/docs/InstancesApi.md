# MorpheusApi.InstancesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addInstance**](InstancesApi.md#addInstance) | **POST** /api/instances | Create an Instance
[**backupInstance**](InstancesApi.md#backupInstance) | **PUT** /api/instances/{id}/backup | Backup an instance
[**backupsInstance**](InstancesApi.md#backupsInstance) | **GET** /api/instances/{id}/backups | Get list of backups for an Instance
[**cancelExpirationInstance**](InstancesApi.md#cancelExpirationInstance) | **PUT** /api/instances/{id}/cancel-expiration | Cancel Expiration of an Instance
[**cancelRemovalInstance**](InstancesApi.md#cancelRemovalInstance) | **PUT** /api/instances/{id}/cancel-removal | Cancel Removal of an Instance
[**cancelShutdownInstance**](InstancesApi.md#cancelShutdownInstance) | **PUT** /api/instances/{id}/cancel-shutdown | Cancel Shutdown of an Instance
[**cloneImageInstance**](InstancesApi.md#cloneImageInstance) | **PUT** /api/instances/{id}/clone-image | Clone to Image
[**cloneInstance**](InstancesApi.md#cloneInstance) | **PUT** /api/instances/{id}/clone | Clone an Instance
[**createInstanceSchedule**](InstancesApi.md#createInstanceSchedule) | **POST** /api/instances/{id}/schedules | Create a new Instance Schedule
[**deleteAllSnapshotsInstance**](InstancesApi.md#deleteAllSnapshotsInstance) | **DELETE** /api/instances/{id}/delete-all-snapshots | Delete All Snapshots of Instance
[**deleteAllSnapshotsInstanceContainer**](InstancesApi.md#deleteAllSnapshotsInstanceContainer) | **DELETE** /api/instances/{id}/delete-container-snapshots/{containerId} | Delete All Snapshots of Instance Container
[**deleteInstance**](InstancesApi.md#deleteInstance) | **DELETE** /api/instances/{id} | Delete an instance
[**deleteInstanceSchedule**](InstancesApi.md#deleteInstanceSchedule) | **DELETE** /api/instances/{id}/schedules/{scheduleId} | Deletes an Instance Schedule
[**deleteSnapshotInstance**](InstancesApi.md#deleteSnapshotInstance) | **DELETE** /api/snapshots/{id} | Delete Snapshot of Instance
[**ejectInstance**](InstancesApi.md#ejectInstance) | **PUT** /api/instances/{id}/eject | Eject an instance
[**extendExpirationInstance**](InstancesApi.md#extendExpirationInstance) | **PUT** /api/instances/{id}/extend-expiration | Extend Expiration of an Instance
[**extendShutdownInstance**](InstancesApi.md#extendShutdownInstance) | **PUT** /api/instances/{id}/extend-shutdown | Extend Shutdown of an Instance
[**getEnvVariables**](InstancesApi.md#getEnvVariables) | **GET** /api/instances/{id}/envs | Get Env Variables
[**getInstance**](InstancesApi.md#getInstance) | **GET** /api/instances/{id} | Retrieves a Specific Instance
[**getInstanceContainers**](InstancesApi.md#getInstanceContainers) | **GET** /api/instances/{id}/containers | Get Container Details
[**getInstanceHistory**](InstancesApi.md#getInstanceHistory) | **GET** /api/instances/{id}/history | Get Instance History
[**getInstanceSchedule**](InstancesApi.md#getInstanceSchedule) | **GET** /api/instances/{id}/schedules/{scheduleId} | Get a Specific Instance Schedule
[**getInstanceSchedules**](InstancesApi.md#getInstanceSchedules) | **GET** /api/instances/{id}/schedules | Get all Instance Schedules
[**getInstanceThreshold**](InstancesApi.md#getInstanceThreshold) | **GET** /api/instances/{id}/threshold | Get an Instance Scale Threshold
[**getInstanceTypeProvisioning**](InstancesApi.md#getInstanceTypeProvisioning) | **GET** /api/instance-types/{id} | Get Specific Instance Type for Provisioning
[**getPrepareApplyInstance**](InstancesApi.md#getPrepareApplyInstance) | **GET** /api/instances/{id}/prepare-apply | Prepare To Apply an Instance
[**getSnapshotInstance**](InstancesApi.md#getSnapshotInstance) | **GET** /api/snapshots/{id} | Get a Specific Snapshot
[**getStateInstance**](InstancesApi.md#getStateInstance) | **GET** /api/instances/{id}/state | Get State of an Instance
[**getValidateApplyInstance**](InstancesApi.md#getValidateApplyInstance) | **POST** /api/instances/{id}/validate-apply | Validate Apply State for an Instance
[**getWikiInstance**](InstancesApi.md#getWikiInstance) | **GET** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**importSnapshotInstance**](InstancesApi.md#importSnapshotInstance) | **PUT** /api/instances/{id}/import-snapshot | Import Snapshot of an Instance
[**linkedCloneSnapshotInstance**](InstancesApi.md#linkedCloneSnapshotInstance) | **PUT** /api/instances/{id}/linked-clone/{snapshotId} | Create Linked Clone of Instance Snapshot
[**listInstanceServicePlans**](InstancesApi.md#listInstanceServicePlans) | **GET** /api/instances/service-plans | Get Available Service Plans for an Instance
[**listInstanceTypesProvisioning**](InstancesApi.md#listInstanceTypesProvisioning) | **GET** /api/instance-types | Get All Instance Types for Provisioning
[**listInstances**](InstancesApi.md#listInstances) | **GET** /api/instances | Get All Instances
[**listSecurityGroupsInstance**](InstancesApi.md#listSecurityGroupsInstance) | **GET** /api/instances/{id}/security-groups | Get Security Groups for an Instance
[**lockInstance**](InstancesApi.md#lockInstance) | **PUT** /api/instances/{id}/lock | Lock an Instance
[**refreshStateInstance**](InstancesApi.md#refreshStateInstance) | **POST** /api/instances/{id}/refresh | Refresh State of an Instance
[**removeInstancesFromControl**](InstancesApi.md#removeInstancesFromControl) | **POST** /api/instances/remove-from-control | Remove From Control
[**resizeInstance**](InstancesApi.md#resizeInstance) | **PUT** /api/instances/{id}/resize | Resize an Instance
[**restartInstance**](InstancesApi.md#restartInstance) | **PUT** /api/instances/{id}/restart | Restart an instance
[**revertSnapshotInstance**](InstancesApi.md#revertSnapshotInstance) | **PUT** /api/instances/{id}/revert-snapshot/{snapshotId} | Revert Instance to Snapshot
[**runWorkflowInstance**](InstancesApi.md#runWorkflowInstance) | **PUT** /api/instances/{id}/workflow | Run Workflow on an Instance
[**setApplyInstance**](InstancesApi.md#setApplyInstance) | **POST** /api/instances/{id}/apply | Apply State of an Instance
[**setInstanceSecurityGroups**](InstancesApi.md#setInstanceSecurityGroups) | **POST** /api/instances/{id}/security-groups | Set Security Groups for an Instance
[**snapshotInstance**](InstancesApi.md#snapshotInstance) | **PUT** /api/instances/{id}/snapshot | Snapshot an Instance
[**snapshotsInstance**](InstancesApi.md#snapshotsInstance) | **GET** /api/instances/{id}/snapshots | Get list of snapshots for an Instance
[**startInstance**](InstancesApi.md#startInstance) | **PUT** /api/instances/{id}/start | Start an instance
[**stopInstance**](InstancesApi.md#stopInstance) | **PUT** /api/instances/{id}/stop | Stop an instance
[**suspendInstance**](InstancesApi.md#suspendInstance) | **PUT** /api/instances/{id}/suspend | Suspend an instance
[**unlockInstance**](InstancesApi.md#unlockInstance) | **PUT** /api/instances/{id}/unlock | Unlock an Instance
[**updateInstance**](InstancesApi.md#updateInstance) | **PUT** /api/instances/{id} | Updating an Instance
[**updateInstanceNetworkInterface**](InstancesApi.md#updateInstanceNetworkInterface) | **PUT** /api/instances/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for an Instance&#39;s Network
[**updateInstanceSchedule**](InstancesApi.md#updateInstanceSchedule) | **PUT** /api/instances/{id}/schedules/{scheduleId} | Updating an Instance Schedule
[**updateInstanceThreshold**](InstancesApi.md#updateInstanceThreshold) | **PUT** /api/instances/{id}/threshold | Updates an Instance Scale Threshold
[**updateWikiInstance**](InstancesApi.md#updateWikiInstance) | **PUT** /api/instances/{id}/wiki | Update an Instance Wiki Page



## addInstance

> Object addInstance(opts)

Create an Instance

Create an Instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let opts = {
  'instanceCreate': {"zoneId":15,"instance":{"name":"test2","cloud":"vcenter1 - homelab","site":{"id":1},"type":"ubuntu","instanceType":{"code":"ubuntu"},"instanceContext":"dev","labels":["ubuntuserver"],"layout":{"id":1399,"code":"vmware-ubuntu-20.04-single"},"plan":{"id":361,"code":"vm-512","name":"1 CPU, 512MB Memory"},"userGroup":{"id":1}},"config":{"resourcePoolId":75,"hostId":"","vmwareFolderId":"group-v90002","nestedVirtualization":"off","createUser":true},"volumes":[{"id":-1,"rootVolume":true,"name":"root","size":10,"sizeId":null,"storageType":1,"datastoreId":59}],"networkInterfaces":[{"network":{"id":"network-110"},"networkInterfaceTypeId":4}],"layoutSize":1} // InstanceCreate | 
};
apiInstance.addInstance(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceCreate** | [**InstanceCreate**](InstanceCreate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## backupInstance

> Model200Success backupInstance(id)

Backup an instance

Backup an instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.backupInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## backupsInstance

> InstanceBackups backupsInstance(id)

Get list of backups for an Instance

Get list of backups for an Instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.backupsInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InstanceBackups**](InstanceBackups.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cancelExpirationInstance

> Model200Success cancelExpirationInstance(id)

Cancel Expiration of an Instance

This operation will cancel the expiration of an instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.cancelExpirationInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cancelRemovalInstance

> Model200Success cancelRemovalInstance(id)

Cancel Removal of an Instance

This operation will undo the delete of an instance that is pending removal.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.cancelRemovalInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cancelShutdownInstance

> Model200Success cancelShutdownInstance(id)

Cancel Shutdown of an Instance

This operation will cancel the shutdown of an instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.cancelShutdownInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## cloneImageInstance

> Model200Success cloneImageInstance(id, opts)

Clone to Image

This endpoint allows creating an image template from an existing instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'instancesCloneImage': {"templateName":"Example Image","zoneFolder":"group-v81"} // InstancesCloneImage | 
};
apiInstance.cloneImageInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **instancesCloneImage** | [**InstancesCloneImage**](InstancesCloneImage.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## cloneInstance

> Model200Success cloneInstance(id, opts)

Clone an Instance

One can easily clone an instance and all containers within that instance. The containers are backed up via the backup services and used as a snapshot to produce a clone of the instance. It is possible to clone this app instance into an entirely different availability zone.  This endpoint also supports all of the same parameters as instance creation, so you can override any configuration options when provisioning the clone. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'instanceClone': {"name":"New Name","group":{"id":1}} // InstanceClone | 
};
apiInstance.cloneInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **instanceClone** | [**InstanceClone**](InstanceClone.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createInstanceSchedule

> Object createInstanceSchedule(id, opts)

Create a new Instance Schedule

Create a new schedule for a specific instance.  This creates an instance scaling threshold that only applies during a defined schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 42; // Number | Instance ID
let opts = {
  'inlineObject96': new MorpheusApi.InlineObject96() // InlineObject96 | 
};
apiInstance.createInstanceSchedule(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Instance ID | 
 **inlineObject96** | [**InlineObject96**](InlineObject96.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteAllSnapshotsInstance

> Model200Success deleteAllSnapshotsInstance(id)

Delete All Snapshots of Instance

Delete All Snapshots attached to Instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteAllSnapshotsInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteAllSnapshotsInstanceContainer

> Model200Success deleteAllSnapshotsInstanceContainer(id, containerId)

Delete All Snapshots of Instance Container

Delete All Snapshots attached to Instance Container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let containerId = 4; // Number | Container ID
apiInstance.deleteAllSnapshotsInstanceContainer(id, containerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **containerId** | **Number**| Container ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteInstance

> Model200Success deleteInstance(id, opts)

Delete an instance

Will delete an instance and all associated monitors and backups.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'preserveVolumes': on, // String | Preserve Volumes
  'keepBackups': on, // String | Preserve copy of backups
  'releaseFloatingIps': off, // String | Release Floating IPs
  'releaseEIPs': off, // String | Alias for releaseFloatingIps
  'force': on // String | Force Delete
};
apiInstance.deleteInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **keepBackups** | **String**| Preserve copy of backups | [optional] [default to &#39;off&#39;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteInstanceSchedule

> Model200Success deleteInstanceSchedule(id, scheduleId)

Deletes an Instance Schedule

Deletes a specified instance scaling schedule 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 42; // Number | Instance ID
let scheduleId = 1; // Number | Instance Schedule ID
apiInstance.deleteInstanceSchedule(id, scheduleId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Instance ID | 
 **scheduleId** | **Number**| Instance Schedule ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteSnapshotInstance

> Model200Success deleteSnapshotInstance(id)

Delete Snapshot of Instance

Delete snapshot of instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteSnapshotInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## ejectInstance

> Model200Success ejectInstance(id)

Eject an instance

This will eject any ISO media on all containers in the instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.ejectInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## extendExpirationInstance

> Model200Success extendExpirationInstance(id)

Extend Expiration of an Instance

This operation will extend the expiration of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.extendExpirationInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## extendShutdownInstance

> Model200Success extendShutdownInstance(id)

Extend Shutdown of an Instance

This operation will extend the shutdown of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.extendShutdownInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getEnvVariables

> InlineResponse20057 getEnvVariables(id)

Get Env Variables

This gets all the environment variables associated with the instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getEnvVariables(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20057**](InlineResponse20057.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstance

> InlineResponse20056 getInstance(id, opts)

Retrieves a Specific Instance

Retrieves a specific instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'details': true // Boolean | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2.
};
apiInstance.getInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **details** | **Boolean**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] [default to false]

### Return type

[**InlineResponse20056**](InlineResponse20056.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceContainers

> Object getInstanceContainers(id)

Get Container Details

This can be valuable for evaluating the details of the compute server(s) running on an instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getInstanceContainers(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceHistory

> Object getInstanceHistory(id, opts)

Get Instance History

This endpoint retrieves the process history for a specific instance. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'containerId': 135, // Number | The Container ID for Filtering
  'serverId': 97, // Number | The Server ID for Filtering
  'zoneId': 3 // Number | The Zone ID for Filtering
};
apiInstance.getInstanceHistory(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **containerId** | **Number**| The Container ID for Filtering | [optional] 
 **serverId** | **Number**| The Server ID for Filtering | [optional] 
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceSchedule

> InlineResponse20059 getInstanceSchedule(id, scheduleId)

Get a Specific Instance Schedule

This endpoint retrieves a specific instance scaling schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 42; // Number | Instance ID
let scheduleId = 1; // Number | Instance Schedule ID
apiInstance.getInstanceSchedule(id, scheduleId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Instance ID | 
 **scheduleId** | **Number**| Instance Schedule ID | 

### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceSchedules

> Object getInstanceSchedules(id)

Get all Instance Schedules

This endpoint retrieves all the scaling threshold schedules for a specific instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 42; // Number | Instance ID
apiInstance.getInstanceSchedules(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Instance ID | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceThreshold

> InlineResponse20058 getInstanceThreshold(id)

Get an Instance Scale Threshold

Retrieves the scale threshold settings for a specific instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getInstanceThreshold(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20058**](InlineResponse20058.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceTypeProvisioning

> Object getInstanceTypeProvisioning(id)

Get Specific Instance Type for Provisioning

Fetch an instance type by ID. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getInstanceTypeProvisioning(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getPrepareApplyInstance

> Object getPrepareApplyInstance(id)

Prepare To Apply an Instance

This endpoint provides a way to view the current instance configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPrepareApplyInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getSnapshotInstance

> Snapshot getSnapshotInstance(id)

Get a Specific Snapshot

This endpoint retrieves a specific snapshot.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getSnapshotInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getStateInstance

> Object getStateInstance(id)

Get State of an Instance

This endpoint provides a way to view the state of an instance. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getStateInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getValidateApplyInstance

> Object getValidateApplyInstance(id, opts)

Validate Apply State for an Instance

This endpoint provides a way to validate instance configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject98': new MorpheusApi.InlineObject98() // InlineObject98 | 
};
apiInstance.getValidateApplyInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject98** | [**InlineObject98**](InlineObject98.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getWikiInstance

> InlineResponse200168 getWikiInstance(id)

Retrieves an Instance Wiki Page

This endpoint retrieves an instance Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getWikiInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## importSnapshotInstance

> Model200Success importSnapshotInstance(id, opts)

Import Snapshot of an Instance

It is possible to import a snapshot of an instance. This creates a Virtual Image of the instance as it currently exists.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject93': new MorpheusApi.InlineObject93() // InlineObject93 | 
};
apiInstance.importSnapshotInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject93** | [**InlineObject93**](InlineObject93.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## linkedCloneSnapshotInstance

> Model200Success linkedCloneSnapshotInstance(id, snapshotId)

Create Linked Clone of Instance Snapshot

It is possible to create a linked clone of an Instance Snapshot.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let snapshotId = 4; // Number | Snapshot ID
apiInstance.linkedCloneSnapshotInstance(id, snapshotId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **snapshotId** | **Number**| Snapshot ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listInstanceServicePlans

> InlineResponse20060 listInstanceServicePlans(opts)

Get Available Service Plans for an Instance

This endpoint retrieves all the Service Plans available for the specified cloud and instance layout. The response includes details about the plans and their configuration options. It may be used to get the list of available plans when creating a new instance or resizing an existing instance. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let opts = {
  'zoneId': 3, // Number | The Zone ID for Filtering
  'layoutId': 98, // Number | The Layout ID for Filtering
  'siteId': 7 // Number | The Site ID for Filtering
};
apiInstance.listInstanceServicePlans(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **layoutId** | **Number**| The Layout ID for Filtering | [optional] 
 **siteId** | **Number**| The Site ID for Filtering | [optional] 

### Return type

[**InlineResponse20060**](InlineResponse20060.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listInstanceTypesProvisioning

> Object listInstanceTypesProvisioning(opts)

Get All Instance Types for Provisioning

Fetch the list of available instance types. These can vary in range from database containers, to web containers, to custom containers. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'featured': false, // Boolean | Filter by featured
  'details': true // Boolean | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default.
};
apiInstance.listInstanceTypesProvisioning(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **featured** | **Boolean**| Filter by featured | [optional] 
 **details** | **Boolean**| Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listInstances

> Object listInstances(opts)

Get All Instances

This endpoint retrieves a paginated list of instances. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'instanceType': ubuntu, // String | The Instance Type Code for Filtering
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'createdBy': 63, // Number | The User ID for Filtering
  'agentInstalled': true, // Boolean | Filter instances by if agent is installed or not on the associated servers.
  'status': running, // String | The instance status for filtering.
  'environment': lab, // String | The environment for filtering.
  'showDeleted': true, // Boolean | If true, includes instances in pending removal status.
  'deleted': true, // Boolean | If true, only deleted resources or instances in pending removal status are returned.
  'expireDate': 2019-01-01, // String | Filter by expireDate less than or equal to specified date
  'expireDateMin': 2019-01-01, // String | Filter expireDate greater than or equal to the specified date
  'expireDays': 20, // String | Filter by expireDays less than or equal to the specified value
  'expireDaysMin': 20, // String | Filter by expireDays greater than or equal to the specified value
  'shutdownDate': 2019-01-01, // String | Filter by shutdownDate less than equal to the specified date
  'shutdownDateMin': 2019-01-01, // String | Filter by shutdownDate greater than or equal to the specified date
  'shutdownDays': 20, // String | Filter by shutdownDays less than or equal to the specified value
  'shutdownDaysMin': 20, // String | Filter by shutdownDays greater than or equal to the specified value
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'tags': tags.env=qa&tags.env=test, // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
  'metadata': tags.env=qa&tags.env=test, // String | Alias for tags
  'details': true // Boolean | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2.
};
apiInstance.listInstances(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **instanceType** | **String**| The Instance Type Code for Filtering | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **createdBy** | **Number**| The User ID for Filtering | [optional] 
 **agentInstalled** | **Boolean**| Filter instances by if agent is installed or not on the associated servers. | [optional] 
 **status** | **String**| The instance status for filtering. | [optional] 
 **environment** | **String**| The environment for filtering. | [optional] 
 **showDeleted** | **Boolean**| If true, includes instances in pending removal status. | [optional] [default to false]
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional] 
 **expireDate** | **String**| Filter by expireDate less than or equal to specified date | [optional] 
 **expireDateMin** | **String**| Filter expireDate greater than or equal to the specified date | [optional] 
 **expireDays** | **String**| Filter by expireDays less than or equal to the specified value | [optional] 
 **expireDaysMin** | **String**| Filter by expireDays greater than or equal to the specified value | [optional] 
 **shutdownDate** | **String**| Filter by shutdownDate less than equal to the specified date | [optional] 
 **shutdownDateMin** | **String**| Filter by shutdownDate greater than or equal to the specified date | [optional] 
 **shutdownDays** | **String**| Filter by shutdownDays less than or equal to the specified value | [optional] 
 **shutdownDaysMin** | **String**| Filter by shutdownDays greater than or equal to the specified value | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **tags** | **String**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional] 
 **metadata** | **String**| Alias for tags | [optional] 
 **details** | **Boolean**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSecurityGroupsInstance

> Object listSecurityGroupsInstance(id)

Get Security Groups for an Instance

This returns a list of all of the security groups applied to an instance and whether the firewall is enabled.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.listSecurityGroupsInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## lockInstance

> Model200Success lockInstance(id)

Lock an Instance

This will lock the instance. While locked, instances may not be removed.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.lockInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## refreshStateInstance

> Model200Success refreshStateInstance(id)

Refresh State of an Instance

This endpoint provides a way to refresh the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.refreshStateInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeInstancesFromControl

> SuccessMessage removeInstancesFromControl(opts)

Remove From Control

Will delete a brownfield instance (or instances) asynchronously (Only deletes records local to Morpheus, actual VMs remain unchanged).

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let opts = {
  'inlineObject99': new MorpheusApi.InlineObject99() // InlineObject99 | 
};
apiInstance.removeInstancesFromControl(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject99** | [**InlineObject99**](InlineObject99.md)|  | [optional] 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## resizeInstance

> Object resizeInstance(id, opts)

Resize an Instance

It is possible to resize containers within an instance by increasing their memory plan or storage limit. This is done by assigning a new service plan to the container. This endpoint also allows for NIC reconfiguration by passing a new array of networkInterfaces

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'instanceResize': {"instance":{"id":1,"plan":{"id":15}},"volumes":[{"id":-1,"rootVolume":true,"name":"root","size":20,"sizeId":null,"storageType":null,"datastoreId":null}],"deleteOriginalVolumes":true,"servicePlanOptions":{"maxCores":1,"coresPerSocket":1,"maxMemory":536870912},"networkInterfaces":[{"id":534,"name":"eth0","network":{"id":"network-20"}},{"name":"eth1","network":{"id":"network-20"}}]} // InstanceResize | 
};
apiInstance.resizeInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **instanceResize** | [**InstanceResize**](InstanceResize.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## restartInstance

> Object restartInstance(id)

Restart an instance

This will restart all containers running within an instance. This includes rebuilding the environment variables and applying settings to the docker containers.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## revertSnapshotInstance

> Model200Success revertSnapshotInstance(id, snapshotId)

Revert Instance to Snapshot

It is possible to restore an Instance to a snapshot.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let snapshotId = 4; // Number | Snapshot ID
apiInstance.revertSnapshotInstance(id, snapshotId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **snapshotId** | **Number**| Snapshot ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## runWorkflowInstance

> Model200Success runWorkflowInstance(id, opts)

Run Workflow on an Instance

This will run a provisioning workflow on all containers in an instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'workflowId': 3, // Number | ID of the workflow to execute
  'workflowName': myworkflow, // String | Name of the workflow to execute
  'instanceWorkflow': {"taskSet":{"customOptions":{"foo":"bar"}}} // InstanceWorkflow | 
};
apiInstance.runWorkflowInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **workflowId** | **Number**| ID of the workflow to execute | [optional] 
 **workflowName** | **String**| Name of the workflow to execute | [optional] 
 **instanceWorkflow** | [**InstanceWorkflow**](InstanceWorkflow.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## setApplyInstance

> Model200Success setApplyInstance(id, opts)

Apply State of an Instance

This endpoint provides a way to apply the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject91': new MorpheusApi.InlineObject91() // InlineObject91 | 
};
apiInstance.setApplyInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject91** | [**InlineObject91**](InlineObject91.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## setInstanceSecurityGroups

> Object setInstanceSecurityGroups(id, opts)

Set Security Groups for an Instance

Set Security Groups for an Instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject94': new MorpheusApi.InlineObject94() // InlineObject94 | 
};
apiInstance.setInstanceSecurityGroups(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject94** | [**InlineObject94**](InlineObject94.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## snapshotInstance

> Model200Success snapshotInstance(id, opts)

Snapshot an Instance

This endpoint will create a snapshot of an instance. This is done asychronously, so the ID of the snapshot is not returned.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'instanceSnapshot': {"snapshot":{"name":"snapshot-test","description":"A snapshot created via the Morpheus api"}} // InstanceSnapshot | 
};
apiInstance.snapshotInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **instanceSnapshot** | [**InstanceSnapshot**](InstanceSnapshot.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## snapshotsInstance

> InstanceSnapshots snapshotsInstance(id)

Get list of snapshots for an Instance

Get list of snapshots for an Instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.snapshotsInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InstanceSnapshots**](InstanceSnapshots.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## startInstance

> Object startInstance(id)

Start an instance

This will start all containers running within an instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.startInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## stopInstance

> Object stopInstance(id)

Stop an instance

This will stop all containers running within an instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.stopInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## suspendInstance

> Model200Success suspendInstance(id)

Suspend an instance

This will suspend all containers in the instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.suspendInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## unlockInstance

> Model200Success unlockInstance(id)

Unlock an Instance

This will unlock the instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.unlockInstance(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateInstance

> Object updateInstance(id, opts)

Updating an Instance

Updating an Instance

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'instanceUpdate': {"instance":{"description":"my new redis","addTags":[{"name":"hello","value":"world"},{"name":"flash","value":"bang"}],"removeTags":[{"name":"oldTag"}]},"config":{"customOptions":{"dbfoldername":"data"}}} // InstanceUpdate | 
};
apiInstance.updateInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **instanceUpdate** | [**InstanceUpdate**](InstanceUpdate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateInstanceNetworkInterface

> Object updateInstanceNetworkInterface(id, networkInterfaceId, opts)

Updating a label for an Instance&#39;s Network

Updating an Instance&#39;s Network&#39;s Label

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let networkInterfaceId = 7; // Number | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
let opts = {
  'networkInterfaceUpdate': {"name":"newLabelName"} // NetworkInterfaceUpdate | 
};
apiInstance.updateInstanceNetworkInterface(id, networkInterfaceId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **networkInterfaceId** | **Number**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced | 
 **networkInterfaceUpdate** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateInstanceSchedule

> Object updateInstanceSchedule(id, scheduleId, opts)

Updating an Instance Schedule

This endpoint provides updating of an instance schedule

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 9; // Number | Instance ID
let scheduleId = 1; // Number | Instance Schedule ID
let opts = {
  'inlineObject97': new MorpheusApi.InlineObject97() // InlineObject97 | 
};
apiInstance.updateInstanceSchedule(id, scheduleId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Instance ID | 
 **scheduleId** | **Number**| Instance Schedule ID | 
 **inlineObject97** | [**InlineObject97**](InlineObject97.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateInstanceThreshold

> Object updateInstanceThreshold(id, opts)

Updates an Instance Scale Threshold

Updates the scale threshold settings for a specific instance 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject95': new MorpheusApi.InlineObject95() // InlineObject95 | 
};
apiInstance.updateInstanceThreshold(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject95** | [**InlineObject95**](InlineObject95.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateWikiInstance

> Object updateWikiInstance(id, opts)

Update an Instance Wiki Page

Updates an instance Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InstancesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject270': new MorpheusApi.InlineObject270() // InlineObject270 | 
};
apiInstance.updateWikiInstance(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject270** | [**InlineObject270**](InlineObject270.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

