# openapi_client.ArchivesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_archive_bucket**](ArchivesApi.md#add_archive_bucket) | **POST** /api/archives/buckets | Create an Archive Bucket
[**add_archive_file**](ArchivesApi.md#add_archive_file) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
[**add_archive_file_link**](ArchivesApi.md#add_archive_file_link) | **POST** /api/archives/files/{id}/links | Create an Archive File Link
[**delete_archive_bucket**](ArchivesApi.md#delete_archive_bucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket
[**delete_archive_file**](ArchivesApi.md#delete_archive_file) | **DELETE** /api/archives/files/{id} | Delete Archive File
[**delete_archive_file_link**](ArchivesApi.md#delete_archive_file_link) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
[**get_archive_bucket**](ArchivesApi.md#get_archive_bucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket
[**get_archive_file**](ArchivesApi.md#get_archive_file) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File
[**get_archive_file_detail**](ArchivesApi.md#get_archive_file_detail) | **GET** /api/archives/files/{id} | Get Archive File Details
[**get_archive_file_links**](ArchivesApi.md#get_archive_file_links) | **GET** /api/archives/files/{id}/links | Get Archive File Links
[**get_archive_public_file**](ArchivesApi.md#get_archive_public_file) | **GET** /api/public-archives/download/{bucket}/{filepath} | Download a Public Archive File
[**get_archive_public_file_link**](ArchivesApi.md#get_archive_public_file_link) | **GET** /api/public-archives/link | Download an Archive File Link
[**list_archive_buckets**](ArchivesApi.md#list_archive_buckets) | **GET** /api/archives/buckets | Get All Archive Buckets
[**list_archive_files**](ArchivesApi.md#list_archive_files) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
[**update_archive_bucket**](ArchivesApi.md#update_archive_bucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket


# **add_archive_bucket**
> AddArchiveBucket200Response add_archive_bucket()

Create an Archive Bucket

Create an Archive Bucket

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_archive_bucket_request import AddArchiveBucketRequest
from openapi_client.model.add_archive_bucket200_response import AddArchiveBucket200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    add_archive_bucket_request = AddArchiveBucketRequest(
        archive_bucket=ArchiveBucketCreate(
            name="name_example",
            description="description_example",
            storage_provider=ArchiveBucketCreateStorageProvider(
                id=1,
            ),
            visibility="private",
            is_public=False,
            accounts=UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                id=1,
            ),
        ),
    ) # AddArchiveBucketRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Archive Bucket
        api_response = api_instance.add_archive_bucket(add_archive_bucket_request=add_archive_bucket_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->add_archive_bucket: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_archive_bucket_request** | [**AddArchiveBucketRequest**](AddArchiveBucketRequest.md)|  | [optional]

### Return type

[**AddArchiveBucket200Response**](AddArchiveBucket200Response.md)

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

# **add_archive_file**
> AddArchiveFile200Response add_archive_file(bucket, )

Upload Archive File

Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.add_archive_file200_response import AddArchiveFile200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    bucket = "bucket_example" # str | Bucket name
    filename = "/path/to/file" # str | Specify a filename for archive file. The base filename of the uploaded file is the default. (optional)
    file = open('/path/to/file', 'rb') # file_type |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Upload Archive File
        api_response = api_instance.add_archive_file(bucket, )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->add_archive_file: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Upload Archive File
        api_response = api_instance.add_archive_file(bucket, filename=filename, file=file)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->add_archive_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **str**| Bucket name |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"
 **filename** | **str**| Specify a filename for archive file. The base filename of the uploaded file is the default. | [optional]
 **file** | **file_type**|  | [optional]

### Return type

[**AddArchiveFile200Response**](AddArchiveFile200Response.md)

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

# **add_archive_file_link**
> AddArchiveFileLink200Response add_archive_file_link(id)

Create an Archive File Link

This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.add_archive_file_link200_response import AddArchiveFileLink200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    expire_seconds = 600 # int | Time to live in seconds. 0 means do not expire. (optional) if omitted the server will use the default value of 0

    # example passing only required values which don't have defaults set
    try:
        # Create an Archive File Link
        api_response = api_instance.add_archive_file_link(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->add_archive_file_link: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Archive File Link
        api_response = api_instance.add_archive_file_link(id, expire_seconds=expire_seconds)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->add_archive_file_link: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **expire_seconds** | **int**| Time to live in seconds. 0 means do not expire. | [optional] if omitted the server will use the default value of 0

### Return type

[**AddArchiveFileLink200Response**](AddArchiveFileLink200Response.md)

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

# **delete_archive_bucket**
> Model200Success delete_archive_bucket(id)

Delete an Archive Bucket

Will delete an archive bucket from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete an Archive Bucket
        api_response = api_instance.delete_archive_bucket(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->delete_archive_bucket: %s\n" % e)
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

# **delete_archive_file**
> Model200Success delete_archive_file(id)

Delete Archive File

Permanently delete a file or directory.  Deleting a directory will also delete all the files under it. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete Archive File
        api_response = api_instance.delete_archive_file(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->delete_archive_file: %s\n" % e)
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

# **delete_archive_file_link**
> Model200Success delete_archive_file_link(id, link_id)

Delete an Archive File Link

This will delete the link from the system, so it can no longer be used.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    link_id = 1 # int | The ID of the archive file link.

    # example passing only required values which don't have defaults set
    try:
        # Delete an Archive File Link
        api_response = api_instance.delete_archive_file_link(id, link_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->delete_archive_file_link: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **link_id** | **int**| The ID of the archive file link. |

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

# **get_archive_bucket**
> GetArchiveBucket200Response get_archive_bucket(id)

Get a Specific Archive Bucket

This endpoint retrieves a specific archive bucket.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.get_archive_bucket200_response import GetArchiveBucket200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Archive Bucket
        api_response = api_instance.get_archive_bucket(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->get_archive_bucket: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetArchiveBucket200Response**](GetArchiveBucket200Response.md)

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

# **get_archive_file**
> get_archive_file(bucket, )

Download an Archive File

Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
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
    api_instance = archives_api.ArchivesApi(api_client)
    bucket = "bucket_example" # str | Bucket name

    # example passing only required values which don't have defaults set
    try:
        # Download an Archive File
        api_instance.get_archive_file(bucket, )
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->get_archive_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **str**| Bucket name |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Returns the contents of the specified file as an attachment with Content-Type dictated by the file |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_archive_file_detail**
> GetArchiveFileDetail200Response get_archive_file_detail(id)

Get Archive File Details

Get details about a specific archive file.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.get_archive_file_detail200_response import GetArchiveFileDetail200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Archive File Details
        api_response = api_instance.get_archive_file_detail(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->get_archive_file_detail: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetArchiveFileDetail200Response**](GetArchiveFileDetail200Response.md)

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

# **get_archive_file_links**
> GetArchiveFileLinks200Response get_archive_file_links(id)

Get Archive File Links

This endpoint retrieves the links that have been created for the specified file.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.get_archive_file_links200_response import GetArchiveFileLinks200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Archive File Links
        api_response = api_instance.get_archive_file_links(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->get_archive_file_links: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetArchiveFileLinks200Response**](GetArchiveFileLinks200Response.md)

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

# **get_archive_public_file**
> get_archive_public_file(bucket, )

Download a Public Archive File

Files in an archive bucket that has Public URL enabled can be downloaded via this endpoint without any authentication, anonymously.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
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
    api_instance = archives_api.ArchivesApi(api_client)
    bucket = "bucket_example" # str | Bucket name

    # example passing only required values which don't have defaults set
    try:
        # Download a Public Archive File
        api_instance.get_archive_public_file(bucket, )
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->get_archive_public_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **str**| Bucket name |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Returns the contents of the specified file as an attachment with Content-Type dicated by the file |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_archive_public_file_link**
> get_archive_public_file_link(s)

Download an Archive File Link

Download an archive file link.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
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
    api_instance = archives_api.ArchivesApi(api_client)
    s = "45a214fce9a546b9" # str | The secret access token for the archive file being downloaded.

    # example passing only required values which don't have defaults set
    try:
        # Download an Archive File Link
        api_instance.get_archive_public_file_link(s)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->get_archive_public_file_link: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **str**| The secret access token for the archive file being downloaded. |

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Returns the contents of the specified file as an attachment with Content-Type dicated by the file |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_archive_buckets**
> ListArchiveBuckets200Response list_archive_buckets()

Get All Archive Buckets

This endpoint retrieves all archive buckets associated with the account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.list_archive_buckets200_response import ListArchiveBuckets200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Archive Buckets
        api_response = api_instance.list_archive_buckets(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->list_archive_buckets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListArchiveBuckets200Response**](ListArchiveBuckets200Response.md)

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

# **list_archive_files**
> ListArchiveFiles200Response list_archive_files(bucket, )

Get All Archive Files

This endpoint retrieves all files in an archive bucket under the specified `filePath`.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_archive_files200_response import ListArchiveFiles200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    bucket = "bucket_example" # str | Bucket name
    name = "test-config.yaml" # str | If specified will return an exact match on file name (optional)
    phrase = "test-%.yaml" # str | Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.  (optional)
    full_tree = False # bool | Include files under sub directories too. This is always true when searching with phrase. (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Get All Archive Files
        api_response = api_instance.list_archive_files(bucket, )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->list_archive_files: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Archive Files
        api_response = api_instance.list_archive_files(bucket, name=name, phrase=phrase, full_tree=full_tree)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->list_archive_files: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **str**| Bucket name |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"
 **name** | **str**| If specified will return an exact match on file name | [optional]
 **phrase** | **str**| Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.  | [optional]
 **full_tree** | **bool**| Include files under sub directories too. This is always true when searching with phrase. | [optional] if omitted the server will use the default value of False

### Return type

[**ListArchiveFiles200Response**](ListArchiveFiles200Response.md)

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

# **update_archive_bucket**
> AddArchiveBucket200Response update_archive_bucket(id)

Update an Archive Bucket

Update an Archive Bucket

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import archives_api
from openapi_client.model.update_archive_bucket_request import UpdateArchiveBucketRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_archive_bucket200_response import AddArchiveBucket200Response
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
    api_instance = archives_api.ArchivesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_archive_bucket_request = UpdateArchiveBucketRequest(
        archive_bucket=ArchiveBucketUpdate(
            name="name_example",
            description="description_example",
            visibility="private",
            is_public=False,
            accounts=UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                id=1,
            ),
        ),
    ) # UpdateArchiveBucketRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Archive Bucket
        api_response = api_instance.update_archive_bucket(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->update_archive_bucket: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Archive Bucket
        api_response = api_instance.update_archive_bucket(id, update_archive_bucket_request=update_archive_bucket_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ArchivesApi->update_archive_bucket: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_archive_bucket_request** | [**UpdateArchiveBucketRequest**](UpdateArchiveBucketRequest.md)|  | [optional]

### Return type

[**AddArchiveBucket200Response**](AddArchiveBucket200Response.md)

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

