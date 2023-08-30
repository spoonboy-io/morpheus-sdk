# openapi_client.MonitoringSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_monitoring_settings**](MonitoringSettingsApi.md#get_monitoring_settings) | **GET** /api/monitoring-settings | Get Monitoring Settings
[**update_monitoring_settings**](MonitoringSettingsApi.md#update_monitoring_settings) | **PUT** /api/monitoring-settings | Update Monitoring Settings


# **get_monitoring_settings**
> GetMonitoringSettings200Response get_monitoring_settings()

Get Monitoring Settings

This endpoint retrieves monitoring settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import monitoring_settings_api
from openapi_client.model.get_monitoring_settings200_response import GetMonitoringSettings200Response
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
    api_instance = monitoring_settings_api.MonitoringSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Monitoring Settings
        api_response = api_instance.get_monitoring_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling MonitoringSettingsApi->get_monitoring_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetMonitoringSettings200Response**](GetMonitoringSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Monitoring Settings |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_monitoring_settings**
> UpdateGuidanceSettings200Response update_monitoring_settings()

Update Monitoring Settings

Update Monitoring Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import monitoring_settings_api
from openapi_client.model.get_monitoring_settings200_response import GetMonitoringSettings200Response
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
    api_instance = monitoring_settings_api.MonitoringSettingsApi(api_client)
    get_monitoring_settings200_response = GetMonitoringSettings200Response(
        monitoring_settings=MonitoringSettings(
            auto_manage_checks=True,
            availability_time_frame=30,
            availability_precision=3,
            default_check_interval=5,
            service_now=MonitoringSettingsServiceNow(
                enabled=True,
                integration=MonitoringSettingsServiceNowIntegration(
                    id=1,
                    name="name_example",
                ),
                new_incident_action="create",
                close_incident_action="close",
                info_mapping="low",
                warning_mapping="medium",
                critical_mapping="high",
            ),
        ),
    ) # GetMonitoringSettings200Response |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Monitoring Settings
        api_response = api_instance.update_monitoring_settings(get_monitoring_settings200_response=get_monitoring_settings200_response)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling MonitoringSettingsApi->update_monitoring_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **get_monitoring_settings200_response** | [**GetMonitoringSettings200Response**](GetMonitoringSettings200Response.md)|  | [optional]

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
**200** | Update Monitoring Settings Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

