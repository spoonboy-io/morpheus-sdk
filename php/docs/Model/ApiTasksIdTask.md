# # ApiTasksIdTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A unique name for the task | [optional]
**code** | **string** | A unique code for the task | [optional]
**visibility** | **string** | Visibility | [optional] [default to 'private']
**task_type** | [**\OpenAPI\Client\Model\ApiTasksTaskTaskType**](ApiTasksTaskTaskType.md) |  | [optional]
**labels** | **string[]** | An array of Category labels for filtering | [optional]
**task_options** | **object** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional]
**result_type** | **string** |  | [optional]
**execute_target** | **string** | The execution target. eg. local,remote,resource. The default value varies by task type. | [optional]
**retryable** | **bool** | If the task should be retried or not. | [optional] [default to false]
**retry_count** | **int** | The number of times to retry. | [optional]
**retry_delay_seconds** | **float** | The delay, between retries. | [optional]
**file** | [**\OpenAPI\Client\Model\ApiTasksTaskFile**](ApiTasksTaskFile.md) |  | [optional]
**credential** | [**OneOfObjectObject**](OneOfObjectObject.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
