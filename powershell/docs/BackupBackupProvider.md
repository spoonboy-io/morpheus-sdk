# BackupBackupProvider
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup Provider ID | [optional] 
**Name** | **String** | Backup Provider Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupBackupProvider = Initialize-PSOpenAPIToolsBackupBackupProvider  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupBackupProvider | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

