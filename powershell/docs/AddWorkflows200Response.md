# AddWorkflows200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSet** | [**Workflow**](Workflow.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddWorkflows200Response = Initialize-PSOpenAPIToolsAddWorkflows200Response  -TaskSet null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddWorkflows200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

