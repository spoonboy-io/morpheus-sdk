# Client
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**ClientId** | **String** |  | [optional] 
**AccessTokenValiditySeconds** | **Int64** |  | [optional] 
**RefreshTokenValiditySeconds** | **Int64** |  | [optional] 
**Authorities** | **String[]** |  | [optional] 
**AuthorizedGrantTypes** | **String[]** |  | [optional] 
**Scopes** | **String[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Client = Initialize-PSOpenAPIToolsClient  -Id null `
 -ClientId null `
 -AccessTokenValiditySeconds null `
 -RefreshTokenValiditySeconds null `
 -Authorities null `
 -AuthorizedGrantTypes null `
 -Scopes null
```

- Convert the resource to JSON
```powershell
$Client | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

