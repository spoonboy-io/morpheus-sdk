# openapi_client.WikiApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_wiki**](WikiApi.md#add_wiki) | **POST** /api/wiki/pages | Create a Wiki Page
[**get_wiki**](WikiApi.md#get_wiki) | **GET** /api/wiki/pages/{id} | Retrieves a specific Wiki page
[**get_wiki_app**](WikiApi.md#get_wiki_app) | **GET** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**get_wiki_categories**](WikiApi.md#get_wiki_categories) | **GET** /api/wiki/categories | Retrieves all Wiki categories associated with the account
[**get_wiki_cloud**](WikiApi.md#get_wiki_cloud) | **GET** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**get_wiki_cluster**](WikiApi.md#get_wiki_cluster) | **GET** /api/clusters/{clusterId}/wiki | Retrieves a Cluster Wiki Page
[**get_wiki_group**](WikiApi.md#get_wiki_group) | **GET** /api/groups/{id}/wiki | Retrieves a Group Wiki Page
[**get_wiki_instance**](WikiApi.md#get_wiki_instance) | **GET** /api/instances/{id}/wiki | Retrieves an Instance Wiki Page
[**get_wiki_server**](WikiApi.md#get_wiki_server) | **GET** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**list_wiki**](WikiApi.md#list_wiki) | **GET** /api/wiki/pages | Retrieves wiki pages associated with the account.
[**remove_wiki**](WikiApi.md#remove_wiki) | **DELETE** /api/wiki/pages/{id} | Deletes a Wiki Page
[**update_wiki**](WikiApi.md#update_wiki) | **PUT** /api/wiki/pages/{id} | Updates a Wiki Page
[**update_wiki_app**](WikiApi.md#update_wiki_app) | **PUT** /api/apps/{id}/wiki | Update an App Wiki Page
[**update_wiki_cloud**](WikiApi.md#update_wiki_cloud) | **PUT** /api/zones/{id}/wiki | Update a Cloud Wiki Page
[**update_wiki_cluster**](WikiApi.md#update_wiki_cluster) | **PUT** /api/clusters/{clusterId}/wiki | Update a Cluster Wiki Page
[**update_wiki_group**](WikiApi.md#update_wiki_group) | **PUT** /api/groups/{id}/wiki | Update a Group Wiki Page
[**update_wiki_instance**](WikiApi.md#update_wiki_instance) | **PUT** /api/instances/{id}/wiki | Update an Instance Wiki Page
[**update_wiki_server**](WikiApi.md#update_wiki_server) | **PUT** /api/servers/{id}/wiki | Update a Server Wiki Page


# **add_wiki**
> UpdateWikiApp200Response add_wiki()

Create a Wiki Page

Creates a Wiki Page 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.add_wiki_request import AddWikiRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    add_wiki_request = AddWikiRequest(
        page=AddWikiRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.",
        ),
    ) # AddWikiRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Wiki Page
        api_response = api_instance.add_wiki(add_wiki_request=add_wiki_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->add_wiki: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_wiki_request** | [**AddWikiRequest**](AddWikiRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **get_wiki**
> GetWikiApp200Response get_wiki(id)

Retrieves a specific Wiki page

This endpoint retrieves a specific wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a specific Wiki page
        api_response = api_instance.get_wiki(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **get_wiki_app**
> GetWikiApp200Response get_wiki_app(id)

Retrieves an App Wiki Page

This endpoint retrieves an app Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves an App Wiki Page
        api_response = api_instance.get_wiki_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **get_wiki_categories**
> GetWikiCategories200Response get_wiki_categories()

Retrieves all Wiki categories associated with the account

This endpoint retrieves all Wiki categories associated with the account. The results are not paginated. The categories returned are those of the found pages. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_categories200_response import GetWikiCategories200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    page_phrase = "pagePhrase_example" # str | If specified will return a partial match on page name (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Wiki categories associated with the account
        api_response = api_instance.get_wiki_categories(phrase=phrase, page_phrase=page_phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_categories: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **page_phrase** | **str**| If specified will return a partial match on page name | [optional]

### Return type

[**GetWikiCategories200Response**](GetWikiCategories200Response.md)

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

# **get_wiki_cloud**
> GetWikiApp200Response get_wiki_cloud(id)

Retrieves a Cloud Wiki Page

This endpoint retrieves a cloud Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Cloud Wiki Page
        api_response = api_instance.get_wiki_cloud(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **get_wiki_cluster**
> GetWikiApp200Response get_wiki_cluster(cluster_id)

Retrieves a Cluster Wiki Page

This endpoint retrieves a cluster Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    cluster_id = 5 # int | The ID of the cluster

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Cluster Wiki Page
        api_response = api_instance.get_wiki_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **get_wiki_group**
> GetWikiApp200Response get_wiki_group(id)

Retrieves a Group Wiki Page

This endpoint retrieves a group Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Group Wiki Page
        api_response = api_instance.get_wiki_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **get_wiki_instance**
> GetWikiApp200Response get_wiki_instance(id)

Retrieves an Instance Wiki Page

This endpoint retrieves an instance Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves an Instance Wiki Page
        api_response = api_instance.get_wiki_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **get_wiki_server**
> GetWikiApp200Response get_wiki_server(id)

Retrieves a Server Wiki Page

This endpoint retrieves a server Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Server Wiki Page
        api_response = api_instance.get_wiki_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->get_wiki_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **list_wiki**
> GetWikiApp200Response list_wiki()

Retrieves wiki pages associated with the account.

This endpoint retrieves wiki pages associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = wiki_api.WikiApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves wiki pages associated with the account.
        api_response = api_instance.list_wiki(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->list_wiki: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **remove_wiki**
> Model200Success remove_wiki(id)

Deletes a Wiki Page

Deletes the specified Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Wiki Page
        api_response = api_instance.remove_wiki(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->remove_wiki: %s\n" % e)
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

# **update_wiki**
> UpdateWikiApp200Response update_wiki(id)

Updates a Wiki Page

Updates a Wiki Page 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Wiki Page
        api_response = api_instance.update_wiki(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Wiki Page
        api_response = api_instance.update_wiki(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **update_wiki_app**
> UpdateWikiApp200Response update_wiki_app(id)

Update an App Wiki Page

Updates an app Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an App Wiki Page
        api_response = api_instance.update_wiki_app(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_app: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an App Wiki Page
        api_response = api_instance.update_wiki_app(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_app: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **update_wiki_cloud**
> UpdateWikiApp200Response update_wiki_cloud(id)

Update a Cloud Wiki Page

Updates a cloud Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Cloud Wiki Page
        api_response = api_instance.update_wiki_cloud(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_cloud: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Cloud Wiki Page
        api_response = api_instance.update_wiki_cloud(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **update_wiki_cluster**
> UpdateWikiApp200Response update_wiki_cluster(cluster_id)

Update a Cluster Wiki Page

Updates a cluster Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    cluster_id = 5 # int | The ID of the cluster
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Cluster Wiki Page
        api_response = api_instance.update_wiki_cluster(cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_cluster: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Cluster Wiki Page
        api_response = api_instance.update_wiki_cluster(cluster_id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **int**| The ID of the cluster |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **update_wiki_group**
> UpdateWikiApp200Response update_wiki_group(id)

Update a Group Wiki Page

Updates a group Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Group Wiki Page
        api_response = api_instance.update_wiki_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Group Wiki Page
        api_response = api_instance.update_wiki_group(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **update_wiki_instance**
> UpdateWikiApp200Response update_wiki_instance(id)

Update an Instance Wiki Page

Updates an instance Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Instance Wiki Page
        api_response = api_instance.update_wiki_instance(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_instance: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Instance Wiki Page
        api_response = api_instance.update_wiki_instance(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_instance: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

# **update_wiki_server**
> UpdateWikiApp200Response update_wiki_server(id)

Update a Server Wiki Page

Updates a server Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import wiki_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = wiki_api.WikiApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Server Wiki Page
        api_response = api_instance.update_wiki_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Server Wiki Page
        api_response = api_instance.update_wiki_server(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WikiApi->update_wiki_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

