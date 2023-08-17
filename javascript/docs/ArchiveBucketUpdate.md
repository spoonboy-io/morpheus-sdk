# MorpheusApi.ArchiveBucketUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the archive bucket. Must be globally unique. | [optional] 
**description** | **String** | A description for the archive bucket | [optional] 
**visibility** | **String** | Visibility - Set to public to allow all tenants | [optional] [default to &#39;private&#39;]
**isPublic** | **Boolean** | Public URL - Set to true to allow anonymous access | [optional] [default to false]
**accounts** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 



## Enum: VisibilityEnum


* `public` (value: `"public"`)

* `private` (value: `"private"`)




