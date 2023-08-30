# openapi_client.BackupSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_backup_settings**](BackupSettingsApi.md#list_backup_settings) | **GET** /api/backup-settings | Get Backup Settings
[**update_backup_settings**](BackupSettingsApi.md#update_backup_settings) | **PUT** /api/backup-settings | Update Backup Settings


# **list_backup_settings**
> ListBackupSettings200Response list_backup_settings()

Get Backup Settings

This endpoint retrieves backup settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backup_settings_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_backup_settings200_response import ListBackupSettings200Response
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
    api_instance = backup_settings_api.BackupSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Backup Settings
        api_response = api_instance.list_backup_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupSettingsApi->list_backup_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListBackupSettings200Response**](ListBackupSettings200Response.md)

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

# **update_backup_settings**
> Model200Success update_backup_settings()

Update Backup Settings

Update Backup Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backup_settings_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_backup_settings_request import UpdateBackupSettingsRequest
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
    api_instance = backup_settings_api.BackupSettingsApi(api_client)
    update_backup_settings_request = UpdateBackupSettingsRequest(
        backup_settings=BackupSettingsUpdate(
            backups_enabled=True,
            retention_count=1,
            create_backups=True,
            backup_appliance=True,
            update_existing=True,
            default_schedule=BackupSettingsUpdateDefaultSchedule(
                id=1,
            ),
            clear_default_schedule=True,
            default_storage_bucket=BackupSettingsUpdateDefaultStorageBucket(
                id=1,
            ),
            clear_default_storage_bucket=True,
        ),
    ) # UpdateBackupSettingsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Backup Settings
        api_response = api_instance.update_backup_settings(update_backup_settings_request=update_backup_settings_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupSettingsApi->update_backup_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_backup_settings_request** | [**UpdateBackupSettingsRequest**](UpdateBackupSettingsRequest.md)|  | [optional]

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

