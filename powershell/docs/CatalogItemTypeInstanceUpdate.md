# CatalogItemTypeInstanceUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Catalog Item Type name | [optional] 
**Code** | **String** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | **String** | Catalog Item Type category | [optional] 
**Description** | **String** | Catalog Item Type description | [optional] 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Type** | **String** | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. | [optional] 
**Visibility** | **String** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**LayoutCode** | **String** | Identifier primarily used for Plugin Catalog Item Types | [optional] 
**IconPath** | **String** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**Enabled** | **Boolean** | Can be used to enable / disable the catalog item type. | [optional] [default to $true]
**Featured** | **Boolean** | Can be used to feature the catalog item type. | [optional] [default to $false]
**AllowQuantity** | **Boolean** | Can users order more than one of this item at a time. | [optional] [default to $false]
**Config** | [**CatalogItemTypeInstanceScribe**](CatalogItemTypeInstanceScribe.md) |  | [optional] 
**OptionTypes** | **Int64[]** | Array of option type IDs. Only applies to type instance and blueprint. | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItemTypeInstanceUpdate = Initialize-PSOpenAPIToolsCatalogItemTypeInstanceUpdate  -Name null `
 -Code null `
 -Category null `
 -Description null `
 -Labels null `
 -Type null `
 -Visibility null `
 -LayoutCode null `
 -IconPath null `
 -Enabled null `
 -Featured null `
 -AllowQuantity null `
 -Config null `
 -OptionTypes null
```

- Convert the resource to JSON
```powershell
$CatalogItemTypeInstanceUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

