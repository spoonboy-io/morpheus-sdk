# ContainerTypeContainerPorts
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | [optional] 
**Port** | **Int64** |  | [optional] 
**LoadBalanceProtocol** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerTypeContainerPorts = Initialize-PSOpenAPIToolsContainerTypeContainerPorts  -Name null `
 -Port null `
 -LoadBalanceProtocol null
```

- Convert the resource to JSON
```powershell
$ContainerTypeContainerPorts | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

