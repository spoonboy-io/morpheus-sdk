# ArchiveBucketFile
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**FilePath** | **String** |  | [optional] 
**ArchiveBucket** | [**ArchiveBucketFileArchiveBucket**](ArchiveBucketFileArchiveBucket.md) |  | [optional] 
**CreatedBy** | [**ArchiveBucketFileCreatedBy**](ArchiveBucketFileCreatedBy.md) |  | [optional] 
**IsDirectory** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**RawSize** | **Int64** |  | [optional] 
**ContentType** | **String** |  | [optional] 
**DownloadCount** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ArchiveBucketFile = Initialize-PSOpenAPIToolsArchiveBucketFile  -Id null `
 -Name null `
 -FilePath null `
 -ArchiveBucket null `
 -CreatedBy null `
 -IsDirectory null `
 -Status null `
 -RawSize null `
 -ContentType null `
 -DownloadCount null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$ArchiveBucketFile | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

