# ApiTaskSetsIdTaskSet
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name for the workflow | [optional] 
**Description** | **String** | A description of the workflow | [optional] 
**Labels** | **String[]** | An array of Category labels for filtering | [optional] 
**Type** | **String** | Workflow type | [optional] [default to "provision"]
**OptionTypes** | **Int64[]** | List of option type IDs for use with operational workflow configuration. | [optional] 
**Tasks** | [**ApiTaskSetsTaskSetTasks**](ApiTaskSetsTaskSetTasks.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiTaskSetsIdTaskSet = Initialize-PSOpenAPIToolsApiTaskSetsIdTaskSet  -Name Sample Workflow `
 -Description null `
 -Labels null `
 -Type null `
 -OptionTypes null `
 -Tasks null
```

- Convert the resource to JSON
```powershell
$ApiTaskSetsIdTaskSet | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

