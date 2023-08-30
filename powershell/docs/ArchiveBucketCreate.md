# ArchiveBucketCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the archive bucket. Must be globally unique. | [optional] 
**Description** | **String** | A description for the archive bucket | [optional] 
**StorageProvider** | [**ArchiveBucketCreateStorageProvider**](ArchiveBucketCreateStorageProvider.md) |  | [optional] 
**Visibility** | **String** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**IsPublic** | **Boolean** | Public URL - Set to true to allow anonymous access | [optional] [default to $false]
**Accounts** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ArchiveBucketCreate = Initialize-PSOpenAPIToolsArchiveBucketCreate  -Name null `
 -Description null `
 -StorageProvider null `
 -Visibility null `
 -IsPublic null `
 -Accounts null
```

- Convert the resource to JSON
```powershell
$ArchiveBucketCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

