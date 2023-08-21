# ApiTaskSetsTaskSet
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name for the workflow | 
**Description** | **String** | A description of the workflow | [optional] 
**Labels** | **String[]** | An array of Category labels for filtering | [optional] 
**Type** | **String** | Workflow type | [optional] [default to "provision"]
**Visibility** | **String** | private or public | [optional] [default to "private"]
**OptionTypes** | **Int64[]** | List of option type IDs for use with operational workflow configuration. | [optional] 
**Tasks** | [**ApiTaskSetsTaskSetTasks**](ApiTaskSetsTaskSetTasks.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiTaskSetsTaskSet = Initialize-PSOpenAPIToolsApiTaskSetsTaskSet  -Name Sample Workflow `
 -Description null `
 -Labels null `
 -Type null `
 -Visibility null `
 -OptionTypes null `
 -Tasks null
```

- Convert the resource to JSON
```powershell
$ApiTaskSetsTaskSet | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

