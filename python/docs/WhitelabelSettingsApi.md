# openapi_client.WhitelabelSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_whitelabel_image**](WhitelabelSettingsApi.md#get_whitelabel_image) | **GET** /api/whitelabel-settings/images/{imageType} | Download Image
[**list_whitelabel_settings**](WhitelabelSettingsApi.md#list_whitelabel_settings) | **GET** /api/whitelabel-settings | Get Whitelabel Settings
[**remove_whitelabel_image**](WhitelabelSettingsApi.md#remove_whitelabel_image) | **DELETE** /api/whitelabel-settings/images/{imageType} | Reset Image
[**update_whitelabel_images**](WhitelabelSettingsApi.md#update_whitelabel_images) | **POST** /api/whitelabel-settings/images | Update Images
[**update_whitelabel_settings**](WhitelabelSettingsApi.md#update_whitelabel_settings) | **PUT** /api/whitelabel-settings | Update Whitelabel Settings


# **get_whitelabel_image**
> file_type get_whitelabel_image(image_type)

Download Image

Downloads the specified image.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import whitelabel_settings_api
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
    api_instance = whitelabel_settings_api.WhitelabelSettingsApi(api_client)
    image_type = "headerLogo" # str | Valid image types

    # example passing only required values which don't have defaults set
    try:
        # Download Image
        api_response = api_instance.get_whitelabel_image(image_type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WhitelabelSettingsApi->get_whitelabel_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_type** | **str**| Valid image types |

### Return type

**file_type**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/ico, image/jpeg, image/png, image/svg+xml, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_whitelabel_settings**
> ListWhitelabelSettings200Response list_whitelabel_settings()

Get Whitelabel Settings

This endpoint retrieves whitelabel settings.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import whitelabel_settings_api
from openapi_client.model.list_whitelabel_settings200_response import ListWhitelabelSettings200Response
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
    api_instance = whitelabel_settings_api.WhitelabelSettingsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get Whitelabel Settings
        api_response = api_instance.list_whitelabel_settings()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WhitelabelSettingsApi->list_whitelabel_settings: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListWhitelabelSettings200Response**](ListWhitelabelSettings200Response.md)

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

# **remove_whitelabel_image**
> Model200Success remove_whitelabel_image(image_type)

Reset Image

Resets the specified image to the Morpheus default.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import whitelabel_settings_api
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
    api_instance = whitelabel_settings_api.WhitelabelSettingsApi(api_client)
    image_type = "headerLogo" # str | Valid image types

    # example passing only required values which don't have defaults set
    try:
        # Reset Image
        api_response = api_instance.remove_whitelabel_image(image_type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WhitelabelSettingsApi->remove_whitelabel_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image_type** | **str**| Valid image types |

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

# **update_whitelabel_images**
> Model200Success update_whitelabel_images()

Update Images

Uploads whitelabel images. Expects multipart form data as the request format, not JSON.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import whitelabel_settings_api
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
    api_instance = whitelabel_settings_api.WhitelabelSettingsApi(api_client)
    header_logo_file = open('/path/to/file', 'rb') # file_type | Header logo image file, valid image types `png|jpg|svg` (optional)
    reset_header_logo = True # bool | Resets header logo to default (optional)
    footer_logo_file = open('/path/to/file', 'rb') # file_type | Footer logo image file, valid image types `png|jpg|svg` (optional)
    reset_footer_logo = True # bool | Resets footer logo to default (optional)
    login_logo_file = open('/path/to/file', 'rb') # file_type | Login logo image file, valid image types `png|jpg|svg` (optional)
    reset_login_logo = True # bool | Resets login logo to default (optional)
    favicon_file = open('/path/to/file', 'rb') # file_type | Favicon image file, valid image type ico (optional)
    reset_favicon_logo = True # bool | Resets favicon logo to default (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Images
        api_response = api_instance.update_whitelabel_images(header_logo_file=header_logo_file, reset_header_logo=reset_header_logo, footer_logo_file=footer_logo_file, reset_footer_logo=reset_footer_logo, login_logo_file=login_logo_file, reset_login_logo=reset_login_logo, favicon_file=favicon_file, reset_favicon_logo=reset_favicon_logo)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WhitelabelSettingsApi->update_whitelabel_images: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **header_logo_file** | **file_type**| Header logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **reset_header_logo** | **bool**| Resets header logo to default | [optional]
 **footer_logo_file** | **file_type**| Footer logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **reset_footer_logo** | **bool**| Resets footer logo to default | [optional]
 **login_logo_file** | **file_type**| Login logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **reset_login_logo** | **bool**| Resets login logo to default | [optional]
 **favicon_file** | **file_type**| Favicon image file, valid image type ico | [optional]
 **reset_favicon_logo** | **bool**| Resets favicon logo to default | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_whitelabel_settings**
> Model200Success update_whitelabel_settings()

Update Whitelabel Settings

Update Whitelabel Settings

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import whitelabel_settings_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_whitelabel_settings_request import UpdateWhitelabelSettingsRequest
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
    api_instance = whitelabel_settings_api.WhitelabelSettingsApi(api_client)
    update_whitelabel_settings_request = UpdateWhitelabelSettingsRequest(
        whitelabel_settings=WhitelabelSettingsUpdate(
            enabled=True,
            appliance_name="appliance_name_example",
            disable_support_menu=True,
            reset_header_logo=True,
            reset_footer_logo=True,
            reset_login_logo=True,
            reset_favicon=True,
            header_bg_color="header_bg_color_example",
            header_fg_color="header_fg_color_example",
            nav_bg_color="nav_bg_color_example",
            nav_fg_color="nav_fg_color_example",
            nav_hover_color="nav_hover_color_example",
            primary_button_bg_color="primary_button_bg_color_example",
            primary_button_fg_color="primary_button_fg_color_example",
            primary_button_hover_bg_color="primary_button_hover_bg_color_example",
            primary_button_hover_fg_color="primary_button_hover_fg_color_example",
            footer_bg_color="footer_bg_color_example",
            footer_fg_color="footer_fg_color_example",
            login_bg_color="login_bg_color_example",
            copyright_string="copyright_string_example",
            override_css="override_css_example",
            terms_of_use="terms_of_use_example",
            privacy_policy="privacy_policy_example",
            support_menu_links=[
                WhitelabelSettingsUpdateSupportMenuLinksInner(
                    url="url_example",
                    label="label_example",
                    label_code="label_code_example",
                ),
            ],
        ),
    ) # UpdateWhitelabelSettingsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Whitelabel Settings
        api_response = api_instance.update_whitelabel_settings(update_whitelabel_settings_request=update_whitelabel_settings_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling WhitelabelSettingsApi->update_whitelabel_settings: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_whitelabel_settings_request** | [**UpdateWhitelabelSettingsRequest**](UpdateWhitelabelSettingsRequest.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

