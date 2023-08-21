# BackupJob
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup Job ID | [optional] 
**Name** | **String** | Backup Job Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupJob = Initialize-PSOpenAPIToolsBackupJob  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupJob | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

