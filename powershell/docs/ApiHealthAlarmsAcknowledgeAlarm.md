# ApiHealthAlarmsAcknowledgeAlarm
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | **Boolean** | Pass &#x60;true&#x60; to ackowledge an alarm, or pass &#x60;false&#x60; to unacknowledge it. | 
**Ids** | **Int64[]** | Array of Alarm ID(s)to be updated. | [optional] 
**All** | **Boolean** | Pass &#x60;true&#x60; to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged.  | [optional] [default to $false]

## Examples

- Prepare the resource
```powershell
$ApiHealthAlarmsAcknowledgeAlarm = Initialize-PSOpenAPIToolsApiHealthAlarmsAcknowledgeAlarm  -Acknowledged null `
 -Ids null `
 -All null
```

- Convert the resource to JSON
```powershell
$ApiHealthAlarmsAcknowledgeAlarm | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

