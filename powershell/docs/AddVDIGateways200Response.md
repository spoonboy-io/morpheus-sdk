# AddVDIGateways200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiGateway** | [**VdiGateway**](VdiGateway.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddVDIGateways200Response = Initialize-PSOpenAPIToolsAddVDIGateways200Response  -VdiGateway null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddVDIGateways200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

