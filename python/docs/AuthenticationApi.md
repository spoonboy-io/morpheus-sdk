# openapi_client.AuthenticationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**forgot_password**](AuthenticationApi.md#forgot_password) | **POST** /api/forgot/send-email | Request a reset password email
[**get_access_token**](AuthenticationApi.md#get_access_token) | **POST** /oauth/token | Provides authentication via username and password
[**reset_password**](AuthenticationApi.md#reset_password) | **POST** /api/forgot/reset-password | Reset user password
[**whoami**](AuthenticationApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions


# **forgot_password**
> ForgotPassword200Response forgot_password()

Request a reset password email

This endpoint will trigger the Reset your password email to be sent to the specified user.  The User is identified by `username` and, if they exist, will be notified via their configured email address. The email notification will indicate a Reset Password Request was made and it will include a token.  Once you obtain the token from the email, it may be used to reset the password of your user.  **Note**: This is an unauthorized endpoint and the response will always appear successful, it is not possible determine from the response whether the user exists or if an email was sent. 

### Example


```python
import time
import openapi_client
from openapi_client.api import authentication_api
from openapi_client.model.forgot_password_request import ForgotPasswordRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.forgot_password200_response import ForgotPassword200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = authentication_api.AuthenticationApi(api_client)
    forgot_password_request = ForgotPasswordRequest(
        username="username_example",
    ) # ForgotPasswordRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Request a reset password email
        api_response = api_instance.forgot_password(forgot_password_request=forgot_password_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AuthenticationApi->forgot_password: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgot_password_request** | [**ForgotPasswordRequest**](ForgotPasswordRequest.md)|  | [optional]

### Return type

[**ForgotPassword200Response**](ForgotPassword200Response.md)

### Authorization

No authorization required

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

# **get_access_token**
> AccessToken get_access_token(grant_type, )

Provides authentication via username and password

This endpoint provides authentication via `username` and `password` of a Morpheus User. The response includes a valid access token. If your current token is expired, a new one will be created and returned.  Subtenant users will need to pass their subdomain prefix like subdomain\\username. The default subdomain is the tenant account ID.  This endpoint also allows refreshing your current access token to get a new token. This is done by passing your current `refresh_token`. This provides a way to renew your client’s session with the API, and extend the expiration date.  This will render your current access token invalid, so you will need to update any scripts relying on it. 

### Example


```python
import time
import openapi_client
from openapi_client.api import authentication_api
from openapi_client.model.access_token import AccessToken
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = authentication_api.AuthenticationApi(api_client)
    grant_type = "password" # str | OAuth Grant Type, use password.
    username = "username_example" # str | Specified as \\\"username\\\" or \\\"tenantId\\\\username\\\" with proper HTML encoding and used in conjunction with `password`. Not utilized with `refresh_token`. (optional)
    password = "password_example" # str | The Password for defined `username`. Must have proper HTML encoding (optional)
    refresh_token = None # bool, date, datetime, dict, float, int, list, str, none_type | The `refresh_token` value from a previous API generation can be utilized instead of `username` and `password`. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Provides authentication via username and password
        api_response = api_instance.get_access_token(grant_type, )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AuthenticationApi->get_access_token: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Provides authentication via username and password
        api_response = api_instance.get_access_token(grant_type, username=username, password=password, refresh_token=refresh_token)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AuthenticationApi->get_access_token: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grant_type** | **str**| OAuth Grant Type, use password. |
 **client_id** | **str**| Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli. | defaults to "morph-api"
 **scope** | **str**| OAuth token scope, use write. | defaults to "write"
 **username** | **str**| Specified as \\\&quot;username\\\&quot; or \\\&quot;tenantId\\\\username\\\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;. Not utilized with &#x60;refresh_token&#x60;. | [optional]
 **password** | **str**| The Password for defined &#x60;username&#x60;. Must have proper HTML encoding | [optional]
 **refresh_token** | **bool, date, datetime, dict, float, int, list, str, none_type**| The &#x60;refresh_token&#x60; value from a previous API generation can be utilized instead of &#x60;username&#x60; and &#x60;password&#x60;. | [optional]

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reset_password**
> ResetPassword200Response reset_password()

Reset user password

This endpoint will reset the password for a user, updating it to the specified value. A secret token must be passed to identify the user who is being updated.  **Note**: You can obtain this token by inspecting the URL of the “Click here to reset” link seen in the email. 

### Example


```python
import time
import openapi_client
from openapi_client.api import authentication_api
from openapi_client.model.reset_password200_response import ResetPassword200Response
from openapi_client.model.reset_password_request import ResetPasswordRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = authentication_api.AuthenticationApi(api_client)
    reset_password_request = ResetPasswordRequest(
        token="token_example",
        password="password_example",
    ) # ResetPasswordRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Reset user password
        api_response = api_instance.reset_password(reset_password_request=reset_password_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AuthenticationApi->reset_password: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reset_password_request** | [**ResetPasswordRequest**](ResetPasswordRequest.md)|  | [optional]

### Return type

[**ResetPassword200Response**](ResetPassword200Response.md)

### Authorization

No authorization required

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

# **whoami**
> Whoami200Response whoami()

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import authentication_api
from openapi_client.model.whoami200_response import Whoami200Response
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
    api_instance = authentication_api.AuthenticationApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves information about current user roles and permissions
        api_response = api_instance.whoami()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AuthenticationApi->whoami: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**Whoami200Response**](Whoami200Response.md)

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

