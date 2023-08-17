# PSOpenAPITools.PSOpenAPITools/Api.AuthenticationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Invoke-ForgotPassword**](AuthenticationApi.md#Invoke-ForgotPassword) | **POST** /api/forgot/send-email | Request a reset password email
[**Reset-Password**](AuthenticationApi.md#Reset-Password) | **POST** /api/forgot/reset-password | Reset user password


<a id="Invoke-ForgotPassword"></a>
# **Invoke-ForgotPassword**
> ForgotPassword200Response Invoke-ForgotPassword<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ForgotPasswordRequest] <PSCustomObject><br>

Request a reset password email

This endpoint will trigger the Reset your password email to be sent to the specified user.  The User is identified by `username` and, if they exist, will be notified via their configured email address. The email notification will indicate a Reset Password Request was made and it will include a token.  Once you obtain the token from the email, it may be used to reset the password of your user.  **Note**: This is an unauthorized endpoint and the response will always appear successful, it is not possible determine from the response whether the user exists or if an email was sent. 

### Example
```powershell
$ForgotPasswordRequest = Initialize-ForgotPasswordRequest -Username "MyUsername" # ForgotPasswordRequest |  (optional)

# Request a reset password email
try {
    $Result = Invoke-ForgotPassword -ForgotPasswordRequest $ForgotPasswordRequest
} catch {
    Write-Host ("Exception occurred when calling Invoke-ForgotPassword: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ForgotPasswordRequest** | [**ForgotPasswordRequest**](ForgotPasswordRequest.md)|  | [optional] 

### Return type

[**ForgotPassword200Response**](ForgotPassword200Response.md) (PSCustomObject)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Reset-Password"></a>
# **Reset-Password**
> ResetPassword200Response Reset-Password<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ResetPasswordRequest] <PSCustomObject><br>

Reset user password

This endpoint will reset the password for a user, updating it to the specified value. A secret token must be passed to identify the user who is being updated.  **Note**: You can obtain this token by inspecting the URL of the “Click here to reset” link seen in the email. 

### Example
```powershell
$ResetPasswordRequest = Initialize-ResetPasswordRequest -Token "MyToken" -Password "MyPassword" # ResetPasswordRequest |  (optional)

# Reset user password
try {
    $Result = Reset-Password -ResetPasswordRequest $ResetPasswordRequest
} catch {
    Write-Host ("Exception occurred when calling Reset-Password: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ResetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md)|  | [optional] 

### Return type

[**ResetPassword200Response**](ResetPassword200Response.md) (PSCustomObject)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

