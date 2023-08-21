# # ApiJobsIdJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the Job | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**enabled** | **bool** | Use this to set enabled state | [optional] [default to true]
**task** | [**\OpenAPI\Client\Model\ApiJobsIdJobTask**](ApiJobsIdJobTask.md) |  | [optional]
**workflow** | [**\OpenAPI\Client\Model\ApiJobsIdJobTask**](ApiJobsIdJobTask.md) |  | [optional]
**scan_path** | **string** | Scan Checklist. Only applies to type scap-package. | [optional]
**security_profile** | **string** | Security Profile. Only applies to type scap-package. | [optional]
**target_type** | **string** | Target type where job will execute | [optional]
**targets** | [**\OpenAPI\Client\Model\ApiJobsIdJobTargets[]**](ApiJobsIdJobTargets.md) |  | [optional]
**instance_label** | **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional]
**server_label** | **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional]
**schedule_mode** | [**OneOfStringLong**](OneOfStringLong.md) |  | [optional]
**custom_options** | **object** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional]
**custom_config** | **string** | Job custom configuration (String in JSON format) | [optional]
**date_time** | [**\DateTime**](\DateTime.md) | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional]
**run** | **bool** | If true, executes job | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
