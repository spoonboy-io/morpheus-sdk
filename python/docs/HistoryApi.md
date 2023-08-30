# openapi_client.HistoryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_history**](HistoryApi.md#get_history) | **GET** /api/processes/{id} | Retrieves a Specific Process
[**get_history_event**](HistoryApi.md#get_history_event) | **GET** /api/processes/events/{id} | Retrieves a Specific Process Event
[**list_history**](HistoryApi.md#list_history) | **GET** /api/processes | Retrieves Process History


# **get_history**
> GetHistory200Response get_history(id)

Retrieves a Specific Process

Retrieves a specific process. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import history_api
from openapi_client.model.get_history200_response import GetHistory200Response
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
    api_instance = history_api.HistoryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Process
        api_response = api_instance.get_history(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HistoryApi->get_history: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetHistory200Response**](GetHistory200Response.md)

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

# **get_history_event**
> GetHistoryEvent200Response get_history_event(id)

Retrieves a Specific Process Event

Retrieves a specific process event. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import history_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_history_event200_response import GetHistoryEvent200Response
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
    api_instance = history_api.HistoryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Process Event
        api_response = api_instance.get_history_event(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HistoryApi->get_history_event: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetHistoryEvent200Response**](GetHistoryEvent200Response.md)

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

# **list_history**
> ListHistory200Response list_history()

Retrieves Process History

Retrieves process history for objects 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import history_api
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
    api_instance = history_api.HistoryApi(api_client)
    instance_id = 94 # int | The Instance ID for Filtering (optional)
    container_id = 135 # int | The Container ID for Filtering (optional)
    server_id = 97 # int | The Server ID for Filtering (optional)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    app_id = 5 # int | The App ID for Filtering (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Process History
        api_response = api_instance.list_history(instance_id=instance_id, container_id=container_id, server_id=server_id, zone_id=zone_id, app_id=app_id, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HistoryApi->list_history: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **app_id** | **int**| The App ID for Filtering | [optional]
 **phrase** | **str**| Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error | [optional]

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

