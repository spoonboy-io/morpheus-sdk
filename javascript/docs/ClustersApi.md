# MorpheusApi.ClustersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCluster**](ClustersApi.md#addCluster) | **POST** /api/clusters | Create a Cluster
[**addClusterNamespace**](ClustersApi.md#addClusterNamespace) | **POST** /api/clusters/{clusterId}/namespaces | Add Namespace (Kubernetes)
[**addClusterWorker**](ClustersApi.md#addClusterWorker) | **POST** /api/clusters/{clusterId}/servers | Add Worker
[**applyTemplate**](ClustersApi.md#applyTemplate) | **POST** /api/clusters/{clusterId}/apply-template | Apply Template to Cluster (Kubernetes)
[**deleteCluster**](ClustersApi.md#deleteCluster) | **DELETE** /api/clusters/{clusterId} | Delete a Cluster
[**deleteClusterContainer**](ClustersApi.md#deleteClusterContainer) | **DELETE** /api/clusters/{clusterId}/containers/{id} | Delete Container
[**deleteClusterDeployment**](ClustersApi.md#deleteClusterDeployment) | **DELETE** /api/clusters/{clusterId}/deployments/{id} | Delete Deployment
[**deleteClusterJob**](ClustersApi.md#deleteClusterJob) | **DELETE** /api/clusters/{clusterId}/jobs/{id} | Delete a Job
[**deleteClusterNamespace**](ClustersApi.md#deleteClusterNamespace) | **DELETE** /api/clusters/{clusterId}/namespaces/{id} | Delete a Namespace (Kubernetes)
[**deleteClusterService**](ClustersApi.md#deleteClusterService) | **DELETE** /api/clusters/{clusterId}/services/{id} | Delete a Service
[**deleteClusterStatefulSet**](ClustersApi.md#deleteClusterStatefulSet) | **DELETE** /api/clusters/{clusterId}/statefulsets/{id} | Delete a Stateful Set
[**deleteClusterVolume**](ClustersApi.md#deleteClusterVolume) | **DELETE** /api/clusters/{clusterId}/volumes/{id} | Delete a Volume
[**deleteClusterWorker**](ClustersApi.md#deleteClusterWorker) | **DELETE** /api/clusters/{clusterId}/servers/{id} | Delete a Worker
[**getCluster**](ClustersApi.md#getCluster) | **GET** /api/clusters/{clusterId} | Get a Specific Cluster
[**getClusterApiConfig**](ClustersApi.md#getClusterApiConfig) | **GET** /api/clusters/{clusterId}/api-config | Get API Config
[**getClusterDatastore**](ClustersApi.md#getClusterDatastore) | **GET** /api/clusters/{clusterId}/datastores/{id} | Get a Specific Datastore
[**getClusterHistory**](ClustersApi.md#getClusterHistory) | **GET** /api/clusters/{clusterId}/history | Get Cluster History
[**getClusterHistoryDetail**](ClustersApi.md#getClusterHistoryDetail) | **GET** /api/clusters/{clusterId}/history/{id} | Get Cluster History Details
[**getClusterHistoryEventDetail**](ClustersApi.md#getClusterHistoryEventDetail) | **GET** /api/clusters/{clusterId}/history/events/{id} | Get Cluster History Event
[**getClusterMasters**](ClustersApi.md#getClusterMasters) | **GET** /api/clusters/{clusterId}/masters | Get Masters (Kubernetes)
[**getClusterNamespace**](ClustersApi.md#getClusterNamespace) | **GET** /api/clusters/{clusterId}/namespaces/{id} | Get Namespace (Kubernetes)
[**getClusterNamespaces**](ClustersApi.md#getClusterNamespaces) | **GET** /api/clusters/{clusterId}/namespaces | List Namespaces (Kubernetes)
[**getClusterUpgradeVersions**](ClustersApi.md#getClusterUpgradeVersions) | **GET** /api/clusters/{clusterId}/upgrade-cluster | Get Cluster Upgrade Versions (Kubernetes)
[**getWikiCluster**](ClustersApi.md#getWikiCluster) | **GET** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**listClusterContainers**](ClustersApi.md#listClusterContainers) | **GET** /api/clusters/{clusterId}/containers | Get Containers
[**listClusterDatastores**](ClustersApi.md#listClusterDatastores) | **GET** /api/clusters/{clusterId}/datastores | Get Datastores
[**listClusterDeployments**](ClustersApi.md#listClusterDeployments) | **GET** /api/clusters/{clusterId}/deployments | Get Deployments
[**listClusterJobs**](ClustersApi.md#listClusterJobs) | **GET** /api/clusters/{clusterId}/jobs | Get Jobs
[**listClusterPods**](ClustersApi.md#listClusterPods) | **GET** /api/clusters/{clusterId}/pods | Get Pods
[**listClusterServices**](ClustersApi.md#listClusterServices) | **GET** /api/clusters/{clusterId}/services | Get Services
[**listClusterStatefulSets**](ClustersApi.md#listClusterStatefulSets) | **GET** /api/clusters/{clusterId}/statefulsets | Get Stateful Sets
[**listClusterTypes**](ClustersApi.md#listClusterTypes) | **GET** /api/cluster-types | Get All Cluster Types
[**listClusterVolumes**](ClustersApi.md#listClusterVolumes) | **GET** /api/clusters/{clusterId}/volumes | Get Volumes
[**listClusterWorkers**](ClustersApi.md#listClusterWorkers) | **GET** /api/clusters/{clusterId}/workers | Get Workers
[**listClusters**](ClustersApi.md#listClusters) | **GET** /api/clusters | Get All Clusters
[**restartClusterContainer**](ClustersApi.md#restartClusterContainer) | **PUT** /api/clusters/{clusterId}/containers/{id}/restart | Restart a Container
[**restartClusterDeployment**](ClustersApi.md#restartClusterDeployment) | **PUT** /api/clusters/{clusterId}/deployments/{id}/restart | Restart a Deployment
[**restartClusterPod**](ClustersApi.md#restartClusterPod) | **PUT** /api/clusters/{clusterId}/pods/{id}/restart | Restart a Pod
[**restartClusterStatefulSet**](ClustersApi.md#restartClusterStatefulSet) | **PUT** /api/clusters/{clusterId}/statefulsets/{id}/restart | Restart a Stateful Set
[**updateCluster**](ClustersApi.md#updateCluster) | **PUT** /api/clusters/{clusterId} | Update Cluster
[**updateClusterDatastore**](ClustersApi.md#updateClusterDatastore) | **PUT** /api/clusters/{clusterId}/datastores/{id} | Update Datastore
[**updateClusterNamespace**](ClustersApi.md#updateClusterNamespace) | **PUT** /api/clusters/{clusterId}/namespaces/{id} | Update Namespace (Kubernetes)
[**updateClusterPermissions**](ClustersApi.md#updateClusterPermissions) | **PUT** /api/clusters/{clusterId}/permissions | Update Cluster Permissions
[**updateClusterUpgradeVersions**](ClustersApi.md#updateClusterUpgradeVersions) | **POST** /api/clusters/{clusterId}/upgrade-cluster | Upgrade a Cluster (Kubernetes)
[**updateClusterWorkerCount**](ClustersApi.md#updateClusterWorkerCount) | **PUT** /api/clusters/{clusterId}/worker-count | Update Worker Count
[**updateWikiCluster**](ClustersApi.md#updateWikiCluster) | **PUT** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page



## addCluster

> Object addCluster(opts)

Create a Cluster

This endpoint will create a cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let opts = {
  'inlineObject52': new MorpheusApi.InlineObject52() // InlineObject52 | 
};
apiInstance.addCluster(opts, (error, data, response) => {
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
 **inlineObject52** | [**InlineObject52**](InlineObject52.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addClusterNamespace

> Object addClusterNamespace(clusterId, opts)

Add Namespace (Kubernetes)

Add Namespace (Kubernetes)

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'inlineObject56': new MorpheusApi.InlineObject56() // InlineObject56 | 
};
apiInstance.addClusterNamespace(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **inlineObject56** | [**InlineObject56**](InlineObject56.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addClusterWorker

> Object addClusterWorker(clusterId, opts)

Add Worker

Add Worker

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'inlineObject59': new MorpheusApi.InlineObject59() // InlineObject59 | 
};
apiInstance.addClusterWorker(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **inlineObject59** | [**InlineObject59**](InlineObject59.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## applyTemplate

> ClusterApplyTemplate applyTemplate(clusterId, opts)

Apply Template to Cluster (Kubernetes)

This endpoint applies the requested template, via Service Url, YAML, or Spec Template name/id, to a Kubernetes cluster.  **Note**: The success response informs of status of submission of request. Results of the actual template application can be assesed with the returned execution id via [/api/execution-request/{uniqueId}](/reference/getexecutionrequest) 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'inlineObject54': new MorpheusApi.InlineObject54() // InlineObject54 | 
};
apiInstance.applyTemplate(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **inlineObject54** | [**InlineObject54**](InlineObject54.md)|  | [optional] 

### Return type

[**ClusterApplyTemplate**](ClusterApplyTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteCluster

> Model200Success deleteCluster(clusterId, opts)

Delete a Cluster

Will delete a cluster and associated resources, hosts, volumes asynchronously

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'removeInstances': on, // String | Remove Instances
  'removeResources': off, // String | Remove Resources
  'preserveVolumes': on, // String | Preserve Volumes
  'releaseFloatingIps': off, // String | Release Floating IPs
  'releaseEIPs': off, // String | Alias for releaseFloatingIps
  'force': on // String | Force Delete
};
apiInstance.deleteCluster(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **removeInstances** | **String**| Remove Instances | [optional] [default to &#39;off&#39;]
 **removeResources** | **String**| Remove Resources | [optional] [default to &#39;on&#39;]
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &#39;off&#39;]
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


## deleteClusterContainer

> Model200Success deleteClusterContainer(clusterId, id, opts)

Delete Container

This endpoint deletes a specified container from a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterContainer(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterDeployment

> Model200Success deleteClusterDeployment(clusterId, id, opts)

Delete Deployment

This endpoint deletes a specified deployment from a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterDeployment(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterJob

> Model200Success deleteClusterJob(clusterId, id, opts)

Delete a Job

This endpoint deletes a specified job from a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterJob(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterNamespace

> Model200Success deleteClusterNamespace(clusterId, id, opts)

Delete a Namespace (Kubernetes)

Will delete a namespace from the specified cluster

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterNamespace(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterService

> Model200Success deleteClusterService(clusterId, id, opts)

Delete a Service

This endpoint deletes a specified service from a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterService(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterStatefulSet

> SuccessError deleteClusterStatefulSet(clusterId, id, opts)

Delete a Stateful Set

Will delete a stateful set from the specified cluster

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterStatefulSet(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**SuccessError**](SuccessError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterVolume

> Model200Success deleteClusterVolume(clusterId, id, opts)

Delete a Volume

Will delete a volume from the specified cluster

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterVolume(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteClusterWorker

> Model200Success deleteClusterWorker(clusterId, id, opts)

Delete a Worker

This endpoint deletes a specified worker from a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteClusterWorker(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCluster

> InlineResponse20026 getCluster(clusterId)

Get a Specific Cluster

This endpoint retrieves a specific cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.getCluster(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterApiConfig

> ClusterApiConfig getClusterApiConfig(clusterId)

Get API Config

This endpoint retrieves the API configuration for a specified cluster. The configuration is cluster type specific.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.getClusterApiConfig(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

[**ClusterApiConfig**](ClusterApiConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterDatastore

> InlineResponse20027 getClusterDatastore(clusterId, id)

Get a Specific Datastore

This endpoint retrieves a specific cluster datastore.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClusterDatastore(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterHistory

> Object getClusterHistory(clusterId)

Get Cluster History

This endpoint retrieves the process history for a specific cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.getClusterHistory(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterHistoryDetail

> InlineResponse20028 getClusterHistoryDetail(clusterId, id)

Get Cluster History Details

This endpoint retrieves the history for a specific cluster process.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClusterHistoryDetail(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterHistoryEventDetail

> InlineResponse20029 getClusterHistoryEventDetail(clusterId, id)

Get Cluster History Event

This endpoint retrieves the process event for a specific cluster process event.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClusterHistoryEventDetail(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterMasters

> InlineResponse20030 getClusterMasters(clusterId, opts)

Get Masters (Kubernetes)

This endpoint retrieves masters of a specified kubernetes cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getClusterMasters(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterNamespace

> InlineResponse20031 getClusterNamespace(clusterId, id)

Get Namespace (Kubernetes)

This endpoint retrieves a specific namespace of a Kubernetes cluster

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClusterNamespace(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20031**](InlineResponse20031.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterNamespaces

> Object getClusterNamespaces(clusterId)

List Namespaces (Kubernetes)

List Namespaces (Kubernetes)

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.getClusterNamespaces(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClusterUpgradeVersions

> InlineResponse20032 getClusterUpgradeVersions(clusterId)

Get Cluster Upgrade Versions (Kubernetes)

This endpoint returns valid version targets for upgrading kubectl and kubeadm on the cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.getClusterUpgradeVersions(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getWikiCluster

> InlineResponse200168 getWikiCluster(clusterId)

Retrieves a Cluster Wiki Page

This endpoint retrieves a cluster Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.getWikiCluster(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterContainers

> Object listClusterContainers(clusterId, opts)

Get Containers

This endpoint retrieves containers of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'resourceLevel': "resourceLevel_example" // String | Resource level filter
};
apiInstance.listClusterContainers(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **resourceLevel** | **String**| Resource level filter | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterDatastores

> Object listClusterDatastores(clusterId, opts)

Get Datastores

This endpoint retrieves datastores of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'hideInactive': true // Boolean | If true restricts query to only load active datastores
};
apiInstance.listClusterDatastores(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **hideInactive** | **Boolean**| If true restricts query to only load active datastores | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterDeployments

> Object listClusterDeployments(clusterId, opts)

Get Deployments

This endpoint retrieves deployments of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'resourceLevel': "resourceLevel_example" // String | Resource level filter
};
apiInstance.listClusterDeployments(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **resourceLevel** | **String**| Resource level filter | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterJobs

> Object listClusterJobs(clusterId, opts)

Get Jobs

This endpoint retrieves jobs of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listClusterJobs(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterPods

> Object listClusterPods(clusterId, opts)

Get Pods

This endpoint retrieves pods of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'resourceLevel': "resourceLevel_example" // String | Resource level filter
};
apiInstance.listClusterPods(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **resourceLevel** | **String**| Resource level filter | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterServices

> Object listClusterServices(clusterId, opts)

Get Services

This endpoint retrieves services of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listClusterServices(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterStatefulSets

> Object listClusterStatefulSets(clusterId, opts)

Get Stateful Sets

This endpoint retrieves stateful sets of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'resourceLevel': "resourceLevel_example" // String | Resource level filter
};
apiInstance.listClusterStatefulSets(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **resourceLevel** | **String**| Resource level filter | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterTypes

> Object listClusterTypes(opts)

Get All Cluster Types

Fetch a list of available cluster types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'providerType': "providerType_example" // String | Filter by `Provider Type` code. 
};
apiInstance.listClusterTypes(opts, (error, data, response) => {
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
 **providerType** | **String**| Filter by &#x60;Provider Type&#x60; code.  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterVolumes

> Object listClusterVolumes(clusterId)

Get Volumes

This endpoint retrieves volumes of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.listClusterVolumes(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterWorkers

> Object listClusterWorkers(clusterId)

Get Workers

This endpoint retrieves workers of a specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
apiInstance.listClusterWorkers(clusterId, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusters

> Object listClusters(opts)

Get All Clusters

This endpoint retrieves all clusters and a list of clusters associated with the zone by id.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'zoneId': 3, // Number | The Zone ID for Filtering
  'typeId': 3, // Number | Type filter, restricts query to only load clusters of a specified cluster type
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listClusters(opts, (error, data, response) => {
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
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **typeId** | **Number**| Type filter, restricts query to only load clusters of a specified cluster type | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## restartClusterContainer

> SuccessError restartClusterContainer(clusterId, id)

Restart a Container

Will restart a container in the specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartClusterContainer(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**SuccessError**](SuccessError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## restartClusterDeployment

> SuccessError restartClusterDeployment(clusterId, id)

Restart a Deployment

Will restart a deployment in the specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartClusterDeployment(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**SuccessError**](SuccessError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## restartClusterPod

> SuccessError restartClusterPod(clusterId, id)

Restart a Pod

Will restart a pod in the specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartClusterPod(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**SuccessError**](SuccessError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## restartClusterStatefulSet

> SuccessError restartClusterStatefulSet(clusterId, id)

Restart a Stateful Set

Will restart a stateful set in the specified cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartClusterStatefulSet(clusterId, id, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**SuccessError**](SuccessError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateCluster

> Object updateCluster(clusterId, opts)

Update Cluster

Update Cluster

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'inlineObject53': new MorpheusApi.InlineObject53() // InlineObject53 | 
};
apiInstance.updateCluster(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **inlineObject53** | [**InlineObject53**](InlineObject53.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateClusterDatastore

> Object updateClusterDatastore(clusterId, id, opts)

Update Datastore

Update Datastore

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject55': new MorpheusApi.InlineObject55() // InlineObject55 | 
};
apiInstance.updateClusterDatastore(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject55** | [**InlineObject55**](InlineObject55.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateClusterNamespace

> Object updateClusterNamespace(clusterId, id, opts)

Update Namespace (Kubernetes)

Update Namespace (Kubernetes)

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject57': new MorpheusApi.InlineObject57() // InlineObject57 | 
};
apiInstance.updateClusterNamespace(clusterId, id, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject57** | [**InlineObject57**](InlineObject57.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateClusterPermissions

> Object updateClusterPermissions(clusterId, opts)

Update Cluster Permissions

Update Cluster Permissions

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'inlineObject58': new MorpheusApi.InlineObject58() // InlineObject58 | 
};
apiInstance.updateClusterPermissions(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **inlineObject58** | [**InlineObject58**](InlineObject58.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateClusterUpgradeVersions

> Model200Success updateClusterUpgradeVersions(clusterId, targetVersion)

Upgrade a Cluster (Kubernetes)

This endpoint updates the kubectl and kudeadm versions on a Kubernetes cluster to the specified version. Use Get Cluster Upgrade Versions to list valid version targets for the cluster.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let targetVersion = 1.21.14; // String | Target version for cluster after upgrade
apiInstance.updateClusterUpgradeVersions(clusterId, targetVersion, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **targetVersion** | **String**| Target version for cluster after upgrade | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateClusterWorkerCount

> Model200Success updateClusterWorkerCount(clusterId, workerCount)

Update Worker Count

This endpoint resizes a cluster to the specified number of worker nodes (only supports Azure AKS, Google GKE, and Amazon EKS clusters).

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let workerCount = 5; // Number | The target number of worker nodes
apiInstance.updateClusterWorkerCount(clusterId, workerCount, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **workerCount** | **Number**| The target number of worker nodes | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateWikiCluster

> Object updateWikiCluster(clusterId, opts)

Update a Cluster Wiki Page

Updates a cluster Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClustersApi();
let clusterId = 5; // Number | The ID of the cluster
let opts = {
  'inlineObject268': new MorpheusApi.InlineObject268() // InlineObject268 | 
};
apiInstance.updateWikiCluster(clusterId, opts, (error, data, response) => {
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
 **clusterId** | **Number**| The ID of the cluster | 
 **inlineObject268** | [**InlineObject268**](InlineObject268.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

