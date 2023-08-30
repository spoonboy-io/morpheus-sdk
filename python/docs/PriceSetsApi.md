# openapi_client.PriceSetsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_price_sets**](PriceSetsApi.md#add_price_sets) | **POST** /api/price-sets | Creates a Price Set
[**deactivate_price_sets**](PriceSetsApi.md#deactivate_price_sets) | **PUT** /api/price-sets/{id}/deactivate | Deactivates a Price Set
[**get_price_sets**](PriceSetsApi.md#get_price_sets) | **GET** /api/price-sets/{id} | Retrieves a Specific Price Set
[**list_price_sets**](PriceSetsApi.md#list_price_sets) | **GET** /api/price-sets | Retrieves all Price Sets
[**update_price_sets**](PriceSetsApi.md#update_price_sets) | **PUT** /api/price-sets/{id} | Updates a Price Set


# **add_price_sets**
> AddPriceSets200Response add_price_sets()

Creates a Price Set

Creates a price set. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import price_sets_api
from openapi_client.model.add_price_sets200_response import AddPriceSets200Response
from openapi_client.model.add_price_sets_request import AddPriceSetsRequest
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
    api_instance = price_sets_api.PriceSetsApi(api_client)
    add_price_sets_request = AddPriceSetsRequest(
        price_set=AddPriceSetsRequestPriceSet(
            name="testName",
            code="priceSet1",
            region_code="region.code.1",
            zone=AddPriceSetsRequestPriceSetZone(
                id=12,
            ),
            zone_pool=AddPriceSetsRequestPriceSetZonePool(
                id=52,
            ),
            price_unit="minute",
            type="fixed",
            prices=[
                1,
            ],
        ),
    ) # AddPriceSetsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Price Set
        api_response = api_instance.add_price_sets(add_price_sets_request=add_price_sets_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PriceSetsApi->add_price_sets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_price_sets_request** | [**AddPriceSetsRequest**](AddPriceSetsRequest.md)|  | [optional]

### Return type

[**AddPriceSets200Response**](AddPriceSets200Response.md)

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

# **deactivate_price_sets**
> Model200Success deactivate_price_sets(id)

Deactivates a Price Set

Deactivates a price set. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import price_sets_api
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
    api_instance = price_sets_api.PriceSetsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deactivates a Price Set
        api_response = api_instance.deactivate_price_sets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PriceSetsApi->deactivate_price_sets: %s\n" % e)
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

# **get_price_sets**
> GetPriceSets200Response get_price_sets(id)

Retrieves a Specific Price Set

Retrieves a specific price set. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import price_sets_api
from openapi_client.model.get_price_sets200_response import GetPriceSets200Response
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
    api_instance = price_sets_api.PriceSetsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Price Set
        api_response = api_instance.get_price_sets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PriceSetsApi->get_price_sets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPriceSets200Response**](GetPriceSets200Response.md)

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

# **list_price_sets**
> ListPriceSets200Response list_price_sets()

Retrieves all Price Sets

Retrieves all price sets. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import price_sets_api
from openapi_client.model.list_price_sets200_response import ListPriceSets200Response
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
    api_instance = price_sets_api.PriceSetsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    include_inactive = True # bool | If true, include inactive prices in the results (optional)
    type = "type_example" # str | Filter by type code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Price Sets
        api_response = api_instance.list_price_sets(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, include_inactive=include_inactive, type=type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PriceSetsApi->list_price_sets: %s\n" % e)
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
 **include_inactive** | **bool**| If true, include inactive prices in the results | [optional]
 **type** | **str**| Filter by type code | [optional]

### Return type

[**ListPriceSets200Response**](ListPriceSets200Response.md)

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

# **update_price_sets**
> AddPriceSets200Response update_price_sets(id)

Updates a Price Set

Updates a price set. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import price_sets_api
from openapi_client.model.add_price_sets200_response import AddPriceSets200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_price_sets_request import UpdatePriceSetsRequest
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
    api_instance = price_sets_api.PriceSetsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_price_sets_request = UpdatePriceSetsRequest(
        price_set=UpdatePriceSetsRequestPriceSet(
            name="testName",
            code="priceSet1",
            region_code="region.code.1",
            zone=AddPriceSetsRequestPriceSetZone(
                id=12,
            ),
            zone_pool=AddPriceSetsRequestPriceSetZonePool(
                id=52,
            ),
            price_unit="minute",
            type="fixed",
            prices=[
                1,
            ],
        ),
    ) # UpdatePriceSetsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Price Set
        api_response = api_instance.update_price_sets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PriceSetsApi->update_price_sets: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Price Set
        api_response = api_instance.update_price_sets(id, update_price_sets_request=update_price_sets_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PriceSetsApi->update_price_sets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_price_sets_request** | [**UpdatePriceSetsRequest**](UpdatePriceSetsRequest.md)|  | [optional]

### Return type

[**AddPriceSets200Response**](AddPriceSets200Response.md)

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

