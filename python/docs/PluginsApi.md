# openapi_client.PluginsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_plugin**](PluginsApi.md#get_plugin) | **GET** /api/plugins/{id} | Retrieves a Specific Plugin
[**list_plugins**](PluginsApi.md#list_plugins) | **GET** /api/plugins | Retrieves all Plugins
[**remove_plugin**](PluginsApi.md#remove_plugin) | **DELETE** /api/plugins/{id} | Deletes a Plugin
[**update_plugin**](PluginsApi.md#update_plugin) | **PUT** /api/plugins/{id} | Updates a Plugin
[**upload_plugin**](PluginsApi.md#upload_plugin) | **POST** /api/plugins/upload | Upload Plugin


# **get_plugin**
> GetPlugin200Response get_plugin(id)

Retrieves a Specific Plugin

Retrieves a specific plugin. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import plugins_api
from openapi_client.model.get_plugin200_response import GetPlugin200Response
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
    api_instance = plugins_api.PluginsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Plugin
        api_response = api_instance.get_plugin(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PluginsApi->get_plugin: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPlugin200Response**](GetPlugin200Response.md)

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

# **list_plugins**
> ListPlugins200Response list_plugins()

Retrieves all Plugins

Retrieves all plugins. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import plugins_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_plugins200_response import ListPlugins200Response
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
    api_instance = plugins_api.PluginsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Plugins
        api_response = api_instance.list_plugins(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PluginsApi->list_plugins: %s\n" % e)
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

### Return type

[**ListPlugins200Response**](ListPlugins200Response.md)

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

# **remove_plugin**
> Model200Success remove_plugin(id)

Deletes a Plugin

Deletes a specified plugin. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import plugins_api
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
    api_instance = plugins_api.PluginsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Plugin
        api_response = api_instance.remove_plugin(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PluginsApi->remove_plugin: %s\n" % e)
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

# **update_plugin**
> UploadPlugin200Response update_plugin(id)

Updates a Plugin

Updates a plugin settings. See Upload Plugin for installing a new version of a plugin. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import plugins_api
from openapi_client.model.upload_plugin200_response import UploadPlugin200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_plugin_request import UpdatePluginRequest
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
    api_instance = plugins_api.PluginsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_plugin_request = UpdatePluginRequest(
        plugin=UpdatePluginRequestPlugin(
            enabled=True,
            config={},
        ),
    ) # UpdatePluginRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Plugin
        api_response = api_instance.update_plugin(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PluginsApi->update_plugin: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Plugin
        api_response = api_instance.update_plugin(id, update_plugin_request=update_plugin_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PluginsApi->update_plugin: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_plugin_request** | [**UpdatePluginRequest**](UpdatePluginRequest.md)|  | [optional]

### Return type

[**UploadPlugin200Response**](UploadPlugin200Response.md)

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

# **upload_plugin**
> UploadPlugin200Response upload_plugin()

Upload Plugin

Upload a plugin .jar file to install a new plugin or update an existing plugin to a new version. This endpoint parses the plugin file and starts the asynchronous process of installing and reloading the plugin. The plugin status will be `installing` or `updating` until the reload is complete and then the status changes to `loaded` or `error` if the installation fails. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import plugins_api
from openapi_client.model.upload_plugin200_response import UploadPlugin200Response
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
    api_instance = plugins_api.PluginsApi(api_client)
    file = open('/path/to/file', 'rb') # file_type | Plugin .jar file contents (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Upload Plugin
        api_response = api_instance.upload_plugin(file=file)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PluginsApi->upload_plugin: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **file_type**| Plugin .jar file contents | [optional]

### Return type

[**UploadPlugin200Response**](UploadPlugin200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

