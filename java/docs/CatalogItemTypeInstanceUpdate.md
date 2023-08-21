

# CatalogItemTypeInstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Catalog Item Type name |  [optional]
**code** | **String** | Useful shortcode for provisioning naming schemes and export reference. |  [optional]
**category** | **String** | Catalog Item Type category |  [optional]
**description** | **String** | Catalog Item Type description |  [optional]
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. |  [optional]
**visibility** | **String** | Visibility - Set to public to allow all tenants |  [optional]
**layoutCode** | **String** | Identifier primarily used for Plugin Catalog Item Types |  [optional]
**iconPath** | **String** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. |  [optional]
**enabled** | **Boolean** | Can be used to enable / disable the catalog item type. |  [optional]
**featured** | **Boolean** | Can be used to feature the catalog item type. |  [optional]
**allowQuantity** | **Boolean** | Can users order more than one of this item at a time. |  [optional]
**config** | [**CatalogItemTypeInstanceScribe**](CatalogItemTypeInstanceScribe.md) |  |  [optional]
**optionTypes** | **List&lt;Long&gt;** | Array of option type IDs. Only applies to type instance and blueprint. |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
INSTANCE | &quot;instance&quot;



