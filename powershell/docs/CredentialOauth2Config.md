# CredentialOauth2Config
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Credential Type Code | 
**Name** | **String** | A unique name scoped to your account for the credential | 
**Description** | **String** | Optional Description | [optional] 
**Enabled** | **Boolean** | Credential enabled | [optional] [default to $true]
**Integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**Username** | **String** | Username | [optional] 
**Password** | **String** | Password | [optional] 
**Config** | [**CredentialOauth2ConfigConfig**](CredentialOauth2ConfigConfig.md) |  | 

## Examples

- Prepare the resource
```powershell
$CredentialOauth2Config = Initialize-PSOpenAPIToolsCredentialOauth2Config  -Type null `
 -Name null `
 -Description null `
 -Enabled null `
 -Integration null `
 -Username jsmith `
 -Password myPassword `
 -Config null
```

- Convert the resource to JSON
```powershell
$CredentialOauth2Config | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

