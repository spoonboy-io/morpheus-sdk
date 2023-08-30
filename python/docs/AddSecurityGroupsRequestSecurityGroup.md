# AddSecurityGroupsRequestSecurityGroup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name for your security group | 
**zone_id** | **int** | Scoped Cloud ID | 
**description** | **str** | Optional description field | [optional] 
**active** | **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**custom_options** | [**AddSecurityGroupsRequestSecurityGroupCustomOptions**](AddSecurityGroupsRequestSecurityGroupCustomOptions.md) |  | [optional] 
**tenant_permissions** | [**[AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner]**](AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner.md) |  | [optional] 
**resource_permissions** | [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


