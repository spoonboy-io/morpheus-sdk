# ContainerPlan
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerPlan = Initialize-PSOpenAPIToolsContainerPlan  -Id null `
 -Code null `
 -Name null
```

- Convert the resource to JSON
```powershell
$ContainerPlan | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

