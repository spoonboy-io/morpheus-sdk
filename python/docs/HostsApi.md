# openapi_client.HostsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_host**](HostsApi.md#get_host) | **GET** /api/servers/{id} | Get a Specific Host
[**get_host_snpshots**](HostsApi.md#get_host_snpshots) | **GET** /api/servers/{id}/snapshots | Get list of snapshots for a Host
[**get_host_type**](HostsApi.md#get_host_type) | **GET** /api/server-types/{id} | Get a Specific Host Type
[**get_wiki_server**](HostsApi.md#get_wiki_server) | **GET** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**list_host_types**](HostsApi.md#list_host_types) | **GET** /api/server-types | Host Types
[**list_hosts**](HostsApi.md#list_hosts) | **GET** /api/servers | Get All Hosts
[**list_server_service_plans**](HostsApi.md#list_server_service_plans) | **GET** /api/servers/service-plans | Get Available Service Plans for a Host
[**remove_host**](HostsApi.md#remove_host) | **DELETE** /api/servers/{id} | Delete a Host
[**restart_host**](HostsApi.md#restart_host) | **PUT** /api/servers/{id}/restart | Restart a Host
[**start_host**](HostsApi.md#start_host) | **PUT** /api/servers/{id}/start | Start a Host
[**stop_host**](HostsApi.md#stop_host) | **PUT** /api/servers/{id}/stop | Stop a Host
[**update_host**](HostsApi.md#update_host) | **PUT** /api/servers/{id} | Updating a Host
[**update_host_assign_tenant**](HostsApi.md#update_host_assign_tenant) | **PUT** /api/servers/{id}/assign-account | Assign To Tenant
[**update_host_cloud**](HostsApi.md#update_host_cloud) | **PUT** /api/servers/change-cloud | Change Server Cloud
[**update_host_execute_workflow**](HostsApi.md#update_host_execute_workflow) | **PUT** /api/servers/{id}/workflow | Run Workflow on a Host
[**update_host_install_agent**](HostsApi.md#update_host_install_agent) | **PUT** /api/servers/{id}/install-agent | Install Agent
[**update_host_managed**](HostsApi.md#update_host_managed) | **PUT** /api/servers/{id}/make-managed | Convert To Managed
[**update_host_resize**](HostsApi.md#update_host_resize) | **PUT** /api/servers/{id}/resize | Resize a Host
[**update_host_upgrade_agent**](HostsApi.md#update_host_upgrade_agent) | **PUT** /api/servers/{id}/upgrade | Upgrade Agent
[**update_server_network_interface**](HostsApi.md#update_server_network_interface) | **PUT** /api/servers/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for a Server&#39;s Network
[**update_wiki_server**](HostsApi.md#update_wiki_server) | **PUT** /api/servers/{id}/wiki | Update a Server Wiki Page


# **get_host**
> GetHost200Response get_host(id)

Get a Specific Host

This endpoint retrieves a specific host.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.get_host200_response import GetHost200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Host
        api_response = api_instance.get_host(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->get_host: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetHost200Response**](GetHost200Response.md)

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

# **get_host_snpshots**
> GetHostSnpshots200Response get_host_snpshots(id)

Get list of snapshots for a Host

Get list of snapshots for a Host

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.get_host_snpshots200_response import GetHostSnpshots200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get list of snapshots for a Host
        api_response = api_instance.get_host_snpshots(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->get_host_snpshots: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetHostSnpshots200Response**](GetHostSnpshots200Response.md)

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

# **get_host_type**
> GetHostType200Response get_host_type(id)

Get a Specific Host Type

This endpoint will retrieve a specific host type by id

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.get_host_type200_response import GetHostType200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Host Type
        api_response = api_instance.get_host_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->get_host_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetHostType200Response**](GetHostType200Response.md)

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

# **get_wiki_server**
> GetWikiApp200Response get_wiki_server(id)

Retrieves a Server Wiki Page

This endpoint retrieves a server Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.get_wiki_app200_response import GetWikiApp200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Server Wiki Page
        api_response = api_instance.get_wiki_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->get_wiki_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWikiApp200Response**](GetWikiApp200Response.md)

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

# **list_host_types**
> ListHostTypes200Response list_host_types()

Host Types

Fetch a paginated list of available host types. This returns the configuration options for each type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_host_types200_response import ListHostTypes200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    provision_type = "AKS" # str | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    zone_type = "alibaba" # str | Filter by Cloud Type code. (optional)
    creatable = True # bool | Filter by creatable flag. This is whether or not it can be provisioned. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Host Types
        api_response = api_instance.list_host_types(max=max, offset=offset, sort=sort, direction=direction, name=name, code=code, phrase=phrase, provision_type=provision_type, zone_type=zone_type, creatable=creatable)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->list_host_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **provision_type** | **str**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional]
 **zone_type** | **str**| Filter by Cloud Type code. | [optional]
 **creatable** | **bool**| Filter by creatable flag. This is whether or not it can be provisioned. | [optional]

### Return type

[**ListHostTypes200Response**](ListHostTypes200Response.md)

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

# **list_hosts**
> ListHosts200Response list_hosts()

Get All Hosts

This endpoint retrieves a paginated list of hosts.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_hosts200_response import ListHosts200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    site_id = 7 # int | The Site ID for Filtering (optional)
    cluster_id = 135 # int | The Cluster ID(s) for filtering. Accepts multiple values. (optional)
    managed = False # bool | Filter by managed (true) or unmanaged (false) (optional)
    server_type = "vmwareHypervisor" # str | Filter by server type code (optional)
    power_state = "off" # str | Filter by power status (optional)
    ip = "192.168.1.45" # str | Filter by IP address (optional)
    vm = False # bool | Filter to show only Virtual Machines (true) (optional)
    vm_hypervisor = False # bool | Filter to show only VM Hypervisors (true) (optional)
    bare_metal_host = False # bool | Filter to show only Baremetal Servers (optional)
    status = "status_example" # str | Filter by status (optional)
    agent_installed = True # bool | Filter by agent installed (true) (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    created_by = 63 # int | The User ID for Filtering (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    tags = "tags.env=qa&tags.env=test" # str | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
    metadata = "tags.env=qa&tags.env=test" # str | Alias for tags (optional)
    uuid = "uuid_example" # str | Filter by UUID (optional)
    external_id = "externalId_example" # str | Filter by External ID (optional)
    internal_id = "internalId_example" # str | Filter by Internal ID (optional)
    external_uniquel_id = "externalUniquelId_example" # str | Filter by External Unique ID (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Hosts
        api_response = api_instance.list_hosts(name=name, phrase=phrase, zone_id=zone_id, site_id=site_id, cluster_id=cluster_id, managed=managed, server_type=server_type, power_state=power_state, ip=ip, vm=vm, vm_hypervisor=vm_hypervisor, bare_metal_host=bare_metal_host, status=status, agent_installed=agent_installed, max=max, offset=offset, last_updated=last_updated, created_by=created_by, labels=labels, all_labels=all_labels, tags=tags, metadata=metadata, uuid=uuid, external_id=external_id, internal_id=internal_id, external_uniquel_id=external_uniquel_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->list_hosts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]
 **cluster_id** | **int**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional]
 **managed** | **bool**| Filter by managed (true) or unmanaged (false) | [optional]
 **server_type** | **str**| Filter by server type code | [optional]
 **power_state** | **str**| Filter by power status | [optional]
 **ip** | **str**| Filter by IP address | [optional]
 **vm** | **bool**| Filter to show only Virtual Machines (true) | [optional]
 **vm_hypervisor** | **bool**| Filter to show only VM Hypervisors (true) | [optional]
 **bare_metal_host** | **bool**| Filter to show only Baremetal Servers | [optional]
 **status** | **str**| Filter by status | [optional]
 **agent_installed** | **bool**| Filter by agent installed (true) | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **created_by** | **int**| The User ID for Filtering | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **tags** | **str**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional]
 **metadata** | **str**| Alias for tags | [optional]
 **uuid** | **str**| Filter by UUID | [optional]
 **external_id** | **str**| Filter by External ID | [optional]
 **internal_id** | **str**| Filter by Internal ID | [optional]
 **external_uniquel_id** | **str**| Filter by External Unique ID | [optional]

### Return type

[**ListHosts200Response**](ListHosts200Response.md)

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

# **list_server_service_plans**
> ListServerServicePlans200Response list_server_service_plans(zone_id)

Get Available Service Plans for a Host

This endpoint retrieves all the Service Plans available for the specified cloud and host type. It may be used to get the list of available plans when creating a new host or resizing an existing host.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.list_server_service_plans200_response import ListServerServicePlans200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    zone_id = 3 # int | The Zone ID for Filtering
    server_type_id = 5 # int | The ID of the Host Type (optional)
    site_id = 7 # int | The Site ID for Filtering (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get Available Service Plans for a Host
        api_response = api_instance.list_server_service_plans(zone_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->list_server_service_plans: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get Available Service Plans for a Host
        api_response = api_instance.list_server_service_plans(zone_id, server_type_id=server_type_id, site_id=site_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->list_server_service_plans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **int**| The Zone ID for Filtering |
 **server_type_id** | **int**| The ID of the Host Type | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]

### Return type

[**ListServerServicePlans200Response**](ListServerServicePlans200Response.md)

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

# **remove_host**
> Model200Success remove_host(id)

Delete a Host

Will delete a host asynchronously.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    remove_resources = "off" # str | Remove Resources (optional) if omitted the server will use the default value of "on"
    remove_instances = "on" # str | Remove Instances (optional) if omitted the server will use the default value of "off"
    preserve_volumes = "on" # str | Preserve Volumes (optional) if omitted the server will use the default value of "off"
    release_floating_ips = "off" # str | Release Floating IPs (optional) if omitted the server will use the default value of "on"
    release_eips = "off" # str | Alias for releaseFloatingIps (optional) if omitted the server will use the default value of "on"
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Host
        api_response = api_instance.remove_host(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->remove_host: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Host
        api_response = api_instance.remove_host(id, remove_resources=remove_resources, remove_instances=remove_instances, preserve_volumes=preserve_volumes, release_floating_ips=release_floating_ips, release_eips=release_eips, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->remove_host: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_resources** | **str**| Remove Resources | [optional] if omitted the server will use the default value of "on"
 **remove_instances** | **str**| Remove Instances | [optional] if omitted the server will use the default value of "off"
 **preserve_volumes** | **str**| Preserve Volumes | [optional] if omitted the server will use the default value of "off"
 **release_floating_ips** | **str**| Release Floating IPs | [optional] if omitted the server will use the default value of "on"
 **release_eips** | **str**| Alias for releaseFloatingIps | [optional] if omitted the server will use the default value of "on"
 **force** | **str**| Force Delete | [optional] if omitted the server will use the default value of "off"

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

# **restart_host**
> UpdateHostAssignTenant200Response restart_host(id)

Restart a Host

This will restart a host.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_assign_tenant200_response import UpdateHostAssignTenant200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Restart a Host
        api_response = api_instance.restart_host(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->restart_host: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**UpdateHostAssignTenant200Response**](UpdateHostAssignTenant200Response.md)

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

# **start_host**
> Model200Success start_host(id)

Start a Host

This will start a host.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Start a Host
        api_response = api_instance.start_host(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->start_host: %s\n" % e)
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

# **stop_host**
> Model200Success stop_host(id)

Stop a Host

This will stop a host.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Stop a Host
        api_response = api_instance.stop_host(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->stop_host: %s\n" % e)
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

# **update_host**
> GetHost200Response update_host(id)

Updating a Host

Updating a Host

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_request import UpdateHostRequest
from openapi_client.model.get_host200_response import GetHost200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_host_request = UpdateHostRequest(
        server=HostUpdate(
            name="myserver",
            description="my new server",
            ssh_username="myuser",
            ssh_password="mypassword",
            power_schedule_type=1,
            labels=[
                "labels_example",
            ],
            tags=[
                UpdateHostManagedRequestServerTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            add_tags=[
                UpdateHostManagedRequestServerTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            remove_tags=[
                InstanceUpdateInstanceRemoveTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            guest_console_type="guest_console_type_example",
            guest_console_username="guest_console_username_example",
            guest_console_password="guest_console_password_example",
            guest_console_port="guest_console_port_example",
            guest_console_preferred=True,
        ),
    ) # UpdateHostRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a Host
        api_response = api_instance.update_host(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a Host
        api_response = api_instance.update_host(id, update_host_request=update_host_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_host_request** | [**UpdateHostRequest**](UpdateHostRequest.md)|  | [optional]

### Return type

[**GetHost200Response**](GetHost200Response.md)

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

# **update_host_assign_tenant**
> UpdateHostAssignTenant200Response update_host_assign_tenant(id)

Assign To Tenant

This will change the ownership of the host to the specified Tenant account. This is only available to Master Tenant users.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_assign_tenant200_response import UpdateHostAssignTenant200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_host_assign_tenant_request import UpdateHostAssignTenantRequest
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    account_id = 3 # int | ID of the Tenant (optional)
    update_host_assign_tenant_request = UpdateHostAssignTenantRequest(
        move_associated_instances=True,
    ) # UpdateHostAssignTenantRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Assign To Tenant
        api_response = api_instance.update_host_assign_tenant(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_assign_tenant: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Assign To Tenant
        api_response = api_instance.update_host_assign_tenant(id, account_id=account_id, update_host_assign_tenant_request=update_host_assign_tenant_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_assign_tenant: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **account_id** | **int**| ID of the Tenant | [optional]
 **update_host_assign_tenant_request** | [**UpdateHostAssignTenantRequest**](UpdateHostAssignTenantRequest.md)|  | [optional]

### Return type

[**UpdateHostAssignTenant200Response**](UpdateHostAssignTenant200Response.md)

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

# **update_host_cloud**
> UpdateHostCloud200Response update_host_cloud()

Change Server Cloud

This api call is reserved for migrating servers from one cloud to another. This could be due to moving clusters or resource pool scoping of a server without losing the data.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_cloud_request import UpdateHostCloudRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_host_cloud200_response import UpdateHostCloud200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    update_host_cloud_request = UpdateHostCloudRequest(
        cloud_id=1,
        servers=[
            UpdateHostCloudRequestServersInner(
                source=1,
                target=1,
            ),
        ],
    ) # UpdateHostCloudRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Change Server Cloud
        api_response = api_instance.update_host_cloud(update_host_cloud_request=update_host_cloud_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_cloud: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_host_cloud_request** | [**UpdateHostCloudRequest**](UpdateHostCloudRequest.md)|  | [optional]

### Return type

[**UpdateHostCloud200Response**](UpdateHostCloud200Response.md)

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

# **update_host_execute_workflow**
> Model200Success update_host_execute_workflow(id)

Run Workflow on a Host

This will run a provisioning workflow on a host.  For operational workflows, see Execute a Workflow. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_execute_workflow_request import UpdateHostExecuteWorkflowRequest
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    workflow_id = 3 # int | ID of the workflow to execute (optional)
    workflow_name = "myworkflow" # str | Name of the workflow to execute (optional)
    update_host_execute_workflow_request = UpdateHostExecuteWorkflowRequest(
        task_set=UpdateHostExecuteWorkflowRequestTaskSet(
            custom_options={},
        ),
        task_phase="provision",
    ) # UpdateHostExecuteWorkflowRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Run Workflow on a Host
        api_response = api_instance.update_host_execute_workflow(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_execute_workflow: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Run Workflow on a Host
        api_response = api_instance.update_host_execute_workflow(id, workflow_id=workflow_id, workflow_name=workflow_name, update_host_execute_workflow_request=update_host_execute_workflow_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_execute_workflow: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **workflow_id** | **int**| ID of the workflow to execute | [optional]
 **workflow_name** | **str**| Name of the workflow to execute | [optional]
 **update_host_execute_workflow_request** | [**UpdateHostExecuteWorkflowRequest**](UpdateHostExecuteWorkflowRequest.md)|  | [optional]

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

# **update_host_install_agent**
> UpdateHostInstallAgent200Response update_host_install_agent(id)

Install Agent

This will make the host a managed server, and install the agent.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_install_agent_request import UpdateHostInstallAgentRequest
from openapi_client.model.update_host_install_agent200_response import UpdateHostInstallAgent200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_host_install_agent_request = UpdateHostInstallAgentRequest(
        server=UpdateHostInstallAgentRequestServer(
            ssh_username="ssh_username_example",
            ssh_password="ssh_password_example",
            server_os=UpdateHostInstallAgentRequestServerServerOs(
                id=1,
            ),
        ),
    ) # UpdateHostInstallAgentRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Install Agent
        api_response = api_instance.update_host_install_agent(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_install_agent: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Install Agent
        api_response = api_instance.update_host_install_agent(id, update_host_install_agent_request=update_host_install_agent_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_install_agent: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_host_install_agent_request** | [**UpdateHostInstallAgentRequest**](UpdateHostInstallAgentRequest.md)|  | [optional]

### Return type

[**UpdateHostInstallAgent200Response**](UpdateHostInstallAgent200Response.md)

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

# **update_host_managed**
> UpdateHostInstallAgent200Response update_host_managed(id)

Convert To Managed

This will make the host a managed server, and install the agent.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_install_agent200_response import UpdateHostInstallAgent200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_host_managed_request import UpdateHostManagedRequest
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_host_managed_request = UpdateHostManagedRequest(
        server=UpdateHostManagedRequestServer(
            ssh_username="ssh_username_example",
            ssh_password="ssh_password_example",
            server_os=UpdateHostInstallAgentRequestServerServerOs(
                id=1,
            ),
            plan=UpdateHostManagedRequestServerPlan(
                id=1,
            ),
            account=UpdateHostManagedRequestServerAccount(
                id=1,
            ),
            provision_site_id=1,
            tags=[
                UpdateHostManagedRequestServerTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            config=UpdateHostManagedRequestServerConfig(
                custom_options=UpdateHostManagedRequestServerConfigCustomOptions(
                    dbfoldername="dbfoldername_example",
                ),
            ),
        ),
        install_agent=True,
        instance_type_id=1,
        layout=1,
    ) # UpdateHostManagedRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Convert To Managed
        api_response = api_instance.update_host_managed(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_managed: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Convert To Managed
        api_response = api_instance.update_host_managed(id, update_host_managed_request=update_host_managed_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_managed: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_host_managed_request** | [**UpdateHostManagedRequest**](UpdateHostManagedRequest.md)|  | [optional]

### Return type

[**UpdateHostInstallAgent200Response**](UpdateHostInstallAgent200Response.md)

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

# **update_host_resize**
> UpdateHostAssignTenant200Response update_host_resize(id)

Resize a Host

Will resize a host asynchronously. This endpoint also allows for NIC reconfiguration by passing a new array of `networkInterfaces`.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_host_resize_request import UpdateHostResizeRequest
from openapi_client.model.update_host_assign_tenant200_response import UpdateHostAssignTenant200Response
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_host_resize_request = UpdateHostResizeRequest(
        server=UpdateHostResizeRequestServer(
            plan=UpdateHostResizeRequestServerPlan(
                id=1,
            ),
        ),
        service_plan_options=UpdateHostResizeRequestServicePlanOptions(
            max_cores=1,
            cores_per_socket=1,
            max_memory=1,
        ),
        volumes=[
            InstanceCreateVolume(
                id=1,
                root_volume=True,
                name="data",
                size=5,
                size_id=1,
                storage_type=1,
                datastore_id=InstanceCreateVolumeDatastoreId(None),
            ),
        ],
        delete_original_volumes=False,
        network_interfaces=[
            InstanceCreateNetwork(
                network=InstanceCreateNetworkNetwork(
                    id="id_example",
                ),
                network_interface_type_id=1,
                ip_address="ip_address_example",
                id=1,
            ),
        ],
    ) # UpdateHostResizeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Resize a Host
        api_response = api_instance.update_host_resize(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_resize: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Resize a Host
        api_response = api_instance.update_host_resize(id, update_host_resize_request=update_host_resize_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_resize: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_host_resize_request** | [**UpdateHostResizeRequest**](UpdateHostResizeRequest.md)|  | [optional]

### Return type

[**UpdateHostAssignTenant200Response**](UpdateHostAssignTenant200Response.md)

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

# **update_host_upgrade_agent**
> Model200Success update_host_upgrade_agent(id)

Upgrade Agent

This will upgrade the version of the agent installed on the host.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Upgrade Agent
        api_response = api_instance.update_host_upgrade_agent(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_host_upgrade_agent: %s\n" % e)
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

# **update_server_network_interface**
> UpdateInstanceNetworkInterface200Response update_server_network_interface(id, network_interface_id)

Updating a label for a Server's Network

Updating a Server's Network's Label

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_instance_network_interface200_response import UpdateInstanceNetworkInterface200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.network_interface_update import NetworkInterfaceUpdate
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    network_interface_id = 7 # float | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
    network_interface_update = NetworkInterfaceUpdate(
        name="eth0-new",
    ) # NetworkInterfaceUpdate |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a label for a Server's Network
        api_response = api_instance.update_server_network_interface(id, network_interface_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_server_network_interface: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a label for a Server's Network
        api_response = api_instance.update_server_network_interface(id, network_interface_id, network_interface_update=network_interface_update)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_server_network_interface: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **network_interface_id** | **float**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced |
 **network_interface_update** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md)|  | [optional]

### Return type

[**UpdateInstanceNetworkInterface200Response**](UpdateInstanceNetworkInterface200Response.md)

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

# **update_wiki_server**
> UpdateWikiApp200Response update_wiki_server(id)

Update a Server Wiki Page

Updates a server Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import hosts_api
from openapi_client.model.update_wiki_app200_response import UpdateWikiApp200Response
from openapi_client.model.update_wiki_app_request import UpdateWikiAppRequest
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
    api_instance = hosts_api.HostsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_wiki_app_request = UpdateWikiAppRequest(
        page=UpdateWikiAppRequestPage(
            name="Sample Doc",
            category="info",
            content="A sample document in **markdown**.`",
        ),
    ) # UpdateWikiAppRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Server Wiki Page
        api_response = api_instance.update_wiki_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_wiki_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Server Wiki Page
        api_response = api_instance.update_wiki_server(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling HostsApi->update_wiki_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_wiki_app_request** | [**UpdateWikiAppRequest**](UpdateWikiAppRequest.md)|  | [optional]

### Return type

[**UpdateWikiApp200Response**](UpdateWikiApp200Response.md)

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

