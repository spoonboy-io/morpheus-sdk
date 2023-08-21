# CatalogCart
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Items** | [**CatalogItem[]**](CatalogItem.md) |  | [optional] 
**Stats** | [**CatalogCartStats**](CatalogCartStats.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogCart = Initialize-PSOpenAPIToolsCatalogCart  -Id null `
 -Name null `
 -Items null `
 -Stats null
```

- Convert the resource to JSON
```powershell
$CatalogCart | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

