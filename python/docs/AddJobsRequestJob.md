# AddJobsRequestJob


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**enabled** | **bool** | Use this to set enabled state | [optional]  if omitted the server will use the default value of True
**instance_label** | **str** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**server_label** | **str** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**custom_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**custom_config** | **str** | Job custom configuration (String in JSON format) | [optional] 
**date_time** | **datetime** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**run** | **bool** | If true, executes job | [optional] 
**scan_path** | **str, none_type** | Scan Checklist | [optional] 
**security_profile** | **str, none_type** | Security Profile | [optional] 
**name** | **str** | A name for the Job | [optional] 
**task** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | [optional] 
**workflow** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | [optional] 
**target_type** | **str** | Target type where job will execute | [optional] 
**targets** | [**[WorkflowJobPayloadTargetsInner], none_type**](WorkflowJobPayloadTargetsInner.md) |  | [optional] 
**schedule_mode** | [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | [optional] 
**security_package** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


