# openapi_client.EnvironmentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_environments**](EnvironmentsApi.md#add_environments) | **POST** /api/environments | Create a New Environment
[**delete_environments**](EnvironmentsApi.md#delete_environments) | **DELETE** /api/environments/{id} | Delete a Specific Environment
[**get_environments**](EnvironmentsApi.md#get_environments) | **GET** /api/environments/{id} | Get a Specific Environment
[**list_environments**](EnvironmentsApi.md#list_environments) | **GET** /api/environments | List All Environments
[**update_environments**](EnvironmentsApi.md#update_environments) | **PUT** /api/environments/{id} | Update Environment
[**update_environments_active**](EnvironmentsApi.md#update_environments_active) | **PUT** /api/environments/{id}/toggle-active | Toggle Active State of Environment


# **add_environments**
> AddEnvironments200Response add_environments()

Create a New Environment

Create a new environment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import environments_api
from openapi_client.model.add_environments200_response import AddEnvironments200Response
from openapi_client.model.add_environments_request import AddEnvironmentsRequest
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
    api_instance = environments_api.EnvironmentsApi(api_client)
    add_environments_request = AddEnvironmentsRequest(
        environment=AddEnvironmentsRequestEnvironment(
            name="Production",
            code="prod",
            description="Tag used for Production environments",
            visibility="private",
            sort_order=0,
        ),
    ) # AddEnvironmentsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Environment
        api_response = api_instance.add_environments(add_environments_request=add_environments_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->add_environments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_environments_request** | [**AddEnvironmentsRequest**](AddEnvironmentsRequest.md)|  | [optional]

### Return type

[**AddEnvironments200Response**](AddEnvironments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Environment Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_environments**
> SuccessError delete_environments(id)

Delete a Specific Environment

Delete an existing environment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import environments_api
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
    api_instance = environments_api.EnvironmentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Environment
        api_response = api_instance.delete_environments(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->delete_environments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_environments**
> AddEnvironments200ResponseAllOf get_environments(id)

Get a Specific Environment

This endpoint will retrieve a specific environment by id.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import environments_api
from openapi_client.model.add_environments200_response_all_of import AddEnvironments200ResponseAllOf
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
    api_instance = environments_api.EnvironmentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Environment
        api_response = api_instance.get_environments(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->get_environments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddEnvironments200ResponseAllOf**](AddEnvironments200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Environment Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_environments**
> ListEnvironments200Response list_environments()

List All Environments

This endpoint retrieves all environments.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import environments_api
from openapi_client.model.list_environments200_response import ListEnvironments200Response
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
    api_instance = environments_api.EnvironmentsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Environments
        api_response = api_instance.list_environments(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->list_environments: %s\n" % e)
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

### Return type

[**ListEnvironments200Response**](ListEnvironments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Environments |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_environments**
> AddEnvironments200Response update_environments(id)

Update Environment

Update an existing environment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import environments_api
from openapi_client.model.add_environments200_response import AddEnvironments200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_environments_request import UpdateEnvironmentsRequest
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
    api_instance = environments_api.EnvironmentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_environments_request = UpdateEnvironmentsRequest(
        environment=UpdateEnvironmentsRequestEnvironment(
            name="Production",
            description="Tag used for Production environments",
            visibility="private",
            sort_order=0,
            active=True,
        ),
    ) # UpdateEnvironmentsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Environment
        api_response = api_instance.update_environments(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->update_environments: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Environment
        api_response = api_instance.update_environments(id, update_environments_request=update_environments_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->update_environments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_environments_request** | [**UpdateEnvironmentsRequest**](UpdateEnvironmentsRequest.md)|  | [optional]

### Return type

[**AddEnvironments200Response**](AddEnvironments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Environment Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_environments_active**
> AddEnvironments200Response update_environments_active(id)

Toggle Active State of Environment

Toggle Active State of Environment. Default is to toggle the current value.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import environments_api
from openapi_client.model.add_environments200_response import AddEnvironments200Response
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
    api_instance = environments_api.EnvironmentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    active = False # bool | True or False flag for Active (optional)

    # example passing only required values which don't have defaults set
    try:
        # Toggle Active State of Environment
        api_response = api_instance.update_environments_active(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->update_environments_active: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Toggle Active State of Environment
        api_response = api_instance.update_environments_active(id, active=active)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling EnvironmentsApi->update_environments_active: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **active** | **bool**| True or False flag for Active | [optional]

### Return type

[**AddEnvironments200Response**](AddEnvironments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Environment Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

