# MorpheusApi.ApiTasksIdTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name for the task | [optional] 
**code** | **String** | A unique code for the task | [optional] 
**visibility** | **String** | Visibility | [optional] [default to &#39;private&#39;]
**taskType** | [**ApiTasksTaskTaskType**](ApiTasksTaskTaskType.md) |  | [optional] 
**labels** | **[String]** | An array of Category labels for filtering | [optional] 
**taskOptions** | **Object** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**resultType** | **String** |  | [optional] 
**executeTarget** | **String** | The execution target. eg. local,remote,resource. The default value varies by task type.  | [optional] 
**retryable** | **Boolean** | If the task should be retried or not. | [optional] [default to false]
**retryCount** | **Number** | The number of times to retry. | [optional] 
**retryDelaySeconds** | **Number** | The delay, between retries. | [optional] 
**file** | [**ApiTasksTaskFile**](ApiTasksTaskFile.md) |  | [optional] 
**credential** | [**OneOfobjectobject**](OneOfobjectobject.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential.  | [optional] 



## Enum: ResultTypeEnum


* `exitCode` (value: `"exitCode"`)

* `keyValue` (value: `"keyValue"`)

* `json` (value: `"json"`)





## Enum: ExecuteTargetEnum


* `local` (value: `"local"`)

* `remote` (value: `"remote"`)

* `resource` (value: `"resource"`)




