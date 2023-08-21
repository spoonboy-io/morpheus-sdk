# # Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional]
**name** | **string** |  | [optional]
**labels** | **string[]** |  | [optional]
**type** | [**\OpenAPI\Client\Model\InlineResponse20094Network**](InlineResponse20094Network.md) |  | [optional]
**workflow** | [**\OpenAPI\Client\Model\JobWorkflow**](JobWorkflow.md) |  | [optional]
**task** | [**\OpenAPI\Client\Model\JobTask**](JobTask.md) |  | [optional]
**security_package** | [**\OpenAPI\Client\Model\JobSecurityPackage**](JobSecurityPackage.md) |  | [optional]
**job_summary** | **string** |  | [optional]
**schedule_mode** | [**OneOfStringLong**](OneOfStringLong.md) |  | [optional]
**date_time** | **string** |  | [optional]
**status** | **string** |  | [optional]
**namespace** | **string** |  | [optional]
**category** | **string** |  | [optional]
**description** | **string** |  | [optional]
**enabled** | **bool** |  | [optional]
**date_created** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_updated** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_run** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_result** | **string** |  | [optional]
**created_by** | [**\OpenAPI\Client\Model\JobCreatedBy**](JobCreatedBy.md) |  | [optional]
**target_type** | **string** |  | [optional]
**targets** | [**\OpenAPI\Client\Model\JobTargets[]**](JobTargets.md) |  | [optional]
**scan_path** | **string** | Scan Checklist. Only applies to type scap-package. | [optional]
**security_profile** | **string** | Security Profile. Only applies to type scap-package. | [optional]
**custom_config** | **string** |  | [optional]
**custom_options** | [**\OpenAPI\Client\Model\JobCustomOptions**](JobCustomOptions.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
