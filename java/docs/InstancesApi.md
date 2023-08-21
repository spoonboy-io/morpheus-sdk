# InstancesApi

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


<a name="addInstance"></a>
# **addInstance**
> Object addInstance(instanceCreate)

Create an Instance

Create an Instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    InstanceCreate instanceCreate = {"zoneId":15,"instance":{"name":"test2","cloud":"vcenter1 - homelab","site":{"id":1},"type":"ubuntu","instanceType":{"code":"ubuntu"},"instanceContext":"dev","labels":["ubuntuserver"],"layout":{"id":1399,"code":"vmware-ubuntu-20.04-single"},"plan":{"id":361,"code":"vm-512","name":"1 CPU, 512MB Memory"},"userGroup":{"id":1}},"config":{"resourcePoolId":75,"hostId":"","vmwareFolderId":"group-v90002","nestedVirtualization":"off","createUser":true},"volumes":[{"id":-1,"rootVolume":true,"name":"root","size":10,"sizeId":null,"storageType":1,"datastoreId":59}],"networkInterfaces":[{"network":{"id":"network-110"},"networkInterfaceTypeId":4}],"layoutSize":1}; // InstanceCreate | 
    try {
      Object result = apiInstance.addInstance(instanceCreate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#addInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="backupInstance"></a>
# **backupInstance**
> Model200Success backupInstance(id)

Backup an instance

Backup an instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.backupInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#backupInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="backupsInstance"></a>
# **backupsInstance**
> InstanceBackups backupsInstance(id)

Get list of backups for an Instance

Get list of backups for an Instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InstanceBackups result = apiInstance.backupsInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#backupsInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InstanceBackups**](InstanceBackups.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="cancelExpirationInstance"></a>
# **cancelExpirationInstance**
> Model200Success cancelExpirationInstance(id)

Cancel Expiration of an Instance

This operation will cancel the expiration of an instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.cancelExpirationInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#cancelExpirationInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="cancelRemovalInstance"></a>
# **cancelRemovalInstance**
> Model200Success cancelRemovalInstance(id)

Cancel Removal of an Instance

This operation will undo the delete of an instance that is pending removal.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.cancelRemovalInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#cancelRemovalInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="cancelShutdownInstance"></a>
# **cancelShutdownInstance**
> Model200Success cancelShutdownInstance(id)

Cancel Shutdown of an Instance

This operation will cancel the shutdown of an instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.cancelShutdownInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#cancelShutdownInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="cloneImageInstance"></a>
# **cloneImageInstance**
> Model200Success cloneImageInstance(id, instancesCloneImage)

Clone to Image

This endpoint allows creating an image template from an existing instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InstancesCloneImage instancesCloneImage = {"templateName":"Example Image","zoneFolder":"group-v81"}; // InstancesCloneImage | 
    try {
      Model200Success result = apiInstance.cloneImageInstance(id, instancesCloneImage);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#cloneImageInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **instancesCloneImage** | [**InstancesCloneImage**](InstancesCloneImage.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="cloneInstance"></a>
# **cloneInstance**
> Model200Success cloneInstance(id, instanceClone)

Clone an Instance

One can easily clone an instance and all containers within that instance. The containers are backed up via the backup services and used as a snapshot to produce a clone of the instance. It is possible to clone this app instance into an entirely different availability zone.  This endpoint also supports all of the same parameters as instance creation, so you can override any configuration options when provisioning the clone. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InstanceClone instanceClone = {"name":"New Name","group":{"id":1}}; // InstanceClone | 
    try {
      Model200Success result = apiInstance.cloneInstance(id, instanceClone);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#cloneInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **instanceClone** | [**InstanceClone**](InstanceClone.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createInstanceSchedule"></a>
# **createInstanceSchedule**
> Object createInstanceSchedule(id, inlineObject96)

Create a new Instance Schedule

Create a new schedule for a specific instance.  This creates an instance scaling threshold that only applies during a defined schedule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 42L; // Long | Instance ID
    InlineObject96 inlineObject96 = new InlineObject96(); // InlineObject96 | 
    try {
      Object result = apiInstance.createInstanceSchedule(id, inlineObject96);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#createInstanceSchedule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Instance ID |
 **inlineObject96** | [**InlineObject96**](InlineObject96.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteAllSnapshotsInstance"></a>
# **deleteAllSnapshotsInstance**
> Model200Success deleteAllSnapshotsInstance(id)

Delete All Snapshots of Instance

Delete All Snapshots attached to Instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteAllSnapshotsInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#deleteAllSnapshotsInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteAllSnapshotsInstanceContainer"></a>
# **deleteAllSnapshotsInstanceContainer**
> Model200Success deleteAllSnapshotsInstanceContainer(id, containerId)

Delete All Snapshots of Instance Container

Delete All Snapshots attached to Instance Container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal containerId = new BigDecimal(78); // BigDecimal | Container ID
    try {
      Model200Success result = apiInstance.deleteAllSnapshotsInstanceContainer(id, containerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#deleteAllSnapshotsInstanceContainer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **containerId** | **BigDecimal**| Container ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteInstance"></a>
# **deleteInstance**
> Model200Success deleteInstance(id, preserveVolumes, keepBackups, releaseFloatingIps, releaseEIPs, force)

Delete an instance

Will delete an instance and all associated monitors and backups.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String preserveVolumes = "\"off\""; // String | Preserve Volumes
    String keepBackups = "\"off\""; // String | Preserve copy of backups
    String releaseFloatingIps = "\"on\""; // String | Release Floating IPs
    String releaseEIPs = "\"on\""; // String | Alias for releaseFloatingIps
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteInstance(id, preserveVolumes, keepBackups, releaseFloatingIps, releaseEIPs, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#deleteInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &quot;off&quot;]
 **keepBackups** | **String**| Preserve copy of backups | [optional] [default to &quot;off&quot;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &quot;on&quot;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &quot;on&quot;]
 **force** | **String**| Force Delete | [optional] [default to &quot;off&quot;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteInstanceSchedule"></a>
# **deleteInstanceSchedule**
> Model200Success deleteInstanceSchedule(id, scheduleId)

Deletes an Instance Schedule

Deletes a specified instance scaling schedule 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 42L; // Long | Instance ID
    Long scheduleId = 1L; // Long | Instance Schedule ID
    try {
      Model200Success result = apiInstance.deleteInstanceSchedule(id, scheduleId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#deleteInstanceSchedule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Instance ID |
 **scheduleId** | **Long**| Instance Schedule ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteSnapshotInstance"></a>
# **deleteSnapshotInstance**
> Model200Success deleteSnapshotInstance(id)

Delete Snapshot of Instance

Delete snapshot of instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteSnapshotInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#deleteSnapshotInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="ejectInstance"></a>
# **ejectInstance**
> Model200Success ejectInstance(id)

Eject an instance

This will eject any ISO media on all containers in the instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.ejectInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#ejectInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="extendExpirationInstance"></a>
# **extendExpirationInstance**
> Model200Success extendExpirationInstance(id)

Extend Expiration of an Instance

This operation will extend the expiration of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.extendExpirationInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#extendExpirationInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="extendShutdownInstance"></a>
# **extendShutdownInstance**
> Model200Success extendShutdownInstance(id)

Extend Shutdown of an Instance

This operation will extend the shutdown of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.extendShutdownInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#extendShutdownInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getEnvVariables"></a>
# **getEnvVariables**
> InlineResponse20057 getEnvVariables(id)

Get Env Variables

This gets all the environment variables associated with the instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20057 result = apiInstance.getEnvVariables(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getEnvVariables");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20057**](InlineResponse20057.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstance"></a>
# **getInstance**
> InlineResponse20056 getInstance(id, details)

Retrieves a Specific Instance

Retrieves a specific instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Boolean details = false; // Boolean | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2.
    try {
      InlineResponse20056 result = apiInstance.getInstance(id, details);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **details** | **Boolean**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] [default to false]

### Return type

[**InlineResponse20056**](InlineResponse20056.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstanceContainers"></a>
# **getInstanceContainers**
> Object getInstanceContainers(id)

Get Container Details

This can be valuable for evaluating the details of the compute server(s) running on an instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getInstanceContainers(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstanceContainers");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Instance Containers Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstanceHistory"></a>
# **getInstanceHistory**
> Object getInstanceHistory(id, phrase, containerId, serverId, zoneId)

Get Instance History

This endpoint retrieves the process history for a specific instance. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long containerId = 135L; // Long | The Container ID for Filtering
    Long serverId = 97L; // Long | The Server ID for Filtering
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    try {
      Object result = apiInstance.getInstanceHistory(id, phrase, containerId, serverId, zoneId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstanceHistory");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **containerId** | **Long**| The Container ID for Filtering | [optional]
 **serverId** | **Long**| The Server ID for Filtering | [optional]
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstanceSchedule"></a>
# **getInstanceSchedule**
> InlineResponse20059 getInstanceSchedule(id, scheduleId)

Get a Specific Instance Schedule

This endpoint retrieves a specific instance scaling schedule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 42L; // Long | Instance ID
    Long scheduleId = 1L; // Long | Instance Schedule ID
    try {
      InlineResponse20059 result = apiInstance.getInstanceSchedule(id, scheduleId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstanceSchedule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Instance ID |
 **scheduleId** | **Long**| Instance Schedule ID |

### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstanceSchedules"></a>
# **getInstanceSchedules**
> Object getInstanceSchedules(id)

Get all Instance Schedules

This endpoint retrieves all the scaling threshold schedules for a specific instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 42L; // Long | Instance ID
    try {
      Object result = apiInstance.getInstanceSchedules(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstanceSchedules");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Instance ID |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstanceThreshold"></a>
# **getInstanceThreshold**
> InlineResponse20058 getInstanceThreshold(id)

Get an Instance Scale Threshold

Retrieves the scale threshold settings for a specific instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20058 result = apiInstance.getInstanceThreshold(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstanceThreshold");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20058**](InlineResponse20058.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInstanceTypeProvisioning"></a>
# **getInstanceTypeProvisioning**
> Object getInstanceTypeProvisioning(id)

Get Specific Instance Type for Provisioning

Fetch an instance type by ID. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getInstanceTypeProvisioning(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getInstanceTypeProvisioning");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getPrepareApplyInstance"></a>
# **getPrepareApplyInstance**
> Object getPrepareApplyInstance(id)

Prepare To Apply an Instance

This endpoint provides a way to view the current instance configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getPrepareApplyInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getPrepareApplyInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getSnapshotInstance"></a>
# **getSnapshotInstance**
> Snapshot getSnapshotInstance(id)

Get a Specific Snapshot

This endpoint retrieves a specific snapshot.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Snapshot result = apiInstance.getSnapshotInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getSnapshotInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getStateInstance"></a>
# **getStateInstance**
> Object getStateInstance(id)

Get State of an Instance

This endpoint provides a way to view the state of an instance. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getStateInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getStateInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getValidateApplyInstance"></a>
# **getValidateApplyInstance**
> Object getValidateApplyInstance(id, inlineObject98)

Validate Apply State for an Instance

This endpoint provides a way to validate instance configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject98 inlineObject98 = new InlineObject98(); // InlineObject98 | 
    try {
      Object result = apiInstance.getValidateApplyInstance(id, inlineObject98);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getValidateApplyInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject98** | [**InlineObject98**](InlineObject98.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getWikiInstance"></a>
# **getWikiInstance**
> InlineResponse200168 getWikiInstance(id)

Retrieves an Instance Wiki Page

This endpoint retrieves an instance Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200168 result = apiInstance.getWikiInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#getWikiInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="importSnapshotInstance"></a>
# **importSnapshotInstance**
> Model200Success importSnapshotInstance(id, inlineObject93)

Import Snapshot of an Instance

It is possible to import a snapshot of an instance. This creates a Virtual Image of the instance as it currently exists.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject93 inlineObject93 = new InlineObject93(); // InlineObject93 | 
    try {
      Model200Success result = apiInstance.importSnapshotInstance(id, inlineObject93);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#importSnapshotInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject93** | [**InlineObject93**](InlineObject93.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="linkedCloneSnapshotInstance"></a>
# **linkedCloneSnapshotInstance**
> Model200Success linkedCloneSnapshotInstance(id, snapshotId)

Create Linked Clone of Instance Snapshot

It is possible to create a linked clone of an Instance Snapshot.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal snapshotId = new BigDecimal(78); // BigDecimal | Snapshot ID
    try {
      Model200Success result = apiInstance.linkedCloneSnapshotInstance(id, snapshotId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#linkedCloneSnapshotInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **snapshotId** | **BigDecimal**| Snapshot ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listInstanceServicePlans"></a>
# **listInstanceServicePlans**
> InlineResponse20060 listInstanceServicePlans(zoneId, layoutId, siteId)

Get Available Service Plans for an Instance

This endpoint retrieves all the Service Plans available for the specified cloud and instance layout. The response includes details about the plans and their configuration options. It may be used to get the list of available plans when creating a new instance or resizing an existing instance. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long layoutId = 98L; // Long | The Layout ID for Filtering
    Long siteId = 7L; // Long | The Site ID for Filtering
    try {
      InlineResponse20060 result = apiInstance.listInstanceServicePlans(zoneId, layoutId, siteId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#listInstanceServicePlans");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **layoutId** | **Long**| The Layout ID for Filtering | [optional]
 **siteId** | **Long**| The Site ID for Filtering | [optional]

### Return type

[**InlineResponse20060**](InlineResponse20060.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listInstanceTypesProvisioning"></a>
# **listInstanceTypesProvisioning**
> Object listInstanceTypesProvisioning(max, offset, sort, direction, phrase, name, code, featured, details)

Get All Instance Types for Provisioning

Fetch the list of available instance types. These can vary in range from database containers, to web containers, to custom containers. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    Boolean featured = false; // Boolean | Filter by featured
    Boolean details = true; // Boolean | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default.
    try {
      Object result = apiInstance.listInstanceTypesProvisioning(max, offset, sort, direction, phrase, name, code, featured, details);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#listInstanceTypesProvisioning");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listInstances"></a>
# **listInstances**
> Object listInstances(max, offset, name, phrase, instanceType, lastUpdated, createdBy, agentInstalled, status, environment, showDeleted, deleted, expireDate, expireDateMin, expireDays, expireDaysMin, shutdownDate, shutdownDateMin, shutdownDays, shutdownDaysMin, labels, allLabels, tags, metadata, details)

Get All Instances

This endpoint retrieves a paginated list of instances. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String instanceType = "ubuntu"; // String | The Instance Type Code for Filtering
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Long createdBy = 63L; // Long | The User ID for Filtering
    Boolean agentInstalled = true; // Boolean | Filter instances by if agent is installed or not on the associated servers.
    String status = "running"; // String | The instance status for filtering.
    String environment = "lab"; // String | The environment for filtering.
    Boolean showDeleted = false; // Boolean | If true, includes instances in pending removal status.
    Boolean deleted = true; // Boolean | If true, only deleted resources or instances in pending removal status are returned.
    String expireDate = "2019-01-01"; // String | Filter by expireDate less than or equal to specified date
    String expireDateMin = "2019-01-01"; // String | Filter expireDate greater than or equal to the specified date
    String expireDays = "20"; // String | Filter by expireDays less than or equal to the specified value
    String expireDaysMin = "20"; // String | Filter by expireDays greater than or equal to the specified value
    String shutdownDate = "2019-01-01"; // String | Filter by shutdownDate less than equal to the specified date
    String shutdownDateMin = "2019-01-01"; // String | Filter by shutdownDate greater than or equal to the specified date
    String shutdownDays = "20"; // String | Filter by shutdownDays less than or equal to the specified value
    String shutdownDaysMin = "20"; // String | Filter by shutdownDays greater than or equal to the specified value
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    String tags = "tags.env=qa&tags.env=test"; // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
    String metadata = "tags.env=qa&tags.env=test"; // String | Alias for tags
    Boolean details = false; // Boolean | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2.
    try {
      Object result = apiInstance.listInstances(max, offset, name, phrase, instanceType, lastUpdated, createdBy, agentInstalled, status, environment, showDeleted, deleted, expireDate, expireDateMin, expireDays, expireDaysMin, shutdownDate, shutdownDateMin, shutdownDays, shutdownDaysMin, labels, allLabels, tags, metadata, details);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#listInstances");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **instanceType** | **String**| The Instance Type Code for Filtering | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **createdBy** | **Long**| The User ID for Filtering | [optional]
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listSecurityGroupsInstance"></a>
# **listSecurityGroupsInstance**
> Object listSecurityGroupsInstance(id)

Get Security Groups for an Instance

This returns a list of all of the security groups applied to an instance and whether the firewall is enabled.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.listSecurityGroupsInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#listSecurityGroupsInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="lockInstance"></a>
# **lockInstance**
> Model200Success lockInstance(id)

Lock an Instance

This will lock the instance. While locked, instances may not be removed.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.lockInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#lockInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="refreshStateInstance"></a>
# **refreshStateInstance**
> Model200Success refreshStateInstance(id)

Refresh State of an Instance

This endpoint provides a way to refresh the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.refreshStateInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#refreshStateInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="removeInstancesFromControl"></a>
# **removeInstancesFromControl**
> SuccessMessage removeInstancesFromControl(inlineObject99)

Remove From Control

Will delete a brownfield instance (or instances) asynchronously (Only deletes records local to Morpheus, actual VMs remain unchanged).

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    InlineObject99 inlineObject99 = new InlineObject99(); // InlineObject99 | 
    try {
      SuccessMessage result = apiInstance.removeInstancesFromControl(inlineObject99);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#removeInstancesFromControl");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="resizeInstance"></a>
# **resizeInstance**
> Object resizeInstance(id, instanceResize)

Resize an Instance

It is possible to resize containers within an instance by increasing their memory plan or storage limit. This is done by assigning a new service plan to the container. This endpoint also allows for NIC reconfiguration by passing a new array of networkInterfaces

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InstanceResize instanceResize = {"instance":{"id":1,"plan":{"id":15}},"volumes":[{"id":-1,"rootVolume":true,"name":"root","size":20,"sizeId":null,"storageType":null,"datastoreId":null}],"deleteOriginalVolumes":true,"servicePlanOptions":{"maxCores":1,"coresPerSocket":1,"maxMemory":536870912},"networkInterfaces":[{"id":534,"name":"eth0","network":{"id":"network-20"}},{"name":"eth1","network":{"id":"network-20"}}]}; // InstanceResize | 
    try {
      Object result = apiInstance.resizeInstance(id, instanceResize);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#resizeInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **instanceResize** | [**InstanceResize**](InstanceResize.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="restartInstance"></a>
# **restartInstance**
> Object restartInstance(id)

Restart an instance

This will restart all containers running within an instance. This includes rebuilding the environment variables and applying settings to the docker containers.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.restartInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#restartInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="revertSnapshotInstance"></a>
# **revertSnapshotInstance**
> Model200Success revertSnapshotInstance(id, snapshotId)

Revert Instance to Snapshot

It is possible to restore an Instance to a snapshot.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal snapshotId = new BigDecimal(78); // BigDecimal | Snapshot ID
    try {
      Model200Success result = apiInstance.revertSnapshotInstance(id, snapshotId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#revertSnapshotInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **snapshotId** | **BigDecimal**| Snapshot ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="runWorkflowInstance"></a>
# **runWorkflowInstance**
> Model200Success runWorkflowInstance(id, workflowId, workflowName, instanceWorkflow)

Run Workflow on an Instance

This will run a provisioning workflow on all containers in an instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long workflowId = 3L; // Long | ID of the workflow to execute
    String workflowName = "myworkflow"; // String | Name of the workflow to execute
    InstanceWorkflow instanceWorkflow = {"taskSet":{"customOptions":{"foo":"bar"}}}; // InstanceWorkflow | 
    try {
      Model200Success result = apiInstance.runWorkflowInstance(id, workflowId, workflowName, instanceWorkflow);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#runWorkflowInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **workflowId** | **Long**| ID of the workflow to execute | [optional]
 **workflowName** | **String**| Name of the workflow to execute | [optional]
 **instanceWorkflow** | [**InstanceWorkflow**](InstanceWorkflow.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="setApplyInstance"></a>
# **setApplyInstance**
> Model200Success setApplyInstance(id, inlineObject91)

Apply State of an Instance

This endpoint provides a way to apply the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject91 inlineObject91 = new InlineObject91(); // InlineObject91 | 
    try {
      Model200Success result = apiInstance.setApplyInstance(id, inlineObject91);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#setApplyInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject91** | [**InlineObject91**](InlineObject91.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="setInstanceSecurityGroups"></a>
# **setInstanceSecurityGroups**
> Object setInstanceSecurityGroups(id, inlineObject94)

Set Security Groups for an Instance

Set Security Groups for an Instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject94 inlineObject94 = new InlineObject94(); // InlineObject94 | 
    try {
      Object result = apiInstance.setInstanceSecurityGroups(id, inlineObject94);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#setInstanceSecurityGroups");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject94** | [**InlineObject94**](InlineObject94.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="snapshotInstance"></a>
# **snapshotInstance**
> Model200Success snapshotInstance(id, instanceSnapshot)

Snapshot an Instance

This endpoint will create a snapshot of an instance. This is done asychronously, so the ID of the snapshot is not returned.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InstanceSnapshot instanceSnapshot = {"snapshot":{"name":"snapshot-test","description":"A snapshot created via the Morpheus api"}}; // InstanceSnapshot | 
    try {
      Model200Success result = apiInstance.snapshotInstance(id, instanceSnapshot);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#snapshotInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **instanceSnapshot** | [**InstanceSnapshot**](InstanceSnapshot.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="snapshotsInstance"></a>
# **snapshotsInstance**
> InstanceSnapshots snapshotsInstance(id)

Get list of snapshots for an Instance

Get list of snapshots for an Instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InstanceSnapshots result = apiInstance.snapshotsInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#snapshotsInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InstanceSnapshots**](InstanceSnapshots.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="startInstance"></a>
# **startInstance**
> Object startInstance(id)

Start an instance

This will start all containers running within an instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.startInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#startInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="stopInstance"></a>
# **stopInstance**
> Object stopInstance(id)

Stop an instance

This will stop all containers running within an instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.stopInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#stopInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="suspendInstance"></a>
# **suspendInstance**
> Model200Success suspendInstance(id)

Suspend an instance

This will suspend all containers in the instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.suspendInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#suspendInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="unlockInstance"></a>
# **unlockInstance**
> Model200Success unlockInstance(id)

Unlock an Instance

This will unlock the instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.unlockInstance(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#unlockInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateInstance"></a>
# **updateInstance**
> Object updateInstance(id, instanceUpdate)

Updating an Instance

Updating an Instance

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InstanceUpdate instanceUpdate = {"instance":{"description":"my new redis","addTags":[{"name":"hello","value":"world"},{"name":"flash","value":"bang"}],"removeTags":[{"name":"oldTag"}]},"config":{"customOptions":{"dbfoldername":"data"}}}; // InstanceUpdate | 
    try {
      Object result = apiInstance.updateInstance(id, instanceUpdate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#updateInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **instanceUpdate** | [**InstanceUpdate**](InstanceUpdate.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateInstanceNetworkInterface"></a>
# **updateInstanceNetworkInterface**
> Object updateInstanceNetworkInterface(id, networkInterfaceId, networkInterfaceUpdate)

Updating a label for an Instance&#39;s Network

Updating an Instance&#39;s Network&#39;s Label

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal networkInterfaceId = new BigDecimal(78); // BigDecimal | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
    NetworkInterfaceUpdate networkInterfaceUpdate = {"name":"newLabelName"}; // NetworkInterfaceUpdate | 
    try {
      Object result = apiInstance.updateInstanceNetworkInterface(id, networkInterfaceId, networkInterfaceUpdate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#updateInstanceNetworkInterface");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **networkInterfaceId** | **BigDecimal**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced |
 **networkInterfaceUpdate** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateInstanceSchedule"></a>
# **updateInstanceSchedule**
> Object updateInstanceSchedule(id, scheduleId, inlineObject97)

Updating an Instance Schedule

This endpoint provides updating of an instance schedule

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 9L; // Long | Instance ID
    Long scheduleId = 1L; // Long | Instance Schedule ID
    InlineObject97 inlineObject97 = new InlineObject97(); // InlineObject97 | 
    try {
      Object result = apiInstance.updateInstanceSchedule(id, scheduleId, inlineObject97);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#updateInstanceSchedule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Instance ID |
 **scheduleId** | **Long**| Instance Schedule ID |
 **inlineObject97** | [**InlineObject97**](InlineObject97.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateInstanceThreshold"></a>
# **updateInstanceThreshold**
> Object updateInstanceThreshold(id, inlineObject95)

Updates an Instance Scale Threshold

Updates the scale threshold settings for a specific instance 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject95 inlineObject95 = new InlineObject95(); // InlineObject95 | 
    try {
      Object result = apiInstance.updateInstanceThreshold(id, inlineObject95);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#updateInstanceThreshold");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject95** | [**InlineObject95**](InlineObject95.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateWikiInstance"></a>
# **updateWikiInstance**
> Object updateWikiInstance(id, inlineObject270)

Update an Instance Wiki Page

Updates an instance Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InstancesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InstancesApi apiInstance = new InstancesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject270 inlineObject270 = new InlineObject270(); // InlineObject270 | 
    try {
      Object result = apiInstance.updateWikiInstance(id, inlineObject270);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InstancesApi#updateWikiInstance");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject270** | [**InlineObject270**](InlineObject270.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

