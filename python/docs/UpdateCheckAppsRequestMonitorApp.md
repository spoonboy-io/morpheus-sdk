# UpdateCheckAppsRequestMonitorApp

Payload for creating a new monitoring check app

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the check app | [optional] 
**description** | **str** | Optional description field | [optional] 
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional]  if omitted the server will use the default value of True
**severity** | **str** | Severity level of incidents that are created when this check fails | [optional]  if omitted the server will use the default value of "critical"
**active** | **bool** | Used to determine if check app is active | [optional]  if omitted the server will use the default value of True
**checks** | **[int]** |  | [optional] 
**check_groups** | **[int]** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


