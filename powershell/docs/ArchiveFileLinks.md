# ArchiveFileLinks
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**SecretAccessKey** | **String** |  | [optional] 
**ArchiveFile** | [**ArchiveFileLinksArchiveFile**](ArchiveFileLinksArchiveFile.md) |  | [optional] 
**CreatedBy** | [**ArchiveBucketFileCreatedBy**](ArchiveBucketFileCreatedBy.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**LastAccessDate** | **System.DateTime** |  | [optional] 
**ExpirationDate** | **System.DateTime** |  | [optional] 
**DownloadCount** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ArchiveFileLinks = Initialize-PSOpenAPIToolsArchiveFileLinks  -Id null `
 -SecretAccessKey null `
 -ArchiveFile null `
 -CreatedBy null `
 -DateCreated null `
 -LastUpdated null `
 -LastAccessDate null `
 -ExpirationDate null `
 -DownloadCount null
```

- Convert the resource to JSON
```powershell
$ArchiveFileLinks | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

