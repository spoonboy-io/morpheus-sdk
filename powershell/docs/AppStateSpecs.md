# AppStateSpecs
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Template** | [**AppStateTemplate**](AppStateTemplate.md) |  | [optional] 
**Isolated** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppStateSpecs = Initialize-PSOpenAPIToolsAppStateSpecs  -Id null `
 -Name null `
 -Template null `
 -Isolated null
```

- Convert the resource to JSON
```powershell
$AppStateSpecs | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

