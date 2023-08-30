# openapi_client.BackupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_backup_jobs**](BackupsApi.md#add_backup_jobs) | **POST** /api/backups/jobs | Creates a Backup Job
[**add_backups**](BackupsApi.md#add_backups) | **POST** /api/backups | Creates a Backup
[**execute_backup_jobs**](BackupsApi.md#execute_backup_jobs) | **POST** /api/backups/jobs/{id}/execute | Executes a Backup Job
[**execute_backups**](BackupsApi.md#execute_backups) | **POST** /api/backups/{id}/execute | Executes a Backup
[**get_backup_jobs**](BackupsApi.md#get_backup_jobs) | **GET** /api/backups/jobs/{id} | Retrieves a Specific Backup Job
[**get_backup_restores**](BackupsApi.md#get_backup_restores) | **GET** /api/backups/restores/{id} | Retrieves a Specific Backup Restore
[**get_backup_results**](BackupsApi.md#get_backup_results) | **GET** /api/backups/results/{id} | Retrieves a Specific Backup Result
[**get_backups**](BackupsApi.md#get_backups) | **GET** /api/backups/{id} | Retrieves a Specific Backup
[**list_backup_jobs**](BackupsApi.md#list_backup_jobs) | **GET** /api/backups/jobs | Retrieves all Backup Jobs
[**list_backup_restores**](BackupsApi.md#list_backup_restores) | **GET** /api/backups/restores | Retrieves all Backup Restores
[**list_backup_results**](BackupsApi.md#list_backup_results) | **GET** /api/backups/results | Retrieves all Backup Results
[**list_backups**](BackupsApi.md#list_backups) | **GET** /api/backups | Retrieves all Backups
[**remove_backup_jobs**](BackupsApi.md#remove_backup_jobs) | **DELETE** /api/backups/jobs/{id} | Deletes a Backup Job
[**remove_backup_restores**](BackupsApi.md#remove_backup_restores) | **DELETE** /api/backups/restores/{id} | Deletes a Backup Restore
[**remove_backup_results**](BackupsApi.md#remove_backup_results) | **DELETE** /api/backups/results/{id} | Deletes a Backup Result
[**remove_backups**](BackupsApi.md#remove_backups) | **DELETE** /api/backups/{id} | Deletes a Backup
[**update_backup_jobs**](BackupsApi.md#update_backup_jobs) | **PUT** /api/backups/jobs/{id} | Updates a Backup Job
[**update_backups**](BackupsApi.md#update_backups) | **PUT** /api/backups/{id} | Updates a Backup


# **add_backup_jobs**
> AddBackupJobs200Response add_backup_jobs()

Creates a Backup Job

This endpoint allows creating a Backup Job. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backup_jobs200_response import AddBackupJobs200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_backup_jobs_request import AddBackupJobsRequest
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
    api_instance = backups_api.BackupsApi(api_client)
    add_backup_jobs_request = AddBackupJobsRequest(
        job=AddBackupJobsRequestJob(
            name="name_example",
            code="code_example",
            schedule_id=1,
            retention_count=1,
        ),
    ) # AddBackupJobsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Backup Job
        api_response = api_instance.add_backup_jobs(add_backup_jobs_request=add_backup_jobs_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->add_backup_jobs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_backup_jobs_request** | [**AddBackupJobsRequest**](AddBackupJobsRequest.md)|  | [optional]

### Return type

[**AddBackupJobs200Response**](AddBackupJobs200Response.md)

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

# **add_backups**
> AddBackups200Response add_backups()

Creates a Backup

This endpoint allows creating a Backup. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backups_request import AddBackupsRequest
from openapi_client.model.add_backups200_response import AddBackups200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    add_backups_request = AddBackupsRequest(
        backup=AddBackupsRequestBackup(),
    ) # AddBackupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Backup
        api_response = api_instance.add_backups(add_backups_request=add_backups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->add_backups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_backups_request** | [**AddBackupsRequest**](AddBackupsRequest.md)|  | [optional]

### Return type

[**AddBackups200Response**](AddBackups200Response.md)

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

# **execute_backup_jobs**
> AddBackupJobs200Response execute_backup_jobs(id)

Executes a Backup Job

Executes a backup job. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backup_jobs200_response import AddBackupJobs200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    body = {} # {str: (bool, date, datetime, dict, float, int, list, str, none_type)} |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Executes a Backup Job
        api_response = api_instance.execute_backup_jobs(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->execute_backup_jobs: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Executes a Backup Job
        api_response = api_instance.execute_backup_jobs(id, body=body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->execute_backup_jobs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **body** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}**|  | [optional]

### Return type

[**AddBackupJobs200Response**](AddBackupJobs200Response.md)

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

# **execute_backups**
> AddBackups200Response execute_backups(id)

Executes a Backup

Executes a backup. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backups200_response import AddBackups200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    body = {} # {str: (bool, date, datetime, dict, float, int, list, str, none_type)} |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Executes a Backup
        api_response = api_instance.execute_backups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->execute_backups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Executes a Backup
        api_response = api_instance.execute_backups(id, body=body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->execute_backups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **body** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}**|  | [optional]

### Return type

[**AddBackups200Response**](AddBackups200Response.md)

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

# **get_backup_jobs**
> AddBackupJobs200ResponseAllOf get_backup_jobs(id)

Retrieves a Specific Backup Job

Retrieves a specific backup job. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_backup_jobs200_response_all_of import AddBackupJobs200ResponseAllOf
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Backup Job
        api_response = api_instance.get_backup_jobs(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->get_backup_jobs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddBackupJobs200ResponseAllOf**](AddBackupJobs200ResponseAllOf.md)

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

# **get_backup_restores**
> GetBackupRestores200Response get_backup_restores(id)

Retrieves a Specific Backup Restore

Retrieves a specific backup restore. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.get_backup_restores200_response import GetBackupRestores200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Backup Restore
        api_response = api_instance.get_backup_restores(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->get_backup_restores: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetBackupRestores200Response**](GetBackupRestores200Response.md)

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

# **get_backup_results**
> GetBackupResults200Response get_backup_results(id)

Retrieves a Specific Backup Result

Retrieves a specific backup result. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.get_backup_results200_response import GetBackupResults200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Backup Result
        api_response = api_instance.get_backup_results(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->get_backup_results: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetBackupResults200Response**](GetBackupResults200Response.md)

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

# **get_backups**
> AddBackups200ResponseAllOf get_backups(id)

Retrieves a Specific Backup

Retrieves a specific backup. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backups200_response_all_of import AddBackups200ResponseAllOf
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Backup
        api_response = api_instance.get_backups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->get_backups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddBackups200ResponseAllOf**](AddBackups200ResponseAllOf.md)

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

# **list_backup_jobs**
> ListBackupJobs200Response list_backup_jobs()

Retrieves all Backup Jobs

Retrieves all backup jobs. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.list_backup_jobs200_response import ListBackupJobs200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    external_id = "externalId_example" # str | Filter by External ID (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Backup Jobs
        api_response = api_instance.list_backup_jobs(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, external_id=external_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->list_backup_jobs: %s\n" % e)
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
 **code** | **str**| If specified will return an exact match on code | [optional]
 **external_id** | **str**| Filter by External ID | [optional]

### Return type

[**ListBackupJobs200Response**](ListBackupJobs200Response.md)

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

# **list_backup_restores**
> ListBackupRestores200Response list_backup_restores()

Retrieves all Backup Restores

Retrieves all backup restores. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_backup_restores200_response import ListBackupRestores200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "sort_example" # str | Sort order, the name of the property to sort by. The default sort order is `startDate` descending (optional)
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "test-backup" # str | Filter by backup name (optional)
    phrase = "test" # str | Search phrase for partial matches on backup name, wildcard may be specified as %, eg. example-%  (optional)
    backup_id = 1 # int | Filter by backup ID (optional)
    backup_result_id = "backupResultId_example" # str | Filter by backup result ID (optional)
    container_id = 1 # int | Filter by container ID (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Backup Restores
        api_response = api_instance.list_backup_restores(max=max, offset=offset, sort=sort, direction=direction, name=name, phrase=phrase, backup_id=backup_id, backup_result_id=backup_result_id, container_id=container_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->list_backup_restores: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by. The default sort order is &#x60;startDate&#x60; descending | [optional]
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by backup name | [optional]
 **phrase** | **str**| Search phrase for partial matches on backup name, wildcard may be specified as %, eg. example-%  | [optional]
 **backup_id** | **int**| Filter by backup ID | [optional]
 **backup_result_id** | **str**| Filter by backup result ID | [optional]
 **container_id** | **int**| Filter by container ID | [optional]

### Return type

[**ListBackupRestores200Response**](ListBackupRestores200Response.md)

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

# **list_backup_results**
> ListBackupResults200Response list_backup_results()

Retrieves all Backup Results

Retrieves all backup results. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.list_backup_results200_response import ListBackupResults200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "test-backup" # str | Filter by backupName (optional)
    phrase = "test" # str | Search phrase for partial matches on backupName, wildcard may be specified as %, eg. example-%  (optional)
    backup_id = 1 # int | Filter by backup ID (optional)
    backup_set_id = "backupSetId_example" # str | Filter by backupSetId (optional)
    instance_id = 1 # int | Filter by instance ID (optional)
    container_id = 1 # int | Filter by container ID (optional)
    server_id = 1 # int | Filter by server ID (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Backup Results
        api_response = api_instance.list_backup_results(max=max, offset=offset, sort=sort, direction=direction, name=name, phrase=phrase, backup_id=backup_id, backup_set_id=backup_set_id, instance_id=instance_id, container_id=container_id, server_id=server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->list_backup_results: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by backupName | [optional]
 **phrase** | **str**| Search phrase for partial matches on backupName, wildcard may be specified as %, eg. example-%  | [optional]
 **backup_id** | **int**| Filter by backup ID | [optional]
 **backup_set_id** | **str**| Filter by backupSetId | [optional]
 **instance_id** | **int**| Filter by instance ID | [optional]
 **container_id** | **int**| Filter by container ID | [optional]
 **server_id** | **int**| Filter by server ID | [optional]

### Return type

[**ListBackupResults200Response**](ListBackupResults200Response.md)

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

# **list_backups**
> ListBackups200Response list_backups()

Retrieves all Backups

Retrieves all backups. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_backups200_response import ListBackups200Response
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
    api_instance = backups_api.BackupsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Backups
        api_response = api_instance.list_backups(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->list_backups: %s\n" % e)
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

### Return type

[**ListBackups200Response**](ListBackups200Response.md)

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

# **remove_backup_jobs**
> Model200Success remove_backup_jobs(id)

Deletes a Backup Job

Deletes a specified backup job. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Backup Job
        api_response = api_instance.remove_backup_jobs(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->remove_backup_jobs: %s\n" % e)
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

# **remove_backup_restores**
> Model200Success remove_backup_restores(id)

Deletes a Backup Restore

Deletes a specified backup restore. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Backup Restore
        api_response = api_instance.remove_backup_restores(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->remove_backup_restores: %s\n" % e)
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

# **remove_backup_results**
> Model200Success remove_backup_results(id)

Deletes a Backup Result

Deletes a specified backup result. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Backup Result
        api_response = api_instance.remove_backup_results(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->remove_backup_results: %s\n" % e)
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

# **remove_backups**
> Model200Success remove_backups(id)

Deletes a Backup

Deletes a specified backup. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Backup
        api_response = api_instance.remove_backups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->remove_backups: %s\n" % e)
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

# **update_backup_jobs**
> AddBackupJobs200Response update_backup_jobs(id)

Updates a Backup Job

Updates a backup job. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backup_jobs200_response import AddBackupJobs200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_backup_jobs_request import UpdateBackupJobsRequest
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_backup_jobs_request = UpdateBackupJobsRequest(
        job=UpdateBackupJobsRequestJob(
            name="name_example",
            code="code_example",
            schedule_id=1,
            retention_count=1,
        ),
    ) # UpdateBackupJobsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Backup Job
        api_response = api_instance.update_backup_jobs(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->update_backup_jobs: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Backup Job
        api_response = api_instance.update_backup_jobs(id, update_backup_jobs_request=update_backup_jobs_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->update_backup_jobs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_backup_jobs_request** | [**UpdateBackupJobsRequest**](UpdateBackupJobsRequest.md)|  | [optional]

### Return type

[**AddBackupJobs200Response**](AddBackupJobs200Response.md)

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

# **update_backups**
> AddBackups200Response update_backups(id)

Updates a Backup

Updates a backup. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import backups_api
from openapi_client.model.add_backups200_response import AddBackups200Response
from openapi_client.model.update_backups_request import UpdateBackupsRequest
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
    api_instance = backups_api.BackupsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_backups_request = UpdateBackupsRequest(
        backup=UpdateBackupsRequestBackup(
            name="name_example",
            job_id=1,
            enabled=True,
        ),
    ) # UpdateBackupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Backup
        api_response = api_instance.update_backups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->update_backups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Backup
        api_response = api_instance.update_backups(id, update_backups_request=update_backups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BackupsApi->update_backups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_backups_request** | [**UpdateBackupsRequest**](UpdateBackupsRequest.md)|  | [optional]

### Return type

[**AddBackups200Response**](AddBackups200Response.md)

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

