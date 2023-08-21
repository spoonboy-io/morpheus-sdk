# # ApiZonesZoneIdResourcePoolsResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name of Resource Pool |
**default_pool** | **bool** | Set as the Default Pool | [optional] [default to false]
**default_image** | **bool** | Set as the Default Image Target | [optional] [default to false]
**active** | **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] [default to true]
**visibility** | **string** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to 'private']
**display_name** | **string** | Optional Display Name (VMware only) | [optional]
**inventory** | **bool** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] [default to true]
**config** | [**AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig**](AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig.md) |  |
**tenant_permissions** | [**\OpenAPI\Client\Model\ApiZonesZoneIdFoldersIdFolderTenantPermissions[]**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional]
**resource_permissions** | [**\OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
