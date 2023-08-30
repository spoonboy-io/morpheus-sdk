# AddCypherKey200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VarData** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**LeaseDuration** | **Int64** |  | [optional] 
**Cypher** | [**Cypher**](Cypher.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCypherKey200Response = Initialize-PSOpenAPIToolsAddCypherKey200Response  -VarData null `
 -Type null `
 -LeaseDuration null `
 -Cypher null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddCypherKey200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

