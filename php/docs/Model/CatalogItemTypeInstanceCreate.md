# # CatalogItemTypeInstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Catalog Item Type name | [optional]
**code** | **string** | Useful shortcode for provisioning naming schemes and export reference. | [optional]
**category** | **string** | Catalog Item Type category | [optional]
**description** | **string** | Catalog Item Type description | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**type** | **string** | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. | [optional]
**visibility** | **string** | Visibility - Set to public to allow all tenants | [optional] [default to 'private']
**layout_code** | **string** | Identifier primarily used for Plugin Catalog Item Types | [optional]
**icon_path** | **string** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional]
**enabled** | **bool** | Can be used to enable / disable the catalog item type. | [optional] [default to true]
**featured** | **bool** | Can be used to feature the catalog item type. | [optional] [default to false]
**allow_quantity** | **bool** | Can users order more than one of this item at a time. | [optional] [default to false]
**config** | [**\OpenAPI\Client\Model\CatalogItemTypeInstanceScribe**](CatalogItemTypeInstanceScribe.md) |  |
**option_types** | **int[]** | Array of option type IDs. Only applies to type instance and blueprint. | [optional]
**content** | **string** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
