# IdentitySourcesSAMLConfig
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
**Account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**DefaultAccountRole** | [**IdentitySourcesLDAPConfigDefaultAccountRole**](IdentitySourcesLDAPConfigDefaultAccountRole.md) |  | [optional] 
**Config** | [**IdentitySourcesSAMLConfigConfig**](IdentitySourcesSAMLConfigConfig.md) |  | [optional] 
**RoleMappings** | [**IdentitySourcesSAMLConfigRoleMappingsInner[]**](IdentitySourcesSAMLConfigRoleMappingsInner.md) |  | [optional] 
**Subdomain** | **String** |  | [optional] 
**LoginURL** | **String** |  | [optional] 
**ProviderSettings** | [**IdentitySourcesSAMLConfigProviderSettings**](IdentitySourcesSAMLConfigProviderSettings.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesSAMLConfig = Initialize-PSOpenAPIToolsIdentitySourcesSAMLConfig  -Id null `
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
$IdentitySourcesSAMLConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

