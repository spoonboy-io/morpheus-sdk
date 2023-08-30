# openapi_client.AlertsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_alerts**](AlertsApi.md#add_alerts) | **POST** /api/monitoring/alerts | Create a New Alert
[**delete_alerts**](AlertsApi.md#delete_alerts) | **DELETE** /api/monitoring/alerts/{id} | Delete a Specific Alert
[**get_alerts**](AlertsApi.md#get_alerts) | **GET** /api/monitoring/alerts/{id} | Get a Specific Alert
[**list_alerts**](AlertsApi.md#list_alerts) | **GET** /api/monitoring/alerts | List All Alerts
[**update_alerts**](AlertsApi.md#update_alerts) | **PUT** /api/monitoring/alerts/{id} | Update Alert


# **add_alerts**
> AddAlerts200Response add_alerts()

Create a New Alert

Create a new monitoring alert.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import alerts_api
from openapi_client.model.add_alerts200_response import AddAlerts200Response
from openapi_client.model.add_alerts_request import AddAlertsRequest
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
    api_instance = alerts_api.AlertsApi(api_client)
    add_alerts_request = AddAlertsRequest(
        alert=AddAlertsRequestAlert(
            name="My Alert",
            min_duration=0,
            min_severity="critical",
            active=True,
            all_checks=False,
            all_groups=False,
            all_apps=False,
            checks=[
                1,
            ],
            groups=[
                1,
            ],
            apps=[
                1,
            ],
            contacts=[
                AddAlertsRequestAlertContactsInner(
                    id=1,
                    name="name_example",
                    method="method_example",
                    notify=True,
                    close=True,
                ),
            ],
        ),
    ) # AddAlertsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Alert
        api_response = api_instance.add_alerts(add_alerts_request=add_alerts_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AlertsApi->add_alerts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_alerts_request** | [**AddAlertsRequest**](AddAlertsRequest.md)|  | [optional]

### Return type

[**AddAlerts200Response**](AddAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Alert Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_alerts**
> Model200Success delete_alerts(id)

Delete a Specific Alert

Delete an existing monitoring alert.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import alerts_api
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
    api_instance = alerts_api.AlertsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Specific Alert
        api_response = api_instance.delete_alerts(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AlertsApi->delete_alerts: %s\n" % e)
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

# **get_alerts**
> GetAlerts200Response get_alerts(id)

Get a Specific Alert

Get details about a specific monitoring alert.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import alerts_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_alerts200_response import GetAlerts200Response
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
    api_instance = alerts_api.AlertsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Alert
        api_response = api_instance.get_alerts(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AlertsApi->get_alerts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetAlerts200Response**](GetAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Alert Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_alerts**
> ListAlerts200Response list_alerts()

List All Alerts

Get a list of monitoring alerts.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import alerts_api
from openapi_client.model.list_alerts200_response import ListAlerts200Response
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
    api_instance = alerts_api.AlertsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Alerts
        api_response = api_instance.list_alerts(max=max, offset=offset, last_updated=last_updated)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AlertsApi->list_alerts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

### Return type

[**ListAlerts200Response**](ListAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Alerts |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_alerts**
> UpdateAlerts200Response update_alerts(id)

Update Alert

Update an existing monitoring alert.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import alerts_api
from openapi_client.model.update_alerts_request import UpdateAlertsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_alerts200_response import UpdateAlerts200Response
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
    api_instance = alerts_api.AlertsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_alerts_request = UpdateAlertsRequest(
        alert=UpdateAlertsRequestAlert(
            name="My Alert",
            min_duration=0,
            min_severity="critical",
            active=True,
            all_checks=False,
            all_groups=False,
            all_apps=False,
            checks=[
                1,
            ],
            groups=[
                1,
            ],
            apps=[
                1,
            ],
            contacts=[
                AddAlertsRequestAlertContactsInner(
                    id=1,
                    name="name_example",
                    method="method_example",
                    notify=True,
                    close=True,
                ),
            ],
        ),
    ) # UpdateAlertsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Alert
        api_response = api_instance.update_alerts(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AlertsApi->update_alerts: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Alert
        api_response = api_instance.update_alerts(id, update_alerts_request=update_alerts_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AlertsApi->update_alerts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_alerts_request** | [**UpdateAlertsRequest**](UpdateAlertsRequest.md)|  | [optional]

### Return type

[**UpdateAlerts200Response**](UpdateAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Alert Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

