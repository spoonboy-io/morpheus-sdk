# openapi_client.ActivityApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_activity**](ActivityApi.md#list_activity) | **GET** /api/activity | Retrieves Activity


# **list_activity**
> ListActivity200Response list_activity()

Retrieves Activity

Retrieves activity. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import activity_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_activity200_response import ListActivity200Response
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
    api_instance = activity_api.ActivityApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    order = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    user_id = 6 # int | Filter by User ID (optional)
    tenant_id = 3 # float | Filter by Tenant ID. Only available to the master account. (optional)
    timeframe = "month" # str | Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional) if omitted the server will use the default value of "month"
    start = dateutil_parser('1970-01-01T00:00:00.00Z') # datetime | Filter by activity on or after a date(time). Default is 1 month prior (optional)
    end = dateutil_parser('1970-01-01T00:00:00.00Z') # datetime | Filter by activity on or before a date(time). Default is current date (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Activity
        api_response = api_instance.list_activity(max=max, offset=offset, sort=sort, order=order, phrase=phrase, name=name, user_id=user_id, tenant_id=tenant_id, timeframe=timeframe, start=start, end=end)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ActivityApi->list_activity: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **order** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **user_id** | **int**| Filter by User ID | [optional]
 **tenant_id** | **float**| Filter by Tenant ID. Only available to the master account. | [optional]
 **timeframe** | **str**| Filter by a timeframe (ex - today, yesterday, week, month, 3months) | [optional] if omitted the server will use the default value of "month"
 **start** | **datetime**| Filter by activity on or after a date(time). Default is 1 month prior | [optional]
 **end** | **datetime**| Filter by activity on or before a date(time). Default is current date | [optional]

### Return type

[**ListActivity200Response**](ListActivity200Response.md)

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

