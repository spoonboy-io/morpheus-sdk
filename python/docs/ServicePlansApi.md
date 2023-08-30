# openapi_client.ServicePlansApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**activate_service_plans**](ServicePlansApi.md#activate_service_plans) | **PUT** /api/service-plans/{id}/activate | Activates a Service Plan
[**add_service_plans**](ServicePlansApi.md#add_service_plans) | **POST** /api/service-plans | Creates a Service Plan
[**deactivate_service_plans**](ServicePlansApi.md#deactivate_service_plans) | **PUT** /api/service-plans/{id}/deactivate | Deactivates a Service Plan
[**get_service_plans**](ServicePlansApi.md#get_service_plans) | **GET** /api/service-plans/{id} | Retrieves a Specific Service Plan
[**list_service_plans**](ServicePlansApi.md#list_service_plans) | **GET** /api/service-plans | Retrieves all Service Plans
[**remove_service_plans**](ServicePlansApi.md#remove_service_plans) | **DELETE** /api/service-plans/{id} | Deletes a Service Plan
[**update_service_plans**](ServicePlansApi.md#update_service_plans) | **PUT** /api/service-plans/{id} | Updates a Service Plan


# **activate_service_plans**
> Model200Success activate_service_plans(id)

Activates a Service Plan

Activates a service plan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Activates a Service Plan
        api_response = api_instance.activate_service_plans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->activate_service_plans: %s\n" % e)
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

# **add_service_plans**
> AddServicePlans200Response add_service_plans()

Creates a Service Plan

Creates a service plan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_service_plans_request import AddServicePlansRequest
from openapi_client.model.add_service_plans200_response import AddServicePlans200Response
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    add_service_plans_request = AddServicePlansRequest(
        service_plan=AddServicePlansRequestServicePlan(
            name="name_example",
            code="code_example",
            description="description_example",
            editable=True,
            max_storage=3.14,
            max_memory=3.14,
            max_cores=3.14,
            max_disks=3.14,
            provision_type=[
                AddServicePlansRequestServicePlanProvisionTypeInner(
                    id=1,
                ),
            ],
            custom_cores=False,
            custom_max_storage=False,
            custom_max_data_storage=False,
            custom_max_memory=False,
            add_volumes=False,
            sort_order=3.14,
            price_sets=[
                1,
            ],
            config=AddServicePlansRequestServicePlanConfig(
                storage_size_type="gb",
                memory_size_type="mb",
                ranges=AddServicePlansRequestServicePlanConfigRanges(
                    min_storage="min_storage_example",
                    max_storage="max_storage_example",
                    min_per_disk_size="min_per_disk_size_example",
                    max_per_disk_size="max_per_disk_size_example",
                    min_memory=1,
                    max_memory=1,
                    min_cores="min_cores_example",
                    max_cores="max_cores_example",
                    min_sockets="min_sockets_example",
                    max_sockets="max_sockets_example",
                    min_cores_per_socket="min_cores_per_socket_example",
                    max_cores_per_socket="max_cores_per_socket_example",
                ),
            ),
        ),
    ) # AddServicePlansRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Service Plan
        api_response = api_instance.add_service_plans(add_service_plans_request=add_service_plans_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->add_service_plans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_service_plans_request** | [**AddServicePlansRequest**](AddServicePlansRequest.md)|  | [optional]

### Return type

[**AddServicePlans200Response**](AddServicePlans200Response.md)

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

# **deactivate_service_plans**
> Model200Success deactivate_service_plans(id)

Deactivates a Service Plan

Deactivates a service plan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deactivates a Service Plan
        api_response = api_instance.deactivate_service_plans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->deactivate_service_plans: %s\n" % e)
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

# **get_service_plans**
> AddServicePlans200ResponseAllOf get_service_plans(id)

Retrieves a Specific Service Plan

Retrieves a specific service plan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
from openapi_client.model.add_service_plans200_response_all_of import AddServicePlans200ResponseAllOf
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Service Plan
        api_response = api_instance.get_service_plans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->get_service_plans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddServicePlans200ResponseAllOf**](AddServicePlans200ResponseAllOf.md)

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

# **list_service_plans**
> ListServicePlans200Response list_service_plans()

Retrieves all Service Plans

Retrieves all service plans. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
from openapi_client.model.list_service_plans200_response import ListServicePlans200Response
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    include_zones = False # bool | Indicates including zones in the service plan response payload (optional) if omitted the server will use the default value of False
    provision_type_id = 22 # int | Provision type filter, restricts query to only load service plans of specified provision type (optional)
    include_inactive = True # bool | If true, include inactive prices in the results (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Service Plans
        api_response = api_instance.list_service_plans(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, include_zones=include_zones, provision_type_id=provision_type_id, include_inactive=include_inactive)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->list_service_plans: %s\n" % e)
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
 **include_zones** | **bool**| Indicates including zones in the service plan response payload | [optional] if omitted the server will use the default value of False
 **provision_type_id** | **int**| Provision type filter, restricts query to only load service plans of specified provision type | [optional]
 **include_inactive** | **bool**| If true, include inactive prices in the results | [optional]

### Return type

[**ListServicePlans200Response**](ListServicePlans200Response.md)

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

# **remove_service_plans**
> Model200Success remove_service_plans(id)

Deletes a Service Plan

Deletes a specified service plan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Service Plan
        api_response = api_instance.remove_service_plans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->remove_service_plans: %s\n" % e)
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

# **update_service_plans**
> AddServicePlans200Response update_service_plans(id)

Updates a Service Plan

Updates a service plan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_plans_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_service_plans200_response import AddServicePlans200Response
from openapi_client.model.update_service_plans_request import UpdateServicePlansRequest
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
    api_instance = service_plans_api.ServicePlansApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_service_plans_request = UpdateServicePlansRequest(
        service_plan=UpdateServicePlansRequestServicePlan(
            name="name_example",
            code="code_example",
            description="description_example",
            editable=True,
            max_storage=3.14,
            max_memory=3.14,
            max_cores=3.14,
            max_disks=3.14,
            provision_type=[
                UpdateServicePlansRequestServicePlanProvisionTypeInner(
                    id=1,
                ),
            ],
            custom_cores=False,
            custom_max_storage=False,
            custom_max_data_storage=False,
            custom_max_memory=False,
            add_volumes=False,
            sort_order=3.14,
            price_sets=[
                1,
            ],
            config=UpdateServicePlansRequestServicePlanConfig(
                storage_size_type="gb",
                memory_size_type="mb",
                ranges=UpdateServicePlansRequestServicePlanConfigRanges(
                    min_storage="min_storage_example",
                    max_storage="max_storage_example",
                    min_memory=1,
                    max_memory=1,
                    min_cores="min_cores_example",
                    max_cores="max_cores_example",
                    min_sockets="min_sockets_example",
                    max_sockets="max_sockets_example",
                    min_cores_per_socket="min_cores_per_socket_example",
                    max_cores_per_socket="max_cores_per_socket_example",
                ),
            ),
        ),
    ) # UpdateServicePlansRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Service Plan
        api_response = api_instance.update_service_plans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->update_service_plans: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Service Plan
        api_response = api_instance.update_service_plans(id, update_service_plans_request=update_service_plans_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServicePlansApi->update_service_plans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_service_plans_request** | [**UpdateServicePlansRequest**](UpdateServicePlansRequest.md)|  | [optional]

### Return type

[**AddServicePlans200Response**](AddServicePlans200Response.md)

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

