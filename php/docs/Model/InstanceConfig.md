# # InstanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**create_user** | **bool** |  | [optional]
**is_ec2** | **bool** |  | [optional]
**is_vpc_selectable** | **bool** |  | [optional]
**no_agent** | **bool** |  | [optional]
**security_groups** | [**\OpenAPI\Client\Model\InstanceConfigSecurityGroups[]**](InstanceConfigSecurityGroups.md) |  | [optional]
**smbios_asset_tag** | **string** |  | [optional]
**nested_virtualization** | **string** |  | [optional]
**vmware_folder_id** | **string** |  | [optional]
**custom_options** | **object** |  | [optional]
**resource_pool_id** | **int** |  | [optional]
**pool_provider_type** | **string** |  | [optional]
**user_group** | [**\OpenAPI\Client\Model\InstanceConfigSecurityGroups**](InstanceConfigSecurityGroups.md) |  | [optional]
**expire_days** | **string** |  | [optional]
**shutdown_days** | **string** |  | [optional]
**name** | **string** |  | [optional]
**host_name** | **string** |  | [optional]
**instance_type** | [**\OpenAPI\Client\Model\InstanceConfigInstanceType**](InstanceConfigInstanceType.md) |  | [optional]
**site** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional]
**environment_prefix** | **string** |  | [optional]
**layout** | [**\OpenAPI\Client\Model\InstanceConfigLayout**](InstanceConfigLayout.md) |  | [optional]
**type** | **string** |  | [optional]
**instance_context** | **string** |  | [optional]
**memory_display** | **string** |  | [optional]
**expose** | **object[]** |  | [optional]
**create_backup** | **bool** |  | [optional]
**backup** | [**\OpenAPI\Client\Model\InstanceConfigBackup**](InstanceConfigBackup.md) |  | [optional]
**replication_group** | [**\OpenAPI\Client\Model\InstanceConfigReplicationGroup**](InstanceConfigReplicationGroup.md) |  | [optional]
**layout_size** | **int** |  | [optional]
**lb_instances** | **object[]** |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
