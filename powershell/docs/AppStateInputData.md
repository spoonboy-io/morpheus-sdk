# AppStateInputData
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **String** |  | [optional] 
**Name** | [**OneOfstringobject**](OneOfstringobject.md) |  | [optional] 
**Type** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppStateInputData = Initialize-PSOpenAPIToolsAppStateInputData  -Key null `
 -Name null `
 -Type null
```

- Convert the resource to JSON
```powershell
$AppStateInputData | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

