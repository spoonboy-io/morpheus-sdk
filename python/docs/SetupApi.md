# openapi_client.SetupApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**setup**](SetupApi.md#setup) | **POST** /api/setup | Setup appliance


# **setup**
> Setup setup()

Setup appliance

Initialize a freshly installed appliance to create the master tenant and System Admin user.  Authorization is not required.  This operation can only be executed successfully once. 

### Example


```python
import time
import openapi_client
from openapi_client.api import setup_api
from openapi_client.model.setup_request import SetupRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.setup import Setup
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = setup_api.SetupApi(api_client)
    setup_request = SetupRequest(None) # SetupRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Setup appliance
        api_response = api_instance.setup(setup_request=setup_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SetupApi->setup: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setup_request** | [**SetupRequest**](SetupRequest.md)|  | [optional]

### Return type

[**Setup**](Setup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Appliance setup complete. The master tenant and system admin user were created successfully. |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

