# openapi_client.LicenseApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_license**](LicenseApi.md#get_license) | **GET** /api/license | Get license
[**install_license**](LicenseApi.md#install_license) | **POST** /api/license | Install license key
[**test_license**](LicenseApi.md#test_license) | **POST** /api/license/test | Test license key
[**uninstall_license**](LicenseApi.md#uninstall_license) | **DELETE** /api/license | Uninstall license key


# **get_license**
> License get_license()

Get license

Get details about the license that is currently installed on the appliance. This returns the license details, but not the key itself. Your license key will never be returned and should always be kept secret.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import license_api
from openapi_client.model.license import License
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
    api_instance = license_api.LicenseApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get license
        api_response = api_instance.get_license()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LicenseApi->get_license: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Get details about the currently installed license |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **install_license**
> License install_license()

Install license key

Install a new license key. This will potentially change the enabled features and capabilities of your Morpheus appliance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import license_api
from openapi_client.model.license import License
from openapi_client.model.install_license_request import InstallLicenseRequest
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
    api_instance = license_api.LicenseApi(api_client)
    install_license_request = InstallLicenseRequest(
        license="license_example",
    ) # InstallLicenseRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Install license key
        api_response = api_instance.install_license(install_license_request=install_license_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LicenseApi->install_license: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **install_license_request** | [**InstallLicenseRequest**](InstallLicenseRequest.md)|  | [optional]

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | License installed successfully. |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **test_license**
> License test_license()

Test license key

This endpoint can be used to decode a license to see if it is valid and inspect the license settings, such as who it belongs to and the enabled features. This is only a test, it does not install the key, or make any changes to your appliance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import license_api
from openapi_client.model.license import License
from openapi_client.model.install_license_request import InstallLicenseRequest
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
    api_instance = license_api.LicenseApi(api_client)
    install_license_request = InstallLicenseRequest(
        license="license_example",
    ) # InstallLicenseRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Test license key
        api_response = api_instance.test_license(install_license_request=install_license_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LicenseApi->test_license: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **install_license_request** | [**InstallLicenseRequest**](InstallLicenseRequest.md)|  | [optional]

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | License test succeeded, the license details are returned |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **uninstall_license**
> Model200Success uninstall_license()

Uninstall license key

Uninstall your appliance license, leaving the appliance with no license installed. This will change the enabled features and capabilities of your Morpheus appliance.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import license_api
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
    api_instance = license_api.LicenseApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Uninstall license key
        api_response = api_instance.uninstall_license()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LicenseApi->uninstall_license: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

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
**200** | License uninstalled successfully. |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

