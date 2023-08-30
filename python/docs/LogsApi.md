# openapi_client.LogsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_logs**](LogsApi.md#list_logs) | **GET** /api/logs | Retrieves Logs


# **list_logs**
> ListLogs200Response list_logs()

Retrieves Logs

Retrieves logs based on filters provided. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import logs_api
from openapi_client.model.list_logs200_response import ListLogs200Response
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
    api_instance = logs_api.LogsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    query = "query_example" # str | Alias for phrase (optional)
    message = "message_example" # str | Filter by message (optional)
    source_type = "sourceType_example" # str | Filter by source type (optional)
    type_code = "typeCode_example" # str | Filter by code type (optional)
    object_id = 15 # int | Filter by objectId (optional)
    token = "token_example" # str | Filter by token (optional)
    level = "ERROR" # str | Filter by log level. Multiple values can be passed pipe delimited. (optional)
    start_ms = 1657309873 # int | Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. (optional)
    end_ms = 1657309873 # int | Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. (optional)
    start_date_time = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Start Date timestamp (ISO 8601) (optional)
    end_date_time = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | End Date timestamp (ISO 8601) (optional)
    containers = 135 # int | The Container ID(s) for filtering. Accepts multiple values. (optional)
    servers = 135 # int | The Server ID(s) for filtering. Accepts multiple values. (optional)
    cluster_id = 135 # int | The Cluster ID(s) for filtering. Accepts multiple values. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Logs
        api_response = api_instance.list_logs(max=max, offset=offset, sort=sort, order=order, query=query, message=message, source_type=source_type, type_code=type_code, object_id=object_id, token=token, level=level, start_ms=start_ms, end_ms=end_ms, start_date_time=start_date_time, end_date_time=end_date_time, containers=containers, servers=servers, cluster_id=cluster_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LogsApi->list_logs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **query** | **str**| Alias for phrase | [optional]
 **message** | **str**| Filter by message | [optional]
 **source_type** | **str**| Filter by source type | [optional]
 **type_code** | **str**| Filter by code type | [optional]
 **object_id** | **int**| Filter by objectId | [optional]
 **token** | **str**| Filter by token | [optional]
 **level** | **str**| Filter by log level. Multiple values can be passed pipe delimited. | [optional]
 **start_ms** | **int**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. | [optional]
 **end_ms** | **int**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. | [optional]
 **start_date_time** | **datetime**| Start Date timestamp (ISO 8601) | [optional]
 **end_date_time** | **datetime**| End Date timestamp (ISO 8601) | [optional]
 **containers** | **int**| The Container ID(s) for filtering. Accepts multiple values. | [optional]
 **servers** | **int**| The Server ID(s) for filtering. Accepts multiple values. | [optional]
 **cluster_id** | **int**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional]

### Return type

[**ListLogs200Response**](ListLogs200Response.md)

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

