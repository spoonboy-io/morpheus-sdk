# # ApiTasksIdExecuteJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the execution job. Can be used to find execution results with &#x60;/api/processes?name&#x3D;&#x60; | [optional]
**target_type** | **string** | The target context for task execution. This is required for tasks with &#x60;executeTarget&#x60; set to &#x60;resource&#x60;. | [optional]
**instances** | **int[]** | Array of Instance IDs. Only applicable if &#x60;targetType&#x60; is instance. | [optional]
**servers** | **int[]** | Array of Server IDs. Only applicable if &#x60;targetType&#x60; is &#x60;server&#x60;. | [optional]
**instance_label** | **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional]
**server_label** | **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional]
**custom_options** | **object** | Map of options to be used as values in the task. These correspond to option types. | [optional]
**custom_config** | **string** | String of custom configuration values as JSON. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
