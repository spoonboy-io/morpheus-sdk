# # ApiRolesIdRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authority** | **string** | Authority (Name) | [optional]
**description** | **string** | Description | [optional]
**default_persona** | **string** | Set the default persona by code. | [optional]
**permissions** | [**\OpenAPI\Client\Model\ApiRolesRolePermissions[]**](ApiRolesRolePermissions.md) | Set the access level for the specified permissions. | [optional]
**global_site_access** | **string** | Set the default access level for for groups (sites). Only applies to user roles. | [optional]
**sites** | [**\OpenAPI\Client\Model\ApiRolesRoleSites[]**](ApiRolesRoleSites.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional]
**global_zone_access** | **string** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional]
**zones** | [**\OpenAPI\Client\Model\ApiRolesRoleZones[]**](ApiRolesRoleZones.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional]
**global_instance_type_access** | **string** | Set the default access level for for instance types | [optional]
**instance_types** | [**\OpenAPI\Client\Model\ApiRolesRoleInstanceTypes[]**](ApiRolesRoleInstanceTypes.md) | Set the access level for the specified instance types | [optional]
**global_app_template_access** | **string** | Set the default access level for blueprints | [optional]
**app_templates** | [**\OpenAPI\Client\Model\ApiRolesRoleAppTemplates[]**](ApiRolesRoleAppTemplates.md) | Set the access level for the specified blueprints (appTemplates) | [optional]
**global_catalog_item_type_access** | **string** | Set the default access level for catalog item types | [optional]
**catalog_item_types** | [**\OpenAPI\Client\Model\ApiRolesRoleCatalogItemTypes[]**](ApiRolesRoleCatalogItemTypes.md) | Set the access level for the specified catalog item types | [optional]
**global_persona_access** | **string** | Set the default access level for personas | [optional]
**personas** | [**\OpenAPI\Client\Model\ApiRolesRolePersonas[]**](ApiRolesRolePersonas.md) | Set the access level for the specified personas | [optional]
**global_vdi_pool_access** | **string** | Set the default access level for VDI pools | [optional]
**vdi_pools** | [**\OpenAPI\Client\Model\ApiRolesRoleVdiPools[]**](ApiRolesRoleVdiPools.md) | Set the access level for the specified VDI pools | [optional]
**global_report_type_access** | **string** | Set the default access level for report types | [optional]
**report_types** | [**\OpenAPI\Client\Model\ApiRolesRoleReportTypes[]**](ApiRolesRoleReportTypes.md) | Set the access level for the specified report types | [optional]
**global_task_access** | **string** | Set the default access level for tasks | [optional]
**tasks** | [**\OpenAPI\Client\Model\ApiRolesRoleTasks[]**](ApiRolesRoleTasks.md) | Set the access level for the specified tasks | [optional]
**global_task_set_access** | **string** | Set the default access level for workflows (taskSets) | [optional]
**task_sets** | [**\OpenAPI\Client\Model\ApiRolesRoleTaskSets[]**](ApiRolesRoleTaskSets.md) | Set the access level for the specified workflows (taskSets) | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
