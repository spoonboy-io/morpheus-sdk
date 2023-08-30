# AddVdiAllocation200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desktop** | [**Vdi**](Vdi.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddVdiAllocation200Response = Initialize-PSOpenAPIToolsAddVdiAllocation200Response  -Desktop null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddVdiAllocation200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

