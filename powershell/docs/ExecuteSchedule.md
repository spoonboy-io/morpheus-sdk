# ExecuteSchedule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**ScheduleType** | **String** |  | [optional] 
**ScheduleTimezone** | **String** |  | [optional] 
**Cron** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ExecuteSchedule = Initialize-PSOpenAPIToolsExecuteSchedule  -Id null `
 -Name null `
 -Description null `
 -Visibility null `
 -Enabled null `
 -ScheduleType null `
 -ScheduleTimezone null `
 -Cron null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$ExecuteSchedule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

