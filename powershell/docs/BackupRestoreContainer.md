# BackupRestoreContainer
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Container ID | [optional] 
**Name** | **String** | Container Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupRestoreContainer = Initialize-PSOpenAPIToolsBackupRestoreContainer  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupRestoreContainer | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

