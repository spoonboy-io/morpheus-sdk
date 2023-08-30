# openapi_client.RolesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_roles**](RolesApi.md#add_roles) | **POST** /api/roles | Create role
[**delete_role**](RolesApi.md#delete_role) | **DELETE** /api/roles/{id} | Delete role
[**get_role**](RolesApi.md#get_role) | **GET** /api/roles/{id} | Get role
[**list_roles**](RolesApi.md#list_roles) | **GET** /api/roles | List roles
[**update_role**](RolesApi.md#update_role) | **PUT** /api/roles/{id} | Update role
[**update_role_blueprint_access**](RolesApi.md#update_role_blueprint_access) | **PUT** /api/roles/{id}/update-blueprint | Customizing Blueprint Access
[**update_role_catalog_item_type_access**](RolesApi.md#update_role_catalog_item_type_access) | **PUT** /api/roles/{id}/update-catalog-item-type | Customizing Catalog Item Type Access
[**update_role_cloud_access**](RolesApi.md#update_role_cloud_access) | **PUT** /api/roles/{id}/update-cloud | Customizing Cloud Access
[**update_role_group_access**](RolesApi.md#update_role_group_access) | **PUT** /api/roles/{id}/update-group | Customizing Group Access
[**update_role_instance_type_access**](RolesApi.md#update_role_instance_type_access) | **PUT** /api/roles/{id}/update-instance-type | Customizing Instance Type Access
[**update_role_permission**](RolesApi.md#update_role_permission) | **PUT** /api/roles/{id}/update-permission | Updating Role Permissions
[**update_role_persona_access**](RolesApi.md#update_role_persona_access) | **PUT** /api/roles/{id}/update-persona | Customizing Persona Access
[**update_role_report_type_access**](RolesApi.md#update_role_report_type_access) | **PUT** /api/roles/{id}/update-report-type | Customizing Report Type Access
[**update_role_task_access**](RolesApi.md#update_role_task_access) | **PUT** /api/roles/{id}/update-task | Customizing Task Access
[**update_role_vdi_pool_access**](RolesApi.md#update_role_vdi_pool_access) | **PUT** /api/roles/{id}/update-vdi-pool | Customizing VDI Pool Access
[**update_role_workflow_access**](RolesApi.md#update_role_workflow_access) | **PUT** /api/roles/{id}/update-task-set | Customizing Workflow Access


# **add_roles**
> AddRoles200Response add_roles()

Create role

Create a new role.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.add_roles200_response import AddRoles200Response
from openapi_client.model.add_roles_request import AddRolesRequest
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
    api_instance = roles_api.RolesApi(api_client)
    include_default_access = True # bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)
    add_roles_request = AddRolesRequest(
        role=AddRolesRequestRole(
            authority="authority_example",
            description="description_example",
            role_type="user",
            base_role_id=1,
            multitenant=False,
            multitenant_locked=False,
            default_persona="standard",
            permissions=[
                AddRolesRequestRolePermissionsInner(
                    code="account-usage",
                    access="full",
                ),
            ],
            global_site_access="default",
            sites=[
                AddRolesRequestRoleSitesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_zone_access="default",
            zones=[
                AddRolesRequestRoleZonesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_instance_type_access="full",
            instance_types=[
                AddRolesRequestRoleInstanceTypesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_app_template_access="full",
            app_templates=[
                AddRolesRequestRoleAppTemplatesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_catalog_item_type_access="full",
            catalog_item_types=[
                AddRolesRequestRoleCatalogItemTypesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_persona_access="full",
            personas=[
                AddRolesRequestRolePersonasInner(
                    code="standard",
                    access="default",
                ),
            ],
            global_vdi_pool_access="full",
            vdi_pools=[
                AddRolesRequestRoleVdiPoolsInner(
                    id=1,
                    access="default",
                ),
            ],
            global_report_type_access="full",
            report_types=[
                AddRolesRequestRoleReportTypesInner(
                    code="code_example",
                    access="default",
                ),
            ],
            global_task_access="full",
            tasks=[
                AddRolesRequestRoleTasksInner(
                    id=1,
                    access="default",
                ),
            ],
            global_task_set_access="full",
            task_sets=[
                AddRolesRequestRoleTaskSetsInner(
                    id=1,
                    access="default",
                ),
            ],
            owner=1,
        ),
    ) # AddRolesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create role
        api_response = api_instance.add_roles(include_default_access=include_default_access, add_roles_request=add_roles_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->add_roles: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include_default_access** | **bool**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | [optional]
 **add_roles_request** | [**AddRolesRequest**](AddRolesRequest.md)|  | [optional]

### Return type

[**AddRoles200Response**](AddRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Role Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_role**
> Model200Success delete_role(id)

Delete role

Delete an existing role. A role cannot be deleted while it is still in use.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete role
        api_response = api_instance.delete_role(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->delete_role: %s\n" % e)
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
**200** | Role Object |  -  |
**400** | Role still has managed resources |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_role**
> Role get_role(id)

Get role

Get details about a role

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.role import Role
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    include_default_access = True # bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get role
        api_response = api_instance.get_role(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->get_role: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get role
        api_response = api_instance.get_role(id, include_default_access=include_default_access)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->get_role: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **include_default_access** | **bool**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | [optional]

### Return type

[**Role**](Role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Role Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_roles**
> ListRoles200Response list_roles()

List roles

Get a list of roles.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.list_roles200_response import ListRoles200Response
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
    api_instance = roles_api.RolesApi(api_client)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    authority = "authority_example" # str | Filter by authority (optional)
    role_type = "user" # str | Filter by roleType (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List roles
        api_response = api_instance.list_roles(phrase=phrase, max=max, offset=offset, sort=sort, direction=direction, authority=authority, role_type=role_type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->list_roles: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **authority** | **str**| Filter by authority | [optional]
 **role_type** | **str**| Filter by roleType | [optional]

### Return type

[**ListRoles200Response**](ListRoles200Response.md)

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

# **update_role**
> AddRoles200Response update_role(id)

Update role

Update an existing role.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.add_roles200_response import AddRoles200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_role_request import UpdateRoleRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    include_default_access = True # bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)
    update_role_request = UpdateRoleRequest(
        role=UpdateRoleRequestRole(
            authority="authority_example",
            description="description_example",
            default_persona="standard",
            permissions=[
                AddRolesRequestRolePermissionsInner(
                    code="account-usage",
                    access="full",
                ),
            ],
            global_site_access="full",
            sites=[
                AddRolesRequestRoleSitesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_zone_access="full",
            zones=[
                AddRolesRequestRoleZonesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_instance_type_access="full",
            instance_types=[
                AddRolesRequestRoleInstanceTypesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_app_template_access="full",
            app_templates=[
                AddRolesRequestRoleAppTemplatesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_catalog_item_type_access="full",
            catalog_item_types=[
                AddRolesRequestRoleCatalogItemTypesInner(
                    id=1,
                    access="default",
                ),
            ],
            global_persona_access="full",
            personas=[
                AddRolesRequestRolePersonasInner(
                    code="standard",
                    access="default",
                ),
            ],
            global_vdi_pool_access="full",
            vdi_pools=[
                AddRolesRequestRoleVdiPoolsInner(
                    id=1,
                    access="default",
                ),
            ],
            global_report_type_access="full",
            report_types=[
                AddRolesRequestRoleReportTypesInner(
                    code="code_example",
                    access="default",
                ),
            ],
            global_task_access="full",
            tasks=[
                AddRolesRequestRoleTasksInner(
                    id=1,
                    access="default",
                ),
            ],
            global_task_set_access="full",
            task_sets=[
                AddRolesRequestRoleTaskSetsInner(
                    id=1,
                    access="default",
                ),
            ],
        ),
    ) # UpdateRoleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update role
        api_response = api_instance.update_role(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update role
        api_response = api_instance.update_role(id, include_default_access=include_default_access, update_role_request=update_role_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **include_default_access** | **bool**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | [optional]
 **update_role_request** | [**UpdateRoleRequest**](UpdateRoleRequest.md)|  | [optional]

### Return type

[**AddRoles200Response**](AddRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Role Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_role_blueprint_access**
> UpdateRoleBlueprintAccess200Response update_role_blueprint_access(id)

Customizing Blueprint Access

Customizing Blueprint Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.update_role_blueprint_access_request import UpdateRoleBlueprintAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_blueprint_access_request = UpdateRoleBlueprintAccessRequest(None) # UpdateRoleBlueprintAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Blueprint Access
        api_response = api_instance.update_role_blueprint_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_blueprint_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Blueprint Access
        api_response = api_instance.update_role_blueprint_access(id, update_role_blueprint_access_request=update_role_blueprint_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_blueprint_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_blueprint_access_request** | [**UpdateRoleBlueprintAccessRequest**](UpdateRoleBlueprintAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_catalog_item_type_access**
> UpdateRoleBlueprintAccess200Response update_role_catalog_item_type_access(id)

Customizing Catalog Item Type Access

Customizing Catalog Item Type Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_catalog_item_type_access_request import UpdateRoleCatalogItemTypeAccessRequest
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_catalog_item_type_access_request = UpdateRoleCatalogItemTypeAccessRequest(None) # UpdateRoleCatalogItemTypeAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Catalog Item Type Access
        api_response = api_instance.update_role_catalog_item_type_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_catalog_item_type_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Catalog Item Type Access
        api_response = api_instance.update_role_catalog_item_type_access(id, update_role_catalog_item_type_access_request=update_role_catalog_item_type_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_catalog_item_type_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_catalog_item_type_access_request** | [**UpdateRoleCatalogItemTypeAccessRequest**](UpdateRoleCatalogItemTypeAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_cloud_access**
> UpdateRoleBlueprintAccess200Response update_role_cloud_access(id)

Customizing Cloud Access

Customizing Cloud Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.update_role_cloud_access_request import UpdateRoleCloudAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_cloud_access_request = UpdateRoleCloudAccessRequest(None) # UpdateRoleCloudAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Cloud Access
        api_response = api_instance.update_role_cloud_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_cloud_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Cloud Access
        api_response = api_instance.update_role_cloud_access(id, update_role_cloud_access_request=update_role_cloud_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_cloud_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_cloud_access_request** | [**UpdateRoleCloudAccessRequest**](UpdateRoleCloudAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_group_access**
> UpdateRoleBlueprintAccess200Response update_role_group_access(id)

Customizing Group Access

Customizing Group Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_group_access_request import UpdateRoleGroupAccessRequest
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_group_access_request = UpdateRoleGroupAccessRequest(None) # UpdateRoleGroupAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Group Access
        api_response = api_instance.update_role_group_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_group_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Group Access
        api_response = api_instance.update_role_group_access(id, update_role_group_access_request=update_role_group_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_group_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_group_access_request** | [**UpdateRoleGroupAccessRequest**](UpdateRoleGroupAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_instance_type_access**
> UpdateRoleBlueprintAccess200Response update_role_instance_type_access(id)

Customizing Instance Type Access

Customizing Instance Type Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.update_role_instance_type_access_request import UpdateRoleInstanceTypeAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_instance_type_access_request = UpdateRoleInstanceTypeAccessRequest(None) # UpdateRoleInstanceTypeAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Instance Type Access
        api_response = api_instance.update_role_instance_type_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_instance_type_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Instance Type Access
        api_response = api_instance.update_role_instance_type_access(id, update_role_instance_type_access_request=update_role_instance_type_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_instance_type_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_instance_type_access_request** | [**UpdateRoleInstanceTypeAccessRequest**](UpdateRoleInstanceTypeAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_permission**
> UpdateRoleBlueprintAccess200Response update_role_permission(id)

Updating Role Permissions

Update a feature permission or default permission category (group, cloud, persona, ect.)

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.update_role_permission_request import UpdateRolePermissionRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_permission_request = UpdateRolePermissionRequest(None) # UpdateRolePermissionRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating Role Permissions
        api_response = api_instance.update_role_permission(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_permission: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating Role Permissions
        api_response = api_instance.update_role_permission(id, update_role_permission_request=update_role_permission_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_permission: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_permission_request** | [**UpdateRolePermissionRequest**](UpdateRolePermissionRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_persona_access**
> UpdateRoleBlueprintAccess200Response update_role_persona_access(id)

Customizing Persona Access

Customizing Persona Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.update_role_persona_access_request import UpdateRolePersonaAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_persona_access_request = UpdateRolePersonaAccessRequest(None) # UpdateRolePersonaAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Persona Access
        api_response = api_instance.update_role_persona_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_persona_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Persona Access
        api_response = api_instance.update_role_persona_access(id, update_role_persona_access_request=update_role_persona_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_persona_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_persona_access_request** | [**UpdateRolePersonaAccessRequest**](UpdateRolePersonaAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_report_type_access**
> UpdateRoleBlueprintAccess200Response update_role_report_type_access(id)

Customizing Report Type Access

Customizing Report Type Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_report_type_access_request import UpdateRoleReportTypeAccessRequest
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_report_type_access_request = UpdateRoleReportTypeAccessRequest(None) # UpdateRoleReportTypeAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Report Type Access
        api_response = api_instance.update_role_report_type_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_report_type_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Report Type Access
        api_response = api_instance.update_role_report_type_access(id, update_role_report_type_access_request=update_role_report_type_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_report_type_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_report_type_access_request** | [**UpdateRoleReportTypeAccessRequest**](UpdateRoleReportTypeAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_task_access**
> UpdateRoleBlueprintAccess200Response update_role_task_access(id)

Customizing Task Access

Customizing Task Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.update_role_task_access_request import UpdateRoleTaskAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_task_access_request = UpdateRoleTaskAccessRequest(None) # UpdateRoleTaskAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Task Access
        api_response = api_instance.update_role_task_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_task_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Task Access
        api_response = api_instance.update_role_task_access(id, update_role_task_access_request=update_role_task_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_task_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_task_access_request** | [**UpdateRoleTaskAccessRequest**](UpdateRoleTaskAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_vdi_pool_access**
> UpdateRoleBlueprintAccess200Response update_role_vdi_pool_access(id)

Customizing VDI Pool Access

Customizing VDI Pool Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_role_vdi_pool_access_request import UpdateRoleVDIPoolAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_vdi_pool_access_request = UpdateRoleVDIPoolAccessRequest(
        vdi_pool_id=1,
        access="full",
    ) # UpdateRoleVDIPoolAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing VDI Pool Access
        api_response = api_instance.update_role_vdi_pool_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_vdi_pool_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing VDI Pool Access
        api_response = api_instance.update_role_vdi_pool_access(id, update_role_vdi_pool_access_request=update_role_vdi_pool_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_vdi_pool_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_vdi_pool_access_request** | [**UpdateRoleVDIPoolAccessRequest**](UpdateRoleVDIPoolAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

# **update_role_workflow_access**
> UpdateRoleBlueprintAccess200Response update_role_workflow_access(id)

Customizing Workflow Access

Customizing Workflow Access

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import roles_api
from openapi_client.model.update_role_blueprint_access200_response import UpdateRoleBlueprintAccess200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_role_workflow_access_request import UpdateRoleWorkflowAccessRequest
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
    api_instance = roles_api.RolesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_role_workflow_access_request = UpdateRoleWorkflowAccessRequest(None) # UpdateRoleWorkflowAccessRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Customizing Workflow Access
        api_response = api_instance.update_role_workflow_access(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_workflow_access: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Customizing Workflow Access
        api_response = api_instance.update_role_workflow_access(id, update_role_workflow_access_request=update_role_workflow_access_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling RolesApi->update_role_workflow_access: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_role_workflow_access_request** | [**UpdateRoleWorkflowAccessRequest**](UpdateRoleWorkflowAccessRequest.md)|  | [optional]

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

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

