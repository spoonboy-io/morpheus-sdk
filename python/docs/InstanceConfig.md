# InstanceConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**create_user** | **bool** |  | [optional] 
**is_ec2** | **bool** |  | [optional] 
**is_vpc_selectable** | **bool** |  | [optional] 
**no_agent** | **bool** |  | [optional] 
**security_groups** | [**[InstanceConfigSecurityGroupsInner], none_type**](InstanceConfigSecurityGroupsInner.md) |  | [optional] 
**smbios_asset_tag** | **str, none_type** |  | [optional] 
**nested_virtualization** | **str, none_type** |  | [optional] 
**vmware_folder_id** | **str** |  | [optional] 
**custom_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**resource_pool_id** | **int** |  | [optional] 
**pool_provider_type** | **str, none_type** |  | [optional] 
**user_group** | [**InstanceConfigSecurityGroupsInner**](InstanceConfigSecurityGroupsInner.md) |  | [optional] 
**expire_days** | **str** |  | [optional] 
**shutdown_days** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**host_name** | **str** |  | [optional] 
**instance_type** | [**InstanceConfigInstanceType**](InstanceConfigInstanceType.md) |  | [optional] 
**site** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**environment_prefix** | **str, none_type** |  | [optional] 
**layout** | [**InstanceConfigLayout**](InstanceConfigLayout.md) |  | [optional] 
**type** | **str** |  | [optional] 
**instance_context** | **str** |  | [optional] 
**memory_display** | **str** |  | [optional] 
**expose** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**create_backup** | **bool** |  | [optional] 
**backup** | [**InstanceConfigBackup**](InstanceConfigBackup.md) |  | [optional] 
**replication_group** | [**InstanceConfigReplicationGroup**](InstanceConfigReplicationGroup.md) |  | [optional] 
**layout_size** | **int** |  | [optional] 
**lb_instances** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


