# MorpheusApi.CatalogOrderCreateItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**CatalogOrderCreateType**](CatalogOrderCreateType.md) |  | [optional] 
**config** | **Object** | Config Object, required options depend on the catalog item type&#39;s associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type.  | 
**context** | **String** | Context Type for running the workflow, determines if a target resource must be selected. &#x60;instance&#x60;, &#x60;server&#x60;, or &#x60;appliance&#x60;. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type &#x60;workflow&#x60;.  | [optional] 
**target** | **Number** | Resource (Instance or Server) ID for context when running the &#x60;workflow&#x60;. Only applies to type &#x60;workflow&#x60; and only required when context is &#x60;instance&#x60; or &#x60;server&#x60;.  | [optional] 


