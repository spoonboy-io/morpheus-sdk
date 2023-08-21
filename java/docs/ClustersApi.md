# ClustersApi

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


<a name="addCluster"></a>
# **addCluster**
> Object addCluster(inlineObject52)

Create a Cluster

This endpoint will create a cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    InlineObject52 inlineObject52 = new InlineObject52(); // InlineObject52 | 
    try {
      Object result = apiInstance.addCluster(inlineObject52);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#addCluster");
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
 **inlineObject52** | [**InlineObject52**](InlineObject52.md)|  | [optional]

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

<a name="addClusterNamespace"></a>
# **addClusterNamespace**
> Object addClusterNamespace(clusterId, inlineObject56)

Add Namespace (Kubernetes)

Add Namespace (Kubernetes)

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    InlineObject56 inlineObject56 = new InlineObject56(); // InlineObject56 | 
    try {
      Object result = apiInstance.addClusterNamespace(clusterId, inlineObject56);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#addClusterNamespace");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **inlineObject56** | [**InlineObject56**](InlineObject56.md)|  | [optional]

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

<a name="addClusterWorker"></a>
# **addClusterWorker**
> Object addClusterWorker(clusterId, inlineObject59)

Add Worker

Add Worker

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    InlineObject59 inlineObject59 = new InlineObject59(); // InlineObject59 | 
    try {
      Object result = apiInstance.addClusterWorker(clusterId, inlineObject59);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#addClusterWorker");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **inlineObject59** | [**InlineObject59**](InlineObject59.md)|  | [optional]

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

<a name="applyTemplate"></a>
# **applyTemplate**
> ClusterApplyTemplate applyTemplate(clusterId, inlineObject54)

Apply Template to Cluster (Kubernetes)

This endpoint applies the requested template, via Service Url, YAML, or Spec Template name/id, to a Kubernetes cluster.  **Note**: The success response informs of status of submission of request. Results of the actual template application can be assesed with the returned execution id via [/api/execution-request/{uniqueId}](/reference/getexecutionrequest) 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    InlineObject54 inlineObject54 = new InlineObject54(); // InlineObject54 | 
    try {
      ClusterApplyTemplate result = apiInstance.applyTemplate(clusterId, inlineObject54);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#applyTemplate");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **inlineObject54** | [**InlineObject54**](InlineObject54.md)|  | [optional]

### Return type

[**ClusterApplyTemplate**](ClusterApplyTemplate.md)

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

<a name="deleteCluster"></a>
# **deleteCluster**
> Model200Success deleteCluster(clusterId, removeInstances, removeResources, preserveVolumes, releaseFloatingIps, releaseEIPs, force)

Delete a Cluster

Will delete a cluster and associated resources, hosts, volumes asynchronously

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    String removeInstances = "\"off\""; // String | Remove Instances
    String removeResources = "\"on\""; // String | Remove Resources
    String preserveVolumes = "\"off\""; // String | Preserve Volumes
    String releaseFloatingIps = "\"on\""; // String | Release Floating IPs
    String releaseEIPs = "\"on\""; // String | Alias for releaseFloatingIps
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteCluster(clusterId, removeInstances, removeResources, preserveVolumes, releaseFloatingIps, releaseEIPs, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteCluster");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **removeInstances** | **String**| Remove Instances | [optional] [default to &quot;off&quot;]
 **removeResources** | **String**| Remove Resources | [optional] [default to &quot;on&quot;]
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &quot;off&quot;]
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

<a name="deleteClusterContainer"></a>
# **deleteClusterContainer**
> Model200Success deleteClusterContainer(clusterId, id, force)

Delete Container

This endpoint deletes a specified container from a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterContainer(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterContainer");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="deleteClusterDeployment"></a>
# **deleteClusterDeployment**
> Model200Success deleteClusterDeployment(clusterId, id, force)

Delete Deployment

This endpoint deletes a specified deployment from a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterDeployment(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterDeployment");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="deleteClusterJob"></a>
# **deleteClusterJob**
> Model200Success deleteClusterJob(clusterId, id, force)

Delete a Job

This endpoint deletes a specified job from a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterJob(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterJob");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="deleteClusterNamespace"></a>
# **deleteClusterNamespace**
> Model200Success deleteClusterNamespace(clusterId, id, force)

Delete a Namespace (Kubernetes)

Will delete a namespace from the specified cluster

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterNamespace(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterNamespace");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="deleteClusterService"></a>
# **deleteClusterService**
> Model200Success deleteClusterService(clusterId, id, force)

Delete a Service

This endpoint deletes a specified service from a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterService(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterService");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="deleteClusterStatefulSet"></a>
# **deleteClusterStatefulSet**
> SuccessError deleteClusterStatefulSet(clusterId, id, force)

Delete a Stateful Set

Will delete a stateful set from the specified cluster

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      SuccessError result = apiInstance.deleteClusterStatefulSet(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterStatefulSet");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **force** | **String**| Force Delete | [optional] [default to &quot;off&quot;]

### Return type

[**SuccessError**](SuccessError.md)

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

<a name="deleteClusterVolume"></a>
# **deleteClusterVolume**
> Model200Success deleteClusterVolume(clusterId, id, force)

Delete a Volume

Will delete a volume from the specified cluster

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterVolume(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterVolume");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="deleteClusterWorker"></a>
# **deleteClusterWorker**
> Model200Success deleteClusterWorker(clusterId, id, force)

Delete a Worker

This endpoint deletes a specified worker from a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteClusterWorker(clusterId, id, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#deleteClusterWorker");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
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

<a name="getCluster"></a>
# **getCluster**
> InlineResponse20026 getCluster(clusterId)

Get a Specific Cluster

This endpoint retrieves a specific cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      InlineResponse20026 result = apiInstance.getCluster(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getCluster");
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
 **clusterId** | **Integer**| The ID of the cluster |

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

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

<a name="getClusterApiConfig"></a>
# **getClusterApiConfig**
> ClusterApiConfig getClusterApiConfig(clusterId)

Get API Config

This endpoint retrieves the API configuration for a specified cluster. The configuration is cluster type specific.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      ClusterApiConfig result = apiInstance.getClusterApiConfig(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterApiConfig");
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
 **clusterId** | **Integer**| The ID of the cluster |

### Return type

[**ClusterApiConfig**](ClusterApiConfig.md)

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

<a name="getClusterDatastore"></a>
# **getClusterDatastore**
> InlineResponse20027 getClusterDatastore(clusterId, id)

Get a Specific Datastore

This endpoint retrieves a specific cluster datastore.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20027 result = apiInstance.getClusterDatastore(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterDatastore");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

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

<a name="getClusterHistory"></a>
# **getClusterHistory**
> Object getClusterHistory(clusterId)

Get Cluster History

This endpoint retrieves the process history for a specific cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      Object result = apiInstance.getClusterHistory(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterHistory");
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
 **clusterId** | **Integer**| The ID of the cluster |

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

<a name="getClusterHistoryDetail"></a>
# **getClusterHistoryDetail**
> InlineResponse20028 getClusterHistoryDetail(clusterId, id)

Get Cluster History Details

This endpoint retrieves the history for a specific cluster process.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20028 result = apiInstance.getClusterHistoryDetail(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterHistoryDetail");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

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

<a name="getClusterHistoryEventDetail"></a>
# **getClusterHistoryEventDetail**
> InlineResponse20029 getClusterHistoryEventDetail(clusterId, id)

Get Cluster History Event

This endpoint retrieves the process event for a specific cluster process event.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20029 result = apiInstance.getClusterHistoryEventDetail(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterHistoryEventDetail");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

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

<a name="getClusterMasters"></a>
# **getClusterMasters**
> InlineResponse20030 getClusterMasters(clusterId, phrase)

Get Masters (Kubernetes)

This endpoint retrieves masters of a specified kubernetes cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      InlineResponse20030 result = apiInstance.getClusterMasters(clusterId, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterMasters");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

[**InlineResponse20030**](InlineResponse20030.md)

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

<a name="getClusterNamespace"></a>
# **getClusterNamespace**
> InlineResponse20031 getClusterNamespace(clusterId, id)

Get Namespace (Kubernetes)

This endpoint retrieves a specific namespace of a Kubernetes cluster

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20031 result = apiInstance.getClusterNamespace(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterNamespace");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20031**](InlineResponse20031.md)

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

<a name="getClusterNamespaces"></a>
# **getClusterNamespaces**
> Object getClusterNamespaces(clusterId)

List Namespaces (Kubernetes)

List Namespaces (Kubernetes)

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      Object result = apiInstance.getClusterNamespaces(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterNamespaces");
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
 **clusterId** | **Integer**| The ID of the cluster |

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

<a name="getClusterUpgradeVersions"></a>
# **getClusterUpgradeVersions**
> InlineResponse20032 getClusterUpgradeVersions(clusterId)

Get Cluster Upgrade Versions (Kubernetes)

This endpoint returns valid version targets for upgrading kubectl and kubeadm on the cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      InlineResponse20032 result = apiInstance.getClusterUpgradeVersions(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getClusterUpgradeVersions");
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
 **clusterId** | **Integer**| The ID of the cluster |

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

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

<a name="getWikiCluster"></a>
# **getWikiCluster**
> InlineResponse200168 getWikiCluster(clusterId)

Retrieves a Cluster Wiki Page

This endpoint retrieves a cluster Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      InlineResponse200168 result = apiInstance.getWikiCluster(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#getWikiCluster");
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
 **clusterId** | **Integer**| The ID of the cluster |

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

<a name="listClusterContainers"></a>
# **listClusterContainers**
> Object listClusterContainers(clusterId, max, offset, sort, order, phrase, resourceLevel)

Get Containers

This endpoint retrieves containers of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String resourceLevel = "resourceLevel_example"; // String | Resource level filter
    try {
      Object result = apiInstance.listClusterContainers(clusterId, max, offset, sort, order, phrase, resourceLevel);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterContainers");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **resourceLevel** | **String**| Resource level filter | [optional] [enum: app, system, storage, logging]

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

<a name="listClusterDatastores"></a>
# **listClusterDatastores**
> Object listClusterDatastores(clusterId, max, offset, sort, order, phrase, name, code, hideInactive)

Get Datastores

This endpoint retrieves datastores of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    Boolean hideInactive = false; // Boolean | If true restricts query to only load active datastores
    try {
      Object result = apiInstance.listClusterDatastores(clusterId, max, offset, sort, order, phrase, name, code, hideInactive);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterDatastores");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listClusterDeployments"></a>
# **listClusterDeployments**
> Object listClusterDeployments(clusterId, max, offset, sort, order, phrase, resourceLevel)

Get Deployments

This endpoint retrieves deployments of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String resourceLevel = "resourceLevel_example"; // String | Resource level filter
    try {
      Object result = apiInstance.listClusterDeployments(clusterId, max, offset, sort, order, phrase, resourceLevel);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterDeployments");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **resourceLevel** | **String**| Resource level filter | [optional] [enum: app, system, storage, logging]

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

<a name="listClusterJobs"></a>
# **listClusterJobs**
> Object listClusterJobs(clusterId, max, offset, sort, order, phrase)

Get Jobs

This endpoint retrieves jobs of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listClusterJobs(clusterId, max, offset, sort, order, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterJobs");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listClusterPods"></a>
# **listClusterPods**
> Object listClusterPods(clusterId, max, offset, sort, order, phrase, resourceLevel)

Get Pods

This endpoint retrieves pods of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String resourceLevel = "resourceLevel_example"; // String | Resource level filter
    try {
      Object result = apiInstance.listClusterPods(clusterId, max, offset, sort, order, phrase, resourceLevel);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterPods");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **resourceLevel** | **String**| Resource level filter | [optional] [enum: app, system, storage, logging]

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

<a name="listClusterServices"></a>
# **listClusterServices**
> Object listClusterServices(clusterId, max, offset, sort, order, phrase)

Get Services

This endpoint retrieves services of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listClusterServices(clusterId, max, offset, sort, order, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterServices");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listClusterStatefulSets"></a>
# **listClusterStatefulSets**
> Object listClusterStatefulSets(clusterId, max, offset, sort, order, phrase, resourceLevel)

Get Stateful Sets

This endpoint retrieves stateful sets of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String resourceLevel = "resourceLevel_example"; // String | Resource level filter
    try {
      Object result = apiInstance.listClusterStatefulSets(clusterId, max, offset, sort, order, phrase, resourceLevel);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterStatefulSets");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **resourceLevel** | **String**| Resource level filter | [optional] [enum: app, system, storage, logging]

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

<a name="listClusterTypes"></a>
# **listClusterTypes**
> Object listClusterTypes(max, offset, sort, direction, phrase, name, code, providerType)

Get All Cluster Types

Fetch a list of available cluster types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    String providerType = "providerType_example"; // String | Filter by `Provider Type` code. 
    try {
      Object result = apiInstance.listClusterTypes(max, offset, sort, direction, phrase, name, code, providerType);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterTypes");
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
 **providerType** | **String**| Filter by &#x60;Provider Type&#x60; code.  | [optional] [enum: morpheus, kubernetes, kvm, docker]

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

<a name="listClusterVolumes"></a>
# **listClusterVolumes**
> Object listClusterVolumes(clusterId)

Get Volumes

This endpoint retrieves volumes of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      Object result = apiInstance.listClusterVolumes(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterVolumes");
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
 **clusterId** | **Integer**| The ID of the cluster |

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

<a name="listClusterWorkers"></a>
# **listClusterWorkers**
> Object listClusterWorkers(clusterId)

Get Workers

This endpoint retrieves workers of a specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    try {
      Object result = apiInstance.listClusterWorkers(clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusterWorkers");
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
 **clusterId** | **Integer**| The ID of the cluster |

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

<a name="listClusters"></a>
# **listClusters**
> Object listClusters(max, offset, sort, direction, phrase, name, zoneId, typeId, labels, allLabels)

Get All Clusters

This endpoint retrieves all clusters and a list of clusters associated with the zone by id.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long typeId = 3L; // Long | Type filter, restricts query to only load clusters of a specified cluster type
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      Object result = apiInstance.listClusters(max, offset, sort, direction, phrase, name, zoneId, typeId, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#listClusters");
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
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **typeId** | **Long**| Type filter, restricts query to only load clusters of a specified cluster type | [optional]
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional]

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

<a name="restartClusterContainer"></a>
# **restartClusterContainer**
> SuccessError restartClusterContainer(clusterId, id)

Restart a Container

Will restart a container in the specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessError result = apiInstance.restartClusterContainer(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#restartClusterContainer");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessError**](SuccessError.md)

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

<a name="restartClusterDeployment"></a>
# **restartClusterDeployment**
> SuccessError restartClusterDeployment(clusterId, id)

Restart a Deployment

Will restart a deployment in the specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessError result = apiInstance.restartClusterDeployment(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#restartClusterDeployment");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessError**](SuccessError.md)

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

<a name="restartClusterPod"></a>
# **restartClusterPod**
> SuccessError restartClusterPod(clusterId, id)

Restart a Pod

Will restart a pod in the specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessError result = apiInstance.restartClusterPod(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#restartClusterPod");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessError**](SuccessError.md)

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

<a name="restartClusterStatefulSet"></a>
# **restartClusterStatefulSet**
> SuccessError restartClusterStatefulSet(clusterId, id)

Restart a Stateful Set

Will restart a stateful set in the specified cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessError result = apiInstance.restartClusterStatefulSet(clusterId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#restartClusterStatefulSet");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessError**](SuccessError.md)

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

<a name="updateCluster"></a>
# **updateCluster**
> Object updateCluster(clusterId, inlineObject53)

Update Cluster

Update Cluster

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    InlineObject53 inlineObject53 = new InlineObject53(); // InlineObject53 | 
    try {
      Object result = apiInstance.updateCluster(clusterId, inlineObject53);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateCluster");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **inlineObject53** | [**InlineObject53**](InlineObject53.md)|  | [optional]

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

<a name="updateClusterDatastore"></a>
# **updateClusterDatastore**
> Object updateClusterDatastore(clusterId, id, inlineObject55)

Update Datastore

Update Datastore

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject55 inlineObject55 = new InlineObject55(); // InlineObject55 | 
    try {
      Object result = apiInstance.updateClusterDatastore(clusterId, id, inlineObject55);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateClusterDatastore");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject55** | [**InlineObject55**](InlineObject55.md)|  | [optional]

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

<a name="updateClusterNamespace"></a>
# **updateClusterNamespace**
> Object updateClusterNamespace(clusterId, id, inlineObject57)

Update Namespace (Kubernetes)

Update Namespace (Kubernetes)

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject57 inlineObject57 = new InlineObject57(); // InlineObject57 | 
    try {
      Object result = apiInstance.updateClusterNamespace(clusterId, id, inlineObject57);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateClusterNamespace");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject57** | [**InlineObject57**](InlineObject57.md)|  | [optional]

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

<a name="updateClusterPermissions"></a>
# **updateClusterPermissions**
> Object updateClusterPermissions(clusterId, inlineObject58)

Update Cluster Permissions

Update Cluster Permissions

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    InlineObject58 inlineObject58 = new InlineObject58(); // InlineObject58 | 
    try {
      Object result = apiInstance.updateClusterPermissions(clusterId, inlineObject58);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateClusterPermissions");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **inlineObject58** | [**InlineObject58**](InlineObject58.md)|  | [optional]

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

<a name="updateClusterUpgradeVersions"></a>
# **updateClusterUpgradeVersions**
> Model200Success updateClusterUpgradeVersions(clusterId, targetVersion)

Upgrade a Cluster (Kubernetes)

This endpoint updates the kubectl and kudeadm versions on a Kubernetes cluster to the specified version. Use Get Cluster Upgrade Versions to list valid version targets for the cluster.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    String targetVersion = "1.21.14"; // String | Target version for cluster after upgrade
    try {
      Model200Success result = apiInstance.updateClusterUpgradeVersions(clusterId, targetVersion);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateClusterUpgradeVersions");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **targetVersion** | **String**| Target version for cluster after upgrade |

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

<a name="updateClusterWorkerCount"></a>
# **updateClusterWorkerCount**
> Model200Success updateClusterWorkerCount(clusterId, workerCount)

Update Worker Count

This endpoint resizes a cluster to the specified number of worker nodes (only supports Azure AKS, Google GKE, and Amazon EKS clusters).

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    Long workerCount = 5L; // Long | The target number of worker nodes
    try {
      Model200Success result = apiInstance.updateClusterWorkerCount(clusterId, workerCount);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateClusterWorkerCount");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **workerCount** | **Long**| The target number of worker nodes |

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

<a name="updateWikiCluster"></a>
# **updateWikiCluster**
> Object updateWikiCluster(clusterId, inlineObject268)

Update a Cluster Wiki Page

Updates a cluster Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClustersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClustersApi apiInstance = new ClustersApi(defaultClient);
    Integer clusterId = 5; // Integer | The ID of the cluster
    InlineObject268 inlineObject268 = new InlineObject268(); // InlineObject268 | 
    try {
      Object result = apiInstance.updateWikiCluster(clusterId, inlineObject268);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClustersApi#updateWikiCluster");
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
 **clusterId** | **Integer**| The ID of the cluster |
 **inlineObject268** | [**InlineObject268**](InlineObject268.md)|  | [optional]

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

