# IdentitySourcesJumpCloudConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **String** |  | [optional] 
**BindingUsername** | **String** |  | [optional] 
**BindingPassword** | **String** |  | [optional] 
**RequiredRole** | **String** |  | [optional] 
**BindingPasswordHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesJumpCloudConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesJumpCloudConfigConfig  -OrganizationId null `
 -BindingUsername null `
 -BindingPassword null `
 -RequiredRole null `
 -BindingPasswordHash null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesJumpCloudConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

