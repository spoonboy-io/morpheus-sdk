# AddCypherKey200ResponseAllOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VarData** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**LeaseDuration** | **Int64** |  | [optional] 
**Cypher** | [**Cypher**](Cypher.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCypherKey200ResponseAllOf = Initialize-PSOpenAPIToolsAddCypherKey200ResponseAllOf  -VarData null `
 -Type null `
 -LeaseDuration null `
 -Cypher null
```

- Convert the resource to JSON
```powershell
$AddCypherKey200ResponseAllOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

