# CatalogType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Context** | **String** |  | [optional] 
**Featured** | **Boolean** |  | [optional] 
**AllowQuantity** | **Boolean** | Can users order more than one of this item at a time. | [optional] 
**ImagePath** | **String** |  | [optional] 
**DarkImagePath** | **String** |  | [optional] 
**OptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogType = Initialize-PSOpenAPIToolsCatalogType  -Id null `
 -Name null `
 -Description null `
 -Type null `
 -Context null `
 -Featured null `
 -AllowQuantity null `
 -ImagePath null `
 -DarkImagePath null `
 -OptionTypes null
```

- Convert the resource to JSON
```powershell
$CatalogType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

