# MorpheusApi.AuthenticationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**forgotPassword**](AuthenticationApi.md#forgotPassword) | **POST** /api/forgot/send-email | Request a reset password email
[**resetPassword**](AuthenticationApi.md#resetPassword) | **POST** /api/forgot/reset-password | Reset user password



## forgotPassword

> ForgotPassword200Response forgotPassword(opts)

Request a reset password email

This endpoint will trigger the Reset your password email to be sent to the specified user.  The User is identified by &#x60;username&#x60; and, if they exist, will be notified via their configured email address. The email notification will indicate a Reset Password Request was made and it will include a token.  Once you obtain the token from the email, it may be used to reset the password of your user.  **Note**: This is an unauthorized endpoint and the response will always appear successful, it is not possible determine from the response whether the user exists or if an email was sent. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';

let apiInstance = new MorpheusApi.AuthenticationApi();
let opts = {
  'forgotPasswordRequest': new MorpheusApi.ForgotPasswordRequest() // ForgotPasswordRequest | 
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
 **forgotPasswordRequest** | [**ForgotPasswordRequest**](ForgotPasswordRequest.md)|  | [optional] 

### Return type

[**ForgotPassword200Response**](ForgotPassword200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## resetPassword

> ResetPassword200Response resetPassword(opts)

Reset user password

This endpoint will reset the password for a user, updating it to the specified value. A secret token must be passed to identify the user who is being updated.  **Note**: You can obtain this token by inspecting the URL of the “Click here to reset” link seen in the email. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';

let apiInstance = new MorpheusApi.AuthenticationApi();
let opts = {
  'resetPasswordRequest': new MorpheusApi.ResetPasswordRequest() // ResetPasswordRequest | 
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
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md)|  | [optional] 

### Return type

[**ResetPassword200Response**](ResetPassword200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

