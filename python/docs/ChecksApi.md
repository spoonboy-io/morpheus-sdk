# openapi_client.ChecksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_check_apps**](ChecksApi.md#add_check_apps) | **POST** /api/monitoring/apps | Create a New Check App
[**add_check_groups**](ChecksApi.md#add_check_groups) | **POST** /api/monitoring/groups | Create a New Check Group
[**add_checks**](ChecksApi.md#add_checks) | **POST** /api/monitoring/checks | Create a New Check
[**delete_check_apps**](ChecksApi.md#delete_check_apps) | **DELETE** /api/monitoring/apps/{id} | Delete a Specific Check App
[**delete_check_groups**](ChecksApi.md#delete_check_groups) | **DELETE** /api/monitoring/groups/{id} | Delete a Specific Check Group
[**delete_checks**](ChecksApi.md#delete_checks) | **DELETE** /api/monitoring/checks/{id} | Delete a Specific Check
[**get_check_apps**](ChecksApi.md#get_check_apps) | **GET** /api/monitoring/apps/{id} | Get a Specific Check App
[**get_check_groups**](ChecksApi.md#get_check_groups) | **GET** /api/monitoring/groups/{id} | Get a Specific Check Group
[**get_check_types**](ChecksApi.md#get_check_types) | **GET** /api/monitoring/check-types/{id} | Get a Specific Check Type
[**get_checks**](ChecksApi.md#get_checks) | **GET** /api/monitoring/checks/{id} | Get a Specific Check
[**list_check_apps**](ChecksApi.md#list_check_apps) | **GET** /api/monitoring/apps | List All Check Apps
[**list_check_groups**](ChecksApi.md#list_check_groups) | **GET** /api/monitoring/groups | List All Check Groups
[**list_check_types**](ChecksApi.md#list_check_types) | **GET** /api/monitoring/check-types | List All Check Types
[**list_checks**](ChecksApi.md#list_checks) | **GET** /api/monitoring/checks | List All Checks
[**update_check_apps**](ChecksApi.md#update_check_apps) | **PUT** /api/monitoring/apps/{id} | Update Check App
[**update_check_groups**](ChecksApi.md#update_check_groups) | **PUT** /api/monitoring/groups/{id} | Update Check Group
[**update_checks**](ChecksApi.md#update_checks) | **PUT** /api/monitoring/checks/{id} | Updates a Check
[**update_mute_all_check_apps**](ChecksApi.md#update_mute_all_check_apps) | **PUT** /api/monitoring/apps/mute-all | Mute All Check Apps
[**update_mute_all_check_groups**](ChecksApi.md#update_mute_all_check_groups) | **PUT** /api/monitoring/groups/mute-all | Mute All Check Groups
[**update_mute_all_checks**](ChecksApi.md#update_mute_all_checks) | **PUT** /api/monitoring/checks/mute-all | Mute All Checks
[**update_mute_check_apps**](ChecksApi.md#update_mute_check_apps) | **PUT** /api/monitoring/apps/{id}/mute | Mute Check App
[**update_mute_check_groups**](ChecksApi.md#update_mute_check_groups) | **PUT** /api/monitoring/groups/{id}/mute | Mute Check Group
[**update_mute_checks**](ChecksApi.md#update_mute_checks) | **PUT** /api/monitoring/checks/{id}/mute | Mute Check


# **add_check_apps**
> AddCheckApps200Response add_check_apps()

Create a New Check App

Create a new check app.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.add_check_apps200_response import AddCheckApps200Response
from openapi_client.model.add_check_apps_request import AddCheckAppsRequest
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
    api_instance = checks_api.ChecksApi(api_client)
    add_check_apps_request = AddCheckAppsRequest(
        monitor_app=AddCheckAppsRequestMonitorApp(
            name="My Check App",
            description="My cool description",
            in_uptime=True,
            severity="critical",
            active=True,
            checks=[
                1,
            ],
            check_groups=[
                1,
            ],
        ),
    ) # AddCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Check App
        api_response = api_instance.add_check_apps(add_check_apps_request=add_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->add_check_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_check_apps_request** | [**AddCheckAppsRequest**](AddCheckAppsRequest.md)|  | [optional]

### Return type

[**AddCheckApps200Response**](AddCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_check_groups**
> AddCheckGroups200Response add_check_groups()

Create a New Check Group

Create a new check group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.add_check_groups200_response import AddCheckGroups200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_check_groups_request import AddCheckGroupsRequest
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
    api_instance = checks_api.ChecksApi(api_client)
    add_check_groups_request = AddCheckGroupsRequest(
        check_group=AddCheckGroupsRequestCheckGroup(
            name="My Check Group",
            description="My cool description",
            min_happy=1,
            in_uptime=True,
            severity="critical",
            active=True,
            checks=[
                1,
            ],
        ),
    ) # AddCheckGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Check Group
        api_response = api_instance.add_check_groups(add_check_groups_request=add_check_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->add_check_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_check_groups_request** | [**AddCheckGroupsRequest**](AddCheckGroupsRequest.md)|  | [optional]

### Return type

[**AddCheckGroups200Response**](AddCheckGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_checks**
> AddChecks200Response add_checks()

Create a New Check

Create a new monitoring check.  SSH tunneling options allow the different check types to tunnel to a host via a proxy, and execute checks relative to the proxy. A SSH tunnel can use your account generated public and private key-pairs or SSH password. It is strongly recommended to use a key-pair.  To enable SSH tunneling for a check, add `tunnelOn`, `sshHost`, `sshUser`, and optionally, `sshPort` and `sshPassword` parameters to any check type config as seen earlier in the Check Types section. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.add_checks_request import AddChecksRequest
from openapi_client.model.add_checks200_response import AddChecks200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    add_checks_request = AddChecksRequest(
        check=AddChecksRequestCheck(),
    ) # AddChecksRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Check
        api_response = api_instance.add_checks(add_checks_request=add_checks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->add_checks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_checks_request** | [**AddChecksRequest**](AddChecksRequest.md)|  | [optional]

### Return type

[**AddChecks200Response**](AddChecks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_check_apps**
> Model200Success delete_check_apps(id)

Delete a Specific Check App

Delete an existing monitoring check app.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Check App
        api_response = api_instance.delete_check_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->delete_check_apps: %s\n" % e)
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_check_groups**
> Model200Success delete_check_groups(id)

Delete a Specific Check Group

Delete an existing monitoring check group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Check Group
        api_response = api_instance.delete_check_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->delete_check_groups: %s\n" % e)
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_checks**
> Model200Success delete_checks(id)

Delete a Specific Check

Delete an existing monitoring check.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Check
        api_response = api_instance.delete_checks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->delete_checks: %s\n" % e)
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_check_apps**
> GetCheckApps200Response get_check_apps(id)

Get a Specific Check App

Get details about a specific monitoring check app.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.get_check_apps200_response import GetCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Check App
        api_response = api_instance.get_check_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->get_check_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCheckApps200Response**](GetCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_check_groups**
> GetCheckGroups200Response get_check_groups(id)

Get a Specific Check Group

Get details about a specific monitoring check group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.get_check_groups200_response import GetCheckGroups200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Check Group
        api_response = api_instance.get_check_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->get_check_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCheckGroups200Response**](GetCheckGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_check_types**
> GetCheckTypes200Response get_check_types(id)

Get a Specific Check Type

Get details about a specific monitoring check type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_check_types200_response import GetCheckTypes200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Check Type
        api_response = api_instance.get_check_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->get_check_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCheckTypes200Response**](GetCheckTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Types Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_checks**
> AddChecks200ResponseAllOf get_checks(id)

Get a Specific Check

Get details about a specific monitoring check.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.add_checks200_response_all_of import AddChecks200ResponseAllOf
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Check
        api_response = api_instance.get_checks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->get_checks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddChecks200ResponseAllOf**](AddChecks200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Checks Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_check_apps**
> ListCheckApps200Response list_check_apps()

List All Check Apps

Get a list of check apps.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.list_check_apps200_response import ListCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    status = "running" # str | The instance status for filtering. (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted = True # bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Check Apps
        api_response = api_instance.list_check_apps(max=max, offset=offset, sort=sort, name=name, phrase=phrase, status=status, last_updated=last_updated, deleted=deleted)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->list_check_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **status** | **str**| The instance status for filtering. | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

### Return type

[**ListCheckApps200Response**](ListCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Check Apps |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_check_groups**
> ListCheckGroups200Response list_check_groups()

List All Check Groups

Get a list of check groups.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_check_groups200_response import ListCheckGroups200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    status = "running" # str | The instance status for filtering. (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted = True # bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Check Groups
        api_response = api_instance.list_check_groups(max=max, offset=offset, sort=sort, name=name, phrase=phrase, status=status, last_updated=last_updated, deleted=deleted)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->list_check_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **status** | **str**| The instance status for filtering. | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

### Return type

[**ListCheckGroups200Response**](ListCheckGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Check Groups |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_check_types**
> ListCheckTypes200Response list_check_types()

List All Check Types

Get a list of monitoring check types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_check_types200_response import ListCheckTypes200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Check Types
        api_response = api_instance.list_check_types(max=max, offset=offset, sort=sort, name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->list_check_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListCheckTypes200Response**](ListCheckTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Check Types |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_checks**
> ListChecks200Response list_checks()

List All Checks

Get a list of monitoring checks.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_checks200_response import ListChecks200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    status = "healthy" # str | Filter by health status. (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deleted = True # bool | If true, only deleted resources or instances in pending removal status are returned. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Checks
        api_response = api_instance.list_checks(max=max, offset=offset, sort=sort, name=name, phrase=phrase, status=status, last_updated=last_updated, deleted=deleted)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->list_checks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **status** | **str**| Filter by health status. | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **bool**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

### Return type

[**ListChecks200Response**](ListChecks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Checks |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_check_apps**
> UpdateCheckApps200Response update_check_apps(id)

Update Check App

Update an existing monitoring check app.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_check_apps_request import UpdateCheckAppsRequest
from openapi_client.model.update_check_apps200_response import UpdateCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_check_apps_request = UpdateCheckAppsRequest(
        monitor_app=UpdateCheckAppsRequestMonitorApp(
            name="My Check App",
            description="My cool description",
            in_uptime=True,
            severity="critical",
            active=True,
            checks=[
                1,
            ],
            check_groups=[
                1,
            ],
        ),
    ) # UpdateCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Check App
        api_response = api_instance.update_check_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_check_apps: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Check App
        api_response = api_instance.update_check_apps(id, update_check_apps_request=update_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_check_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_check_apps_request** | [**UpdateCheckAppsRequest**](UpdateCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateCheckApps200Response**](UpdateCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_check_groups**
> AddCheckGroups200Response update_check_groups(id)

Update Check Group

Update an existing monitoring check group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.add_check_groups200_response import AddCheckGroups200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_check_groups_request import UpdateCheckGroupsRequest
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_check_groups_request = UpdateCheckGroupsRequest(
        check_group=UpdateCheckGroupsRequestCheckGroup(
            name="My Check Group",
            description="My cool description",
            min_happy=1,
            in_uptime=True,
            severity="critical",
            active=True,
            checks=[
                1,
            ],
        ),
    ) # UpdateCheckGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Check Group
        api_response = api_instance.update_check_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_check_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Check Group
        api_response = api_instance.update_check_groups(id, update_check_groups_request=update_check_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_check_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_check_groups_request** | [**UpdateCheckGroupsRequest**](UpdateCheckGroupsRequest.md)|  | [optional]

### Return type

[**AddCheckGroups200Response**](AddCheckGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_checks**
> AddChecks200Response update_checks(id)

Updates a Check

Updates a monitoring check.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.add_checks200_response import AddChecks200Response
from openapi_client.model.update_checks_request import UpdateChecksRequest
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_checks_request = UpdateChecksRequest(
        check=UpdateChecksRequestCheck(),
    ) # UpdateChecksRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Check
        api_response = api_instance.update_checks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_checks: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Check
        api_response = api_instance.update_checks(id, update_checks_request=update_checks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_checks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_checks_request** | [**UpdateChecksRequest**](UpdateChecksRequest.md)|  | [optional]

### Return type

[**AddChecks200Response**](AddChecks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_all_check_apps**
> UpdateMuteAllCheckApps200Response update_mute_all_check_apps()

Mute All Check Apps

Mute all existing check apps.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_all_check_apps200_response import UpdateMuteAllCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute All Check Apps
        api_response = api_instance.update_mute_all_check_apps(update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_all_check_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteAllCheckApps200Response**](UpdateMuteAllCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mute All Check Apps Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_all_check_groups**
> UpdateMuteAllCheckApps200Response update_mute_all_check_groups()

Mute All Check Groups

Mute all existing check groups.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_all_check_apps200_response import UpdateMuteAllCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute All Check Groups
        api_response = api_instance.update_mute_all_check_groups(update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_all_check_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteAllCheckApps200Response**](UpdateMuteAllCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mute All Check Groups Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_all_checks**
> UpdateMuteAllCheckApps200Response update_mute_all_checks()

Mute All Checks

Mute all existing checks.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_all_check_apps200_response import UpdateMuteAllCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute All Checks
        api_response = api_instance.update_mute_all_checks(update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_all_checks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteAllCheckApps200Response**](UpdateMuteAllCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mute All Checks Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_check_apps**
> UpdateMuteCheckApps200Response update_mute_check_apps(id)

Mute Check App

Mute an existing check app.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_check_apps200_response import UpdateMuteCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Mute Check App
        api_response = api_instance.update_mute_check_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_check_apps: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute Check App
        api_response = api_instance.update_mute_check_apps(id, update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_check_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteCheckApps200Response**](UpdateMuteCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mute Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_check_groups**
> UpdateMuteCheckApps200Response update_mute_check_groups(id)

Mute Check Group

Mute an existing check group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_check_apps200_response import UpdateMuteCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Mute Check Group
        api_response = api_instance.update_mute_check_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_check_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute Check Group
        api_response = api_instance.update_mute_check_groups(id, update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_check_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteCheckApps200Response**](UpdateMuteCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mute Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_checks**
> UpdateMuteCheckApps200Response update_mute_checks(id)

Mute Check

Mute an existing check.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import checks_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_check_apps200_response import UpdateMuteCheckApps200Response
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
    api_instance = checks_api.ChecksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Mute Check
        api_response = api_instance.update_mute_checks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_checks: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute Check
        api_response = api_instance.update_mute_checks(id, update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ChecksApi->update_mute_checks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteCheckApps200Response**](UpdateMuteCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Mute Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

