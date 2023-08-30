# ContainerTypeCreateContainerPortsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | 
**Port** | **Int64** |  | 
**LoadBalanceProtocol** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerTypeCreateContainerPortsInner = Initialize-PSOpenAPIToolsContainerTypeCreateContainerPortsInner  -Name null `
 -Port null `
 -LoadBalanceProtocol null
```

- Convert the resource to JSON
```powershell
$ContainerTypeCreateContainerPortsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

