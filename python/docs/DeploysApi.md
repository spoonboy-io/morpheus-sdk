# openapi_client.DeploysApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_instance_deploy**](DeploysApi.md#add_instance_deploy) | **POST** /api/instances/{id}/deploys | Deploy to an Instance
[**deletedeploy**](DeploysApi.md#deletedeploy) | **DELETE** /api/deploys/{id} | Delete a Deploy
[**get_instance_deploys**](DeploysApi.md#get_instance_deploys) | **GET** /api/instances/{id}/deploys | Get all Deploys for an Instance
[**list_deploys**](DeploysApi.md#list_deploys) | **GET** /api/deploys | Get all Deploys
[**run_deploy**](DeploysApi.md#run_deploy) | **POST** /api/deploys/{id}/deploy | Run a Deploy
[**update_deploy**](DeploysApi.md#update_deploy) | **PUT** /api/deploys/{id} | Update a Deploy


# **add_instance_deploy**
> UpdateDeploy200Response add_instance_deploy(id)

Deploy to an Instance

This endpoint will deploy the specified deployment version to specified instance. The version to deploy can be identified with deploymentId and version or with versionId alone.  By default, the deployment is executed right away. To prevent this so that it can be run manually later on. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deploys_api
from openapi_client.model.update_deploy200_response import UpdateDeploy200Response
from openapi_client.model.add_instance_deploy_request import AddInstanceDeployRequest
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
    api_instance = deploys_api.DeploysApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_instance_deploy_request = AddInstanceDeployRequest(
        app_deploy=AddInstanceDeployRequestAppDeploy(
            deployment_id=1,
            version="version_example",
            version_id=1,
            config={},
            stage_only=False,
        ),
    ) # AddInstanceDeployRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Deploy to an Instance
        api_response = api_instance.add_instance_deploy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->add_instance_deploy: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Deploy to an Instance
        api_response = api_instance.add_instance_deploy(id, add_instance_deploy_request=add_instance_deploy_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->add_instance_deploy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_instance_deploy_request** | [**AddInstanceDeployRequest**](AddInstanceDeployRequest.md)|  | [optional]

### Return type

[**UpdateDeploy200Response**](UpdateDeploy200Response.md)

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

# **deletedeploy**
> Model200Success deletedeploy(id)

Delete a Deploy

This endpoint will delete an archived instance deploy.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deploys_api
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
    api_instance = deploys_api.DeploysApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Deploy
        api_response = api_instance.deletedeploy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->deletedeploy: %s\n" % e)
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

# **get_instance_deploys**
> ListDeploys200Response get_instance_deploys(id)

Get all Deploys for an Instance

This endpoint retrieves all deploys for a specific instance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deploys_api
from openapi_client.model.list_deploys200_response import ListDeploys200Response
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
    api_instance = deploys_api.DeploysApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    deployment_id = 5 # int | Filter by deployment id (optional)
    instance_name = "instance1" # str | Filter by instance name (optional)
    instance_id = 94 # int | The Instance ID for Filtering (optional)
    version = 5 # int | Filter by version number (userVersion) (optional)
    version_id = 5 # int | Filter by deployment version id (optional)
    created_by_id = 63 # int | Filter by owner (user) id (optional)
    deploy_type = "file" # str | Filter by type (deployType), file, git, fetch (optional)
    date_created = "2019-01-01" # str | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deploy_date = "2020-01-01" # str | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified (optional)
    status = "deploying" # str | Filter by status (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Deploys for an Instance
        api_response = api_instance.get_instance_deploys(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->get_instance_deploys: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Deploys for an Instance
        api_response = api_instance.get_instance_deploys(id, max=max, offset=offset, phrase=phrase, name=name, deployment_id=deployment_id, instance_name=instance_name, instance_id=instance_id, version=version, version_id=version_id, created_by_id=created_by_id, deploy_type=deploy_type, date_created=date_created, last_updated=last_updated, deploy_date=deploy_date, status=status)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->get_instance_deploys: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **deployment_id** | **int**| Filter by deployment id | [optional]
 **instance_name** | **str**| Filter by instance name | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **version_id** | **int**| Filter by deployment version id | [optional]
 **created_by_id** | **int**| Filter by owner (user) id | [optional]
 **deploy_type** | **str**| Filter by type (deployType), file, git, fetch | [optional]
 **date_created** | **str**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deploy_date** | **str**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional]
 **status** | **str**| Filter by status | [optional]

### Return type

[**ListDeploys200Response**](ListDeploys200Response.md)

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

# **list_deploys**
> ListDeploys200Response list_deploys()

Get all Deploys

This endpoint retrieves all deploys.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deploys_api
from openapi_client.model.list_deploys200_response import ListDeploys200Response
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
    api_instance = deploys_api.DeploysApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    deployment_id = 5 # int | Filter by deployment id (optional)
    instance_name = "instance1" # str | Filter by instance name (optional)
    instance_id = 94 # int | The Instance ID for Filtering (optional)
    version = 5 # int | Filter by version number (userVersion) (optional)
    version_id = 5 # int | Filter by deployment version id (optional)
    created_by_id = 63 # int | Filter by owner (user) id (optional)
    deploy_type = "file" # str | Filter by type (deployType), file, git, fetch (optional)
    date_created = "2019-01-01" # str | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deploy_date = "2020-01-01" # str | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified (optional)
    status = "deploying" # str | Filter by status (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Deploys
        api_response = api_instance.list_deploys(max=max, offset=offset, phrase=phrase, name=name, deployment_id=deployment_id, instance_name=instance_name, instance_id=instance_id, version=version, version_id=version_id, created_by_id=created_by_id, deploy_type=deploy_type, date_created=date_created, last_updated=last_updated, deploy_date=deploy_date, status=status)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->list_deploys: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **deployment_id** | **int**| Filter by deployment id | [optional]
 **instance_name** | **str**| Filter by instance name | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **version** | **int**| Filter by version number (userVersion) | [optional]
 **version_id** | **int**| Filter by deployment version id | [optional]
 **created_by_id** | **int**| Filter by owner (user) id | [optional]
 **deploy_type** | **str**| Filter by type (deployType), file, git, fetch | [optional]
 **date_created** | **str**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deploy_date** | **str**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional]
 **status** | **str**| Filter by status | [optional]

### Return type

[**ListDeploys200Response**](ListDeploys200Response.md)

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

# **run_deploy**
> UpdateDeploy200Response run_deploy(id)

Run a Deploy

This endpoint will run an existing instance deploy. This is for running a new staged deploy or to rollback to previous version by re-running a deploy that is archived.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deploys_api
from openapi_client.model.update_deploy200_response import UpdateDeploy200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_deploy_request import UpdateDeployRequest
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
    api_instance = deploys_api.DeploysApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_deploy_request = UpdateDeployRequest(
        app_deploy=UpdateDeployRequestAppDeploy(
            config={},
        ),
    ) # UpdateDeployRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Run a Deploy
        api_response = api_instance.run_deploy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->run_deploy: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Run a Deploy
        api_response = api_instance.run_deploy(id, update_deploy_request=update_deploy_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->run_deploy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_deploy_request** | [**UpdateDeployRequest**](UpdateDeployRequest.md)|  | [optional]

### Return type

[**UpdateDeploy200Response**](UpdateDeploy200Response.md)

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

# **update_deploy**
> UpdateDeploy200Response update_deploy(id)

Update a Deploy

This endpoint will update an existing deploy. This is typically only needed to change settings on a deploy that is staged, before it is run.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import deploys_api
from openapi_client.model.update_deploy200_response import UpdateDeploy200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_deploy_request import UpdateDeployRequest
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
    api_instance = deploys_api.DeploysApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_deploy_request = UpdateDeployRequest(
        app_deploy=UpdateDeployRequestAppDeploy(
            config={},
        ),
    ) # UpdateDeployRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Deploy
        api_response = api_instance.update_deploy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->update_deploy: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Deploy
        api_response = api_instance.update_deploy(id, update_deploy_request=update_deploy_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling DeploysApi->update_deploy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_deploy_request** | [**UpdateDeployRequest**](UpdateDeployRequest.md)|  | [optional]

### Return type

[**UpdateDeploy200Response**](UpdateDeploy200Response.md)

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

