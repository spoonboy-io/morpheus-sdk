# openapi_client.SearchApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**search**](SearchApi.md#search) | **GET** /api/search | Provides global search for all types of objects


# **search**
> Search search()

Provides global search for all types of objects

This endpoint provides global search for all types of objects, with the results sorted in order of relevance.  The `phrase` parameter can be specified as part of the URL or as a query parameter. If phrase is not specified, then 0 results (hits) will be returned. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import search_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.search import Search
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
    api_instance = search_api.SearchApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    query = "query_example" # str | Alias for phrase (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Provides global search for all types of objects
        api_response = api_instance.search(max=max, offset=offset, phrase=phrase, query=query)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling SearchApi->search: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **query** | **str**| Alias for phrase | [optional]

### Return type

[**Search**](Search.md)

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

