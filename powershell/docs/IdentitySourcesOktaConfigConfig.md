# IdentitySourcesOktaConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **String** |  | [optional] 
**AdministratorAPIToken** | **String** |  | [optional] 
**RequiredGroup** | **String** |  | [optional] 
**RequiredGroupId** | **String** |  | [optional] 
**AdministratorAPITokenHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesOktaConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesOktaConfigConfig  -Url null `
 -AdministratorAPIToken null `
 -RequiredGroup null `
 -RequiredGroupId null `
 -AdministratorAPITokenHash null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesOktaConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

