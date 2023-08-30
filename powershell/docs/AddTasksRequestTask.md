# AddTasksRequestTask
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name for the task | 
**Code** | **String** | A unique code for the task | [optional] 
**Visibility** | **String** | private or public | [optional] [default to "private"]
**TaskType** | [**AddTasksRequestTaskTaskType**](AddTasksRequestTaskTaskType.md) |  | 
**Labels** | **String[]** | An array of Category labels for filtering | [optional] 
**TaskOptions** | [**SystemCollectionsHashtable**](.md) | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**ResultType** | **String** |  | [optional] 
**ExecuteTarget** | **String** | The execution target. eg. local,remote,resource. The default value varies by task type.  | 
**Retryable** | **Boolean** | If the task should be retried or not. | [optional] [default to $false]
**RetryCount** | **Int32** | The number of times to retry. | [optional] 
**RetryDelaySeconds** | **Decimal** | The delay, between retries. | [optional] 
**File** | [**AddTasksRequestTaskFile**](AddTasksRequestTaskFile.md) |  | [optional] 
**Credential** | [**AddTasksRequestTaskCredential**](AddTasksRequestTaskCredential.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddTasksRequestTask = Initialize-PSOpenAPIToolsAddTasksRequestTask  -Name Sample Task `
 -Code aTaskCode `
 -Visibility null `
 -TaskType null `
 -Labels null `
 -TaskOptions null `
 -ResultType null `
 -ExecuteTarget null `
 -Retryable null `
 -RetryCount null `
 -RetryDelaySeconds null `
 -File null `
 -Credential null
```

- Convert the resource to JSON
```powershell
$AddTasksRequestTask | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

