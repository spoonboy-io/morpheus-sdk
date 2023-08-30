# BackupCreationPolicyTypeConfiguration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateBackupType** | **String** |  | [optional] 
**CreateBackup** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupCreationPolicyTypeConfiguration = Initialize-PSOpenAPIToolsBackupCreationPolicyTypeConfiguration  -CreateBackupType null `
 -CreateBackup null
```

- Convert the resource to JSON
```powershell
$BackupCreationPolicyTypeConfiguration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

