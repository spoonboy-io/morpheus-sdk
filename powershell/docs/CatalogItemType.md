# CatalogItemType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | **String** | Catalog Item Type category | [optional] 
**Description** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**Type** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Featured** | **Boolean** |  | [optional] 
**AllowQuantity** | **Boolean** | Can users order more than one of this item at a time. | [optional] 
**IconPath** | **String** |  | [optional] 
**ImagePath** | **String** |  | [optional] 
**DarkImagePath** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**LayoutCode** | **String** |  | [optional] 
**Blueprint** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**AppSpec** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Workflow** | [**CheckGroupInstance**](CheckGroupInstance.md) |  | [optional] 
**Content** | **String** |  | [optional] 
**OptionTypes** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**CreatedBy** | **String** |  | [optional] 
**Owner** | [**ApplianceSettingsEnabledZoneTypesInner**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItemType = Initialize-PSOpenAPIToolsCatalogItemType  -Id null `
 -Name null `
 -Code null `
 -Category null `
 -Description null `
 -Labels null `
 -Type null `
 -Enabled null `
 -Featured null `
 -AllowQuantity null `
 -IconPath null `
 -ImagePath null `
 -DarkImagePath null `
 -Visibility null `
 -LayoutCode null `
 -Blueprint null `
 -AppSpec null `
 -Config null `
 -Workflow null `
 -Content null `
 -OptionTypes null `
 -CreatedBy null `
 -Owner null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$CatalogItemType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

