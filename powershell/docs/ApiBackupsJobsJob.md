# ApiBackupsJobsJob
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the backup job | 
**Code** | **String** | A code for the backup job | [optional] 
**ScheduleId** | **Int64** | Execute Schedule ID to use for the backup job | [optional] 
**RetentionCount** | **Int64** | Retention Count. By default the backup settings value will be used. | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiBackupsJobsJob = Initialize-PSOpenAPIToolsApiBackupsJobsJob  -Name null `
 -Code null `
 -ScheduleId null `
 -RetentionCount null
```

- Convert the resource to JSON
```powershell
$ApiBackupsJobsJob | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

