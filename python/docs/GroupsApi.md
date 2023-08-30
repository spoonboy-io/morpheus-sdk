# openapi_client.GroupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_groups**](GroupsApi.md#add_groups) | **POST** /api/groups | Creates a Group
[**get_groups**](GroupsApi.md#get_groups) | **GET** /api/groups/{id} | Retrieves a Specific Group
[**get_wiki_group**](GroupsApi.md#get_wiki_group) | **GET** /api/groups/{id}/wiki | Retrieves a Group Wiki Page
[**list_groups**](GroupsApi.md#list_groups) | **GET** /api/groups | Retrieves all Groups
[**remove_groups**](GroupsApi.md#remove_groups) | **DELETE** /api/groups/{id} | Deletes a Group
[**update_groups**](GroupsApi.md#update_groups) | **PUT** /api/groups/{id} | Updates a Group
[**update_groups_zones**](GroupsApi.md#update_groups_zones) | **PUT** /api/groups/{id}/update-zones | Updates a Group&#39;s Zones
[**update_wiki_group**](GroupsApi.md#update_wiki_group) | **PUT** /api/groups/{id}/wiki | Update a Group Wiki Page


# **add_groups**
> AddGroups200Response add_groups()

Creates a Group

Creates a group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
from openapi_client.model.add_groups_request import AddGroupsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_groups200_response import AddGroups200Response
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
    api_instance = groups_api.GroupsApi(api_client)
    add_groups_request = AddGroupsRequest(
        group=GroupCreate(
            name="name_example",
            code="code_example",
            location="location_example",
            config=GroupCreateConfig(
                dns_integration_id="dns_integration_id_example",
                config_cmdb_id="config_cmdb_id_example",
                config_cm_id="config_cm_id_example",
                service_registry_id="service_registry_id_example",
                config_management_id="config_management_id_example",
                config_cmdb_discovery=True,
            ),
        ),
    ) # AddGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Group
        api_response = api_instance.add_groups(add_groups_request=add_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->add_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_groups_request** | [**AddGroupsRequest**](AddGroupsRequest.md)|  | [optional]

### Return type

[**AddGroups200Response**](AddGroups200Response.md)

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

# **get_groups**
> AddGroups200ResponseAllOf get_groups(id)

Retrieves a Specific Group

Retrieves a specific group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
from openapi_client.model.add_groups200_response_all_of import AddGroups200ResponseAllOf
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
    api_instance = groups_api.GroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Group
        api_response = api_instance.get_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->get_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddGroups200ResponseAllOf**](AddGroups200ResponseAllOf.md)

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

# **get_wiki_group**
> GetWikiApp200Response get_wiki_group(id)

Retrieves a Group Wiki Page

This endpoint retrieves a group Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
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
    api_instance = groups_api.GroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Group Wiki Page
        api_response = api_instance.get_wiki_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->get_wiki_group: %s\n" % e)
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

# **list_groups**
> ListGroups200Response list_groups()

Retrieves all Groups

Retrieves all groups. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
from openapi_client.model.list_groups200_response import ListGroups200Response
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
    api_instance = groups_api.GroupsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Groups
        api_response = api_instance.list_groups(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->list_groups: %s\n" % e)
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

[**ListGroups200Response**](ListGroups200Response.md)

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

# **remove_groups**
> Model200Success remove_groups(id)

Deletes a Group

Deletes a specified Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
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
    api_instance = groups_api.GroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Group
        api_response = api_instance.remove_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->remove_groups: %s\n" % e)
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

# **update_groups**
> UpdateGroups200Response update_groups(id)

Updates a Group

Updates a group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
from openapi_client.model.update_groups200_response import UpdateGroups200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_groups_request import UpdateGroupsRequest
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
    api_instance = groups_api.GroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_groups_request = UpdateGroupsRequest(
        group=GroupUpdate(
            name="name_example",
            code="code_example",
            location="location_example",
            config=GroupCreateConfig(
                dns_integration_id="dns_integration_id_example",
                config_cmdb_id="config_cmdb_id_example",
                config_cm_id="config_cm_id_example",
                service_registry_id="service_registry_id_example",
                config_management_id="config_management_id_example",
                config_cmdb_discovery=True,
            ),
        ),
    ) # UpdateGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Group
        api_response = api_instance.update_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->update_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Group
        api_response = api_instance.update_groups(id, update_groups_request=update_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->update_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_groups_request** | [**UpdateGroupsRequest**](UpdateGroupsRequest.md)|  | [optional]

### Return type

[**UpdateGroups200Response**](UpdateGroups200Response.md)

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

# **update_groups_zones**
> Model200Success update_groups_zones(id)

Updates a Group's Zones

Updates a group's zones. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_groups_zones_request import UpdateGroupsZonesRequest
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
    api_instance = groups_api.GroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_groups_zones_request = UpdateGroupsZonesRequest(
        group=UpdateGroupsZonesRequestGroup(
            zones=[{"id":1},{"id":5}],
        ),
    ) # UpdateGroupsZonesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Group's Zones
        api_response = api_instance.update_groups_zones(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->update_groups_zones: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Group's Zones
        api_response = api_instance.update_groups_zones(id, update_groups_zones_request=update_groups_zones_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->update_groups_zones: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_groups_zones_request** | [**UpdateGroupsZonesRequest**](UpdateGroupsZonesRequest.md)|  | [optional]

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

# **update_wiki_group**
> UpdateWikiApp200Response update_wiki_group(id)

Update a Group Wiki Page

Updates a group Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import groups_api
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
    api_instance = groups_api.GroupsApi(api_client)
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
        # Update a Group Wiki Page
        api_response = api_instance.update_wiki_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->update_wiki_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Group Wiki Page
        api_response = api_instance.update_wiki_group(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GroupsApi->update_wiki_group: %s\n" % e)
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

