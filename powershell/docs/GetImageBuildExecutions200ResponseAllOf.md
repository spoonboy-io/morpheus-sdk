# GetImageBuildExecutions200ResponseAllOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuildExecutions** | [**ImageBuildExecution[]**](ImageBuildExecution.md) |  | [optional] 
**ImageBuildExecutionCount** | **Int64** |  | [optional] 
**ImageBuild** | [**ImageBuild**](ImageBuild.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetImageBuildExecutions200ResponseAllOf = Initialize-PSOpenAPIToolsGetImageBuildExecutions200ResponseAllOf  -ImageBuildExecutions null `
 -ImageBuildExecutionCount null `
 -ImageBuild null
```

- Convert the resource to JSON
```powershell
$GetImageBuildExecutions200ResponseAllOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

