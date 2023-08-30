# openapi_client.GuidanceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_guidance_settings**](GuidanceSettingsApi.md#get_guidance_settings) | **GET** /api/guidance-settings | Get Guidance Settings
[**update_guidance_settings**](GuidanceSettingsApi.md#update_guidance_settings) | **PUT** /api/guidance-settings | Update Guidance Settings


# **get_guidance_settings**
> GetGuidanceSettings200Response get_guidance_settings()

Get Guidance Settings

This endpoint retrieves guidance settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_settings_api
from openapi_client.model.get_guidance_settings200_response import GetGuidanceSettings200Response
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
    api_instance = guidance_settings_api.GuidanceSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Guidance Settings
        api_response = api_instance.get_guidance_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceSettingsApi->get_guidance_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetGuidanceSettings200Response**](GetGuidanceSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Guidance Settings |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_guidance_settings**
> UpdateGuidanceSettings200Response update_guidance_settings()

Update Guidance Settings

Update Guidance Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import guidance_settings_api
from openapi_client.model.get_guidance_settings200_response import GetGuidanceSettings200Response
from openapi_client.model.update_guidance_settings200_response import UpdateGuidanceSettings200Response
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
    api_instance = guidance_settings_api.GuidanceSettingsApi(api_client)
    get_guidance_settings200_response = GetGuidanceSettings200Response(
        guidance_settings=GuidanceSettings(
            cpu_avg_cutoff_power=75,
            cpu_max_cutoff_power=500,
            network_cutoff_power=2000,
            cpu_up_avg_standard_cutoff_right_size=50,
            cpu_up_max_standard_cutoff_right_size=99,
            memory_up_avg_standard_cutoff_right_size=10,
            memory_down_avg_standard_cutoff_right_size=60,
            memory_down_max_standard_cutoff_right_size=30,
        ),
    ) # GetGuidanceSettings200Response |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Guidance Settings
        api_response = api_instance.update_guidance_settings(get_guidance_settings200_response=get_guidance_settings200_response)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling GuidanceSettingsApi->update_guidance_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **get_guidance_settings200_response** | [**GetGuidanceSettings200Response**](GetGuidanceSettings200Response.md)|  | [optional]

### Return type

[**UpdateGuidanceSettings200Response**](UpdateGuidanceSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Update Guidance Settings Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

