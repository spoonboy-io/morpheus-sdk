# CypherAccessPolicyTypeConfiguration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPattern** | **String** |  | [optional] 
**Read** | **Boolean** |  | [optional] 
**Write** | **Boolean** |  | [optional] 
**Update** | **Boolean** |  | [optional] 
**Delete** | **Boolean** |  | [optional] 
**List** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CypherAccessPolicyTypeConfiguration = Initialize-PSOpenAPIToolsCypherAccessPolicyTypeConfiguration  -KeyPattern null `
 -Read null `
 -Write null `
 -Update null `
 -Delete null `
 -List null
```

- Convert the resource to JSON
```powershell
$CypherAccessPolicyTypeConfiguration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

