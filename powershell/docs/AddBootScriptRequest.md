# AddBootScriptRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootScript** | [**BootScriptsCreate**](BootScriptsCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddBootScriptRequest = Initialize-PSOpenAPIToolsAddBootScriptRequest  -BootScript null
```

- Convert the resource to JSON
```powershell
$AddBootScriptRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

