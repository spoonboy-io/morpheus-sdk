# AccessToken
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **String** | Token that grants API Access | [optional] 
**RefreshToken** | **String** | Token that can request an new API access_token | [optional] 
**ExpiresIn** | **Decimal** | Epoch time when token expires | [optional] 
**TokenType** | **String** | Token type granted | [optional] 
**Scope** | **String** | Scope granted | [optional] 

## Examples

- Prepare the resource
```powershell
$AccessToken = Initialize-PSOpenAPIToolsAccessToken  -AccessToken null `
 -RefreshToken null `
 -ExpiresIn null `
 -TokenType null `
 -Scope null
```

- Convert the resource to JSON
```powershell
$AccessToken | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

