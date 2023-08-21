# ApiBackupsJobsIdJob
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the backup job | [optional] 
**Code** | **String** | A code for the backup job | [optional] 
**ScheduleId** | **Int64** | Execute Schedule ID to use for the backup job | [optional] 
**RetentionCount** | **Int64** | Retention Count | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiBackupsJobsIdJob = Initialize-PSOpenAPIToolsApiBackupsJobsIdJob  -Name null `
 -Code null `
 -ScheduleId null `
 -RetentionCount null
```

- Convert the resource to JSON
```powershell
$ApiBackupsJobsIdJob | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

