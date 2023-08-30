# openapi_client.IntegrationsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_integration_snow_objects**](IntegrationsApi.md#add_integration_snow_objects) | **POST** /api/integrations/{id}/objects | Creates an Exposed ServiceNow Catalog Item
[**add_integrations**](IntegrationsApi.md#add_integrations) | **POST** /api/integrations | Creates an Integration
[**get_integration_inventory**](IntegrationsApi.md#get_integration_inventory) | **GET** /api/integrations/{id}/inventory/{inventoryId} | Get a Specific Integration Inventory
[**get_integration_objects**](IntegrationsApi.md#get_integration_objects) | **GET** /api/integrations/{id}/objects/{objectId} | Get a Specific ServiceNow Integration Object
[**get_integration_type_option_types**](IntegrationsApi.md#get_integration_type_option_types) | **GET** /api/integration-types/{id}/option-types | Retrieves a Option Types for a Specific Integration Type
[**get_integration_types**](IntegrationsApi.md#get_integration_types) | **GET** /api/integration-types/{id} | Retrieves a Specific Integration Type
[**get_integrations**](IntegrationsApi.md#get_integrations) | **GET** /api/integrations/{id} | Retrieves a Specific Integration
[**list_integration_inventory**](IntegrationsApi.md#list_integration_inventory) | **GET** /api/integrations/{id}/inventory | Get All Integration Inventory
[**list_integration_objects**](IntegrationsApi.md#list_integration_objects) | **GET** /api/integrations/{id}/objects | Get ServiceNow Integration Objects
[**list_integration_types**](IntegrationsApi.md#list_integration_types) | **GET** /api/integration-types | Retrieves all Integration Types
[**list_integrations**](IntegrationsApi.md#list_integrations) | **GET** /api/integrations | Retrieves all Integrations
[**refresh_integrations**](IntegrationsApi.md#refresh_integrations) | **POST** /api/integrations/{id}/refresh | Refresh an Integration
[**remove_integration_objects**](IntegrationsApi.md#remove_integration_objects) | **DELETE** /api/integrations/{id}/objects/{objectId} | Deletes a ServiceNow Integration object
[**remove_integrations**](IntegrationsApi.md#remove_integrations) | **DELETE** /api/integrations/{id} | Deletes an Integration
[**update_integration_inventory**](IntegrationsApi.md#update_integration_inventory) | **PUT** /api/integrations/{id}/inventory/{inventoryId} | Updating an Integration Inventory
[**update_integrations**](IntegrationsApi.md#update_integrations) | **PUT** /api/integrations/{id} | Updates an Integration


# **add_integration_snow_objects**
> AddIntegrationSnowObjects200Response add_integration_snow_objects(id)

Creates an Exposed ServiceNow Catalog Item

This endpoint creates an Exposed Catalog Item. This is an integration object of type `catalog` that references a `Catalog Item Type.` 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.add_integration_snow_objects200_response import AddIntegrationSnowObjects200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_integration_snow_objects_request import AddIntegrationSnowObjectsRequest
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_integration_snow_objects_request = AddIntegrationSnowObjectsRequest(
        object=AddIntegrationSnowObjectsRequestObject(
            name="name_example",
            type="catalog",
            catalog_item_type=27,
        ),
    ) # AddIntegrationSnowObjectsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates an Exposed ServiceNow Catalog Item
        api_response = api_instance.add_integration_snow_objects(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->add_integration_snow_objects: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates an Exposed ServiceNow Catalog Item
        api_response = api_instance.add_integration_snow_objects(id, add_integration_snow_objects_request=add_integration_snow_objects_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->add_integration_snow_objects: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_integration_snow_objects_request** | [**AddIntegrationSnowObjectsRequest**](AddIntegrationSnowObjectsRequest.md)|  | [optional]

### Return type

[**AddIntegrationSnowObjects200Response**](AddIntegrationSnowObjects200Response.md)

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

# **add_integrations**
> AddIntegrations200Response add_integrations()

Creates an Integration

Creates an integration. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.add_integrations200_response import AddIntegrations200Response
from openapi_client.model.add_integrations_request import AddIntegrationsRequest
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    add_integrations_request = AddIntegrationsRequest(None) # AddIntegrationsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates an Integration
        api_response = api_instance.add_integrations(add_integrations_request=add_integrations_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->add_integrations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_integrations_request** | [**AddIntegrationsRequest**](AddIntegrationsRequest.md)|  | [optional]

### Return type

[**AddIntegrations200Response**](AddIntegrations200Response.md)

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

# **get_integration_inventory**
> GetIntegrationInventory200Response get_integration_inventory(id, inventory_id)

Get a Specific Integration Inventory

This endpoint retrieves a specific integration inventory. Only certain types of integrations support this operation, such as Ansible Tower. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_integration_inventory200_response import GetIntegrationInventory200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 9 # int | Morpheus ID of the Integration
    inventory_id = 1 # int | Morpheus ID of the Integration Inventory

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Integration Inventory
        api_response = api_instance.get_integration_inventory(id, inventory_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->get_integration_inventory: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Integration |
 **inventory_id** | **int**| Morpheus ID of the Integration Inventory |

### Return type

[**GetIntegrationInventory200Response**](GetIntegrationInventory200Response.md)

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

# **get_integration_objects**
> GetIntegrationObjects200Response get_integration_objects(id, object_id)

Get a Specific ServiceNow Integration Object

This endpoint retrieves a specific ServiceNow integration object.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.get_integration_objects200_response import GetIntegrationObjects200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    object_id = 122 # int | Morpheus ID of the Object being created or referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific ServiceNow Integration Object
        api_response = api_instance.get_integration_objects(id, object_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->get_integration_objects: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **object_id** | **int**| Morpheus ID of the Object being created or referenced |

### Return type

[**GetIntegrationObjects200Response**](GetIntegrationObjects200Response.md)

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

# **get_integration_type_option_types**
> GetIntegrationTypeOptionTypes200Response get_integration_type_option_types(id)

Retrieves a Option Types for a Specific Integration Type

This endpoint will retrieve the list of option types for a specific integration type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.get_integration_type_option_types200_response import GetIntegrationTypeOptionTypes200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Option Types for a Specific Integration Type
        api_response = api_instance.get_integration_type_option_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->get_integration_type_option_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetIntegrationTypeOptionTypes200Response**](GetIntegrationTypeOptionTypes200Response.md)

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

# **get_integration_types**
> GetIntegrationTypes200Response get_integration_types(id)

Retrieves a Specific Integration Type

Retrieves a specific integration type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_integration_types200_response import GetIntegrationTypes200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    optiontypes = False # bool | Pass `true` to include optionTypes in the response for each integration type (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Integration Type
        api_response = api_instance.get_integration_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->get_integration_types: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves a Specific Integration Type
        api_response = api_instance.get_integration_types(id, optiontypes=optiontypes)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->get_integration_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **optiontypes** | **bool**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] if omitted the server will use the default value of False

### Return type

[**GetIntegrationTypes200Response**](GetIntegrationTypes200Response.md)

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

# **get_integrations**
> AddIntegrations200Response get_integrations(id)

Retrieves a Specific Integration

Retrieves a specific integration. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.add_integrations200_response import AddIntegrations200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Integration
        api_response = api_instance.get_integrations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->get_integrations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddIntegrations200Response**](AddIntegrations200Response.md)

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

# **list_integration_inventory**
> ListIntegrationInventory200Response list_integration_inventory(id)

Get All Integration Inventory

This endpoint retrieves a list of inventory for a specific integration. Only certain types of integrations support this operation, such as Ansible Tower. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.list_integration_inventory200_response import ListIntegrationInventory200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 9 # int | Morpheus ID of the Integration
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Integration Inventory
        api_response = api_instance.list_integration_inventory(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->list_integration_inventory: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Integration Inventory
        api_response = api_instance.list_integration_inventory(id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->list_integration_inventory: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Integration |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListIntegrationInventory200Response**](ListIntegrationInventory200Response.md)

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

# **list_integration_objects**
> ListIntegrationObjects200Response list_integration_objects(id)

Get ServiceNow Integration Objects

This endpoint retrieves a list of exposed objects for a specific ServiceNow integration. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.list_integration_objects200_response import ListIntegrationObjects200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    value = "string" # str | The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing `type=string`. This means the data value returned by the API will be a string instead of an object.  (optional)
    ref_id = 3 # int | If specified will return an exact match on refId (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get ServiceNow Integration Objects
        api_response = api_instance.list_integration_objects(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->list_integration_objects: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get ServiceNow Integration Objects
        api_response = api_instance.list_integration_objects(id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, value=value, ref_id=ref_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->list_integration_objects: %s\n" % e)
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
 **value** | **str**| The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the data value returned by the API will be a string instead of an object.  | [optional]
 **ref_id** | **int**| If specified will return an exact match on refId | [optional]

### Return type

[**ListIntegrationObjects200Response**](ListIntegrationObjects200Response.md)

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

# **list_integration_types**
> ListIntegrationTypes200Response list_integration_types()

Retrieves all Integration Types

An Integration Type is specific third party software that the appliance can be configured to integrate with, such as Ansible, Chef, Git, ServiceNOW, etc. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.list_integration_types200_response import ListIntegrationTypes200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    optiontypes = False # bool | Pass `true` to include optionTypes in the response for each integration type (optional) if omitted the server will use the default value of False
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    category = "category_example" # str | If specified will return an exact match on category (optional)
    creatable = True # bool | Filter by creatable (optional)
    enabled = False # bool | Filter by enabled (optional)
    has_cmdb = True # bool | Filter by hasCMDB (optional)
    has_cmdb_discovery = True # bool | Filter by hasCMDBDiscovery (optional)
    has_cm = True # bool | Filter by hasCM (optional)
    has_dns = True # bool | Filter by hasDNS (optional)
    has_approvals = True # bool | Filter by hasApprovals (optional)
    is_plugin = True # bool | Filter by isPlugin (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Integration Types
        api_response = api_instance.list_integration_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, optiontypes=optiontypes, description=description, category=category, creatable=creatable, enabled=enabled, has_cmdb=has_cmdb, has_cmdb_discovery=has_cmdb_discovery, has_cm=has_cm, has_dns=has_dns, has_approvals=has_approvals, is_plugin=is_plugin)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->list_integration_types: %s\n" % e)
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
 **code** | **str**| If specified will return an exact match on code | [optional]
 **optiontypes** | **bool**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] if omitted the server will use the default value of False
 **description** | **str**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **category** | **str**| If specified will return an exact match on category | [optional]
 **creatable** | **bool**| Filter by creatable | [optional]
 **enabled** | **bool**| Filter by enabled | [optional]
 **has_cmdb** | **bool**| Filter by hasCMDB | [optional]
 **has_cmdb_discovery** | **bool**| Filter by hasCMDBDiscovery | [optional]
 **has_cm** | **bool**| Filter by hasCM | [optional]
 **has_dns** | **bool**| Filter by hasDNS | [optional]
 **has_approvals** | **bool**| Filter by hasApprovals | [optional]
 **is_plugin** | **bool**| Filter by isPlugin | [optional]

### Return type

[**ListIntegrationTypes200Response**](ListIntegrationTypes200Response.md)

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

# **list_integrations**
> ListIntegrations200Response list_integrations()

Retrieves all Integrations

Retrieves all integrations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.list_integrations200_response import ListIntegrations200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    id = 7 # int | Morpheus ID of the Object being created or referenced (optional)
    url = "https://example.com/testimage.ovf" # str | Download the file from a remote url. This can be used instead of uploading a local file. (optional)
    host = "host_example" # str | Filter by integration Host (optional)
    port = "port_example" # str | Filter by integration Port (optional)
    username = "administrator" # str | Username (optional)
    version = 5 # int | Filter by version number (userVersion) (optional)
    windows_version = "windowsVersion_example" # str | Filter by integration Windows Version (optional)
    status = "running" # str | The instance status for filtering. (optional)
    objects = False # bool | Include `objects=true` to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Integrations
        api_response = api_instance.list_integrations(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, id=id, url=url, host=host, port=port, username=username, version=version, windows_version=windows_version, status=status, objects=objects)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->list_integrations: %s\n" % e)
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
 **id** | **int**| Morpheus ID of the Object being created or referenced | [optional]
 **url** | **str**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional]
 **host** | **str**| Filter by integration Host | [optional]
 **port** | **str**| Filter by integration Port | [optional]
 **username** | **str**| Username | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **windows_version** | **str**| Filter by integration Windows Version | [optional]
 **status** | **str**| The instance status for filtering. | [optional]
 **objects** | **bool**| Include &#x60;objects&#x3D;true&#x60; to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  | [optional] if omitted the server will use the default value of False

### Return type

[**ListIntegrations200Response**](ListIntegrations200Response.md)

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

# **refresh_integrations**
> AddIntegrations200Response refresh_integrations(id)

Refresh an Integration

This endpoint will refresh an existing Integration. Only some types support this and will actually perform an action as a result. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.add_integrations200_response import AddIntegrations200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Refresh an Integration
        api_response = api_instance.refresh_integrations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->refresh_integrations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddIntegrations200Response**](AddIntegrations200Response.md)

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

# **remove_integration_objects**
> Model200Success remove_integration_objects(id, object_id)

Deletes a ServiceNow Integration object

Deletes a specified ServiceNow integration object. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    object_id = 122 # int | Morpheus ID of the Object being created or referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a ServiceNow Integration object
        api_response = api_instance.remove_integration_objects(id, object_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->remove_integration_objects: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **object_id** | **int**| Morpheus ID of the Object being created or referenced |

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

# **remove_integrations**
> Model200Success remove_integrations(id)

Deletes an Integration

Deletes a specified integration. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes an Integration
        api_response = api_instance.remove_integrations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->remove_integrations: %s\n" % e)
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

# **update_integration_inventory**
> UpdateIntegrationInventory200Response update_integration_inventory(id, inventory_id)

Updating an Integration Inventory

This endpoint provides updating of integration inventory

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.update_integration_inventory_request import UpdateIntegrationInventoryRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_integration_inventory200_response import UpdateIntegrationInventory200Response
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 9 # int | Morpheus ID of the Integration
    inventory_id = 1 # int | Morpheus ID of the Integration Inventory
    update_integration_inventory_request = UpdateIntegrationInventoryRequest(
        inventory=UpdateIntegrationInventoryRequestInventory(
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
        ),
    ) # UpdateIntegrationInventoryRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating an Integration Inventory
        api_response = api_instance.update_integration_inventory(id, inventory_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->update_integration_inventory: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating an Integration Inventory
        api_response = api_instance.update_integration_inventory(id, inventory_id, update_integration_inventory_request=update_integration_inventory_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->update_integration_inventory: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Integration |
 **inventory_id** | **int**| Morpheus ID of the Integration Inventory |
 **update_integration_inventory_request** | [**UpdateIntegrationInventoryRequest**](UpdateIntegrationInventoryRequest.md)|  | [optional]

### Return type

[**UpdateIntegrationInventory200Response**](UpdateIntegrationInventory200Response.md)

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

# **update_integrations**
> AddIntegrations200Response update_integrations(id)

Updates an Integration

Updates an integration. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import integrations_api
from openapi_client.model.add_integrations200_response import AddIntegrations200Response
from openapi_client.model.add_integrations_request import AddIntegrationsRequest
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
    api_instance = integrations_api.IntegrationsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_integrations_request = AddIntegrationsRequest(None) # AddIntegrationsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates an Integration
        api_response = api_instance.update_integrations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->update_integrations: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates an Integration
        api_response = api_instance.update_integrations(id, add_integrations_request=add_integrations_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IntegrationsApi->update_integrations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_integrations_request** | [**AddIntegrationsRequest**](AddIntegrationsRequest.md)|  | [optional]

### Return type

[**AddIntegrations200Response**](AddIntegrations200Response.md)

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

