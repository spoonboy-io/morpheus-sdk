# ArchiveBucketCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the archive bucket. Must be globally unique. | [optional] 
**description** | **str** | A description for the archive bucket | [optional] 
**storage_provider** | [**ArchiveBucketCreateStorageProvider**](ArchiveBucketCreateStorageProvider.md) |  | [optional] 
**visibility** | **str** | Visibility - Set to public to allow all tenants | [optional]  if omitted the server will use the default value of "private"
**is_public** | **bool** | Public URL - Set to true to allow anonymous access | [optional]  if omitted the server will use the default value of False
**accounts** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


