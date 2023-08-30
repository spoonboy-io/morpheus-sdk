# AppBlueprint
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppBlueprint = Initialize-PSOpenAPIToolsAppBlueprint  -Id null `
 -Name null `
 -Type null
```

- Convert the resource to JSON
```powershell
$AppBlueprint | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

