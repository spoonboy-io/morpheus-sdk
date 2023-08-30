# IdentitySourcesADConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **String** |  | [optional] 
**Domain** | **String** |  | [optional] 
**UseSSL** | **String** |  | [optional] 
**BindingUsername** | **String** |  | [optional] 
**BindingPassword** | **String** |  | [optional] 
**RequiredGroup** | **String** |  | [optional] 
**RequiredGroupDN** | **String** |  | [optional] 
**SearchMemberGroups** | **Boolean** |  | [optional] 
**BindingPasswordHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesADConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesADConfigConfig  -Url null `
 -Domain null `
 -UseSSL null `
 -BindingUsername null `
 -BindingPassword null `
 -RequiredGroup null `
 -RequiredGroupDN null `
 -SearchMemberGroups null `
 -BindingPasswordHash null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesADConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

