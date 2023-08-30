# ClusterDatastores


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**code** | **str, none_type** |  | [optional] 
**type** | **str** |  | [optional] 
**visibility** | **str** |  | [optional] 
**storage_size** | **int, none_type** |  | [optional] 
**free_space** | **int, none_type** |  | [optional] 
**drs_enabled** | **bool** |  | [optional] 
**active** | **bool** |  | [optional] 
**allow_write** | **bool** |  | [optional] 
**default_store** | **bool** |  | [optional] 
**online** | **bool** |  | [optional] 
**allow_read** | **bool** |  | [optional] 
**allow_provision** | **bool** |  | [optional] 
**ref_type** | **str** |  | [optional] 
**ref_id** | **int** |  | [optional] 
**external_id** | **str** |  | [optional] 
**zone** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**zone_pool** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**owner** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**tenants** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**datastores** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**permissions** | [**ClusterDatastoresPermissions**](ClusterDatastoresPermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


