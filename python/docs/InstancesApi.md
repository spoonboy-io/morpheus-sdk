# openapi_client.InstancesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_instance**](InstancesApi.md#add_instance) | **POST** /api/instances | Create an Instance
[**backup_instance**](InstancesApi.md#backup_instance) | **PUT** /api/instances/{id}/backup | Backup an instance
[**backups_instance**](InstancesApi.md#backups_instance) | **GET** /api/instances/{id}/backups | Get list of backups for an Instance
[**cancel_expiration_instance**](InstancesApi.md#cancel_expiration_instance) | **PUT** /api/instances/{id}/cancel-expiration | Cancel Expiration of an Instance
[**cancel_removal_instance**](InstancesApi.md#cancel_removal_instance) | **PUT** /api/instances/{id}/cancel-removal | Cancel Removal of an Instance
[**cancel_shutdown_instance**](InstancesApi.md#cancel_shutdown_instance) | **PUT** /api/instances/{id}/cancel-shutdown | Cancel Shutdown of an Instance
[**clone_image_instance**](InstancesApi.md#clone_image_instance) | **PUT** /api/instances/{id}/clone-image | Clone to Image
[**clone_instance**](InstancesApi.md#clone_instance) | **PUT** /api/instances/{id}/clone | Clone an Instance
[**create_instance_schedule**](InstancesApi.md#create_instance_schedule) | **POST** /api/instances/{id}/schedules | Create a new Instance Schedule
[**delete_all_snapshots_instance**](InstancesApi.md#delete_all_snapshots_instance) | **DELETE** /api/instances/{id}/delete-all-snapshots | Delete All Snapshots of Instance
[**delete_all_snapshots_instance_container**](InstancesApi.md#delete_all_snapshots_instance_container) | **DELETE** /api/instances/{id}/delete-container-snapshots/{containerId} | Delete All Snapshots of Instance Container
[**delete_instance**](InstancesApi.md#delete_instance) | **DELETE** /api/instances/{id} | Delete an instance
[**delete_instance_schedule**](InstancesApi.md#delete_instance_schedule) | **DELETE** /api/instances/{id}/schedules/{scheduleId} | Deletes an Instance Schedule
[**delete_snapshot_instance**](InstancesApi.md#delete_snapshot_instance) | **DELETE** /api/snapshots/{id} | Delete Snapshot of Instance
[**eject_instance**](InstancesApi.md#eject_instance) | **PUT** /api/instances/{id}/eject | Eject an instance
[**extend_expiration_instance**](InstancesApi.md#extend_expiration_instance) | **PUT** /api/instances/{id}/extend-expiration | Extend Expiration of an Instance
[**extend_shutdown_instance**](InstancesApi.md#extend_shutdown_instance) | **PUT** /api/instances/{id}/extend-shutdown | Extend Shutdown of an Instance
[**get_env_variables**](InstancesApi.md#get_env_variables) | **GET** /api/instances/{id}/envs | Get Env Variables
[**get_instance**](InstancesApi.md#get_instance) | **GET** /api/instances/{id} | Retrieves a Specific Instance
[**get_instance_containers**](InstancesApi.md#get_instance_containers) | **GET** /api/instances/{id}/containers | Get Container Details
[**get_instance_history**](InstancesApi.md#get_instance_history) | **GET** /api/instances/{id}/history | Get Instance History
[**get_instance_schedule**](InstancesApi.md#get_instance_schedule) | **GET** /api/instances/{id}/schedules/{scheduleId} | Get a Specific Instance Schedule
[**get_instance_schedules**](InstancesApi.md#get_instance_schedules) | **GET** /api/instances/{id}/schedules | Get all Instance Schedules
[**get_instance_threshold**](InstancesApi.md#get_instance_threshold) | **GET** /api/instances/{id}/threshold | Get an Instance Scale Threshold
[**get_instance_type_provisioning**](InstancesApi.md#get_instance_type_provisioning) | **GET** /api/instance-types/{id} | Get Specific Instance Type for Provisioning
[**get_prepare_apply_instance**](InstancesApi.md#get_prepare_apply_instance) | **GET** /api/instances/{id}/prepare-apply | Prepare To Apply an Instance
[**get_snapshot_instance**](InstancesApi.md#get_snapshot_instance) | **GET** /api/snapshots/{id} | Get a Specific Snapshot
[**get_state_instance**](InstancesApi.md#get_state_instance) | **GET** /api/instances/{id}/state | Get State of an Instance
[**get_validate_apply_instance**](InstancesApi.md#get_validate_apply_instance) | **POST** /api/instances/{id}/validate-apply | Validate Apply State for an Instance
[**get_wiki_instance**](InstancesApi.md#get_wiki_instance) | **GET** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**import_snapshot_instance**](InstancesApi.md#import_snapshot_instance) | **PUT** /api/instances/{id}/import-snapshot | Import Snapshot of an Instance
[**linked_clone_snapshot_instance**](InstancesApi.md#linked_clone_snapshot_instance) | **PUT** /api/instances/{id}/linked-clone/{snapshotId} | Create Linked Clone of Instance Snapshot
[**list_instance_service_plans**](InstancesApi.md#list_instance_service_plans) | **GET** /api/instances/service-plans | Get Available Service Plans for an Instance
[**list_instance_types_provisioning**](InstancesApi.md#list_instance_types_provisioning) | **GET** /api/instance-types | Get All Instance Types for Provisioning
[**list_instances**](InstancesApi.md#list_instances) | **GET** /api/instances | Get All Instances
[**list_security_groups_instance**](InstancesApi.md#list_security_groups_instance) | **GET** /api/instances/{id}/security-groups | Get Security Groups for an Instance
[**lock_instance**](InstancesApi.md#lock_instance) | **PUT** /api/instances/{id}/lock | Lock an Instance
[**refresh_state_instance**](InstancesApi.md#refresh_state_instance) | **POST** /api/instances/{id}/refresh | Refresh State of an Instance
[**remove_instances_from_control**](InstancesApi.md#remove_instances_from_control) | **POST** /api/instances/remove-from-control | Remove From Control
[**resize_instance**](InstancesApi.md#resize_instance) | **PUT** /api/instances/{id}/resize | Resize an Instance
[**restart_instance**](InstancesApi.md#restart_instance) | **PUT** /api/instances/{id}/restart | Restart an instance
[**revert_snapshot_instance**](InstancesApi.md#revert_snapshot_instance) | **PUT** /api/instances/{id}/revert-snapshot/{snapshotId} | Revert Instance to Snapshot
[**run_workflow_instance**](InstancesApi.md#run_workflow_instance) | **PUT** /api/instances/{id}/workflow | Run Workflow on an Instance
[**set_apply_instance**](InstancesApi.md#set_apply_instance) | **POST** /api/instances/{id}/apply | Apply State of an Instance
[**set_instance_security_groups**](InstancesApi.md#set_instance_security_groups) | **POST** /api/instances/{id}/security-groups | Set Security Groups for an Instance
[**snapshot_instance**](InstancesApi.md#snapshot_instance) | **PUT** /api/instances/{id}/snapshot | Snapshot an Instance
[**snapshots_instance**](InstancesApi.md#snapshots_instance) | **GET** /api/instances/{id}/snapshots | Get list of snapshots for an Instance
[**start_instance**](InstancesApi.md#start_instance) | **PUT** /api/instances/{id}/start | Start an instance
[**stop_instance**](InstancesApi.md#stop_instance) | **PUT** /api/instances/{id}/stop | Stop an instance
[**suspend_instance**](InstancesApi.md#suspend_instance) | **PUT** /api/instances/{id}/suspend | Suspend an instance
[**unlock_instance**](InstancesApi.md#unlock_instance) | **PUT** /api/instances/{id}/unlock | Unlock an Instance
[**update_instance**](InstancesApi.md#update_instance) | **PUT** /api/instances/{id} | Updating an Instance
[**update_instance_network_interface**](InstancesApi.md#update_instance_network_interface) | **PUT** /api/instances/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for an Instance&#39;s Network
[**update_instance_schedule**](InstancesApi.md#update_instance_schedule) | **PUT** /api/instances/{id}/schedules/{scheduleId} | Updating an Instance Schedule
[**update_instance_threshold**](InstancesApi.md#update_instance_threshold) | **PUT** /api/instances/{id}/threshold | Updates an Instance Scale Threshold
[**update_wiki_instance**](InstancesApi.md#update_wiki_instance) | **PUT** /api/instances/{id}/wiki | Update an Instance Wiki Page


# **add_instance**
> AddInstance200Response add_instance()

Create an Instance

Create an Instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_create import InstanceCreate
from openapi_client.model.add_instance200_response import AddInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    instance_create = InstanceCreate(
        instance=InstanceCreateInstance(
            name="myinstance",
            site=InstanceCreateInstanceSite(
                id=2,
            ),
            instance_type=InstanceCreateInstanceInstanceType(
                code="Ubuntu",
            ),
            layout=InstanceCreateInstanceLayout(
                id=105,
            ),
            plan=InstanceCreateInstancePlan(
                id=75,
            ),
            instance_context="dev",
        ),
        zone_id=6,
        evars=[
            UpdateHostManagedRequestServerTagsInner(
                name="name_example",
                value="value_example",
            ),
        ],
        copies=1,
        layout_size=1,
        service_plan_options={},
        security_groups=[
            {},
        ],
        volumes=[
            InstanceCreateVolume(
                id=1,
                root_volume=True,
                name="data",
                size=5,
                size_id=1,
                storage_type=1,
                datastore_id=InstanceCreateVolumeDatastoreId(None),
            ),
        ],
        network_interfaces=[
            InstanceCreateNetwork(
                network=InstanceCreateNetworkNetwork(
                    id="id_example",
                ),
                network_interface_type_id=1,
                ip_address="ip_address_example",
                id=1,
            ),
        ],
        config=CatalogItemTypeInstanceScribeConfig(None),
        labels=[
            "labels_example",
        ],
        tags=[
            UpdateHostManagedRequestServerTagsInner(
                name="name_example",
                value="value_example",
            ),
        ],
        metadata=[
            UpdateHostManagedRequestServerTagsInner(
                name="name_example",
                value="value_example",
            ),
        ],
        ports=[
            InstanceCreatePortsInner(
                port=8080,
                name="name_example",
                lb="lb_example",
            ),
        ],
        task_set_id=1,
        task_set_name="task_set_name_example",
    ) # InstanceCreate |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Instance
        api_response = api_instance.add_instance(instance_create=instance_create)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->add_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_create** | [**InstanceCreate**](InstanceCreate.md)|  | [optional]

### Return type

[**AddInstance200Response**](AddInstance200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **backup_instance**
> Model200Success backup_instance(id)

Backup an instance

Backup an instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Backup an instance
        api_response = api_instance.backup_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->backup_instance: %s\n" % e)
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

# **backups_instance**
> InstanceBackups backups_instance(id)

Get list of backups for an Instance

Get list of backups for an Instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_backups import InstanceBackups
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get list of backups for an Instance
        api_response = api_instance.backups_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->backups_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cancel_expiration_instance**
> Model200Success cancel_expiration_instance(id)

Cancel Expiration of an Instance

This operation will cancel the expiration of an instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Cancel Expiration of an Instance
        api_response = api_instance.cancel_expiration_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->cancel_expiration_instance: %s\n" % e)
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

# **cancel_removal_instance**
> Model200Success cancel_removal_instance(id)

Cancel Removal of an Instance

This operation will undo the delete of an instance that is pending removal.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Cancel Removal of an Instance
        api_response = api_instance.cancel_removal_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->cancel_removal_instance: %s\n" % e)
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

# **cancel_shutdown_instance**
> Model200Success cancel_shutdown_instance(id)

Cancel Shutdown of an Instance

This operation will cancel the shutdown of an instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Cancel Shutdown of an Instance
        api_response = api_instance.cancel_shutdown_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->cancel_shutdown_instance: %s\n" % e)
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

# **clone_image_instance**
> Model200Success clone_image_instance(id)

Clone to Image

This endpoint allows creating an image template from an existing instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.instances_clone_image import InstancesCloneImage
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    instances_clone_image = InstancesCloneImage(
        template_name="{server.name}-{timestamp}",
        zone_folder="group-v81",
    ) # InstancesCloneImage |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Clone to Image
        api_response = api_instance.clone_image_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->clone_image_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Clone to Image
        api_response = api_instance.clone_image_instance(id, instances_clone_image=instances_clone_image)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->clone_image_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instances_clone_image** | [**InstancesCloneImage**](InstancesCloneImage.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **clone_instance**
> Model200Success clone_instance(id)

Clone an Instance

One can easily clone an instance and all containers within that instance. The containers are backed up via the backup services and used as a snapshot to produce a clone of the instance. It is possible to clone this app instance into an entirely different availability zone.  This endpoint also supports all of the same parameters as instance creation, so you can override any configuration options when provisioning the clone. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_clone import InstanceClone
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    instance_clone = InstanceClone(
        name="my-cloned-instance",
        group=InstanceCloneGroup(
            id=2,
        ),
    ) # InstanceClone |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Clone an Instance
        api_response = api_instance.clone_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->clone_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Clone an Instance
        api_response = api_instance.clone_instance(id, instance_clone=instance_clone)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->clone_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_clone** | [**InstanceClone**](InstanceClone.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_instance_schedule**
> CreateInstanceSchedule200Response create_instance_schedule(id)

Create a new Instance Schedule

Create a new schedule for a specific instance.  This creates an instance scaling threshold that only applies during a defined schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.create_instance_schedule200_response import CreateInstanceSchedule200Response
from openapi_client.model.create_instance_schedule_request import CreateInstanceScheduleRequest
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 42 # int | Instance ID
    create_instance_schedule_request = CreateInstanceScheduleRequest(
        instance_schedule=InstanceScheduleCreate(
            schedule_type="dayOfWeek",
            schedule_timezone="UTC",
            start_day_of_week=1,
            start_time="01:00",
            end_day_of_week=7,
            end_time="20:15",
            start_date=dateutil_parser('2022-12-25T00:00:00Z'),
            end_date=dateutil_parser('2023-01-01T00:00:00Z'),
            threshold=InstanceScheduleCreateThreshold(
                source_threshold_id=1,
                auto_up=False,
                auto_down=False,
                min_count=1,
                max_count=3,
                cpu_enabled=False,
                min_cpu=0,
                max_cpu=0,
                memory_enabled=False,
                min_memory=0,
                max_memory=0,
                disk_enabled=False,
                min_disk=0,
                max_disk=0,
            ),
        ),
    ) # CreateInstanceScheduleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a new Instance Schedule
        api_response = api_instance.create_instance_schedule(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->create_instance_schedule: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a new Instance Schedule
        api_response = api_instance.create_instance_schedule(id, create_instance_schedule_request=create_instance_schedule_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->create_instance_schedule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **create_instance_schedule_request** | [**CreateInstanceScheduleRequest**](CreateInstanceScheduleRequest.md)|  | [optional]

### Return type

[**CreateInstanceSchedule200Response**](CreateInstanceSchedule200Response.md)

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

# **delete_all_snapshots_instance**
> Model200Success delete_all_snapshots_instance(id)

Delete All Snapshots of Instance

Delete All Snapshots attached to Instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete All Snapshots of Instance
        api_response = api_instance.delete_all_snapshots_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->delete_all_snapshots_instance: %s\n" % e)
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

# **delete_all_snapshots_instance_container**
> Model200Success delete_all_snapshots_instance_container(id, container_id)

Delete All Snapshots of Instance Container

Delete All Snapshots attached to Instance Container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    container_id = 4 # float | Container ID

    # example passing only required values which don't have defaults set
    try:
        # Delete All Snapshots of Instance Container
        api_response = api_instance.delete_all_snapshots_instance_container(id, container_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->delete_all_snapshots_instance_container: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **container_id** | **float**| Container ID |

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

# **delete_instance**
> Model200Success delete_instance(id)

Delete an instance

Will delete an instance and all associated monitors and backups.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    preserve_volumes = "on" # str | Preserve Volumes (optional) if omitted the server will use the default value of "off"
    keep_backups = "on" # str | Preserve copy of backups (optional) if omitted the server will use the default value of "off"
    release_floating_ips = "off" # str | Release Floating IPs (optional) if omitted the server will use the default value of "on"
    release_eips = "off" # str | Alias for releaseFloatingIps (optional) if omitted the server will use the default value of "on"
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete an instance
        api_response = api_instance.delete_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->delete_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete an instance
        api_response = api_instance.delete_instance(id, preserve_volumes=preserve_volumes, keep_backups=keep_backups, release_floating_ips=release_floating_ips, release_eips=release_eips, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->delete_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **preserve_volumes** | **str**| Preserve Volumes | [optional] if omitted the server will use the default value of "off"
 **keep_backups** | **str**| Preserve copy of backups | [optional] if omitted the server will use the default value of "off"
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

# **delete_instance_schedule**
> Model200Success delete_instance_schedule(id, schedule_id)

Deletes an Instance Schedule

Deletes a specified instance scaling schedule 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 42 # int | Instance ID
    schedule_id = 1 # int | Instance Schedule ID

    # example passing only required values which don't have defaults set
    try:
        # Deletes an Instance Schedule
        api_response = api_instance.delete_instance_schedule(id, schedule_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->delete_instance_schedule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **schedule_id** | **int**| Instance Schedule ID |

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

# **delete_snapshot_instance**
> Model200Success delete_snapshot_instance(id)

Delete Snapshot of Instance

Delete snapshot of instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete Snapshot of Instance
        api_response = api_instance.delete_snapshot_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->delete_snapshot_instance: %s\n" % e)
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

# **eject_instance**
> Model200Success eject_instance(id)

Eject an instance

This will eject any ISO media on all containers in the instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Eject an instance
        api_response = api_instance.eject_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->eject_instance: %s\n" % e)
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

# **extend_expiration_instance**
> Model200Success extend_expiration_instance(id)

Extend Expiration of an Instance

This operation will extend the expiration of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Extend Expiration of an Instance
        api_response = api_instance.extend_expiration_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->extend_expiration_instance: %s\n" % e)
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

# **extend_shutdown_instance**
> Model200Success extend_shutdown_instance(id)

Extend Shutdown of an Instance

This operation will extend the shutdown of an instance. The period of time it is extended is equal to the number of renewal days in the expiration policy.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Extend Shutdown of an Instance
        api_response = api_instance.extend_shutdown_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->extend_shutdown_instance: %s\n" % e)
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

# **get_env_variables**
> GetEnvVariables200Response get_env_variables(id)

Get Env Variables

This gets all the environment variables associated with the instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_env_variables200_response import GetEnvVariables200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Env Variables
        api_response = api_instance.get_env_variables(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_env_variables: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetEnvVariables200Response**](GetEnvVariables200Response.md)

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

# **get_instance**
> GetInstance200Response get_instance(id)

Retrieves a Specific Instance

Retrieves a specific instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_instance200_response import GetInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    details = True # bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Instance
        api_response = api_instance.get_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves a Specific Instance
        api_response = api_instance.get_instance(id, details=details)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **details** | **bool**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] if omitted the server will use the default value of False

### Return type

[**GetInstance200Response**](GetInstance200Response.md)

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

# **get_instance_containers**
> GetInstanceContainers200Response get_instance_containers(id)

Get Container Details

This can be valuable for evaluating the details of the compute server(s) running on an instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_instance_containers200_response import GetInstanceContainers200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Container Details
        api_response = api_instance.get_instance_containers(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_containers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetInstanceContainers200Response**](GetInstanceContainers200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_instance_history**
> ListHistory200Response get_instance_history(id)

Get Instance History

This endpoint retrieves the process history for a specific instance. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.list_history200_response import ListHistory200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    container_id = 135 # int | The Container ID for Filtering (optional)
    server_id = 97 # int | The Server ID for Filtering (optional)
    zone_id = 3 # int | The Zone ID for Filtering (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Instance History
        api_response = api_instance.get_instance_history(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_history: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Instance History
        api_response = api_instance.get_instance_history(id, phrase=phrase, container_id=container_id, server_id=server_id, zone_id=zone_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_history: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]

### Return type

[**ListHistory200Response**](ListHistory200Response.md)

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

# **get_instance_schedule**
> CreateInstanceSchedule200ResponseAllOf get_instance_schedule(id, schedule_id)

Get a Specific Instance Schedule

This endpoint retrieves a specific instance scaling schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.create_instance_schedule200_response_all_of import CreateInstanceSchedule200ResponseAllOf
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 42 # int | Instance ID
    schedule_id = 1 # int | Instance Schedule ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Instance Schedule
        api_response = api_instance.get_instance_schedule(id, schedule_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_schedule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **schedule_id** | **int**| Instance Schedule ID |

### Return type

[**CreateInstanceSchedule200ResponseAllOf**](CreateInstanceSchedule200ResponseAllOf.md)

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

# **get_instance_schedules**
> GetInstanceSchedules200Response get_instance_schedules(id)

Get all Instance Schedules

This endpoint retrieves all the scaling threshold schedules for a specific instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_instance_schedules200_response import GetInstanceSchedules200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 42 # int | Instance ID

    # example passing only required values which don't have defaults set
    try:
        # Get all Instance Schedules
        api_response = api_instance.get_instance_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |

### Return type

[**GetInstanceSchedules200Response**](GetInstanceSchedules200Response.md)

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

# **get_instance_threshold**
> GetInstanceThreshold200Response get_instance_threshold(id)

Get an Instance Scale Threshold

Retrieves the scale threshold settings for a specific instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_instance_threshold200_response import GetInstanceThreshold200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get an Instance Scale Threshold
        api_response = api_instance.get_instance_threshold(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_threshold: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetInstanceThreshold200Response**](GetInstanceThreshold200Response.md)

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

# **get_instance_type_provisioning**
> GetInstanceTypeProvisioning200Response get_instance_type_provisioning(id)

Get Specific Instance Type for Provisioning

Fetch an instance type by ID. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_instance_type_provisioning200_response import GetInstanceTypeProvisioning200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Specific Instance Type for Provisioning
        api_response = api_instance.get_instance_type_provisioning(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_instance_type_provisioning: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetInstanceTypeProvisioning200Response**](GetInstanceTypeProvisioning200Response.md)

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

# **get_prepare_apply_instance**
> GetPrepareApplyInstance200Response get_prepare_apply_instance(id)

Prepare To Apply an Instance

This endpoint provides a way to view the current instance configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_prepare_apply_instance200_response import GetPrepareApplyInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Prepare To Apply an Instance
        api_response = api_instance.get_prepare_apply_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_prepare_apply_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPrepareApplyInstance200Response**](GetPrepareApplyInstance200Response.md)

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

# **get_snapshot_instance**
> Snapshot get_snapshot_instance(id)

Get a Specific Snapshot

This endpoint retrieves a specific snapshot.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.snapshot import Snapshot
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Snapshot
        api_response = api_instance.get_snapshot_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_snapshot_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_state_instance**
> GetStateInstance200Response get_state_instance(id)

Get State of an Instance

This endpoint provides a way to view the state of an instance. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.get_state_instance200_response import GetStateInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get State of an Instance
        api_response = api_instance.get_state_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_state_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetStateInstance200Response**](GetStateInstance200Response.md)

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

# **get_validate_apply_instance**
> ValidateAppState200Response get_validate_apply_instance(id)

Validate Apply State for an Instance

This endpoint provides a way to validate instance configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.apply_app_state_request import ApplyAppStateRequest
from openapi_client.model.validate_app_state200_response import ValidateAppState200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    apply_app_state_request = ApplyAppStateRequest(
        template_parameter={},
    ) # ApplyAppStateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Validate Apply State for an Instance
        api_response = api_instance.get_validate_apply_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_validate_apply_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Validate Apply State for an Instance
        api_response = api_instance.get_validate_apply_instance(id, apply_app_state_request=apply_app_state_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_validate_apply_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **apply_app_state_request** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional]

### Return type

[**ValidateAppState200Response**](ValidateAppState200Response.md)

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

# **get_wiki_instance**
> GetWikiApp200Response get_wiki_instance(id)

Retrieves an Instance Wiki Page

This endpoint retrieves an instance Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves an Instance Wiki Page
        api_response = api_instance.get_wiki_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->get_wiki_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

# **import_snapshot_instance**
> Model200Success import_snapshot_instance(id)

Import Snapshot of an Instance

It is possible to import a snapshot of an instance. This creates a Virtual Image of the instance as it currently exists.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.import_snapshot_instance_request import ImportSnapshotInstanceRequest
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    import_snapshot_instance_request = ImportSnapshotInstanceRequest(
        storage_provider_id=1,
    ) # ImportSnapshotInstanceRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Import Snapshot of an Instance
        api_response = api_instance.import_snapshot_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->import_snapshot_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Import Snapshot of an Instance
        api_response = api_instance.import_snapshot_instance(id, import_snapshot_instance_request=import_snapshot_instance_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->import_snapshot_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **import_snapshot_instance_request** | [**ImportSnapshotInstanceRequest**](ImportSnapshotInstanceRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **linked_clone_snapshot_instance**
> Model200Success linked_clone_snapshot_instance(id, snapshot_id)

Create Linked Clone of Instance Snapshot

It is possible to create a linked clone of an Instance Snapshot.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    snapshot_id = 4 # float | Snapshot ID

    # example passing only required values which don't have defaults set
    try:
        # Create Linked Clone of Instance Snapshot
        api_response = api_instance.linked_clone_snapshot_instance(id, snapshot_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->linked_clone_snapshot_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **snapshot_id** | **float**| Snapshot ID |

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

# **list_instance_service_plans**
> ListInstanceServicePlans200Response list_instance_service_plans()

Get Available Service Plans for an Instance

This endpoint retrieves all the Service Plans available for the specified cloud and instance layout. The response includes details about the plans and their configuration options. It may be used to get the list of available plans when creating a new instance or resizing an existing instance. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.list_instance_service_plans200_response import ListInstanceServicePlans200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    layout_id = 98 # int | The Layout ID for Filtering (optional)
    site_id = 7 # int | The Site ID for Filtering (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Available Service Plans for an Instance
        api_response = api_instance.list_instance_service_plans(zone_id=zone_id, layout_id=layout_id, site_id=site_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->list_instance_service_plans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **layout_id** | **int**| The Layout ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]

### Return type

[**ListInstanceServicePlans200Response**](ListInstanceServicePlans200Response.md)

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

# **list_instance_types_provisioning**
> ListInstanceTypesProvisioning200Response list_instance_types_provisioning()

Get All Instance Types for Provisioning

Fetch the list of available instance types. These can vary in range from database containers, to web containers, to custom containers. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.list_instance_types_provisioning200_response import ListInstanceTypesProvisioning200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    featured = False # bool | Filter by featured (optional)
    details = True # bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Instance Types for Provisioning
        api_response = api_instance.list_instance_types_provisioning(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, featured=featured, details=details)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->list_instance_types_provisioning: %s\n" % e)
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
 **featured** | **bool**| Filter by featured | [optional]
 **details** | **bool**| Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | [optional]

### Return type

[**ListInstanceTypesProvisioning200Response**](ListInstanceTypesProvisioning200Response.md)

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

# **list_instances**
> ListInstances200Response list_instances()

Get All Instances

This endpoint retrieves a paginated list of instances. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.list_instances200_response import ListInstances200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    instance_type = "ubuntu" # str | The Instance Type Code for Filtering (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    created_by = 63 # int | The User ID for Filtering (optional)
    agent_installed = True # bool | Filter instances by if agent is installed or not on the associated servers. (optional)
    status = "running" # str | The instance status for filtering. (optional)
    environment = "lab" # str | The environment for filtering. (optional)
    show_deleted = True # bool | If true, includes instances in pending removal status. (optional) if omitted the server will use the default value of False
    deleted = True # bool | If true, only deleted resources or instances in pending removal status are returned. (optional)
    expire_date = "2019-01-01" # str | Filter by expireDate less than or equal to specified date (optional)
    expire_date_min = "2019-01-01" # str | Filter expireDate greater than or equal to the specified date (optional)
    expire_days = "20" # str | Filter by expireDays less than or equal to the specified value (optional)
    expire_days_min = "20" # str | Filter by expireDays greater than or equal to the specified value (optional)
    shutdown_date = "2019-01-01" # str | Filter by shutdownDate less than equal to the specified date (optional)
    shutdown_date_min = "2019-01-01" # str | Filter by shutdownDate greater than or equal to the specified date (optional)
    shutdown_days = "20" # str | Filter by shutdownDays less than or equal to the specified value (optional)
    shutdown_days_min = "20" # str | Filter by shutdownDays greater than or equal to the specified value (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    tags = "tags.env=qa&tags.env=test" # str | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
    metadata = "tags.env=qa&tags.env=test" # str | Alias for tags (optional)
    details = True # bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Instances
        api_response = api_instance.list_instances(max=max, offset=offset, name=name, phrase=phrase, instance_type=instance_type, last_updated=last_updated, created_by=created_by, agent_installed=agent_installed, status=status, environment=environment, show_deleted=show_deleted, deleted=deleted, expire_date=expire_date, expire_date_min=expire_date_min, expire_days=expire_days, expire_days_min=expire_days_min, shutdown_date=shutdown_date, shutdown_date_min=shutdown_date_min, shutdown_days=shutdown_days, shutdown_days_min=shutdown_days_min, labels=labels, all_labels=all_labels, tags=tags, metadata=metadata, details=details)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->list_instances: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **instance_type** | **str**| The Instance Type Code for Filtering | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **created_by** | **int**| The User ID for Filtering | [optional]
 **agent_installed** | **bool**| Filter instances by if agent is installed or not on the associated servers. | [optional]
 **status** | **str**| The instance status for filtering. | [optional]
 **environment** | **str**| The environment for filtering. | [optional]
 **show_deleted** | **bool**| If true, includes instances in pending removal status. | [optional] if omitted the server will use the default value of False
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]
 **expire_date** | **str**| Filter by expireDate less than or equal to specified date | [optional]
 **expire_date_min** | **str**| Filter expireDate greater than or equal to the specified date | [optional]
 **expire_days** | **str**| Filter by expireDays less than or equal to the specified value | [optional]
 **expire_days_min** | **str**| Filter by expireDays greater than or equal to the specified value | [optional]
 **shutdown_date** | **str**| Filter by shutdownDate less than equal to the specified date | [optional]
 **shutdown_date_min** | **str**| Filter by shutdownDate greater than or equal to the specified date | [optional]
 **shutdown_days** | **str**| Filter by shutdownDays less than or equal to the specified value | [optional]
 **shutdown_days_min** | **str**| Filter by shutdownDays greater than or equal to the specified value | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **tags** | **str**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional]
 **metadata** | **str**| Alias for tags | [optional]
 **details** | **bool**| Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [optional] if omitted the server will use the default value of False

### Return type

[**ListInstances200Response**](ListInstances200Response.md)

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

# **list_security_groups_instance**
> ListSecurityGroupsInstance200Response list_security_groups_instance(id)

Get Security Groups for an Instance

This returns a list of all of the security groups applied to an instance and whether the firewall is enabled.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.list_security_groups_instance200_response import ListSecurityGroupsInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Security Groups for an Instance
        api_response = api_instance.list_security_groups_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->list_security_groups_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**ListSecurityGroupsInstance200Response**](ListSecurityGroupsInstance200Response.md)

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

# **lock_instance**
> Model200Success lock_instance(id)

Lock an Instance

This will lock the instance. While locked, instances may not be removed.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Lock an Instance
        api_response = api_instance.lock_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->lock_instance: %s\n" % e)
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

# **refresh_state_instance**
> Model200Success refresh_state_instance(id)

Refresh State of an Instance

This endpoint provides a way to refresh the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Refresh State of an Instance
        api_response = api_instance.refresh_state_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->refresh_state_instance: %s\n" % e)
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

# **remove_instances_from_control**
> SuccessMessage remove_instances_from_control()

Remove From Control

Will delete a brownfield instance (or instances) asynchronously (Only deletes records local to Morpheus, actual VMs remain unchanged).

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.success_message import SuccessMessage
from openapi_client.model.remove_instances_from_control_request import RemoveInstancesFromControlRequest
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
    api_instance = instances_api.InstancesApi(api_client)
    remove_instances_from_control_request = RemoveInstancesFromControlRequest(
        ids=[
            1,
        ],
    ) # RemoveInstancesFromControlRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Remove From Control
        api_response = api_instance.remove_instances_from_control(remove_instances_from_control_request=remove_instances_from_control_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->remove_instances_from_control: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remove_instances_from_control_request** | [**RemoveInstancesFromControlRequest**](RemoveInstancesFromControlRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **resize_instance**
> UpdateInstance200Response resize_instance(id)

Resize an Instance

It is possible to resize containers within an instance by increasing their memory plan or storage limit. This is done by assigning a new service plan to the container. This endpoint also allows for NIC reconfiguration by passing a new array of networkInterfaces

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.update_instance200_response import UpdateInstance200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.instance_resize import InstanceResize
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    instance_resize = InstanceResize(
        instance=InstanceResizeInstance(
            plan=InstanceResizeInstancePlan(
                id=75,
            ),
        ),
        service_plan_options=ServicePlanOptions(
            max_cores=2,
            cores_per_socket=1,
            max_memory=536870912,
        ),
        volumes=[
            InstanceCreateVolume(
                id=1,
                root_volume=True,
                name="data",
                size=5,
                size_id=1,
                storage_type=1,
                datastore_id=InstanceCreateVolumeDatastoreId(None),
            ),
        ],
        delete_original_volumes=False,
        network_interfaces=[
            InstanceCreateNetwork(
                network=InstanceCreateNetworkNetwork(
                    id="id_example",
                ),
                network_interface_type_id=1,
                ip_address="ip_address_example",
                id=1,
            ),
        ],
    ) # InstanceResize |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Resize an Instance
        api_response = api_instance.resize_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->resize_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Resize an Instance
        api_response = api_instance.resize_instance(id, instance_resize=instance_resize)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->resize_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_resize** | [**InstanceResize**](InstanceResize.md)|  | [optional]

### Return type

[**UpdateInstance200Response**](UpdateInstance200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_instance**
> RestartInstance200Response restart_instance(id)

Restart an instance

This will restart all containers running within an instance. This includes rebuilding the environment variables and applying settings to the docker containers.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.restart_instance200_response import RestartInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart an instance
        api_response = api_instance.restart_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->restart_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**RestartInstance200Response**](RestartInstance200Response.md)

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

# **revert_snapshot_instance**
> Model200Success revert_snapshot_instance(id, snapshot_id)

Revert Instance to Snapshot

It is possible to restore an Instance to a snapshot.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    snapshot_id = 4 # float | Snapshot ID

    # example passing only required values which don't have defaults set
    try:
        # Revert Instance to Snapshot
        api_response = api_instance.revert_snapshot_instance(id, snapshot_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->revert_snapshot_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **snapshot_id** | **float**| Snapshot ID |

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

# **run_workflow_instance**
> Model200Success run_workflow_instance(id)

Run Workflow on an Instance

This will run a provisioning workflow on all containers in an instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_workflow import InstanceWorkflow
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    workflow_id = 3 # int | ID of the workflow to execute (optional)
    workflow_name = "myworkflow" # str | Name of the workflow to execute (optional)
    instance_workflow = InstanceWorkflow(
        task_set=InstanceWorkflowTaskSet(
            custom_options={},
        ),
        task_phase="provision",
    ) # InstanceWorkflow |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Run Workflow on an Instance
        api_response = api_instance.run_workflow_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->run_workflow_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Run Workflow on an Instance
        api_response = api_instance.run_workflow_instance(id, workflow_id=workflow_id, workflow_name=workflow_name, instance_workflow=instance_workflow)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->run_workflow_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **workflow_id** | **int**| ID of the workflow to execute | [optional]
 **workflow_name** | **str**| Name of the workflow to execute | [optional]
 **instance_workflow** | [**InstanceWorkflow**](InstanceWorkflow.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_apply_instance**
> Model200Success set_apply_instance(id)

Apply State of an Instance

This endpoint provides a way to apply the state of an instance. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.apply_app_state_request import ApplyAppStateRequest
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    apply_app_state_request = ApplyAppStateRequest(
        template_parameter={},
    ) # ApplyAppStateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Apply State of an Instance
        api_response = api_instance.set_apply_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->set_apply_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Apply State of an Instance
        api_response = api_instance.set_apply_instance(id, apply_app_state_request=apply_app_state_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->set_apply_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **apply_app_state_request** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_instance_security_groups**
> SetInstanceSecurityGroups200Response set_instance_security_groups(id)

Set Security Groups for an Instance

Set Security Groups for an Instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.set_instance_security_groups200_response import SetInstanceSecurityGroups200Response
from openapi_client.model.set_instance_security_groups_request import SetInstanceSecurityGroupsRequest
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    set_instance_security_groups_request = SetInstanceSecurityGroupsRequest(
        security_group_ids=[
            1,
        ],
    ) # SetInstanceSecurityGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set Security Groups for an Instance
        api_response = api_instance.set_instance_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->set_instance_security_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set Security Groups for an Instance
        api_response = api_instance.set_instance_security_groups(id, set_instance_security_groups_request=set_instance_security_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->set_instance_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **set_instance_security_groups_request** | [**SetInstanceSecurityGroupsRequest**](SetInstanceSecurityGroupsRequest.md)|  | [optional]

### Return type

[**SetInstanceSecurityGroups200Response**](SetInstanceSecurityGroups200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **snapshot_instance**
> Model200Success snapshot_instance(id)

Snapshot an Instance

This endpoint will create a snapshot of an instance. This is done asychronously, so the ID of the snapshot is not returned.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_snapshot import InstanceSnapshot
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    instance_snapshot = InstanceSnapshot(
        snapshot=InstanceSnapshotSnapshot(
            name="{serverName}.{timestamp}",
            description="description_example",
        ),
    ) # InstanceSnapshot |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Snapshot an Instance
        api_response = api_instance.snapshot_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->snapshot_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Snapshot an Instance
        api_response = api_instance.snapshot_instance(id, instance_snapshot=instance_snapshot)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->snapshot_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_snapshot** | [**InstanceSnapshot**](InstanceSnapshot.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **snapshots_instance**
> InstanceSnapshots snapshots_instance(id)

Get list of snapshots for an Instance

Get list of snapshots for an Instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_snapshots import InstanceSnapshots
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get list of snapshots for an Instance
        api_response = api_instance.snapshots_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->snapshots_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **start_instance**
> RestartInstance200Response start_instance(id)

Start an instance

This will start all containers running within an instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.restart_instance200_response import RestartInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Start an instance
        api_response = api_instance.start_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->start_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**RestartInstance200Response**](RestartInstance200Response.md)

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

# **stop_instance**
> RestartInstance200Response stop_instance(id)

Stop an instance

This will stop all containers running within an instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.restart_instance200_response import RestartInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Stop an instance
        api_response = api_instance.stop_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->stop_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**RestartInstance200Response**](RestartInstance200Response.md)

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

# **suspend_instance**
> Model200Success suspend_instance(id)

Suspend an instance

This will suspend all containers in the instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Suspend an instance
        api_response = api_instance.suspend_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->suspend_instance: %s\n" % e)
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

# **unlock_instance**
> Model200Success unlock_instance(id)

Unlock an Instance

This will unlock the instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Unlock an Instance
        api_response = api_instance.unlock_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->unlock_instance: %s\n" % e)
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

# **update_instance**
> UpdateInstance200Response update_instance(id)

Updating an Instance

Updating an Instance

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.instance_update import InstanceUpdate
from openapi_client.model.update_instance200_response import UpdateInstance200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    instance_update = InstanceUpdate(
        instance=InstanceUpdateInstance(
            name="myinstance",
            description="my new instance",
            instance_context="dev",
            labels=[
                "labels_example",
            ],
            tags=[
                UpdateHostManagedRequestServerTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            add_tags=[
                UpdateHostManagedRequestServerTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            remove_tags=[
                InstanceUpdateInstanceRemoveTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            power_schedule_type=1,
            site=InstanceUpdateInstanceSite(
                id=2,
            ),
            owner_id=5,
            display_name="myInstanceDisplayName",
        ),
        config=InstancesConfigCustomOptions(
            custom_options={},
        ),
    ) # InstanceUpdate |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating an Instance
        api_response = api_instance.update_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating an Instance
        api_response = api_instance.update_instance(id, instance_update=instance_update)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instance_update** | [**InstanceUpdate**](InstanceUpdate.md)|  | [optional]

### Return type

[**UpdateInstance200Response**](UpdateInstance200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_instance_network_interface**
> UpdateInstanceNetworkInterface200Response update_instance_network_interface(id, network_interface_id)

Updating a label for an Instance's Network

Updating an Instance's Network's Label

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.update_instance_network_interface200_response import UpdateInstanceNetworkInterface200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.network_interface_update import NetworkInterfaceUpdate
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    network_interface_id = 7 # float | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
    network_interface_update = NetworkInterfaceUpdate(
        name="eth0-new",
    ) # NetworkInterfaceUpdate |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a label for an Instance's Network
        api_response = api_instance.update_instance_network_interface(id, network_interface_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance_network_interface: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a label for an Instance's Network
        api_response = api_instance.update_instance_network_interface(id, network_interface_id, network_interface_update=network_interface_update)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance_network_interface: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **network_interface_id** | **float**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced |
 **network_interface_update** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md)|  | [optional]

### Return type

[**UpdateInstanceNetworkInterface200Response**](UpdateInstanceNetworkInterface200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_instance_schedule**
> CreateInstanceSchedule200Response update_instance_schedule(id, schedule_id)

Updating an Instance Schedule

This endpoint provides updating of an instance schedule

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.create_instance_schedule200_response import CreateInstanceSchedule200Response
from openapi_client.model.update_instance_schedule_request import UpdateInstanceScheduleRequest
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 9 # int | Instance ID
    schedule_id = 1 # int | Instance Schedule ID
    update_instance_schedule_request = UpdateInstanceScheduleRequest(
        instance_schedule=InstanceScheduleUpdate(
            schedule_type="dayOfWeek",
            schedule_timezone="UTC",
            start_day_of_week=1,
            start_time="00:00",
            end_day_of_week=7,
            end_time="23:59",
            start_date=dateutil_parser('2022-12-25T00:00:00Z'),
            end_date=dateutil_parser('2023-01-01T00:00:00Z'),
            threshold=InstanceScheduleUpdateThreshold(
                source_threshold_id=1,
                auto_up=True,
                auto_down=True,
                min_count=1,
                max_count=3,
                cpu_enabled=True,
                min_cpu=3.14,
                max_cpu=3.14,
                memory_enabled=True,
                min_memory=3.14,
                max_memory=3.14,
                disk_enabled=True,
                min_disk=3.14,
                max_disk=3.14,
            ),
        ),
    ) # UpdateInstanceScheduleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating an Instance Schedule
        api_response = api_instance.update_instance_schedule(id, schedule_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance_schedule: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating an Instance Schedule
        api_response = api_instance.update_instance_schedule(id, schedule_id, update_instance_schedule_request=update_instance_schedule_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance_schedule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Instance ID |
 **schedule_id** | **int**| Instance Schedule ID |
 **update_instance_schedule_request** | [**UpdateInstanceScheduleRequest**](UpdateInstanceScheduleRequest.md)|  | [optional]

### Return type

[**CreateInstanceSchedule200Response**](CreateInstanceSchedule200Response.md)

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

# **update_instance_threshold**
> UpdateInstanceThreshold200Response update_instance_threshold(id)

Updates an Instance Scale Threshold

Updates the scale threshold settings for a specific instance 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
from openapi_client.model.update_instance_threshold_request import UpdateInstanceThresholdRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_instance_threshold200_response import UpdateInstanceThreshold200Response
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_instance_threshold_request = UpdateInstanceThresholdRequest(
        instance_threshold=UpdateInstanceThresholdRequestInstanceThreshold(
            auto_up=False,
            auto_down=False,
            min_count=1,
            max_count=3,
            cpu_enabled=False,
            min_cpu=0,
            max_cpu=0,
            memory_enabled=False,
            min_memory=0,
            max_memory=0,
            disk_enabled=False,
            min_disk=0,
            max_disk=0,
        ),
    ) # UpdateInstanceThresholdRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates an Instance Scale Threshold
        api_response = api_instance.update_instance_threshold(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance_threshold: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates an Instance Scale Threshold
        api_response = api_instance.update_instance_threshold(id, update_instance_threshold_request=update_instance_threshold_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_instance_threshold: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_instance_threshold_request** | [**UpdateInstanceThresholdRequest**](UpdateInstanceThresholdRequest.md)|  | [optional]

### Return type

[**UpdateInstanceThreshold200Response**](UpdateInstanceThreshold200Response.md)

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

# **update_wiki_instance**
> UpdateWikiApp200Response update_wiki_instance(id)

Update an Instance Wiki Page

Updates an instance Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import instances_api
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
    api_instance = instances_api.InstancesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Instance Wiki Page
        api_response = api_instance.update_wiki_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_wiki_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Instance Wiki Page
        api_response = api_instance.update_wiki_instance(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InstancesApi->update_wiki_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
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

