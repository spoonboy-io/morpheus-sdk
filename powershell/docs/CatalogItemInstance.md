# CatalogItemInstance
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**Locations** | **String[]** |  | [optional] 
**VirtualMachines** | **Int64** |  | [optional] 
**Version** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItemInstance = Initialize-PSOpenAPIToolsCatalogItemInstance  -Id null `
 -Name null `
 -Status null `
 -Locations null `
 -VirtualMachines null `
 -Version null
```

- Convert the resource to JSON
```powershell
$CatalogItemInstance | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

