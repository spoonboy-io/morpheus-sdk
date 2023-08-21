

# ApiRolesIdRole

Payload for updating an existing role
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authority** | **String** | Authority (Name) |  [optional]
**description** | **String** | Description |  [optional]
**defaultPersona** | [**DefaultPersonaEnum**](#DefaultPersonaEnum) | Set the default persona by code. |  [optional]
**permissions** | [**List&lt;ApiRolesRolePermissions&gt;**](ApiRolesRolePermissions.md) | Set the access level for the specified permissions. |  [optional]
**globalSiteAccess** | [**GlobalSiteAccessEnum**](#GlobalSiteAccessEnum) | Set the default access level for for groups (sites). Only applies to user roles. |  [optional]
**sites** | [**List&lt;ApiRolesRoleSites&gt;**](ApiRolesRoleSites.md) | Set the access level for the specified groups (sites). Only applies to user roles. |  [optional]
**globalZoneAccess** | [**GlobalZoneAccessEnum**](#GlobalZoneAccessEnum) | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. |  [optional]
**zones** | [**List&lt;ApiRolesRoleZones&gt;**](ApiRolesRoleZones.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. |  [optional]
**globalInstanceTypeAccess** | [**GlobalInstanceTypeAccessEnum**](#GlobalInstanceTypeAccessEnum) | Set the default access level for for instance types |  [optional]
**instanceTypes** | [**List&lt;ApiRolesRoleInstanceTypes&gt;**](ApiRolesRoleInstanceTypes.md) | Set the access level for the specified instance types |  [optional]
**globalAppTemplateAccess** | [**GlobalAppTemplateAccessEnum**](#GlobalAppTemplateAccessEnum) | Set the default access level for blueprints |  [optional]
**appTemplates** | [**List&lt;ApiRolesRoleAppTemplates&gt;**](ApiRolesRoleAppTemplates.md) | Set the access level for the specified blueprints (appTemplates) |  [optional]
**globalCatalogItemTypeAccess** | [**GlobalCatalogItemTypeAccessEnum**](#GlobalCatalogItemTypeAccessEnum) | Set the default access level for catalog item types |  [optional]
**catalogItemTypes** | [**List&lt;ApiRolesRoleCatalogItemTypes&gt;**](ApiRolesRoleCatalogItemTypes.md) | Set the access level for the specified catalog item types |  [optional]
**globalPersonaAccess** | [**GlobalPersonaAccessEnum**](#GlobalPersonaAccessEnum) | Set the default access level for personas |  [optional]
**personas** | [**List&lt;ApiRolesRolePersonas&gt;**](ApiRolesRolePersonas.md) | Set the access level for the specified personas |  [optional]
**globalVdiPoolAccess** | [**GlobalVdiPoolAccessEnum**](#GlobalVdiPoolAccessEnum) | Set the default access level for VDI pools |  [optional]
**vdiPools** | [**List&lt;ApiRolesRoleVdiPools&gt;**](ApiRolesRoleVdiPools.md) | Set the access level for the specified VDI pools |  [optional]
**globalReportTypeAccess** | [**GlobalReportTypeAccessEnum**](#GlobalReportTypeAccessEnum) | Set the default access level for report types |  [optional]
**reportTypes** | [**List&lt;ApiRolesRoleReportTypes&gt;**](ApiRolesRoleReportTypes.md) | Set the access level for the specified report types |  [optional]
**globalTaskAccess** | [**GlobalTaskAccessEnum**](#GlobalTaskAccessEnum) | Set the default access level for tasks |  [optional]
**tasks** | [**List&lt;ApiRolesRoleTasks&gt;**](ApiRolesRoleTasks.md) | Set the access level for the specified tasks |  [optional]
**globalTaskSetAccess** | [**GlobalTaskSetAccessEnum**](#GlobalTaskSetAccessEnum) | Set the default access level for workflows (taskSets) |  [optional]
**taskSets** | [**List&lt;ApiRolesRoleTaskSets&gt;**](ApiRolesRoleTaskSets.md) | Set the access level for the specified workflows (taskSets) |  [optional]



## Enum: DefaultPersonaEnum

Name | Value
---- | -----
STANDARD | &quot;standard&quot;
SERVICECATALOG | &quot;serviceCatalog&quot;
VDI | &quot;vdi&quot;



## Enum: GlobalSiteAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
READ | &quot;read&quot;
NONE | &quot;none&quot;



## Enum: GlobalZoneAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
READ | &quot;read&quot;
NONE | &quot;none&quot;



## Enum: GlobalInstanceTypeAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalAppTemplateAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalCatalogItemTypeAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalPersonaAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalVdiPoolAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalReportTypeAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalTaskAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



## Enum: GlobalTaskSetAccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
NONE | &quot;none&quot;



