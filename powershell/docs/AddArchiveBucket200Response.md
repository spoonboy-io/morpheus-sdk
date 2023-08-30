# AddArchiveBucket200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveBucket** | [**ArchiveBucket**](ArchiveBucket.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddArchiveBucket200Response = Initialize-PSOpenAPIToolsAddArchiveBucket200Response  -ArchiveBucket null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddArchiveBucket200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

