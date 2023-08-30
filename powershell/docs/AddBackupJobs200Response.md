# AddBackupJobs200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**BackupJob**](BackupJob.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddBackupJobs200Response = Initialize-PSOpenAPIToolsAddBackupJobs200Response  -Job null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddBackupJobs200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

