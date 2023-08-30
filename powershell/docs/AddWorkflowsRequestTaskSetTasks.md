# AddWorkflowsRequestTaskSetTasks
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **Int64** | Task ID | [optional] 
**TaskPhase** | **String** | Task Phase. Pass operation for &#x60;operational&#x60; workflows. | [optional] [default to "provision"]

## Examples

- Prepare the resource
```powershell
$AddWorkflowsRequestTaskSetTasks = Initialize-PSOpenAPIToolsAddWorkflowsRequestTaskSetTasks  -TaskId null `
 -TaskPhase null
```

- Convert the resource to JSON
```powershell
$AddWorkflowsRequestTaskSetTasks | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

