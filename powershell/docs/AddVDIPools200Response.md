# AddVDIPools200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiPool** | [**VdiPool**](VdiPool.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddVDIPools200Response = Initialize-PSOpenAPIToolsAddVDIPools200Response  -VdiPool null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddVDIPools200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

