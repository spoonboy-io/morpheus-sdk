# BackupInstance
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Instance ID | [optional] 
**Name** | **String** | Instance Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupInstance = Initialize-PSOpenAPIToolsBackupInstance  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupInstance | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

