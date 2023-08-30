# GetBackupRestores200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restore** | [**BackupRestore**](BackupRestore.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBackupRestores200Response = Initialize-PSOpenAPIToolsGetBackupRestores200Response  -Restore null
```

- Convert the resource to JSON
```powershell
$GetBackupRestores200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

