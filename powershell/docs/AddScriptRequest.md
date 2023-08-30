# AddScriptRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerScript** | [**ScriptCreate**](ScriptCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddScriptRequest = Initialize-PSOpenAPIToolsAddScriptRequest  -ContainerScript null
```

- Convert the resource to JSON
```powershell
$AddScriptRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

