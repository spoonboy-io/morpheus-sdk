# ApiClientsClient
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **String** | ClientId | 
**ClientSecret** | **String** | ClientSecret | [optional] 
**AccessTokenValiditySeconds** | **Int32** | Length of time accessToken is valid in seconds. | 
**RefreshTokenValiditySeconds** | **Int32** | Length of time refreshToken is valid in seconds. | 

## Examples

- Prepare the resource
```powershell
$ApiClientsClient = Initialize-PSOpenAPIToolsApiClientsClient  -ClientId Test Client `
 -ClientSecret thisIsaClientSecretKeyPhrase `
 -AccessTokenValiditySeconds 43200 `
 -RefreshTokenValiditySeconds 43200
```

- Convert the resource to JSON
```powershell
$ApiClientsClient | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

