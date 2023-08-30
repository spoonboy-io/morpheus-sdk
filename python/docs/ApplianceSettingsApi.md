# openapi_client.ApplianceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_appliance_settings**](ApplianceSettingsApi.md#list_appliance_settings) | **GET** /api/appliance-settings | Get Appliance Settings
[**reindex**](ApplianceSettingsApi.md#reindex) | **POST** /api/appliance-settings/reindex | Reindex Search
[**set_appliance_settings_maintenance_mode**](ApplianceSettingsApi.md#set_appliance_settings_maintenance_mode) | **POST** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**update_appliance_settings**](ApplianceSettingsApi.md#update_appliance_settings) | **PUT** /api/appliance-settings | Update Appliance Settings


# **list_appliance_settings**
> ListApplianceSettings200Response list_appliance_settings()

Get Appliance Settings

This endpoint retrieves appliance settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import appliance_settings_api
from openapi_client.model.list_appliance_settings200_response import ListApplianceSettings200Response
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
    api_instance = appliance_settings_api.ApplianceSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Appliance Settings
        api_response = api_instance.list_appliance_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ApplianceSettingsApi->list_appliance_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListApplianceSettings200Response**](ListApplianceSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reindex**
> Model200Success reindex()

Reindex Search

Reindex all searchable data

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import appliance_settings_api
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
    api_instance = appliance_settings_api.ApplianceSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Reindex Search
        api_response = api_instance.reindex()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ApplianceSettingsApi->reindex: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_appliance_settings_maintenance_mode**
> Model200Success set_appliance_settings_maintenance_mode()

Toggle Maintenance Mode

This endpoint allows toggling the appliance maintenance mode.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import appliance_settings_api
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
    api_instance = appliance_settings_api.ApplianceSettingsApi(api_client)
    enabled = True # bool | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Toggle Maintenance Mode
        api_response = api_instance.set_appliance_settings_maintenance_mode(enabled=enabled)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ApplianceSettingsApi->set_appliance_settings_maintenance_mode: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool**| Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_appliance_settings**
> Model200Success update_appliance_settings()

Update Appliance Settings

Update Appliance Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import appliance_settings_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_appliance_settings_request import UpdateApplianceSettingsRequest
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
    api_instance = appliance_settings_api.ApplianceSettingsApi(api_client)
    update_appliance_settings_request = UpdateApplianceSettingsRequest(
        appliance_settings=ApplianceSettingsUpdate(
            appliance_url="appliance_url_example",
            internal_appliance_url="internal_appliance_url_example",
            cors_allowed="cors_allowed_example",
            registration_enabled=True,
            default_role_id=1,
            default_user_role_id=1,
            docker_privileged_mode=True,
            password_min_length="password_min_length_example",
            password_min_upper_case="password_min_upper_case_example",
            password_min_numbers="password_min_numbers_example",
            password_min_symbols="password_min_symbols_example",
            user_browser_session_timeout="user_browser_session_timeout_example",
            user_browser_session_warning="user_browser_session_warning_example",
            expire_pwd_days=1,
            disable_after_attempts=1,
            disable_after_days_inactive=1,
            warn_user_days_before=1,
            smtp_mail_from="smtp_mail_from_example",
            smtp_server="smtp_server_example",
            smtp_port=1,
            smtp_ssl=True,
            smtp_tls=True,
            smtp_user="smtp_user_example",
            smtp_password="smtp_password_example",
            proxy_host="proxy_host_example",
            proxy_port="proxy_port_example",
            proxy_user="proxy_user_example",
            proxy_password="proxy_password_example",
            proxy_domain="proxy_domain_example",
            proxy_workstation="proxy_workstation_example",
            currency_provider="currency_provider_example",
            currency_key="currency_key_example",
            enable_all_zone_types=True,
            enable_zone_types=[
                1,
            ],
            disable_zone_types=[
                1,
            ],
            disable_all_zone_types=True,
        ),
    ) # UpdateApplianceSettingsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Appliance Settings
        api_response = api_instance.update_appliance_settings(update_appliance_settings_request=update_appliance_settings_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ApplianceSettingsApi->update_appliance_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_appliance_settings_request** | [**UpdateApplianceSettingsRequest**](UpdateApplianceSettingsRequest.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

