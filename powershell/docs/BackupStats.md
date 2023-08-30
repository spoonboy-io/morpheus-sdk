# BackupStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSize** | **Int64** | Total size of all backups in bytes | [optional] 
**AvgSize** | **Int64** | Average size of each backup in bytes | [optional] 
**TotalCompleted** | **Int64** | Total completed count | [optional] 
**Success** | **Int64** | Successful backup count | [optional] 
**Failed** | **Int64** | Failed backup count | [optional] 
**SuccessRate** | **Double** | Success rate 0-1 | [optional] 
**FailRate** | **Double** | Failure rate 0-1 | [optional] 
**LastFiveResults** | **String[]** | List of the last 5 backup result statuses | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupStats = Initialize-PSOpenAPIToolsBackupStats  -TotalSize null `
 -AvgSize null `
 -TotalCompleted null `
 -Success null `
 -Failed null `
 -SuccessRate null `
 -FailRate null `
 -LastFiveResults null
```

- Convert the resource to JSON
```powershell
$BackupStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

