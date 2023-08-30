# AddPriceSets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | [**PriceSet**](PriceSet.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddPriceSets200Response = Initialize-PSOpenAPIToolsAddPriceSets200Response  -Budget null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddPriceSets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

