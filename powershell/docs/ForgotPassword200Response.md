# ForgotPassword200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ForgotPassword200Response = Initialize-PSOpenAPIToolsForgotPassword200Response  -Success true `
 -Msg Reset password instructions have been sent to the user &#39;jsmith&#39;, if they exist.
```

- Convert the resource to JSON
```powershell
$ForgotPassword200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

