# ApiPluginsIdPlugin
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **Boolean** | Can be used to disable the plugin | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object that contains settings for the applicable option types. | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiPluginsIdPlugin = Initialize-PSOpenAPIToolsApiPluginsIdPlugin  -Enabled null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ApiPluginsIdPlugin | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

