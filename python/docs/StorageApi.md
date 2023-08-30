# openapi_client.StorageApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_storage_buckets**](StorageApi.md#add_storage_buckets) | **POST** /api/storage-buckets | Creates a Storage Bucket
[**add_storage_servers**](StorageApi.md#add_storage_servers) | **POST** /api/storage-servers | Creates a Storage Server
[**add_storage_volumes**](StorageApi.md#add_storage_volumes) | **POST** /api/storage-volumes | Creates a Storage Volume
[**get_storage_buckets**](StorageApi.md#get_storage_buckets) | **GET** /api/storage-buckets/{id} | Retrieves a Specific Storage Bucket
[**get_storage_server_types**](StorageApi.md#get_storage_server_types) | **GET** /api/storage-server-types/{id} | Retrieves a Specific Storage Server Type
[**get_storage_servers**](StorageApi.md#get_storage_servers) | **GET** /api/storage-servers/{id} | Retrieves a Specific Storage Server
[**get_storage_volume_types**](StorageApi.md#get_storage_volume_types) | **GET** /api/storage-volume-types/{id} | Retrieves a Specific Storage Volume Type
[**get_storage_volumes**](StorageApi.md#get_storage_volumes) | **GET** /api/storage-volumes/{id} | Retrieves a Specific Storage Volume
[**list_storage_buckets**](StorageApi.md#list_storage_buckets) | **GET** /api/storage-buckets | Retrieves all Storage Buckets
[**list_storage_server_types**](StorageApi.md#list_storage_server_types) | **GET** /api/storage-server-types | Retrieves all Storage Server Types
[**list_storage_servers**](StorageApi.md#list_storage_servers) | **GET** /api/storage-servers | Retrieves all Storage Servers
[**list_storage_volume_types**](StorageApi.md#list_storage_volume_types) | **GET** /api/storage-volume-types | Retrieves all Storage Volume Types
[**list_storage_volumes**](StorageApi.md#list_storage_volumes) | **GET** /api/storage-volumes | Retrieves all Storage Volumes
[**remove_storage_buckets**](StorageApi.md#remove_storage_buckets) | **DELETE** /api/storage-buckets/{id} | Deletes a Storage Bucket
[**remove_storage_servers**](StorageApi.md#remove_storage_servers) | **DELETE** /api/storage-servers/{id} | Deletes a Storage Server
[**remove_storage_volumes**](StorageApi.md#remove_storage_volumes) | **DELETE** /api/storage-volumes/{id} | Deletes a Storage Volume
[**update_storage_buckets**](StorageApi.md#update_storage_buckets) | **PUT** /api/storage-buckets/{id} | Updates a Storage Bucket
[**update_storage_servers**](StorageApi.md#update_storage_servers) | **PUT** /api/storage-servers/{id} | Updates a Storage Server
[**update_storage_volumes**](StorageApi.md#update_storage_volumes) | **PUT** /api/storage-volumes/{id} | Updates a Storage Volume


# **add_storage_buckets**
> AddStorageBuckets200Response add_storage_buckets()

Creates a Storage Bucket

Creates a storage bucket. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_buckets_request import AddStorageBucketsRequest
from openapi_client.model.add_storage_buckets200_response import AddStorageBuckets200Response
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
    api_instance = storage_api.StorageApi(api_client)
    add_storage_buckets_request = AddStorageBucketsRequest(
        storage_bucket=AddStorageBucketsRequestStorageBucket(
            name="name_example",
            provider_type="s3",
            default_backup_target=False,
            copy_to_store=True,
            default_deployment_target=False,
            default_virtual_image_target=False,
            retention_policy_type="none",
            retention_policy_days=1,
            retention_provider="retention_provider_example",
            bucket_name="myBucket",
            create_bucket=False,
            config=AddStorageBucketsRequestStorageBucketConfig(None),
        ),
    ) # AddStorageBucketsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Storage Bucket
        api_response = api_instance.add_storage_buckets(add_storage_buckets_request=add_storage_buckets_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->add_storage_buckets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_storage_buckets_request** | [**AddStorageBucketsRequest**](AddStorageBucketsRequest.md)|  | [optional]

### Return type

[**AddStorageBuckets200Response**](AddStorageBuckets200Response.md)

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

# **add_storage_servers**
> AddStorageServers200Response add_storage_servers()

Creates a Storage Server

Creates a storage server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_servers200_response import AddStorageServers200Response
from openapi_client.model.add_storage_servers_request import AddStorageServersRequest
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
    api_instance = storage_api.StorageApi(api_client)
    add_storage_servers_request = AddStorageServersRequest(
        storage_server=AddStorageServersRequestStorageServer(
            name="name_example",
            type="type_example",
            description="description_example",
            enabled=True,
            config={},
            visibility="private",
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
        ),
    ) # AddStorageServersRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Storage Server
        api_response = api_instance.add_storage_servers(add_storage_servers_request=add_storage_servers_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->add_storage_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_storage_servers_request** | [**AddStorageServersRequest**](AddStorageServersRequest.md)|  | [optional]

### Return type

[**AddStorageServers200Response**](AddStorageServers200Response.md)

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

# **add_storage_volumes**
> AddStorageVolumes200Response add_storage_volumes()

Creates a Storage Volume

This endpoint allows creating a Storage Volume. Only certain types of storage servers support creating and deleting storage volumes, such as 3Par and Isilon. Configuration options vary by `Storage Volume Type`. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_volumes200_response import AddStorageVolumes200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_storage_volumes_request import AddStorageVolumesRequest
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
    api_instance = storage_api.StorageApi(api_client)
    add_storage_volumes_request = AddStorageVolumesRequest(
        storage_volume=AddStorageVolumesRequestStorageVolume(
            name="name_example",
            type="type_example",
            config={},
            storage_server=AddStorageVolumesRequestStorageVolumeStorageServer(
                id=1,
            ),
            storage_group=AddStorageVolumesRequestStorageVolumeStorageServer(
                id=1,
            ),
        ),
    ) # AddStorageVolumesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Storage Volume
        api_response = api_instance.add_storage_volumes(add_storage_volumes_request=add_storage_volumes_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->add_storage_volumes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_storage_volumes_request** | [**AddStorageVolumesRequest**](AddStorageVolumesRequest.md)|  | [optional]

### Return type

[**AddStorageVolumes200Response**](AddStorageVolumes200Response.md)

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

# **get_storage_buckets**
> AddStorageBuckets200ResponseAllOf get_storage_buckets(id)

Retrieves a Specific Storage Bucket

Retrieves a specific storage bucket. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_buckets200_response_all_of import AddStorageBuckets200ResponseAllOf
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Storage Bucket
        api_response = api_instance.get_storage_buckets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->get_storage_buckets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddStorageBuckets200ResponseAllOf**](AddStorageBuckets200ResponseAllOf.md)

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

# **get_storage_server_types**
> GetStorageServerTypes200Response get_storage_server_types(id)

Retrieves a Specific Storage Server Type

Retrieves a specific storage server type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_storage_server_types200_response import GetStorageServerTypes200Response
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Storage Server Type
        api_response = api_instance.get_storage_server_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->get_storage_server_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetStorageServerTypes200Response**](GetStorageServerTypes200Response.md)

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

# **get_storage_servers**
> AddStorageServers200ResponseAllOf get_storage_servers(id)

Retrieves a Specific Storage Server

Retrieves a specific storage server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_servers200_response_all_of import AddStorageServers200ResponseAllOf
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Storage Server
        api_response = api_instance.get_storage_servers(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->get_storage_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddStorageServers200ResponseAllOf**](AddStorageServers200ResponseAllOf.md)

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

# **get_storage_volume_types**
> GetStorageVolumeTypes200Response get_storage_volume_types(id)

Retrieves a Specific Storage Volume Type

Retrieves a specific storage volume type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.get_storage_volume_types200_response import GetStorageVolumeTypes200Response
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Storage Volume Type
        api_response = api_instance.get_storage_volume_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->get_storage_volume_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetStorageVolumeTypes200Response**](GetStorageVolumeTypes200Response.md)

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

# **get_storage_volumes**
> AddStorageVolumes200ResponseAllOf get_storage_volumes(id)

Retrieves a Specific Storage Volume

Retrieves a specific storage volume. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_volumes200_response_all_of import AddStorageVolumes200ResponseAllOf
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Storage Volume
        api_response = api_instance.get_storage_volumes(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->get_storage_volumes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddStorageVolumes200ResponseAllOf**](AddStorageVolumes200ResponseAllOf.md)

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

# **list_storage_buckets**
> ListStorageBuckets200Response list_storage_buckets()

Retrieves all Storage Buckets

Retrieves all storage buckets. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_storage_buckets200_response import ListStorageBuckets200Response
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
    api_instance = storage_api.StorageApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Storage Buckets
        api_response = api_instance.list_storage_buckets(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->list_storage_buckets: %s\n" % e)
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

[**ListStorageBuckets200Response**](ListStorageBuckets200Response.md)

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

# **list_storage_server_types**
> ListStorageServerTypes200Response list_storage_server_types()

Retrieves all Storage Server Types

Fetch a paginated list of available storage server types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.list_storage_server_types200_response import ListStorageServerTypes200Response
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
    api_instance = storage_api.StorageApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Storage Server Types
        api_response = api_instance.list_storage_server_types(max=max, offset=offset, sort=sort, direction=direction, name=name, code=code, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->list_storage_server_types: %s\n" % e)
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

### Return type

[**ListStorageServerTypes200Response**](ListStorageServerTypes200Response.md)

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

# **list_storage_servers**
> ListStorageServers200Response list_storage_servers()

Retrieves all Storage Servers

Retrieves all storage servers. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.list_storage_servers200_response import ListStorageServers200Response
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
    api_instance = storage_api.StorageApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Storage Servers
        api_response = api_instance.list_storage_servers(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->list_storage_servers: %s\n" % e)
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

[**ListStorageServers200Response**](ListStorageServers200Response.md)

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

# **list_storage_volume_types**
> ListStorageVolumeTypes200Response list_storage_volume_types()

Retrieves all Storage Volume Types

Fetch a paginated list of available storage volume types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_storage_volume_types200_response import ListStorageVolumeTypes200Response
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
    api_instance = storage_api.StorageApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Storage Volume Types
        api_response = api_instance.list_storage_volume_types(max=max, offset=offset, sort=sort, direction=direction, name=name, code=code, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->list_storage_volume_types: %s\n" % e)
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

### Return type

[**ListStorageVolumeTypes200Response**](ListStorageVolumeTypes200Response.md)

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

# **list_storage_volumes**
> ListStorageVolumes200Response list_storage_volumes()

Retrieves all Storage Volumes

Retrieves all storage volumes. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.list_storage_volumes200_response import ListStorageVolumes200Response
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
    api_instance = storage_api.StorageApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Storage Volumes
        api_response = api_instance.list_storage_volumes(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->list_storage_volumes: %s\n" % e)
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

[**ListStorageVolumes200Response**](ListStorageVolumes200Response.md)

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

# **remove_storage_buckets**
> Model200Success remove_storage_buckets(id)

Deletes a Storage Bucket

Deletes a specified storage bucket. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Storage Bucket
        api_response = api_instance.remove_storage_buckets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->remove_storage_buckets: %s\n" % e)
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

# **remove_storage_servers**
> Model200Success remove_storage_servers(id)

Deletes a Storage Server

Deletes a specified storage server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Storage Server
        api_response = api_instance.remove_storage_servers(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->remove_storage_servers: %s\n" % e)
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

# **remove_storage_volumes**
> Model200Success remove_storage_volumes(id)

Deletes a Storage Volume

Deletes a specified storage volume. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Storage Volume
        api_response = api_instance.remove_storage_volumes(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->remove_storage_volumes: %s\n" % e)
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

# **update_storage_buckets**
> AddStorageBuckets200Response update_storage_buckets(id)

Updates a Storage Bucket

Updates a storage bucket. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_buckets200_response import AddStorageBuckets200Response
from openapi_client.model.update_storage_buckets_request import UpdateStorageBucketsRequest
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_storage_buckets_request = UpdateStorageBucketsRequest(
        storage_bucket=UpdateStorageBucketsRequestStorageBucket(
            name="name_example",
            provider_type="s3",
            default_backup_target=False,
            copy_to_store=True,
            default_deployment_target=False,
            default_virtual_image_target=False,
            retention_policy_type="none",
            retention_policy_days=1,
            retention_provider="retention_provider_example",
            bucket_name="myBucket",
            create_bucket=False,
            config=AddStorageBucketsRequestStorageBucketConfig(None),
        ),
    ) # UpdateStorageBucketsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Storage Bucket
        api_response = api_instance.update_storage_buckets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->update_storage_buckets: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Storage Bucket
        api_response = api_instance.update_storage_buckets(id, update_storage_buckets_request=update_storage_buckets_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->update_storage_buckets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_storage_buckets_request** | [**UpdateStorageBucketsRequest**](UpdateStorageBucketsRequest.md)|  | [optional]

### Return type

[**AddStorageBuckets200Response**](AddStorageBuckets200Response.md)

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

# **update_storage_servers**
> AddStorageServers200Response update_storage_servers(id)

Updates a Storage Server

Updates a storage server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_servers200_response import AddStorageServers200Response
from openapi_client.model.update_storage_servers_request import UpdateStorageServersRequest
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_storage_servers_request = UpdateStorageServersRequest(
        storage_server=UpdateStorageServersRequestStorageServer(
            name="name_example",
            type="type_example",
            description="description_example",
            enabled=True,
            config={},
            visibility="private",
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
        ),
    ) # UpdateStorageServersRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Storage Server
        api_response = api_instance.update_storage_servers(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->update_storage_servers: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Storage Server
        api_response = api_instance.update_storage_servers(id, update_storage_servers_request=update_storage_servers_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->update_storage_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_storage_servers_request** | [**UpdateStorageServersRequest**](UpdateStorageServersRequest.md)|  | [optional]

### Return type

[**AddStorageServers200Response**](AddStorageServers200Response.md)

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

# **update_storage_volumes**
> AddStorageVolumes200Response update_storage_volumes(id)

Updates a Storage Volume

Updates a storage volume. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import storage_api
from openapi_client.model.add_storage_volumes200_response import AddStorageVolumes200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_storage_volumes_request import UpdateStorageVolumesRequest
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
    api_instance = storage_api.StorageApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_storage_volumes_request = UpdateStorageVolumesRequest(
        storage_volume=UpdateStorageVolumesRequestStorageVolume(
            name="name_example",
            type="type_example",
            config={},
            storage_server=AddStorageVolumesRequestStorageVolumeStorageServer(
                id=1,
            ),
            storage_group=AddStorageVolumesRequestStorageVolumeStorageServer(
                id=1,
            ),
        ),
    ) # UpdateStorageVolumesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Storage Volume
        api_response = api_instance.update_storage_volumes(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->update_storage_volumes: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Storage Volume
        api_response = api_instance.update_storage_volumes(id, update_storage_volumes_request=update_storage_volumes_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling StorageApi->update_storage_volumes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_storage_volumes_request** | [**UpdateStorageVolumesRequest**](UpdateStorageVolumesRequest.md)|  | [optional]

### Return type

[**AddStorageVolumes200Response**](AddStorageVolumes200Response.md)

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

