# AuthenticationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**forgotPassword**](AuthenticationApi.md#forgotPassword) | **POST** /api/forgot/send-email | Request a reset password email
[**getAccessToken**](AuthenticationApi.md#getAccessToken) | **POST** /oauth/token | Provides authentication via username and password
[**resetPassword**](AuthenticationApi.md#resetPassword) | **POST** /api/forgot/reset-password | Reset user password
[**whoami**](AuthenticationApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions


<a name="forgotPassword"></a>
# **forgotPassword**
> InlineResponse2007 forgotPassword(inlineObject11)

Request a reset password email

This endpoint will trigger the Reset your password email to be sent to the specified user.  The User is identified by &#x60;username&#x60; and, if they exist, will be notified via their configured email address. The email notification will indicate a Reset Password Request was made and it will include a token.  Once you obtain the token from the email, it may be used to reset the password of your user.  **Note**: This is an unauthorized endpoint and the response will always appear successful, it is not possible determine from the response whether the user exists or if an email was sent. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    InlineObject11 inlineObject11 = new InlineObject11(); // InlineObject11 | 
    try {
      InlineResponse2007 result = apiInstance.forgotPassword(inlineObject11);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#forgotPassword");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject11** | [**InlineObject11**](InlineObject11.md)|  | [optional]

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

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

<a name="getAccessToken"></a>
# **getAccessToken**
> AccessToken getAccessToken(clientId, grantType, scope, username, password, refreshToken)

Provides authentication via username and password

This endpoint provides authentication via &#x60;username&#x60; and &#x60;password&#x60; of a Morpheus User. The response includes a valid access token. If your current token is expired, a new one will be created and returned.  Subtenant users will need to pass their subdomain prefix like subdomain\\username. The default subdomain is the tenant account ID.  This endpoint also allows refreshing your current access token to get a new token. This is done by passing your current &#x60;refresh_token&#x60;. This provides a way to renew your client’s session with the API, and extend the expiration date.  This will render your current access token invalid, so you will need to update any scripts relying on it. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    String clientId = "clientId_example"; // String | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli.
    String grantType = "grantType_example"; // String | OAuth Grant Type, use password.
    String scope = "scope_example"; // String | OAuth token scope, use write.
    String username = "username_example"; // String | Specified as \\\"username\\\" or \\\"tenantId\\\\username\\\" with proper HTML encoding and used in conjunction with `password`. Not utilized with `refresh_token`.
    String password = "password_example"; // String | The Password for defined `username`. Must have proper HTML encoding
    Object refreshToken = null; // Object | The `refresh_token` value from a previous API generation can be utilized instead of `username` and `password`.
    try {
      AccessToken result = apiInstance.getAccessToken(clientId, grantType, scope, username, password, refreshToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#getAccessToken");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **String**| Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli. | [enum: morph-api]
 **grantType** | **String**| OAuth Grant Type, use password. | [enum: password, refresh_token]
 **scope** | **String**| OAuth token scope, use write. | [enum: write]
 **username** | **String**| Specified as \\\&quot;username\\\&quot; or \\\&quot;tenantId\\\\username\\\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;. Not utilized with &#x60;refresh_token&#x60;. | [optional]
 **password** | **String**| The Password for defined &#x60;username&#x60;. Must have proper HTML encoding | [optional]
 **refreshToken** | [**Object**](Object.md)| The &#x60;refresh_token&#x60; value from a previous API generation can be utilized instead of &#x60;username&#x60; and &#x60;password&#x60;. | [optional] [default to null]

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

<a name="resetPassword"></a>
# **resetPassword**
> InlineResponse2006 resetPassword(inlineObject10)

Reset user password

This endpoint will reset the password for a user, updating it to the specified value. A secret token must be passed to identify the user who is being updated.  **Note**: You can obtain this token by inspecting the URL of the “Click here to reset” link seen in the email. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    InlineObject10 inlineObject10 = new InlineObject10(); // InlineObject10 | 
    try {
      InlineResponse2006 result = apiInstance.resetPassword(inlineObject10);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#resetPassword");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject10** | [**InlineObject10**](InlineObject10.md)|  | [optional]

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

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

<a name="whoami"></a>
# **whoami**
> InlineResponse200167 whoami()

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AuthenticationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AuthenticationApi apiInstance = new AuthenticationApi(defaultClient);
    try {
      InlineResponse200167 result = apiInstance.whoami();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AuthenticationApi#whoami");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200167**](InlineResponse200167.md)

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

