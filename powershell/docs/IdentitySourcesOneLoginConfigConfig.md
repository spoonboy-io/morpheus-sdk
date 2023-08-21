# IdentitySourcesOneLoginConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subdomain** | **String** |  | [optional] 
**Region** | **String** |  | [optional] 
**ClientSecret** | **String** |  | [optional] 
**ClientId** | **String** |  | [optional] 
**RequiredRole** | **String** |  | [optional] 
**RequiredRoleId** | **String** |  | [optional] 
**ClientSecretHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesOneLoginConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesOneLoginConfigConfig  -Subdomain null `
 -Region null `
 -ClientSecret null `
 -ClientId null `
 -RequiredRole null `
 -RequiredRoleId null `
 -ClientSecretHash null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesOneLoginConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

