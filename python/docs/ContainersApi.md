# openapi_client.ContainersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**clone_image_container_action**](ContainersApi.md#clone_image_container_action) | **PUT** /api/containers/{id}/clone-image | Clone Specific Container to Image
[**containers_attach_floating_ip**](ContainersApi.md#containers_attach_floating_ip) | **PUT** /api/containers/{id}/attach-floating-ip | Attach Floating IP to Container
[**containers_detach_floating_ip**](ContainersApi.md#containers_detach_floating_ip) | **PUT** /api/containers/{id}/detach-floating-ip | Detach Floating IP from Container
[**eject_container_action**](ContainersApi.md#eject_container_action) | **PUT** /api/containers/{id}/eject | Eject a Specific Container
[**execute_container_action**](ContainersApi.md#execute_container_action) | **PUT** /api/containers/{id}/action | Execute Container Action
[**get_container**](ContainersApi.md#get_container) | **GET** /api/containers/{id} | Get a Specific Container
[**get_container_actions**](ContainersApi.md#get_container_actions) | **GET** /api/containers/{id}/actions | List Container Actions
[**import_container_action**](ContainersApi.md#import_container_action) | **PUT** /api/containers/{id}/import | Import a Specific Container
[**restart_container_action**](ContainersApi.md#restart_container_action) | **PUT** /api/containers/{id}/restart | Restart a Specific Container
[**start_container_action**](ContainersApi.md#start_container_action) | **PUT** /api/containers/{id}/start | Start a Specific Container
[**stop_container_action**](ContainersApi.md#stop_container_action) | **PUT** /api/containers/{id}/stop | Stop a Specific Container
[**suspend_container_action**](ContainersApi.md#suspend_container_action) | **PUT** /api/containers/{id}/suspend | Suspend a Specific Container


# **clone_image_container_action**
> SuccessMessage clone_image_container_action(id)

Clone Specific Container to Image

This endpoint clones and converts a Container in its current state to image in the source Cloud and adds Virtual Image Record with metadata matching the source configuration.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    instances_clone_image = InstancesCloneImage(
        template_name="{server.name}-{timestamp}",
        zone_folder="group-v81",
    ) # InstancesCloneImage |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Clone Specific Container to Image
        api_response = api_instance.clone_image_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->clone_image_container_action: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Clone Specific Container to Image
        api_response = api_instance.clone_image_container_action(id, instances_clone_image=instances_clone_image)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->clone_image_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instances_clone_image** | [**InstancesCloneImage**](InstancesCloneImage.md)|  | [optional]

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

# **containers_attach_floating_ip**
> Model200Success containers_attach_floating_ip(id)

Attach Floating IP to Container

This endpoint attaches a floating IP to a container (node/VM).

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.containers_attach_floating_ip_request import ContainersAttachFloatingIpRequest
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    containers_attach_floating_ip_request = ContainersAttachFloatingIpRequest(
        config=ContainersAttachFloatingIpRequestConfig(
            os_external_network_id="ip-42",
            floating_ip_bandwidth=1024,
        ),
    ) # ContainersAttachFloatingIpRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Attach Floating IP to Container
        api_response = api_instance.containers_attach_floating_ip(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->containers_attach_floating_ip: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Attach Floating IP to Container
        api_response = api_instance.containers_attach_floating_ip(id, containers_attach_floating_ip_request=containers_attach_floating_ip_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->containers_attach_floating_ip: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **containers_attach_floating_ip_request** | [**ContainersAttachFloatingIpRequest**](ContainersAttachFloatingIpRequest.md)|  | [optional]

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

# **containers_detach_floating_ip**
> Model200Success containers_detach_floating_ip(id)

Detach Floating IP from Container

This endpoint detaches a floating IP from a container (node/VM).

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Detach Floating IP from Container
        api_response = api_instance.containers_detach_floating_ip(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->containers_detach_floating_ip: %s\n" % e)
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

# **eject_container_action**
> SuccessMessage eject_container_action(id)

Eject a Specific Container

This endpoint ejects attached disk/iso of a specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Eject a Specific Container
        api_response = api_instance.eject_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->eject_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

# **execute_container_action**
> SuccessMessage execute_container_action(id, code)

Execute Container Action

This endpoint executes a container action on specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    code = "apache-remove-node" # str | Action code to be executed on a specific container.

    # example passing only required values which don't have defaults set
    try:
        # Execute Container Action
        api_response = api_instance.execute_container_action(id, code)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->execute_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **code** | **str**| Action code to be executed on a specific container. |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

# **get_container**
> GetContainer200Response get_container(id)

Get a Specific Container

This endpoint retrieves a specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_container200_response import GetContainer200Response
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Container
        api_response = api_instance.get_container(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->get_container: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetContainer200Response**](GetContainer200Response.md)

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

# **get_container_actions**
> GetContainerActions200Response get_container_actions(id)

List Container Actions

This endpoint lists available actions for specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.get_container_actions200_response import GetContainerActions200Response
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # List Container Actions
        api_response = api_instance.get_container_actions(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->get_container_actions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetContainerActions200Response**](GetContainerActions200Response.md)

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

# **import_container_action**
> SuccessMessage import_container_action(id)

Import a Specific Container

This endpoint clones and exports a Container in its current state to target Storage provider and adds Virtual Image Record with metadata matching the source configuration.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.import_container_action_request import ImportContainerActionRequest
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    import_container_action_request = ImportContainerActionRequest(
        storage_provider_id=1,
    ) # ImportContainerActionRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Import a Specific Container
        api_response = api_instance.import_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->import_container_action: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Import a Specific Container
        api_response = api_instance.import_container_action(id, import_container_action_request=import_container_action_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->import_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **import_container_action_request** | [**ImportContainerActionRequest**](ImportContainerActionRequest.md)|  | [optional]

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

# **restart_container_action**
> SuccessMessage restart_container_action(id)

Restart a Specific Container

This endpoint restarts a specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart a Specific Container
        api_response = api_instance.restart_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->restart_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

# **start_container_action**
> SuccessMessage start_container_action(id)

Start a Specific Container

This endpoint starts a specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Start a Specific Container
        api_response = api_instance.start_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->start_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

# **stop_container_action**
> SuccessMessage stop_container_action(id)

Stop a Specific Container

This endpoint stops a specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Stop a Specific Container
        api_response = api_instance.stop_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->stop_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

# **suspend_container_action**
> SuccessMessage suspend_container_action(id)

Suspend a Specific Container

This endpoint suspends a specific container.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import containers_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = containers_api.ContainersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Suspend a Specific Container
        api_response = api_instance.suspend_container_action(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ContainersApi->suspend_container_action: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

