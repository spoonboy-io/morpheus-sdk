# AddIntegrations200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | [**AddIntegrations200ResponseAllOfIntegration**](AddIntegrations200ResponseAllOfIntegration.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddIntegrations200Response = Initialize-PSOpenAPIToolsAddIntegrations200Response  -Integration null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddIntegrations200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

