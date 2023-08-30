# openapi_client.GuidanceApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**execute_guidances**](GuidanceApi.md#execute_guidances) | **PUT** /api/guidance/{id}/execute | Executes a Specific Guidance Recommendation
[**get_guidance_stats**](GuidanceApi.md#get_guidance_stats) | **GET** /api/guidance/stats | Retrieves Guidance Stats
[**get_guidance_types**](GuidanceApi.md#get_guidance_types) | **GET** /api/guidance/types | Retrieves Guidance Types
[**get_guidances**](GuidanceApi.md#get_guidances) | **GET** /api/guidance/{id} | Retrieves a Specific Guidance Recommendation
[**ignore_guidances**](GuidanceApi.md#ignore_guidances) | **PUT** /api/guidance/{id}/ignore | Ignores a Specific Guidance Recommendation
[**list_guidances**](GuidanceApi.md#list_guidances) | **GET** /api/guidance | Retrieves all Guidance Recommendations


# **execute_guidances**
> Model200Success execute_guidances(id)

Executes a Specific Guidance Recommendation

Executes a specific guidance recommendation. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_api
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
    api_instance = guidance_api.GuidanceApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Executes a Specific Guidance Recommendation
        api_response = api_instance.execute_guidances(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceApi->execute_guidances: %s\n" % e)
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

# **get_guidance_stats**
> GetGuidanceStats200Response get_guidance_stats()

Retrieves Guidance Stats

This endpoint retrieves a summary of actionable guidance. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_api
from openapi_client.model.get_guidance_stats200_response import GetGuidanceStats200Response
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
    api_instance = guidance_api.GuidanceApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves Guidance Stats
        api_response = api_instance.get_guidance_stats()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceApi->get_guidance_stats: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetGuidanceStats200Response**](GetGuidanceStats200Response.md)

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

# **get_guidance_types**
> GetGuidanceTypes200Response get_guidance_types()

Retrieves Guidance Types

This endpoint retrieves all guidance types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_api
from openapi_client.model.get_guidance_types200_response import GetGuidanceTypes200Response
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
    api_instance = guidance_api.GuidanceApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves Guidance Types
        api_response = api_instance.get_guidance_types()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceApi->get_guidance_types: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetGuidanceTypes200Response**](GetGuidanceTypes200Response.md)

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

# **get_guidances**
> GetGuidances200Response get_guidances(id)

Retrieves a Specific Guidance Recommendation

Retrieves a specific guidance recommendation. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_api
from openapi_client.model.get_guidances200_response import GetGuidances200Response
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
    api_instance = guidance_api.GuidanceApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Guidance Recommendation
        api_response = api_instance.get_guidances(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceApi->get_guidances: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetGuidances200Response**](GetGuidances200Response.md)

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

# **ignore_guidances**
> Model200Success ignore_guidances(id)

Ignores a Specific Guidance Recommendation

Ignores a specific guidance recommendation. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_api
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
    api_instance = guidance_api.GuidanceApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Ignores a Specific Guidance Recommendation
        api_response = api_instance.ignore_guidances(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceApi->ignore_guidances: %s\n" % e)
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

# **list_guidances**
> ListGuidances200Response list_guidances()

Retrieves all Guidance Recommendations

Retrieves all guidance recommendations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_guidances200_response import ListGuidances200Response
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
    api_instance = guidance_api.GuidanceApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    severity = "INFO" # str | Filter by severity (optional)
    type = "Shutdown" # str | Filter by Guidance Type.  See `Retrieves Guidance Types` for most up to date list of types. (optional)
    state = "any" # str | Filter by state, restricts query to only load discoveries by state (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Guidance Recommendations
        api_response = api_instance.list_guidances(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, severity=severity, type=type, state=state)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceApi->list_guidances: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **severity** | **str**| Filter by severity | [optional]
 **type** | **str**| Filter by Guidance Type.  See &#x60;Retrieves Guidance Types&#x60; for most up to date list of types. | [optional]
 **state** | **str**| Filter by state, restricts query to only load discoveries by state | [optional]

### Return type

[**ListGuidances200Response**](ListGuidances200Response.md)

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

