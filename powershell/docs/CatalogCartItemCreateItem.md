# CatalogCartItemCreateItem
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogCartItemCreateItemType**](CatalogCartItemCreateItemType.md) |  | [optional] 
**Quantity** | **Int64** | Quantity for this catalog item. Will be overridden to 1 if quantity not allowed by the item type.  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Config Object, required options depend on the catalog item type&#39;s associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type.  | [optional] 
**Context** | **String** | Context Type for running the workflow, determines if a target resource must be selected. &#x60;instance&#x60;, &#x60;server&#x60;, or &#x60;appliance&#x60;. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type &#x60;workflow&#x60;.  | [optional] 
**Target** | **Int64** | Resource (Instance or Server) ID for context when running the &#x60;workflow&#x60;. Only applies to type &#x60;workflow&#x60; and only required when context is &#x60;instance&#x60; or &#x60;server&#x60;.  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogCartItemCreateItem = Initialize-PSOpenAPIToolsCatalogCartItemCreateItem  -Type null `
 -Quantity null `
 -Config null `
 -Context null `
 -Target null
```

- Convert the resource to JSON
```powershell
$CatalogCartItemCreateItem | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

