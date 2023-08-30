# CatalogOrderCreateItemsInner

Array of Item objects, each item in the order can be configured individually.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Config Object, required options depend on the catalog item type&#39;s associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type.  | 
**type** | [**CatalogOrderCreateItemsInnerType**](CatalogOrderCreateItemsInnerType.md) |  | [optional] 
**context** | **str** | Context Type for running the workflow, determines if a target resource must be selected. &#x60;instance&#x60;, &#x60;server&#x60;, or &#x60;appliance&#x60;. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type &#x60;workflow&#x60;.  | [optional] 
**target** | **int** | Resource (Instance or Server) ID for context when running the &#x60;workflow&#x60;. Only applies to type &#x60;workflow&#x60; and only required when context is &#x60;instance&#x60; or &#x60;server&#x60;.  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


