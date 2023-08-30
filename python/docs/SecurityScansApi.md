# openapi_client.SecurityScansApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_security_scans**](SecurityScansApi.md#get_security_scans) | **GET** /api/security-scans/{id} | Retrieves a Specific Security Scan
[**list_security_scans**](SecurityScansApi.md#list_security_scans) | **GET** /api/security-scans | Retrieves all Security Scans


# **get_security_scans**
> GetSecurityScans200Response get_security_scans(id)

Retrieves a Specific Security Scan

Retrieves a specific security scan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_scans_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_security_scans200_response import GetSecurityScans200Response
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
    api_instance = security_scans_api.SecurityScansApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    results = True # bool | Include the `results` object in the response under the security scan. This is a potentially very large object containing the raw results of the scan. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Security Scan
        api_response = api_instance.get_security_scans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityScansApi->get_security_scans: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves a Specific Security Scan
        api_response = api_instance.get_security_scans(id, results=results)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityScansApi->get_security_scans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **results** | **bool**| Include the &#x60;results&#x60; object in the response under the security scan. This is a potentially very large object containing the raw results of the scan. | [optional] if omitted the server will use the default value of False

### Return type

[**GetSecurityScans200Response**](GetSecurityScans200Response.md)

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

# **list_security_scans**
> ListSecurityScans200Response list_security_scans()

Retrieves all Security Scans

Retrieves all security scans. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import security_scans_api
from openapi_client.model.list_security_scans200_response import ListSecurityScans200Response
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
    api_instance = security_scans_api.SecurityScansApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "scanDate" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "scanDate"
    direction = "desc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "desc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description of security package (optional)
    security_package_id = 1 # int | Filter results by security package id(s). This parameter can be passed multiple times to match more than one id. (optional)
    server_id = 97 # int | The Server ID for Filtering (optional)
    results = True # bool | Include the `results` object in the response under each security scan. This is a potentially very large object containing the raw results of the scan. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Security Scans
        api_response = api_instance.list_security_scans(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, security_package_id=security_package_id, server_id=server_id, results=results)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SecurityScansApi->list_security_scans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "scanDate"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "desc"
 **phrase** | **str**| Search phrase for partial matches on name or description of security package | [optional]
 **security_package_id** | **int**| Filter results by security package id(s). This parameter can be passed multiple times to match more than one id. | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **results** | **bool**| Include the &#x60;results&#x60; object in the response under each security scan. This is a potentially very large object containing the raw results of the scan. | [optional] if omitted the server will use the default value of False

### Return type

[**ListSecurityScans200Response**](ListSecurityScans200Response.md)

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

