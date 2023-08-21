# CredentialUsernameAPIKeyConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Credential Type Code | 
**Name** | **String** | A unique name scoped to your account for the credential | 
**Description** | **String** | Optional Description | [optional] 
**Enabled** | **Boolean** | Credential enabled | [optional] [default to $true]
**Integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**Username** | **String** | Username | 
**Password** | **String** | API Key | 

## Examples

- Prepare the resource
```powershell
$CredentialUsernameAPIKeyConfig = Initialize-PSOpenAPIToolsCredentialUsernameAPIKeyConfig  -Type null `
 -Name null `
 -Description null `
 -Enabled null `
 -Integration null `
 -Username your_user `
 -Password your_api_key
```

- Convert the resource to JSON
```powershell
$CredentialUsernameAPIKeyConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

