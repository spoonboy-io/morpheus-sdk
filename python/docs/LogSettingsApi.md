# openapi_client.LogSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_log_settings_syslog_rules**](LogSettingsApi.md#add_log_settings_syslog_rules) | **POST** /api/log-settings/syslog-rules | Create a New Syslog Rule
[**delete_log_settings_syslog_rules**](LogSettingsApi.md#delete_log_settings_syslog_rules) | **DELETE** /api/log-settings/syslog-rules/{id} | Delete a Specific Syslog Rule
[**list_log_settings**](LogSettingsApi.md#list_log_settings) | **GET** /api/log-settings | List All Log Settings
[**update_log_settings**](LogSettingsApi.md#update_log_settings) | **PUT** /api/log-settings | Update Log Settings


# **add_log_settings_syslog_rules**
> Model200Success add_log_settings_syslog_rules()

Create a New Syslog Rule

Creates a new syslog rule. This command will also update existing syslog rule by specified name if already exists

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import log_settings_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_log_settings_syslog_rules_request import AddLogSettingsSyslogRulesRequest
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
    api_instance = log_settings_api.LogSettingsApi(api_client)
    add_log_settings_syslog_rules_request = AddLogSettingsSyslogRulesRequest(
        syslog_rule=AddLogSettingsSyslogRulesRequestSyslogRule(
            name="foo",
            rule="*.* @@bar.com:8080",
        ),
    ) # AddLogSettingsSyslogRulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Syslog Rule
        api_response = api_instance.add_log_settings_syslog_rules(add_log_settings_syslog_rules_request=add_log_settings_syslog_rules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LogSettingsApi->add_log_settings_syslog_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_log_settings_syslog_rules_request** | [**AddLogSettingsSyslogRulesRequest**](AddLogSettingsSyslogRulesRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Syslog Rule Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_log_settings_syslog_rules**
> Model200Success delete_log_settings_syslog_rules(id)

Delete a Specific Syslog Rule

Will delete the syslog rule matching the specified name.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import log_settings_api
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
    api_instance = log_settings_api.LogSettingsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Syslog Rule
        api_response = api_instance.delete_log_settings_syslog_rules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LogSettingsApi->delete_log_settings_syslog_rules: %s\n" % e)
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_log_settings**
> ListLogSettings200Response list_log_settings()

List All Log Settings

This endpoint retrieves log settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import log_settings_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_log_settings200_response import ListLogSettings200Response
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
    api_instance = log_settings_api.LogSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # List All Log Settings
        api_response = api_instance.list_log_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LogSettingsApi->list_log_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListLogSettings200Response**](ListLogSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Log Settings |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_log_settings**
> UpdateGuidanceSettings200Response update_log_settings()

Update Log Settings

Update Log Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import log_settings_api
from openapi_client.model.update_guidance_settings200_response import UpdateGuidanceSettings200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_log_settings200_response import ListLogSettings200Response
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
    api_instance = log_settings_api.LogSettingsApi(api_client)
    list_log_settings200_response = ListLogSettings200Response(
        log_settings=LogSettings(
            enabled=True,
            retention_days="retention_days_example",
            syslog_rules=[{"name":"foo","rule":"*.* @@localhost:4567"}],
            integrations=[
                {},
            ],
        ),
    ) # ListLogSettings200Response |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Log Settings
        api_response = api_instance.update_log_settings(list_log_settings200_response=list_log_settings200_response)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LogSettingsApi->update_log_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list_log_settings200_response** | [**ListLogSettings200Response**](ListLogSettings200Response.md)|  | [optional]

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
**200** | Update Log Settings Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

