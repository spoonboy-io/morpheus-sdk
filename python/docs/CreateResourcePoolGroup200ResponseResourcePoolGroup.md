# CreateResourcePoolGroup200ResponseResourcePoolGroup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**description** | **str** |  | [optional] 
**visibility** | **str** |  | [optional] 
**mode** | **str** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**pools** | **[int]** | Array of Resource Pool IDs | [optional] 
**tenants** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**resource_permission** | [**GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission**](GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


