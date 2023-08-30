# openapi_client.CloudsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_cloud_resource_pool**](CloudsApi.md#add_cloud_resource_pool) | **POST** /api/zones/{zoneId}/resource-pools | Creates a Specified Resource Pool for Specified Cloud
[**add_clouds**](CloudsApi.md#add_clouds) | **POST** /api/zones | Creates a Cloud
[**get_cloud_datastores**](CloudsApi.md#get_cloud_datastores) | **GET** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**get_cloud_folders**](CloudsApi.md#get_cloud_folders) | **GET** /api/zones/{zoneId}/folders/{id} | Retrieves a Resource Folder for Specified Cloud
[**get_cloud_resource_pools**](CloudsApi.md#get_cloud_resource_pools) | **GET** /api/zones/{zoneId}/resource-pools/{id} | Retrieves a Resource Pool for Specified Cloud
[**get_cloud_types**](CloudsApi.md#get_cloud_types) | **GET** /api/zone-types/{id} | Retrieves a Specific Cloud Type
[**get_clouds**](CloudsApi.md#get_clouds) | **GET** /api/zones/{id} | Retrieves a Specific Cloud
[**get_wiki_cloud**](CloudsApi.md#get_wiki_cloud) | **GET** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**list_cloud_datastores**](CloudsApi.md#list_cloud_datastores) | **GET** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**list_cloud_folders**](CloudsApi.md#list_cloud_folders) | **GET** /api/zones/{zoneId}/folders | Retrieves all resource folders for Specified Cloud
[**list_cloud_resource_pools**](CloudsApi.md#list_cloud_resource_pools) | **GET** /api/zones/{zoneId}/resource-pools | Retrieves all Resource Pools for Specified Cloud
[**list_cloud_security_groups**](CloudsApi.md#list_cloud_security_groups) | **GET** /api/zones/{id}/security-groups | Retrieves all Security Groups for a Cloud
[**list_cloud_types**](CloudsApi.md#list_cloud_types) | **GET** /api/zone-types | Retrieves all Cloud Types
[**list_clouds**](CloudsApi.md#list_clouds) | **GET** /api/zones | Retrieves all Clouds
[**refresh_clouds**](CloudsApi.md#refresh_clouds) | **POST** /api/zones/{id}/refresh | Refreshes a Cloud
[**remove_cloud_resource_pools**](CloudsApi.md#remove_cloud_resource_pools) | **DELETE** /api/zones/{zoneId}/resource-pools/{id} | Deletes a Resource Pool for Specified Cloud
[**remove_clouds**](CloudsApi.md#remove_clouds) | **DELETE** /api/zones/{id} | Deletes a Cloud
[**update_cloud_datastores**](CloudsApi.md#update_cloud_datastores) | **PUT** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**update_cloud_folders**](CloudsApi.md#update_cloud_folders) | **PUT** /api/zones/{zoneId}/folders/{id} | Updates a Resource Folder for Specified Cloud
[**update_cloud_logo**](CloudsApi.md#update_cloud_logo) | **POST** /api/zones/{id}/update-logo | Update Logo For Cloud
[**update_cloud_resource_pool**](CloudsApi.md#update_cloud_resource_pool) | **PUT** /api/zones/{zoneId}/resource-pools/{id} | Updates a Specified Resource Pool for Specified Cloud
[**update_cloud_security_groups**](CloudsApi.md#update_cloud_security_groups) | **POST** /api/zones/{id}/security-groups | Sets Security Groups for a Cloud
[**update_clouds**](CloudsApi.md#update_clouds) | **PUT** /api/zones/{id} | Updates a Cloud
[**update_wiki_cloud**](CloudsApi.md#update_wiki_cloud) | **PUT** /api/zones/{id}/wiki | Update a Cloud Wiki Page


# **add_cloud_resource_pool**
> AddCloudResourcePool200Response add_cloud_resource_pool(zone_id)

Creates a Specified Resource Pool for Specified Cloud

Creates a resource pool for specified cloud. Only certain types of clouds support creating and deleting resource pools. Configuration options vary by type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.add_cloud_resource_pool200_response import AddCloudResourcePool200Response
from openapi_client.model.add_cloud_resource_pool_request import AddCloudResourcePoolRequest
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    add_cloud_resource_pool_request = AddCloudResourcePoolRequest(
        resource_pool=AddCloudResourcePoolRequestResourcePool(
            name="name_example",
            default_pool=False,
            default_image=False,
            active=True,
            visibility="private",
            display_name="display_name_example",
            inventory=True,
            config=AddCloudResourcePoolRequestResourcePoolConfig(None),
            tenant_permissions=[
                UpdateCloudFoldersRequestFolderTenantPermissionsInner(
                    accounts=[1,3],
                ),
            ],
            resource_permissions=UpdateCloudDatastoresRequestDatastoreResourcePermissions(
                all=True,
                sites=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
                all_plans=True,
                plans=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
            ),
        ),
    ) # AddCloudResourcePoolRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Creates a Specified Resource Pool for Specified Cloud
        api_response = api_instance.add_cloud_resource_pool(zone_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->add_cloud_resource_pool: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Specified Resource Pool for Specified Cloud
        api_response = api_instance.add_cloud_resource_pool(zone_id, add_cloud_resource_pool_request=add_cloud_resource_pool_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->add_cloud_resource_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **add_cloud_resource_pool_request** | [**AddCloudResourcePoolRequest**](AddCloudResourcePoolRequest.md)|  | [optional]

### Return type

[**AddCloudResourcePool200Response**](AddCloudResourcePool200Response.md)

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

# **add_clouds**
> AddClouds200Response add_clouds()

Creates a Cloud

Creates a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.add_clouds200_response import AddClouds200Response
from openapi_client.model.add_clouds_request import AddCloudsRequest
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
    api_instance = clouds_api.CloudsApi(api_client)
    add_clouds_request = AddCloudsRequest(
        zone=ZoneCreate(
            name="My Cloud",
            description="description_example",
            code="mycloud",
            location="US East",
            visibility="private",
            zone_type=ZoneCreateZoneType(None),
            group_id=3,
            account_id=1,
            enabled=True,
            auto_recover_power_state=False,
            scale_priority=1,
            linked_account_id=1,
            config={},
            security_mode="off",
            credential=ZoneCreateCredential(
                type="local",
                id=1,
            ),
        ),
    ) # AddCloudsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Cloud
        api_response = api_instance.add_clouds(add_clouds_request=add_clouds_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->add_clouds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_clouds_request** | [**AddCloudsRequest**](AddCloudsRequest.md)|  | [optional]

### Return type

[**AddClouds200Response**](AddClouds200Response.md)

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

# **get_cloud_datastores**
> GetCloudDatastores200Response get_cloud_datastores(zone_id, id)

Retrieves a Datastore for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific datastore under a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.get_cloud_datastores200_response import GetCloudDatastores200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Datastore for Specified Cloud
        api_response = api_instance.get_cloud_datastores(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->get_cloud_datastores: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCloudDatastores200Response**](GetCloudDatastores200Response.md)

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

# **get_cloud_folders**
> GetCloudFolders200Response get_cloud_folders(zone_id, id)

Retrieves a Resource Folder for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific folder under a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cloud_folders200_response import GetCloudFolders200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Resource Folder for Specified Cloud
        api_response = api_instance.get_cloud_folders(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->get_cloud_folders: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCloudFolders200Response**](GetCloudFolders200Response.md)

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

# **get_cloud_resource_pools**
> GetCloudResourcePools200Response get_cloud_resource_pools(zone_id, id)

Retrieves a Resource Pool for Specified Cloud

This endpoint retrieves a specific resource pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.get_cloud_resource_pools200_response import GetCloudResourcePools200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Resource Pool for Specified Cloud
        api_response = api_instance.get_cloud_resource_pools(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->get_cloud_resource_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCloudResourcePools200Response**](GetCloudResourcePools200Response.md)

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

# **get_cloud_types**
> GetCloudTypes200Response get_cloud_types(id)

Retrieves a Specific Cloud Type

Retrieves a specific cloud type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cloud_types200_response import GetCloudTypes200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Cloud Type
        api_response = api_instance.get_cloud_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->get_cloud_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetCloudTypes200Response**](GetCloudTypes200Response.md)

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

# **get_clouds**
> AddClouds200ResponseAllOf get_clouds(id)

Retrieves a Specific Cloud

Retrieves a specific cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.add_clouds200_response_all_of import AddClouds200ResponseAllOf
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Cloud
        api_response = api_instance.get_clouds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->get_clouds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddClouds200ResponseAllOf**](AddClouds200ResponseAllOf.md)

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

# **get_wiki_cloud**
> GetWikiApp200Response get_wiki_cloud(id)

Retrieves a Cloud Wiki Page

This endpoint retrieves a cloud Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Cloud Wiki Page
        api_response = api_instance.get_wiki_cloud(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->get_wiki_cloud: %s\n" % e)
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

# **list_cloud_datastores**
> ListCloudDatastores200Response list_cloud_datastores(zone_id)

Retrieves all Datastores for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all data stores under a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.list_cloud_datastores200_response import ListCloudDatastores200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"

    # example passing only required values which don't have defaults set
    try:
        # Retrieves all Datastores for Specified Cloud
        api_response = api_instance.list_cloud_datastores(zone_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_datastores: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Datastores for Specified Cloud
        api_response = api_instance.list_cloud_datastores(zone_id, name=name, phrase=phrase, max=max, sort=sort, direction=direction)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_datastores: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"

### Return type

[**ListCloudDatastores200Response**](ListCloudDatastores200Response.md)

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

# **list_cloud_folders**
> ListCloudFolders200Response list_cloud_folders(zone_id)

Retrieves all resource folders for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource folders under a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_cloud_folders200_response import ListCloudFolders200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25

    # example passing only required values which don't have defaults set
    try:
        # Retrieves all resource folders for Specified Cloud
        api_response = api_instance.list_cloud_folders(zone_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_folders: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all resource folders for Specified Cloud
        api_response = api_instance.list_cloud_folders(zone_id, name=name, phrase=phrase, max=max)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_folders: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25

### Return type

[**ListCloudFolders200Response**](ListCloudFolders200Response.md)

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

# **list_cloud_resource_pools**
> ListCloudResourcePools200Response list_cloud_resource_pools(zone_id)

Retrieves all Resource Pools for Specified Cloud

Resource Pools can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource pools under a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.list_cloud_resource_pools200_response import ListCloudResourcePools200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25

    # example passing only required values which don't have defaults set
    try:
        # Retrieves all Resource Pools for Specified Cloud
        api_response = api_instance.list_cloud_resource_pools(zone_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_resource_pools: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Resource Pools for Specified Cloud
        api_response = api_instance.list_cloud_resource_pools(zone_id, name=name, phrase=phrase, max=max)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_resource_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25

### Return type

[**ListCloudResourcePools200Response**](ListCloudResourcePools200Response.md)

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

# **list_cloud_security_groups**
> ListCloudSecurityGroups200Response list_cloud_security_groups(id)

Retrieves all Security Groups for a Cloud

Retrieves all security groups for a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.list_cloud_security_groups200_response import ListCloudSecurityGroups200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves all Security Groups for a Cloud
        api_response = api_instance.list_cloud_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**ListCloudSecurityGroups200Response**](ListCloudSecurityGroups200Response.md)

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

# **list_cloud_types**
> ListCloudTypes200Response list_cloud_types()

Retrieves all Cloud Types

Fetch a paginated list of available cloud types. This returns the configuration options for each type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.list_cloud_types200_response import ListCloudTypes200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    provision_type = "AKS" # str | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Cloud Types
        api_response = api_instance.list_cloud_types(max=max, offset=offset, sort=sort, direction=direction, name=name, code=code, phrase=phrase, provision_type=provision_type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_cloud_types: %s\n" % e)
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

### Return type

[**ListCloudTypes200Response**](ListCloudTypes200Response.md)

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

# **list_clouds**
> ListClouds200Response list_clouds()

Retrieves all Clouds

Retrieves all clouds. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_clouds200_response import ListClouds200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    type = "alibaba" # str | If specified will return all zones by cloud type code. Refer to `Zone Types` API for up to date listings.  (optional)
    group_id = 13 # int | If specified will return all zones assigned to a server group by id. Accepts multiple values. (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Clouds
        api_response = api_instance.list_clouds(last_updated=last_updated, type=type, group_id=group_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->list_clouds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **type** | **str**| If specified will return all zones by cloud type code. Refer to &#x60;Zone Types&#x60; API for up to date listings.  | [optional]
 **group_id** | **int**| If specified will return all zones assigned to a server group by id. Accepts multiple values. | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListClouds200Response**](ListClouds200Response.md)

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

# **refresh_clouds**
> Model200Success refresh_clouds(id)

Refreshes a Cloud

Refreshes a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.refresh_clouds_request import RefreshCloudsRequest
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    refresh_clouds_request = RefreshCloudsRequest(
        mode="mode_example",
        rebuild="rebuild_example",
        period="period_example",
    ) # RefreshCloudsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Refreshes a Cloud
        api_response = api_instance.refresh_clouds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->refresh_clouds: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Refreshes a Cloud
        api_response = api_instance.refresh_clouds(id, refresh_clouds_request=refresh_clouds_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->refresh_clouds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **refresh_clouds_request** | [**RefreshCloudsRequest**](RefreshCloudsRequest.md)|  | [optional]

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

# **remove_cloud_resource_pools**
> Model200Success remove_cloud_resource_pools(zone_id, id)

Deletes a Resource Pool for Specified Cloud

Deletes a resource pool for specified Cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Resource Pool for Specified Cloud
        api_response = api_instance.remove_cloud_resource_pools(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->remove_cloud_resource_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
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

# **remove_clouds**
> Model200Success remove_clouds(id)

Deletes a Cloud

Deletes a specified Cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    remove_resources = False # bool | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Cloud
        api_response = api_instance.remove_clouds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->remove_clouds: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Deletes a Cloud
        api_response = api_instance.remove_clouds(id, remove_resources=remove_resources)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->remove_clouds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **remove_resources** | **bool**| Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [optional] if omitted the server will use the default value of False

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

# **update_cloud_datastores**
> UpdateCloudDatastores200Response update_cloud_datastores(zone_id, id)

Updates a Specified Datastore for Specified Cloud

Updates a datastore for specified cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.update_cloud_datastores200_response import UpdateCloudDatastores200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_cloud_datastores_request import UpdateCloudDatastoresRequest
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cloud_datastores_request = UpdateCloudDatastoresRequest(
        datastore=UpdateCloudDatastoresRequestDatastore(
            active=True,
            visibility="private",
            tenant_permissions=[
                UpdateCloudDatastoresRequestDatastoreTenantPermissionsInner(
                    accounts=[1,3],
                    default_target=[1,3],
                    default_store=[1,3],
                ),
            ],
            resource_permissions=UpdateCloudDatastoresRequestDatastoreResourcePermissions(
                all=True,
                sites=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
                all_plans=True,
                plans=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
            ),
        ),
    ) # UpdateCloudDatastoresRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Specified Datastore for Specified Cloud
        api_response = api_instance.update_cloud_datastores(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_datastores: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Specified Datastore for Specified Cloud
        api_response = api_instance.update_cloud_datastores(zone_id, id, update_cloud_datastores_request=update_cloud_datastores_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_datastores: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cloud_datastores_request** | [**UpdateCloudDatastoresRequest**](UpdateCloudDatastoresRequest.md)|  | [optional]

### Return type

[**UpdateCloudDatastores200Response**](UpdateCloudDatastores200Response.md)

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

# **update_cloud_folders**
> UpdateCloudFolders200Response update_cloud_folders(zone_id, id)

Updates a Resource Folder for Specified Cloud

Updates a resource folder for specified cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.update_cloud_folders200_response import UpdateCloudFolders200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_cloud_folders_request import UpdateCloudFoldersRequest
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cloud_folders_request = UpdateCloudFoldersRequest(
        folder=UpdateCloudFoldersRequestFolder(
            default_folder=False,
            default_image=False,
            active=True,
            visibility="private",
            tenant_permissions=[
                UpdateCloudFoldersRequestFolderTenantPermissionsInner(
                    accounts=[1,3],
                ),
            ],
            resource_permissions=UpdateCloudDatastoresRequestDatastoreResourcePermissions(
                all=True,
                sites=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
                all_plans=True,
                plans=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
            ),
        ),
    ) # UpdateCloudFoldersRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Resource Folder for Specified Cloud
        api_response = api_instance.update_cloud_folders(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_folders: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Resource Folder for Specified Cloud
        api_response = api_instance.update_cloud_folders(zone_id, id, update_cloud_folders_request=update_cloud_folders_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_folders: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cloud_folders_request** | [**UpdateCloudFoldersRequest**](UpdateCloudFoldersRequest.md)|  | [optional]

### Return type

[**UpdateCloudFolders200Response**](UpdateCloudFolders200Response.md)

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

# **update_cloud_logo**
> Model200Success update_cloud_logo(id)

Update Logo For Cloud

Use this command to update the logo and dark logo images for a cloud. This endpoint expects multipart form data as the request format, not JSON. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    logo = open('/path/to/file', 'rb') # file_type | Logo File png,jpg,svg (optional)
    dark_logo = open('/path/to/file', 'rb') # file_type | Dark Logo File png,jpg,svg (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Logo For Cloud
        api_response = api_instance.update_cloud_logo(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_logo: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Logo For Cloud
        api_response = api_instance.update_cloud_logo(id, logo=logo, dark_logo=dark_logo)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_logo: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **logo** | **file_type**| Logo File png,jpg,svg | [optional]
 **dark_logo** | **file_type**| Dark Logo File png,jpg,svg | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_cloud_resource_pool**
> AddCloudResourcePool200Response update_cloud_resource_pool(zone_id, id)

Updates a Specified Resource Pool for Specified Cloud

Updates a resource pool for specified cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.update_cloud_resource_pool_request import UpdateCloudResourcePoolRequest
from openapi_client.model.add_cloud_resource_pool200_response import AddCloudResourcePool200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    zone_id = 7 # float | The ID of the cloud
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cloud_resource_pool_request = UpdateCloudResourcePoolRequest(
        resource_pool=UpdateCloudResourcePoolRequestResourcePool(
            active=True,
            visibility="private",
            display_name="display_name_example",
            inventory=True,
            tenant_permissions=[
                UpdateCloudFoldersRequestFolderTenantPermissionsInner(
                    accounts=[1,3],
                ),
            ],
            resource_permissions=UpdateCloudDatastoresRequestDatastoreResourcePermissions(
                all=True,
                sites=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
                all_plans=True,
                plans=[
                    UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner(
                        id=1,
                    ),
                ],
            ),
        ),
    ) # UpdateCloudResourcePoolRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Specified Resource Pool for Specified Cloud
        api_response = api_instance.update_cloud_resource_pool(zone_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_resource_pool: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Specified Resource Pool for Specified Cloud
        api_response = api_instance.update_cloud_resource_pool(zone_id, id, update_cloud_resource_pool_request=update_cloud_resource_pool_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_resource_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone_id** | **float**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cloud_resource_pool_request** | [**UpdateCloudResourcePoolRequest**](UpdateCloudResourcePoolRequest.md)|  | [optional]

### Return type

[**AddCloudResourcePool200Response**](AddCloudResourcePool200Response.md)

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

# **update_cloud_security_groups**
> UpdateCloudSecurityGroups200Response update_cloud_security_groups(id)

Sets Security Groups for a Cloud

Sets security groups for acloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.update_cloud_security_groups_request import UpdateCloudSecurityGroupsRequest
from openapi_client.model.update_cloud_security_groups200_response import UpdateCloudSecurityGroups200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_cloud_security_groups_request = UpdateCloudSecurityGroupsRequest(
        security_group_ids=[19,2],
    ) # UpdateCloudSecurityGroupsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Sets Security Groups for a Cloud
        api_response = api_instance.update_cloud_security_groups(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_security_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Sets Security Groups for a Cloud
        api_response = api_instance.update_cloud_security_groups(id, update_cloud_security_groups_request=update_cloud_security_groups_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_cloud_security_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_cloud_security_groups_request** | [**UpdateCloudSecurityGroupsRequest**](UpdateCloudSecurityGroupsRequest.md)|  | [optional]

### Return type

[**UpdateCloudSecurityGroups200Response**](UpdateCloudSecurityGroups200Response.md)

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

# **update_clouds**
> AddClouds200Response update_clouds(id)

Updates a Cloud

Updates a cloud. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
from openapi_client.model.update_clouds_request import UpdateCloudsRequest
from openapi_client.model.add_clouds200_response import AddClouds200Response
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
    api_instance = clouds_api.CloudsApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_clouds_request = UpdateCloudsRequest(
        zone=UpdateCloudsRequestZone(
            name="My Cloud",
            description="description_example",
            code="mycloud",
            location="US East",
            visibility="private",
            zone_type="standard",
            group_id=3,
            account_id=1,
            enabled=True,
            auto_recover_power_state=False,
            scale_priority=1,
            linked_account_id=1,
            config={},
            security_mode="off",
            default_cloud_logos=True,
            credential={},
        ),
    ) # UpdateCloudsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Cloud
        api_response = api_instance.update_clouds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_clouds: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Cloud
        api_response = api_instance.update_clouds(id, update_clouds_request=update_clouds_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_clouds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_clouds_request** | [**UpdateCloudsRequest**](UpdateCloudsRequest.md)|  | [optional]

### Return type

[**AddClouds200Response**](AddClouds200Response.md)

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

# **update_wiki_cloud**
> UpdateWikiApp200Response update_wiki_cloud(id)

Update a Cloud Wiki Page

Updates a cloud Wiki page. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import clouds_api
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
    api_instance = clouds_api.CloudsApi(api_client)
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
        # Update a Cloud Wiki Page
        api_response = api_instance.update_wiki_cloud(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_wiki_cloud: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Cloud Wiki Page
        api_response = api_instance.update_wiki_cloud(id, update_wiki_app_request=update_wiki_app_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CloudsApi->update_wiki_cloud: %s\n" % e)
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

