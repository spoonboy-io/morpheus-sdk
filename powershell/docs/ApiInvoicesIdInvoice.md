# ApiInvoicesIdInvoice
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) | This adds or updates the specified Metadata tags and removes any tags not specified. Array of objects having a name and value.  | [optional] 
**AddTags** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) | Add or update value of Metadata tags. Array of objects having a name and value.  | [optional] 
**RemoveTags** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) | This removes the specified Metadata tags matching name and optionally value (if provided). Array of objects having a name and value.  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiInvoicesIdInvoice = Initialize-PSOpenAPIToolsApiInvoicesIdInvoice  -Tags [{&quot;name&quot;:&quot;hello&quot;,&quot;value&quot;:&quot;world&quot;},{&quot;name&quot;:&quot;foo&quot;,&quot;value&quot;:&quot;bar&quot;}] `
 -AddTags [{&quot;name&quot;:&quot;hello&quot;,&quot;value&quot;:&quot;world&quot;},{&quot;name&quot;:&quot;foo&quot;,&quot;value&quot;:&quot;bar&quot;}] `
 -RemoveTags [{&quot;name&quot;:&quot;hello&quot;,&quot;value&quot;:&quot;world&quot;},{&quot;name&quot;:&quot;foo&quot;,&quot;value&quot;:&quot;bar&quot;}]
```

- Convert the resource to JSON
```powershell
$ApiInvoicesIdInvoice | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

