# AddArchiveFile200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**ArchiveFile** | [**ArchiveBucketFile**](ArchiveBucketFile.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddArchiveFile200Response = Initialize-PSOpenAPIToolsAddArchiveFile200Response  -Success null `
 -ArchiveFile null
```

- Convert the resource to JSON
```powershell
$AddArchiveFile200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

