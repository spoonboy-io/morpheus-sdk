# CatalogItem
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | [**AppBlueprint**](AppBlueprint.md) |  | [optional] 
**Quantity** | **Int64** |  | [optional] 
**Status** | **String** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**Instance** | [**CatalogItemInstance**](CatalogItemInstance.md) |  | [optional] 
**OrderDate** | **System.DateTime** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItem = Initialize-PSOpenAPIToolsCatalogItem  -Id null `
 -Name null `
 -Type null `
 -Quantity null `
 -Status null `
 -StatusMessage null `
 -RefType null `
 -Instance null `
 -OrderDate null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$CatalogItem | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

