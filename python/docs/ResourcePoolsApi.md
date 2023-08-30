# openapi_client.ResourcePoolsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_resource_pool_group**](ResourcePoolsApi.md#create_resource_pool_group) | **POST** /api/resource-pools/groups | Create a Resource Pool Group
[**delete_resource_pool_group**](ResourcePoolsApi.md#delete_resource_pool_group) | **DELETE** /api/resource-pools/groups/{id} | Delete a Resource Pool Group
[**get_resource_pool_groups**](ResourcePoolsApi.md#get_resource_pool_groups) | **GET** /api/resource-pools/groups | Get all Resource Pool Groups
[**getresource_pool_group**](ResourcePoolsApi.md#getresource_pool_group) | **GET** /api/resource-pools/groups/{id} | Get a Specific Resource Pool Group
[**update_resource_pool_group**](ResourcePoolsApi.md#update_resource_pool_group) | **PUT** /api/resource-pools/groups/{id} | Update a Resource Pool Group


# **create_resource_pool_group**
> CreateResourcePoolGroup200Response create_resource_pool_group()

Create a Resource Pool Group

Use this command to create a resource pool group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import resource_pools_api
from openapi_client.model.create_resource_pool_group_request import CreateResourcePoolGroupRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_resource_pool_group200_response import CreateResourcePoolGroup200Response
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
    api_instance = resource_pools_api.ResourcePoolsApi(api_client)
    create_resource_pool_group_request = CreateResourcePoolGroupRequest(
        resource_pool_group=ResourcePoolGroupsCreateInput(
            name="name_example",
            description="description_example",
            visibility="visibility_example",
            mode="mode_example",
            pools=[
                1,
            ],
            tenants=[
                ListDeploys200ResponseAllOfAppDeploysInnerInstance(
                    id=1,
                    name="name_example",
                ),
            ],
            resource_permission=GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission(
                all=True,
                sites=[
                    GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner(
                        id=1,
                        default=True,
                    ),
                ],
            ),
        ),
    ) # CreateResourcePoolGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Resource Pool Group
        api_response = api_instance.create_resource_pool_group(create_resource_pool_group_request=create_resource_pool_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ResourcePoolsApi->create_resource_pool_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_resource_pool_group_request** | [**CreateResourcePoolGroupRequest**](CreateResourcePoolGroupRequest.md)|  | [optional]

### Return type

[**CreateResourcePoolGroup200Response**](CreateResourcePoolGroup200Response.md)

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

# **delete_resource_pool_group**
> Model200Success delete_resource_pool_group(id)

Delete a Resource Pool Group

Will delete a Resource Pool Group from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import resource_pools_api
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
    api_instance = resource_pools_api.ResourcePoolsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Resource Pool Group
        api_response = api_instance.delete_resource_pool_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ResourcePoolsApi->delete_resource_pool_group: %s\n" % e)
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

# **get_resource_pool_groups**
> GetResourcePoolGroups200Response get_resource_pool_groups()

Get all Resource Pool Groups

This endpoint retrieves all Resource Pool Groups associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import resource_pools_api
from openapi_client.model.get_resource_pool_groups200_response import GetResourcePoolGroups200Response
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
    api_instance = resource_pools_api.ResourcePoolsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get all Resource Pool Groups
        api_response = api_instance.get_resource_pool_groups()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ResourcePoolsApi->get_resource_pool_groups: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetResourcePoolGroups200Response**](GetResourcePoolGroups200Response.md)

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

# **getresource_pool_group**
> CreateResourcePoolGroup200Response getresource_pool_group(id)

Get a Specific Resource Pool Group

This endpoint retrieves a specific Resource Pool Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import resource_pools_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_resource_pool_group200_response import CreateResourcePoolGroup200Response
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
    api_instance = resource_pools_api.ResourcePoolsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Resource Pool Group
        api_response = api_instance.getresource_pool_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ResourcePoolsApi->getresource_pool_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateResourcePoolGroup200Response**](CreateResourcePoolGroup200Response.md)

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

# **update_resource_pool_group**
> CreateResourcePoolGroup200Response update_resource_pool_group(id)

Update a Resource Pool Group

Use this command to update an existing resource pool Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import resource_pools_api
from openapi_client.model.create_resource_pool_group_request import CreateResourcePoolGroupRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_resource_pool_group200_response import CreateResourcePoolGroup200Response
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
    api_instance = resource_pools_api.ResourcePoolsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_resource_pool_group_request = CreateResourcePoolGroupRequest(
        resource_pool_group=ResourcePoolGroupsCreateInput(
            name="name_example",
            description="description_example",
            visibility="visibility_example",
            mode="mode_example",
            pools=[
                1,
            ],
            tenants=[
                ListDeploys200ResponseAllOfAppDeploysInnerInstance(
                    id=1,
                    name="name_example",
                ),
            ],
            resource_permission=GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission(
                all=True,
                sites=[
                    GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner(
                        id=1,
                        default=True,
                    ),
                ],
            ),
        ),
    ) # CreateResourcePoolGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Resource Pool Group
        api_response = api_instance.update_resource_pool_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ResourcePoolsApi->update_resource_pool_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Resource Pool Group
        api_response = api_instance.update_resource_pool_group(id, create_resource_pool_group_request=create_resource_pool_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ResourcePoolsApi->update_resource_pool_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_resource_pool_group_request** | [**CreateResourcePoolGroupRequest**](CreateResourcePoolGroupRequest.md)|  | [optional]

### Return type

[**CreateResourcePoolGroup200Response**](CreateResourcePoolGroup200Response.md)

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

