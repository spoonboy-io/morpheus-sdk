# BackupLastResult
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Last Result ID | [optional] 
**Name** | **String** | Last Result Status | [optional] 
**DateCreated** | **System.DateTime** | Last Result Date Created | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupLastResult = Initialize-PSOpenAPIToolsBackupLastResult  -Id null `
 -Name null `
 -DateCreated null
```

- Convert the resource to JSON
```powershell
$BackupLastResult | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

