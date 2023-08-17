

# ArchiveBucketCreate


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | A name for the archive bucket. Must be globally unique. |  [optional] |
|**description** | **String** | A description for the archive bucket |  [optional] |
|**storageProvider** | [**ArchiveBucketCreateStorageProvider**](ArchiveBucketCreateStorageProvider.md) |  |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) | Visibility - Set to public to allow all tenants |  [optional] |
|**isPublic** | **Boolean** | Public URL - Set to true to allow anonymous access |  [optional] |
|**accounts** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  |  [optional] |



## Enum: VisibilityEnum

| Name | Value |
|---- | -----|
| PUBLIC | &quot;public&quot; |
| PRIVATE | &quot;private&quot; |



