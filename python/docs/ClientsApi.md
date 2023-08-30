# openapi_client.ClientsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_client**](ClientsApi.md#add_client) | **POST** /api/clients | Create an Oauth Client
[**get_clients**](ClientsApi.md#get_clients) | **GET** /api/clients/{id} | Retrieves a Specific Oauth Client
[**list_clients**](ClientsApi.md#list_clients) | **GET** /api/clients | Get All Oauth Clients
[**remove_clients**](ClientsApi.md#remove_clients) | **DELETE** /api/clients/{id} | Deletes an Oauth Client
[**update_clients**](ClientsApi.md#update_clients) | **PUT** /api/clients/{id} | Updates an Oauth Client


# **add_client**
> AddClient200Response add_client()

Create an Oauth Client

Create a new Oauth Client.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clients_api
from openapi_client.model.add_client_request import AddClientRequest
from openapi_client.model.add_client200_response import AddClient200Response
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
    api_instance = clients_api.ClientsApi(api_client)
    add_client_request = AddClientRequest(
        client=AddClientRequestClient(
            client_id="Test Client",
            client_secret="thisIsaClientSecretKeyPhrase",
            access_token_validity_seconds=43200,
            refresh_token_validity_seconds=43200,
        ),
    ) # AddClientRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Oauth Client
        api_response = api_instance.add_client(add_client_request=add_client_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClientsApi->add_client: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_client_request** | [**AddClientRequest**](AddClientRequest.md)|  | [optional]

### Return type

[**AddClient200Response**](AddClient200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Client Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_clients**
> GetClients200Response get_clients(id)

Retrieves a Specific Oauth Client

Retrieves a specific oauth client. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clients_api
from openapi_client.model.get_clients200_response import GetClients200Response
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
    api_instance = clients_api.ClientsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Oauth Client
        api_response = api_instance.get_clients(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClientsApi->get_clients: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetClients200Response**](GetClients200Response.md)

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

# **list_clients**
> ListClients200Response list_clients()

Get All Oauth Clients

This endpoint retrieves a paginated list of oauth clients.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clients_api
from openapi_client.model.list_clients200_response import ListClients200Response
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
    api_instance = clients_api.ClientsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "clientId" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "clientId"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on clientId (optional)
    client_id = "clientId_example" # str | Search phrase for partial matches on clientId (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Oauth Clients
        api_response = api_instance.list_clients(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, client_id=client_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClientsApi->list_clients: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "clientId"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on clientId | [optional]
 **client_id** | **str**| Search phrase for partial matches on clientId | [optional]

### Return type

[**ListClients200Response**](ListClients200Response.md)

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

# **remove_clients**
> Model200Success remove_clients(id)

Deletes an Oauth Client

Deletes a specified oauth client. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clients_api
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
    api_instance = clients_api.ClientsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes an Oauth Client
        api_response = api_instance.remove_clients(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClientsApi->remove_clients: %s\n" % e)
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

# **update_clients**
> AddClient200Response update_clients(id)

Updates an Oauth Client

Updates an oauth client. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clients_api
from openapi_client.model.update_clients_request import UpdateClientsRequest
from openapi_client.model.add_client200_response import AddClient200Response
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
    api_instance = clients_api.ClientsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_clients_request = UpdateClientsRequest(
        client=ClientUpdate(
            client_id="client_id_example",
            access_token_validity_seconds=1,
            refresh_token_validity_seconds=1,
        ),
    ) # UpdateClientsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates an Oauth Client
        api_response = api_instance.update_clients(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClientsApi->update_clients: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates an Oauth Client
        api_response = api_instance.update_clients(id, update_clients_request=update_clients_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ClientsApi->update_clients: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_clients_request** | [**UpdateClientsRequest**](UpdateClientsRequest.md)|  | [optional]

### Return type

[**AddClient200Response**](AddClient200Response.md)

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

