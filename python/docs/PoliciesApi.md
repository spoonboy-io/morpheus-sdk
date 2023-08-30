# openapi_client.PoliciesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_policies**](PoliciesApi.md#add_policies) | **POST** /api/policies | Creates a Policy
[**add_policies_cloud**](PoliciesApi.md#add_policies_cloud) | **POST** /api/zones/{cloudId}/policies | Creates a Policy for a Cloud
[**add_policies_group**](PoliciesApi.md#add_policies_group) | **POST** /api/groups/{groupId}/policies | Creates a Policy for a Group
[**get_policies**](PoliciesApi.md#get_policies) | **GET** /api/policies/{id} | Retrieves a Specific Policy
[**get_policies_cloud**](PoliciesApi.md#get_policies_cloud) | **GET** /api/zones/{cloudId}/policies/{id} | Retrieves a Specific Policy for a Cloud
[**get_policies_group**](PoliciesApi.md#get_policies_group) | **GET** /api/groups/{groupId}/policies/{id} | Retrieves a Specific Policy for a Group
[**list_policies**](PoliciesApi.md#list_policies) | **GET** /api/policies | Retrieves all Policies
[**list_policies_cloud**](PoliciesApi.md#list_policies_cloud) | **GET** /api/zones/{cloudId}/policies | Retrieves Policies for a Cloud
[**list_policies_group**](PoliciesApi.md#list_policies_group) | **GET** /api/groups/{groupId}/policies | Retrieves Policies for a Group
[**list_policy_types**](PoliciesApi.md#list_policy_types) | **GET** /api/policy-types | Retrieves all Policy Types
[**remove_policies**](PoliciesApi.md#remove_policies) | **DELETE** /api/policies/{id} | Deletes a Policy
[**remove_policies_cloud**](PoliciesApi.md#remove_policies_cloud) | **DELETE** /api/zones/{cloudId}/policies/{id} | Deletes a Policy for a Cloud
[**remove_policies_group**](PoliciesApi.md#remove_policies_group) | **DELETE** /api/groups/{groupId}/policies/{id} | Deletes a Policy for a Group
[**update_policies**](PoliciesApi.md#update_policies) | **PUT** /api/policies/{id} | Updates a Policy
[**update_policies_cloud**](PoliciesApi.md#update_policies_cloud) | **PUT** /api/zones/{cloudId}/policies/{id} | Updates a Policy for a Cloud
[**update_policies_group**](PoliciesApi.md#update_policies_group) | **PUT** /api/groups/{groupId}/policies/{id} | Updates a Policy for a Group


# **add_policies**
> AddPolicies200Response add_policies()

Creates a Policy

Creates a policy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.add_policies_request import AddPoliciesRequest
from openapi_client.model.add_policies200_response import AddPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    add_policies_request = AddPoliciesRequest(
        policy=PolicyCreate(
            name="Sample Policy",
            description="description_example",
            policy_type=PolicyCreatePolicyType(
                code="deleteApproval",
            ),
            config=PolicyCreateConfig(None),
            enabled=True,
            ref_type="ref_type_example",
            ref_id=1,
            accounts=[
                1,
            ],
            each_user=True,
        ),
    ) # AddPoliciesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Policy
        api_response = api_instance.add_policies(add_policies_request=add_policies_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->add_policies: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_policies_request** | [**AddPoliciesRequest**](AddPoliciesRequest.md)|  | [optional]

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

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

# **add_policies_cloud**
> AddPolicies200Response add_policies_cloud(cloud_id)

Creates a Policy for a Cloud

Creates a policy for a Cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.add_policies200_response import AddPolicies200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_policies_cloud_request import AddPoliciesCloudRequest
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
    api_instance = policies_api.PoliciesApi(api_client)
    cloud_id = 7 # int | The ID of the cloud
    add_policies_cloud_request = AddPoliciesCloudRequest(
        policy=PolicyCloudCreate(
            name="Sample Policy",
            description="description_example",
            policy_type=PolicyCloudCreatePolicyType(
                code="Approve Delete",
                config=PolicyGroupCreatePolicyTypeConfig(None),
                enabled=True,
                ref_type="ComputeZone",
                ref_id=1,
                accounts=[
                    1,
                ],
                each_user=True,
            ),
        ),
    ) # AddPoliciesCloudRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates a Policy for a Cloud
        api_response = api_instance.add_policies_cloud(cloud_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->add_policies_cloud: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Policy for a Cloud
        api_response = api_instance.add_policies_cloud(cloud_id, add_policies_cloud_request=add_policies_cloud_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->add_policies_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **add_policies_cloud_request** | [**AddPoliciesCloudRequest**](AddPoliciesCloudRequest.md)|  | [optional]

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

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

# **add_policies_group**
> AddPolicies200Response add_policies_group(group_id)

Creates a Policy for a Group

Creates a policy for a Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.add_policies200_response import AddPolicies200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_policies_group_request import AddPoliciesGroupRequest
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
    api_instance = policies_api.PoliciesApi(api_client)
    group_id = 7 # int | The ID of the group
    add_policies_group_request = AddPoliciesGroupRequest(
        policy=PolicyGroupCreate(
            name="Sample Policy",
            description="description_example",
            policy_type=PolicyGroupCreatePolicyType(
                code="Approve Delete",
                config=PolicyGroupCreatePolicyTypeConfig(None),
                enabled=True,
                ref_type="ComputeSite",
                ref_id=1,
                accounts=[
                    1,
                ],
                each_user=True,
            ),
        ),
    ) # AddPoliciesGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates a Policy for a Group
        api_response = api_instance.add_policies_group(group_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->add_policies_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Policy for a Group
        api_response = api_instance.add_policies_group(group_id, add_policies_group_request=add_policies_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->add_policies_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **add_policies_group_request** | [**AddPoliciesGroupRequest**](AddPoliciesGroupRequest.md)|  | [optional]

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

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

# **get_policies**
> GetPolicies200Response get_policies(id)

Retrieves a Specific Policy

Retrieves a specific policy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.get_policies200_response import GetPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Policy
        api_response = api_instance.get_policies(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->get_policies: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPolicies200Response**](GetPolicies200Response.md)

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

# **get_policies_cloud**
> GetPolicies200Response get_policies_cloud(cloud_id, id)

Retrieves a Specific Policy for a Cloud

Retrieves a specific policy for a Cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.get_policies200_response import GetPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    cloud_id = 7 # int | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Policy for a Cloud
        api_response = api_instance.get_policies_cloud(cloud_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->get_policies_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPolicies200Response**](GetPolicies200Response.md)

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

# **get_policies_group**
> GetPolicies200Response get_policies_group(group_id, id)

Retrieves a Specific Policy for a Group

Retrieves a specific policy for a Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.get_policies200_response import GetPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    group_id = 7 # int | The ID of the group
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Policy for a Group
        api_response = api_instance.get_policies_group(group_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->get_policies_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPolicies200Response**](GetPolicies200Response.md)

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

# **list_policies**
> ListPolicies200Response list_policies()

Retrieves all Policies

Retrieves all policies. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.list_policies200_response import ListPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Policies
        api_response = api_instance.list_policies(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->list_policies: %s\n" % e)
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

[**ListPolicies200Response**](ListPolicies200Response.md)

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

# **list_policies_cloud**
> ListPolicies200Response list_policies_cloud(cloud_id)

Retrieves Policies for a Cloud

Retrieves policies for a specific cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.list_policies200_response import ListPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    cloud_id = 7 # int | The ID of the cloud
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    try:
        # Retrieves Policies for a Cloud
        api_response = api_instance.list_policies_cloud(cloud_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->list_policies_cloud: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Policies for a Cloud
        api_response = api_instance.list_policies_cloud(cloud_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->list_policies_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListPolicies200Response**](ListPolicies200Response.md)

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

# **list_policies_group**
> ListPolicies200Response list_policies_group(group_id)

Retrieves Policies for a Group

Retrieves policies for a specific group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.list_policies200_response import ListPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    group_id = 7 # int | The ID of the group
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    try:
        # Retrieves Policies for a Group
        api_response = api_instance.list_policies_group(group_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->list_policies_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Policies for a Group
        api_response = api_instance.list_policies_group(group_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->list_policies_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListPolicies200Response**](ListPolicies200Response.md)

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

# **list_policy_types**
> ListPolicyTypes200Response list_policy_types()

Retrieves all Policy Types

Retrieves all Policy Types 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_policy_types200_response import ListPolicyTypes200Response
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
    api_instance = policies_api.PoliciesApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves all Policy Types
        api_response = api_instance.list_policy_types()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->list_policy_types: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListPolicyTypes200Response**](ListPolicyTypes200Response.md)

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

# **remove_policies**
> Model200Success remove_policies(id)

Deletes a Policy

Deletes a specified policy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
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
    api_instance = policies_api.PoliciesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Policy
        api_response = api_instance.remove_policies(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->remove_policies: %s\n" % e)
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

# **remove_policies_cloud**
> Model200Success remove_policies_cloud(cloud_id, id)

Deletes a Policy for a Cloud

Deletes a specified policy for a Cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
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
    api_instance = policies_api.PoliciesApi(api_client)
    cloud_id = 7 # int | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Policy for a Cloud
        api_response = api_instance.remove_policies_cloud(cloud_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->remove_policies_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
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

# **remove_policies_group**
> Model200Success remove_policies_group(group_id, id)

Deletes a Policy for a Group

Deletes a specified policy for a Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
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
    api_instance = policies_api.PoliciesApi(api_client)
    group_id = 7 # int | The ID of the group
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Policy for a Group
        api_response = api_instance.remove_policies_group(group_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->remove_policies_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
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

# **update_policies**
> AddPolicies200Response update_policies(id)

Updates a Policy

Updates a policy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.update_policies_request import UpdatePoliciesRequest
from openapi_client.model.add_policies200_response import AddPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_policies_request = UpdatePoliciesRequest(
        policy=PolicyUpdate(
            name="Sample Policy",
            description="description_example",
            config=PolicyUpdateConfig(None),
            enabled=True,
            ref_type="ComputeSite",
            ref_id=1,
            accounts=[
                1,
            ],
            each_user=True,
        ),
    ) # UpdatePoliciesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Policy
        api_response = api_instance.update_policies(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->update_policies: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Policy
        api_response = api_instance.update_policies(id, update_policies_request=update_policies_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->update_policies: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_policies_request** | [**UpdatePoliciesRequest**](UpdatePoliciesRequest.md)|  | [optional]

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

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

# **update_policies_cloud**
> AddPolicies200Response update_policies_cloud(cloud_id, id)

Updates a Policy for a Cloud

Updates a policy for a Cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.add_policies200_response import AddPolicies200Response
from openapi_client.model.update_policies_cloud_request import UpdatePoliciesCloudRequest
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
    api_instance = policies_api.PoliciesApi(api_client)
    cloud_id = 7 # int | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced
    update_policies_cloud_request = UpdatePoliciesCloudRequest(
        policy=PolicyCloudUpdate(
            name="Sample Policy",
            description="description_example",
            policy_type=PolicyCloudUpdatePolicyType(
                code="Approve Delete",
                config=PolicyGroupUpdatePolicyTypeConfig(None),
                enabled=True,
                ref_type="ComputeZone",
                ref_id=1,
                accounts=[
                    1,
                ],
                each_user=True,
            ),
        ),
    ) # UpdatePoliciesCloudRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Policy for a Cloud
        api_response = api_instance.update_policies_cloud(cloud_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->update_policies_cloud: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Policy for a Cloud
        api_response = api_instance.update_policies_cloud(cloud_id, id, update_policies_cloud_request=update_policies_cloud_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->update_policies_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_policies_cloud_request** | [**UpdatePoliciesCloudRequest**](UpdatePoliciesCloudRequest.md)|  | [optional]

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

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

# **update_policies_group**
> AddPolicies200Response update_policies_group(group_id, id)

Updates a Policy for a Group

Updates a policy for a Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import policies_api
from openapi_client.model.update_policies_group_request import UpdatePoliciesGroupRequest
from openapi_client.model.add_policies200_response import AddPolicies200Response
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
    api_instance = policies_api.PoliciesApi(api_client)
    group_id = 7 # int | The ID of the group
    id = 1 # int | Morpheus ID of the Object being referenced
    update_policies_group_request = UpdatePoliciesGroupRequest(
        policy=PolicyGroupUpdate(
            name="Sample Policy",
            description="description_example",
            policy_type=PolicyGroupUpdatePolicyType(
                code="Approve Delete",
                config=PolicyGroupUpdatePolicyTypeConfig(None),
                enabled=True,
                ref_type="ComputeSite",
                ref_id=1,
                accounts=[
                    1,
                ],
                each_user=True,
            ),
        ),
    ) # UpdatePoliciesGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Policy for a Group
        api_response = api_instance.update_policies_group(group_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->update_policies_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Policy for a Group
        api_response = api_instance.update_policies_group(group_id, id, update_policies_group_request=update_policies_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PoliciesApi->update_policies_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_policies_group_request** | [**UpdatePoliciesGroupRequest**](UpdatePoliciesGroupRequest.md)|  | [optional]

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

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

