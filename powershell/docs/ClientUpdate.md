# ClientUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **String** |  | [optional] 
**AccessTokenValiditySeconds** | **Int64** |  | [optional] 
**RefreshTokenValiditySeconds** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClientUpdate = Initialize-PSOpenAPIToolsClientUpdate  -ClientId null `
 -AccessTokenValiditySeconds null `
 -RefreshTokenValiditySeconds null
```

- Convert the resource to JSON
```powershell
$ClientUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

