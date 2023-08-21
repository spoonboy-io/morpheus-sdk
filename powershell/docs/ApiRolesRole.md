# ApiRolesRole
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **String** | Authority (Name) | 
**Description** | **String** | Description | [optional] 
**RoleType** | **String** | Role type | [optional] [default to "user"]
**BaseRoleId** | **Int64** | Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions. | [optional] 
**Multitenant** | **Boolean** | Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant* | [optional] [default to $false]
**MultitenantLocked** | **Boolean** | Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant* | [optional] [default to $false]
**DefaultPersona** | **String** |  | [optional] 
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
**Owner** | **Int64** | Set the role owner (tenant) by ID. *Only available to master tenant* | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiRolesRole = Initialize-PSOpenAPIToolsApiRolesRole  -Authority null `
 -Description null `
 -RoleType null `
 -BaseRoleId null `
 -Multitenant null `
 -MultitenantLocked null `
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
 -TaskSets null `
 -Owner null
```

- Convert the resource to JSON
```powershell
$ApiRolesRole | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

