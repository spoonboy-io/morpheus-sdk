# ApiTasksTaskFile
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **String** | File Source i.e. &#x60;local&#x60;, &#x60;repository&#x60;, &#x60;url&#x60;. Default is &#x60;local&#x60;. | [default to "local"]
**Content** | **String** | File content, the script text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**ContentPath** | **String** | Content Path, the repo file location or url. Required when sourceType is &#x60;repository&#x60; or &#x60;url&#x60;. | [optional] 
**ContentRef** | **String** | Content Ref, the branch/tag. Only used when sourceType is &#x60;repository&#x60;. | [optional] 
**Repository** | [**ApiTasksTaskFileRepository**](ApiTasksTaskFileRepository.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiTasksTaskFile = Initialize-PSOpenAPIToolsApiTasksTaskFile  -SourceType null `
 -Content null `
 -ContentPath null `
 -ContentRef null `
 -Repository null
```

- Convert the resource to JSON
```powershell
$ApiTasksTaskFile | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

