# MorpheusApi.AuthenticationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**forgotPassword**](AuthenticationApi.md#forgotPassword) | **POST** /api/forgot/send-email | Request a reset password email
[**getAccessToken**](AuthenticationApi.md#getAccessToken) | **POST** /oauth/token | Provides authentication via username and password
[**resetPassword**](AuthenticationApi.md#resetPassword) | **POST** /api/forgot/reset-password | Reset user password
[**whoami**](AuthenticationApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions



## forgotPassword

> InlineResponse2007 forgotPassword(opts)

Request a reset password email

This endpoint will trigger the Reset your password email to be sent to the specified user.  The User is identified by &#x60;username&#x60; and, if they exist, will be notified via their configured email address. The email notification will indicate a Reset Password Request was made and it will include a token.  Once you obtain the token from the email, it may be used to reset the password of your user.  **Note**: This is an unauthorized endpoint and the response will always appear successful, it is not possible determine from the response whether the user exists or if an email was sent. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';

let apiInstance = new MorpheusApi.AuthenticationApi();
let opts = {
  'inlineObject11': new MorpheusApi.InlineObject11() // InlineObject11 | 
};
apiInstance.forgotPassword(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
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


## getAccessToken

> AccessToken getAccessToken(clientId, grantType, scope, opts)

Provides authentication via username and password

This endpoint provides authentication via &#x60;username&#x60; and &#x60;password&#x60; of a Morpheus User. The response includes a valid access token. If your current token is expired, a new one will be created and returned.  Subtenant users will need to pass their subdomain prefix like subdomain\\username. The default subdomain is the tenant account ID.  This endpoint also allows refreshing your current access token to get a new token. This is done by passing your current &#x60;refresh_token&#x60;. This provides a way to renew your client’s session with the API, and extend the expiration date.  This will render your current access token invalid, so you will need to update any scripts relying on it. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';

let apiInstance = new MorpheusApi.AuthenticationApi();
let clientId = "clientId_example"; // String | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli.
let grantType = "grantType_example"; // String | OAuth Grant Type, use password.
let scope = "scope_example"; // String | OAuth token scope, use write.
let opts = {
  'username': "username_example", // String | Specified as \\\"username\\\" or \\\"tenantId\\\\username\\\" with proper HTML encoding and used in conjunction with `password`. Not utilized with `refresh_token`.
  'password': "password_example", // String | The Password for defined `username`. Must have proper HTML encoding
  'refreshToken': null // Object | The `refresh_token` value from a previous API generation can be utilized instead of `username` and `password`.
};
apiInstance.getAccessToken(clientId, grantType, scope, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **String**| Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli. | 
 **grantType** | **String**| OAuth Grant Type, use password. | 
 **scope** | **String**| OAuth token scope, use write. | 
 **username** | **String**| Specified as \\\&quot;username\\\&quot; or \\\&quot;tenantId\\\\username\\\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;. Not utilized with &#x60;refresh_token&#x60;. | [optional] 
 **password** | **String**| The Password for defined &#x60;username&#x60;. Must have proper HTML encoding | [optional] 
 **refreshToken** | [**Object**](Object.md)| The &#x60;refresh_token&#x60; value from a previous API generation can be utilized instead of &#x60;username&#x60; and &#x60;password&#x60;. | [optional] 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json


## resetPassword

> InlineResponse2006 resetPassword(opts)

Reset user password

This endpoint will reset the password for a user, updating it to the specified value. A secret token must be passed to identify the user who is being updated.  **Note**: You can obtain this token by inspecting the URL of the “Click here to reset” link seen in the email. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';

let apiInstance = new MorpheusApi.AuthenticationApi();
let opts = {
  'inlineObject10': new MorpheusApi.InlineObject10() // InlineObject10 | 
};
apiInstance.resetPassword(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
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


## whoami

> InlineResponse200167 whoami()

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AuthenticationApi();
apiInstance.whoami((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
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

