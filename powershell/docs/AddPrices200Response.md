# AddPrices200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | [**Price**](Price.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddPrices200Response = Initialize-PSOpenAPIToolsAddPrices200Response  -Price null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddPrices200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

