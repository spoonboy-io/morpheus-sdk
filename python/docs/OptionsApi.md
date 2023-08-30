# openapi_client.OptionsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_option_source_data**](OptionsApi.md#get_option_source_data) | **GET** /api/options/{optionSource} | Get Option Source Data
[**list_code_repositories**](OptionsApi.md#list_code_repositories) | **GET** /api/options/codeRepositories | Retrieves a list of Code/GIT Repositories
[**list_option_network_options**](OptionsApi.md#list_option_network_options) | **GET** /api/options/zoneNetworkOptions | Retrieves network options by zone/cloud
[**list_option_values**](OptionsApi.md#list_option_values) | **GET** /api/options/list | Retrieves input option values


# **get_option_source_data**
> GetOptionSourceData200Response get_option_source_data(option_source)

Get Option Source Data

Returns a list of name/value pairs for option-type models. Some option-types depend on input data for proper representation. This typically includes zoneId or siteId for the item being provisioned as request parameters or sometimes previous option type parameters. Each option returned has a `value`, which is often the `id`, but may be a `code` or other attribute. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import options_api
from openapi_client.model.get_option_source_data200_response import GetOptionSourceData200Response
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
    api_instance = options_api.OptionsApi(api_client)
    option_source = "keyPairs" # str | `optionSource` to be listed

    # example passing only required values which don't have defaults set
    try:
        # Get Option Source Data
        api_response = api_instance.get_option_source_data(option_source)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OptionsApi->get_option_source_data: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **option_source** | **str**| &#x60;optionSource&#x60; to be listed |

### Return type

[**GetOptionSourceData200Response**](GetOptionSourceData200Response.md)

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

# **list_code_repositories**
> ListCodeRepositories200Response list_code_repositories()

Retrieves a list of Code/GIT Repositories

Retrieves a list of Code/GIT Repositories 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import options_api
from openapi_client.model.list_code_repositories200_response import ListCodeRepositories200Response
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
    api_instance = options_api.OptionsApi(api_client)
    integration_id = 33 # int | Filter by an integration Id. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves a list of Code/GIT Repositories
        api_response = api_instance.list_code_repositories(integration_id=integration_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OptionsApi->list_code_repositories: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integration_id** | **int**| Filter by an integration Id. | [optional]

### Return type

[**ListCodeRepositories200Response**](ListCodeRepositories200Response.md)

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

# **list_option_network_options**
> ListOptionNetworkOptions200Response list_option_network_options()

Retrieves network options by zone/cloud

This endpoint can be used to see which network options are available for a given cloud (zoneId) and provision type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import options_api
from openapi_client.model.list_option_network_options200_response import ListOptionNetworkOptions200Response
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
    api_instance = options_api.OptionsApi(api_client)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    provision_type_id = 22 # int | Provision type filter, restricts query to only load service plans of specified provision type (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves network options by zone/cloud
        api_response = api_instance.list_option_network_options(zone_id=zone_id, provision_type_id=provision_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OptionsApi->list_option_network_options: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **provision_type_id** | **int**| Provision type filter, restricts query to only load service plans of specified provision type | [optional]

### Return type

[**ListOptionNetworkOptions200Response**](ListOptionNetworkOptions200Response.md)

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

# **list_option_values**
> ListOptionValues200Response list_option_values(option_type_id)

Retrieves input option values

Retrieves all input option values.  Can be used with parameters to supply dependent input values. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import options_api
from openapi_client.model.list_option_values200_response import ListOptionValues200Response
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
    api_instance = options_api.OptionsApi(api_client)
    option_type_id = 1 # int | Input or Option Type ID
    config = {} # {str: (bool, date, datetime, dict, float, int, list, str, none_type)} | Input parameters are required if the input is dependent on them.  Fields must be prefixed with `config.` (optional)

    # example passing only required values which don't have defaults set
    try:
        # Retrieves input option values
        api_response = api_instance.list_option_values(option_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OptionsApi->list_option_values: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves input option values
        api_response = api_instance.list_option_values(option_type_id, config=config)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OptionsApi->list_option_values: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **option_type_id** | **int**| Input or Option Type ID |
 **config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}**| Input parameters are required if the input is dependent on them.  Fields must be prefixed with &#x60;config.&#x60; | [optional]

### Return type

[**ListOptionValues200Response**](ListOptionValues200Response.md)

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

