# openapi_client.IncidentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_incident**](IncidentsApi.md#add_incident) | **POST** /api/monitoring/incidents | Create a New Incident
[**delete_incidents**](IncidentsApi.md#delete_incidents) | **DELETE** /api/monitoring/incidents/{id} | Close a Specific Incident
[**get_incidents**](IncidentsApi.md#get_incidents) | **GET** /api/monitoring/incidents/{id} | Get a Specific Incident
[**list_incidents**](IncidentsApi.md#list_incidents) | **GET** /api/monitoring/incidents | List All Incidents
[**update_incidents**](IncidentsApi.md#update_incidents) | **PUT** /api/monitoring/incidents/{id} | Update Incident
[**update_incidents_reopen**](IncidentsApi.md#update_incidents_reopen) | **GET** /api/monitoring/incidents/{id}/reopen | Reopen a Specific Incident
[**update_mute_all_incidents**](IncidentsApi.md#update_mute_all_incidents) | **PUT** /api/monitoring/incidents/mute-all | Mute All Incidents
[**update_mute_incidents**](IncidentsApi.md#update_mute_incidents) | **PUT** /api/monitoring/incidents/{id}/mute | Mute Incident


# **add_incident**
> AddIncident200Response add_incident()

Create a New Incident

Create a new incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.add_incident200_response import AddIncident200Response
from openapi_client.model.add_incident_request import AddIncidentRequest
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
    api_instance = incidents_api.IncidentsApi(api_client)
    add_incident_request = AddIncidentRequest(
        incident=AddIncidentRequestIncident(
            resolution="I plugged it back in",
            comment="This is a summary of the incident",
            status="open",
            severity="info",
            name="Incident Name",
            start_date=dateutil_parser('2019-10-20T19:42:00Z'),
            end_date=dateutil_parser('2019-10-20T19:42:00Z'),
            in_uptime=True,
        ),
    ) # AddIncidentRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a New Incident
        api_response = api_instance.add_incident(add_incident_request=add_incident_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->add_incident: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_incident_request** | [**AddIncidentRequest**](AddIncidentRequest.md)|  | [optional]

### Return type

[**AddIncident200Response**](AddIncident200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_incidents**
> Model200Success delete_incidents(id)

Close a Specific Incident

Close an existing monitoring incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
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
    api_instance = incidents_api.IncidentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Close a Specific Incident
        api_response = api_instance.delete_incidents(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->delete_incidents: %s\n" % e)
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

# **get_incidents**
> AddIncident200ResponseAllOf get_incidents(id)

Get a Specific Incident

Get details about a specific incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_incident200_response_all_of import AddIncident200ResponseAllOf
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
    api_instance = incidents_api.IncidentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Incident
        api_response = api_instance.get_incidents(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->get_incidents: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddIncident200ResponseAllOf**](AddIncident200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_incidents**
> ListIncidents200Response list_incidents()

List All Incidents

Get a list of monitoring incidents.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.list_incidents200_response import ListIncidents200Response
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
    api_instance = incidents_api.IncidentsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    status = "running" # str | The instance status for filtering. (optional)
    severity = "INFO" # str | Filter by severity (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Incidents
        api_response = api_instance.list_incidents(max=max, offset=offset, status=status, severity=severity)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->list_incidents: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **status** | **str**| The instance status for filtering. | [optional]
 **severity** | **str**| Filter by severity | [optional]

### Return type

[**ListIncidents200Response**](ListIncidents200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Incidents |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_incidents**
> AddIncident200Response update_incidents(id)

Update Incident

Update an existing incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.update_incidents_request import UpdateIncidentsRequest
from openapi_client.model.add_incident200_response import AddIncident200Response
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
    api_instance = incidents_api.IncidentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_incidents_request = UpdateIncidentsRequest(
        incident=UpdateIncidentsRequestIncident(
            resolution="I plugged it back in",
            comment="This is a summary of the incident",
            status="open",
            severity="info",
            name="Incident Name",
            start_date=dateutil_parser('2019-10-20T19:42:00Z'),
            end_date=dateutil_parser('2019-10-20T19:42:00Z'),
            in_uptime=True,
        ),
    ) # UpdateIncidentsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Incident
        api_response = api_instance.update_incidents(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->update_incidents: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Incident
        api_response = api_instance.update_incidents(id, update_incidents_request=update_incidents_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->update_incidents: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_incidents_request** | [**UpdateIncidentsRequest**](UpdateIncidentsRequest.md)|  | [optional]

### Return type

[**AddIncident200Response**](AddIncident200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_incidents_reopen**
> SuccessMessage update_incidents_reopen(id)

Reopen a Specific Incident

Get details about a specific incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.success_message import SuccessMessage
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
    api_instance = incidents_api.IncidentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Reopen a Specific Incident
        api_response = api_instance.update_incidents_reopen(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->update_incidents_reopen: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_all_incidents**
> UpdateMuteAllCheckApps200Response update_mute_all_incidents()

Mute All Incidents

Mute all existing incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_all_check_apps200_response import UpdateMuteAllCheckApps200Response
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
    api_instance = incidents_api.IncidentsApi(api_client)
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute All Incidents
        api_response = api_instance.update_mute_all_incidents(update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->update_mute_all_incidents: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteAllCheckApps200Response**](UpdateMuteAllCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mute All Incidents Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_mute_incidents**
> UpdateMuteCheckApps200Response update_mute_incidents(id)

Mute Incident

Mute an existing incident.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import incidents_api
from openapi_client.model.update_mute_all_check_apps_request import UpdateMuteAllCheckAppsRequest
from openapi_client.model.update_mute_check_apps200_response import UpdateMuteCheckApps200Response
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
    api_instance = incidents_api.IncidentsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_mute_all_check_apps_request = UpdateMuteAllCheckAppsRequest(
        muted=True,
    ) # UpdateMuteAllCheckAppsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Mute Incident
        api_response = api_instance.update_mute_incidents(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->update_mute_incidents: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Mute Incident
        api_response = api_instance.update_mute_incidents(id, update_mute_all_check_apps_request=update_mute_all_check_apps_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncidentsApi->update_mute_incidents: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_mute_all_check_apps_request** | [**UpdateMuteAllCheckAppsRequest**](UpdateMuteAllCheckAppsRequest.md)|  | [optional]

### Return type

[**UpdateMuteCheckApps200Response**](UpdateMuteCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Mute Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

