# AddAlertsRequestAlert

Payload for creating a new monitoring alert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the alert | 
**min_duration** | **int** | Duration in minutes of the delay before sending notification(s) | [optional]  if omitted the server will use the default value of 0
**min_severity** | **str** | Severity level threshold for sending notifications. | [optional]  if omitted the server will use the default value of "critical"
**active** | **bool** | Set to false to disable notifications | [optional]  if omitted the server will use the default value of True
**all_checks** | **bool** | Trigger for all checks | [optional]  if omitted the server will use the default value of False
**all_groups** | **bool** | Trigger for all check groups | [optional]  if omitted the server will use the default value of False
**all_apps** | **bool** | Trigger for all monitor apps | [optional]  if omitted the server will use the default value of False
**checks** | **[int]** |  | [optional] 
**groups** | **[int]** |  | [optional] 
**apps** | **[int]** |  | [optional] 
**contacts** | [**[AddAlertsRequestAlertContactsInner]**](AddAlertsRequestAlertContactsInner.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


