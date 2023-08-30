# openapi_client.ClusterLayoutsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_cluster_layout_clone**](ClusterLayoutsApi.md#add_cluster_layout_clone) | **POST** /api/library/cluster-layouts/{id}/clone | Clone a Cluster Layout
[**add_cluster_layouts**](ClusterLayoutsApi.md#add_cluster_layouts) | **POST** /api/library/cluster-layouts | Create a Cluster Layout
[**delete_cluster_layout**](ClusterLayoutsApi.md#delete_cluster_layout) | **DELETE** /api/library/cluster-layouts/{id} | Delete a Cluster Layout
[**get_cluster_layout**](ClusterLayoutsApi.md#get_cluster_layout) | **GET** /api/library/cluster-layouts/{id} | Get a Specific Cluster Layout
[**list_cluster_layouts**](ClusterLayoutsApi.md#list_cluster_layouts) | **GET** /api/library/cluster-layouts | Get All Cluster Layouts
[**update_cluster_layout**](ClusterLayoutsApi.md#update_cluster_layout) | **PUT** /api/library/cluster-layouts/{id} | Update a Cluster Layout


# **add_cluster_layout_clone**
> SuccessId add_cluster_layout_clone(id)

Clone a Cluster Layout

Use this command to clone a cluster layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cluster_layouts_api
from openapi_client.model.success_id import SuccessId
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
    api_instance = cluster_layouts_api.ClusterLayoutsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    name = "New Name" # str | Name of cluster layout. Defaults to Copy of <cloned layout name> (optional)
    description = "New Description" # str | Description of cluster layout. Defaults to cloned layout description (optional)
    compute_version = "2.0" # str | Version of cluster layout. Defaults to cloned layout version (optional)

    # example passing only required values which don't have defaults set
    try:
        # Clone a Cluster Layout
        api_response = api_instance.add_cluster_layout_clone(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->add_cluster_layout_clone: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Clone a Cluster Layout
        api_response = api_instance.add_cluster_layout_clone(id, name=name, description=description, compute_version=compute_version)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->add_cluster_layout_clone: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **name** | **str**| Name of cluster layout. Defaults to Copy of &lt;cloned layout name&gt; | [optional]
 **description** | **str**| Description of cluster layout. Defaults to cloned layout description | [optional]
 **compute_version** | **str**| Version of cluster layout. Defaults to cloned layout version | [optional]

### Return type

[**SuccessId**](SuccessId.md)

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

# **add_cluster_layouts**
> SuccessId add_cluster_layouts()

Create a Cluster Layout

Use this command to create a cluster layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cluster_layouts_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_cluster_layouts_request import AddClusterLayoutsRequest
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
    api_instance = cluster_layouts_api.ClusterLayoutsApi(api_client)
    add_cluster_layouts_request = AddClusterLayoutsRequest(
        layout=ClusterLayoutCreate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            compute_version="compute_version_example",
            creatable=True,
            has_auto_scale=False,
            install_container_runtime=False,
            memory_requirement=1,
            group_type=ClusterLayoutCreateGroupType(
                id=1,
            ),
            provision_type=AddServicePlansRequestServicePlanProvisionTypeInner(
                id=1,
            ),
            option_types=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            task_sets=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            masters=[
                ClusterLayoutCreateMastersInner(
                    node_count=1,
                    container_type=AddStorageVolumesRequestStorageVolumeStorageServer(
                        id=1,
                    ),
                ),
            ],
            workers=[
                ClusterLayoutCreateMastersInner(
                    node_count=1,
                    container_type=AddStorageVolumesRequestStorageVolumeStorageServer(
                        id=1,
                    ),
                ),
            ],
        ),
    ) # AddClusterLayoutsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Cluster Layout
        api_response = api_instance.add_cluster_layouts(add_cluster_layouts_request=add_cluster_layouts_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->add_cluster_layouts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_cluster_layouts_request** | [**AddClusterLayoutsRequest**](AddClusterLayoutsRequest.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

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

# **delete_cluster_layout**
> Model200Success delete_cluster_layout(id)

Delete a Cluster Layout

Will delete a cluster layout

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cluster_layouts_api
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
    api_instance = cluster_layouts_api.ClusterLayoutsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Cluster Layout
        api_response = api_instance.delete_cluster_layout(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->delete_cluster_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

# **get_cluster_layout**
> GetClusterLayout200Response get_cluster_layout(id)

Get a Specific Cluster Layout

This endpoint retrieves a specific cluster layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cluster_layouts_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cluster_layout200_response import GetClusterLayout200Response
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
    api_instance = cluster_layouts_api.ClusterLayoutsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Cluster Layout
        api_response = api_instance.get_cluster_layout(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->get_cluster_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetClusterLayout200Response**](GetClusterLayout200Response.md)

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

# **list_cluster_layouts**
> ListClusterLayouts200Response list_cluster_layouts()

Get All Cluster Layouts

This endpoint retrieves all cluster layouts.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cluster_layouts_api
from openapi_client.model.list_cluster_layouts200_response import ListClusterLayouts200Response
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
    api_instance = cluster_layouts_api.ClusterLayoutsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    provision_type = "AKS" # str | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Cluster Layouts
        api_response = api_instance.list_cluster_layouts(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, provision_type=provision_type, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->list_cluster_layouts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **provision_type** | **str**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListClusterLayouts200Response**](ListClusterLayouts200Response.md)

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

# **update_cluster_layout**
> SuccessId update_cluster_layout(id)

Update a Cluster Layout

Use this command to update an existing cluster layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cluster_layouts_api
from openapi_client.model.update_cluster_layout_request import UpdateClusterLayoutRequest
from openapi_client.model.success_id import SuccessId
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
    api_instance = cluster_layouts_api.ClusterLayoutsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cluster_layout_request = UpdateClusterLayoutRequest(
        layout=ClusterLayoutUpdate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            compute_version="compute_version_example",
            creatable=True,
            has_auto_scale=False,
            install_container_runtime=False,
            memory_requirement=1,
            group_type=ClusterLayoutCreateGroupType(
                id=1,
            ),
            provision_type=AddServicePlansRequestServicePlanProvisionTypeInner(
                id=1,
            ),
            option_types=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            task_sets=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            masters=[
                ClusterLayoutCreateMastersInner(
                    node_count=1,
                    container_type=AddStorageVolumesRequestStorageVolumeStorageServer(
                        id=1,
                    ),
                ),
            ],
            workers=[
                ClusterLayoutCreateMastersInner(
                    node_count=1,
                    container_type=AddStorageVolumesRequestStorageVolumeStorageServer(
                        id=1,
                    ),
                ),
            ],
        ),
    ) # UpdateClusterLayoutRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Cluster Layout
        api_response = api_instance.update_cluster_layout(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->update_cluster_layout: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Cluster Layout
        api_response = api_instance.update_cluster_layout(id, update_cluster_layout_request=update_cluster_layout_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClusterLayoutsApi->update_cluster_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cluster_layout_request** | [**UpdateClusterLayoutRequest**](UpdateClusterLayoutRequest.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

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

