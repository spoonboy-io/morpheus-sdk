# GetImageBuildExecutions200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuildExecutions** | [**ImageBuildExecution[]**](ImageBuildExecution.md) |  | [optional] 
**ImageBuildExecutionCount** | **Int64** |  | [optional] 
**ImageBuild** | [**ImageBuild**](ImageBuild.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetImageBuildExecutions200Response = Initialize-PSOpenAPIToolsGetImageBuildExecutions200Response  -ImageBuildExecutions null `
 -ImageBuildExecutionCount null `
 -ImageBuild null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetImageBuildExecutions200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

