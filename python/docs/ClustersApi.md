# openapi_client.ClustersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_cluster**](ClustersApi.md#add_cluster) | **POST** /api/clusters | Create a Cluster
[**add_cluster_namespace**](ClustersApi.md#add_cluster_namespace) | **POST** /api/clusters/{clusterId}/namespaces | Add Namespace (Kubernetes)
[**add_cluster_worker**](ClustersApi.md#add_cluster_worker) | **POST** /api/clusters/{clusterId}/servers | Add Worker
[**apply_template**](ClustersApi.md#apply_template) | **POST** /api/clusters/{clusterId}/apply-template | Apply Template to Cluster (Kubernetes)
[**delete_cluster**](ClustersApi.md#delete_cluster) | **DELETE** /api/clusters/{clusterId} | Delete a Cluster
[**delete_cluster_container**](ClustersApi.md#delete_cluster_container) | **DELETE** /api/clusters/{clusterId}/containers/{id} | Delete Container
[**delete_cluster_deployment**](ClustersApi.md#delete_cluster_deployment) | **DELETE** /api/clusters/{clusterId}/deployments/{id} | Delete Deployment
[**delete_cluster_job**](ClustersApi.md#delete_cluster_job) | **DELETE** /api/clusters/{clusterId}/jobs/{id} | Delete a Job
[**delete_cluster_namespace**](ClustersApi.md#delete_cluster_namespace) | **DELETE** /api/clusters/{clusterId}/namespaces/{id} | Delete a Namespace (Kubernetes)
[**delete_cluster_service**](ClustersApi.md#delete_cluster_service) | **DELETE** /api/clusters/{clusterId}/services/{id} | Delete a Service
[**delete_cluster_stateful_set**](ClustersApi.md#delete_cluster_stateful_set) | **DELETE** /api/clusters/{clusterId}/statefulsets/{id} | Delete a Stateful Set
[**delete_cluster_volume**](ClustersApi.md#delete_cluster_volume) | **DELETE** /api/clusters/{clusterId}/volumes/{id} | Delete a Volume
[**delete_cluster_worker**](ClustersApi.md#delete_cluster_worker) | **DELETE** /api/clusters/{clusterId}/servers/{id} | Delete a Worker
[**get_cluster**](ClustersApi.md#get_cluster) | **GET** /api/clusters/{clusterId} | Get a Specific Cluster
[**get_cluster_api_config**](ClustersApi.md#get_cluster_api_config) | **GET** /api/clusters/{clusterId}/api-config | Get API Config
[**get_cluster_datastore**](ClustersApi.md#get_cluster_datastore) | **GET** /api/clusters/{clusterId}/datastores/{id} | Get a Specific Datastore
[**get_cluster_history**](ClustersApi.md#get_cluster_history) | **GET** /api/clusters/{clusterId}/history | Get Cluster History
[**get_cluster_history_detail**](ClustersApi.md#get_cluster_history_detail) | **GET** /api/clusters/{clusterId}/history/{id} | Get Cluster History Details
[**get_cluster_history_event_detail**](ClustersApi.md#get_cluster_history_event_detail) | **GET** /api/clusters/{clusterId}/history/events/{id} | Get Cluster History Event
[**get_cluster_masters**](ClustersApi.md#get_cluster_masters) | **GET** /api/clusters/{clusterId}/masters | Get Masters (Kubernetes)
[**get_cluster_namespace**](ClustersApi.md#get_cluster_namespace) | **GET** /api/clusters/{clusterId}/namespaces/{id} | Get Namespace (Kubernetes)
[**get_cluster_namespaces**](ClustersApi.md#get_cluster_namespaces) | **GET** /api/clusters/{clusterId}/namespaces | List Namespaces (Kubernetes)
[**get_cluster_upgrade_versions**](ClustersApi.md#get_cluster_upgrade_versions) | **GET** /api/clusters/{clusterId}/upgrade-cluster | Get Cluster Upgrade Versions (Kubernetes)
[**get_wiki_cluster**](ClustersApi.md#get_wiki_cluster) | **GET** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**list_cluster_containers**](ClustersApi.md#list_cluster_containers) | **GET** /api/clusters/{clusterId}/containers | Get Containers
[**list_cluster_datastores**](ClustersApi.md#list_cluster_datastores) | **GET** /api/clusters/{clusterId}/datastores | Get Datastores
[**list_cluster_deployments**](ClustersApi.md#list_cluster_deployments) | **GET** /api/clusters/{clusterId}/deployments | Get Deployments
[**list_cluster_jobs**](ClustersApi.md#list_cluster_jobs) | **GET** /api/clusters/{clusterId}/jobs | Get Jobs
[**list_cluster_pods**](ClustersApi.md#list_cluster_pods) | **GET** /api/clusters/{clusterId}/pods | Get Pods
[**list_cluster_services**](ClustersApi.md#list_cluster_services) | **GET** /api/clusters/{clusterId}/services | Get Services
[**list_cluster_stateful_sets**](ClustersApi.md#list_cluster_stateful_sets) | **GET** /api/clusters/{clusterId}/statefulsets | Get Stateful Sets
[**list_cluster_types**](ClustersApi.md#list_cluster_types) | **GET** /api/cluster-types | Get All Cluster Types
[**list_cluster_volumes**](ClustersApi.md#list_cluster_volumes) | **GET** /api/clusters/{clusterId}/volumes | Get Volumes
[**list_cluster_workers**](ClustersApi.md#list_cluster_workers) | **GET** /api/clusters/{clusterId}/workers | Get Workers
[**list_clusters**](ClustersApi.md#list_clusters) | **GET** /api/clusters | Get All Clusters
[**restart_cluster_container**](ClustersApi.md#restart_cluster_container) | **PUT** /api/clusters/{clusterId}/containers/{id}/restart | Restart a Container
[**restart_cluster_deployment**](ClustersApi.md#restart_cluster_deployment) | **PUT** /api/clusters/{clusterId}/deployments/{id}/restart | Restart a Deployment
[**restart_cluster_pod**](ClustersApi.md#restart_cluster_pod) | **PUT** /api/clusters/{clusterId}/pods/{id}/restart | Restart a Pod
[**restart_cluster_stateful_set**](ClustersApi.md#restart_cluster_stateful_set) | **PUT** /api/clusters/{clusterId}/statefulsets/{id}/restart | Restart a Stateful Set
[**update_cluster**](ClustersApi.md#update_cluster) | **PUT** /api/clusters/{clusterId} | Update Cluster
[**update_cluster_datastore**](ClustersApi.md#update_cluster_datastore) | **PUT** /api/clusters/{clusterId}/datastores/{id} | Update Datastore
[**update_cluster_namespace**](ClustersApi.md#update_cluster_namespace) | **PUT** /api/clusters/{clusterId}/namespaces/{id} | Update Namespace (Kubernetes)
[**update_cluster_permissions**](ClustersApi.md#update_cluster_permissions) | **PUT** /api/clusters/{clusterId}/permissions | Update Cluster Permissions
[**update_cluster_upgrade_versions**](ClustersApi.md#update_cluster_upgrade_versions) | **POST** /api/clusters/{clusterId}/upgrade-cluster | Upgrade a Cluster (Kubernetes)
[**update_cluster_worker_count**](ClustersApi.md#update_cluster_worker_count) | **PUT** /api/clusters/{clusterId}/worker-count | Update Worker Count
[**update_wiki_cluster**](ClustersApi.md#update_wiki_cluster) | **PUT** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page


# **add_cluster**
> AddCluster200Response add_cluster()

Create a Cluster

This endpoint will create a cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.add_cluster_request import AddClusterRequest
from openapi_client.model.add_cluster200_response import AddCluster200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    add_cluster_request = AddClusterRequest(
        cluster=ClusterCreate(
            type=ClusterCreateType(None),
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            group=ClusterCreateGroup(
                id=1,
            ),
            cloud=ClusterCreateCloud(
                id=1,
            ),
            layout=ClusterCreateLayout(
                id=1,
            ),
            server=ClusterServerCreate(
                config={},
                server_type=ClusterServerCreateServerType(
                    id=1,
                ),
                name="name_example",
                plan=ClusterServerCreatePlan(
                    id=1,
                    code="code_example",
                    options={},
                ),
                volumes=[
                    ClusterServerCreateVolumesInner(
                        id=-1,
                        root_volume=True,
                        name="root",
                        size=1,
                        size_id="size_id_example",
                        storage_type=1,
                        datastore_id="datastore_id_example",
                    ),
                ],
                network_interfaces=[
                    ClusterServerCreateNetworkInterfacesInner(
                        network=ClusterServerCreateNetworkInterfacesInnerNetwork(
                            id=ClusterServerCreateNetworkInterfacesInnerNetworkId(None),
                        ),
                        network_interface_type_id=1,
                        ip_address="ip_address_example",
                    ),
                ],
                security_groups=[
                    "security_groups_example",
                ],
                visibility="private",
                user_group=ClusterServerCreateUserGroup(
                    id=1,
                ),
                network_domain="network_domain_example",
                hostname="hostname_example",
                node_count=1,
                tags=[
                    UpdateHostManagedRequestServerTagsInner(
                        name="name_example",
                        value="value_example",
                    ),
                ],
                labels=[
                    "labels_example",
                ],
            ),
        ),
    ) # AddClusterRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Cluster
        api_response = api_instance.add_cluster(add_cluster_request=add_cluster_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->add_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_cluster_request** | [**AddClusterRequest**](AddClusterRequest.md)|  | [optional]

### Return type

[**AddCluster200Response**](AddCluster200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_cluster_namespace**
> AddClusterNamespace200Response add_cluster_namespace(cluster_id)

Add Namespace (Kubernetes)

Add Namespace (Kubernetes)

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.add_cluster_namespace_request import AddClusterNamespaceRequest
from openapi_client.model.add_cluster_namespace200_response import AddClusterNamespace200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    add_cluster_namespace_request = AddClusterNamespaceRequest(
        namespace=ClusterNamespaceCreate(
            name="name_example",
            description="description_example",
            active=False,
            resource_permissions=ClusterNamespaceCreateResourcePermissions(
                all=True,
                sites=[
                    ClusterUpdatePermissionsResourcePermissionsSitesInner(
                        id=1,
                        default=True,
                    ),
                ],
                all_plans=True,
                plans=[
                    ClusterUpdatePermissionsResourcePermissionsSitesInner(
                        id=1,
                        default=True,
                    ),
                ],
            ),
        ),
    ) # AddClusterNamespaceRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Add Namespace (Kubernetes)
        api_response = api_instance.add_cluster_namespace(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->add_cluster_namespace: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Add Namespace (Kubernetes)
        api_response = api_instance.add_cluster_namespace(cluster_id, add_cluster_namespace_request=add_cluster_namespace_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->add_cluster_namespace: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **add_cluster_namespace_request** | [**AddClusterNamespaceRequest**](AddClusterNamespaceRequest.md)|  | [optional]

### Return type

[**AddClusterNamespace200Response**](AddClusterNamespace200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_cluster_worker**
> AddClusterWorker200Response add_cluster_worker(cluster_id)

Add Worker

Add Worker

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.add_cluster_worker_request import AddClusterWorkerRequest
from openapi_client.model.add_cluster_worker200_response import AddClusterWorker200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    add_cluster_worker_request = AddClusterWorkerRequest(
        server=ClusterServerCreate(
            config={},
            server_type=ClusterServerCreateServerType(
                id=1,
            ),
            name="name_example",
            plan=ClusterServerCreatePlan(
                id=1,
                code="code_example",
                options={},
            ),
            volumes=[
                ClusterServerCreateVolumesInner(
                    id=-1,
                    root_volume=True,
                    name="root",
                    size=1,
                    size_id="size_id_example",
                    storage_type=1,
                    datastore_id="datastore_id_example",
                ),
            ],
            network_interfaces=[
                ClusterServerCreateNetworkInterfacesInner(
                    network=ClusterServerCreateNetworkInterfacesInnerNetwork(
                        id=ClusterServerCreateNetworkInterfacesInnerNetworkId(None),
                    ),
                    network_interface_type_id=1,
                    ip_address="ip_address_example",
                ),
            ],
            security_groups=[
                "security_groups_example",
            ],
            visibility="private",
            user_group=ClusterServerCreateUserGroup(
                id=1,
            ),
            network_domain="network_domain_example",
            hostname="hostname_example",
            node_count=1,
            tags=[
                UpdateHostManagedRequestServerTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            labels=[
                "labels_example",
            ],
        ),
    ) # AddClusterWorkerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Add Worker
        api_response = api_instance.add_cluster_worker(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->add_cluster_worker: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Add Worker
        api_response = api_instance.add_cluster_worker(cluster_id, add_cluster_worker_request=add_cluster_worker_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->add_cluster_worker: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **add_cluster_worker_request** | [**AddClusterWorkerRequest**](AddClusterWorkerRequest.md)|  | [optional]

### Return type

[**AddClusterWorker200Response**](AddClusterWorker200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apply_template**
> ClusterApplyTemplate apply_template(cluster_id)

Apply Template to Cluster (Kubernetes)

This endpoint applies the requested template, via Service Url, YAML, or Spec Template name/id, to a Kubernetes cluster.  **Note**: The success response informs of status of submission of request. Results of the actual template application can be assesed with the returned execution id via [/api/execution-request/{uniqueId}](/reference/getexecutionrequest) 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.apply_template_request import ApplyTemplateRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.cluster_apply_template import ClusterApplyTemplate
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    apply_template_request = ApplyTemplateRequest(
        service_url="service_url_example",
        spec_template="spec_template_example",
        spec_yaml="spec_yaml_example",
    ) # ApplyTemplateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Apply Template to Cluster (Kubernetes)
        api_response = api_instance.apply_template(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->apply_template: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Apply Template to Cluster (Kubernetes)
        api_response = api_instance.apply_template(cluster_id, apply_template_request=apply_template_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->apply_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **apply_template_request** | [**ApplyTemplateRequest**](ApplyTemplateRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster**
> Model200Success delete_cluster(cluster_id)

Delete a Cluster

Will delete a cluster and associated resources, hosts, volumes asynchronously

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    remove_instances = "on" # str | Remove Instances (optional) if omitted the server will use the default value of "off"
    remove_resources = "off" # str | Remove Resources (optional) if omitted the server will use the default value of "on"
    preserve_volumes = "on" # str | Preserve Volumes (optional) if omitted the server will use the default value of "off"
    release_floating_ips = "off" # str | Release Floating IPs (optional) if omitted the server will use the default value of "on"
    release_eips = "off" # str | Alias for releaseFloatingIps (optional) if omitted the server will use the default value of "on"
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Cluster
        api_response = api_instance.delete_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Cluster
        api_response = api_instance.delete_cluster(cluster_id, remove_instances=remove_instances, remove_resources=remove_resources, preserve_volumes=preserve_volumes, release_floating_ips=release_floating_ips, release_eips=release_eips, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **remove_instances** | **str**| Remove Instances | [optional] if omitted the server will use the default value of "off"
 **remove_resources** | **str**| Remove Resources | [optional] if omitted the server will use the default value of "on"
 **preserve_volumes** | **str**| Preserve Volumes | [optional] if omitted the server will use the default value of "off"
 **release_floating_ips** | **str**| Release Floating IPs | [optional] if omitted the server will use the default value of "on"
 **release_eips** | **str**| Alias for releaseFloatingIps | [optional] if omitted the server will use the default value of "on"
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_container**
> Model200Success delete_cluster_container(cluster_id, id)

Delete Container

This endpoint deletes a specified container from a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete Container
        api_response = api_instance.delete_cluster_container(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_container: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete Container
        api_response = api_instance.delete_cluster_container(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_container: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_deployment**
> Model200Success delete_cluster_deployment(cluster_id, id)

Delete Deployment

This endpoint deletes a specified deployment from a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete Deployment
        api_response = api_instance.delete_cluster_deployment(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_deployment: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete Deployment
        api_response = api_instance.delete_cluster_deployment(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_job**
> Model200Success delete_cluster_job(cluster_id, id)

Delete a Job

This endpoint deletes a specified job from a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Job
        api_response = api_instance.delete_cluster_job(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_job: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Job
        api_response = api_instance.delete_cluster_job(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_job: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_namespace**
> Model200Success delete_cluster_namespace(cluster_id, id)

Delete a Namespace (Kubernetes)

Will delete a namespace from the specified cluster

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Namespace (Kubernetes)
        api_response = api_instance.delete_cluster_namespace(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_namespace: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Namespace (Kubernetes)
        api_response = api_instance.delete_cluster_namespace(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_namespace: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_service**
> Model200Success delete_cluster_service(cluster_id, id)

Delete a Service

This endpoint deletes a specified service from a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Service
        api_response = api_instance.delete_cluster_service(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_service: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Service
        api_response = api_instance.delete_cluster_service(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_service: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_stateful_set**
> SuccessError delete_cluster_stateful_set(cluster_id, id)

Delete a Stateful Set

Will delete a stateful set from the specified cluster

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.success_error import SuccessError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Stateful Set
        api_response = api_instance.delete_cluster_stateful_set(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_stateful_set: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Stateful Set
        api_response = api_instance.delete_cluster_stateful_set(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_stateful_set: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_volume**
> Model200Success delete_cluster_volume(cluster_id, id)

Delete a Volume

Will delete a volume from the specified cluster

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Volume
        api_response = api_instance.delete_cluster_volume(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_volume: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Volume
        api_response = api_instance.delete_cluster_volume(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_volume: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cluster_worker**
> Model200Success delete_cluster_worker(cluster_id, id)

Delete a Worker

This endpoint deletes a specified worker from a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Worker
        api_response = api_instance.delete_cluster_worker(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_worker: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Worker
        api_response = api_instance.delete_cluster_worker(cluster_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->delete_cluster_worker: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster**
> AddCluster200ResponseAllOf get_cluster(cluster_id)

Get a Specific Cluster

This endpoint retrieves a specific cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.add_cluster200_response_all_of import AddCluster200ResponseAllOf
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Cluster
        api_response = api_instance.get_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**AddCluster200ResponseAllOf**](AddCluster200ResponseAllOf.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_api_config**
> ClusterApiConfig get_cluster_api_config(cluster_id)

Get API Config

This endpoint retrieves the API configuration for a specified cluster. The configuration is cluster type specific.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.cluster_api_config import ClusterApiConfig
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Get API Config
        api_response = api_instance.get_cluster_api_config(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_api_config: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_datastore**
> GetClusterDatastore200Response get_cluster_datastore(cluster_id, id)

Get a Specific Datastore

This endpoint retrieves a specific cluster datastore.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.get_cluster_datastore200_response import GetClusterDatastore200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Datastore
        api_response = api_instance.get_cluster_datastore(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_datastore: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetClusterDatastore200Response**](GetClusterDatastore200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_history**
> GetClusterHistory200Response get_cluster_history(cluster_id)

Get Cluster History

This endpoint retrieves the process history for a specific cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cluster_history200_response import GetClusterHistory200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Get Cluster History
        api_response = api_instance.get_cluster_history(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_history: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**GetClusterHistory200Response**](GetClusterHistory200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_history_detail**
> GetClusterHistoryDetail200Response get_cluster_history_detail(cluster_id, id)

Get Cluster History Details

This endpoint retrieves the history for a specific cluster process.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.get_cluster_history_detail200_response import GetClusterHistoryDetail200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Cluster History Details
        api_response = api_instance.get_cluster_history_detail(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_history_detail: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetClusterHistoryDetail200Response**](GetClusterHistoryDetail200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_history_event_detail**
> GetClusterHistoryEventDetail200Response get_cluster_history_event_detail(cluster_id, id)

Get Cluster History Event

This endpoint retrieves the process event for a specific cluster process event.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.get_cluster_history_event_detail200_response import GetClusterHistoryEventDetail200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Cluster History Event
        api_response = api_instance.get_cluster_history_event_detail(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_history_event_detail: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetClusterHistoryEventDetail200Response**](GetClusterHistoryEventDetail200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_masters**
> GetClusterMasters200Response get_cluster_masters(cluster_id)

Get Masters (Kubernetes)

This endpoint retrieves masters of a specified kubernetes cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.get_cluster_masters200_response import GetClusterMasters200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Masters (Kubernetes)
        api_response = api_instance.get_cluster_masters(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_masters: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Masters (Kubernetes)
        api_response = api_instance.get_cluster_masters(cluster_id, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_masters: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetClusterMasters200Response**](GetClusterMasters200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_namespace**
> GetClusterNamespace200Response get_cluster_namespace(cluster_id, id)

Get Namespace (Kubernetes)

This endpoint retrieves a specific namespace of a Kubernetes cluster

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cluster_namespace200_response import GetClusterNamespace200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Namespace (Kubernetes)
        api_response = api_instance.get_cluster_namespace(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_namespace: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetClusterNamespace200Response**](GetClusterNamespace200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_namespaces**
> GetClusterNamespaces200Response get_cluster_namespaces(cluster_id)

List Namespaces (Kubernetes)

List Namespaces (Kubernetes)

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cluster_namespaces200_response import GetClusterNamespaces200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # List Namespaces (Kubernetes)
        api_response = api_instance.get_cluster_namespaces(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_namespaces: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**GetClusterNamespaces200Response**](GetClusterNamespaces200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cluster_upgrade_versions**
> GetClusterUpgradeVersions200Response get_cluster_upgrade_versions(cluster_id)

Get Cluster Upgrade Versions (Kubernetes)

This endpoint returns valid version targets for upgrading kubectl and kubeadm on the cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.get_cluster_upgrade_versions200_response import GetClusterUpgradeVersions200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Get Cluster Upgrade Versions (Kubernetes)
        api_response = api_instance.get_cluster_upgrade_versions(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_cluster_upgrade_versions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**GetClusterUpgradeVersions200Response**](GetClusterUpgradeVersions200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_wiki_cluster**
> GetWikiApp200Response get_wiki_cluster(cluster_id)

Retrieves a Cluster Wiki Page

This endpoint retrieves a cluster Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Cluster Wiki Page
        api_response = api_instance.get_wiki_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->get_wiki_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_containers**
> ListClusterContainers200Response list_cluster_containers(cluster_id)

Get Containers

This endpoint retrieves containers of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_containers200_response import ListClusterContainers200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    resource_level = "app" # str | Resource level filter (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Containers
        api_response = api_instance.list_cluster_containers(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_containers: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Containers
        api_response = api_instance.list_cluster_containers(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase, resource_level=resource_level)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_containers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **str**| Resource level filter | [optional]

### Return type

[**ListClusterContainers200Response**](ListClusterContainers200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_datastores**
> ListClusterDatastores200Response list_cluster_datastores(cluster_id)

Get Datastores

This endpoint retrieves datastores of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_datastores200_response import ListClusterDatastores200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    hide_inactive = True # bool | If true restricts query to only load active datastores (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Get Datastores
        api_response = api_instance.list_cluster_datastores(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_datastores: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Datastores
        api_response = api_instance.list_cluster_datastores(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase, name=name, code=code, hide_inactive=hide_inactive)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_datastores: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]
 **hide_inactive** | **bool**| If true restricts query to only load active datastores | [optional] if omitted the server will use the default value of False

### Return type

[**ListClusterDatastores200Response**](ListClusterDatastores200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_deployments**
> ListClusterDeployments200Response list_cluster_deployments(cluster_id)

Get Deployments

This endpoint retrieves deployments of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_cluster_deployments200_response import ListClusterDeployments200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    resource_level = "app" # str | Resource level filter (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Deployments
        api_response = api_instance.list_cluster_deployments(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_deployments: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Deployments
        api_response = api_instance.list_cluster_deployments(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase, resource_level=resource_level)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_deployments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **str**| Resource level filter | [optional]

### Return type

[**ListClusterDeployments200Response**](ListClusterDeployments200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_jobs**
> ListClusterJobs200Response list_cluster_jobs(cluster_id)

Get Jobs

This endpoint retrieves jobs of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_jobs200_response import ListClusterJobs200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Jobs
        api_response = api_instance.list_cluster_jobs(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_jobs: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Jobs
        api_response = api_instance.list_cluster_jobs(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_jobs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListClusterJobs200Response**](ListClusterJobs200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_pods**
> ListClusterPods200Response list_cluster_pods(cluster_id)

Get Pods

This endpoint retrieves pods of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_cluster_pods200_response import ListClusterPods200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    resource_level = "app" # str | Resource level filter (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Pods
        api_response = api_instance.list_cluster_pods(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_pods: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Pods
        api_response = api_instance.list_cluster_pods(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase, resource_level=resource_level)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_pods: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **str**| Resource level filter | [optional]

### Return type

[**ListClusterPods200Response**](ListClusterPods200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_services**
> ListClusterServices200Response list_cluster_services(cluster_id)

Get Services

This endpoint retrieves services of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_services200_response import ListClusterServices200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Services
        api_response = api_instance.list_cluster_services(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_services: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Services
        api_response = api_instance.list_cluster_services(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_services: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListClusterServices200Response**](ListClusterServices200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_stateful_sets**
> ListClusterStatefulSets200Response list_cluster_stateful_sets(cluster_id)

Get Stateful Sets

This endpoint retrieves stateful sets of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_cluster_stateful_sets200_response import ListClusterStatefulSets200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    resource_level = "app" # str | Resource level filter (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Stateful Sets
        api_response = api_instance.list_cluster_stateful_sets(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_stateful_sets: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Stateful Sets
        api_response = api_instance.list_cluster_stateful_sets(cluster_id, max=max, offset=offset, sort=sort, order=order, phrase=phrase, resource_level=resource_level)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_stateful_sets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **resource_level** | **str**| Resource level filter | [optional]

### Return type

[**ListClusterStatefulSets200Response**](ListClusterStatefulSets200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_types**
> ListClusterTypes200Response list_cluster_types()

Get All Cluster Types

Fetch a list of available cluster types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_types200_response import ListClusterTypes200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    provider_type = "morpheus" # str | Filter by `Provider Type` code.  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Cluster Types
        api_response = api_instance.list_cluster_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, provider_type=provider_type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]
 **provider_type** | **str**| Filter by &#x60;Provider Type&#x60; code.  | [optional]

### Return type

[**ListClusterTypes200Response**](ListClusterTypes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_volumes**
> ListClusterVolumes200Response list_cluster_volumes(cluster_id)

Get Volumes

This endpoint retrieves volumes of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_volumes200_response import ListClusterVolumes200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Get Volumes
        api_response = api_instance.list_cluster_volumes(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_volumes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**ListClusterVolumes200Response**](ListClusterVolumes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cluster_workers**
> ListClusterWorkers200Response list_cluster_workers(cluster_id)

Get Workers

This endpoint retrieves workers of a specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_cluster_workers200_response import ListClusterWorkers200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Get Workers
        api_response = api_instance.list_cluster_workers(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_cluster_workers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**ListClusterWorkers200Response**](ListClusterWorkers200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_clusters**
> ListClusters200Response list_clusters()

Get All Clusters

This endpoint retrieves all clusters and a list of clusters associated with the zone by id.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.list_clusters200_response import ListClusters200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    type_id = 3 # int | Type filter, restricts query to only load clusters of a specified cluster type (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Clusters
        api_response = api_instance.list_clusters(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, zone_id=zone_id, type_id=type_id, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->list_clusters: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **type_id** | **int**| Type filter, restricts query to only load clusters of a specified cluster type | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_cluster_container**
> SuccessError restart_cluster_container(cluster_id, id)

Restart a Container

Will restart a container in the specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.success_error import SuccessError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart a Container
        api_response = api_instance.restart_cluster_container(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->restart_cluster_container: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_cluster_deployment**
> SuccessError restart_cluster_deployment(cluster_id, id)

Restart a Deployment

Will restart a deployment in the specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.success_error import SuccessError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart a Deployment
        api_response = api_instance.restart_cluster_deployment(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->restart_cluster_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_cluster_pod**
> SuccessError restart_cluster_pod(cluster_id, id)

Restart a Pod

Will restart a pod in the specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.success_error import SuccessError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart a Pod
        api_response = api_instance.restart_cluster_pod(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->restart_cluster_pod: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_cluster_stateful_set**
> SuccessError restart_cluster_stateful_set(cluster_id, id)

Restart a Stateful Set

Will restart a stateful set in the specified cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.success_error import SuccessError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart a Stateful Set
        api_response = api_instance.restart_cluster_stateful_set(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->restart_cluster_stateful_set: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cluster**
> AddCluster200Response update_cluster(cluster_id)

Update Cluster

Update Cluster

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.add_cluster200_response import AddCluster200Response
from openapi_client.model.update_cluster_request import UpdateClusterRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    update_cluster_request = UpdateClusterRequest(
        cluster=ClusterUpdate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            enabled=True,
            service_url="service_url_example",
            service_token="service_token_example",
            refresh=True,
            managed=True,
        ),
    ) # UpdateClusterRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Cluster
        api_response = api_instance.update_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Cluster
        api_response = api_instance.update_cluster(cluster_id, update_cluster_request=update_cluster_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **update_cluster_request** | [**UpdateClusterRequest**](UpdateClusterRequest.md)|  | [optional]

### Return type

[**AddCluster200Response**](AddCluster200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cluster_datastore**
> UpdateClusterDatastore200Response update_cluster_datastore(cluster_id, id)

Update Datastore

Update Datastore

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.update_cluster_datastore_request import UpdateClusterDatastoreRequest
from openapi_client.model.update_cluster_datastore200_response import UpdateClusterDatastore200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cluster_datastore_request = UpdateClusterDatastoreRequest(
        datastore=ClusterDatastoreUpdate(
            active=True,
            permissions=ClusterUpdatePermissions(
                resource_permissions=ClusterUpdatePermissionsResourcePermissions(
                    all=True,
                    sites=[
                        ClusterUpdatePermissionsResourcePermissionsSitesInner(
                            id=1,
                            default=True,
                        ),
                    ],
                    all_plans=True,
                    plans=[
                        ClusterUpdatePermissionsResourcePermissionsSitesInner(
                            id=1,
                            default=True,
                        ),
                    ],
                ),
                resource_pool=ClusterUpdatePermissionsResourcePool(
                    visibility="visibility_example",
                ),
                tenant_permissions=ClusterUpdatePermissionsTenantPermissions(
                    accounts=[
                        1,
                    ],
                ),
            ),
            visibility="private",
        ),
    ) # UpdateClusterDatastoreRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Datastore
        api_response = api_instance.update_cluster_datastore(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_datastore: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Datastore
        api_response = api_instance.update_cluster_datastore(cluster_id, id, update_cluster_datastore_request=update_cluster_datastore_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_datastore: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cluster_datastore_request** | [**UpdateClusterDatastoreRequest**](UpdateClusterDatastoreRequest.md)|  | [optional]

### Return type

[**UpdateClusterDatastore200Response**](UpdateClusterDatastore200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cluster_namespace**
> AddClusterNamespace200Response update_cluster_namespace(cluster_id, id)

Update Namespace (Kubernetes)

Update Namespace (Kubernetes)

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.update_cluster_namespace_request import UpdateClusterNamespaceRequest
from openapi_client.model.add_cluster_namespace200_response import AddClusterNamespace200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cluster_namespace_request = UpdateClusterNamespaceRequest(
        namespace=ClusterNamespaceUpdate(
            name="name_example",
            description="description_example",
            active=False,
            permissions=ClusterNamespaceUpdatePermissions(
                resource_permissions=ClusterNamespaceCreateResourcePermissions(
                    all=True,
                    sites=[
                        ClusterUpdatePermissionsResourcePermissionsSitesInner(
                            id=1,
                            default=True,
                        ),
                    ],
                    all_plans=True,
                    plans=[
                        ClusterUpdatePermissionsResourcePermissionsSitesInner(
                            id=1,
                            default=True,
                        ),
                    ],
                ),
            ),
        ),
    ) # UpdateClusterNamespaceRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Namespace (Kubernetes)
        api_response = api_instance.update_cluster_namespace(cluster_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_namespace: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Namespace (Kubernetes)
        api_response = api_instance.update_cluster_namespace(cluster_id, id, update_cluster_namespace_request=update_cluster_namespace_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_namespace: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cluster_namespace_request** | [**UpdateClusterNamespaceRequest**](UpdateClusterNamespaceRequest.md)|  | [optional]

### Return type

[**AddClusterNamespace200Response**](AddClusterNamespace200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cluster_permissions**
> AddCluster200Response update_cluster_permissions(cluster_id)

Update Cluster Permissions

Update Cluster Permissions

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.update_cluster_permissions_request import UpdateClusterPermissionsRequest
from openapi_client.model.add_cluster200_response import AddCluster200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    update_cluster_permissions_request = UpdateClusterPermissionsRequest(
        permissions=ClusterUpdatePermissions(
            resource_permissions=ClusterUpdatePermissionsResourcePermissions(
                all=True,
                sites=[
                    ClusterUpdatePermissionsResourcePermissionsSitesInner(
                        id=1,
                        default=True,
                    ),
                ],
                all_plans=True,
                plans=[
                    ClusterUpdatePermissionsResourcePermissionsSitesInner(
                        id=1,
                        default=True,
                    ),
                ],
            ),
            resource_pool=ClusterUpdatePermissionsResourcePool(
                visibility="visibility_example",
            ),
            tenant_permissions=ClusterUpdatePermissionsTenantPermissions(
                accounts=[
                    1,
                ],
            ),
        ),
    ) # UpdateClusterPermissionsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Cluster Permissions
        api_response = api_instance.update_cluster_permissions(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_permissions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Cluster Permissions
        api_response = api_instance.update_cluster_permissions(cluster_id, update_cluster_permissions_request=update_cluster_permissions_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_permissions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **update_cluster_permissions_request** | [**UpdateClusterPermissionsRequest**](UpdateClusterPermissionsRequest.md)|  | [optional]

### Return type

[**AddCluster200Response**](AddCluster200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cluster_upgrade_versions**
> Model200Success update_cluster_upgrade_versions(cluster_id, target_version)

Upgrade a Cluster (Kubernetes)

This endpoint updates the kubectl and kudeadm versions on a Kubernetes cluster to the specified version. Use Get Cluster Upgrade Versions to list valid version targets for the cluster.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    target_version = "1.21.14" # str | Target version for cluster after upgrade

    # example passing only required values which don't have defaults set
    try:
        # Upgrade a Cluster (Kubernetes)
        api_response = api_instance.update_cluster_upgrade_versions(cluster_id, target_version)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_upgrade_versions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **target_version** | **str**| Target version for cluster after upgrade |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cluster_worker_count**
> Model200Success update_cluster_worker_count(cluster_id, worker_count)

Update Worker Count

This endpoint resizes a cluster to the specified number of worker nodes (only supports Azure AKS, Google GKE, and Amazon EKS clusters).

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    worker_count = 5 # int | The target number of worker nodes

    # example passing only required values which don't have defaults set
    try:
        # Update Worker Count
        api_response = api_instance.update_cluster_worker_count(cluster_id, worker_count)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_cluster_worker_count: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **worker_count** | **int**| The target number of worker nodes |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_wiki_cluster**
> UpdateWikiApp200Response update_wiki_cluster(cluster_id)

Update a Cluster Wiki Page

Updates a cluster Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clusters_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = clusters_api.ClustersApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Cluster Wiki Page
        api_response = api_instance.update_wiki_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_wiki_cluster: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Cluster Wiki Page
        api_response = api_instance.update_wiki_cluster(cluster_id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClustersApi->update_wiki_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

