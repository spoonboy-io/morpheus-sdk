# AddCatalogItemTypeRequestCatalogItemType


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Catalog Item Type name | [optional] 
**code** | **str, none_type** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**category** | **str, none_type** | Catalog Item Type category | [optional] 
**description** | **str** | Catalog Item Type description | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**type** | **str** | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. | [optional]  if omitted the server will use the default value of "workflow"
**visibility** | **str** | Visibility - Set to public to allow all tenants | [optional]  if omitted the server will use the default value of "private"
**layout_code** | **str, none_type** | Identifier primarily used for Plugin Catalog Item Types | [optional] 
**icon_path** | **str** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**enabled** | **bool** | Can be used to enable / disable the catalog item type. | [optional]  if omitted the server will use the default value of True
**featured** | **bool** | Can be used to feature the catalog item type. | [optional]  if omitted the server will use the default value of False
**allow_quantity** | **bool** | Can users order more than one of this item at a time. | [optional]  if omitted the server will use the default value of False
**option_types** | **[int]** | Array of option type IDs, see Inputs. Only applies to type instance and blueprint. | [optional] 
**content** | **str** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional] 
**app_spec** | **str** | The appSpec for blueprint type catalog items is a string in the Scribe YAML format with fields | [optional] 
**context** | **str** | Context for running the workflow, determines if a target resource must be selected. | [optional] 
**workflow_config** | **str** | Configuration object that contains settings for the workflow. | [optional] 
**config** | [**CatalogItemTypeInstanceScribe**](CatalogItemTypeInstanceScribe.md) |  | [optional] 
**blueprint** | [**CatalogItemTypeBlueprintCreateBlueprint**](CatalogItemTypeBlueprintCreateBlueprint.md) |  | [optional] 
**workflow** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


