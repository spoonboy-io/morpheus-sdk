# BackupBackupType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup Type ID | [optional] 
**Code** | **String** | Backup Type Code | [optional] 
**Name** | **String** | Backup Type Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupBackupType = Initialize-PSOpenAPIToolsBackupBackupType  -Id null `
 -Code null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupBackupType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

