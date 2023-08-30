# CredentialAccessSecretKeyConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Credential Type Code | 
**Name** | **String** | A unique name scoped to your account for the credential | 
**Description** | **String** | Optional Description | [optional] 
**Enabled** | **Boolean** | Credential enabled | [optional] [default to $true]
**Integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**Username** | **String** | Access Key | 
**Password** | **String** | Secret Key | 

## Examples

- Prepare the resource
```powershell
$CredentialAccessSecretKeyConfig = Initialize-PSOpenAPIToolsCredentialAccessSecretKeyConfig  -Type null `
 -Name null `
 -Description null `
 -Enabled null `
 -Integration null `
 -Username 72c54d9b-1c73-4c73-91b9-1fb895f5fe5a `
 -Password 2b3450f3-b563-4a5f-ba96-91af0212fd69
```

- Convert the resource to JSON
```powershell
$CredentialAccessSecretKeyConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

