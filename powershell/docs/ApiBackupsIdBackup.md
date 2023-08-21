# ApiBackupsIdBackup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the backup | [optional] 
**JobId** | **Int64** | The Backup Job ID to assign the backup to. This determines when the backup is run. | [optional] 
**Enabled** | **Boolean** | Can be used to enable or disable the backup | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiBackupsIdBackup = Initialize-PSOpenAPIToolsApiBackupsIdBackup  -Name null `
 -JobId null `
 -Enabled null
```

- Convert the resource to JSON
```powershell
$ApiBackupsIdBackup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

