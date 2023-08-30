# openapi_client.DashboardApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard**](DashboardApi.md#dashboard) | **GET** /api/dashboard | Overview of Dashboard information


# **dashboard**
> Dashboard dashboard()

Overview of Dashboard information

This endpoint can be used to view dashboard information about the remote Morpheus appliance. This is an overview and summary of data available to the user that can be used to render a dashboard. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import dashboard_api
from openapi_client.model.dashboard import Dashboard
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
    api_instance = dashboard_api.DashboardApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Overview of Dashboard information
        api_response = api_instance.dashboard()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DashboardApi->dashboard: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**Dashboard**](Dashboard.md)

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

