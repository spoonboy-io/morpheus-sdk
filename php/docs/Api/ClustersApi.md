# OpenAPI\Client\ClustersApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCluster()**](ClustersApi.md#addCluster) | **POST** /api/clusters | Create a Cluster
[**addClusterNamespace()**](ClustersApi.md#addClusterNamespace) | **POST** /api/clusters/{clusterId}/namespaces | Add Namespace (Kubernetes)
[**addClusterWorker()**](ClustersApi.md#addClusterWorker) | **POST** /api/clusters/{clusterId}/servers | Add Worker
[**applyTemplate()**](ClustersApi.md#applyTemplate) | **POST** /api/clusters/{clusterId}/apply-template | Apply Template to Cluster (Kubernetes)
[**deleteCluster()**](ClustersApi.md#deleteCluster) | **DELETE** /api/clusters/{clusterId} | Delete a Cluster
[**deleteClusterContainer()**](ClustersApi.md#deleteClusterContainer) | **DELETE** /api/clusters/{clusterId}/containers/{id} | Delete Container
[**deleteClusterDeployment()**](ClustersApi.md#deleteClusterDeployment) | **DELETE** /api/clusters/{clusterId}/deployments/{id} | Delete Deployment
[**deleteClusterJob()**](ClustersApi.md#deleteClusterJob) | **DELETE** /api/clusters/{clusterId}/jobs/{id} | Delete a Job
[**deleteClusterNamespace()**](ClustersApi.md#deleteClusterNamespace) | **DELETE** /api/clusters/{clusterId}/namespaces/{id} | Delete a Namespace (Kubernetes)
[**deleteClusterService()**](ClustersApi.md#deleteClusterService) | **DELETE** /api/clusters/{clusterId}/services/{id} | Delete a Service
[**deleteClusterStatefulSet()**](ClustersApi.md#deleteClusterStatefulSet) | **DELETE** /api/clusters/{clusterId}/statefulsets/{id} | Delete a Stateful Set
[**deleteClusterVolume()**](ClustersApi.md#deleteClusterVolume) | **DELETE** /api/clusters/{clusterId}/volumes/{id} | Delete a Volume
[**deleteClusterWorker()**](ClustersApi.md#deleteClusterWorker) | **DELETE** /api/clusters/{clusterId}/servers/{id} | Delete a Worker
[**getCluster()**](ClustersApi.md#getCluster) | **GET** /api/clusters/{clusterId} | Get a Specific Cluster
[**getClusterApiConfig()**](ClustersApi.md#getClusterApiConfig) | **GET** /api/clusters/{clusterId}/api-config | Get API Config
[**getClusterDatastore()**](ClustersApi.md#getClusterDatastore) | **GET** /api/clusters/{clusterId}/datastores/{id} | Get a Specific Datastore
[**getClusterHistory()**](ClustersApi.md#getClusterHistory) | **GET** /api/clusters/{clusterId}/history | Get Cluster History
[**getClusterHistoryDetail()**](ClustersApi.md#getClusterHistoryDetail) | **GET** /api/clusters/{clusterId}/history/{id} | Get Cluster History Details
[**getClusterHistoryEventDetail()**](ClustersApi.md#getClusterHistoryEventDetail) | **GET** /api/clusters/{clusterId}/history/events/{id} | Get Cluster History Event
[**getClusterMasters()**](ClustersApi.md#getClusterMasters) | **GET** /api/clusters/{clusterId}/masters | Get Masters (Kubernetes)
[**getClusterNamespace()**](ClustersApi.md#getClusterNamespace) | **GET** /api/clusters/{clusterId}/namespaces/{id} | Get Namespace (Kubernetes)
[**getClusterNamespaces()**](ClustersApi.md#getClusterNamespaces) | **GET** /api/clusters/{clusterId}/namespaces | List Namespaces (Kubernetes)
[**getClusterUpgradeVersions()**](ClustersApi.md#getClusterUpgradeVersions) | **GET** /api/clusters/{clusterId}/upgrade-cluster | Get Cluster Upgrade Versions (Kubernetes)
[**getWikiCluster()**](ClustersApi.md#getWikiCluster) | **GET** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**listClusterContainers()**](ClustersApi.md#listClusterContainers) | **GET** /api/clusters/{clusterId}/containers | Get Containers
[**listClusterDatastores()**](ClustersApi.md#listClusterDatastores) | **GET** /api/clusters/{clusterId}/datastores | Get Datastores
[**listClusterDeployments()**](ClustersApi.md#listClusterDeployments) | **GET** /api/clusters/{clusterId}/deployments | Get Deployments
[**listClusterJobs()**](ClustersApi.md#listClusterJobs) | **GET** /api/clusters/{clusterId}/jobs | Get Jobs
[**listClusterPods()**](ClustersApi.md#listClusterPods) | **GET** /api/clusters/{clusterId}/pods | Get Pods
[**listClusterServices()**](ClustersApi.md#listClusterServices) | **GET** /api/clusters/{clusterId}/services | Get Services
[**listClusterStatefulSets()**](ClustersApi.md#listClusterStatefulSets) | **GET** /api/clusters/{clusterId}/statefulsets | Get Stateful Sets
[**listClusterTypes()**](ClustersApi.md#listClusterTypes) | **GET** /api/cluster-types | Get All Cluster Types
[**listClusterVolumes()**](ClustersApi.md#listClusterVolumes) | **GET** /api/clusters/{clusterId}/volumes | Get Volumes
[**listClusterWorkers()**](ClustersApi.md#listClusterWorkers) | **GET** /api/clusters/{clusterId}/workers | Get Workers
[**listClusters()**](ClustersApi.md#listClusters) | **GET** /api/clusters | Get All Clusters
[**restartClusterContainer()**](ClustersApi.md#restartClusterContainer) | **PUT** /api/clusters/{clusterId}/containers/{id}/restart | Restart a Container
[**restartClusterDeployment()**](ClustersApi.md#restartClusterDeployment) | **PUT** /api/clusters/{clusterId}/deployments/{id}/restart | Restart a Deployment
[**restartClusterPod()**](ClustersApi.md#restartClusterPod) | **PUT** /api/clusters/{clusterId}/pods/{id}/restart | Restart a Pod
[**restartClusterStatefulSet()**](ClustersApi.md#restartClusterStatefulSet) | **PUT** /api/clusters/{clusterId}/statefulsets/{id}/restart | Restart a Stateful Set
[**updateCluster()**](ClustersApi.md#updateCluster) | **PUT** /api/clusters/{clusterId} | Update Cluster
[**updateClusterDatastore()**](ClustersApi.md#updateClusterDatastore) | **PUT** /api/clusters/{clusterId}/datastores/{id} | Update Datastore
[**updateClusterNamespace()**](ClustersApi.md#updateClusterNamespace) | **PUT** /api/clusters/{clusterId}/namespaces/{id} | Update Namespace (Kubernetes)
[**updateClusterPermissions()**](ClustersApi.md#updateClusterPermissions) | **PUT** /api/clusters/{clusterId}/permissions | Update Cluster Permissions
[**updateClusterUpgradeVersions()**](ClustersApi.md#updateClusterUpgradeVersions) | **POST** /api/clusters/{clusterId}/upgrade-cluster | Upgrade a Cluster (Kubernetes)
[**updateClusterWorkerCount()**](ClustersApi.md#updateClusterWorkerCount) | **PUT** /api/clusters/{clusterId}/worker-count | Update Worker Count
[**updateWikiCluster()**](ClustersApi.md#updateWikiCluster) | **PUT** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page


## `addCluster()`

```php
addCluster($inline_object52): object
```

Create a Cluster

This endpoint will create a cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object52 = new \OpenAPI\Client\Model\InlineObject52(); // \OpenAPI\Client\Model\InlineObject52

try {
    $result = $apiInstance->addCluster($inline_object52);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->addCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object52** | [**\OpenAPI\Client\Model\InlineObject52**](../Model/InlineObject52.md)|  | [optional]

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

## `addClusterNamespace()`

```php
addClusterNamespace($cluster_id, $inline_object56): object
```

Add Namespace (Kubernetes)

Add Namespace (Kubernetes)

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object56 = new \OpenAPI\Client\Model\InlineObject56(); // \OpenAPI\Client\Model\InlineObject56

try {
    $result = $apiInstance->addClusterNamespace($cluster_id, $inline_object56);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->addClusterNamespace: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object56** | [**\OpenAPI\Client\Model\InlineObject56**](../Model/InlineObject56.md)|  | [optional]

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

## `addClusterWorker()`

```php
addClusterWorker($cluster_id, $inline_object59): object
```

Add Worker

Add Worker

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object59 = new \OpenAPI\Client\Model\InlineObject59(); // \OpenAPI\Client\Model\InlineObject59

try {
    $result = $apiInstance->addClusterWorker($cluster_id, $inline_object59);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->addClusterWorker: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object59** | [**\OpenAPI\Client\Model\InlineObject59**](../Model/InlineObject59.md)|  | [optional]

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

## `applyTemplate()`

```php
applyTemplate($cluster_id, $inline_object54): \OpenAPI\Client\Model\ClusterApplyTemplate
```

Apply Template to Cluster (Kubernetes)

This endpoint applies the requested template, via Service Url, YAML, or Spec Template name/id, to a Kubernetes cluster.  **Note**: The success response informs of status of submission of request. Results of the actual template application can be assesed with the returned execution id via [/api/execution-request/{uniqueId}](/reference/getexecutionrequest)

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object54 = new \OpenAPI\Client\Model\InlineObject54(); // \OpenAPI\Client\Model\InlineObject54

try {
    $result = $apiInstance->applyTemplate($cluster_id, $inline_object54);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->applyTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object54** | [**\OpenAPI\Client\Model\InlineObject54**](../Model/InlineObject54.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\ClusterApplyTemplate**](../Model/ClusterApplyTemplate.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteCluster()`

```php
deleteCluster($cluster_id, $remove_instances, $remove_resources, $preserve_volumes, $release_floating_ips, $release_ei_ps, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Cluster

Will delete a cluster and associated resources, hosts, volumes asynchronously

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$remove_instances = on; // string | Remove Instances
$remove_resources = off; // string | Remove Resources
$preserve_volumes = on; // string | Preserve Volumes
$release_floating_ips = off; // string | Release Floating IPs
$release_ei_ps = off; // string | Alias for releaseFloatingIps
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteCluster($cluster_id, $remove_instances, $remove_resources, $preserve_volumes, $release_floating_ips, $release_ei_ps, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **remove_instances** | **string**| Remove Instances | [optional] [default to &#39;off&#39;]
 **remove_resources** | **string**| Remove Resources | [optional] [default to &#39;on&#39;]
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

## `deleteClusterContainer()`

```php
deleteClusterContainer($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete Container

This endpoint deletes a specified container from a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterContainer($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterContainer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `deleteClusterDeployment()`

```php
deleteClusterDeployment($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete Deployment

This endpoint deletes a specified deployment from a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterDeployment($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterDeployment: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `deleteClusterJob()`

```php
deleteClusterJob($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Job

This endpoint deletes a specified job from a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterJob($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterJob: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `deleteClusterNamespace()`

```php
deleteClusterNamespace($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Namespace (Kubernetes)

Will delete a namespace from the specified cluster

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterNamespace($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterNamespace: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `deleteClusterService()`

```php
deleteClusterService($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Service

This endpoint deletes a specified service from a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterService($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterService: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `deleteClusterStatefulSet()`

```php
deleteClusterStatefulSet($cluster_id, $id, $force): \OpenAPI\Client\Model\SuccessError
```

Delete a Stateful Set

Will delete a stateful set from the specified cluster

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterStatefulSet($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterStatefulSet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **string**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**\OpenAPI\Client\Model\SuccessError**](../Model/SuccessError.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteClusterVolume()`

```php
deleteClusterVolume($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Volume

Will delete a volume from the specified cluster

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterVolume($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterVolume: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `deleteClusterWorker()`

```php
deleteClusterWorker($cluster_id, $id, $force): \OpenAPI\Client\Model\Model200Success
```

Delete a Worker

This endpoint deletes a specified worker from a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$force = on; // string | Force Delete

try {
    $result = $apiInstance->deleteClusterWorker($cluster_id, $id, $force);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->deleteClusterWorker: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `getCluster()`

```php
getCluster($cluster_id): \OpenAPI\Client\Model\InlineResponse20026
```

Get a Specific Cluster

This endpoint retrieves a specific cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getCluster($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20026**](../Model/InlineResponse20026.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterApiConfig()`

```php
getClusterApiConfig($cluster_id): \OpenAPI\Client\Model\ClusterApiConfig
```

Get API Config

This endpoint retrieves the API configuration for a specified cluster. The configuration is cluster type specific.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getClusterApiConfig($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterApiConfig: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**\OpenAPI\Client\Model\ClusterApiConfig**](../Model/ClusterApiConfig.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterDatastore()`

```php
getClusterDatastore($cluster_id, $id): \OpenAPI\Client\Model\InlineResponse20027
```

Get a Specific Datastore

This endpoint retrieves a specific cluster datastore.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getClusterDatastore($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterDatastore: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20027**](../Model/InlineResponse20027.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterHistory()`

```php
getClusterHistory($cluster_id): object
```

Get Cluster History

This endpoint retrieves the process history for a specific cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getClusterHistory($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterHistory: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

## `getClusterHistoryDetail()`

```php
getClusterHistoryDetail($cluster_id, $id): \OpenAPI\Client\Model\InlineResponse20028
```

Get Cluster History Details

This endpoint retrieves the history for a specific cluster process.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getClusterHistoryDetail($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterHistoryDetail: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20028**](../Model/InlineResponse20028.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterHistoryEventDetail()`

```php
getClusterHistoryEventDetail($cluster_id, $id): \OpenAPI\Client\Model\InlineResponse20029
```

Get Cluster History Event

This endpoint retrieves the process event for a specific cluster process event.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getClusterHistoryEventDetail($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterHistoryEventDetail: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20029**](../Model/InlineResponse20029.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterMasters()`

```php
getClusterMasters($cluster_id, $phrase): \OpenAPI\Client\Model\InlineResponse20030
```

Get Masters (Kubernetes)

This endpoint retrieves masters of a specified kubernetes cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getClusterMasters($cluster_id, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterMasters: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20030**](../Model/InlineResponse20030.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterNamespace()`

```php
getClusterNamespace($cluster_id, $id): \OpenAPI\Client\Model\InlineResponse20031
```

Get Namespace (Kubernetes)

This endpoint retrieves a specific namespace of a Kubernetes cluster

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getClusterNamespace($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterNamespace: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20031**](../Model/InlineResponse20031.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getClusterNamespaces()`

```php
getClusterNamespaces($cluster_id): object
```

List Namespaces (Kubernetes)

List Namespaces (Kubernetes)

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getClusterNamespaces($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterNamespaces: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

## `getClusterUpgradeVersions()`

```php
getClusterUpgradeVersions($cluster_id): \OpenAPI\Client\Model\InlineResponse20032
```

Get Cluster Upgrade Versions (Kubernetes)

This endpoint returns valid version targets for upgrading kubectl and kubeadm on the cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getClusterUpgradeVersions($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getClusterUpgradeVersions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20032**](../Model/InlineResponse20032.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getWikiCluster()`

```php
getWikiCluster($cluster_id): \OpenAPI\Client\Model\InlineResponse200168
```

Retrieves a Cluster Wiki Page

This endpoint retrieves a cluster Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->getWikiCluster($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->getWikiCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

## `listClusterContainers()`

```php
listClusterContainers($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level): object
```

Get Containers

This endpoint retrieves containers of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$resource_level = 'resource_level_example'; // string | Resource level filter

try {
    $result = $apiInstance->listClusterContainers($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterContainers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **string**| Resource level filter | [optional]

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

## `listClusterDatastores()`

```php
listClusterDatastores($cluster_id, $max, $offset, $sort, $order, $phrase, $name, $code, $hide_inactive): object
```

Get Datastores

This endpoint retrieves datastores of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code
$hide_inactive = true; // bool | If true restricts query to only load active datastores

try {
    $result = $apiInstance->listClusterDatastores($cluster_id, $max, $offset, $sort, $order, $phrase, $name, $code, $hide_inactive);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterDatastores: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]
 **hide_inactive** | **bool**| If true restricts query to only load active datastores | [optional] [default to false]

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

## `listClusterDeployments()`

```php
listClusterDeployments($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level): object
```

Get Deployments

This endpoint retrieves deployments of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$resource_level = 'resource_level_example'; // string | Resource level filter

try {
    $result = $apiInstance->listClusterDeployments($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterDeployments: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **string**| Resource level filter | [optional]

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

## `listClusterJobs()`

```php
listClusterJobs($cluster_id, $max, $offset, $sort, $order, $phrase): object
```

Get Jobs

This endpoint retrieves jobs of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listClusterJobs($cluster_id, $max, $offset, $sort, $order, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterJobs: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

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

## `listClusterPods()`

```php
listClusterPods($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level): object
```

Get Pods

This endpoint retrieves pods of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$resource_level = 'resource_level_example'; // string | Resource level filter

try {
    $result = $apiInstance->listClusterPods($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterPods: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **string**| Resource level filter | [optional]

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

## `listClusterServices()`

```php
listClusterServices($cluster_id, $max, $offset, $sort, $order, $phrase): object
```

Get Services

This endpoint retrieves services of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listClusterServices($cluster_id, $max, $offset, $sort, $order, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterServices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

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

## `listClusterStatefulSets()`

```php
listClusterStatefulSets($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level): object
```

Get Stateful Sets

This endpoint retrieves stateful sets of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$order = 'asc'; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$resource_level = 'resource_level_example'; // string | Resource level filter

try {
    $result = $apiInstance->listClusterStatefulSets($cluster_id, $max, $offset, $sort, $order, $phrase, $resource_level);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterStatefulSets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **order** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **string**| Resource level filter | [optional]

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

## `listClusterTypes()`

```php
listClusterTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $provider_type): object
```

Get All Cluster Types

Fetch a list of available cluster types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
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
$provider_type = 'provider_type_example'; // string | Filter by `Provider Type` code.

try {
    $result = $apiInstance->listClusterTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $provider_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterTypes: ', $e->getMessage(), PHP_EOL;
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
 **provider_type** | **string**| Filter by &#x60;Provider Type&#x60; code. | [optional]

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

## `listClusterVolumes()`

```php
listClusterVolumes($cluster_id): object
```

Get Volumes

This endpoint retrieves volumes of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->listClusterVolumes($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterVolumes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

## `listClusterWorkers()`

```php
listClusterWorkers($cluster_id): object
```

Get Workers

This endpoint retrieves workers of a specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster

try {
    $result = $apiInstance->listClusterWorkers($cluster_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusterWorkers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

## `listClusters()`

```php
listClusters($max, $offset, $sort, $direction, $phrase, $name, $zone_id, $type_id, $labels, $all_labels): object
```

Get All Clusters

This endpoint retrieves all clusters and a list of clusters associated with the zone by id.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
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
$zone_id = 3; // int | The Zone ID for Filtering
$type_id = 3; // int | Type filter, restricts query to only load clusters of a specified cluster type
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listClusters($max, $offset, $sort, $direction, $phrase, $name, $zone_id, $type_id, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->listClusters: ', $e->getMessage(), PHP_EOL;
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
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **type_id** | **int**| Type filter, restricts query to only load clusters of a specified cluster type | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]

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

## `restartClusterContainer()`

```php
restartClusterContainer($cluster_id, $id): \OpenAPI\Client\Model\SuccessError
```

Restart a Container

Will restart a container in the specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartClusterContainer($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->restartClusterContainer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\SuccessError**](../Model/SuccessError.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `restartClusterDeployment()`

```php
restartClusterDeployment($cluster_id, $id): \OpenAPI\Client\Model\SuccessError
```

Restart a Deployment

Will restart a deployment in the specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartClusterDeployment($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->restartClusterDeployment: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\SuccessError**](../Model/SuccessError.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `restartClusterPod()`

```php
restartClusterPod($cluster_id, $id): \OpenAPI\Client\Model\SuccessError
```

Restart a Pod

Will restart a pod in the specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartClusterPod($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->restartClusterPod: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\SuccessError**](../Model/SuccessError.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `restartClusterStatefulSet()`

```php
restartClusterStatefulSet($cluster_id, $id): \OpenAPI\Client\Model\SuccessError
```

Restart a Stateful Set

Will restart a stateful set in the specified cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartClusterStatefulSet($cluster_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->restartClusterStatefulSet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\SuccessError**](../Model/SuccessError.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateCluster()`

```php
updateCluster($cluster_id, $inline_object53): object
```

Update Cluster

Update Cluster

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object53 = new \OpenAPI\Client\Model\InlineObject53(); // \OpenAPI\Client\Model\InlineObject53

try {
    $result = $apiInstance->updateCluster($cluster_id, $inline_object53);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object53** | [**\OpenAPI\Client\Model\InlineObject53**](../Model/InlineObject53.md)|  | [optional]

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

## `updateClusterDatastore()`

```php
updateClusterDatastore($cluster_id, $id, $inline_object55): object
```

Update Datastore

Update Datastore

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object55 = new \OpenAPI\Client\Model\InlineObject55(); // \OpenAPI\Client\Model\InlineObject55

try {
    $result = $apiInstance->updateClusterDatastore($cluster_id, $id, $inline_object55);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateClusterDatastore: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object55** | [**\OpenAPI\Client\Model\InlineObject55**](../Model/InlineObject55.md)|  | [optional]

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

## `updateClusterNamespace()`

```php
updateClusterNamespace($cluster_id, $id, $inline_object57): object
```

Update Namespace (Kubernetes)

Update Namespace (Kubernetes)

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object57 = new \OpenAPI\Client\Model\InlineObject57(); // \OpenAPI\Client\Model\InlineObject57

try {
    $result = $apiInstance->updateClusterNamespace($cluster_id, $id, $inline_object57);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateClusterNamespace: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object57** | [**\OpenAPI\Client\Model\InlineObject57**](../Model/InlineObject57.md)|  | [optional]

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

## `updateClusterPermissions()`

```php
updateClusterPermissions($cluster_id, $inline_object58): object
```

Update Cluster Permissions

Update Cluster Permissions

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object58 = new \OpenAPI\Client\Model\InlineObject58(); // \OpenAPI\Client\Model\InlineObject58

try {
    $result = $apiInstance->updateClusterPermissions($cluster_id, $inline_object58);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateClusterPermissions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object58** | [**\OpenAPI\Client\Model\InlineObject58**](../Model/InlineObject58.md)|  | [optional]

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

## `updateClusterUpgradeVersions()`

```php
updateClusterUpgradeVersions($cluster_id, $target_version): \OpenAPI\Client\Model\Model200Success
```

Upgrade a Cluster (Kubernetes)

This endpoint updates the kubectl and kudeadm versions on a Kubernetes cluster to the specified version. Use Get Cluster Upgrade Versions to list valid version targets for the cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$target_version = 1.21.14; // string | Target version for cluster after upgrade

try {
    $result = $apiInstance->updateClusterUpgradeVersions($cluster_id, $target_version);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateClusterUpgradeVersions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **target_version** | **string**| Target version for cluster after upgrade |

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

## `updateClusterWorkerCount()`

```php
updateClusterWorkerCount($cluster_id, $worker_count): \OpenAPI\Client\Model\Model200Success
```

Update Worker Count

This endpoint resizes a cluster to the specified number of worker nodes (only supports Azure AKS, Google GKE, and Amazon EKS clusters).

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$worker_count = 5; // int | The target number of worker nodes

try {
    $result = $apiInstance->updateClusterWorkerCount($cluster_id, $worker_count);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateClusterWorkerCount: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **worker_count** | **int**| The target number of worker nodes |

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

## `updateWikiCluster()`

```php
updateWikiCluster($cluster_id, $inline_object268): object
```

Update a Cluster Wiki Page

Updates a cluster Wiki page.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ClustersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cluster_id = 5; // int | The ID of the cluster
$inline_object268 = new \OpenAPI\Client\Model\InlineObject268(); // \OpenAPI\Client\Model\InlineObject268

try {
    $result = $apiInstance->updateWikiCluster($cluster_id, $inline_object268);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ClustersApi->updateWikiCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **inline_object268** | [**\OpenAPI\Client\Model\InlineObject268**](../Model/InlineObject268.md)|  | [optional]

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
