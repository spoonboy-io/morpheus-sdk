# ListArchiveBuckets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**ArchiveBuckets** | [**ArchiveBucket[]**](ArchiveBucket.md) |  | [optional] 
**ArchiveBucketCount** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListArchiveBuckets200Response = Initialize-PSOpenAPIToolsListArchiveBuckets200Response  -Meta null `
 -ArchiveBuckets null `
 -ArchiveBucketCount null
```

- Convert the resource to JSON
```powershell
$ListArchiveBuckets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

