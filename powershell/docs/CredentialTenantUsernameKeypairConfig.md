# CredentialTenantUsernameKeypairConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Credential Type Code | 
**Name** | **String** | A unique name scoped to your account for the credential | 
**Description** | **String** | Optional Description | [optional] 
**Enabled** | **Boolean** | Credential enabled | [optional] [default to $true]
**Integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**AuthPath** | **String** | Tenant name | 
**Username** | **String** | Username | 
**AuthKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](CredentialEmailPrivateKeyConfigAuthKey.md) |  | 

## Examples

- Prepare the resource
```powershell
$CredentialTenantUsernameKeypairConfig = Initialize-PSOpenAPIToolsCredentialTenantUsernameKeypairConfig  -Type null `
 -Name null `
 -Description null `
 -Enabled null `
 -Integration null `
 -AuthPath My Tenant `
 -Username jsmith `
 -AuthKey null
```

- Convert the resource to JSON
```powershell
$CredentialTenantUsernameKeypairConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

