# CredentialConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **String** |  | [optional] 
**ClientId** | **String** |  | [optional] 
**ClientAuth** | **String** |  | [optional] 
**Scope** | **String** |  | [optional] 
**GrantType** | **String** |  | [optional] 
**AccessTokenUrl** | **String** |  | [optional] 
**ClientSecretHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CredentialConfig = Initialize-PSOpenAPIToolsCredentialConfig  -ClientSecret null `
 -ClientId null `
 -ClientAuth null `
 -Scope null `
 -GrantType null `
 -AccessTokenUrl null `
 -ClientSecretHash null
```

- Convert the resource to JSON
```powershell
$CredentialConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

