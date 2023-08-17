# ResetPassword200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ResetPassword200Response = Initialize-PSOpenAPIToolsResetPassword200Response  -Success true `
 -Msg Password has been updated and account unlocked
```

- Convert the resource to JSON
```powershell
$ResetPassword200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

