# openapi_client.BlueprintsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_blueprint**](BlueprintsApi.md#add_blueprint) | **POST** /api/blueprints | Create a Blueprint
[**delete_blueprint**](BlueprintsApi.md#delete_blueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint
[**get_blueprint**](BlueprintsApi.md#get_blueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint
[**list_blueprints**](BlueprintsApi.md#list_blueprints) | **GET** /api/blueprints | Get All Blueprints
[**update_blueprint**](BlueprintsApi.md#update_blueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint
[**update_blueprint_image**](BlueprintsApi.md#update_blueprint_image) | **POST** /api/blueprints/{id}/image | Update Blueprint Image
[**update_blueprint_permissions**](BlueprintsApi.md#update_blueprint_permissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions


# **add_blueprint**
> AddBlueprint200Response add_blueprint()

Create a Blueprint

Create a Blueprint

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
from openapi_client.model.add_blueprint200_response import AddBlueprint200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_blueprint_request import AddBlueprintRequest
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    add_blueprint_request = AddBlueprintRequest(None) # AddBlueprintRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Blueprint
        api_response = api_instance.add_blueprint(add_blueprint_request=add_blueprint_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->add_blueprint: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_blueprint_request** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional]

### Return type

[**AddBlueprint200Response**](AddBlueprint200Response.md)

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

# **delete_blueprint**
> Model200Success delete_blueprint(id)

Delete a Blueprint

Delete a Blueprint

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Blueprint
        api_response = api_instance.delete_blueprint(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->delete_blueprint: %s\n" % e)
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

# **get_blueprint**
> GetBlueprint200Response get_blueprint(id)

Get a Specific Blueprint

This endpoint retrieves a specific blueprint.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
from openapi_client.model.get_blueprint200_response import GetBlueprint200Response
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Blueprint
        api_response = api_instance.get_blueprint(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->get_blueprint: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

# **list_blueprints**
> ListBlueprints200Response list_blueprints()

Get All Blueprints

This endpoint retrieves all blueprints.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
from openapi_client.model.list_blueprints200_response import ListBlueprints200Response
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Blueprints
        api_response = api_instance.list_blueprints(max=max, offset=offset, name=name, phrase=phrase, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->list_blueprints: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListBlueprints200Response**](ListBlueprints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_blueprint**
> GetBlueprint200Response update_blueprint(id)

Updating a Blueprint

Update a Blueprint. This overwrites the entire config, so the entire blueprint config should be passed.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
from openapi_client.model.get_blueprint200_response import GetBlueprint200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_blueprint_request import AddBlueprintRequest
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_blueprint_request = AddBlueprintRequest(None) # AddBlueprintRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a Blueprint
        api_response = api_instance.update_blueprint(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->update_blueprint: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a Blueprint
        api_response = api_instance.update_blueprint(id, add_blueprint_request=add_blueprint_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->update_blueprint: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_blueprint_request** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional]

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

# **update_blueprint_image**
> GetBlueprint200Response update_blueprint_image(id)

Update Blueprint Image

Update Blueprint Image

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
from openapi_client.model.get_blueprint200_response import GetBlueprint200Response
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    template_image = open('/path/to/file', 'rb') # file_type |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Blueprint Image
        api_response = api_instance.update_blueprint_image(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->update_blueprint_image: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Blueprint Image
        api_response = api_instance.update_blueprint_image(id, template_image=template_image)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->update_blueprint_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **template_image** | **file_type**|  | [optional]

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

# **update_blueprint_permissions**
> GetBlueprint200Response update_blueprint_permissions(id)

Update Blueprint Permissions

Update Blueprint Permissions

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import blueprints_api
from openapi_client.model.get_blueprint200_response import GetBlueprint200Response
from openapi_client.model.update_blueprint_permissions_request import UpdateBlueprintPermissionsRequest
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
    api_instance = blueprints_api.BlueprintsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_blueprint_permissions_request = UpdateBlueprintPermissionsRequest(
        resource_permission=UpdateBlueprintPermissionsRequestResourcePermission(
            all=True,
            sites=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            owner_id=1,
        ),
    ) # UpdateBlueprintPermissionsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Blueprint Permissions
        api_response = api_instance.update_blueprint_permissions(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->update_blueprint_permissions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Blueprint Permissions
        api_response = api_instance.update_blueprint_permissions(id, update_blueprint_permissions_request=update_blueprint_permissions_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BlueprintsApi->update_blueprint_permissions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_blueprint_permissions_request** | [**UpdateBlueprintPermissionsRequest**](UpdateBlueprintPermissionsRequest.md)|  | [optional]

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

