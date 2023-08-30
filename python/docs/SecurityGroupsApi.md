# openapi_client.SecurityGroupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_security_group_locations**](SecurityGroupsApi.md#add_security_group_locations) | **POST** /api/security-groups/{id}/locations | Creates a Security Group Location
[**add_security_group_rules**](SecurityGroupsApi.md#add_security_group_rules) | **POST** /api/security-groups/{id}/rules | Creates a Security Group Rule
[**add_security_groups**](SecurityGroupsApi.md#add_security_groups) | **POST** /api/security-groups | Creates a Security Group
[**get_security_group_rules**](SecurityGroupsApi.md#get_security_group_rules) | **GET** /api/security-groups/{id}/rules/{sgId} | Retrieves a Specific Security Group Rule
[**get_security_groups**](SecurityGroupsApi.md#get_security_groups) | **GET** /api/security-groups/{id} | Retrieves a Specific Security Group
[**list_security_group_rules**](SecurityGroupsApi.md#list_security_group_rules) | **GET** /api/security-groups/{id}/rules | Retrieves all Security Group Rules
[**list_security_groups**](SecurityGroupsApi.md#list_security_groups) | **GET** /api/security-groups | Retrieves all Security Groups
[**remove_security_group_locations**](SecurityGroupsApi.md#remove_security_group_locations) | **DELETE** /api/security-groups/{id}/locations/{locationId} | Deletes a Security Group Location
[**remove_security_group_rules**](SecurityGroupsApi.md#remove_security_group_rules) | **DELETE** /api/security-groups/{id}/rules/{sgId} | Deletes a Security Group Rule
[**remove_security_groups**](SecurityGroupsApi.md#remove_security_groups) | **DELETE** /api/security-groups/{id} | Deletes a Security Group
[**update_security_group_rules**](SecurityGroupsApi.md#update_security_group_rules) | **PUT** /api/security-groups/{id}/rules/{sgId} | Updates a Security Group Rule
[**update_security_groups**](SecurityGroupsApi.md#update_security_groups) | **PUT** /api/security-groups/{id} | Updating a Security Group


# **add_security_group_locations**
> AddSecurityGroupLocations200Response add_security_group_locations(id)

Creates a Security Group Location

Creates a security group location. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.add_security_group_locations200_response import AddSecurityGroupLocations200Response
from openapi_client.model.add_security_group_locations_request import AddSecurityGroupLocationsRequest
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_security_group_locations_request = AddSecurityGroupLocationsRequest(
        security_group_location=AddSecurityGroupLocationsRequestSecurityGroupLocation(
            zone_id=1,
            custom_options=AddSecurityGroupsRequestSecurityGroupCustomOptions(None),
        ),
    ) # AddSecurityGroupLocationsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates a Security Group Location
        api_response = api_instance.add_security_group_locations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->add_security_group_locations: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Security Group Location
        api_response = api_instance.add_security_group_locations(id, add_security_group_locations_request=add_security_group_locations_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->add_security_group_locations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_security_group_locations_request** | [**AddSecurityGroupLocationsRequest**](AddSecurityGroupLocationsRequest.md)|  | [optional]

### Return type

[**AddSecurityGroupLocations200Response**](AddSecurityGroupLocations200Response.md)

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

# **add_security_group_rules**
> AddSecurityGroupRules200Response add_security_group_rules(id)

Creates a Security Group Rule

Creates a security group rule on specified security group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.add_security_group_rules200_response import AddSecurityGroupRules200Response
from openapi_client.model.add_security_group_rules_request import AddSecurityGroupRulesRequest
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_security_group_rules_request = AddSecurityGroupRulesRequest(
        rule=AddSecurityGroupRulesRequestRule(
            name="name_example",
            direction="ingress",
            source_type="cidr",
            source="50.22.10.10/32",
            source_group=AddSecurityGroupRulesRequestRuleSourceGroup(
                id=56496,
            ),
            source_tier=AddSecurityGroupRulesRequestRuleSourceTier(
                id=56496,
            ),
            port_range="55-72",
            protocol="tcp",
            destination_type="cidr",
            destination="50.22.10.10/32",
            destination_group=AddSecurityGroupRulesRequestRuleDestinationGroup(
                id=56496,
            ),
            destination_tier=AddSecurityGroupRulesRequestRuleDestinationTier(
                id=56496,
            ),
            rule_type="customRule",
            policy="accept",
            instance_type_id=54568,
        ),
    ) # AddSecurityGroupRulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates a Security Group Rule
        api_response = api_instance.add_security_group_rules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->add_security_group_rules: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Security Group Rule
        api_response = api_instance.add_security_group_rules(id, add_security_group_rules_request=add_security_group_rules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->add_security_group_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_security_group_rules_request** | [**AddSecurityGroupRulesRequest**](AddSecurityGroupRulesRequest.md)|  | [optional]

### Return type

[**AddSecurityGroupRules200Response**](AddSecurityGroupRules200Response.md)

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

# **add_security_groups**
> AddSecurityGroups200Response add_security_groups()

Creates a Security Group

Creates a security group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.add_security_groups200_response import AddSecurityGroups200Response
from openapi_client.model.add_security_groups_request import AddSecurityGroupsRequest
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    add_security_groups_request = AddSecurityGroupsRequest(
        security_group=AddSecurityGroupsRequestSecurityGroup(
            name="name_example",
            description="description_example",
            zone_id=3,
            active=True,
            custom_options=AddSecurityGroupsRequestSecurityGroupCustomOptions(None),
            tenant_permissions=[
                AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner(
                    accounts=[1,3],
                    can_manage_accounts=[1,3],
                ),
            ],
            resource_permissions=UpdateCloudDatastoresRequestDatastoreResourcePermissions(
                all=True,
                sites=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
                all_plans=True,
                plans=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
            ),
        ),
    ) # AddSecurityGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Security Group
        api_response = api_instance.add_security_groups(add_security_groups_request=add_security_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->add_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_security_groups_request** | [**AddSecurityGroupsRequest**](AddSecurityGroupsRequest.md)|  | [optional]

### Return type

[**AddSecurityGroups200Response**](AddSecurityGroups200Response.md)

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

# **get_security_group_rules**
> AddSecurityGroupRules200ResponseAllOf get_security_group_rules(id, sg_id)

Retrieves a Specific Security Group Rule

Retrieves a specific security group rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.add_security_group_rules200_response_all_of import AddSecurityGroupRules200ResponseAllOf
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    sg_id = 2352 # float | Morpheus ID of the security group rule being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Security Group Rule
        api_response = api_instance.get_security_group_rules(id, sg_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->get_security_group_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **sg_id** | **float**| Morpheus ID of the security group rule being referenced |

### Return type

[**AddSecurityGroupRules200ResponseAllOf**](AddSecurityGroupRules200ResponseAllOf.md)

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

# **get_security_groups**
> GetSecurityGroups200Response get_security_groups(id)

Retrieves a Specific Security Group

Retrieves a specific security group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_security_groups200_response import GetSecurityGroups200Response
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Security Group
        api_response = api_instance.get_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->get_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetSecurityGroups200Response**](GetSecurityGroups200Response.md)

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

# **list_security_group_rules**
> ListSecurityGroupRules200Response list_security_group_rules(id)

Retrieves all Security Group Rules

Retrieves all security group rules for specified security group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_security_group_rules200_response import ListSecurityGroupRules200Response
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    try:
        # Retrieves all Security Group Rules
        api_response = api_instance.list_security_group_rules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->list_security_group_rules: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Security Group Rules
        api_response = api_instance.list_security_group_rules(id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->list_security_group_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListSecurityGroupRules200Response**](ListSecurityGroupRules200Response.md)

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

# **list_security_groups**
> ListSecurityGroups200Response list_security_groups()

Retrieves all Security Groups

This endpoint retrieves all security groups and their JSON encoded configuration attributes. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_security_groups200_response import ListSecurityGroups200Response
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Security Groups
        api_response = api_instance.list_security_groups(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->list_security_groups: %s\n" % e)
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

[**ListSecurityGroups200Response**](ListSecurityGroups200Response.md)

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

# **remove_security_group_locations**
> Model200Success remove_security_group_locations(id, location_id)

Deletes a Security Group Location

Deletes a security group location. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    location_id = 2 # float | The ID of the location

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Security Group Location
        api_response = api_instance.remove_security_group_locations(id, location_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->remove_security_group_locations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **location_id** | **float**| The ID of the location |

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

# **remove_security_group_rules**
> Model200Success remove_security_group_rules(id, sg_id)

Deletes a Security Group Rule

Deletes a security group rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    sg_id = 2352 # float | Morpheus ID of the security group rule being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Security Group Rule
        api_response = api_instance.remove_security_group_rules(id, sg_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->remove_security_group_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **sg_id** | **float**| Morpheus ID of the security group rule being referenced |

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

# **remove_security_groups**
> Model200Success remove_security_groups(id)

Deletes a Security Group

Deletes a specified security group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Security Group
        api_response = api_instance.remove_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->remove_security_groups: %s\n" % e)
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

# **update_security_group_rules**
> AddSecurityGroupRules200Response update_security_group_rules(id, sg_id)

Updates a Security Group Rule

Updates a security group rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.add_security_group_rules200_response import AddSecurityGroupRules200Response
from openapi_client.model.update_security_group_rules_request import UpdateSecurityGroupRulesRequest
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    sg_id = 2352 # float | Morpheus ID of the security group rule being referenced
    update_security_group_rules_request = UpdateSecurityGroupRulesRequest(
        rule=UpdateSecurityGroupRulesRequestRule(
            name="name_example",
            direction="ingress",
            source_type="cidr",
            source="50.22.10.10/32",
            source_group=AddSecurityGroupRulesRequestRuleSourceGroup(
                id=56496,
            ),
            source_tier=AddSecurityGroupRulesRequestRuleSourceTier(
                id=56496,
            ),
            port_range="55-72",
            protocol="tcp",
            destination_type="cidr",
            destination="50.22.10.10/32",
            destination_group=AddSecurityGroupRulesRequestRuleDestinationGroup(
                id=56496,
            ),
            destination_tier=AddSecurityGroupRulesRequestRuleDestinationTier(
                id=56496,
            ),
            rule_type="customRule",
            policy="accept",
            instance_type_id=54568,
        ),
    ) # UpdateSecurityGroupRulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Security Group Rule
        api_response = api_instance.update_security_group_rules(id, sg_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->update_security_group_rules: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Security Group Rule
        api_response = api_instance.update_security_group_rules(id, sg_id, update_security_group_rules_request=update_security_group_rules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->update_security_group_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **sg_id** | **float**| Morpheus ID of the security group rule being referenced |
 **update_security_group_rules_request** | [**UpdateSecurityGroupRulesRequest**](UpdateSecurityGroupRulesRequest.md)|  | [optional]

### Return type

[**AddSecurityGroupRules200Response**](AddSecurityGroupRules200Response.md)

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

# **update_security_groups**
> AddSecurityGroups200Response update_security_groups(id)

Updating a Security Group

Updating a Security Group 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_groups_api
from openapi_client.model.add_security_groups200_response import AddSecurityGroups200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_security_groups_request import UpdateSecurityGroupsRequest
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
    api_instance = security_groups_api.SecurityGroupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_security_groups_request = UpdateSecurityGroupsRequest(
        security_group=UpdateSecurityGroupsRequestSecurityGroup(
            name="name_example",
            description="description_example",
            active=True,
            tenant_permissions=[
                AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner(
                    accounts=[1,3],
                    can_manage_accounts=[1,3],
                ),
            ],
            resource_permissions=UpdateCloudDatastoresRequestDatastoreResourcePermissions(
                all=True,
                sites=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
                all_plans=True,
                plans=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
            ),
        ),
    ) # UpdateSecurityGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a Security Group
        api_response = api_instance.update_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->update_security_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a Security Group
        api_response = api_instance.update_security_groups(id, update_security_groups_request=update_security_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityGroupsApi->update_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_security_groups_request** | [**UpdateSecurityGroupsRequest**](UpdateSecurityGroupsRequest.md)|  | [optional]

### Return type

[**AddSecurityGroups200Response**](AddSecurityGroups200Response.md)

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

