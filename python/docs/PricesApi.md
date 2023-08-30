# openapi_client.PricesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_prices**](PricesApi.md#add_prices) | **POST** /api/prices | Creates a Price
[**deactivate_prices**](PricesApi.md#deactivate_prices) | **PUT** /api/prices/{id}/deactivate | Deactivates a Price
[**get_prices**](PricesApi.md#get_prices) | **GET** /api/prices/{id} | Retrieves a Specific Price
[**list_prices**](PricesApi.md#list_prices) | **GET** /api/prices | Retrieves all Prices
[**update_prices**](PricesApi.md#update_prices) | **PUT** /api/prices/{id} | Updates a Price


# **add_prices**
> AddPrices200Response add_prices()

Creates a Price

Creates a price. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import prices_api
from openapi_client.model.add_prices_request import AddPricesRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_prices200_response import AddPrices200Response
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
    api_instance = prices_api.PricesApi(api_client)
    add_prices_request = AddPricesRequest(
        price=AddPricesRequestPrice(
            name="name_example",
            code="code_example",
            account=AddPricesRequestPriceAccount(
                id=1,
            ),
            price_type="fixed",
            price_unit="minute",
            incur_charges="running",
            currency="USD",
            cost=10.5,
            markup_type="fixed",
            markup=2.5,
            markup_percent=13.5,
            custom_price=12.25,
            platform="linux",
            software="software_example",
            volume_type=AddPricesRequestPriceVolumeType(
                id=1,
            ),
            datastore=AddPricesRequestPriceDatastore(
                id=1,
            ),
            cross_cloud_apply=True,
        ),
    ) # AddPricesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Price
        api_response = api_instance.add_prices(add_prices_request=add_prices_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PricesApi->add_prices: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_prices_request** | [**AddPricesRequest**](AddPricesRequest.md)|  | [optional]

### Return type

[**AddPrices200Response**](AddPrices200Response.md)

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

# **deactivate_prices**
> Model200Success deactivate_prices(id)

Deactivates a Price

Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import prices_api
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
    api_instance = prices_api.PricesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deactivates a Price
        api_response = api_instance.deactivate_prices(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PricesApi->deactivate_prices: %s\n" % e)
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

# **get_prices**
> AddPrices200ResponseAllOf get_prices(id)

Retrieves a Specific Price

Retrieves a specific price. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import prices_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_prices200_response_all_of import AddPrices200ResponseAllOf
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
    api_instance = prices_api.PricesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Price
        api_response = api_instance.get_prices(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PricesApi->get_prices: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddPrices200ResponseAllOf**](AddPrices200ResponseAllOf.md)

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

# **list_prices**
> ListPrices200Response list_prices()

Retrieves all Prices

Retrieves all prices. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import prices_api
from openapi_client.model.list_prices200_response import ListPrices200Response
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
    api_instance = prices_api.PricesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    include_inactive = True # bool | If true, include inactive prices in the results (optional)
    price_type = "fixed" # str | Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software  (optional)
    platform = "linux" # str | Restricts query to only load only prices with specified platform (optional)
    currency = "USD" # str | Restricts query to only load only prices with specified currency (optional)
    type = "type_example" # str | Filter by type code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Prices
        api_response = api_instance.list_prices(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, include_inactive=include_inactive, price_type=price_type, platform=platform, currency=currency, type=type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PricesApi->list_prices: %s\n" % e)
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
 **price_type** | **str**| Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  | [optional]
 **platform** | **str**| Restricts query to only load only prices with specified platform | [optional]
 **currency** | **str**| Restricts query to only load only prices with specified currency | [optional]
 **type** | **str**| Filter by type code | [optional]

### Return type

[**ListPrices200Response**](ListPrices200Response.md)

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

# **update_prices**
> AddPrices200Response update_prices(id)

Updates a Price

Updates a price. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import prices_api
from openapi_client.model.update_prices_request import UpdatePricesRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_prices200_response import AddPrices200Response
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
    api_instance = prices_api.PricesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_prices_request = UpdatePricesRequest(
        price=UpdatePricesRequestPrice(
            name="name_example",
            code="code_example",
            account=AddPricesRequestPriceAccount(
                id=1,
            ),
            price_type="fixed",
            price_unit="minute",
            incur_charges="running",
            currency="USD",
            cost=10.5,
            markup_type="fixed",
            markup=2.5,
            markup_percent=13.5,
            custom_price=12.25,
            platform="linux",
            software="software_example",
            volume_type=AddPricesRequestPriceVolumeType(
                id=1,
            ),
            datastore=AddPricesRequestPriceDatastore(
                id=1,
            ),
            cross_cloud_apply=True,
        ),
    ) # UpdatePricesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Price
        api_response = api_instance.update_prices(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PricesApi->update_prices: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Price
        api_response = api_instance.update_prices(id, update_prices_request=update_prices_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling PricesApi->update_prices: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_prices_request** | [**UpdatePricesRequest**](UpdatePricesRequest.md)|  | [optional]

### Return type

[**AddPrices200Response**](AddPrices200Response.md)

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

