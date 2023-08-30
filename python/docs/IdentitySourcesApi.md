# openapi_client.IdentitySourcesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_identity_sources**](IdentitySourcesApi.md#add_identity_sources) | **POST** /api/user-sources | Creates an Identity Source
[**get_identity_sources**](IdentitySourcesApi.md#get_identity_sources) | **GET** /api/user-sources/{id} | Retrieves a Specific Identity Source
[**list_identity_sources**](IdentitySourcesApi.md#list_identity_sources) | **GET** /api/user-sources | Retrieves all Identity Sources
[**remove_identity_sources**](IdentitySourcesApi.md#remove_identity_sources) | **DELETE** /api/user-sources/{id} | Deletes an Identity Source
[**update_identity_source_subdomains**](IdentitySourcesApi.md#update_identity_source_subdomains) | **PUT** /api/user-sources/{id}/subdomain | Updates an Identity Source Subdomain
[**update_identity_sources**](IdentitySourcesApi.md#update_identity_sources) | **PUT** /api/user-sources/{id} | Updates an Identity Source


# **add_identity_sources**
> AddIdentitySources200Response add_identity_sources()

Creates an Identity Source

Creates an identity source. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import identity_sources_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_identity_sources200_response import AddIdentitySources200Response
from openapi_client.model.user_source_create import UserSourceCreate
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
    api_instance = identity_sources_api.IdentitySourcesApi(api_client)
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    user_source_create = UserSourceCreate(
        user_source=UserSourceCreateUserSource(
            name="mydomain AD",
            type="type_example",
            description="description_example",
            default_account_role=UserSourceCreateUserSourceDefaultAccountRole(
                id=19,
            ),
            role_mappings=UserSourceCreateUserSourceRoleMappings(None),
            role_mapping_names={
                "key": "key_example",
            },
            allow_custom_mappings=True,
            manual_role_assignment=True,
            config=UserSourceCreateUserSourceConfig(None),
        ),
    ) # UserSourceCreate |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates an Identity Source
        api_response = api_instance.add_identity_sources(account_id=account_id, user_source_create=user_source_create)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->add_identity_sources: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **user_source_create** | [**UserSourceCreate**](UserSourceCreate.md)|  | [optional]

### Return type

[**AddIdentitySources200Response**](AddIdentitySources200Response.md)

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

# **get_identity_sources**
> GetIdentitySources200Response get_identity_sources(id)

Retrieves a Specific Identity Source

Retrieves a specific identity source. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import identity_sources_api
from openapi_client.model.get_identity_sources200_response import GetIdentitySources200Response
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
    api_instance = identity_sources_api.IdentitySourcesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Identity Source
        api_response = api_instance.get_identity_sources(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->get_identity_sources: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetIdentitySources200Response**](GetIdentitySources200Response.md)

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

# **list_identity_sources**
> ListIdentitySources200Response list_identity_sources()

Retrieves all Identity Sources

Retrieves all identity sources. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import identity_sources_api
from openapi_client.model.list_identity_sources200_response import ListIdentitySources200Response
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
    api_instance = identity_sources_api.IdentitySourcesApi(api_client)
    type = "ansibleTask" # str | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.  (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Identity Sources
        api_response = api_instance.list_identity_sources(type=type, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, account_id=account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->list_identity_sources: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type** | **str**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

### Return type

[**ListIdentitySources200Response**](ListIdentitySources200Response.md)

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

# **remove_identity_sources**
> Model200Success remove_identity_sources(id)

Deletes an Identity Source

Deletes a specified identity source. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import identity_sources_api
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
    api_instance = identity_sources_api.IdentitySourcesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes an Identity Source
        api_response = api_instance.remove_identity_sources(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->remove_identity_sources: %s\n" % e)
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

# **update_identity_source_subdomains**
> UpdateIdentitySourceSubdomains200Response update_identity_source_subdomains(id)

Updates an Identity Source Subdomain

Updates an identity source subdomain. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import identity_sources_api
from openapi_client.model.update_identity_source_subdomains_request import UpdateIdentitySourceSubdomainsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_identity_source_subdomains200_response import UpdateIdentitySourceSubdomains200Response
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
    api_instance = identity_sources_api.IdentitySourcesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_identity_source_subdomains_request = UpdateIdentitySourceSubdomainsRequest(
        subdomain="mynewdomain",
    ) # UpdateIdentitySourceSubdomainsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates an Identity Source Subdomain
        api_response = api_instance.update_identity_source_subdomains(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->update_identity_source_subdomains: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates an Identity Source Subdomain
        api_response = api_instance.update_identity_source_subdomains(id, update_identity_source_subdomains_request=update_identity_source_subdomains_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->update_identity_source_subdomains: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_identity_source_subdomains_request** | [**UpdateIdentitySourceSubdomainsRequest**](UpdateIdentitySourceSubdomainsRequest.md)|  | [optional]

### Return type

[**UpdateIdentitySourceSubdomains200Response**](UpdateIdentitySourceSubdomains200Response.md)

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

# **update_identity_sources**
> AddIdentitySources200Response update_identity_sources(id)

Updates an Identity Source

Updates an identity source. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import identity_sources_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_identity_sources200_response import AddIdentitySources200Response
from openapi_client.model.user_source_create import UserSourceCreate
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
    api_instance = identity_sources_api.IdentitySourcesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    user_source_create = UserSourceCreate(
        user_source=UserSourceCreateUserSource(
            name="mydomain AD",
            type="type_example",
            description="description_example",
            default_account_role=UserSourceCreateUserSourceDefaultAccountRole(
                id=19,
            ),
            role_mappings=UserSourceCreateUserSourceRoleMappings(None),
            role_mapping_names={
                "key": "key_example",
            },
            allow_custom_mappings=True,
            manual_role_assignment=True,
            config=UserSourceCreateUserSourceConfig(None),
        ),
    ) # UserSourceCreate |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates an Identity Source
        api_response = api_instance.update_identity_sources(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->update_identity_sources: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates an Identity Source
        api_response = api_instance.update_identity_sources(id, user_source_create=user_source_create)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IdentitySourcesApi->update_identity_sources: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **user_source_create** | [**UserSourceCreate**](UserSourceCreate.md)|  | [optional]

### Return type

[**AddIdentitySources200Response**](AddIdentitySources200Response.md)

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

