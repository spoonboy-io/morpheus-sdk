# IdentitySourcesOneLoginConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**Deleted** | **Boolean** |  | [optional] 
**AutoSyncOnLogin** | **Boolean** |  | [optional] 
**ExternalLogin** | **Boolean** |  | [optional] 
**AllowCustomMappings** | **Boolean** |  | [optional] 
**ManualRoleAssignment** | **Boolean** |  | [optional] 
**Account** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**DefaultAccountRole** | [**IdentitySourcesLDAPConfigDefaultAccountRole**](IdentitySourcesLDAPConfigDefaultAccountRole.md) |  | [optional] 
**Config** | [**IdentitySourcesOneLoginConfigConfig**](IdentitySourcesOneLoginConfigConfig.md) |  | [optional] 
**RoleMappings** | [**IdentitySourcesJumpCloudConfigRoleMappings[]**](IdentitySourcesJumpCloudConfigRoleMappings.md) |  | [optional] 
**Subdomain** | **String** |  | [optional] 
**LoginURL** | **String** |  | [optional] 
**ProviderSettings** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesOneLoginConfig = Initialize-PSOpenAPIToolsIdentitySourcesOneLoginConfig  -Id null `
 -Name null `
 -Description null `
 -Code null `
 -Type null `
 -Active null `
 -Deleted null `
 -AutoSyncOnLogin null `
 -ExternalLogin null `
 -AllowCustomMappings null `
 -ManualRoleAssignment null `
 -Account null `
 -DefaultAccountRole null `
 -Config null `
 -RoleMappings null `
 -Subdomain null `
 -LoginURL null `
 -ProviderSettings null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesOneLoginConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

