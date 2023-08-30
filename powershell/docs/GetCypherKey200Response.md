# GetCypherKey200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**VarData** | [**GetCypherKey200ResponseAllOfData**](GetCypherKey200ResponseAllOfData.md) |  | [optional] 
**Type** | **String** | Type of data that was written to the key | [optional] 
**LeaseDuration** | **Int32** | Lease duration in seconds, 0 means no expiry. | [optional] 
**Cypher** | [**Cypher**](Cypher.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetCypherKey200Response = Initialize-PSOpenAPIToolsGetCypherKey200Response  -Success null `
 -VarData null `
 -Type null `
 -LeaseDuration null `
 -Cypher null
```

- Convert the resource to JSON
```powershell
$GetCypherKey200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

