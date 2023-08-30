# GetCypherKey200ResponseAllOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VarData** | [**GetCypherKey200ResponseAllOfData**](GetCypherKey200ResponseAllOfData.md) |  | [optional] 
**Type** | **String** | Type of data that was written to the key | [optional] 
**LeaseDuration** | **Int32** | Lease duration in seconds, 0 means no expiry. | [optional] 

## Examples

- Prepare the resource
```powershell
$GetCypherKey200ResponseAllOf = Initialize-PSOpenAPIToolsGetCypherKey200ResponseAllOf  -VarData null `
 -Type null `
 -LeaseDuration null
```

- Convert the resource to JSON
```powershell
$GetCypherKey200ResponseAllOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

