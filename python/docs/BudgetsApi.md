# openapi_client.BudgetsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_budgets**](BudgetsApi.md#add_budgets) | **POST** /api/budgets | Creates a Budget
[**get_budgets**](BudgetsApi.md#get_budgets) | **GET** /api/budgets/{id} | Retrieves a Specific Budget
[**list_budgets**](BudgetsApi.md#list_budgets) | **GET** /api/budgets | Retrieves all Budgets
[**remove_budgets**](BudgetsApi.md#remove_budgets) | **DELETE** /api/budgets/{id} | Deletes a Budget
[**update_budgets**](BudgetsApi.md#update_budgets) | **PUT** /api/budgets/{id} | Updates a Budget


# **add_budgets**
> AddBudgets200Response add_budgets()

Creates a Budget

Creates a budget. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import budgets_api
from openapi_client.model.add_budgets200_response import AddBudgets200Response
from openapi_client.model.add_budgets_request import AddBudgetsRequest
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
    api_instance = budgets_api.BudgetsApi(api_client)
    add_budgets_request = AddBudgetsRequest(
        budget=AddBudgetsRequestBudget(
            name="name_example",
            description="description_example",
            scope="account",
            period="year",
            year=2020,
            start_date=dateutil_parser('1970-01-01T00:00:00.00Z'),
            end_date=dateutil_parser('1970-01-01T00:00:00.00Z'),
            interval="year",
            scope_tenant_id=1,
            scope_group_id=1,
            scope_cloud_id=1,
            scope_user_id=1,
            costs=[100,100,110,120],
            enabled=True,
        ),
    ) # AddBudgetsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Budget
        api_response = api_instance.add_budgets(add_budgets_request=add_budgets_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BudgetsApi->add_budgets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_budgets_request** | [**AddBudgetsRequest**](AddBudgetsRequest.md)|  | [optional]

### Return type

[**AddBudgets200Response**](AddBudgets200Response.md)

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

# **get_budgets**
> GetBudgets200Response get_budgets(id)

Retrieves a Specific Budget

Retrieves a specific budget. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import budgets_api
from openapi_client.model.get_budgets200_response import GetBudgets200Response
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
    api_instance = budgets_api.BudgetsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Budget
        api_response = api_instance.get_budgets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BudgetsApi->get_budgets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetBudgets200Response**](GetBudgets200Response.md)

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

# **list_budgets**
> ListBudgets200Response list_budgets()

Retrieves all Budgets

Retrieves all budgets. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import budgets_api
from openapi_client.model.list_budgets200_response import ListBudgets200Response
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
    api_instance = budgets_api.BudgetsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Budgets
        api_response = api_instance.list_budgets(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BudgetsApi->list_budgets: %s\n" % e)
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

[**ListBudgets200Response**](ListBudgets200Response.md)

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

# **remove_budgets**
> Model200Success remove_budgets(id)

Deletes a Budget

Deletes a specified Budget. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import budgets_api
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
    api_instance = budgets_api.BudgetsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Budget
        api_response = api_instance.remove_budgets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BudgetsApi->remove_budgets: %s\n" % e)
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

# **update_budgets**
> AddBudgets200Response update_budgets(id)

Updates a Budget

Updates a budget. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import budgets_api
from openapi_client.model.add_budgets200_response import AddBudgets200Response
from openapi_client.model.add_budgets_request import AddBudgetsRequest
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
    api_instance = budgets_api.BudgetsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_budgets_request = AddBudgetsRequest(
        budget=AddBudgetsRequestBudget(
            name="name_example",
            description="description_example",
            scope="account",
            period="year",
            year=2020,
            start_date=dateutil_parser('1970-01-01T00:00:00.00Z'),
            end_date=dateutil_parser('1970-01-01T00:00:00.00Z'),
            interval="year",
            scope_tenant_id=1,
            scope_group_id=1,
            scope_cloud_id=1,
            scope_user_id=1,
            costs=[100,100,110,120],
            enabled=True,
        ),
    ) # AddBudgetsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Budget
        api_response = api_instance.update_budgets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BudgetsApi->update_budgets: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Budget
        api_response = api_instance.update_budgets(id, add_budgets_request=add_budgets_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BudgetsApi->update_budgets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_budgets_request** | [**AddBudgetsRequest**](AddBudgetsRequest.md)|  | [optional]

### Return type

[**AddBudgets200Response**](AddBudgets200Response.md)

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

