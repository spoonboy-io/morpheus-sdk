# # ApiSecurityGroupsSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name for your security group |
**description** | **string** | Optional description field | [optional]
**zone_id** | **int** | Scoped Cloud ID |
**active** | **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional]
**custom_options** | [**AnyOfSecurityGroupLocationAzureCustomOptionsSecurityGroupLocationAwsCustomOptionsSecurityGroupLocationOpenstackCustomOptions**](AnyOfSecurityGroupLocationAzureCustomOptionsSecurityGroupLocationAwsCustomOptionsSecurityGroupLocationOpenstackCustomOptions.md) |  | [optional]
**tenant_permissions** | [**\OpenAPI\Client\Model\ApiSecurityGroupsSecurityGroupTenantPermissions[]**](ApiSecurityGroupsSecurityGroupTenantPermissions.md) |  | [optional]
**resource_permissions** | [**\OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
