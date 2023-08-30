# GetStateInstance200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**InstanceState**](InstanceState.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetStateInstance200Response = Initialize-PSOpenAPIToolsGetStateInstance200Response  -Instance null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetStateInstance200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

