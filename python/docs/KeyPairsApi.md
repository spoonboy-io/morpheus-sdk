# openapi_client.KeyPairsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_key_pairs**](KeyPairsApi.md#add_key_pairs) | **POST** /api/key-pairs | Creates a Key Pair
[**generate_key_pairs**](KeyPairsApi.md#generate_key_pairs) | **POST** /api/key-pairs/generate | Generates a Key Pair
[**get_key_pairs**](KeyPairsApi.md#get_key_pairs) | **GET** /api/key-pairs/{id} | Retrieves a Specific Key Pair
[**remove_key_pairs**](KeyPairsApi.md#remove_key_pairs) | **DELETE** /api/key-pairs/{id} | Deletes a Key Pair


# **add_key_pairs**
> AddKeyPairs200Response add_key_pairs()

Creates a Key Pair

Creates a Key Pair.  Private keys are typically optional. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import key_pairs_api
from openapi_client.model.add_key_pairs_request import AddKeyPairsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_key_pairs200_response import AddKeyPairs200Response
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
    api_instance = key_pairs_api.KeyPairsApi(api_client)
    add_key_pairs_request = AddKeyPairsRequest(
        key_pair=AddKeyPairsRequestKeyPair(
            name="name_example",
            public_key="public_key_example",
            private_key="private_key_example",
            passphrase="passphrase_example",
        ),
    ) # AddKeyPairsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Key Pair
        api_response = api_instance.add_key_pairs(add_key_pairs_request=add_key_pairs_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling KeyPairsApi->add_key_pairs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_key_pairs_request** | [**AddKeyPairsRequest**](AddKeyPairsRequest.md)|  | [optional]

### Return type

[**AddKeyPairs200Response**](AddKeyPairs200Response.md)

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

# **generate_key_pairs**
> GenerateKeyPairs200Response generate_key_pairs()

Generates a Key Pair

Generates a Key Pair. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import key_pairs_api
from openapi_client.model.generate_key_pairs_request import GenerateKeyPairsRequest
from openapi_client.model.generate_key_pairs200_response import GenerateKeyPairs200Response
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
    api_instance = key_pairs_api.KeyPairsApi(api_client)
    generate_key_pairs_request = GenerateKeyPairsRequest(
        key_pair=GenerateKeyPairsRequestKeyPair(
            name="name_example",
        ),
    ) # GenerateKeyPairsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Generates a Key Pair
        api_response = api_instance.generate_key_pairs(generate_key_pairs_request=generate_key_pairs_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling KeyPairsApi->generate_key_pairs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generate_key_pairs_request** | [**GenerateKeyPairsRequest**](GenerateKeyPairsRequest.md)|  | [optional]

### Return type

[**GenerateKeyPairs200Response**](GenerateKeyPairs200Response.md)

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

# **get_key_pairs**
> AddKeyPairs200ResponseAllOf get_key_pairs(id)

Retrieves a Specific Key Pair

Retrieves a specific key pair. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import key_pairs_api
from openapi_client.model.add_key_pairs200_response_all_of import AddKeyPairs200ResponseAllOf
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
    api_instance = key_pairs_api.KeyPairsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Key Pair
        api_response = api_instance.get_key_pairs(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling KeyPairsApi->get_key_pairs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddKeyPairs200ResponseAllOf**](AddKeyPairs200ResponseAllOf.md)

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

# **remove_key_pairs**
> Model200Success remove_key_pairs(id)

Deletes a Key Pair

Deletes a specified key pair. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import key_pairs_api
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
    api_instance = key_pairs_api.KeyPairsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Key Pair
        api_response = api_instance.remove_key_pairs(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling KeyPairsApi->remove_key_pairs: %s\n" % e)
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

