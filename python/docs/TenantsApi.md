# openapi_client.TenantsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_tenant**](TenantsApi.md#add_tenant) | **POST** /api/accounts | Create a Tenant
[**add_user_tenant**](TenantsApi.md#add_user_tenant) | **POST** /api/accounts/{accountId}/users | Create a User For a Tenant
[**create_tenant_subtenant_group**](TenantsApi.md#create_tenant_subtenant_group) | **POST** /api/accounts/{accountId}/groups | Create a Group for Subtenant
[**get_tenant**](TenantsApi.md#get_tenant) | **GET** /api/accounts/{id} | Get tenant
[**get_tenant_subtenant_group**](TenantsApi.md#get_tenant_subtenant_group) | **GET** /api/accounts/{accountId}/groups/{id} | Get a Specific Group for Subtenant
[**list_tenant_subtenant_groups**](TenantsApi.md#list_tenant_subtenant_groups) | **GET** /api/accounts/{accountId}/groups | Get Subtenant Groups
[**list_tenants**](TenantsApi.md#list_tenants) | **GET** /api/accounts | List All Tenants
[**list_tenants_available_roles**](TenantsApi.md#list_tenants_available_roles) | **GET** /api/accounts/available-roles | List available roles for a tenant
[**remove_tenant**](TenantsApi.md#remove_tenant) | **DELETE** /api/accounts/{id} | Delete a Specific Tenant
[**remove_tenant_subtenant_group**](TenantsApi.md#remove_tenant_subtenant_group) | **DELETE** /api/accounts/{accountId}/groups/{id} | Delete a Group for Subtenant
[**update_tenant**](TenantsApi.md#update_tenant) | **PUT** /api/accounts/{id} | Update tenant
[**update_tenant_subtenant_group**](TenantsApi.md#update_tenant_subtenant_group) | **PUT** /api/accounts/{accountId}/groups/{id} | Updating a Group for Subtenant
[**update_tenant_subtenant_group_zones**](TenantsApi.md#update_tenant_subtenant_group_zones) | **PUT** /api/accounts/{accountId}/groups/{id}/update-zones | Updating Group Zones for Subtenant


# **add_tenant**
> AddTenant200Response add_tenant()

Create a Tenant

Create a new tenant. This new account will be a sub-tenant with the master tenant as its parent.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.add_tenant200_response import AddTenant200Response
from openapi_client.model.add_tenant_request import AddTenantRequest
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
    api_instance = tenants_api.TenantsApi(api_client)
    add_tenant_request = AddTenantRequest(
        account=AddTenantRequestAccount(
            name="Test Tenant",
            description="Testing Tenant Creation",
            role=AddTenantRequestAccountRole(
                id=3,
            ),
            subdomain="tt",
            currency=CurrencyCode("USD"),
        ),
    ) # AddTenantRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Tenant
        api_response = api_instance.add_tenant(add_tenant_request=add_tenant_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->add_tenant: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_tenant_request** | [**AddTenantRequest**](AddTenantRequest.md)|  | [optional]

### Return type

[**AddTenant200Response**](AddTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_user_tenant**
> AddUserTenant200Response add_user_tenant(account_id)

Create a User For a Tenant

Create a User For a Tenant.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
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
    try:
        # Create a User For a Tenant
        api_response = api_instance.add_user_tenant(account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->add_user_tenant: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a User For a Tenant
        api_response = api_instance.add_user_tenant(account_id, add_user_tenant_request=add_user_tenant_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->add_user_tenant: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
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

# **create_tenant_subtenant_group**
> CreateTenantSubtenantGroup200Response create_tenant_subtenant_group(account_id)

Create a Group for Subtenant

Create a Group for Subtenant.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.create_tenant_subtenant_group200_response import CreateTenantSubtenantGroup200Response
from openapi_client.model.create_tenant_subtenant_group_request import CreateTenantSubtenantGroupRequest
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
    create_tenant_subtenant_group_request = CreateTenantSubtenantGroupRequest(
        group=CreateTenantSubtenantGroupRequestGroup(
            name="name_example",
            description="description_example",
            code="code_example",
            location="location_example",
        ),
    ) # CreateTenantSubtenantGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Group for Subtenant
        api_response = api_instance.create_tenant_subtenant_group(account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->create_tenant_subtenant_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Group for Subtenant
        api_response = api_instance.create_tenant_subtenant_group(account_id, create_tenant_subtenant_group_request=create_tenant_subtenant_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->create_tenant_subtenant_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **create_tenant_subtenant_group_request** | [**CreateTenantSubtenantGroupRequest**](CreateTenantSubtenantGroupRequest.md)|  | [optional]

### Return type

[**CreateTenantSubtenantGroup200Response**](CreateTenantSubtenantGroup200Response.md)

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

# **get_tenant**
> AddTenant200ResponseAllOf get_tenant(id)

Get tenant

Get details about a tenant

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.add_tenant200_response_all_of import AddTenant200ResponseAllOf
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
    api_instance = tenants_api.TenantsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get tenant
        api_response = api_instance.get_tenant(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->get_tenant: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddTenant200ResponseAllOf**](AddTenant200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_tenant_subtenant_group**
> GetTenantSubtenantGroup200Response get_tenant_subtenant_group(account_id, id)

Get a Specific Group for Subtenant

This endpoint retrieves a specific group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.get_tenant_subtenant_group200_response import GetTenantSubtenantGroup200Response
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Group for Subtenant
        api_response = api_instance.get_tenant_subtenant_group(account_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->get_tenant_subtenant_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetTenantSubtenantGroup200Response**](GetTenantSubtenantGroup200Response.md)

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

# **list_tenant_subtenant_groups**
> ListTenantSubtenantGroups200Response list_tenant_subtenant_groups(account_id)

Get Subtenant Groups

Groups belonging to a subtenant can be managed by the master account.  This endpoint retrieves all groups and a list of zones associated with the group by id. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.list_tenant_subtenant_groups200_response import ListTenantSubtenantGroups200Response
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Subtenant Groups
        api_response = api_instance.list_tenant_subtenant_groups(account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->list_tenant_subtenant_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Subtenant Groups
        api_response = api_instance.list_tenant_subtenant_groups(account_id, phrase=phrase, name=name, last_updated=last_updated)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->list_tenant_subtenant_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

### Return type

[**ListTenantSubtenantGroups200Response**](ListTenantSubtenantGroups200Response.md)

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

# **list_tenants**
> ListTenants200Response list_tenants()

List All Tenants

Get a list of tenants. A tenant is also referred to as an account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.list_tenants200_response import ListTenants200Response
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
    api_instance = tenants_api.TenantsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Tenants
        api_response = api_instance.list_tenants(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, last_updated=last_updated)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->list_tenants: %s\n" % e)
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
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

### Return type

[**ListTenants200Response**](ListTenants200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Tenants |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_tenants_available_roles**
> TenantsAvailableRoles list_tenants_available_roles()

List available roles for a tenant

Get a list of available roles that can be assigned as the default base role for a sub tenant account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.tenants_available_roles import TenantsAvailableRoles
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
    api_instance = tenants_api.TenantsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # List available roles for a tenant
        api_response = api_instance.list_tenants_available_roles()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->list_tenants_available_roles: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**TenantsAvailableRoles**](TenantsAvailableRoles.md)

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

# **remove_tenant**
> Model200Success remove_tenant(id)

Delete a Specific Tenant

Delete an existing tenant. This action is not reversible and will result in the removal of all data pertaining to this tenant as well as potentially any provisioned assets depending on the value of `removeResources`.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.error import Error
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
    api_instance = tenants_api.TenantsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    remove_resources = False # bool | Remove Resources. This will delete all the managed resources in the tenant. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Tenant
        api_response = api_instance.remove_tenant(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->remove_tenant: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Specific Tenant
        api_response = api_instance.remove_tenant(id, remove_resources=remove_resources)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->remove_tenant: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_resources** | **bool**| Remove Resources. This will delete all the managed resources in the tenant. | [optional] if omitted the server will use the default value of False

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
**200** | Tenant Deletion Object |  -  |
**400** | Tenant still has managed resources |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_tenant_subtenant_group**
> Model200Success remove_tenant_subtenant_group(account_id, id)

Delete a Group for Subtenant

If a group has zones or servers still tied to it, a delete action will fail.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Group for Subtenant
        api_response = api_instance.remove_tenant_subtenant_group(account_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->remove_tenant_subtenant_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
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

# **update_tenant**
> UpdateTenant200Response update_tenant(id)

Update tenant

Update an existing tenant.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.update_tenant_request import UpdateTenantRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_tenant200_response import UpdateTenant200Response
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
    api_instance = tenants_api.TenantsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_tenant_request = UpdateTenantRequest(
        account=UpdateTenantRequestAccount(None),
    ) # UpdateTenantRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update tenant
        api_response = api_instance.update_tenant(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->update_tenant: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update tenant
        api_response = api_instance.update_tenant(id, update_tenant_request=update_tenant_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->update_tenant: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_tenant_request** | [**UpdateTenantRequest**](UpdateTenantRequest.md)|  | [optional]

### Return type

[**UpdateTenant200Response**](UpdateTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_tenant_subtenant_group**
> CreateTenantSubtenantGroup200Response update_tenant_subtenant_group(account_id, id)

Updating a Group for Subtenant

Updating a Group for Subtenant.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.create_tenant_subtenant_group200_response import CreateTenantSubtenantGroup200Response
from openapi_client.model.update_tenant_subtenant_group_request import UpdateTenantSubtenantGroupRequest
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
    id = 1 # int | Morpheus ID of the Object being referenced
    update_tenant_subtenant_group_request = UpdateTenantSubtenantGroupRequest(
        group=UpdateTenantSubtenantGroupRequestGroup(
            name="name_example",
            description="description_example",
            code="code_example",
            location="location_example",
        ),
    ) # UpdateTenantSubtenantGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a Group for Subtenant
        api_response = api_instance.update_tenant_subtenant_group(account_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->update_tenant_subtenant_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a Group for Subtenant
        api_response = api_instance.update_tenant_subtenant_group(account_id, id, update_tenant_subtenant_group_request=update_tenant_subtenant_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->update_tenant_subtenant_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_tenant_subtenant_group_request** | [**UpdateTenantSubtenantGroupRequest**](UpdateTenantSubtenantGroupRequest.md)|  | [optional]

### Return type

[**CreateTenantSubtenantGroup200Response**](CreateTenantSubtenantGroup200Response.md)

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

# **update_tenant_subtenant_group_zones**
> Model200Success update_tenant_subtenant_group_zones(account_id, id)

Updating Group Zones for Subtenant

This will update the zones that are assigned to the group. Any zones that are not passed in the zones parameter will be removed from the group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import tenants_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_tenant_subtenant_group_zones_request import UpdateTenantSubtenantGroupZonesRequest
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
    api_instance = tenants_api.TenantsApi(api_client)
    account_id = 7 # int | The ID of the subtenant account
    id = 1 # int | Morpheus ID of the Object being referenced
    update_tenant_subtenant_group_zones_request = UpdateTenantSubtenantGroupZonesRequest(
        group=UpdateTenantSubtenantGroupZonesRequestGroup(
            zones=[
                AddStorageVolumesRequestStorageVolumeStorageServer(
                    id=1,
                ),
            ],
        ),
    ) # UpdateTenantSubtenantGroupZonesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating Group Zones for Subtenant
        api_response = api_instance.update_tenant_subtenant_group_zones(account_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->update_tenant_subtenant_group_zones: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating Group Zones for Subtenant
        api_response = api_instance.update_tenant_subtenant_group_zones(account_id, id, update_tenant_subtenant_group_zones_request=update_tenant_subtenant_group_zones_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TenantsApi->update_tenant_subtenant_group_zones: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| The ID of the subtenant account |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_tenant_subtenant_group_zones_request** | [**UpdateTenantSubtenantGroupZonesRequest**](UpdateTenantSubtenantGroupZonesRequest.md)|  | [optional]

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

