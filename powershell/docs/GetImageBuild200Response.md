# GetImageBuild200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuild** | [**ImageBuild**](ImageBuild.md) |  | [optional] 
**ImageBuildExecutions** | [**ImageBuildExecution[]**](ImageBuildExecution.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetImageBuild200Response = Initialize-PSOpenAPIToolsGetImageBuild200Response  -ImageBuild null `
 -ImageBuildExecutions null
```

- Convert the resource to JSON
```powershell
$GetImageBuild200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

