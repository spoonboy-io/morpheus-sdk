# openapi_client.ProvisioningLicensesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_provisioning_license**](ProvisioningLicensesApi.md#add_provisioning_license) | **POST** /api/provisioning-licenses | Create a License
[**get_provisioning_license**](ProvisioningLicensesApi.md#get_provisioning_license) | **GET** /api/provisioning-licenses/{id} | Get a Specific License
[**get_provisioning_license_reservations**](ProvisioningLicensesApi.md#get_provisioning_license_reservations) | **GET** /api/provisioning-licenses/{id}/reservations | Get Reservations for Specific License
[**list_provisioning_licenses**](ProvisioningLicensesApi.md#list_provisioning_licenses) | **GET** /api/provisioning-licenses | Get All Licenses
[**remove_provisioning_license**](ProvisioningLicensesApi.md#remove_provisioning_license) | **DELETE** /api/provisioning-licenses/{id} | Delete a License
[**update_provisioning_license**](ProvisioningLicensesApi.md#update_provisioning_license) | **PUT** /api/provisioning-licenses/{id} | Update a License


# **add_provisioning_license**
> AddProvisioningLicense200Response add_provisioning_license()

Create a License

Use this command to create a new license.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_licenses_api
from openapi_client.model.add_provisioning_license_request import AddProvisioningLicenseRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_provisioning_license200_response import AddProvisioningLicense200Response
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
    api_instance = provisioning_licenses_api.ProvisioningLicensesApi(api_client)
    add_provisioning_license_request = AddProvisioningLicenseRequest(
        license=ProvisioningLicensesCreate(
            name="name_example",
            license_type="win",
            license_key="license_key_example",
            org_name="org_name_example",
            full_name="full_name_example",
            license_version="license_version_example",
            copies=1,
            description="description_example",
            virtual_images=[
                1,
            ],
            tenants=[
                1,
            ],
        ),
    ) # AddProvisioningLicenseRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a License
        api_response = api_instance.add_provisioning_license(add_provisioning_license_request=add_provisioning_license_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->add_provisioning_license: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_provisioning_license_request** | [**AddProvisioningLicenseRequest**](AddProvisioningLicenseRequest.md)|  | [optional]

### Return type

[**AddProvisioningLicense200Response**](AddProvisioningLicense200Response.md)

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

# **get_provisioning_license**
> GetProvisioningLicense200Response get_provisioning_license(id)

Get a Specific License

This endpoint retrieves a specific license.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_licenses_api
from openapi_client.model.get_provisioning_license200_response import GetProvisioningLicense200Response
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
    api_instance = provisioning_licenses_api.ProvisioningLicensesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific License
        api_response = api_instance.get_provisioning_license(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->get_provisioning_license: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetProvisioningLicense200Response**](GetProvisioningLicense200Response.md)

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

# **get_provisioning_license_reservations**
> GetProvisioningLicenseReservations200Response get_provisioning_license_reservations(id)

Get Reservations for Specific License

This endpoint retrieves all reservations for a specific license. Each time a license is applied to a new server, a reservation is created, reducing the available copies for the license.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_licenses_api
from openapi_client.model.get_provisioning_license_reservations200_response import GetProvisioningLicenseReservations200Response
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
    api_instance = provisioning_licenses_api.ProvisioningLicensesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Reservations for Specific License
        api_response = api_instance.get_provisioning_license_reservations(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->get_provisioning_license_reservations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetProvisioningLicenseReservations200Response**](GetProvisioningLicenseReservations200Response.md)

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

# **list_provisioning_licenses**
> ListProvisioningLicenses200Response list_provisioning_licenses()

Get All Licenses

This endpoint retrieves all licenses.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_licenses_api
from openapi_client.model.list_provisioning_licenses200_response import ListProvisioningLicenses200Response
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
    api_instance = provisioning_licenses_api.ProvisioningLicensesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    license_type = "win" # str | If specified will return an exact match on licenseType code (optional)
    license_version = "2012 R2 Standard" # str | If specified will return an exact match on licenseVersion (optional)
    org_name = "Acme Motors" # str | If specified will return an exact match on orgName (optional)
    full_name = "Bugs Bunny" # str | If specified will return an exact match on fullName (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Licenses
        api_response = api_instance.list_provisioning_licenses(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, license_type=license_type, license_version=license_version, org_name=org_name, full_name=full_name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->list_provisioning_licenses: %s\n" % e)
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
 **license_type** | **str**| If specified will return an exact match on licenseType code | [optional]
 **license_version** | **str**| If specified will return an exact match on licenseVersion | [optional]
 **org_name** | **str**| If specified will return an exact match on orgName | [optional]
 **full_name** | **str**| If specified will return an exact match on fullName | [optional]

### Return type

[**ListProvisioningLicenses200Response**](ListProvisioningLicenses200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_provisioning_license**
> Model200Success remove_provisioning_license(id)

Delete a License

Will delete a license.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_licenses_api
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
    api_instance = provisioning_licenses_api.ProvisioningLicensesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a License
        api_response = api_instance.remove_provisioning_license(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->remove_provisioning_license: %s\n" % e)
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

# **update_provisioning_license**
> AddProvisioningLicense200Response update_provisioning_license(id)

Update a License

Use this command to update an existing license.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import provisioning_licenses_api
from openapi_client.model.update_provisioning_license_request import UpdateProvisioningLicenseRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_provisioning_license200_response import AddProvisioningLicense200Response
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
    api_instance = provisioning_licenses_api.ProvisioningLicensesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_provisioning_license_request = UpdateProvisioningLicenseRequest(
        license=ProvisioningLicensesUpdate(
            name="name_example",
            license_version="license_version_example",
            copies=1,
            description="description_example",
            virtual_images=[
                1,
            ],
            tenants=[
                1,
            ],
        ),
    ) # UpdateProvisioningLicenseRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a License
        api_response = api_instance.update_provisioning_license(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->update_provisioning_license: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a License
        api_response = api_instance.update_provisioning_license(id, update_provisioning_license_request=update_provisioning_license_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling ProvisioningLicensesApi->update_provisioning_license: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_provisioning_license_request** | [**UpdateProvisioningLicenseRequest**](UpdateProvisioningLicenseRequest.md)|  | [optional]

### Return type

[**AddProvisioningLicense200Response**](AddProvisioningLicense200Response.md)

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

