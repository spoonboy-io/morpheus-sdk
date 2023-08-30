# openapi_client.SSLCertificatesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_certificate**](SSLCertificatesApi.md#add_certificate) | **POST** /api/certificates | Create a Certificate
[**delete_certificate**](SSLCertificatesApi.md#delete_certificate) | **DELETE** /api/certificates/{id} | Delete a Certificate
[**get_certificate**](SSLCertificatesApi.md#get_certificate) | **GET** /api/certificates/{id} | Get a Specific Certificate
[**list_certificates**](SSLCertificatesApi.md#list_certificates) | **GET** /api/certificates | Get All SSL Certificates
[**update_certificate**](SSLCertificatesApi.md#update_certificate) | **PUT** /api/certificates/{id} | Update a Certificate


# **add_certificate**
> AddCertificate200Response add_certificate()

Create a Certificate

Create a Certificate

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import ssl_certificates_api
from openapi_client.model.add_certificate_request import AddCertificateRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_certificate200_response import AddCertificate200Response
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
    api_instance = ssl_certificates_api.SSLCertificatesApi(api_client)
    add_certificate_request = AddCertificateRequest(
        certificate=AddCertificateRequestCertificate(
            name="name_example",
            cert_file="cert_file_example",
            key_file="key_file_example",
            domain_name="domain_name_example",
            wildcard=False,
        ),
    ) # AddCertificateRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Certificate
        api_response = api_instance.add_certificate(add_certificate_request=add_certificate_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SSLCertificatesApi->add_certificate: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_certificate_request** | [**AddCertificateRequest**](AddCertificateRequest.md)|  | [optional]

### Return type

[**AddCertificate200Response**](AddCertificate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_certificate**
> Model200Success delete_certificate(id)

Delete a Certificate

Will delete a certificate from the system and make it no longer usable.  If a certificate is actively in use, a delete will fail.  

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import ssl_certificates_api
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
    api_instance = ssl_certificates_api.SSLCertificatesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Certificate
        api_response = api_instance.delete_certificate(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SSLCertificatesApi->delete_certificate: %s\n" % e)
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

# **get_certificate**
> AddCertificate200ResponseAllOf get_certificate(id)

Get a Specific Certificate

This endpoint retrieves a specific certificate.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import ssl_certificates_api
from openapi_client.model.add_certificate200_response_all_of import AddCertificate200ResponseAllOf
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
    api_instance = ssl_certificates_api.SSLCertificatesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Certificate
        api_response = api_instance.get_certificate(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SSLCertificatesApi->get_certificate: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddCertificate200ResponseAllOf**](AddCertificate200ResponseAllOf.md)

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

# **list_certificates**
> ListCertificates200Response list_certificates()

Get All SSL Certificates

This endpoint retrieves all SSL certificates associated with the account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import ssl_certificates_api
from openapi_client.model.list_certificates200_response import ListCertificates200Response
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
    api_instance = ssl_certificates_api.SSLCertificatesApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All SSL Certificates
        api_response = api_instance.list_certificates(name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SSLCertificatesApi->list_certificates: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListCertificates200Response**](ListCertificates200Response.md)

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

# **update_certificate**
> AddCertificate200ResponseAllOf update_certificate(id)

Update a Certificate

Update a Certificate.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import ssl_certificates_api
from openapi_client.model.add_certificate_request import AddCertificateRequest
from openapi_client.model.add_certificate200_response_all_of import AddCertificate200ResponseAllOf
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
    api_instance = ssl_certificates_api.SSLCertificatesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_certificate_request = AddCertificateRequest(
        certificate=AddCertificateRequestCertificate(
            name="name_example",
            cert_file="cert_file_example",
            key_file="key_file_example",
            domain_name="domain_name_example",
            wildcard=False,
        ),
    ) # AddCertificateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Certificate
        api_response = api_instance.update_certificate(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SSLCertificatesApi->update_certificate: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Certificate
        api_response = api_instance.update_certificate(id, add_certificate_request=add_certificate_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SSLCertificatesApi->update_certificate: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_certificate_request** | [**AddCertificateRequest**](AddCertificateRequest.md)|  | [optional]

### Return type

[**AddCertificate200ResponseAllOf**](AddCertificate200ResponseAllOf.md)

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

