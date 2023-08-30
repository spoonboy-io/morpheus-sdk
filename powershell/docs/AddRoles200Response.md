# AddRoles200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**RoleRole**](RoleRole.md) |  | [optional] 
**FeaturePermissions** | [**RoleFeaturePermissionsInner[]**](RoleFeaturePermissionsInner.md) |  | [optional] 
**GlobalSiteAccess** | **String** |  | [optional] 
**Sites** | [**RoleSitesInner[]**](RoleSitesInner.md) |  | [optional] 
**GlobalZoneAccess** | **String** |  | [optional] 
**Zones** | [**RoleSitesInner[]**](RoleSitesInner.md) |  | [optional] 
**GlobalInstanceTypeAccess** | **String** |  | [optional] 
**InstanceTypePermissions** | [**RoleInstanceTypePermissionsInner[]**](RoleInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalAppTemplateAccess** | **String** |  | [optional] 
**AppTemplatePermissions** | [**RoleAppTemplatePermissionsInner[]**](RoleAppTemplatePermissionsInner.md) |  | [optional] 
**GlobalCatalogItemTypeAccess** | **String** |  | [optional] 
**CatalogItemTypePermissions** | [**RoleSitesInner[]**](RoleSitesInner.md) |  | [optional] 
**GlobalPersonaAccess** | **String** |  | [optional] 
**PersonaPermissions** | [**RoleInstanceTypePermissionsInner[]**](RoleInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalVdiPoolAccess** | **String** |  | [optional] 
**VdiPoolPermissions** | [**RoleSitesInner[]**](RoleSitesInner.md) |  | [optional] 
**GlobalReportTypeAccess** | **String** |  | [optional] 
**ReportTypePermissions** | [**RoleInstanceTypePermissionsInner[]**](RoleInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalTaskAccess** | **String** |  | [optional] 
**TaskPermissions** | [**RoleAppTemplatePermissionsInner[]**](RoleAppTemplatePermissionsInner.md) |  | [optional] 
**GlobalTaskSetAccess** | **String** |  | [optional] 
**TaskSetPermissions** | [**RoleAppTemplatePermissionsInner[]**](RoleAppTemplatePermissionsInner.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddRoles200Response = Initialize-PSOpenAPIToolsAddRoles200Response  -Role null `
 -FeaturePermissions null `
 -GlobalSiteAccess null `
 -Sites null `
 -GlobalZoneAccess null `
 -Zones null `
 -GlobalInstanceTypeAccess null `
 -InstanceTypePermissions null `
 -GlobalAppTemplateAccess null `
 -AppTemplatePermissions null `
 -GlobalCatalogItemTypeAccess null `
 -CatalogItemTypePermissions null `
 -GlobalPersonaAccess null `
 -PersonaPermissions null `
 -GlobalVdiPoolAccess null `
 -VdiPoolPermissions null `
 -GlobalReportTypeAccess null `
 -ReportTypePermissions null `
 -GlobalTaskAccess null `
 -TaskPermissions null `
 -GlobalTaskSetAccess null `
 -TaskSetPermissions null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddRoles200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

