# ArchiveBucket
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**StorageProvider** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Owner** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**CreatedBy** | [**ArchiveBucketCreatedBy**](ArchiveBucketCreatedBy.md) |  | [optional] 
**IsPublic** | **Boolean** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**FilePath** | **String** |  | [optional] 
**RawSize** | **Int64** |  | [optional] 
**FileCount** | **Int64** |  | [optional] 
**Accounts** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ArchiveBucket = Initialize-PSOpenAPIToolsArchiveBucket  -Id null `
 -Name null `
 -Description null `
 -StorageProvider null `
 -Owner null `
 -CreatedBy null `
 -IsPublic null `
 -Visibility null `
 -Code null `
 -FilePath null `
 -RawSize null `
 -FileCount null `
 -Accounts null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$ArchiveBucket | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

