# openapi_client.SecurityPackagesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_security_packages**](SecurityPackagesApi.md#add_security_packages) | **POST** /api/security-packages | Creates a Security Package
[**get_security_packages**](SecurityPackagesApi.md#get_security_packages) | **GET** /api/security-packages/{id} | Retrieves a Specific Security Package
[**list_security_packages**](SecurityPackagesApi.md#list_security_packages) | **GET** /api/security-packages | Retrieves all Security Packages
[**remove_security_packages**](SecurityPackagesApi.md#remove_security_packages) | **DELETE** /api/security-packages/{id} | Deletes a Security Package
[**update_security_packages**](SecurityPackagesApi.md#update_security_packages) | **PUT** /api/security-packages/{id} | Updates a Security Package


# **add_security_packages**
> AddSecurityPackages200Response add_security_packages()

Creates a Security Package

Creates a security package. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_packages_api
from openapi_client.model.add_security_packages200_response import AddSecurityPackages200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_security_packages_request import AddSecurityPackagesRequest
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
    api_instance = security_packages_api.SecurityPackagesApi(api_client)
    add_security_packages_request = AddSecurityPackagesRequest(
        security_package=AddSecurityPackagesRequestSecurityPackage(
            name="Sample Security Package",
            labels=[
                "labels_example",
            ],
            type="scap-package",
            description="description_example",
            url="http://10.0.2.2:8080/public-archives/download/SCAP/scap-security-guide-0.1.51.zip",
            enabled=True,
        ),
    ) # AddSecurityPackagesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Security Package
        api_response = api_instance.add_security_packages(add_security_packages_request=add_security_packages_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityPackagesApi->add_security_packages: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_security_packages_request** | [**AddSecurityPackagesRequest**](AddSecurityPackagesRequest.md)|  | [optional]

### Return type

[**AddSecurityPackages200Response**](AddSecurityPackages200Response.md)

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

# **get_security_packages**
> GetSecurityPackages200Response get_security_packages(id)

Retrieves a Specific Security Package

Retrieves a specific security package. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_packages_api
from openapi_client.model.get_security_packages200_response import GetSecurityPackages200Response
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
    api_instance = security_packages_api.SecurityPackagesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Security Package
        api_response = api_instance.get_security_packages(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityPackagesApi->get_security_packages: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetSecurityPackages200Response**](GetSecurityPackages200Response.md)

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

# **list_security_packages**
> ListSecurityPackages200Response list_security_packages()

Retrieves all Security Packages

Retrieves all security packages. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_packages_api
from openapi_client.model.list_security_packages200_response import ListSecurityPackages200Response
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
    api_instance = security_packages_api.SecurityPackagesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Security Packages
        api_response = api_instance.list_security_packages(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityPackagesApi->list_security_packages: %s\n" % e)
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
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListSecurityPackages200Response**](ListSecurityPackages200Response.md)

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

# **remove_security_packages**
> Model200Success remove_security_packages(id)

Deletes a Security Package

Deletes a specified security package. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_packages_api
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
    api_instance = security_packages_api.SecurityPackagesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Security Package
        api_response = api_instance.remove_security_packages(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityPackagesApi->remove_security_packages: %s\n" % e)
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

# **update_security_packages**
> AddSecurityPackages200Response update_security_packages(id)

Updates a Security Package

Updates a security package. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_packages_api
from openapi_client.model.add_security_packages200_response import AddSecurityPackages200Response
from openapi_client.model.update_security_packages_request import UpdateSecurityPackagesRequest
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
    api_instance = security_packages_api.SecurityPackagesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_security_packages_request = UpdateSecurityPackagesRequest(
        security_package=UpdateSecurityPackagesRequestSecurityPackage(
            name="Sample Security Package",
            labels=[
                "labels_example",
            ],
            type="scap-package",
            description="description_example",
            url="http://10.0.2.2:8080/public-archives/download/SCAP/scap-security-guide-0.1.51.zip",
            enabled=True,
        ),
    ) # UpdateSecurityPackagesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Security Package
        api_response = api_instance.update_security_packages(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityPackagesApi->update_security_packages: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Security Package
        api_response = api_instance.update_security_packages(id, update_security_packages_request=update_security_packages_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityPackagesApi->update_security_packages: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_security_packages_request** | [**UpdateSecurityPackagesRequest**](UpdateSecurityPackagesRequest.md)|  | [optional]

### Return type

[**AddSecurityPackages200Response**](AddSecurityPackages200Response.md)

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

