# openapi_client.AppsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_app_instance**](AppsApi.md#add_app_instance) | **POST** /api/apps/{id}/add-instance | Add Existing Instance to App
[**add_app_undo_delete**](AppsApi.md#add_app_undo_delete) | **PUT** /api/apps/{id}/cancel-removal | Undo Delete of an App
[**add_apps**](AppsApi.md#add_apps) | **POST** /api/apps | Create an App
[**apply_app_state**](AppsApi.md#apply_app_state) | **POST** /api/apps/{id}/apply | Apply State of an App
[**delete_app**](AppsApi.md#delete_app) | **DELETE** /api/apps/{id} | Delete an App
[**get_app**](AppsApi.md#get_app) | **GET** /api/apps/{id} | Get a Specific App
[**get_app_security_groups**](AppsApi.md#get_app_security_groups) | **GET** /api/apps/{id}/security-groups | Get Security Groups for an App
[**get_app_state**](AppsApi.md#get_app_state) | **GET** /api/apps/{id}/state | Get State of an App
[**get_wiki_app**](AppsApi.md#get_wiki_app) | **GET** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**list_apps**](AppsApi.md#list_apps) | **GET** /api/apps | Get All Apps
[**prepare_app_apply**](AppsApi.md#prepare_app_apply) | **GET** /api/apps/{id}/prepare-apply | Prepare To Apply an App
[**refresh_app_state**](AppsApi.md#refresh_app_state) | **POST** /api/apps/{id}/refresh | Refresh State of an App
[**remove_app_instance**](AppsApi.md#remove_app_instance) | **POST** /api/apps/{id}/remove-instance | Remove Instance from App
[**set_app_security_groups**](AppsApi.md#set_app_security_groups) | **POST** /api/apps/{id}/security-groups | Set Security Groups for an App
[**update_app**](AppsApi.md#update_app) | **PUT** /api/apps/{id} | Updating an App
[**update_wiki_app**](AppsApi.md#update_wiki_app) | **PUT** /api/apps/{id}/wiki | Update an App Wiki Page
[**validate_app_state**](AppsApi.md#validate_app_state) | **POST** /api/apps/{id}/validate-apply | Validate Apply State for an App


# **add_app_instance**
> GetApp200Response add_app_instance(id)

Add Existing Instance to App

Add Existing Instance to App

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.add_app_instance_request import AddAppInstanceRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app200_response import GetApp200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_app_instance_request = AddAppInstanceRequest(
        instance_id=1,
        tier_name="tier_name_example",
    ) # AddAppInstanceRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Add Existing Instance to App
        api_response = api_instance.add_app_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->add_app_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Add Existing Instance to App
        api_response = api_instance.add_app_instance(id, add_app_instance_request=add_app_instance_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->add_app_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_app_instance_request** | [**AddAppInstanceRequest**](AddAppInstanceRequest.md)|  | [optional]

### Return type

[**GetApp200Response**](GetApp200Response.md)

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

# **add_app_undo_delete**
> GetApp200Response add_app_undo_delete(id)

Undo Delete of an App

This operation will undo the delete of an app that is pending removal.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app200_response import GetApp200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Undo Delete of an App
        api_response = api_instance.add_app_undo_delete(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->add_app_undo_delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetApp200Response**](GetApp200Response.md)

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

# **add_apps**
> AddApps200Response add_apps()

Create an App

Create an App

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.app_create import AppCreate
from openapi_client.model.add_apps200_response import AddApps200Response
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
    api_instance = apps_api.AppsApi(api_client)
    app_create = AppCreate(
        template_id=1,
        blueprint_id=AppCreateBlueprintId(None),
        name="name_example",
        description="description_example",
        labels=[
            "labels_example",
        ],
        group=AppCreateGroup(
            id=1,
        ),
        default_cloud=AppCreateDefaultCloud(
            id=1,
        ),
        environment="environment_example",
        tiers={},
    ) # AppCreate |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an App
        api_response = api_instance.add_apps(app_create=app_create)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->add_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_create** | [**AppCreate**](AppCreate.md)|  | [optional]

### Return type

[**AddApps200Response**](AddApps200Response.md)

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

# **apply_app_state**
> Model200Success apply_app_state(id)

Apply State of an App

This endpoint provides a way to apply the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.apply_app_state_request import ApplyAppStateRequest
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    apply_app_state_request = ApplyAppStateRequest(
        template_parameter={},
    ) # ApplyAppStateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Apply State of an App
        api_response = api_instance.apply_app_state(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->apply_app_state: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Apply State of an App
        api_response = api_instance.apply_app_state(id, apply_app_state_request=apply_app_state_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->apply_app_state: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **apply_app_state_request** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional]

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

# **delete_app**
> Model200Success delete_app(id)

Delete an App

Will delete an app. Use removeInstances=on to also delete the instances in the app and all associated monitors and backups.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    remove_instances = "on" # str | Remove Instances (optional) if omitted the server will use the default value of "off"
    preserve_volumes = "on" # str | Preserve Volumes (optional) if omitted the server will use the default value of "off"
    keep_backups = "on" # str | Preserve copy of backups (optional) if omitted the server will use the default value of "off"
    release_floating_ips = "off" # str | Release Floating IPs (optional) if omitted the server will use the default value of "on"
    release_eips = "off" # str | Alias for releaseFloatingIps (optional) if omitted the server will use the default value of "on"
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete an App
        api_response = api_instance.delete_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->delete_app: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete an App
        api_response = api_instance.delete_app(id, remove_instances=remove_instances, preserve_volumes=preserve_volumes, keep_backups=keep_backups, release_floating_ips=release_floating_ips, release_eips=release_eips, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->delete_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_instances** | **str**| Remove Instances | [optional] if omitted the server will use the default value of "off"
 **preserve_volumes** | **str**| Preserve Volumes | [optional] if omitted the server will use the default value of "off"
 **keep_backups** | **str**| Preserve copy of backups | [optional] if omitted the server will use the default value of "off"
 **release_floating_ips** | **str**| Release Floating IPs | [optional] if omitted the server will use the default value of "on"
 **release_eips** | **str**| Alias for releaseFloatingIps | [optional] if omitted the server will use the default value of "on"
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

# **get_app**
> GetApp200Response get_app(id)

Get a Specific App

This endpoint retrieves a specific app.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app200_response import GetApp200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific App
        api_response = api_instance.get_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->get_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetApp200Response**](GetApp200Response.md)

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

# **get_app_security_groups**
> GetAppSecurityGroups200Response get_app_security_groups(id)

Get Security Groups for an App

This returns a list of all of the security groups applied to an app and whether the firewall is enabled.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app_security_groups200_response import GetAppSecurityGroups200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Security Groups for an App
        api_response = api_instance.get_app_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->get_app_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetAppSecurityGroups200Response**](GetAppSecurityGroups200Response.md)

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

# **get_app_state**
> GetAppState200Response get_app_state(id)

Get State of an App

This endpoint provides a way to view the state of an app. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app_state200_response import GetAppState200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get State of an App
        api_response = api_instance.get_app_state(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->get_app_state: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetAppState200Response**](GetAppState200Response.md)

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

# **get_wiki_app**
> GetWikiApp200Response get_wiki_app(id)

Retrieves an App Wiki Page

This endpoint retrieves an app Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves an App Wiki Page
        api_response = api_instance.get_wiki_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->get_wiki_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **list_apps**
> ListApps200Response list_apps()

Get All Apps

This endpoint retrieves a paginated list of apps. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.list_apps200_response import ListApps200Response
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
    api_instance = apps_api.AppsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    created_by = 63 # int | The User ID for Filtering (optional)
    show_deleted = True # bool | If true, includes instances in pending removal status. (optional) if omitted the server will use the default value of False
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Apps
        api_response = api_instance.list_apps(max=max, offset=offset, name=name, phrase=phrase, created_by=created_by, show_deleted=show_deleted, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->list_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **created_by** | **int**| The User ID for Filtering | [optional]
 **show_deleted** | **bool**| If true, includes instances in pending removal status. | [optional] if omitted the server will use the default value of False
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListApps200Response**](ListApps200Response.md)

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

# **prepare_app_apply**
> PrepareAppApply200Response prepare_app_apply(id)

Prepare To Apply an App

This endpoint provides a way to view the current app configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.prepare_app_apply200_response import PrepareAppApply200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Prepare To Apply an App
        api_response = api_instance.prepare_app_apply(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->prepare_app_apply: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**PrepareAppApply200Response**](PrepareAppApply200Response.md)

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

# **refresh_app_state**
> Model200Success refresh_app_state(id)

Refresh State of an App

This endpoint provides a way to refresh the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.   

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    body = {} # {str: (bool, date, datetime, dict, float, int, list, str, none_type)} |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Refresh State of an App
        api_response = api_instance.refresh_app_state(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->refresh_app_state: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Refresh State of an App
        api_response = api_instance.refresh_app_state(id, body=body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->refresh_app_state: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **body** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}**|  | [optional]

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

# **remove_app_instance**
> GetApp200Response remove_app_instance(id)

Remove Instance from App

Remove Instance from App

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.remove_app_instance_request import RemoveAppInstanceRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app200_response import GetApp200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    remove_app_instance_request = RemoveAppInstanceRequest(
        instance_id=1,
    ) # RemoveAppInstanceRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Remove Instance from App
        api_response = api_instance.remove_app_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->remove_app_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Remove Instance from App
        api_response = api_instance.remove_app_instance(id, remove_app_instance_request=remove_app_instance_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->remove_app_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_app_instance_request** | [**RemoveAppInstanceRequest**](RemoveAppInstanceRequest.md)|  | [optional]

### Return type

[**GetApp200Response**](GetApp200Response.md)

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

# **set_app_security_groups**
> SetAppSecurityGroups200Response set_app_security_groups(id)

Set Security Groups for an App

Set Security Groups for an App

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.set_app_security_groups200_response import SetAppSecurityGroups200Response
from openapi_client.model.set_app_security_groups_request import SetAppSecurityGroupsRequest
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    set_app_security_groups_request = SetAppSecurityGroupsRequest(None) # SetAppSecurityGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Set Security Groups for an App
        api_response = api_instance.set_app_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->set_app_security_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Set Security Groups for an App
        api_response = api_instance.set_app_security_groups(id, set_app_security_groups_request=set_app_security_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->set_app_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **set_app_security_groups_request** | [**SetAppSecurityGroupsRequest**](SetAppSecurityGroupsRequest.md)|  | [optional]

### Return type

[**SetAppSecurityGroups200Response**](SetAppSecurityGroups200Response.md)

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

# **update_app**
> GetApp200Response update_app(id)

Updating an App

This endpoint provides updating of some basic app settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.app_update import AppUpdate
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_app200_response import GetApp200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    app_update = AppUpdate(
        name="name_example",
        description="description_example",
        labels=[
            "labels_example",
        ],
        environment="environment_example",
        owner_id=1,
    ) # AppUpdate |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating an App
        api_response = api_instance.update_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->update_app: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating an App
        api_response = api_instance.update_app(id, app_update=app_update)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->update_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **app_update** | [**AppUpdate**](AppUpdate.md)|  | [optional]

### Return type

[**GetApp200Response**](GetApp200Response.md)

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

# **update_wiki_app**
> UpdateWikiApp200Response update_wiki_app(id)

Update an App Wiki Page

Updates an app Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an App Wiki Page
        api_response = api_instance.update_wiki_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->update_wiki_app: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an App Wiki Page
        api_response = api_instance.update_wiki_app(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->update_wiki_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **validate_app_state**
> ValidateAppState200Response validate_app_state(id)

Validate Apply State for an App

This endpoint provides a way to validate app configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import apps_api
from openapi_client.model.apply_app_state_request import ApplyAppStateRequest
from openapi_client.model.validate_app_state200_response import ValidateAppState200Response
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
    api_instance = apps_api.AppsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    apply_app_state_request = ApplyAppStateRequest(
        template_parameter={},
    ) # ApplyAppStateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Validate Apply State for an App
        api_response = api_instance.validate_app_state(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->validate_app_state: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Validate Apply State for an App
        api_response = api_instance.validate_app_state(id, apply_app_state_request=apply_app_state_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AppsApi->validate_app_state: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **apply_app_state_request** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional]

### Return type

[**ValidateAppState200Response**](ValidateAppState200Response.md)

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

