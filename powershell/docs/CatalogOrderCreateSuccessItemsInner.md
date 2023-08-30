# CatalogOrderCreateSuccessItemsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Type** | [**AppBlueprint**](AppBlueprint.md) |  | [optional] 
**Quantity** | **Int64** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Unit** | **String** |  | [optional] 
**Valid** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogOrderCreateSuccessItemsInner = Initialize-PSOpenAPIToolsCatalogOrderCreateSuccessItemsInner  -Id null `
 -Type null `
 -Quantity null `
 -Price null `
 -Currency null `
 -Unit null `
 -Valid null `
 -Status null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$CatalogOrderCreateSuccessItemsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

