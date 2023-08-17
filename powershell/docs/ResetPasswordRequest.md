# ResetPasswordRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **String** | The secret Reset Password token that was included in the **Forgot Password Email**. | 
**Password** | **String** | User new password. This is the new password for your user. | 

## Examples

- Prepare the resource
```powershell
$ResetPasswordRequest = Initialize-PSOpenAPIToolsResetPasswordRequest  -Token null `
 -Password null
```

- Convert the resource to JSON
```powershell
$ResetPasswordRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

