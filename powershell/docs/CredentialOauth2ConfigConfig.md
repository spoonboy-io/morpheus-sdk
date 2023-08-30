# CredentialOauth2ConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **String** | OAuth 2.0 grant type | 
**AccessTokenUrl** | **String** | Token endpoint | 
**ClientId** | **String** | Client ID | [optional] 
**ClientSecret** | **String** | Client Secret | [optional] 
**Scope** | **String** | Scope | [optional] 
**ClientAuth** | **String** | Client Authentication Method | 

## Examples

- Prepare the resource
```powershell
$CredentialOauth2ConfigConfig = Initialize-PSOpenAPIToolsCredentialOauth2ConfigConfig  -GrantType null `
 -AccessTokenUrl https://example.com `
 -ClientId myclientid12345 `
 -ClientSecret myclientsecret12345 `
 -Scope read `
 -ClientAuth null
```

- Convert the resource to JSON
```powershell
$CredentialOauth2ConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

