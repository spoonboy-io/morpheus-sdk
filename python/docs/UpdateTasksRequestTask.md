# UpdateTasksRequestTask


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A unique name for the task | [optional] 
**code** | **str** | A unique code for the task | [optional] 
**visibility** | **str** | Visibility | [optional]  if omitted the server will use the default value of "private"
**task_type** | [**AddTasksRequestTaskTaskType**](AddTasksRequestTaskTaskType.md) |  | [optional] 
**labels** | **[str]** | An array of Category labels for filtering | [optional] 
**task_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**result_type** | **str, none_type** |  | [optional] 
**execute_target** | **str** | The execution target. eg. local,remote,resource. The default value varies by task type.  | [optional] 
**retryable** | **bool** | If the task should be retried or not. | [optional]  if omitted the server will use the default value of False
**retry_count** | **int** | The number of times to retry. | [optional] 
**retry_delay_seconds** | **float** | The delay, between retries. | [optional] 
**file** | [**AddTasksRequestTaskFile**](AddTasksRequestTaskFile.md) |  | [optional] 
**credential** | [**UpdateTasksRequestTaskCredential**](UpdateTasksRequestTaskCredential.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


