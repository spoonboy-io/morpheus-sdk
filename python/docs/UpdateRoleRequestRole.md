# UpdateRoleRequestRole

Payload for updating an existing role

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authority** | **str** | Authority (Name) | [optional] 
**description** | **str, none_type** | Description | [optional] 
**default_persona** | **str, none_type** | Set the default persona by code. | [optional] 
**permissions** | [**[AddRolesRequestRolePermissionsInner]**](AddRolesRequestRolePermissionsInner.md) | Set the access level for the specified permissions. | [optional] 
**global_site_access** | **str** | Set the default access level for for groups (sites). Only applies to user roles. | [optional] 
**sites** | [**[AddRolesRequestRoleSitesInner]**](AddRolesRequestRoleSitesInner.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional] 
**global_zone_access** | **str** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**zones** | [**[AddRolesRequestRoleZonesInner]**](AddRolesRequestRoleZonesInner.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**global_instance_type_access** | **str** | Set the default access level for for instance types | [optional] 
**instance_types** | [**[AddRolesRequestRoleInstanceTypesInner]**](AddRolesRequestRoleInstanceTypesInner.md) | Set the access level for the specified instance types | [optional] 
**global_app_template_access** | **str** | Set the default access level for blueprints | [optional] 
**app_templates** | [**[AddRolesRequestRoleAppTemplatesInner]**](AddRolesRequestRoleAppTemplatesInner.md) | Set the access level for the specified blueprints (appTemplates) | [optional] 
**global_catalog_item_type_access** | **str** | Set the default access level for catalog item types | [optional] 
**catalog_item_types** | [**[AddRolesRequestRoleCatalogItemTypesInner]**](AddRolesRequestRoleCatalogItemTypesInner.md) | Set the access level for the specified catalog item types | [optional] 
**global_persona_access** | **str** | Set the default access level for personas | [optional] 
**personas** | [**[AddRolesRequestRolePersonasInner]**](AddRolesRequestRolePersonasInner.md) | Set the access level for the specified personas | [optional] 
**global_vdi_pool_access** | **str** | Set the default access level for VDI pools | [optional] 
**vdi_pools** | [**[AddRolesRequestRoleVdiPoolsInner]**](AddRolesRequestRoleVdiPoolsInner.md) | Set the access level for the specified VDI pools | [optional] 
**global_report_type_access** | **str** | Set the default access level for report types | [optional] 
**report_types** | [**[AddRolesRequestRoleReportTypesInner]**](AddRolesRequestRoleReportTypesInner.md) | Set the access level for the specified report types | [optional] 
**global_task_access** | **str** | Set the default access level for tasks | [optional] 
**tasks** | [**[AddRolesRequestRoleTasksInner]**](AddRolesRequestRoleTasksInner.md) | Set the access level for the specified tasks | [optional] 
**global_task_set_access** | **str** | Set the default access level for workflows (taskSets) | [optional] 
**task_sets** | [**[AddRolesRequestRoleTaskSetsInner]**](AddRolesRequestRoleTaskSetsInner.md) | Set the access level for the specified workflows (taskSets) | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


