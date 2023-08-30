# openapi_client.ServiceCatalogApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_catalog_cart**](ServiceCatalogApi.md#add_catalog_cart) | **POST** /api/catalog/checkout | Checkout Catalog Cart
[**add_catalog_cart_item**](ServiceCatalogApi.md#add_catalog_cart_item) | **PUT** /api/catalog/cart/items | Add Catalog Item to Cart
[**add_catalog_order**](ServiceCatalogApi.md#add_catalog_order) | **POST** /api/catalog/orders | Place Catalog Order
[**delete_catalog_cart**](ServiceCatalogApi.md#delete_catalog_cart) | **DELETE** /api/catalog/cart | Clear Catalog Cart
[**delete_catalog_cart_item**](ServiceCatalogApi.md#delete_catalog_cart_item) | **DELETE** /api/catalog/cart/items/{id} | Remove a Catalog Item From Cart
[**delete_catalog_item**](ServiceCatalogApi.md#delete_catalog_item) | **DELETE** /api/catalog/items/{id} | Delete a Catalog Inventory Item
[**get_catalog_item**](ServiceCatalogApi.md#get_catalog_item) | **GET** /api/catalog/items/{id} | Get a Specific Catalog Inventory Item
[**get_catalog_type**](ServiceCatalogApi.md#get_catalog_type) | **GET** /api/catalog/types/{id} | Get a Specific Catalog Type
[**list_catalog_cart**](ServiceCatalogApi.md#list_catalog_cart) | **GET** /api/catalog/cart | Get Catalog Cart
[**list_catalog_items**](ServiceCatalogApi.md#list_catalog_items) | **GET** /api/catalog/items | List Catalog Inventory Items
[**list_catalog_types**](ServiceCatalogApi.md#list_catalog_types) | **GET** /api/catalog/types | List Catalog Types


# **add_catalog_cart**
> AddCatalogCart200Response add_catalog_cart()

Checkout Catalog Cart

Use this command to checkout, finalizing your cart and placing an order. This converts each item in the cart to an inventory item, changing the status from IN_CART to ORDERED and potentially starts the provisioning process for each item.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.add_catalog_cart200_response import AddCatalogCart200Response
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Checkout Catalog Cart
        api_response = api_instance.add_catalog_cart()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->add_catalog_cart: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**AddCatalogCart200Response**](AddCatalogCart200Response.md)

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

# **add_catalog_cart_item**
> AddCatalogCartItem200Response add_catalog_cart_item()

Add Catalog Item to Cart

Use this command to add an item to your service catalog cart.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.add_catalog_cart_item200_response import AddCatalogCartItem200Response
from openapi_client.model.catalog_cart_item_create import CatalogCartItemCreate
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    validate = False # bool | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory (optional) if omitted the server will use the default value of False
    catalog_cart_item_create = CatalogCartItemCreate(
        item=CatalogCartItemCreateItem(
            type=CatalogCartItemCreateItemType(
                name="name_example",
            ),
            quantity=1,
            config={},
            context="context_example",
            target=1,
        ),
    ) # CatalogCartItemCreate |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Add Catalog Item to Cart
        api_response = api_instance.add_catalog_cart_item(validate=validate, catalog_cart_item_create=catalog_cart_item_create)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->add_catalog_cart_item: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] if omitted the server will use the default value of False
 **catalog_cart_item_create** | [**CatalogCartItemCreate**](CatalogCartItemCreate.md)|  | [optional]

### Return type

[**AddCatalogCartItem200Response**](AddCatalogCartItem200Response.md)

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

# **add_catalog_order**
> AddCatalogOrder200Response add_catalog_order()

Place Catalog Order

This will place an order for the specified items, adding items to the inventory right away, without using the cart.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.add_catalog_order_request import AddCatalogOrderRequest
from openapi_client.model.add_catalog_order200_response import AddCatalogOrder200Response
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    validate = False # bool | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory (optional) if omitted the server will use the default value of False
    add_catalog_order_request = AddCatalogOrderRequest(
        order=CatalogOrderCreate(
            items=[
                CatalogOrderCreateItemsInner(
                    type=CatalogOrderCreateItemsInnerType(
                        name="name_example",
                    ),
                    config={},
                    context="context_example",
                    target=1,
                ),
            ],
        ),
    ) # AddCatalogOrderRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Place Catalog Order
        api_response = api_instance.add_catalog_order(validate=validate, add_catalog_order_request=add_catalog_order_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->add_catalog_order: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **bool**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] if omitted the server will use the default value of False
 **add_catalog_order_request** | [**AddCatalogOrderRequest**](AddCatalogOrderRequest.md)|  | [optional]

### Return type

[**AddCatalogOrder200Response**](AddCatalogOrder200Response.md)

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

# **delete_catalog_cart**
> Model200Success delete_catalog_cart()

Clear Catalog Cart

Use this command to empty your cart, deleting all the items in it.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Clear Catalog Cart
        api_response = api_instance.delete_catalog_cart()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->delete_catalog_cart: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

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

# **delete_catalog_cart_item**
> Model200Success delete_catalog_cart_item(id)

Remove a Catalog Item From Cart

Will remove a catalog item that is currently in the cart.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Remove a Catalog Item From Cart
        api_response = api_instance.delete_catalog_cart_item(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->delete_catalog_cart_item: %s\n" % e)
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

# **delete_catalog_item**
> Model200Success delete_catalog_item(id)

Delete a Catalog Inventory Item

Will delete a catalog inventory item, which by default will deprovision any associated any instances and servers.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    preserve_volumes = "on" # str | Preserve Volumes (optional) if omitted the server will use the default value of "off"
    keep_backups = "on" # str | Preserve copy of backups (optional) if omitted the server will use the default value of "off"
    release_floating_ips = "off" # str | Release Floating IPs (optional) if omitted the server will use the default value of "on"
    release_eips = "off" # str | Alias for releaseFloatingIps (optional) if omitted the server will use the default value of "on"
    remove_instances = "off" # str | Remove Instances. Only applies to type `blueprint` (Apps) (optional) if omitted the server will use the default value of "on"
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Catalog Inventory Item
        api_response = api_instance.delete_catalog_item(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->delete_catalog_item: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Catalog Inventory Item
        api_response = api_instance.delete_catalog_item(id, preserve_volumes=preserve_volumes, keep_backups=keep_backups, release_floating_ips=release_floating_ips, release_eips=release_eips, remove_instances=remove_instances, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->delete_catalog_item: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **preserve_volumes** | **str**| Preserve Volumes | [optional] if omitted the server will use the default value of "off"
 **keep_backups** | **str**| Preserve copy of backups | [optional] if omitted the server will use the default value of "off"
 **release_floating_ips** | **str**| Release Floating IPs | [optional] if omitted the server will use the default value of "on"
 **release_eips** | **str**| Alias for releaseFloatingIps | [optional] if omitted the server will use the default value of "on"
 **remove_instances** | **str**| Remove Instances. Only applies to type &#x60;blueprint&#x60; (Apps) | [optional] if omitted the server will use the default value of "on"
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

# **get_catalog_item**
> AddCatalogCartItem200ResponseAllOf get_catalog_item(id)

Get a Specific Catalog Inventory Item

This endpoint retrieves a specific catalog inventory item.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.add_catalog_cart_item200_response_all_of import AddCatalogCartItem200ResponseAllOf
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Catalog Inventory Item
        api_response = api_instance.get_catalog_item(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->get_catalog_item: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddCatalogCartItem200ResponseAllOf**](AddCatalogCartItem200ResponseAllOf.md)

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

# **get_catalog_type**
> GetCatalogType200Response get_catalog_type(id)

Get a Specific Catalog Type

This endpoint retrieves a specific catalog item type. This also returns an array of associated optionTypes that are used to configure the catalog item.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.get_catalog_type200_response import GetCatalogType200Response
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Catalog Type
        api_response = api_instance.get_catalog_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->get_catalog_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCatalogType200Response**](GetCatalogType200Response.md)

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

# **list_catalog_cart**
> ListCatalogCart200Response list_catalog_cart()

Get Catalog Cart

This endpoint retrieves the current catalog cart and all the items in it.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.list_catalog_cart200_response import ListCatalogCart200Response
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Catalog Cart
        api_response = api_instance.list_catalog_cart()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->list_catalog_cart: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListCatalogCart200Response**](ListCatalogCart200Response.md)

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

# **list_catalog_items**
> ListCatalogItems200Response list_catalog_items()

List Catalog Inventory Items

This endpoint retrieves a list of the catalog inventory items.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.list_catalog_items200_response import ListCatalogItems200Response
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List Catalog Inventory Items
        api_response = api_instance.list_catalog_items(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->list_catalog_items: %s\n" % e)
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

[**ListCatalogItems200Response**](ListCatalogItems200Response.md)

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

# **list_catalog_types**
> ListCatalogTypes200Response list_catalog_types()

List Catalog Types

This endpoint retrieves the types available for ordering.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import service_catalog_api
from openapi_client.model.list_catalog_types200_response import ListCatalogTypes200Response
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
    api_instance = service_catalog_api.ServiceCatalogApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    featured = False # bool | Filter by featured (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List Catalog Types
        api_response = api_instance.list_catalog_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, featured=featured)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ServiceCatalogApi->list_catalog_types: %s\n" % e)
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
 **featured** | **bool**| Filter by featured | [optional]

### Return type

[**ListCatalogTypes200Response**](ListCatalogTypes200Response.md)

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

