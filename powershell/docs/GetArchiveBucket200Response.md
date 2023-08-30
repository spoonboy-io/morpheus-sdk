# GetArchiveBucket200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveBucket** | [**ArchiveBucket**](ArchiveBucket.md) |  | [optional] 
**IsOwner** | **Boolean** |  | [optional] 
**ParentDirectory** | **String** |  | [optional] 
**ArchiveFiles** | [**ArchiveBucketFile[]**](ArchiveBucketFile.md) |  | [optional] 
**ArchiveFileCount** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetArchiveBucket200Response = Initialize-PSOpenAPIToolsGetArchiveBucket200Response  -ArchiveBucket null `
 -IsOwner null `
 -ParentDirectory null `
 -ArchiveFiles null `
 -ArchiveFileCount null
```

- Convert the resource to JSON
```powershell
$GetArchiveBucket200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

