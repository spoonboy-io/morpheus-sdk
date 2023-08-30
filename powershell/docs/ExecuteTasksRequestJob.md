# ExecuteTasksRequestJob
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the execution job. Can be used to find execution results with &#x60;/api/processes?name&#x3D;&#x60; | [optional] 
**TargetType** | **String** | The target context for task execution. This is required for tasks with &#x60;executeTarget&#x60; set to &#x60;resource&#x60;. | [optional] 
**Instances** | **Int64[]** | Array of Instance IDs. Only applicable if &#x60;targetType&#x60; is instance. | [optional] 
**Servers** | **Int64[]** | Array of Server IDs. Only applicable if &#x60;targetType&#x60; is &#x60;server&#x60;. | [optional] 
**InstanceLabel** | **String** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | **String** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**CustomOptions** | [**SystemCollectionsHashtable**](.md) | Map of options to be used as values in the task. These correspond to option types. | [optional] 
**CustomConfig** | **String** | String of custom configuration values as JSON. | [optional] 

## Examples

- Prepare the resource
```powershell
$ExecuteTasksRequestJob = Initialize-PSOpenAPIToolsExecuteTasksRequestJob  -Name null `
 -TargetType null `
 -Instances null `
 -Servers null `
 -InstanceLabel null `
 -ServerLabel null `
 -CustomOptions null `
 -CustomConfig null
```

- Convert the resource to JSON
```powershell
$ExecuteTasksRequestJob | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

