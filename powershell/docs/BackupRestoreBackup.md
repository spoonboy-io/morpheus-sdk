# BackupRestoreBackup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup ID | [optional] 
**Name** | **String** | Backup Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupRestoreBackup = Initialize-PSOpenAPIToolsBackupRestoreBackup  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupRestoreBackup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

