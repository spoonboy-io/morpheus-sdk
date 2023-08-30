# CatalogTypesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Context** | **String** |  | [optional] 
**Featured** | **Boolean** |  | [optional] 
**AllowQuantity** | **Boolean** |  | [optional] 
**ImagePath** | **String** |  | [optional] 
**DarkImagePath** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogTypesInner = Initialize-PSOpenAPIToolsCatalogTypesInner  -Id null `
 -Name null `
 -Description null `
 -Type null `
 -Context null `
 -Featured null `
 -AllowQuantity null `
 -ImagePath null `
 -DarkImagePath null
```

- Convert the resource to JSON
```powershell
$CatalogTypesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

