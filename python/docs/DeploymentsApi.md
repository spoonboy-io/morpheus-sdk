# openapi_client.DeploymentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_deployment_file**](DeploymentsApi.md#add_deployment_file) | **POST** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Upload a Deployment File
[**add_deployment_version**](DeploymentsApi.md#add_deployment_version) | **POST** /api/deployments/{deploymentId}/versions | Create a new Deployment Version
[**add_deployments**](DeploymentsApi.md#add_deployments) | **POST** /api/deployments | Create a new Deployment
[**delete_deployment**](DeploymentsApi.md#delete_deployment) | **DELETE** /api/deployments/{deploymentId} | Delete a Deployment
[**delete_deployment_file**](DeploymentsApi.md#delete_deployment_file) | **DELETE** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Delete a Deployment File
[**delete_deployment_version**](DeploymentsApi.md#delete_deployment_version) | **DELETE** /api/deployments/{deploymentId}/versions/{id} | Delete a Deployment Version
[**get_deployment**](DeploymentsApi.md#get_deployment) | **GET** /api/deployments/{deploymentId} | Get a Specific Deployment
[**get_deployment_version**](DeploymentsApi.md#get_deployment_version) | **GET** /api/deployments/{deploymentId}/versions/{id} | Get a Specific Deployment Version
[**list_deployment_files**](DeploymentsApi.md#list_deployment_files) | **GET** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | List Deployment Files
[**list_deployment_versions**](DeploymentsApi.md#list_deployment_versions) | **GET** /api/deployments/{deploymentId}/versions | Get All Versions For a Deployment
[**list_deployments**](DeploymentsApi.md#list_deployments) | **GET** /api/deployments | Get All Deployments
[**update_deployment**](DeploymentsApi.md#update_deployment) | **PUT** /api/deployments/{deploymentId} | Updating a Deployment
[**update_deployment_version**](DeploymentsApi.md#update_deployment_version) | **PUT** /api/deployments/{deploymentId}/versions/{id} | Updating a Deployment Version


# **add_deployment_file**
> Model200Success add_deployment_file(deployment_id, id, )

Upload a Deployment File

This endpoint will upload a file for a specific deployment version. This will overwrite the file if one with the same name exists already.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    id = 1 # int | Morpheus ID of the Object being referenced
    file = open('/path/to/file', 'rb') # file_type |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Upload a Deployment File
        api_response = api_instance.add_deployment_file(deployment_id, id, )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->add_deployment_file: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Upload a Deployment File
        api_response = api_instance.add_deployment_file(deployment_id, id, file=file)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->add_deployment_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"
 **file** | **file_type**|  | [optional]

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

# **add_deployment_version**
> AddDeploymentVersion200Response add_deployment_version(deployment_id)

Create a new Deployment Version

This endpoint will create a new deployment version that is ready to have files uploaded to it. The default type is file, which has files directly uploaded via Morpheus. Alternatively, the type git or fetch can be used to just point to a repository or remote url.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.add_deployment_version200_response import AddDeploymentVersion200Response
from openapi_client.model.add_deployment_version_request import AddDeploymentVersionRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    add_deployment_version_request = AddDeploymentVersionRequest(
        version=DeploymentVersionCreate(
            version="version_example",
            user_version="user_version_example",
            deploy_type="file",
            git_url="git_url_example",
            git_ref="git_ref_example",
            fetch_url="fetch_url_example",
        ),
    ) # AddDeploymentVersionRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a new Deployment Version
        api_response = api_instance.add_deployment_version(deployment_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->add_deployment_version: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a new Deployment Version
        api_response = api_instance.add_deployment_version(deployment_id, add_deployment_version_request=add_deployment_version_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->add_deployment_version: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **add_deployment_version_request** | [**AddDeploymentVersionRequest**](AddDeploymentVersionRequest.md)|  | [optional]

### Return type

[**AddDeploymentVersion200Response**](AddDeploymentVersion200Response.md)

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

# **add_deployments**
> AddDeployments200Response add_deployments()

Create a new Deployment

This endpoint will create a new deployment that is ready to have versions added to it.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.add_deployments_request import AddDeploymentsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_deployments200_response import AddDeployments200Response
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    add_deployments_request = AddDeploymentsRequest(
        deployment=DeploymentCreate(
            name="name_example",
            description="description_example",
        ),
    ) # AddDeploymentsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a new Deployment
        api_response = api_instance.add_deployments(add_deployments_request=add_deployments_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->add_deployments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_deployments_request** | [**AddDeploymentsRequest**](AddDeploymentsRequest.md)|  | [optional]

### Return type

[**AddDeployments200Response**](AddDeployments200Response.md)

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

# **delete_deployment**
> Model200Success delete_deployment(deployment_id)

Delete a Deployment

This endpoint will delete an existing deployment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Deployment
        api_response = api_instance.delete_deployment(deployment_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->delete_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |

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

# **delete_deployment_file**
> Model200Success delete_deployment_file(deployment_id, id, )

Delete a Deployment File

This endpoint will delete an existing deployment file. To recursively delete a directory and all of its contents, the force parameter must be specified.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    id = 1 # int | Morpheus ID of the Object being referenced
    force = "on" # str | Force Delete (optional) if omitted the server will use the default value of "off"

    # example passing only required values which don't have defaults set
    try:
        # Delete a Deployment File
        api_response = api_instance.delete_deployment_file(deployment_id, id, )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->delete_deployment_file: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete a Deployment File
        api_response = api_instance.delete_deployment_file(deployment_id, id, force=force)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->delete_deployment_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"
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

# **delete_deployment_version**
> Model200Success delete_deployment_version(deployment_id, id)

Delete a Deployment Version

This endpoint will delete an existing deployment version.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Deployment Version
        api_response = api_instance.delete_deployment_version(deployment_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->delete_deployment_version: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
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

# **get_deployment**
> GetDeployment200Response get_deployment(deployment_id)

Get a Specific Deployment

This endpoint retrieves a specific deployment. By default the 5 most recent versions are returned, more can be returned by specifying the maxVersions parameter.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.get_deployment200_response import GetDeployment200Response
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    max_versions = 6 # int | Max number of recent versions to return. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Deployment
        api_response = api_instance.get_deployment(deployment_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a Specific Deployment
        api_response = api_instance.get_deployment(deployment_id, max_versions=max_versions)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **max_versions** | **int**| Max number of recent versions to return. | [optional]

### Return type

[**GetDeployment200Response**](GetDeployment200Response.md)

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

# **get_deployment_version**
> AddDeploymentVersion200ResponseAllOf get_deployment_version(deployment_id, id)

Get a Specific Deployment Version

This endpoint retrieves a specific deployment version.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.add_deployment_version200_response_all_of import AddDeploymentVersion200ResponseAllOf
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Deployment Version
        api_response = api_instance.get_deployment_version(deployment_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment_version: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddDeploymentVersion200ResponseAllOf**](AddDeploymentVersion200ResponseAllOf.md)

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

# **list_deployment_files**
> ListDeploymentVersions200Response list_deployment_files(deployment_id, id, )

List Deployment Files

This endpoint returns a list of files for a specific deployment version. This only applies to deploy type file. Files are sorted alphabetically, with directories appearing at the beginning of the list.  The filepath parameter can be specified to search for specific files or directories. To list files under a directory, use a trailing / in the filepath parameter.  To list a specific file, provide it's full path. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_deployment_versions200_response import ListDeploymentVersions200Response
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    id = 1 # int | Morpheus ID of the Object being referenced
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    version = 5 # int | Filter by version number (userVersion) (optional)

    # example passing only required values which don't have defaults set
    try:
        # List Deployment Files
        api_response = api_instance.list_deployment_files(deployment_id, id, )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->list_deployment_files: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List Deployment Files
        api_response = api_instance.list_deployment_files(deployment_id, id, max=max, offset=offset, phrase=phrase, version=version)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->list_deployment_files: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **filepath** | **str**| The path to to search for files under. Default is the root directory /. | defaults to "/"
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]

### Return type

[**ListDeploymentVersions200Response**](ListDeploymentVersions200Response.md)

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

# **list_deployment_versions**
> ListDeploymentVersions200Response list_deployment_versions(deployment_id)

Get All Versions For a Deployment

This endpoint returns a paginated list of versions for a specific deployment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_deployment_versions200_response import ListDeploymentVersions200Response
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    version = 5 # int | Filter by version number (userVersion) (optional)
    type = "file" # str | Filter by type (deployType), file, git, fetch (optional)
    date_created = "2019-01-01" # str | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Versions For a Deployment
        api_response = api_instance.list_deployment_versions(deployment_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->list_deployment_versions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Versions For a Deployment
        api_response = api_instance.list_deployment_versions(deployment_id, max=max, offset=offset, phrase=phrase, version=version, type=type, date_created=date_created, last_updated=last_updated)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->list_deployment_versions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **type** | **str**| Filter by type (deployType), file, git, fetch | [optional]
 **date_created** | **str**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

### Return type

[**ListDeploymentVersions200Response**](ListDeploymentVersions200Response.md)

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

# **list_deployments**
> ListDeployments200Response list_deployments()

Get All Deployments

This endpoint returns a paginated list of deployments.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.list_deployments200_response import ListDeployments200Response
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    description = "The desription of my cool object" # str | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    date_created = "2019-01-01" # str | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Deployments
        api_response = api_instance.list_deployments(max=max, offset=offset, phrase=phrase, name=name, description=description, date_created=date_created, last_updated=last_updated)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->list_deployments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **description** | **str**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **date_created** | **str**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

### Return type

[**ListDeployments200Response**](ListDeployments200Response.md)

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

# **update_deployment**
> UpdateDeployment200Response update_deployment(deployment_id)

Updating a Deployment

This endpoint will update an existing deployment.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.update_deployment200_response import UpdateDeployment200Response
from openapi_client.model.add_deployments_request import AddDeploymentsRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    add_deployments_request = AddDeploymentsRequest(
        deployment=DeploymentCreate(
            name="name_example",
            description="description_example",
        ),
    ) # AddDeploymentsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a Deployment
        api_response = api_instance.update_deployment(deployment_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->update_deployment: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a Deployment
        api_response = api_instance.update_deployment(deployment_id, add_deployments_request=add_deployments_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->update_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **add_deployments_request** | [**AddDeploymentsRequest**](AddDeploymentsRequest.md)|  | [optional]

### Return type

[**UpdateDeployment200Response**](UpdateDeployment200Response.md)

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

# **update_deployment_version**
> AddDeploymentVersion200Response update_deployment_version(deployment_id, id)

Updating a Deployment Version

This endpoint will update an existing deployment version.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deployments_api
from openapi_client.model.add_deployment_version200_response import AddDeploymentVersion200Response
from openapi_client.model.add_deployment_version_request import AddDeploymentVersionRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = 4 # int | Deployment ID
    id = 1 # int | Morpheus ID of the Object being referenced
    add_deployment_version_request = AddDeploymentVersionRequest(
        version=DeploymentVersionCreate(
            version="version_example",
            user_version="user_version_example",
            deploy_type="file",
            git_url="git_url_example",
            git_ref="git_ref_example",
            fetch_url="fetch_url_example",
        ),
    ) # AddDeploymentVersionRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updating a Deployment Version
        api_response = api_instance.update_deployment_version(deployment_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->update_deployment_version: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updating a Deployment Version
        api_response = api_instance.update_deployment_version(deployment_id, id, add_deployment_version_request=add_deployment_version_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploymentsApi->update_deployment_version: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **int**| Deployment ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_deployment_version_request** | [**AddDeploymentVersionRequest**](AddDeploymentVersionRequest.md)|  | [optional]

### Return type

[**AddDeploymentVersion200Response**](AddDeploymentVersion200Response.md)

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

