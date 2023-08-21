

# ApiTasksTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name for the task | 
**code** | **String** | A unique code for the task |  [optional]
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | private or public |  [optional]
**taskType** | [**ApiTasksTaskTaskType**](ApiTasksTaskTaskType.md) |  | 
**labels** | **List&lt;String&gt;** | An array of Category labels for filtering |  [optional]
**taskOptions** | **Object** | Map of options specific to each &#x60;task type&#x60;. eg. script |  [optional]
**resultType** | [**ResultTypeEnum**](#ResultTypeEnum) |  |  [optional]
**executeTarget** | [**ExecuteTargetEnum**](#ExecuteTargetEnum) | The execution target. eg. local,remote,resource. The default value varies by task type.  | 
**retryable** | **Boolean** | If the task should be retried or not. |  [optional]
**retryCount** | **Integer** | The number of times to retry. |  [optional]
**retryDelaySeconds** | **BigDecimal** | The delay, between retries. |  [optional]
**file** | [**ApiTasksTaskFile**](ApiTasksTaskFile.md) |  |  [optional]
**credential** | [**OneOfobjectobject**](OneOfobjectobject.md) | Map containing Credential ID or the default &#x60;{\&quot;type\&quot;: \&quot;local\&quot;}&#x60; which means use the values set in the local task options username and password instead of associating a credential.  |  [optional]



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



## Enum: ResultTypeEnum

Name | Value
---- | -----
EXITCODE | &quot;exitCode&quot;
KEYVALUE | &quot;keyValue&quot;
JSON | &quot;json&quot;



## Enum: ExecuteTargetEnum

Name | Value
---- | -----
LOCAL | &quot;local&quot;
REMOTE | &quot;remote&quot;
RESOURCE | &quot;resource&quot;



