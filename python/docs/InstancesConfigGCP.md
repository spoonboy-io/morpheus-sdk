# InstancesConfigGCP


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_agent** | **bool, none_type** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional]  if omitted the server will use the default value of False
**google_zone_id** | **int** | External ID of the Google zone to use for instance. | [optional] 
**external_ip_type** | **int** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional] 
**network_tags** | **str** | Network Tags | [optional] 
**service_account** | **str** | Service Account | [optional] 
**access_scope** | **str** | Access Scope | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


