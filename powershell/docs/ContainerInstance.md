# ContainerInstance
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**Name** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerInstance = Initialize-PSOpenAPIToolsContainerInstance  -Id null `
 -Name null
```

- Convert the resource to JSON
```powershell
$ContainerInstance | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

