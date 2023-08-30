# Job


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**labels** | **[str], none_type** |  | [optional] 
**type** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**workflow** | [**JobWorkflow**](JobWorkflow.md) |  | [optional] 
**task** | [**JobTask**](JobTask.md) |  | [optional] 
**security_package** | [**JobSecurityPackage**](JobSecurityPackage.md) |  | [optional] 
**job_summary** | **str, none_type** |  | [optional] 
**schedule_mode** | [**JobScheduleMode**](JobScheduleMode.md) |  | [optional] 
**date_time** | **str, none_type** |  | [optional] 
**status** | **str, none_type** |  | [optional] 
**namespace** | **str, none_type** |  | [optional] 
**category** | **str, none_type** |  | [optional] 
**description** | **str, none_type** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**last_run** | **datetime** |  | [optional] 
**last_result** | **str, none_type** |  | [optional] 
**created_by** | [**JobCreatedBy**](JobCreatedBy.md) |  | [optional] 
**target_type** | **str, none_type** |  | [optional] 
**targets** | [**[JobTargetsInner], none_type**](JobTargetsInner.md) |  | [optional] 
**scan_path** | **str, none_type** | Scan Checklist. Only applies to type scap-package. | [optional] 
**security_profile** | **str, none_type** | Security Profile. Only applies to type scap-package. | [optional] 
**custom_config** | **str, none_type** |  | [optional] 
**custom_options** | [**JobCustomOptions**](JobCustomOptions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


