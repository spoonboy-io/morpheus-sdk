# AddBackupsRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**AddBackupsRequestBackup**](AddBackupsRequestBackup.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddBackupsRequest = Initialize-PSOpenAPIToolsAddBackupsRequest  -Backup null
```

- Convert the resource to JSON
```powershell
$AddBackupsRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

