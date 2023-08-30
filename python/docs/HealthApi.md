# openapi_client.HealthApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**acknowledge_health_alarm**](HealthApi.md#acknowledge_health_alarm) | **PUT** /api/health/alarms/{id}/acknowledge | Acknowledge a Health Alarm
[**acknowledge_health_alarms**](HealthApi.md#acknowledge_health_alarms) | **PUT** /api/health/alarms/acknowledge | Acknowledge Many Health Alarms
[**export_health_logs**](HealthApi.md#export_health_logs) | **GET** /api/health/logs/export | Export Appliance Health Logs
[**get_health_alarms**](HealthApi.md#get_health_alarms) | **GET** /api/health/alarms/{id} | Retrieves a Specific Appliance Health Alarm
[**list_health**](HealthApi.md#list_health) | **GET** /api/health | Retrieves Appliance Health
[**list_health_alarms**](HealthApi.md#list_health_alarms) | **GET** /api/health/alarms | Retrieves Appliance Health Alarms
[**list_health_logs**](HealthApi.md#list_health_logs) | **GET** /api/health/logs | Retrieves Appliance Health Logs
[**ping**](HealthApi.md#ping) | **GET** /api/ping | Basic information about current Morpheus Installation


# **acknowledge_health_alarm**
> Model200Success acknowledge_health_alarm(id)

Acknowledge a Health Alarm

Acknowledge a specific health alarm.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.acknowledge_health_alarm_request import AcknowledgeHealthAlarmRequest
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
    api_instance = health_api.HealthApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    acknowledge_health_alarm_request = AcknowledgeHealthAlarmRequest(
        alarm=AcknowledgeHealthAlarmRequestAlarm(
            acknowledged=True,
        ),
    ) # AcknowledgeHealthAlarmRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Acknowledge a Health Alarm
        api_response = api_instance.acknowledge_health_alarm(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->acknowledge_health_alarm: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Acknowledge a Health Alarm
        api_response = api_instance.acknowledge_health_alarm(id, acknowledge_health_alarm_request=acknowledge_health_alarm_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->acknowledge_health_alarm: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **acknowledge_health_alarm_request** | [**AcknowledgeHealthAlarmRequest**](AcknowledgeHealthAlarmRequest.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **acknowledge_health_alarms**
> Model200Success acknowledge_health_alarms()

Acknowledge Many Health Alarms

Acknowledge health alarms.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.acknowledge_health_alarms_request import AcknowledgeHealthAlarmsRequest
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
    api_instance = health_api.HealthApi(api_client)
    acknowledge_health_alarms_request = AcknowledgeHealthAlarmsRequest(
        alarm=AcknowledgeHealthAlarmsRequestAlarm(
            acknowledged=True,
            ids=[],
            all=False,
        ),
    ) # AcknowledgeHealthAlarmsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Acknowledge Many Health Alarms
        api_response = api_instance.acknowledge_health_alarms(acknowledge_health_alarms_request=acknowledge_health_alarms_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->acknowledge_health_alarms: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acknowledge_health_alarms_request** | [**AcknowledgeHealthAlarmsRequest**](AcknowledgeHealthAlarmsRequest.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **export_health_logs**
> file_type export_health_logs()

Export Appliance Health Logs

This endpoint downloads the morpheus appliance logs as a file attachment. By default, the most recent 10,000 log entries are returned, with the newest at the end of the file. The format for each log entry is `timestamp` `level` `message`.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
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
    api_instance = health_api.HealthApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    acknowledged = False # bool | True or False flag for Acknowledged items (optional)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    reverse = True # bool | Reverse order of records. This `true` by default when sort and direction are not passed, but `false` by default if either is passed. This means that by default the newest log entries are the bottom of the file.  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Export Appliance Health Logs
        api_response = api_instance.export_health_logs(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, acknowledged=acknowledged, start_date=start_date, end_date=end_date, reverse=reverse)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->export_health_logs: %s\n" % e)
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
 **acknowledged** | **bool**| True or False flag for Acknowledged items | [optional]
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **reverse** | **bool**| Reverse order of records. This &#x60;true&#x60; by default when sort and direction are not passed, but &#x60;false&#x60; by default if either is passed. This means that by default the newest log entries are the bottom of the file.  | [optional]

### Return type

**file_type**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  * Content-Disposition -  <br>  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_health_alarms**
> GetHealthAlarms200Response get_health_alarms(id)

Retrieves a Specific Appliance Health Alarm

This endpoint will retrieve a specific health alarm by ID.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.get_health_alarms200_response import GetHealthAlarms200Response
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
    api_instance = health_api.HealthApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Appliance Health Alarm
        api_response = api_instance.get_health_alarms(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->get_health_alarms: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetHealthAlarms200Response**](GetHealthAlarms200Response.md)

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

# **list_health**
> ListHealth200Response list_health()

Retrieves Appliance Health

This endpoint retrieves health info about the appliance such as cpu, memory and database usage. Elasticsearch statistics and queue usage are also returned.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.list_health200_response import ListHealth200Response
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
    api_instance = health_api.HealthApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves Appliance Health
        api_response = api_instance.list_health()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->list_health: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListHealth200Response**](ListHealth200Response.md)

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

# **list_health_alarms**
> ListHealthAlarms200Response list_health_alarms()

Retrieves Appliance Health Alarms

This endpoint retrieves all health alarms, which are Operation notifications from Cloud and other Service Integrations. These alarms are not generated by the appliance, but synced and displayed for visibility. By default only open alarms are returned. Open alarms are those that have not yet been acknowledged.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.list_health_alarms200_response import ListHealthAlarms200Response
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
    api_instance = health_api.HealthApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    acknowledged = False # bool | True or False flag for Acknowledged items (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Appliance Health Alarms
        api_response = api_instance.list_health_alarms(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, acknowledged=acknowledged)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->list_health_alarms: %s\n" % e)
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
 **acknowledged** | **bool**| True or False flag for Acknowledged items | [optional]

### Return type

[**ListHealthAlarms200Response**](ListHealthAlarms200Response.md)

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

# **list_health_logs**
> ListHealthLogs200Response list_health_logs()

Retrieves Appliance Health Logs

This endpoint retrieves all health logs. These are the logs of the remote appliance itself. These logs show all ui activity and are useful for troubleshooting and auditing. Stack traces are filtered for Morpheus services. Complete stack traces can be found in `/var/log/morpheus/morpheus-ui/current`.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.list_health_logs200_response import ListHealthLogs200Response
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
    api_instance = health_api.HealthApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    acknowledged = False # bool | True or False flag for Acknowledged items (optional)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves Appliance Health Logs
        api_response = api_instance.list_health_logs(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, acknowledged=acknowledged, start_date=start_date, end_date=end_date)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->list_health_logs: %s\n" % e)
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
 **acknowledged** | **bool**| True or False flag for Acknowledged items | [optional]
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]

### Return type

[**ListHealthLogs200Response**](ListHealthLogs200Response.md)

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

# **ping**
> Ping ping()

Basic information about current Morpheus Installation

This endpoint can be used to check the remote appliance build version and some other basic information.  This is an unsecured endpoint and does not require authorization. However, build version will not be returned unless you are authenticated with a valid access token. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import health_api
from openapi_client.model.ping import Ping
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
    api_instance = health_api.HealthApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Basic information about current Morpheus Installation
        api_response = api_instance.ping()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HealthApi->ping: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**Ping**](Ping.md)

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

