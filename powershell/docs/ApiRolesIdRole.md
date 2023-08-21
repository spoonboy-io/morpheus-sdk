# ApiRolesIdRole
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **String** | Authority (Name) | [optional] 
**Description** | **String** | Description | [optional] 
**DefaultPersona** | **String** | Set the default persona by code. | [optional] 
**Permissions** | [**ApiRolesRolePermissions[]**](ApiRolesRolePermissions.md) | Set the access level for the specified permissions. | [optional] 
**GlobalSiteAccess** | **String** | Set the default access level for for groups (sites). Only applies to user roles. | [optional] 
**Sites** | [**ApiRolesRoleSites[]**](ApiRolesRoleSites.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional] 
**GlobalZoneAccess** | **String** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**Zones** | [**ApiRolesRoleZones[]**](ApiRolesRoleZones.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**GlobalInstanceTypeAccess** | **String** | Set the default access level for for instance types | [optional] 
**InstanceTypes** | [**ApiRolesRoleInstanceTypes[]**](ApiRolesRoleInstanceTypes.md) | Set the access level for the specified instance types | [optional] 
**GlobalAppTemplateAccess** | **String** | Set the default access level for blueprints | [optional] 
**AppTemplates** | [**ApiRolesRoleAppTemplates[]**](ApiRolesRoleAppTemplates.md) | Set the access level for the specified blueprints (appTemplates) | [optional] 
**GlobalCatalogItemTypeAccess** | **String** | Set the default access level for catalog item types | [optional] 
**CatalogItemTypes** | [**ApiRolesRoleCatalogItemTypes[]**](ApiRolesRoleCatalogItemTypes.md) | Set the access level for the specified catalog item types | [optional] 
**GlobalPersonaAccess** | **String** | Set the default access level for personas | [optional] 
**Personas** | [**ApiRolesRolePersonas[]**](ApiRolesRolePersonas.md) | Set the access level for the specified personas | [optional] 
**GlobalVdiPoolAccess** | **String** | Set the default access level for VDI pools | [optional] 
**VdiPools** | [**ApiRolesRoleVdiPools[]**](ApiRolesRoleVdiPools.md) | Set the access level for the specified VDI pools | [optional] 
**GlobalReportTypeAccess** | **String** | Set the default access level for report types | [optional] 
**ReportTypes** | [**ApiRolesRoleReportTypes[]**](ApiRolesRoleReportTypes.md) | Set the access level for the specified report types | [optional] 
**GlobalTaskAccess** | **String** | Set the default access level for tasks | [optional] 
**Tasks** | [**ApiRolesRoleTasks[]**](ApiRolesRoleTasks.md) | Set the access level for the specified tasks | [optional] 
**GlobalTaskSetAccess** | **String** | Set the default access level for workflows (taskSets) | [optional] 
**TaskSets** | [**ApiRolesRoleTaskSets[]**](ApiRolesRoleTaskSets.md) | Set the access level for the specified workflows (taskSets) | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiRolesIdRole = Initialize-PSOpenAPIToolsApiRolesIdRole  -Authority null `
 -Description null `
 -DefaultPersona null `
 -Permissions null `
 -GlobalSiteAccess null `
 -Sites null `
 -GlobalZoneAccess null `
 -Zones null `
 -GlobalInstanceTypeAccess null `
 -InstanceTypes null `
 -GlobalAppTemplateAccess null `
 -AppTemplates null `
 -GlobalCatalogItemTypeAccess null `
 -CatalogItemTypes null `
 -GlobalPersonaAccess null `
 -Personas null `
 -GlobalVdiPoolAccess null `
 -VdiPools null `
 -GlobalReportTypeAccess null `
 -ReportTypes null `
 -GlobalTaskAccess null `
 -Tasks null `
 -GlobalTaskSetAccess null `
 -TaskSets null
```

- Convert the resource to JSON
```powershell
$ApiRolesIdRole | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

