# CredentialUsernameKeypairConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Credential Type Code | 
**Name** | **String** | A unique name scoped to your account for the credential | 
**Description** | **String** | Optional Description | [optional] 
**Enabled** | **Boolean** | Credential enabled | [optional] [default to $true]
**Integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**Username** | **String** | Username | 
**AuthKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](CredentialEmailPrivateKeyConfigAuthKey.md) |  | 

## Examples

- Prepare the resource
```powershell
$CredentialUsernameKeypairConfig = Initialize-PSOpenAPIToolsCredentialUsernameKeypairConfig  -Type null `
 -Name null `
 -Description null `
 -Enabled null `
 -Integration null `
 -Username your_user `
 -AuthKey null
```

- Convert the resource to JSON
```powershell
$CredentialUsernameKeypairConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

