# AddBackups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**Backup**](Backup.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddBackups200Response = Initialize-PSOpenAPIToolsAddBackups200Response  -Backup null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddBackups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

