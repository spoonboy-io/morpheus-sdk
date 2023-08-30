# AddVirtualImage200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualImage** | [**VirtualImage**](VirtualImage.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddVirtualImage200Response = Initialize-PSOpenAPIToolsAddVirtualImage200Response  -VirtualImage null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddVirtualImage200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

