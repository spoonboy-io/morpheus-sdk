# openapi_client.ImageBuildsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_boot_script**](ImageBuildsApi.md#add_boot_script) | **POST** /api/boot-scripts | Create a Boot Script
[**add_image_build**](ImageBuildsApi.md#add_image_build) | **POST** /api/image-builds | Create an Image Build
[**add_preseed_script**](ImageBuildsApi.md#add_preseed_script) | **POST** /api/preseed-scripts | Create a Preseed Script
[**delete_boot_script**](ImageBuildsApi.md#delete_boot_script) | **DELETE** /api/boot-scripts/{id} | Delete a Boot Script
[**delete_image_build**](ImageBuildsApi.md#delete_image_build) | **DELETE** /api/image-builds/{id} | Delete an Image Build
[**delete_preseed_script**](ImageBuildsApi.md#delete_preseed_script) | **DELETE** /api/preseed-scripts/{id} | Delete a Preseed Script
[**execute_image_build**](ImageBuildsApi.md#execute_image_build) | **POST** /api/image-builds/{id}/run | Run an Image Build
[**get_boot_script**](ImageBuildsApi.md#get_boot_script) | **GET** /api/boot-scripts/{id} | Get a Specific Boot Script
[**get_image_build**](ImageBuildsApi.md#get_image_build) | **GET** /api/image-builds/{id} | Get a Specific Image Build
[**get_image_build_executions**](ImageBuildsApi.md#get_image_build_executions) | **GET** /api/image-builds/{id}/list-executions | List Image Build Executions
[**get_preseed_script**](ImageBuildsApi.md#get_preseed_script) | **GET** /api/preseed-scripts/{id} | Get a Specific Preseed Script
[**list_boot_scripts**](ImageBuildsApi.md#list_boot_scripts) | **GET** /api/boot-scripts | Boot Scripts
[**list_image_builds**](ImageBuildsApi.md#list_image_builds) | **GET** /api/image-builds | Get All Image Builds
[**list_preseed_scripts**](ImageBuildsApi.md#list_preseed_scripts) | **GET** /api/preseed-scripts | Preseed Scripts
[**update_boot_script**](ImageBuildsApi.md#update_boot_script) | **PUT** /api/boot-scripts/{id} | Update a Boot Script
[**update_image_build**](ImageBuildsApi.md#update_image_build) | **PUT** /api/image-builds/{id} | Update an Image Build
[**update_preseed_script**](ImageBuildsApi.md#update_preseed_script) | **PUT** /api/preseed-scripts/{id} | Update a Preseed Script


# **add_boot_script**
> AddBootScript200Response add_boot_script()

Create a Boot Script

Create a Boot Script

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.add_boot_script200_response import AddBootScript200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_boot_script_request import AddBootScriptRequest
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    add_boot_script_request = AddBootScriptRequest(
        boot_script=BootScriptsCreate(
            file_name="file_name_example",
            content="content_example",
        ),
    ) # AddBootScriptRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Boot Script
        api_response = api_instance.add_boot_script(add_boot_script_request=add_boot_script_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->add_boot_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_boot_script_request** | [**AddBootScriptRequest**](AddBootScriptRequest.md)|  | [optional]

### Return type

[**AddBootScript200Response**](AddBootScript200Response.md)

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

# **add_image_build**
> AddImageBuild200Response add_image_build()

Create an Image Build

Create an Image Build

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.add_image_build200_response import AddImageBuild200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_image_build_request import AddImageBuildRequest
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    add_image_build_request = AddImageBuildRequest(
        image_build=ImageBuildCreate(
            name="name_example",
            description="description_example",
            type="vmware",
            site=ImageBuildCreateSite(
                id=1,
            ),
            zone=ImageBuildCreateZone(
                id=1,
            ),
            config={},
            boot_script=ImageBuildCreateBootScript(
                id=1,
            ),
            preseed_script=ImageBuildCreatePreseedScript(
                id=1,
            ),
            ssh_username="ssh_username_example",
            ssh_password="ssh_password_example",
            storage_provider="storage_provider_example",
            is_cloud_init="is_cloud_init_example",
            build_output_name="build_output_name_example",
            conversion_formats="conversion_formats_example",
            keep_results=0,
        ),
    ) # AddImageBuildRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Image Build
        api_response = api_instance.add_image_build(add_image_build_request=add_image_build_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->add_image_build: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_image_build_request** | [**AddImageBuildRequest**](AddImageBuildRequest.md)|  | [optional]

### Return type

[**AddImageBuild200Response**](AddImageBuild200Response.md)

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

# **add_preseed_script**
> AddPreseedScript200Response add_preseed_script()

Create a Preseed Script

Create a Preseed Script

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.add_preseed_script200_response import AddPreseedScript200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_preseed_script_request import AddPreseedScriptRequest
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    add_preseed_script_request = AddPreseedScriptRequest(
        preseed_script=PreseedScriptsCreate(
            file_name="file_name_example",
            content="content_example",
        ),
    ) # AddPreseedScriptRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Preseed Script
        api_response = api_instance.add_preseed_script(add_preseed_script_request=add_preseed_script_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->add_preseed_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_preseed_script_request** | [**AddPreseedScriptRequest**](AddPreseedScriptRequest.md)|  | [optional]

### Return type

[**AddPreseedScript200Response**](AddPreseedScript200Response.md)

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

# **delete_boot_script**
> Model200Success delete_boot_script(id)

Delete a Boot Script

Will delete a boot script from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Boot Script
        api_response = api_instance.delete_boot_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->delete_boot_script: %s\n" % e)
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

# **delete_image_build**
> Model200Success delete_image_build(id)

Delete an Image Build

Will delete an image build from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete an Image Build
        api_response = api_instance.delete_image_build(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->delete_image_build: %s\n" % e)
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

# **delete_preseed_script**
> Model200Success delete_preseed_script(id)

Delete a Preseed Script

Will delete a preseed script from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Preseed Script
        api_response = api_instance.delete_preseed_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->delete_preseed_script: %s\n" % e)
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

# **execute_image_build**
> Model200Success execute_image_build(id)

Run an Image Build

Running an image build is done asynchronously. This api will kick off the new execution and update the image build status.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Run an Image Build
        api_response = api_instance.execute_image_build(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->execute_image_build: %s\n" % e)
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

# **get_boot_script**
> AddBootScript200ResponseAllOf get_boot_script(id)

Get a Specific Boot Script

This endpoint retrieves a specific boot script.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.add_boot_script200_response_all_of import AddBootScript200ResponseAllOf
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Boot Script
        api_response = api_instance.get_boot_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->get_boot_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddBootScript200ResponseAllOf**](AddBootScript200ResponseAllOf.md)

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

# **get_image_build**
> GetImageBuild200Response get_image_build(id)

Get a Specific Image Build

This endpoint retrieves a specific image build.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.get_image_build200_response import GetImageBuild200Response
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Image Build
        api_response = api_instance.get_image_build(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->get_image_build: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetImageBuild200Response**](GetImageBuild200Response.md)

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

# **get_image_build_executions**
> GetImageBuildExecutions200Response get_image_build_executions(id)

List Image Build Executions

List all executions for an image build. This same info is also returned by the image build GET api, which returns the 100 most recent executions.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.get_image_build_executions200_response import GetImageBuildExecutions200Response
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    build_number = 4 # int | If specified will return an exact match on buildNumber (optional)
    status = "running" # str | Filter by status (optional)

    # example passing only required values which don't have defaults set
    try:
        # List Image Build Executions
        api_response = api_instance.get_image_build_executions(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->get_image_build_executions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List Image Build Executions
        api_response = api_instance.get_image_build_executions(id, build_number=build_number, status=status)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->get_image_build_executions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **build_number** | **int**| If specified will return an exact match on buildNumber | [optional]
 **status** | **str**| Filter by status | [optional]

### Return type

[**GetImageBuildExecutions200Response**](GetImageBuildExecutions200Response.md)

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

# **get_preseed_script**
> GetPreseedScript200Response get_preseed_script(id)

Get a Specific Preseed Script

This endpoint retrieves a specific preseed script.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.get_preseed_script200_response import GetPreseedScript200Response
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Preseed Script
        api_response = api_instance.get_preseed_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->get_preseed_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPreseedScript200Response**](GetPreseedScript200Response.md)

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

# **list_boot_scripts**
> ListBootScripts200Response list_boot_scripts()

Boot Scripts

Boot Scripts are used in the Image Builder service. See Image Builds.  This endpoint retrieves all boot scripts associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_boot_scripts200_response import ListBootScripts200Response
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Boot Scripts
        api_response = api_instance.list_boot_scripts(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->list_boot_scripts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListBootScripts200Response**](ListBootScripts200Response.md)

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

# **list_image_builds**
> ListImageBuilds200Response list_image_builds()

Get All Image Builds

This endpoint retrieves all image builds associated with the account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.list_image_builds200_response import ListImageBuilds200Response
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Image Builds
        api_response = api_instance.list_image_builds(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->list_image_builds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListImageBuilds200Response**](ListImageBuilds200Response.md)

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

# **list_preseed_scripts**
> ListPreseedScripts200Response list_preseed_scripts()

Preseed Scripts

Preseed Scripts are used in the Image Builder service. See Image Builds.  This endpoint retrieves all preseed scripts associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.list_preseed_scripts200_response import ListPreseedScripts200Response
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Preseed Scripts
        api_response = api_instance.list_preseed_scripts(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->list_preseed_scripts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListPreseedScripts200Response**](ListPreseedScripts200Response.md)

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

# **update_boot_script**
> AddBootScript200Response update_boot_script(id)

Update a Boot Script

Update a Boot Script

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.add_boot_script200_response import AddBootScript200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_boot_script_request import AddBootScriptRequest
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_boot_script_request = AddBootScriptRequest(
        boot_script=BootScriptsCreate(
            file_name="file_name_example",
            content="content_example",
        ),
    ) # AddBootScriptRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Boot Script
        api_response = api_instance.update_boot_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->update_boot_script: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Boot Script
        api_response = api_instance.update_boot_script(id, add_boot_script_request=add_boot_script_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->update_boot_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_boot_script_request** | [**AddBootScriptRequest**](AddBootScriptRequest.md)|  | [optional]

### Return type

[**AddBootScript200Response**](AddBootScript200Response.md)

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

# **update_image_build**
> UpdateImageBuild200Response update_image_build(id)

Update an Image Build

Update an Image Build

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.update_image_build200_response import UpdateImageBuild200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_image_build_request import AddImageBuildRequest
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_image_build_request = AddImageBuildRequest(
        image_build=ImageBuildCreate(
            name="name_example",
            description="description_example",
            type="vmware",
            site=ImageBuildCreateSite(
                id=1,
            ),
            zone=ImageBuildCreateZone(
                id=1,
            ),
            config={},
            boot_script=ImageBuildCreateBootScript(
                id=1,
            ),
            preseed_script=ImageBuildCreatePreseedScript(
                id=1,
            ),
            ssh_username="ssh_username_example",
            ssh_password="ssh_password_example",
            storage_provider="storage_provider_example",
            is_cloud_init="is_cloud_init_example",
            build_output_name="build_output_name_example",
            conversion_formats="conversion_formats_example",
            keep_results=0,
        ),
    ) # AddImageBuildRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Image Build
        api_response = api_instance.update_image_build(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->update_image_build: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Image Build
        api_response = api_instance.update_image_build(id, add_image_build_request=add_image_build_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->update_image_build: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_image_build_request** | [**AddImageBuildRequest**](AddImageBuildRequest.md)|  | [optional]

### Return type

[**UpdateImageBuild200Response**](UpdateImageBuild200Response.md)

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

# **update_preseed_script**
> AddPreseedScript200Response update_preseed_script(id)

Update a Preseed Script

Update a Preseed Script

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import image_builds_api
from openapi_client.model.add_preseed_script200_response import AddPreseedScript200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_preseed_script_request import AddPreseedScriptRequest
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
    api_instance = image_builds_api.ImageBuildsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_preseed_script_request = AddPreseedScriptRequest(
        preseed_script=PreseedScriptsCreate(
            file_name="file_name_example",
            content="content_example",
        ),
    ) # AddPreseedScriptRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Preseed Script
        api_response = api_instance.update_preseed_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->update_preseed_script: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Preseed Script
        api_response = api_instance.update_preseed_script(id, add_preseed_script_request=add_preseed_script_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ImageBuildsApi->update_preseed_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_preseed_script_request** | [**AddPreseedScriptRequest**](AddPreseedScriptRequest.md)|  | [optional]

### Return type

[**AddPreseedScript200Response**](AddPreseedScript200Response.md)

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

