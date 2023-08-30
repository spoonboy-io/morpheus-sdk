# BootScriptsCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **String** | A name for the boot script | [optional] 
**Content** | **String** | The script content | [optional] 

## Examples

- Prepare the resource
```powershell
$BootScriptsCreate = Initialize-PSOpenAPIToolsBootScriptsCreate  -FileName null `
 -Content null
```

- Convert the resource to JSON
```powershell
$BootScriptsCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

