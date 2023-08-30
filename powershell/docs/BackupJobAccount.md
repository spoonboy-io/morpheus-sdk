# BackupJobAccount
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Tenant ID | [optional] 
**Name** | **String** | Tenant Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupJobAccount = Initialize-PSOpenAPIToolsBackupJobAccount  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$BackupJobAccount | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

