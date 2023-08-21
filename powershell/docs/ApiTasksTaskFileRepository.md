# ApiTasksTaskFileRepository
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Code Repository ID, required for type &#x60;repository&#x60;. Use &#x60;/api/options/codeRepositories&#x60; to see available repositories. | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiTasksTaskFileRepository = Initialize-PSOpenAPIToolsApiTasksTaskFileRepository  -Id null
```

- Convert the resource to JSON
```powershell
$ApiTasksTaskFileRepository | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

