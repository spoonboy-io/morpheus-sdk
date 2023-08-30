# openapi_client.UsersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_user**](UsersApi.md#add_user) | **POST** /api/users | Create a New User
[**delete_user**](UsersApi.md#delete_user) | **DELETE** /api/users/{id} | Delete a User
[**delete_user_settings_access_token**](UsersApi.md#delete_user_settings_access_token) | **PUT** /api/user-settings/clear-access-token | Revoke API Access Token
[**delete_user_settings_avatar**](UsersApi.md#delete_user_settings_avatar) | **DELETE** /api/user-settings/avatar | Delete Avatar
[**delete_user_settings_desktop_background**](UsersApi.md#delete_user_settings_desktop_background) | **DELETE** /api/user-settings/desktop-background | Delete Desktop Background
[**get_user**](UsersApi.md#get_user) | **GET** /api/users/{id} | Get a Specific User
[**get_user_permissions**](UsersApi.md#get_user_permissions) | **GET** /api/users/{id}/permissions | Get a Specific User Permissions
[**get_user_settings_api_clients**](UsersApi.md#get_user_settings_api_clients) | **GET** /api/user-settings/api-clients | Get Available API Clients
[**list_user_settings**](UsersApi.md#list_user_settings) | **GET** /api/user-settings | User Settings
[**list_users**](UsersApi.md#list_users) | **GET** /api/users | List All Users
[**list_users_available_roles**](UsersApi.md#list_users_available_roles) | **GET** /api/users/available-roles | List available roles for a user
[**update_user_settings**](UsersApi.md#update_user_settings) | **PUT** /api/user-settings | Update User Settings
[**update_user_settings_access_token**](UsersApi.md#update_user_settings_access_token) | **PUT** /api/user-settings/regenerate-access-token | Regenerate API Access Token
[**update_user_settings_avatar**](UsersApi.md#update_user_settings_avatar) | **POST** /api/user-settings/avatar | Update Avatar
[**update_user_settings_desktop_background**](UsersApi.md#update_user_settings_desktop_background) | **POST** /api/user-settings/desktop-background | Update Desktop Background
[**update_users**](UsersApi.md#update_users) | **PUT** /api/users/{id} | Update user
[**whoami**](UsersApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions


# **add_user**
> AddUserTenant200Response add_user()

Create a New User

Create a new user.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.add_user_tenant_request import AddUserTenantRequest
from openapi_client.model.add_user_tenant200_response import AddUserTenant200Response
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
    api_instance = users_api.UsersApi(api_client)
    account_id = 1 # int | Tenant ID, create user in a sub tenant account instead of your own. (optional)
    add_user_tenant_request = AddUserTenantRequest(
        user=UserCreate(
            first_name="John",
            last_name="Smith",
            username="jsmith",
            email="jsmith@email.com",
            password="Passw0rd!",
            roles=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            receive_notifications=True,
            linux_username="linux_username_example",
            linux_password="linux_password_example",
            linux_key_pair_id=1,
            windows_username="windows_username_example",
            windows_password="windows_password_example",
        ),
    ) # AddUserTenantRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New User
        api_response = api_instance.add_user(account_id=account_id, add_user_tenant_request=add_user_tenant_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->add_user: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Tenant ID, create user in a sub tenant account instead of your own. | [optional]
 **add_user_tenant_request** | [**AddUserTenantRequest**](AddUserTenantRequest.md)|  | [optional]

### Return type

[**AddUserTenant200Response**](AddUserTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_user**
> Model200Success delete_user(id)

Delete a User

Delete an existing user.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
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
    api_instance = users_api.UsersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a User
        api_response = api_instance.delete_user(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->delete_user: %s\n" % e)
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

# **delete_user_settings_access_token**
> Model200Success delete_user_settings_access_token()

Revoke API Access Token

This endpoint revokes your API access token for the specified client.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    client_id = "morph-cli" # str | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Revoke API Access Token
        api_response = api_instance.delete_user_settings_access_token(user_id=user_id, client_id=client_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->delete_user_settings_access_token: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **client_id** | **str**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_user_settings_avatar**
> Model200Success delete_user_settings_avatar()

Delete Avatar

Delete your avatar image.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete Avatar
        api_response = api_instance.delete_user_settings_avatar(user_id=user_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->delete_user_settings_avatar: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

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

# **delete_user_settings_desktop_background**
> Model200Success delete_user_settings_desktop_background()

Delete Desktop Background

Delete your desktop background image.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete Desktop Background
        api_response = api_instance.delete_user_settings_desktop_background(user_id=user_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->delete_user_settings_desktop_background: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

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

# **get_user**
> AddUserTenant200ResponseAllOf get_user(id)

Get a Specific User

This endpoint will retrieve a specific user by id if the user has permission to access the user.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_user_tenant200_response_all_of import AddUserTenant200ResponseAllOf
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
    api_instance = users_api.UsersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    include_access = True # bool | Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific User
        api_response = api_instance.get_user(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->get_user: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a Specific User
        api_response = api_instance.get_user(id, include_access=include_access)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->get_user: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **include_access** | **bool**| Include &#x60;access&#x60; information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). | [optional]

### Return type

[**AddUserTenant200ResponseAllOf**](AddUserTenant200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_permissions**
> GetUserPermissions200Response get_user_permissions(id)

Get a Specific User Permissions

This will list all the permissions for a specific user.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.get_user_permissions200_response import GetUserPermissions200Response
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
    api_instance = users_api.UsersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific User Permissions
        api_response = api_instance.get_user_permissions(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->get_user_permissions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetUserPermissions200Response**](GetUserPermissions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_settings_api_clients**
> GetUserSettingsApiClients200Response get_user_settings_api_clients()

Get Available API Clients

This endpoint retrieves a list of available API clients.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.get_user_settings_api_clients200_response import GetUserSettingsApiClients200Response
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Available API Clients
        api_response = api_instance.get_user_settings_api_clients(user_id=user_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->get_user_settings_api_clients: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**GetUserSettingsApiClients200Response**](GetUserSettingsApiClients200Response.md)

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

# **list_user_settings**
> ListUserSettings200Response list_user_settings()

User Settings

Provides API for managing your own user settings and api access tokens.  This endpoint retrieves your user settings and API access token information. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.list_user_settings200_response import ListUserSettings200Response
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # User Settings
        api_response = api_instance.list_user_settings(user_id=user_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->list_user_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**ListUserSettings200Response**](ListUserSettings200Response.md)

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

# **list_users**
> ListUsers200Response list_users()

List All Users

This endpoint retrieves all users in the current user's tenant account. Master tenant users with permission to manage subtenants can use `global=true` to find users across all tenants. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.list_users200_response import ListUsers200Response
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
    api_instance = users_api.UsersApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on username or email (optional)
    username = "administrator" # str | Username (optional)
    email = "mytest@email.com" # str | E-Mail Address (optional)
    role_id = 13 # int | Filter by Role ID (supports multiple values) (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    _global = False # bool | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Users
        api_response = api_instance.list_users(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, username=username, email=email, role_id=role_id, last_updated=last_updated, account_id=account_id, _global=_global)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->list_users: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on username or email | [optional]
 **username** | **str**| Username | [optional]
 **email** | **str**| E-Mail Address | [optional]
 **role_id** | **int**| Filter by Role ID (supports multiple values) | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **_global** | **bool**| Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] if omitted the server will use the default value of False

### Return type

[**ListUsers200Response**](ListUsers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Users |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_users_available_roles**
> UsersAvailableRoles list_users_available_roles()

List available roles for a user

Get a list of roles that can be assigned to a user.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.users_available_roles import UsersAvailableRoles
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
    api_instance = users_api.UsersApi(api_client)
    account_id = 1 # int | Tenant ID, find roles available to users of a sub tenant account. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List available roles for a user
        api_response = api_instance.list_users_available_roles(account_id=account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->list_users_available_roles: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Tenant ID, find roles available to users of a sub tenant account. | [optional]

### Return type

[**UsersAvailableRoles**](UsersAvailableRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of roles |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_user_settings**
> Model200Success update_user_settings()

Update User Settings

This endpoint allows updating user settings. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.update_user_settings_request1 import UpdateUserSettingsRequest1
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_user_settings_request import UpdateUserSettingsRequest
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    update_user_settings_request = UpdateUserSettingsRequest(
        user=UserSettingsUpdate(
            username="username_example",
            email="email_example",
            first_name="first_name_example",
            last_name="last_name_example",
            password="password_example",
            linux_username="linux_username_example",
            linux_password="linux_password_example",
            linux_key_pair_id=1,
            windows_username="windows_username_example",
            windows_password="windows_password_example",
            receive_notifications=True,
            default_group=UserSettingsUpdateDefaultGroup(
                id=1,
            ),
            default_cloud=UserSettingsUpdateDefaultCloud(
                id=1,
            ),
            default_persona=UserSettingsUpdateDefaultPersona(
                code="standard",
            ),
        ),
    ) # UpdateUserSettingsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update User Settings
        api_response = api_instance.update_user_settings(user_id=user_id, update_user_settings_request=update_user_settings_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->update_user_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **update_user_settings_request** | [**UpdateUserSettingsRequest**](UpdateUserSettingsRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_user_settings_access_token**
> UpdateUserSettingsAccessToken200Response update_user_settings_access_token()

Regenerate API Access Token

This endpoint regenerates your API access token for the specified client. If a current token exists, it is revoked and a new token is returned.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.update_user_settings_access_token200_response import UpdateUserSettingsAccessToken200Response
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    client_id = "morph-cli" # str | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Regenerate API Access Token
        api_response = api_instance.update_user_settings_access_token(user_id=user_id, client_id=client_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->update_user_settings_access_token: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **client_id** | **str**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional]

### Return type

[**UpdateUserSettingsAccessToken200Response**](UpdateUserSettingsAccessToken200Response.md)

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

# **update_user_settings_avatar**
> Model200Success update_user_settings_avatar()

Update Avatar

This endpoint updates your avatar image. Expects multipart form data as the request format, not JSON.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    user_avatar = "/path/to/file.png" # str |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Avatar
        api_response = api_instance.update_user_settings_avatar(user_id=user_id, user_avatar=user_avatar)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->update_user_settings_avatar: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **user_avatar** | **str**|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

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

# **update_user_settings_desktop_background**
> Model200Success update_user_settings_desktop_background()

Update Desktop Background

This endpoint updates your desktop background image that is used in the Virtual Desktop persona. Expects multipart form data as the request format, not JSON.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
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
    api_instance = users_api.UsersApi(api_client)
    user_id = 1 # int | ID of User (Only available to the master tenant.)  Default is the current user (optional)
    user_desktop_background = "/path/to/file.png" # str |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Desktop Background
        api_response = api_instance.update_user_settings_desktop_background(user_id=user_id, user_desktop_background=user_desktop_background)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->update_user_settings_desktop_background: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **user_desktop_background** | **str**|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

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

# **update_users**
> AddUserTenant200Response update_users(id)

Update user

Update an existing user.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.update_users_request import UpdateUsersRequest
from openapi_client.model.add_user_tenant200_response import AddUserTenant200Response
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
    api_instance = users_api.UsersApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_users_request = UpdateUsersRequest(
        user=UpdateUsersRequestUser(
            first_name="first_name_example",
            last_name="last_name_example",
            username="username_example",
            email="email_example",
            password="password_example",
            roles=[
                ReferenceObject(
                    id=1,
                    name="name_example",
                    code="code_example",
                ),
            ],
        ),
    ) # UpdateUsersRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update user
        api_response = api_instance.update_users(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->update_users: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update user
        api_response = api_instance.update_users(id, update_users_request=update_users_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->update_users: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_users_request** | [**UpdateUsersRequest**](UpdateUsersRequest.md)|  | [optional]

### Return type

[**AddUserTenant200Response**](AddUserTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **whoami**
> Whoami200Response whoami()

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import users_api
from openapi_client.model.whoami200_response import Whoami200Response
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
    api_instance = users_api.UsersApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves information about current user roles and permissions
        api_response = api_instance.whoami()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UsersApi->whoami: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**Whoami200Response**](Whoami200Response.md)

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

