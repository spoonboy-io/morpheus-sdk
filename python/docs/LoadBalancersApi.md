# openapi_client.LoadBalancersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_load_balancer**](LoadBalancersApi.md#create_load_balancer) | **POST** /api/load-balancers | Create a Load Balancer
[**create_load_balancer_monitor**](LoadBalancersApi.md#create_load_balancer_monitor) | **POST** /api/load-balancers/{loadBalancerId}/monitors | Create a Load Balancer Monitor
[**create_load_balancer_pool**](LoadBalancersApi.md#create_load_balancer_pool) | **POST** /api/load-balancers/{loadBalancerId}/pools | Create a Load Balancer Pool
[**create_load_balancer_pool_node**](LoadBalancersApi.md#create_load_balancer_pool_node) | **POST** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Create a Load Balancer Pool Node
[**create_load_balancer_profile**](LoadBalancersApi.md#create_load_balancer_profile) | **POST** /api/load-balancers/{loadBalancerId}/profiles | Create a Load Balancer Profile
[**create_load_balancer_virtual_server**](LoadBalancersApi.md#create_load_balancer_virtual_server) | **POST** /api/load-balancers/{loadBalancerId}/virtual-servers | Create a Load Balancer Virtual Server
[**delete_load_balancer**](LoadBalancersApi.md#delete_load_balancer) | **DELETE** /api/load-balancers/{loadBalancerId} | Delete a Load Balancer
[**delete_load_balancer_monitor**](LoadBalancersApi.md#delete_load_balancer_monitor) | **DELETE** /api/load-balancers/{loadBalancerId}/monitors/{id} | Delete a Load Balancer Monitor
[**delete_load_balancer_pool**](LoadBalancersApi.md#delete_load_balancer_pool) | **DELETE** /api/load-balancers/{loadBalancerId}/pools/{id} | Delete a Load Balancer Pool
[**delete_load_balancer_pool_node**](LoadBalancersApi.md#delete_load_balancer_pool_node) | **DELETE** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Delete a Load Balancer Pool Node
[**delete_load_balancer_profile**](LoadBalancersApi.md#delete_load_balancer_profile) | **DELETE** /api/load-balancers/{loadBalancerId}/profiles/{id} | Delete a Load Balancer Profile
[**delete_load_balancer_virtual_server**](LoadBalancersApi.md#delete_load_balancer_virtual_server) | **DELETE** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Delete a Load Balancer Virtual Server
[**get_load_balancer**](LoadBalancersApi.md#get_load_balancer) | **GET** /api/load-balancers/{loadBalancerId} | Get a Specific Load Balancer
[**get_load_balancer_monitor**](LoadBalancersApi.md#get_load_balancer_monitor) | **GET** /api/load-balancers/{loadBalancerId}/monitors/{id} | Get a Specific Load Balancer Monitor
[**get_load_balancer_pool**](LoadBalancersApi.md#get_load_balancer_pool) | **GET** /api/load-balancers/{loadBalancerId}/pools/{id} | Get a Specific Load Balancer Pool
[**get_load_balancer_pool_node**](LoadBalancersApi.md#get_load_balancer_pool_node) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Get a Specific Load Balancer Pool Node
[**get_load_balancer_profile**](LoadBalancersApi.md#get_load_balancer_profile) | **GET** /api/load-balancers/{loadBalancerId}/profiles/{id} | Get a Specific Load Balancer Profile
[**get_load_balancer_type**](LoadBalancersApi.md#get_load_balancer_type) | **GET** /api/load-balancer-types/{id} | Get a Specific Load Balancer Type
[**get_load_balancer_virtual_server**](LoadBalancersApi.md#get_load_balancer_virtual_server) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Get a Specific Load Balancer Virtual Server
[**list_load_balancer_monitors**](LoadBalancersApi.md#list_load_balancer_monitors) | **GET** /api/load-balancers/{loadBalancerId}/monitors | Get All Load Balancer Monitors For Load Balancer
[**list_load_balancer_pool_nodes**](LoadBalancersApi.md#list_load_balancer_pool_nodes) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Get All Load Balancer Pool Nodes For Load Balancer Pool
[**list_load_balancer_pools**](LoadBalancersApi.md#list_load_balancer_pools) | **GET** /api/load-balancers/{loadBalancerId}/pools | Get All Load Balancer Pools For Load Balancer
[**list_load_balancer_profiles**](LoadBalancersApi.md#list_load_balancer_profiles) | **GET** /api/load-balancers/{loadBalancerId}/profiles | Get All Load Balancer Profiles For Load Balancer
[**list_load_balancer_types**](LoadBalancersApi.md#list_load_balancer_types) | **GET** /api/load-balancer-types | Get All Load Balancer Types
[**list_load_balancer_virtual_servers**](LoadBalancersApi.md#list_load_balancer_virtual_servers) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers | Get All Load Balancer Virtual Servers For Load Balancer
[**list_load_balancers**](LoadBalancersApi.md#list_load_balancers) | **GET** /api/load-balancers | Get All Load Balancers
[**refresh_load_balancer**](LoadBalancersApi.md#refresh_load_balancer) | **PUT** /api/load-balancers/{loadBalancerId}/refresh | Refresh a Load Balancer
[**update_load_balancer**](LoadBalancersApi.md#update_load_balancer) | **PUT** /api/load-balancers/{loadBalancerId} | Update a Load Balancer
[**update_load_balancer_monitor**](LoadBalancersApi.md#update_load_balancer_monitor) | **PUT** /api/load-balancers/{loadBalancerId}/monitors/{id} | Update a Load Balancer Monitor
[**update_load_balancer_pool**](LoadBalancersApi.md#update_load_balancer_pool) | **PUT** /api/load-balancers/{loadBalancerId}/pools/{id} | Update a Load Balancer Pool
[**update_load_balancer_pool_node**](LoadBalancersApi.md#update_load_balancer_pool_node) | **PUT** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Update a Load Balancer Pool Node
[**update_load_balancer_profile**](LoadBalancersApi.md#update_load_balancer_profile) | **PUT** /api/load-balancers/{loadBalancerId}/profiles/{id} | Update a Load Balancer Profile
[**update_load_balancer_virtual_server**](LoadBalancersApi.md#update_load_balancer_virtual_server) | **PUT** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Update a Load Balancer Virtual Server


# **create_load_balancer**
> CreateLoadBalancer200Response create_load_balancer()

Create a Load Balancer

Available for NSX load balancers only  Use this command to create a load balancer. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer200_response import CreateLoadBalancer200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_request import CreateLoadBalancerRequest
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    create_load_balancer_request = CreateLoadBalancerRequest(
        load_balancer=LoadBalancerCreate(
            name="name_example",
            description="description_example",
            network_server_id=1,
            config={},
            visibility="public",
            tenants=[
                GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner(
                    id=1,
                ),
            ],
            resource_permission=LoadBalancerCreateResourcePermission(
                all=True,
                sites=[
                    1,
                ],
            ),
        ),
    ) # CreateLoadBalancerRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Load Balancer
        api_response = api_instance.create_load_balancer(create_load_balancer_request=create_load_balancer_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_load_balancer_request** | [**CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancer200Response**](CreateLoadBalancer200Response.md)

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

# **create_load_balancer_monitor**
> CreateLoadBalancerMonitor200Response create_load_balancer_monitor(load_balancer_id)

Create a Load Balancer Monitor

Use this command to create a load balancer Monitor.  This endpoint allows creating a Load Balancer Monitor. Configuration options vary by Load Balancer Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_monitor200_response import CreateLoadBalancerMonitor200Response
from openapi_client.model.create_load_balancer_monitor_request import CreateLoadBalancerMonitorRequest
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    create_load_balancer_monitor_request = CreateLoadBalancerMonitorRequest(
        load_balancer_monitor=CreateLoadBalancerMonitorRequestLoadBalancerMonitor(
            name="name_example",
            description="description_example",
            monitor_type="monitor_type_example",
            monitor_timeout=1,
            config={},
        ),
    ) # CreateLoadBalancerMonitorRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Load Balancer Monitor
        api_response = api_instance.create_load_balancer_monitor(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_monitor: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Load Balancer Monitor
        api_response = api_instance.create_load_balancer_monitor(load_balancer_id, create_load_balancer_monitor_request=create_load_balancer_monitor_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_monitor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **create_load_balancer_monitor_request** | [**CreateLoadBalancerMonitorRequest**](CreateLoadBalancerMonitorRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerMonitor200Response**](CreateLoadBalancerMonitor200Response.md)

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

# **create_load_balancer_pool**
> CreateLoadBalancerPool200Response create_load_balancer_pool(load_balancer_id)

Create a Load Balancer Pool

Use this command to create a load balancer pool.  This endpoint allows creating a Load Balancer Pool. Configuration options vary by Load Balancer Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_pool_request import CreateLoadBalancerPoolRequest
from openapi_client.model.create_load_balancer_pool200_response import CreateLoadBalancerPool200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    create_load_balancer_pool_request = CreateLoadBalancerPoolRequest(
        load_balancer_pool=CreateLoadBalancerPoolRequestLoadBalancerPool(
            name="name_example",
            description="description_example",
            vip_balance="vip_balance_example",
            min_active=1,
            config={},
        ),
    ) # CreateLoadBalancerPoolRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Load Balancer Pool
        api_response = api_instance.create_load_balancer_pool(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_pool: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Load Balancer Pool
        api_response = api_instance.create_load_balancer_pool(load_balancer_id, create_load_balancer_pool_request=create_load_balancer_pool_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **create_load_balancer_pool_request** | [**CreateLoadBalancerPoolRequest**](CreateLoadBalancerPoolRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerPool200Response**](CreateLoadBalancerPool200Response.md)

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

# **create_load_balancer_pool_node**
> CreateLoadBalancerPoolNode200Response create_load_balancer_pool_node(load_balancer_pool_id)

Create a Load Balancer Pool Node

Use this command to create a load balancer pool node.  This endpoint allows creating a Load Balancer Pool Node. Configuration options vary by Load Balancer Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_pool_node200_response import CreateLoadBalancerPoolNode200Response
from openapi_client.model.create_load_balancer_pool_node_request import CreateLoadBalancerPoolNodeRequest
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_pool_id = 4 # float | Load Balancer Pool ID
    create_load_balancer_pool_node_request = CreateLoadBalancerPoolNodeRequest(
        load_balancer_node=CreateLoadBalancerPoolNodeRequestLoadBalancerNode(
            name="name_example",
            description="description_example",
            ip_address="ip_address_example",
            port=1,
            weight=1,
            config={},
        ),
    ) # CreateLoadBalancerPoolNodeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Load Balancer Pool Node
        api_response = api_instance.create_load_balancer_pool_node(load_balancer_pool_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_pool_node: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Load Balancer Pool Node
        api_response = api_instance.create_load_balancer_pool_node(load_balancer_pool_id, create_load_balancer_pool_node_request=create_load_balancer_pool_node_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_pool_node: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **create_load_balancer_pool_node_request** | [**CreateLoadBalancerPoolNodeRequest**](CreateLoadBalancerPoolNodeRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerPoolNode200Response**](CreateLoadBalancerPoolNode200Response.md)

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

# **create_load_balancer_profile**
> CreateLoadBalancerProfile200Response create_load_balancer_profile(load_balancer_id)

Create a Load Balancer Profile

Use this command to create a load balancer Profile.  This endpoint allows creating a Load Balancer Profile. Configuration options vary by Load Balancer Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_profile_request import CreateLoadBalancerProfileRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_profile200_response import CreateLoadBalancerProfile200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    create_load_balancer_profile_request = CreateLoadBalancerProfileRequest(
        load_balancer_profile=CreateLoadBalancerProfileRequestLoadBalancerProfile(
            name="name_example",
            description="description_example",
            service_type="service_type_example",
            config={},
        ),
    ) # CreateLoadBalancerProfileRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Load Balancer Profile
        api_response = api_instance.create_load_balancer_profile(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_profile: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Load Balancer Profile
        api_response = api_instance.create_load_balancer_profile(load_balancer_id, create_load_balancer_profile_request=create_load_balancer_profile_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_profile: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **create_load_balancer_profile_request** | [**CreateLoadBalancerProfileRequest**](CreateLoadBalancerProfileRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerProfile200Response**](CreateLoadBalancerProfile200Response.md)

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

# **create_load_balancer_virtual_server**
> CreateLoadBalancerVirtualServer200Response create_load_balancer_virtual_server(load_balancer_id)

Create a Load Balancer Virtual Server

Use this command to create a load balancer virtual server.  This endpoint allows creating a Load Balancer Virtual Server. Configuration options vary by Load Balancer Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_virtual_server_request import CreateLoadBalancerVirtualServerRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_virtual_server200_response import CreateLoadBalancerVirtualServer200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    create_load_balancer_virtual_server_request = CreateLoadBalancerVirtualServerRequest(
        load_balancer_instance=CreateLoadBalancerVirtualServerRequestLoadBalancerInstance(
            vip_name="vip_name_example",
            description="description_example",
            vip_address="vip_address_example",
            vip_port="vip_port_example",
            vip_protocol="vip_protocol_example",
            vip_hostname="vip_hostname_example",
            ssl_cert=1,
            ssl_server_cert=1,
            config={},
        ),
    ) # CreateLoadBalancerVirtualServerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Load Balancer Virtual Server
        api_response = api_instance.create_load_balancer_virtual_server(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_virtual_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Load Balancer Virtual Server
        api_response = api_instance.create_load_balancer_virtual_server(load_balancer_id, create_load_balancer_virtual_server_request=create_load_balancer_virtual_server_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->create_load_balancer_virtual_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **create_load_balancer_virtual_server_request** | [**CreateLoadBalancerVirtualServerRequest**](CreateLoadBalancerVirtualServerRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerVirtualServer200Response**](CreateLoadBalancerVirtualServer200Response.md)

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

# **delete_load_balancer**
> Model200Success delete_load_balancer(load_balancer_id)

Delete a Load Balancer

Will delete a Load Balancer from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Load Balancer
        api_response = api_instance.delete_load_balancer(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->delete_load_balancer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |

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

# **delete_load_balancer_monitor**
> Model200Success delete_load_balancer_monitor(load_balancer_id, id)

Delete a Load Balancer Monitor

Will delete a Load Balancer Monitor from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Load Balancer Monitor
        api_response = api_instance.delete_load_balancer_monitor(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->delete_load_balancer_monitor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

# **delete_load_balancer_pool**
> Model200Success delete_load_balancer_pool(load_balancer_id, id)

Delete a Load Balancer Pool

Will delete a Load Balancer Pool from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Load Balancer Pool
        api_response = api_instance.delete_load_balancer_pool(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->delete_load_balancer_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

# **delete_load_balancer_pool_node**
> Model200Success delete_load_balancer_pool_node(load_balancer_pool_id, id)

Delete a Load Balancer Pool Node

Will delete a Load Balancer Pool Node from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_pool_id = 4 # float | Load Balancer Pool ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Load Balancer Pool Node
        api_response = api_instance.delete_load_balancer_pool_node(load_balancer_pool_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->delete_load_balancer_pool_node: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
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

# **delete_load_balancer_profile**
> Model200Success delete_load_balancer_profile(load_balancer_id, id)

Delete a Load Balancer Profile

Will delete a Load Balancer Profile from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Load Balancer Profile
        api_response = api_instance.delete_load_balancer_profile(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->delete_load_balancer_profile: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

# **delete_load_balancer_virtual_server**
> Model200Success delete_load_balancer_virtual_server(load_balancer_id, id)

Delete a Load Balancer Virtual Server

Will delete a Load Balancer Virtual Server from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Load Balancer Virtual Server
        api_response = api_instance.delete_load_balancer_virtual_server(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->delete_load_balancer_virtual_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

# **get_load_balancer**
> CreateLoadBalancer200Response get_load_balancer(load_balancer_id)

Get a Specific Load Balancer

This endpoint retrieves a specific Load Balancer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer200_response import CreateLoadBalancer200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer
        api_response = api_instance.get_load_balancer(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |

### Return type

[**CreateLoadBalancer200Response**](CreateLoadBalancer200Response.md)

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

# **get_load_balancer_monitor**
> CreateLoadBalancerMonitor200ResponseAllOf get_load_balancer_monitor(load_balancer_id, id)

Get a Specific Load Balancer Monitor

This endpoint retrieves a specific Load Balancer Monitor.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_monitor200_response_all_of import CreateLoadBalancerMonitor200ResponseAllOf
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer Monitor
        api_response = api_instance.get_load_balancer_monitor(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer_monitor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateLoadBalancerMonitor200ResponseAllOf**](CreateLoadBalancerMonitor200ResponseAllOf.md)

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

# **get_load_balancer_pool**
> CreateLoadBalancerPool200ResponseAllOf get_load_balancer_pool(load_balancer_id, id)

Get a Specific Load Balancer Pool

This endpoint retrieves a specific Load Balancer Pool.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_pool200_response_all_of import CreateLoadBalancerPool200ResponseAllOf
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer Pool
        api_response = api_instance.get_load_balancer_pool(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateLoadBalancerPool200ResponseAllOf**](CreateLoadBalancerPool200ResponseAllOf.md)

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

# **get_load_balancer_pool_node**
> CreateLoadBalancerPoolNode200ResponseAllOf get_load_balancer_pool_node(load_balancer_pool_id, id)

Get a Specific Load Balancer Pool Node

This endpoint retrieves a specific Load Balancer Pool Node.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_pool_node200_response_all_of import CreateLoadBalancerPoolNode200ResponseAllOf
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_pool_id = 4 # float | Load Balancer Pool ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer Pool Node
        api_response = api_instance.get_load_balancer_pool_node(load_balancer_pool_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer_pool_node: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateLoadBalancerPoolNode200ResponseAllOf**](CreateLoadBalancerPoolNode200ResponseAllOf.md)

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

# **get_load_balancer_profile**
> CreateLoadBalancerProfile200ResponseAllOf get_load_balancer_profile(load_balancer_id, id)

Get a Specific Load Balancer Profile

This endpoint retrieves a specific Load Balancer Profile.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_profile200_response_all_of import CreateLoadBalancerProfile200ResponseAllOf
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer Profile
        api_response = api_instance.get_load_balancer_profile(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer_profile: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateLoadBalancerProfile200ResponseAllOf**](CreateLoadBalancerProfile200ResponseAllOf.md)

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

# **get_load_balancer_type**
> GetLoadBalancerType200Response get_load_balancer_type(id)

Get a Specific Load Balancer Type

This endpoint will retrieve a specific load balancer type by id.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.get_load_balancer_type200_response import GetLoadBalancerType200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer Type
        api_response = api_instance.get_load_balancer_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetLoadBalancerType200Response**](GetLoadBalancerType200Response.md)

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

# **get_load_balancer_virtual_server**
> CreateLoadBalancerVirtualServer200Response get_load_balancer_virtual_server(load_balancer_id, id)

Get a Specific Load Balancer Virtual Server

This endpoint retrieves a specific Load Balancer Virtual Server.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_virtual_server200_response import CreateLoadBalancerVirtualServer200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Load Balancer Virtual Server
        api_response = api_instance.get_load_balancer_virtual_server(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->get_load_balancer_virtual_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateLoadBalancerVirtualServer200Response**](CreateLoadBalancerVirtualServer200Response.md)

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

# **list_load_balancer_monitors**
> ListLoadBalancerMonitors200Response list_load_balancer_monitors(load_balancer_id)

Get All Load Balancer Monitors For Load Balancer

This endpoint retrieves all load balancer monitors associated with a specified load balancer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.list_load_balancer_monitors200_response import ListLoadBalancerMonitors200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Load Balancer Monitors For Load Balancer
        api_response = api_instance.list_load_balancer_monitors(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_monitors: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancer Monitors For Load Balancer
        api_response = api_instance.list_load_balancer_monitors(load_balancer_id, max=max, offset=offset, sort=sort, direction=direction, name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_monitors: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListLoadBalancerMonitors200Response**](ListLoadBalancerMonitors200Response.md)

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

# **list_load_balancer_pool_nodes**
> ListLoadBalancerPoolNodes200Response list_load_balancer_pool_nodes(load_balancer_pool_id)

Get All Load Balancer Pool Nodes For Load Balancer Pool

This endpoint retrieves all load balancer pool nodes associated with a specified load balancer pool.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.list_load_balancer_pool_nodes200_response import ListLoadBalancerPoolNodes200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_pool_id = 4 # float | Load Balancer Pool ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Load Balancer Pool Nodes For Load Balancer Pool
        api_response = api_instance.list_load_balancer_pool_nodes(load_balancer_pool_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_pool_nodes: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancer Pool Nodes For Load Balancer Pool
        api_response = api_instance.list_load_balancer_pool_nodes(load_balancer_pool_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_pool_nodes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListLoadBalancerPoolNodes200Response**](ListLoadBalancerPoolNodes200Response.md)

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

# **list_load_balancer_pools**
> ListLoadBalancerPools200Response list_load_balancer_pools(load_balancer_id)

Get All Load Balancer Pools For Load Balancer

This endpoint retrieves all load balancer pools associated with a specified load balancer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.list_load_balancer_pools200_response import ListLoadBalancerPools200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Load Balancer Pools For Load Balancer
        api_response = api_instance.list_load_balancer_pools(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_pools: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancer Pools For Load Balancer
        api_response = api_instance.list_load_balancer_pools(load_balancer_id, max=max, offset=offset, sort=sort, direction=direction, name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListLoadBalancerPools200Response**](ListLoadBalancerPools200Response.md)

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

# **list_load_balancer_profiles**
> ListLoadBalancerProfiles200Response list_load_balancer_profiles(load_balancer_id)

Get All Load Balancer Profiles For Load Balancer

This endpoint retrieves all load balancer profiles associated with a specified load balancer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.list_load_balancer_profiles200_response import ListLoadBalancerProfiles200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Load Balancer Profiles For Load Balancer
        api_response = api_instance.list_load_balancer_profiles(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_profiles: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancer Profiles For Load Balancer
        api_response = api_instance.list_load_balancer_profiles(load_balancer_id, max=max, offset=offset, sort=sort, direction=direction, name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_profiles: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListLoadBalancerProfiles200Response**](ListLoadBalancerProfiles200Response.md)

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

# **list_load_balancer_types**
> ListLoadBalancerTypes200Response list_load_balancer_types()

Get All Load Balancer Types

This endpoint retrieves all Load Balancer Types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.list_load_balancer_types200_response import ListLoadBalancerTypes200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    option_types = False # bool | Pass true to include optionTypes in the response for each entry. (optional) if omitted the server will use the default value of False
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancer Types
        api_response = api_instance.list_load_balancer_types(max=max, offset=offset, sort=sort, direction=direction, option_types=option_types, phrase=phrase, name=name, code=code)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **option_types** | **bool**| Pass true to include optionTypes in the response for each entry. | [optional] if omitted the server will use the default value of False
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]

### Return type

[**ListLoadBalancerTypes200Response**](ListLoadBalancerTypes200Response.md)

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

# **list_load_balancer_virtual_servers**
> ListLoadBalancerVirtualServers200Response list_load_balancer_virtual_servers(load_balancer_id)

Get All Load Balancer Virtual Servers For Load Balancer

This endpoint retrieves load balancer virtual servers associated with a specified load balancer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_load_balancer_virtual_servers200_response import ListLoadBalancerVirtualServers200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    vip_name = "lb-114" # str | If specified will return an exact match on vipName (optional)
    vip_address = "192.168.123.44" # str | If specified will return an exact match on vipAddress (optional)
    vip_hostname = "mylb" # str | If specified will return an exact match on vipHostname (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Load Balancer Virtual Servers For Load Balancer
        api_response = api_instance.list_load_balancer_virtual_servers(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_virtual_servers: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancer Virtual Servers For Load Balancer
        api_response = api_instance.list_load_balancer_virtual_servers(load_balancer_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, vip_name=vip_name, vip_address=vip_address, vip_hostname=vip_hostname)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancer_virtual_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **vip_name** | **str**| If specified will return an exact match on vipName | [optional]
 **vip_address** | **str**| If specified will return an exact match on vipAddress | [optional]
 **vip_hostname** | **str**| If specified will return an exact match on vipHostname | [optional]

### Return type

[**ListLoadBalancerVirtualServers200Response**](ListLoadBalancerVirtualServers200Response.md)

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

# **list_load_balancers**
> ListLoadBalancers200Response list_load_balancers()

Get All Load Balancers

This endpoint retrieves all load balancers associated with the account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.list_load_balancers200_response import ListLoadBalancers200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Load Balancers
        api_response = api_instance.list_load_balancers(max=max, offset=offset, sort=sort, direction=direction, name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->list_load_balancers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListLoadBalancers200Response**](ListLoadBalancers200Response.md)

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

# **refresh_load_balancer**
> RefreshLoadBalancer200Response refresh_load_balancer(load_balancer_id)

Refresh a Load Balancer

Will refresh a Load Balancer.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.refresh_load_balancer200_response import RefreshLoadBalancer200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID

    # example passing only required values which don't have defaults set
    try:
        # Refresh a Load Balancer
        api_response = api_instance.refresh_load_balancer(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->refresh_load_balancer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |

### Return type

[**RefreshLoadBalancer200Response**](RefreshLoadBalancer200Response.md)

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

# **update_load_balancer**
> CreateLoadBalancer200Response update_load_balancer(load_balancer_id)

Update a Load Balancer

Available for NSX load balancers only  Use this command to update an existing load balancer. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer200_response import CreateLoadBalancer200Response
from openapi_client.model.update_load_balancer_request import UpdateLoadBalancerRequest
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    update_load_balancer_request = UpdateLoadBalancerRequest(
        load_balancer=LoadBalancerUpdate(
            name="name_example",
            description="description_example",
            enabled=True,
            config={},
            visibility="public",
            tenants=[
                GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner(
                    id=1,
                ),
            ],
            resource_permission=LoadBalancerCreateResourcePermission(
                all=True,
                sites=[
                    1,
                ],
            ),
        ),
    ) # UpdateLoadBalancerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Load Balancer
        api_response = api_instance.update_load_balancer(load_balancer_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Load Balancer
        api_response = api_instance.update_load_balancer(load_balancer_id, update_load_balancer_request=update_load_balancer_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **update_load_balancer_request** | [**UpdateLoadBalancerRequest**](UpdateLoadBalancerRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancer200Response**](CreateLoadBalancer200Response.md)

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

# **update_load_balancer_monitor**
> CreateLoadBalancerMonitor200Response update_load_balancer_monitor(load_balancer_id, id)

Update a Load Balancer Monitor

Use this command to update an existing load balancer monitor.  This endpoint allows updating a Load Balancer Monitor. Configuration options vary by Load Balancer Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_monitor200_response import CreateLoadBalancerMonitor200Response
from openapi_client.model.create_load_balancer_monitor_request import CreateLoadBalancerMonitorRequest
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced
    create_load_balancer_monitor_request = CreateLoadBalancerMonitorRequest(
        load_balancer_monitor=CreateLoadBalancerMonitorRequestLoadBalancerMonitor(
            name="name_example",
            description="description_example",
            monitor_type="monitor_type_example",
            monitor_timeout=1,
            config={},
        ),
    ) # CreateLoadBalancerMonitorRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Load Balancer Monitor
        api_response = api_instance.update_load_balancer_monitor(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_monitor: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Load Balancer Monitor
        api_response = api_instance.update_load_balancer_monitor(load_balancer_id, id, create_load_balancer_monitor_request=create_load_balancer_monitor_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_monitor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_load_balancer_monitor_request** | [**CreateLoadBalancerMonitorRequest**](CreateLoadBalancerMonitorRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerMonitor200Response**](CreateLoadBalancerMonitor200Response.md)

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

# **update_load_balancer_pool**
> CreateLoadBalancerPool200Response update_load_balancer_pool(load_balancer_id, id)

Update a Load Balancer Pool

Use this command to update an existing load balancer pool.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_pool_request import CreateLoadBalancerPoolRequest
from openapi_client.model.create_load_balancer_pool200_response import CreateLoadBalancerPool200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced
    create_load_balancer_pool_request = CreateLoadBalancerPoolRequest(
        load_balancer_pool=CreateLoadBalancerPoolRequestLoadBalancerPool(
            name="name_example",
            description="description_example",
            vip_balance="vip_balance_example",
            min_active=1,
            config={},
        ),
    ) # CreateLoadBalancerPoolRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Load Balancer Pool
        api_response = api_instance.update_load_balancer_pool(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_pool: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Load Balancer Pool
        api_response = api_instance.update_load_balancer_pool(load_balancer_id, id, create_load_balancer_pool_request=create_load_balancer_pool_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_load_balancer_pool_request** | [**CreateLoadBalancerPoolRequest**](CreateLoadBalancerPoolRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerPool200Response**](CreateLoadBalancerPool200Response.md)

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

# **update_load_balancer_pool_node**
> CreateLoadBalancerPoolNode200Response update_load_balancer_pool_node(load_balancer_pool_id, id)

Update a Load Balancer Pool Node

Use this command to update an existing load balancer pool node.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_pool_node200_response import CreateLoadBalancerPoolNode200Response
from openapi_client.model.create_load_balancer_pool_node_request import CreateLoadBalancerPoolNodeRequest
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_pool_id = 4 # float | Load Balancer Pool ID
    id = 1 # int | Morpheus ID of the Object being referenced
    create_load_balancer_pool_node_request = CreateLoadBalancerPoolNodeRequest(
        load_balancer_node=CreateLoadBalancerPoolNodeRequestLoadBalancerNode(
            name="name_example",
            description="description_example",
            ip_address="ip_address_example",
            port=1,
            weight=1,
            config={},
        ),
    ) # CreateLoadBalancerPoolNodeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Load Balancer Pool Node
        api_response = api_instance.update_load_balancer_pool_node(load_balancer_pool_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_pool_node: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Load Balancer Pool Node
        api_response = api_instance.update_load_balancer_pool_node(load_balancer_pool_id, id, create_load_balancer_pool_node_request=create_load_balancer_pool_node_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_pool_node: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_load_balancer_pool_node_request** | [**CreateLoadBalancerPoolNodeRequest**](CreateLoadBalancerPoolNodeRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerPoolNode200Response**](CreateLoadBalancerPoolNode200Response.md)

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

# **update_load_balancer_profile**
> CreateLoadBalancerProfile200Response update_load_balancer_profile(load_balancer_id, id)

Update a Load Balancer Profile

Use this command to update an existing load balancer Profile.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.create_load_balancer_profile_request import CreateLoadBalancerProfileRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_profile200_response import CreateLoadBalancerProfile200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced
    create_load_balancer_profile_request = CreateLoadBalancerProfileRequest(
        load_balancer_profile=CreateLoadBalancerProfileRequestLoadBalancerProfile(
            name="name_example",
            description="description_example",
            service_type="service_type_example",
            config={},
        ),
    ) # CreateLoadBalancerProfileRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Load Balancer Profile
        api_response = api_instance.update_load_balancer_profile(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_profile: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Load Balancer Profile
        api_response = api_instance.update_load_balancer_profile(load_balancer_id, id, create_load_balancer_profile_request=create_load_balancer_profile_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_profile: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_load_balancer_profile_request** | [**CreateLoadBalancerProfileRequest**](CreateLoadBalancerProfileRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerProfile200Response**](CreateLoadBalancerProfile200Response.md)

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

# **update_load_balancer_virtual_server**
> CreateLoadBalancerVirtualServer200Response update_load_balancer_virtual_server(load_balancer_id, id)

Update a Load Balancer Virtual Server

Use this command to update an existing load balancer virtual server.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import load_balancers_api
from openapi_client.model.update_load_balancer_virtual_server_request import UpdateLoadBalancerVirtualServerRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_load_balancer_virtual_server200_response import CreateLoadBalancerVirtualServer200Response
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
    api_instance = load_balancers_api.LoadBalancersApi(api_client)
    load_balancer_id = 4 # float | Load Balancer ID
    id = 1 # int | Morpheus ID of the Object being referenced
    update_load_balancer_virtual_server_request = UpdateLoadBalancerVirtualServerRequest(
        load_balancer_instance=LoadBalancerInstanceUpdate(
            vip_name="vip_name_example",
            description="description_example",
            vip_address="vip_address_example",
            vip_port="vip_port_example",
            vip_protocol="vip_protocol_example",
            vip_hostname="vip_hostname_example",
            pool=1,
            ssl_cert=1,
            ssl_server_cert=1,
            config={},
        ),
    ) # UpdateLoadBalancerVirtualServerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Load Balancer Virtual Server
        api_response = api_instance.update_load_balancer_virtual_server(load_balancer_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_virtual_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Load Balancer Virtual Server
        api_response = api_instance.update_load_balancer_virtual_server(load_balancer_id, id, update_load_balancer_virtual_server_request=update_load_balancer_virtual_server_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LoadBalancersApi->update_load_balancer_virtual_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_load_balancer_virtual_server_request** | [**UpdateLoadBalancerVirtualServerRequest**](UpdateLoadBalancerVirtualServerRequest.md)|  | [optional]

### Return type

[**CreateLoadBalancerVirtualServer200Response**](CreateLoadBalancerVirtualServer200Response.md)

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

