# CredentialUser
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Username** | **String** |  | [optional] 
**DisplayName** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CredentialUser = Initialize-PSOpenAPIToolsCredentialUser  -Id null `
 -Username null `
 -DisplayName null
```

- Convert the resource to JSON
```powershell
$CredentialUser | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

