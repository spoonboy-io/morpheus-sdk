# ListArchiveFiles200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**ArchiveBucket** | [**ArchiveBucket**](ArchiveBucket.md) |  | [optional] 
**IsOwner** | **Boolean** |  | [optional] 
**ParentDirectory** | **String** |  | [optional] 
**ArchiveFiles** | [**ArchiveBucketFile[]**](ArchiveBucketFile.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListArchiveFiles200Response = Initialize-PSOpenAPIToolsListArchiveFiles200Response  -Meta null `
 -ArchiveBucket null `
 -IsOwner null `
 -ParentDirectory null `
 -ArchiveFiles null
```

- Convert the resource to JSON
```powershell
$ListArchiveFiles200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

