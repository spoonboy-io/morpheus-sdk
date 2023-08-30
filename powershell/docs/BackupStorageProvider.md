# BackupStorageProvider
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Storage Provider ID | [optional] 
**Name** | **String** | Storage Provider Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupStorageProvider = Initialize-PSOpenAPIToolsBackupStorageProvider  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupStorageProvider | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

