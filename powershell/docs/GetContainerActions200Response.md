# GetContainerActions200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerIds** | **Int64[]** |  | [optional] 
**Actions** | [**GetContainerActions200ResponseActionsInner[]**](GetContainerActions200ResponseActionsInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetContainerActions200Response = Initialize-PSOpenAPIToolsGetContainerActions200Response  -ContainerIds null `
 -Actions null
```

- Convert the resource to JSON
```powershell
$GetContainerActions200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

