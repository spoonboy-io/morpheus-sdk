# MorpheusApi.ApiRolesRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authority** | **String** | Authority (Name) | 
**description** | **String** | Description | [optional] 
**roleType** | **String** | Role type | [optional] [default to &#39;user&#39;]
**baseRoleId** | **Number** | Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions. | [optional] 
**multitenant** | **Boolean** | Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant* | [optional] [default to false]
**multitenantLocked** | **Boolean** | Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant* | [optional] [default to false]
**defaultPersona** | **String** |  | [optional] 
**permissions** | [**[ApiRolesRolePermissions]**](ApiRolesRolePermissions.md) | Set the access level for the specified permissions. | [optional] 
**globalSiteAccess** | **String** | Set the default access level for for groups (sites). Only applies to user roles. | [optional] 
**sites** | [**[ApiRolesRoleSites]**](ApiRolesRoleSites.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional] 
**globalZoneAccess** | **String** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**zones** | [**[ApiRolesRoleZones]**](ApiRolesRoleZones.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**globalInstanceTypeAccess** | **String** | Set the default access level for for instance types | [optional] 
**instanceTypes** | [**[ApiRolesRoleInstanceTypes]**](ApiRolesRoleInstanceTypes.md) | Set the access level for the specified instance types | [optional] 
**globalAppTemplateAccess** | **String** | Set the default access level for blueprints | [optional] 
**appTemplates** | [**[ApiRolesRoleAppTemplates]**](ApiRolesRoleAppTemplates.md) | Set the access level for the specified blueprints (appTemplates) | [optional] 
**globalCatalogItemTypeAccess** | **String** | Set the default access level for catalog item types | [optional] 
**catalogItemTypes** | [**[ApiRolesRoleCatalogItemTypes]**](ApiRolesRoleCatalogItemTypes.md) | Set the access level for the specified catalog item types | [optional] 
**globalPersonaAccess** | **String** | Set the default access level for personas | [optional] 
**personas** | [**[ApiRolesRolePersonas]**](ApiRolesRolePersonas.md) | Set the access level for the specified personas | [optional] 
**globalVdiPoolAccess** | **String** | Set the default access level for VDI pools | [optional] 
**vdiPools** | [**[ApiRolesRoleVdiPools]**](ApiRolesRoleVdiPools.md) | Set the access level for the specified VDI pools | [optional] 
**globalReportTypeAccess** | **String** | Set the default access level for report types | [optional] 
**reportTypes** | [**[ApiRolesRoleReportTypes]**](ApiRolesRoleReportTypes.md) | Set the access level for the specified report types | [optional] 
**globalTaskAccess** | **String** | Set the default access level for tasks | [optional] 
**tasks** | [**[ApiRolesRoleTasks]**](ApiRolesRoleTasks.md) | Set the access level for the specified tasks | [optional] 
**globalTaskSetAccess** | **String** | Set the default access level for workflows (taskSets) | [optional] 
**taskSets** | [**[ApiRolesRoleTaskSets]**](ApiRolesRoleTaskSets.md) | Set the access level for the specified workflows (taskSets) | [optional] 
**owner** | **Number** | Set the role owner (tenant) by ID. *Only available to master tenant* | [optional] 



## Enum: RoleTypeEnum


* `user` (value: `"user"`)

* `account` (value: `"account"`)





## Enum: DefaultPersonaEnum


* `standard` (value: `"standard"`)

* `serviceCatalog` (value: `"serviceCatalog"`)

* `vdi` (value: `"vdi"`)





## Enum: GlobalSiteAccessEnum


* `default` (value: `"default"`)

* `full` (value: `"full"`)

* `read` (value: `"read"`)

* `none` (value: `"none"`)





## Enum: GlobalZoneAccessEnum


* `default` (value: `"default"`)

* `full` (value: `"full"`)

* `read` (value: `"read"`)

* `none` (value: `"none"`)





## Enum: GlobalInstanceTypeAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalAppTemplateAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalCatalogItemTypeAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalPersonaAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalVdiPoolAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalReportTypeAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalTaskAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)





## Enum: GlobalTaskSetAccessEnum


* `full` (value: `"full"`)

* `none` (value: `"none"`)




