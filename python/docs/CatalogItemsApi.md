# openapi_client.CatalogItemsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_catalog_item_type**](CatalogItemsApi.md#add_catalog_item_type) | **POST** /api/catalog-item-types | Create a Catalog Item Type
[**get_catalog_item_type**](CatalogItemsApi.md#get_catalog_item_type) | **GET** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type
[**list_catalog_item_types**](CatalogItemsApi.md#list_catalog_item_types) | **GET** /api/catalog-item-types | Get All Catalog Item Types
[**remove_catalog_item_type**](CatalogItemsApi.md#remove_catalog_item_type) | **DELETE** /api/catalog-item-types/{id} | Delete a Catalog Item Type
[**update_catalog_item_type**](CatalogItemsApi.md#update_catalog_item_type) | **PUT** /api/catalog-item-types/{id} | Update a Catalog Item Type
[**update_catalog_item_type_logo**](CatalogItemsApi.md#update_catalog_item_type_logo) | **PUT** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type


# **add_catalog_item_type**
> AddCatalogItemType200Response add_catalog_item_type()

Create a Catalog Item Type

Use this command to create a catalog item type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import catalog_items_api
from openapi_client.model.add_catalog_item_type200_response import AddCatalogItemType200Response
from openapi_client.model.add_catalog_item_type_request import AddCatalogItemTypeRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_catalog_item_type_request1 import AddCatalogItemTypeRequest1
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
    api_instance = catalog_items_api.CatalogItemsApi(api_client)
    add_catalog_item_type_request = AddCatalogItemTypeRequest(
        catalog_item_type=AddCatalogItemTypeRequestCatalogItemType(None),
    ) # AddCatalogItemTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Catalog Item Type
        api_response = api_instance.add_catalog_item_type(add_catalog_item_type_request=add_catalog_item_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->add_catalog_item_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_catalog_item_type_request** | [**AddCatalogItemTypeRequest**](AddCatalogItemTypeRequest.md)|  | [optional]

### Return type

[**AddCatalogItemType200Response**](AddCatalogItemType200Response.md)

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

# **get_catalog_item_type**
> GetCatalogItemType200Response get_catalog_item_type(id)

Get a Specific Catalog Item Type

This endpoint retrieves a specific category item type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import catalog_items_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_catalog_item_type200_response import GetCatalogItemType200Response
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
    api_instance = catalog_items_api.CatalogItemsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Catalog Item Type
        api_response = api_instance.get_catalog_item_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->get_catalog_item_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCatalogItemType200Response**](GetCatalogItemType200Response.md)

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

# **list_catalog_item_types**
> ListCatalogItemTypes200Response list_catalog_item_types()

Get All Catalog Item Types

This endpoint retrieves all catalog item types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import catalog_items_api
from openapi_client.model.list_catalog_item_types200_response import ListCatalogItemTypes200Response
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
    api_instance = catalog_items_api.CatalogItemsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    enabled = False # bool | Filter by enabled (optional)
    featured = False # bool | Filter by featured (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Catalog Item Types
        api_response = api_instance.list_catalog_item_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, description=description, enabled=enabled, featured=featured, labels=labels, all_labels=all_labels, code=code)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->list_catalog_item_types: %s\n" % e)
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
 **featured** | **bool**| Filter by featured | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]

### Return type

[**ListCatalogItemTypes200Response**](ListCatalogItemTypes200Response.md)

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

# **remove_catalog_item_type**
> Model200Success remove_catalog_item_type(id)

Delete a Catalog Item Type

Will delete a catalog item type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import catalog_items_api
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
    api_instance = catalog_items_api.CatalogItemsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Catalog Item Type
        api_response = api_instance.remove_catalog_item_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->remove_catalog_item_type: %s\n" % e)
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

# **update_catalog_item_type**
> UpdateCatalogItemType200Response update_catalog_item_type(id)

Update a Catalog Item Type

Use this command to update an existing catalog item type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import catalog_items_api
from openapi_client.model.update_catalog_item_type200_response import UpdateCatalogItemType200Response
from openapi_client.model.update_catalog_item_type_request import UpdateCatalogItemTypeRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_catalog_item_type_request1 import AddCatalogItemTypeRequest1
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
    api_instance = catalog_items_api.CatalogItemsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_catalog_item_type_request = UpdateCatalogItemTypeRequest(
        catalog_item_type=UpdateCatalogItemTypeRequestCatalogItemType(None),
    ) # UpdateCatalogItemTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Catalog Item Type
        api_response = api_instance.update_catalog_item_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->update_catalog_item_type: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Catalog Item Type
        api_response = api_instance.update_catalog_item_type(id, update_catalog_item_type_request=update_catalog_item_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->update_catalog_item_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_catalog_item_type_request** | [**UpdateCatalogItemTypeRequest**](UpdateCatalogItemTypeRequest.md)|  | [optional]

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md)

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

# **update_catalog_item_type_logo**
> UpdateCatalogItemType200Response update_catalog_item_type_logo(id)

Update Logo For Catalog Item Type

Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import catalog_items_api
from openapi_client.model.update_catalog_item_type200_response import UpdateCatalogItemType200Response
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
    api_instance = catalog_items_api.CatalogItemsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    catalog_item_type_logo = open('/path/to/file', 'rb') # file_type | Logo File png,jpg,svg (optional)
    catalog_item_type_dark_logo = open('/path/to/file', 'rb') # file_type | Dark Logo File png,jpg,svg (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Logo For Catalog Item Type
        api_response = api_instance.update_catalog_item_type_logo(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->update_catalog_item_type_logo: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Logo For Catalog Item Type
        api_response = api_instance.update_catalog_item_type_logo(id, catalog_item_type_logo=catalog_item_type_logo, catalog_item_type_dark_logo=catalog_item_type_dark_logo)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CatalogItemsApi->update_catalog_item_type_logo: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **catalog_item_type_logo** | **file_type**| Logo File png,jpg,svg | [optional]
 **catalog_item_type_dark_logo** | **file_type**| Dark Logo File png,jpg,svg | [optional]

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

