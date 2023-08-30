# openapi_client.ProvisioningSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_provisioning_settings**](ProvisioningSettingsApi.md#list_provisioning_settings) | **GET** /api/provisioning-settings | Get Provisioning Settings
[**update_provisioning_settings**](ProvisioningSettingsApi.md#update_provisioning_settings) | **PUT** /api/provisioning-settings | Update Provisioning Settings


# **list_provisioning_settings**
> ListProvisioningSettings200Response list_provisioning_settings()

Get Provisioning Settings

This endpoint retrieves provisioning settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_settings_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_provisioning_settings200_response import ListProvisioningSettings200Response
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
    api_instance = provisioning_settings_api.ProvisioningSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Provisioning Settings
        api_response = api_instance.list_provisioning_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningSettingsApi->list_provisioning_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListProvisioningSettings200Response**](ListProvisioningSettings200Response.md)

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

# **update_provisioning_settings**
> Model200Success update_provisioning_settings()

Update Provisioning Settings

Update Provisioning Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_settings_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_provisioning_settings_request import UpdateProvisioningSettingsRequest
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
    api_instance = provisioning_settings_api.ProvisioningSettingsApi(api_client)
    update_provisioning_settings_request = UpdateProvisioningSettingsRequest(
        provisioning_settings=ProvisioningSettingsUpdate(
            allow_zone_selection=True,
            allow_server_selection=True,
            require_environments=True,
            show_pricing=True,
            hide_datastore_stats=True,
            cross_tenant_naming_policies=True,
            reuse_sequence=True,
            cloud_init_username="cloud_init_username_example",
            cloud_init_password="cloud_init_password_example",
            cloud_init_key_pair=ProvisioningSettingsUpdateCloudInitKeyPair(
                id=1,
            ),
            deploy_storage_provider=ProvisioningSettingsUpdateDeployStorageProvider(
                id=1,
            ),
            windows_password="windows_password_example",
            pxe_root_password="pxe_root_password_example",
            default_template_type=ProvisioningSettingsUpdateDefaultTemplateType(
                id=1,
            ),
        ),
    ) # UpdateProvisioningSettingsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Provisioning Settings
        api_response = api_instance.update_provisioning_settings(update_provisioning_settings_request=update_provisioning_settings_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningSettingsApi->update_provisioning_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_provisioning_settings_request** | [**UpdateProvisioningSettingsRequest**](UpdateProvisioningSettingsRequest.md)|  | [optional]

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

