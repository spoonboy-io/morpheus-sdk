# AcknowledgeHealthAlarmsRequestAlarm


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**acknowledged** | **bool** | Pass &#x60;true&#x60; to ackowledge an alarm, or pass &#x60;false&#x60; to unacknowledge it. | 
**ids** | **[int]** | Array of Alarm ID(s)to be updated. | [optional]  if omitted the server will use the default value of []
**all** | **bool** | Pass &#x60;true&#x60; to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged.  | [optional]  if omitted the server will use the default value of False
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


