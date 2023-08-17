# UpdateBackupSettingsRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSettings** | [**BackupSettingsUpdate**](BackupSettingsUpdate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$UpdateBackupSettingsRequest = Initialize-PSOpenAPIToolsUpdateBackupSettingsRequest  -BackupSettings null
```

- Convert the resource to JSON
```powershell
$UpdateBackupSettingsRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

