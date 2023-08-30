# openapi_client.VDIApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_vdi_allocation**](VDIApi.md#add_vdi_allocation) | **POST** /api/vdi/{id}/allocate | Allocate Virtual Desktop
[**add_vdi_apps**](VDIApi.md#add_vdi_apps) | **POST** /api/vdi-apps | Creates a VDI App
[**add_vdi_gateways**](VDIApi.md#add_vdi_gateways) | **POST** /api/vdi-gateways | Creates a VDI Gateway
[**add_vdi_pools**](VDIApi.md#add_vdi_pools) | **POST** /api/vdi-pools | Creates a VDI Pool
[**get_vdi**](VDIApi.md#get_vdi) | **GET** /api/vdi/{id} | Get a Specific Virtual Desktop
[**get_vdi_allocations**](VDIApi.md#get_vdi_allocations) | **GET** /api/vdi-allocations/{id} | Retrieves a Specific VDI Allocation
[**get_vdi_apps**](VDIApi.md#get_vdi_apps) | **GET** /api/vdi-apps/{id} | Retrieves a Specific VDI App
[**get_vdi_gateways**](VDIApi.md#get_vdi_gateways) | **GET** /api/vdi-gateways/{id} | Retrieves a Specific VDI Gateway
[**get_vdi_pools**](VDIApi.md#get_vdi_pools) | **GET** /api/vdi-pools/{id} | Retrieves a Specific VDI Pool
[**list_vdi**](VDIApi.md#list_vdi) | **GET** /api/vdi | List Virtual Desktops
[**list_vdi_allocations**](VDIApi.md#list_vdi_allocations) | **GET** /api/vdi-allocations | Retrieves all VDI Allocations
[**list_vdi_apps**](VDIApi.md#list_vdi_apps) | **GET** /api/vdi-apps | Retrieves all VDI Apps
[**list_vdi_gateways**](VDIApi.md#list_vdi_gateways) | **GET** /api/vdi-gateways | Retrieves all VDI Gateways
[**list_vdi_pools**](VDIApi.md#list_vdi_pools) | **GET** /api/vdi-pools | Retrieves all VDI Pools
[**remove_vdi_apps**](VDIApi.md#remove_vdi_apps) | **DELETE** /api/vdi-apps/{id} | Deletes a VDI App
[**remove_vdi_gateways**](VDIApi.md#remove_vdi_gateways) | **DELETE** /api/vdi-gateways/{id} | Deletes a VDI Gateway
[**remove_vdi_pools**](VDIApi.md#remove_vdi_pools) | **DELETE** /api/vdi-pools/{id} | Deletes a VDI Pool
[**update_vdi_apps**](VDIApi.md#update_vdi_apps) | **PUT** /api/vdi-apps/{id} | Updates a VDI App Configuration or Icon
[**update_vdi_gateways**](VDIApi.md#update_vdi_gateways) | **PUT** /api/vdi-gateways/{id} | Updates a VDI Gateway Configuration
[**update_vdi_pools**](VDIApi.md#update_vdi_pools) | **PUT** /api/vdi-pools/{id} | Updates a VDI Pool Configuration or Icon


# **add_vdi_allocation**
> AddVdiAllocation200Response add_vdi_allocation(id)

Allocate Virtual Desktop

This endpoint allocates a specific virtual desktop for use by your user. It will return the desktop and its allocation for your user, or an error if allocation fails, which will occur if the desktop is fully allocated already. If your user already has an allocation, the desktop and allocation will still be returned succesfully and the server does not make any changes. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_allocation200_response import AddVdiAllocation200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Allocate Virtual Desktop
        api_response = api_instance.add_vdi_allocation(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->add_vdi_allocation: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddVdiAllocation200Response**](AddVdiAllocation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | VDI Allocate Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_vdi_apps**
> AddVDIApps200Response add_vdi_apps()

Creates a VDI App

Creates a VDI app. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_apps_request1 import AddVDIAppsRequest1
from openapi_client.model.add_vdi_apps_request import AddVDIAppsRequest
from openapi_client.model.add_vdi_apps200_response import AddVDIApps200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    add_vdi_apps_request = AddVDIAppsRequest(
        vdi_app=AddVDIAppsRequestVdiApp(None),
    ) # AddVDIAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a VDI App
        api_response = api_instance.add_vdi_apps(add_vdi_apps_request=add_vdi_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->add_vdi_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_vdi_apps_request** | [**AddVDIAppsRequest**](AddVDIAppsRequest.md)|  | [optional]

### Return type

[**AddVDIApps200Response**](AddVDIApps200Response.md)

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

# **add_vdi_gateways**
> AddVDIGateways200Response add_vdi_gateways()

Creates a VDI Gateway

Creates a VDI gateway. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_gateways200_response import AddVDIGateways200Response
from openapi_client.model.add_vdi_gateways_request import AddVDIGatewaysRequest
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
    api_instance = vdi_api.VDIApi(api_client)
    add_vdi_gateways_request = AddVDIGatewaysRequest(
        vdi_gateway=AddVDIGatewaysRequestVdiGateway(None),
    ) # AddVDIGatewaysRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a VDI Gateway
        api_response = api_instance.add_vdi_gateways(add_vdi_gateways_request=add_vdi_gateways_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->add_vdi_gateways: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_vdi_gateways_request** | [**AddVDIGatewaysRequest**](AddVDIGatewaysRequest.md)|  | [optional]

### Return type

[**AddVDIGateways200Response**](AddVDIGateways200Response.md)

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

# **add_vdi_pools**
> AddVDIPools200Response add_vdi_pools()

Creates a VDI Pool

Creates a VDI pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_pools200_response import AddVDIPools200Response
from openapi_client.model.add_vdi_pools_request import AddVDIPoolsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_vdi_pools_request1 import AddVDIPoolsRequest1
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
    api_instance = vdi_api.VDIApi(api_client)
    add_vdi_pools_request = AddVDIPoolsRequest(
        vdi_pool=AddVDIPoolsRequestVdiPool(None),
    ) # AddVDIPoolsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a VDI Pool
        api_response = api_instance.add_vdi_pools(add_vdi_pools_request=add_vdi_pools_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->add_vdi_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_vdi_pools_request** | [**AddVDIPoolsRequest**](AddVDIPoolsRequest.md)|  | [optional]

### Return type

[**AddVDIPools200Response**](AddVDIPools200Response.md)

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

# **get_vdi**
> GetVdi200Response get_vdi(id)

Get a Specific Virtual Desktop

This endpoint retrieves a specific virtual desktop along with the allocation for your user if one exists. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.get_vdi200_response import GetVdi200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Virtual Desktop
        api_response = api_instance.get_vdi(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->get_vdi: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetVdi200Response**](GetVdi200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | VDI Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_vdi_allocations**
> GetVDIAllocations200Response get_vdi_allocations(id)

Retrieves a Specific VDI Allocation

Retrieves a specific VDI allocation. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.get_vdi_allocations200_response import GetVDIAllocations200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific VDI Allocation
        api_response = api_instance.get_vdi_allocations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->get_vdi_allocations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetVDIAllocations200Response**](GetVDIAllocations200Response.md)

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

# **get_vdi_apps**
> AddVDIApps200ResponseAnyOf get_vdi_apps(id)

Retrieves a Specific VDI App

Retrieves a specific VDI app. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_apps200_response_any_of import AddVDIApps200ResponseAnyOf
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific VDI App
        api_response = api_instance.get_vdi_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->get_vdi_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddVDIApps200ResponseAnyOf**](AddVDIApps200ResponseAnyOf.md)

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

# **get_vdi_gateways**
> AddVDIGateways200ResponseAnyOf get_vdi_gateways(id)

Retrieves a Specific VDI Gateway

Retrieves a specific VDI gateway. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_vdi_gateways200_response_any_of import AddVDIGateways200ResponseAnyOf
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific VDI Gateway
        api_response = api_instance.get_vdi_gateways(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->get_vdi_gateways: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddVDIGateways200ResponseAnyOf**](AddVDIGateways200ResponseAnyOf.md)

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

# **get_vdi_pools**
> AddVDIPools200ResponseAnyOf get_vdi_pools(id)

Retrieves a Specific VDI Pool

Retrieves a specific VDI pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_pools200_response_any_of import AddVDIPools200ResponseAnyOf
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific VDI Pool
        api_response = api_instance.get_vdi_pools(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->get_vdi_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddVDIPools200ResponseAnyOf**](AddVDIPools200ResponseAnyOf.md)

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

# **list_vdi**
> ListVdi200Response list_vdi()

List Virtual Desktops

This endpoint retrieves all virtual desktops along with the allocation for your user if one exists. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.list_vdi200_response import ListVdi200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List Virtual Desktops
        api_response = api_instance.list_vdi(phrase=phrase, max=max, offset=offset, sort=sort, direction=direction, name=name, description=description)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->list_vdi: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **description** | **str**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]

### Return type

[**ListVdi200Response**](ListVdi200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of VDI |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_vdi_allocations**
> ListVDIAllocations200Response list_vdi_allocations()

Retrieves all VDI Allocations

Retrieves all VDI allocations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.list_vdi_allocations200_response import ListVDIAllocations200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    id = "preparing" # str | Filter by allocation ID (optional)
    status = "preparing" # str | Filter by allocation status (optional)
    pool_id = 1 # int | Filter by `VDI Pool` ID (optional)
    user_id = 6 # int | Filter by User ID (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all VDI Allocations
        api_response = api_instance.list_vdi_allocations(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, id=id, status=status, pool_id=pool_id, user_id=user_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->list_vdi_allocations: %s\n" % e)
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
 **id** | **str**| Filter by allocation ID | [optional]
 **status** | **str**| Filter by allocation status | [optional]
 **pool_id** | **int**| Filter by &#x60;VDI Pool&#x60; ID | [optional]
 **user_id** | **int**| Filter by User ID | [optional]

### Return type

[**ListVDIAllocations200Response**](ListVDIAllocations200Response.md)

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

# **list_vdi_apps**
> ListVDIApps200Response list_vdi_apps()

Retrieves all VDI Apps

Retrieves all VDI apps. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_vdi_apps200_response import ListVDIApps200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all VDI Apps
        api_response = api_instance.list_vdi_apps(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, description=description)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->list_vdi_apps: %s\n" % e)
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
 **description** | **str**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]

### Return type

[**ListVDIApps200Response**](ListVDIApps200Response.md)

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

# **list_vdi_gateways**
> ListVDIGateways200Response list_vdi_gateways()

Retrieves all VDI Gateways

Retrieves all VDI gateways. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.list_vdi_gateways200_response import ListVDIGateways200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all VDI Gateways
        api_response = api_instance.list_vdi_gateways(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, description=description)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->list_vdi_gateways: %s\n" % e)
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
 **description** | **str**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]

### Return type

[**ListVDIGateways200Response**](ListVDIGateways200Response.md)

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

# **list_vdi_pools**
> ListVDIPools200Response list_vdi_pools()

Retrieves all VDI Pools

Retrieves all VDI pools. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_vdi_pools200_response import ListVDIPools200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    enabled = False # bool | Filter by enabled (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all VDI Pools
        api_response = api_instance.list_vdi_pools(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, description=description, enabled=enabled)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->list_vdi_pools: %s\n" % e)
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
 **description** | **str**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **enabled** | **bool**| Filter by enabled | [optional]

### Return type

[**ListVDIPools200Response**](ListVDIPools200Response.md)

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

# **remove_vdi_apps**
> Model200Success remove_vdi_apps(id)

Deletes a VDI App

Deletes a specified VDI App. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a VDI App
        api_response = api_instance.remove_vdi_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->remove_vdi_apps: %s\n" % e)
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

# **remove_vdi_gateways**
> Model200Success remove_vdi_gateways(id)

Deletes a VDI Gateway

Deletes a specified VDI Gateway. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a VDI Gateway
        api_response = api_instance.remove_vdi_gateways(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->remove_vdi_gateways: %s\n" % e)
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

# **remove_vdi_pools**
> Model200Success remove_vdi_pools(id)

Deletes a VDI Pool

Deletes a specified VDI Pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a VDI Pool
        api_response = api_instance.remove_vdi_pools(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->remove_vdi_pools: %s\n" % e)
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

# **update_vdi_apps**
> AddVDIApps200Response update_vdi_apps(id)

Updates a VDI App Configuration or Icon

Updates a VDI App configuration or icon. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.update_vdi_apps_request import UpdateVDIAppsRequest
from openapi_client.model.add_vdi_apps_request1 import AddVDIAppsRequest1
from openapi_client.model.add_vdi_apps200_response import AddVDIApps200Response
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_vdi_apps_request = UpdateVDIAppsRequest(
        vdi_app=UpdateVDIAppsRequestVdiApp(None),
    ) # UpdateVDIAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a VDI App Configuration or Icon
        api_response = api_instance.update_vdi_apps(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->update_vdi_apps: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a VDI App Configuration or Icon
        api_response = api_instance.update_vdi_apps(id, update_vdi_apps_request=update_vdi_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->update_vdi_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_vdi_apps_request** | [**UpdateVDIAppsRequest**](UpdateVDIAppsRequest.md)|  | [optional]

### Return type

[**AddVDIApps200Response**](AddVDIApps200Response.md)

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

# **update_vdi_gateways**
> AddVDIGateways200Response update_vdi_gateways(id)

Updates a VDI Gateway Configuration

Updates a VDI Gateway configuration or icon. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.add_vdi_gateways200_response import AddVDIGateways200Response
from openapi_client.model.update_vdi_gateways_request import UpdateVDIGatewaysRequest
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_vdi_gateways_request = UpdateVDIGatewaysRequest(
        vdi_gateway=UpdateVDIGatewaysRequestVdiGateway(None),
    ) # UpdateVDIGatewaysRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a VDI Gateway Configuration
        api_response = api_instance.update_vdi_gateways(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->update_vdi_gateways: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a VDI Gateway Configuration
        api_response = api_instance.update_vdi_gateways(id, update_vdi_gateways_request=update_vdi_gateways_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->update_vdi_gateways: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_vdi_gateways_request** | [**UpdateVDIGatewaysRequest**](UpdateVDIGatewaysRequest.md)|  | [optional]

### Return type

[**AddVDIGateways200Response**](AddVDIGateways200Response.md)

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

# **update_vdi_pools**
> AddVDIPools200Response update_vdi_pools(id)

Updates a VDI Pool Configuration or Icon

Updates a VDI Pool configuration or icon. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import vdi_api
from openapi_client.model.update_vdi_pools_request import UpdateVDIPoolsRequest
from openapi_client.model.add_vdi_pools200_response import AddVDIPools200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_vdi_pools_request1 import AddVDIPoolsRequest1
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
    api_instance = vdi_api.VDIApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_vdi_pools_request = UpdateVDIPoolsRequest(
        vdi_pool=UpdateVDIPoolsRequestVdiPool(
            name="name_example",
            description="description_example",
            owner=21,
            min_idle=5,
            initial_pool_size=10,
            max_idle=2,
            max_pool_size=50,
            allocation_timeout_minutes=20,
            persistent_user=False,
            recyclable=False,
            allow_copy=False,
            allow_printer=False,
            allow_fileshare=False,
            allow_hypervisor_console=False,
            auto_create_local_user_on_reservation=False,
            enabled=True,
            icon_path="/assets/containers-png/windows.png",
            apps=[
                1,
            ],
            gateway=1,
            instance_config="instance_config_example",
            config=AddVDIPoolsRequestVdiPoolOneOfConfig(
                name="name_example",
                group=AddVDIPoolsRequestVdiPoolOneOfConfigGroup(None),
                cloud=AddVDIPoolsRequestVdiPoolOneOfConfigCloud(None),
                type=AddVDIPoolsRequestVdiPoolOneOfConfigType(None),
                layout=AddVDIPoolsRequestVdiPoolOneOfConfigLayout(None),
                plan=AddVDIPoolsRequestVdiPoolOneOfConfigPlan(None),
            ),
            guest_console_jump_host="guest_console_jump_host_example",
            guest_console_jump_port=1,
            guest_console_jump_username="guest_console_jump_username_example",
            guest_console_jump_password="guest_console_jump_password_example",
            guest_console_jump_keypair=1,
        ),
    ) # UpdateVDIPoolsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a VDI Pool Configuration or Icon
        api_response = api_instance.update_vdi_pools(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->update_vdi_pools: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a VDI Pool Configuration or Icon
        api_response = api_instance.update_vdi_pools(id, update_vdi_pools_request=update_vdi_pools_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling VDIApi->update_vdi_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_vdi_pools_request** | [**UpdateVDIPoolsRequest**](UpdateVDIPoolsRequest.md)|  | [optional]

### Return type

[**AddVDIPools200Response**](AddVDIPools200Response.md)

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

