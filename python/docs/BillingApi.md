# openapi_client.BillingApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_billing_account**](BillingApi.md#get_billing_account) | **GET** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it
[**get_billing_instances_identifier**](BillingApi.md#get_billing_instances_identifier) | **GET** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.
[**get_billing_servers_identifier**](BillingApi.md#get_billing_servers_identifier) | **GET** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.
[**get_billing_zone_identifier**](BillingApi.md#get_billing_zone_identifier) | **GET** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.
[**list_billing_account**](BillingApi.md#list_billing_account) | **GET** /api/billing/account | Retrieves billing information for the requesting user&#39;s account.
[**list_billing_instances**](BillingApi.md#list_billing_instances) | **GET** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account.
[**list_billing_servers**](BillingApi.md#list_billing_servers) | **GET** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.
[**list_billing_zone**](BillingApi.md#list_billing_zone) | **GET** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account.


# **get_billing_account**
> ListBillingAccount200Response get_billing_account(id)

This endpoint will retrieve a specific account by id if the user has permission to access it

Will retrieve billing information for a specific tenant, if it is the current account or a sub account of the requesting user's account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.list_billing_account200_response import ListBillingAccount200Response
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
    api_instance = billing_api.BillingApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_compute_servers = True # bool | Optional ability to exclude compute servers (optional) if omitted the server will use the default value of True
    include_instances = True # bool | Optional ability to exclude instances (optional) if omitted the server will use the default value of True
    include_discovered_servers = True # bool | Optional ability to exclude discovered servers (optional) if omitted the server will use the default value of True
    include_load_balancers = True # bool | Optional ability to exclude load balancers (optional) if omitted the server will use the default value of True
    include_virtual_images = True # bool | Optional ability to exclude virtual images (optional) if omitted the server will use the default value of True
    include_snapshots = True # bool | Optional ability to exclude snapshots (optional) if omitted the server will use the default value of True

    # example passing only required values which don't have defaults set
    try:
        # This endpoint will retrieve a specific account by id if the user has permission to access it
        api_response = api_instance.get_billing_account(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_account: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # This endpoint will retrieve a specific account by id if the user has permission to access it
        api_response = api_instance.get_billing_account(id, start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_compute_servers=include_compute_servers, include_instances=include_instances, include_discovered_servers=include_discovered_servers, include_load_balancers=include_load_balancers, include_virtual_images=include_virtual_images, include_snapshots=include_snapshots)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_account: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] if omitted the server will use the default value of True
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] if omitted the server will use the default value of True
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] if omitted the server will use the default value of True
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] if omitted the server will use the default value of True
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] if omitted the server will use the default value of True
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] if omitted the server will use the default value of True

### Return type

[**ListBillingAccount200Response**](ListBillingAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_billing_instances_identifier**
> GetBillingInstancesIdentifier200Response get_billing_instances_identifier(identifier)

Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.get_billing_instances_identifier200_response import GetBillingInstancesIdentifier200Response
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
    api_instance = billing_api.BillingApi(api_client)
    identifier = "identifier_example" # str | Morpheus UUID or ID of the Object being created or referenced
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_tenants = False # bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) if omitted the server will use the default value of False
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.
        api_response = api_instance.get_billing_instances_identifier(identifier)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_instances_identifier: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.
        api_response = api_instance.get_billing_instances_identifier(identifier, start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_tenants=include_tenants, account_id=account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_instances_identifier: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **str**| Morpheus UUID or ID of the Object being created or referenced |
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] if omitted the server will use the default value of False
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

### Return type

[**GetBillingInstancesIdentifier200Response**](GetBillingInstancesIdentifier200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_billing_servers_identifier**
> GetBillingServersIdentifier200Response get_billing_servers_identifier(identifier)

Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.get_billing_servers_identifier200_response import GetBillingServersIdentifier200Response
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
    api_instance = billing_api.BillingApi(api_client)
    identifier = "identifier_example" # str | Morpheus UUID or ID of the Object being created or referenced
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_tenants = False # bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) if omitted the server will use the default value of False
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.
        api_response = api_instance.get_billing_servers_identifier(identifier)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_servers_identifier: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.
        api_response = api_instance.get_billing_servers_identifier(identifier, start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_tenants=include_tenants, account_id=account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_servers_identifier: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **str**| Morpheus UUID or ID of the Object being created or referenced |
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] if omitted the server will use the default value of False
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

### Return type

[**GetBillingServersIdentifier200Response**](GetBillingServersIdentifier200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_billing_zone_identifier**
> GetBillingZoneIdentifier200Response get_billing_zone_identifier(identifier)

Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.get_billing_zone_identifier200_response import GetBillingZoneIdentifier200Response
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
    api_instance = billing_api.BillingApi(api_client)
    identifier = "identifier_example" # str | Morpheus UUID or ID of the Object being created or referenced
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_compute_servers = True # bool | Optional ability to exclude compute servers (optional) if omitted the server will use the default value of True
    include_instances = True # bool | Optional ability to exclude instances (optional) if omitted the server will use the default value of True
    include_discovered_servers = True # bool | Optional ability to exclude discovered servers (optional) if omitted the server will use the default value of True
    include_load_balancers = True # bool | Optional ability to exclude load balancers (optional) if omitted the server will use the default value of True
    include_virtual_images = True # bool | Optional ability to exclude virtual images (optional) if omitted the server will use the default value of True
    include_snapshots = True # bool | Optional ability to exclude snapshots (optional) if omitted the server will use the default value of True

    # example passing only required values which don't have defaults set
    try:
        # Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.
        api_response = api_instance.get_billing_zone_identifier(identifier)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_zone_identifier: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.
        api_response = api_instance.get_billing_zone_identifier(identifier, start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_compute_servers=include_compute_servers, include_instances=include_instances, include_discovered_servers=include_discovered_servers, include_load_balancers=include_load_balancers, include_virtual_images=include_virtual_images, include_snapshots=include_snapshots)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->get_billing_zone_identifier: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **str**| Morpheus UUID or ID of the Object being created or referenced |
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] if omitted the server will use the default value of True
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] if omitted the server will use the default value of True
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] if omitted the server will use the default value of True
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] if omitted the server will use the default value of True
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] if omitted the server will use the default value of True
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] if omitted the server will use the default value of True

### Return type

[**GetBillingZoneIdentifier200Response**](GetBillingZoneIdentifier200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_billing_account**
> ListBillingAccount200Response list_billing_account()

Retrieves billing information for the requesting user's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.list_billing_account200_response import ListBillingAccount200Response
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
    api_instance = billing_api.BillingApi(api_client)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_compute_servers = True # bool | Optional ability to exclude compute servers (optional) if omitted the server will use the default value of True
    include_instances = True # bool | Optional ability to exclude instances (optional) if omitted the server will use the default value of True
    include_discovered_servers = True # bool | Optional ability to exclude discovered servers (optional) if omitted the server will use the default value of True
    include_load_balancers = True # bool | Optional ability to exclude load balancers (optional) if omitted the server will use the default value of True
    include_virtual_images = True # bool | Optional ability to exclude virtual images (optional) if omitted the server will use the default value of True
    include_snapshots = True # bool | Optional ability to exclude snapshots (optional) if omitted the server will use the default value of True

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for the requesting user's account.
        api_response = api_instance.list_billing_account(start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_compute_servers=include_compute_servers, include_instances=include_instances, include_discovered_servers=include_discovered_servers, include_load_balancers=include_load_balancers, include_virtual_images=include_virtual_images, include_snapshots=include_snapshots)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->list_billing_account: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] if omitted the server will use the default value of True
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] if omitted the server will use the default value of True
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] if omitted the server will use the default value of True
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] if omitted the server will use the default value of True
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] if omitted the server will use the default value of True
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] if omitted the server will use the default value of True

### Return type

[**ListBillingAccount200Response**](ListBillingAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_billing_instances**
> ListBillingInstances200Response list_billing_instances()

Retrieves billing information for all instances on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_billing_instances200_response import ListBillingInstances200Response
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
    api_instance = billing_api.BillingApi(api_client)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_tenants = False # bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) if omitted the server will use the default value of False
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for all instances on the requestor's account.
        api_response = api_instance.list_billing_instances(start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_tenants=include_tenants, account_id=account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->list_billing_instances: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] if omitted the server will use the default value of False
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

### Return type

[**ListBillingInstances200Response**](ListBillingInstances200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_billing_servers**
> ListBillingServers200Response list_billing_servers()

Retrieves billing information for all servers (container hosts) on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_billing_servers200_response import ListBillingServers200Response
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
    api_instance = billing_api.BillingApi(api_client)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_tenants = False # bool | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) if omitted the server will use the default value of False
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for all servers (container hosts) on the requestor's account.
        api_response = api_instance.list_billing_servers(start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_tenants=include_tenants, account_id=account_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->list_billing_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_tenants** | **bool**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] if omitted the server will use the default value of False
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

### Return type

[**ListBillingServers200Response**](ListBillingServers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_billing_zone**
> ListBillingZone200Response list_billing_zone()

Retrieves billing information for all zones on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import billing_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_billing_zone200_response import ListBillingZone200Response
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
    api_instance = billing_api.BillingApi(api_client)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    include_usages = True # bool | Optional ability to suppress the usage records (optional) if omitted the server will use the default value of True
    max_usages = 1 # int | Optional ability to limit the usages returned (optional)
    offset_usages = 1 # int | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
    include_compute_servers = True # bool | Optional ability to exclude compute servers (optional) if omitted the server will use the default value of True
    include_instances = True # bool | Optional ability to exclude instances (optional) if omitted the server will use the default value of True
    include_discovered_servers = True # bool | Optional ability to exclude discovered servers (optional) if omitted the server will use the default value of True
    include_load_balancers = True # bool | Optional ability to exclude load balancers (optional) if omitted the server will use the default value of True
    include_virtual_images = True # bool | Optional ability to exclude virtual images (optional) if omitted the server will use the default value of True
    include_snapshots = True # bool | Optional ability to exclude snapshots (optional) if omitted the server will use the default value of True

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves billing information for all zones on the requestor's account.
        api_response = api_instance.list_billing_zone(start_date=start_date, end_date=end_date, include_usages=include_usages, max_usages=max_usages, offset_usages=offset_usages, include_compute_servers=include_compute_servers, include_instances=include_instances, include_discovered_servers=include_discovered_servers, include_load_balancers=include_load_balancers, include_virtual_images=include_virtual_images, include_snapshots=include_snapshots)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling BillingApi->list_billing_zone: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **include_usages** | **bool**| Optional ability to suppress the usage records | [optional] if omitted the server will use the default value of True
 **max_usages** | **int**| Optional ability to limit the usages returned | [optional]
 **offset_usages** | **int**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional]
 **include_compute_servers** | **bool**| Optional ability to exclude compute servers | [optional] if omitted the server will use the default value of True
 **include_instances** | **bool**| Optional ability to exclude instances | [optional] if omitted the server will use the default value of True
 **include_discovered_servers** | **bool**| Optional ability to exclude discovered servers | [optional] if omitted the server will use the default value of True
 **include_load_balancers** | **bool**| Optional ability to exclude load balancers | [optional] if omitted the server will use the default value of True
 **include_virtual_images** | **bool**| Optional ability to exclude virtual images | [optional] if omitted the server will use the default value of True
 **include_snapshots** | **bool**| Optional ability to exclude snapshots | [optional] if omitted the server will use the default value of True

### Return type

[**ListBillingZone200Response**](ListBillingZone200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Billing |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

