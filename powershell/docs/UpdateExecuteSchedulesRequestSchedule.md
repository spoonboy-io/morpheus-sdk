# UpdateExecuteSchedulesRequestSchedule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the execute schedule | [optional] 
**Description** | **String** | A description for the execute schedule | [optional] 
**ScheduleType** | **String** | Type of schedule | [optional] 
**ScheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to "UTC"]
**Cron** | **String** | Cron Expression. The default is daily at midnight | [optional] [default to "0 0 * * *"]
**Enabled** | **Boolean** | Is enabled | [optional] [default to $true]

## Examples

- Prepare the resource
```powershell
$UpdateExecuteSchedulesRequestSchedule = Initialize-PSOpenAPIToolsUpdateExecuteSchedulesRequestSchedule  -Name Sample Execution `
 -Description null `
 -ScheduleType null `
 -ScheduleTimezone null `
 -Cron null `
 -Enabled null
```

- Convert the resource to JSON
```powershell
$UpdateExecuteSchedulesRequestSchedule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

